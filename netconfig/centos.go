// Package netconfig provides ...
package netconfig

const CentOSTemplate = `
DEVICE={{ Ifname }}
BOOTPROTO=none
ONBOOT={{ OnBoot }}
TYPE={{ Type }}
IPADDR={{ IP4 }}
NETMASK={{ Netmask4 }}
{% if Gateway4 %}
GATEWAY={{ Gateway4 }}
{% endif %}
IPV6INIT=no
USERCTL=no

{% if Master %}
MASTER={{ Master }}
SLAVE=yes
{% endif %}
`
