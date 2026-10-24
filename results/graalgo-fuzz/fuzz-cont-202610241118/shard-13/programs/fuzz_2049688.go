package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 int16, p1 uint32, p2 uint64) int16 {
    v6 := []int{-9223372036854775808, 0, -9223372036854775808, 767}
    var v7 uint64 = (p2 | p2)
    if false {
        p1 = uint32(p2)
        p0 = -1
        p0 -= (p0 + p0)
    } else {
        if len(v6) > 0 {
            v6[0] = 32767
        }
        v6 = append(v6, -1)
        _ = len(v6)
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: int(p2), F1: 41}
        if (!false) {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 string = ""
                var v12 int8 = 127
                _ = fi10
                _ = v11
                _ = v12
            }
            p1 += (p1 * uint32(v7))
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 42
    var v3 bool = false
    var v4 int8 = 100
    var v5 uint = 4294967295
    var v13 int8 = (v4 + (v4 << (((v5 + 65535)) & 63)))
    v14 := GenS2{F0: 1}
    var v15 uint32 = (^(v2))
    fmt.Println(v15, v4, v5)
    for fi16 := 0; fi16 < 15; fi16++ {
        v13 -= (v13 % ((v4) | 1))
        _ = fi16
    }
    v17 := GenS1{F0: -1, F1: 42}
    v18 := GenS1{F0: -1, F1: 0}
    v19 := []int{1}
    for fi20 := 0; fi20 < 14; fi20++ {
        v4 -= (v13 | v4)
        _ = fi20
    }
    delete(v1, (v2 >> ((v5) & 63)))
    if (!false) {
        _ = len(v1)
        switch 128 {
        case 70:
            _ = len(v19)
        }
    } else {
        var v21 int16 = 365
        _ = v21
    }
    delete(v1, 0)
    if v3 {
        var v22 uint16 = 334
        v23 := GenS3{F0: (v5 | v5), F1: 0}
        _ = v22
        _ = v23
    } else {
        v5 = 65535
        _ = len(v1)
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        switch fi24 {
        case 128:
            v19 = append(v19, 0)
        case 127:
            if (903 < 32767) {
                var v25 int16 = -611
                if v3 {
                    fmt.Println(v15, v4)
                    v4 = v4
                    var v26 float64 = (0.0 / float64((-2.25 - 0.1)))
                    _ = v26
                }
                v27 := GenS4{F0: 1, F1: v5, F2: 255}
                _ = v25
                _ = v27
            } else {
                fmt.Println(v13, v4, v2)
            }
        }
        _ = fi24
    }
    v15 -= 231
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
