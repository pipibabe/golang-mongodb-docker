# golang-mongodb-docker

#  golang docker development environment

	cd ~/projects/golang-mongodb-docker/golang_test \
	docker run --rm -it --name go-test -v .:/usr/local/go/src/golang_test:rw --workdir /usr/local/go/src/golang_test/cmd golang_test golang