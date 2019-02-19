package zone

type zone struct {
	Name string `xml:"name"`
}

type ResourceRecord struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
	TTL   uint   `xml:"ttl"`
}
