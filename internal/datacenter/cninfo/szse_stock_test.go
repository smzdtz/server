package cninfo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStockList(t *testing.T) {
	results, err := _c.StockList(_ctx)
	require.Nil(t, err)
	t.Log(results)
}
