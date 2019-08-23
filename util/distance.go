package util

import "math"

// return size: KM
// lat1 := 116.378972
//	lng1 := 39.991926
//	lat2 := 116.441638
//	lng2 := 39.993695
//	fmt.Println(GetDistance(lat1, lat2, lng1, lng2))
// 6.968690540084304 KM
func GetDistance(lat1, lat2, lng1, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
