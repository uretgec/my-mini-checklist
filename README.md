# My Mini Checklist
How to make simple key value based checklist api with golang stdlib for not master users?

## Feature
- InMemory database (with bg save option).
- Set, Get, GetAll, Del, Stats, Flush methods allowed.
- POST and GET methods are equal. Use the method which you prefer most. (whatever you want)
- Logged all requests
- Postman Collections, Docker, Local and DO Server build option is ready to use
- No UI only restful api (Sorry)
- Thats its all :)

## Run Quickly

Without build

```
go run . --addr ":3000" --dbpath "./store.db" --bgsave 1m0s --logpath "./service-http.log"
```

Open your browser and go to localhost:3000
```
Welcome Home
```

Ready to use

## Installation

Select below one of them for easy installation.

1. Docker Build
First Have to download and install Docker Desktop App

If you want build image only
```
docker build -t my-mini-checklist .
```

After run the command
```
docker-compose up --build
```

Finish the install and only one container is up
Open your browser and go to localhost:3000
```
Welcome Home
```

Ready to use

2. Local Build

Run sh command for generate executable "my-mini-checklist" file to build folder
```
./build.sh my-mini-checklist
```

Run the server with options
```
cd build
chmod +x my-mini-checklist
./my-mini-checklist --addr ":3000" --dbpath "./store.db" --bgsave 1m0s --logpath "./service-http.log"
```

Open your browser and go to localhost:3000
```
Welcome Home
```

Ready to use

3. Github Actions Deploy DO Server (ASAP)

> NOT READY

## Tests

1. Postman collection file for all routes is ready to import. 

    1. Pls go to docs folder and import "My-Mini-Checklist.postman_collection.json" file to Postman App -> Collection Tab in your computer
    2. Go to Postman App -> Enviroments tab. Import all environments "My-Mini-Checklist Dev Env.postman_environment.json", "My-Mini-Checklist Prod Env.postman_environment.json", "My Workspace.postman_globals.json" 
    3. Run the "my-mini-checklist" server which deploy process you choose
    4. Select "My-Mini-Checklist" collection folder in Postman App
    5. Select "Prod or Dev Environment" (right up side the header)
    6. Click the "Run" button (right up side the header)
    7. Good lock

> NOTE: If not selected environment, Postman collections always request localhost:3000

2. Run Go test command


```
go test 

or more detail option:

go test -v

or 

go test -v --race
```

## Licence
Distributed under the MIT License. See LICENSE for more details.


# Swagger UI

```
./swagger generate spec -o /path/docs/swagger.json -w /path
```

```
./swagger serve -F swagger /path/docs/swagger.json 
```