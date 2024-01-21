package main

import (
	"net/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"software_containerization/Models"
	"software_containerization/DTOs"
	"software_containerization/Handlers"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"os"
	"fmt"
)

func InitializeRoles(db *gorm.DB) {
	roles := []string{"Organizer", "Participant", "Guest"}
	for _, roleName := range roles {
		var role Models.Role
		result := db.First(&role, "name = ?", roleName)
		if result.Error == gorm.ErrRecordNotFound {
			// Role not found, create it
			newRole := Models.Role{Name: roleName}
			db.Create(&newRole)
			log.Printf("Role '%s' created.\n", roleName)
		}
	}
}

// ValidateUser checks the provided credentials against the database.
// If the credentials are valid, it returns the user object. Otherwise, it returns an error.
func ValidateUser(db *gorm.DB, credentials DTOs.UserCredentials) (*Models.User, error) {
	var user Models.User

	// Retrieve the user from the database by the provided email.
	if err := db.Preload("Role").Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		// Handle user not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid login credentials")
		}
		// Handle other potential errors
		return nil, err
	}

	// Compare the provided password with the stored hashed password.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		// The password does not match.
		return nil, errors.New("invalid login credentials")
	}

	// The credentials are valid.
	return &user, nil
}

func LoginHandler(db *gorm.DB, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials DTOs.UserCredentials // Define UserCredentials struct
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate user credentials
		user, err := ValidateUser(db, credentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Create JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"role":  user.Role.Name,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		
		// Extract and validate the JWT token
		// Assume token parsing returns userID and role
		userID, role, err := parseToken(tokenString, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Set the user information in Gin context
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

func parseToken(tokenString string, secretKey string) (userID uint, role string, err error) {
	// Check if the tokenString is empty
	if tokenString == "" {
		return 0, "", errors.New("no token provided")
	}

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token's signing method corresponds to the signing method you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user ID and role from the token's claims
		userIDfloat, ok := claims["id"].(float64)
		if !ok {
			return 0, "", errors.New("Invalid user id")
		}

		userID := uint(userIDfloat)

		role, ok := claims["role"].(string)
		if !ok {
			return 0, "", errors.New("invalid role")
		}

		return userID, role, nil
	}

	return 0, "", errors.New("invalid token")
}

func IsOrganizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract role from Gin context (set in AuthMiddleware)
		role, exists := c.Get("role")
		if !exists || role != "Organizer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func main() {
	// Retrieve JWT secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT secret key must be set")
	}

	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting sever")
	db := database_init()

	r := gin.Default()

	// Initialize roles
	InitializeRoles(db)

	r.GET("/roles", Handlers.GetRolesHandler(db))
	r.GET("/roles/:id", Handlers.GetRoleByIdHandler(db))

	// r.POST("/user", Handlers.CreateUserHandler(db))
	r.POST("/auth/register", Handlers.CreateUserHandler(db))
	r.POST("/auth/login", LoginHandler(db, secretKey)) // /auth/login
	// auth/logout
	r.GET("/users", Handlers.GetUsersHandler(db)) // get all users, get user by id, or by email query params
	r.PUT("/users/:id", Handlers.UpdateUserHandler(db))

	r.POST("/events", AuthMiddleware(secretKey), IsOrganizer(), Handlers.CreateEventHandler(db)) // Only organizers can create events
	r.GET("/events", Handlers.GetEventsHandler(db)) // get all events, get event by id, or by userId query params
	r.DELETE("/events/:id", Handlers.DeleteEventHandler(db))
	r.PUT("/events/:id", AuthMiddleware(secretKey), IsOrganizer(), Handlers.UpdateEventHandler(db))

	r.Run(":8080")
}

func database_init() *gorm.DB {
	log.Println("connecting to database")
	log.Println("waiting for database to start")
	time.Sleep(time.Duration(5) * time.Second)

	dsn := "host=database user=postgres password=postgres dbname=Events port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connection to database established")
	log.Println("migrating database")
	err = db.AutoMigrate(&Models.Event{}, &Models.User{}, &Models.Role{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("database migration finished")
	//return gorm.DB{}
	return db
}