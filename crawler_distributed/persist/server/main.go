package server

import (
	"GoTest/crawler_distributed/persist"
	"GoTest/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func ServeRpc(host, index string) error {
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	return nil
}
