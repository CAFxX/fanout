package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }

func main() {
    var v1 int8 = -1
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 float32 = 0.0
    var v6 int32 = -2147483648
    fmt.Println(v2, v2, v2)
    v5 = (3.14159 / v5)
    v2 = v2
    delete(v3, 973)
    for fi7 := 0; fi7 < 8; fi7++ {
        v3[(fi7 - fi7)] = (!(false || (119 <= 0)))
        v2 = 32767
        if false {
            var v8 string = "\n\t"
            _ = v8
        } else {
            v9 := GenS1{F0: (v2 ^ v2), F1: (fi7 != fi7), F2: 164}
            _ = v9
        }
        _ = fi7
    }
    if false {
        var v10 uint = 18446744073709551615
        v11 := []int{}
        var v12 uint8 = 2
        _ = v10
        _ = v11
        _ = v12
    }
    v1 = -1
    if (v1 <= (v1 & v1)) {
        fmt.Println(v5)
        fmt.Println(v2)
        v13 := GenS2{F0: 502, F1: 0, F2: (^(v1))}
        _ = v13
    } else {
        v2 = v2
        switch "" {
        case "a":
            v1 = (v1 - (^(v1)))
        }
        var v14 uint = 18446744073709551615
        _ = v14
    }
    v15 := []int{-1, -2147483648, -708, 2147483647}
    v6 += (v6 - (v6 ^ (v6 & v6)))
    delete(v4, 1)
    fmt.Println(v1, v1)
    switch 981 {
    case 872:
        var v16 int = 42
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
