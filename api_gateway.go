package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gemcook/pagination-go"
)

// NewResponse はレスポンス情報を返す
func NewResponse(body string, statusCode int, optionalHeaders ...map[string]string) (events.APIGatewayProxyResponse, error) {

	// ヘッダーを設定
	headers := NewResponseHeaders(statusCode)
	for _, h := range optionalHeaders {
		headers = merge(headers, h)
	}

	// エラーの場合のレスポンスフォーマットに合わせる
	if statusCode >= 400 {
		body = NewErrorResponseBody(body, statusCode)
	}

	// カスタムヘッダのヘッダを作成
	customHeaders(headers)

	return events.APIGatewayProxyResponse{
		Body:       body,
		Headers:    headers,
		StatusCode: statusCode,
	}, nil
}

// NewErrorResponse はエラーレスポンスを返す
func NewErrorResponse(err error, statusCode int, optionalHeaders ...map[string]string) (events.APIGatewayProxyResponse, error) {

	// ヘッダーを設定
	headers := NewResponseHeaders(statusCode)
	for _, h := range optionalHeaders {
		headers = merge(headers, h)
	}

	// カスタムヘッダのヘッダを作成
	customHeaders(headers)

	return events.APIGatewayProxyResponse{
		Body:       NewErrorResponseBody(err.Error(), statusCode),
		Headers:    headers,
		StatusCode: statusCode,
	}, nil
}

// NewPagingResponse creates 200 response for paging.
func NewPagingResponse(res *pagination.PagingResponse, totalCount, pageCount int) (events.APIGatewayProxyResponse, error) {
	headers := NewResponseHeaders(http.StatusOK)
	TotalCountHeaders(totalCount, headers)
	PageCountHeaders(pageCount, headers)

	// カスタムヘッダのヘッダを作成
	customHeaders(headers)

	b, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		Headers:    headers,
		StatusCode: http.StatusOK,
	}, nil
}

// TotalCountHeaders は取得件数を返す
func TotalCountHeaders(count int, h map[string]string) {
	h["X-Total-Count"] = strconv.Itoa(count)
	customHeaders(h)
}

// PageCountHeaders はページ数を返す
func PageCountHeaders(count int, h map[string]string) {
	h["X-Total-Pages"] = strconv.Itoa(count)
	customHeaders(h)
}

func merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}
	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

func customHeaders(h map[string]string) {
	buf := []byte{}

	for key, value := range h {
		if key[0] == 'X' && value != "" {
			if len(buf) > 0 {
				buf = append(buf, ',', ' ')
			}
			buf = append(buf, key...)
		}
	}
	h["Access-Control-Expose-Headers"] = string(buf)
}

// NewResponseHeaders はステータスに応じたレスポンスヘッダを返す
func NewResponseHeaders(statusCode int) map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Access-Control-Allow-Origin"] = "*"
	switch statusCode {
	case http.StatusOK:
		headers["Cache-Control"] = "max-age=" + os.Getenv("MAX_AGE")
	case http.StatusBadRequest:
		headers["Cache-Control"] = "no-cache"
	}
	return headers
}

// NewErrorResponseBody は ResponseBody を作成する
func NewErrorResponseBody(s string, statusCode int) string {
	// 成功の場合は何もしない
	if statusCode == http.StatusOK {
		return s
	}
	errorResponse := &ErrorResponse{
		Message: s,
	}
	jsonString, err := json.Marshal(errorResponse)
	if err != nil {
		fmt.Printf("json.Marshal error: %v", err)
		return ""
	}
	body := string(jsonString)
	return body
}

// ConvertQueryStringToStruct converts
// events.APIGatewayProxyRequest.QueryStringParameters into given struct.
func ConvertQueryStringToStruct(qs map[string]string, s interface{}) error {
	jsonStr, err := json.Marshal(qs)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonStr, s)
}

// ConvertQueryStringToPaginationQuery converts
// events.APIGatewayProxyRequest.QueryStringParameters into pagination.Query.
func ConvertQueryStringToPaginationQuery(qs map[string]string) *pagination.Query {
	type rangeQuery struct {
		Limit string `json:"limit"`
		Page  string `json:"page"`
	}

	r := rangeQuery{}

	ConvertQueryStringToStruct(qs, &r)

	orders := []*pagination.Order{}

	if sort, ok := qs["sort"]; ok {
		orders = pagination.ParseOrders(sort)
	}

	limit, _ := strconv.Atoi(r.Limit)
	page, _ := strconv.Atoi(r.Page)

	q := &pagination.Query{
		Limit: limit,
		Page:  page,
		Sort:  orders,
	}
	return q
}
