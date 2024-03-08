package main

import (
	pjstealthgo "pjstealth-go"
	"time"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	defer browser.Close()

	ctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"),
		Proxy: &playwright.Proxy{
			Server: "socks5://127.0.0.1:7890",
		},
	})
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	page, err := ctx.NewPage()
	if err != nil {
		panic(err)
	}
	err = pjstealthgo.StealthSync(page)
	if err != nil {
		panic(err)
	}
	defer page.Close()

	_, err = page.Goto("https://web.uutool.cn/", playwright.PageGotoOptions{
		Timeout:   playwright.Float(60000),
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		panic(err)
	}
	page.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
		Path:     playwright.String("1.png"),
	})
	time.Sleep(time.Second * 10)
}
