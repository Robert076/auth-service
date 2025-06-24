package hashing_service

import "testing"

func TestHashPassword(t *testing.T) {
	password := "secure123"
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	if err1 != nil || err2 != nil {
		t.Fatalf("HashPassword failed: %v, %v", err1, err2)
	}
	if hash1 == "" || hash2 == "" {
		t.Fatal("HashPassword returned empty hash")
	}
	if hash1 == hash2 {
		t.Error("Hashes for the same password should be different due to salting")
	}
}

func TestCompareHash(t *testing.T) {
	password := "secure123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if !CompareHash(hash, password) {
		t.Error("CompareHash should return true for correct password")
	}
	if CompareHash(hash, "wrongPassword") {
		t.Error("CompareHash should return false for incorrect password")
	}
}
