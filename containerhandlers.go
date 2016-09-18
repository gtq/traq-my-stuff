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


func ContainerList(w http.ResponseWriter, r *http.Request) {
	tableData := EntryList("select * from containers")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(tableData); err != nil {
                //if err := json.Marshal(rows); err != nil {
        	panic(err)
        }

}

func ContainerDelete(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var containerid int
        var err error
        if containerid, err = strconv.Atoi(vars["containerId"]); err != nil {
                panic(err)
        }
        t := DeleteEntry("delete from containers where id=?",containerid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func ContainerUpdate(w http.ResponseWriter, r *http.Request) {
        var msgMapTemplate interface{}
        var containerid int
        var err error
        vars := mux.Vars(r)
        if containerid, err = strconv.Atoi(vars["containerId"]); err != nil {
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


	t := UpdateEntry("update containers set id=?,name=?,desc=?,currloc=? where id=?", containerid, []string {"newid","name","desc","currloc","id"}, msgMapTemplate.(map[string]interface{}))

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(t); err != nil {
                panic(err)
        }
}


func ContainerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var containerid int
	var err error
	if containerid, err = strconv.Atoi(vars["containerId"]); err != nil {
		panic(err)
	}

	tableData := ReadEntry("select * from containers where id=?",containerid)

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
func ContainerCreate(w http.ResponseWriter, r *http.Request) {
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
	t := CreateEntry("insert into containers(id,name,desc,currloc) values (?,?,?,?)",[]string {"id","name","desc","currloc"}, msgMapTemplate.(map[string]interface{}))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
