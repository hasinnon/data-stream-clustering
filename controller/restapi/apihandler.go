package restapi

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	den "my/ar/399/datastream/controller/clustering/denstream"
	io "my/ar/399/datastream/controller/clustering/io"
	"my/ar/399/datastream/controller/security/jwt"
	"my/ar/399/datastream/controller/utility/ptime"
	"my/ar/399/datastream/controller/utility/strrand"
	"my/ar/399/datastream/datalayer"
	template "my/ar/399/datastream/view/gotemplate"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// UserAPIHandler stores db
type UserAPIHandler struct {
	dbhandler datalayer.MyDB
}

func newUserAPIHandler(db datalayer.MyDB) *UserAPIHandler {
	return &UserAPIHandler{
		dbhandler: db,
	}
}

// *******************   Error page   ********************

// notFound get 404 error page
func (handler UserAPIHandler) notFound(w http.ResponseWriter, r *http.Request) {
	template.PageNotFoundHandler(w)
}

// UnderConstruction get login page
func (handler UserAPIHandler) UnderConstruction(w http.ResponseWriter, r *http.Request) {
	template.UnderConstructionHandler(w)
}

// *******************   Dashboard   ********************

// Dashboard get login page
func (handler UserAPIHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	template.DashboardHandler(user, w)
}

// *******************   Login   ********************

// Login get login page
func (handler UserAPIHandler) Login(w http.ResponseWriter, r *http.Request) {
	template.LoginHandler(w)
}

