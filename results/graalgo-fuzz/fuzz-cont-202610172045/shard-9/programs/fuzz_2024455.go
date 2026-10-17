package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 int64; F1 uint16 }
func genFn0() int {
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 = v7
        _ = fi8
    }
    v7 -= 0
    _ = v7
    return 568
}
func genFn1() uint8 {
    v9 := GenS1{F0: -140, F1: 42, F2: 775}
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 uint8 = 128
    _ = v11
    _ = v9
    return 128
}

func main() {
    var v1 int = -839
    v2 := []int{(v1 >> ((1) & 63))}
    v3 := make(map[string]uint32)
    v4 := make(map[string]string)
    var v5 int64 = -800
    var v6 uint16 = 2
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int8 = -82
        _ = fi12
        _ = v13
    }
    var v14 uint = 227
    v1 = v1
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 bool = (v14 == (v14 - v14))
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v6 -= (^(v6))
        var v18 uint = (v14 ^ (v14 ^ (v14 - (v14 - (v14 - 403)))))
        _ = fi17
        _ = v18
    }
    var v19 float64 = 1.5
    if true {
        v20 := GenS1{F0: 1, F1: 1, F2: 1}
        _ = v20
    }
    v1 = 553
    v21 := GenS1{F0: -1, F1: 18446744073709551615, F2: 724}
    if true {
        v22 := GenS1{F0: 0, F1: 4294967295, F2: 953}
        _ = v22
    }
    v23 := GenS1{F0: 32767, F1: 0, F2: 0}
    var v24 float64 = 0.0
    for fi25 := 0; fi25 < 9; fi25++ {
        if (0 == 128) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v26 uint8 = 127
            var v27 int16 = 42
            _ = v26
            _ = v27
        }
        if true {
            v28 := GenS1{F0: 1000, F1: 1, F2: uint32(v14)}
            if (false && false) {
                v29 := v4[("hello" + "foo bar")]
                var v30 float32 = -2.25
                v29 = v29
                _ = v29
                _ = v30
            } else {
                var v31 uint64 = 545
                var v32 int8 = -128
                for fi33 := 0; fi33 < 19; fi33++ {
                    if ((^(v5)) > -182) {
                        _ = len(v3)
                        if false {
                            if true {
                                v34 := GenS1{F0: 609, F1: (v31 & (^(v31))), F2: 171}
                                _ = v34
                            } else {
                                v32 += 42
                                v1 = v1
                                var v35 float64 = 0.0
                                _ = v35
                            }
                            var v36 string = ((("" + "") + "foo bar") + "foo bar")
                            _ = v36
                        } else {
                            switch (fi33 ^ fi25) {
                            case -9223372036854775808:
                                if true {
                                    v32 = int8(v5)
                                    v5 = int64(v32)
                                }
                            case 128:
                                var v37 int32 = -1
                                _ = v37
                            case 780:
                                if false {
                                    if true {
                                        if ((fi33 % ((fi33) | 1)) < -2147483648) {
                                            _ = len(v3)
                                            var v38 int64 = v5
                                            var v39 uint = (^(v14))
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v1 -= -2147483648
                                            fmt.Println(v6)
                                            if (uint(v31) < v14) {
                                                if true {
                                                    if (true && ((v5 * 0) > (v5 / (((v5 & -487)) | 1)))) {
                                                        switch "" {
                                                        case "hello":
                                                            var v40 int = (fi25 << (((^(v14))) & 63))
                                                            _ = v40
                                                        default:
                                                            if ((false && (((^(v1)) < (fi25 * genFn0())) && ((!false) || true))) || true) {
                                                                v41 := GenS1{F0: -337, F1: v31, F2: uint32(v31)}
                                                                fmt.Println(v14, fi25)
                                                                _ = v41
                                                            }
                                                        }
                                                        v42 := GenS1{F0: 32767, F1: v31, F2: 262}
                                                        if (-2.25 < -2.25) {
                                                            var v43 uint8 = 130
                                                            switch v1 {
                                                            case 4294967295:
                                                                v32 = v32
                                                            case 238:
                                                                var v44 int64 = v5
                                                                v45 := GenS2{F0: int64(v31), F1: v31, F2: -567}
                                                                v46 := GenS3{F0: (v31 - v31), F1: 2, F2: (v31 + v31)}
                                                                _ = v44
                                                                _ = v45
                                                                _ = v46
                                                            default:
                                                                var v47 int = fi25
                                                                _ = v47
                                                            }
                                                            _ = v43
                                                        }
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v48 int32 = -2147483648
                                                            _ = v48
                                                        case "hello":
                                                            v4[("hello" + "\n\t")] = ((("a" + "hello") + (("a" + "foo bar") + "a")) + "hello")
                                                        }
                                                        v32 += -128
                                                        v6 -= (v6 << (((v14 ^ (v14 - 767))) & 63))
                                                        _ = v42
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if (v6 != (v6 + 32767)) {
                                                        if false {
                                                            v5 -= v5
                                                        }
                                                        v49 := GenS1{F0: 32767, F1: (v31 * v31), F2: 191}
                                                        _ = v49
                                                    }
                                                    var v50 uint16 = (v6 * (v6 >> ((634) & 63)))
                                                    _ = v50
                                                }
                                            }
                                        }
                                        v14 += (v14 ^ 844)
                                        v19 = (0.1 / (((v24 + float64((3.14159 / (3.14159 - 0.0)))) * 3.14159) - v24))
                                    } else {
                                        v4["hello"] = ("foo bar" + "foo bar")
                                        if (!false) {
                                            var v51 uint32 = 2147483648
                                            var v52 uint8 = 0
                                            _ = len(v4)
                                            _ = v51
                                            _ = v52
                                        } else {
                                            v53 := GenS3{F0: v31, F1: 71, F2: 880}
                                            _ = v53
                                        }
                                    }
                                    if (v31 > (v31 << (((v14 >> ((v14) & 63))) & 63))) {
                                        var v54 uint = (v14 ^ 1)
                                        _ = v54
                                    }
                                    if (32767 >= 32767) {
                                        var v55 bool = false
                                        var v56 int32 = -1
                                        v32 += v32
                                        _ = v55
                                        _ = v56
                                    } else {
                                        var v57 bool = ((!((((float32(v24) / (1e10 - (float32((v19 * v24)) / 3.14159))) - 1e10) * 1e10) <= 0.1)) && (-32768 <= 1))
                                        v58 := GenS3{F0: (v31 | v31), F1: 728, F2: v31}
                                        _ = v57
                                        _ = v58
                                    }
                                }
                            default:
                                fmt.Println(v1)
                            }
                            var v59 int = 470
                            switch (fi25 - genFn0()) {
                            case -1:
                                var v60 bool = (-1 >= int16(v14))
                                _ = v60
                            case -9223372036854775808:
                                v32 = v32
                            default:
                                var v61 int32 = -424
                                _ = v61
                            }
                            _ = v59
                        }
                    }
                    _ = len(v4)
                    _ = fi33
                }
                _ = v31
                _ = v32
            }
            var v62 uint16 = (v6 >> ((v14) & 63))
            _ = v28
            _ = v62
        } else {
            var v63 string = "a"
            v2 = append(v2, fi25)
            _ = v63
        }
        if (v14 == (v14 >> (((v14 >> (((v14 & v14)) & 63))) & 63))) {
            switch ("foo bar" + "a") {
            case "foo bar":
                fmt.Println(fi25, v19, v14)
            case "a":
                v64 := GenS2{F0: 1, F1: 18446744073709551615, F2: 761}
                _ = v64
            case "hello":
                v65 := GenS1{F0: -921, F1: 0, F2: 139}
                _ = v65
            default:
                v66 := GenS4{F0: ((("a" + "\n\t") + (("a" + "hello") + "foo bar")) + ((("" + "\n\t") + (("foo bar" + "\n\t") + ("" + ("\n\t" + (("\n\t" + "foo bar") + "foo bar"))))) + "a"))}
                _ = v66
            }
        } else {
            if false {
                fmt.Println(v5, fi25)
                fmt.Println(v6, v19)
                var v67 int8 = 1
                _ = v67
            }
            switch -441 {
            case 0:
                for fi68 := 0; fi68 < 1; fi68++ {
                    _ = len(v3)
                    _ = fi68
                }
            }
        }
        _ = fi25
    }
    if false {
        v69 := v4[""]
        _ = v69
    }
    var v70 int32 = 361
    for fi71 := 0; fi71 < 15; fi71++ {
        v24 -= (v19 - v24)
        _ = fi71
    }
    for fi72 := 0; fi72 < 8; fi72++ {
        var v73 uint64 = 42
        switch "foo bar" {
        case "foo bar":
            if false {
                v19 = v24
            }
        case "\n\t":
            for fi74 := 0; fi74 < 12; fi74++ {
                switch "hello" {
                case "":
                    v19 = (-2.25 - v19)
                }
                v75 := GenS5{F0: ((("foo bar" + ("a" + "a")) + ("\n\t" + "a")) + ("\n\t" + ("a" + "foo bar"))), F1: v70, F2: (-(v70))}
                _ = fi74
                _ = v75
            }
        case "a":
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v5 += (v5 | (v5 % ((v5) | 1)))
            case "":
                var v76 uint16 = (^(v6))
                _ = v76
            case "a":
                if false {
                    v6 += v6
                    for fi77 := 0; fi77 < 14; fi77++ {
                        var v78 int16 = 1000
                        if (!(1e10 == (1e10 - (3.14159 * float32((1.5 + 0.1)))))) {
                            v19 = (0.1 * float64(1.5))
                        } else {
                            fmt.Println(v78, v70)
                            if false {
                                switch "foo bar" {
                                case "a":
                                    var v79 uint8 = 0
                                    _ = v79
                                case "foo bar":
                                    var v80 uint = (v14 / ((v14) | 1))
                                    _ = v80
                                case "":
                                    fmt.Println(v78)
                                }
                                if (v19 >= ((v24 / ((v19 - 0.0) / v24)) - v19)) {
                                    v81 := GenS6{F0: 0, F1: (^(v6))}
                                    _ = v81
                                }
                            } else {
                                var v82 uint = (^(v14))
                                _ = v82
                            }
                        }
                        _ = fi77
                        _ = v78
                    }
                } else {
                    v1 -= (-(v1))
                    var v83 float32 = 1e10
                    _ = v83
                }
            }
        default:
            v5 -= v5
        }
        if ((v24 / v24) == 0.0) {
            var v84 float32 = (float32(-2.25) * (1.5 + 0.1))
            switch "hello" {
            case "foo bar":
                switch 922 {
                case 2147483647:
                    if false {
                        v4[("" + ("\n\t" + ("foo bar" + "foo bar")))] = ""
                    }
                case -701:
                    fmt.Println(v24)
                }
            case "a":
                fmt.Println(v5, v73, v73)
            case "\n\t":
                v85 := GenS6{F0: 42, F1: v6}
                _ = v85
            default:
                var v86 uint64 = v73
                _ = v86
            }
            _ = v84
        }
        _ = fi72
        _ = v73
    }
    v14 += v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
