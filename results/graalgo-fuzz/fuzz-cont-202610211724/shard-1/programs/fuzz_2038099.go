package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]bool)
    var v3 int64 = -1
    v4 := make(map[int]string)
    var v5 uint = 18446744073709551615
    _ = len(v1)
    if false {
        fmt.Println(v3)
        if len(v1) > 0 {
            v1[0] = 0
        }
    } else {
        fmt.Println(v5, v5)
    }
    v3 += (v3 * 926)
    var v6 int32 = -213
    if ((!((v6 / ((-2147483648) | 1)) > (v6 ^ (v6 & v6)))) || (127 <= 1)) {
        v7 := GenS1{F0: 1, F1: 4294967295}
        switch 32767 {
        case 794:
            v3 = -1
        default:
            var v8 int16 = -503
            _ = v8
        }
        _ = v7
    }
    v5 = v5
    v6 = 344
    var v9 int8 = 0
    switch "hello" {
    case "\n\t":
        for fi10 := 0; fi10 < 20; fi10++ {
            v5 += v5
            _ = fi10
        }
    case "a":
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int16 = 1000
            _ = fi11
            _ = v12
        }
    }
    var v13 bool = (v6 <= v6)
    v14 := GenS1{F0: v9, F1: 0}
    v15 := GenS1{F0: (v9 - v9), F1: 0}
    var v16 int8 = (v9 * v9)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
