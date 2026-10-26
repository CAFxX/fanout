package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }

func main() {
    var v1 int8 = 42
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    var v5 uint = 2
    var v6 uint64 = 42
    var v7 float64 = ((1e10 * float64((1.5 / (float32(3.14159) + (-2.25 * 0.0))))) / (-2.25 * ((1.5 / 1.5) + 0.0)))
    v8 := v4[4294967295]
    v6 = v6
    if (v7 > float64(-2.25)) {
        for fi9 := 0; fi9 < 16; fi9++ {
            v1 = 93
            var v10 uint8 = 42
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 1}
        var v12 uint8 = 8
        _ = v11
        _ = v12
    } else {
        v2 = append(v2, 2147483648)
    }
    if ((v8 | (v8 & (v8 + (v8 & 1)))) < (v8 & 1)) {
        if ("foo bar" < ("" + "a")) {
            if ("foo bar" >= "foo bar") {
                var v13 uint32 = 2
                _ = v13
            } else {
                if (float64(1e10) != v7) {
                    v6 -= 1
                }
            }
        }
        fmt.Println(v6, v5)
    }
    v14 := v4[969]
    v15 := v4[2147483647]
    for fi16 := 0; fi16 < 1; fi16++ {
        v5 = uint(v6)
        if (!true) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                fmt.Println(fi16)
                _ = fi17
            }
        } else {
            v7 = (0.0 + v7)
        }
        _ = fi16
    }
    if true {
        var v18 uint32 = 445
        var v19 uint = v5
        fmt.Println(v18, v19, v14)
        _ = v18
        _ = v19
    }
    _ = len(v4)
    _ = len(v2)
    var v20 uint16 = 878
    v21 := GenS2{F0: -128, F1: (^(v6)), F2: -68}
    if true {
        if len(v2) > 0 {
            v2[0] = 834
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
