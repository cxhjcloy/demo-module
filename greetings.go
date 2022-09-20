package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	innerShow "github.com/cxhjcloy/demo-module/showinfos"
)

var greetingsData = []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
	"你好! %v 热烈欢迎!",
}

// package 初始化 自动调用
func init() {
	fmt.Println("greetings init....")
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() (rnData string) {
	rnData = greetingsData[rand.Intn(len(greetingsData))]
	return
}

// Hello returns a greeting for the named person.
func Hello(name string) (message string, err error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		message = ""
		err = errors.New("empty name")
		return
	}
	err = nil
	message = fmt.Sprintf(randomFormat(), name)
	fmt.Println(innerShow.Show())
	return
}

func Hellos(names []string) (retData map[string]string, err error) {
	if len(names) < 1 {
		return nil, errors.New("names empty")
	}
	retData = make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			//return nil, errors.New("names empty")
			continue
		}
		retData[name] = message
	}
	return
}
