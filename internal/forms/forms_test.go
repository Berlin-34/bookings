package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "a")
	r.PostForm = postedData
	form := New(r.PostForm)

	if !form.Has("a") {
		t.Error("Has returns false when it should return true")
	}
	if form.Has("d") {
		t.Error("Has returns true when it should return false")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("a", "four")
	postedData.Add("b", "five")
	r.PostForm = postedData
	form := New(r.PostForm)

	valid := form.MinLength("a", 3)
	if !valid {
		t.Error("string greater than minimum length but returning not valid")
	}
	valid = form.MinLength("b", 5)
	if valid {
		t.Error("string lesser than minimum length but returning valid")
	}
	isErr := form.Errors.Get("b")
	if isErr == "" {
		t.Error("should have an error, but did not get one")
	}

	isErr = form.Errors.Get("a")
	if isErr != "" {
		t.Error("should not have an error, but did get one")
	}

	if form.Valid() {
		t.Error("form is valid when should've been invalid")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("a", "valid@email.com")
	postedData.Add("b", "invalid@email")
	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("Email showing invalid when it is valid")
	}

	form.IsEmail("b")
	if form.Valid() {
		t.Error("Email showing valid when it is invalid")
	}
}
