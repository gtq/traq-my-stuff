package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)


func LocationList(w http.ResponseWriter, r *http.Request) {
	tableData := EntryList("select * from locations")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(tableData); err != nil {
                //if err := json.Marshal(rows); err != nil {
        	panic(err)
        }

}

func LocationDelete(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var locationid int
        var err error
        if locationid, err = strconv.Atoi(vars["locationId"]); err != nil {
                panic(err)
        }
        t := DeleteEntry("delete from locations where id=?",locationid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func LocationUpdate(w http.ResponseWriter, r *http.Request) {
        var msgMapTemplate interface{}
        var locationid int
        var err error
        vars := mux.Vars(r)
        if locationid, err = strconv.Atoi(vars["locationId"]); err != nil {
                panic(err)
        }
 
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
        if err != nil {
                panic(err)
        }
        if err := r.Body.Close(); err != nil {
                panic(err)
        }
        if err := json.Unmarshal(body, &msgMapTemplate); err != nil {
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(422) // unprocessable entity
                if err := json.NewEncoder(w).Encode(err); err != nil {
                        panic(err)
                }
        }


	t := UpdateEntry("update locations set id=?,name=?,desc=? where id=?", locationid, []string {"newid","name","desc","id"}, msgMapTemplate.(map[string]interface{}))

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(t); err != nil {
                panic(err)
        }
}


func LocationGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var locationid int
	var err error
	if locationid, err = strconv.Atoi(vars["locationId"]); err != nil {
		panic(err)
	}

	tableData := ReadEntry("select * from locations where id=?",locationid)

	//todo := RepoFindTodo(todoId)
	//if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(tableData); err != nil {
		//if err := json.Marshal(rows); err != nil {
			panic(err)
		}
		return
	//}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func LocationCreate(w http.ResponseWriter, r *http.Request) {
	var msgMapTemplate interface{}
 
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &msgMapTemplate); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := CreateEntry("insert into locations(id,name,desc) values (?,?,?)",[]string {"id","name","desc"}, msgMapTemplate.(map[string]interface{}))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
