import requests
r = requests.post('localhost:8080/pictures/', files={'tempfile': open('tempfile', 'rb')})
print r.text
