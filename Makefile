.PHONY: up
up: 
	docker-compose up --build 
	
.PHONY: down
down: 
	docker image prune --filter label=stage=build1 -f
	docker image prune --filter label=stage=build2 -f
	docker-compose down
