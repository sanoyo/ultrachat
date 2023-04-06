package graph

import (
	"github.com/sanoyo/ultrachat/aws"
	"github.com/sanoyo/ultrachat/repository"
)

type Resolver struct {
	dynamoClient aws.DynamoDBRepository
	spaceClient  repository.SpaceRepository
}

func NewResolver(dynamoDBRepo aws.DynamoDBRepository, spaceRepo repository.SpaceRepository) *Resolver {
	return &Resolver{
		dynamoClient: dynamoDBRepo,
		spaceClient:  spaceRepo,
	}
}
