package mahasiswacontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

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

func GetForm(w http.ResponseWriter, r *http.Request) {

	//dapat kan data id
	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	//jika err, atau tidak ada id, maka new data
	if err != nil {
		data = map[string]interface{}{
			"title": "Add New Data",
		}
	} else {
		var mahasiswa entities.Mahasiswa
		err := mahasiswaModel.Find(id, &mahasiswa)

		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title":     "Edit Data",
			"mahasiswa": mahasiswa,
		}
	}

	temp, _ := template.ParseFiles("views/mahasiswa/form.html")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	// jika requestnya adalah POST maka proses form nya
	if r.Method == http.MethodPost {
		r.ParseForm()
		var mahasiswa entities.Mahasiswa

		mahasiswa.NamaLengkap = r.Form.Get("nama_lengkap")
		mahasiswa.JenisKelamin = r.Form.Get("jenis_kelamin")
		mahasiswa.TanggalLahir = r.Form.Get("tanggal_lahir")
		mahasiswa.TempatLahir = r.Form.Get("tempat_lahir")
		mahasiswa.Alamat = r.Form.Get("alamat")

		//proses insert atau update data
		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
		var data map[string]interface{}

		if err != nil {
			//insert data
			err := mahasiswaModel.Create(&mahasiswa)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data Saved",
				"data":    template.HTML(GetData()),
			}
		} else {
			//update data
			mahasiswa.Id = id
			err := mahasiswaModel.Update(mahasiswa)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data updated",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	err = mahasiswaModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data Deleted",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
