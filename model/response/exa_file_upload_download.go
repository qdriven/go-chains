package response

import "go-chains/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
