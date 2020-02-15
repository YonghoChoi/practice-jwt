package token

import "testing"

func TestMakeJWT(t *testing.T) {
	token, err := MakeJWT("yongho")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(token)
}
