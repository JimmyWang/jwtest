package metadata

type IDataSource interface {
	DataSourceTesting(connections []Connection) bool
}

type DataSource struct {
	Type        string
	Name        string
	Version     string
	Connections []Connection
}

type Connection struct {
	Name string // 表单实际数据字段名称
	DisplayName string // 页面显示名称
	Editable bool // 是否可编辑
	Value string // 值 (默认值到真实值)
	ValueType  string // 数据类型
}

type EntityResource struct {
	Name string
	Desc string
}

type FieldResource struct {
	Name string
	Desc string
	Type string
}

type PageInfo struct {
	PageNum  int
	PageSize int
}

type TimeWindow struct {
	realTime string
	history  string
}

type QueryData struct {
	fieldId     string
	Connection  map[string]interface{}
	dataKey     string     // 数据key，针对于同一个客户端发出的请求保持唯一
	aggregation string     // 聚合类型
	timeWindow  TimeWindow // 时间窗口
}

type QueryResult struct {
	timestamp string
	value     string
}
