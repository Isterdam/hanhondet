package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Info struct {
	Ord       string
	Genus     string
	Sokningar string
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	info := Info{}

	if err := r.ParseForm(); err != nil {
		fmt.Println(fmt.Errorf("Error: %v"), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	info.Ord = strings.ToLower(r.Form.Get("ord"))

	infos, err := page.SearchInfo(&info)
	if err != nil {
		http.Redirect(w, r, "/assets/mainPage.html?ord=Error&genus=Error&sokningar=Error", 301)
	} else {
		url := "/assets/mainPage.html?ord=" + infos[0].Ord + "&genus=" + infos[0].Genus + "&sokningar=" + infos[0].Sokningar
		for i := 0; i < len(infos); i++ {
			url += "&ord=" + infos[i].Ord + "&genus=" + infos[i].Genus + "&sokningar=" + infos[i].Sokningar
		}
		http.Redirect(w, r, url, 301)
	}
}
