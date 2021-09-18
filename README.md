# OVH DynHost Updater

This tool is here to help with updating the IP address of OVH DynHost.
To manage DynHost in your OVH manager, go to [your domain](https://www.ovh.com/manager/#/domain), select tab DynHost. There you can create a DynHost and credentials.

### Example

`ovh-dynhost --username $USERNAME --password $PASSWORD update-record $HOSTNAME`

### Cross compile

```bash
make linux
make osx
make windows
make armv7
```

### Install on Raspberry Pi

```bash
curl -LO https://github.com/yene/ovh-dynhost/releases/download/v0.0.1/ovh-dynhost-armv7
sudo mv ovh-dynhost-armv7 /usr/local/bin/ovh-dynhost
sudo chmod +x /usr/local/bin/ovh-dynhost
sudo mv systemd/ovh-dynhost.* /etc/systemd/system/
# update credentials
sudo vim /etc/systemd/system/ovh-dynhost.service
sudo systemctl enable ovh-dynhost.timer
sudo systemctl start ovh-dynhost.timer
sudo systemctl enable ovh-dynhost.service
```

Verify with `systemctl list-timers`
See the logs with `journalctl -f -u ovh-dynhost`
