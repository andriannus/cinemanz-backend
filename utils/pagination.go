package utils

import (
	"net/http"
	"strconv"

	"cinemanz/constants"
)

// GetSkipAndLimit return skip and limit for pagination
func GetSkipAndLimit(r *http.Request) (skip int64, limit int64) {
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)

	if err != nil {
		skip = constants.SkipPerPage
	}

	limit, err = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	if err != nil {
		limit = constants.LimitPerPage
	}

	return skip, limit
}
