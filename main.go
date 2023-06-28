package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// BucketInfo represents the information about an S3 bucket
type BucketInfo struct {
	Name                     string `json:"name"`
	AccelerateConfig         *s3.GetBucketAccelerateConfigurationOutput
	ACL                      *s3.GetBucketAclOutput
	AnalyticsConfig          *s3.GetBucketAnalyticsConfigurationOutput
	CORSConfig               *s3.GetBucketCorsOutput
	EncryptionConfig         *s3.GetBucketEncryptionOutput
	IntelligentTieringConfig *s3.GetBucketIntelligentTieringConfigurationOutput
	InventoryConfig          *s3.GetBucketInventoryConfigurationOutput
	LifecycleConfig          *s3.GetBucketLifecycleConfigurationOutput
	Location                 *s3.GetBucketLocationOutput
	LoggingConfig            *s3.GetBucketLoggingOutput
	MetricsConfig            *s3.GetBucketMetricsConfigurationOutput
	NotificationConfig       *s3.GetBucketNotificationConfigurationOutput
	OwnershipControlsConfig  *s3.GetBucketOwnershipControlsOutput
	Policy                   *s3.GetBucketPolicyOutput
	PolicyStatus             *s3.GetBucketPolicyStatusOutput
	ReplicationConfig        *s3.GetBucketReplicationOutput
	RequestPaymentConfig     *s3.GetBucketRequestPaymentOutput
	TaggingConfig            *s3.GetBucketTaggingOutput
	VersioningConfig         *s3.GetBucketVersioningOutput
}

func main() {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List S3 buckets
	resp, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("failed to list S3 buckets: %v", err)
	}

	// Create a slice to hold bucket information
	bucketInfos := []BucketInfo{}

	// Iterate over buckets and retrieve information
	for _, bucket := range resp.Buckets {
		bucketInfo := BucketInfo{
			Name: aws.ToString(bucket.Name),
		}

		// Get bucket accelerate configuration
		accelerateConfigResp, err := client.GetBucketAccelerateConfiguration(
			context.TODO(),
			&s3.GetBucketAccelerateConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Fatalf("failed to get bucket accelerate configuration: %v", err)
		} else {
			bucketInfo.AccelerateConfig = accelerateConfigResp
		}

		// Get bucket ACL
		aclResp, err := client.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Fatalf("failed to get bucket ACL: %v", err)
		} else {
			bucketInfo.ACL = aclResp
		}

		// Get bucket analytics configuration
		analyticsConfigResp, err := client.GetBucketAnalyticsConfiguration(
			context.TODO(),
			&s3.GetBucketAnalyticsConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket analytics configuration: %v", err)
		} else {
			bucketInfo.AnalyticsConfig = analyticsConfigResp
		}

		// Get bucket CORS configuration
		corsConfigResp, err := client.GetBucketCors(context.TODO(), &s3.GetBucketCorsInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("failed to get bucket CORS configuration: %v", err)
		} else {
			bucketInfo.CORSConfig = corsConfigResp
		}

		// Get bucket encryption configuration
		encryptionConfigResp, err := client.GetBucketEncryption(
			context.TODO(),
			&s3.GetBucketEncryptionInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket encryption configuration: %v", err)
		} else {
			bucketInfo.EncryptionConfig = encryptionConfigResp
		}

		// Get bucket intelligent tiering configuration
		intelligentTieringConfigResp, err := client.GetBucketIntelligentTieringConfiguration(
			context.TODO(),
			&s3.GetBucketIntelligentTieringConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket intelligent tiering configuration: %v", err)
		} else {
			bucketInfo.IntelligentTieringConfig = intelligentTieringConfigResp
		}

		// Get bucket inventory configuration
		inventoryConfigResp, err := client.GetBucketInventoryConfiguration(
			context.TODO(),
			&s3.GetBucketInventoryConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket inventory configuration: %v", err)
		} else {
			bucketInfo.InventoryConfig = inventoryConfigResp
		}

		// Get bucket lifecycle configuration
		lifecycleConfigResp, err := client.GetBucketLifecycleConfiguration(
			context.TODO(),
			&s3.GetBucketLifecycleConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket lifecycle configuration: %v", err)
		} else {
			bucketInfo.LifecycleConfig = lifecycleConfigResp
		}

		// Get bucket location
		locationResp, err := client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("failed to get bucket location: %v", err)
		} else {
			bucketInfo.Location = locationResp
		}

		// Get bucket logging configuration
		loggingConfigResp, err := client.GetBucketLogging(context.TODO(), &s3.GetBucketLoggingInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("failed to get bucket logging configuration: %v", err)
		} else {
			bucketInfo.LoggingConfig = loggingConfigResp
		}

		// Get bucket metrics configuration
		metricsConfigResp, err := client.GetBucketMetricsConfiguration(
			context.TODO(),
			&s3.GetBucketMetricsConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket metrics configuration: %v", err)
		} else {
			bucketInfo.MetricsConfig = metricsConfigResp
		}

		// Get bucket notification configuration
		notificationConfigResp, err := client.GetBucketNotificationConfiguration(
			context.TODO(),
			&s3.GetBucketNotificationConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket notification configuration: %v", err)
		} else {
			bucketInfo.NotificationConfig = notificationConfigResp
		}

		// Get bucket ownership controls configuration
		ownershipControlsConfigResp, err := client.GetBucketOwnershipControls(
			context.TODO(),
			&s3.GetBucketOwnershipControlsInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket ownership controls configuration: %v", err)
		} else {
			bucketInfo.OwnershipControlsConfig = ownershipControlsConfigResp
		}

		// Get bucket policy
		policyResp, err := client.GetBucketPolicy(context.TODO(), &s3.GetBucketPolicyInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("failed to get bucket policy: %v", err)
		} else {
			bucketInfo.Policy = policyResp
		}

		// Get bucket policy status
		policyStatusResp, err := client.GetBucketPolicyStatus(
			context.TODO(),
			&s3.GetBucketPolicyStatusInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket policy status: %v", err)
		} else {
			bucketInfo.PolicyStatus = policyStatusResp
		}

		// Get bucket replication configuration
		replicationConfigResp, err := client.GetBucketReplication(
			context.TODO(),
			&s3.GetBucketReplicationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket replication configuration: %v", err)
		} else {
			bucketInfo.ReplicationConfig = replicationConfigResp
		}

		// Get bucket request payment configuration
		requestPaymentConfigResp, err := client.GetBucketRequestPayment(
			context.TODO(),
			&s3.GetBucketRequestPaymentInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket request payment configuration: %v", err)
		} else {
			bucketInfo.RequestPaymentConfig = requestPaymentConfigResp
		}

		// Get bucket tagging configuration
		taggingConfigResp, err := client.GetBucketTagging(context.TODO(), &s3.GetBucketTaggingInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("failed to get bucket tagging configuration: %v", err)
		} else {
			bucketInfo.TaggingConfig = taggingConfigResp
		}

		// Get bucket versioning configuration
		versioningConfigResp, err := client.GetBucketVersioning(
			context.TODO(),
			&s3.GetBucketVersioningInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Printf("failed to get bucket versioning configuration: %v", err)
		} else {
			bucketInfo.VersioningConfig = versioningConfigResp
		}

		bucketInfos = append(bucketInfos, bucketInfo)
	}

	// Create the output file
	file, err := os.Create("bucket_info.json")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer file.Close()

	// Write bucket information to JSON file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(bucketInfos)
	if err != nil {
		log.Fatalf("failed to encode bucket information: %v", err)
	}

	fmt.Println("Bucket information written to bucket_info.json")
}
