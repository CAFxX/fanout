package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }

func main() {
    var v1 string = "hello"
    v2 := []string{}
    var v3 int64 = 42
    var v4 uint16 = 65535
    v5 := []int{-472, -1}
    v6 := GenS1{F0: int16(v3), F1: 65535, F2: v4}
    fmt.Println(v3)
    var v7 int16 = 349
    fmt.Println(v1)
    v5 = append(v5, 4294967295)
    if ((true && true) && (42 < 42)) {
        var v8 float32 = -2.25
        _ = v8
    }
    v7 = v7
    if false {
        if (128 != 255) {
            v1 = v1
            var v9 uint64 = 2
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                v7 -= v7
                v3 -= 42
                var v11 uint16 = v4
                _ = fi10
                _ = v11
            }
        }
        if true {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 string = v1
                var v14 uint16 = v4
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 int = (-(fi15))
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            v17 := GenS1{F0: int16(v3), F1: (v4 + v4), F2: (v4 % (((v4 ^ v4)) | 1))}
            _ = len(v5)
            _ = v17
        }
        var v18 uint16 = (v4 >> ((18446744073709551615) & 63))
        _ = v18
    } else {
        switch 2147483647 {
        case 127:
            var v19 uint16 = (^(v4))
            _ = v19
        case 309:
            fmt.Println(v3, v3, v1)
        default:
            var v20 bool = (2 >= 160)
            _ = v20
        }
    }
    var v21 int32 = 385
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v22 int16 = int16(v21)
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS1{F0: (v7 ^ (v7 + v22)), F1: v4, F2: 1}
        v7 -= v7
        _ = fi23
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = ("a" + "foo bar")
    }
    switch 113 {
    case 127:
        fmt.Println(v4, v22, v21)
    case -951:
        v25 := GenS2{F0: v22, F1: 1000, F2: v4}
        _ = v25
    case 327:
        _ = len(v5)
    default:
        fmt.Println(v1, v21, v3)
    }
    v21 = (v21 + (v21 % ((v21) | 1)))
    var v26 float32 = 0.1
    v5 = append(v5, int(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
