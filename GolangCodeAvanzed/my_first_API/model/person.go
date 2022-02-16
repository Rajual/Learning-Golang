package model

//Community estructura de una comunidad
type Community struct {
	//Name nomber de una comunidad. Ej:Tensor
	Name string `json:"name"`
}

//Communities slice de comunidades
type Communities []Community

//Person estreuctura de una persona
type Person struct {
	//Name nombre de la persona
	Name string `json:"name"`
	//Age edad de la persona Ej:48
	Age uint8 `json:"age"`
	//Communities comunidades a alas que perteneces una persona
	Communities Communities `json:"communities"`
}

//Persons slice de personas
type Persons []Person
