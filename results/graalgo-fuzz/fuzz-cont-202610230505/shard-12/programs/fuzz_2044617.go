package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = float32((3.14159 - 0.1))
    var v3 int = -1
    var v4 uint8 = 42
    v5 := []uint32{65535, 2147483648, 966}
    var v6 float32 = v2
    var v7 int64 = -1
    v8 := v1[(v3 % ((200) | 1))]
    v4 += (v4 + (v4 / ((v4) | 1)))
    var v9 uint = 170
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 14, F1: uint(v4)}
            v9 = 1
            v6 -= v2
            v2 += 0.1
            var v13 int8 = 100
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int16 = 32767
        _ = fi10
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: (^(v9))}
    switch "hello" {
    case "":
        v16 := GenS1{F0: -1, F1: v9}
        _ = v16
    default:
        var v17 float64 = (1.5 + 3.14159)
        _ = v17
    }
    v18 := GenS2{F0: 42}
    for fi19 := 0; fi19 < 7; fi19++ {
        if (0.0 == ((3.14159 + (1e10 * float64(0.1))) - 0.0)) {
            v4 -= uint8(v9)
            v20 := v1[v8]
            v20 = -2147483648
            _ = v20
        }
        _ = fi19
    }
    var v21 string = ("" + "")
    fmt.Println(v6, v21, v6)
    v22 := GenS2{F0: 18446744073709551615}
    var v23 int32 = int32(v4)
    if (277 != 2147483648) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v4 = v4
        var v24 int32 = v23
        _ = v24
    }
    v23 += 825
    var v25 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