// PostLogin post login
func (handler UserAPIHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	type msgerror struct {
		ErrState bool   `json:"errstate"`
		ErrMsg   string `json:"errmsg"`
	}
	err := r.ParseForm()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		msg := msgerror{ErrState: true, ErrMsg: "خطایی رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var info jwt.LoginForm
	info.Username = r.Form["username"][0]
	info.Password = r.Form["password"][0]

	validerr := info.IsValid()
	if validerr != nil {
		msg := msgerror{ErrState: true, ErrMsg: validerr.Error()}
		json.NewEncoder(w).Encode(msg)
		return
	}

	ok := jwt.UserLogin(w, info, handler.dbhandler)
	if !ok {
		msg := msgerror{ErrState: true, ErrMsg: "کاربری با این مشخصات یافت نشد"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := msgerror{ErrState: false, ErrMsg: "/Dashboard"}
	json.NewEncoder(w).Encode(msg)
	return
}

// *******************   Signup   ********************

// Login get login page
func (handler UserAPIHandler) Signup(w http.ResponseWriter, r *http.Request) {
	template.SignupHandler(w)
}

// PostLogin post login
func (handler UserAPIHandler) PostSignup(w http.ResponseWriter, r *http.Request) {
	type msgerror struct {
		ErrState bool   `json:"errstate"`
		ErrMsg   string `json:"errmsg"`
	}
	err := r.ParseForm()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		msg := msgerror{ErrState: true, ErrMsg: "خطایی رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var info jwt.LoginForm
	info.Username = r.Form["username"][0]
	info.Password = r.Form["password"][0]

	validerr := info.IsValid()
	if validerr != nil {
		msg := msgerror{ErrState: true, ErrMsg: validerr.Error()}
		json.NewEncoder(w).Encode(msg)
		return
	}

	ok := jwt.UserLogin(w, info, handler.dbhandler)
	if !ok {
		msg := msgerror{ErrState: true, ErrMsg: "کاربری با این مشخصات یافت نشد"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := msgerror{ErrState: false, ErrMsg: "/Dashboard"}
	json.NewEncoder(w).Encode(msg)
	return
}

// *******************   Logout   ********************

// Logout get Signup page
func (handler UserAPIHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Expires:  time.Now(),
		Value:    "invalid",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/Login", 302)
}

// *******************   Services   ********************

// Services get services list page
func (handler UserAPIHandler) Services(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	servs, err := handler.dbhandler.GetUserAllServices(user.ID)
	if err != nil {
		var nullservs []datalayer.ServiceInfo
		template.ServicesHandler(user, nullservs, w)
	} else {
		template.ServicesHandler(user, servs, w)
	}
}

// *******************   NewService   ********************

// NewService get login page
func (handler UserAPIHandler) NewService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	template.AddServiceHandler(user, w)
}

// *******************   AddStructuredService   ********************

// AddStructuredService get AddStructuredService page
func (handler UserAPIHandler) AddStructuredService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	template.AddStructuredServiceHandler(user, w)
}

// PostAddStructuredService handel form of AddStructuredService page
func (handler UserAPIHandler) PostAddStructuredService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type msgerror struct {
		ErrStatus bool   `json:"errstatus"`
		ErrMsg    string `json:"errmsg"`
	}
	err := r.ParseForm()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	user, ok := jwt.IsLogedin(r)
	if !ok {
		msg := msgerror{ErrStatus: false, ErrMsg: "/Login"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	// step 1
	sname := r.Form["sname"][0]

	//step 2
	count, err := strconv.Atoi(r.Form["count"][0])
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد count رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	var dataname den.PointVars
	dataname.Vars = make(map[string]bool)
	for i := 0; i <= count; i++ {
		temp1, ok := r.Form["data_"+fmt.Sprint(i)]
		if ok {
			dataname.Vars[temp1[0]] = true
		}
	}
	//step 3
	meps := 0.0
	dbeps := 0.0
	lambda := 0.0
	mepstemp, ok := r.Form["meps"]
	if ok {
		meps, err = strconv.ParseFloat(mepstemp[0], 64)
		if err != nil {
			msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد epsilon رخ داده است."}
			json.NewEncoder(w).Encode(msg)
			return
		}
	}

	dbepstemp, ok := r.Form["dbeps"]
	if ok {
		dbeps, err = strconv.ParseFloat(dbepstemp[0], 64)
		if err != nil {
			msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد epsilon رخ داده است."}
			json.NewEncoder(w).Encode(msg)
			return
		}
	}

	lmbdatemp, ok := r.Form["lmbda"]
	if ok {
		lambda, err = strconv.ParseFloat(lmbdatemp[0], 64)
		if err != nil {
			msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد epsilon رخ داده است."}
			json.NewEncoder(w).Encode(msg)
			return
		}
		if lambda <= 0 {
			msg := msgerror{ErrStatus: true, ErrMsg: "مقدار ضریب اهمیت داده های قدیمی حتما باید بزرگ تر از صفر باشد"}
			json.NewEncoder(w).Encode(msg)
			return
		}
	}

	//step 5
	active := false
	if _, ok := r.Form["active"]; ok {
		active = true
	}

	expstr := ""
	var exp time.Time = time.Now().Add(time.Hour * 24 * 90)
	var exptime time.Time = exp
	temp2, timeok := r.Form["exptime"]
	if timeok {
		expstr = temp2[0]
		if expstr != "" {
			temptime, err := ptime.ToGregorian(expstr)
			if err != nil {
				msg := msgerror{ErrStatus: true, ErrMsg: err.Error()}
				json.NewEncoder(w).Encode(msg)
				return
			}
			exptime = temptime
		}
	}
	//
	dbserv := datalayer.Service{
		Sid: sql.NullInt32{
			Int32: -1,
			Valid: false,
		},
		Name: sql.NullString{
			String: sname,
			Valid:  true,
		},
		Created: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Credit: sql.NullInt32{
			Int32: 90 * 24,
			Valid: true,
		},
		HDcount: sql.NullInt32{
			Int32: int32(len(dataname.Vars)),
			Valid: true,
		},
		Deleted: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		Type: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		UID: sql.NullInt32{
			Int32: int32(user.ID),
			Valid: true,
		},
	}
	dbclu := datalayer.ServiceClusters{
		Clusters: sql.NullString{
			String: "",
			Valid:  false,
		},
	}

	lastid, err := handler.dbhandler.InsertService(dbserv, dbclu)
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در ارتباط با پایگاه داده رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	//
	var serv den.DenStream
	serv.Start(meps, dbeps, lambda)
	key := strrand.String(10)
	jvars, _ := json.Marshal(dataname)
	jden, _ := json.Marshal(serv)
	strucserv := datalayer.StructureServices{
		Sid: sql.NullInt32{
			Int32: int32(lastid),
			Valid: true,
		},
		Key: sql.NullString{
			String: key,
			Valid:  true,
		},
		Expiration: sql.NullTime{
			Time:  exp,
			Valid: true,
		},
		ExpScheduled: sql.NullTime{
			Time:  exptime,
			Valid: true,
		},
		Credit: sql.NullInt32{
			Int32: 90 * 24,
			Valid: true,
		},
		NotUsed: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		Status: sql.NullBool{
			Bool:  active,
			Valid: true,
		},
		Variables: sql.NullString{
			String: string(jvars),
			Valid:  true,
		},
		DenStream: sql.NullString{
			String: string(jden),
			Valid:  true,
		},
		Used: 0,
	}

	err = handler.dbhandler.InsertStructureServices(strucserv)
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در ارتباط با پایگاه داده رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := msgerror{ErrStatus: false, ErrMsg: "/Service/" + fmt.Sprint(lastid) + "/Management"}
	json.NewEncoder(w).Encode(msg)
	return
}

// *******************   AddUnStructuredService   ********************

// AddUnStructuredService get AddUnStructuredService page
func (handler UserAPIHandler) AddUnStructuredService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	template.AddUnStructuredServiceHandler(user, w)
}

// PostAddUnStructuredService post AddUnStructuredService
func (handler UserAPIHandler) PostAddUnStructuredService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		fmt.Fprint(w, "/Login")
	}

	r.ParseMultipartForm(40 * 1024)
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}

	defer file.Close()

	rand.Seed(time.Now().UTC().UnixNano())
	randnum := 100 + rand.Intn(899)
	fname := fmt.Sprint(user.ID) + "-" + fmt.Sprint(randnum) + ".csv"
	path := filepath.Join("./datalayer/UploadedFiles/", fname)
	newFile, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}

	defer newFile.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}
	_, err = newFile.Write(fileByte)
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}

	fmt.Fprintln(w, "/Service/"+fmt.Sprint(randnum)+"/AcceptFile/"+header.Filename)

}

