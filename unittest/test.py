import requests

files = {'file': open("C:\\Users\\Vink\\Desktop\\temp\\aaa.mp3", 'rb')}


req=requests.post(url="http://192.168.3.10:6588/YouDaoManager/tool/MusicUpload",files=files)

