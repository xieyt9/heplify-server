# 部署说明

- `mkdir -p /home/homer/config`
- `vi /home/homer/config/env`
```conf
ENV_DB_USER=homer
ENV_DB_PASSWORD=123456
ENV_DB_HOST=test-db.com
ENV_DB_PORT=3306
ENV_LISTEN_HOST=0.0.0.0
ENV_DB_BULK=1000
```
- `sh -x deploy.sh`