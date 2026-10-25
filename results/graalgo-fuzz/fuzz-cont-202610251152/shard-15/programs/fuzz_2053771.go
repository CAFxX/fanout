package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int, p2 uint8) bool {
    p0 = p0
    v7 := []int{(p1 + (p1 >> ((396) & 63)))}
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 = append(v7, (^(p1)))
        var v9 int16 = 474
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 uint64, p1 int32) uint64 {
    if (753 > 42) {
        switch 44 {
        case 32767:
            v10 := GenS1{F0: p0, F1: -472}
            _ = v10
        case 4294967295:
            if ((((-2.25 * float32(1.5)) / 1.0) + 3.14159) > (1.5 / 1.5)) {
                if false {
                    var v11 int64 = 9223372036854775807
                    v12 := GenS2{F0: 0, F1: 4294967295, F2: -32768}
                    for fi13 := 0; fi13 < 3; fi13++ {
                        fmt.Println(fi13, p1)
                        if ((true || false) || false) {
                            p1 -= p1
                            for fi14 := 0; fi14 < 20; fi14++ {
                                p0 -= (p0 + (^(p0)))
                                fmt.Println(v11, p0)
                                _ = fi14
                            }
                            switch fi13 {
                            case 32767:
                                v15 := GenS3{F0: 141, F1: 1, F2: ""}
                                _ = v15
                            case 42:
                                for fi16 := 0; fi16 < 20; fi16++ {
                                    v17 := []int{}
                                    _ = fi16
                                    _ = v17
                                }
                            default:
                                fmt.Println(p0)
                            }
                        } else {
                            p1 = p1
                            var v18 float64 = ((3.14159 - (1e10 + 0.1)) + (((1.5 / float64(-2.25)) + 1e10) + 3.14159))
                            _ = v18
                        }
                        var v19 int = (fi13 | (fi13 - (fi13 ^ fi13)))
                        _ = fi13
                        _ = v19
                    }
                    _ = v11
                    _ = v12
                }
                v20 := []int{-463, 127, 0, 2147483647}
                fmt.Println(p1)
                _ = v20
            }
        default:
            if true {
                if false {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v22 := GenS4{F0: 100, F1: p0}
                        _ = fi21
                        _ = v22
                    }
                    var v23 int8 = 42
                    _ = v23
                }
            }
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            p1 += (p1 % (((p1 | int32(p0))) | 1))
            _ = fi24
        }
    }
    if true {
        v25 := GenS2{F0: 32768, F1: p0, F2: 0}
        var v26 bool = true
        if (!v26) {
            v27 := []int{-363}
            var v28 string = "a"
            var v29 int16 = 603
            _ = v27
            _ = v28
            _ = v29
        } else {
            v30 := []int{1, 9223372036854775807, -258, 127}
            switch 127 {
            case -345:
                fmt.Println(v26, p0)
            case 32767:
                v26 = v26
            }
            _ = v30
        }
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    return (p0 << ((255) & 63))
}

func main() {
    var v1 float64 = (3.14159 / 1.0)
    v2 := make(map[string]uint32)
    var v3 bool = (!false)
    v4 := make(map[uint32]uint32)
    var v5 string = "a"
    var v6 int64 = 42
    delete(v2, "")
    var v31 string = v5
    v32 := GenS5{F0: -1}
    var v33 uint32 = 65535
    fmt.Println(v1, v1)
    fmt.Println(v3, v3)
    v34 := GenS5{F0: 0}
    for fi35 := 0; fi35 < 14; fi35++ {
        v1 -= float64(3.14159)
        var v36 bool = (((!(!false)) && (!v3)) || (false && false))
        _ = fi35
        _ = v36
    }
    v31 = "foo bar"
    if ((v1 < 0.0) || (1 > 255)) {
        var v37 int64 = 563
        for fi38 := 0; fi38 < 1; fi38++ {
            var v39 int8 = 100
            var v40 uint8 = 0
            v5 = ("a" + (v5 + "\n\t"))
            _ = fi38
            _ = v39
            _ = v40
        }
        v5 = "foo bar"
        _ = v37
    } else {
        if (0 < -9223372036854775808) {
            fmt.Println(v6)
            if true {
                for fi41 := 0; fi41 < 19; fi41++ {
                    v42 := v4[2147483647]
                    var v43 int = fi41
                    v5 = v5
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
        fmt.Println(v6, v6)
    }
    v5 = (v31 + "foo bar")
    v44 := []int{32767}
    switch 42 {
    case 42:
        v4[(v33 >> ((2) & 63))] = (v33 << ((0) & 63))
    case 2147483647:
        v45 := GenS4{F0: 127, F1: 0}
        _ = v45
    }
    var v46 float64 = -2.25
    var v47 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
