// Define a type
type TypeName struct {
    // fields
}

// Attach a method to that type
func (t TypeName) methodName() returnType {
    // access t.field, etc.
}

t - receiver. 
you can access the elements of the struct using this receiver

func main() {
variable_name := TypeName{values} // this will assign the values to the fields declared inside the struct 
fmt.Prinln(variable_name.methodName()) // variable name acts as an receiver here when the methodName is called, which is nothing but a normal function call
}
