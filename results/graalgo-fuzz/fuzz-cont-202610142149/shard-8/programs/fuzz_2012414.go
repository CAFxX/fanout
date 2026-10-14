package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint8 = 128
    var v2 bool = (399 != 1000)
    var v3 int16 = 234
    var v4 uint16 = 663
    v5 := []uint32{978, 65535, 291}
    var v6 float64 = float64(-2.25)
    v7 := GenS1{F0: 32767}
    v2 = v2
    var v8 uint = 2
    fmt.Println(v1)
    var v9 uint16 = (v4 ^ (v4 << ((uint(v4)) & 63)))
    var v10 int8 = 1
    v11 := GenS1{F0: -9223372036854775808}
    v3 = v3
    v5 = append(v5, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
