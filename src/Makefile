NAME = goweb
TAG= 0.0.6
DBINSTANCE = userdb

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run  --link userdb -p 10700:10700 -e MYSQL_HOST="$(DBINSTANCE)" -d $(NAME):$(TAG)