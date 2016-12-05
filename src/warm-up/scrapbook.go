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
    k++ // loop does not terminate if the value of k is not incremented.
}

for {
    fmt.Println("test break.. print only once")
    break;
}

if 8%4==0 {
    fmt.Println("int%int is mod operation,8%4 is 0")
}
//const n int =10
if n:=10; n>11 {   // variable declared in if is avilable in all branches.
    fmt.Println(" if block if n>11 where n =10")
} else {
    fmt.Println("else block if n>11 where n =10")
}


i:=10
for j:=1;j<=3; j++ { // j=0 will bring integer divide by 0 error.
    k = i%j
    fmt.Println("value of switch =>",k)
    switch k {
        case 0 :
        fmt.Println(" print 0")
        case 1:
        fmt.Println("print 1")
        case 2:
        fmt.Println("print2")
        case 3:
        fmt.Println("print 3")
        default:
        fmt.Println("default case")
    }
}


}
