/**
 *  @author: yanKoo
 *  @Date: 2019/4/21 20:03
 *  @Description:
 */
package engine

// 放进任务队列的任务
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}
