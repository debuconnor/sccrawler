package sccrawler

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func ParseHtml(html string) (result []Reservation){
	_html, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil{
		panic(err)
	}

	list := _html.Find(".list_box")

	list.Each(func(i int, sel *goquery.Selection){
		num := sel.Find(".reservation_num").Text()
		num = strings.ReplaceAll(num, "예약번호 ", "")

		user := sel.Find(".user").Text()
		user = RemoveBlanks(user)
		user = strings.ReplaceAll(user, "예약자명", "")

		tel := sel.Find(".tel").Text()
		tel = RemoveBlanks(tel)
		tel = strings.ReplaceAll(tel, "전화번호", "")
		tel = strings.ReplaceAll(tel, "-", "")

		place := sel.Find(".place").Text()
		place = RemoveBlanks(place)
		place = strings.ReplaceAll(place, "예약공간", "")
		place = strings.Split(place, "-")[0]

		date := sel.Find(".date").Text()
		date = RemoveBlanks(date)
		date = strings.ReplaceAll(date, "날짜/시간", "")

		result = append(result, Reservation{num, user, tel, place, date})
	})

	return
}

func RemoveBlanks(text string) (result string){
	result = strings.ReplaceAll(text, " ", "")
	return
}