include .env
export $(shell sed 's/=.*//' .env)

deploy: export TAG=$(shell git pull -q && git describe --abbrev=0 --tags)
deploy:
	git pull
	docker-compose pull
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
	docker ps

api:
	cd cmd/api && CONFIG_FILE=config-dev.json go run .

indexer:
	cd cmd/indexer && CONFIG_FILE=config-dev.json go run .

metrics:
	cd cmd/metrics && CONFIG_FILE=config-dev.json go run .

clearmq:
	docker exec -it bcd-mq rabbitmqctl stop_app
	docker exec -it bcd-mq rabbitmqctl reset
	docker exec -it bcd-mq rabbitmqctl start_app

aliases:
	cd scripts/aliases && go run .

migration:
	cd scripts/migration && go run .

upd:
	docker-compose -f docker-compose.yml docker-compose.prod.yml up -d --build

es-aws:
	cd scripts/es-aws && go build .

s3-creds: es-aws
	docker exec -it bcd-elastic bash -c 'bin/elasticsearch-keystore add --stdin s3.client.default.access_key <<< "$$AWS_ACCESS_KEY_ID"'
	docker exec -it bcd-elastic bash -c 'bin/elasticsearch-keystore add --stdin s3.client.default.secret_key <<< "$$AWS_SECRET_ACCESS_KEY"'
	./scripts/es-aws/es-aws -a reload_secure_settings

s3-repo: es-aws
	./scripts/es-aws/es-aws -a create_repository

s3-restore: es-aws
	./scripts/es-aws/es-aws -a delete_indices
	./scripts/es-aws/es-aws -a restore

s3-snapshot: es-aws
	./scripts/es-aws/es-aws -a snapshot

s3-policy: es-aws
	./scripts/es-aws/es-aws -a set_policy

latest:
	git tag latest -f && git push origin latest -f