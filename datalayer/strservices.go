package datalayer

import (
	"database/sql"
	"encoding/json"
	"my/ar/399/datastream/controller/clustering/denstream"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// StructureServices stores information about a service
type StructureServices struct {
	Sid          sql.NullInt32  // int
	Key          sql.NullString // int
	Expiration   sql.NullTime   // int
	ExpScheduled sql.NullTime   // int
	Credit       sql.NullInt32  // int
	NotUsed      sql.NullInt32  // int
	Status       sql.NullBool   // int
	Variables    sql.NullString
	DenStream    sql.NullString
	Used         int
	JKey         string
}

// CService stores information about a structure service with prossecc
type CService struct {
	Sid          sql.NullInt32 // int
	ExpScheduled sql.NullTime  // int
	Status       sql.NullBool  // int
	Vars         PointVars
	DenStream    denstream.DenStream
	JKey         string
}

// PointVars stores information about names and status of variable of service
type PointVars struct {
	Vars map[string]bool `json:"vars"`
}

// ServiceLogin stores information about a login service
type ServiceLogin struct {
	SID int32
	Key string
}

// GetStructureServicesInfoByID returns a service with the given service ID.
func (handler *MyDB) GetStructureServicesInfoByID(userID, serviceID int) (ServiceInfo, error) {
	row := handler.db.QueryRow("select services.sid,services.name,services.created,services.credit,services.deleted,services.type,"+
		" strservices.expiration,strservices.scheduled,strservices.notused,strservices.status,strservices.variables"+
		" from clustream_sru.services right join clustream_sru.strservices on services.sid=strservices.sid"+
		" where services.uid=? and strservices.sid=?", userID, serviceID)

	return getRowDataStructureInfoServices(row)
}

func getRowDataStructureInfoServices(row *sql.Row) (ServiceInfo, error) {
	u := ServiceInfo{}

	err := row.Scan(&u.BServ.Sid, &u.BServ.Name, &u.BServ.Created, &u.BServ.Credit, &u.BServ.Deleted, &u.BServ.Type, &u.SServ.Expiration, &u.SServ.ExpScheduled, &u.SServ.NotUsed, &u.SServ.Status, &u.SServ.Variables)
	if err != nil {
		return u, err
	}
	u.SServ.UpdateKey()
	return u, nil
}

// GetStructureServicesInfoByID returns a service with the given service ID.
func (handler *MyDB) GetCServicesBySID(serviceID int32) (CService, error) {
	row := handler.db.QueryRow("select sid,skey,scheduled,status,variables,denstream  from clustream_sru.strservices where sid=?", serviceID)

	return getRowDataCServices(row)
}

func getRowDataCServices(row *sql.Row) (CService, error) {
	u := CService{}
	var tempkey sql.NullString
	var jvar sql.NullString
	var jdenstream sql.NullString
	err := row.Scan(&u.Sid, &tempkey, &u.ExpScheduled, &u.Status, &jvar, &jdenstream)
	if err != nil {
		return u, err
	}
	u.UpdateKey(tempkey.String)
	err = json.Unmarshal([]byte(jvar.String), &u.Vars)
	if err != nil {
		return u, err
	}
	err = json.Unmarshal([]byte(jdenstream.String), &u.DenStream)
	if err != nil {
		return u, err
	}
	return u, nil
}

// InsertStructureServices insert a doctor to database and returns err.
func (handler *MyDB) InsertStructureServices(serv StructureServices) error {
	_, err := handler.db.Exec("INSERT INTO strservices(sid,skey,expiration,scheduled,notused,status,variables,denstream) VALUES (?,?,?,?,?,?,?,?)", serv.Sid, serv.Key, serv.Expiration, serv.ExpScheduled, serv.NotUsed, serv.Status, serv.Variables, serv.DenStream)

	if err != nil {
		return err
	}

	return nil
}

// UpdateCServices update cache service to database and returns err.
func (handler *MyDB) UpdateCServices(serv *CService) error {
	jtemp1, err := json.Marshal(serv.Vars)
	if err != nil {
		return err
	}
	jvars := sql.NullString{
		String: string(jtemp1),
		Valid:  true,
	}

	jtemp2, err := json.Marshal(serv.DenStream)
	if err != nil {
		return err
	}
	jden := sql.NullString{
		String: string(jtemp2),
		Valid:  true,
	}
	_, err = handler.db.Exec("UPDATE strservices SET scheduled = ? , status = ?, variables = ? ,denstream = ?	WHERE sid=?	", serv.ExpScheduled, serv.Status, jvars, jden, serv.Sid)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

// Value update a  to database and returns err.
func (handler *MyDB) Value(serviceID int32) (*CService, error) {
	val, err := handler.cache.Value(serviceID)
	if err != nil {
		tempserv := CService{}
		serv, err := handler.GetCServicesBySID(serviceID)
		if err != nil {
			return &tempserv, err
		}
		handler.cache.Add(serviceID, 150*time.Second, &serv)
		res, err := handler.cache.Value(serviceID)
		if err != nil {
			return &tempserv, err
		}
		return res.Data().(*CService), nil

	}
	return val.Data().(*CService), nil

}

// JWT

var jwtKey = []byte("my_Secret_Key")

// ServiceClaims srores inforamtion about service jwt
type ServiceClaims struct {
	SID int32
	Key string
	jwt.StandardClaims
}

// ServiceJWTKey returns true if user logined corectlly ***
func ServiceJWTKey(serv StructureServices) string {

	expTime := serv.ExpScheduled.Time

	cliams := &ServiceClaims{
		SID: serv.Sid.Int32,
		Key: serv.Key.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		return ""
	}
	return stringToken
}

func (serv *StructureServices) UpdateKey() string {
	serv.JKey = ServiceJWTKey(*serv)
	return serv.JKey
}

func (serv *CService) UpdateKey(skey string) string {
	expTime := serv.ExpScheduled.Time

	cliams := &ServiceClaims{
		SID: serv.Sid.Int32,
		Key: skey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		return ""
	}
	serv.JKey = stringToken
	return serv.JKey
}
