package api

import (
	"net/http"
	"swift-parser/internal/models"

	"github.com/gin-gonic/gin"
)

// Response structures
type BranchResponse struct {
	Address       string `json:"address"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}

type HeadquarterResponse struct {
	Address       string           `json:"address"`
	BankName      string           `json:"bankName"`
	CountryISO2   string           `json:"countryISO2"`
	CountryName   string           `json:"countryName"`
	IsHeadquarter bool             `json:"isHeadquarter"`
	SwiftCode     string           `json:"swiftCode"`
	Branches      []BranchResponse `json:"branches"`
}

type CountryResponse struct {
	CountryISO2 string           `json:"countryISO2"`
	CountryName string           `json:"countryName"`
	SwiftCodes  []BranchResponse `json:"swiftCodes"`
}

func (r *Router) GetSWIFTCode(c *gin.Context) {
	swiftCode := c.Param("swiftCode")

	code, err := r.db.GetSWIFTCode(swiftCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	}

	if code.IsHeadquarter {
		branches, err := r.db.GetBranches(swiftCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get branches"})
			return
		}

		branchResponses := make([]BranchResponse, len(branches))
		for i, branch := range branches {
			branchResponses[i] = BranchResponse{
				Address:       branch.Address,
				BankName:      branch.BankName,
				CountryISO2:   branch.CountryISO2,
				IsHeadquarter: false,
				SwiftCode:     branch.SwiftCode,
			}
		}

		response := HeadquarterResponse{
			Address:       code.Address,
			BankName:      code.BankName,
			CountryISO2:   code.CountryISO2,
			CountryName:   code.CountryName,
			IsHeadquarter: true,
			SwiftCode:     code.SwiftCode,
			Branches:      branchResponses,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, code)
}

func (r *Router) GetSWIFTCodesByCountry(c *gin.Context) {
	countryCode := c.Param("countryISO2")

	codes, err := r.db.GetSWIFTCodesByCountry(countryCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}

	swiftCodes := make([]BranchResponse, len(codes))
	for i, code := range codes {
		swiftCodes[i] = BranchResponse{
			Address:       code.Address,
			BankName:      code.BankName,
			CountryISO2:   code.CountryISO2,
			IsHeadquarter: code.IsHeadquarter,
			SwiftCode:     code.SwiftCode,
		}
	}

	response := CountryResponse{
		CountryISO2: countryCode,
		CountryName: codes[0].CountryName,
		SwiftCodes:  swiftCodes,
	}
	c.JSON(http.StatusOK, response)
}

func (r *Router) PostSWIFTCode(c *gin.Context) {
	var newCode models.SwiftCode
	if err := c.ShouldBindJSON(&newCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := r.db.AddSWIFTCode(&newCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add SWIFT code"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "SWIFT code added successfully"})
}

func (r *Router) DeleteSWIFTCode(c *gin.Context) {
	swiftCode := c.Param("swiftCode")

	if err := r.db.DeleteSWIFTCode(swiftCode); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SWIFT code deleted successfully"})
}
