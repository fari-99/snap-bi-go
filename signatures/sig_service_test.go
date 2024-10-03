package signatures

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var TestConfigSigServiceData = &Config{
	ClientID:     "962489e9-de5d-4eb7-92a4-b07d44d64bf4",
	ClientSecret: "xS3vNQQgJRemFF0SZfXkZOq3r7kQ9n5YJgK4Wg0tVCQ=",
	PrivateKey: "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tDQpNSUlFdkFJQkFEQU5CZ2txaGtpRzl3MEJBUUVG" +
		"QUFTQ0JLWXdnZ1NpQWdFQUFvSUJBUURVeUllMkluN1g3N1VVDQovYUtRVG1jOFFHYTl2a1pkYnVMQlZ0ekRkbE" +
		"FkNUhKcis1RlNHdFIvUHRJN1ZEelRScVpGTnRVTUJubG1xNTgzDQppV1dSb2NLc0lqWndvZ3FGRDhVRTlDZ0py" +
		"SG5SY1Rvb0U1VTBsQ21iY2NpUnNPQUxMRXVzbW5vMDFQN01ta2J0DQplVlk0cVZ2VEhtK0VpRXUrYWRaM0lWZm" +
		"N2Qm1hbUxFbnJ0cmh3OXFaT0NMcTM2dGJFWDcrcTJZZFEyd1RDRjlGDQpHZUI1cWZ0bVFwWm9zaTRrTk5ZTHBt" +
		"UEFyK2dwU2FPWlNXNUxYdWNqbGJXb1lTdExoQjdJa1pLZVhRRnFYeHlrDQo2QlY0M2NMOGNWenNIWVFiWEx3Wk" +
		"puVi81YVYzRWxpemdnMWxSb0FyV2VkR1htV3FlbEFDRk1nN204M1N0SDRRDQpFK092QVg5NUFnTUJBQUVDZ2Y5" +
		"SnhEdElLUzFRTHduWC9SWWVjUS9aRTNxS3pTd1h1QmxrYlhpYlNxajdzL2IyDQorSCtvdHlSeitKTnBmZ3FZTk" +
		"NYd1dOSURHVHQ5SElKVi9vdjNzTnlNVUZBUFBCM3BnMDRLS1pxM3Y5TjdpSVZnDQozMmtha1VMYllOYm5xeW9B" +
		"dFpxR2ZBZGh5WFdMK1BBZDQ3TEFjZllqZXFtdWlKRzVzWGRvVldiZ1gxV0JBa3lQDQowalpqb2RtOFFUY3N0aU" +
		"E0YytlelFmaXExbUpVMzlVV3B1WlJBSEJCMVgrdzQrZzlxbWhmR0tLM3NWdTBRS3RMDQpQNFBKY0tjbS8waXha" +
		"eXRWejVLckpKeWtHRUVQc2tvY0ZsSDBRbktGTzBxOVk5KzBWeDNOSGtCQjdCVC9zUHAzDQpNOHBZRGV2QXBtT1" +
		"B1dFM1S2QvR0xDTEI4OWtlQUM2MjR5emd4eTBDZ1lFQTdRLzEzcFRQMExJaUxlRGVNTFlPDQp0WTM4emZjZkFD" +
		"M2UwUzdNM1Fnc01ENDc1OThBUURFQ0huRjUwSXE2OTFZNng1eXFVNUdVK05zMnhCWjB3RlVZDQpuT2JmS2pJQl" +
		"VMQmlnYjRBZFFEWHVwdkpJbU9BOUkyWnA0dkF2Q3ZDODZNMkM2VEFPbE5iak5HVGFDWnNoZGg2DQpraE5XLzE5" +
		"NFd1bVUzSldvYkwzOGtLVUNnWUVBNWNnTlNaTlp2MG1Ua3l0amFCZHhlbVZsUS81UGltcUw0R2NjDQp1MndCdn" +
		"Q4Z29Bc1l4SGVzNFZJTTB6VUNiRzg1RE9PZTcxdHprYlpGS3JZcFk4a0wzTTJhaGVRd3VhYW9xZUkyDQpicGVG" +
		"OU5JVU43dEpUYWlNRWZzTm1JUHprQ21iZHliQkZnQldXd2RjemRiRXdQSjNLd0NpcEVqR0hsWnpJMkd6DQpMSj" +
		"d5QjBVQ2dZRUF0V2FDS1BsTWNKRHUwSmFDejFsd3RnTFprUzNwZmZTYnpRdjMzYWNVVVRJK0d2Y3N0UElMDQpn" +
		"Z29wUWFmMjI2OFRPWTJyVkZsUldvQThUeDR6NXJ6M0laRVU4YS9rQzc5OUVYUzMydEJ6UTZ0VDVNbXBjdmtyDQ" +
		"p1K2FYN0NXOGZ6ekQ3WEw3V1daUjV0YXo2bjJFaTRNVWY5VG12SEZzb255YzJaSjNtQVpsVHJVQ2dZQnYzUGxq" +
		"DQpCVUllMXRIcVJxZ3BSNmh0L3FqUzloNkE4eWZQZ2tOQmx1NHVudStDR29UZk9LZklOYXFhcytiUEpVSnYzT3" +
		"hGDQorSTZGdEdkdHlLclRzazk5R2ZjTkVhOEs0bkNmNjhtMkF6d3FtSTlSSm5Na1JGYVVkbFVERjlIZUwvTklp" +
		"V3hjDQpxcU5mdzNEUm9LbStpUXVzdXExekx4Nk1MR2YzSkp3V3IvYUJ4UUtCZ1FDSi9NTUVWUmVISjIyOTc3bj" +
		"ZRbDdoDQprdXlXZ3czSzhLVDJJZDRZVGEzOGx2Z25wVzVDMnZ2YVRncXRsczhOQy9KeVo0WXJPc3VoZWU1bHZE" +
		"WXduYWlCDQp5eDdNdUVYMkluK0xlY3gzMUR6RlZiamJpYmZFSW5hSnhmUmhGM3gwS0ZOSzlnWXNRY2wwODBSTi" +
		"svcGJlaWFFDQo4a29ZMDdwWHJwTjJUNnpBWXlmY0ZRPT0NCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=",

	PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0NCk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NB" +
		"UThBTUlJQkNnS0NBUUVBMU1pSHRpSisxKysxRlAyaWtFNW4NClBFQm12YjVHWFc3aXdWYmN3M1pRSGVSeWEvdV" +
		"JVaHJVZno3U08xUTgwMGFtUlRiVkRBWjVacXVmTjRsbGthSEMNCnJDSTJjS0lLaFEvRkJQUW9DYXg1MFhFNktC" +
		"T1ZOSlFwbTNISWtiRGdDeXhMckpwNk5OVCt6SnBHN1hsV09LbGINCjB4NXZoSWhMdm1uV2R5RlgzTHdabXBpeE" +
		"o2N2E0Y1BhbVRnaTZ0K3JXeEYrL3F0bUhVTnNFd2hmUlJuZ2VhbjcNClprS1dhTEl1SkRUV0M2Wmp3Sy9vS1Vt" +
		"am1VbHVTMTduSTVXMXFHRXJTNFFleUpHU25sMEJhbDhjcE9nVmVOM0MNCi9IRmM3QjJFRzF5OEdTWjFmK1dsZH" +
		"hKWXM0SU5aVWFBSzFublJsNWxxbnBRQWhUSU81dk4wclIrRUJQanJ3Ri8NCmVRSURBUUFCDQotLS0tLUVORCBQ" +
		"VUJMSUMgS0VZLS0tLS0=",
}

