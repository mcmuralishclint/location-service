# Hex Arch FAQ

1. What is an adapter?
`This is the layer that communicates with external entities such as databases, RMQs, APIs, etc`
2. What is a domain?
`Contains the skeleton of the business logic of the application and encapsulate the core logic of the system`
3. What is a service?
`contains the application services which use the domain objects to perform the use cases. Actual implementation of the business logic skeleton resides here.`
4. In the domain layer, Why do we have interfaces in repository.go and service.go?
`To clear the responsibilites of the domain layer and the persistence layer. Repository contains the interfaces to communicate with the persistence layer. Service contains the interface to commnicate with the domain objects.`
5. Flow of logic
```
main -> server -> handler -> service -> repository -> adapter -> external
```