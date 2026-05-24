COMPOSE=sudo docker compose
DOCKER_COMPOSE=infraestruture/docker-compose.yaml
.PHONY: up down build restart logs ps clean app-logs db-logs shell-app shell-db

up:
	$(COMPOSE) -f $(DOCKER_COMPOSE) up

build:
	$(COMPOSE) -f $(DOCKER_COMPOSE) up --build 

down:
	$(COMPOSE) -f $(DOCKER_COMPOSE) down

restart:
	$(COMPOSE) -f $(DOCKER_COMPOSE) restart

logs:
	$(COMPOSE) -f $(DOCKER_COMPOSE) logs -f

app-logs:
	$(COMPOSE) -f $(DOCKER_COMPOSE) logs -f app

db-logs:
	$(COMPOSE) -f $(DOCKER_COMPOSE) logs -f postgres

ps:
	$(COMPOSE) -f $(DOCKER_COMPOSE) ps

clean:
	$(COMPOSE) -f $(DOCKER_COMPOSE) down -v --remove-orphans

shell-app:
	$(COMPOSE) -f $(DOCKER_COMPOSE) exec app sh

shell-db:
	$(COMPOSE) -f $(DOCKER_COMPOSE) exec postgres psql -U admin -d appdb