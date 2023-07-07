package datalayer

import (
	"database/sql"
	"log"
	"time"
)

// Service stores information about a service
type ServiceInfo struct {
	BServ Service
	SServ StructureServices
}

// Service stores information about a service
type Service struct {
	Sid     sql.NullInt32
	Name    sql.NullString
	Created sql.NullTime
	Credit  sql.NullInt32
	HDcount sql.NullInt32
	Deleted sql.NullBool
	Type    sql.NullBool
	UID     sql.NullInt32
}

// ServiceClusters stores information about a key service
type ServiceClusters struct {
	Clusters sql.NullString //json
}

// GetServiceByID returns a service with the given service ID.
func (handler *MyDB) GetServiceByID(serviceID int, user UserLogin) (Service, ServiceClusters, error) {
	row := handler.db.QueryRow("select sid,name,created,credit,hdcount,deleted,type,uid,clusters from clustream_sru.services where sid=? and uid=?", serviceID, user.ID)

	return getRowDataService(row)
}

// GetUserAllServices returns the all services stored in the database.
func (handler *MyDB) GetUserAllServices(user uint) ([]ServiceInfo, error) {
	rows, err := handler.db.Query("select services.sid,services.name,services.created,services.credit,services.deleted,services.type,"+
		" strservices.expiration,strservices.scheduled,strservices.notused,strservices.status"+
		" from clustream_sru.services left join clustream_sru.strservices on services.sid=strservices.sid"+
		" where services.uid=?", user)
	if err != nil {
		// panic(err)
		return nil, err
	}
	defer rows.Close()

	services := []ServiceInfo{}
	for rows.Next() {
		u, err := getRowsDataUserAllServices(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		services = append(services, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return services, nil
}

func getRowsDataUserAllServices(row *sql.Rows) (ServiceInfo, error) {
	u := ServiceInfo{}

	err := row.Scan(&u.BServ.Sid, &u.BServ.Name, &u.BServ.Created, &u.BServ.Credit, &u.BServ.Deleted, &u.BServ.Type, &u.SServ.Expiration, &u.SServ.ExpScheduled, &u.SServ.NotUsed, &u.SServ.Status)
	if err != nil {
		return u, err
	}
	if u.BServ.Type.Bool {
		if u.SServ.Status.Bool {
			use := time.Now().Sub(u.BServ.Created.Time).Hours()
			notuse := u.SServ.NotUsed.Int32
			//u.SServ.Used = int(((int32(time.Now().Sub(u.BServ.Created.Time).Hours()) - u.SServ.NotUsed.Int32) / u.BServ.Credit.Int32) * 100)
			u.SServ.Used = int((use - float64(notuse)) * 100 / float64(u.BServ.Credit.Int32))
		} else {
			u.SServ.Used = int((u.SServ.NotUsed.Int32 * 100 / u.BServ.Credit.Int32))
		}
	}

	return u, nil
}

// InsertService insert a service to database and returns id and err.
func (handler *MyDB) InsertService(serv Service, clu ServiceClusters) (int64, error) {
	res, err := handler.db.Exec("INSERT INTO services(name,created,credit,hdcount,type,uid,clusters) VALUES (?,?,?,?,?,?,?)", serv.Name, serv.Created, serv.Credit, serv.HDcount, serv.Type, serv.UID, clu.Clusters)

	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()

	if err != nil {
		return -2, err
	}
	return lastID, nil
}

func getRowDataService(row *sql.Row) (Service, ServiceClusters, error) {
	u := Service{}
	c := ServiceClusters{}
	err := row.Scan(&u.Sid, &u.Name, &u.Created, &u.Credit, &u.HDcount, &u.Deleted, &u.Type, &u.UID, &c.Clusters)
	if err != nil {
		return u, c, err
	}
	return u, c, nil
}

// GetServiceLogin returns a service Login information with the given service ID that send to jwt
func (handler *MyDB) GetServiceLogin(serviceID string) (ServiceClusters, error) {
	row := handler.db.QueryRow("select clusters from clustream_sru.servicess where sid=?", serviceID)

	return getRowDataServiceClusters(row)
}

func getRowDataServiceClusters(row *sql.Row) (ServiceClusters, error) {
	u := ServiceClusters{}

	err := row.Scan(&u.Clusters)
	if err != nil {
		return u, err
	}
	return u, nil
}
