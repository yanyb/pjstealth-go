package pjstealth_go

import (
	"github.com/playwright-community/playwright-go"
)

func StealthSync(page playwright.Page) error {
	navigatorUserAgent, err := page.Evaluate("navigator.userAgent")
	if err != nil {
		return err
	}
	navigatorPlatform, err := page.Evaluate("navigator.platform")
	if err != nil {
		return err
	}
	sc := NewStealthConfig(navigatorUserAgent.(string), navigatorPlatform.(string), nil)
	for _, script := range sc.enabledScripts() {
		page.AddInitScript(playwright.Script{
			Content: playwright.String(script),
		})
	}
	return nil
}
