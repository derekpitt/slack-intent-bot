# Slack Bot + WitAI

A simple Slack bot for calling out to Wit.AI and then letting you act on the outcome.

## Install

`go get github.com/derekpitt/slack-intent-bot`

## Example

This example uses negroni, but feel free to use anything that can use http.Handler


    package main
    
    import (
    	"net/http"
    
    	"github.com/codegangsta/negroni"
    	"github.com/derekpitt/slack-intent-bot"
    	"github.com/phyber/negroni-gzip/gzip"
    )
    
    func main() {
    	bot := slackintent.NewBot("<Slack Outgoing Token>", "<Slack Incoming URL>", "<Wit Api Token>")
    
    	bot.HandleIntent("deploy", func(context slackintent.IntentContext) {
    		if context.WitOutcome().Confidence > 0.5 {
    			context.Reply("deploy intent!")
    		} else {
    			context.Reply("what did you say?")
    		}
    	})
    
    	mux := http.NewServeMux()
    	mux.Handle("/", bot) // bot is a http.Handler, stick it anywhere!
    
    	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), gzip.Gzip(gzip.DefaultCompression))
    	n.UseHandler(mux)
    	n.Run(":3001")
    }
