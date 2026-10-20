package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int16; F1 uint8; F2 int }
type GenS6 struct { F0 int; F1 uint64 }
func genFn0(p0 int32) bool {
    v4 := GenS1{F0: 42}
    if true {
        v5 := GenS1{F0: 0}
        if (false && false) {
            if (1 == -1) {
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 float64 = 1.5
                    v8 := GenS2{F0: 42}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                fmt.Println(p0)
            } else {
                var v9 uint8 = 159
                v9 += (v9 & v9)
                _ = v9
            }
            fmt.Println(p0)
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int8 = -128
            fmt.Println(v11)
            if (0 >= 447) {
                if (42 >= 1) {
                    v11 += 0
                    v11 = (v11 & v11)
                }
            } else {
                for fi12 := 0; fi12 < 6; fi12++ {
                    p0 = p0
                    _ = fi12
                }
                switch 1 {
                case 32767:
                    p0 = p0
                case 32768:
                    fmt.Println(fi10, p0)
                case 42:
                    if false {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            fmt.Println(fi10, fi13, p0)
                            _ = fi13
                        }
                    }
                default:
                    var v14 bool = (true || false)
                    _ = v14
                }
                var v15 uint16 = 2
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        _ = v5
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(p0, fi16)
            v17 := []int{(fi16 % ((fi16) | 1)), fi16}
            _ = fi16
            _ = v17
        }
        v18 := GenS3{F0: true, F1: 255, F2: 2}
        v19 := []int{-57, 127, -327}
        _ = v18
        _ = v19
    }
    switch -713 {
    case 9223372036854775807:
        fmt.Println(p0)
    case -730:
        v20 := GenS3{F0: (false || false), F1: 128, F2: 255}
        _ = v20
    case 4294967295:
        if (-2.25 <= float64(((1e10 / 1.0) - float32(0.0)))) {
            var v21 int = 32767
            _ = v21
        }
    default:
        var v22 uint32 = 2147483648
        _ = v22
    }
    _ = p0
    _ = v4
    return false
}
func genFn1(p0 uint8, p1 bool, p2 uint64) uint64 {
    if (uint(p0) >= 65535) {
        v23 := GenS4{F0: p0, F1: 1}
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if p1 {
            v25 := []int{(fi24 ^ (fi24 - 263)), (fi24 % ((-1) | 1)), (fi24 ^ (fi24 | (-(fi24))))}
            _ = v25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(p2, p0, fi24)
            _ = fi26
        }
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 int8 = 100
    var v2 int64 = -785
    var v3 uint8 = 42
    var v27 uint16 = 32767
    v28 := GenS5{F0: 32767, F1: (v3 << ((276) & 63)), F2: -1}
    var v29 float32 = -2.25
    v30 := GenS6{F0: 821, F1: genFn1((v3 % ((v3) | 1)), ((true && (true && ("" < "foo bar"))) || (v3 != v3)), 0)}
    var v31 int16 = int16(v1)
    var v32 uint8 = (v3 + (v3 & (v3 - (^(v3)))))
    for fi33 := 0; fi33 < 2; fi33++ {
        v1 -= v1
        _ = fi33
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
