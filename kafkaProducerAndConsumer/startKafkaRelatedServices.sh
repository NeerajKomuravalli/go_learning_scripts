zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
kafka-server-start /usr/local/etc/kafka/server.properties
kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic message-log
kafka-console-producer --broker-list localhost:9092 --topic message-log
kafka-console-consumer --bootstrap-server localhost:9092 --topic message-log --from-beginning