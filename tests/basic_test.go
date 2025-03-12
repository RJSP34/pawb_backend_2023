package tests

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/service"
	"testing"
)

func TestAddition(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Addition test failed. Expected: %d, got: %d", expected, result)
	}
}

func Add(i int, i2 int) interface{} {
	return i2 + i
}

func TestSubtraction(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtraction test failed. Expected: %d, got: %d", expected, result)
	}
}

func Subtract(i int, i2 int) interface{} {
	return i - i2
}

func TestBase64(t *testing.T) {
	plaintext := "Hello, World!"

	// Encrypt the plaintext
	ciphertext := service.EncodeToBase64([]byte(plaintext))

	// Decrypt the ciphertext
	decryptedText, err := service.DecodeBase64(ciphertext)
	assert.NoError(t, err, "Failed to decrypt the ciphertext")

	// Verify that the decrypted text matches the original plaintext
	assert.Equal(t, plaintext, string(decryptedText), "Decrypted text does not match the original plaintext")
}

func TestAESCrypto(t *testing.T) {
	plaintext := "Hello, World!"

	// Encrypt the plaintext
	key := make([]byte, 32) // 256-bit key size for AES-256
	_, err := rand.Read(key)
	assert.NoError(t, err, "Failed to encrypt the plaintext")

	ciphertext, err := service.Encrypt([]byte(plaintext), key)
	assert.NoError(t, err, "Failed to encrypt the plaintext")

	// Decrypt the ciphertext
	decryptedText, err := service.Decrypt(ciphertext, key)
	assert.NoError(t, err, "Failed to decrypt the ciphertext")

	// Verify that the decrypted text matches the original plaintext
	assert.Equal(t, plaintext, string(decryptedText), "Decrypted text does not match the original plaintext")
}

func TestSetCheckPassword(t *testing.T) {
	password := "myPassword123"

	// Set password hash
	hash, err := service.SetPasswordHash(password)
	assert.NoError(t, err, "Error setting password hash")

	// Check password hash
	err = service.CheckPasswordHash(password, hash)
	assert.NoError(t, err, "Error checking password hash")
}

func TestIsValidEmail(t *testing.T) {
	tests := map[string]struct {
		email    string
		expected bool
	}{
		"valid email 1": {
			email:    "test@example.com",
			expected: true,
		},
		"valid email 2": {
			email:    "user.name@example.co.uk",
			expected: true,
		},
		"invalid email 1": {
			email:    "test@example",
			expected: false,
		},
		"invalid email 2": {
			email:    "user@.com",
			expected: false,
		},
		"invalid email 3": {
			email:    "@example.com",
			expected: false,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := service.IsValidEmail(test.email)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCreateToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "your_jwt_secret")
	require.NoError(t, err, "Failed to set environment variable")

	// Test input values
	id := uint64(123)
	email := "test@example.com"
	role := "admin"

	// Create token
	token, err := service.CreateToken(id, email, role)
	require.NoError(t, err, "Error creating token")
	require.NotEmpty(t, token, "Token should not be empty")
}

func TestValidateToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "your_jwt_secret")
	require.NoError(t, err, "Failed to set environment variable")

	// Test input values
	id := uint64(123)
	email := "test@example.com"
	role := "admin"

	// Create token
	token, err := service.CreateToken(id, email, role)
	require.NoError(t, err, "Error creating token")
	require.NotEmpty(t, token, "Token should not be empty")

	// Validate token
	parsedToken, err := service.ValidateToken(token)
	require.NoError(t, err, "Error validating token")
	require.NotNil(t, parsedToken, "Parsed token should not be nil")
}

func TestConvertDTOToEntities(t *testing.T) {
	userID := uint64(1)

	patientClinicianDtos := []dto.PatientClinicianDTO{
		{
			ClinicianID: 3,
		},
		{
			ClinicianID: 5,
		},
	}

	expectedEntities := []entity.PatientClinician{
		{
			PatientID:   userID,
			ClinicianID: 3,
		},
		{
			PatientID:   userID,
			ClinicianID: 5,
		},
	}

	result := service.ConvertDTOToEntities(userID, patientClinicianDtos)

	assert.Equal(t, len(expectedEntities), len(result), "Number of entities does not match the expected result")

	for i := range expectedEntities {
		assert.Equal(t, expectedEntities[i].PatientID, result[i].PatientID, "PatientID does not match the expected result")
		assert.Equal(t, expectedEntities[i].ClinicianID, result[i].ClinicianID, "ClinicianID does not match the expected result")
	}
}

func TestGetItemsToInsert(t *testing.T) {
	currentList := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 2,
			Clinician:   entity.User{ID: 2},
		},
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 4,
			Clinician:   entity.User{ID: 4},
		},
	}

	requestedList := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 2,
			Clinician:   entity.User{ID: 2},
		},
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 6,
			Clinician:   entity.User{ID: 6},
		},
	}

	expected := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 6,
			Clinician:   entity.User{ID: 6},
		},
	}

	result := service.GetItemsToInsert(requestedList, currentList)

	assert.Equal(t, len(expected), len(result), "Number of items to insert is not equal to the expected result")

	for i := range expected {
		assert.Equal(t, expected[i].Patient.ID, result[i].Patient.ID, "Patient ID does not match the expected result")
		assert.Equal(t, expected[i].Clinician.ID, result[i].Clinician.ID, "Clinician ID does not match the expected result")
	}
}

func TestGetItemsToDelete(t *testing.T) {
	currentList := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 2,
			Clinician:   entity.User{ID: 2},
		},
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 4,
			Clinician:   entity.User{ID: 4},
		},
	}

	requestedList := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 2,
			Clinician:   entity.User{ID: 2},
		},
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 6,
			Clinician:   entity.User{ID: 6},
		},
	}

	expected := []entity.PatientClinician{
		{
			PatientID:   1,
			Patient:     entity.User{ID: 1},
			ClinicianID: 4,
			Clinician:   entity.User{ID: 4},
		},
	}

	result := service.GetItemsToDelete(requestedList, currentList)

	assert.Equal(t, len(expected), len(result), "Number of items to insert is not equal to the expected result")

	for i := range expected {
		assert.Equal(t, expected[i].Patient.ID, result[i].Patient.ID, "Patient ID does not match the expected result")
		assert.Equal(t, expected[i].Clinician.ID, result[i].Clinician.ID, "Clinician ID does not match the expected result")
	}
}
