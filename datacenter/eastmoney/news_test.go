package eastmoney

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStockNews(t *testing.T) {
	data, err := _em.GetStockNews(_ctx, "300059.SZ")
	require.Nil(t, err)
	require.Len(t, data, 3)
}
