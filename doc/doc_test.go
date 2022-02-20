package doc

import (
	"fmt"

	"github.com/alcor67/newprojectcalculator/calculator"
)

func Example() {
	result, err := calculator.Calc("+", 300, 500)
	fmt.Println(result, err)
	// Output: 800 <nil>
}
