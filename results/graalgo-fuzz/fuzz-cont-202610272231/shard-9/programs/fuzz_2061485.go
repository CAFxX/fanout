package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 float32 = (1.5 + -2.25)
    var v2 uint32 = 42
    var v3 int = 127
    var v4 uint32 = (^(v2))
    var v5 float32 = -2.25
    var v6 uint32 = 0
    v1 = v1
    var v7 float64 = ((3.14159 - 1e10) - (((3.14159 * 1e10) - 3.14159) * (-2.25 * 1.5)))
    v8 := []int{(v3 - -746), 761, 1}
    fmt.Println(v1, v5, v5)
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint16 = 2
            fmt.Println(fi9)
            v3 = (v3 + (v3 % ((1) | 1)))
            _ = fi9
            _ = v10
        }
        var v11 uint8 = 0
        _ = v11
    }
    v4 -= (v6 - (v4 % ((v2) | 1)))
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v12 := GenS1{F0: "\n\t", F1: 182, F2: true}
        _ = v12
    default:
        v13 := GenS2{F0: 525, F1: 128}
        _ = v13
    }
    var v14 bool = true
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v1, v4)
            var v17 uint64 = 4294967295
            if len(v8) > 0 {
                v8[0] = v3
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v6)
    v18 := GenS1{F0: "hello", F1: 128, F2: v14}
    v19 := GenS1{F0: "", F1: 42, F2: v14}
    var v20 uint64 = 85
    v1 = 0.1
    v8 = append(v8, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
