package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
func genFn0(p0 int, p1 int8) uint {
    if true {
        p1 = p1
        fmt.Println(p0, p1)
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        p1 = p1
        v6 := GenS1{F0: 0, F1: 4294967295}
        v7 := GenS2{F0: true, F1: 19, F2: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint16, p1 int16, p2 uint8) string {
    var v8 uint64 = 1
    if (0.1 != (3.14159 * float64((1e10 / (-2.25 * 1e10))))) {
        var v9 int64 = int64(p0)
        v8 = uint64(p0)
        _ = v9
    }
    if ((^(v8)) <= 866) {
        var v10 uint16 = 0
        fmt.Println(p2)
        p0 -= 154
        _ = v10
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{(fi11 | (fi11 / (((fi11 >> ((genFn0((fi11 % ((fi11) | 1)), 127)) & 63))) | 1))), (fi11 | fi11), (^(fi11))}
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v8 += 223
        _ = fi13
    }
    v14 := GenS1{F0: 945, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v8
    return ("a" + "foo bar")
}

func main() {
    var v1 uint = 290
    v2 := []int{-1}
    var v3 uint32 = 42
    var v4 string = ""
    var v15 uint64 = 534
    fmt.Println(v4)
    v16 := GenS2{F0: (-9 <= 0), F1: (v3 | v3), F2: 255}
    fmt.Println(v1, v3, v4)
    var v17 int8 = 36
    v4 = v4
    v18 := GenS2{F0: true, F1: uint32(v15), F2: 2}
    if len(v2) > 0 {
        v2[0] = 827
    }
    v3 += (v3 & 2)
    if false {
        switch 0 {
        case 2147483647:
            switch v4 {
            case "foo bar":
                switch "a" {
                case "foo bar":
                    var v19 int32 = 42
                    _ = v19
                case "\n\t":
                    if (0.1 == (((-2.25 / (3.14159 - 1.5)) / 1.0) / float32(3.14159))) {
                        if (!(v17 == -104)) {
                            var v20 uint8 = 128
                            var v21 int = 128
                            _ = v20
                            _ = v21
                        } else {
                            v2 = append(v2, -9223372036854775808)
                        }
                        switch 497 {
                        case -9223372036854775808:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 32767:
                            if false {
                                switch ((v4 + "\n\t") + v4) {
                                case "\n\t":
                                    switch (v4 + v4) {
                                    case "a":
                                        if (!true) {
                                            var v22 int32 = 1
                                            _ = v22
                                        } else {
                                            var v23 uint64 = (v15 % (((v15 & (v15 - 42))) | 1))
                                            var v24 uint32 = 2
                                            v25 := GenS1{F0: 728, F1: v1}
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                    case "":
                                        for fi26 := 0; fi26 < 2; fi26++ {
                                            var v27 uint64 = (v15 / ((v15) | 1))
                                            _ = fi26
                                            _ = v27
                                        }
                                    case "hello":
                                        v28 := GenS2{F0: (!false), F1: (v3 * 757), F2: 0}
                                        _ = v28
                                    default:
                                        v4 = ""
                                    }
                                case "hello":
                                    fmt.Println(v15)
                                default:
                                    for fi29 := 0; fi29 < 17; fi29++ {
                                        var v30 uint8 = 255
                                        _ = fi29
                                        _ = v30
                                    }
                                }
                                v3 = (v3 / (((^(v3))) | 1))
                            } else {
                                fmt.Println(v15)
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    v1 -= (v1 >> (((v1 << ((v1) & 63))) & 63))
                                    v4 = v4
                                    fmt.Println(v15, v1, v4)
                                    _ = fi31
                                }
                                v32 := GenS3{F0: v1, F1: v4, F2: 462}
                                _ = v32
                            }
                        default:
                            var v33 int32 = 42
                            _ = v33
                        }
                        fmt.Println(v3)
                    } else {
                        if (-1 > 0) {
                            var v34 int16 = 42
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v4 = "\n\t"
                            v2 = append(v2, (fi35 >> (((v1 + (^(v1)))) & 63)))
                            fmt.Println(v15, v1)
                            _ = fi35
                        }
                    }
                }
            case "\n\t":
                v15 -= v15
            default:
                var v36 uint16 = 65535
                _ = v36
            }
        default:
            v2 = append(v2, -673)
        }
        var v37 float64 = (-2.25 * (1e10 + 0.1))
        _ = v37
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            v15 = 203
            var v39 float32 = float32(float64((0.1 / 1e10)))
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 int8 = 127
                v42 := GenS4{F0: 475, F1: -704, F2: fi40}
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi38
            _ = v39
        }
    }
    v17 -= int8(v1)
    var v43 uint8 = 128
    switch "a" {
    case "a":
        switch (v4 + (v4 + (((((v4 + "hello") + v4) + "foo bar") + (v4 + ("hello" + v4))) + v4))) {
        case "foo bar":
            v17 -= v17
        default:
            fmt.Println(v1, v43, v15)
        }
    case "":
        _ = len(v2)
    }
    var v44 int32 = 1
    if false {
        fmt.Println(v15, v17)
    } else {
        v45 := GenS1{F0: 32767, F1: (^(v1))}
        _ = v45
    }
    fmt.Println(v4)
    var v46 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
}
