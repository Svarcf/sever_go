package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Svarcf/sever_go/internal/models"
	"gorm.io/gorm"
)

type SeedData struct {
	Roles             []models.Role            `json:"roles"`
	Users             []UserSeed               `json:"users"`
	Files             []models.File            `json:"files"`
	MechanicalPresses []models.MechanicalPress `json:"mechanical_presses"`
	RawMaterials      []models.RawMaterial     `json:"raw_materials"`
	StandardParts     []models.StandardPart    `json:"standard_parts"`
	ToolTypes         []models.ToolType        `json:"tool_types"`
	Tools             []ToolSeed               `json:"tools"`
	ToolRepairRecords []ToolRepairRecordSeed   `json:"tool_repair_records"`
}

// For Users, as we need to map by role name instead of role ID
type UserSeed struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	RoleName  string `json:"role_name"`
}

// For Tools, we need to map ToolType by name and ToolAssociation by name
type ToolSeed struct {
	Name                 string   `json:"name"`
	Code                 string   `json:"code"`
	Dimension            string   `json:"dimension"`
	Note                 string   `json:"note"`
	ToolStroke           string   `json:"tool_stroke"`
	WorkpieceDescription string   `json:"workpiece_description"`
	ToolTypeName         string   `json:"tool_type_name"`
	ToolAssociationName  string   `json:"tool_association_name"`
	RawMaterialNames     []string `json:"raw_material_names"`
	StandardPartNames    []string `json:"standard_part_names"`
	MechanicalPressNames []string `json:"mechanical_press_names"`
}

type ToolRepairRecordSeed struct {
	DateStarted            time.Time `json:"date_started"`
	DateEnded              time.Time `json:"date_ended"`
	RepairDescription      string    `json:"repair_description"`
	MalfunctionDescription string    `json:"malfunction_description"`
	DeadlineDate           time.Time `json:"deadline_date"`
	Material               uint      `json:"material"`
	TimeSpent              int       `json:"time_spent"`
	ExternalServices       string    `json:"external_services"`
	Note                   string    `json:"note"`
	ToolCode               string    `json:"tool_code"`
	RepairedByUsername     string    `json:"repaired_by_username"`
}

func seedDatabase(db *gorm.DB) {
	// Load seed data from file
	var seedData SeedData
	file, err := os.ReadFile("seed_data.json")
	if err != nil {
		panic(fmt.Sprintf("failed to read seed data: %v", err))
	}

	err = json.Unmarshal(file, &seedData)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal seed data: %v", err))
	}

	// Seed roles
	for _, role := range seedData.Roles {
		db.Create(&role)
	}

	// Seed users
	for _, user := range seedData.Users {
		var role models.Role
		db.First(&role, "name = ?", user.RoleName)
		db.Create(&models.User{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Password:  user.Password,
			RoleID:    role.Id,
		})
	}

	// Seed files
	for _, file := range seedData.Files {
		db.Create(&file)
	}

	// Seed mechanical presses
	for _, press := range seedData.MechanicalPresses {
		db.Create(&press)
	}

	// Seed raw materials
	for _, material := range seedData.RawMaterials {
		db.Create(&material)
	}

	// Seed standard parts
	for _, part := range seedData.StandardParts {
		db.Create(&part)
	}

	// Seed tool types
	for _, toolType := range seedData.ToolTypes {
		db.Create(&toolType)
	}

	// Seed tools
	for _, toolSeed := range seedData.Tools {
		var toolType models.ToolType
		db.First(&toolType, "name = ?", toolSeed.ToolTypeName)

		var toolAssociation *models.Tool
		if toolSeed.ToolAssociationName != "" {
			var associatedTool models.Tool
			db.First(&associatedTool, "name = ?", toolSeed.ToolAssociationName)
			toolAssociation = &associatedTool
		}

		// Fetch related raw materials
		var rawMaterials []*models.RawMaterial
		for _, materialName := range toolSeed.RawMaterialNames {
			var rawMaterial models.RawMaterial
			db.First(&rawMaterial, "name = ?", materialName)
			rawMaterials = append(rawMaterials, &rawMaterial)
		}

		// Fetch related standard parts
		var standardParts []*models.StandardPart
		for _, partName := range toolSeed.StandardPartNames {
			var standardPart models.StandardPart
			db.First(&standardPart, "name = ?", partName)
			standardParts = append(standardParts, &standardPart)
		}

		// Fetch related mechanical presses
		var mechanicalPresses []*models.MechanicalPress
		for _, pressName := range toolSeed.MechanicalPressNames {
			var mechanicalPress models.MechanicalPress
			db.First(&mechanicalPress, "name = ?", pressName)
			mechanicalPresses = append(mechanicalPresses, &mechanicalPress)
		}

		tool := models.Tool{
			Name:                 toolSeed.Name,
			Code:                 toolSeed.Code,
			Dimension:            toolSeed.Dimension,
			Note:                 toolSeed.Note,
			ToolStroke:           toolSeed.ToolStroke,
			WorkpieceDescription: toolSeed.WorkpieceDescription,
			ToolTypeID:           toolType.Id,
			ToolAssociationID:    nil, // This will stay nil if no association exists
			RawMaterials:         rawMaterials,
			StandardParts:        standardParts,
			MechanicalPresses:    mechanicalPresses,
		}

		if toolAssociation != nil {
			tool.ToolAssociationID = &toolAssociation.Id
		}

		db.Create(&tool)
	}

	// Seed tool repair records
	for _, repairSeed := range seedData.ToolRepairRecords {
		var tool models.Tool
		db.First(&tool, "code = ?", repairSeed.ToolCode)

		var user models.User
		db.First(&user, "username = ?", repairSeed.RepairedByUsername)

		repairRecord := models.ToolRepairRecord{
			DateStarted:            repairSeed.DateStarted,
			DateEnded:              sql.NullTime{Time: repairSeed.DateEnded, Valid: true},
			RepairDescription:      repairSeed.RepairDescription,
			MalfunctionDescription: repairSeed.MalfunctionDescription,
			DeadlineDate:           repairSeed.DeadlineDate,
			MaterialID:             repairSeed.Material,
			TimeSpent:              repairSeed.TimeSpent,
			ExternalServices:       repairSeed.ExternalServices,
			Note:                   repairSeed.Note,
			ToolID:                 tool.Id,
			UserID:                 user.Id,
		}
		db.Create(&repairRecord)
	}
}
