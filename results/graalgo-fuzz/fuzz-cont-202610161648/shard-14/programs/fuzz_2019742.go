package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 0
    v3 := []string{"", "foo bar"}
    var v4 int8 = -20
    var v5 uint8 = 0
    if (42 > 9223372036854775807) {
        switch 0 {
        case -2147483648:
            switch -681 {
            case -1:
                v6 := GenS1{F0: 159}
                _ = v6
            default:
                switch "" {
                case "":
                    v7 := GenS2{F0: -154, F1: v4, F2: 1}
                    _ = v7
                }
            }
        default:
            v8 := GenS3{F0: -9223372036854775808, F1: 0, F2: 470}
            _ = v8
        }
        var v9 uint8 = (v2 + uint8(v1))
        fmt.Println(v5, v4)
        _ = v9
    } else {
        v1 = v1
    }
    var v10 int32 = 408
    var v11 uint16 = v1
    var v12 uint8 = 1
    v12 += 2
    fmt.Println(v2, v2, v4)
    switch 127 {
    case 9223372036854775807:
        var v13 uint64 = 858
        _ = v13
    }
    var v14 int16 = 1000
    var v15 uint32 = 2
    v10 = (-(v10))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v11)
    var v16 uint8 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
