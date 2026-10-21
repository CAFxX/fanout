package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint8 = 0
    v2 := []uint32{42, 42, 2147483648}
    var v3 bool = (((-988 != 1) && false) && false)
    v4 := []int{}
    var v5 int64 = 1
    var v6 bool = true
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v5)
        v6 = (1 == 860)
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = -984
        _ = fi8
        _ = v9
    }
    var v10 float64 = 3.14159
    fmt.Println(v10, v6)
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(fi11, v6, v3)
        v6 = true
        switch "foo bar" {
        case "foo bar":
            var v12 uint8 = v1
            _ = v12
        case "":
            v13 := GenS1{F0: 2147483647, F1: v3}
            _ = v13
        case "\n\t":
            var v14 int64 = (^(v5))
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v5, v6, v3)
    var v15 uint8 = v1
    if len(v2) > 0 {
        v2[0] = 8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
