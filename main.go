/**
 *  @author: yanKoo
 *  @Date: 2019/4/21 17:10
 *  @Description:
 */
package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 500,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
