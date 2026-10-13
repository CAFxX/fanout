package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int32; F2 string }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 uint8 }
func genFn0(p0 int32, p1 uint16) uint8 {
    var v5 float64 = 1e10
    switch -2147483648 {
    case -715:
        v6 := GenS1{F0: -320, F1: 42, F2: 0}
        _ = v6
    }
    var v7 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 220
}

func main() {
    v1 := []string{"foo bar", (("\n\t" + "hello") + "hello"), "foo bar", ""}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    var v8 uint8 = 1
    v8 += v8
    if false {
        v9 := GenS1{F0: -505, F1: 4294967295, F2: 414}
        fmt.Println(v8)
        var v10 uint16 = 966
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -162, F1: 383, F2: -1}
    fmt.Println(v8)
    var v12 uint16 = 42
    v12 += v12
    v13 := GenS2{F0: 2147483647}
    for fi14 := 0; fi14 < 15; fi14++ {
        if true {
            v15 := GenS3{F0: 65535}
            v8 = v8
            var v16 uint8 = (v8 % ((v8) | 1))
            _ = v15
            _ = v16
        } else {
            switch (-(fi14)) {
            case -257:
                _ = len(v4)
            case -9223372036854775808:
                v17 := GenS4{F0: true}
                _ = v17
            case -2147483648:
                v12 -= v12
            default:
                var v18 uint16 = 707
                _ = v18
            }
        }
        v19 := GenS1{F0: 32767, F1: 42, F2: (fi14 * fi14)}
        _ = fi14
        _ = v19
    }
    switch 32768 {
    case 1:
        v20 := GenS1{F0: -48, F1: 707, F2: 360}
        _ = v20
    case -9223372036854775808:
        for fi21 := 0; fi21 < 11; fi21++ {
            switch ("" + "") {
            case "":
                if true {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 uint = 793
                        var v24 int64 = -1
                        v23 += v23
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            case "\n\t":
                v25 := GenS2{F0: 42}
                _ = v25
            default:
                if false {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    v12 -= 65535
                    fmt.Println(fi21, fi21)
                } else {
                    var v26 int = 128
                    v12 = v12
                    _ = v26
                }
            }
            var v27 int64 = 449
            _ = fi21
            _ = v27
        }
    default:
        if false {
            for fi28 := 0; fi28 < 5; fi28++ {
                v12 -= (v12 >> ((2) & 63))
                v12 -= (v12 << ((0) & 63))
                _ = fi28
            }
        } else {
            if true {
                if false {
                    var v29 int32 = 0
                    var v30 string = "hello"
                    if (1 > 42) {
                        _ = len(v3)
                        v29 = v29
                        for fi31 := 0; fi31 < 11; fi31++ {
                            _ = len(v1)
                            v30 = v30
                            _ = fi31
                        }
                    }
                    _ = v29
                    _ = v30
                }
                if false {
                    if true {
                        v8 += (^(v8))
                        var v32 int8 = -1
                        _ = v32
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v33 := GenS4{F0: (true && (!(!true)))}
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 20; fi34++ {
                                if ("foo bar" == "foo bar") {
                                    if true {
                                        v8 = (v8 << ((18446744073709551615) & 63))
                                        for fi35 := 0; fi35 < 9; fi35++ {
                                            switch fi35 {
                                            case 127:
                                                switch 42 {
                                                case -854:
                                                    switch "hello" {
                                                    case "hello":
                                                        v36 := GenS5{F0: 42, F1: true, F2: fi35}
                                                        _ = v36
                                                    case "":
                                                        v37 := GenS6{F0: uint32(v8)}
                                                        _ = v37
                                                    default:
                                                        if ((true || (4294967295 < 712)) && false) {
                                                            switch (fi35 | (fi35 + (fi34 | 323))) {
                                                            case 32767:
                                                                v38 := GenS7{F0: uint32(v12), F1: -1, F2: "\n\t"}
                                                                _ = v38
                                                            default:
                                                                v39 := GenS4{F0: ((2 >= 65535) || true)}
                                                                _ = v39
                                                            }
                                                        }
                                                        _ = len(v2)
                                                        _ = len(v3)
                                                    }
                                                case 128:
                                                    v40 := GenS4{F0: false}
                                                    _ = v40
                                                case 956:
                                                    _ = len(v1)
                                                }
                                            case 128:
                                                v8 += uint8(v12)
                                            default:
                                                if true {
                                                    fmt.Println(fi34, fi35, fi35)
                                                    if len(v1) > 0 {
                                                        v1[0] = ("hello" + "a")
                                                    }
                                                }
                                            }
                                            v12 = v12
                                            _ = fi35
                                        }
                                    } else {
                                        v41 := GenS7{F0: 1, F1: int32(fi34), F2: "a"}
                                        var v42 int = fi34
                                        var v43 uint8 = (v8 * (^(v8)))
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                    if (0.0 > (float32((1e10 - -2.25)) * 0.0)) {
                                        var v44 int16 = 1
                                        _ = v44
                                    }
                                    v2[(("hello" + ((("\n\t" + "hello") + (("foo bar" + "foo bar") + "\n\t")) + ((("a" + ("a" + "hello")) + "foo bar") + ("\n\t" + "hello")))) + (("foo bar" + (("" + "a") + "foo bar")) + ("foo bar" + ("" + ("" + "a")))))] = ("foo bar" + (("" + ("foo bar" + "foo bar")) + "foo bar"))
                                }
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    if true {
                                        var v46 int64 = -9223372036854775808
                                        v47 := GenS5{F0: -1, F1: false, F2: fi45}
                                        var v48 uint64 = 18446744073709551615
                                        var v49 uint = 255
                                        var v50 int8 = 42
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        var v51 bool = (0 > -1)
                                        v52 := GenS4{F0: ((!(0.0 == 1.5)) || v51)}
                                        _ = v51
                                        _ = v52
                                    }
                                    v12 += (^(v12))
                                    v12 -= v12
                                    var v53 float64 = 1.5
                                    v8 = (v8 / ((v8) | 1))
                                    var v54 int32 = 386
                                    var v55 int = 127
                                    v53 = (3.14159 * (v53 / v53))
                                    var v56 int16 = 1
                                    _ = fi45
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                v57 := GenS4{F0: (true && ((float64(1e10) - 0.0) > -2.25))}
                                _ = fi34
                                _ = v57
                            }
                        }
                    }
                    switch "" {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    default:
                        for fi58 := 0; fi58 < 17; fi58++ {
                            var v59 int16 = 1
                            if true {
                                for fi60 := 0; fi60 < 14; fi60++ {
                                    _ = len(v2)
                                    _ = fi60
                                }
                                fmt.Println(v8, fi58)
                            }
                            for fi61 := 0; fi61 < 19; fi61++ {
                                v59 += (v59 * v59)
                                v12 += v12
                                _ = fi61
                            }
                            _ = fi58
                            _ = v59
                        }
                    }
                }
                switch ("" + "foo bar") {
                case "":
                    var v62 int32 = -525
                    _ = v62
                default:
                    v63 := GenS8{F0: (v12 ^ (^(v12))), F1: 128}
                    _ = v63
                }
            } else {
                var v64 uint16 = (v12 >> ((840) & 63))
                switch 0 {
                case 4294967295:
                    if (3.14159 <= (1e10 / 3.14159)) {
                        if false {
                            v65 := GenS1{F0: 42, F1: 18446744073709551615, F2: 293}
                            var v66 float32 = (float32(3.14159) + 1e10)
                            for fi67 := 0; fi67 < 7; fi67++ {
                                var v68 int8 = -42
                                v69 := GenS9{F0: (^(v64))}
                                fmt.Println(v12)
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            _ = v65
                            _ = v66
                        } else {
                            fmt.Println(v8, v8)
                            v12 = (v64 | (v64 | 32767))
                        }
                    }
                default:
                    v12 += (v12 * v12)
                }
                _ = v64
            }
            switch 128 {
            case -9223372036854775808:
                switch 493 {
                case 0:
                    if (!(521 >= 0)) {
                        fmt.Println(v12, v8)
                        for fi70 := 0; fi70 < 3; fi70++ {
                            v71 := GenS9{F0: 32767}
                            if false {
                                var v72 int16 = -456
                                _ = v72
                            } else {
                                if (999 <= 0) {
                                    for fi73 := 0; fi73 < 4; fi73++ {
                                        if false {
                                            if (fi73 == fi70) {
                                                _ = len(v1)
                                                if (true && false) {
                                                    var v74 bool = false
                                                    v8 = (^(v8))
                                                    v12 = v12
                                                    var v75 float64 = 1.5
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v76 int8 = 1
                                                        _ = v76
                                                    case "":
                                                        v77 := GenS2{F0: 42}
                                                        _ = v77
                                                    case "foo bar":
                                                        if v74 {
                                                            var v78 uint = 0
                                                            if (v75 >= ((float64(1.5) + float64(0.0)) / v75)) {
                                                                v79 := GenS10{F0: 232, F1: 444}
                                                                v12 -= (^(v12))
                                                                _ = v79
                                                            } else {
                                                                _ = len(v1)
                                                                fmt.Println(v78)
                                                                v8 -= (^(v8))
                                                            }
                                                            v12 = 1
                                                            var v80 int64 = 352
                                                            _ = v78
                                                            _ = v80
                                                        } else {
                                                            v1 = append(v1, "hello")
                                                            fmt.Println(v8, v8)
                                                        }
                                                    default:
                                                        v8 += (v8 - (v8 / (((v8 - (v8 << ((557) & 63)))) | 1)))
                                                    }
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v81 := GenS11{F0: fi73}
                                                _ = v81
                                            }
                                            var v82 string = "a"
                                            v12 = (v12 + 2)
                                            var v83 bool = false
                                            _ = v82
                                            _ = v83
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v12 = (^(v12))
                                        }
                                        v12 -= v12
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        _ = fi73
                                    }
                                    var v84 uint = 18446744073709551615
                                    _ = v84
                                }
                            }
                            _ = fi70
                            _ = v71
                        }
                        v85 := GenS12{F0: 0, F1: v8}
                        _ = v85
                    }
                case 32768:
                    for fi86 := 0; fi86 < 1; fi86++ {
                        if true {
                            var v87 float64 = -2.25
                            _ = v87
                        }
                        _ = fi86
                    }
                default:
                    fmt.Println(v12)
                }
            case -2147483648:
                if (true && (-1 < 1)) {
                    v8 = v8
                    var v88 uint32 = 0
                    _ = v88
                } else {
                    v89 := GenS1{F0: 42, F1: 138, F2: -2147483648}
                    _ = v89
                }
            case -818:
                for fi90 := 0; fi90 < 8; fi90++ {
                    switch "hello" {
                    case "\n\t":
                        for fi91 := 0; fi91 < 9; fi91++ {
                            v92 := v4[285]
                            _ = fi91
                            _ = v92
                        }
                    default:
                        switch (fi90 | fi90) {
                        case -2147483648:
                            v12 += (^(v12))
                        case 620:
                            v12 = (^(v12))
                        }
                    }
                    _ = fi90
                }
            default:
                if (true && false) {
                    var v93 int8 = -128
                    _ = v93
                } else {
                    v8 += (v8 | v8)
                }
            }
        }
    }
    var v94 uint64 = 771
    v95 := GenS3{F0: 177}
    switch ("hello" + "foo bar") {
    case "":
        var v96 int = 127
        _ = v96
    }
    var v97 uint32 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v94
    _ = v95
    _ = v97
}
