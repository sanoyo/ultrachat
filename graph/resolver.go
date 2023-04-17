package graph

import (
	"github.com/sanoyo/ultrachat/aws"
	"github.com/sanoyo/ultrachat/repository"
)

type Resolver struct {
	dynamoClient aws.DynamoDBRepository
	spaceClient  repository.SpaceRepository
	userClient   repository.UserRepository
}

func NewResolver(dynamoDBRepo aws.DynamoDBRepository,
	spaceRepo repository.SpaceRepository,
	userRepo repository.UserRepository) *Resolver {
	return &Resolver{
		dynamoClient: dynamoDBRepo,
		spaceClient:  spaceRepo,
		userClient:   userRepo,
	}
}
