package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 bool = true
    var v2 uint = 0
    var v3 uint = 644
    var v4 bool = false
    v2 = (v2 ^ (^(v3)))
    fmt.Println(v4)
    v4 = (255 == 67)
    v5 := []int{127, 2147483647}
    v1 = (333 > -1)
    var v6 float32 = 3.14159
    v7 := GenS1{F0: 439}
    v5 = append(v5, 9223372036854775807)
    var v8 float32 = (3.14159 + 1e10)
    var v9 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
