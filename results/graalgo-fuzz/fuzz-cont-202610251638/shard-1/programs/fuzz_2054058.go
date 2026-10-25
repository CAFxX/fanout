package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int16; F1 uint; F2 int8 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint64; F1 int32 }
type GenS14 struct { F0 uint64; F1 int16 }
type GenS15 struct { F0 int64; F1 string }
type GenS16 struct { F0 int32; F1 int }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint32; F1 bool; F2 bool }
type GenS19 struct { F0 uint16; F1 uint64; F2 uint }
func genFn0(p0 int64, p1 uint32) uint8 {
    v6 := []int{-519}
    var v7 bool = true
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return uint8(p1)
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    var v4 bool = (true || true)
    var v5 int32 = -2147483648
    v8 := v2["foo bar"]
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float64 = (0.0 - ((((0.0 * -2.25) / float64((0.1 * (1e10 / float32(-2.25))))) / 1.0) - (float64(3.14159) / 3.14159)))
        if (fi9 == fi9) {
            fmt.Println(fi9, v10)
            if v8 {
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 uint32 = 42
                    _ = fi11
                    _ = v12
                }
                fmt.Println(fi9, v8)
            }
            var v13 float64 = 1.5
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    var v14 int8 = 0
    switch 9223372036854775807 {
    case 2147483647:
        v4 = v8
    default:
        var v15 uint16 = 255
        _ = v15
    }
    v16 := GenS1{F0: 32768, F1: (1e10 != 3.14159), F2: v14}
    var v17 uint = 2
    v8 = true
    v18 := GenS2{F0: -1, F1: 42, F2: -817}
    var v19 uint64 = 344
    v4 = (32767 < 264)
    v20 := GenS1{F0: -2147483648, F1: (4294967295 == 1), F2: v14}
    fmt.Println(v19)
    v5 = v5
    v3 += genFn0(42, uint32(v17))
    if v4 {
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int = fi21
                switch (fi22 | fi21) {
                case 9223372036854775807:
                    fmt.Println(v8, v5, fi21)
                    if v4 {
                        v5 += (v5 << ((212) & 63))
                        if (!true) {
                            v23 += int(v5)
                            if (fi22 >= (^(v23))) {
                                v24 := GenS2{F0: 0, F1: 335, F2: (v5 ^ (v5 + int32(v19)))}
                                if true {
                                    v25 := GenS2{F0: 42, F1: 9223372036854775807, F2: -1}
                                    _ = v25
                                }
                                _ = v24
                            }
                            if v4 {
                                var v26 string = (("foo bar" + "foo bar") + (("a" + ("\n\t" + "foo bar")) + ""))
                                _ = v26
                            } else {
                                fmt.Println(fi22)
                            }
                        } else {
                            if false {
                                v27 := GenS3{F0: v3}
                                v4 = false
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v27
                            }
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                        }
                        fmt.Println(v5)
                    } else {
                        fmt.Println(fi22, v4)
                    }
                default:
                    var v28 uint8 = (^(v3))
                    _ = v28
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    } else {
        v19 += 42
        switch "foo bar" {
        case "a":
            var v29 string = "hello"
            _ = v29
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch ("" + ("\n\t" + ("\n\t" + ""))) {
        case "a":
            _ = len(v1)
        case "\n\t":
            var v30 uint8 = (v3 ^ (^(v3)))
            _ = v30
        case "hello":
            for fi31 := 0; fi31 < 5; fi31++ {
                if (1000 != -888) {
                    if v8 {
                        switch ("hello" + (((((("foo bar" + ("hello" + (("foo bar" + (((("hello" + "a") + "hello") + ("" + "")) + ("a" + "a"))) + (("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + "")))) + "foo bar") + "hello") + "\n\t") + "foo bar") + "hello")) {
                        case "hello":
                            v5 = (v5 | -1)
                        case "\n\t":
                            delete(v2, "a")
                        case "a":
                            var v32 string = "hello"
                            _ = v32
                        }
                        if v8 {
                            if v8 {
                                switch fi31 {
                                case 0:
                                    if (v8 || (!v4)) {
                                        v33 := GenS3{F0: v3}
                                        _ = v33
                                    } else {
                                        fmt.Println(v8, v19, v17)
                                        v14 -= v14
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v34 int64 = 86
                                            _ = v34
                                        case "a":
                                            v17 += (v17 - v17)
                                        case "":
                                            for fi35 := 0; fi35 < 11; fi35++ {
                                                if (v8 || false) {
                                                    var v36 uint = 999
                                                    switch "\n\t" {
                                                    case "":
                                                        if v4 {
                                                            switch fi31 {
                                                            case 891:
                                                                var v37 int32 = v5
                                                                _ = v37
                                                            case 42:
                                                                if (fi31 == -863) {
                                                                    if (v4 && (v19 != v19)) {
                                                                        var v38 uint8 = (v3 ^ (v3 >> (((v17 % ((2) | 1))) & 63)))
                                                                        _ = len(v1)
                                                                        v36 -= 897
                                                                        v8 = (v14 != (v14 & int8(fi35)))
                                                                        switch fi35 {
                                                                        case -2147483648:
                                                                            var v39 string = "foo bar"
                                                                            _ = v39
                                                                        default:
                                                                            if ((42 == 207) || false) {
                                                                                var v40 int16 = 1
                                                                                _ = v40
                                                                            }
                                                                            var v41 int64 = 0
                                                                            if (149 <= genFn0(v41, 480)) {
                                                                                v14 += v14
                                                                                var v42 int = 939
                                                                                fmt.Println(v3, v17, v17)
                                                                                _ = v42
                                                                            }
                                                                            v43 := GenS4{F0: fi35}
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 2147483647
                                                                            }
                                                                            _ = v41
                                                                            _ = v43
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                } else {
                                                                    if v8 {
                                                                        fmt.Println(v19)
                                                                        if (v4 || (v4 || true)) {
                                                                            v17 += v17
                                                                            if false {
                                                                                switch ("hello" + "") {
                                                                                case "":
                                                                                    if true {
                                                                                        v14 = (v14 & v14)
                                                                                        var v44 int16 = 1000
                                                                                        if true {
                                                                                            switch fi35 {
                                                                                            case 4294967295:
                                                                                                fmt.Println(v19, v19, v4)
                                                                                            case -2147483648:
                                                                                                switch (fi31 | fi31) {
                                                                                                case 128:
                                                                                                    v36 = v17
                                                                                                case 32767:
                                                                                                    switch (fi35 << ((v17) & 63)) {
                                                                                                    case 127:
                                                                                                        var v45 int = (fi35 - 42)
                                                                                                        _ = v45
                                                                                                    default:
                                                                                                        if true {
                                                                                                            v46 := GenS5{F0: 194}
                                                                                                            _ = v46
                                                                                                        } else {
                                                                                                            var v47 uint16 = 801
                                                                                                            v48 := GenS3{F0: (v3 - genFn0(0, 2147483647))}
                                                                                                            _ = v47
                                                                                                            _ = v48
                                                                                                        }
                                                                                                    }
                                                                                                case 262:
                                                                                                    v49 := GenS6{F0: ((0.1 - 0.1) <= (((((float32(0.0) + (-2.25 / -2.25)) - 1e10) * -2.25) * 3.14159) + ((float32((3.14159 - ((0.0 * (1e10 / 1.0)) - 1e10))) * 0.0) * -2.25))), F1: v5}
                                                                                                    _ = v49
                                                                                                default:
                                                                                                    fmt.Println(v19, v8, v44)
                                                                                                    var v50 string = ("hello" + ("foo bar" + "foo bar"))
                                                                                                    _ = v50
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(fi35, v36)
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                fmt.Println(v4, fi35, v8)
                                                                                            default:
                                                                                                v44 -= v44
                                                                                            }
                                                                                            fmt.Println(v44, v44, v3)
                                                                                        } else {
                                                                                            fmt.Println(v17, v8, v14)
                                                                                        }
                                                                                        if v8 {
                                                                                            var v51 uint = v17
                                                                                            _ = len(v1)
                                                                                            if true {
                                                                                                var v52 uint16 = 0
                                                                                                v8 = (42 == v44)
                                                                                                _ = v52
                                                                                            }
                                                                                            _ = v51
                                                                                        } else {
                                                                                            if (0 <= int64(fi31)) {
                                                                                                var v53 int = fi35
                                                                                                if v8 {
                                                                                                    var v54 string = "hello"
                                                                                                    if v4 {
                                                                                                        var v55 float64 = 0.1
                                                                                                        fmt.Println(v55)
                                                                                                        v19 += v19
                                                                                                        _ = v55
                                                                                                    }
                                                                                                    var v56 uint32 = uint32(v17)
                                                                                                    _ = v54
                                                                                                    _ = v56
                                                                                                } else {
                                                                                                    fmt.Println(v3, v53, v36)
                                                                                                    var v57 int64 = -9223372036854775808
                                                                                                    _ = v57
                                                                                                }
                                                                                                _ = v53
                                                                                            } else {
                                                                                                var v58 uint8 = v3
                                                                                                _ = v58
                                                                                            }
                                                                                            if true {
                                                                                                v17 += 65535
                                                                                                var v59 uint = (v17 % ((uint(v3)) | 1))
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    v59 += v36
                                                                                                default:
                                                                                                    v17 = v36
                                                                                                }
                                                                                                var v60 bool = true
                                                                                                v61 := GenS7{F0: int64(v3), F1: 42}
                                                                                                v62 := v2[("a" + "")]
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            } else {
                                                                                                var v63 bool = v8
                                                                                                v5 -= 1
                                                                                                v44 = (v44 / (((v44 ^ 680)) | 1))
                                                                                                if (!v4) {
                                                                                                    v14 -= (-(v14))
                                                                                                }
                                                                                                switch ("foo bar" + ("foo bar" + (("\n\t" + (("foo bar" + "") + "a")) + ("\n\t" + ((("a" + "") + "") + "hello"))))) {
                                                                                                case "a":
                                                                                                    v4 = (!(2147483647 > 2))
                                                                                                case "hello":
                                                                                                    switch int(v17) {
                                                                                                    case 2147483647:
                                                                                                        v3 += v3
                                                                                                    case 455:
                                                                                                        switch 146 {
                                                                                                        case -2147483648:
                                                                                                            if (139 > 42) {
                                                                                                                v14 -= (v14 % (((^(v14))) | 1))
                                                                                                            } else {
                                                                                                                v64 := GenS1{F0: (fi31 >> (((^(v17))) & 63)), F1: false, F2: (v14 / ((v14) | 1))}
                                                                                                                _ = v64
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v63
                                                                                            }
                                                                                        }
                                                                                        v65 := GenS1{F0: (fi31 / ((2147483647) | 1)), F1: (v4 || v4), F2: (v14 * v14)}
                                                                                        switch 128 {
                                                                                        case -1:
                                                                                            v2[(("a" + "hello") + "a")] = (!v8)
                                                                                        default:
                                                                                            _ = len(v1)
                                                                                        }
                                                                                        v66 := GenS8{F0: v44, F1: 549, F2: (v14 * 0)}
                                                                                        _ = v44
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    } else {
                                                                                        v4 = true
                                                                                    }
                                                                                case "\n\t":
                                                                                    fmt.Println(v4, v4)
                                                                                case "a":
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        if false {
                                                                                            _ = len(v1)
                                                                                        } else {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            fmt.Println(v19, v14, fi35)
                                                                                        }
                                                                                    case "a":
                                                                                        if true {
                                                                                            var v67 uint16 = 65535
                                                                                            _ = v67
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if ((3.14159 + 0.1) == 0.1) {
                                                                                            v5 -= (v5 - v5)
                                                                                            v4 = (1 >= 2)
                                                                                            if (!(!(v3 > v3))) {
                                                                                                if false {
                                                                                                    fmt.Println(v17, v8)
                                                                                                    if false {
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        var v68 string = ""
                                                                                                        v14 += 42
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    _ = len(v1)
                                                                                                } else {
                                                                                                    _ = len(v2)
                                                                                                }
                                                                                                var v69 uint64 = v19
                                                                                                _ = v69
                                                                                            } else {
                                                                                                v70 := GenS6{F0: false, F1: v5}
                                                                                                _ = v70
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v71 int64 = -423
                                                                                    var v72 uint = 518
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                }
                                                                                var v73 int64 = -440
                                                                                v19 = (v19 >> ((65535) & 63))
                                                                                _ = v73
                                                                            } else {
                                                                                v4 = (!v8)
                                                                                if v4 {
                                                                                    v74 := GenS4{F0: fi31}
                                                                                    var v75 string = (("" + "foo bar") + ("" + ("" + "\n\t")))
                                                                                    v19 -= v19
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                } else {
                                                                                    v3 -= (v3 ^ v3)
                                                                                    v36 = (v17 << ((v17) & 63))
                                                                                    v76 := GenS9{F0: 42, F1: v14}
                                                                                    _ = v76
                                                                                }
                                                                                var v77 uint64 = 245
                                                                                _ = v77
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v17, v8, v4)
                                                                        }
                                                                    }
                                                                    if v8 {
                                                                        if v4 {
                                                                            if v8 {
                                                                                var v78 int64 = 9223372036854775807
                                                                                fmt.Println(v19)
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            var v79 uint8 = (^(v3))
                                                                            if (!(false || (v8 && ((v5 <= 900) && false)))) {
                                                                                var v80 float32 = float32(-2.25)
                                                                                if v4 {
                                                                                    switch fi31 {
                                                                                    case -1:
                                                                                        var v81 int32 = v5
                                                                                        _ = v81
                                                                                    case 0:
                                                                                        fmt.Println(v5)
                                                                                    default:
                                                                                        fmt.Println(v17, v17, v3)
                                                                                    }
                                                                                }
                                                                                var v82 uint = (v17 << ((v36) & 63))
                                                                                _ = len(v1)
                                                                                var v83 uint = (v17 ^ 723)
                                                                                v79 += (^(v3))
                                                                                _ = v80
                                                                                _ = v82
                                                                                _ = v83
                                                                            } else {
                                                                                var v84 bool = (-863 != (v5 & (v5 * 42)))
                                                                                if ("hello" <= "\n\t") {
                                                                                    var v85 float32 = float32(float64(3.14159))
                                                                                    switch 1 {
                                                                                    case 32767:
                                                                                        if true {
                                                                                            var v86 uint8 = 128
                                                                                            switch (fi35 ^ (fi35 - fi31)) {
                                                                                            case -445:
                                                                                                if (100 >= (v14 & v14)) {
                                                                                                    if (270 != uint32(v19)) {
                                                                                                        v1 = append(v1, 2)
                                                                                                        var v87 string = "a"
                                                                                                        delete(v2, v87)
                                                                                                        _ = v87
                                                                                                    }
                                                                                                } else {
                                                                                                    var v88 string = ""
                                                                                                    fmt.Println(v88)
                                                                                                    fmt.Println(fi31, v79)
                                                                                                    _ = v88
                                                                                                }
                                                                                            case -841:
                                                                                                v8 = v8
                                                                                            case -957:
                                                                                                var v89 bool = false
                                                                                                _ = v89
                                                                                            default:
                                                                                                var v90 float32 = float32(-2.25)
                                                                                                switch (-(fi35)) {
                                                                                                case 2147483647:
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                case 128:
                                                                                                    var v91 int32 = (v5 - (v5 / (((^(v5))) | 1)))
                                                                                                    _ = v91
                                                                                                default:
                                                                                                    v92 := GenS1{F0: -745, F1: (!(uint32(v19) >= 2147483648)), F2: v14}
                                                                                                    _ = v92
                                                                                                }
                                                                                                v1 = append(v1, 1)
                                                                                                _ = v90
                                                                                            }
                                                                                            v36 = v17
                                                                                            v93 := GenS2{F0: 0, F1: 1, F2: (v5 - -829)}
                                                                                            _ = v86
                                                                                            _ = v93
                                                                                        } else {
                                                                                            fmt.Println(v19)
                                                                                            var v94 uint8 = v79
                                                                                            v79 += (v3 / ((genFn0(750, 1)) | 1))
                                                                                            _ = v94
                                                                                        }
                                                                                    }
                                                                                    _ = v85
                                                                                } else {
                                                                                    fmt.Println(v8)
                                                                                }
                                                                                v84 = (v5 == v5)
                                                                                _ = v84
                                                                            }
                                                                            v95 := GenS9{F0: 42, F1: -44}
                                                                            var v96 int64 = -385
                                                                            _ = v79
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        var v97 uint32 = 452
                                                                        var v98 float32 = 3.14159
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                }
                                                            default:
                                                                v99 := GenS1{F0: fi31, F1: ((uint32(v17) <= 258) || (1.5 == float32((1e10 - 3.14159)))), F2: (v14 | (v14 | (v14 & (-(v14)))))}
                                                                switch fi31 {
                                                                case 2147483647:
                                                                    if ((v3 & v3) >= v3) {
                                                                        var v100 uint64 = 305
                                                                        v17 += (v17 >> ((v17) & 63))
                                                                        v36 = v36
                                                                        _ = v100
                                                                    }
                                                                case -2147483648:
                                                                    v101 := GenS3{F0: v3}
                                                                    _ = v101
                                                                default:
                                                                    if (v8 || false) {
                                                                        v102 := GenS10{F0: (v3 ^ genFn0(-9223372036854775808, 2147483648))}
                                                                        v103 := GenS5{F0: v19}
                                                                        v1 = append(v1, 65535)
                                                                        _ = v102
                                                                        _ = v103
                                                                    } else {
                                                                        var v104 float32 = ((0.0 / 1.5) - 0.1)
                                                                        var v105 uint8 = genFn0(-42, 55)
                                                                        _ = v104
                                                                        _ = v105
                                                                    }
                                                                }
                                                                _ = v99
                                                            }
                                                            v2[("foo bar" + (("foo bar" + "") + "foo bar"))] = (!v4)
                                                        }
                                                    case "foo bar":
                                                        var v106 bool = v8
                                                        _ = v106
                                                    case "\n\t":
                                                        fmt.Println(v14, v36)
                                                    }
                                                    fmt.Println(v3)
                                                    fmt.Println(v4)
                                                    var v107 int32 = v5
                                                    _ = v107
                                                    _ = v36
                                                } else {
                                                    if ((v3 | 198) >= genFn0(42, 65535)) {
                                                        var v108 float32 = ((-2.25 + (0.1 - 0.0)) + (float32(1e10) / (0.1 / 1.0)))
                                                        _ = v108
                                                    } else {
                                                        var v109 uint64 = v19
                                                        v3 = uint8(v19)
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v110 float64 = 1.5
                                                            _ = v110
                                                        }
                                                        _ = v109
                                                    }
                                                    v17 -= (v17 - uint(v3))
                                                }
                                                fmt.Println(v17, v17)
                                                if v8 {
                                                    var v111 int64 = 9223372036854775807
                                                    v111 -= v111
                                                    if (v4 && false) {
                                                        if false {
                                                            var v112 int32 = (v5 / ((-2147483648) | 1))
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch -2147483648 {
                                                                case 9223372036854775807:
                                                                    v113 := GenS8{F0: 1, F1: (^(v17)), F2: v14}
                                                                    _ = v113
                                                                default:
                                                                    switch (fi35 | fi35) {
                                                                    case 128:
                                                                        _ = len(v2)
                                                                    default:
                                                                        var v114 int64 = -9223372036854775808
                                                                        _ = v114
                                                                    }
                                                                }
                                                            case "":
                                                                v3 += (v3 * v3)
                                                            case "hello":
                                                                if v8 {
                                                                    var v115 float32 = ((3.14159 - (3.14159 / 1e10)) * 0.1)
                                                                    _ = v115
                                                                } else {
                                                                    var v116 uint32 = 2147483647
                                                                    v111 = 9223372036854775807
                                                                    fmt.Println(v14, v8, v19)
                                                                    _ = v116
                                                                }
                                                            default:
                                                                _ = len(v2)
                                                            }
                                                            _ = v112
                                                        } else {
                                                            var v117 int64 = -521
                                                            v4 = ((!false) || true)
                                                            _ = v117
                                                        }
                                                    }
                                                    _ = v111
                                                } else {
                                                    v118 := GenS4{F0: fi35}
                                                    _ = v118
                                                }
                                                _ = fi35
                                            }
                                        default:
                                            var v119 float64 = 3.14159
                                            _ = v119
                                        }
                                    }
                                case 547:
                                    v120 := GenS9{F0: -32768, F1: (v14 << ((4294967295) & 63))}
                                    _ = v120
                                default:
                                    var v121 int64 = 42
                                    _ = v121
                                }
                                for fi122 := 0; fi122 < 14; fi122++ {
                                    if (((v5 << ((v17) & 63)) > v5) || true) {
                                        switch (fi122 % ((fi31) | 1)) {
                                        case 485:
                                            switch int(v5) {
                                            case -9223372036854775808:
                                                var v123 float64 = float64(3.14159)
                                                _ = v123
                                            case 42:
                                                switch (("\n\t" + "\n\t") + "a") {
                                                case "foo bar":
                                                    v124 := GenS6{F0: v4, F1: 1}
                                                    _ = v124
                                                case "hello":
                                                    var v125 uint = v17
                                                    _ = v125
                                                }
                                            default:
                                                fmt.Println(fi31, fi122)
                                                var v126 float64 = float64(0.1)
                                                switch "" {
                                                case "foo bar":
                                                    v127 := GenS2{F0: 263, F1: 1, F2: v5}
                                                    _ = v127
                                                case "hello":
                                                    delete(v2, "")
                                                default:
                                                    v19 -= 42
                                                }
                                                _ = v126
                                            }
                                        case 128:
                                            _ = len(v2)
                                        }
                                        if v4 {
                                            var v128 float64 = 0.0
                                            _ = v128
                                        } else {
                                            fmt.Println(v8)
                                        }
                                        var v129 float64 = 1e10
                                        _ = v129
                                    }
                                    _ = fi122
                                }
                            } else {
                                if (false || true) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v19, fi31, v14)
                                    case "\n\t":
                                        if v8 {
                                            switch "a" {
                                            case "":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                v14 = (v14 | v14)
                                            }
                                        } else {
                                            switch ("\n\t" + "hello") {
                                            case "\n\t":
                                                switch (fi31 * (fi31 >> (((v17 ^ 65535)) & 63))) {
                                                case 622:
                                                    v3 = 23
                                                case 0:
                                                    for fi130 := 0; fi130 < 13; fi130++ {
                                                        v131 := GenS11{F0: v5, F1: (v3 ^ 1), F2: v5}
                                                        _ = fi130
                                                        _ = v131
                                                    }
                                                case 449:
                                                    fmt.Println(v19, v17, v8)
                                                }
                                            case "a":
                                                if ((v5 >> (((v17 & v17)) & 63)) > (v5 << ((0) & 63))) {
                                                    var v132 int32 = -2147483648
                                                    v19 += 42
                                                    _ = v132
                                                }
                                            }
                                        }
                                    default:
                                        var v133 int = fi31
                                        _ = v133
                                    }
                                    switch 32767 {
                                    case -9223372036854775808:
                                        for fi134 := 0; fi134 < 19; fi134++ {
                                            v17 += (v17 * 18446744073709551615)
                                            if (v8 && true) {
                                                v135 := GenS8{F0: -1, F1: (^(v17)), F2: (v14 | (v14 >> ((65535) & 63)))}
                                                _ = v135
                                            } else {
                                                var v136 uint64 = v19
                                                v137 := v2["\n\t"]
                                                if v8 {
                                                    var v138 bool = v4
                                                    _ = v138
                                                } else {
                                                    var v139 uint64 = 36
                                                    v140 := GenS12{F0: 42, F1: (v17 % (((v17 | 255)) | 1)), F2: 18446744073709551615}
                                                    if false {
                                                        var v141 uint16 = 1
                                                        v8 = v4
                                                        _ = v141
                                                    }
                                                    _ = v139
                                                    _ = v140
                                                }
                                                _ = v136
                                                _ = v137
                                            }
                                            v14 += (v14 - 100)
                                            _ = fi134
                                        }
                                    case 1:
                                        for fi142 := 0; fi142 < 12; fi142++ {
                                            fmt.Println(v5, v8)
                                            _ = fi142
                                        }
                                    }
                                    v8 = true
                                } else {
                                    if ("a" != ("" + ((("a" + ("a" + (((("foo bar" + ("" + "a")) + "foo bar") + "\n\t") + ("\n\t" + ("" + "hello"))))) + "a") + ("foo bar" + "a")))) {
                                        v3 = 42
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v8 = (false || (v8 || (("foo bar" + "") < ("hello" + "foo bar"))))
                                    }
                                    if v8 {
                                        switch "hello" {
                                        case "foo bar":
                                            v143 := GenS11{F0: (v5 >> ((v17) & 63)), F1: v3, F2: v5}
                                            _ = v143
                                        case "a":
                                            if v4 {
                                                var v144 uint8 = v3
                                                v8 = v8
                                                _ = v144
                                            } else {
                                                var v145 int64 = -82
                                                _ = v145
                                            }
                                        }
                                        switch fi31 {
                                        case -1000:
                                            delete(v2, ("foo bar" + ("hello" + (("hello" + "") + "foo bar"))))
                                        case 392:
                                            v146 := GenS3{F0: (^(v3))}
                                            _ = v146
                                        case 42:
                                            var v147 bool = ((^(v3)) < uint8(v17))
                                            _ = v147
                                        default:
                                            if v4 {
                                                v3 -= (v3 & genFn0(642, 2147483648))
                                                v148 := GenS13{F0: (v19 | v19), F1: v5}
                                                _ = v148
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = 0
                                                }
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v8)
                                                case "\n\t":
                                                    v149 := GenS14{F0: uint64(v17), F1: -780}
                                                    _ = v149
                                                case "foo bar":
                                                    v150 := v2["hello"]
                                                    _ = v150
                                                }
                                            }
                                        }
                                    } else {
                                        v1 = append(v1, 714)
                                        if (!(1000 >= 42)) {
                                            v17 -= v17
                                        }
                                    }
                                }
                            }
                            var v151 float32 = ((3.14159 / 1e10) - (-2.25 * 1.5))
                            _ = v151
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v4 = (!(0.1 > float64((float32(0.1) * 1e10))))
                    }
                }
                delete(v2, ((("\n\t" + "\n\t") + "") + "hello"))
                _ = fi31
            }
        default:
            v14 += -128
        }
    }
    if v4 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v152 uint32 = 1
        var v153 uint = uint(v19)
        var v154 uint = v17
        _ = v152
        _ = v153
        _ = v154
    }
    if true {
        v19 = 616
        fmt.Println(v14, v4, v19)
        if v8 {
            fmt.Println(v8, v17)
            if ((v3 - uint8(v19)) < v3) {
                for fi155 := 0; fi155 < 9; fi155++ {
                    v14 += (v14 & -71)
                    _ = fi155
                }
                var v156 uint64 = v19
                if v8 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v157 uint = (^(v17))
                    if (42 != 185) {
                        if true {
                            if v8 {
                                fmt.Println(v5)
                            }
                            var v158 string = "a"
                            _ = v158
                        } else {
                            if false {
                                v157 -= (^(v17))
                            }
                            var v159 int64 = 9223372036854775807
                            _ = v159
                        }
                    } else {
                        if (42 >= 4294967295) {
                            fmt.Println(v156, v156, v14)
                            if (false || ("a" <= "\n\t")) {
                                if (!true) {
                                    var v160 int16 = -32768
                                    for fi161 := 0; fi161 < 2; fi161++ {
                                        if false {
                                            for fi162 := 0; fi162 < 5; fi162++ {
                                                delete(v2, "\n\t")
                                                v8 = v4
                                                _ = fi162
                                            }
                                        } else {
                                            switch "" {
                                            case "":
                                                var v163 uint = (v17 - 2)
                                                _ = v163
                                            case "hello":
                                                var v164 uint64 = v19
                                                _ = v164
                                            default:
                                                fmt.Println(v4, v160, fi161)
                                            }
                                            if (2 > 255) {
                                                for fi165 := 0; fi165 < 12; fi165++ {
                                                    var v166 uint16 = 2
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = fi165
                                                    _ = v166
                                                }
                                                for fi167 := 0; fi167 < 12; fi167++ {
                                                    var v168 uint16 = 42
                                                    _ = fi167
                                                    _ = v168
                                                }
                                            } else {
                                                var v169 int16 = v160
                                                var v170 bool = (-9223372036854775808 > -270)
                                                _ = v169
                                                _ = v170
                                            }
                                        }
                                        if true {
                                            switch ("hello" + "a") {
                                            case "a":
                                                v157 -= 18446744073709551615
                                            default:
                                                for fi171 := 0; fi171 < 6; fi171++ {
                                                    var v172 string = "\n\t"
                                                    v173 := GenS15{F0: -1, F1: ((("a" + (v172 + "a")) + v172) + "a")}
                                                    fmt.Println(v160, v3)
                                                    _ = fi171
                                                    _ = v172
                                                    _ = v173
                                                }
                                            }
                                            if (v3 != genFn0(0, 2147483647)) {
                                                fmt.Println(v5)
                                            }
                                        }
                                        for fi174 := 0; fi174 < 12; fi174++ {
                                            switch fi174 {
                                            case 1:
                                                var v175 float64 = 0.1
                                                var v176 bool = v8
                                                v177 := GenS15{F0: 42, F1: "\n\t"}
                                                _ = v175
                                                _ = v176
                                                _ = v177
                                            }
                                            var v178 uint32 = 2147483648
                                            v5 = -1
                                            _ = fi174
                                            _ = v178
                                        }
                                        _ = fi161
                                    }
                                    v179 := GenS16{F0: (v5 ^ v5), F1: -9223372036854775808}
                                    _ = v160
                                    _ = v179
                                }
                                var v180 uint64 = v19
                                _ = v180
                            } else {
                                v1 = append(v1, 2)
                            }
                        } else {
                            var v181 bool = v4
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v181
                        }
                        switch -939 {
                        case 2147483647:
                            v8 = (true || true)
                        case -592:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            v4 = (((((("\n\t" + ("hello" + "hello")) + "\n\t") + (("\n\t" + "foo bar") + "a")) + "a") + "hello") == "a")
                        }
                        for fi182 := 0; fi182 < 20; fi182++ {
                            for fi183 := 0; fi183 < 6; fi183++ {
                                if (fi182 == fi182) {
                                    v184 := GenS17{F0: ("\n\t" + ("a" + ("" + ("a" + ("a" + "a")))))}
                                    v8 = v8
                                    if false {
                                        var v185 bool = v4
                                        fmt.Println(v17, v156, v185)
                                        if v8 {
                                            var v186 bool = false
                                            var v187 bool = v186
                                            var v188 uint16 = 255
                                            if (!true) {
                                                var v189 string = ("\n\t" + "")
                                                var v190 bool = true
                                                _ = v189
                                                _ = v190
                                            } else {
                                                var v191 uint32 = 0
                                                var v192 uint16 = 968
                                                var v193 uint32 = v191
                                                _ = v191
                                                _ = v192
                                                _ = v193
                                            }
                                            _ = v186
                                            _ = v187
                                            _ = v188
                                        }
                                        _ = v185
                                    }
                                    _ = v184
                                } else {
                                    if ((!v4) || (!v4)) {
                                        if (-9223372036854775808 >= fi182) {
                                            var v194 uint32 = 65535
                                            v14 = 100
                                            if (0 >= (v156 & v19)) {
                                                if len(v1) > 0 {
                                                    v1[0] = (v194 >> ((77) & 63))
                                                }
                                                fmt.Println(v5)
                                                var v195 bool = (-207 < -1)
                                                _ = v195
                                            }
                                            _ = v194
                                        }
                                    } else {
                                        var v196 uint64 = 445
                                        var v197 uint32 = 2147483648
                                        _ = v196
                                        _ = v197
                                    }
                                    v19 = (v156 * v156)
                                    delete(v2, "foo bar")
                                }
                                var v198 int8 = (-(v14))
                                _ = fi183
                                _ = v198
                            }
                            v199 := GenS2{F0: 32767, F1: -719, F2: 0}
                            if ((!(-2.25 > 0.0)) || (-1 > int16(v3))) {
                                var v200 uint8 = (^(v3))
                                fmt.Println(v3, v8)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v200
                            } else {
                                v4 = (!false)
                            }
                            _ = fi182
                            _ = v199
                        }
                    }
                    switch 128 {
                    case -9223372036854775808:
                        if (!(v4 || v4)) {
                            fmt.Println(v14)
                        }
                    case 32768:
                        var v201 float32 = 0.1
                        _ = v201
                    case 32767:
                        if v8 {
                            for fi202 := 0; fi202 < 8; fi202++ {
                                v157 = v157
                                v19 += (v156 + (^(v19)))
                                for fi203 := 0; fi203 < 10; fi203++ {
                                    v14 = v14
                                    _ = fi203
                                }
                                _ = fi202
                            }
                            v204 := GenS18{F0: 4294967295, F1: v4, F2: (2 == 255)}
                            v205 := GenS8{F0: 1, F1: (v157 % ((v17) | 1)), F2: 100}
                            _ = v204
                            _ = v205
                        } else {
                            if v4 {
                                v19 = (^(v19))
                                for fi206 := 0; fi206 < 4; fi206++ {
                                    var v207 float32 = (1.5 + 1.5)
                                    v14 = (v14 >> ((v17) & 63))
                                    _ = fi206
                                    _ = v207
                                }
                            } else {
                                v157 += (v157 + 42)
                            }
                            for fi208 := 0; fi208 < 20; fi208++ {
                                var v209 int32 = 2147483647
                                var v210 int64 = -509
                                fmt.Println(v19, v8)
                                _ = fi208
                                _ = v209
                                _ = v210
                            }
                            if (2147483647 >= v5) {
                                switch "" {
                                case "hello":
                                    v4 = v8
                                case "foo bar":
                                    var v211 float64 = (0.0 * float64((1.5 / 0.1)))
                                    _ = v211
                                }
                                v19 = (^(v156))
                                v14 = v14
                            } else {
                                for fi212 := 0; fi212 < 3; fi212++ {
                                    var v213 int8 = v14
                                    _ = fi212
                                    _ = v213
                                }
                            }
                        }
                    default:
                        v214 := GenS19{F0: 42, F1: (v156 << ((v157) & 63)), F2: (v157 - v157)}
                        _ = v214
                    }
                    _ = v157
                }
                _ = v156
            }
            fmt.Println(v5)
        } else {
            delete(v2, "a")
            _ = len(v2)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
