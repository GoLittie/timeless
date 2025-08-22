package timeless

import (
	"github.com/golittie/timeless/pkg/dateformat"
)

type parseOptions struct {
	UTCOffset float32
	dateformat.DateFormat
	AllowNegatives bool
}

type ParseOption func(*parseOptions)

func WithTimezone(offset float32) ParseOption {
	return func(options *parseOptions) {
		options.UTCOffset = offset
	}
}

func WithDateFormat(dateFormat dateformat.DateFormat) ParseOption {
	return func(options *parseOptions) {
		options.DateFormat = dateFormat
	}
}

func WithoutNegatives() ParseOption {
	return func(options *parseOptions) {
		options.AllowNegatives = false
	}
}
