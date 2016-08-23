package platform

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kataras/iris"
)

type englishQuote struct {
	EnQuote  string
	EnAuthor string
	CnQuote  string
	CnAuthor string
}

var allQuotes []englishQuote
var quotesLen int

// Index ...
func Index(ctx *iris.Context) {
	quote := getRandQuote()
	fmt.Println(quote.EnQuote, quote.EnAuthor, quote.CnQuote, quote.CnAuthor)
	ctx.Render("index.html", quote)
}

// BandScore ...
func BandScore(ctx *iris.Context) {
	ctx.Render("band_score.html", nil)
}

// retrieve all english quotes from public/assets/quote/english_quotes.txt
// and convert those information into a global structure slice.
func init() {
	fh, err := os.Open("./public/assets/quote/english_quotes.txt")
	if err != nil {
		fmt.Println("platform init open english_quotes.txt error. ", err)
		return
	}
	defer fh.Close()

	r := bufio.NewReader(fh)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("platform init read line error. ", err)
			return
		}

		if err == io.EOF {
			break
		}

		lineArray := strings.Split(line, " && ")
		allQuotes = append(allQuotes, englishQuote{lineArray[1], lineArray[2], lineArray[3], lineArray[4]})
	}

	quotesLen = len(allQuotes)

	fmt.Println("platform package init successfully. ", quotesLen, " quotes have been loaded.")
}

func getRandQuote() englishQuote {
	i := rand.New(rand.NewSource(time.Now().UnixNano()))

	return allQuotes[i.Intn(quotesLen)-1]
}
