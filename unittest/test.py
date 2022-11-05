import requests
req=requests.post(url="http://192.168.3.10:6588/panic")
print(req.text)