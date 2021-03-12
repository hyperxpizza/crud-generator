package parser

type Query struct {
	Type      string
	TableName string
	Fields    map[string]interface{}
}

func Parse(sql string) (*Query, error) {
	return nil, nil
}
