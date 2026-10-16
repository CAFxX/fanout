package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2147483648
    var v3 int32 = -1
    if (926 != 992) {
        switch 32767 {
        case 127:
            switch "foo bar" {
            case "a":
                v4 := []int{-805}
                _ = v4
            }
        case -1:
            v3 = (v3 - 2147483647)
        }
    }
    if true {
        v5 := []int{}
        v2 = 0
        _ = v5
    }
    v6 := GenS1{F0: 128, F1: -128}
    fmt.Println(v3, v3, v3)
    if (!false) {
        var v7 float32 = ((-2.25 + 1e10) + (1.5 + 0.0))
        _ = v7
    }
    v2 = (v2 % ((v2) | 1))
    var v8 int16 = 7
    v1 -= (^(v1))
    var v9 int64 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 217, F1: 101}
        var v12 uint64 = v1
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{184, 32768}
    fmt.Println(v3, v1)
    var v14 bool = ((18446744073709551615 == 672) && false)
    for fi15 := 0; fi15 < 17; fi15++ {
        v1 -= (v1 - 1)
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
