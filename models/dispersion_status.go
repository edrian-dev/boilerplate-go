package models

type DispersionStatus struct {
	ID              int64  `json:"id"`
	Empresa         string `json:"empresa"`
	FolioOrigen     string `json:"folioOrigen"`
	Estado          string `json:"estado"`
	CausaDevolucion string `json:"causaDevolucion"`
	TSLiquidacion   string `json:"tsLiquidacion"`
}
