package search

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	//"log"
	"net/http"
)
type item struct {
XMLName    xml.Name `xml:"item"`
PubDate    string   `xml:"pubDate"`
Title      string   `xml:"title"`
Description string `xml:"description"`
Link string `xml:"link"`
GUID string `xml:"guid"`
GeoRssPoint string `xml:"georss:point"`
}
type image struct {
XMLName xml.Name `xml:"image"`
URL string `xml:"url"`
Title string `xml:"title"`
Link string `xml:"link"`
}
type channel struct {
	XMLName xml.Name `xml:"channel"`
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	LastBuildDate string `xml:"lastBuildDate"`
	TTL string `xml:"ttl"`
	Language string `xml:"language"`
	ManagingEditor string `xml:"managingEditor"`
	WebMaster string `xml:"webMaster"`
	Image image `xml:"image"`
	Item []item `xml:"item"`
}

type rssDocument struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel `xml:"channel"`
}

func Search(feed *Feed,searchTerm string)([]*Result,error){
	var results []*Result
	log.Printf("Search Feed Type[#{feed.Type}] Site[#{feed.Name}] For Url[#{feed.URL}]\n")

	document,err:=retrieve(feed)
	if err!=nil{
		return nil,err
	}
	for _,channelItem:=range document.Channel.Item{
		
		results = append(results, &Result{
			Filed: "Title",
			Content: channelItem.Title,})
		results=append(results,&Result{
			Filed:   "Description",
			Content: channelItem.Description,
		})
	}
	return results,nil
}
func retrieve(feed *Feed)(*rssDocument,error){
	if feed.URL == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// 从网络获得 rss 数据源文档
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}

	// 一旦从函数返回，关闭返回的响应链接
	defer resp.Body.Close()

	// 检查状态码是不是 200，这样就能知道
	// 是不是收到了正确的响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// 将 rss 数据源文档解码到我们定义的结构类型里
	// 不需要检查错误，调用者会做这件事
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}