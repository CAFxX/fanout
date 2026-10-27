package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint16 {
    var v5 int64 = -790
    v6 := GenS1{F0: -1}
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 4
}
func genFn1(p0 int32, p1 uint32) int64 {
    var v7 float64 = 3.14159
    if true {
        v8 := []int{911, 700, 1, 32768}
        v8 = append(v8, 2147483647)
        p0 += p0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return -679
}

func main() {
    v1 := []string{"a"}
    var v2 string = ("a" + ("hello" + "\n\t"))
    var v3 uint8 = 2
    var v4 int64 = -1
    v9 := GenS1{F0: 0}
    v4 -= (v4 + v4)
    var v10 int = 448
    var v11 int = (^(v10))
    var v12 float64 = 3.14159
    var v13 int32 = 2147483647
    var v14 uint32 = 42
    var v15 uint32 = 207
    v10 += (-(v10))
    fmt.Println(v10, v2, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
