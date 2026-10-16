package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 bool, p2 uint) int {
    var v6 int = 127
    fmt.Println(p1, p1, p0)
    var v7 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(v6))
}

func main() {
    var v1 float64 = (0.0 / -2.25)
    var v2 uint8 = 1
    var v3 int64 = -459
    var v4 string = "foo bar"
    var v5 uint = 777
    var v8 uint16 = 32767
    v3 = v3
    fmt.Println(v5)
    if false {
        v9 := GenS1{F0: (v3 ^ 9223372036854775807), F1: int32(v8)}
        v10 := []int{int(v5)}
        _ = v10
        _ = v9
    }
    var v11 int64 = -955
    v5 = (^(v5))
    v12 := GenS2{F0: 42}
    var v13 bool = (!(2147483647 == 4294967295))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
