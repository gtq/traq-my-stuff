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


func ItemList(w http.ResponseWriter, r *http.Request) {
	tableData := EntryList("select * from items")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(tableData); err != nil {
                //if err := json.Marshal(rows); err != nil {
        	panic(err)
        }

}

func ItemDelete(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var itemid int
        var err error
        if itemid, err = strconv.Atoi(vars["itemId"]); err != nil {
                panic(err)
        }
        t := DeleteEntry("delete from items where id=?",itemid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func ItemUpdate(w http.ResponseWriter, r *http.Request) {
        var msgMapTemplate interface{}
        var itemid int
        var err error
        vars := mux.Vars(r)
        if itemid, err = strconv.Atoi(vars["itemId"]); err != nil {
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


	t := UpdateEntry("update items set id=?,name=?,desc=?,currcon=? where id=?", itemid, []string {"newid","name","desc","currcon","id"}, msgMapTemplate.(map[string]interface{}))

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(t); err != nil {
                panic(err)
        }
}


func ItemGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemid int
	var err error
	if itemid, err = strconv.Atoi(vars["itemId"]); err != nil {
		panic(err)
	}

	tableData := ReadEntry("select * from items where id=?",itemid)

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
func ItemCreate(w http.ResponseWriter, r *http.Request) {
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
	t := CreateEntry("insert into items(id,name,desc,currcon) values (?,?,?,?)",[]string {"id","name","desc","currcon"}, msgMapTemplate.(map[string]interface{}))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
