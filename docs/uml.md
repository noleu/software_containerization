# Architecture

## Sequence Diagram
The following sequence diagram shows the interaction between the user, the frontend, the backend and the database.  
First the user requests the events from the frontend, which in turn requests the events from the backend.  
The backend fetches the events from the database and returns them to the frontend, which then returns them to the user.
The user can also create a new event, which is then sent to the frontend, which sends it to the backend,
which inserts it into the database and returns the updated list of events to the frontend,
which then returns it to the user.

```mermaid
sequenceDiagram
    actor user
    participant fe as Frontend
    participant be as Backend
    participant db as Database
    
    user ->> fe: Request
    fe ->> be: getEvents
    be ->> db: Fetch events
    db -->> be: Events
    be -->> fe: Events
    
    user ->> fe: createNewEvent
    fe ->> be: POST events
    be ->> db: Insert event
    db -->> be: Events
    be -->> fe: Events    
```

The frontend is a static webpage, which calls a backend API written in golang, with gin as the http framework.
The backend uses gorm to to store and retrieve events from a postgres database.

## Kubernetes component diagram
![Kubernetes component diagram](./Kubernetes-architecture.png "component diagram")