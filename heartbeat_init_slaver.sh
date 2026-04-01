#!/usr/bin/env bash

heartbeat_init() {
	if [ -e "./heartbeat.zip" ]; then
		unzip -q -o ./heartbeat.zip -d /home/admin/
		chmod +x /home/admin/heartbeat/*
		chown -R admin:admin /home/admin/heartbeat/

		cat >/etc/systemd/system/heartbeat.service <<EOF
[Unit]
Description=Heart Beat Service
After=network.target
Wants=network.target

[Service]
Type=simple
User=admin
WorkingDirectory=/home/admin/heartbeat/
ExecStart=/home/admin/heartbeat/heart_beat_arm64 server

[Install]
WantedBy=multi-user.target
EOF
		systemctl daemon-reload
		systemctl enable --now heartbeat.service

	else
		echo "heartbeat.zip不在目录下"
	fi
}
heartbeat_init
