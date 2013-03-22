package core

type DAO interface {
	//Creates the given value in the data source and returns the id. 
	Create(value interface{}) string

	//Reads the value with the given id from the data source and returns it.
	Read(id string) interface{}

	//Updates the object with the given id and overrides it with the values given in value.
	Update(id string, value interface{})

	//Deletes the value with the given id from the data source.
	Delete(id string)
}
