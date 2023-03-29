package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

var (
	methodGET     string = "GET"
	methodPOST    string = "POST"
	methodPUT     string = "PUT"
	methodPATCH   string = "PATCH"
	methodDELETE  string = "DELETE"
	methodOPTIONS string = "OPTIONS"
	methodHEAD    string = "HEAD"

	allMethods []*string = []*string{
		&methodPOST,
		&methodHEAD,
		&methodPATCH,
		&methodDELETE,
		&methodPUT,
		&methodGET,
		&methodOPTIONS,
	}
)

func NewCloudfrontClient() (*cloudfront.CloudFront, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})

	if err != nil {
		return nil, err
	}
	client := cloudfront.New(sess)
	return client, nil
}
