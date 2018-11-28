// Package main
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kjelly/netc/netconfig"
	"github.com/kjelly/netc/node"
	"github.com/kjelly/netc/sshlib"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	jsonBype, _ := ioutil.ReadFile("./node.json")

	dec := json.NewDecoder(strings.NewReader(string(jsonBype)))

	var m node.Data
	_ = dec.Decode(&m)

	o := netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[0])
	fmt.Printf("%v\n", o)
	o = netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[1])
	fmt.Printf("%v\n", o)
	o = netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[2])
	fmt.Printf("%v\n", o)

	for _, v := range m.Nodes {
		configContent := netconfig.Generate(netconfig.CentOSTemplate, v)

		for fileName, fileContent := range configContent {
			path := fmt.Sprintf("/tmp/%s", fileName)
			ioutil.WriteFile(path, []byte(fileContent), 0655)
			sshlib.CopyFile(v.TargetIP, v.User, v.KeyPath, path, path)
			command := fmt.Sprintf("sudo mv /tmp/%s /etc/sysconfig/network-scripts/%s", fileName, fileName)
			sshlib.ExecuteCommand(command, v.TargetIP, v.User, v.KeyPath)
		}
	}
}
