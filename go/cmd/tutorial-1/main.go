package main
import "fmt" // for prining
import "unicode/utf8"

// Go does not allow random var hanging around if you create one you have to use it that's why i'm printing it to the console here!

func main(){
	var intNum int32 // In Go should think should what you store where
	// var intNum uint32 For unsigned int
	fmt.Println(intNum)
	
	var floatNum float32 //doesn;t have a float only float32 and float64
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.6
	var intNum32 int32 = 10
	// var result float32 = floatNum32 + intNum32 can't add 2 explict types together
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1/intNum2) // will return 1 will not provide a float value basically default is floor division
	fmt.Println(intNum1%intNum2) // To get remainder

	var myString string = "Hello Twin" // ain't no way we define string as 'string' class
	myString = "Hello" + " " + "Twin" // Also can you concatination
	fmt.Println(myString)
	
	fmt.Println(len(myString)) // Gives the number of bytes not the number of characters
	fmt.Println(utf8.RuneCountInString(myString)) // As special characters outside ascii result 2 as len does not result the number it results the number of bytes to get the right results when we use something outside of ASCII in a string we must use utf8.RuneCountInString to count the nunber of characters in the string also rune is another datatype in Go
	
	var myRune rune = 'a'
	fmt.Println(myRune) // Give 97 (ASCII value) rune is basically a number pointing a character it's just a number that says that this number represent this character like 97 for 'a' basically is you put a character it could be a emoji or letter etc it converts it to a number and store it as utf8 can't use these special characters
	
	var myBoolean bool = false // easy could be true or false
	fmt.Println(myBoolean)

	var myVar = "string" // We don't have to specify the datatype if we set the value right away, then the datatype is inferred here it assumes its a string as we defined its value there and there
	fmt.Println(myVar)

	myVariable := 10 // We can even drop the var keyword and declare a variable directly using := it is called shorthand
	fmt.Println(myVariable)

	var1, var2 := 1, 2 // We can declare multiple variable simoultaneously
	fmt.Println(var1,var2)
	
	// const myConst string // You can't just declare a const you have to initialize it with a value 
	const myConst = "Hey gng" // Everything applies from variables just you can't change a constant's value
	fmt.Println(myConst)
}

//Also you should specify the type explictly when you can and when the entered value is not obvious

// For ints uints floats and rune the default value is 0
// For strings it's '' or empty string
// For boolean it's false

