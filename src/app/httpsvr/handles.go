package httpsvr

import (
	"app/common"
	"net/http"
)

func getCityName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cityId := r.FormValue("cityId")
	//cityName :=common.CityList[cityId]
	if cityName, ok := common.CityList[cityId]; !ok {
		w.Write([]byte(""))
	} else {
		w.Write([]byte(cityName))
	}
}
