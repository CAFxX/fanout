package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
func genFn0(p0 int32) int8 {
    p0 -= (^(p0))
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = 915
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    _ = p0
    return 30
}
func genFn1(p0 int8, p1 uint64) uint32 {
    v8 := []int{1}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = 127
    v3 := make(map[int]uint32)
    var v4 int = 32768
    var v5 int8 = v2
    var v6 string = "foo bar"
    var v9 uint16 = v1
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v5 -= (v2 >> ((4294967295) & 63))
            fmt.Println(v5, v4)
            var v12 float32 = float32((0.1 - (0.1 + float64(0.0))))
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := []int{42}
    v5 += -1
    if (444 == 429) {
        for fi14 := 0; fi14 < 4; fi14++ {
            v13 = append(v13, (-(v4)))
            if len(v13) > 0 {
                v13[0] = (-(fi14))
            }
            v4 = (fi14 << ((18446744073709551615) & 63))
            _ = fi14
        }
    }
    var v15 int64 = 1
    var v16 int8 = 127
    v1 -= 255
    switch v4 {
    case 180:
        var v17 bool = ((v9 * v9) == (^(v1)))
        _ = v17
    case 9223372036854775807:
        if (187 <= 128) {
            v18 := GenS1{F0: 508, F1: int16(v15)}
            v19 := GenS2{F0: (308 == 0), F1: v15, F2: 0}
            _ = v18
            _ = v19
        }
    }
    if ((v15 >= v15) && (1 > 222)) {
        v20 := GenS1{F0: v9, F1: 0}
        var v21 uint64 = 1
        _ = v20
        _ = v21
    } else {
        var v22 int32 = 0
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
