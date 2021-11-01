NAME = userdb
TAG= 0.0.1
INSTANCE = userdb

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run --name $(INSTANCE) -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d $(NAME):$(TAG)