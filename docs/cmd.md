# firewalld-cmd 和 dbus 接口使用示例

**获取所有可配置的service**

命令行

```bash
$ firewall-cmd --get-services
RH-Satellite-6 RH-Satellite-6-capsule amanda-client amanda-k5-client amqp amqps apcupsd audit bacula bacula-client bb bgp bitcoin bitcoin-rpc bitcoin-testnet bitcoin-testnet-rpc bittorrent-lsd ceph ceph-mon cfengine cockpit collectd condor-collector ctdb dhcp dhcpv6 dhcpv6-client distcc dns dns-over-tls docker-registry docker-swarm dropbox-lansync elasticsearch etcd-client etcd-server finger foreman foreman-proxy freeipa-4 freeipa-ldap freeipa-ldaps freeipa-replication freeipa-trust ftp galera ganglia-client ganglia-master git grafana gre high-availability http https imap imaps ipp ipp-client ipsec irc ircs iscsi-target isns jenkins kadmin kdeconnect kerberos kibana klogin kpasswd kprop kshell kube-api kube-apiserver kube-control-plane kube-controller-manager kube-scheduler kubelet-worker ldap ldaps libvirt libvirt-tls lightning-network llmnr managesieve matrix mdns memcache minidlna mongodb mosh mountd mqtt mqtt-tls ms-wbt mssql murmur mysql nbd netbios-ns nfs nfs3 nmea-0183 nrpe ntp nut openvpn ovirt-imageio ovirt-storageconsole ovirt-vmconsole plex pmcd pmproxy pmwebapi pmwebapis pop3 pop3s postgresql privoxy prometheus proxy-dhcp ptp pulseaudio puppetmaster quassel radius rdp redis redis-sentinel rpc-bind rquotad rsh rsyncd rtsp salt-master samba samba-client samba-dc sane sip sips slp smtp smtp-submission smtps snmp snmptrap spideroak-lansync spotify-sync squid ssdp ssh steam-streaming svdrp svn syncthing syncthing-gui synergy syslog syslog-tls telnet tentacle tftp tile38 tinc tor-socks transmission-client upnp-client vdsm vnc-server wbem-http wbem-https wireguard wsman wsmans xdmcp xmpp-bosh xmpp-client xmpp-local xmpp-server zabbix-agent zabbix-server
```

运行时dbus

```bash
$ dbus-send --system --dest=org.fedoraproject.FirewallD1    --print-reply --type=method_call    /org/fedoraproject/FirewallD1   org.fedoraproject.FirewallD1.listServices
```

配置dbus

```bash
[root@localhost 003listservices]# dbus-send --system --dest=org.fedoraproject.FirewallD1  --print-reply --type=method_call    /org/fedoraproject/FirewallD1/config  org.fedoraproject.FirewallD1.config.listServices
method return time=1730180069.413534 sender=:1.4 -> destination=:1.29274 serial=334 reply_serial=2
   array [
      object path "/org/fedoraproject/FirewallD1/config/service/0"
      object path "/org/fedoraproject/FirewallD1/config/service/1"
      object path "/org/fedoraproject/FirewallD1/config/service/2"
      object path "/org/fedoraproject/FirewallD1/config/service/3"
	...
      object path "/org/fedoraproject/FirewallD1/config/service/180"
      object path "/org/fedoraproject/FirewallD1/config/service/181"
   ]
```

**获取list-all**

```bash
[root@localhost ~]# dbus-send --system --dest=org.fedoraproject.FirewallD1 --print-reply --type=method_call  /org/fedoraproject/FirewallD1 org.fedoraproject.FirewallD1.getZoneSettings string:"public"
method return time=1730200853.534319 sender=:1.19 -> destination=:1.86275 serial=1346 reply_serial=2
   struct {
      string ""
      string "Public"
      string "For use in public areas. You do not trust the other computers on networks to not harm your computer. Only selected incoming connections are accepted."
      boolean false
      string "default"
      array [
         string "ssh"
         string "mdns"
         string "dhcpv6-client"
      ]
      array [
      ]
      array [
      ]
      boolean false
      array [
      ]
      array [
         string "ens3"
      ]
      array [
      ]
      array [
      ]
      array [
      ]
      array [
      ]
      boolean false
   }
[root@localhost ~]# firewall-cmd --zone=public --list-all
public (active)
  target: default
  icmp-block-inversion: no
  interfaces: ens3
  sources: 
  services: dhcpv6-client mdns ssh
  ports: 
  protocols: 
  masquerade: no
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules
```

