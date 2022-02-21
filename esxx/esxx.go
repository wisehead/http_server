package esxx

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

/*
type Zhb_search_online_db_Item struct {
	Id     string `json:"id"`
	DataId string `json:"dataId"`

	Item string `json:"item"`

	SourceUrl string `json:"sourceUrl"`

	Category string `json:"title"`

	Title string `json:"title"`

	Content string `json:"content"`

	Author string `json:"author"`

	CreateTime string `json:"createTime"`

	DataType string `json:"dataType"`

	PublishedTime string `json:"publishedTime"`

	Summary string `json:"summary"`

	Ext string `json:"ext"`

	ClassifyName string `json:"classifyName"`

	ClassifyDesc string `json:"classifyDesc"`

	Description string `json:"description"`

	Detail string `json:"detail"`

	NCode string `json:"nCode"`

	Photo string `json:"photo"`
}
*/
type Zhb_search_online_db_Item struct {
	DataID        int       `json:"dataId,omitempty"`
	Author        string    `json:"author,omitempty"`
	Category      string    `json:"category,omitempty"`
	Content       string    `json:"content,omitempty"`
	CreateTime    time.Time `json:"createTime,omitempty"`
	DataType      string    `json:"dataType,omitempty"`
	Item          string    `json:"item,omitempty"`
	Ext           string    `json:"ext,omitempty"`
	ID            int       `json:"id,omitempty"`
	PublishedTime time.Time `json:"publishedTime,omitempty"`
	SourceURL     string    `json:"sourceUrl,omitempty"`
	Summary       string    `json:"summary,omitempty"`
	Title         string    `json:"title,omitempty"`
	Photo         string    `json:"photo,omitempty"`
	Detail        string    `json:"detail,omitempty"`
	Description   string    `json:"description,omitempty"`
	ClassifyName  string    `json:"classifyName,omitempty"`
	ClassifyDesc  string    `json:"classifyDesc,omitempty"`
	NCode         string    `json:"nCode,omitempty"`
}

/*
curl -H "Content-Type:application/json"  -X PUT 'localhost:9200/accounts/person/1' -d '
{
  "user": "lisi",
  "title": "cleaner",
  "desc": "{\"nProductId\": 85, \"cName\": \"xxx\", \"cKind\": \"accident\", \"cDesc\": \"aaaaa\", \"cCode\": \"1949a\"}"
}'
*/
type Accounts struct {
	User  string `json:"user,omitempty"`
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func Test_es() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Person{Name: "lmh", Age: 18, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func Es_search() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 创建term查询条件，用于精确查询
	//termQuery := elastic.NewTermQuery("description", "nProductId")

	searchResult, err := client.Search().
		Index("zhb_search_online_db"). // 设置索引名
		//Query(termQuery).              // 设置查询条件
		Sort("createTime", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).                  // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).                 // 设置分页参数 - 每页大小
		Pretty(true).             // 查询结果返回可读性较好的JSON格式
		Do(ctx)                   // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	fmt.Println(searchResult)

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Zhb_search_online_db_Item
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			/*
				if t, ok := item.(Zhb_search_online_db_Item); ok {
					//fmt.Println(t.Title)
					//fmt.Println(t)
				}
			*/
			t, _ := item.(Zhb_search_online_db_Item)
			fmt.Println(t.ID)
		}
	}

	//fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
func Es_test() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 创建term查询条件，用于精确查询
	//termQuery := elastic.NewTermQuery("description", "nProductId")

	searchResult, err := client.Search().
		Index("accounts"). // 设置索引名
		//Query(termQuery).              // 设置查询条件
		//Sort("createTime", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).      // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).     // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(ctx)       // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	fmt.Println(searchResult)

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Accounts
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			/*
				if t, ok := item.(Zhb_search_online_db_Item); ok {
					//fmt.Println(t.Title)
					//fmt.Println(t)
				}
			*/
			t, _ := item.(Accounts)
			fmt.Println(t)
		}
	}

	//fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
