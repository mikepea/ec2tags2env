package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getInstanceID() string {
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, err := svc.GetInstanceIdentityDocument()
	if err != nil {
		panic(err)
	}
	return doc.InstanceID
}

func getRegion() string {
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, err := svc.GetInstanceIdentityDocument()
	if err != nil {
		panic(err)
	}
	return doc.Region
}

func getTagsForInstanceID(id string) *ec2.DescribeTagsOutput {
	region := getRegion()
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []*string{
					aws.String(id),
				},
			},
		},
	}
	result, err := svc.DescribeTags(input)
	if err != nil {
		panic(err)
	}
	return result

}

func main() {
	id := getInstanceID()
	fmt.Println(id)
	fmt.Println(getTagsForInstanceID(id))
}
