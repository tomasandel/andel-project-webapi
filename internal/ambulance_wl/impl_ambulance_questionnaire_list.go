package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implAmbulanceQuestionnaireListAPI struct {
}

func NewAmbulanceQuestionnaireListApi() QuestionnaireAPI {
	return &implAmbulanceQuestionnaireListAPI{}
}

func (o implAmbulanceQuestionnaireListAPI) CreateQuestionnaireEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceQuestionnaireListAPI) DeleteQuestionnaireEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceQuestionnaireListAPI) GetQuestionnaireEntries(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceQuestionnaireListAPI) GetQuestionnaireEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceQuestionnaireListAPI) UpdateQuestionnaireEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
