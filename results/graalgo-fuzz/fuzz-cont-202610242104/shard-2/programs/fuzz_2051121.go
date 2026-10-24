package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 int8, p1 int8) uint8 {
    p1 -= p0
    v7 := []int{1}
    if false {
        if false {
            switch -1 {
            case 479:
                _ = len(v7)
            case 128:
                if false {
                    var v8 float32 = 0.0
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        switch fi9 {
                        case 4294967295:
                            p1 = (^(p0))
                        case 9223372036854775807:
                            if (577 >= 42) {
                                p1 = (p1 & (p1 + p0))
                                p1 += 42
                                p0 -= (p0 >> ((4294967295) & 63))
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = fi9
                                }
                                var v10 float64 = (0.0 - float64(1e10))
                                _ = v10
                            }
                        default:
                            var v11 string = ("\n\t" + "")
                            _ = v11
                        }
                        _ = fi9
                    }
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := GenS1{F0: 428}
                        v14 := GenS2{F0: 32767, F1: 407}
                        var v15 uint16 = 0
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
            case -1:
                if true {
                    switch -578 {
                    case 32767:
                        for fi16 := 0; fi16 < 13; fi16++ {
                            for fi17 := 0; fi17 < 9; fi17++ {
                                p0 += (p1 + p0)
                                if (0 == 878) {
                                    fmt.Println(p0, fi17)
                                    v18 := GenS1{F0: -894}
                                    p1 += (-(p1))
                                    _ = v18
                                } else {
                                    var v19 uint = 625
                                    var v20 uint16 = 482
                                    switch (("a" + "a") + "foo bar") {
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            var v21 int8 = p1
                                            _ = v21
                                        case "":
                                            if false {
                                                p0 += -127
                                            }
                                        default:
                                            if false {
                                                v19 = 809
                                                fmt.Println(v20)
                                                v20 = v20
                                            }
                                        }
                                    case "foo bar":
                                        v20 = 412
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                _ = fi17
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi16
                        }
                    case -130:
                        var v22 int16 = -1
                        _ = v22
                    }
                    if (!true) {
                        p1 = 100
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    p0 -= p1
                } else {
                    var v23 string = "hello"
                    var v24 int8 = -128
                    _ = v23
                    _ = v24
                }
            default:
                p1 += (p1 - (p1 % ((24) | 1)))
            }
            fmt.Println(p1, p0)
        }
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            if true {
                p1 = -93
                fmt.Println(p1)
            } else {
                var v26 int64 = 184
                v7 = append(v7, fi25)
                if len(v7) > 0 {
                    v7[0] = -9223372036854775808
                }
                _ = v26
            }
            _ = fi25
        }
    }
    p1 = (p0 & p1)
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := []int{0, 283, -59}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 int = -9223372036854775808
    var v5 uint32 = 2147483648
    v6 := []uint32{(^(v5)), 923, (v5 - (v5 + v5))}
    fmt.Println(v4, v5)
    var v27 float64 = float64(((1.5 + 1.5) * -2.25))
    var v28 int32 = 42
    for fi29 := 0; fi29 < 8; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            delete(v2, (fi29 + 42))
            _ = fi30
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := GenS3{F0: "hello", F1: 18446744073709551615}
        v28 = -1
        if (0 > 2) {
            var v33 int16 = 157
            fmt.Println(fi31, v4, v33)
            _ = v33
        } else {
            v5 = 51
            var v34 string = "a"
            for fi35 := 0; fi35 < 15; fi35++ {
                v27 += (v27 / (v27 * float64(0.1)))
                var v36 uint = 823
                v37 := GenS4{F0: v5, F1: (v28 ^ 50), F2: v34}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    var v38 int64 = 0
    fmt.Println(v28, v28, v38)
    v39 := v2[v4]
    if false {
        var v40 float64 = (v27 / v27)
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
