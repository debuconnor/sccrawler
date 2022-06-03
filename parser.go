package main

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func parseHtml(html string) (result [][]string){
	_html, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil{
		panic(err)
	}

	list := _html.Find(".list_box")

	list.Each(func(i int, sel *goquery.Selection){
		num := sel.Find(".reservation_num").Text()
		num = strings.ReplaceAll(num, "예약번호 ", "")

		user := sel.Find(".user").Text()
		user = removeBlanks(user)
		user = strings.ReplaceAll(user, "예약자명", "")

		tel := sel.Find(".tel").Text()
		tel = removeBlanks(tel)
		tel = strings.ReplaceAll(tel, "전화번호", "")
		tel = strings.ReplaceAll(tel, "-", "")

		place := sel.Find(".place").Text()
		place = removeBlanks(place)
		place = strings.ReplaceAll(place, "예약공간", "")
		place = strings.Split(place, "-")[0]

		date := sel.Find(".date").Text()
		date = removeBlanks(date)
		date = strings.ReplaceAll(date, "날짜/시간", "")

		result = append(result, []string{num, user, tel, place, date})
	})

	return
}

func removeBlanks(text string) (result string){
	result = strings.ReplaceAll(text, " ", "")
	return
}