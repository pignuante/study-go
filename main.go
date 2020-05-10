package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/pignuante/test-crawler/crawler"
	"github.com/pignuante/test-crawler/utils"
	"github.com/schollz/progressbar"
)

func main() {
	totalPages := getPages()
	c := make(chan []extractedJob)
	var jobs []extractedJob

	// bar := progressbar.Default(int64(totalPages))
	bar := pb.StartNew(totalPages)
	for i := 0; i < totalPages; i++ {
		bar.Increment()
		go crawler.getPage(i, c)
		// jobs = append(jobs, ...)
	}
	bar.Finish()

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobsToCsv(jobs)
}

func writeJobsToCsv(jobs []crawler.ExtractedJob) {
	file, err := os.Create("./jobs.csv")
	utils.CheckErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	utils.CheckErr(wErr)

	bar := progressbar.Default(int64(len(jobs)))
	for _, job := range jobs {
		bar.Add(1)
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		utils.CheckErr(jwErr)
	}
	fmt.Println("Done")
}
