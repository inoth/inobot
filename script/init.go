package script

import "fmt"

var ss *ScriptScheduler

func Init(poolSize int) {
	ss = &ScriptScheduler{
		SchedulerChan: make(chan string, poolSize),
	}
}

func CallScheduler() *ScriptScheduler {
	return ss
}

type ScriptScheduler struct {
	// 修改点，包含脚本，执行类型等信息
	SchedulerChan chan string
}

func (s *ScriptScheduler) SetQeust(args string) {
	s.SchedulerChan <- args
}

func (s *ScriptScheduler) Run() error {
	for {
		select {
		case m := <-s.SchedulerChan:
			fmt.Printf("脚本调度处理： %v\n", m)
		}
	}
}
