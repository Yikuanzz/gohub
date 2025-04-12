package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGetReqCtx(req interface{}, handlerFunc gin.HandlerFunc) (isSuccess bool, resp string) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	encode := structToURLValues(req).Encode()
	c.Request, _ = http.NewRequest("GET", "/?"+encode, nil)
	handlerFunc(c)
	return w.Code == http.StatusOK, w.Body.String()
}

func CreatePostReqCtx(req interface{}, handlerFunc gin.HandlerFunc) (isSuccess bool, resp string) {
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	body, _ := json.Marshal(req)
	ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBuffer(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	handlerFunc(ctx)
	return responseRecorder.Code == http.StatusOK, responseRecorder.Body.String()
}

func structToURLValues(s interface{}) (values url.Values) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	values = url.Values{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("form")
		if tag == "" {
			continue
		}
		value := v.Field(i).Interface()
		values.Set(tag, valueToString(value))
	}

	return values
}

func valueToString(v interface{}) string {
	switch v := v.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return ""
	}
}
