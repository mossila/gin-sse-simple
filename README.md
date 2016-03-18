# gin-sse-simple

usage
```
go run time_stream.go
```

test
```
curl http://localhost:8080/time/
```

website client
```
http://localhost:8080/client/
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

Tested with [sse-pref][https://github.com/matthiasn/sse-perf]

####Hardware
Yeah, that my macbook.

![alt tag](img/hardware.png)

####Connection

touch the max transferate 

![alt tag](img/connection.png)

####CPU

Cpu usage Instant drop after close all connection 

![alt tag](img/cpu-usage.png)