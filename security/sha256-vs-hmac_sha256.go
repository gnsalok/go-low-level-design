package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// sha256Digest returns the hex-encoded SHA-256 of msg (no authentication).
func sha256Digest(msg []byte) string {
	sum := sha256.Sum256(msg)
	return hex.EncodeToString(sum[:])
}

// hmacSHA256 computes HMAC-SHA256(key, msg) and returns hex string.
func hmacSHA256(key, msg []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return hex.EncodeToString(mac.Sum(nil))
}

// verifyHMAC returns true if providedMAC is a valid HMAC-SHA256 for (key,msg).
func verifyHMAC(key, msg []byte, providedMACHex string) bool {
	expectedMAC := hmac.New(sha256.New, key)
	expectedMAC.Write(msg)
	expected := expectedMAC.Sum(nil)

	provided, err := hex.DecodeString(providedMACHex)
	if err != nil {
		return false
	}
	// Constant-time comparison to avoid timing leaks.
	return hmac.Equal(provided, expected)
}

// helper to get a random key (32 bytes ~ 256-bit).
func randomKey(n int) []byte {
	k := make([]byte, n)
	if _, err := rand.Read(k); err != nil {
		panic(err)
	}
	return k
}

func main() {
	message := []byte("Hello")
	// In production, store and distribute this key securely.
	key := randomKey(32)

	// Plain SHA-256 (no key, not authenticated)
	sha := sha256Digest(message)
	fmt.Println("SHA-256(message):", sha)

	// HMAC-SHA256 (keyed, authenticated)
	tag := hmacSHA256(key, message)
	fmt.Println("HMAC-SHA256(key, message):", tag)

	// Verification succeeds with the same message & key
	ok := verifyHMAC(key, message, tag)
	fmt.Println("Verify(original):", ok)

	// Tamper with the message
	tampered := []byte("H3llo")
	ok = verifyHMAC(key, tampered, tag)
	fmt.Println("Verify(tampered):", ok)

	// Tamper with the tag (first hex nibble)
	if tag[0] != '0' {
		tagBad := "0" + tag[1:]
		fmt.Println("Verify(tag modified):", verifyHMAC(key, message, tagBad))
	}
}
