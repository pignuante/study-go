package scrapper

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/pignuante/test-crawler/crawler"
	"github.com/pignuante/test-crawler/utils"
)

// Scrape run scrapping
func Scrape(term string) {
	var baseURI string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"

	totalPages := crawler.GetPages(baseURI)
	c := make(chan []crawler.ExtractedJob)
	var jobs []crawler.ExtractedJob

	// bar := progressbar.Default(int64(totalPages))
	bar := pb.StartNew(totalPages)
	for i := 0; i < totalPages; i++ {
		bar.Increment()
		go crawler.GetPage(baseURI, i, c)
		// jobs = append(jobs, ...)
	}
	bar.Finish()

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	WriteJobsToCsv(jobs)

}

// WriteJobsToCsv write jobs to csv
func WriteJobsToCsv(jobs []crawler.ExtractedJob) {
	file, err := os.Create("./jobs.csv")
	utils.CheckErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	utils.CheckErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.ID, job.Title, job.Location, job.Salary, job.Summary}
		jwErr := w.Write(jobSlice)
		utils.CheckErr(jwErr)
	}
	fmt.Println("Done")
}
