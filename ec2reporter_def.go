package ec2reporter

import (
	"fmt"
	"os"
	"time"
)

// Ec2Instance : struct for EC2 Instances
type Ec2Instance struct {
	Name       string
	InstID     string
	ImageID    string
	InstType   string
	Launch     time.Time
	State      string
	PublicIP   string
	PrivateIP  string
	Monitoring string
	Region     string
}

// InstancesList : array with struct from Instances
type InstancesList []Ec2Instance

// Ec2List : var from InstancesList
var Ec2List InstancesList
var ec2format string

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
