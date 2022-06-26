package test_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"reflect"
	"shape/entities"
	"strings"
	"testing"
)

var (
	loginJSON = `
{"username":"dung",
"password":"dung123456"}
`
	baseURL = `http://localhost:8800/api/shape/v1/`
)

func Login() (token string, err error) {
	rs := make(map[string]interface{})
	request, _ := http.NewRequest(http.MethodPost, baseURL+"login", strings.NewReader(loginJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	json.Unmarshal(body, &rs)
	v := reflect.ValueOf(rs["data"])
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			token = strct.Interface().(string)
		}
	}
	return token, nil
}

func TestCreateUser(t *testing.T) {
	user := entities.User{
		Username: "dung2",
		Password: "dung123456",
		Fullname: "Tran Hoang Dung",
		Email:    "admin@mail.com",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	request, _ := http.NewRequest(http.MethodPost, baseURL+"register", &buf)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestGetUser(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	request, _ := http.NewRequest(http.MethodGet, baseURL+"user", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestUpdateUser(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	newUser := entities.User{
		Fullname: "Dung updated",
		Email:    "dung@gmail.com",
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(newUser)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	request, _ := http.NewRequest(http.MethodPut, baseURL+"user/6c8df72c-21e1-41c6-ac5a-5e4cd51d6b39", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestDeleteUser(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	request, _ := http.NewRequest(http.MethodDelete, baseURL+"user/6c8df72c-21e1-41c6-ac5a-5e4cd51d6b39", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPostTriangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	triangle := entities.Triangle{
		FirstSide:  3,
		SecondSide: 4,
		ThirdSide:  6,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(triangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPost, baseURL+"triangle", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestPutTriangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	triangle := entities.Triangle{
		FirstSide:  2,
		SecondSide: 3,
		ThirdSide:  4,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(triangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPut, baseURL+"triangle", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}
func TestDeleteTriangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodDelete, baseURL+"triangle", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestGetTriangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	triangle := entities.Triangle{
		FirstSide:  3,
		SecondSide: 4,
		ThirdSide:  6,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(triangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"triangle", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

	//get info
	request, _ := http.NewRequest(http.MethodGet, baseURL+"triangle", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)
	//area
	request2, _ := http.NewRequest(http.MethodGet, baseURL+"triangle/area", nil)
	request2.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp2, err := http.DefaultClient.Do(request2)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body2)
	assert.Equal(t, 200, resp2.StatusCode)
	//area
	request3, _ := http.NewRequest(http.MethodGet, baseURL+"triangle/perimeter", nil)
	request3.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp3, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body3)
	assert.Equal(t, 200, resp2.StatusCode)

}
func TestPostSquare(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	square := entities.Square{
		Length: 10,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(square)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"square", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

}

func TestPutSquare(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	square := entities.Square{
		Length: 20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(square)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPut, baseURL+"square", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}
func TestDeleteSquare(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodDelete, baseURL+"square", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestGetSquare(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	square := entities.Square{
		Length: 10,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(square)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"square", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

	//get info
	request, _ := http.NewRequest(http.MethodGet, baseURL+"square", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)
	//area
	request2, _ := http.NewRequest(http.MethodGet, baseURL+"square/area", nil)
	request2.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp2, err := http.DefaultClient.Do(request2)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body2)
	assert.Equal(t, 200, resp2.StatusCode)
	//area
	request3, _ := http.NewRequest(http.MethodGet, baseURL+"square/perimeter", nil)
	request3.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp3, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body3)
	assert.Equal(t, 200, resp2.StatusCode)

}
func TestPostDiamond(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	diamond := entities.Diamond{
		Length: 10,
		Height: 10,
		Side:   20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(diamond)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"diamond", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

}

func TestPutDiamond(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	diamond := entities.Diamond{
		Length: 12,
		Height: 12,
		Side:   22,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(diamond)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPut, baseURL+"diamond", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}
func TestDeleteDiamond(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodDelete, baseURL+"diamond", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestGetDiamond(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	diamond := entities.Diamond{
		Length: 10,
		Height: 10,
		Side:   20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(diamond)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"diamond", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

	//get info
	request, _ := http.NewRequest(http.MethodGet, baseURL+"diamond", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)
	//area
	request2, _ := http.NewRequest(http.MethodGet, baseURL+"diamond/area", nil)
	request2.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp2, err := http.DefaultClient.Do(request2)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body2)
	assert.Equal(t, 200, resp2.StatusCode)
	//area
	request3, _ := http.NewRequest(http.MethodGet, baseURL+"diamond/perimeter", nil)
	request3.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp3, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body3)
	assert.Equal(t, 200, resp2.StatusCode)

}
func TestPostRectangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rectangle := entities.Rectangle{
		Length: 10,
		Width:  20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(rectangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"rectangle", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

}

func TestPutRectangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rectangle := entities.Rectangle{
		Length: 10,
		Width:  20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(rectangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPut, baseURL+"rectangle", &buf)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}
func TestDeleteRectangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodDelete, baseURL+"rectangle", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)

}

func TestGetRectangle(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rectangle := entities.Rectangle{
		Length: 10,
		Width:  20,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(rectangle)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	requestPost, _ := http.NewRequest(http.MethodPost, baseURL+"rectangle", &buf)
	requestPost.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	respPost, err := http.DefaultClient.Do(requestPost)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(bodyPost)
	assert.Equal(t, 200, respPost.StatusCode)

	//get info
	request, _ := http.NewRequest(http.MethodGet, baseURL+"rectangle", nil)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body)
	assert.Equal(t, 200, resp.StatusCode)
	//area
	request2, _ := http.NewRequest(http.MethodGet, baseURL+"rectangle/area", nil)
	request2.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp2, err := http.DefaultClient.Do(request2)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body2)
	assert.Equal(t, 200, resp2.StatusCode)
	//area
	request3, _ := http.NewRequest(http.MethodGet, baseURL+"rectangle/perimeter", nil)
	request3.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	resp3, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	t.Log(body3)
	assert.Equal(t, 200, resp2.StatusCode)

}
