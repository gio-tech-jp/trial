package controller_test

import (
	"encoding/json"
	searchHistoryController "ess/api/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

/**
 * TestSearchHistoryController_Get_001
 * 概要：正常系
 * 条件：正常パターン
 * 期待結果：正常終了（200）
 */
func TestSearchHistoryController_Get_001(t *testing.T) {

	// レスポンス生成
	response := httptest.NewRecorder()

	// コンテキスト作成
	context, _ := gin.CreateTestContext(response)

	// リクエスト設定
	url := "api/v1/searchHistory/getSearchHistory"
	param := "?userId=000000001"

	// リクエスト格納
	context.Request, _ = http.NewRequest(http.MethodGet, url+param, nil)

	// フォーム属性付与
	context.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 実行
	searchHistoryController.GetSearchHistory(context)

	// レスポンス
	var responseBody map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

	// 検証
	assert.Equal(t, response.Code, 200)

}

/**
 * TestSearchHistoryController_Get_101
 * 概要：異常系（validation errorが発生）
 * 条件：userIdがnull
 * 期待結果：異常終了（400）
 */
func TestSearchHistoryController_Get_101(t *testing.T) {

	// レスポンス生成
	response := httptest.NewRecorder()

	// コンテキスト作成
	context, _ := gin.CreateTestContext(response)

	// リクエスト設定
	url := "api/v1/searchHistory/getSearchHistory"
	param := "?userId="

	// リクエスト格納
	context.Request, _ = http.NewRequest(http.MethodGet, url+param, nil)

	// フォーム属性付与
	context.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 実行
	searchHistoryController.GetSearchHistory(context)

	// レスポンス
	var responseBody map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

	// 検証
	assert.Equal(t, response.Code, 400)

}

/**
 * TestSearchHistoryController_Get_102
 * 概要：異常系（validation errorが発生）
 * 条件：userIdが8桁
 * 期待結果：異常終了（400）
 */
func TestSearchHistoryController_Get_102(t *testing.T) {

	// レスポンス生成
	response := httptest.NewRecorder()

	// コンテキスト作成
	context, _ := gin.CreateTestContext(response)

	// リクエスト設定
	url := "api/v1/searchHistory/getSearchHistory"
	param := "?userId=12345678"

	// リクエスト格納
	context.Request, _ = http.NewRequest(http.MethodGet, url+param, nil)

	// フォーム属性付与
	context.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 実行
	searchHistoryController.GetSearchHistory(context)

	// レスポンス
	var responseBody map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

	// 検証
	assert.Equal(t, response.Code, 400)

}

/**
 * TestSearchHistoryController_Get_103
 * 概要：異常系（validation errorが発生）
 * 条件：userIdが10桁
 * 期待結果：異常終了（400）
 */
func TestSearchHistoryController_Get_103(t *testing.T) {

	// レスポンス生成
	response := httptest.NewRecorder()

	// コンテキスト作成
	context, _ := gin.CreateTestContext(response)

	// リクエスト設定
	url := "api/v1/searchHistory/getSearchHistory"
	param := "?userId=1234567890"

	// リクエスト格納
	context.Request, _ = http.NewRequest(http.MethodGet, url+param, nil)

	// フォーム属性付与
	context.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 実行
	searchHistoryController.GetSearchHistory(context)

	// レスポンス
	var responseBody map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

	// 検証
	assert.Equal(t, response.Code, 400)

}
