// Package main provides ...
package node

type Node struct {
	TargetIP string `json:"target_ip"`
	User     string
	KeyPath  string `json:"key_path"`
	Kind     string

	Ifnames  []string
	Type     string
	IP4      string
	Netmask4 string
	Gateway4 string
	Mode     string
	OnBoot   string
}

type Data struct {
	Nodes []Node
}
