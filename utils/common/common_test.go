package common

import (
	"fmt"
	"testing"
	"time"
)

func TestStrftime(t *testing.T) {
	now := time.Now()
	strTime :=  Strftime(now)
	fmt.Println(strTime)

}

func TestStrptime(t *testing.T) {
	strTime := "2019-08-07 09:22:10"
	format := "2006-01-02 15:04:05"
	structTime, err := Strptime(strTime, format)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(structTime)

}