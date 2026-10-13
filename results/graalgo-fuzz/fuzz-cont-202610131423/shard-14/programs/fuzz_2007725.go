package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int32; F2 int }
type GenS7 struct { F0 int8; F1 uint32; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint32 }
func genFn0() uint16 {
    if (65535 < 1) {
        var v7 bool = true
        v8 := []int{-2147483648, 127}
        v7 = (127 != 150)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        switch "\n\t" {
        case "hello":
            var v10 int8 = 42
            _ = v10
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
        }
        fmt.Println(fi9)
        _ = fi9
    }
    switch "foo bar" {
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                switch ((("foo bar" + (("\n\t" + ("hello" + "a")) + "hello")) + "") + ("" + "hello")) {
                case "":
                    var v14 bool = (-2.25 < float32((3.14159 / 1e10)))
                    _ = v14
                case "a":
                    if (0.1 <= ((float64(1.5) * 1.5) + float64(1.5))) {
                        v15 := []int{fi13, fi12, fi13, fi13}
                        var v16 int = (fi13 | fi12)
                        var v17 float32 = 0.0
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
                _ = fi13
            }
            var v18 float32 = ((1e10 / 3.14159) * 0.0)
            _ = fi12
            _ = v18
        }
    default:
        var v19 uint64 = 545
        _ = v19
    }
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 uint32 = 65535
    v4 := []string{}
    v5 := make(map[uint32]bool)
    v6 := []string{"a", ("\n\t" + ("foo bar" + ("foo bar" + ("a" + "\n\t"))))}
    fmt.Println(v3)
    v20 := GenS1{F0: ("\n\t" + "\n\t"), F1: 2, F2: false}
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    var v21 int8 = 42
    if (!false) {
        for fi22 := 0; fi22 < 13; fi22++ {
            v21 += v21
            for fi23 := 0; fi23 < 13; fi23++ {
                if (4294967295 <= 4294967295) {
                    if false {
                        var v24 uint64 = 1
                        if (1.5 > 1e10) {
                            if true {
                                v24 = v24
                            }
                            v4 = append(v4, "hello")
                        }
                        switch "a" {
                        case "":
                            var v25 float32 = 1e10
                            _ = v25
                        case "foo bar":
                            if false {
                                var v26 float64 = 3.14159
                                v3 = (^(v3))
                                v27 := GenS1{F0: "hello", F1: v24, F2: true}
                                _ = v26
                                _ = v27
                            } else {
                                switch ((("" + "foo bar") + "hello") + "\n\t") {
                                case "":
                                    fmt.Println(v21, v21)
                                case "foo bar":
                                    v21 = v21
                                    v28 := GenS2{F0: -684, F1: -124}
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: "hello", F1: v24, F2: (int16(v24) == 0)}
                                    _ = v29
                                }
                                if false {
                                    var v30 int = 127
                                    v3 -= (v3 & (v3 ^ (^(v3))))
                                    v31 := GenS3{F0: 0, F1: 1, F2: 543}
                                    _ = v30
                                    _ = v31
                                }
                            }
                        case "\n\t":
                            var v32 int8 = v21
                            _ = v32
                        }
                        _ = v24
                    }
                    v3 = (v3 << ((255) & 63))
                    if (-1 >= 42) {
                        var v33 uint64 = 42
                        switch -9223372036854775808 {
                        case 2147483647:
                            if len(v6) > 0 {
                                v6[0] = "hello"
                            }
                        case -173:
                            _ = len(v5)
                        case 20:
                            v34 := GenS3{F0: v21, F1: -2147483648, F2: uint(v33)}
                            _ = v34
                        }
                        v35 := GenS1{F0: "\n\t", F1: (v33 / ((v33) | 1)), F2: false}
                        _ = v33
                        _ = v35
                    } else {
                        v36 := GenS4{F0: 494, F1: true, F2: 1}
                        var v37 uint64 = 416
                        _ = v36
                        _ = v37
                    }
                } else {
                    fmt.Println(fi23)
                    delete(v2, ("a" + "hello"))
                    delete(v2, "")
                }
                fmt.Println(fi22, v21, v3)
                v21 = 127
                _ = fi23
            }
            _ = fi22
        }
        var v38 uint16 = 447
        if false {
            v21 += v21
        }
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            if true {
                v4 = append(v4, "hello")
                v40 := GenS2{F0: -1, F1: v21}
                _ = v40
            }
            v41 := GenS5{F0: 255}
            if (((-(fi39)) >= (fi39 | fi39)) || false) {
                _ = len(v4)
            }
            _ = fi39
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        v43 := GenS6{F0: 42, F1: 42, F2: (fi42 >> ((456) & 63))}
        if true {
            v21 = int8(fi42)
        } else {
            v3 -= v3
        }
        v21 = (^(v21))
        _ = fi42
        _ = v43
    }
    if ((false && true) || (1 > 1)) {
        v21 -= -1
    } else {
        for fi44 := 0; fi44 < 5; fi44++ {
            if ((32767 < 127) && false) {
                var v45 float32 = (((1.5 - -2.25) - (float32(1.5) + 0.1)) / 1e10)
                if (!true) {
                    v21 = (v21 << ((65535) & 63))
                }
                _ = v45
            }
            if (0 == 1) {
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 int32 = 2147483647
                    switch fi44 {
                    case 32767:
                        fmt.Println(fi46, v47, v47)
                    case 9223372036854775807:
                        if false {
                            v21 += (v21 / ((1) | 1))
                            var v48 int = (fi46 / ((9223372036854775807) | 1))
                            switch (v48 - (fi46 - 32768)) {
                            case -1:
                                if ("\n\t" != "a") {
                                    fmt.Println(fi44, fi46)
                                }
                            case -189:
                                v49 := GenS4{F0: 65535, F1: (!(-128 == int8(v3))), F2: -1}
                                if true {
                                    v21 = v21
                                    v48 += fi46
                                    v5[677] = true
                                } else {
                                    v21 = 100
                                }
                                if (0 < 255) {
                                    var v50 float64 = -2.25
                                    var v51 int64 = -1
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v52 uint8 = 241
                                        _ = v52
                                    case "":
                                        v5[(v3 * v3)] = (42 < (v47 - 413))
                                    case "a":
                                        _ = len(v2)
                                    default:
                                        v48 = (fi46 >> ((576) & 63))
                                        v53 := GenS4{F0: 65535, F1: (false || (!true)), F2: int16(v51)}
                                        var v54 uint64 = 390
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = v50
                                    _ = v51
                                }
                                _ = v49
                            case 4294967295:
                                var v55 uint64 = 18446744073709551615
                                _ = v55
                            default:
                                var v56 uint16 = 1
                                _ = v56
                            }
                            _ = v48
                        } else {
                            var v57 int64 = 1
                            var v58 int64 = v57
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v57
                            _ = v58
                        }
                    default:
                        delete(v2, "\n\t")
                    }
                    _ = fi46
                    _ = v47
                }
                switch ("a" + "foo bar") {
                case "hello":
                    var v59 uint8 = 7
                    _ = v59
                default:
                    var v60 int16 = 1000
                    _ = v60
                }
            } else {
                var v61 string = "foo bar"
                var v62 uint16 = 65535
                for fi63 := 0; fi63 < 9; fi63++ {
                    var v64 uint8 = 179
                    var v65 uint32 = (v3 & (v3 | (v3 % ((42) | 1))))
                    switch (fi44 % (((^(fi63))) | 1)) {
                    case 128:
                        v5[v3] = (!true)
                    case 127:
                        v66 := GenS7{F0: 1, F1: 2147483648, F2: (fi44 + (^(fi63)))}
                        v3 = v65
                        var v67 uint16 = (v62 & (v62 ^ v62))
                        _ = v66
                        _ = v67
                    default:
                        v21 = (v21 ^ 0)
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = v61
                _ = v62
            }
            _ = fi44
        }
        v3 = v3
        v68 := GenS8{F0: 9223372036854775807}
        _ = v68
    }
    switch "hello" {
    case "foo bar":
        v3 = (v3 ^ (v3 ^ 756))
    case "a":
        v69 := GenS9{F0: 2, F1: v3}
        _ = v69
    default:
        switch 182 {
        case -2147483648:
            var v70 uint16 = 276
            _ = v70
        case -9223372036854775808:
            fmt.Println(v3, v21)
        default:
            var v71 int32 = 0
            _ = v71
        }
    }
    if ((1 > 0) && (63 == -590)) {
        v72 := GenS2{F0: 0, F1: -1}
        v3 += v3
        fmt.Println(v3, v21)
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
