# golang-mongodb-docker

# run golang development environment on docker

	$ cd ~/projects/golang-mongodb-docker/golang_test \
	docker run --rm -it --name go-test -v .:/usr/local/go/src/golang_test:rw --workdir /usr/local/go/src/golang_test/cmd golang

# create go module
	switch to your module folder
  	$ go mod init golang_test/moduleName

# execute go program
	go run .

# run mongo on docker

	$ docker run --rm -d --name mongodb -p 27017:27017 mongo
	$ docker exec -it mongodb mongosh

	$ docker run -it --rm --name mongodb -p 27017:27017 mongo mongosh

	$ docker exec -it some-mongo bash
The MongoDB Server log is available through Docker's container log:

$ docker logs some-mongo
