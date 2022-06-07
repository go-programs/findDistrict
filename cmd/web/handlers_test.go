package main

import (
	"io"
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestChumma(t *testing.T){
	rr:= httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/chumma", nil)
	if err != nil	{
		t.Fatal(err)
	}

	chumma(rr, r)
	rs:= rr.Result()

	if rs.StatusCode != http.StatusOK	{
		t.Errorf("got %d, want %d",rs.StatusCode,http.StatusOK )
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil{
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	if string(body) != "Chumma another page"	{
		t.Errorf("got %s, want %s",string(body),"Chumma another page")
	}	
}