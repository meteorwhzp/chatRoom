package httpsvr

import (
	logger "github.com/shengkehua/xlog4go"
	"net/http"
)

func Init() error {
	http.HandleFunc("/getCityName", getCityName)

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			panic(err)
		}
	}()
	logger.Info("Listening http on port 8000")

	return nil
}
