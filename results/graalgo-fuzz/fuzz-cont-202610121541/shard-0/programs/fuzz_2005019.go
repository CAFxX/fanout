package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 38
    v3 := []uint32{2147483647, 856, 2147483648}
    var v4 uint = 551
    v5 := []uint32{1, 1, 847, 510}
    v6 := []uint32{}
    if false {
        var v7 uint64 = 18446744073709551615
        var v8 uint = v4
        _ = v7
        _ = v8
    }
    fmt.Println(v2)
    var v9 uint64 = 1
    v10 := GenS1{F0: uint8(v4)}
    _ = len(v5)
    var v11 uint8 = 127
    var v12 uint8 = 228
    if true {
        v9 -= v9
    } else {
        v9 -= 955
        v13 := GenS1{F0: (v12 / ((1) | 1))}
        _ = v13
    }
    v11 -= 243
    fmt.Println(v11)
    v12 = 127
    var v14 int = -662
    for fi15 := 0; fi15 < 14; fi15++ {
        switch "hello" {
        case "a":
            fmt.Println(v4)
        case "":
            v12 = (v12 - (v11 / ((255) | 1)))
        default:
            var v16 uint32 = 42
            _ = v16
        }
        _ = fi15
    }
    switch (v14 / ((-680) | 1)) {
    case 32768:
        if len(v5) > 0 {
            v5[0] = 65535
        }
    case -416:
        _ = len(v6)
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    v17 := GenS2{F0: v4, F1: 42}
    var v18 float32 = float32(0.1)
    var v19 uint = (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
