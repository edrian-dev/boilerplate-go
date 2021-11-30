package models

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

const publicKey = `-----BEGIN CERTIFICATE-----
MIIDdzCCAl+gAwIBAgIEfp/OQDANBgkqhkiG9w0BAQsFADBsMRAwDgYDVQQGEwdV
bmtub3duMRAwDgYDVQQIEwdVbmtub3duMRAwDgYDVQQHEwdVbmtub3duMRAwDgYD
VQQKEwdVbmtub3duMRAwDgYDVQQLEwdVbmtub3duMRAwDgYDVQQDEwdVbmtub3du
MB4XDTIxMTExNzE4MjAxNloXDTIyMDIxNTE4MjAxNlowbDEQMA4GA1UEBhMHVW5r
bm93bjEQMA4GA1UECBMHVW5rbm93bjEQMA4GA1UEBxMHVW5rbm93bjEQMA4GA1UE
ChMHVW5rbm93bjEQMA4GA1UECxMHVW5rbm93bjEQMA4GA1UEAxMHVW5rbm93bjCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK3qvjtqrV77THmHu9pd4WlP
1RHu/2G0zsf0Il63gTPQqp2ddIEfAFE3+RKhNeDvhWD17GDx7GB3Tg+TsRPhNjFT
3etUrGH8Z8dIXWvPrZ4ADRSJbOzMFeyjDMO0dtOFgwv8jqqld1+qRqhfQP4/lgYR
JwOrmygzr77bMJ7COlKF1TU3efvRVTMRJSOsRUL7AV+zOG7QysYdKZx7ajdrQ14S
aEFwwTVfKV7Ds6uqsKBGvyF2VtigeIgRP70txLgNB0wcIU+UylelQZ4CjYpVc52j
IhU+0/34r+8mEXratqFHkbcqCXikiAAhqDFpJ5hjwr6Yfg68fHqmmseEt3/rKt8C
AwEAAaMhMB8wHQYDVR0OBBYEFE8mq5sexcuvPOpZxol9O3HipcgCMA0GCSqGSIb3
DQEBCwUAA4IBAQAF/UfSrJeE7ijG5M5kuM7Npbd8wqZXhRyBBU2WJzAhPF5Sq23i
JjVDheAB7iZaZyL7rWpWtnVGLd/vd4uQiqwYO4mvQ0r/qn+4h6F1q+wqEFfhn4li
0qmAH/Ozsxk5YVe7tlax2ev+g//xcXAbf1t6FQ8GFLlyphuuywRVXPSZgU25moE3
PvAOiMMnsVCDr6VGJGoz8P8PtkVg9ysWN6cfXvUtZJbMR2ymQyRr3j0kGZL1z8UI
0rdZa5SIPpmVSGR0NmKIu0ihFdeLq9iwEoSTNWG5ooEi+HAkcSYfGWQY2a3+PugC
6+IhulM6jS5nJ1S18rVKbz3Ilxrlk9y8Bcue
-----END CERTIFICATE-----`

type Dispersion struct {
	Base                   Base
	STPID                  int64   `json:"stpID"`
	Estado                 string  `json:"estado"`
	CausaDevolucion        string  `json:"causaDevolucion"`
	TSLiquidacion          string  `json:"tsLiquidacion"`
	InstitucionContraparte int     `json:"institucionContraparte"`
	Empresa                string  `json:"empresa"`
	FechaOperación         int     `json:"fechaOperacion"`
	FolioOrigen            string  `json:"folioOrigen"`
	ClaveRastreo           string  `json:"claveRastreo"`
	InstitucionOperante    int     `json:"institucionOperante"`
	MontoPago              float64 `json:"monto"`
	TipoPago               int     `json:"tipoPago"`
	TipoCuentaOrdenante    int     `json:"tipoCuentaOrdenante"`
	NombreOrdenante        string  `json:"nombreOrdenante"`
	CuentaOrdenante        string  `json:"cuentaOrdenante"`
	RFCCURPOrdenante       string  `json:"rfcCurpOrdenante"`
	TipoCuentaBeneficiario int     `json:"tipoCuentaBeneficiario"`
	NombreBeneficiario     string  `json:"nombreBeneficiario"`
	CuentaBeneficiario     string  `json:"cuentaBeneficiario"`
	RFCCURPBeneficiario    string  `json:"rfcCurpBeneficiario"`
	ConceptoPago           string  `json:"conceptoPago"`
	ReferenciaNumerica     int     `json:"referenciaNumerica"`
}

