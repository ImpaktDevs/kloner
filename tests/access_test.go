package tests

import (
	"testing"

	"main/access"
	"main/mocker"

	"github.com/stretchr/testify/assert"
)

func TestConnectToServerWithPrivatePublicKeys(t *testing.T) {
	shutdown := mocker.StartMockServer()
	defer shutdown()
	user := "user"
	port := "2022"
	authType := "password-auth"
	pass := "pass"
	addr := "localhost"
	access.ConnectToServerWithPrivatePublicKeys(user, addr, port, authType, "", pass)
	assert.Nil(t, nil, "Test should have no errors")
}
