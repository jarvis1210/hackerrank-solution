// Sample code to perform I/O:
package main
	
import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)
 
//fmt.Scanf("%s", &myname)            // Reading input from STDIN
//fmt.Println("Hello", myname)        // Writing output to STDOUT
 
// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
 
 
// Write your code here
func main(){
 
    var T int
    
 
     fmt.Scanf("%d",&T)
    
 
    for  i := 1; i <= T; i++ {
        
        var N int
        fmt.Scanf("%d",&N)
        minNum := minNum(N)
        fmt.Println(minNum)
    }
}
 
func minNum(num int) int{
    //fmt.Println(num)
 
    var numArr1 [] int
    var subSlice1 [] int
    var subSlice2 [] int
    
    for (num != 0 ){
        
        rem := num % 10
        //fmt.Println("rem", rem)
        numArr1 = append(numArr1,rem)
        num = num / 10
        //fmt.Println("num", num)
 
    }
    //fmt.Println(numArr1)
    sort.Ints(numArr1)
 
    for i, val := range numArr1{
        //fmt.Println("index ", i," value" , val)
        if( i % 2 ==0 ){
            subSlice1 = append(subSlice1,val)
        }else{
            subSlice2 = append(subSlice2,val)
        }
    }
 
    //subSlice1StringSlice := make(subSlice1,len(subSlice1))
    //subSlice2StringSlice := make(subSlice2,len(subSlice1))
 
 
    subSlice1StringSlice := make([]string, len(subSlice1))
    subSlice2StringSlice := make([]string, len(subSlice2))
 
    for ind, val := range subSlice1 {
        subSlice1StringSlice[ind] = strconv.Itoa(val)
    }
 
    for ind, val := range subSlice2 {
        subSlice2StringSlice[ind] = strconv.Itoa(val)
    }
 
    
    firstNum,_ := strconv.Atoi(strings.Join(subSlice1StringSlice, ""))
    secondNum,_ := strconv.Atoi(strings.Join(subSlice2StringSlice, ""))
    
    return firstNum + secondNum
}
Language: Go
