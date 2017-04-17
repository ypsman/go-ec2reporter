package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Flags for output, name, state
var flagFormat = flag.String("out", "table", "table, json or block for Outputformat")
var flagTNames = flag.String("name", "*", "Search for Tag")
var flagStatus = flag.String("state", "*", "Search for state: running, pending, stop")

//var flagInstID = flag.String("instid", "", "search for InstanceID")

func main() {
	flag.Parse()
	fmt.Println("ec2reporter v.02")
	result := getInstances(*flagTNames, *flagStatus)
	resultWorker(result)
	// switch for output
	switch *flagFormat {
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

// getInstances : connect to aws and get results
func getInstances(iname, istatus string) *ec2.DescribeInstancesOutput {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create EC2 client
	ec2sess := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String("*" + iname + "*"),
				},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("*" + istatus + "*")},
			},
		},
		// not possible to search with Wildcard
		//InstanceIds: []*string{
		//	aws.String(""), // Required
		//},
	}
	result, err := ec2sess.DescribeInstances(params)
	checkError(err)
	return result
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
