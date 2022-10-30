package albumRestApi

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Albums []Album

func (albums *Albums) GetAlbumByID(id string) (Album, bool) {
	// For vs For-each:
	// albumLength := len(albums)
	// for currIdx := 0; currIdx < albumLength; currIdx++ {
	// 	var currAlbum = albums[currIdx]

	// Use _ in place of currIdx to skip the index value
	// To skip the slice item, just omit currAlbum
	for currIdx, currAlbum := range *albums {
		if currAlbum.ID == id {
			println("Found at index", currIdx)

			return currAlbum, true
		}
	}

	return Album{}, false
}

func (albums *Albums) AddAlbum(newAlbum Album) {
	*albums = append(*albums, newAlbum)
}
