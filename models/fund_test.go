package models

import (
	"context"
	"smzdtz-server/datacenter/eastmoney"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func TestNewFund(t *testing.T) {
	ctx := context.TODO()
	efund, err := eastmoney.NewEastMoney().QueryFundInfo(ctx, "260104")
	require.Nil(t, err)
	fund := NewFund(ctx, efund)
	b, err := jsoniter.Marshal(fund)
	require.Nil(t, err)
	t.Log(string(b))
}
