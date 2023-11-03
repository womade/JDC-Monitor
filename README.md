# 安装InfluxDB
```
apt install influxdb
```


# 创建数据库
```
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE USER admin WITH PASSWORD 'admin' WITH ALL PRIVILEGES"
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=RETENTION POLICY "30days" ON JDCDB DURATION 30d REPLICATION 1 DEFAULT"
```
