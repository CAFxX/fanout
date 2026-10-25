package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint64) uint8 {
    p0 = p0
    p0 -= (p0 - (p0 - (p0 >> ((741) & 63))))
    var v6 uint = 42
    switch "hello" {
    case "foo bar":
        v6 += v6
    case "":
        v6 = v6
    case "a":
        var v7 int16 = -850
        _ = v7
    default:
        v6 = (v6 * 598)
    }
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 int64, p1 uint8, p2 uint) string {
    if (p0 < 1) {
        var v8 int16 = 1000
        _ = v8
    } else {
        v9 := GenS1{F0: -2147483648, F1: p1}
        fmt.Println(p0, p1)
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{-687}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = -126
    var v3 uint8 = 255
    var v4 uint32 = (^(v1))
    v5 := []int{32767, -499, 1, 32768}
    var v12 bool = true
    v4 = (v1 * v4)
    fmt.Println(v3, v4)
    if (!(1e10 >= float32(0.1))) {
        var v13 int16 = -544
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v13
    } else {
        _ = len(v5)
    }
    if ("a" < ("hello" + ("\n\t" + "\n\t"))) {
        var v14 bool = true
        _ = v14
    }
    var v15 uint32 = v4
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 int = (-(fi16))
        v2 = v2
        v18 := GenS2{F0: v15, F1: 1}
        _ = fi16
        _ = v17
        _ = v18
    }
    v2 = (v2 / ((v2) | 1))
    var v19 int64 = 113
    switch 128 {
    case 128:
        var v20 int8 = v2
        _ = v20
    default:
        v21 := GenS2{F0: (v4 - v15), F1: (v19 % (((v19 << ((65535) & 63))) | 1))}
        _ = v21
    }
    var v22 float64 = 3.14159
    for fi23 := 0; fi23 < 5; fi23++ {
        if v12 {
            v3 -= (v3 << ((uint(v4)) & 63))
            for fi24 := 0; fi24 < 17; fi24++ {
                if v12 {
                    v22 -= (-2.25 + float64(1e10))
                    if v12 {
                        var v25 uint64 = 257
                        var v26 uint16 = 0
                        if true {
                            var v27 int8 = v2
                            fmt.Println(fi23, v19, v2)
                            _ = v27
                        } else {
                            var v28 bool = true
                            if (int64(v3) > v19) {
                                var v29 uint8 = 255
                                var v30 int = (fi23 & (fi23 / ((-844) | 1)))
                                _ = v29
                                _ = v30
                            }
                            _ = v28
                        }
                        _ = v25
                        _ = v26
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                _ = fi24
            }
            var v31 float64 = (1.5 * 0.1)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 15; fi32++ {
                switch 127 {
                case 655:
                    v4 -= v4
                default:
                    if ((v4 + (^(v15))) == 65535) {
                        var v33 uint32 = v15
                        var v34 uint64 = 378
                        v3 = genFn0(752)
                        _ = v33
                        _ = v34
                    }
                }
                v5 = append(v5, 0)
                v35 := GenS2{F0: (v4 * v4), F1: v19}
                _ = fi32
                _ = v35
            }
            var v36 string = ""
            v37 := GenS3{F0: false, F1: v19}
            _ = v36
            _ = v37
        }
        v38 := GenS1{F0: -1, F1: (v3 % (((v3 % ((genFn0(0)) | 1))) | 1))}
        if v12 {
            v4 -= 2147483648
        }
        _ = fi23
        _ = v38
    }
    v39 := GenS4{F0: (("a" + "\n\t") + "")}
    v1 += v1
    v40 := GenS2{F0: v4, F1: (v19 - v19)}
    if v12 {
        v3 = 210
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
