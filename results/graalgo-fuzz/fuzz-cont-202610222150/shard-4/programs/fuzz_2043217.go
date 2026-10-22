package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 int64, p1 uint16) uint32 {
    var v7 uint8 = 255
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    return 73
}
func genFn1(p0 int64, p1 int8) int8 {
    p0 -= (p0 | p0)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: (p1 / (((p1 | 127)) | 1)), F1: 0, F2: 1}
        p0 = (p0 << ((65535) & 63))
        for fi10 := 0; fi10 < 17; fi10++ {
            p1 = p1
            p1 += (p1 % ((p1) | 1))
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int8 = p1
        var v13 int64 = (p0 << ((65535) & 63))
        for fi14 := 0; fi14 < 4; fi14++ {
            if true {
                v15 := []int{(fi11 / (((-(fi11))) | 1)), (fi14 + (-(fi14)))}
                _ = v15
            }
            fmt.Println(fi11, fi11, v13)
            fmt.Println(v13)
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v16 := GenS1{F0: (p1 & p1), F1: 1, F2: 32767}
    _ = p0
    _ = p1
    _ = v16
    return (p1 + p1)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    v3 := make(map[int]bool)
    var v4 string = ("" + "")
    v5 := []string{v4, v4, "hello", ""}
    v6 := make(map[int]string)
    var v17 int8 = -6
    v3[840] = true
    if false {
        v4 = v4
    }
    if (32767 == 818) {
        var v18 uint16 = 42
        if ((-267 != v2) || (65535 == 2147483647)) {
            v2 = (v2 ^ v2)
            v5 = append(v5, "foo bar")
            var v19 int64 = -976
            _ = v19
        } else {
            fmt.Println(v17, v18, v4)
            v20 := GenS1{F0: 42, F1: 1, F2: 0}
            _ = v20
        }
        _ = v18
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int16 = 0
            if true {
                var v23 uint32 = 362
                v24 := GenS2{F0: 13, F1: (v22 ^ (v22 - (v22 + v22)))}
                _ = v23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        v2 -= (-(v2))
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS3{F0: 1, F1: 42, F2: 128}
        v27 := GenS2{F0: 1, F1: 42}
        fmt.Println(v2, v2)
        _ = fi25
        _ = v26
        _ = v27
    }
    _ = len(v5)
    v17 = v17
    var v28 float64 = (1e10 * ((0.1 * 1.5) - -2.25))
    fmt.Println(v28, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
