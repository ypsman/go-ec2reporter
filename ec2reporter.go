package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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

func main() {
	fmt.Println("ec2reporter v.02")
	result := connector()
	resultWorker(result)
	outformat := argcheck()
	outputter(outformat)
}

func argcheck() string {
	if len(os.Args) == 2 {
		outformat := os.Args[1:][0]
		return outformat
	}
	outformat := "table"
	return outformat
}

func outputter(ec2format string) {
	switch ec2format {
	case "json":
		outputjson()
	case "block":
		outputrepot()
	case "table":
		outouttable()
	default:
		outouttable()
	}
}

func outouttable() {
	tablespec := " %-17s %-19s %-8s %-15s %-15s \n"
	fmt.Printf(tablespec, "Name", "InstanceID ", "Status ", "PrivateIP ", "PublicIP ")
	for _, inst := range Ec2List {
		fmt.Printf(tablespec, inst.Name, inst.InstID, inst.State, inst.PrivateIP, inst.PrivateIP)
	}
}

func outputrepot() {
	for _, inst := range Ec2List {
		fmt.Println(" ")
		fmt.Println(" Instance Name: " + inst.Name)
		fmt.Println("   instanceID: " + inst.InstID)
		fmt.Println("   InstanceType: " + inst.InstType)
		fmt.Println("   Instance Name: " + inst.Region)
		fmt.Println("   AMI Image: " + inst.ImageID)
		fmt.Println("   Launched at: " + inst.Launch.String())
		fmt.Println("   Status: " + inst.State)
		fmt.Println("   Public IP: " + inst.PublicIP)
		fmt.Println("   Private IP: " + inst.PrivateIP)
		fmt.Println("   Monitoring: " + inst.Monitoring)
	}
	fmt.Println(" ")
}

func outputjson() {
	marshalled, err := json.Marshal(Ec2List)
	checkError(err)
	fmt.Println(string(marshalled))
}

// Connector : Connect to AWS
func connector() *ec2.DescribeInstancesOutput {
	// Load session from config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create EC2 client
	ec2sess := ec2.New(sess)
	result, err := ec2sess.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	return (result)
}

func resultWorker(result *ec2.DescribeInstancesOutput) {
	var myInst Ec2Instance
	var name string
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			name = "None"
			for _, t := range instance.Tags {
				if *t.Key == "Name" {
					name = *t.Value
				}
			}
			myInst = Ec2Instance{
				Name:       name,
				InstID:     *instance.InstanceId,
				ImageID:    *instance.ImageId,
				InstType:   *instance.InstanceType,
				Launch:     *instance.LaunchTime,
				State:      *instance.State.Name,
				PublicIP:   *instance.PublicIpAddress,
				PrivateIP:  *instance.PrivateIpAddress,
				Monitoring: *instance.Monitoring.State,
				Region:     *instance.Placement.AvailabilityZone,
			}
			Ec2List = append(Ec2List, myInst)
		}
	}
}
