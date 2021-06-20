# go-kafka

Go project for building / consuming Kafka

## local deployment

run `docker-compose up -d` for spinning up Kafka and Zookeeper

run `go run main.go`

## API usage

### Endpoints

#### POST /v1/comments

request body:

```
{
 "text"   : "text-text"
}
```

command:

```
curl --header "Content-Type: application/json" --request POST --data '{"text":"test"}' http://localhost:8080/v1/comments
```