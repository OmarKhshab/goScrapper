package routes

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"example.com/rest-api/models"
	postgresdb "example.com/rest-api/postgres_db"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

var mutex *sync.Mutex = &sync.Mutex{}

func createArticle(context *gin.Context) {
	var articles models.ArticleDTO
	var data = []models.Article{}
	var result = []string{}
	err := context.ShouldBindJSON(&articles)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	c := colly.NewCollector(colly.Async(true))
	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// triggered when a CSS selector matches an element
	c.OnHTML(articles.TitleSelector, func(e *colly.HTMLElement) {
		mutex.Lock()
		defer mutex.Unlock()
		// printing all URLs associated with the <a> tag on the page
		fmt.Println("%v", e.Text)
		data = append(data, models.Article{Title: e.Text, TitleSelector: articles.TitleSelector, Url: e.Request.URL.String()})
		result = append(result, e.Text)
	})
	for _, url := range articles.URLs {
		if err := c.Visit(url); err != nil {
			fmt.Println("error in scraping")
		}
	}
	c.Wait()

	res := postgresdb.PostgresDb.CreateInBatches(&data, 10)
	if res.Error != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{"message": "could not crate event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "article added!", "article": result})
}

func GetAndUpdateArticles() {
	var data = []models.Article{}
	res := postgresdb.PostgresDb.Find(&data)
	if res.Error != nil {
		log.Fatal("bad data")
	}
	c := colly.NewCollector(colly.Async(true), colly.AllowURLRevisit())
	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	for _, page := range data {
		c.OnHTML(page.TitleSelector, func(e *colly.HTMLElement) {
			// printing all URLs associated with the <a> tag on the page
			fmt.Println("%v", e.Text)
			page.Title = e.Text
			res := postgresdb.PostgresDb.Save(&page)
			if res.Error != nil {
				log.Fatal("bad data")
			}
		})
		if err := c.Visit(page.Url); err != nil {
			fmt.Println("error in scraping")
		}
		fmt.Println(page.Url)
	}
	c.Wait()

}
func getArticles(context *gin.Context) {
	var data = []models.Article{}
	res := postgresdb.PostgresDb.Find(&data)
	if res.Error != nil {
		log.Fatal("bad data")
	}

	context.JSON(http.StatusCreated, gin.H{"message": "articles!", "articles": data})
}
