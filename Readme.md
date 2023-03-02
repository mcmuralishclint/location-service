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

# Configs
## config.yml
```
google:
  maps_api_key: 'API_KEY_GOES_HERE'
test:
  maps_api_key: ''
baidu:
  maps_api_key: 'API_KEY_GOES_HERE'
address_provider: "google"
port: 3000
```

# How to start the project
1. Place the config.yml file in the repo root
2. Start the redis server locally on port 6379
```
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
```
3. Start the application
```
go run main.go
```