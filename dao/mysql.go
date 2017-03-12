package dao

import (
	"database/sql"

	"web-service/model"

	"github.com/ngaut/log"
)

const (
	_getSecUserByNameSQL = "SELECT id,name,version,mark,operator,last_login_date,ctime,mtime,status FROM sec_users WHERE name like ? "
)

func (d *Dao) SecUser(name string) (secUser *model.SecUser, err error) {
	//预编译语句时采取该方式
	var row = d.getSecUserStmt.QueryRow(name)
	//非预编译方式
	//var row = d.db.QueryRow(_getSecUserByNameSQL, name)
	secUser = new(model.SecUser)
	if err = row.Scan(&secUser.Id, &secUser.Name, &secUser.Version,
		&secUser.Mark, &secUser.Operator, &secUser.Last_login_date,
		&secUser.Ctime, &secUser.Mtime, &secUser.Status); err != nil {
		if err == sql.ErrNoRows {
			secUser = nil
			err = nil
			return
		}
		log.Errorf("row.Scan() error %v", err)
	}
	return
}
