# Learn Go

![Learning Go](https://mytrashcode.com/wp-content/webp-express/webp-images/doc-root/wp-content/uploads/2019/11/4.jpg.webp "Learning Go")

> This is a ReadMe template to learn basic syntax of Golang.

---

### Table of Contents
You're sections headers will be used to reference location of destination.

- [Description](#description)
- [How To Use](#how-to-use)
- [Basic Syntax](#basic-syntax)
- [Error Handling](#error-handling)
- [Struct in Go](#struct-in-go)
- [Pointers in Go](#pointers-in-go)
- [Interface in Go](#interface-in-go)
- [Go Routines](#go-routiens)
- [References](#references)
- [License](#license)
- [Author Info](#author-info)

---

## Description

Creating ReadMe's for your Github repository can be tedious.  I hope this template can save you time and effort as well as provide you with some consistency across your projects.

#### Technologies

- Technology 1
- Technology 2

[Back To The Top](#learn-go)

---

## How To Use

## Basic Syntax

1) **string(bs)** -> converts Byte Slice to the String type
2) **strings.split(s, separator string)** ->  String.split(String s,char c)
3) **rand.Intn(i int)** -> Random.rand(Integer i)
4) **fmt.Printf("%+v", struct)** -> Prints the value of each field in the struct followed by actual value

5) **Value Types** -> int, float, string, bool, structs (we use pointers to change these things in a function)
6) **Reference Types** -> slice, maps, channels, pointers, functions. Don't worry about pointers with these.
7) **Make Function** -> bs = make([] byte, 9999) - makes a byte slice with 9999 elements in it.
8) **io.Copy** -> io.Copy(os.Stdout, resp.Body) 
#### Installation



#### API Reference

```html
    <p>dummy code</p>
```
[Back To The Top](#learn-go)

---

## References
[Back To The Top](#learn-go)

---
## Error Handling
[Back To The Top](#learn-go)

---
## Unit Tet in Go
[Unit Test in Go](#unit-tests)

- To make a test, create a new file ending in _test.go like **deck_test.go**
- To run all the tests in a package, run the command
```go
    go test
```

[Back To The Top](#learn-go)

---
## Struct in Go

- Data Structure, Collection of properties that are related.
- Declaring a struct 
```html
type person struct {
    firstName string
    lastName string
}

func main() {
    alex := person{firstName: Alex, lastName: Anderson}
    fmt.Println(alex)
}
```
**More Example of Struct**
```html
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
```

[Back To The Top](#learn-go)

---

## Pointers in Go
- **&variable** -> (& is an operator) Give me the memory address of the value this variable is pointing at.
- ***pointer** -> (\* is an operator) Give me the value that this memory addresss is pointing to.
- ***person** -> (person is a struct type) Type description - we are working with a pointer to a person.

>> Turn address into value using *address.  
Turn value into adderss using &value.  
Pointer shortcut is to call a function expecting the pointer to a type using the type only.

```html
    func (pointerToPerson *person) updateName() {
        (*pointerToPerson).firstname = "jimmy"
    }
```

### Pointer Gotcha
The below code will output [Bye There How Are You] which is different than expectation, but behaviour of Pointers is different for slice and struct.

```html
    func main() {
        mySlice := []string{"Hi", "There", "How", "Are", "You"}
        updateSlice(mySlice)
        fmt.Println(mySlice)
    }

    func updateSlice(s []string) {
        s[0] = "Bye"
    }
```
Slice internally stores the elements in an array.
Slice has three components (ptr to head of an array, capacity, length). Hence even when there is a copy of slice the underlying array is the same.

slice is referred to as Reference type.



[Back To The Top](#learn-go)

---

## Map in Go  

Map stores key, value pairs, all keys and all values should be of same type respectively

[Back To The Top](#learn-go)

---

## Interface in Go  
Declaring an interface type in a file, says our program has a new type called bot

```html
    type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

```
If you are a type in this program with a function called 'getGreeting' and you return a string then you are a member of type 'bot'

Now that you are an honorary member of type 'bot' you can now call this function called printGreeting()

- Interfaces are not generic types and Go does not have support for the same.
- Interfaces are **implicit**, as in java we don't manually have to say that our custom type satisfies some interface
- Interfaces are a contract to help us manage types. if our custom type implementation is broken interfaces won't help.
- Interfaces are tough. Step #1 is understanding to read them. Writing our own interfaces is tough and requires experience.
- We can take different interfaces and combine them together to form a new interface.

```html
    type ReadCloser interface {
        Reader 
        Closer
    }

    type Reader interface {
        read([]byte) (int, error)
    }

```

[Back To The Top](#learn-go)

---
## Go Routines
1) Our running program (a process) 
--
## Author Info

- Github - [Kapil Gupta](https://github.com/kapilgupta101292)

[Back To The Top](#learn-go)