package main

import "fmt"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws/ec2metadata"

func getInstanceID() string {
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, err := svc.GetInstanceIdentityDocument()
	if err != nil {
		panic(err)
	}
	return doc.InstanceID
}

func main() {
	fmt.Println(getInstanceID())
}
