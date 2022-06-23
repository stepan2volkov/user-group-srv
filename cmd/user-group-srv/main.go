package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/stepan2volkov/user-group-srv/internal/api/router"
	"github.com/stepan2volkov/user-group-srv/internal/api/server"
	"github.com/stepan2volkov/user-group-srv/internal/app/groupapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/userapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/config"
	"github.com/stepan2volkov/user-group-srv/internal/store/mem/memgroupstore"
	"github.com/stepan2volkov/user-group-srv/internal/store/mem/memusergroupstore"
	"github.com/stepan2volkov/user-group-srv/internal/store/mem/memuserstore"
	"github.com/stepan2volkov/user-group-srv/internal/store/pg/pggroupstore"
	"github.com/stepan2volkov/user-group-srv/internal/store/pg/pgstarter"
	"github.com/stepan2volkov/user-group-srv/internal/store/pg/pgusergroupstore"
	"github.com/stepan2volkov/user-group-srv/internal/store/pg/pguserstore"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config: %+v\n", cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Init storages
	var userStore userapp.UserProvider
	var groupStore groupapp.GroupProvider
	var userGroupStore usergroupapp.UserGroupProvider

	if strings.HasPrefix(cfg.DSN, "postgres://") {
		log.Println("using postgres stores")
		db, err := pgstarter.NewPGStore(cfg.DSN)
		if err != nil {
			log.Fatal(err)
		}
		userStore = pguserstore.New(db)
		groupStore = pggroupstore.New(db)
		userGroupStore = pgusergroupstore.New(db)
	} else {
		log.Println("using memory stores")
		userStore = memuserstore.New()
		groupStore = memgroupstore.New()
		userGroupStore = memusergroupstore.New()
	}

	log.Println("connection with db has been established")

	// Init applications
	userApp := userapp.New(userStore)
	groupApp := groupapp.New(groupStore)
	userGroupApp := usergroupapp.New(userStore, groupStore, userGroupStore)

	// Version
	version := router.VersionInfo{
		BuildCommit: config.BuildCommit,
		BuildTime:   config.BuildTime,
	}

	// Init server
	rt := router.New(version, userApp, groupApp, userGroupApp)
	srv := server.NewServer(cfg, rt)

	srv.Start()
	log.Println("application has been started")
	<-ctx.Done()
	srv.Stop()
}
