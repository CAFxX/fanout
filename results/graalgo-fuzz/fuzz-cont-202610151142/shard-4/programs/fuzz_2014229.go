package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    var v4 uint32 = 490
    if (0 >= 255) {
        if (799 == 102) {
            _ = len(v3)
            if (("" + "a") > "foo bar") {
                v5 := GenS1{F0: 1}
                _ = v5
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v6 uint32 = 2
                _ = v6
            }
            v7 := GenS1{F0: 1}
            _ = v7
        } else {
            v4 = v4
        }
        v2 = 80
    } else {
        var v8 uint8 = 191
        v9 := GenS1{F0: 441}
        v4 = 728
        _ = v8
        _ = v9
    }
    switch 1 {
    case -652:
        v2 = v2
    case 128:
        var v10 string = "\n\t"
        _ = v10
    case -9223372036854775808:
        var v11 int16 = -32768
        _ = v11
    default:
        var v12 float64 = (((1.5 * 0.0) / 1e10) / ((3.14159 + float64((float32((-2.25 / 3.14159)) + (1e10 * float32(0.0))))) + 0.1))
        _ = v12
    }
    var v13 int16 = 0
    if len(v1) > 0 {
        v1[0] = v4
    }
    fmt.Println(v4, v2, v2)
    fmt.Println(v2, v4, v13)
    v14 := GenS1{F0: -215}
    v15 := GenS2{F0: int16(v2)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
