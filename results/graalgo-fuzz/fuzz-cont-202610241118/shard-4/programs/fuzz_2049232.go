package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16 }
func genFn0() int32 {
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int64 = -491
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int = (^(fi9))
            v11 := []int{v10, (fi9 & -727)}
            fmt.Println(fi9, fi9, v10)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    switch 995 {
    case 491:
        v12 := []int{1}
        _ = v12
    case 9223372036854775807:
        var v13 string = "hello"
        _ = v13
    case 4294967295:
        v14 := GenS1{F0: 897, F1: -204, F2: 342}
        _ = v14
    }
    _ = v6
    return 2147483647
}

func main() {
    v1 := []uint32{1, 971, 2147483648, 706}
    var v2 int32 = 0
    v3 := []string{}
    v4 := make(map[string]string)
    var v5 uint32 = 1
    for fi15 := 0; fi15 < 18; fi15++ {
        if (("a" >= ("a" + ("a" + ("\n\t" + "\n\t")))) && false) {
            v5 = v5
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            if true {
                var v17 int64 = 9223372036854775807
                if ((-95 != -32768) || ((false || true) && (!((!(97 == 16)) && (42 != 1000))))) {
                    v18 := GenS2{F0: (fi16 / ((fi16) | 1))}
                    switch ("foo bar" + "foo bar") {
                    case "":
                        var v19 int8 = -128
                        v2 = (v2 * -2147483648)
                        var v20 int32 = (v2 % ((v2) | 1))
                        _ = v19
                        _ = v20
                    default:
                        switch (fi15 | fi16) {
                        case 42:
                            if false {
                                if ((((float64(3.14159) - (1e10 / 1e10)) == float64(1.5)) || (4294967295 == 1)) && (32767 != 65535)) {
                                    v21 := GenS2{F0: 128}
                                    v1 = append(v1, (v5 & (v5 * (v5 << ((1) & 63)))))
                                    if (0 == 4294967295) {
                                        switch "\n\t" {
                                        case "hello":
                                            fmt.Println(v2, v5)
                                        case "":
                                            v22 := v4["hello"]
                                            v23 := v4[((("foo bar" + "hello") + ((v22 + ((v22 + ((v22 + v22) + v22)) + (v22 + "\n\t"))) + "a")) + v22)]
                                            fmt.Println(v17)
                                            _ = v22
                                            _ = v23
                                        case "\n\t":
                                            fmt.Println(v2, fi16, v5)
                                        }
                                        delete(v4, ("foo bar" + ("a" + ("a" + ("hello" + "foo bar")))))
                                    } else {
                                        if true {
                                            v5 += (v5 ^ v5)
                                            var v24 uint8 = 0
                                            _ = v24
                                        }
                                        _ = len(v1)
                                        v25 := GenS2{F0: (fi16 ^ (fi16 << ((42) & 63)))}
                                        _ = v25
                                    }
                                    _ = v21
                                } else {
                                    var v26 int = (fi16 + (fi15 | fi16))
                                    _ = v26
                                }
                                switch (fi16 + fi16) {
                                case -2147483648:
                                    v1 = append(v1, (^(v5)))
                                case 950:
                                    switch (("" + "foo bar") + "") {
                                    case "\n\t":
                                        v2 += v2
                                    case "foo bar":
                                        v27 := GenS3{F0: 255}
                                        _ = v27
                                    default:
                                        v5 = v5
                                    }
                                case 128:
                                    v28 := v4["a"]
                                    _ = v28
                                default:
                                    if false {
                                        v17 = (v17 ^ v17)
                                        fmt.Println(v17)
                                        if (false && (255 > 65535)) {
                                            fmt.Println(v17)
                                        } else {
                                            var v29 uint16 = 2
                                            _ = v29
                                        }
                                        if (4294967295 > 18446744073709551615) {
                                            v1 = append(v1, 4294967295)
                                            v30 := GenS4{F0: 18446744073709551615, F1: -32768}
                                            _ = v30
                                        }
                                    } else {
                                        v5 = v5
                                    }
                                }
                            } else {
                                if false {
                                    v2 = (v2 >> ((65535) & 63))
                                    if ((1e10 / 1e10) != (-2.25 - 0.0)) {
                                        if len(v1) > 0 {
                                            v1[0] = (v5 >> ((65535) & 63))
                                        }
                                        var v31 uint = 4294967295
                                        _ = v31
                                    }
                                }
                            }
                        default:
                            v17 += -281
                        }
                    }
                    _ = v18
                } else {
                    v2 += 1
                    var v32 uint64 = 4294967295
                    _ = v32
                }
                _ = v17
            }
            var v33 int16 = 42
            v34 := GenS3{F0: 62}
            _ = fi16
            _ = v33
            _ = v34
        }
        _ = len(v1)
        _ = fi15
    }
    if (42 >= 32768) {
        var v35 int32 = (v2 << ((255) & 63))
        if len(v1) > 0 {
            v1[0] = v5
        }
        _ = v35
    } else {
        if (205 >= 117) {
            v4[("" + "")] = "foo bar"
            fmt.Println(v2, v5)
        }
    }
    v36 := GenS1{F0: 621, F1: 0, F2: 46}
    v2 = genFn0()
    v37 := GenS5{F0: 42, F1: v2, F2: (v5 - 42)}
    v5 -= 638
    v2 += v2
    v5 -= 4294967295
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if (-9223372036854775808 >= -186) {
        if (!true) {
            v2 += genFn0()
        }
        v3 = append(v3, "")
    }
    v5 += v5
    var v38 int8 = 100
    switch 2147483647 {
    case 32768:
        for fi39 := 0; fi39 < 14; fi39++ {
            v2 -= v2
            _ = fi39
        }
    case 2147483647:
        switch 9223372036854775807 {
        case 9223372036854775807:
            if (!false) {
                var v40 uint32 = v5
                v41 := GenS5{F0: 41, F1: (^(v2)), F2: 1}
                v5 += (^(v40))
                _ = v40
                _ = v41
            }
        case 539:
            fmt.Println(v2, v38, v38)
        case -1:
            switch ("a" + (("hello" + "") + "a")) {
            case "foo bar":
                switch "a" {
                case "\n\t":
                    var v42 int8 = v38
                    _ = v42
                case "hello":
                    switch ("hello" + "") {
                    case "\n\t":
                        v5 = v5
                    default:
                        var v43 int32 = v2
                        _ = v43
                    }
                case "foo bar":
                    if (((1.5 + 1e10) != 3.14159) && (true && (2 <= 582))) {
                        var v44 int8 = -1
                        v5 = v5
                        if true {
                            var v45 float32 = float32((1e10 * -2.25))
                            fmt.Println(v38, v2)
                            var v46 uint32 = 2147483648
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v38)
                            v44 -= v44
                            if true {
                                v47 := GenS6{F0: (v5 << ((776) & 63)), F1: -654, F2: 4294967295}
                                var v48 uint8 = 255
                                _ = v47
                                _ = v48
                            } else {
                                if (v5 == (v5 % ((2) | 1))) {
                                    switch 2147483647 {
                                    case 408:
                                        if (-2.25 > 0.0) {
                                            var v49 int16 = 42
                                            var v50 uint32 = (v5 ^ 2)
                                            _ = v49
                                            _ = v50
                                        }
                                    case 421:
                                        for fi51 := 0; fi51 < 4; fi51++ {
                                            v44 += v38
                                            v2 = v2
                                            _ = fi51
                                        }
                                    case 192:
                                        if (!false) {
                                            fmt.Println(v5, v2, v5)
                                            v5 = (v5 & (v5 - 0))
                                            var v52 uint16 = 255
                                            _ = v52
                                        } else {
                                            for fi53 := 0; fi53 < 20; fi53++ {
                                                v54 := v4["\n\t"]
                                                v54 = ((("\n\t" + v54) + v54) + "foo bar")
                                                fmt.Println(fi53)
                                                _ = fi53
                                                _ = v54
                                            }
                                            delete(v4, ("\n\t" + "\n\t"))
                                            if len(v1) > 0 {
                                                v1[0] = 65535
                                            }
                                        }
                                    }
                                    for fi55 := 0; fi55 < 16; fi55++ {
                                        if false {
                                            v56 := GenS4{F0: 444, F1: 952}
                                            var v57 float32 = 3.14159
                                            v38 -= 11
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = fi55
                                    }
                                } else {
                                    switch "" {
                                    case "a":
                                        if (genFn0() < (v2 << ((443) & 63))) {
                                            var v58 uint64 = 18446744073709551615
                                            var v59 uint16 = 1
                                            _ = v58
                                            _ = v59
                                        } else {
                                            var v60 int8 = v38
                                            _ = v60
                                        }
                                    case "\n\t":
                                        v61 := GenS4{F0: 0, F1: 804}
                                        _ = v61
                                    default:
                                        if (-9223372036854775808 != -153) {
                                            fmt.Println(v2)
                                            var v62 int8 = 100
                                            _ = v62
                                        } else {
                                            var v63 bool = (1 != 4294967295)
                                            var v64 uint = 628
                                            if ((^(v38)) < (v38 / ((v44) | 1))) {
                                                fmt.Println(v2, v44, v5)
                                            } else {
                                                var v65 string = ("hello" + ("a" + "a"))
                                                switch v65 {
                                                case "foo bar":
                                                    v64 = v64
                                                default:
                                                    v66 := GenS7{F0: 42, F1: uint16(v64), F2: 0}
                                                    _ = v66
                                                }
                                                var v67 string = "hello"
                                                _ = v65
                                                _ = v67
                                            }
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    switch 572 {
                                    case 395:
                                        v44 += v38
                                    case 128:
                                        for fi68 := 0; fi68 < 20; fi68++ {
                                            v69 := GenS8{F0: 128}
                                            var v70 int64 = 858
                                            _ = fi68
                                            _ = v69
                                            _ = v70
                                        }
                                    case 0:
                                        v71 := GenS9{F0: 32768, F1: 4294967295}
                                        _ = v71
                                    }
                                    var v72 uint64 = 492
                                    _ = v72
                                }
                                if true {
                                    v5 = 0
                                    v2 -= 0
                                    v2 = 2147483647
                                }
                            }
                        }
                        _ = v44
                    }
                }
            case "\n\t":
                for fi73 := 0; fi73 < 12; fi73++ {
                    v5 -= 1
                    if (!false) {
                        var v74 uint16 = 685
                        if true {
                            v38 -= (v38 - (v38 % ((127) | 1)))
                            if (!(genFn0() != v2)) {
                                var v75 string = ("foo bar" + "foo bar")
                                v76 := GenS7{F0: 310, F1: v74, F2: -842}
                                _ = v75
                                _ = v76
                            }
                        }
                        _ = v74
                    } else {
                        v77 := GenS4{F0: 18446744073709551615, F1: 1000}
                        _ = v77
                    }
                    _ = fi73
                }
            default:
                v78 := GenS8{F0: 255}
                _ = v78
            }
        default:
            var v79 uint32 = v5
            _ = v79
        }
    case 9223372036854775807:
        fmt.Println(v38, v38)
    default:
        fmt.Println(v2, v2, v38)
    }
    v80 := GenS10{F0: -1}
    _ = len(v3)
    v81 := GenS11{F0: 1000}
    var v82 uint64 = 0
    switch "foo bar" {
    case "hello":
        for fi83 := 0; fi83 < 14; fi83++ {
            for fi84 := 0; fi84 < 4; fi84++ {
                v85 := GenS10{F0: 9223372036854775807}
                if len(v1) > 0 {
                    v1[0] = v5
                }
                _ = fi84
                _ = v85
            }
            _ = fi83
        }
    case "\n\t":
        v82 = v82
    case "":
        var v86 uint64 = v82
        _ = v86
    default:
        var v87 float64 = (float64(0.1) / 1.5)
        _ = v87
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
}
