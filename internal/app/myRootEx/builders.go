package myRootEx

func (r *RootEX) resource() {
	r.router.HandleFunc("/grab", r.Grab).Methods("POST")
	r.router.HandleFunc("/solve", r.Solve).Methods("GET")
}

// func (roo *RootEX) OpenJSON(){

// }
