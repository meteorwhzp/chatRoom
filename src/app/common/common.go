package common

import (
	"bufio"
	"fmt"
	logger "github.com/shengkehua/xlog4go"
	"io"
	"os"
	"strings"
)

var CityList map[string]string = make(map[string]string)

func Init(path string) error {
	logger.Info("init city info list with path %s\n", path)
	err := loadPath(path)
	return err

}

func loadPath(path string) (err error) {

	fin, err := os.Open(path)
	if err != nil {
		return
	}
	defer fin.Close()
	buf := bufio.NewReader(fin)

	errCnt := 0

	for {
		line, _, e := buf.ReadLine()
		if e == io.EOF {
			break
		} else if e != nil {
			errCnt += 1
			continue
		}

		token := ","
		s_line := strings.Split(string(line), token)
		if len(s_line) < 8 {
			errCnt += 1
			continue
		}
		cityId := s_line[5]
		cityName := s_line[6]
		fmt.Printf("cityId: %s, cityname: %s\n", cityId, cityName)
		CityList[cityId] = cityName
	}

	if errCnt != 0 {
		logger.Warn("%s errCnt:%v", path, errCnt)
	}

	return nil
}
