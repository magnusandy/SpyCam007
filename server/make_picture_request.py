import requests
r = requests.post('http://laforgesplayground.appspot.com/pictures/', files={'picture': open('img.png', 'rb')})
print r.text
