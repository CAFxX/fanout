package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 uint, p1 int64) int16 {
    v6 := GenS1{F0: (-(p1)), F1: -2147483648, F2: 65535}
    if true {
        if true {
            p1 = (p1 >> ((255) & 63))
        } else {
            v7 := GenS1{F0: p1, F1: 338, F2: 2147483648}
            var v8 string = ""
            _ = v7
            _ = v8
        }
        var v9 bool = (2 <= p0)
        _ = v9
    }
    var v10 bool = true
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return -369
}

func main() {
    v1 := []uint32{}
    v2 := []string{}
    var v3 uint64 = 953
    v4 := []uint32{929}
    v5 := make(map[int]uint32)
    _ = len(v5)
    fmt.Println(v3)
    var v11 int16 = 42
    _ = len(v2)
    var v12 uint32 = 2
    var v13 int8 = 0
    if false {
        if false {
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi14, v12)
                    _ = fi15
                }
                _ = fi14
            }
        } else {
            v16 := GenS1{F0: -1, F1: 47, F2: (v12 | uint32(v3))}
            _ = v16
        }
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    v12 = 0
                    _ = fi18
                }
                v19 := GenS1{F0: -9223372036854775808, F1: 590, F2: (v12 ^ 0)}
                fmt.Println(fi17, fi17, v13)
                _ = fi17
                _ = v19
            }
        }
        v20 := GenS1{F0: 42, F1: -402, F2: (v12 | 2147483647)}
        _ = v20
    } else {
        v13 = (v13 / ((-128) | 1))
        v3 = v3
        v3 -= v3
    }
    fmt.Println(v3, v12, v11)
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            v12 -= (v12 | (^(v12)))
        }
        var v22 float64 = 1e10
        _ = fi21
        _ = v22
    }
    if len(v1) > 0 {
        v1[0] = v12
    }
    var v23 int32 = 42
    fmt.Println(v23)
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS2{F0: v11, F1: (^(v23))}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 int = (fi26 % ((fi26) | 1))
        var v28 int16 = 171
        var v29 uint32 = 1
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v3 -= uint64(v12)
    v30 := GenS2{F0: (v11 >> ((425) & 63)), F1: (v23 | v23)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
