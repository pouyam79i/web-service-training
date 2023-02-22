package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const keyServerAddr = "cal_api.ir"

// errors
var errNoOps = errors.New("No Operation!")
var errBadStr = errors.New("Bad String")
var errZeroDiv = errors.New("Zero Division!")

type jsonRes struct {
	Stat      string    `json:"stat"`
	Res       string    `json:"res"`
	CreatedAt time.Time `json:"createdAt"`
}

// split by operation and convert to float64
func splitByFirstCharAndConvertToFloat64(input, subStr string) (float64, float64, error) {
	index := strings.Index(input, subStr)
	len := len(input)
	a := input[0:index]
	b := input[index+1 : len]
	var af, bf float64 = 0.0, 0.0
	var err error = nil
	if af, err = strconv.ParseFloat(a, 64); err != nil {
		return 0, 0, errBadStr
	}
	if bf, err = strconv.ParseFloat(b, 64); err != nil {
		return 0, 0, errBadStr
	}
	return af, bf, nil
}

// calculate the req of user
func calReqParser(input string) (string, error) {
	var res float64 = 0.0
	var err error = nil
	var a, b float64 = 0.0, 0.0
	if strings.Contains(input, "+") {
		a, b, err = splitByFirstCharAndConvertToFloat64(input, "+")
		res = a + b
	} else if strings.Contains(input, "-") {
		a, b, err = splitByFirstCharAndConvertToFloat64(input, "-")
		res = a - b
	} else if strings.Contains(input, "*") {
		a, b, err = splitByFirstCharAndConvertToFloat64(input, "*")
		res = a * b
	} else if strings.Contains(input, "/") {
		a, b, err = splitByFirstCharAndConvertToFloat64(input, "/")
		if b == 0 && err == nil {
			err = errZeroDiv
		} else {
			res = a / b
		}
	} else {
		err = errNoOps
		res = 0.0
	}
	return strconv.FormatFloat(res, 'E', -1, 64), err
}

// build a proper json respond
func buildResJSON(input string) ([]byte, int, error) {
	var err error = nil
	var stat, res string = "0", "0"
	var code int = 0
	res, err = calReqParser(input)
	if err == nil {
		stat = "successful"
		code = http.StatusOK
	} else {
		stat = err.Error()
		if errors.Is(err, errBadStr) {
			code = http.StatusBadRequest
		} else {
			code = http.StatusConflict
		}
	}
	jr := &jsonRes{
		Stat:      stat,
		Res:       res,
		CreatedAt: time.Now(),
	}
	out, o_err := json.MarshalIndent(jr, "", " ")
	if o_err != nil {
		code = http.StatusInternalServerError
		out = nil
	} else {
		fmt.Println("JR OBJ: \n", string(out))
	}
	return out, code, o_err
}

// Root function
func getRoot(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("x-req-error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Server Failed To Read Req")
		return
	}
	input := string(body)
	jr, code, err := buildResJSON(input)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	c, err := w.Write(jr)
	if err == nil {
		err = errors.New("NO-ERROR")
	}
	fmt.Printf("Data Count: %d,\nError: %s\n", c, err.Error())
}

// server builder
func BuildServer(ip, port string) {
	fmt.Println("Server is booting...")
	// attaching handler functions
	mux := http.NewServeMux() // server mux instead of default http handler
	mux.HandleFunc("/", getRoot)

	// setting server config
	ctx := context.Background()
	addr := ip + ":" + port
	serverOne := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	// running server
	fmt.Println("Server One is On!")
	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server One is down!")
		err = nil
	} else if err != nil {
		fmt.Println("Error on server one: ", err)
	}

}

// running application
func main() {
	BuildServer("", "8085")
}
