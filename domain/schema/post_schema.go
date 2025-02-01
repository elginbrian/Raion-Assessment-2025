package schema

import "github.com/graphql-go/graphql"

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"userId":     &graphql.Field{Type: graphql.ID},
		"caption":    &graphql.Field{Type: graphql.String},
		"imageURL":   &graphql.Field{Type: graphql.String},
		"createdAt":  &graphql.Field{Type: graphql.String},
		"updatedAt":  &graphql.Field{Type: graphql.String},
	},
})

var PostQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllPosts": &graphql.Field{
			Type: graphql.NewList(PostType),
		},
		"getPostByID": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
		},
		"getPostsByUserID": &graphql.Field{
			Type: graphql.NewList(PostType),
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{Type: graphql.ID},
			},
		},
	},
})

var PostMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"caption": &graphql.ArgumentConfig{Type: graphql.String},
				// "image":   &graphql.ArgumentConfig{Type: graphql.Upload}, 
			},
		},
		"updatePostCaption": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"id":      &graphql.ArgumentConfig{Type: graphql.ID},
				"caption": &graphql.ArgumentConfig{Type: graphql.String},
			},
		},
		"deletePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
		},
	},
})