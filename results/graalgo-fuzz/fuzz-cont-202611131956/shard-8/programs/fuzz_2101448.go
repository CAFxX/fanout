package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    var v1 int8 = -128
    var v2 int32 = -2147483648
    var v3 int16 = 42
    var v4 int8 = v1
    fmt.Println(v1, v3, v4)
    v1 = v4
    v5 := []int{4294967295, 127}
    v1 = (^(v4))
    for fi6 := 0; fi6 < 18; fi6++ {
        _ = len(v5)
        _ = fi6
    }
    fmt.Println(v4, v4, v1)
    fmt.Println(v4, v4, v2)
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v7 float64 = -2.25
    if true {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 bool = (v3 != 1000)
            v7 = v7
            fmt.Println(v4)
            _ = fi8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 441, F1: -9223372036854775808}
        _ = v10
    }
    var v11 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
