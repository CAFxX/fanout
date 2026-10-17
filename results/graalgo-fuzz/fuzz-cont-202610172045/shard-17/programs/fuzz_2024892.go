package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0(p0 uint8) uint64 {
    if false {
        p0 = p0
    }
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 255, F1: false, F2: 100}
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 string = "hello"
            v7 = v7
            v8 := GenS1{F0: 509, F1: (1e10 == 0.1), F2: 42}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 uint8 = 163
    var v2 int64 = 551
    var v3 uint64 = 42
    var v9 uint8 = (v1 - v1)
    v10 := []int{9223372036854775807, -1, 9223372036854775807, -869}
    var v11 int32 = -1
    var v12 int64 = (v2 ^ (^(v2)))
    v13 := GenS1{F0: 65535, F1: false, F2: 1}
    var v14 int64 = 42
    if (v2 == (v14 % (((v2 >> ((0) & 63))) | 1))) {
        for fi15 := 0; fi15 < 20; fi15++ {
            switch fi15 {
            case -51:
                fmt.Println(v11, v14, v2)
            }
            fmt.Println(v11)
            _ = fi15
        }
        switch 128 {
        case 32768:
            fmt.Println(v1, v14, v12)
        default:
            var v16 bool = false
            _ = v16
        }
    }
    v17 := GenS1{F0: 65535, F1: (!false), F2: -12}
    v18 := GenS1{F0: 1, F1: (1000 == 32767), F2: -88}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
