package geocoder

type objMetaData struct {
	h string
	l []loc
}

type loc struct {
	src    city
	coslat float64
	sinlat float64
	coslon float64
	sinlon float64
}

type city struct {
	name         string
	lat          float64
	lon          float64
	country_code string
	country      string
	region_code  string
	region       string
}
