build:
	sh start.sh
	docker-compose up

rebuild:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down

log:
	docker logs -f mongo_app

sudo:
	echo "deucandau123" | sudo -S chmod -R 777 . ; \
	
