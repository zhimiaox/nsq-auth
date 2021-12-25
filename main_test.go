package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
)

func TestPluginRootSecret_Authorization(t *testing.T) {
	get, err := http.Get("http://localhost:1325/auth?secret=123")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, get.StatusCode)
	body, _ := ioutil.ReadAll(get.Body)
	_ = get.Body.Close()
	t.Log(string(body))
}

func BenchmarkApi_Auth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get, err := http.Get("http://localhost:1325/auth?secret=123")
		if err != nil {
			b.Fatal(err)
		}
		assert.Equal(b, 200, get.StatusCode)
	}
}

const host = "localhost:4150"

type handle struct {
}

func (h handle) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}

func TestPush(t *testing.T) {
	conf := nsq.NewConfig()
	conf.AuthSecret = "123"
	producer, err := nsq.NewProducer(host, conf)
	if err != nil {
		return
	}
	for i := 0; i < 200000; i++ {
		err := producer.Publish("t1", []byte(fmt.Sprintf("bal-%d", i)))
		if err != nil {
			t.Error(err.Error())
		}
		time.Sleep(3 * time.Second)
	}
}

func TestSub1(t *testing.T) {
	conf := nsq.NewConfig()
	conf.AuthSecret = "123"
	consumer, err := nsq.NewConsumer("t1", "c1", conf)
	if err != nil {
		t.Fatal(err)
	}
	consumer.AddHandler(new(handle))
	err = consumer.ConnectToNSQD(host)
	if err != nil {
		t.Fatal(err)
	}
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
}

func TestSub2(t *testing.T) {
	conf := nsq.NewConfig()
	conf.AuthSecret = "123"
	consumer, err := nsq.NewConsumer("t1", "c1", conf)
	if err != nil {
		t.Fatal(err)
	}
	consumer.AddHandler(new(handle))
	err = consumer.ConnectToNSQD(host)
	if err != nil {
		t.Fatal(err)
	}
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
}
