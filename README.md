
# The Solution

First, I reviewed the task and tried to define the business logic. I believe that to implement the `exploreService`, I also need an underlying domain service, which I named the `decision` module.

In my implementation, I follow a lightweight hexagonal-based architecture, where there are various domain modules in the system, and we can utilize different actors for communication.

The `decision` domain module is essentially a repository layer that stores the defined entity in the database and performs other operations.

The `explore` module is responsible for implementing the service definition specified in the proto file (service), as well as matching the generated gRPC code (controller). The `explore` module does not contain state and is dependent on the `decision` domain module.

I bootstrapped the project using a starter kit maintained by me and followed the structural recommendations of the starter kit. I wrote more about this here: [Github Repo](https://github.com/kazmerdome/best-ever-golang-starter).

Each domain module has a definition file (`decision.go`, `explore.go`), a `module.go` file responsible for exporting providers, and provider files (`controller.go`, `service.go`, `repository.go`) that are responsible for interactions with various actors (db, gRPC).

## Notes
The gRPC definition included a pagination token. I implemented this token using the `created_at` timestamp. This returns the next items in the list in descending order of creation time (`(created_at < sqlc.arg('pagination_token')::date OR sqlc.narg('pagination_token') IS NULL)`). Since the definition did not include a `Limit`, I set a default value of `10` in the repository layer.




## Test

I have prepared unit tests for each domain module. The unit tests are mocked, for which I used Mockery. I aimed to test the units within each module in isolation, so all dependencies were mocked.

## Selected Technologies
- Database: PostgreSQL with SQLC
- Grpc: Protoc
- Configuration Management: Viper
- Logging: Zerolog




# Run the Application in Docker Container

## Option1: Using Docker Compose (recommended)

`docker compose up`

or

`make up`


use ```localhost:4445```


## Option2: Using Docker

### Step 1.
`docker build -t kazmerdome-muzz .`

### Step 2.
`docker run --rm --network="host" -p 4445:4444 --env=POSTGRES_URI=postgresql://muzz:muzz@localhost:5432 --env=POSTGRES_DATABASE=muzz --env=POSTGRES_IS_SSL_DISABLED=true kazmerdome-muzz`

use ```localhost:4445```




# Run the Application Locally

### Step 1.
Run docker compose file for database and its migration

Option 1:
```bash
  `make up`
```

Option 2:
```bash
  `docker compose up`
```

### Step 2.
```bash
  `cp .env.example .env`
```

### Step 3.

Option 1 (Recommended):

```
  start applications using Visual Studio Code Run and Debug
```

Option 1:

```bash
  make run-gateway
```

Option 3:

```bash
  go run cmd/gateway/main.go
```

### Step 4.

use ```localhost:4444```
