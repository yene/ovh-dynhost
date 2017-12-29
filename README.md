# OVH DynHost Updater

This tool is here to help with updating the IP address of OVH DynHost.

### Example
`ovh-dynhost --username $USERNAME --password $PASSWORD update-record $HOSTNAME`

### Cross compile

```bash
make linux
make osx
make windows
make armv7
```

### Install systemd

```bash
sudo mv systemd/ovh-dynhost.* /etc/systemd/system/
sudo systemctl enable ovh-dynhost.timer
sudo systemctl start ovh-dynhost.timer
sudo systemctl enable ovh-dynhost.service
journalctl -f -u ovh-dynhost
```
