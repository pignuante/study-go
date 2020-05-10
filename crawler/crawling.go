package crawler

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pignuante/test-crawler/utils"
)

// ExtractedJob info struct
type ExtractedJob struct {
	ID       string
	Title    string
	Salary   string
	Location string
	Summary  string
}

// GetPage get Page info
func GetPage(baseURI string, page int, mainC chan<- []ExtractedJob) {
	var jobs []ExtractedJob
	c := make(chan ExtractedJob)
	pageURI := baseURI + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURI)
	utils.CheckErr(err)
	utils.CheckCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go ExtractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

// GetPages get pages info
func GetPages(baseURI string) (pages int) {
	pages = 0
	res, err := http.Get(baseURI)
	utils.CheckErr(err)
	utils.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	// fmt.Println("doc")
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// ExtractJob make job info struct
func ExtractJob(card *goquery.Selection, c chan<- ExtractedJob) {
	id, exist := card.Attr("data-jk")
	title := utils.CleanString(card.Find(".title>a").Text())
	location := utils.CleanString(card.Find(".sjcl").Text())
	salary := utils.CleanString(card.Find(".salaryText").Text())
	summary := utils.CleanString(card.Find(".summary").Text())
	if exist {
		job := ExtractedJob{
			ID:       id,
			Title:    title,
			Location: location,
			Salary:   salary,
			Summary:  summary}

		c <- job
	}
}
