package alogsubscriber

import "fmt"

type MQ struct {
}

func (mq *MQ) Send(v ...interface{}) {

	for idx := range v {
		value := v[idx]

		switch value {
		case value.(string):
			fmt.Println("Fake log sending", value)
		}
	}

}
