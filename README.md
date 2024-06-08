# REST API in Go

*Disclaimer*: This is a practice project that follows the tutorial [How To Build a Complete API In Golang (Docker, JWT, MySQL)](https://www.youtube.com/watch?v=2JNUmzuBNV0&list=WL&index=10&t=432s&ab_channel=Tiago)

# Table of Contents
* [STEP 0: Initialize Project](#step-0-initialize-project)
* [STEP 1: API Server](#step-1-api-server)
* [STEP 2: Database](#step-2-database)
* [STEP 3: Environment Variables Sourcing](#step-3-environment-variables-sourcing)

## Third-Party Dependencies
```bash
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```



## STEP 0: Initialize Project
1. Choose a directory on your local machine (e.g. `github`). Create a new directory (e.g. `rest-api-go-practice`). Go into this newly created directory.
```
cd /path/to/github/
mkdir rest-api-go-practice
cd rest-api-go-practice
```
2. Initialize a Go module within this directory.
```
go mod init github.com/username/rest-api-go-practice // go mod init <module-path>
```
NOTE: `github.com/username/` is a part of the module path for Go to manage. It is NOT creating new folders (`github.com`, `username`) on your local machine. Rather, `github.com/username/rest-api-go-practice` is used within the Go ecosystem to help Go uniquely identify this project (aka module). Because domain names (e.g. `github.com`) are unique on the Internet, using this convention ensures global uniqueness of this module path. If you are linking the project to a remote repository hosted on GitHub, it is the common practice to use the full GitHub path `github.com/<username>/<project-name>` as Go module path to be the unique identifier. Two advantages of doing so: 1). using domain name ensures global uniqueness. 2). helps other developers easily find the source repository on GitHub. 

3. Create a new repository on Github. The repository name should be the same as `rest-api-go-practice`. Link the local project to the remote repository by following the instructions provided by GitHub.
    - add and edit `.gitignore`
    - add `README.md`

## STEP 1: API Server
```
project/
├── main.go
└── api.go
    └── APIServer <struct>
        ├── NewAPIServer() <constructor>
        └── Serve()        <method>
```
## STEP 2: Database
```
project/
├── main.go
├── api.go
├── store.go
│   ├── Store <interface>
│   │    └── CreateUser()
│   └── Storage <struct>
│        ├── NewStore()    <constructor>
│        └── CreateUser()  <method>
└── db.go
    └── MySQLStorage <struct>
        ├── NewMySQLStorage() <constructor>
        └── Init()            <method>
```
## STEP 3: Environment Variables Sourcing
```
project/
├── main.go
├── api.go
├── store.go
├── db.go
└── config.go
    ├── Config        <struct>
    ├── Envs          <Config-type global variable>
    ├── initConfig()  <constructor>
    └── getEnv()
```
- **Pause for a second!** Let's be clear about the workflow of the main program!
    1. First, it creates a `mysql.Config` struct, which holds the configuration for connecting to a MySQL database. The configuration values are taken from the `Envs` variable, which holds environment variables. The `Net` field is set to `"tcp"`, indicating that the connection to the MySQL server should be made over TCP. `AllowNativePasswords` is set to `true`, which allows the use of the native MySQL password authentication method. `ParseTime` is set to `true`, which tells the driver to parse `DATE` and `DATETIME` columns into `time.Time` values.
    2. Next, it calls `NewMySQLStorage(cfg)`, which creates a new `MySQLStorage` object. This object is responsible for interacting with the MySQL database. It uses the configuration provided to establish a connection to the database.
    3. The `Init` method of the `MySQLStorage` object is then called to initialize the database. If there's an error during this process, the application will log the error and exit.
    4. After the database has been initialized, a new `Store` is created with `NewStore(db)`. This `Store` object is responsible for performing operations on the database.
    5. Finally, a new `APIServer` is created with `NewAPIServer(":3000", store)`. This server is configured to listen on port 3000 and uses the `Store` for its database operations. The `Serve` method is then called to start the server and begin listening for incoming HTTP requests.
```go
func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// MySQL storage configuration
	sqlStorage := NewMySQLStorage(cfg)

	// Initialize the database
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the store
	store := NewStore(db)

	// Inject the store into the API server
	api := NewAPIServer(":3000", store)
	api.Serve()
}
```