package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0() uint8 {
    var v6 float64 = float64(0.1)
    v7 := GenS1{F0: ((!false) && false), F1: 562}
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int, p1 uint16) uint {
    if (!(42 > 1)) {
        var v8 int = 848
        if false {
            fmt.Println(p1)
        }
        v8 += (p0 | -234)
        _ = v8
    }
    fmt.Println(p1, p0)
    fmt.Println(p0)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 bool = true
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return 631
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (v1 ^ (v1 / (((^(v1))) | 1)))
    v3 := make(map[int]bool)
    var v4 uint32 = (v1 % ((v2) | 1))
    var v5 bool = (0 >= 150)
    switch "a" {
    case "foo bar":
        v1 = 2
    }
    v5 = (!(0.0 >= 1.5))
    v5 = (2147483647 <= 261)
    if (!(false || v5)) {
        fmt.Println(v5, v2, v5)
        switch 9223372036854775807 {
        case -9223372036854775808:
            v1 = (v2 | (v1 | v4))
        case 0:
            fmt.Println(v1, v1)
        case -967:
            var v11 string = ("foo bar" + "foo bar")
            _ = v11
        }
        v12 := v3[-2147483648]
        _ = v12
    }
    switch 481 {
    case 127:
        v13 := GenS1{F0: v5, F1: -2147483648}
        _ = v13
    }
    if v5 {
        if (!v5) {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 string = "foo bar"
                var v16 uint8 = genFn0()
                if (1000 >= 42) {
                    if v5 {
                        var v17 float64 = 1e10
                        v15 = v15
                        _ = v17
                    } else {
                        var v18 int32 = 129
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 uint32 = 759
                            var v21 int = fi14
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        _ = v18
                    }
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 uint64 = 0
                        fmt.Println(v2, v15)
                        v24 := GenS1{F0: (!v5), F1: 1}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    var v25 int = (fi14 % ((0) | 1))
                    _ = v25
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                v1 = v2
                _ = fi26
            }
        }
    }
    v27 := []int{301}
    var v28 float32 = (-2.25 + 0.0)
    switch 480 {
    case 883:
        v2 -= v4
    case 128:
        v2 = v4
    }
    var v29 int = 9223372036854775807
    for fi30 := 0; fi30 < 8; fi30++ {
        _ = len(v3)
        v31 := GenS2{F0: -863, F1: 42}
        _ = fi30
        _ = v31
    }
    var v32 string = "\n\t"
    var v33 uint8 = 127
    var v34 uint64 = 18446744073709551615
    for fi35 := 0; fi35 < 18; fi35++ {
        var v36 int = -81
        if v5 {
            if v5 {
                v37 := GenS2{F0: -32768, F1: 255}
                v4 += v1
                _ = v37
            }
            var v38 int32 = 538
            _ = v38
        } else {
            var v39 uint = 517
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        v41 := GenS3{F0: 32767, F1: 0, F2: -2147483648}
        if (float64(0.0) != (0.1 / -2.25)) {
            v2 = v4
            if v5 {
                if v5 {
                    fmt.Println(v28, v2)
                }
                if v5 {
                    v32 = (("a" + v32) + v32)
                    v33 = genFn0()
                }
            } else {
                v42 := v3[v29]
                var v43 uint8 = 0
                _ = v42
                _ = v43
            }
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v2, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
