package main

import (
    "fmt"

    "log"

    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess := session.Must(session.NewSession(&aws.Config{
        Endpoint:                      aws.String("s3.solera-uk.services"),
        CredentialsChainVerboseErrors: aws.Bool(true),
        Region:                        aws.String("us-east-1")},
    ))

    // Region: aws.String("us-west-2")},

    // Create S3 service client
    svc := s3.New(sess)

    result, err := svc.ListBuckets(&s3.ListBucketsInput{})
    if err != nil {
        log.Println("Failed to list buckets", err)
        return
    }

    log.Println("Buckets:")

    for _, bucket := range result.Buckets {
        log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
    }

    // // Create an S3 Bucket
    // bucket := "testb1"
    // _, err = svc.CreateBucket(&s3.CreateBucketInput{
    //     Bucket: aws.String(bucket),
    // })
    // if err != nil {
    //     exitErrorf("Unable to create bucket %q, %v", bucket, err)
    // }

    // // Wait until bucket is created before finishing
    // fmt.Printf("Waiting for bucket %q to be created...\n", bucket)

    // err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
    //     Bucket: aws.String(bucket),
    // })
    // if err != nil {
    //     exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
    // }

    // fmt.Printf("Bucket %q successfully created\n", bucket)

    // Delete the S3 Bucket
    // It must be empty or else the call fails
    bucket := "testb1"
    _, err = svc.DeleteBucket(&s3.DeleteBucketInput{
        Bucket: aws.String(bucket),
    })
    if err != nil {
        exitErrorf("Unable to delete bucket %q, %v", bucket, err)
    }

    // Wait until bucket is deleted before finishing
    fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

    err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
        Bucket: aws.String(bucket),
    })
    if err != nil {
        exitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
    }

    fmt.Printf("Bucket %q successfully deleted\n", bucket)

}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
