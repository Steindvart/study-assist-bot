package localization

import (
	"context"
	"log"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetText(ctx context.Context, textID string, templateData map[string]any) string {
	localizer, ok := FromContext(ctx)
	if !ok {
		return textID
	}

	config := &goi18n.LocalizeConfig{
		MessageID:    textID,
		TemplateData: templateData,
	}

	message, err := localizer.Localize(config)
	if err != nil {
		log.Println("Localization error: ", err)
		return textID
	}

	return message
}

func GetSimpleText(ctx context.Context, messageID string) string {
	return GetText(ctx, messageID, nil)
}

func GetPluralText(ctx context.Context, messageID string, count int, templateData map[string]any) string {
	localizer, ok := FromContext(ctx)
	if !ok {
		return messageID
	}

	if templateData == nil {
		templateData = make(map[string]any)
	}
	templateData["Count"] = count

	config := &goi18n.LocalizeConfig{
		MessageID:    messageID,
		PluralCount:  count,
		TemplateData: templateData,
	}

	message, err := localizer.Localize(config)
	if err != nil {
		return messageID
	}

	return message
}
