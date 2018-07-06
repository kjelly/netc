// Package main
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kjelly/netc/netconfig"
	"github.com/kjelly/netc/node"
	"io/ioutil"
	"strings"
)

func main() {
	jsonBype, _ := ioutil.ReadFile("./node.json")

	dec := json.NewDecoder(strings.NewReader(string(jsonBype)))

	var m node.Data
	_ = dec.Decode(&m)

	fmt.Printf("%v\n", m)

	netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[0])
	netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[1])
	netconfig.Generate(netconfig.CentOSTemplate, m.Nodes[2])

}
