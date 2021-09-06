# Launches a container from golang in interactive mode
docker run --rm -it --name go-restful golang

docker run --rm -it --name go-restful -v $PWD:/go/src/github.com/onyewuenyi/rest_api golang


docker run -it -p 5432:5432 -d postgres

docker run -p 8080:8080 rest_api:multistage

# Initialize the Go modules with your GitHub repository address:
go mod init github.com/onyewuenyi/zero2prod_golang

# Fetch go modules
go get -u github.com/gorilla/mux 
go get -u github.com/lib/pq


# Create project structure
touch app.go
touch main.go
touch main_test.go
touch model.go


go run main.go
go build
go test

# Docker 
# View local imgs 
docker images



docker build -t rest_api:multistage -f Dockerfile.multistage .

http://localhost:8080/
http://localhost:8080/health_check

curl --data 'name=le%20guin&email=ursula_le_guin%40gmail.com' http://127.0.0.1:8080/subscriptions --verbose


# Docker 
1. To build a new docker image or use an existing one
docker run --rm -it --name go-restful golang
--rm 		Automatically remove the container when it exits
--it create an interactive bash shell in the container
This launches a container from golang in interactive mode so that I can later execute everygo commands


instead of installing the db or using a os native package manager we are going ot use the docker imag of the db 

wire up app and db container 


# Docker DB setup 
# start a test db
$ docker run -it -p 5432:5432 -d postgres




IaC using Terraform 
docker compose: running multiple containers at the same time abstraction

want these containers to comm with eachother in the network 

single entry pt is nginx 
env file with conn str to db

build from an img to have it locally