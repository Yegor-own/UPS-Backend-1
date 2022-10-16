package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"ups/models"
)

func EncodeController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var encodeData models.EncodeData
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("smth went wrong"))
		fmt.Println(err)
		return
	}

	json.Unmarshal(req, &encodeData)
	fmt.Println(encodeData)

	models.RotReportWrite(encodeData.Rot, app)

	var encodedAlphabet models.EncodeResult
	encodedAlphabet.Message = models.EncodeAlphabet(encodeData.Message, encodeData.Rot)
	res, err := json.Marshal(encodedAlphabet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("smth went wrong"))
		fmt.Println(err)
		return
	}
	w.Write(res)
}

func DecodeController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var encodedAlphabet models.EncodeResult
	rot, err := strconv.Atoi(r.FormValue("rot"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("smth went wrong"))
		fmt.Println(err)
		return
	}

	encodedAlphabet.Message = models.DecodeAlphabet(r.FormValue("message"), rot)
	res, err := json.Marshal(encodedAlphabet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("smth went wrong"))
		fmt.Println(err)
		return
	}
	w.Write(res)
}

func StatsController(w http.ResponseWriter, r *http.Request) {
	stats := models.GetStats(app)

	res, err := json.Marshal(stats)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("smth went wrong"))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
