
.PHONY: kafka-up
kafka-up:
	docker-compose -f docker-compose.kafka.yml up -d
	make create-topic

.PHONY: create-topic
create-topic:
	docker-compose -f docker-compose.kafka.yml exec broker kafka-topics --create --topic messages --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 --if-not-exists


.PHONY: kafka-stop
kafka-stop:
	docker-compose -f docker-compose.kafka.yml stop

.PHONY: kafka-down
kafka-down:
	docker-compose -f docker-compose.kafka.yml down

.PHONY: broker-bash
broker-bash:
	docker exec -it broker bash


.PHONY: zookeeper-bash
zookeeper-bash:
	docker exec -it zookeeper bash

.PHONY: topic-list
topic-list:
	docker-compose -f docker-compose.kafka.yml exec broker kafka-topics --list --bootstrap-server localhost:9092