package lynx

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func AwsCreds() {
	// assume role
	svc := sts.New(sts.Options{
		Credentials: nil,
		Region:      "",
	})

	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String("roleArn"),
		RoleSessionName: aws.String("roleSessionName")}

	out, err := svc.AssumeRole(context.TODO(), input)
	if err != nil {
		fmt.Println("aws assume role %s: %v", "roleArn", err)
	}

	fmt.Println(out)

}
