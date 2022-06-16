package products

import (
	"github.com/lordvidex/go-example-server/internal/common/data"
	"strings"
	"testing"
)

type testCase struct {
	obj        data.Codable
	json       string
	shouldFail bool
	name       string
}

var testCases = []testCase{
	{
		obj:        &Product{Id: 1, Name: "test", Description: "test"},
		json:       `{"id":1,"name":"test","description":"test"}`,
		shouldFail: false,
		name:       "correct product with id 1",
	},
	{
		obj: &Product{
			Id:          2,
			Name:        "test2",
			Description: "this is test2",
		},
		json:       `{"id":2,"name":"test2","description":"this is test2"}`,
		shouldFail: false,
		name:       "correct product with id 2",
	},
	{
		obj: &Product{
			Id:          2,
			Name:        "test2",
			Description: "this is test2",
		},
		json:       `{"id":2,"name":"test2_check","description":"this is test2"}`,
		shouldFail: true,
		name:       "wrong name",
	},
	{
		obj: &Product{
			Id:          2,
			Name:        "test2",
			Description: "this is test2",
		},
		json:       `{"id":3,"name":"test2","description":"this is test2"}`,
		shouldFail: true,
		name:       "wrong id",
	},
	{
		obj: &Product{
			Id:          2,
			Name:        "test2",
			Description: "this is test2",
		},
		json:       `{"id":2,"name":"test2","description":"this is test45353"}`,
		shouldFail: true,
		name:       "wrong description",
	},
}

func TestProduct_ToJSON(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder := strings.Builder{}
			err := tc.obj.ToJSON(&builder)
			if err != nil && !tc.shouldFail {
				if tc.shouldFail {
					return
				}
				t.Error("test should not fail")
			}
			if cleanString(builder.String()) != cleanString(tc.json) && !tc.shouldFail {
				t.Errorf("expected %s, got %s", tc.json, builder.String())
			}
		})
	}
}

func TestProduct_FromJSON(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.json)
			obj := &Product{}
			err := obj.FromJSON(reader)
			if err != nil {
				if tc.shouldFail {
					return
				}
				t.Errorf("test should not fail got error: %s", err.Error())
			}
			if obj == tc.obj && !tc.shouldFail {
				t.Errorf("expected %v, got %v", tc.obj, obj)
			}
		})
	}
}

func cleanString(str string) string {
	return strings.TrimSpace(str)
}
