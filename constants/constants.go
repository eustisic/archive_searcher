package constants

const RootUrl string = "https://archive.org/"
const ListUrl string = "https://archive.org/advancedsearch.php?q=subject:%22Miles%20Davis%22&sort=+date&mediatype:audio&output=json&rows=2"
const SearchUrl string = "https://archive.org/advancedsearch.php?q=subject:Miles+Davis&sort=-date&mediatype:audio&output=json&rows=1000"

type ApiResponse struct {
	Response ResponseContent `json:"response"`
}

type ResponseContent struct {
	Docs     []Doc `json:"docs"`
	NumFound int   `json:"numFound"`
	Start    int   `json:"start"`
}

type Doc struct {
	BackupLocation string   `json:"backup_location"`
	Btih           string   `json:"btih"`
	Collection     []string `json:"collection"`
	Creator        string   `json:"creator"`
	Date           string   `json:"date"`
	Description    string   `json:"description"`
	Downloads      int      `json:"downloads"`
	Format         []string `json:"format"`
	Identifier     string   `json:"identifier"`
}
