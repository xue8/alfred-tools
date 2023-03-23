package main

// Package is called aw
import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/deanishe/awgo"
	"github.com/xue8/alfred-tools/pkg/base64"
	"github.com/xue8/alfred-tools/pkg/url"
)

// Workflow is the main API
var wf *aw.Workflow

// Your workflow starts here
func run() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		wf.NewItem("Invalid argument values")
		wf.SendFeedback()
		return
	}

	commandType := args[0]
	query := ""
	if len(args) > 1 {
		query = args[1]
	}

	log.Printf("command: %s, query: %s\n", commandType, query)
	switch commandType {
	case "sq":
		s := strconv.Quote(query)
		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "suq":
		s, err := strconv.Unquote(query)
		if err != nil {
			wf.NewItem(fmt.Sprintf("decode error: %v", err.Error()))
			break
		}

		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "urlencode", "ue":
		s := url.Encode(query)
		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "urldecode", "ud":
		s, err := url.Decode(query)
		if err != nil {
			wf.NewItem(fmt.Sprintf("decode error: %v", err.Error()))
			break
		}

		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "b6e":
		s := base64.Encode(query)
		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "b6d":
		s, err := base64.Decode(query)
		if err != nil {
			wf.NewItem(fmt.Sprintf("decode error: %v", err.Error()))
			break
		}

		wf.NewItem(s).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(s)
	case "dt":
		if len(query) == 0 {
			t := time.Now()
			wf.NewItem(t.String()).
				Subtitle("Press Enter to copy this value to clipboard").
				Valid(true).
				Arg(t.String())
			break
		}

		duration, err := time.ParseDuration(query)
		if err != nil {
			wf.NewItem(fmt.Sprintf("parseDuration error: %v", err.Error()))
			break
		}

		t := time.Now().Add(duration)
		wf.NewItem(t.String()).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(t.String())
	case "ts":
		if len(query) == 0 {
			t := time.Now()
			wf.NewItem(fmt.Sprint(t.Unix())).
				Subtitle("Press Enter to copy this value to clipboard").
				Valid(true).
				Arg(fmt.Sprint(t.Unix()))
			break
		}

		duration, err := time.ParseDuration(query)
		if err != nil {
			wf.NewItem(fmt.Sprintf("parseDuration error: %v", err.Error()))
			break
		}

		t := time.Now().Add(duration)
		wf.NewItem(fmt.Sprint(t.Unix())).
			Subtitle("Press Enter to copy this value to clipboard").
			Valid(true).
			Arg(fmt.Sprint(t.Unix()))
	default:
		wf.NewItem(fmt.Sprintf("command '%s' not support", commandType))
	}

	wf.SendFeedback()
}

func main() {
	wf = aw.New()
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
