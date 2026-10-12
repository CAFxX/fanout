package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }

func main() {
    var v1 uint = 0
    v2 := []int{}
    var v3 uint = (v1 - (^(v1)))
    var v4 float64 = 1e10
    var v5 uint = 18446744073709551615
    v1 += 0
    v6 := GenS1{F0: 32767, F1: 0}
    switch 32768 {
    case 9223372036854775807:
        v7 := GenS2{F0: -2147483648, F1: 1000, F2: 2147483647}
        _ = v7
    case 1:
        fmt.Println(v5, v4)
    }
    var v8 int8 = 42
    v3 -= (^(v3))
    switch -359 {
    case 128:
        v9 := GenS2{F0: 1, F1: 0, F2: 2147483647}
        _ = v9
    }
    fmt.Println(v1, v3)
    v1 = v1
    var v10 bool = false
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if v10 {
        v4 -= (0.1 / (v4 / v4))
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v12 uint32 = 1
        if (((v3 - (v1 | v3)) < (v3 << (((v1 * (v5 << ((v3) & 63)))) & 63))) && v10) {
            if v10 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        _ = fi11
        _ = v12
    }
    var v13 int16 = -32768
    v10 = true
    var v14 uint64 = 18446744073709551615
    var v15 string = ("a" + "foo bar")
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
