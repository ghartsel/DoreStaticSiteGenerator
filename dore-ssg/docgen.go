package main

import (
    "os"
	"os/exec"
	"fmt"
    "io"
    "io/ioutil"
	"strings"
	"strconv"	
//    "bytes"
    "time"
	"github.com/BurntSushi/toml"
    "golang.org/x/net/html"
    "encoding/json"
)

type docConfig struct {
	Meta   		metaInfo
	Params		paramInfo
	Authors		authorInfo
	Reviews 	reviewInfo
	Publishes 	publishInfo
	Options 	optionInfo
	Badges 		badgeInfo
	Domains		domainInfo
}

type metaInfo struct {
	Title 		string
	Subtitle 	string
	Abstract 	string
	Keywords	[]string
	Status 		string 			// Draft-Review-Published
	DocVersion 	string
	Id			string
	Author		string
	Org 		string
	ProdName	string
	ProdVersion string
}

type paramInfo struct {
	Validate 	string
	Continue 	string 			// true = continue on validation error
}

type authorInfo struct {
	Markup 		string 				// rst | md
}

type reviewInfo struct {
	Location 	string
	Reviewers 	[]reviewerInfo
}

type publishInfo struct {
	Stylesheet 	string
}

type reviewerInfo struct {
	Name 	string
	Email 	string
}

type optionInfo struct {
	Header 		string
	Footer 		string
	Search 		string
}

type badgeInfo struct {
	Properties 	[]badgePropInfo
}

type badgePropInfo struct {
	Text 		string 
	Url 	 	string
}

type domainInfo struct {
	Classes 		[]classInfo
}

type classInfo struct {
	Name 		string
	Description	string
	Icon		string
	Topics 		[]topicInfo
}

type contentInfo struct {
	Topics 		[]topicInfo
}

type topicInfo struct {
	Filename 	string
	Author		string
	Status		string
}

var topics map[string]bool
var config docConfig
var mainNav = ""
var docList []string	// useful for indexer and prev|next nav

func cleanPubDir() {
	file, err := os.Open("pub")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    list,_ := file.Readdirnames(0)
    for _, name := range list {
    	if (strings.HasSuffix(name, "html")) {
    		os.Remove("pub/" + name)
    	}
    }
}

func copyImg(src, dst string) error {

    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }
    return out.Close()
}

func getConfig() error {

	if _, err := toml.DecodeFile("docmap.toml", &config); err != nil {
		fmt.Println(err)
		return err
	}

/*
	fmt.Printf("Title: %s\n", config.Meta.Title)
	fmt.Printf("Subtitle: %s\n", config.Meta.Subtitle)
	fmt.Printf("Abstract: %s\n", config.Meta.Abstract)
	fmt.Printf("Keywords:")
	for _, word := range config.Meta.Keywords{
		fmt.Printf(" %s,", word)
	}
	fmt.Printf("\nStatus: %s\n", config.Meta.Status)
	fmt.Printf("Version: %s\n", config.Meta.Version)
	fmt.Printf("Id: %s\n", config.Meta.Id)
	fmt.Printf("Author: %s\n", config.Meta.Author)
	fmt.Printf("Org: %s\n", config.Meta.Org)
	fmt.Printf("ProdName: %s\n", config.Meta.ProdName)
	fmt.Printf("ProdVersion: %s\n", config.Meta.ProdVersion)

	fmt.Printf("Validate: %s\n", config.Params.Validate)
	fmt.Printf("Continue: %s\n", config.Params.Continue)

	fmt.Printf("LandingPageOpt: %s\n", config.Options.Landingpage)
	fmt.Printf("HeaderOpt: %s\n", config.Options.Header)
	fmt.Printf("FooterOpt: %s\n", config.Options.Footer)
	fmt.Printf("SearchOpt: %s\n", config.Options.Search)

	fmt.Printf("Markup: %s\n", config.Contents.Markup)
	fmt.Printf("Stylesheet: %s\n", config.Contents.Stylesheet)
	for x, class := range config.Domains.Classes {
		fmt.Printf("\tTopic Class: %d - %s\n", x, class.Name)
		fmt.Printf("\t\tClass Description: %s\n", class.Description)
		for x, topic := range class.Topics {
			fmt.Printf("\tTopic: %d - %s\n", x, topic.Filename)
			fmt.Printf("\t\tAuthor: %s\n", topic.Author)
			fmt.Printf("\t\tStatus: %s\n", topic.Status)
		}
	}

	fmt.Printf("Review Location: %s\n", config.Reviews.Location)
	for x, reviewer := range config.Reviews.Reviewers {
		fmt.Printf("\tReviewer: %d %s\n", x, reviewer.Name)
		fmt.Printf("\tEmail: %d %s\n", x, reviewer.Email)
	}
*/

	return nil
}

