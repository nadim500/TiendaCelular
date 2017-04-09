package common

/*StartUp inicia todas las funciones por primera vez ...*/
func StartUp() {
	initConfig()
	createDBSession()
}
