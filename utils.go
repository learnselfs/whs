// Package whs @Author Bing
// @Date 2024/2/3 22:00:00
// @Desc
package whs

import "strings"

// parseUrl for parse url from client request
func parseUrl(url string) []string {
	result := strings.Split(url, "/")
	return rangeExcludeString(result, "")
}

func parseUrlExcludeSpecialSymbol(url string) string {
	results := strings.Split(url, "/")
	result := rangeExcludeString(results, "")
	return strings.Join(result, "")
}

// doublePointerExcludeString for double pointer exclude string
func doublePointerExcludeString(list []string, s string) []string {

	return list
}

func rangeExcludeString(list []string, s string) []string {
	var result []string
	for _, v := range list {
		if v != s {
			result = append(result, v)
		}
	}
	return result
}

type Message struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
