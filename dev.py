from kafka import KafkaProducer
from kafka import KafkaConsumer

producer = KafkaProducer(bootstrap_servers='localhost:9092')
consumer = KafkaConsumer('verifiedData', bootstrap_servers='localhost:9092')

#user input
userInput = raw_input('Enter Number:')
producer.send('userData', value = str.encode(userInput))

#print ops result
for message in consumer:
	print(message.value)
	break
