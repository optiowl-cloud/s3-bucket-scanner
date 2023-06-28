package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/exp/slices"
)

// BucketInfo represents the information about an S3 bucket
type BucketInfo struct {
	Name                     string `json:"name"`
	AccelerateConfig         *s3.GetBucketAccelerateConfigurationOutput
	ACL                      *s3.GetBucketAclOutput
	AnalyticsConfig          []*s3.GetBucketAnalyticsConfigurationOutput
	CORSConfig               *s3.GetBucketCorsOutput
	EncryptionConfig         *s3.GetBucketEncryptionOutput
	IntelligentTieringConfig []*s3.GetBucketIntelligentTieringConfigurationOutput
	InventoryConfig          []*s3.GetBucketInventoryConfigurationOutput
	LifecycleConfig          *s3.GetBucketLifecycleConfigurationOutput
	Location                 *s3.GetBucketLocationOutput
	LoggingConfig            *s3.GetBucketLoggingOutput
	MetricsConfig            []*s3.GetBucketMetricsConfigurationOutput
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
		if slices.Contains(
			strings.Split(os.Getenv("EXCLUDED_BUCKETS"), ","),
			aws.ToString(bucket.Name),
		) {
			continue
		}
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
			log.Fatalf("failed to get bucket accelerate configuration: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.AccelerateConfig = accelerateConfigResp
		}

		// Get bucket ACL
		aclResp, err := client.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Fatalf("failed to get bucket ACL: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.ACL = aclResp
		}

		// Get bucket analytics configuration
		analyticsConfigurationsResp, err := client.ListBucketAnalyticsConfigurations(
			context.TODO(),
			&s3.ListBucketAnalyticsConfigurationsInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Fatalf("failed to list bucket analytics configurations: %+v, %s", err, *bucket.Name)
			continue
		}

		for _, analyticsConfiguration := range analyticsConfigurationsResp.AnalyticsConfigurationList {
			analyticsConfigResp, err := client.GetBucketAnalyticsConfiguration(
				context.TODO(),
				&s3.GetBucketAnalyticsConfigurationInput{
					Bucket: bucket.Name,
					Id:     analyticsConfiguration.Id,
				},
			)
			if err != nil {
				log.Fatalf(
					"failed to get bucket analytics configuration: %v, %s",
					err,
					*bucket.Name,
				)
				continue
			}

			bucketInfo.AnalyticsConfig = append(bucketInfo.AnalyticsConfig, analyticsConfigResp)
		}

		// Get bucket CORS configuration
		corsConfigResp, err := client.GetBucketCors(context.TODO(), &s3.GetBucketCorsInput{
			Bucket: bucket.Name,
		})
		if err != nil && !strings.Contains(err.Error(), "NoSuchCORSConfiguration") {
			log.Fatalf("failed to get bucket CORS configuration: %+v, %s", err, *bucket.Name)
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
			log.Fatalf("failed to get bucket encryption configuration: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.EncryptionConfig = encryptionConfigResp
		}
		// Get bucket intelligent tiering configuration
		intelligentTieringConfigurationsResp, err := client.ListBucketIntelligentTieringConfigurations(
			context.TODO(),
			&s3.ListBucketIntelligentTieringConfigurationsInput{
				Bucket: bucket.Name,
			},
		)

		if err != nil {
			log.Fatalf(
				"failed to list bucket intelligent tiering configurations: %v, %s",
				err,
				*bucket.Name,
			)
			continue
		}

		for _, intelligentTieringConfiguration := range intelligentTieringConfigurationsResp.IntelligentTieringConfigurationList {
			intelligentTieringConfigResp, err := client.GetBucketIntelligentTieringConfiguration(
				context.TODO(),
				&s3.GetBucketIntelligentTieringConfigurationInput{
					Bucket: bucket.Name,
					Id:     intelligentTieringConfiguration.Id,
				},
			)
			if err != nil {
				log.Fatalf(
					"failed to get bucket intelligent tiering configuration: %v, %s",
					err,
					*bucket.Name,
				)
				continue
			}

			bucketInfo.IntelligentTieringConfig = append(
				bucketInfo.IntelligentTieringConfig,
				intelligentTieringConfigResp,
			)
		}

		// Get bucket inventory configuration
		inventoryConfigurationsResp, err := client.ListBucketInventoryConfigurations(
			context.TODO(),
			&s3.ListBucketInventoryConfigurationsInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Fatalf("failed to list bucket inventory configurations: %+v, %s", err, *bucket.Name)
			continue
		}

		for _, inventoryConfiguration := range inventoryConfigurationsResp.InventoryConfigurationList {
			inventoryConfigResp, err := client.GetBucketInventoryConfiguration(
				context.TODO(),
				&s3.GetBucketInventoryConfigurationInput{
					Bucket: bucket.Name,
					Id:     inventoryConfiguration.Id,
				},
			)
			if err != nil {
				log.Fatalf(
					"failed to get bucket inventory configuration: %v, %s",
					err,
					*bucket.Name,
				)
				continue
			}

			bucketInfo.InventoryConfig = append(bucketInfo.InventoryConfig, inventoryConfigResp)
		}

		// Get bucket lifecycle configuration
		lifecycleConfigResp, err := client.GetBucketLifecycleConfiguration(
			context.TODO(),
			&s3.GetBucketLifecycleConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil && !strings.Contains(err.Error(), "NoSuchLifecycleConfiguration") {
			log.Fatalf("failed to get bucket lifecycle configuration: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.LifecycleConfig = lifecycleConfigResp
		}

		// Get bucket location
		locationResp, err := client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Fatalf("failed to get bucket location: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.Location = locationResp
		}

		// Get bucket logging configuration
		loggingConfigResp, err := client.GetBucketLogging(context.TODO(), &s3.GetBucketLoggingInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Fatalf("failed to get bucket logging configuration: %+v, %s", err, *bucket.Name)
		} else {
			bucketInfo.LoggingConfig = loggingConfigResp
		}
		// Get bucket metrics configuration
		metricsConfigurationsResp, err := client.ListBucketMetricsConfigurations(
			context.TODO(),
			&s3.ListBucketMetricsConfigurationsInput{
				Bucket: bucket.Name,
			},
		)

		if err != nil {
			log.Fatalf("failed to list bucket metrics configurations: %+v, %s", err, *bucket.Name)
			continue
		}

		for _, metricsConfiguration := range metricsConfigurationsResp.MetricsConfigurationList {
			metricsConfigResp, err := client.GetBucketMetricsConfiguration(
				context.TODO(),
				&s3.GetBucketMetricsConfigurationInput{
					Bucket: bucket.Name,
					Id:     metricsConfiguration.Id,
				},
			)
			if err != nil {
				log.Fatalf("failed to get bucket metrics configuration: %+v, %s", err, *bucket.Name)
				continue
			}

			bucketInfo.MetricsConfig = append(bucketInfo.MetricsConfig, metricsConfigResp)
		}

		// Get bucket notification configuration
		notificationConfigResp, err := client.GetBucketNotificationConfiguration(
			context.TODO(),
			&s3.GetBucketNotificationConfigurationInput{
				Bucket: bucket.Name,
			},
		)
		if err != nil {
			log.Fatalf(
				"failed to get bucket notification configuration: %+v, %s",
				err,
				*bucket.Name,
			)
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
		if err != nil && !strings.Contains(err.Error(), "OwnershipControlsNotFoundError") {
			log.Fatalf(
				"failed to get bucket ownership controls configuration: %v, %s",
				err,
				*bucket.Name,
			)
		} else {
			bucketInfo.OwnershipControlsConfig = ownershipControlsConfigResp
		}

		// Get bucket policy
		policyResp, err := client.GetBucketPolicy(context.TODO(), &s3.GetBucketPolicyInput{
			Bucket: bucket.Name,
		})
		if err != nil && !strings.Contains(err.Error(), "NoSuchBucketPolicy") {
			log.Fatalf("failed to get bucket policy: %+v, %s", err, *bucket.Name)
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
		if err != nil && !strings.Contains(err.Error(), "NoSuchBucketPolicy") {
			log.Fatalf("failed to get bucket policy status: %+v, %s", err, *bucket.Name)
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
		if err != nil && !strings.Contains(err.Error(), "ReplicationConfigurationNotFoundError") {
			log.Fatalf("failed to get bucket replication configuration: %+v, %s", err, *bucket.Name)
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
			log.Fatalf(
				"failed to get bucket request payment configuration: %v, %s",
				err,
				*bucket.Name,
			)
		} else {
			bucketInfo.RequestPaymentConfig = requestPaymentConfigResp
		}

		// Get bucket tagging configuration
		taggingConfigResp, err := client.GetBucketTagging(context.TODO(), &s3.GetBucketTaggingInput{
			Bucket: bucket.Name,
		})
		if err != nil && !strings.Contains(err.Error(), "NoSuchTagSet") {
			log.Fatalf("failed to get bucket tagging configuration: %+v, %s", err, *bucket.Name)
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
			log.Fatalf("failed to get bucket versioning configuration: %+v, %s", err, *bucket.Name)
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
