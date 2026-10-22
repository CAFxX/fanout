package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 1
    v3 := []string{}
    var v4 float64 = 3.14159
    v1 += (1e10 + v1)
    v5 := GenS1{F0: -962}
    v6 := GenS1{F0: -660}
    var v7 int8 = 81
    v3 = append(v3, "")
    if true {
        var v8 uint = 2
        _ = v8
    }
    var v9 float32 = 1.5
    v10 := GenS1{F0: 127}
    v4 -= 1e10
    _ = len(v3)
    fmt.Println(v7, v7)
    if (132 >= -2147483648) {
        if ((42 > -9223372036854775808) || false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = v2
        }
        v2 = 128
        var v11 uint64 = 2
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
