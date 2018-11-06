package config

// Se exportan las funciones para el testing externo
// esto no afecta el ejecutable ya que es un archivo _test.go
var (
	LoadFile  = loadFile
	LoadBytes = loadBytes
)
