run-server:
	@docker-compose --env-file ./server/.env up --build server

# run-parser:
# 	@docker-compose --env-file ./server/.env up --build server