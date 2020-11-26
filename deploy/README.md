# 部署说明
## 部署mysql
docker run --restart always --name mysql -p 3306:3306 -e MYSQL_USER=homer -e MYSQL_PASSWORD=123456 -e MYSQL_DATABASE=homer_data --network=host --log-driver json-file --log-opt max-size=10m --log-opt max-file=7 -d centos/mysql-56-centos7:latest

## 部署homer

- `mkdir -p /home/homer/config`
- `vi /home/homer/config/env`
```conf
ENV_DB_USER=homer
ENV_DB_PASSWORD=123456
ENV_DB_HOST=test-db.com
ENV_DB_PORT=3306
ENV_LISTEN_HOST=0.0.0.0
ENV_DB_BULK=1000
ENV_DB_TIMER=5
```
- `sh -x deploy.sh`
