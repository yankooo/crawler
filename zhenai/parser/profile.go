/**
 *  @author: yanKoo
 *  @Date: 2019/4/21 20:44
 *  @Description:
 */
package parser

import (
	"bytes"
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var nameRe = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]+)</span>`)
var profileRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {

	pInfo := extractAllString(contents, profileRe)

	name := extractString(contents, nameRe)
	var info bytes.Buffer
	for _, v := range pInfo {
		info .WriteString(v + " | ")
	}
	profile := model.Profile{
		Name: name,
		Info: info.String(),
	}

	res := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return res

}

func extractAllString(contents []byte, re *regexp.Regexp) []string {
	matchs := re.FindAllSubmatch(contents, -1)

	res := make([]string, 0)
	if len(matchs) >= 2 {
		for _, v := range matchs {
			res = append(res, string(v[1]))
		}
		return res

	} else {
		return nil
	}

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}