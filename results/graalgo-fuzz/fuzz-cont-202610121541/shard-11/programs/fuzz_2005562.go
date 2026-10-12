package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float64 = (0.0 + 0.1)
    var v2 uint16 = 178
    var v3 uint64 = 0
    var v4 bool = (false && (9223372036854775807 > -9223372036854775808))
    var v5 int8 = 0
    v6 := []uint32{}
    switch "" {
    case "":
        fmt.Println(v4, v1, v4)
    default:
        v6 = append(v6, 318)
    }
    var v7 uint64 = v3
    if false {
        v1 += v1
    } else {
        switch 3 {
        case 880:
            v8 := GenS1{F0: "a", F1: 462, F2: 18446744073709551615}
            _ = v8
        case -145:
            v9 := GenS2{F0: 42}
            _ = v9
        default:
            fmt.Println(v5, v5)
        }
    }
    v10 := GenS3{F0: 42, F1: (v2 | (v2 % ((v2) | 1))), F2: v4}
    _ = len(v6)
    var v11 bool = v4
    var v12 bool = (42 == uint32(v7))
    for fi13 := 0; fi13 < 6; fi13++ {
        if len(v6) > 0 {
            v6[0] = 2147483648
        }
        _ = fi13
    }
    v5 += v5
    fmt.Println(v4, v12)
    v14 := GenS3{F0: -2147483648, F1: (v2 / ((v2) | 1)), F2: true}
    switch -2147483648 {
    case 1:
        if len(v6) > 0 {
            v6[0] = 42
        }
    }
    if false {
        if (true || v11) {
            var v15 int = 2147483647
            _ = v15
        } else {
            if (!(v5 < (v5 - v5))) {
                v12 = (!(255 < v2))
                v1 -= ((1.5 / 3.14159) / 0.1)
            } else {
                v3 -= (v7 ^ (v3 - 414))
            }
            var v16 uint = 65535
            _ = v16
        }
        v6 = append(v6, 1)
    }
    v17 := GenS4{F0: (80 <= 65535), F1: (-1 != 1)}
    for fi18 := 0; fi18 < 13; fi18++ {
        if v12 {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 bool = false
                v11 = (-9223372036854775808 == 1)
                var v21 uint8 = 255
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi18
    }
    v22 := GenS5{F0: 2, F1: -1}
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 string = ""
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
