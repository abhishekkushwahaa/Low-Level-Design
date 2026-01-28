package main

// Prototype interface
type Prototype interface {
	Clone() Prototype
}

// Product
type User struct {
	Name string
	Age  string
	Meta *Metadata
}

type Metadata struct {
	Role string
}

/*
1. Basic Prototype (Shallow Copy)
Benefits:
a)- Fast object creation
b)- Avoids expensive initialization
c)- Simple implementation
Problems:
a)- Shallow copy shares references
b)- Changes in nested objects affect original
*/
func (u *User) Clone() Prototype {
	copy := *u // shallow copy
	return &copy
}

/*
2. Deep Copy Prototype
Benefits:
a)- Fully independent copy
b)- Safe for mutable nested objects
Problems:
a)- More code
b)- Slight performance overhead
*/
func (u *User) DeepClone() *User {
	copy := *u
	if u.Meta != nil {
		copy.Meta = &Metadata{
			Role: u.Meta.Role,
		}
	}
	return &copy
}

/*
3. Prototype Registry (Cache/Map)
Benefits:
a)- Centralized prototype management
b)- Fast object creation from templates
c)- Avoids repeated object construction
Problems:
a)- Registry must be maintained
b)- Can grow large if unmanaged
*/
var prototypeRegistry = map[string]Prototype{}

func RegisterPrototype(key string, prototype Prototype) {
	prototypeRegistry[key] = prototype
}

func GetPrototype(key string) Prototype {
	if prototype, ok := prototypeRegistry[key]; ok {
		return prototype.Clone()
	}
	return nil
}
