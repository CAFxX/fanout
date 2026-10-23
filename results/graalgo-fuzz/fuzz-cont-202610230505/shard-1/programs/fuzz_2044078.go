package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int32; F1 string; F2 int8 }
type GenS9 struct { F0 bool; F1 int; F2 int16 }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 int, p1 int) uint16 {
    v4 := []int{p0, -466, p0}
    v4 = append(v4, p0)
    var v5 int16 = 32767
    var v6 uint32 = 4294967295
    var v7 uint32 = (v6 * (v6 & (v6 + (v6 << ((18446744073709551615) & 63)))))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 972
}
func genFn1(p0 int16, p1 uint32) bool {
    var v8 uint64 = 695
    var v9 string = ("foo bar" + "foo bar")
    if ((!true) && true) {
        v10 := GenS1{F0: p0, F1: (1 < 128)}
        v11 := []int{217}
        var v12 uint8 = 1
        _ = v10
        _ = v11
        _ = v12
    } else {
        var v13 int64 = -9223372036854775808
        v14 := []int{127, 9223372036854775807, 32767}
        _ = v13
        _ = v14
    }
    var v15 int64 = -372
    _ = p0
    _ = p1
    _ = v15
    _ = v8
    _ = v9
    return (!true)
}

func main() {
    var v1 uint16 = 2
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        switch (-(fi17)) {
        case 1:
            fmt.Println(v2, v2, v2)
        case 174:
            v18 := GenS2{F0: (v2 ^ v2)}
            _ = v18
        case -1:
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 bool = (!(((((float32(-2.25) >= 3.14159) || true) && true) || (1 <= 1)) && true))
                var v21 uint64 = 1
                var v22 float32 = (((((float32((0.1 / float64((1.5 + (0.0 * 0.1))))) - ((1e10 / 1.0) + 0.0)) + 1e10) + (float32(((-2.25 + -2.25) - (0.1 + 1.5))) - float32(0.1))) / ((((float32((0.1 - (0.1 + ((0.0 * ((0.0 / 1.0) * (0.1 * (0.1 + (((-2.25 / (0.1 - 3.14159)) + 1.5) + ((1e10 / (-2.25 / 1e10)) + ((1.5 / 3.14159) * float64(((1.5 * (-2.25 * (float32(((((3.14159 / (0.0 + 0.1)) - float64(1e10)) - ((0.0 + float64(1e10)) + 1e10)) * -2.25)) - 1e10))) * 0.0))))))))) - (3.14159 - (3.14159 * ((0.0 - 0.1) + 1e10))))))) - (-2.25 / 1.0)) * (0.1 + (3.14159 + (1e10 / 3.14159)))) + (1e10 * 0.0)) / -2.25)) * 1e10)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        default:
            var v23 uint64 = 42
            _ = v23
        }
        _ = fi17
    }
    var v24 float64 = (float64(((3.14159 / 3.14159) * (0.0 + -2.25))) + 0.0)
    v25 := []int{int(v2)}
    fmt.Println(v1)
    for fi26 := 0; fi26 < 9; fi26++ {
        switch "hello" {
        case "":
            for fi27 := 0; fi27 < 20; fi27++ {
                v25 = append(v25, fi26)
                if (!true) {
                    switch "\n\t" {
                    case "\n\t":
                        v2 += v2
                    case "foo bar":
                        var v28 int16 = 251
                        _ = v28
                    }
                    v2 = v2
                    var v29 bool = false
                    _ = v29
                }
                v1 -= (v1 & v1)
                _ = fi27
            }
        case "\n\t":
            v2 = v2
        case "foo bar":
            v1 += (v1 / ((genFn0(fi26, (fi26 | 127))) | 1))
        default:
            _ = len(v25)
        }
        _ = fi26
    }
    if true {
        if true {
            if (true && ((!true) && (false || (32767 != 1000)))) {
                fmt.Println(v24, v24, v24)
                fmt.Println(v24, v1, v1)
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 string = "foo bar"
                    _ = fi30
                    _ = v31
                }
            } else {
                if (135 <= uint8(v1)) {
                    switch ("" + "a") {
                    case "foo bar":
                        var v32 uint = 1
                        _ = v32
                    case "\n\t":
                        var v33 float64 = v24
                        _ = v33
                    case "hello":
                        v34 := GenS3{F0: 0}
                        _ = v34
                    default:
                        if false {
                            if true {
                                switch -2147483648 {
                                case -9223372036854775808:
                                    var v35 uint8 = 65
                                    _ = v35
                                case 1:
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                default:
                                    if (!true) {
                                        switch 127 {
                                        case 127:
                                            for fi36 := 0; fi36 < 11; fi36++ {
                                                for fi37 := 0; fi37 < 15; fi37++ {
                                                    var v38 uint = 42
                                                    delete(v3, -535)
                                                    v38 = 1
                                                    _ = fi37
                                                    _ = v38
                                                }
                                                switch (fi36 | (-(fi36))) {
                                                case -9223372036854775808:
                                                    var v39 float64 = v24
                                                    _ = v39
                                                case 331:
                                                    switch -660 {
                                                    case 127:
                                                        if true {
                                                            var v40 bool = false
                                                            for fi41 := 0; fi41 < 18; fi41++ {
                                                                if false {
                                                                    var v42 bool = true
                                                                    var v43 uint = 65535
                                                                    v42 = true
                                                                    if true {
                                                                        if v40 {
                                                                            v40 = (-1 == (v2 >> (((v43 << ((42) & 63))) & 63)))
                                                                            if true {
                                                                                if v42 {
                                                                                    var v44 string = ("hello" + ("\n\t" + "a"))
                                                                                    v45 := GenS2{F0: (v2 ^ v2)}
                                                                                    v24 -= 3.14159
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                } else {
                                                                                    v46 := GenS3{F0: 865}
                                                                                    switch fi41 {
                                                                                    case 0:
                                                                                        var v47 float32 = -2.25
                                                                                        _ = v47
                                                                                    default:
                                                                                        if (((v2 - (v2 % ((-1) | 1))) == v2) && (int32(v43) < 660)) {
                                                                                            v42 = v40
                                                                                            fmt.Println(v40, v42, v43)
                                                                                        }
                                                                                    }
                                                                                    _ = v46
                                                                                }
                                                                                v43 -= 0
                                                                                if ((fi41 - (fi36 >> ((v43) & 63))) <= -9223372036854775808) {
                                                                                    fmt.Println(v42, fi41, fi36)
                                                                                    v2 = v2
                                                                                    if len(v25) >= 2 {
                                                                                        v25 = v25[1:len(v25)]
                                                                                    }
                                                                                } else {
                                                                                    if (0 == 100) {
                                                                                        if (!(!(-21 < -295))) {
                                                                                            v48 := GenS1{F0: 42, F1: true}
                                                                                            var v49 uint16 = 42
                                                                                            v2 = v2
                                                                                            if ((fi36 ^ 0) <= (fi41 >> ((v43) & 63))) {
                                                                                                if (v24 > ((v24 - 0.0) * (v24 - -2.25))) {
                                                                                                    var v50 int64 = (^(v2))
                                                                                                    _ = v50
                                                                                                }
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        } else {
                                                                                            var v51 float32 = (-2.25 * ((0.0 / ((3.14159 / 1.5) / 1.0)) / 1.0))
                                                                                            _ = v51
                                                                                        }
                                                                                        if (!true) {
                                                                                            if (int32(fi41) > 0) {
                                                                                                v1 = v1
                                                                                                var v52 float64 = v24
                                                                                                _ = v52
                                                                                            }
                                                                                            var v53 uint64 = 1
                                                                                            v43 = v43
                                                                                            var v54 uint32 = 1
                                                                                            switch fi41 {
                                                                                            case 4294967295:
                                                                                                if (!(v42 || v40)) {
                                                                                                    if ((v53 == v53) || (v42 || v40)) {
                                                                                                        v40 = (("" + "") == "a")
                                                                                                        if false {
                                                                                                            if true {
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    var v55 int = (fi36 * (fi41 / (((fi41 - -154)) | 1)))
                                                                                                                    _ = v55
                                                                                                                case "foo bar":
                                                                                                                    v56 := GenS3{F0: 42}
                                                                                                                    _ = v56
                                                                                                                default:
                                                                                                                    v57 := GenS4{F0: 710, F1: 2147483647, F2: -9223372036854775808}
                                                                                                                    if ((fi36 % ((1) | 1)) >= (fi36 >> ((0) & 63))) {
                                                                                                                        if v42 {
                                                                                                                            if v42 {
                                                                                                                                switch "" {
                                                                                                                                case "a":
                                                                                                                                    v54 = v54
                                                                                                                                case "":
                                                                                                                                    var v58 uint = (v43 % ((v43) | 1))
                                                                                                                                    _ = v58
                                                                                                                                case "\n\t":
                                                                                                                                    v59 := GenS1{F0: -1, F1: (-102 > 100)}
                                                                                                                                    _ = v59
                                                                                                                                default:
                                                                                                                                    fmt.Println(fi41)
                                                                                                                                }
                                                                                                                                switch fi36 {
                                                                                                                                case 957:
                                                                                                                                    if (-128 > -37) {
                                                                                                                                        v42 = v42
                                                                                                                                    } else {
                                                                                                                                        var v60 string = "\n\t"
                                                                                                                                        v42 = (v40 && v42)
                                                                                                                                        _ = v60
                                                                                                                                    }
                                                                                                                                case 4294967295:
                                                                                                                                    v54 += v54
                                                                                                                                default:
                                                                                                                                    fmt.Println(v1, v54, v24)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if v40 {
                                                                                                                                fmt.Println(v42, v24)
                                                                                                                                v61 := GenS4{F0: 42, F1: -1, F2: v2}
                                                                                                                                fmt.Println(v40)
                                                                                                                                _ = v61
                                                                                                                            }
                                                                                                                            if len(v25) > 0 {
                                                                                                                                v25[0] = 508
                                                                                                                            }
                                                                                                                            if ("foo bar" == "foo bar") {
                                                                                                                                var v62 int = fi36
                                                                                                                                var v63 float64 = v24
                                                                                                                                v43 = (^(v43))
                                                                                                                                _ = v62
                                                                                                                                _ = v63
                                                                                                                            } else {
                                                                                                                                if len(v25) >= 2 {
                                                                                                                                    v25 = v25[1:len(v25)]
                                                                                                                                }
                                                                                                                                fmt.Println(v24)
                                                                                                                                v43 += (v43 >> (((^(v43))) & 63))
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v64 := GenS5{F0: v42, F1: (fi36 * 4294967295)}
                                                                                                                        switch "hello" {
                                                                                                                        case "foo bar":
                                                                                                                            switch "a" {
                                                                                                                            case "hello":
                                                                                                                                v53 += (v53 ^ v53)
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            if v42 {
                                                                                                                                v65 := GenS5{F0: true, F1: fi41}
                                                                                                                                _ = v65
                                                                                                                            } else {
                                                                                                                                fmt.Println(fi36, v42)
                                                                                                                                var v66 int32 = 1
                                                                                                                                var v67 int64 = 491
                                                                                                                                _ = v66
                                                                                                                                _ = v67
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            fmt.Println(v24)
                                                                                                                        }
                                                                                                                        if ((^(v1)) < uint16(v54)) {
                                                                                                                            v68 := GenS6{F0: -432, F1: 2147483647}
                                                                                                                            if v40 {
                                                                                                                                if (-1 != (v2 & v2)) {
                                                                                                                                    switch 0 {
                                                                                                                                    case 572:
                                                                                                                                        var v69 float64 = float64(0.0)
                                                                                                                                        _ = v69
                                                                                                                                    default:
                                                                                                                                        if (!v40) {
                                                                                                                                            switch fi41 {
                                                                                                                                            case 943:
                                                                                                                                                var v70 uint = (^(v43))
                                                                                                                                                _ = v70
                                                                                                                                            case 1:
                                                                                                                                                var v71 uint64 = 418
                                                                                                                                                _ = v71
                                                                                                                                            }
                                                                                                                                            switch fi41 {
                                                                                                                                            case 128:
                                                                                                                                                if (!(198 < 42)) {
                                                                                                                                                    v54 -= 1
                                                                                                                                                    if v40 {
                                                                                                                                                        switch "" {
                                                                                                                                                        case "":
                                                                                                                                                            if v40 {
                                                                                                                                                                _ = len(v3)
                                                                                                                                                                if v42 {
                                                                                                                                                                    var v72 int64 = v2
                                                                                                                                                                    _ = v72
                                                                                                                                                                } else {
                                                                                                                                                                    if len(v25) >= 2 {
                                                                                                                                                                        v25 = v25[1:len(v25)]
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                var v73 float32 = 1.5
                                                                                                                                                                _ = v73
                                                                                                                                                            } else {
                                                                                                                                                                fmt.Println(fi36)
                                                                                                                                                                var v74 string = ""
                                                                                                                                                                fmt.Println(v74)
                                                                                                                                                                _ = v74
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        v75 := GenS7{F0: v42, F1: "a"}
                                                                                                                                                        var v76 float64 = v24
                                                                                                                                                        _ = v75
                                                                                                                                                        _ = v76
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    var v77 float32 = 0.1
                                                                                                                                                    _ = len(v25)
                                                                                                                                                    _ = v77
                                                                                                                                                }
                                                                                                                                            case 32767:
                                                                                                                                                if (genFn0(fi41, (fi41 + 1)) != (v1 | (v1 % ((v1) | 1)))) {
                                                                                                                                                    _ = len(v25)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v78 uint32 = (v54 ^ (^(v54)))
                                                                                                                                            _ = v78
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v42 = (!false)
                                                                                                                                    v79 := GenS8{F0: 812, F1: "\n\t", F2: 21}
                                                                                                                                    _ = v79
                                                                                                                                } else {
                                                                                                                                    v25 = append(v25, (fi36 | fi41))
                                                                                                                                }
                                                                                                                                var v80 uint32 = (v54 >> ((uint(v54)) & 63))
                                                                                                                                _ = v80
                                                                                                                            } else {
                                                                                                                                if len(v25) > 0 {
                                                                                                                                    v25[0] = (fi36 ^ fi41)
                                                                                                                                }
                                                                                                                                v2 -= v2
                                                                                                                                _ = len(v25)
                                                                                                                            }
                                                                                                                            fmt.Println(v54)
                                                                                                                            var v81 float32 = 3.14159
                                                                                                                            v40 = v40
                                                                                                                            _ = v68
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        _ = v64
                                                                                                                    }
                                                                                                                    _ = v57
                                                                                                                }
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    if (!(!false)) {
                                                                                                                        if (!((v40 && (985 < v43)) || (-2.25 <= 1e10))) {
                                                                                                                            if ((!v40) && v40) {
                                                                                                                                var v82 bool = v42
                                                                                                                                v43 = (v43 << ((v43) & 63))
                                                                                                                                _ = v82
                                                                                                                            } else {
                                                                                                                                if (!(!v40)) {
                                                                                                                                    v83 := GenS9{F0: (128 > 0), F1: (fi41 - 0), F2: 1000}
                                                                                                                                    _ = v83
                                                                                                                                } else {
                                                                                                                                    if v42 {
                                                                                                                                        fmt.Println(v2, v42, fi36)
                                                                                                                                    } else {
                                                                                                                                        v54 -= 2147483648
                                                                                                                                    }
                                                                                                                                    switch 128 {
                                                                                                                                    case 42:
                                                                                                                                        if v40 {
                                                                                                                                            v42 = (-58 != -128)
                                                                                                                                        } else {
                                                                                                                                            _ = len(v25)
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v84 int8 = 1
                                                                                                                                        var v85 int64 = v2
                                                                                                                                        v2 -= (v85 >> ((0) & 63))
                                                                                                                                        v85 += (v2 << ((18446744073709551615) & 63))
                                                                                                                                        var v86 int64 = -9223372036854775808
                                                                                                                                        if v42 {
                                                                                                                                            var v87 uint8 = 0
                                                                                                                                            _ = v87
                                                                                                                                        } else {
                                                                                                                                            var v88 bool = (32767 <= int16(v43))
                                                                                                                                            v54 -= v54
                                                                                                                                            v1 += uint16(v53)
                                                                                                                                            if true {
                                                                                                                                                var v89 bool = v42
                                                                                                                                                _ = v89
                                                                                                                                            }
                                                                                                                                            _ = v88
                                                                                                                                        }
                                                                                                                                        v1 += genFn0(fi41, (fi36 & (fi36 & 318)))
                                                                                                                                        _ = v84
                                                                                                                                        _ = v85
                                                                                                                                        _ = v86
                                                                                                                                    }
                                                                                                                                    v53 += (v53 | v53)
                                                                                                                                    fmt.Println(v40, v1, v53)
                                                                                                                                }
                                                                                                                                if len(v25) >= 2 {
                                                                                                                                    v25 = v25[1:len(v25)]
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v90 int16 = 1000
                                                                                                                            _ = v90
                                                                                                                        } else {
                                                                                                                            var v91 float32 = 0.0
                                                                                                                            _ = v91
                                                                                                                        }
                                                                                                                        var v92 uint64 = v53
                                                                                                                        switch -1 {
                                                                                                                        case -9223372036854775808:
                                                                                                                            var v93 int16 = 42
                                                                                                                            _ = v93
                                                                                                                        default:
                                                                                                                            v24 = v24
                                                                                                                        }
                                                                                                                        _ = v92
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    _ = len(v3)
                                                                                                                    fmt.Println(v42)
                                                                                                                }
                                                                                                                v2 = (v2 & 502)
                                                                                                                if false {
                                                                                                                    v43 = v43
                                                                                                                    fmt.Println(v43, v53, fi36)
                                                                                                                } else {
                                                                                                                    v94 := GenS9{F0: false, F1: fi36, F2: 776}
                                                                                                                    v95 := GenS10{F0: 467, F1: 255}
                                                                                                                    _ = v94
                                                                                                                    _ = v95
                                                                                                                }
                                                                                                            }
                                                                                                            v96 := GenS9{F0: (!true), F1: (^(fi41)), F2: 91}
                                                                                                            var v97 uint = (^(v43))
                                                                                                            _ = v96
                                                                                                            _ = v97
                                                                                                        } else {
                                                                                                            var v98 int = (fi36 >> ((v43) & 63))
                                                                                                            var v99 float32 = (1.5 * ((((0.1 + 1.5) - 3.14159) * 1e10) + 3.14159))
                                                                                                            _ = v98
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        v53 = (v53 - 955)
                                                                                                    }
                                                                                                    v3[fi36] = v54
                                                                                                }
                                                                                            case 1:
                                                                                                if v40 {
                                                                                                    if v40 {
                                                                                                        var v100 int8 = 38
                                                                                                        _ = v100
                                                                                                    } else {
                                                                                                        v101 := GenS9{F0: (((v43 + v43) != (v43 ^ v43)) && (false || (0.0 <= v24))), F1: int(v43), F2: -267}
                                                                                                        if (475 < 42) {
                                                                                                            var v102 uint = v43
                                                                                                            _ = v102
                                                                                                        } else {
                                                                                                            v103 := GenS3{F0: 1000}
                                                                                                            v104 := GenS7{F0: (-1 < 1), F1: "foo bar"}
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                        }
                                                                                                        _ = v101
                                                                                                    }
                                                                                                    fmt.Println(v54, v1)
                                                                                                    v105 := GenS11{F0: (v53 % (((^(v53))) | 1))}
                                                                                                    _ = v105
                                                                                                } else {
                                                                                                    v106 := GenS4{F0: v1, F1: -2147483648, F2: 803}
                                                                                                    var v107 string = "hello"
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                }
                                                                                            }
                                                                                            _ = v53
                                                                                            _ = v54
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v40, v1, fi36)
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v42 = true
                                                                            case "a":
                                                                                var v108 float64 = ((v24 + v24) / v24)
                                                                                _ = v108
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v40, v40)
                                                                            fmt.Println(v2)
                                                                            var v109 uint = v43
                                                                            fmt.Println(v1)
                                                                            _ = v109
                                                                        }
                                                                        v2 = (v2 * -890)
                                                                        v110 := GenS12{F0: (v2 << ((v43) & 63)), F1: 158}
                                                                        _ = v110
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                                if true {
                                                                    v111 := GenS5{F0: ((872 > 1000) && (v40 || v40)), F1: fi36}
                                                                    _ = v111
                                                                } else {
                                                                    if len(v25) >= 2 {
                                                                        v25 = v25[1:len(v25)]
                                                                    }
                                                                    var v112 uint32 = 417
                                                                    _ = v112
                                                                }
                                                                _ = fi41
                                                            }
                                                            fmt.Println(v2)
                                                            _ = v40
                                                        } else {
                                                            v25 = append(v25, fi36)
                                                        }
                                                    case -2147483648:
                                                        for fi113 := 0; fi113 < 13; fi113++ {
                                                            if (0 == 127) {
                                                                fmt.Println(v24)
                                                                fmt.Println(fi36, fi36, v24)
                                                            } else {
                                                                v114 := GenS4{F0: v1, F1: 634, F2: int64(fi113)}
                                                                _ = v114
                                                            }
                                                            _ = fi113
                                                        }
                                                    case -224:
                                                        var v115 float64 = (float64(float32(v24)) / 3.14159)
                                                        _ = v115
                                                    }
                                                }
                                                _ = fi36
                                            }
                                        case -428:
                                            v116 := GenS1{F0: 192, F1: false}
                                            _ = v116
                                        case -1:
                                            v117 := GenS13{F0: (!((71 == 0) || true)), F1: (v1 * (v1 << ((216) & 63)))}
                                            _ = v117
                                        }
                                        v1 = v1
                                    }
                                }
                            }
                        } else {
                            v24 -= (1.5 - float64(1e10))
                            v118 := GenS1{F0: -1, F1: false}
                            v1 += v1
                            _ = v118
                        }
                    }
                    var v119 float32 = -2.25
                    _ = v119
                } else {
                    fmt.Println(v24, v24, v1)
                    if false {
                        for fi120 := 0; fi120 < 14; fi120++ {
                            var v121 string = "foo bar"
                            _ = fi120
                            _ = v121
                        }
                    }
                    if (true && (true && (false || (!false)))) {
                        var v122 int32 = -295
                        _ = v122
                    }
                }
                for fi123 := 0; fi123 < 4; fi123++ {
                    switch int(v2) {
                    case 32767:
                        for fi124 := 0; fi124 < 16; fi124++ {
                            if false {
                                v125 := v3[(^(fi124))]
                                _ = v125
                            } else {
                                var v126 float64 = ((v24 + v24) - v24)
                                _ = v126
                            }
                            _ = fi124
                        }
                    case -696:
                        var v127 uint64 = 18446744073709551615
                        _ = v127
                    case 9223372036854775807:
                        fmt.Println(v2)
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v24 -= v24
                        case "foo bar":
                            if true {
                                if true {
                                    v128 := GenS9{F0: ((1e10 * 3.14159) > 0.0), F1: fi123, F2: 32767}
                                    _ = v128
                                } else {
                                    fmt.Println(fi123, v1, v24)
                                    v129 := GenS12{F0: v2, F1: 2}
                                    v1 -= (v1 & v1)
                                    _ = v129
                                }
                                v1 += (^(v1))
                            }
                        }
                    }
                    _ = fi123
                }
                for fi130 := 0; fi130 < 12; fi130++ {
                    _ = len(v25)
                    _ = fi130
                }
            }
            fmt.Println(v1, v24, v24)
            v2 -= v2
        }
    }
    if true {
        v2 = (v2 / ((v2) | 1))
        v24 += v24
        switch 0 {
        case 636:
            if (false || true) {
                var v131 int = 2147483647
                fmt.Println(v131)
                _ = v131
            } else {
                _ = len(v25)
            }
        }
    } else {
        for fi132 := 0; fi132 < 16; fi132++ {
            var v133 uint32 = 4294967295
            var v134 string = ("foo bar" + "hello")
            _ = fi132
            _ = v133
            _ = v134
        }
        v24 = 1e10
        v135 := GenS4{F0: genFn0(0, 42), F1: 42, F2: v2}
        _ = v135
    }
    switch 4294967295 {
    case 2147483647:
        if (!(false || true)) {
            var v136 uint8 = 255
            for fi137 := 0; fi137 < 17; fi137++ {
                v1 = v1
                v1 += v1
                _ = fi137
            }
            _ = v136
        }
    }
    fmt.Println(v2)
    fmt.Println(v1, v1, v1)
    v138 := GenS12{F0: (v2 << ((4294967295) & 63)), F1: 119}
    if (!true) {
        v139 := GenS10{F0: v2, F1: 128}
        _ = v139
    } else {
        v140 := GenS14{F0: 42, F1: 42, F2: 142}
        _ = v140
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v138
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