var TestConfigSigServiceData_PEM = &Config{
	ClientID:       "962489e9-de5d-4eb7-92a4-b07d44d64bf4",
	ClientSecret:   "xS3vNQQgJRemFF0SZfXkZOq3r7kQ9n5YJgK4Wg0tVCQ=",
	PrivateKeyPath: "./private_key_test.pem",
	PublicKeyPath:  "./public_key_test.pem",
}

var AccessTokenTest = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJhOGQ2YmVkNS05Mz" +
	"dkLTQzZTUtYTlkMi1hYWY0ODFlZjc2YTIiLCJjbGllbnRJZCI6IjZhZTk1N2M0LTI4NjMtNDcxMy1hY2N" +
	"lLWJhMTJkZTYzNmNmYyIsIm5iZiI6MTYxMTM4NjM4NywiZXhwIjoxNjExMzg3Mjg3LCJpYXQiOjE2MTEz" +
	"ODYzODd9.nUillb6567_zkM6Ys35OOG-YWGoo7Ik1odPJn1tR-ao"

var TimestampSigServiceTest, _ = time.Parse(TimestampFormat, "2020-01-01T00:00:00+07:00")

var SignatureServiceAsymmetricResult = "pixoDayhjY8R4Jo4mXWHSOIdu094IFD64AxZHzGGwRGJFggVoFfIy+FU" +
	"bzZIATa8FRliuiU6UZVsuLrpnttZARb71cuWUrJG+ktl/cpb46f2RUU8fL6f3I1nQH4M5T0mRr1vC4Bhp/1SCZyEjQB" +
	"JPAHTc/NzPUVvBfF4CT8GUvijJii1bVQ0nvl8IX/yTeNgWKiC24XLOZt4esFsYNKx2cQEZ8ViDrtc/CgBaPregGbmCO" +
	"zp9q7ioEp4PHP1XgZVeMU4gkfeXZa/H9IQYsTpBr5Hp10lWoO+v/+aeAAhvcuPHgiUi2O+s+VGaoYdWQ1EFImBhqN48" +
	"XQZypKkeZzzLQ=="

