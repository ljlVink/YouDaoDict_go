@echo off

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
SET GOARM=7

go build 
adb push YouDaoManager /userdata
adb shell chmod 777 /userdata/YouDaoManager
adb shell killall YouDaoManager
del YouDaoManager

adb shell /userdata/YouDaoManager