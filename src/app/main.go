package main

func main() {
	r := registerRoutes() //create a default router that has all of the basic middlewares
	// like logging gzip middleware!
	// http.Handle("/img/", http.FileServer(http.Dir("public")))
	// http.Handle("/css/", http.FileServer(http.Dir("public")))
	// http.Handle("/js/", http.FileServer(http.Dir("public")))
	r.Run(":3000")
}
