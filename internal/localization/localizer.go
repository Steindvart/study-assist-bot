package localization

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localizer interface {
	Localize(config *i18n.LocalizeConfig) (string, error)
}

type localizerKey struct{}

func WithLocalizer(ctx context.Context, localizer Localizer) context.Context {
	return context.WithValue(ctx, localizerKey{}, localizer)
}

func FromContext(ctx context.Context) (Localizer, bool) {
	localizer, ok := ctx.Value(localizerKey{}).(Localizer)
	return localizer, ok
}

func MustFromContext(ctx context.Context) Localizer {
	localizer, ok := FromContext(ctx)
	if !ok {
		panic("localizer not found in context")
	}
	return localizer
}
