cat input.log | tmt replace ":" " " columselect 4 trim suffix : columnorder 1,2,4,3 columndeselect replace ": " ":"