var SignatureServiceSymmetricResult = "aqFyJg5631pb1qegHgkgFUNWdaD7fY6MCBbfYbIdcbyTO6Y" +
	"NEYT+wMpSmn1bnGiEjPBkSPGBPb42JK6gIalc+Q=="

func TestBase_SignatureService_Symmetric(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigServiceData)

	inputData := SignatureServiceInput{
		HttpMethod:  http.MethodPost,
		Url:         "/api/v1/balance-inquiry",
		AccessToken: AccessTokenTest,
		RequestBody: nil,
		Timestamp:   TimestampSigServiceTest,
	}

	signature, err := sigBase.SignatureService(SignatureAlgSymmetric, inputData)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}

	if !assert.Equal(t, SignatureServiceSymmetricResult,
		signature.StdEncoding()) {
		t.Fatalf("signature is not the same")
		return
	}
}

func TestBase_SignatureService_Asymmetric(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigServiceData)

	inputData := SignatureServiceInput{
		HttpMethod:  http.MethodPost,
		Url:         "/api/v1/balance-inquiry",
		AccessToken: AccessTokenTest,
		RequestBody: nil,
		Timestamp:   TimestampSigServiceTest,
	}

	signature, err := sigBase.SignatureService(SignatureAlgAsymmetric, inputData)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}

	if !assert.Equal(t, SignatureServiceAsymmetricResult,
		signature.StdEncoding()) {
		t.Fatalf("signature is not the same")
		return
	}
}

func TestBase_SignatureService_Asymmetric_PEM(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigServiceData_PEM)

	inputData := SignatureServiceInput{
		HttpMethod:  http.MethodPost,
		Url:         "/api/v1/balance-inquiry",
		AccessToken: AccessTokenTest,
		RequestBody: nil,
		Timestamp:   TimestampSigServiceTest,
	}

	signature, err := sigBase.SignatureService(SignatureAlgAsymmetric, inputData)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}

	if !assert.Equal(t, SignatureServiceAsymmetricResult,
		signature.StdEncoding()) {
		t.Fatalf("signature is not the same")
		return
	}
}

func TestBase_SignatureService_Symmetric_Validation(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigAccessTokenData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/balance-inquiry", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TIMESTAMP", TimestampSigAccessTokenTest.Format(TimestampFormat))
	req.Header.Set("Authorization", "Bearer "+AccessTokenTest)
	req.Header.Set("X-EXTERNAL-ID", time.Now().String())
	req.Header.Set("X-SIGNATURE", SignatureServiceSymmetricResult)
	req.Header.Set("X-PARTNER-ID", TestConfigSigAccessTokenData.ClientID)

	err := sigBase.VerifySignatureService(SignatureAlgSymmetric, req)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}
}

func TestBase_SignatureService_Asymmetric_Validation(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigAccessTokenData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/balance-inquiry", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TIMESTAMP", TimestampSigAccessTokenTest.Format(TimestampFormat))
	req.Header.Set("Authorization", "Bearer "+AccessTokenTest)
	req.Header.Set("X-EXTERNAL-ID", time.Now().String())
	req.Header.Set("X-SIGNATURE", SignatureServiceAsymmetricResult)
	req.Header.Set("X-PARTNER-ID", TestConfigSigAccessTokenData.ClientID)

	err := sigBase.VerifySignatureService(SignatureAlgAsymmetric, req)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}
}

func TestBase_SignatureService_Asymmetric_Validation_PEM(t *testing.T) {
	sigBase := NewBase()
	sigBase.SetConfig(TestConfigSigAccessTokenData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/balance-inquiry", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TIMESTAMP", TimestampSigAccessTokenTest.Format(TimestampFormat))
	req.Header.Set("Authorization", "Bearer "+AccessTokenTest)
	req.Header.Set("X-EXTERNAL-ID", time.Now().String())
	req.Header.Set("X-SIGNATURE", SignatureServiceAsymmetricResult)
	req.Header.Set("X-PARTNER-ID", TestConfigSigAccessTokenData.ClientID)

	err := sigBase.VerifySignatureService(SignatureAlgAsymmetric, req)
	if err != nil {
		t.Fatalf("Error in SignatureAccessToken Symmetric: %s", err.Error())
		return
	}
}
