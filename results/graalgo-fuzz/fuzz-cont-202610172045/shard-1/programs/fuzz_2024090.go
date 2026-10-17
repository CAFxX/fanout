package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
func genFn0(p0 bool) bool {
    var v7 float32 = 3.14159
    switch 209 {
    case 0:
        switch -9223372036854775808 {
        case 32767:
            v7 = float32(0.0)
        case -1:
            v8 := GenS1{F0: 42, F1: 255, F2: "a"}
            _ = v8
        default:
            if ((!p0) || p0) {
                var v9 float32 = -2.25
                p0 = ((812 != 0) && p0)
                _ = v9
            }
        }
    }
    if p0 {
        for fi10 := 0; fi10 < 9; fi10++ {
            v7 += v7
            _ = fi10
        }
        var v11 int8 = 106
        _ = v11
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    v1 := []uint32{1, 2147483647, 2147483648, 716}
    v2 := make(map[int]bool)
    var v3 bool = (100 == -1)
    v4 := []int{117, 32768, -760}
    var v5 uint16 = 2
    var v6 uint16 = v5
    if (1 != 4294967295) {
        var v12 uint64 = 1
        switch -405 {
        case -1:
            var v13 uint16 = 278
            _ = v13
        }
        _ = v12
    }
    v14 := GenS1{F0: 42, F1: 1, F2: "a"}
    for fi15 := 0; fi15 < 8; fi15++ {
        if true {
            if len(v4) > 0 {
                v4[0] = fi15
            }
            if true {
                fmt.Println(v3, v5)
            }
        } else {
            var v16 int8 = 100
            var v17 int64 = 0
            if v3 {
                v17 = -730
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 bool = ((-2147483648 <= -2147483648) || true)
                    var v20 float32 = 0.0
                    v21 := v2[fi15]
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v22 := v2[fi15]
                _ = v22
            }
            _ = v16
            _ = v17
        }
        var v23 uint64 = 1
        var v24 int = fi15
        _ = fi15
        _ = v23
        _ = v24
    }
    v5 = (v6 + v5)
    v25 := GenS1{F0: 1, F1: 185, F2: "a"}
    v3 = v3
    fmt.Println(v3)
    v26 := GenS2{F0: 451, F1: "\n\t", F2: -1}
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
