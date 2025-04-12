package repo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/yikuanzz/unitest/entity"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

type TestDBSetting struct {
	Driver       string
	ImageName    string
	ImageVersion string
	ENV          []string
	PortID       string
	Connection   string
}

var (
	mysqlDBSetting = TestDBSetting{
		Driver:       string(schemas.MYSQL),
		ImageName:    "mariadb",
		ImageVersion: "10.4.7",
		ENV:          []string{"MYSQL_ROOT_PASSWORD=root", "MYSQL_DATABASE=test", "MYSQL_ROOT_HOST=%"},
		PortID:       "3306/tcp",
		Connection:   "root:root@(localhost:%s)/test?parseTime=true",
	}
	tearDown       func()
	testDataSource *xorm.Engine
)

func TestMain(t *testing.M) {
	defer func() {
		if tearDown != nil {
			tearDown()
		}
	}()
	if err := initTestDataSource(mysqlDBSetting); err != nil {
		panic(err)
	}
	if ret := t.Run(); ret != 0 {
		panic(ret)
	}
}

func Test_userRepo_AddUser(t *testing.T) {
	ur := NewUserRepo(testDataSource)
	user := &entity.User{Name: "test", Email: "test@test.com", Password: "test"}
	err := ur.AddUser(context.TODO(), user)
	assert.NoError(t, err)

	dbUser, exist, err := ur.GetUser(context.TODO(), user.ID)
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, user.Name, dbUser.Name)

	err = ur.DelUser(context.TODO(), user.ID)
	assert.NoError(t, err)
}

func initTestDataSource(dbSetting TestDBSetting) (err error) {
	connection, imageCleanUp, err := initDatabaseImage(dbSetting)
	if err != nil {
		return err
	}
	dbSetting.Connection = connection

	testDataSource, err = initDatabase(dbSetting)
	if err != nil {
		return err
	}

	tearDown = func() {
		testDataSource.Close()
		imageCleanUp()
	}
	return nil
}

func initDatabaseImage(dbSetting TestDBSetting) (connection string, cleanup func(), err error) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 5
	if err != nil {
		return "", nil, fmt.Errorf("could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: dbSetting.ImageName,
		Tag:        dbSetting.ImageVersion,
		Env:        dbSetting.ENV,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return "", nil, fmt.Errorf("could not pull resource: %s", err)
	}

	connection = fmt.Sprintf(dbSetting.Connection, resource.GetPort(dbSetting.PortID))
	if err := pool.Retry(func() error {
		db, err := sql.Open(dbSetting.Driver, connection)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return db.Ping()
	}); err != nil {
		return "", nil, fmt.Errorf("could not connect to database: %s", err)
	}
	return connection, func() { _ = pool.Purge(resource) }, nil
}

func initDatabase(dbSetting TestDBSetting) (dbEngine *xorm.Engine, err error) {
	dbEngine, err = xorm.NewEngine(dbSetting.Driver, dbSetting.Connection)
	if err != nil {
		return nil, err
	}
	err = initDatabaseData(dbEngine)
	if err != nil {
		return nil, fmt.Errorf("init database data failed: %s", err)
	}
	return dbEngine, nil
}

func initDatabaseData(dbEngine *xorm.Engine) error {
	return dbEngine.Sync(new(entity.User))
}
