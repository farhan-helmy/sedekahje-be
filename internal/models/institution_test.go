package models

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestInstitutionValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		inst    Institution
		wantErr bool
	}{
		{
			name: "valid institution",
			inst: Institution{
				Name:             "Masjid Al-Test",
				Category:         MOSQUE,
				State:            "Selangor",
				City:             "Shah Alam",
				QRContent:        "test-qr-content",
				SupportedPayment: []string{"fpx", "card"},
				Coords:           []float64{3.0731, 101.5183},
			},
			wantErr: false,
		},
		{
			name: "missing required name",
			inst: Institution{
				Category:         MOSQUE,
				State:            "Selangor",
				City:             "Shah Alam",
				QRContent:        "test-qr-content",
				SupportedPayment: []string{"fpx"},
				Coords:           []float64{3.0731, 101.5183},
			},
			wantErr: true,
		},
		{
			name: "invalid category",
			inst: Institution{
				Name:             "Masjid Al-Test",
				Category:         "invalid",
				State:            "Selangor",
				City:             "Shah Alam",
				QRContent:        "test-qr-content",
				SupportedPayment: []string{"fpx"},
				Coords:           []float64{3.0731, 101.5183},
			},
			wantErr: true,
		},
		{
			name: "missing coordinates",
			inst: Institution{
				Name:             "Masjid Al-Test",
				Category:         MOSQUE,
				State:            "Selangor",
				City:             "Shah Alam",
				QRContent:        "test-qr-content",
				SupportedPayment: []string{"fpx"},
			},
			wantErr: true,
		},
		{
			name: "invalid category #2",
			inst: Institution{
				Name:             "Masjid Al-Test",
				Category:         "TEST",
				State:            "Selangor",
				City:             "Shah Alam",
				QRContent:        "test-qr-content",
				SupportedPayment: []string{"fpx"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.inst)
			if (err != nil) != tt.wantErr {
				t.Errorf("Institution validation error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
