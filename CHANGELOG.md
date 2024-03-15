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
