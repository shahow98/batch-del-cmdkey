package core

import "testing"

func TestGetAllCmdKeys(t *testing.T) {
	list, err := GetAllCmdKeys()
	if err != nil {
		t.Error(err)
	}
	for _, v := range list {
		t.Log(v)
	}
}

func TestDelCmdkeys(t *testing.T) {
	err := DelCmdkeys([]string{"LegacyGeneric:target=git:https://zehongke@bitbucket.org"})
	if err != nil {
		t.Error(err)
	}
}
