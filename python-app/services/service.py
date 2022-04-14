
from time import sleep

from kafka import KafkaProducer, KafkaConsumer

bootstrap_servers = ['localhost:9091', 'localhost:9092']
topicName = 'my-topic-three'


def run_producer():
    producer = KafkaProducer(bootstrap_servers=bootstrap_servers)
    producer = KafkaProducer()

    for e in range(1000):
        # data = f'number {e}'
        try:
            producer.send(topicName, b'Algo Aew!')
            producer.flush()
        except Exception as e:
            print(e)
            raise
        sleep(5)
        print(f"send data number: {e+1}")


def run_consumer():
    try:
        consumer = KafkaConsumer(topicName, bootstrap_servers=bootstrap_servers)
        print('foi')
    except Exception as e:
        print(e)
        raise
    print(consumer)
    for message in consumer:
        print(message)