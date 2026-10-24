package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 774
    var v3 uint16 = v2
    if ((-4 >= 100) && (101 >= 0)) {
        for fi4 := 0; fi4 < 10; fi4++ {
            fmt.Println(v2, v2, v1)
            v5 := GenS1{F0: -83}
            _ = fi4
            _ = v5
        }
        v6 := []int{32767, 1, 127}
        switch -9223372036854775808 {
        case -853:
            v1 += 0
        default:
            if false {
                v7 := GenS2{F0: 4294967295, F1: (^(v1))}
                _ = v7
            }
        }
        _ = v6
    }
    v2 = (v3 ^ v3)
    v8 := GenS1{F0: -2147483648}
    v9 := []int{32768, 32768, -655}
    v3 += v3
    v9 = append(v9, 387)
    switch -9223372036854775808 {
    case 4294967295:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    case 127:
        var v10 int32 = 42
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint8 = 42
        v12 -= v12
        _ = fi11
        _ = v12
    }
    v1 = (v1 ^ v1)
    var v13 int32 = 2147483647
    fmt.Println(v13, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
