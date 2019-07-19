# geoip
Local http api for transfering ip address to geo info.

Imported from:
github.com/oschwald/geoip2-golang

Usage:

curl http://127.0.0.1:8000/myip

curl http://127.0.0.1:8000/geoip/8.8.8.8
Response:
{
  "ip": "8.8.8.8",
  "geoip": {
    "location": {
      "lat": 37.751,
      "lon": -97.822
    },
    "region_name": "United States",
    "country_iso_code": "US",
    "city_name": "",
    "continent_name": "North America"
  }
}
