# golang-user-microservice

Create a micro service for handling user accounts.
Written in golang.


# Getting started

Minimum requirements:

- Go 1.15 

Enable go modules and go dep integration.

## Running the server

    go run main.go


## Testing the server
    
In a browser, navigate to:

    http://localhost:80/api/ping
    
You should see:

    {"ok":true,"pong":true}
    
# Running with docker

Simply run in the background with a docker-enabled machine:

    docker-compose up -d
    
Again, the ping check with this URL should return a valid JSON response.

    http://localhost:80/api/ping
    
    
## Updating `swagger.yml`

The APIs here should conform to OpenAPI as much as possible.

1. In a web browser, go to:

    https://editor.swagger.io
    
2. Paste the contents of `swagger.yml` to the editor.
3. Make the changes on the editor.
4. Copy the contents on the editor and paste in `swagger.wml`.
5. Commit the new codes and submit a pull request.   


# Contributing

The goal of this microservice to handle only API calls for user accounts.

The first part involves CRUD operations. (completed)

The second part involves authentication and authorization.

The third part involves integration with external microservices and authenticate over JWT tokens.

Some areas of enhancements for this microservice may be:

- Using a MySQL/Sqlite3 database for storage. (completed)
- Enable CORS for cross-domain access over REST calls. (completed)
- Generating and authenting JWT tokens.
- Handling refresh tokens, and session timeouts and extension.
- Production deployment with Docker, with support for environment-level settings. (completed)
    
