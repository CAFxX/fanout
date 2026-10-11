package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int64; F2 int32 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int32; F1 uint }

func main() {
    v1 := []string{"foo bar", "", "a", ""}
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 int32 = 1
    var v5 uint8 = 0
    switch 128 {
    case -2147483648:
        v3[304] = -905
    case -837:
        if false {
            v6 := GenS1{F0: 587, F1: 1, F2: 32767}
            var v7 float32 = 0.1
            _ = v6
            _ = v7
        } else {
            var v8 float32 = 1.5
            v2 = true
            v9 := v3[4294967295]
            _ = v8
            _ = v9
        }
    case -10:
        var v10 uint64 = 115
        _ = v10
    default:
        delete(v3, 4294967295)
    }
    v4 = -175
    if false {
        var v11 float32 = (3.14159 / 1e10)
        if (2 < uint16(v5)) {
            switch "hello" {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                fmt.Println(v2)
            case "\n\t":
                v4 = 2147483647
            default:
                v12 := GenS1{F0: 18446744073709551615, F1: -32768, F2: -1}
                _ = v12
            }
        }
        _ = v11
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            v5 = (v5 | v5)
            _ = len(v3)
            _ = fi13
        }
    }
    v14 := GenS1{F0: 1, F1: -32768, F2: -1}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int32 = v4
        _ = len(v3)
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 bool = (581 >= 42)
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 int8 = 25
    v19 = v19
    switch "" {
    case "hello":
        var v20 uint16 = 65535
        _ = v20
    }
    delete(v3, 2147483647)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v19, v19)
    case "":
        v4 += 2147483647
    default:
        for fi21 := 0; fi21 < 3; fi21++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            var v22 int32 = (-(v4))
            for fi23 := 0; fi23 < 11; fi23++ {
                v24 := GenS2{F0: (v19 & v19), F1: "a", F2: uint32(v5)}
                var v25 int = -9223372036854775808
                if v2 {
                    v26 := GenS2{F0: -128, F1: "\n\t", F2: 0}
                    v1 = append(v1, "")
                    var v27 uint8 = (v5 * 255)
                    _ = v26
                    _ = v27
                } else {
                    _ = len(v3)
                    v2 = ((v2 && (2 <= 491)) || (true && (1 != (-(v19)))))
                }
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                if (fi21 < fi21) {
                    fmt.Println(fi21, v5)
                    var v28 bool = true
                    _ = v28
                } else {
                    if true {
                        v5 = (v5 % ((v5) | 1))
                    } else {
                        var v29 uint = 0
                        v22 -= v22
                        _ = len(v3)
                        fmt.Println(v5, v2)
                        var v30 uint64 = 0
                        _ = v29
                        _ = v30
                    }
                    v2 = (837 >= 65535)
                    v22 += v4
                    fmt.Println(fi21, v4)
                    v31 := GenS1{F0: 255, F1: 32767, F2: 0}
                    v32 := GenS2{F0: v19, F1: "foo bar", F2: 2}
                    switch "" {
                    case "a":
                        var v33 int = 127
                        _ = v33
                    case "hello":
                        fmt.Println(fi23, fi23, v4)
                    case "\n\t":
                        switch "foo bar" {
                        case "foo bar":
                            var v34 int64 = 1
                            _ = v34
                        case "a":
                            switch v25 {
                            case -295:
                                fmt.Println(fi23, v5)
                            case 2147483647:
                                if v2 {
                                    var v35 uint16 = 156
                                    v36 := GenS1{F0: uint(v5), F1: -76, F2: 0}
                                    v22 += 1
                                    _ = v35
                                    _ = v36
                                }
                            case -2147483648:
                                if ((v19 / ((v19) | 1)) != v19) {
                                    v37 := GenS1{F0: 42, F1: -904, F2: 32767}
                                    var v38 uint8 = 70
                                    v39 := GenS1{F0: 1, F1: 1000, F2: -32768}
                                    var v40 int32 = v4
                                    v41 := v3[2147483648]
                                    v2 = false
                                    if false {
                                        v19 += (v19 - v19)
                                        fmt.Println(v41)
                                        var v42 string = "hello"
                                        _ = v42
                                    } else {
                                        var v43 int64 = 711
                                        _ = v43
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    if v2 {
                                        v4 -= 2147483647
                                        var v44 uint16 = 1
                                        _ = v44
                                    }
                                    if (v2 && (v2 && (813 <= 2))) {
                                        v45 := GenS3{F0: 65535, F1: true}
                                        var v46 int16 = -912
                                        v47 := GenS4{F0: "foo bar"}
                                        var v48 int8 = -128
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            default:
                                var v49 int16 = 1
                                _ = v49
                            }
                        case "":
                            v4 += (v4 % ((v22) | 1))
                        default:
                            switch ("foo bar" + "a") {
                            case "hello":
                                v50 := GenS4{F0: ("hello" + "a")}
                                if (v2 || (!(1e10 != 0.1))) {
                                    if false {
                                        if (!(!v2)) {
                                            fmt.Println(v5, v4)
                                            v5 -= (v5 + 42)
                                            v51 := GenS5{F0: true, F1: -1, F2: -961}
                                            if (v2 || (!true)) {
                                                var v52 bool = (4294967295 == 65535)
                                                _ = v52
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = "foo bar"
                                                }
                                                fmt.Println(fi21, v19, fi21)
                                                if (1 < 4294967295) {
                                                    v5 = (v5 / (((v5 >> ((18446744073709551615) & 63))) | 1))
                                                    v2 = v2
                                                } else {
                                                    v4 += (v22 & (v22 * (v22 * 1)))
                                                    if v2 {
                                                        var v53 bool = false
                                                        var v54 uint64 = 4294967295
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    v19 -= (v19 | v19)
                                                }
                                            }
                                            v55 := v3[751]
                                            _ = v51
                                            _ = v55
                                        }
                                        v4 = v4
                                        fmt.Println(v19, v25)
                                        var v56 uint64 = 788
                                        _ = v56
                                    }
                                } else {
                                    v22 = (v4 % ((v22) | 1))
                                }
                                if v2 {
                                    switch "\n\t" {
                                    case "hello":
                                        v57 := GenS6{F0: 2147483647, F1: 42}
                                        _ = v57
                                    case "":
                                        var v58 int64 = 42
                                        _ = v58
                                    }
                                    var v59 float32 = ((((-2.25 / (1.5 / 1.0)) * float32(1.5)) * -2.25) * float32((0.1 - ((float64((1e10 + ((1e10 * 0.0) + (1.5 - (0.0 / -2.25))))) - 0.0) / (((0.0 + 1.5) / 1.0) - (float64((0.1 * -2.25)) + (0.1 * (1e10 - 0.1))))))))
                                    _ = v59
                                }
                                _ = v50
                            }
                        }
                    default:
                        switch (fi23 / ((v25) | 1)) {
                        case 127:
                            fmt.Println(v22, v5)
                        case 246:
                            v60 := v3[743]
                            _ = v60
                        default:
                            var v61 uint8 = (v5 / (((v5 % ((v5) | 1))) | 1))
                            v2 = (float32(0.0) <= 1e10)
                            if (((-1 < 1) || (!(false || (!((v25 ^ (fi21 ^ -9223372036854775808)) >= (v25 * (v25 % (((-(fi23))) | 1)))))))) || v2) {
                                var v62 int16 = 1000
                                if (!false) {
                                    v63 := GenS3{F0: 2, F1: false}
                                    v4 = 1
                                    v19 += -128
                                    _ = v63
                                } else {
                                    var v64 uint16 = 2
                                    fmt.Println(v62, v19, fi21)
                                    v22 += v4
                                    _ = v64
                                }
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                _ = v62
                            } else {
                                switch fi21 {
                                case 32767:
                                    _ = len(v1)
                                default:
                                    v65 := GenS3{F0: 2147483648, F1: (0 > (fi21 << ((42) & 63)))}
                                    _ = v65
                                }
                                var v66 float32 = 0.1
                                v67 := GenS5{F0: ((v5 ^ (v61 % ((v5) | 1))) == v5), F1: 9223372036854775807, F2: v22}
                                _ = v66
                                _ = v67
                            }
                            _ = v61
                        }
                        var v68 uint32 = 2147483648
                        switch -887 {
                        case -9223372036854775808:
                            if (0.0 != (-2.25 * -2.25)) {
                                v69 := GenS4{F0: "foo bar"}
                                v25 -= -37
                                fmt.Println(v25)
                                _ = v69
                            }
                        case 4294967295:
                            switch "hello" {
                            case "a":
                                v19 -= 120
                            case "foo bar":
                                var v70 uint32 = 715
                                _ = v70
                            default:
                                if (("foo bar" + "hello") == "hello") {
                                    fmt.Println(v4, v5, v68)
                                    var v71 uint16 = 2
                                    _ = v71
                                } else {
                                    v19 -= (-(v19))
                                    v25 += (^(v25))
                                    fmt.Println(fi21, fi21)
                                }
                            }
                            v72 := GenS7{F0: -1, F1: 2}
                            var v73 uint = uint(v5)
                            _ = v72
                            _ = v73
                        case 1:
                            v74 := GenS3{F0: v68, F1: true}
                            _ = v74
                        }
                        _ = v68
                    }
                    _ = v31
                    _ = v32
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    }
    for fi75 := 0; fi75 < 6; fi75++ {
        v5 += v5
        _ = fi75
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
