## commands - kafka in docker
 -  kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
 -  kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
 -  consumir topics : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
 -  consumir topics do inicio : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning
 -  consumir topics do inicio + grupo : kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning --group=x
 -  kafka-topics --describe --bootstrap-server=localhost:9092 --topic=teste
 -  criar msgs : kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

 ## installs - go
 -  go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka

 ## project
   - docker exec -it kafka_kafka1_1 kafka-topics --zookeeper zookeeper:2181 --create --topic my-topic --partitions 1 --replication-factor 1
   - docker exec -it kafka_kafka1_1 kafka-topics --zookeeper zookeeper:2181 --create --topic my-topic-three --partitions 1 --replication-factor 2
   - kafdrop : localhost:9000

