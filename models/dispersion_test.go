package models

import "testing"

func TestDispersionInput_GenerateFirma(t *testing.T) {
	type fields struct {
		InstitucionContraparte int
		Empresa                string
		FechaOperación         int
		FolioOrigen            string
		ClaveRastreo           string
		InstitucionOperante    int
		MontoPago              float64
		TipoPago               int
		TipoCuentaOrdenante    int
		NombreOrdenante        string
		CuentaOrdenante        string
		RFCCURPOrdenante       string
		TipoCuentaBeneficiario int
		NombreBeneficiario     string
		CuentaBeneficiario     string
		RFCCURPBeneficiario    string
		EmailBeneficiario      string
		ConceptoPago           string
		ReferenciaNumerica     int
	}
	type args struct {
		privateKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{{
		name: "Prueba Firma",
		fields: fields{
			ClaveRastreo:           "Levita001",
			ConceptoPago:           "Prueba deposito",
			CuentaBeneficiario:     "846180000400000001",
			CuentaOrdenante:        "846180000400000002",
			Empresa:                "Peubas deposito",
			InstitucionContraparte: 846,
			InstitucionOperante:    90646,
			MontoPago:              99.99,
			NombreBeneficiario:     "Jose Rodriguez",
			NombreOrdenante:        "Adrian Tejeda",
			ReferenciaNumerica:     1,
			RFCCURPBeneficiario:    "ND",
			RFCCURPOrdenante:       "TERA960601II7",
			TipoCuentaBeneficiario: 40,
			TipoCuentaOrdenante:    40,
			TipoPago:               1,
			FolioOrigen:            "SDCF41DC9C",
		},
		args: args{
			privateKey: `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAreq+O2qtXvtMeYe72l3haU/VEe7/YbTOx/QiXreBM9CqnZ10
gR8AUTf5EqE14O+FYPXsYPHsYHdOD5OxE+E2MVPd61SsYfxnx0hda8+tngANFIls
7MwV7KMMw7R204WDC/yOqqV3X6pGqF9A/j+WBhEnA6ubKDOvvtswnsI6UoXVNTd5
+9FVMxElI6xFQvsBX7M4btDKxh0pnHtqN2tDXhJoQXDBNV8pXsOzq6qwoEa/IXZW
2KB4iBE/vS3EuA0HTBwhT5TKV6VBngKNilVznaMiFT7T/fiv7yYRetq2oUeRtyoJ
eKSIACGoMWknmGPCvph+Drx8eqaax4S3f+sq3wIDAQABAoIBAQCYO8Bklg5H88SP
JKbkkAS9gCnL1t4oknHmSObE7e/kgSS3bpoKZC6M6WfWTNKyoDaqFtngSiNUlVn4
/JimBB7Bst2wHp/97UiSAd/0fQdGJFlOvrRf2ozeZCLKM+9HW7jIBzyCVvBVBuFu
hGKJqvMRQHZcGhxoZkXR1F+MuAsBK7lEUphKLr2Z/Qw9f6tF8p+uspjQ/+Azemq+
bacLApPhqa5Whg5l5I2oD07ybVCH1aGUt+GtYMle7LM7uWBKO9ItqNKVNU5NrZzj
X+VoMDXhlBr+W2+33QmJoQr/RCItcN16/V+rmDgc48TmiG9nWHTk6LM1gsQq0AFM
iDDwaTXxAoGBAP1yezNSp7/F0gzp6dcV6IM5JX1oLdh4vOTbTLJ54Eb81RUwmG0G
Nqblpdpu9xAdpyaaGu58K6X/0l2+pwm7WTRjBVZwOStHzbGgSw7WxA4KLjr/EFAV
X1lYsyCYD9S7Zi4oVLe2Gdau77p5IOg5FZMecW0tyrNndkKydMCKzCv1AoGBAK+r
MQg1Q8n09+E+iOwVtzCu8AQ2xEkVl1OqLjz2jvZmH6Fo6qgwG+aDKXk7ATWCrR65
xxHoFEsoKD0uOn3jJjTrfEUzTu6UnzBPION1AR8YVZTwAcXuV+kgcCdNGEruVjlQ
3Lci6DYuTqH/X70jND+dBY0to2yRU8Iw4mLlQqsDAoGAb2FdNSvQ3p1P9y7g/g35
tZlqmVcsNqKw2J5rcU/QOUIpXnuRsO/3Gpd3sKvtVZ7Cc1tsFHxrLjNjvqDnhYZY
6IXPtbHx3cxbYAFCmw3U8RqWQIURJTCPS5OIfP12j5WfjQ9aU7XqWwdcAQX0aTmU
+qd5T3K9TuRb+2mzl14n5fUCgYALh1Ofc0dwoJgH+z+nJsGv+zsbO7DkDvuwbPG9
5Hx7ZTZcNHN5+DZMiX0WTK4Gof3Uj2KmJTH3wLnrUfZavqaKQI4WSIi8lZpg5ECH
TO126fTr4lhrcfno/cz+d0vJ1xxnhIwHpM3SiHV2ojjZQ88xAZAUtNMN0/fIPccN
sHgQmwKBgQCEot2o+/vrAxJRxk7/IHwGZ8GUeU4frINSqbXCy5HtI9KKwIn08wXa
zXO+REMtJIlTOYeSsew3jBR3NbMiV124+09Q7jrY3U8SV777Uip/8ItVXh+F+KwN
yobeV8V5SqSPnfl9J3cz2TF+bwTdOW1J8g34vPK1DvPGd5xUkbI5QQ==
-----END RSA PRIVATE KEY-----
`,
		},
		want:    "",
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &DispersionInput{
				InstitucionContraparte: tt.fields.InstitucionContraparte,
				Empresa:                tt.fields.Empresa,
				FechaOperación:         tt.fields.FechaOperación,
				FolioOrigen:            tt.fields.FolioOrigen,
				ClaveRastreo:           tt.fields.ClaveRastreo,
				InstitucionOperante:    tt.fields.InstitucionOperante,
				MontoPago:              tt.fields.MontoPago,
				TipoPago:               tt.fields.TipoPago,
				TipoCuentaOrdenante:    tt.fields.TipoCuentaOrdenante,
				NombreOrdenante:        tt.fields.NombreOrdenante,
				CuentaOrdenante:        tt.fields.CuentaOrdenante,
				RFCCURPOrdenante:       tt.fields.RFCCURPOrdenante,
				TipoCuentaBeneficiario: tt.fields.TipoCuentaBeneficiario,
				NombreBeneficiario:     tt.fields.NombreBeneficiario,
				CuentaBeneficiario:     tt.fields.CuentaBeneficiario,
				RFCCURPBeneficiario:    tt.fields.RFCCURPBeneficiario,
				EmailBeneficiario:      tt.fields.EmailBeneficiario,
				ConceptoPago:           tt.fields.ConceptoPago,
				ReferenciaNumerica:     tt.fields.ReferenciaNumerica,
			}
			got, err := doc.GenerateFirma(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DispersionInput.GenerateFirma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DispersionInput.GenerateFirma() = %v, want %v", got, tt.want)
			}
		})
	}
}
