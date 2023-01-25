package sccrawler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

// Run crwaler to get SC reservations
func Start(email, pw, pageNum string, c chan<- string, channum int) {
	var result string

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://partner.spacecloud.kr/auth/login"),
		chromedp.WaitVisible("section.member_wrap"),
		chromedp.SendKeys("#email", email, chromedp.ByQuery),
		chromedp.SendKeys("#pw", pw, chromedp.ByQuery),
		chromedp.Click(".member_inner > form > fieldset > button", chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#all_menu", chromedp.ByQuery),
		chromedp.Navigate("https://partner.spacecloud.kr/reservation?RSV_STAT_CD=RSCMP&page="+pageNum),
		chromedp.WaitVisible(".list_box_wrap", chromedp.ByQuery),
		chromedp.OuterHTML(".list_box_wrap", &result, chromedp.NodeVisible),
	)

	checkError(err, "Error chan# "+strconv.Itoa(channum))

	c <- result
}

func ChromedpTest(email, pw string) {
	fmt.Println("*** chromedp Test ***")
	var result string

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://partner.spacecloud.kr/auth/login`),
		chromedp.WaitVisible(`#email`),
		chromedp.WaitVisible(`#pw`),
		chromedp.SendKeys(`#email`, email, chromedp.ByID),
		chromedp.SendKeys(`#pw`, pw, chromedp.ByID),
		chromedp.Click(`.member_inner > form > fieldset > button`, chromedp.NodeVisible),
		chromedp.WaitVisible("#all_menu"),
		chromedp.OuterHTML(`.text-info`, &result, chromedp.NodeVisible),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
