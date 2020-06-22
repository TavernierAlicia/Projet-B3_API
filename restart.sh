#!/bin/bash
PID=`sudo netstat -plutn | grep "/./main" | grep ":9999" | tr -s ' ' | cut -d ' ' -f7 | cut -d '/' -f1`
# echo $PID
sudo kill -9 $PID
sudo ./main &

