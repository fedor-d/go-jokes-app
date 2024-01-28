go-jokes-app

# A simple Golang app
[![Chuck Norris](resources/images/image.jpg)](https://en.wikipedia.org/wiki/Chuck_Norris)

## Running app via Docker
```shell
docker pull fedorddev/go-jokes-app:v0.0.1
docker run -p 8080:8080 fedorddev/go-jokes-app:v0.0.1
```

## App requests
| Request           | Example                               |
|-------------------|---------------------------------------|
| GET Jokes         | http://localhost:8080/v1/jokes        |
| GET Joke by id    | http://localhost:8080/v1/jokes/$id    |
| GET random Joke   | http://localhost:8080/v1/jokes/random |
| POST Joke         | http://localhost:8080/v1/jokes        |
| GET Health Live   | http://localhost:8080/health/live     |
| POST Health Ready | http://localhost:8080/health/ready    |

## Request example 
```shell
curl --location --request POST 'http://localhost:8080/v1/jokes' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": "1000",
"jokeText": "Joke 1000"
}'
```

## Postman collection
Refer to the project _auxiliary_ folder