package slashing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgUnrevokeGetSignBytes(t *testing.T) {
	addr := sdk.Address("abcd")
	msg := NewMsgUnrevoke(addr)
	bytes := msg.GetSignBytes()
	assert.Equal(t, string(bytes), `{"address":"ichainvaladdr1v93xxeql7nzj5"}`)
}
