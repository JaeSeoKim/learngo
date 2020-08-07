package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var mu sync.RWMutex

func main() {
	var jobs []extractedJob

	totalPages := getPages()
	pageChanel := make(chan []extractedJob)
	for i := 0; i < totalPages; i++ {
		go getpage(i, pageChanel)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-pageChanel
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
}
func writeJobs(jobs []extractedJob) {
	defer fmt.Println("Done! extracted all jobs!")
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	work := new(sync.WaitGroup)
	work.Add(len(jobs))
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		go writeCSV(w, jobSlice, work)
	}
	work.Wait()
	w.Flush()
}

func writeCSV(w *csv.Writer, jobSlice []string, work *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	jwErr := (*w).Write(jobSlice)
	checkErr(jwErr)
	work.Done()
}

func getpage(i int, c chan<- []extractedJob) {
	var jobs []extractedJob
	getURL := baseURL + "&start=" + strconv.Itoa(i*50)
	fmt.Println("get Pages... :", getURL)
	res, err := http.Get(getURL)
	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	c <- jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return (pages)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
