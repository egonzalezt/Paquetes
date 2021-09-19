package greet //debe corresponder al nombre de la carpeta del paquete

//en este caso hay variables globales como esta
//lo que significa que cualquier paquete componente del paquete la puede usar
var txt = "sample" //no funciona el asignador de variable corta

//entre paquetes se pueden compartir funciones,var,const,estructuras,etc
//yo puedo determinar que identificador pueda ser exportada o no para que desde otro paquete no se acceda a eso
//por ejemplo no exporte txt pero si las funciones

//si la primera palabra de la funcion o variable empieza en mayuscula es porque se debe exportar si no es porque
//es una variable privada

//evitar que los paquetes sean paquetes que ya existan porque si yo creo un paquete fmt voy a tener un confilcto

func English() string {
	return "Hi "
}

func Spanish() string {
	return "Hola "
}

func Italian() string {
	return "Ciao "
}
