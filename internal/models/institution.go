package models

type InstitutionCategory string

const (
	MOSQUE InstitutionCategory = "mosque"
	SURAU  InstitutionCategory = "surau"
	OTHERS InstitutionCategory = "others"
)

type Institution struct {
	OldId            int32               `bson:"oldId,omitempty" json:"oldId,omitempty"`
	Name             string              `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Category         InstitutionCategory `bson:"category,omitempty" json:"category,omitempty" validate:"required,oneof=mosque surau others"`
	State            string              `bson:"state,omitempty" json:"state,omitempty" validate:"required"`
	City             string              `bson:"city,omitempty" json:"city,omitempty" validate:"required"`
	QRImage          string              `bson:"qrImage,omitempty" json:"qrImage,omitempty"`
	QRContent        string              `bson:"qrContent,omitempty" json:"qrContent,omitempty" validate:"required"`
	SupportedPayment []string            `bson:"supportedPayment,omitempty" json:"supportedPayment,omitempty" validate:"required"`
	Coords           []float64           `bson:"coords,omitempty" json:"coords,omitempty" validate:"required,min=2,max=2"`
	Slug             string              `bson:"slug,omitempty" json:"slug,omitempty" validate:"required"`
}
