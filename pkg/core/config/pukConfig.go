package config

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"

const (
	PubK_SM = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT
Vyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==
-----END PUBLIC KEY-----`

	PubK_R1 = `-----BEGIN CERTIFICATE-----
MIIC+zCCAqGgAwIBAgIUARhAfFSyhzcx9q4LdiYKl2UHo1YwCgYIKoZIzj0EAwIw
TjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxDDAKBgNVBAoTA0JTTjEP
MA0GA1UECxMGY2xpZW50MQ4wDAYDVQQDEwVic25jYTAgFw0xOTA5MjYxMDI0MDBa
GA8yMDk5MDkwNTAyMDQwMFowgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhDaGFu
Z3NoYTEOMAwGA1UEChMFQ21QYXkxPTALBgNVBAsTBHVzZXIwEgYDVQQLEwtob25n
emFvbm9kZTAOBgNVBAsTB2JzbmJhc2UwCgYDVQQLEwNjb20xJTAjBgNVBAMMHG5v
ZGVAaG9uZ3phb25vZGUuYnNuYmFzZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAAQ/X2w5+pJoZXNCO81T4xMR+TxmFoYk6eG1kYML8HBPrUT6QflxtDXYsE9h
SzVAovq5DHww3vD8Xft/mxwsAXyuo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgeAMAwG
A1UdEwEB/wQCMAAwHQYDVR0OBBYEFDPVPRqPANJavkNOg/WhPkUkH6wqMB8GA1Ud
IwQYMBaAFJuwcYba1G07p1ySkpzyes8L79OPMCUGA1UdEQQeMByCGmNhLmhvbmd6
YW9ub2RlLmJzbmJhc2UuY29tMIGEBggqAwQFBgcIAQR4eyJhdHRycyI6eyJoZi5B
ZmZpbGlhdGlvbiI6Imhvbmd6YW9ub2RlLmJzbmJhc2UuY29tIiwiaGYuRW5yb2xs
bWVudElEIjoibm9kZUBob25nemFvbm9kZS5ic25iYXNlLmNvbSIsImhmLlR5cGUi
OiJ1c2VyIn19MAoGCCqGSM49BAMCA0gAMEUCIQD7FBAQJsgS0uhaL4mjJgILdFfY
RKXvNutyKz/MqJ54RQIgNS67sSUCbOZRx1rWDqYEcBF1zypEFik25fNgY3zk5gM=
-----END CERTIFICATE-----`

	Pubk_K1 = `-----BEGIN PUBLIC KEY-----
MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAETXyl2gs9ChpfEIpJqgNj6ob6GrNp5Zax
TRpgP7zigIEEwAUEQNL8cQoMS+yzPXTFrWlZubE4GFWFoi/Nxk2jdA==
-----END PUBLIC KEY-----`

	PubK_SM_test = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECwJ5ftuqndO9H3ks1hD8cB6IA9lx
/b0Z2hnFZ77rgRm9Q4lY1aqIhkM63Lh6X7uwPsoRC1xkS0PMp5x/jnRWcw==
-----END PUBLIC KEY-----`

	PubK_R1_test = `-----BEGIN CERTIFICATE-----
MIIC+zCCAqGgAwIBAgIUARhAfFSyhzcx9q4LdiYKl2UHo1YwCgYIKoZIzj0EAwIw
TjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxDDAKBgNVBAoTA0JTTjEP
MA0GA1UECxMGY2xpZW50MQ4wDAYDVQQDEwVic25jYTAgFw0xOTA5MjYxMDI0MDBa
GA8yMDk5MDkwNTAyMDQwMFowgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhDaGFu
Z3NoYTEOMAwGA1UEChMFQ21QYXkxPTALBgNVBAsTBHVzZXIwEgYDVQQLEwtob25n
emFvbm9kZTAOBgNVBAsTB2JzbmJhc2UwCgYDVQQLEwNjb20xJTAjBgNVBAMMHG5v
ZGVAaG9uZ3phb25vZGUuYnNuYmFzZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAAQ/X2w5+pJoZXNCO81T4xMR+TxmFoYk6eG1kYML8HBPrUT6QflxtDXYsE9h
SzVAovq5DHww3vD8Xft/mxwsAXyuo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgeAMAwG
A1UdEwEB/wQCMAAwHQYDVR0OBBYEFDPVPRqPANJavkNOg/WhPkUkH6wqMB8GA1Ud
IwQYMBaAFJuwcYba1G07p1ySkpzyes8L79OPMCUGA1UdEQQeMByCGmNhLmhvbmd6
YW9ub2RlLmJzbmJhc2UuY29tMIGEBggqAwQFBgcIAQR4eyJhdHRycyI6eyJoZi5B
ZmZpbGlhdGlvbiI6Imhvbmd6YW9ub2RlLmJzbmJhc2UuY29tIiwiaGYuRW5yb2xs
bWVudElEIjoibm9kZUBob25nemFvbm9kZS5ic25iYXNlLmNvbSIsImhmLlR5cGUi
OiJ1c2VyIn19MAoGCCqGSM49BAMCA0gAMEUCIQD7FBAQJsgS0uhaL4mjJgILdFfY
RKXvNutyKz/MqJ54RQIgNS67sSUCbOZRx1rWDqYEcBF1zypEFik25fNgY3zk5gM=
-----END CERTIFICATE-----`

	Pubk_K1_test = `-----BEGIN PUBLIC KEY-----
MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEh4WlY4pCv814i3WY5aRhtR3PoiIXOM1I
5xBGylyQTedo6DzJUdLfYZSZLs4py70D8FJtNICMVQCfezA7whHzUw==
-----END PUBLIC KEY-----`
)

var test = false

//Enable test net parameters
//set test gateway public key
func SetTest() {
	test = true
}

func getSmKey() string {
	if test {
		return PubK_SM_test
	} else {
		return PubK_SM
	}

}

func getR1Key() string {
	if test {
		return PubK_R1_test
	} else {
		return PubK_R1
	}

}

func getK1Key() string {
	if test {
		return Pubk_K1_test
	} else {
		return Pubk_K1
	}

}

func GetGatewayPublicKey(algorithmType enum.App_AlgorithmType) string {
	switch algorithmType {
	case enum.AppAlgorithmType_SM2:
		return getSmKey()
	case enum.AppAlgorithmType_R1:
		return getR1Key()
	case enum.AppAlgorithmType_K1:
		return getK1Key()
	}

	return ""
}
