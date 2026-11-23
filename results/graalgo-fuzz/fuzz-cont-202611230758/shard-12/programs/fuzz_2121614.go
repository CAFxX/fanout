package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
func genFn0(p0 uint, p1 int32) string {
    v7 := []int{9223372036854775807, -282}
    v8 := GenS1{F0: 841, F1: 1, F2: 65535}
    var v9 int32 = 1
    if (1 <= 1) {
        var v10 int = 241
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 867
    var v3 int = 700
    var v4 int8 = 30
    var v5 string = "hello"
    var v6 uint8 = 127
    var v11 float64 = (1e10 + 1.5)
    var v12 uint = 18446744073709551615
    v13 := []int{(-(v3)), (^(v3)), -2147483648}
    switch v3 {
    case 127:
        v14 := GenS1{F0: v2, F1: 2, F2: v2}
        _ = v14
    default:
        switch v5 {
        case "foo bar":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            v4 = v4
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v2, v1, v2)
        _ = fi15
    }
    v4 = (v4 - v4)
    if len(v13) > 0 {
        v13[0] = (v3 % (((v3 ^ 902)) | 1))
    }
    var v16 int64 = -1
    switch v5 {
    case "foo bar":
        fmt.Println(v2, v11)
    case "hello":
        fmt.Println(v4, v11)
    default:
        v17 := GenS1{F0: (v2 | 65535), F1: 2, F2: 255}
        _ = v17
    }
    _ = len(v13)
    var v18 uint = v12
    v6 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
