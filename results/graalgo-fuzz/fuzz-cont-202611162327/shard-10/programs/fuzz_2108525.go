package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(fi6, p0)
        _ = fi6
    }
    p0 = p0
    v7 := GenS1{F0: -1, F1: 0, F2: 42}
    _ = p0
    _ = v7
    return 469
}

func main() {
    var v1 int8 = -1
    var v2 float64 = 1e10
    var v3 string = "a"
    v4 := []string{"\n\t", "a"}
    v5 := make(map[string]uint32)
    v1 = v1
    if (2147483647 < -1) {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi8)
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 bool = (false || (false || (false && true)))
                var v11 uint32 = 1
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        fmt.Println(v1, v1)
        if ((!false) && (1 < 351)) {
            fmt.Println(v2, v1, v3)
        } else {
            if true {
                v2 += (v2 + (float64(((float32((float64((float32(v2) * 1.5)) + v2)) / 1e10) / 1.5)) / ((v2 + v2) + (0.1 / v2))))
            } else {
                v12 := GenS1{F0: 42, F1: 255, F2: -1}
                for fi13 := 0; fi13 < 18; fi13++ {
                    v3 = "a"
                    v2 = 0.1
                    var v14 int = (fi13 * (fi13 / (((fi13 << ((450) & 63))) | 1)))
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
        }
    } else {
        v1 -= 127
    }
    var v15 float32 = 0.0
    var v16 int16 = 1
    v17 := GenS1{F0: (v16 << ((genFn0(v1)) & 63)), F1: 128, F2: 9223372036854775807}
    var v18 uint8 = 58
    var v19 uint8 = (v18 * v18)
    var v20 int16 = (-(v16))
    if (1 < -1) {
        var v21 int64 = -9223372036854775808
        for fi22 := 0; fi22 < 20; fi22++ {
            _ = len(v4)
            if true {
                if false {
                    var v23 int64 = v21
                    v19 += (^(v18))
                    fmt.Println(v18, v3)
                    _ = v23
                } else {
                    v18 = (v18 | v19)
                    for fi24 := 0; fi24 < 17; fi24++ {
                        v16 = (^(v16))
                        _ = fi24
                    }
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS2{F0: (((false || false) && true) || (18446744073709551615 < 2)), F1: true}
                        var v27 int16 = v20
                        var v28 int16 = -32768
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 18; fi29++ {
                    fmt.Println(v20, v19, v19)
                    var v30 int16 = (v20 >> ((876) & 63))
                    v18 = v18
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi22
        }
        _ = v21
    }
    if false {
        v19 += 1
    }
    if len(v4) > 0 {
        v4[0] = (((((v3 + ("\n\t" + v3)) + v3) + ("foo bar" + v3)) + "") + v3)
    }
    if true {
        v31 := GenS3{F0: v16, F1: 785, F2: v3}
        v19 = 255
        _ = v31
    } else {
        if (1 >= 1) {
            v3 = ("" + "\n\t")
        }
        v32 := GenS4{F0: v18, F1: false, F2: 42}
        _ = v32
    }
    v33 := GenS5{F0: 65535, F1: genFn0((v1 & 1))}
    v19 -= 128
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := GenS6{F0: 0, F1: 249, F2: 681}
            if false {
                v2 -= float64((float32(1e10) / v15))
                var v37 uint = 643
                var v38 uint32 = 2147483648
                if (!(674 < 255)) {
                    var v39 float32 = v15
                    _ = v39
                } else {
                    _ = len(v4)
                }
                _ = v37
                _ = v38
            } else {
                fmt.Println(v3)
                var v40 uint32 = 2147483648
                _ = v40
            }
            var v41 uint = 417
            _ = fi35
            _ = v36
            _ = v41
        }
        _ = fi34
    }
    v3 = v3
    for fi42 := 0; fi42 < 8; fi42++ {
        switch fi42 {
        case 0:
            var v43 int = (fi42 & 2147483647)
            _ = v43
        case 456:
            fmt.Println(v20, v15, v20)
        }
        _ = fi42
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
