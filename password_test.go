package passwordonion

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	const pepper = "XISUbiKpMwp6XeCeE2eJDpnP4e5wpT0g"
	const password = "O99xwLK1O2hTD9C8K4djxprTI9Yga1o6bjEjYIy8tos8Eg8IvUCCqxyd1XJIkFrAW"

	_, err := Encrypt(pepper, password)
	if err != nil {
		t.Errorf("Encrypt failed, %v", err)
	}
}

func TestCompare(t *testing.T) {
	const pepper = "n!wX5&b^UjG^G3j!Gsih6aSzGDUI$$@E"
	const password = "0GYp#dyR^WSTL#Dv1Rx6$dBtMgt05*j6l@Z4K50ksjjQ5J65YE@F!&!Fu^l*NhllYE@iVhYvM@t"

	p, err := Encrypt(pepper, password)
	if err != nil {
		t.Errorf("Encrypt failed, %v", err)
	}

	if err := Compare(pepper, string(p), password); err != nil {
		t.Errorf("Decrypt failed, %v", err)
	}
}
