package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	NewFirmwareParams struct {
		ProductKey   string    `json:"product_key" binding:"required"`
		FirmwareName string    `json:"firmware_name" binding:"required"`
		Version      string    `json:"version" binding:"required"`
		Remark       string    `json:"remark" binding:"max=20"`
		FirmwareType int       `json:"firmware_type" binding:"required,min=0,max=1"`
		URL          string    `json:"url" binding:"required,max=256"`
		SignMethod   Algorithm `json:"sign_method" binding:"min=0,max=4"`
		Enable       bool      `json:"enable"`
		Tags         []string  `json:"tags"`
	}

	NewFirmwareVersionParams struct {
		FirmwareInfoID primitive.ObjectID `json:"firmware_info_id" binding:"required"`
		Version        string             `json:"version" binding:"required"`
		Remark         string             `json:"remark" binding:"max=20"`
		URL            string             `json:"url" binding:"required,max=256"`
		Enable         bool               `json:"enable"`
		SignMethod     Algorithm          `json:"sign_method" binding:"min=0,max=4"`
		Tags           []string           `json:"tags"`
	}

	Algorithm uint8
)
