package otpgo

import (
	"testing"
)

func TestHOTP_Generate(t *testing.T) {
	t.Run("Normal Generation", testHOTPNormalGeneration)
	t.Run("Bad Key", testHOTPBadKey)
	t.Run("Default Params", testHOTPDefaultParams)
	t.Run("Autogenerated Key", testHOTPAutogeneratedKey)
	t.Run("Lower Case Key", testHOTPLowerCaseKey)
}

func testHOTPNormalGeneration(t *testing.T) {
	h := &HOTP{
		Key:       "73QK7D3A3PIZ6NUQQBF4BNFYQBRVUHUQ",
		Counter:   363,
		Algorithm: HmacSHA256,
		Length:    Length6,
	}

	otp, err := h.Generate()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if otp != "561655" {
		t.Errorf("wrong hotp\nexpected: %s\n  actual: %s", "561655", otp)
	}
}

func testHOTPBadKey(t *testing.T) {
	h := &HOTP{
		Key:       "invalid-base-32",
		Counter:   363,
		Algorithm: HmacSHA256,
		Length:    Length6,
	}

	_, err := h.Generate()
	expectedErr := ErrorInvalidKey{msg: "illegal base32 data at input byte 7"}
	if err != expectedErr {
		t.Errorf("unexpected error: %s", err)
	}
}

func testHOTPDefaultParams(t *testing.T) {
	h := &HOTP{Key: "73QK7D3A3PIZ6NUQQBF4BNFYQBRVUHUQ"}

	expectedAlg := HmacSHA1
	expectedLength := Length6
	expectedOtp := "769784"

	otp, err := h.Generate()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if h.Algorithm != expectedAlg {
		t.Errorf("unexpected hash algorithm\nexpected: %d (SHA1)\n  actual: %d", expectedAlg, h.Algorithm)
	}

	if h.Length != expectedLength {
		t.Errorf("unexpected length\nexpected: %d\n  actual: %d", expectedLength, h.Length)
	}

	if otp != expectedOtp {
		t.Errorf("unexpected hotp\nexpected: %s\n  actual: %s", expectedOtp, otp)
	}
}

func testHOTPAutogeneratedKey(t *testing.T) {
	h := &HOTP{}

	_, err := h.Generate()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if h.Key == "" {
		t.Error("expected Key to be populated")
	}
}

func testHOTPLowerCaseKey(t *testing.T) {
	h := &HOTP{Key: "73qk7d3a3piz6nuqqbf4bnfyqbrvuhuq"}

	_, err := h.Generate()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if h.Key == "" {
		t.Error("expected Key to be populated")
	}
}
