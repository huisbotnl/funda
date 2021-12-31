package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	"os"
	"strings"
)

var rooms []Room

func setupDataMaps() {
	setupAmenitiesMap()
	setupBedroomsMap()
	setupCitiesMap()
	setupFacilitiesMap()
	setupFurnituresMap()
	setupNeighborhoodsMap()
	setupTypesMap()
}

func grabWithMap() {
	url := "/en/koop/heel-nederland/sorteer-datum-af/"
	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})
	c.OnHTML(`div[class=search-content-output]`, func(e *colly.HTMLElement) {
		e.ForEach("li[class=top-position-item-container]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName: os.Getenv("SCRAPER_NAME"),
			}
			room.SetFurniture("No")
			el.ForEachWithBreak("a", func(i int, el *colly.HTMLElement) bool {
				if el.Attr("class") == "top-position-object-link top-position-object is-backgroundcover" {
					href := el.Attr("href")
					room.Url = os.Getenv("BASE_SCRAPER_URL") + href
					DB.Where("url = ?", room.Url).First(&room)
					if room.ID != 0 {
						return false
					}
					parts := strings.Split(strings.Split(href, "?")[0], "/")
					room.SetCity(parts[2])
					hp := strings.Split(parts[len(parts)-2], "-")
					room.SetType(hp[0])
					room.Districts = hp[2]
					for i := 3; i < len(hp); i++ {
						room.Districts += " " + hp[i]
					}
					room.SetDistrict(room.Districts)
					el.ForEach("span[class=top-position-object-description]", func(i int, el *colly.HTMLElement) {
						s := strings.Split(el.Text, " ")
						index := 0
						for i, s2 := range s {
							if strings.Contains(s2, ",") {
								index = i + 2
								break
							}
						}
						room.PriceCurrency = "EUR"
						room.Price = s[index]
						//room.PricePeriod = s[index+1]
						room.PricePeriod = "month"
					})
					return false
				}
				return true
			})
			if room.ID == 0 {
				//rooms = append(rooms, room)
				room.Create()
			}
		})
		e.ForEach("li[class=search-result]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName:   os.Getenv("SCRAPER_NAME"),
				PriceCurrency: "EUR",
			}
			room.SetFurniture("No")
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
						room.SetCity(parts[2])
						hp := strings.Split(parts[len(parts)-2], "-")
						room.SetType(hp[0])
						room.Districts = hp[2]
						for i := 3; i < len(hp); i++ {
							room.Districts += " " + hp[i]
						}
						room.SetDistrict(room.Districts)
						return false
					}
					return true
				})
			})
			el.ForEach("li", func(i int, el *colly.HTMLElement) {
				if i == 2 {
					room.SetBedroom(el.Text)
				}
			})
			if room.ID != 0 {
				goto endOfIteration
			}
			el.ForEach("span[class=search-result-price]", func(i int, el *colly.HTMLElement) {
				priceParts := strings.Split(el.Text, " ")
				room.Price = priceParts[1]
				room.PricePeriod = priceParts[2]
			})
			//rooms = append(rooms, room)
			room.Create()
		endOfIteration:
		})
		e.ForEach("div[class=search-result-content-info--object]", func(i int, el *colly.HTMLElement) {
			room := Room{
				ScraperName:   os.Getenv("SCRAPER_NAME"),
				PriceCurrency: "EUR",
			}
			room.SetFurniture("No")
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
						room.SetCity(parts[2])
						hp := strings.Split(parts[len(parts)-2], "-")
						room.SetType(hp[0])
						room.Districts += hp[2]
						for i := 3; i < len(hp); i++ {
							room.Districts += " " + hp[i]
						}
						room.SetDistrict(room.Districts)
						return false
					}
					return true
				})
			})
			el.ForEach("li", func(i int, el *colly.HTMLElement) {
				if i == 1 {
					room.SetBedroom(el.Text)
				}
			})
			if room.ID != 0 {
				goto endOfIteration
			}
			el.ForEach("span[class=search-result-price]", func(i int, el *colly.HTMLElement) {
				priceParts := strings.Split(el.Text, " ")
				room.Price = priceParts[1]
				room.PricePeriod = priceParts[2]
			})
			//rooms = append(rooms, room)
			room.Create()
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
	for _, room := range rooms {
		cn := colly.NewCollector()
		cn.OnHTML(`div`, func(e *colly.HTMLElement) {
			room.SetDistrict(after(e.ChildText("h1[class=article__title]"), ", "))
			found := false
			e.ForEach("div[class=field]", func(i int, el *colly.HTMLElement) {
				if el.ChildAttr("span", "class") == "property__label label label--interior" {
					room.SetFurniture(strings.TrimSpace(after(el.Text, "Interior: ")))
					found = true
				}
				if el.ChildAttr("span", "class") == "property__label label label--bedrooms" {
					room.SetBedroom(strings.TrimSpace(after(el.Text, "Bedrooms: ")))
				}
				if el.ChildAttr("div", "class") == "label-above" {
					room.AddFacility(before(el.ChildText("div"), ":"))
				}
			})
			if !found {
				room.SetFurniture("No")
			}
			room.Create()
		})
		cn.OnResponse(func(response *colly.Response) {
			//fmt.Println(string(response.Body))
		})
		cn.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
			(*r.Headers)["User-Agent"] = []string{"*"}
		})
		cn.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
		errr := cn.Visit(room.Url)
		if errr != nil {
			fmt.Println("errr", errr.Error())
		}
	}
}

func jobs() {
	gocron.Every(2).Hours().From(gocron.NextTick()).Do(grabWithMap)
	gocron.Start()
}
