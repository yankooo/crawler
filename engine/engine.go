/**
 *  @author: yanKoo
 *  @Date: 2019/4/21 20:14
 *  @Description:
 */
package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: fetching url %s with error : %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)
	return parseResult, nil
}