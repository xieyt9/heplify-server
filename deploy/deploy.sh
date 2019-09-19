#!/bin/bash

function usage(){
    echo "----------------------------------------usage---------------------------------------------------"
    echo "-----------------please select config type------------------------------------------------------"
    echo "-----------------0: out --------- --------------------------------------------------------------"
    echo "-----------------1: deploy mysql----------------------------------------------------------------"
    echo "-----------------2: deploy homer----------------------------------------------------------------"
    echo "-----------------3: update homer----------------------------------------------------------------"
    echo "-----------------9: start ----------------------------------------------------------------------"
}

function homer_default_env(){
    if [ ! -d /home/homer/config ];then
        isConfig="true"
        mkdir -p /home/homer/config
    fi
    envfile=/home/homer/config/env
    if [ ! -f $envfile ];then
      touch $envfile
	  echo "请输入数据库密码:"
	  read  PASSWORD
      echo 'ENV_DB_USER=root' >> $envfile
      echo 'ENV_DB_PASSWORD='$PASSWORD >> $envfile
      echo 'ENV_DB_HOST=127.0.0.1' >> $envfile
      echo 'ENV_DB_PORT=3306' >> $envfile
      echo 'ENV_LISTEN_HOST=0.0.0.0' >> $envfile
      echo 'ENV_DB_BULK=800' >> $envfile
      echo 'ENV_DB_TIMER=5' >> $envfile
    fi
}
function deploy_mysql(){
    echo "请输入数据库密码:"
	read  PASSWORD
    docker run  --restart always  \
    --privileged=true \
    --name mysql-56 \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=$PASSWORD \
    -d registry.cn-beijing.aliyuncs.com/tinet-hub/homer-mysql-56-centos7:latest
    
    docker ps -a
}
DEFAULT_ENV_FILE=/home/homer/config/env
DEFAULT_DOCKER_TAG=latest

function deploy_homer(){
    echo "----------------set homer TAG defalt [latest]---------------------------------------------------------"
    read CONFIG_DOCKER_TAG
    if [  "X$CONFIG_DOCKER_TAG" = "X" ];then
      CONFIG_DOCKER_TAG=$DEFAULT_DOCKER_TAG
    fi
    docker run --ulimit nofile=90000:90000 \
    --restart always  --name homer --network="host" \
    --env-file /home/homer/config/env \
    --log-driver json-file --log-opt max-size=10m --log-opt max-file=7 \
    -d registry.cn-beijing.aliyuncs.com/tinet-hub/homer:$CONFIG_DOCKER_TAG
    
    docker ps -a
}

function main()
{
    for((; ; ))  
    do  
        usage
        echo "please set config type:"
        read  TYPE
        case $TYPE in
        0|exit)
            echo 'You about quit'
            exit
        ;;
        1)
            echo 'deploy mysql 5.6'
            deploy_mysql
        ;;
        2)
            echo 'deploy homer'
            homer_default_env
            deploy_homer
			      exit
        ;;
		    3)
            echo 'update homer'
            homer_default_env
			      docker rm -f homer
            deploy_homer
			      exit
        ;;
        9)
            start
            exit
        ;;
        *)
            echo 'set error type'
        ;;
        esac
    done 
}
main