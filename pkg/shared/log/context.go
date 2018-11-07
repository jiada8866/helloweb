package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

var (
	// G is an alias for GetLogger.
	//
	// We may want to define this locally to a package to get package tagged log
	// messages.
	G = GetLogger

	// L is an alias for the the standard logger.
	// 直接使用 logrus.NewEntry 在 context 里没有 loggerKey 对应的 logger 时，
	// 都是直接返回的 L，绕过了 logrus 内置的 pool
	// L = logrus.NewEntry(logrus.StandardLogger())
	// 使用 WithFields 是想充分利用 logrus 里的 pool
	L = logrus.WithFields(logrus.Fields{})
)

type (
	loggerKey struct{}
)

// RFC3339NanoFixed is time.RFC3339Nano with nanoseconds padded using zeros to
// ensure the formatted time is always the same number of characters.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})

	if logger == nil {
		return L
	}

	return logger.(*logrus.Entry)
}
