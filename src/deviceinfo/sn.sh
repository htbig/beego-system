#!/bin/bash
snline=$(cat /proc/cpuinfo|grep Serial)
sn=${snline##*:}
newstr=`tr '[a-z]' '[A-Z]' <<<"$sn"`
echo $newstr
