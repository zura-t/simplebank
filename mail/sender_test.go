package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zura-t/simplebank/utils"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Test email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test message <a href="http://google.com></a></p>
	`

	to := []string{"zura.terloeva395@gmail.com"}
	attachFiles := []string{}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
