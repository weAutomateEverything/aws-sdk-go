// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/aws/awserr"
	"github.com/weAutomateEverything/aws-sdk-go/aws/session"
	"github.com/weAutomateEverything/aws-sdk-go/service/dynamodb"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To retrieve multiple items from a table
//
// This example reads multiple items from the Music table using a batch of three GetItem
// requests.  Only the AlbumTitle attribute is returned.
func ExampleDynamoDB_BatchGetItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			"Music": {
				Keys: []map[string]*dynamodb.AttributeValue{
					{
						"Artist": &dynamodb.AttributeValue{
							S: aws.String("No One You Know"),
						},
						"SongTitle": &dynamodb.AttributeValue{
							S: aws.String("Call Me Today"),
						},
					},
					{
						"Artist": &dynamodb.AttributeValue{
							S: aws.String("Acme Band"),
						},
						"SongTitle": &dynamodb.AttributeValue{
							S: aws.String("Happy Day"),
						},
					},
					{
						"Artist": &dynamodb.AttributeValue{
							S: aws.String("No One You Know"),
						},
						"SongTitle": &dynamodb.AttributeValue{
							S: aws.String("Scared of My Shadow"),
						},
					},
				},
				ProjectionExpression: aws.String("AlbumTitle"),
			},
		},
	}

	result, err := svc.BatchGetItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add multiple items to a table
//
// This example adds three new items to the Music table using a batch of three PutItem
// requests.
func ExampleDynamoDB_BatchWriteItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"Music": {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"AlbumTitle": {
								S: aws.String("Somewhat Famous"),
							},
							"Artist": {
								S: aws.String("No One You Know"),
							},
							"SongTitle": {
								S: aws.String("Call Me Today"),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"AlbumTitle": {
								S: aws.String("Songs About Life"),
							},
							"Artist": {
								S: aws.String("Acme Band"),
							},
							"SongTitle": {
								S: aws.String("Happy Day"),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"AlbumTitle": {
								S: aws.String("Blue Sky Blues"),
							},
							"Artist": {
								S: aws.String("No One You Know"),
							},
							"SongTitle": {
								S: aws.String("Scared of My Shadow"),
							},
						},
					},
				},
			},
		},
	}

	result, err := svc.BatchWriteItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a table
//
// This example creates a table named Music.
func ExampleDynamoDB_CreateTable_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Artist"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("SongTitle"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Artist"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("SongTitle"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Music"),
	}

	result, err := svc.CreateTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete an item
//
// This example deletes an item from the Music table.
func ExampleDynamoDB_DeleteItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Artist": {
				S: aws.String("No One You Know"),
			},
			"SongTitle": {
				S: aws.String("Scared of My Shadow"),
			},
		},
		TableName: aws.String("Music"),
	}

	result, err := svc.DeleteItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a table
//
// This example deletes the Music table.
func ExampleDynamoDB_DeleteTable_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String("Music"),
	}

	result, err := svc.DeleteTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To determine capacity limits per table and account, in the current AWS region
//
// The following example returns the maximum read and write capacity units per table,
// and for the AWS account, in the current AWS region.
func ExampleDynamoDB_DescribeLimits_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.DescribeLimitsInput{}

	result, err := svc.DescribeLimits(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a table
//
// This example describes the Music table.
func ExampleDynamoDB_DescribeTable_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String("Music"),
	}

	result, err := svc.DescribeTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To read an item from a table
//
// This example retrieves an item from the Music table. The table has a partition key
// and a sort key (Artist and SongTitle), so you must specify both of these attributes.
func ExampleDynamoDB_GetItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Artist": {
				S: aws.String("Acme Band"),
			},
			"SongTitle": {
				S: aws.String("Happy Day"),
			},
		},
		TableName: aws.String("Music"),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list tables
//
// This example lists all of the tables associated with the current AWS account and
// endpoint.
func ExampleDynamoDB_ListTables_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add an item to a table
//
// This example adds a new item to the Music table.
func ExampleDynamoDB_PutItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"AlbumTitle": {
				S: aws.String("Somewhat Famous"),
			},
			"Artist": {
				S: aws.String("No One You Know"),
			},
			"SongTitle": {
				S: aws.String("Call Me Today"),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("Music"),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To query an item
//
// This example queries items in the Music table. The table has a partition key and
// sort key (Artist and SongTitle), but this query only specifies the partition key
// value. It returns song titles by the artist named "No One You Know".
func ExampleDynamoDB_Query_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String("No One You Know"),
			},
		},
		KeyConditionExpression: aws.String("Artist = :v1"),
		ProjectionExpression:   aws.String("SongTitle"),
		TableName:              aws.String("Music"),
	}

	result, err := svc.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To scan a table
//
// This example scans the entire Music table, and then narrows the results to songs
// by the artist "No One You Know". For each item, only the album title and song title
// are returned.
func ExampleDynamoDB_Scan_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames: map[string]*string{
			"AT": aws.String("AlbumTitle"),
			"ST": aws.String("SongTitle"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String("No One You Know"),
			},
		},
		FilterExpression:     aws.String("Artist = :a"),
		ProjectionExpression: aws.String("#ST, #AT"),
		TableName:            aws.String("Music"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update an item in a table
//
// This example updates an item in the Music table. It adds a new attribute (Year) and
// modifies the AlbumTitle attribute.  All of the attributes in the item, as they appear
// after the update, are returned in the response.
func ExampleDynamoDB_UpdateItem_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#AT": aws.String("AlbumTitle"),
			"#Y":  aws.String("Year"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":t": {
				S: aws.String("Louder Than Ever"),
			},
			":y": {
				N: aws.String("2015"),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Artist": {
				S: aws.String("Acme Band"),
			},
			"SongTitle": {
				S: aws.String("Happy Day"),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		TableName:        aws.String("Music"),
		UpdateExpression: aws.String("SET #Y = :y, #AT = :t"),
	}

	result, err := svc.UpdateItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To modify a table's provisioned throughput
//
// This example increases the provisioned read and write capacity on the Music table.
func ExampleDynamoDB_UpdateTable_shared00() {
	svc := dynamodb.New(session.New())
	input := &dynamodb.UpdateTableInput{
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("MusicCollection"),
	}

	result, err := svc.UpdateTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
