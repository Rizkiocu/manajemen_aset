package controller

import (
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/usecase"

	"github.com/gin-gonic/gin"
)

type ManageAssetController struct {
	manageAssetUC usecase.ManageAssetUsecase
	g             *gin.RouterGroup
}

// show assets handler
func (m ManageAssetController) ShowAllAssetHandler(c *gin.Context) {

	mAssets, err := m.manageAssetUC.ShowAllAsset()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Message": "Success", "Data": mAssets})
}

//create a new asset handler

func (m ManageAssetController) CreateNewAssetHandler(c *gin.Context) {

	var manageAssetReq dto.ManageAssetRequest
	if err := c.ShouldBindJSON(&manageAssetReq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format", "error": err.Error()})
		return
	}

	if err := m.manageAssetUC.CreateTransaction(manageAssetReq); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Success", "Data": manageAssetReq})
	return

}
func (m ManageAssetController) Route() {
	m.g.GET("/manage-assets/show-all", m.ShowAllAssetHandler)
	m.g.POST("/manage-assets/create-new", m.CreateNewAssetHandler)
}

func NewManageAssetController(maUC usecase.ManageAssetUsecase, g *gin.RouterGroup) *ManageAssetController {
	return &ManageAssetController{
		manageAssetUC: maUC,
		g:             g,
	}
}