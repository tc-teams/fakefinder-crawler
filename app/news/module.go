package news

import (
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
	validate "github.com/idasilva/dtk-knowledge/app/news/valid"
	"github.com/idasilva/dtk-knowledge/collector"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//Limiting Colly to parse only links that are on the clienturl.com domain
//Turning on Async processing of links (this is where we get a HUGE speed increase as we'll talk about in a bit)

//HandlerFakeFinder instance a new collector of news
func HandlerFakeFinder(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)

	validation := validate.NewValidate("Validate")

	c := collector.NewColly(colly.NewCollector(
		colly.AllowedDomains(collector.Folha, collector.G1, collector.Uol),
		colly.Async(true),
		colly.AllowURLRevisit(),
	),
		&log.Logger{
			Out:       os.Stdout,
			Formatter: &log.JSONFormatter{},
			Level:     log.DebugLevel,
		}, validation.Valid, param["content"],
	)

	log.WithFields(log.Fields{"Text": param["content"]}).Warn("Search by content input")

	c.SearchAndInputNews()

	w.WriteHeader(http.StatusOK)

}