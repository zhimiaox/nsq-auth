package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type pluginCSVSecret struct {
	Split string
}

func (p *pluginCSVSecret) Init() Plugin {
	p.Split = " "
	if SystemOpts.CSV == "" {
		return nil
	}
	info, err := os.Stat(SystemOpts.CSV)
	if err != nil {
		log.Printf("csv plugin init err %s\n", err.Error())
		return nil
	}
	if info.IsDir() {
		log.Println("csv plugin init err path is dir")
		return nil
	}
	return p
}

func (p *pluginCSVSecret) Authorization() map[string][]Authorization {
	data, err := os.ReadFile(SystemOpts.CSV)
	if err != nil {
		log.Printf("csv file load err %s\n", err.Error())
		return nil
	}
	csvReader := csv.NewReader(bytes.NewReader(data))
	authResp := make(map[string][]Authorization)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) != 4 {
			log.Printf("csv file parse err %s\n", err.Error())
			continue
		}
		// secret, topic, channel, allow
		auth, ok := authResp[record[0]]
		if !ok {
			auth = make([]Authorization, 0)
		}
		authResp[record[0]] = append(auth, Authorization{
			Topic:       record[1],
			Channels:    strings.Split(strings.TrimSpace(record[2]), p.Split),
			Permissions: strings.Split(strings.TrimSpace(record[3]), p.Split),
		})
	}
	return authResp
}
