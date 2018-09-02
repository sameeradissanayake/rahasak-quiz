import re
from kafka import KafkaConsumer
from kafka import KafkaProducer

consumer = KafkaConsumer('userData', bootstrap_servers='localhost:9092')
producer = KafkaProducer(bootstrap_servers='localhost:9092')

try:
    for message in consumer:
    	rule = re.compile(r'(^[+0-9]{1,3})*([0-9]{10,11}$)')
        if rule.search(message.value):
        	producer.send('verifiedData', value = b'yes')
        else:
        	producer.send('verifiedData', value = b'no')

except:
	print "Error"

