#!/bin/sh

logfile=/var/log/safe_hepipe.log

while :
do
  count_hepipe_proc=`ps -ef|grep -w hepipe.js|wc -l`

  if [ ${count_hepipe_proc}  -lt 2 ]; then
    echo "$(date +%Y-%m-%d@%H:%M:%S) Warning! hepipe is crash." >> ${logfile}
    killall -9 node
    forever --sourceDir=/usr/src/ -p /var/log/forever start -o /var/log/hepipe/hepipe.log \
         -e /var/log/hepipe/error.log -a --minUptime 1000 --spinSleepTime 3000 hepipe.js
    sleep 1
  else
    echo "$(date +%Y-%m-%d@%H:%M:%S) hepipe is running" >> ${logfile}.$(date +%Y%m%d)
    sleep 10
  fi
done
