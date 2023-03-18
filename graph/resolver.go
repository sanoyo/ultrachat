package graph

import (
	"github.com/sanoyo/ultrachat/aws"
)

type Resolver struct {
	dynamoClient aws.DynamoDBRepository
}

func NewResolver(dynamoDBRepo aws.DynamoDBRepository) *Resolver {
	return &Resolver{
		dynamoClient: dynamoDBRepo,
	}
}
