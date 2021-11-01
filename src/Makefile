NAME = goweb
TAG= 0.0.6
DBINSTANCE = userdb

default: docker

docker:
	docker build -t $(NAME):$(TAG) .

run:
	docker run  --link userdb -p 10707:8080 -e MYSQL_HOST="$(DBINSTANCE)" -d $(NAME):$(TAG)