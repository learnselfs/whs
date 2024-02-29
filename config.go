// Package whs @Author Bing
// @Date 2024/2/27 10:44:00
// @Desc
package whs

import "github.com/learnselfs/wlog"

var (
	accessLog *wlog.Log
)

const ()

type ()

func init() {
	// logger
	accessLog = wlog.New()
	accessLog.SetLevel(wlog.PanicLevel)

	// accessLog

	access := make(wlog.Fields)
	access["type"] = "access"
	accessLog.WithFields(access)
}
