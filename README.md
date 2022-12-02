# CRM Backend

The following workspaces provides a sample go application that runs an HTTP server utilizing gorilla/mux. This sample application includes the following capabilites which allows a user to create, read, update, and delete resources specified in a mock database backend. 

## How to Use

1. Download the `main.go` and `main_test.go` from the associated workspace
2. Download the external 3rd Party Modules
    - github.com/google/uuid - Used to generate a unique id
    - github.com/gorilla/mux - Used as the http router

    ```bash
    go get github.com/google/uuid
    go get github.com/gorilla/mux
    ```
3. Run the application

    ```bash
    go run main.go
    ```
4. Test the application

    ```bash
    go test
    ```
5. Build the application

    ```bash
    go build
    ```
6. Run the application

    ```bash
    ./crm
    ```