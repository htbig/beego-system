deviceinfo-arm:
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm CGO_ENABLED=1 go build -gcflags "-N -l" -i  *.go
deviceinfo:
	go build -gcflags "-N -l" -i *.go
install:
	if [ ! -d "/usr/lib/vega/deviceinfo/conf" ]; then mkdir -p /usr/lib/vega/deviceinfo/conf; fi
	if [ ! -d "/usr/lib/vega/deviceinfo/views" ]; then mkdir -p /usr/lib/vega/deviceinfo/views; fi
	if [ ! -d "/usr/lib/vega/deviceinfo/static/js" ]; then mkdir -p /usr/lib/vega/deviceinfo/static/js; fi
	cp ./main /usr/lib/vega/deviceinfo/deviceinfo
	cp ./conf/app.conf /usr/lib/vega/deviceinfo/conf/
	cp ./static/js/reload.min.js /usr/lib/vega/deviceifo/static/js/
	cp ./views/index.tpl /usr/lib/vega/deviceinfo/views/
	cp ./sn.sh /usr/lib/vega/deviceinfo/
	if [ ! -d "/etc/systemd/system/multi-user.target.wants/" ]; then mkdir -p /etc/systemd/system/multi-user.target.wants; fi
	if [ ! -d "/lib/systemd/system/" ]; then mkdir -p /lib/systemd/system/; fi
	cp ./deviceinfo.service /lib/systemd/system/
	if [ -f "/etc/systemd/system/multi-user.target.wants/deviceinfo.service" ]; then rm -f /etc/systemd/system/multi-user.target.wants/deviceinfo.service; fi
	ln -s /lib/systemd/system/deviceinfo.service /etc/systemd/system/multi-user.target.wants/deviceinfo.service;
deb:
	checkinstall -D -A armhf --pkgname=deviceinfo --pkgversion=2017-12-6 --install=no  --pkgsource=../deviceinfo
clean:
	-rm main deviceinfo
