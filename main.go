package main

import (
	//	"encoding/json"
	//	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"

	"github.com/kataras/iris"
)

type Geoinfo struct {
	Geoip Geo    `json:"geoip"`
	Ip    string `json:"ip"`
}

type Geo struct {
	//	Ip               string      `json:"ip"`
	Continent_name   string      `json:"continent_name"`
	City_name        string      `json:"city_name"`
	Country_iso_code string      `json:"country_iso_code"`
	Region_name      string      `json:"region_name"`
	Location         Geolocation `json:"location"`
}

type Geolocation struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func main() {
	app := iris.New()
	db, err := geoip2.Open("./db/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.Get("/geoip/{ip:string}", func(ctx iris.Context) {
		ipaddr := ctx.Params().Get("ip")
		ip := net.ParseIP(ipaddr)
		record, err := db.City(ip)
		if err != nil {
			log.Println(err)
		}
		output_geolocation := Geolocation{
			Lon: record.Location.Longitude,
			Lat: record.Location.Latitude,
		}

		Geoipstuct := Geo{
			Continent_name:   record.Continent.Names["en"],
			City_name:        record.City.Names["en"],
			Country_iso_code: record.Country.IsoCode,
			Region_name:      record.Country.Names["en"],
			Location:         output_geolocation,
		}
		output := Geoinfo{
			Ip:    ipaddr,
			Geoip: Geoipstuct,
		}
		ctx.JSON(output)
	})
	app.Get("/myip", func(ctx iris.Context) {
		ipaddr := ctx.GetHeader("X-Forwarded-For")
		if ipaddr == "" {
			ipaddr = ctx.RemoteAddr()
		}
		ip := net.ParseIP(ipaddr)
		record, err := db.City(ip)
		if err != nil {
			log.Println(err)
		}

		output_geolocation := Geolocation{
			Lon: record.Location.Longitude,
			Lat: record.Location.Latitude,
		}
		Geoipstuct := Geo{
			Continent_name:   record.Continent.Names["en"],
			City_name:        record.City.Names["en"],
			Country_iso_code: record.Country.IsoCode,
			Region_name:      record.Country.Names["en"],
			Location:         output_geolocation,
		}
		output := Geoinfo{
			Ip:    ipaddr,
			Geoip: Geoipstuct,
		}
		ctx.JSON(output)
	})

	app.Run(iris.Addr(":8000"), iris.WithPostMaxMemory(50<<20))

}
