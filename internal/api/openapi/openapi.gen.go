// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// Defines values for GroupType.
const (
	Community      GroupType = "community"
	CorporateGroup GroupType = "corporate_group"
	Org            GroupType = "org"
	Project        GroupType = "project"
)

// Group defines model for Group.
type Group struct {
	CreatedAt *string             `json:"created_at,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Title     string              `json:"title"`
	Type      GroupType           `json:"type"`
}

// GroupType defines model for Group.Type.
type GroupType string

// User defines model for User.
type User struct {
	CreatedAt *string             `json:"created_at,omitempty"`
	Email     string              `json:"email"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Nickname  string              `json:"nickname"`
}

// UserGroup defines model for UserGroup.
type UserGroup struct {
	GroupId openapi_types.UUID `json:"group_id"`
	UserId  openapi_types.UUID `json:"user_id"`
}

// CreateGroupJSONBody defines parameters for CreateGroup.
type CreateGroupJSONBody = Group

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody = User

// AddUserToGroupJSONBody defines parameters for AddUserToGroup.
type AddUserToGroupJSONBody = UserGroup

// CreateGroupJSONRequestBody defines body for CreateGroup for application/json ContentType.
type CreateGroupJSONRequestBody = CreateGroupJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUserJSONBody

// AddUserToGroupJSONRequestBody defines body for AddUserToGroup for application/json ContentType.
type AddUserToGroupJSONRequestBody = AddUserToGroupJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Group
	// (POST /group)
	CreateGroup(w http.ResponseWriter, r *http.Request)
	// Getting info about group by its title
	// (GET /group/title/{title})
	FindGroupByTitle(w http.ResponseWriter, r *http.Request, title string)
	// Getting info about group
	// (GET /group/{groupID})
	GetGroupByID(w http.ResponseWriter, r *http.Request, groupID openapi_types.UUID)
	// Create User
	// (POST /user)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Getting info about user by his/her nickname
	// (GET /user/nickname/{nickname})
	FindUserByNickname(w http.ResponseWriter, r *http.Request, nickname string)
	// Getting info about user
	// (GET /user/{userID})
	GetUserByID(w http.ResponseWriter, r *http.Request, userID openapi_types.UUID)
	// Add user to group
	// (POST /usergroup)
	AddUserToGroup(w http.ResponseWriter, r *http.Request)
	// Getting user groups
	// (GET /usergroup/groups/{userID})
	FindGroupsByUserID(w http.ResponseWriter, r *http.Request, userID openapi_types.UUID)
	// Getting users in group
	// (GET /usergroup/users/{groupID})
	FindUsersByGroupID(w http.ResponseWriter, r *http.Request, groupID openapi_types.UUID)
	// Drop user from group
	// (DELETE /usergroup/{userID}/{groupID})
	DropUserFromGroup(w http.ResponseWriter, r *http.Request, userID openapi_types.UUID, groupID openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// CreateGroup operation middleware
func (siw *ServerInterfaceWrapper) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateGroup(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindGroupByTitle operation middleware
func (siw *ServerInterfaceWrapper) FindGroupByTitle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "title" -------------
	var title string

	err = runtime.BindStyledParameter("simple", false, "title", chi.URLParam(r, "title"), &title)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "title", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindGroupByTitle(w, r, title)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetGroupByID operation middleware
func (siw *ServerInterfaceWrapper) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "groupID" -------------
	var groupID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "groupID", chi.URLParam(r, "groupID"), &groupID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "groupID", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGroupByID(w, r, groupID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindUserByNickname operation middleware
func (siw *ServerInterfaceWrapper) FindUserByNickname(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "nickname" -------------
	var nickname string

	err = runtime.BindStyledParameter("simple", false, "nickname", chi.URLParam(r, "nickname"), &nickname)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nickname", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindUserByNickname(w, r, nickname)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUserByID operation middleware
func (siw *ServerInterfaceWrapper) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "userID", chi.URLParam(r, "userID"), &userID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserByID(w, r, userID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AddUserToGroup operation middleware
func (siw *ServerInterfaceWrapper) AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddUserToGroup(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindGroupsByUserID operation middleware
func (siw *ServerInterfaceWrapper) FindGroupsByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "userID", chi.URLParam(r, "userID"), &userID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindGroupsByUserID(w, r, userID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindUsersByGroupID operation middleware
func (siw *ServerInterfaceWrapper) FindUsersByGroupID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "groupID" -------------
	var groupID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "groupID", chi.URLParam(r, "groupID"), &groupID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "groupID", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindUsersByGroupID(w, r, groupID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DropUserFromGroup operation middleware
func (siw *ServerInterfaceWrapper) DropUserFromGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "userID", chi.URLParam(r, "userID"), &userID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	// ------------- Path parameter "groupID" -------------
	var groupID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "groupID", chi.URLParam(r, "groupID"), &groupID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "groupID", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DropUserFromGroup(w, r, userID, groupID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/group", wrapper.CreateGroup)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/group/title/{title}", wrapper.FindGroupByTitle)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/group/{groupID}", wrapper.GetGroupByID)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/user", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user/nickname/{nickname}", wrapper.FindUserByNickname)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user/{userID}", wrapper.GetUserByID)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/usergroup", wrapper.AddUserToGroup)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/usergroup/groups/{userID}", wrapper.FindGroupsByUserID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/usergroup/users/{groupID}", wrapper.FindUsersByGroupID)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/usergroup/{userID}/{groupID}", wrapper.DropUserFromGroup)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xXTW8bNxD9K8S0x4VWcdKLTo1qRDAQ9NDap8YwqN2RxFhLbodcoQtB/70Ycj/0sZI2",
	"sWXrEnGZIYd8b97jeA2JyXKjUTsLozXYZIGZ9MMJmSLnQU4mR3IK/XRCKB2mT9Lx18xQxiNIpUOnMoQI",
	"XJkjjMA6UnoOmwhUuhNaFCrtCnPKLZEjD//HT6wBdZHB6B8+0XdMHERgaA4RJIZyQ9Lh09wfmmeyrNDK",
	"lfB4kGkTAeG/hSJMebOQtopqo83Up9hE8GCRXgUHzKRadl6wJ0JaJc9aZl0g7V2qiayzHrvYEZY9jk89",
	"j1VYpH6xe6esF0ZtvsNz8iKlZ8aDbrSTiUe8AhOsw1zqm5VZPpvV76XUKf43oILPlaJNSOVOGQ0juF8o",
	"K5QVboGiKiAxMyQmiM9jkkpbkZiCLIpvMDFiLJNn1Km4xRUuTZ6hdgPxlT/EzTeAplw9isLDKP5GWqmE",
	"QV8h2ZD2w2A4GPJpTI5a5gpG8HEwHHyECHLpFh7teN6wYKy/HXMh+eB3KYzgD19rk6q2GUG0bmzSssYE",
	"tV8l83ypEr8u/m45fa1oHv1KOIMR/BK3ko8rvcdh780uQY4K9BM2N9qGwrgZfrhE0l2qPByiUhhj92k4",
	"5M12o6YyFRUWHPNbV4zSDknLpbBIKySBRIZ8HdoiyySVDbqihtfJueXiDN+PHBv4iT3j8dr/bLxKsIOs",
	"L0qnfu24vK+sJZckM3RIvPG687K1Cyme4sKACILUG4PaJSbaAnlfZY8HpA3firSZKXRF2adDOkKMNq6N",
	"+2naJuic0nPB3iDk1BROhO2npVDONoge53Ptf+5uj3M5QVdReXd7jke7MOTEw19fvasQpoow4RN2s1rl",
	"PsnrOTe9Dp7PS/PdauEI/UX9op9wW//qX8Zs/dZv7LVtzl10GYq3dtoK2poZ/9kSE9e9S7yuR6fNlteP",
	"yz/bjmdPph3i22qPrsNVT9Jz2lN9yMVk5HeflmKhbLzgRC1yx/hb879nXDVw1mWqHWyF/a7WKfty95M+",
	"eXl+j3N5pi/9nHrx3ZtLtqbtXyn9LXPvZS6SBK2dFUvRnP8o3K+D9Oc0DdJx5uAlai+0B3RoS+x5ATUt",
	"ph2XD7U4XtZkXoHGlMPM9mxLmuySSJav1Ka8WFee8UBiL755ZHu0ofUrZ8flpGkbX8b3NbSfvQgP7nq9",
	"fFuh9A9IvNb2LuspLtHhIfG3ZHLe6wuZrPbYd9d5dE3F1tfpL1gNTFKQ/oxMdroWNs38fqY7HRBQRm89",
	"zbYFsRLC+XWNAVULa8vsl3F/9dbz+7j5PwAA///+LP2GLxYAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
