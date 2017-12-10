#!/bin/bash
snline=$(cat /proc/cpuinfo|grep vendor_id)
sn=${snline##*:}
newstr=`tr '[a-z]' '[A-Z]' <<<"$sn"`
echo $newstr