// *******************   AcceptFileUS   ********************

// AcceptFileUS get AcceptFile page
func (handler UserAPIHandler) AcceptFileUS(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	v := mux.Vars(r)
	file, err := os.Open("./datalayer/UploadedFiles/" + fmt.Sprint(user.ID) + "-" + v["sid"] + ".csv")
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}
	table, tr, tc, _ := io.InfoCSV(file)
	template.AcceptFileUSHandler(user, table, v["filename"], tr, tc, w)
}

// *******************   ClusteredUS   ********************

// ClusteredUS get Clustered page
func (handler UserAPIHandler) ClusteredUS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type msgerror struct {
		ErrStatus bool   `json:"errstatus"`
		ErrMsg    string `json:"errmsg"`
	}

	user, ok := jwt.IsLogedin(r)
	if !ok {
		msg := msgerror{ErrStatus: false, ErrMsg: "/Login"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := r.ParseForm()
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	meps, err := strconv.ParseFloat(r.Form["micro_eps"][0], 64)
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد epsilon رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	deps, err := strconv.ParseFloat(r.Form["db_eps"][0], 64)
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد epsilon رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	v := mux.Vars(r)
	file, err := os.Open("./datalayer/UploadedFiles/" + fmt.Sprint(user.ID) + "-" + v["sid"] + ".csv") // For read access.
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فایل داده‌ها رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}
	//
	countCol, err := strconv.Atoi(r.Form["col"][0])
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: "خطایی در خواندن فیلد col رخ داده است."}
		json.NewEncoder(w).Encode(msg)
		return
	}
	Ccheck := make([]bool, countCol)
	for i := 0; i < len(Ccheck); i++ {
		if _, ok := r.Form["check_"+fmt.Sprint(i)]; ok {
			Ccheck[i] = true
		} else {
			Ccheck[i] = false
		}
	}
	Hcheck := false
	if v, ok := r.Form["check_hdr"]; ok {
		if v[0] == "true" {
			Hcheck = true
		}
	}

	start := time.Now()
	print("start")
	skey := fmt.Sprint(user.ID) + "-" + v["sid"]
	infoInput := io.InfoTable{
		Check:    Ccheck,
		HdrCheck: Hcheck,
		Sname:    v["filename"],
		ServKey:  skey,
		UID:      user.ID,
		Meps:     meps,
		DBeps:    deps,
	}

	rsid, err := io.Clusteringfile(handler.dbhandler, file, infoInput)
	duration := time.Since(start)
	print("finish", " | ", float32(duration.Seconds()))
	if err != nil {
		msg := msgerror{ErrStatus: true, ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := msgerror{ErrStatus: false, ErrMsg: "/Service/" + fmt.Sprint(rsid) + "/Result?t=" + fmt.Sprintf("%.2f", duration.Seconds())}
	json.NewEncoder(w).Encode(msg)
	return
}
func (handler UserAPIHandler) progressUS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, ok := jwt.IsLogedin(r)
	if !ok {
		fmt.Fprintf(w, "")
		return
	}
	v := mux.Vars(r)
	file, err := os.Open("./datalayer/UploadedFiles/" + fmt.Sprint(user.ID) + "-" + v["sid"] + ".csv") // For read access.
	if err != nil {
		fmt.Fprintln(w, "err :", err)
		return
	}
	scanner := bufio.NewScanner(bufio.NewReader(file))
	counter := 0
	for scanner.Scan() {
		counter++
	}
	io.Progressbar[fmt.Sprint(user.ID)+"-"+v["sid"]] = make(chan int)
	var progress int
	var cr int
	var pr int
	for true {
		f, t := w.(http.Flusher)
		if t {
			cr = <-io.Progressbar[fmt.Sprint(user.ID)+"-"+v["sid"]]
			pr = int((float32(cr) / float32(counter)) * 100)
			if pr < 0 {
				break
			}
			if pr != progress {
				progress = pr
				fmt.Fprintf(w, "data: %d%%\n\n", progress)
				f.Flush()
			}

			if pr > 99 {
				break
			}

		} else {
			println("error in flush")
		}
	}
}

