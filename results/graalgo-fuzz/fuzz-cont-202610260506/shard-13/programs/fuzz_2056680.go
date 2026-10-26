package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int }

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{2147483648, 1, 2147483648}
    v3 := []uint32{}
    var v4 uint32 = 740
    v1 += (v1 ^ 1)
    var v5 uint = 2
    var v6 uint64 = 4294967295
    var v7 uint8 = 42
    if false {
        var v8 int64 = 989
        _ = v8
    }
    fmt.Println(v4, v4)
    var v9 int64 = 205
    var v10 float64 = ((float64(float32((0.1 * ((1.5 + (-2.25 * 3.14159)) * 1.5)))) * ((3.14159 + 1.5) - (float64((1.5 - 1.5)) * float64(1.5)))) - float64(1e10))
    if false {
        v11 := GenS1{F0: 2, F1: v6, F2: 32767}
        _ = v11
    } else {
        v6 += v6
        v12 := GenS1{F0: 699, F1: v6, F2: 0}
        _ = v12
    }
    v13 := GenS2{F0: v5, F1: ((v5 & v5) <= 469), F2: 1}
    v10 += (-2.25 + 0.1)
    v6 = v6
    if true {
        switch "hello" {
        case "\n\t":
            v1 += -328
        case "a":
            switch int(v1) {
            case 1:
                v9 -= int64(v4)
            case 418:
                v4 += v4
            case 4294967295:
                fmt.Println(v7, v6)
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v14 float32 = 3.14159
        _ = v14
    }
    var v15 float64 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
