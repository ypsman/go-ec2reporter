package ec2reporter

import (
	"encoding/json"
	"fmt"
)

func outouttable() {
	tablespec := " %-17s %-19s %-8s %-15s %-15s \n"
	fmt.Printf(tablespec, "Name", "InstanceID ", "Status ", "PrivateIP ", "PublicIP ")
	for _, inst := range Ec2List {
		fmt.Printf(tablespec, inst.Name, inst.InstID, inst.State, inst.PrivateIP, inst.PublicIP)
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
		fmt.Println("   Private IP: " + inst.PrivateIP)
		fmt.Println("   Public IP: " + inst.PublicIP)
		fmt.Println("   Monitoring: " + inst.Monitoring)
	}
	fmt.Println(" ")
}

func outputjson() {
	marshalled, err := json.Marshal(Ec2List)
	checkError(err)
	fmt.Println(string(marshalled))
}
