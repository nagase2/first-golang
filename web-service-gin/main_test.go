package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
}

func (suite *GameTestSuite) BeforeTest(_, _ string) {
	// execute code before test starts
	fmt.Println("BEFORE")
}

func (suite *GameTestSuite) AfterTest(_, _ string) {
	// execute code after test finishes
	fmt.Println("AFTER")
}

func (suite *GameTestSuite) SetupSuite() {
	// create relevant resources
	fmt.Println("SETUP")
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

func TestPingRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestAlbumsRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// assert.JSONEq(w.Body.String(),
	// 	`[    {        "id": "3",        "title": "Sarah Vaughan and Clifford Brown",        "artist": "Sarah Vaughan",        "price": 39.99    }	]`)
}

func TestJSONEq(t *testing.T) {
	expected := `{"name": "Nozomi", "age": 17, "blood": "O"}`
	actual := `{"age": 17, "blood": "O", "name": "Nozomi"}`

	assert.JSONEq(t, expected, actual)
}

func TestAssertion1(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b, "The two words should be the same.")
}

// オブジェクトの比較をしてみる
func TestAssertion2(t *testing.T) {
	t.Run("あああああ", func(t *testing.T) {
		assert := assert.New(t)

		// JSONが同じかどうか検証
		assert.JSONEq(`[{"aa": "dd" } ] `, `[{"aa": "dd"}]`)

		//assert.ObjectsAreEqualValues(obj, obj2, "The two words should be the same.")
	})
}

// オブジェクトの部分一致比較をする
func (suite *GameTestSuite) TestAssertionPertialCompare() {
	suite.T().Run("🐮sad path: invalid stats", func(t *testing.T) {
		assert := assert.New(t)

		// JSONが同じかどうか検証
		assert.JSONEq(`[{"aa": "dd" } ] `, `[{"aa": "dd"}]`)

		//assert.Contains(t, "Hello World", "World")

		//assert.Contains(t, `{"Hello": "World"}`, "Hello")
	})
	suite.T().Run("🐵sad2path: invalid stats", func(t *testing.T) {
		assert := assert.New(t)

		// JSONが同じかどうか検証
		assert.JSONEq(`[{"aa": "dd" } ] `, `[{"aa": "dd"}]`)

	})
}

// Object完全一致パターン
func TestNewCompanyHandler(t *testing.T) {
	r := SetupRouter()
	// r.POST("/company", NewCompanyHandler)
	req, _ := http.NewRequest("GET", "/albums", nil)
	// companyId := xid.New().String()
	// company := Company{
	// 	ID:      companyId,
	// 	Name:    "Demo Company",
	// 	CEO:     "Demo CEO",
	// 	Revenue: "35 million",
	// }
	// jsonValue, _ := json.Marshal(company)
	// req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))
	mockResponse := `[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]`
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//fmt.Print("🈲🈲")
	// fmt.Println(w.Body)
	responseData, _ := ioutil.ReadAll(w.Body)
	//fmt.Println(responseData)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// レスポンスに特定の文字列を含んでいたら成功
func TestNewCompanyHandler2(t *testing.T) {
	r := SetupRouter()
	// r.POST("/company", NewCompanyHandler)
	req, _ := http.NewRequest("GET", "/albums", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//fmt.Print("🈲🈲")
	// fmt.Println(w.Body)

	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(responseData))
	// こう書くことでレスポンスの中に含まれている文字列のチェックが可能
	assert.Contains(t, string(responseData), "Coltrane")
	assert.Contains(t, "Maki Nishikino", "Maki")
	assert.Contains(t, []string{"Maki", "Hanayo", "Rin"}, "Maki")
	assert.Contains(t, map[string]int{"Maki": 1, "Honoka": 2, "Niko": 3}, "Maki")
	// map の value は対象外
	assert.NotContains(t, map[string]int{"Maki": 1, "Honoka": 2, "Niko": 3}, 3)
}

// 特定のメソッドをMockする