func publishBaseHTML() error {

	// only transform topics changed since last build

	lastBuildTime := (int64)(0)
	file, err1 := os.Stat("pub/index.html")
	if err1 == nil {
		lastBuildTime = file.ModTime().Unix()
	}

	xform := config.Authors.Markup
    css := "--stylesheet=static/css/" + config.Publishes.Stylesheet

	for _, class := range config.Domains.Classes {
		for _, topic := range class.Topics {

			topics[topic.Filename] = false

			file, err1 = os.Stat("src/" + topic.Filename + ".rst")
		    if err1 != nil {
		        fmt.Println(err1)
		    }

		    modifiedtime := file.ModTime().Unix()

		    if (modifiedtime > lastBuildTime) {

		    	topics[topic.Filename] = true

			    _, err := exec.Command(xform, "--link-stylesheet", "--no-doc-title", css, "src/" + topic.Filename + ".rst", "pub/" + topic.Filename + ".html").Output()

			    if err != nil {
			        fmt.Printf("Error: %s %s\n", topic.Filename, err)
		        	return err
			    }
		    }

		    // update source file list for search indexer
		    docList = append (docList, topic.Filename)
		}
	}

	return nil
}

func searchResultsPageGen() {

	//
	// populate html except results content
	//

    newhtml := ""
    newhtml += string(Prefix)
    newhtml += string(PrefixNav)

    // inject nav sidebar and prev-next divs

    newhtml += mainNav
    newhtml += string(PreContent)
    newhtml += string(PreContentEnd)
    newhtml += string(PostfixNav) + "\n"
    newhtml += string(PostfixNavEnd) + "\n"
    newhtml += string(Postfix) + "\n"

    // insert badges

    newhtml += `<ul class="header__links">`
    for _, badge := range config.Badges.Properties {
        newhtml += `<li><a href="` + badge.Url + `">` + badge.Text + `</a></li>`
    }
    newhtml += `</ul>`

    newhtml += string(PostfixTerminal)

	//
    // write search results page to file
 	//

    srpage, err := os.Create("pub/searchResults.html")
    if err != nil {
        fmt.Println(err)
    }
    defer srpage.Close()

    if _, err := srpage.Write([]byte(newhtml)); err != nil {
        fmt.Println("Failed to create search results page", err)
    }
}

func landingPageGen() {

    lpContent := ""

    lpage, err := os.Create("pub/index.html")
    if err != nil {
        fmt.Println("File create error", err)
    }
    defer lpage.Close()

    lpContent += string(Prefix)
    lpContent += `<div class="landing"><h1 class="landing-title">` + config.Meta.Title + `</h1>`
    lpContent += `<div class="landing-version"> Version ` + config.Meta.DocVersion + `</div>`

    lpContent += string(landingPageSearch)

	lpContent += `<div class="abstract">`
	lpContent += `<p>` + config.Meta.Abstract + `</p></div>`

    lpContent += `<div class="tiles-flex">`

    for _, class := range config.Domains.Classes {
    	lpContent += `<a class="tile" href="` + class.Topics[0].Filename + `.html">`
//    	fmt.Println(class.Topics[0])
    	if class.Icon != "" {
	     	lpContent += `<img class="raw_img" src="static/img/` + class.Icon + `"/>`
	   	}
    	lpContent += `<h2 class="tile-title">` + class.Name + `</h2>`
    	lpContent += `<div class="tile-divider"></div><p class="tile-description">` + class.Description + `</p></a>`
    }

    lpContent += `</div>`
    lpContent += `<hr/>`
    if (config.Meta.Org != "") {
    	lpContent += `<div class="landing-copyright">`
    	time := time.Now()
    	year := time.Year()
    	lpContent += `Copyright ` + strconv.Itoa(year) + ` ` + config.Meta.Org
    	lpContent += `</div>`
    }
    lpContent += `</div>`
    
    lpContent += string(Postfix)

	lpContent += `<ul class="header__links">`
	for _, badge := range config.Badges.Properties {
		lpContent += `<li><a href="` + badge.Url + `">` + badge.Text + `</a></li>`
	}
	lpContent += `</ul>`
	lpContent += string(PostfixTerminal)
	lpContent += `</html>`

    if _, err := lpage.Write([]byte(lpContent)); err != nil {
        fmt.Println("Failed to create landing page", err)
    }
}

