package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 uint8 = 128
    var v2 uint8 = 42
    var v3 float32 = float32((-2.25 + -2.25))
    var v4 uint64 = 318
    var v5 int32 = 42
    var v6 uint = 255
    v4 = 2
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint64 = v4
        var v9 int16 = 724
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v4)
    v10 := []int{}
    switch "\n\t" {
    case "foo bar":
        switch 2147483647 {
        case 9223372036854775807:
            fmt.Println(v1, v2)
        default:
            v11 := GenS1{F0: "foo bar", F1: int8(v2)}
            _ = v11
        }
    }
    if len(v10) > 0 {
        v10[0] = int(v4)
    }
    fmt.Println(v4, v6)
    var v12 int = 127
    v13 := GenS1{F0: "a", F1: 42}
    if (1 <= (^(v5))) {
        v14 := GenS2{F0: v2, F1: (v12 ^ (v12 & (v12 >> (((v6 * v6)) & 63))))}
        var v15 int8 = -1
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: (v1 | 0), F1: 804}
    fmt.Println(v6, v5, v4)
    if false {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS2{F0: (^(v2)), F1: (^(v12))}
            v10 = append(v10, 49)
            _ = fi17
            _ = v18
        }
        var v19 int64 = 7
        for fi20 := 0; fi20 < 20; fi20++ {
            v2 = 1
            _ = fi20
        }
        _ = v19
    } else {
        v4 -= (^(v4))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
