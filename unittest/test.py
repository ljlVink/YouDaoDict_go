import requests
import json
'''
files_mp3 = {'file': open("C:\\Users\\Vink\\Desktop\\temp\\aaa.mp3", 'rb')}
req=requests.post(url="http://192.168.3.10:6588/YouDaoManager/tool/MusicUpload",files=files_mp3)
data=json.dumps({"filename":"aaa.mp3"})
print(data)
req=requests.post(url="http://192.168.3.10:6588/YouDaoManager/tool/MusicRemove",data=data)
print(req.text)
'''

files_zip = {'file': open("C:\\Users\\Vink\\Desktop\\imagescat.zip", 'rb')}
req=requests.post(url="http://192.168.3.10:6588/YouDaoManager/tool/ApplySkin",files=files_zip)

