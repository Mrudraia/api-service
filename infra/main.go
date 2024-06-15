package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		sgArgs := &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(8080),
					ToPort:     pulumi.Int(8080),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
				ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(22),
					ToPort:     pulumi.Int(22),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				ec2.SecurityGroupEgressArgs{
					Protocol:   pulumi.String("-1"),
					FromPort:   pulumi.Int(0),
					ToPort:     pulumi.Int(0),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
		}
		sg, err := ec2.NewSecurityGroup(ctx, "jenkins-sg", sgArgs)
		if err != nil {
			return nil
		}

		kp, err := ec2.NewKeyPair(ctx, "local-ssh", &ec2.KeyPairArgs{
			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCpc9kKT9Z8Wy8xW0246gitPLyC836FG5tJ3+xs8HzJbwNUOjDurmnegq0wy7OGy36074vUmfOv4oIYmZVVFBlULKUTA81ihNpyEyK9NAlC+ME/7Fhj3ffPC0xdsYLdcO1oVbYGbv2grXtIoEqOke/lOd5Cc2Sl2C25HbXL2IxVg5XDcawkUPNtLkxvh2+Hdcu60eFN3apsOj2PhywYrOowchdeFNBzadwg0SfOx0gcLAJnqvEnABzPdWkJ/eXhXuvXUe7ue4iLysCO+kfmH53S4vMj0KMl+eOqXml5m0UtluAaVAKOov+EcPpzp5/0Mga1a+348wj+dkg+w61eC+sIDTuJjLCRIkP8w4YagrUYVlERXgJtJ9DwEWG3BOTsAGC1363oJd4B1A/gZcsrJTgiiCxrVLf8AuUGM/MJ+8Z4jZWVLA+nQmIXo2heDJeLAte9aVe0/fF5pHtUYSHZcbSfhsqFbNI8Xxk5J5rTXswSLs20Jcx2gEK2xVGu4Eo//9M= mrudraia@li-df9ef54c-2877-11b2-a85c-9a6a73b854ef.ibm.com"),
		})

		if err != nil {
			return nil
		}

		jenkinsServer, err := ec2.NewInstance(ctx, "jenkins-server", &ec2.InstanceArgs{
			InstanceType:        pulumi.String("t2.micro"),
			VpcSecurityGroupIds: pulumi.StringArray{sg.ID()},
			Ami:                 pulumi.String("ami-0583d8c7a9c35822c"),
			KeyName:             kp.KeyName,
		})

		fmt.Println(jenkinsServer.PublicIp)
		fmt.Println(jenkinsServer.PublicDns)

		ctx.Export("publicIp", jenkinsServer.PublicIp)
		ctx.Export("publicHostName", jenkinsServer.PublicDns)

		return nil
	})
}
