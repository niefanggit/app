package crawler

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func crawler() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not launch playwright: %v", err)
	}

	headless := false
	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless})
	if err != nil {
		log.Fatalf("Could not launch chromium: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}

	if _, err := page.Goto("https://www.3658801.com/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}
}
