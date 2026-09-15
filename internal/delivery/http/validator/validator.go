// Package validator gom các helper validate dùng chung cho delivery và usecase.
package validator

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	uuidRe = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	supiRe = regexp.MustCompile(`^imsi-[0-9]{5,15}$`)
)

// IsUUID kiểm tra định dạng NfInstanceId (TS 29.571: UUID version 4 string).
func IsUUID(s string) bool { return uuidRe.MatchString(s) }

// IsSupi kiểm tra pattern imsi-<5..15 digits>.
func IsSupi(s string) bool { return supiRe.MatchString(s) }

// IsAbsoluteHTTPURL yêu cầu URL tuyệt đối với scheme http/https và có host.
func IsAbsoluteHTTPURL(raw string) bool {
	if raw == "" {
		return false
	}
	u, err := url.Parse(raw)
	if err != nil || !u.IsAbs() || u.Host == "" {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// IsJSONPointer kiểm tra RFC 6901: rỗng hoặc chuỗi các segment bắt đầu bằng '/'.
func IsJSONPointer(p string) bool {
	if p == "" {
		return true
	}
	if !strings.HasPrefix(p, "/") {
		return false
	}
	for _, seg := range strings.Split(p[1:], "/") {
		for i := 0; i < len(seg); i++ {
			if seg[i] != '~' {
				continue
			}
			if i+1 >= len(seg) || (seg[i+1] != '0' && seg[i+1] != '1') {
				return false
			}
			i++
		}
	}
	return true
}

// EscapePointerSegment escape một segment theo RFC 6901 để ghép JSON Pointer.
func EscapePointerSegment(seg string) string {
	seg = strings.ReplaceAll(seg, "~", "~0")
	return strings.ReplaceAll(seg, "/", "~1")
}
