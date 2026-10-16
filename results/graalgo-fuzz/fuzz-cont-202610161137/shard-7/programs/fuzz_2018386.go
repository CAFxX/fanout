package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 int16) bool {
    v6 := []int{4294967295, -9223372036854775808, p0, (p0 - p0)}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 0
        var v9 uint32 = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    return false
}
func genFn1(p0 int64, p1 bool, p2 uint16) uint {
    if p1 {
        for fi10 := 0; fi10 < 12; fi10++ {
            if (!(!((642 <= 0) && p1))) {
                var v11 int64 = (p0 ^ (p0 >> ((42) & 63)))
                var v12 int8 = 0
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 128, F1: false, F2: p2}
                _ = v13
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS2{F0: uint64(p2)}
                _ = fi14
                _ = v15
            }
            _ = fi10
        }
    } else {
        var v16 string = ("foo bar" + "\n\t")
        v16 = "a"
        _ = v16
    }
    switch 1 {
    case 9223372036854775807:
        v17 := []int{42, 719, 42}
        _ = v17
    case -742:
        p0 = p0
    case -148:
        fmt.Println(p2, p0, p2)
    }
    p1 = (p1 && ((-2.25 == (0.1 + (((1.5 + float32(3.14159)) + (0.0 / 1e10)) / 0.1))) || true))
    var v18 uint32 = 65535
    var v19 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 240
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 825
    v3 := make(map[int]uint32)
    var v4 uint8 = 127
    var v5 float64 = (3.14159 / 1.5)
    v20 := v3[0]
    v21 := v3[int(v2)]
    v22 := GenS2{F0: (v2 + v2)}
    v23 := GenS2{F0: v2}
    var v24 uint32 = 104
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(fi25, v2)
        var v26 int16 = 1000
        _ = fi25
        _ = v26
    }
    var v27 int16 = -31
    v21 = v24
    delete(v3, 32768)
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int16 = 64
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 65535
    fmt.Println(v20, v20, v5)
    v31 := []int{-9223372036854775808, -9223372036854775808, 128}
    v2 += v2
    v24 += v21
    for fi32 := 0; fi32 < 13; fi32++ {
        v20 = 4294967295
        _ = fi32
    }
    var v33 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
}
