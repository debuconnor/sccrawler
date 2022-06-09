package sccrawler

import (
	"time"
	"context"
	"github.com/chromedp/chromedp"
)

func Start(email, pw, pageNum string, c chan<- string) {
	var result string

	opts :=append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,		
		chromedp.Flag("headless", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second) // recommand time = page * 4 sec
	defer cancel()

	err := chromedp.Run(ctx,
						chromedp.Navigate(`https://partner.spacecloud.kr/auth/login`),
						chromedp.WaitVisible(`#email`),
						chromedp.WaitVisible(`#pw`),
						chromedp.SendKeys(`#email`, email, chromedp.ByID),
						chromedp.SendKeys(`#pw`, pw, chromedp.ByID),
						chromedp.Click(`.member_inner > form > fieldset > button`, chromedp.NodeVisible),
						chromedp.Sleep(1*time.Second),
						chromedp.Navigate(`https://partner.spacecloud.kr/reservation?page=`+pageNum),
						chromedp.Sleep(1*time.Second),
						chromedp.OuterHTML(`.list_box_wrap`, &result, chromedp.NodeVisible),
					)

	if err != nil{
		panic(err)
	}

	c <- result
}