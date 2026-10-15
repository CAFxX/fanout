package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 string; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16) int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        if false {
            if (true || (true && false)) {
                switch "" {
                case "\n\t":
                    if ((p0 != (p0 & 255)) || true) {
                        v7 := []int{-646}
                        if ((^(p0)) < (p0 - 42)) {
                            if len(v7) > 0 {
                                v7[0] = -876
                            }
                        }
                        if (!true) {
                            for fi8 := 0; fi8 < 17; fi8++ {
                                v7 = append(v7, 0)
                                switch -1 {
                                case -9223372036854775808:
                                    if (!true) {
                                        v9 := GenS1{F0: -32768}
                                        fmt.Println(fi8, p0)
                                        var v10 uint8 = 19
                                        _ = v10
                                        _ = v9
                                    }
                                }
                                var v11 string = "foo bar"
                                _ = fi8
                                _ = v11
                            }
                        } else {
                            if true {
                                p0 -= (^(p0))
                                for fi12 := 0; fi12 < 19; fi12++ {
                                    if true {
                                        p0 = (p0 ^ 431)
                                        _ = len(v7)
                                    }
                                    _ = fi12
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                        }
                        _ = v7
                    } else {
                        var v13 bool = false
                        _ = v13
                    }
                case "hello":
                    v14 := GenS1{F0: -32768}
                    _ = v14
                }
                v15 := GenS1{F0: 1000}
                _ = v15
            }
        } else {
            p0 -= (p0 | p0)
            p0 = (p0 + p0)
            switch (^(fi6)) {
            case -955:
                for fi16 := 0; fi16 < 11; fi16++ {
                    if (1 <= 1) {
                        p0 -= p0
                    } else {
                        v17 := GenS2{F0: 0, F1: 1}
                        _ = v17
                    }
                    p0 -= p0
                    p0 += 255
                    _ = fi16
                }
            case -9223372036854775808:
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := []int{}
                    _ = fi18
                    _ = v19
                }
            }
        }
        _ = fi6
    }
    fmt.Println(p0)
    var v20 float64 = 0.1
    _ = p0
    _ = v20
    return 110
}
func genFn1(p0 uint32, p1 int32, p2 int64) int {
    v21 := []int{}
    fmt.Println(p0, p1, p1)
    var v22 int32 = p1
    v21 = append(v21, -1)
    switch ("foo bar" + "a") {
    case "hello":
        switch "foo bar" {
        case "a":
            var v23 uint = 313
            _ = v23
        case "hello":
            fmt.Println(v22, p0)
        default:
            for fi24 := 0; fi24 < 19; fi24++ {
                _ = len(v21)
                _ = fi24
            }
        }
    default:
        v25 := GenS1{F0: 0}
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return -457
}

func main() {
    v1 := []int{-2147483648, 255}
    var v2 uint64 = 1
    var v3 int = 42
    v4 := make(map[uint32]uint32)
    var v5 float32 = float32((3.14159 + 3.14159))
    v26 := v4[2147483648]
    if ((!(v26 > (^(v26)))) || true) {
        v27 := GenS3{F0: (("a" + ("" + ("" + "a"))) + ("\n\t" + "")), F1: (v2 % ((669) | 1)), F2: 32767}
        var v28 float64 = float64((float32((-2.25 + 1e10)) - 3.14159))
        _ = v27
        _ = v28
    } else {
        _ = len(v4)
        var v29 float64 = (((-2.25 * 0.0) + 1.5) / (float64((1.5 / (v5 + 0.1))) / 1.0))
        if len(v1) > 0 {
            v1[0] = (v3 * v3)
        }
        _ = v29
    }
    var v30 uint32 = v26
    var v31 float64 = (1.5 / 3.14159)
    for fi32 := 0; fi32 < 4; fi32++ {
        if false {
            fmt.Println(v30)
        } else {
            v33 := v4[(v26 << ((556) & 63))]
            _ = v33
        }
        var v34 uint32 = (v30 % (((^(v30))) | 1))
        for fi35 := 0; fi35 < 13; fi35++ {
            if true {
                v30 = v34
            } else {
                _ = len(v4)
            }
            if false {
                switch genFn1(v30, 1, 42) {
                case -9223372036854775808:
                    var v36 float32 = ((v5 + ((float32((v31 * (v31 - v31))) / 1e10) * v5)) * float32(v31))
                    _ = v36
                default:
                    var v37 float64 = ((float64((v5 * float32(v31))) - v31) - (v31 / v31))
                    _ = v37
                }
                var v38 float32 = (1.5 + v5)
                _ = v38
            } else {
                v39 := GenS4{F0: 2}
                if len(v1) > 0 {
                    v1[0] = v3
                }
                v40 := GenS5{F0: "", F1: 32767}
                _ = v39
                _ = v40
            }
            v3 += (fi35 - genFn1((^(v30)), 1, 1))
            var v41 int16 = 233
            _ = fi35
            _ = v41
        }
        _ = fi32
        _ = v34
    }
    if true {
        delete(v4, v26)
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        if false {
            v43 := GenS6{F0: 43, F1: 1}
            _ = v43
        } else {
            for fi44 := 0; fi44 < 4; fi44++ {
                fmt.Println(v2, v26)
                var v45 uint16 = 1
                var v46 string = "hello"
                v47 := GenS2{F0: 483, F1: v26}
                v2 -= 2
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
        }
        var v48 uint16 = 67
        _ = fi42
        _ = v48
    }
    var v49 uint32 = 4294967295
    v3 = -232
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
}
