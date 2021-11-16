package cron

import (
	"testing"

	"github.com/spf13/viper"
)

func _TestSyncFund(t *testing.T) {
	viper.SetDefault("app.chan_size", 500)
	SyncFund()
}
