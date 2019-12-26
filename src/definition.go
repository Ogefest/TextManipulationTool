package main

func GetDefinitions() map[string]CommandDefinition {

	var result = map[string]CommandDefinition{
		"replace": CommandDefinition{
			Command:        "replace",
			NumberOfParams: 2,
			Function:       Replace,
		},
		"lower": CommandDefinition{
			Command:        "lower",
			NumberOfParams: 0,
			Function:       Lower,
		},
		"upper": CommandDefinition{
			Command:        "upper",
			NumberOfParams: 0,
			Function:       Upper,
		},
		"title": CommandDefinition{
			Command:        "title",
			NumberOfParams: 0,
			Function:       Title,
		},
		"trim": CommandDefinition{
			Command:        "trim",
			NumberOfParams: 0,
			Function:       Trim,
		},
		"duplicate": CommandDefinition{
			Command:        "duplicate",
			NumberOfParams: 0,
			Function:       Duplicate,
		},
		"filter": CommandDefinition{
			Command:        "filter",
			NumberOfParams: 1,
			Function:       Filter,
		},
		"prefix": CommandDefinition{
			Command:        "prefix",
			NumberOfParams: 1,
			Function:       Filter,
		},
		"suffix": CommandDefinition{
			Command:        "suffix",
			NumberOfParams: 1,
			Function:       Filter,
		},
		"cut": CommandDefinition{
			Command:        "cut",
			NumberOfParams: 1,
			Function:       Cut,
		},
		"columnseparator": CommandDefinition{
			Command:        "columnseparator",
			NumberOfParams: 1,
			Function:       ColumnSeparator,
		},
		"columnremove": CommandDefinition{
			Command:        "columnremove",
			NumberOfParams: 1,
			Function:       ColumnRemove,
		},
		"columnadd": CommandDefinition{
			Command:        "columnadd",
			NumberOfParams: 2,
			Function:       ColumnAdd,
		},
	}

	return result
}
