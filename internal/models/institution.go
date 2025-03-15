package models

type InstitutionCategory string

const (
	MOSQUE InstitutionCategory = "mosque"
	SURAU  InstitutionCategory = "surau"
	OTHERS InstitutionCategory = "others"
)

type Institution struct {
	OldId            int32               `bson:"oldId,omitempty" json:"oldId,omitempty"`
	Name             string              `bson:"name,omitempty" json:"name,omitempty"`
	Category         InstitutionCategory `bson:"category,omitempty" json:"category,omitempty"`
	State            string              `bson:"state,omitempty" json:"state,omitempty"`
	City             string              `bson:"city,omitempty" json:"city,omitempty"`
	QRImage          string              `bson:"qrImage,omitempty" json:"qrImage,omitempty"`
	QRContent        string              `bson:"qrContent,omitempty" json:"qrContent,omitempty"`
	SupportedPayment []string            `bson:"supportedPayment,omitempty" json:"supportedPayment,omitempty"`
	Coords           []float64           `bson:"coords,omitempty" json:"coords,omitempty"`
}
