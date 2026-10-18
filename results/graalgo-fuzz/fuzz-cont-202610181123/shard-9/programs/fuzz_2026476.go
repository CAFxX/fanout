package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 bool; F1 uint64; F2 uint }
type GenS11 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8; F1 int }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 string }
func genFn0(p0 string) uint64 {
    v7 := GenS1{F0: 42, F1: 123}
    if (!(-1 <= -2147483648)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS2{F0: 65535, F1: 65535}
            p0 = (p0 + p0)
            _ = fi8
            _ = v9
        }
        var v10 int16 = 32767
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = 42
            v13 := GenS3{F0: (v12 / (((v12 | v12)) | 1)), F1: 370, F2: -1}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        v14 := GenS4{F0: 404}
        v15 := []int{}
        if false {
            var v16 uint64 = 18446744073709551615
            if (true && (false && (-1 > -128))) {
                for fi17 := 0; fi17 < 17; fi17++ {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v15 = append(v15, (fi17 >> ((889) & 63)))
                        v16 -= 42
                        var v19 bool = false
                        var v20 float32 = 0.0
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 1; fi21++ {
                        var v22 uint8 = uint8(v16)
                        var v23 float64 = float64((0.0 + (((((((1e10 - 0.1) + 1.5) / (0.1 - -2.25)) - (float32(1.5) / 1.5)) - 0.0) - 0.0) - float32(1e10))))
                        var v24 uint = 201
                        v25 := GenS5{F0: -37, F1: (false && false), F2: v24}
                        v26 := GenS4{F0: 0}
                        if (((1.5 + 1e10) * float64(3.14159)) > (-2.25 + 0.0)) {
                            var v27 bool = ((!((!(!((v24 / ((v24) | 1)) > v24))) || true)) && (!(1e10 >= 0.0)))
                            v22 += v22
                            switch fi17 {
                            case 127:
                                var v28 float32 = 0.0
                                _ = v28
                            case 131:
                                switch p0 {
                                case "foo bar":
                                    if v27 {
                                        v24 += v24
                                        var v29 int32 = -278
                                        _ = v29
                                    } else {
                                        var v30 uint8 = 128
                                        var v31 uint16 = 480
                                        _ = v30
                                        _ = v31
                                    }
                                case "":
                                    v32 := GenS2{F0: 18446744073709551615, F1: 65535}
                                    p0 = p0
                                    _ = len(v15)
                                    _ = v32
                                default:
                                    fmt.Println(v16, fi21, v16)
                                }
                            case 32767:
                                if true {
                                    v15 = append(v15, 32767)
                                    p0 = p0
                                }
                            default:
                                p0 = p0
                            }
                            p0 = p0
                            _ = v27
                        } else {
                            var v33 float64 = (0.1 * ((v23 + 0.0) * v23))
                            v34 := GenS6{F0: 42, F1: "\n\t"}
                            v22 -= 128
                            _ = v33
                            _ = v34
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(v16, p0, v16)
                    _ = fi17
                }
                switch 862 {
                case 4294967295:
                    var v35 bool = true
                    _ = v35
                case 9223372036854775807:
                    v16 += (v16 << ((2) & 63))
                case 637:
                    if true {
                        fmt.Println(v16, v16)
                        var v36 uint16 = 2
                        _ = v36
                    }
                default:
                    v37 := GenS7{F0: 311, F1: 1000}
                    _ = v37
                }
            }
            _ = v16
        }
        _ = v14
        _ = v15
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS1{F0: 0, F1: 106}
        p0 = (((p0 + p0) + ((p0 + ("hello" + "hello")) + p0)) + p0)
        _ = fi38
        _ = v39
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1() int {
    var v40 int64 = 427
    if ((!((42 <= 323) || (("foo bar" + "a") == ("" + "\n\t")))) || (((1e10 / 0.1) - -2.25) <= 3.14159)) {
        fmt.Println(v40)
    } else {
        v40 = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        if (v40 > (v40 ^ v40)) {
            v42 := []int{(fi41 * 0), (fi41 | fi41), 108}
            var v43 int16 = 0
            _ = v42
            _ = v43
        } else {
            var v44 uint32 = 319
            _ = v44
        }
        for fi45 := 0; fi45 < 2; fi45++ {
            v46 := GenS8{F0: (fi45 | -9223372036854775808), F1: (false && true)}
            var v47 int8 = 42
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi41
    }
    v48 := []int{2147483647, 769}
    if false {
        if (!false) {
            var v49 int64 = v40
            _ = v49
        }
        var v50 int8 = 22
        var v51 uint8 = 1
        _ = v50
        _ = v51
    }
    _ = v40
    _ = v48
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 42
    v3 := []int{0, 9223372036854775807, 42, -172}
    var v4 int64 = 555
    var v5 uint = 2
    var v6 uint64 = (v2 - 461)
    switch -294 {
    case -2147483648:
        if (v4 == (v4 / ((1) | 1))) {
            v5 = v5
            var v52 int32 = 0
            switch 965 {
            case 9223372036854775807:
                _ = len(v1)
            case 2147483647:
                var v53 int = -39
                _ = v53
            default:
                if ("hello" >= "a") {
                    _ = len(v3)
                    for fi54 := 0; fi54 < 1; fi54++ {
                        switch "\n\t" {
                        case "\n\t":
                            for fi55 := 0; fi55 < 1; fi55++ {
                                v2 -= (v2 - (v6 ^ (^(v6))))
                                v5 = v5
                                v56 := GenS9{F0: v5, F1: (1000 != 32767)}
                                if len(v3) > 0 {
                                    v3[0] = fi54
                                }
                                _ = fi55
                                _ = v56
                            }
                        case "a":
                            var v57 bool = true
                            _ = v57
                        default:
                            if true {
                                var v58 float32 = 1.5
                                if (!true) {
                                    fmt.Println(v52, v52, v4)
                                    if (true || true) {
                                        for fi59 := 0; fi59 < 10; fi59++ {
                                            v60 := GenS10{F0: false, F1: (v6 + genFn0("\n\t")), F2: v5}
                                            _ = fi59
                                            _ = v60
                                        }
                                    } else {
                                        for fi61 := 0; fi61 < 5; fi61++ {
                                            v62 := GenS7{F0: v2, F1: -338}
                                            var v63 int8 = 42
                                            if true {
                                                var v64 uint16 = 2
                                                if ((1000 <= 42) || ((v5 & (v5 | (v5 >> ((v5) & 63)))) != 2)) {
                                                    switch 969 {
                                                    case 32767:
                                                        if (!((v64 % ((v64) | 1)) >= 395)) {
                                                            v65 := GenS11{F0: 4294967295, F1: 71, F2: 127}
                                                            v66 := GenS1{F0: 65535, F1: 127}
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    if true {
                                                        fmt.Println(v63, fi61, v64)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if true {
                                                            var v67 uint16 = 59
                                                            switch 128 {
                                                            case 128:
                                                                var v68 float32 = v58
                                                                switch ("hello" + "") {
                                                                case "foo bar":
                                                                    v63 = (v63 >> ((uint(v6)) & 63))
                                                                case "hello":
                                                                    var v69 uint8 = 0
                                                                    _ = v69
                                                                default:
                                                                    if ((fi61 / (((fi61 ^ -9223372036854775808)) | 1)) != genFn1()) {
                                                                        v1[65535] = ("" + "")
                                                                    } else {
                                                                        var v70 string = "hello"
                                                                        if ((0.0 - (0.0 / (((((1e10 * (0.0 - 1.5)) / -2.25) + 1e10) - 1.5) + ((1.5 - 0.1) + 1.5)))) < 1e10) {
                                                                            fmt.Println(v64, fi61)
                                                                        }
                                                                        fmt.Println(v2)
                                                                        if false {
                                                                            switch v70 {
                                                                            case "\n\t":
                                                                                v71 := GenS8{F0: (fi61 | fi61), F1: (uint(v2) == v5)}
                                                                                _ = v71
                                                                            case "a":
                                                                                switch -9223372036854775808 {
                                                                                case 32768:
                                                                                    var v72 bool = false
                                                                                    _ = v72
                                                                                case 128:
                                                                                    var v73 int16 = 374
                                                                                    v4 = (v4 - int64(v73))
                                                                                    switch (v70 + (v70 + v70)) {
                                                                                    case "foo bar":
                                                                                        switch (("" + ((v70 + v70) + "\n\t")) + ((((v70 + v70) + (("hello" + "a") + (v70 + v70))) + v70) + "\n\t")) {
                                                                                        case "\n\t":
                                                                                            v74 := GenS6{F0: v4, F1: "hello"}
                                                                                            _ = v74
                                                                                        case "a":
                                                                                            v67 = 65535
                                                                                        }
                                                                                    }
                                                                                    _ = v73
                                                                                default:
                                                                                    v75 := GenS12{F0: (v6 + (v2 - 4294967295))}
                                                                                    var v76 int16 = -1
                                                                                    if true {
                                                                                        v67 = 348
                                                                                        if true {
                                                                                            var v77 int64 = (v4 % (((v4 + (v4 | 9223372036854775807))) | 1))
                                                                                            fmt.Println(v67)
                                                                                            var v78 string = "\n\t"
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        }
                                                                                    }
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                            case "":
                                                                                if (!(!true)) {
                                                                                    v4 = -1
                                                                                }
                                                                            default:
                                                                                if (!(((1e10 * 0.1) / 1.0) > (1e10 + 3.14159))) {
                                                                                    fmt.Println(v5, v67, fi54)
                                                                                    switch v70 {
                                                                                    case "":
                                                                                        if ((v63 % ((v63) | 1)) != v63) {
                                                                                            delete(v1, 2)
                                                                                            switch ("a" + v70) {
                                                                                            case "hello":
                                                                                                var v79 float64 = float64(0.1)
                                                                                                _ = v79
                                                                                            case "a":
                                                                                                v63 = v63
                                                                                            case "\n\t":
                                                                                                fmt.Println(v6, v70, v63)
                                                                                            default:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = fi61
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            switch (fi61 ^ (fi61 >> ((v5) & 63))) {
                                                                                            case 32767:
                                                                                                var v80 uint16 = v64
                                                                                                _ = v80
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            v64 -= (^(v64))
                                                                                        }
                                                                                    case "foo bar":
                                                                                        fmt.Println(v4, v5)
                                                                                    default:
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    }
                                                                                    v52 += 2147483647
                                                                                } else {
                                                                                    var v81 int8 = v63
                                                                                    v68 += (v58 * (1e10 + 3.14159))
                                                                                    _ = v81
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v82 int8 = (v63 & 127)
                                                                            v83 := GenS4{F0: v2}
                                                                            v84 := GenS3{F0: 2, F1: 297, F2: 1}
                                                                            _ = v82
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (-(fi61))
                                                                }
                                                                _ = v68
                                                            case 1:
                                                                if (2147483647 != -90) {
                                                                    fmt.Println(v2, v2)
                                                                    var v85 int16 = -826
                                                                    _ = v85
                                                                } else {
                                                                    fmt.Println(v6, v2, v52)
                                                                    v5 -= (v5 >> ((383) & 63))
                                                                    v2 += (v2 % ((v6) | 1))
                                                                }
                                                            }
                                                            var v86 int32 = v52
                                                            v87 := GenS13{F0: (v63 | v63), F1: (fi54 / (((fi54 << (((v5 & (^(v5)))) & 63))) | 1))}
                                                            _ = v67
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            v88 := v1[2]
                                                            _ = len(v1)
                                                            var v89 float32 = -2.25
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                    if false {
                                                        v90 := GenS14{F0: v5}
                                                        _ = v90
                                                    }
                                                }
                                                _ = v64
                                            }
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                        var v91 uint16 = 1
                                        _ = v91
                                    }
                                    v92 := v1[2147483648]
                                    _ = v92
                                }
                                _ = v58
                            } else {
                                var v93 uint16 = uint16(v6)
                                _ = v93
                            }
                        }
                        _ = fi54
                    }
                    for fi94 := 0; fi94 < 15; fi94++ {
                        _ = len(v1)
                        v6 = genFn0(("foo bar" + "a"))
                        _ = fi94
                    }
                } else {
                    v95 := GenS4{F0: (^(v6))}
                    _ = v95
                }
            }
            _ = v52
        }
    }
    fmt.Println(v2, v4)
    v5 = (v5 ^ (v5 & v5))
    switch 32767 {
    case 32768:
        var v96 int64 = v4
        _ = v96
    case 866:
        v6 += v6
    case 1:
        v5 = 490
    }
    fmt.Println(v2, v5)
    if true {
        if (false || false) {
            v97 := v1[uint32(v6)]
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v98 uint64 = (v2 << ((951) & 63))
            _ = v97
            _ = v98
        } else {
            var v99 string = "hello"
            v5 = v5
            _ = v99
        }
        v2 += 2
    } else {
        v5 -= 409
    }
    v100 := GenS15{F0: "\n\t"}
    var v101 int8 = 0
    for fi102 := 0; fi102 < 7; fi102++ {
        for fi103 := 0; fi103 < 11; fi103++ {
            if (!(-1 <= 1)) {
                v101 += v101
                v4 += (-(v4))
            } else {
                var v104 uint8 = 128
                if false {
                    var v105 int64 = (v4 + (-(v4)))
                    var v106 string = "foo bar"
                    v1[997] = v106
                    v107 := GenS6{F0: v4, F1: ("" + v106)}
                    fmt.Println(v104, v6, fi103)
                    _ = v105
                    _ = v106
                    _ = v107
                }
                _ = v104
            }
            v4 -= 42
            _ = fi103
        }
        _ = fi102
    }
    var v108 int = -388
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v100
    _ = v101
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
