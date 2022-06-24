package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/stepan2volkov/user-group-srv/internal/api/openapi"
	"github.com/stepan2volkov/user-group-srv/internal/app/groupapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/userapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
)

var _ openapi.ServerInterface = &Router{}

type VersionInfo struct {
	BuildCommit string
	BuildTime   string
}

type Router struct {
	http.Handler
	version      VersionInfo
	userApp      *userapp.App
	groupApp     *groupapp.App
	userGroupApp *usergroupapp.App
}

func New(
	version VersionInfo,
	userApp *userapp.App,
	groupApp *groupapp.App,
	userGroupApp *usergroupapp.App,
) *Router {
	r := chi.NewRouter()

	rt := &Router{
		version:      version,
		userApp:      userApp,
		groupApp:     groupApp,
		userGroupApp: userGroupApp,
	}

	r.Mount("/", openapi.Handler(rt))
	r.Get("/__version__", rt.versionHandler)
	r.Get("/__heartbeat__", rt.heartbeatHandler)

	rt.Handler = r

	return rt
}

func (rt Router) versionHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&rt.version); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (rt Router) heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
