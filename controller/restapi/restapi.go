package restapi

import (
	"fmt"
	"my/ar/399/datastream/datalayer"
	"net/http"

	"github.com/gorilla/mux"
)

// STATIC_DIR g*****
const STATIC_DIR = "/Static/"

// STATICUI_DIR g*****
const STATICUI_DIR = "/StaticUI/"

// STATICLogin_DIR g*****
const STATICLogin_DIR = "/StaticLogin/"

// JsonLogin_DIR g*****
const JsonLogin_DIR = "/JsonLogin/"

// IMAGE_DIR g*****
const IMAGE_DIR = "/Images/"

const plot = "/Plot/"

//Http Get search /api/person/name/{name}
//http POST add or Edit /api/person/add | localhost:8181/api/person/edit

// RunAPI Run Api
func RunAPI(endpoint string, db datalayer.MyDB) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	fmt.Println("Server Started ...")
	return http.ListenAndServe(endpoint, r)
}

// RunAPIOnRouter Run Routers
func RunAPIOnRouter(r *mux.Router, db datalayer.MyDB) {
	handler := newUserAPIHandler(db)

	r.Methods("Get").Path("/Login").HandlerFunc(handler.Login)
	r.Methods("POST").Path("/Login").HandlerFunc(handler.PostLogin)

	r.Methods("Get").Path("/Signup").HandlerFunc(handler.Signup)
	r.Methods("POST").Path("/Signup").HandlerFunc(handler.PostSignup)

	r.Methods("Get").Path("/Logout").HandlerFunc(handler.Logout)

	r.Methods("Get").Path("/Dashboard").HandlerFunc(handler.Dashboard)

	r.Methods("Get").Path("/Services").HandlerFunc(handler.Services)

	r.Methods("Get").Path("/NewService").HandlerFunc(handler.NewService)

	r.Methods("Get").Path("/AddUnStructuredService").HandlerFunc(handler.AddUnStructuredService)
	r.Methods("POST").Path("/AddUnStructuredService").HandlerFunc(handler.PostAddUnStructuredService)

	r.Methods("Get").Path("/Service/{sid:[0-9]+}/AcceptFile/{filename}").HandlerFunc(handler.AcceptFileUS)
	r.Methods("POST").Path("/Service/{sid:[0-9]+}/AcceptFile/{filename}/Clustering").HandlerFunc(handler.ClusteredUS)
	r.Path("/Service/{sid:[0-9]+}/AcceptFile/{filename}/progress").HandlerFunc(handler.progressUS)

	r.Methods("Get").Path("/Service/{sid:[0-9]+}/Result").HandlerFunc(handler.ResultService)
	r.Methods("POST").Path("/Service/{sid:[0-9]+}/Result/FindCluster").HandlerFunc(handler.FindCluster)
	r.Methods("POST").Path("/Service/{sid:[0-9]+}/Result/CheckData").HandlerFunc(handler.CheckData)

	r.Methods("Get").Path("/Service/{sid:[0-9]+}/View").HandlerFunc(handler.ResultStrService)

	r.Methods("Get").Path("/Service/{sid:[0-9]+}/Management").HandlerFunc(handler.ManagementService)

	r.Methods("Get").Path("/AddStructuredService").HandlerFunc(handler.AddStructuredService)
	r.Methods("POST").Path("/AddStructuredService").HandlerFunc(handler.PostAddStructuredService)

	r.Methods("Get").Path("/Credit").HandlerFunc(handler.UnderConstruction)
	r.Methods("POST").Path("/Credit").HandlerFunc(handler.UnderConstruction)
	r.Methods("Get").Path("/IncreaseCredit").HandlerFunc(handler.UnderConstruction)

	r.Methods("Get").Path("/Settings").HandlerFunc(handler.UnderConstruction)
	r.Methods("POST").Path("/Settings").HandlerFunc(handler.UnderConstruction)

	r.Methods("Get").Path("/Messages").HandlerFunc(handler.UnderConstruction)
	r.Methods("Get").Path("/ContactUs").HandlerFunc(handler.UnderConstruction)

	r.Methods("POST").Path("/Clustering").HandlerFunc(handler.OnlineClustering)

	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("./mainTemplate/"))))
	r.PathPrefix(STATICUI_DIR).Handler(http.StripPrefix(STATICUI_DIR, http.FileServer(http.Dir("./view/ui/"))))
	r.PathPrefix(STATICLogin_DIR).Handler(http.StripPrefix(STATICLogin_DIR, http.FileServer(http.Dir("./ui/login/static/"))))
	r.PathPrefix(JsonLogin_DIR).Handler(http.StripPrefix(JsonLogin_DIR, http.FileServer(http.Dir("./ui/login/"))))
	r.PathPrefix(IMAGE_DIR).Handler(http.StripPrefix(IMAGE_DIR, http.FileServer(http.Dir("./Content/Images"))))
	r.PathPrefix(plot).Handler(http.StripPrefix(plot, http.FileServer(http.Dir("./datalayer/plot"))))

	r.NotFoundHandler = http.HandlerFunc(handler.notFound)

}
