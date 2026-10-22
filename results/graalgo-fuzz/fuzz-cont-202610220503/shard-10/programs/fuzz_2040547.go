package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
func genFn0(p0 int16) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int32 = int32(fi7)
        fmt.Println(p0, v8)
        var v9 int8 = -128
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = -2147483648
    switch "foo bar" {
    case "\n\t":
        var v11 uint32 = 2147483647
        _ = v11
    default:
        p0 += (-(p0))
    }
    var v12 uint32 = 297
    v10 = v10
    _ = p0
    _ = v10
    _ = v12
    return v12
}
func genFn1(p0 int, p1 string) int16 {
    v13 := GenS1{F0: -212, F1: (p0 | (p0 >> ((84) & 63)))}
    fmt.Println(p1)
    var v14 bool = false
    for fi15 := 0; fi15 < 14; fi15++ {
        p1 = p1
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 uint32 = 949
    v2 := make(map[string]string)
    var v3 string = "a"
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 0
    v6 := []int{}
    v3 = (v3 + "a")
    var v16 uint64 = (v4 ^ v4)
    v16 = (^(v4))
    v2[(v3 + ("\n\t" + v3))] = v3
    if true {
        var v17 int32 = -136
        for fi18 := 0; fi18 < 5; fi18++ {
            v16 += (v16 * (v4 / (((^(v4))) | 1)))
            if false {
                v4 += (^(v16))
            }
            _ = fi18
        }
        v2["hello"] = "foo bar"
        _ = v17
    }
    v5 = v5
    if ((v4 << ((255) & 63)) == v4) {
        v1 -= v1
        var v19 int64 = int64(v4)
        _ = v19
    } else {
        switch 611 {
        case -779:
            v20 := GenS2{F0: 1, F1: 0, F2: (0 < genFn1(32768, v3))}
            _ = v20
        default:
            var v21 float64 = -2.25
            _ = v21
        }
    }
    var v22 int64 = -289
    var v23 int = 889
    var v24 int64 = (v22 >> ((65535) & 63))
    var v25 bool = true
    fmt.Println(v1, v5, v25)
    v22 -= -134
    v26 := GenS2{F0: -1, F1: (v24 + (-(v24))), F2: v25}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
