// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"github.com/hongsyang/req/v3/internal/ascii"
	"net/http"
	"sync"
)

var (
	commonBuildOnce   sync.Once
	commonLowerHeader map[string]string // Go-Canonical-Case -> lower-case
	commonCanonHeader map[string]string // lower-case -> Go-Canonical-Case
)

func buildCommonHeaderMapsOnce() {
	commonBuildOnce.Do(buildCommonHeaderMaps)
}

func buildCommonHeaderMaps() {
	common := []string{
		"accept",
		"accept-charset",
		"accept-encoding",
		"accept-language",
		"accept-ranges",
		"age",
		"access-control-allow-credentials",
		"access-control-allow-headers",
		"access-control-allow-methods",
		"access-control-allow-origin",
		"access-control-expose-headers",
		"access-control-max-age",
		"access-control-request-headers",
		"access-control-request-method",
		"allow",
		"authorization",
		"cache-control",
		"content-disposition",
		"content-encoding",
		"content-language",
		"content-length",
		"content-location",
		"content-range",
		"content-type",
		"cookie",
		"date",
		"etag",
		"expect",
		"expires",
		"from",
		"host",
		"if-match",
		"if-modified-since",
		"if-none-match",
		"if-unmodified-since",
		"last-modified",
		"link",
		"location",
		"max-forwards",
		"origin",
		"proxy-authenticate",
		"proxy-authorization",
		"range",
		"referer",
		"refresh",
		"retry-after",
		"server",
		"set-cookie",
		"strict-transport-security",
		"trailer",
		"transfer-encoding",
		"user-agent",
		"vary",
		"via",
		"www-authenticate",
		"x-forwarded-for",
		"x-forwarded-proto",
	}
	commonLowerHeader = make(map[string]string, len(common))
	commonCanonHeader = make(map[string]string, len(common))
	for _, v := range common {
		chk := http.CanonicalHeaderKey(v)
		commonLowerHeader[chk] = v
		commonCanonHeader[v] = chk
	}
}

func lowerHeader(v string) (lower string, isAscii bool) {
	buildCommonHeaderMapsOnce()
	if s, ok := commonLowerHeader[v]; ok {
		return s, true
	}
	return ascii.ToLower(v)
}

func canonicalHeader(v string) string {
	buildCommonHeaderMapsOnce()
	if s, ok := commonCanonHeader[v]; ok {
		return s
	}
	return http.CanonicalHeaderKey(v)
}
