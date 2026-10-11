package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }

func main() {
    var v1 uint32 = 42
    var v2 float64 = -2.25
    var v3 string = ("foo bar" + "hello")
    var v4 int8 = -67
    var v5 uint32 = (v1 * (v1 + v1))
    var v6 int8 = (v4 ^ (v4 / (((v4 + -1)) | 1)))
    v7 := []int{32767, 127, -710}
    switch 179 {
    case 2147483647:
        var v8 uint32 = 1
        _ = v8
    default:
        var v9 uint64 = 251
        _ = v9
    }
    var v10 uint16 = 888
    var v11 uint8 = 255
    var v12 uint = 572
    v1 = (v1 ^ 0)
    if false {
        var v13 int16 = 1
        switch 42 {
        case 1:
            var v14 uint = v12
            _ = v14
        case 0:
            v1 = (v5 >> ((v12) & 63))
        case 42:
            v3 = v3
        default:
            v15 := GenS1{F0: 9223372036854775807, F1: (v12 * (v12 * 960)), F2: 42}
            _ = v15
        }
        if ((((v3 + v3) + v3) + v3) > "hello") {
            var v16 uint8 = (^(v11))
            v17 := GenS1{F0: -121, F1: uint(v1), F2: 1}
            var v18 string = ("a" + (v3 + v3))
            _ = v16
            _ = v17
            _ = v18
        } else {
            var v19 int = 32
            _ = v19
        }
        _ = v13
    } else {
        var v20 uint32 = v1
        var v21 int8 = int8(v10)
        var v22 int8 = v4
        _ = v20
        _ = v21
        _ = v22
    }
    v7 = append(v7, 1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v23 int64 = 564
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
