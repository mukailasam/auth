package routers

func (r *Router) Route() {
	// Home
	r.Mux.HandleFunc("/api/home", r.Handler.Index)

	// Register user
	r.Mux.HandleFunc("/api/auth/register", r.Handler.Register)
	// verify user
	r.Mux.HandleFunc("/api/auth/verify/{username}/{token}", r.Handler.VerifyUser)
	// login user
	r.Mux.HandleFunc("/api/auth/login", r.Handler.Login)
	// Logout User
	r.Mux.HandleFunc("/api/auth/logout", r.Handler.Logout)
	// Delete User
	r.Mux.HandleFunc("/api/auth/delete", r.Handler.DeleteUserAccount)

}
