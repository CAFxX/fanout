package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0() uint {
    var v5 uint64 = 0
    v6 := []int{0, -9223372036854775808, 4294967295}
    v5 += 2
    v6 = append(v6, 0)
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = 42
    var v3 int32 = 42
    var v4 int64 = 471
    var v7 uint16 = 446
    fmt.Println(v7)
    var v8 int8 = 127
    switch "hello" {
    case "hello":
        v9 := []int{127, -2147483648, 2147483647}
        _ = v9
    case "a":
        v7 = (v7 + v7)
    case "":
        var v10 uint16 = v7
        _ = v10
    }
    switch 0 {
    case -451:
        var v11 float64 = (((0.0 * -2.25) - (0.0 - -2.25)) - (1.5 + (3.14159 * (1.5 + 0.0))))
        _ = v11
    case 127:
        v12 := GenS1{F0: v4}
        _ = v12
    case 4294967295:
        switch ("a" + "hello") {
        case "":
            v7 -= v7
        case "\n\t":
            v7 -= (v7 & v7)
        default:
            var v13 float32 = 1e10
            _ = v13
        }
    default:
        v14 := GenS1{F0: v4}
        _ = v14
    }
    var v15 int64 = 42
    var v16 int8 = v8
    if false {
        v17 := []int{-2147483648, -1, 32768}
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 255, F1: (v7 - 32767), F2: (v15 % (((v15 + v4)) | 1))}
        v19 := GenS1{F0: -9223372036854775808}
        _ = v18
        _ = v19
    } else {
        fmt.Println(v4, v2)
        v20 := GenS1{F0: (v4 | (v4 - v4))}
        _ = v20
    }
    var v21 uint16 = v7
    if false {
        v15 += (v4 << ((65535) & 63))
    }
    v22 := []int{887, 42, 42}
    v1 += (1.5 + 1.5)
    if true {
        if true {
            v23 := GenS1{F0: (v15 >> ((42) & 63))}
            v21 = (v7 ^ v21)
            _ = v23
        } else {
            var v24 int8 = (v8 - v16)
            _ = v24
        }
        var v25 uint32 = 909
        switch 2147483647 {
        case 683:
            switch -1 {
            case -1:
                var v26 uint32 = (^(v25))
                _ = v26
            default:
                var v27 uint64 = 211
                _ = v27
            }
        case 42:
            v7 = (v21 / ((v21) | 1))
        default:
            v28 := GenS1{F0: v15}
            _ = v28
        }
        _ = v25
    }
    v15 = (v4 - int64(v8))
    var v29 uint = 490
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
