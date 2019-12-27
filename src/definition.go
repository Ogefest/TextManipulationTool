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
		"if": CommandDefinition{
			Command:        "if",
			NumberOfParams: 1,
			Function:       FIf,
		},
		"ifnot": CommandDefinition{
			Command:        "ifnot",
			NumberOfParams: 1,
			Function:       FIfNot,
		},
		"prefix": CommandDefinition{
			Command:        "prefix",
			NumberOfParams: 1,
			Function:       Prefix,
		},
		"suffix": CommandDefinition{
			Command:        "suffix",
			NumberOfParams: 1,
			Function:       Suffix,
		},
		"cut": CommandDefinition{
			Command:        "cut",
			NumberOfParams: 1,
			Function:       Cut,
		},
		"column": CommandDefinition{
			Command:        "column",
			NumberOfParams: 1,
			Function:       Column,
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
		"columnselect": CommandDefinition{
			Command:        "columnselect",
			NumberOfParams: 1,
			Function:       ColumnSelect,
		},
		"columndeselect": CommandDefinition{
			Command:        "columndeselect",
			NumberOfParams: 0,
			Function:       ColumnDeselect,
		},
		"columnorder": CommandDefinition{
			Command:        "columnorder",
			NumberOfParams: 1,
			Function:       ColumnOrder,
		},
	}

	return result
}
