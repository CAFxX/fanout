package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 226
    var v3 uint = 671
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    delete(v4, 476)
    fmt.Println(v1)
    var v6 int16 = -425
    var v7 uint8 = v2
    var v8 uint32 = 2147483648
    v9 := GenS1{F0: v6}
    v10 := v4[v8]
    var v11 float64 = 0.0
    v4[0] = 1
    v11 -= v11
    var v12 string = "hello"
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
