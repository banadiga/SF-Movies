package search

import (
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"github.com/yanzay/log"
	"golang.org/x/net/context"
	"encoding/json"
	"sync/atomic"
	"bytes"
)

const ELASTIC SearchEngineType = "elasticsearch"

type ElasticsearchEngine struct {
	client *elastic.Client
}

type CountingDecoder struct {
	dec json.Decoder
	N   int64
}

func (d *CountingDecoder) Decode(data []byte, v interface{}) error {
	atomic.AddInt64(&d.N, 10000)
	return json.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func NewElasticsearchEngine(url string) (ISearchEngine, error) {
	log.Println(url)

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetDecoder(&CountingDecoder{}),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticsearchEngine{
		client : client,
	}, nil
}

func (engine *ElasticsearchEngine) IndexExists(name string) (bool, error) {
	ctx := context.Background()
	exists, err := engine.client.IndexExists(name).
		Do(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (engine *ElasticsearchEngine) CreateIndex(name string, mapping string) (error) {
	ctx := context.Background()
	createIndex, err := engine.client.CreateIndex(name).
		BodyString(mapping).
		Do(ctx)
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		return errors.New("Not acknowledged")
	}
	return nil
}

func (engine *ElasticsearchEngine) Add(index string, body interface{}) (string, error) {
	ctx := context.Background()
	responce, err := engine.client.Index().
		Index(index).
		BodyJson(body).
		Do(ctx)
	if err != nil {
		return "", err
	}

	return responce.Id, nil
}

func (engine *ElasticsearchEngine) Search(index string, name string, value string, limit int) (int64, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery(name, value)
	searchResult, err := engine.client.Search().
		Index(index).
		Query(termQuery).
		Sort(name, true).
		From(0).Size(limit).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return 0, err
	}

	return searchResult.Hits.TotalHits, nil
}
