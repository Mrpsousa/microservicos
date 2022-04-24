
from time import sleep
from json import dumps
from kafka import KafkaProducer, KafkaConsumer
from classes import Quotes

bootstrap_servers = ['localhost:9091', 'localhost:9092']
topicName = 'my-topic-three'

quotes: Quotes = Quotes()


def run_producer():
    try:
        producer = KafkaProducer(bootstrap_servers=bootstrap_servers)
        producer = KafkaProducer()

        for e in range(1000):
            data = quotes.run()    
            producer.send(topicName, value=dumps(data).encode('ascii'))
            producer.flush()
            print(f"send data number: {e+1}")

            sleep(5)
    except Exception as e:
        print(e)
        raise
    

def run_consumer():
    try:
        consumer = KafkaConsumer('my-topic-three',
                                 bootstrap_servers=['localhost:9091'],
                                 auto_offset_reset='earliest',
                                 enable_auto_commit=True)
    except Exception as e:
        print(e)
        raise
    for message in consumer:
        print(message.value)


def run_quotes():
    try:
        return quotes.run()
    except Exception as e:
        print(e)
        raise