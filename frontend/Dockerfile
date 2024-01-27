# Stage 1: Development
FROM oven/bun:latest as dev

WORKDIR /app

COPY package.json /app
COPY bun.lockb /app

RUN bun install --silent

COPY . /app

CMD ["bun", "run", "dev"]

# Stage 2: Build for production
FROM oven/bun:latest as builder

WORKDIR /app

COPY package.json /app
COPY bun.lockb /app

RUN bun install --silent

COPY . /app

RUN bun run build

# Stage 3: Production environment with Nginx
FROM nginx:alpine

# Install wget
RUN apk --no-cache add wget

# Copy static assets from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
