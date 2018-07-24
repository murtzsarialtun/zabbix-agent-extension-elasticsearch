package main

import (
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

//NodeIndices - indices stats
type NodeIndices struct {
	Docs struct {
		Count   int64 `json:"count"`
		Deleted int64 `json:"deleted"`
	} `json:"docs"`

	Store struct {
		SizeInBytes          int64 `json:"size_in_bytes"`
		ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
	} `json:"store"`

	Indexing IndicesIndexingStats `json:"indexing"`
	Get      IndicesGetStats      `json:"get"`
	Search   IndicesSearchStats   `json:"search"`
	Merges   IndicesMergesStats   `json:"merges"`
}

// IndicesIndexingStats - indices indexing stats
type IndicesIndexingStats struct {
	IndexTotal           int64 `json:"index_total"`
	IndexTimeInMillis    int64 `json:"index_time_in_millis"`
	IndexCurrent         int64 `json:"index_current"`
	IndexFailed          int64 `json:"index_failed"`
	DeleteTotal          int64 `json:"delete_total"`
	DeleteTimeInMillis   int64 `json:"delete_time_in_millis"`
	DeleteCurrent        int64 `json:"delete_current"`
	NoopUpdateTotal      int64 `json:"noop_update_total"`
	IsThrottled          bool  `json:"is_throttled"`
	ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
}

// IndicesGetStats - indices get stats
type IndicesGetStats struct {
	Total               int64 `json:"total"`
	TimeInMillis        int64 `json:"time_in_millis"`
	ExistsTotal         int64 `json:"exists_total"`
	ExistsTimeInMillis  int64 `json:"exists_time_in_millis"`
	MissingTotal        int64 `json:"missing_total"`
	MissingTimeInMillis int64 `json:"missing_time_in_millis"`
	Current             int64 `json:"current"`
}

type IndicesSearchStats struct {
	OpenContexts        int64 `json:"open_contexts"`
	QueryTotal          int64 `json:"query_total"`
	QueryTimeInMillis   int64 `json:"query_time_in_millis"`
	QueryCurrent        int64 `json:"query_current"`
	FetchTotal          int64 `json:"fetch_total"`
	FetchTimeInMillis   int64 `json:"fetch_time_in_millis"`
	FetchCurrent        int64 `json:"fetch_current"`
	ScrollTotal         int64 `json:"scroll_total"`
	ScrollTimeInMillis  int64 `json:"scroll_time_in_millis"`
	ScrollCurrent       int64 `json:"scroll_current"`
	SuggestTotal        int64 `json:"suggest_total"`
	SuggestTimeInMillis int64 `json:"suggest_time_in_millis"`
	SuggestCurrent      int64 `json:"suggest_current"`
}

type IndicesMergesStats struct {
	Current                    int64 `json:"current"`
	CurrentDocs                int64 `json:"current_docs"`
	CurrentSizeInBytes         int64 `json:"current_size_in_bytes"`
	Total                      int64 `json:"total"`
	TotalTimeInMillis          int64 `json:"total_time_in_millis"`
	TotalDocs                  int64 `json:"total_docs"`
	TotalSizeInBytes           int64 `json:"total_size_in_bytes"`
	TotalStoppedTimeInMillis   int64 `json:"total_stopped_time_in_millis"`
	TotalThrottledTimeInMillis int64 `json:"total_throttled_time_in_millis"`
	TotalAutoThrottleInBytes   int64 `json:"total_auto_throttle_in_bytes"`
}

func createNodeStatsIndices(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	var nodeStats ElasticNodeStats

	for _, nodeStat := range nodesStats.Nodes {
		nodeStats = nodeStat
		break
	}

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.docs.count",
			),
			strconv.Itoa(int(nodeStats.Indices.Docs.Count)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.docs.deleted",
			),
			strconv.Itoa(int(nodeStats.Indices.Docs.Deleted)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.store.size_in_bytes",
			),
			strconv.Itoa(int(nodeStats.Indices.Store.SizeInBytes)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.store.throttle_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Store.ThrottleTimeInMillis)),
		),
	)

	metrics = createNodeStatsIndicesIndexing(
		hostname,
		&nodeStats,
		metrics,
		prefix,
	)

	metrics = createNodeStatsIndicesSearch(
		hostname,
		&nodeStats,
		metrics,
		prefix,
	)

	return metrics
}

func createNodeStatsIndicesIndexing(
	hostname string,
	nodeStats *ElasticNodeStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.index_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.IndexTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.index_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.IndexTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.index_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.IndexCurrent)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.index_failed",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.IndexFailed)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.delete_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.DeleteTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.delete_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.DeleteTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.delete_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.DeleteCurrent)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.noop_update_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.NoopUpdateTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.is_throttled",
			),
			strconv.FormatBool(nodeStats.Indices.Indexing.IsThrottled),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.indexing.throttle_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Indexing.ThrottleTimeInMillis)),
		),
	)

	return metrics
}

func createNodeStatsIndicesSearch(
	hostname string,
	nodeStats *ElasticNodeStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.open_contexts",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.OpenContexts)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.query_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.QueryTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.query_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.QueryTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.query_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.QueryCurrent)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.fetch_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.FetchTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.fetch_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.FetchTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.fetch_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.FetchCurrent)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.scroll_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.ScrollTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.scroll_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.ScrollTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.scroll_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.ScrollCurrent)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.suggest_total",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.SuggestTotal)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.suggest_time_in_millis",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.SuggestTimeInMillis)),
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.indices.search.suggest_current",
			),
			strconv.Itoa(int(nodeStats.Indices.Search.SuggestCurrent)),
		),
	)

	return metrics
}