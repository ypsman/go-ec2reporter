package ec2reporter

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	flagout := flag.String("out", "json", "table, json or block")
	flag.Parse()
	fmt.Println(*flagout)
	fmt.Println("ec2reporter v.02")
	result := connector()
	resultWorker(result)
	//outformat := argcheck()

	outputter(*flagout)
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
	fmt.Println(ec2format)
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
