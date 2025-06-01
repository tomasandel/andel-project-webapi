package ambulance_wl

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type implAmbulanceQuestionnaireListAPI struct {
}

func NewAmbulanceQuestionnaireListApi() QuestionnaireAPI {
	return &implAmbulanceQuestionnaireListAPI{}
}

func (o implAmbulanceQuestionnaireListAPI) CreateQuestionnaireEntry(c *gin.Context) {
	updateAmbulanceFunc(c, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		var entry Questionnaire

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if entry.PatientId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Patient ID is required",
			}, http.StatusBadRequest
		}

		if entry.Id == "" || entry.Id == "@new" {
			entry.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(ambulance.Questionnaires, func(existing Questionnaire) bool {
			return entry.Id == existing.Id || entry.PatientId == existing.PatientId
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Entry already exists",
			}, http.StatusConflict
		}

		ambulance.Questionnaires = append(ambulance.Questionnaires, entry)
		// ambulance.reconcileWaitingList()
		// entry was copied by value return reconciled value from the list
		entryIndx := slices.IndexFunc(ambulance.Questionnaires, func(existing Questionnaire) bool {
			return entry.Id == existing.Id
		})
		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save entry",
			}, http.StatusInternalServerError
		}
		return ambulance, ambulance.Questionnaires[entryIndx], http.StatusOK
	})
}

func (o implAmbulanceQuestionnaireListAPI) DeleteQuestionnaireEntry(c *gin.Context) {
	updateAmbulanceFunc(c, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.Questionnaires, func(existing Questionnaire) bool {
			return entryId == existing.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		ambulance.Questionnaires = append(ambulance.Questionnaires[:entryIndx], ambulance.Questionnaires[entryIndx+1:]...)
		return ambulance, nil, http.StatusNoContent
	})
}

func (o implAmbulanceQuestionnaireListAPI) GetQuestionnaireEntries(c *gin.Context) {
	updateAmbulanceFunc(c, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		result := ambulance.Questionnaires
		if result == nil {
			result = []Questionnaire{}
		}
		return nil, result, http.StatusOK
	})
}

func (o implAmbulanceQuestionnaireListAPI) GetQuestionnaireEntry(c *gin.Context) {
	updateAmbulanceFunc(c, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.Questionnaires, func(existing Questionnaire) bool {
			return entryId == existing.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		// return nil ambulance - no need to update it in db
		return nil, ambulance.Questionnaires[entryIndx], http.StatusOK
	})
}

func (o implAmbulanceQuestionnaireListAPI) UpdateQuestionnaireEntry(c *gin.Context) {
	updateAmbulanceFunc(c, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		var entry Questionnaire

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.Questionnaires, func(existing Questionnaire) bool {
			return entryId == existing.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		if entry.PatientId != "" {
			ambulance.Questionnaires[entryIndx].PatientId = entry.PatientId
		}

		if entry.Id != "" {
			ambulance.Questionnaires[entryIndx].Id = entry.Id
		}

		return ambulance, ambulance.Questionnaires[entryIndx], http.StatusOK
	})
}
