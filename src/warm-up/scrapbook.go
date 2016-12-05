package main
    import "fmt"
    import "math"
    
func main() {
//const s:="Bipin" does not work
s:="Bipin"
//const  s string = "constant"  error as s is already declared
const constant string = "constant"

fmt.Println("value of s => ",s)
fmt.Println("value of constant => ",constant)
fmt.Println("value of Pi=>", 22/7)
fmt.Println("sin value 30 =>",math.Sin(30) )
//wrong
/*
for i:=0; i<= 3 {
    fmt.Println(i)
    
}
*/


for i:=0; i<= 3 ;i++{
    fmt.Println(i)
    
}
for j:=7; j<=9;j++{
    fmt.Println(j);
}

k:=0;
for k<3 {
    fmt.Println("value of k =>", k)
    k++ // loop does not terminate if the value of k is not encremented but program compiles.
}

for {
    fmt.Println("test break.. print only once")
    break;
}

if 8%4==0 {
    fmt.Println("int%int is mod operation,8%4 is 0")
}
const n int =10
if n>11 {
    fmt.Println(" if block if n>11 where n =10")
} else {
    fmt.Println("else block if n>11 where n =10")
}

}
