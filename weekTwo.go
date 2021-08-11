package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1.The calculate function should be a variadic function
//2. The passed in strings are checked for numeric characters they contain
//3. According to the numeric characters contained in the strings, the function relating to that operation is called.
//4. The values in that component is passed into the operation to be calculated and the result is returned in an empty array
// declared in the calculate function
//5. The array of results is returned
//
func division(divArray []float64) float64{
	var divResult float64 = divArray[0]
	for i:=1;i < len(divArray); i++{
		divResult /= divArray[i]
	}
	return divResult
}

////calculate function
func multiplication(mulArray []float64) float64 {
	var mulResult float64 = 1

	for _,v := range mulArray{
		mulResult *= v
	}

	return mulResult
}

////// subtraction function
func subtraction(subArray []float64) float64{
	var subResult float64

	for i, v:= range subArray{
		if i == 0{
			subResult = v
		}else{
			subResult -= v
		}
	}
	return subResult
}

////addition function
func addition(addArray []float64) float64 {
	var addResult float64
	for _,v := range addArray{
		addResult += v
	}
	return addResult
}



func calculator (values ...string) []float64 {
	var result []float64
	//fmt.Println(multiplyArray)

	for _,value := range values{

		if strings.Contains(value,"*" ){
			var multiplyArray []float64
			multiply := strings.Split(value, "*")
			for _,v := range multiply{
				if n, err := strconv.ParseFloat(v, 64); err == nil {
					multiplyArray = append(multiplyArray, n)
				}
			}
			answerMultiply := multiplication(multiplyArray)
			result = append(result, answerMultiply)
		}else if strings.Contains(string(value),"+"){
			var addArray []float64
			add := strings.Split(value,"+")
			for _,v := range add{
				if n,err := strconv.ParseFloat(v,64); err == nil{
					addArray = append(addArray,n)
				}
			}
			addAnswer := addition(addArray)
			result = append(result,addAnswer)

		}else if strings.Contains(string(value),"-"){
			var subArray []float64
			sub := strings.Split(value,"-")
			for _,v := range sub{
				if n,err := strconv.ParseFloat(v,64); err == nil{
					subArray = append(subArray,n)
				}
			}
			subAnswer := subtraction(subArray)
			result = append(result,subAnswer)
		}else {
			strings.Contains(string(value),"/")
				var divArray []float64
				div := strings.Split(value,"/")
				for _,v := range div{
					if n,err := strconv.ParseFloat(v,64);err == nil{
						divArray = append(divArray,n)
					}
				}
				divAnswer := division(divArray)
				result = append(result,divAnswer)
			}
	}
	return result

}

func main() {

	checker :=calculator("2+3","2-3", "3*2","4-5","2/2","5*4")
	fmt.Println(checker)
}