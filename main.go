package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sort"
	// Import this so we don't have to use qm.Limit etc.
	"github.com/marekswiecznik/ohiboka2-sql-data-to-mongo/models"
	"github.com/marekswiecznik/ohiboka2-sql-data-to-mongo/models/mongo"
	"github.com/vattle/sqlboiler/boil"
	. "github.com/vattle/sqlboiler/queries/qm"
)

func main() {
	url := os.Args[1]
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)
	bracelets, err := models.BraceletBracelets(db,
		Load("Category"),
		Load("Photo"),
		Load("User"),
		Load("BraceletBraceletBraceletknots"),
		Load("BraceletBraceletBraceletstrings"),
		Load("BraceletBraceletBraceletstrings.Color"),
		Load("BraceletBraceletPhotos"),
		Load("BraceletBraceletPhotos.User"),
		Load("BraceletBraceletRates"),
		Load("BraceletBraceletRates.User"),
		//Where("URL = 'kwadraty'"),
	).All()
	if err != nil {
		log.Fatal(err)
	}
	var mongoBracelets []*mongo.Bracelet
	for _, b := range bracelets {
		bracelet := createBracelet(b)
		if bracelet != nil {
			mongoBracelets = append(mongoBracelets, bracelet)
			jsonString, _ := json.MarshalIndent(bracelet, "", "  ")
			fmt.Println(string(jsonString))
		}
	}
	//jsonString, _ := json.MarshalIndent(mongoBracelets, "", "  ")
	//fmt.Println(string(jsonString))
}

func createBracelet(b *models.BraceletBracelet) *mongo.Bracelet {
	bracelet := &mongo.Bracelet{}
	bracelet.ID = b.ID
	bracelet.Created = b.Date
	bracelet.Name = b.Name
	bracelet.Accepted = b.Accepted == 1
	bracelet.Difficulty = b.Difficulty
	if b.R.Category != nil {
		bracelet.Category = &b.R.Category.Name
	}
	bracelet.Rate = b.Rate
	bracelet.Public = b.Public == 1
	bracelet.Slug = b.URL
	bracelet.Deleted = b.Deleted == 1
	bracelet.Author = createUser(b.R.User)
	bracelet.Photo = createPhoto(b.R.Photo)
	bracelet.Rates = createRates(b.R.BraceletBraceletRates)
	bracelet.Photos = createPhotos(b.R.BraceletBraceletPhotos)
	bracelet.Strings = createStrings(b.R.BraceletBraceletBraceletstrings)
	stringsLength := len(bracelet.Strings)
	if stringsLength < 2 {
		return nil
	}
	if b.Type == 1 {
		bracelet.Type = "standard"
		bracelet.Rows = createStandardRows(stringsLength, b.R.BraceletBraceletBraceletknots)
	} else {
		bracelet.Type = "text"
		bracelet.Rows = createTextRows(stringsLength, b.R.BraceletBraceletBraceletknots)
	}

	return bracelet
}

func createUser(user *models.AuthUser) mongo.User {
	return mongo.User{
		ID:   user.ID,
		Name: user.Username,
	}
}

func createPhoto(photo *models.BraceletPhoto) *mongo.Photo {
	if photo == nil {
		return nil
	}
	user, _ := models.AuthUsersG(Where("id=?", photo.UserID)).One()
	return &mongo.Photo{
		Filename: photo.Name,
		Accepted: photo.Accepted == 1,
		Author:   createUser(user),
	}
}

func createRates(rates models.BraceletRateSlice) []mongo.Rate {
	arr := make([]mongo.Rate, len(rates))
	for i, r := range rates {
		arr[i] = mongo.Rate{
			Rate: r.Rate,
			User: createUser(r.R.User),
		}
	}
	return arr
}

func createPhotos(rates models.BraceletPhotoSlice) []mongo.Photo {
	arr := make([]mongo.Photo, len(rates))
	for i, p := range rates {
		arr[i] = mongo.Photo{
			Filename: p.Name,
			Accepted: p.Accepted == 1,
			Author:   createUser(p.R.User),
		}
	}
	return arr
}

type ByStringIndex models.BraceletBraceletstringSlice

func (s ByStringIndex) Len() int {
	return len(s)
}
func (s ByStringIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByStringIndex) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

func createStrings(strings models.BraceletBraceletstringSlice) []string {
	sort.Sort(ByStringIndex(strings))

	arr := make([]string, len(strings))
	for i, s := range strings {
		color := s.R.Color
		if color != nil {
			arr[i] = fmt.Sprintf("#%06X", color.Hexcolor)
		} else {
			arr[i] = "#000000"
		}
	}
	return arr
}

type ByKnotIndex models.BraceletBraceletknotSlice

func (s ByKnotIndex) Len() int {
	return len(s)
}
func (s ByKnotIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByKnotIndex) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

func createStandardRows(stringsLength int, knots models.BraceletBraceletknotSlice) []mongo.Row {
	sort.Sort(ByKnotIndex(knots))

	nofrows := 2 * len(knots) / (stringsLength - 1)
	if len(knots)-nofrows*(stringsLength/2.0) > 0 {
		nofrows += 1
	}
	arr := make([]mongo.Row, nofrows)
	nofcols := stringsLength / 2
	evenStrings := stringsLength%2 == 0

	knotsInRow := make([]mongo.Knot, nofcols)
	// first row
	for j := 0; j < nofcols; j += 1 {
		knotsInRow[j] = knotForTypeId(knots[j].KnottypeID)
	}
	arr[0] = mongo.Row{
		Odd:   false,
		Knots: knotsInRow,
	}
	// other rows
	index := nofcols
	var noc int
	for i := 1; i < nofrows; i += 1 {
		if evenStrings {
			noc = nofcols - (i % 2)
		} else {
			noc = nofcols
		}
		knotsInRow := make([]mongo.Knot, noc)
		for j := 0; j < noc; j += 1 {
			knotsInRow[j] = knotForTypeId(knots[index].KnottypeID)
			index += 1
		}
		arr[i] = mongo.Row{
			Odd:   i%2 == 1,
			Knots: knotsInRow,
		}
	}
	return arr
}

func createTextRows(stringsLength int, knots models.BraceletBraceletknotSlice) []mongo.Row {
	sort.Sort(ByKnotIndex(knots))

	nofrows := (len(knots) + 1) / (stringsLength - 1)
	arr := make([]mongo.Row, nofrows)
	nofcols := stringsLength - 1
	for i := 0; i < nofrows; i += 1 {
		knotsInRow := make([]mongo.Knot, nofcols)
		for j := 0; j < nofcols; j += 1 {
			knotsInRow[j] = knotForTypeId(knots[i*nofcols+j].KnottypeID)
		}
		arr[i] = mongo.Row{
			Odd:   i%2 == 0,
			Knots: knotsInRow,
		}
	}

	return arr
}

func knotForTypeId(id int) mongo.Knot {
	return mongo.Knot{Type: knotTypeForId(id)}
}

func knotTypeForId(id int) string {
	switch id {
	case 1:
		return "right-regular"
	case 2:
		return "left-regular"
	case 3:
		return "right-backward"
	case 4:
		return "left-backward"
	case 5:
		return "horizontal"
	case 6:
		return "vertical"
	default:
		return "" // FIXME fatal
	}
}
