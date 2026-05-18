#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "Root required"
    exit 1
fi

echo "Creating service user"
useradd --system --shell /bin/nologin doh-proxyd
passwd -l doh-proxyd

echo "Compiling and installing"
go build
install -m 700 -o doh-proxyd -g doh-proxyd doh-proxy /usr/local/bin/doh-proxyd
rm -f doh-proxy

echo "Creating systemd service"
cat > /etc/systemd/system/doh-proxy.service << EOF
[Unit]
Description=DNS-DoH proxy
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/doh-proxyd
Restart=on-failure
User=doh-proxyd
AmbientCapabilities=CAP_NET_BIND_SERVICE
CapabilityBoundingSet=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
#systemctl enable --now doh-proxy.service

echo -e "Done\n"
echo "Enable and start service: $ systemctl enable --now doh-proxy"
