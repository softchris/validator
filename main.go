package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

// Person is ...
type Person struct {
	Id      int
	Name    string
	Age     int
	Email   string
	Address string
}

// Validator is ...
type Validator struct {
	Fields []*FieldType
}

// create constructor for Validator
func NewValidator() *Validator {
	return &Validator{[]*FieldType{}}
}

type FieldType struct {
	fieldName string
	fieldType string
	required  bool
}

func (f *FieldType) String() *FieldType {
	f.fieldType = "string"
	return f
}

func (f *FieldType) Int() *FieldType {
	f.fieldType = "int"
	return f
}

func (f *FieldType) IsRequired() *FieldType {
	f.required = true

	return f
}

func (f *FieldType) Email() *FieldType {
	f.fieldType = "email"
	return f
}

func (v *Validator) Field(name string) *FieldType {
	f := &FieldType{fieldName: name}
	// add field to Fields
	v.Fields = append(v.Fields, f)

	return f
}

func (v *Validator) Validate(obj interface{}) bool {
	// for each field in fields
	for _, field := range v.Fields {
		// check if field is required
		// does obj have field with field.fieldName
		if obj == nil {
			return false
		}
		// CHECK if fieldName in obj

		v := reflect.ValueOf(obj)
		f := v.FieldByName(field.fieldName)
		if f.IsValid() {
		} else {

			err := fmt.Sprintf("obj does not have field.fieldName: (%s)", field.fieldName)
			panic(err)
		}
		// print field name
		fmt.Printf("Field: %q\n", field.fieldName)
		// print field required
		fmt.Printf("Required: %t\n", field.required)
		// CHECK if field is required and if it is empty
		// if field is required and empty, return false

		if field.required == true {
			// check if v is a primitive type
			fmt.Print("checking required\n")
			// fmt.Printf("Kind: %q", f.Kind())

			if f.IsZero() {
				err := fmt.Sprintf("obj fieldName (%s) is required but lacks value", field.fieldName)
				panic(err)
			}
		}

		// get field type from f
		// fmt.Printf("Type: %q\n", f.Type())
		if field.fieldType != f.Type().String() {
			// ignore if email, we check that differently
			if !(f.Type().String() == "string" && field.fieldType == "email") {
				err := fmt.Sprintf("obj fieldName (%s) is not of type %s", field.fieldName, field.fieldType)
				panic(err)
			}
		}

		if field.fieldType == "email" {
			// create regex for email
			r := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
			// print field value (f.String())
			fmt.Printf("Value: %q\n", f.String())
			if r.MatchString(f.String()) == false {
				err := fmt.Sprintf("obj fieldName (%s) is not a valid email", field.fieldName)
				panic(err)
			}

		}
	}

	return true
}

// print all fields
func (v *Validator) PrintAllFields() {

	for _, field := range v.Fields {
		fmt.Printf("Field: %q\n", field)
	}
}
