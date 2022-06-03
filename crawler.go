package main

import (
	"time"
	"context"
	"github.com/chromedp/chromedp"
)

func start(email, pw string) (result string){
	opts :=append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,		
		chromedp.Flag("headless", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	err := chromedp.Run(ctx,
						chromedp.Navigate(`https://partner.spacecloud.kr/auth/login`),
						chromedp.WaitVisible(`#email`),
						chromedp.WaitVisible(`#pw`),
						chromedp.SendKeys(`#email`, email, chromedp.ByID),
						chromedp.SendKeys(`#pw`, pw, chromedp.ByID),
						chromedp.Click(`.member_inner > form > fieldset > button`, chromedp.NodeVisible),
						chromedp.Sleep(1*time.Second),
						chromedp.Navigate(`https://partner.spacecloud.kr/reservation`),
						chromedp.Sleep(1*time.Second),
						chromedp.OuterHTML(`.list_box_wrap`, &result, chromedp.NodeVisible),
						chromedp.Sleep(1*time.Second),
					)

	if err != nil{
		panic(err)
	}

	return
}