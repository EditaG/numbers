Installation and setup
------------
Requirements: working docker and docker-compose installation. Tested with docker-compose v1.26.0 and docker v20.10.11

- Start project with:
```bash
docker-compose up
```
Session manager API is available at "localhost:8081" and session API is available at: "localhost:8080"

- After starting the project, run tests with:
```bash
docker-compose exec session_manager go test ./...
docker-compose exec converter go test ./...
```


Some example requests
------------

More request examples can be found in test cases

```bash
# login
curl --location --request POST 'localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "demo_user",
    "password" : "Pa33m0rD*&!"
}'

# convert roman numeral
curl --location --request POST 'localhost:8080/api/convert/roman' \
--header 'Authorization: {token}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "input": "I"
}'

# logout
curl --location --request POST 'localhost:8080/api/logout' \
--header 'Authorization: {token}' \
--header 'Content-Type: application/json' \
--data-raw ''
```


TODO
------------

- Consider to store session in Cookie header, using Custom authorization header has several security issues

- Hide internal errors from http response data, might have sensitive information

- Custom validation messages

- Generate API docs (f.e. https://github.com/swaggo/gin-swagger)

- Cover all project with tests

- Add ability to pass and save additional session data to session_manager

- Stop http server process gracefully

- Use encrypted session token id

- Add code style fixer

- Add semantic commit linter, git commit hooks, versioning, changelog generation from commit diff

- Install monitoring packages (sentry, newrelic etc.)


