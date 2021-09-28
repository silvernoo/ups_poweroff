# ups_poweroff

ping脚本指向市电路由，无ping断电关机。市电路由加开机唤醒脚本，来电开机

# build

```
go build -o /bin/ups_poweroff .
```

# usage

```
-g string
  	Gateway
-i int
  	Interval (Second) (default 3)
-r uint
  	Retry count (default 10)
```

# wol脚本
https://gist.github.com/vincascm/7fa72a8c27933736d802
