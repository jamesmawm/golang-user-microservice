# golang-user-microservice

Create a micro service for handling user accounts.
Written in golang.


# Getting started

Minimum requirements:

- Go 1.15 

Enable go modules and go dep integration.

## Running the server

Run the following commands using go modules:

    go mod init v1
    go mod download
    go mod vendor
    
This creates a `go.mod` in the root directory.

Then run:

    go run main.go


## Testing the server
    
In a browser, navigate to:

    http://localhost:8080/api/ping
    
You should see:

    {"ok":true,"pong":true}
    
    
## Updating `swagger.yml`

1. In a web browser, go to:

    https://editor.swagger.io
    
2. Paste the contents of `swagger.yml` to the editor.
3. Make the changes on the editor.
4. Copy the contents on the editor and paste in `swagger.wml`.
5. Commit the new codes and submit a pull request.   


# Contributing

The goal of this microservice to handle only API calls for user accounts.
The first part involves CRUD operations.
The second part involves authentication and authorization.
The third part involves integration with external microservices and authenticate over JWT tokens.

Some aras of enhancements:

- Using a MySQL database for storage.
- Enable CORS for cross-domain access over REST calls.
- Generating and authenting JWT tokens.
- Handling refresh tokens and session extension.
    