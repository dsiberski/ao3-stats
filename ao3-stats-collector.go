package ao3stats

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ao3_site := "https://archiveofourown.org/works?commit=Sort+and+Filter&work_search%5Bsort_column%5D=kudos_count&work_search%5Bother_tag_names%5D=&work_search%5Bexcluded_tag_names%5D=&work_search%5Bcrossover%5D=&work_search%5Bcomplete%5D=&work_search%5Bwords_from%5D=&work_search%5Bwords_to%5D=&work_search%5Bdate_from%5D=&work_search%5Bdate_to%5D=&work_search%5Bquery%5D=&work_search%5Blanguage_id%5D=&tag_id=One+Piece+%28Anime+*a*+Manga%29"

	//one_piece_site := "https://archiveofourown.org/tags/One%20Piece%20(Anime%20*a*%20Manga)/works"

	response, err := http.Get(ao3_site)
	if err != nil {
		log.Fatalln(err)
	}

	// Read response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Save body data
	ao3_data := string(body)

	data_file, err := os.Create("ao3-data.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := data_file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = data_file.WriteString(ao3_data)
	if err != nil {
		log.Fatalln(err)
	}
}
