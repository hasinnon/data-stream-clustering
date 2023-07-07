package datalayer

import (
	"database/sql"
	"log"
)

// Admin stores information about a admin
type Admin struct {
	Pid sql.NullString // int
	Nid sql.NullString // int
}

// GetAdminByPID returns a admin with the given admin ID.
func (handler *MyDB) GetAdminByPID(adminID int) (Admin, error) {
	row := handler.db.QueryRow("select pid,nid from emr.admins where pid=?", adminID)

	return getRowDataAdmin(row)
}

// GetAdminByNID returns a admin with the given national ID
func (handler *MyDB) GetAdminByNID(nationalID int) (Admin, error) {
	row := handler.db.QueryRow("select pid,nid from emr.admins where nid=?", nationalID)

	return getRowDataAdmin(row)
}

// GetAdmins returns the all admins stored in the database.
func (handler *MyDB) GetAdmins() ([]Admin, error) {
	rows, err := handler.db.Query("SELECT * FROM `emr`.`admins`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	admins := []Admin{}
	for rows.Next() {
		u, err := getRowsDataAdmin(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		admins = append(admins, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return admins, nil
}

func getRowsDataAdmin(row *sql.Rows) (Admin, error) {
	u := Admin{}

	err := row.Scan(&u.Pid, &u.Nid)
	if err != nil {
		return u, err
	}
	return u, nil
}

func getRowDataAdmin(row *sql.Row) (Admin, error) {
	u := Admin{}

	err := row.Scan(&u.Pid, &u.Nid)
	if err != nil {
		return u, err
	}
	return u, nil
}
