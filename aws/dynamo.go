package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBRepository interface {
	PutItemWithContext(ctx context.Context, tableName string, model interface{}) error
	GetItemsWithContext(ctx context.Context, tableName string) ([]map[string]*dynamodb.AttributeValue, error)
}

type DynamoClient struct {
	dynamoClient *dynamodb.DynamoDB
}

func NewDynamoClient(endpoint string) *DynamoClient {
	client := dynamodb.New(
		session.Must(session.NewSession()),
		&aws.Config{
			Endpoint: aws.String(endpoint),
			Region:   aws.String("ap-northeast-1"),
		})
	return &DynamoClient{dynamoClient: client}
}

func (d *DynamoClient) GetItemsWithContext(ctx context.Context, tableName string) ([]map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := d.dynamoClient.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (d *DynamoClient) PutItemWithContext(ctx context.Context, tableName string, model interface{}) error {
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = d.dynamoClient.PutItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
