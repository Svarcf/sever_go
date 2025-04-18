package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Svarcf/sever_go/internal/graph/model"
	"github.com/Svarcf/sever_go/internal/models"
)

// CreateToolRepairRecord is the resolver for the createToolRepairRecord field.
func (r *mutationResolver) CreateToolRepairRecord(ctx context.Context, createToolRepairRecordInput *model.CreateToolRepairRecordInput) (*models.ToolRepairRecordDTO, error) {
	// Validate required fields
	if createToolRepairRecordInput.Tool == nil {
		return nil, fmt.Errorf("tool is required")
	}
	toolID := *createToolRepairRecordInput.Tool

	//TODO: Uncomment and implement user validation after authentication is done
	// if createToolRepairRecordInput.User == nil {
	//	return nil, fmt.Errorf("user is required")
	//}
	userID := uint(1)

	// For string pointers, use empty string if nil
	repairDescription := ""
	if createToolRepairRecordInput.RepairDescription != nil {
		repairDescription = *createToolRepairRecordInput.RepairDescription
	}

	note := ""
	if createToolRepairRecordInput.Note != nil {
		note = *createToolRepairRecordInput.Note
	}

	externalService := ""
	if createToolRepairRecordInput.ExternalServices != nil {
		externalService = *createToolRepairRecordInput.ExternalServices
	}

	var materialID uint
	if createToolRepairRecordInput.Material != nil {
		materialID = *createToolRepairRecordInput.Material
	}

	toolRepairRecord := models.ToolRepairRecord{
		RepairDescription:      repairDescription,
		Note:                   note,
		ExternalServices:       externalService,
		ToolID:                 toolID,
		MaterialID:             materialID,
		UserID:                 userID,
		MalfunctionDescription: *createToolRepairRecordInput.MalfunctionDescription,
		TimeSpent:              *createToolRepairRecordInput.TimeSpent,
	}

	// Handle date fields carefully - these might need to be nil
	if createToolRepairRecordInput.DateStarted != nil {
		toolRepairRecord.DateStarted = *createToolRepairRecordInput.DateStarted
	}

	if createToolRepairRecordInput.DateEnded != nil {
		toolRepairRecord.DateEnded = sql.NullTime{Time: *createToolRepairRecordInput.DateEnded, Valid: true}
	} else {
		toolRepairRecord.DateEnded = sql.NullTime{Valid: false}
	}

	if createToolRepairRecordInput.DeadlineDate != nil {
		toolRepairRecord.DeadlineDate = *createToolRepairRecordInput.DeadlineDate
	}

	return r.ToolRepairRecordService.CreateToolRepairRecord(&toolRepairRecord)
}

// ToolRepairRecords is the resolver for the toolRepairRecords field.
func (r *queryResolver) ToolRepairRecords(ctx context.Context) ([]*models.ToolRepairRecordDTO, error) {
	return r.ToolRepairRecordService.GetToolRepairRecords()
}

// ToolRepairRecord is the resolver for the toolRepairRecord field.
func (r *queryResolver) ToolRepairRecord(ctx context.Context, id uint) (*models.ToolRepairRecordDTO, error) {
	return r.ToolRepairRecordService.GetToolRepairRecordById(id)
}

// RawMaterial is the resolver for the rawMaterial field.
func (r *toolRepairRecordResolver) RawMaterial(ctx context.Context, obj *models.ToolRepairRecordDTO) (*models.RawMaterialDTO, error) {
	return r.RawMaterialService.GetRawMaterialById(obj.MaterialID)
}

// Tool is the resolver for the tool field.
func (r *toolRepairRecordResolver) Tool(ctx context.Context, obj *models.ToolRepairRecordDTO) (*models.ToolDTO, error) {
	return r.ToolService.GetToolById(obj.ToolID)
}

// User is the resolver for the user field.
func (r *toolRepairRecordResolver) User(ctx context.Context, obj *models.ToolRepairRecordDTO) (*models.UserDTO, error) {
	return r.UserService.GetUserById(obj.UserID)
}

// ToolRepairRecord returns ToolRepairRecordResolver implementation.
func (r *Resolver) ToolRepairRecord() ToolRepairRecordResolver { return &toolRepairRecordResolver{r} }

type toolRepairRecordResolver struct{ *Resolver }
