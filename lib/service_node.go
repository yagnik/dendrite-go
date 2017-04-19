package io

// package io

// import "strings"
// import "fmt"
// import "encoding/json"

// type ServiceNode struct {
// 	Organization   string            `json:"organization"`
// 	Component      string            `json:"component"`
// 	Name           string            `json:"name"`
// 	LeadEmail      string            `json:"lead_email"`
// 	TeamEmail      string            `json:"team_email"`
// 	ServiceType    string            `json:"type"`
// 	Deploy         Deploy            `json:"deploy"`
// 	DefaultServers []DefaultServer   `json:"default_servers"`
// 	Scale          Scale             `json:"scale"`
// 	Ports          []Port            `json:"ports"`
// 	Dependencies   []Dependency      `json:"dependencies"`
// 	Telemetry      Telemetry         `json:"telemetry"`
// 	Metadata       map[string]string `json:"metadata"`
// }

// type Deploy struct {
// 	Repository string `json:"repository"`
// 	Package    string `json:"package"`
// }

// type DefaultServer struct {
// 	Environment string `json:"environment"`
// 	Host        string `json:"host"`
// 	Port        int    `json:"port"`
// }

// type Scale struct {
// 	MaxInstanceCount int `json:"max_instance_count"`
// 	MinInstanceCount int `json:"min_instance_count"`
// 	MinMemory        int `json:"min_memory"`
// 	MinCpu           int `json:"min_cpu"`
// }

// type Port struct {
// 	Name string `json:"name"`
// 	Port int    `json:"port"`
// }

// type Dependency struct {
// 	Identifier  string `json:"identifier"`
// 	Latency     int    `json:"latency"`
// 	ServiceNode *ServiceNode
// }

// type Telemetry struct {
// 	HealthEndpoint     string   `json:"health_endpoint"`
// 	NotificationEmails []string `json:"notification_emails"`
// }

// func (servicenode ServiceNode) FullName() string {
// 	return strings.Join([]string{servicenode.Organization, servicenode.Component, servicenode.Name}, "_")
// }

// func (servicenode ServiceNode) Validate() bool {
// 	return false
// }

// func main() {
// 	s := &ServiceNode{
// 		Component:    "sd",
// 		Organization: "coms",
// 		Name:         "admin",
// 		Deploy:       Deploy{Package: "foobar"},
// 	}
// 	b, _ := json.Marshal(s)
// 	fmt.Println(string(b))
// }
