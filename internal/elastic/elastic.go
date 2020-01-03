package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

// Elastic -
type Elastic struct {
	*elasticsearch.Client
}

// New -
func New(addresses []string) (*Elastic, error) {
	elasticConfig := elasticsearch.Config{
		Addresses: addresses,
	}
	es, err := elasticsearch.NewClient(elasticConfig)
	if err != nil {
		return nil, err
	}
	e := &Elastic{es}
	r, err := e.TestConnection()
	if err != nil {
		return nil, err
	}
	log.Printf("Elasticsearch Server: %s", r.Get("version.number").String())

	return e, nil
}

func (e *Elastic) getResponse(resp *esapi.Response) (*gjson.Result, error) {
	if resp.IsError() {
		return nil, fmt.Errorf(resp.String())
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(b)
	return &r, nil
}

func (e *Elastic) query(index string, query map[string]interface{}, source ...string) (r *gjson.Result, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return
	}

	// log.Printf("Body: %s", buf.String())

	// Perform the search request.
	var resp *esapi.Response
	options := []func(*esapi.SearchRequest){
		e.Search.WithContext(context.Background()),
		e.Search.WithIndex(index),
		e.Search.WithBody(&buf),
		e.Search.WithSource(source...),
	}

	if resp, err = e.Search(
		options...,
	); err != nil {
		return
	}
	defer resp.Body.Close()

	return e.getResponse(resp)
}

// TestConnection -
func (e *Elastic) TestConnection() (*gjson.Result, error) {
	res, err := e.Info()
	if err != nil {
		return nil, err
	}

	return e.getResponse(res)
}

// CreateIndexIfNotExists -
func (e *Elastic) CreateIndexIfNotExists(index string) error {
	_, err := e.MatchAll(index)
	if err != nil {
		if !strings.Contains(err.Error(), IndexNotFoundError) {
			return err
		}
	} else {
		return nil
	}

	jsonFile, err := os.Open(fmt.Sprintf("mappings/%s.json", index))
	if err != nil {
		log.Printf("Can't open %s.json file. Loading default config.", index)
		return err
	}
	defer jsonFile.Close()

	res, err := e.Indices.Create(index, e.Indices.Create.WithBody(jsonFile))
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("%s", res)
	}
	return nil

}

// AddDocument -
func (e *Elastic) AddDocument(v interface{}, index string) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	req := esapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(b),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), e)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	r, err := e.getResponse(res)
	if err != nil {
		return "", err
	}
	return r.Get("_id").String(), nil
}

// BulkInsertArray -
func (e *Elastic) BulkInsertArray(index string, v interface{}) error {
	bulk := bytes.NewBuffer([]byte{})
	arr := v.([]interface{})
	for i := range arr {
		id := uuid.New().String()
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id": "%s"} }%s`, id, "\n"))
		data, err := json.Marshal(arr[i])
		if err != nil {
			log.Println(err)
			continue
		}
		data = append(data, "\n"...)

		bulk.Grow(len(meta) + len(data))
		bulk.Write(meta)
		bulk.Write(data)
	}
	return e.BulkInsert(index, bulk)
}

// CreateIndex -
func (e *Elastic) CreateIndex(index string) error {
	resp, err := e.Indices.Create(index, e.Indices.Create.WithContext(context.Background()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = e.getResponse(resp)
	return err
}

// GetByID -
func (e *Elastic) GetByID(index, id string) (r *gjson.Result, err error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}
	resp, err := req.Do(context.Background(), e)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return e.getResponse(resp)
}

// Match - returns data by match filter
func (e *Elastic) Match(index string, match map[string]interface{}) (r *gjson.Result, err error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": match,
		},
	}
	return e.query(index, query)
}

// MatchAll - returns all data
func (e *Elastic) MatchAll(index string) (r *gjson.Result, err error) {
	return e.query(index, queryAll)
}

// UpdateDoc - updates document by ID
func (e *Elastic) UpdateDoc(index, id string, v interface{}) (r *gjson.Result, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(b),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), e)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return e.getResponse(res)
}

// BulkInsert -
func (e *Elastic) BulkInsert(index string, buf *bytes.Buffer) error {
	req := esapi.BulkRequest{
		Index:   index,
		Body:    bytes.NewReader(buf.Bytes()),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), e)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = e.getResponse(res)
	return err
}