package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

const banner = `                                                                                                                                                                                                                                                           
                     00              0000000                                                                                                                                                            
                  000               00000000000                                                                                                                                                         
                000                000000000000                                                                                                                                                         
              0000                0000000000000                                                                                                                                                         
            00000                 000 0000000000                                                                                                                                                        
           00000                     0000000000                                                                                                                                                         
          000000          0000000  00000000000                                                                                                                                                          
         000000        000  000000000000000000                                                                                                                                                          
         000000       00  00    000000000000                000000      000000   000000  000000        00000 00000      000000 00000000000000  00000               000000       00000000000             
        0000000           00      000000000000               000000    000000   00000000  000000      00000  00000      000000 00000000000000  00000              00000000      000000000000000         
        0000000           0         00000000000000            000000  00000    0000000000  00000     000000  00000      000000        0000000  00000             000000000      000000   0000000        
        0000000                      000000000000000  00000    00000000000    00000 00000   00000    00000   00000      000000       000000    00000            00000 00000     000000     00000        
        00000000                       00000000000000000000      00000000     00000  00000  000000  00000    00000      000000     000000      00000           000000  00000    000000    000000        
        00000000                            000000000 0000        000000     00000    00000  00000 00000     00000      000000   0000000       00000           00000    00000   000000000000000         
        000000000                           00000000 000000       00000     0000000000000000  0000000000     000000     00000   000000         00000          0000000000000000  0000000000000           
        0000000000                           00000   0000000      00000    000000000000000000  00000000       000000000000000  00000000000000  0000000000000 00000000000000000  000000   000000         
         0000000000                            0000   00000000    00000   000000        00000  0000000          00000000000    00000000000000  0000000000000000000        00000 000000    000000        
          00000000000                           0000    00                                                                                                                                              
           000000000000                         000    00                                                                                                                                               
            0000000000000                      0  0 0000                                                                                                                                                
              000000000000000                00  000000                                                                                                                                                 
                0000000000000000000        0000000000                                                                                                                                                   
                  00000000000000000000000000000000                                                                                                                                                      
                     00000000000000000000000000                            000  000   00 000000    00000000 0000000   000000  000                                                                       
                         000000000000000000                                 00  0000 0000000         000   00    000000    000000                                                                       
                                                                            00000 00000    00000     000  000    000000    000000                                                                       
                                                                             0000  000   0000000     000   00000000  00000000 0000000                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             
`

func main() {
	var i int
	site1, _ := os.Create("site1.txt")
	site2, _ := os.Create("site2.txt")
	site3, _ := os.Create("site3.txt")
	defer site1.Close()
	defer site2.Close()
	defer site3.Close()
	description := "-description"
	date := "-date"

	fmt.Printf(banner)
	fmt.Println("-1 Display the first news site")
	fmt.Println("-2 Display the second news site")
	fmt.Println("-3 Display the third news site")
	fmt.Println("-date")
	fmt.Println("        Filters the date part")
	fmt.Println("-description")
	fmt.Println("        Filters the description part")
	fmt.Println("-4 For exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number and filter: ")
	text, _ := reader.ReadString('\n')

	re := regexp.MustCompile("[1-4]+")
	sayi := re.FindAllString(text, -1)

	if len(sayi) > 0 {
		secim := sayi[0]
		i, _ = strconv.Atoi(secim)
	}

	switch i {

	case 1:

		site, _ := http.Get("https://thehackernews.com/")
		if site.StatusCode != 200 {
			fmt.Println("Hata", site.StatusCode)
		}
		doc, _ := goquery.NewDocumentFromReader(site.Body)

		doc.Find(".clear.home-right").Each(func(i int, selection *goquery.Selection) {
			haber := selection.Find("h2").Text()
			tarih := selection.Find(".h-datetime").Text()
			desc := selection.Find(".home-desc").Text()

			if strings.Contains(text, description) && strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, description) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			} else {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			}
			veri := fmt.Sprintf("%d. HABER: %s\nAÇIKLAMA: %s\nTARİH: %s\n\n\n", i+1, haber, desc, tarih)
			site1.WriteString(veri)
		})

	case 2:
		req, _ := http.NewRequest("GET", "https://www.bleepingcomputer.com", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		client := &http.Client{}
		site, err := client.Do(req)

		if err != nil {
			fmt.Println("Bağlantı hatası:", err)
			return
		}
		defer site.Body.Close()

		if site.StatusCode != 200 {
			fmt.Println("Hata", site.StatusCode)
			return
		}

		doc, _ := goquery.NewDocumentFromReader(site.Body)

		doc.Find(".bc_latest_news_text").Each(func(i int, selection *goquery.Selection) {
			haber := selection.Find("h4").Text()
			tarih := selection.Find(".bc_news_date").Text()
			desc := selection.Find("p").Text()

			if strings.Contains(text, description) && strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, description) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			} else {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			}
			veri := fmt.Sprintf("%d. HABER: %s\nAÇIKLAMA: %s\nTARİH: %s\n\n\n", i+1, haber, desc, tarih)
			site2.WriteString(veri)
		})
	case 3:
		req, _ := http.NewRequest("GET", "https://www.theregister.com/security/", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		client := &http.Client{}
		site, err := client.Do(req)

		if err != nil {
			fmt.Println("Bağlantı hatası:", err)
			return
		}
		defer site.Body.Close()

		if site.StatusCode != 200 {
			fmt.Println("Hata", site.StatusCode)
			return
		}

		doc, _ := goquery.NewDocumentFromReader(site.Body)

		doc.Find("article").Each(func(i int, selection *goquery.Selection) {
			haber := selection.Find("h4").Text()
			tarih := selection.Find(".time_stamp").Text()
			desc := selection.Find(".standfirst").Text()

			if strings.Contains(text, description) && strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, date) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)
				fmt.Println("\n------------------")
			} else if strings.Contains(text, description) {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			} else {
				color.Red("%d. Haber", i+1)
				fmt.Println(haber)

				color.Blue("Açıklama")
				fmt.Println(desc)

				color.Magenta("Tarih : %s", tarih)
				fmt.Println("\n------------------")
			}
			veri := fmt.Sprintf("%d. HABER: %s\nAÇIKLAMA: %s\nTARİH: %s\n\n\n", i+1, haber, desc, tarih)
			site3.WriteString(veri)
		})
	case 4:
		return

	}

}
