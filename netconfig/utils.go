// Package netconfig provides ...
package netconfig

import (
	"encoding/json"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/kjelly/netc/node"
)

func Generate(template string, n node.Node) map[string]string {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	var context pongo2.Context
	r := make(map[string]string)

	firstDeviceName := ""
	for i, ifname := range n.Ifnames {
		nodeJSONString, _ := json.Marshal(n)
		json.Unmarshal(nodeJSONString, &context)
		context["Ifname"] = ifname
		switch context["Kind"] {
		case "eth":
			context["Type"] = "Ethernet"
		case "vlan":
			context["Type"] = "vlan"

		}
		if i == 0 {
			firstDeviceName = context["Ifname"].(string)
			if context["Kind"] == "bond" {
				context["Type"] = "bond"
			}
		} else {
			if context["Kind"] == "bond" {
				context["Master"] = firstDeviceName
				context["Type"] = "Ethernet"
			}
		}
		out, err := tpl.Execute(context)
		if err != nil {
			panic(err)
		}
		r[fmt.Sprintf("ifcfg-%s", ifname)] = out

	}
	return r
}
