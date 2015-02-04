package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/rds"
)

var region = "us-west-2"

// Connect will provide a valid RDS client
func Connect() *rds.RDS {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	return rds.New(creds, region, nil)
}

// ListDBInstances will print out the RDS instances
func ListDBInstances(db *rds.RDS) {
	desc, err := db.DescribeDBInstances(nil)
	if err != nil {
		log.Fatalln(err)
	}
	for i, d := range desc.DBInstances {
		fmt.Printf("%2d: %v\n", i+1, *d.DBInstanceIdentifier)
	}
}

var rfs = &rds.RestoreDBInstanceFromDBSnapshotMessage{
	AutoMinorVersionUpgrade: aws.Boolean(false),
	AvailabilityZone:        aws.String("us-west-2a"),
	DBInstanceClass:         aws.String("db.t2.micro"),
	DBInstanceIdentifier:    aws.String("eolastester"),
	DBName:                  aws.String("eolas"),
	DBSnapshotIdentifier:    aws.String("sg-e7f4a782"),
	DBSubnetGroupName:       aws.String("default"),
	Engine:                  aws.String("MySQL 5.6.19a"),
	LicenseModel:            aws.String("General Public License"),
	MultiAZ:                 aws.Boolean(false),
	OptionGroupName:         aws.String("default:mysql-5-6"),
	Port:                    aws.Integer(3306),
	PubliclyAccessible:      aws.Boolean(true),
	//IOPS:                    aws.Integer(0),
	//StorageType:             aws.String(""),
	//Tags:                    []aws.Tag{},
	//TDECredentialARN:      aws.String(""),
	//TDECredentialPassword: aws.String(""),
}

// RestoreSnapshot will restore the DB instance from a snapshot
func RestoreFromSnapshot(db *rds.RDS) {
	resp, err := db.RestoreDBInstanceFromDBSnapshot(rfs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*resp)
}

func main() {

}
