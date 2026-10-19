package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 string; F1 int8; F2 uint }

func main() {
    var v1 float32 = -2.25
    v2 := []int{-90, -2147483648, 32768}
    v3 := []int{-697, -296, -548}
    v4 := make(map[int]uint32)
    v5 := []string{}
    var v6 string = ("foo bar" + "foo bar")
    var v7 string = v6
    var v8 uint32 = 787
    switch 42 {
    case 9223372036854775807:
        var v9 uint = uint(v8)
        _ = v9
    default:
        v8 = 4294967295
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        if true {
            if (!true) {
                var v11 uint = 73
                v11 -= (^(v11))
                _ = v11
            }
            switch fi10 {
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -2147483648:
                var v12 float32 = (((((float32(0.0) * v1) * v1) - float32(1e10)) + ((v1 - (float32(-2.25) * 3.14159)) - 1e10)) * v1)
                _ = v12
            default:
                var v13 int16 = -1
                _ = v13
            }
        }
        v6 = v7
        _ = fi10
    }
    v7 = v6
    v6 = v7
    v1 += 0.1
    if (0 > 776) {
        v8 += v8
        v14 := v4[1]
        if (false || true) {
            v15 := GenS1{F0: (v14 ^ (v14 & 65535)), F1: false}
            fmt.Println(v14)
            v16 := GenS1{F0: (v14 | 151), F1: true}
            _ = v15
            _ = v16
        } else {
            v4[682] = (^(v14))
        }
        _ = v14
    } else {
        if false {
            v1 += v1
            if true {
                if (1e10 != -2.25) {
                    var v17 bool = true
                    v18 := GenS1{F0: v8, F1: (v17 && (v17 && true))}
                    v1 = v1
                    _ = v17
                    _ = v18
                } else {
                    if ((118 != 1) || (255 <= 602)) {
                        var v19 uint32 = (v8 << ((248) & 63))
                        var v20 uint64 = 4294967295
                        _ = v19
                        _ = v20
                    }
                }
            }
        } else {
            v21 := GenS1{F0: (v8 + 1), F1: (1 >= 1000)}
            _ = v21
        }
    }
    v6 = (("a" + ((("hello" + v7) + "") + v7)) + v7)
    switch 1 {
    case 32768:
        fmt.Println(v6)
    case 128:
        fmt.Println(v6, v8)
    case -885:
        var v22 uint32 = v8
        _ = v22
    default:
        fmt.Println(v6, v1)
    }
    _ = len(v5)
    v4[int(v8)] = 2147483647
    var v23 uint8 = 42
    v24 := GenS2{F0: (v7 + "foo bar"), F1: 100, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