// *******************   ResultService   ********************

// ResultService get ResultService page
func (handler UserAPIHandler) ResultService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	v := mux.Vars(r)
	sid, _ := strconv.Atoi(v["sid"])
	serv, jclu, err := handler.dbhandler.GetServiceByID(sid, user)
	if err != nil {
		handler.notFound(w, r)
		return
	}
	clu := den.USresult{}
	json.Unmarshal([]byte(jclu.Clusters.String), &clu)
	template.ResultServiceHandler(user, v["sid"], serv, clu, r.FormValue("t"), w)
}

// ResultService get ResultService page
func (handler UserAPIHandler) ResultStrService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	v := mux.Vars(r)
	sid, _ := strconv.Atoi(v["sid"])
	serv, jclu, err := handler.dbhandler.GetServiceByID(sid, user)
	if err != nil {
		handler.notFound(w, r)
		return
	}
	clu := den.USresult{}
	json.Unmarshal([]byte(jclu.Clusters.String), &clu)

	template.ViewResultHandler(user, v["sid"], serv, clu, w)
}

// *******************   FindCluster   ********************

func (handler UserAPIHandler) FindCluster(w http.ResponseWriter, r *http.Request) {

	user, ok := jwt.IsLogedin(r)
	if !ok {
		fmt.Fprint(w, "نشست کاری به پایان رسیده است")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "خطایی رخ داده است.")
		return
	}

	v := mux.Vars(r)
	//
	DataCount, err := strconv.Atoi(r.Form["count"][0])
	if err != nil {
		fmt.Fprint(w, "خطایی در خواندن فیلد count رخ داده است.")
		return
	}
	input := make([]float64, DataCount)
	for i := 0; i < len(input); i++ {
		rec, err := strconv.ParseFloat(r.Form["data0_"+fmt.Sprint(i)][0], 64)
		if err != nil {
			input[i] = 0
		} else {
			input[i] = rec
		}
	}

	sid, _ := strconv.Atoi(v["sid"])
	_, jclu, err := handler.dbhandler.GetServiceByID(sid, user)
	if err != nil {
		return //????????
	}
	clu := den.USresult{}
	json.Unmarshal([]byte(jclu.Clusters.String), &clu)
	clus := clu.FindClustersIDs(input)
	result := make(map[int]int)
	for i := 0; i < len(clus); i++ {
		result[clus[i][0]] = clus[i][0] + 1
	}
	switch len(result) {
	case 0:
		fmt.Fprint(w, "<span class=\"price-tax\">"+"این داده متعلق به هیچ خوشه‌ای نیست و به عنوان noise شناخته می‌شود."+"</span>")
		return
	case 1:
		fmt.Fprint(w, "<span class=\"price-tax\">"+"این داده متعلق به خوشه‌ی زیر می‌باشد."+"</span>")
		for _, v := range result {
			fmt.Fprint(w, "<h1 class=\"price\">"+"cluster "+fmt.Sprint(v)+"</h1>")
		}

		return
	default:
		fmt.Fprint(w, "<span class=\"price-tax\">"+"این داده می‌تواند به خوشه‌ها‌ی زیر تعلق داشته باشد."+"</span>")
		for _, v := range result {
			fmt.Fprint(w, "<h1 class=\"price\">"+"cluster "+fmt.Sprint(v)+"</h1>")
		}

	}
}

