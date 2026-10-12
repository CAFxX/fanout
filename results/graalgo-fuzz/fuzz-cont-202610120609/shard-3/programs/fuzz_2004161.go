package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 uint16, p1 int) uint16 {
    v5 := GenS1{F0: 42}
    if (p1 <= p1) {
        v6 := []int{9223372036854775807, p1, int(p0), (p1 + 9223372036854775807)}
        var v7 uint16 = p0
        _ = v6
        _ = v7
    }
    var v8 int16 = 42
    if false {
        var v9 int32 = 1
        fmt.Println(v9, p1, p0)
        _ = v9
    } else {
        p1 = (-(p1))
    }
    v10 := GenS2{F0: 0, F1: 212, F2: p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 127
    var v2 int16 = 554
    var v3 bool = true
    v4 := []uint32{0, uint32(v1), 2147483648}
    fmt.Println(v2, v2, v2)
    if v3 {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS2{F0: v2, F1: 42, F2: 3}
            switch "foo bar" {
            case "hello":
                var v13 bool = (v3 || v3)
                _ = v13
            case "":
                _ = len(v4)
            default:
                var v14 float64 = 0.0
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    } else {
        v15 := GenS2{F0: v2, F1: int32(v1), F2: 645}
        v1 = (v1 * 173)
        v2 += (v2 / (((v2 - (v2 & v2))) | 1))
        _ = v15
    }
    v1 -= 255
    var v16 uint64 = 1
    switch 1 {
    case 128:
        var v17 uint8 = (v1 ^ v1)
        _ = v17
    case 0:
        if false {
            _ = len(v4)
            v1 = (v1 - uint8(v16))
            fmt.Println(v3, v3)
        }
    }
    v4 = append(v4, 42)
    if v3 {
        v2 = 1000
    }
    if (4294967295 >= 1) {
        if v3 {
            if v3 {
                switch 32768 {
                case 793:
                    v16 = 2
                default:
                    v18 := GenS3{F0: -1}
                    _ = v18
                }
                v3 = ((!(629 == 1)) || ("a" != ("a" + "")))
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v16 = (^(v16))
        }
        var v19 uint32 = 42
        _ = v19
    } else {
        fmt.Println(v2, v2)
        if false {
            for fi20 := 0; fi20 < 17; fi20++ {
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
                if (1 != -1) {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v16 -= v16
                        fmt.Println(v1, v16, v3)
                        v22 := GenS4{F0: 32768}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        var v24 int = fi20
                        fmt.Println(fi23, v16)
                        var v25 int16 = (^(v2))
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v2 = v2
                }
                _ = fi20
            }
        } else {
            v3 = ((v1 * v1) != v1)
        }
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
