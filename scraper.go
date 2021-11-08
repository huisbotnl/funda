package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
)

var roomTypes = map[string]string{
	"huis":               "house",
	"appartement":        "apartment",
	"bouwgrond":          "land",
	"parkeergelegenheid": "parking",
	"object":             "object",
}

type Room struct {
	gorm.Model
	Price         string
	PricePeriod   string
	PriceCurrency string
	Type          string
	City          string
	Districts     string
	Url           string
	ScraperName   string
}

var DB *gorm.DB = nil
var err error = nil

//var rooms = make([]Room, 0)

/**
* connect with data base with ..env file params
* just edit all data in ..env file
 */
func ConnectToDatabase() {
	if DB == nil {
		DB, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local&character_set_server=utf8mb4")
	}
	if err != nil {
		fmt.Println("Connect To Database Error:", err.Error())
		return
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if os.Getenv("APP_ENV") == "local" {
		DB.LogMode(debug)
	}
}

func grabWithMap() {
	url := "/koop/heel-nederland/sorteer-datum-af/"
	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})
	c.OnHTML(`div[class=search-content-output]`, func(e *colly.HTMLElement) {
		e.ForEach("li[class=top-position-item-container]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName: os.Getenv("SCRAPER_NAME"),
			}
			el.ForEachWithBreak("a", func(i int, el *colly.HTMLElement) bool {
				if el.Attr("class") == "top-position-object-link top-position-object is-backgroundcover" {
					href := el.Attr("href")
					room.Url = os.Getenv("BASE_SCRAPER_URL") + href
					DB.Where("url = ?", room.Url).First(&room)
					if room.ID != 0 {
						return false
					}
					hp := strings.Split(strings.Split(strings.Split(href, "?")[0], "/")[3], "-")
					room.Type = roomTypes[hp[0]]
					room.Districts += hp[2]
					for i := 3; i < len(hp); i++ {
						room.Districts += " " + hp[i]
					}
					el.ForEach("span[class=top-position-object-description]", func(i int, el *colly.HTMLElement) {
						s := strings.Split(el.Text, " ")
						index := 0
						for i, s2 := range s {
							if !strings.Contains(s2, ",") {
								room.City += s2 + ""
							} else {
								room.City += strings.Trim(s2, ",")
								index = i + 2
								break
							}
						}
						room.PriceCurrency = "EUR"
						room.Price = s[index]
						room.PricePeriod = s[index+1]
					})
					return false
				}
				return true
			})
			if room.ID == 0 {
				DB.Create(&room)
			}
		})
		e.ForEach("li[class=search-result]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName:   os.Getenv("SCRAPER_NAME"),
				PriceCurrency: "EUR",
			}
			el.ForEach("div[class=search-result__header-title-col]", func(i int, el *colly.HTMLElement) {
				el.ForEachWithBreak("a", func(i int, el *colly.HTMLElement) bool {
					if i == 0 {
						href := el.Attr("href")
						room.Url = os.Getenv("BASE_SCRAPER_URL") + href
						DB.Where("url = ?", room.Url).First(&room)
						if room.ID != 0 {
							return false
						}
						parts := strings.Split(strings.Split(href, "?")[0], "/")
						room.City = parts[2]
						hp := strings.Split(parts[3], "-")
						room.Type = roomTypes[hp[0]]
						room.Districts += hp[2]
						for i := 3; i < len(hp); i++ {
							room.Districts += " " + hp[i]
						}
						return false
					}
					return true
				})
			})
			if room.ID != 0 {
				goto endOfIteration
			}
			el.ForEach("span[class=search-result-price]", func(i int, el *colly.HTMLElement) {
				priceParts := strings.Split(el.Text, " ")
				room.Price = priceParts[1]
				room.PricePeriod = priceParts[2]
			})
			DB.Create(&room)
		endOfIteration:
		})
		e.ForEach("div[class=search-result-content-info--object]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName:   os.Getenv("SCRAPER_NAME"),
				PriceCurrency: "EUR",
			}
			el.ForEach("div[class=search-result__header-title-col]", func(i int, el *colly.HTMLElement) {
				el.ForEachWithBreak("a", func(i int, el *colly.HTMLElement) bool {
					if i == 0 {
						href := el.Attr("href")
						room.Url = os.Getenv("BASE_SCRAPER_URL") + href
						DB.Where("url = ?", room.Url).First(&room)
						if room.ID != 0 {
							return false
						}
						parts := strings.Split(strings.Split(href, "?")[0], "/")
						room.City = parts[2]
						hp := strings.Split(parts[3], "-")
						room.Type = roomTypes[hp[0]]
						room.Districts += hp[2]
						for i := 3; i < len(hp); i++ {
							room.Districts += " " + hp[i]
						}
						return false
					}
					return true
				})
			})
			if room.ID != 0 {
				goto endOfIteration
			}
			el.ForEach("span[class=search-result-price]", func(i int, el *colly.HTMLElement) {
				priceParts := strings.Split(el.Text, " ")
				room.Price = priceParts[1]
				room.PricePeriod = priceParts[2]
			})
			DB.Create(&room)
		endOfIteration:
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		(*r.Headers)["User-Agent"] = []string{"*"}
	})
	errr := c.Visit(os.Getenv("BASE_SCRAPER_URL") + url)
	if errr != nil {
		fmt.Println("errr", errr.Error())
	}
}

func jobs() {
	gocron.Every(2).Hours().From(gocron.NextTick()).Do(grabWithMap)
	gocron.Start()
}
