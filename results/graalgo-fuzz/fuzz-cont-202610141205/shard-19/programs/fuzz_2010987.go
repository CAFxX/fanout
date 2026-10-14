package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 int64) int32 {
    p0 -= p0
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 = p0
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    _ = p0
    return 226
}
func genFn1(p0 int, p1 bool, p2 int32) uint8 {
    v9 := []int{713}
    if ((p2 + (p2 << ((65535) & 63))) != (^(p2))) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: ("\n\t" + ("" + "foo bar"))}
            _ = fi10
            _ = v11
        }
        if (!p1) {
            for fi12 := 0; fi12 < 11; fi12++ {
                if p1 {
                    p2 = (-(p2))
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                _ = fi12
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if true {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                for fi13 := 0; fi13 < 12; fi13++ {
                    p0 += (fi13 / ((fi13) | 1))
                    if len(v9) > 0 {
                        v9[0] = -336
                    }
                    _ = fi13
                }
            }
        }
    } else {
        v14 := GenS1{F0: ""}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        _ = len(v9)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 122
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    v3 := []int{32768, 0}
    var v4 float64 = 0.0
    v5 := make(map[uint32]uint32)
    var v6 int64 = 9223372036854775807
    v4 = (v4 - 1.5)
    v1 -= v1
    if (false && (1 == int32(v6))) {
        v16 := GenS2{F0: 428}
        var v17 int8 = 1
        for fi18 := 0; fi18 < 15; fi18++ {
            if ((!(!true)) && true) {
                v19 := GenS3{F0: uint16(v1), F1: 0}
                var v20 uint64 = v1
                _ = v19
                _ = v20
            }
            v21 := GenS1{F0: ""}
            _ = fi18
            _ = v21
        }
        _ = v16
        _ = v17
    } else {
        if false {
            fmt.Println(v1, v4, v6)
        }
    }
    v22 := GenS4{F0: 0}
    v2[32767] = false
    v23 := GenS1{F0: "foo bar"}
    var v24 float32 = (-2.25 * -2.25)
    v25 := GenS4{F0: 4294967295}
    if true {
        v6 += v6
        v24 = (v24 / v24)
    }
    v1 = 448
    delete(v5, 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
