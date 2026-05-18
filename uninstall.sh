#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "Root required"
    exit 1
fi

echo "Stopping and disabling service"
systemctl stop doh-proxy
systemctl disable doh-proxy

echo "Removing unit file"
rm -rf /etc/systemd/system/doh-proxy.service
systemctl daemon-reload

echo "Removing binary"
rm /usr/local/bin/doh-proxyd

echo "Removing service user"
userdel doh-proxyd

echo -e "\nDone"
