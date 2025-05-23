package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"

	"github.com/Svarcf/sever_go/internal/graph/model"
	"github.com/Svarcf/sever_go/internal/models"
)

// CreateRawMaterial is the resolver for the createRawMaterial field.
func (r *mutationResolver) CreateRawMaterial(ctx context.Context, createRawMaterialInput *model.CreateRawMaterialInput) (*models.RawMaterialDTO, error) {
	rawMaterial := models.NewRawMaterial(
		createRawMaterialInput.Code,
		createRawMaterialInput.Name,
		createRawMaterialInput.Number)
	return r.RawMaterialService.CreateRawMaterial(rawMaterial)
}

// UpdateRawMaterial is the resolver for the updateRawMaterial field.
func (r *mutationResolver) UpdateRawMaterial(ctx context.Context, updateRawMaterialInput *model.UpdateRawMaterialInput) (*models.RawMaterialDTO, error) {
	rawMaterial := models.NewRawMaterial(
		updateRawMaterialInput.Code,
		updateRawMaterialInput.Name,
		updateRawMaterialInput.Number)
	rawMaterial.Id = updateRawMaterialInput.ID
	return r.RawMaterialService.UpdateRawMaterial(rawMaterial)
}

// RawMaterials is the resolver for the rawMaterials field.
func (r *queryResolver) RawMaterials(ctx context.Context) ([]*models.RawMaterialDTO, error) {
	return r.RawMaterialService.GetRawMaterials()
}

// RawMaterial is the resolver for the rawMaterial field.
func (r *queryResolver) RawMaterial(ctx context.Context, id uint) (*models.RawMaterialDTO, error) {
	return r.RawMaterialService.GetRawMaterialById(id)
}
