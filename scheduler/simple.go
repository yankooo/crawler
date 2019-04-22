/**
 *  @author: yanKoo
 *  @Date: 2019/4/22 22:56
 *  @Description:
 */
package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit( r engine.Request) {
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

