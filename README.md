# go-ec2reporter
CLI to get usefull Informations from AWS EC2.<br>
Its possible to search by Name and Status and define outputformats.

### Example usage:
Show all Instance with "web" in Name and State "running"

    go-ec2reporter -name web -state running

flags:

    -name       # search instances by Name
    -out        # output format (table, json, block)
    -state      # search instance with state (running, pending, stop)

###Example outputs:
table:

     Name        InstanceID   Status   PrivateIP        PublicIP     
     webserver1  i-123456787  running  xxx.xxx.xxx.xxx  xxx.xxx.xxx.xxx
     webserver2  i-123456788  stop     xxx.xxx.xxx.xxx  xxx.xxx.xxx.xxx
     databse1    i-123456789  running  xxx.xxx.xxx.xxx  xxx.xxx.xxx.xxx

block:

    Instance Name: webserver1
      instanceID: 123456787
      InstanceType: t2.micro
      Instance Name: eu-central-1a
      AMI Image: ami-12345
      Launched at: 2017-01-01 09:15:28 +0000 UTC
      Status: running
      Private IP: xxx.xxx.xxx.xxx
      Public IP: xxx.xxx.xxx.xxx
      Monitoring: enabled

    Instance Name: webserver2
      instanceID: 123456788
      InstanceType: t2.micro
      Instance Name: eu-central-1a
      AMI Image: ami-12345
      Launched at: 2017-01-01 09:15:28 +0000 UTC
      Status: stop
      Private IP: xxx.xxx.xxx.xxx
      Public IP: xxx.xxx.xxx.xxx
      Monitoring: disabled

### install:
dependencies:

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"


    go get github.com/ypsman/go-ec2reporter
    go install github.com/ypsman/go-ec2reporter
