package schema

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"name":       &graphql.Field{Type: graphql.String},
		"email":      &graphql.Field{Type: graphql.String},
		"bio":        &graphql.Field{Type: graphql.String},
		"imageURL":   &graphql.Field{Type: graphql.String},
		"createdAt":  &graphql.Field{Type: graphql.String},
		"updatedAt":  &graphql.Field{Type: graphql.String},
	},
})

var UserQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllUsers": &graphql.Field{
			Type: graphql.NewList(UserType),
		},
		"getUserByID": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
		},
		"searchUsers": &graphql.Field{
			Type: graphql.NewList(UserType),
			Args: graphql.FieldConfigArgument{
				"query": &graphql.ArgumentConfig{Type: graphql.String},
			},
		},
	},
})

var UserMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"updateUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{Type: graphql.String},
				"bio":      &graphql.ArgumentConfig{Type: graphql.String},
				// "image":    &graphql.ArgumentConfig{Type: graphql.Upload},
			},
		},
	},
})