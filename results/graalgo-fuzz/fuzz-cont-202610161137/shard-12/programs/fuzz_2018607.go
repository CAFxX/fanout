package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64; F1 bool }
func genFn0(p0 string, p1 string, p2 uint32) string {
    p2 += 2
    p2 -= 2
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: 0, F1: (-529 < (-(fi6))), F2: 32767}
            var v8 uint64 = 350
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    var v9 bool = ((0.1 + float64(float32((1.5 * (0.1 * 3.14159))))) == 0.0)
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return ""
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 int = 32767
    var v5 int16 = -1
    v4 -= (v4 >> ((1) & 63))
    fmt.Println(v4, v4)
    for fi10 := 0; fi10 < 1; fi10++ {
        v5 += (v5 & 0)
        v11 := GenS1{F0: v5, F1: true, F2: int16(fi10)}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (-(v5)), F1: true, F2: -825}
    v1 = (^(v1))
    switch "hello" {
    case "":
        v5 -= (v5 << ((327) & 63))
    case "\n\t":
        v13 := GenS1{F0: (v5 >> ((333) & 63)), F1: (!true), F2: (v5 & int16(v1))}
        _ = v13
    }
    v14 := GenS2{F0: (v1 & 1), F1: 744}
    v15 := GenS3{F0: (v1 - (v1 % ((v1) | 1))), F1: (true && false)}
    v1 = v1
    if (!true) {
        v1 = (v1 - 9223372036854775807)
        v16 := GenS2{F0: 1, F1: 435}
        _ = v16
    } else {
        v17 := []int{(^(v4)), v4, (v4 | v4), (v4 ^ v4)}
        _ = v17
    }
    if (128 < 255) {
        var v18 bool = (60 != v5)
        var v19 uint8 = 1
        _ = v18
        _ = v19
    } else {
        v20 := v3[975]
        var v21 int64 = v1
        var v22 uint16 = 0
        _ = v20
        _ = v21
        _ = v22
    }
    v5 += (v5 / ((42) | 1))
    v23 := GenS2{F0: v1, F1: v5}
    if false {
        v5 -= v5
        v4 += (v4 ^ (v4 >> ((18446744073709551615) & 63)))
    } else {
        var v24 int64 = v1
        var v25 uint16 = 32767
        v5 -= 1000
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: v1, F1: (^(v5))}
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 uint = 255
            v29 -= 157
            if false {
                if false {
                    var v30 uint32 = 1
                    _ = v30
                } else {
                    v5 -= v5
                }
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v1 += (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
