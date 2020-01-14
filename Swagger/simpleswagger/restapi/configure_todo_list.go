// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/swag"
	"io"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"webservice/Swagger/simpleswagger/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"webservice/Swagger/simpleswagger/restapi/operations"
	"webservice/Swagger/simpleswagger/restapi/operations/todos"
)

//go:generate swagger generate server --target ../../Swagger --name TodoList --spec ../swagger.yml

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var items = make(map[int64]*models.Item)
var lastID int64

var itemsLock = &sync.Mutex{}


func deleteItem(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(items, id)
	return nil
}


func addItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	items[newID] = item

	return nil
}

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}


func allItems(since int64, limit int32) (result []*models.Item) {
	result = make([]*models.Item, 0)
	for id, item := range items {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
}

func downloadjenkisnfile(rw http.ResponseWriter){

	// Get the data
	httpclient:=http.Client{}
	url := url.URL{
		Host:"127.0.0.1:8080",
		//Path:"/job/zpyu/job/zpyutest/3/artifact/zpyu2",
		Path:"/job/zpyu/job/zpyutest/3/artifact/*zip*/archive.zip",

		Scheme: "http",
		//http://127.0.0.1:8080/job/zpyu/job/zpyutest/2/artifact/*zip*/archive.zip
	}
	request := http.Request{
		Method: http.MethodGet,
		URL: &url,
		Header:  http.Header{"Authorization":[]string{"Basic YWRtaW46MTFlOTg4NGI5MjI5MTRhNjc0Njk1MjY2N2Y3NjI2YWIyZg=="}},
	}


	resp, err :=httpclient.Do(&request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 这里将data写入了rw中，应该也就是一个reader之类的东西吧
	_, err = io.Copy(rw,resp.Body)

	if err != nil {
		panic(err)
	}
	return

}


func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams) middleware.Responder {
		if err := deleteItem(params.ID); err != nil {
			return todos.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return todos.NewDestroyOneNoContent()
	})

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams) middleware.Responder {
		if err := addItem(params.Body); err != nil {
			return todos.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return todos.NewAddOneCreated().WithPayload(params.Body)
	})

	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
		mergedParams := todos.NewFindTodosParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return todos.NewFindTodosOK().WithPayload(allItems(*mergedParams.Since, *mergedParams.Limit))
	})

	api.GetIDHandler = operations.GetIDHandlerFunc(func(operations.GetIDParams) middleware.Responder{


		httpclient:=http.Client{}
		url := url.URL{
			Host:"127.0.0.1:8080",
			Path:"/job/zpyu/job/zpyutest/3/artifact/*zip*/archive.zip",
			Scheme: "http",
		}
		request := http.Request{
			Method: http.MethodGet,
			URL: &url,
			Header:  http.Header{"Authorization":[]string{"Basic YWRtaW46MTFlOTg4NGI5MjI5MTRhNjc0Njk1MjY2N2Y3NjI2YWIyZg=="}},
		}


		resp, err :=httpclient.Do(&request)
		resp.Header.Set("Content-Disposition", "attachment; filename=zipname.zip")

		resp.Header.Set("Content-Type", "application/zip")
		if err != nil {
			panic(err)
		}

		return operations.NewGetIDOK().WithPayload(resp.Body)

	})


	if api.TodosAddOneHandler == nil {
		api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.AddOne has not yet been implemented")
		})
	}
	if api.TodosDestroyOneHandler == nil {
		api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.DestroyOne has not yet been implemented")
		})
	}
	if api.TodosFindTodosHandler == nil {
		api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.FindTodos has not yet been implemented")
		})
	}
	if api.TodosUpdateOneHandler == nil {
		api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation todos.UpdateOne has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
