IP = 39.107.35.228


default: run

run:
	docker run  -p 80:80 --add-host=myserver:$(IP)  --name nginx-proxy  -v $(PWD)/www  -v $(PWD)/conf/conf.d:/etc/nginx/conf.d  -v $(PWD)/logs:/wwwlogs  -v /etc/localtime:/etc/localtime  --restart=always  -d nginx