package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 678
    var v3 int = 2147483647
    var v4 int8 = 127
    v5 := make(map[int]bool)
    var v6 uint8 = 231
    v7 := GenS1{F0: -1, F1: v2}
    fmt.Println(v4, v3, v1)
    if (v4 < (v4 / ((v4) | 1))) {
        v4 = 1
        fmt.Println(v4)
    }
    var v8 int16 = -1
    v1 += (v1 + v1)
    _ = len(v5)
    if true {
        fmt.Println(v4, v8)
        if (4294967295 != 1) {
            v9 := GenS2{F0: 2147483647}
            _ = v9
        }
        var v10 uint16 = 1
        _ = v10
    }
    var v11 uint64 = (v2 % (((v2 & (v2 - (v2 & (v2 + 0))))) | 1))
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = 7
        _ = fi12
        _ = v13
    }
    var v14 int64 = 486
    switch -534 {
    case 127:
        v15 := GenS3{F0: v14, F1: 2, F2: v4}
        _ = v15
    default:
        var v16 uint64 = 42
        _ = v16
    }
    var v17 uint64 = (v2 - 18446744073709551615)
    var v18 uint = 18446744073709551615
    v19 := []int{v3, 160}
    for fi20 := 0; fi20 < 10; fi20++ {
        if false {
            v14 -= (v14 / ((-9223372036854775808) | 1))
            var v21 string = "foo bar"
            _ = v21
        }
        var v22 uint16 = 42
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi20
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