type DispersionInput struct {
	InstitucionContraparte int     `json:"institucionContraparte"`
	Empresa                string  `json:"empresa"`
	FechaOperación         int     `json:"fechaOperacion"`
	FolioOrigen            string  `json:"folioOrigen"`
	ClaveRastreo           string  `json:"claveRastreo"`
	InstitucionOperante    int     `json:"institucionOperante"`
	MontoPago              float64 `json:"monto"`
	TipoPago               int     `json:"tipoPago"`
	TipoCuentaOrdenante    int     `json:"tipoCuentaOrdenante"`
	NombreOrdenante        string  `json:"nombreOrdenante"`
	CuentaOrdenante        string  `json:"cuentaOrdenante"`
	RFCCURPOrdenante       string  `json:"rfcCurpOrdenante"`
	TipoCuentaBeneficiario int     `json:"tipoCuentaBeneficiario"`
	NombreBeneficiario     string  `json:"nombreBeneficiario"`
	CuentaBeneficiario     string  `json:"cuentaBeneficiario"`
	RFCCURPBeneficiario    string  `json:"rfcCurpBeneficiario"`
	EmailBeneficiario      string  `json:"emailBeneficiario"`
	ConceptoPago           string  `json:"conceptoPago"`
	ReferenciaNumerica     int     `json:"referenciaNumerica"`
}

func (doc *DispersionInput) GetFirma() string {
	return fmt.Sprintf("||%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v||",
		doc.InstitucionContraparte,
		doc.Empresa,
		doc.FechaOperación,
		doc.FolioOrigen,
		doc.ClaveRastreo,
		doc.InstitucionOperante,
		doc.MontoPago,
		doc.TipoPago,
		doc.TipoCuentaOrdenante,
		doc.NombreOrdenante,
		doc.CuentaOrdenante,
		doc.RFCCURPOrdenante,
		doc.TipoCuentaBeneficiario,
		doc.NombreBeneficiario,
		doc.CuentaBeneficiario,
		doc.RFCCURPBeneficiario,
		doc.EmailBeneficiario,
		"", // Tipo de cuenta del beneficiario 2
		"", // Nombre del beneficiario 2
		"", // Cuenta del beneficiario 2
		"", // RFC / Curp del beneficiario 2
		doc.ConceptoPago,
		"", // Concepto del pago 2
		"", // Clave del catálogo de usuario 1
		"", // Clave del catálogo de usuario 2
		"", // Clave del pago
		"", // Referencia de cobranza
		"", // Referencia numérica
		"", // Tipo de operación
		"", // Topología
		"", // Usuario
		"", // Medio de entrega
		"", // Prioridad
		"", // IVA
	)
}

func ParseCertificate(certificateBytes []byte, privateKeyBytes []byte) (tls.Certificate, error) {

	var cert tls.Certificate
	var err error

	cert, err = tls.X509KeyPair([]byte(certificateBytes), []byte(privateKeyBytes))
	if err != nil {
		return cert, err
	}

	if len(cert.Certificate) > 1 {
		return cert, errors.New("PEM file contains multiple certificates")
	}

	c, err := x509.ParseCertificate(cert.Certificate[0])
	if c != nil && err == nil {
		cert.Leaf = c
	}

	return cert, nil

}

func LoadCertficateAndKeyFromFile(path string) (*tls.Certificate, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cert tls.Certificate
	for {
		block, rest := pem.Decode(raw)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		} else {
			cert.PrivateKey, err = parsePrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("Failure reading private key from \"%s\": %s", path, err)
			}
		}
		raw = rest
	}

	if len(cert.Certificate) == 0 {
		return nil, fmt.Errorf("No certificate found in \"%s\"", path)
	} else if cert.PrivateKey == nil {
		return nil, fmt.Errorf("No private key found in \"%s\"", path)
	}

	return &cert, nil
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, fmt.Errorf("Found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, fmt.Errorf("Failed to parse private key")
}

func (doc *DispersionInput) GenerateFirma(privateKey string) (string, error) {
	cadenaOriginal := doc.GetFirma()
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(cadenaOriginal))
	if err != nil {
		return "", err
	}

	// certificate, err := ParseCertificate([]byte(privateKey), []byte("levita"))
	// if err != nil {
	// 	return "", err
	// }

	// certificate, err := LoadCertficateAndKeyFromFile("/Users/intela/go/src/levita/levita-stp/llavePrivadaDev.pem")
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Printf("private key: %v", certificate.PrivateKey)

	block, _ := pem.Decode([]byte(privateKey))

	// result, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	// if err != nil {
	// 	return "", err
	// }
	// key, isOk := result.(*rsa.PrivateKey)
	// if !isOk {
	// 	return "", fmt.Errorf("type: %T", result)
	// }

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	msgHashSum := msgHash.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString(signature)

	publicBlock, _ := pem.Decode([]byte(publicKey))
	c, err := x509.ParseCertificate(publicBlock.Bytes)
	if err == nil {
		fmt.Printf("type: %T", c.PublicKey)
		err = rsa.VerifyPSS(c.PublicKey.(*rsa.PublicKey), crypto.SHA256, msgHashSum, signature, nil)
		if err != nil {
			fmt.Println("could not verify signature: ", err)
		}
		// If we don't get any error from the `VerifyPSS` method, that means our
		// signature is valid
		fmt.Println("signature verified")
	}

	return encoded, nil
}
