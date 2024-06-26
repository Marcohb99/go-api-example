## 2.2.0 (2024-04-12)

### Feat

- **auth**: add api keys to docker compose env
- **auth**: add auth middleware that loads in memory api keys and checks if the user has given a valid api key

## 2.1.0 (2024-03-29)

### Feat

- ALL working!
- **repository**: implement GetAll method
- **all**: implement and test the service
- **all**: add endpoint, handler and command

## 2.0.0 (2024-03-22)

### BREAKING CHANGE

- removed a method

### Feat

- remove all method from repo and fix tests

## 1.0.0 (2024-03-22)

### BREAKING CHANGE

- first major

### Feat

- all working

## 0.7.0 (2024-03-15)

### Feat

- **event-handler-and-subscriber**: implemented the in memory event bus and added a handler to increase the total releases when a release is created
- **events**: add domain events, record events on entity creation and publish in event bus after creation
- **middleware**: add logging and recovery middlewares and test them

## 0.6.0 (2024-02-09)

### Feat

- **repo**: add db timeout and execute with timeout context

## 0.5.0 (2024-02-05)

### Feat

- **server**: add support for graceful shutdown

## 0.4.1 (2024-01-29)

### Refactor

- **command-bus**: add logic to refactor process using command pattern
- **command**: add command domain

## 0.4.0 (2024-01-29)

### Feat

- **useCase**: add application service and refactor logic
- **repository**: add all function to repository and test it

### Fix

- update repository mocj

## 0.3.0 (2024-01-03)

### Feat

- **value-objects**: use value objects to validate domain rules
- **WIP**: add code to insert release in database

### Refactor

- **add-release**: inject repository

## 0.2.1 (2023-12-29)

### Refactor

- **server**: use better architecture

## 0.2.0 (2023-12-29)

### Feat

- **gin**: use gin and define /hello endpoint

## 0.1.0 (2023-12-29)

### Feat

- **main.go**: hc endpoint
