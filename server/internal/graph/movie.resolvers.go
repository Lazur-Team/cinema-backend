package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"server/internal/graph/model"
	"server/internal/models"

	"github.com/google/uuid"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.Movie, error) {
	// Generate a new UUID for the movie
	newMovie := models.Movie{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		ReleaseYear: int(input.ReleaseYear),
	}

	// Save to database
	if err := r.DB.Create(&newMovie).Error; err != nil {
		return nil, fmt.Errorf("failed to create movie: %w", err)
	}

	// Convert GORM model to GraphQL model
	return &model.Movie{
		ID:          newMovie.ID.String(),
		Title:       newMovie.Title,
		Description: newMovie.Description,
		ReleaseYear: int32(newMovie.ReleaseYear),
	}, nil
}

// UpdateMovie is the resolver for the updateMovie field.
func (r *mutationResolver) UpdateMovie(ctx context.Context, id string, input model.MovieInput) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented: UpdateMovie - updateMovie"))
}

// DeleteMovie is the resolver for the deleteMovie field.
func (r *mutationResolver) DeleteMovie(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteMovie - deleteMovie"))
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	// panic(fmt.Errorf("not implemented: Movies - movies"))
	return []*model.Movie{}, nil
}

// Movie is the resolver for the movie field.
func (r *queryResolver) Movie(ctx context.Context, id string) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented: Movie - movie"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
