// Package whs @Author Bing
// @Date 2024/2/27 10:44:00
// @Desc
package whs

import (
	"github.com/learnselfs/wlog"
)

var (
	accessLog *wlog.Log
	// pattern
	pattern []string
)

func init() {
	// logger
	accessLog = createLog(accessLog, wlog.InfoLevel, map[string]any{"type": "access"})
}

func createLog(log *wlog.Log, level wlog.Level, kv map[string]any) *wlog.Log {
	log = wlog.New()
	fields := make(wlog.Fields)
	for k, v := range kv {
		fields[k] = v
	}
	log.WithFields(fields)
	return log
}