func makeNavEntry(prevlevel string, level string, heading *string, fname string) (string) {
	naventry := ""

	// normalize heading to create anchor

	anchor := strings.Replace(strings.ToLower(*heading), " ", "-", -1)
	anchor = strings.Replace(anchor, ".", "", -1)
	anchor = strings.Replace(anchor, ":", "", -1)
	anchor = strings.Replace(anchor, "_", "-", -1)

	// apply html according to heading level

	if level == "1"{
//		naventry = `<li class="toctree-l` + level + ` current"><a class="reference internal" href="` + fname + `.html">` + *heading + `</a>`
		naventry = `<li class="toctree-l` + level + `"><a class="reference internal" href="` + fname + `.html">` + *heading + `</a>`
	} else {
		if (level > prevlevel) {
			naventry = `<ul><li class="toctree-l` + level + `"><a class="reference internal" href="` + fname + `.html#` + anchor + `">` + *heading + `</a>`
		} else if (level < prevlevel) {
			naventry = `</li></ul></li><li class="toctree-l` + level + `"><a class="reference internal" href="` + fname + `.html#` + anchor + `">` + *heading + `</a>`
		} else {
			naventry = `</li><li class="toctree-l` + level + `"><a class="reference internal" href="` + fname + `.html#` + anchor + `">` + *heading + `</a>`
		}
	}

	return naventry
}

func navGen() {

	//
	// generate navigation div
	// config file is already parsed
	//

	heading :=  ""

	mainNav += string(navPrefix)

	for _, class := range config.Domains.Classes {

		// build category heading
		mainNav += `<p class="caption"><span class="caption-text">` + class.Name + `</span></p>`
		mainNav += `<ul class="current">`
//		mainNav += `<ul>`

		// make nav entries for category topics
		for _, topic := range class.Topics {

			// parse transformed HTML file
		    doc, err := os.Open("pub/" + topic.Filename + ".html")
		    if err != nil {
		        fmt.Println("File reading error", err)
		        return
		    }
		    defer doc.Close()

		    tokenizer := html.NewTokenizer(doc)
		    prevlevel := "0"

		    for {
		        //get next token type
		        tokenType := tokenizer.Next()
		//        fmt.Printf("%s\n", tokenType)
		//        fmt.Printf("%s\n", tokenizer.Token().Attr)

		        if tokenType == html.ErrorToken {
		            err := tokenizer.Err()
		            if err == io.EOF {
		                //end of the file, exit
		                break
		            }

		            fmt.Printf("error tokenizing HTML: %v", tokenizer.Err())
		        }

		        //process the token according to the token type...

		        if tokenType == html.StartTagToken {
		            //get the token and the next
		            token := tokenizer.Token()
	    			tokenType = tokenizer.Next()
		            heading = tokenizer.Token().Data

		            switch token.Data {
	    			case "h1", "h2", "h3", "h4":
	//	            	fmt.Println(token.Data)
						// have a heading, build the nav entry
	    				mainNav += makeNavEntry (prevlevel, token.Data[1:2], &heading, topic.Filename) + "\n"
	    				prevlevel = token.Data[1:2]
		            default:
		            	continue
		            }
		        }

		        // copy images from /src to /pub

		        if tokenType == html.SelfClosingTagToken {
		        	if itoken := tokenizer.Token(); itoken.Data == "img" {
		        		imgPath := itoken.Attr[0].Val
		        		if (len(itoken.Attr) > 3) {
		        			imgPath = itoken.Attr[2].Val
		        		}
						if err := copyImg("src/" + imgPath, "pub/" + imgPath); err != nil {
							fmt.Println(err.Error())
						}
		        	}
		        }

				if err := copyImg("src/img/bannerLogo.png", "pub/img/bannerLogo.png"); err != nil {
					fmt.Println(err.Error())
				}
		    }

		    // unwind heading levels

		    switch prevlevel {
		    case "4":
		    	mainNav += "</li></ul>"
		    	fallthrough
		    case "3":
		    	mainNav += "</li></ul>"
		    	fallthrough
		    case "2":
		    	mainNav += "</li></ul>"
		    }
		    mainNav += "</li>"
		}
		
		mainNav += "</ul>"
	}

	mainNav += string(navPostfix)
}

