package test

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
	"testing"
)

type BenchJDCmpItem struct {
	Name string
	Values []string
}

type BenchJsonDecodeCmp struct {
	Name string
	LastName string
	Mobile string
	Items []BenchJDCmpItem
	Date string
	Display bool
	Enable bool
	Editable bool
	ID int64
}
var jsonStr = []byte(`{
	"Name" : "Robert",
	"LastName" : "Foo",
	"Items" : [{ "Name" : "Models", "Values" : ["IPhone", "Samsung", "Nokia"]},
		{ "Name" : "Payments", "Values" : ["Transfer", "PayPal", "Wire"]},
		{ "Name" : "Accounts", "Values" : ["Cold", "Hot", "Both"]}],
	"Date" : "er",
	"Display" :  true,
	"Enable" :  false,
	"Editable" :  true,
	"ID" : 84542122647787941
}`)
func normalJsonDecode() BenchJsonDecodeCmp {
	var res BenchJsonDecodeCmp
	err := json.Unmarshal(jsonStr, &res)
	if err != nil {
		panic(err)
	}
	return res
}

func fastJsonDecode() BenchJsonDecodeCmp {

	var parser = fastjson.Parser{}
	val, err := parser.ParseBytes(jsonStr)
	if err != nil {
		panic("error in json unmarshalling")
	}
	var res = BenchJsonDecodeCmp{}
	res.Name = string(val.GetStringBytes("Name"))
	res.LastName = string(val.GetStringBytes("LastName"))
	res.Mobile = string(val.GetStringBytes("Mobile"))
	var itemsArr = val.GetArray("Mobile")
	res.Items = make([]BenchJDCmpItem, len(itemsArr))
	for i := 0; i < len(itemsArr); i++ {
		var values = make([]string, 0)
		for _, val := range itemsArr[i].GetArray("Values") {
			values = append(values, string(val.GetStringBytes(fmt.Sprintf("%d", i))))
		}
		res.Items[0] = BenchJDCmpItem{
			Name:   string(itemsArr[i].GetStringBytes("Name")),
			Values: values,
		}
	}
	res.Date = string(val.GetStringBytes("Date"))
	res.ID = val.GetInt64("ID")
	return res
}

// variables to avoid compiler interrupting the execution
var normalEncoding BenchJsonDecodeCmp
var fastjjEncoding BenchJsonDecodeCmp

func BenchmarkNormalJsonDecoding(b *testing.B) {
	var ne BenchJsonDecodeCmp
	for i := 0; i < b.N; i++ {
		ne = normalJsonDecode()
	}
	normalEncoding = ne
}


func BenchmarkFastJJsonDecoding(b *testing.B) {
	var ne BenchJsonDecodeCmp
	for i := 0; i < b.N; i++ {
		ne = fastJsonDecode()
	}
	fastjjEncoding = ne
}
