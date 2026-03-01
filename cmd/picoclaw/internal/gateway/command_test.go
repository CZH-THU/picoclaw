package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGatewayCommand(t *testing.T) {
	cmd := NewGatewayCommand()

	require.NotNil(t, cmd)

	assert.Equal(t, "gateway", cmd.Use)
	assert.Equal(t, "Start picoclaw gateway", cmd.Short)

	assert.Len(t, cmd.Aliases, 1)
	assert.True(t, cmd.HasAlias("g"))

	assert.Nil(t, cmd.Run)
	assert.NotNil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	// Gateway command now accepts up to 1 argument (for "status" subcommand)
	assert.NotNil(t, cmd.Args)

	assert.True(t, cmd.HasFlags())
	assert.NotNil(t, cmd.Flags().Lookup("debug"))
}

func TestGatewayCommandStatusArgument(t *testing.T) {
	cmd := NewGatewayCommand()

	// Test that "status" argument is accepted and routes correctly
	err := cmd.RunE(cmd, []string{"status"})
	// StatusCmd() doesn't return an error, so this should succeed
	assert.NoError(t, err)
}

func TestGatewayCommandUnknownArgument(t *testing.T) {
	cmd := NewGatewayCommand()

	// Test that unknown arguments are rejected
	err := cmd.RunE(cmd, []string{"foo"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown argument: foo")
	assert.Contains(t, err.Error(), "did you mean 'picoclaw status'?")
}
