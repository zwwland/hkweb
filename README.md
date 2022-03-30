cgo c reference to the C library

[]: # Language: markdown
[]: # Path: README.md

https://blog.csdn.net/dong_beijing/article/details/79721222

http://222.134.214.134:8005/

222.134.214.134
// 0.9
8105
//0.106
8102


转码命令

ffmpeg -i audio.webm -ac 1 -ar 64000 -acodec pcm_s16le -f s16le output.raw


可以通过注册设备（NET_DVR_Login_V40）返回的设备信息NET_DVR_DEVICEINFO_V30获取
模拟通道个数（byChanNum）、
模拟通道起始通道号（byStartChan）和
设备支持的最大IP通道数（byIPChanNum+ byHighDChanNum*256）、
数字通道起始通道号（byStartDChan）。 

从byStartChan到byStartChan+byChanNum-1对应为模拟通道的通道号，
IP通道的通道号为byStartDChan到byStartDChan+ (byIPChanNum + byHighDChanNum*256) -1。
DVR、IPC、IPD只有模拟通道，NVR只有IP通道，混合型DVR同时支持模拟通道和IP通道。 一般IPC/IPD通道号为1，
32路以及以下路数的NVR的IP通道通道号从33开始，
64路及以上路数的NVR的IP通道通道号从1开始。

```json
set 1001:2:22 '{"startAt":1648029660,"lastAt":1648029660,"limitTime":300,"position":[110.22,168.4,-2],"regionId":12,"energyLeft":98,"nickName":"姓名","personnelId":123,"cardId":"卡ID","cardNumber":"11122","equipmentId":12,"equipmentSerial":"设备序列号"}'

set 1001:2:22 '{"startAt":1648029660,"lastAt":1648029660,"limitTime":300,"position":[110.22,168.4,-2],"regionId":12,"energyLeft":98,"nickName":"姓名","carNumber":"车牌","personnelId":123,"cardId":"卡ID","cardNumber":"11122","equipmentId":12,"equipmentSerial":"设备序列号"}'
```