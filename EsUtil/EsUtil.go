package esutil

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
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

type Zhb_search_online_db_Item struct {
	DataID        string `json:"-"` //`json:"dataId,omitempty"`
	Author        string `json:"author,omitempty"`
	Category      string `json:"category,omitempty"`
	Content       string `json:"content,omitempty"`
	CreateTime    string `json:"createTime,omitempty"` //time.Time
	DataType      string `json:"dataType,omitempty"`
	Item          string `json:"item,omitempty"`
	Ext           string `json:"ext,omitempty"`
	ID            string `json:"-"`                       //`json:"id,omitempty"`
	PublishedTime string `json:"publishedTime,omitempty"` //time.Time
	SourceURL     string `json:"sourceUrl,omitempty"`
	Summary       string `json:"summary,omitempty"`
	Title         string `json:"title,omitempty"`
	Photo         string `json:"photo,omitempty"`
	Detail        string `json:"detail,omitempty"`
	Description   string `json:"description,omitempty"`
	ClassifyName  string `json:"classifyName,omitempty"`
	ClassifyDesc  string `json:"classifyDesc,omitempty"`
	NCode         string `json:"nCode,omitempty"`
	RedirectUrl   string `json:"redirectUrl,omitempty"`
	Class         string `json:"_class,omitempty"`
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

/*
Test_case:
curl -G -d 'sortedType=dataId' -d 'pageNumber=0' -d 'pageSize=100' -d 'searchType=PRODUCT' localhost:8000/search
*/
func Es_search(keyWord string, sortedType string, pageNumber string, pageSize string, searchType string) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 创建term查询条件，用于精确查询
	termQuery := elastic.NewTermsQuery("dataType", searchType)
	//matchQuery := elastic.NewMatchQuery("dataType", searchType)
	// 创建bool查询
	boolQuery := elastic.NewBoolQuery().Must()

	if keyWord != "" {
		multiMatchQuery := elastic.NewMultiMatchQuery(keyWord, "title", "content", "summary")
		boolQuery.Must(termQuery, multiMatchQuery)
	} else {
		matchAllQ := elastic.NewMatchAllQuery()
		boolQuery.Must(termQuery, matchAllQ)
	}
	/*
		src, err := q.Source()
		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(src)
		if err != nil {
			fmt.Printf("marshaling to JSON failed: %v", err)
			panic(err)
		}
		got := string(data)
		expected := `{"multi_match":{"fields":["subject","message"],"query":"this is a test"}}`
		if got != expected {
			fmt.Printf("expected\n%s\n,got:\n%s", expected, got)
			panic(err)
		}
	*/

	pagenum, _ := strconv.Atoi(pageNumber)
	pagesize, _ := strconv.Atoi(pageSize)
	fmt.Printf("sortedType %s\n", sortedType)

	searchResult, err := client.Search().
		Index("zhb_search_online_db"). // 设置索引名
		//Query(termQuery).              // 设置查询条件
		//Query(matchQuery).      // 设置查询条件
		//Query(multiMatchQuery). // 设置查询条件
		Query(boolQuery).       // 设置查询条件
		Sort(sortedType, true). // 设置排序字段，根据字段升序排序，第二个参数false表示逆序
		//Sort("dataId", true). // 设置排序字段，根据字段升序排序，第二个参数false表示逆序
		From(pagenum).  // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(pagesize). // 设置分页参数 - 每页大小
		Pretty(true).   // 查询结果返回可读性较好的JSON格式
		Do(ctx)         // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	//fmt.Println(searchResult)

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Zhb_search_online_db_Item
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Zhb_search_online_db_Item对象
			if t, ok := item.(Zhb_search_online_db_Item); ok {
				fmt.Println("===============================================")
				//fmt.Println(item)
				fmt.Printf("NCode %s, Title: %s\n", t.NCode, t.Title)
				//fmt.Println(t.Title)
				//fmt.Println(t)
			}
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
