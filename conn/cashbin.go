package conn

import (
	"Asm/moled"
	"database/sql"
	"runtime"
	"time"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
)

func finalizer(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func AccStatus(user *moled.User,url string)bool {
	// connect to the database first.
	db, err := sql.Open("mysql", "root:Root123456@tcp(127.0.0.1:3306)/new")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	m := user
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	// need to control by user, not the package
	runtime.SetFinalizer(db, finalizer)

	// Initialize an adapter and use it in a Casbin enforcer:
	// The adapter will use the Sqlite3 table name "casbin_rule_test",
	// the default table name is "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := sqladapter.NewAdapter(db, "mysql", "casbin_rule_test")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	if err != nil {
		panic(err)
	}

	ok, err := e.Enforce(m.Name,m.Telephone,url)
	if err!=nil {
		println("err: e.enforce" )
	}
	if ok!=true{
		return false
	}
	// Check the permission.
	return true
}
func CasbinGetPermission(s1 string,s2 string,s3 string)  {
	db, err := sql.Open("mysql", "root:Root123456@tcp(127.0.0.1:3306)/new")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	// need to control by user, not the package
	runtime.SetFinalizer(db, finalizer)

	// Initialize an adapter and use it in a Casbin enforcer:
	// The adapter will use the Sqlite3 table name "casbin_rule_test",
	// the default table name is "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := sqladapter.NewAdapter(db, "mysql", "casbin_rule_test")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	if err != nil {
		panic(err)
	}
	e.AddPolicy(s1,s2,s3)
}