func (handler UserAPIHandler) CheckData(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		fmt.Fprint(w, "نشست کاری به پایان رسیده است")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "خطایی رخ داده است.")
		return
	}

	v := mux.Vars(r)
	//
	DataCount, err := strconv.Atoi(r.Form["count"][0])
	if err != nil {
		fmt.Fprint(w, "خطایی در خواندن فیلد count رخ داده است.")
		return
	}
	input1 := make([]float64, DataCount)
	for i := 0; i < len(input1); i++ {
		rec, err := strconv.ParseFloat(r.Form["data1_"+fmt.Sprint(i)][0], 64)
		if err != nil {
			input1[i] = 0
		} else {
			input1[i] = rec
		}
	}

	input2 := make([]float64, DataCount)
	for i := 0; i < len(input2); i++ {
		rec, err := strconv.ParseFloat(r.Form["data2_"+fmt.Sprint(i)][0], 64)
		if err != nil {
			input2[i] = 0
		} else {
			input2[i] = rec
		}
	}

	sid, _ := strconv.Atoi(v["sid"])
	_, jclu, err := handler.dbhandler.GetServiceByID(sid, user)
	if err != nil {
		return //????????
	}
	clu := den.USresult{}
	json.Unmarshal([]byte(jclu.Clusters.String), &clu)
	clus1 := clu.FindClustersIDs(input1)
	clus2 := clu.FindClustersIDs(input2)
	isaCluster := false
	isaMcluster := false
	for i := 0; i < len(clus1); i++ {
		for j := 0; j < len(clus2); j++ {
			if clus1[i][0] == clus2[j][0] {
				isaCluster = true
				if clus1[i][1] == clus2[j][1] {
					isaMcluster = true
				}
			}
		}
	}
	if isaCluster {
		fmt.Fprint(w, "<span class=\"price-tax\">"+"این دو داده در یک خوشه قرار دارند."+"</span>")
		if isaMcluster {
			fmt.Fprint(w, "<span class=\"price-tax\">"+"و همچنین این دو داده در یک میکرو خوشه قرار دارند."+"</span>")
		} else {
			fmt.Fprint(w, "<span class=\"price-tax\">"+"اما این دو داده در یک میکرو خوشه قرار ندارند."+"</span>")
		}
	} else {
		fmt.Fprint(w, "<span class=\"price-tax\">"+"این دو داده در یک خوشه قرار ندارند."+"</span>")
	}
}

// *******************   Management Service   ********************

// ResultService get ResultService page
func (handler UserAPIHandler) ManagementService(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.IsLogedin(r)
	if !ok {
		http.Redirect(w, r, "/Login", 302)
	}
	v := mux.Vars(r)
	sid, _ := strconv.Atoi(v["sid"])
	serv, err := handler.dbhandler.GetStructureServicesInfoByID(int(user.ID), sid)
	if err != nil {
		handler.notFound(w, r)
		return
	}
	var svars den.PointVars
	json.Unmarshal([]byte(serv.SServ.Variables.String), &svars)
	template.ManagementServiceHandler(user, serv, svars, v["sid"], w)
	// serv.BServ.Credit.Int32

}

// *******************   Online Clustering  ********************

// ResultService get ResultService page
func (handler UserAPIHandler) OnlineClustering(w http.ResponseWriter, r *http.Request) {
	req, ok := jwt.IsValidServiceRequest(r)
	if !ok {
		fmt.Fprintln(w, "درخواست نامعتبر")
	}
	serv, err := handler.dbhandler.Value(req.SID)
	if err != nil {
		fmt.Fprintln(w, "سرویسی با این شناسه موجود نیست")
	}
	if len(serv.Vars.Vars) == 0 {

	}

}
