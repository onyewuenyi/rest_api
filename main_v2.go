

// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

// func main() {
// 	a := App{}
// 	a.Initialize(
// 		os.Getenv("USER"),
// 		os.Getenv("PASSWORD"),
// 		os.Getenv("DBNAME"),
// 		os.Getenv("DBHOST"),
// 		os.Getenv("DBPORT"))
// 	a.Run("0.0.0.0", 8080)
// }

// type App struct {
// 	Router *mux.Router
// 	DB     *sql.DB
// }

// // Define interfaces of App struct
// func (a *App) Initialize(user string, password string, dbname string, dbhost string, dbport string) {}
// func (a *App) Run(address string, port int)                                                         {}

// type subscription struct {
// 	Email string
// 	Name  string
// }
