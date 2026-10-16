package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 uint8, p1 bool, p2 uint32) int64 {
    switch ((("\n\t" + "") + "hello") + "") {
    case "foo bar":
        var v7 int32 = 513
        _ = v7
    case "a":
        v8 := GenS1{F0: 0, F1: 42, F2: -115}
        _ = v8
    default:
        v9 := GenS1{F0: -12, F1: int64(p2), F2: -591}
        _ = v9
    }
    p0 += (p0 << ((624) & 63))
    var v10 bool = p1
    v11 := GenS1{F0: 42, F1: 0, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 uint8) int64 {
    fmt.Println(p2)
    for fi12 := 0; fi12 < 4; fi12++ {
        p2 = (^(p2))
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    var v14 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return int64(p1)
}

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 7
    v3 := make(map[string]int)
    var v4 int32 = 0
    v5 := make(map[int]bool)
    var v6 int8 = 68
    v15 := GenS1{F0: (-(v6)), F1: -9223372036854775808, F2: 218}
    v16 := GenS2{F0: 255, F1: 1}
    var v17 int64 = 337
    for fi18 := 0; fi18 < 11; fi18++ {
        v17 += (v17 % ((9223372036854775807) | 1))
        _ = fi18
    }
    delete(v5, 42)
    if true {
        var v19 int = 114
        _ = v19
    }
    v20 := []int{32767, 128, -1, -2147483648}
    v4 += (v4 | v4)
    v21 := GenS3{F0: 0, F1: v2, F2: 42}
    v22 := GenS4{F0: ("\n\t" + ""), F1: 812}
    v17 -= v17
    for fi23 := 0; fi23 < 7; fi23++ {
        if false {
            v5[(fi23 - (fi23 | (fi23 << ((255) & 63))))] = true
            if ((false || ((((v1 + (v1 - (v1 / float64((((3.14159 + (1e10 + (3.14159 - (float32(v1) / (-2.25 * 1.5))))) - float32(0.1)) - 1.5))))) / -2.25) >= v1) || true)) && ((v17 * v17) >= 1)) {
                if true {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v17 = (-(v17))
                        v25 := GenS3{F0: 123, F1: (v2 | (v2 - (v2 / ((v2) | 1)))), F2: (v4 | 0)}
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(fi23)
                    v2 -= (v2 * v2)
                }
                switch (fi23 % ((4294967295) | 1)) {
                case -73:
                    v1 -= 0.1
                case 43:
                    v26 := GenS1{F0: v6, F1: (v17 % (((v17 ^ v17)) | 1)), F2: 360}
                    _ = v26
                case 128:
                    v17 = (v17 >> ((596) & 63))
                }
                fmt.Println(v2)
            }
            if false {
                v27 := GenS5{F0: (v17 * v17), F1: 42}
                for fi28 := 0; fi28 < 7; fi28++ {
                    v2 += (v2 * 0)
                    v4 = v4
                    v6 = -1
                    _ = fi28
                }
                _ = v27
            }
        }
        _ = fi23
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        if (false || true) {
            fmt.Println(v4, fi29)
        }
        _ = fi29
    }
    if (v6 < (v6 ^ 0)) {
        fmt.Println(v2, v17)
    }
    v1 -= v1
    var v30 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