func main() {

	now := time.Now()
	startTime := now.UnixNano()

	topics = make(map[string]bool)

	//
	// parse TOML config file, validating settings
	// and xform rst to html to a well-known location
	//

	getConfig()
	publishBaseHTML() // check error rtn

	//
	// PASS 1:
	// 1. traverse doc map and generate nav div
	// 2. validate content: spellcheck, style, links, etc.
	//

	navGen()

	//
	// PASS 2:
	// augment HTML
	//

    newhtml := ""
	docListIndex := 0

	for _, class := range config.Domains.Classes {
		for _, topic := range class.Topics {

			//
			// inject theme elements into HTML
			//

			if topics[topic.Filename] == true {
				doc, err := ioutil.ReadFile("pub/" + topic.Filename + ".html")
				if err != nil {
				    fmt.Println(err)
				}

				file_content := string(doc)
				lines := strings.Split(file_content, "\n")
				contentBlock := false

				for _, y := range lines{
				    if len (y) > 1 {
				        if contentBlock && y != "</body>" {
				           	// report errors reported in transformed content
				           	if (y == `<div class="system-message">`) {
				           		fmt.Printf ("system-message: Check %s.rst for invalid markup\n", topic.Filename)
				        	}
				            newhtml += y + "\n"
				        }
				        if y == "<body>" {
				            // inject <head>
				            newhtml += string(Prefix)
				            newhtml += string(PrefixNav)

				            // generate search div

				            // inject nav sidebar and prev-next divs
				            newhtml += mainNav
				            newhtml += string(PreContent)
				            if (docListIndex < (len(docList) - 1)) {
				               	newhtml += `<a href="` + docList[docListIndex + 1] + `.html` + `" class="stealth-btn float-right" title="Accesskey Alt(+Shift)+n">Next</a><span class="bar  float-right">|</span>`
				            }
				            if docListIndex > 0 {
				               	newhtml += `<a href="` + docList[docListIndex - 1] + `.html` + `" class="stealth-btn float-right" title="Accesskey Alt(+Shift)+p">Previous</a>`
				            }
				            newhtml += string(PreContentEnd)

				            // change state
				            contentBlock = true
				        }
				        if y == "</body>" {
				            // inject <footer> and js

				            newhtml += string(PostfixNav) + "\n"
				            if docListIndex > 0 {
					            newhtml += `<a href="` + docList[docListIndex - 1] + `.html` + `" title="Accesskey Alt(+Shift)+p" accesskey="p">Previous</a><span class="bar">|</span>`
					        }
				            if (docListIndex < (len(docList) - 1)) {
					            newhtml += `<a href="` + docList[docListIndex + 1] + `.html` + `" title="Accesskey Alt(+Shift)+n" accesskey="n">Next</a>`
					        }
				            newhtml += string(PostfixNavEnd) + "\n"
				            newhtml += string(Postfix) + "\n"

							// insert badges
							newhtml += `<ul class="header__links">`
							for _, badge := range config.Badges.Properties {
								newhtml += `<li><a href="` + badge.Url + `">` + badge.Text + `</a></li>`
							}
							newhtml += `</ul>`
							newhtml += string(PostfixTerminal)
		                    break
		                }
			        }
			    }
			    err = ioutil.WriteFile("pub/" + topic.Filename + ".html", []byte(newhtml), 0777)
			    if err != nil {
			        fmt.Println(err)
			    }

			}
		    newhtml = ""
		    docListIndex++
		}
	}

	//
	// build landing page
	//

	landingPageGen()

	//
	// build search results page
	//

	searchResultsPageGen()

	//
	// publish search index
	//

	// create inverted index
	invertedIndex = GenerateInvertedIndex(docList)

	// encode index as json
    jsonData, err := json.Marshal(invertedIndex)
    if err != nil {
        fmt.Println(err)
    }

    // write index to file
    out, err := os.Create("pub/searchIndex.json")
    if err != nil {
        fmt.Println(err)
    }
    defer out.Close()

 	if _, err := out.Write(jsonData); err != nil {
        fmt.Println(err)
    }

	//
	// if in review status, publish doc to review site and generate review notice
	//

	if (config.Meta.Status == "Review") {
	    fmt.Printf("Generating review notice to reviewers\n")
	}

	now = time.Now()
	finishTime := now.UnixNano()
	fmt.Println(finishTime - startTime)
}
