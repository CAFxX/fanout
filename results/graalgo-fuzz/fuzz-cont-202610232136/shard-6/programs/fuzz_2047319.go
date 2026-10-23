package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 bool; F2 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 int) uint64 {
    var v5 bool = ((1 < p1) || (!(42 >= 0)))
    var v6 uint = 42
    fmt.Println(v5, v6, p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 540
}

func main() {
    var v1 int16 = -350
    var v2 uint64 = 4294967295
    var v3 bool = false
    v4 := []uint32{}
    if v3 {
        v1 = (v1 - (v1 + (v1 * 1000)))
        if len(v4) > 0 {
            v4[0] = 65535
        }
        if true {
            var v7 uint = 1
            v8 := GenS1{F0: 2147483648, F1: ((v7 & (v7 % (((v7 << ((v7) & 63))) | 1))) != v7), F2: v3}
            switch -99 {
            case -810:
                if ((false || ((v1 ^ v1) >= v1)) || (-940 != 0)) {
                    var v9 uint32 = 0
                    v7 += v7
                    _ = v9
                }
            default:
                for fi10 := 0; fi10 < 6; fi10++ {
                    fmt.Println(v3)
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 bool = v3
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v14 := GenS1{F0: 2147483648, F1: v3, F2: (0 <= 42)}
                        _ = fi13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            _ = v7
            _ = v8
        }
    } else {
        v15 := GenS2{F0: -2147483648}
        v4 = append(v4, 852)
        v16 := GenS1{F0: 1, F1: v3, F2: (v3 || true)}
        _ = v15
        _ = v16
    }
    var v17 int32 = 1
    if (42 == 153) {
        if v3 {
            if false {
                var v18 int = 2147483647
                if false {
                    fmt.Println(v3, v3, v2)
                } else {
                    v4 = append(v4, 2)
                    var v19 int32 = 1
                    v20 := GenS1{F0: 1, F1: (!(0 <= 758)), F2: (false && ((!(42 == 255)) && false))}
                    _ = v19
                    _ = v20
                }
                _ = v18
            } else {
                v3 = v3
                if (!((1.5 / float32(1e10)) >= 1e10)) {
                    if (!((1 >= 2) || (834 < -1))) {
                        v3 = v3
                    }
                }
            }
            if (42 < 496) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 float32 = (((3.14159 / 1.0) - -2.25) / 1.0)
                        var v24 string = ("foo bar" + "hello")
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 int16 = v1
                    _ = fi21
                    _ = v25
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                for fi26 := 0; fi26 < 18; fi26++ {
                    v4 = append(v4, 980)
                    v3 = ((2 <= 0) || false)
                    if v3 {
                        fmt.Println(v1, v2, v17)
                        var v27 bool = v3
                        _ = v27
                    }
                    _ = fi26
                }
            }
            if (108 >= 0) {
                if (42 <= 32768) {
                    switch 265 {
                    case -407:
                        v17 += 1
                    }
                } else {
                    v28 := GenS3{F0: 4294967295, F1: v3, F2: -1}
                    _ = v28
                }
                v29 := GenS2{F0: 2147483647}
                _ = v29
            }
        }
    }
    v30 := GenS4{F0: (v17 >> ((18446744073709551615) & 63))}
    v2 -= (^(v2))
    v4 = append(v4, 2147483647)
    var v31 bool = (false || (-681 != 0))
    v32 := GenS1{F0: 4294967295, F1: (-109 >= -1), F2: v31}
    v33 := GenS2{F0: -355}
    _ = len(v4)
    var v34 uint64 = (v2 | v2)
    v35 := GenS2{F0: -335}
    v36 := GenS1{F0: 799, F1: (!((-(v17)) < (v17 * (v17 & (v17 / (((v17 | v17)) | 1)))))), F2: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
