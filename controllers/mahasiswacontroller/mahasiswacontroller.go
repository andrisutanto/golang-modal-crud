package mahasiswacontroller

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/andrisutanto/golang-modal-crud/entities"
	"github.com/andrisutanto/golang-modal-crud/models/mahasiswamodel"
)

var mahasiswaModel = mahasiswamodel.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/mahasiswa/index.html")
	temp.Execute(w, data)
}

func GetData() string {
	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/mahasiswa/data.html")

	var mahasiswa []entities.Mahasiswa
	err := mahasiswaModel.FindAll(&mahasiswa)

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()
}
