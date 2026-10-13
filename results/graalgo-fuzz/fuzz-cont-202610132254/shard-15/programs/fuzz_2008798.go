package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    v1 := []uint32{65535, 29}
    var v2 uint = 255
    v3 := make(map[uint32]uint32)
    var v4 uint = (v2 & (v2 ^ v2))
    v5 := make(map[int]uint32)
    var v6 uint8 = 255
    if true {
        v2 += 630
    } else {
        v7 := GenS1{F0: 846}
        _ = v7
    }
    switch "\n\t" {
    case "a":
        var v8 int64 = 42
        _ = v8
    case "\n\t":
        var v9 int16 = 42
        _ = v9
    case "":
        fmt.Println(v4, v2)
    default:
        v4 = v2
    }
    v10 := GenS1{F0: 284}
    var v11 uint = 0
    var v12 uint32 = 1
    var v13 uint32 = v12
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v4 = 4294967295
        if (true && false) {
            v6 = (v6 + uint8(v13))
            var v14 int16 = -1
            if ((true && true) && (-9223372036854775808 != -540)) {
                v15 := GenS1{F0: 771}
                v11 = (v4 % ((360) | 1))
                _ = v15
            } else {
                fmt.Println(v14, v14)
            }
            _ = v14
        }
    }
    var v16 float32 = (0.1 * float32((0.0 * (1.5 / 0.1))))
    var v17 string = "\n\t"
    v18 := GenS1{F0: 42}
    if true {
        delete(v3, (v12 ^ v12))
        switch 532 {
        case 128:
            fmt.Println(v6)
        }
        if (42 == -1) {
            switch ((v17 + ((((v17 + "a") + (v17 + (v17 + "foo bar"))) + v17) + v17)) + (v17 + v17)) {
            case "hello":
                v1 = append(v1, v13)
            default:
                v19 := GenS2{F0: 32767, F1: 32767}
                _ = v19
            }
        } else {
            v2 += v2
        }
    }
    v2 -= (v11 * v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
