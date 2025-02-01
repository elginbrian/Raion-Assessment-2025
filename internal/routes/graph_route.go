package routes

import (
	"bytes"
	"log"
	"net/http"
	schema "raion-assessment/domain/schema"
	"raion-assessment/internal/di"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func SetupGraphQLRoute(app *fiber.App, container di.Container) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: mergeFields(
			convertFieldDefinitionMap(schema.UserQueryType.Fields()),
			convertFieldDefinitionMap(schema.PostQueryType.Fields()),
		),
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: mergeFields(
			convertFieldDefinitionMap(schema.UserMutationType.Fields()),
			convertFieldDefinitionMap(schema.PostMutationType.Fields()),
		),
	})

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema: &graphqlSchema,
		Pretty: true,
	})

	app.Post("/api/v1/graphql", func(c *fiber.Ctx) error {
		body := c.Body()
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create request")
		}

		req.Header = make(http.Header)
		c.Request().Header.VisitAll(func(k, v []byte) {
			req.Header.Set(string(k), string(v))
		})

		rec := &responseRecorder{
			header:     make(http.Header),
			statusCode: http.StatusOK,
		}

		graphQLHandler.ServeHTTP(rec, req)

		for k, v := range rec.header {
			c.Set(k, v[0])
		}

		return c.Status(rec.statusCode).Send(rec.body)
	})
}

func mergeFields(fieldMaps ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}
	for _, fields := range fieldMaps {
		for key, value := range fields {
			merged[key] = value
		}
	}
	return merged
}

func convertFieldDefinitionMap(fieldMap graphql.FieldDefinitionMap) graphql.Fields {
	fields := graphql.Fields{}
	for key, value := range fieldMap {
		fields[key] = &graphql.Field{
			Type:              value.Type,
			Args:              convertArgs(value.Args),
			Resolve:           value.Resolve,
			Description:       value.Description,
			DeprecationReason: value.DeprecationReason,
		}
	}
	return fields
}

func convertArgs(args []*graphql.Argument) graphql.FieldConfigArgument {
	fieldConfigArgs := graphql.FieldConfigArgument{}
	for _, arg := range args {
		fieldConfigArgs[arg.Name()] = &graphql.ArgumentConfig{
			Type:         arg.Type,
			DefaultValue: arg.DefaultValue,
			Description:  arg.Description(),
		}
	}
	return fieldConfigArgs
}

type responseRecorder struct {
	header     http.Header
	body       []byte
	statusCode int
}

func (r *responseRecorder) Header() http.Header {
	return r.header
}

func (r *responseRecorder) Write(data []byte) (int, error) {
	r.body = data
	return len(data), nil
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}