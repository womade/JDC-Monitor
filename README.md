# 安装InfluxDB
```
apt install influxdb
```


# 创建数据库
```
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE USER admin WITH PASSWORD 'admin' WITH ALL PRIVILEGES"
```

```
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE DATABASE JDCDB"
```

```
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE RETENTION POLICY month ON JDCDB DURATION 30d REPLICATION 1 DEFAULT"
```
