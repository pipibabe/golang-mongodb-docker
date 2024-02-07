# golang-mongodb-docker


# 事前準備
    $ docker network create --driver bridge dockernetwork

# run golang development environment on docker

	$ cd ~/projects/golang/golang-mongodb-docker/golang_test \
	&& docker run --rm -it --name go-test --network dockernetwork \
	-v .:/usr/local/go/src/golang_test:rw --workdir /usr/local/go/src/golang_test golang

# create go module
	switch to your module folder
  	$ go mod init golang_test/moduleName

# execute go program
	go run .

# run mongo on docker  --network test-network

	$ docker run --rm -d --name mongodb -p 27017:27017 --network dockernetwork mongo
	$ docker exec -it mongodb mongosh

The MongoDB Server log is available through Docker's container log:

$ docker logs some-mongo


# 查詢 docker container ip 指令

	$ docker network inspect bridge
