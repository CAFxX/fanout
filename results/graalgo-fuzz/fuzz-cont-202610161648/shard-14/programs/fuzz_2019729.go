package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 1
    var v2 int16 = -1
    v3 := make(map[uint32]bool)
    var v4 int32 = -1
    v5 := []int{-2147483648, -9223372036854775808}
    v4 -= v4
    v1 = 961
    v1 = (v1 ^ (^(v1)))
    v5 = append(v5, 4294967295)
    for fi6 := 0; fi6 < 5; fi6++ {
        v2 -= -32768
        var v7 bool = false
        fmt.Println(v2, v4, v1)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int16 = v2
        var v10 int32 = (v4 << ((2) & 63))
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: int8(v2), F1: 42, F2: (42 >= 9223372036854775807)}
        var v13 int = fi11
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = v4
    fmt.Println(v1, v2, v1)
    var v15 string = "hello"
    if (!false) {
        switch -1 {
        case -889:
            if ((((v15 + v15) + v15) + (v15 + "a")) != v15) {
                var v16 string = ((v15 + v15) + v15)
                _ = v16
            }
        case 774:
            if true {
                var v17 bool = false
                _ = v17
            }
        }
        var v18 uint = 603
        _ = v18
    }
    v19 := GenS1{F0: 1, F1: v1, F2: true}
    v20 := GenS1{F0: 1, F1: (v1 & v1), F2: true}
    if len(v5) > 0 {
        v5[0] = 1
    }
    switch v15 {
    case "\n\t":
        v2 = 1
    case "hello":
        v21 := GenS1{F0: -26, F1: (v1 % (((v1 >> ((4294967295) & 63))) | 1)), F2: ((0.1 / (3.14159 + 3.14159)) >= ((0.1 + 0.1) + 1e10))}
        _ = v21
    default:
        var v22 int8 = int8(v2)
        _ = v22
    }
    switch 361 {
    case 32767:
        v23 := GenS2{F0: 42}
        _ = v23
    case -2147483648:
        if true {
            fmt.Println(v1, v14, v14)
        }
    default:
        v4 -= (v4 - (v4 + (-(v14))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
