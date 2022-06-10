package main

import (
	"github.com/windeal/utilsgo/log"
)

func main() {
	c := &log.Config{
		Outputs: []log.OutputConfig{
			{
				Writer: "file",
				WriterConfig: log.WriterConfig{
					LogPath:    "./logs",
					Filename:   "example.log",
					MaxSize:    20,
					MaxAge:     10,
					MaxBackups: 10,
				},
				Level:      "debug",
				CallerSkip: 2,
			},
		},
	}
	log.InitLogger(c)

	for i := 0; i < 1000*1000; i++ {
		log.Info("this is some log: ", i)
	}

	log.Errorf("hello")
}
