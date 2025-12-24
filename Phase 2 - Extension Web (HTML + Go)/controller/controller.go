package controller

import (
	"convert/logic"
	"html/template"
	"net/http"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Resultat1": "",
		"Entree1":   "",
		"Erreur1":   "",
		"Resultat2": "",
		"Entree2":   "",
		"Erreur2":   "",
		"Resultat3": "",
		"Entree3":   "",
		"Erreur3":   "",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
func Km2miles(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Resultat1": "",
		"Entree1":   "",
		"Erreur1":   ""}
	if r.Method == http.MethodPost {
		kilometers := r.FormValue("kilometers")
		km, err1 := strconv.ParseFloat(kilometers, 64)
		if err1 != nil {
			data["Erreur1"] = "Veuillez entrer un nombre valide"
		} else if km == 0 || km < 0 {
			data["Erreur1"] = "La valeur ne peut pas être 0 ou en dessous"
		} else {
			miles := logic.Kilometers2Miles(km)
			data["Resultat1"] = miles
			data["Entree1"] = km
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html")) // ✅ Ajouter
	tmpl.Execute(w, data)
}

func Kg2pounds(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Resultat2": "",
		"Entree2":   "",
		"Erreur2":   ""}
	if r.Method == http.MethodPost {
	}

	if r.Method == http.MethodPost {
		kilograms := r.FormValue("kilograms")
		kg, err2 := strconv.ParseFloat(kilograms, 64)
		if err2 != nil {
			data["Erreur2"] = "Veuillez entrer un nombre valide"
		} else if kg == 0 || kg < 0 {
			data["Erreur2"] = "La valeur ne peut pas être 0 ou en dessous"
		} else {
			pounds := logic.Kilograms2pounds(kg)
			data["Resultat2"] = pounds
			data["Entree2"] = kg
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html")) // ✅ Ajouter
	tmpl.Execute(w, data)
}

func Ctofahrenheit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{

		"Resultat3": "",
		"Entree3":   "",
		"Erreur3":   ""}
	if r.Method == http.MethodPost {
		degresc := r.FormValue("degresc")
		deg, err3 := strconv.ParseFloat(degresc, 64)
		if err3 != nil {
			data["Erreur3"] = "Veuillez entrer un nombre valide"
		} else {
			degresf := logic.Celsius2Fahreneit(deg)
			data["Resultat3"] = degresf
			data["Entree3"] = degresc
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)

}
