package io

import "github.com/ghodss/yaml"
import "fmt"

type ServiceNode struct {
	Organization string `json:"organization"`
	Component    string `json:"component"`
}

func Run() {
	// Marshal a Person struct to YAML.
	p := ServiceNode{"FOO", "BAR"}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	   age: 30
	   name: John
	*/

	// Unmarshal the YAML back into a Person struct.
	var p2 interface{}
	err = yaml.Unmarshal(y, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)
	/* Output:
	   {John 30}
	*/
}
