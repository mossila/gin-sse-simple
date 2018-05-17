# gin-sse-simple




install 

```
go get github.com/mossila/gin-sse-simple
```

Or clone and sync with `govendor`
```
git clone https://github.com/mossila/gin-sse-simple.git
cd gin-sse-simple/
govendor sync
```

usage

```
go run time_stream.go
```


website client

```
http://localhost:8888/
```


test with raw data.

```
curl http://localhost:8888/time/
```


result

```
event:message
data:2015-09-09 13:01:30.278489263 +0700 ICT

event:message
data:2015-09-09 13:01:31.275659459 +0700 ICT

event:message
data:2015-09-09 13:01:32.27942686 +0700 ICT

event:message
data:2015-09-09 13:01:33.279488214 +0700 ICT

event:message
data:2015-09-09 13:01:34.277990689 +0700 ICT
```

## Performance
 
Tested with [sse-pref](https://github.com/matthiasn/sse-perf)

####Hardware
Yeah, that my macbook.

![alt tag](img/hardware.png)

## Connection

* max connection is about 2555 (maybe OSX 's max connections configured?).
* max speed about 50 MB/s (I tested on localhost. Why it has limits?)
* max msg/s is  about 80-90K msg/s.

![alt tag](img/connection.png)

This image test configured with larger message size.
![alt tag](img/max-connection.png)

## CPU

Cpu usage Instant drop after close all connection 

![alt tag](img/cpu-usage.png)
