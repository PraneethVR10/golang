//Basic Struct Template

type TypeName struct {
fields datatype/ void
}

func main() {

single_variable := TypeName_struct{fields: value} // or you could just use TypeName_struct{values} - this is automatically assing the values to the fields based on the number of fields and values used. 
//to print the values
fmt.Println(single_variable.fields)
}
//Struct with methods 

type TypeName struct {
fields datatype/ void
}

func (receiver TypeName) methodName() datatype/void {
 return receiver.fields
}

func main() {

variable := TypeName_struct{fields: value} // which should always point to the TypeName as we are directly assigning the value to the fields declared in the Struct
 
//to print the values
fmt.Println(variable.method) while printing, you can always call the method where the action is completed via the return statement. 

}
