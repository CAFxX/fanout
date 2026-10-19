package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint16; F2 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 int8; F2 int16 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint; F1 int8; F2 uint8 }
type GenS13 struct { F0 int64; F1 uint64; F2 int8 }
type GenS14 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 int32, p1 uint8) bool {
    var v5 int32 = (p0 & -312)
    p0 = v5
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    return true
}
func genFn1(p0 int16, p1 bool) bool {
    p0 = p0
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 = p0
        var v7 uint = 355
        var v8 uint32 = 405
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 0 {
    case 127:
        switch ("a" + "\n\t") {
        case "a":
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 float32 = 1.5
                _ = fi9
                _ = v10
            }
        case "\n\t":
            v11 := GenS1{F0: "a"}
            _ = v11
        case "foo bar":
            var v12 uint8 = 255
            _ = v12
        default:
            if (255 == 730) {
                v13 := GenS2{F0: p0}
                p1 = (float64(1e10) < (-2.25 + (-2.25 + 1e10)))
                _ = v13
            } else {
                if false {
                    fmt.Println(p1)
                    var v14 int = 408
                    p1 = (453 < 380)
                    _ = v14
                }
                switch 42 {
                case -739:
                    if p1 {
                        fmt.Println(p0, p0)
                        var v15 uint64 = 806
                        _ = v15
                    }
                case -531:
                    switch "a" {
                    case "\n\t":
                        fmt.Println(p0)
                    }
                case 127:
                    var v16 int16 = p0
                    _ = v16
                }
            }
        }
    }
    var v17 float32 = 1e10
    p1 = (false || p1)
    _ = p0
    _ = p1
    _ = v17
    return p1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = 42
    v3 := []string{}
    var v4 int64 = -1
    var v18 uint32 = 42
    var v19 float32 = ((3.14159 - 0.0) + 1e10)
    fmt.Println(v18, v1)
    var v20 bool = true
    var v21 uint = 967
    fmt.Println(v21)
    var v22 int16 = 42
    var v23 int16 = v22
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := GenS2{F0: (v22 >> ((v21) & 63))}
        var v26 string = ("hello" + "")
        var v27 uint = (^(v1))
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: "foo bar"}
    var v29 bool = (-846 > -1)
    var v30 int64 = v2
    var v31 uint64 = 1
    v32 := GenS3{F0: -1, F1: 127}
    v31 = v31
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 bool = v20
            if v29 {
                if v29 {
                    if v35 {
                        v36 := GenS4{F0: v31, F1: int32(v2)}
                        switch "foo bar" {
                        case "hello":
                            if v35 {
                                if v20 {
                                    fmt.Println(v31)
                                    if (v20 || ((!(!false)) && (!((0 == 65535) || (v31 == v31))))) {
                                        v3 = append(v3, "a")
                                        if (!(!v35)) {
                                            var v37 int32 = 2147483647
                                            _ = v37
                                        }
                                    }
                                    v38 := GenS2{F0: v23}
                                    _ = v38
                                }
                                var v39 uint32 = (v18 >> ((v21) & 63))
                                _ = v39
                            } else {
                                if (!(1e10 <= (v19 - (v19 * 0.1)))) {
                                    var v40 uint16 = 255
                                    if ((18446744073709551615 <= (v21 << ((42) & 63))) && (!(v22 > 0))) {
                                        fmt.Println(fi33)
                                        _ = len(v3)
                                    } else {
                                        _ = len(v3)
                                        _ = len(v3)
                                        v41 := GenS5{F0: 0}
                                        var v42 uint64 = (v31 ^ (^(v31)))
                                        if (1e10 <= 1e10) {
                                            if ((((fi33 | (fi33 * 521)) == (fi33 / (((fi34 % ((fi33) | 1))) | 1))) && (true && v35)) && (false && ((v30 << ((v1) & 63)) < int64(fi33)))) {
                                                var v43 int16 = 42
                                                _ = len(v3)
                                                _ = v43
                                            }
                                        }
                                        if (true || v29) {
                                            if ((!false) && (true || v20)) {
                                                var v44 int = fi33
                                                _ = v44
                                            }
                                            if v35 {
                                                var v45 uint = (v21 / ((65535) | 1))
                                                _ = v45
                                            }
                                            switch -408 {
                                            case 32768:
                                                var v46 int64 = v2
                                                _ = v46
                                            case 531:
                                                switch fi33 {
                                                case 343:
                                                    _ = len(v3)
                                                    var v47 int8 = -1
                                                    _ = v47
                                                case 42:
                                                    if len(v3) > 0 {
                                                        v3[0] = "hello"
                                                    }
                                                default:
                                                    var v48 int32 = -2147483648
                                                    _ = v48
                                                }
                                            case 128:
                                                var v49 float64 = float64(v19)
                                                _ = v49
                                            default:
                                                var v50 uint32 = v18
                                                _ = v50
                                            }
                                        } else {
                                            v51 := GenS1{F0: ""}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch (((("hello" + "hello") + "a") + ("foo bar" + ("a" + ("hello" + "foo bar")))) + (("hello" + "foo bar") + ("hello" + "foo bar"))) {
                                            case "":
                                                if v35 {
                                                    if (v40 == 90) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v52 := GenS4{F0: v31, F1: -2147483648}
                                                        var v53 int32 = 2147483647
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        if true {
                                                            var v54 uint16 = (v40 * 402)
                                                            _ = len(v3)
                                                            _ = v54
                                                        }
                                                        var v55 int16 = v23
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v56 uint16 = v40
                                                            _ = v56
                                                        case "a":
                                                            fmt.Println(v30)
                                                        case "foo bar":
                                                            var v57 int = fi33
                                                            _ = v57
                                                        default:
                                                            var v58 int8 = 0
                                                            _ = v58
                                                        }
                                                        v30 = (v4 << ((v1) & 63))
                                                        _ = v55
                                                    }
                                                    v2 -= (v4 * 689)
                                                    var v59 string = ("\n\t" + "hello")
                                                    _ = v59
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v35 = ((v18 | (v18 + (v18 >> ((v1) & 63)))) <= v18)
                                            default:
                                                v2 += v30
                                            }
                                            _ = v51
                                        }
                                        if v29 {
                                            switch (fi34 ^ fi33) {
                                            case 1:
                                                v20 = ((fi34 ^ (^(fi34))) > (^(fi34)))
                                            case 128:
                                                var v60 float64 = (-2.25 - float64((0.1 - (float32(-2.25) + -2.25))))
                                                _ = v60
                                            case -9223372036854775808:
                                                var v61 string = "foo bar"
                                                _ = v61
                                            }
                                            if v29 {
                                                switch 42 {
                                                case 128:
                                                    if true {
                                                        if v20 {
                                                            var v62 uint32 = 42
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v62
                                                        } else {
                                                            v3 = append(v3, (("a" + "") + (("foo bar" + "") + (("hello" + "hello") + ((("a" + "") + "hello") + "hello")))))
                                                            var v63 bool = v29
                                                            v64 := GenS6{F0: -483, F1: (((v40 != v40) && (!(!((v35 || v29) || (v20 || true))))) && false)}
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        v3 = append(v3, "foo bar")
                                                        v65 := GenS7{F0: int32(v40)}
                                                        _ = v65
                                                    } else {
                                                        if (-1 < -1) {
                                                            v3 = append(v3, "a")
                                                            v66 := GenS4{F0: (^(v31)), F1: -2147483648}
                                                            v67 := GenS8{F0: 127, F1: (v40 / ((uint16(v1)) | 1)), F2: 1}
                                                            v22 += (v23 + 1)
                                                            var v68 uint = (^(v21))
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            if v35 {
                                                                v69 := GenS4{F0: (v42 & (v42 << ((v1) & 63))), F1: 390}
                                                                _ = v69
                                                            }
                                                        }
                                                        v4 = 42
                                                    }
                                                }
                                                _ = len(v3)
                                            }
                                            v4 += (v30 - v4)
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                    v70 := GenS9{F0: v22}
                                    _ = v40
                                    _ = v70
                                } else {
                                    switch (fi33 ^ (-(fi34))) {
                                    case 42:
                                        fmt.Println(v2, v1)
                                    case 32767:
                                        v71 := GenS10{F0: v31, F1: 1, F2: 1}
                                        _ = v71
                                    default:
                                        if v29 {
                                            var v72 uint32 = 527
                                            v1 += v1
                                            _ = v72
                                        }
                                    }
                                    if false {
                                        if v20 {
                                            fmt.Println(v19, fi34, v19)
                                            v73 := GenS2{F0: int16(v21)}
                                            _ = v73
                                        }
                                    }
                                }
                            }
                        default:
                            var v74 uint8 = 127
                            _ = v74
                        }
                        switch ("foo bar" + "\n\t") {
                        case "":
                            if v20 {
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                        case "foo bar":
                            if (v21 > uint(v18)) {
                                if v35 {
                                    v30 -= -9223372036854775808
                                    var v75 int16 = v22
                                    var v76 uint = (v1 & 2)
                                    _ = v75
                                    _ = v76
                                }
                                fmt.Println(v21, v1)
                                var v77 bool = v35
                                _ = v77
                            } else {
                                v29 = v29
                            }
                        case "hello":
                            fmt.Println(v19, v31, v20)
                        default:
                            if (0 <= 255) {
                                v2 -= 9223372036854775807
                            }
                        }
                        if ((v18 % (((^(v18))) | 1)) != v18) {
                            if (v31 <= 2) {
                                v2 -= (v2 - 42)
                                if (!v20) {
                                    if false {
                                        fmt.Println(v23, v21, v4)
                                        fmt.Println(v4, fi33, v19)
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                    } else {
                                        switch ("\n\t" + (("\n\t" + "") + (("hello" + "hello") + ("" + "\n\t")))) {
                                        case "hello":
                                            var v78 uint32 = (v18 >> (((^(v21))) & 63))
                                            _ = v78
                                        case "\n\t":
                                            v2 = v30
                                        default:
                                            switch (("hello" + "\n\t") + ("hello" + "foo bar")) {
                                            case "a":
                                                v3 = append(v3, "\n\t")
                                            case "":
                                                if (((!(true || v29)) && v20) || v20) {
                                                    if ((1.5 + 1.5) <= float64(v19)) {
                                                        var v79 uint64 = v31
                                                        var v80 uint64 = 24
                                                        var v81 uint64 = v31
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        v22 += (v23 + v22)
                                                        if (uint(v31) < uint(v18)) {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    fmt.Println(v23, v22)
                                                    v31 = v31
                                                } else {
                                                    switch (fi33 % (((fi33 << ((65535) & 63))) | 1)) {
                                                    case 1:
                                                        var v82 uint32 = (^(v18))
                                                        _ = v82
                                                    case 0:
                                                        switch fi33 {
                                                        case -2147483648:
                                                            switch (("hello" + (("foo bar" + ("a" + (("a" + "foo bar") + "a"))) + ("a" + "a"))) + "a") {
                                                            case "\n\t":
                                                                _ = len(v3)
                                                            case "hello":
                                                                var v83 uint = 2
                                                                _ = v83
                                                            default:
                                                                var v84 bool = v35
                                                                _ = v84
                                                            }
                                                        default:
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                        }
                                                    }
                                                    if ((true && (v18 <= v18)) && ((v31 >> ((v1) & 63)) >= (v31 % ((2) | 1)))) {
                                                        var v85 int64 = (^(v30))
                                                        switch (^(fi34)) {
                                                        case 675:
                                                            switch "\n\t" {
                                                            case "":
                                                                var v86 uint16 = 2
                                                                _ = v86
                                                            case "foo bar":
                                                                v18 = (v18 + (v18 / (((v18 / (((v18 ^ 186)) | 1))) | 1)))
                                                            default:
                                                                v1 = v21
                                                            }
                                                        case -9223372036854775808:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case -544:
                                                            var v87 float32 = 0.1
                                                            _ = v87
                                                        default:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        _ = v85
                                                    }
                                                    v19 = float32(float64(v19))
                                                }
                                            }
                                        }
                                    }
                                    v31 = 42
                                }
                                v31 += 843
                            }
                            var v88 uint64 = (v31 | (v31 & 1))
                            v19 += v19
                            _ = v88
                        } else {
                            fmt.Println(v22)
                            fmt.Println(v30, v2, v19)
                            switch "a" {
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v29 = v20
                                if v20 {
                                    switch "hello" {
                                    case "a":
                                        v89 := GenS11{F0: 1, F1: (v23 | 32767)}
                                        _ = v89
                                    case "":
                                        switch "" {
                                        case "hello":
                                            v90 := GenS5{F0: 127}
                                            _ = v90
                                        case "a":
                                            v1 = (v1 | (^(v21)))
                                        case "":
                                            fmt.Println(v21, v30, v20)
                                        default:
                                            if ((-(fi34)) < fi34) {
                                                var v91 uint = (v21 | v1)
                                                v2 = -1
                                                _ = v91
                                            } else {
                                                switch (-(fi33)) {
                                                case 888:
                                                    var v92 int8 = -95
                                                    _ = v92
                                                case 9223372036854775807:
                                                    v3 = append(v3, "\n\t")
                                                }
                                                v18 -= (^(v18))
                                                v93 := GenS12{F0: v21, F1: 42, F2: 127}
                                                switch fi33 {
                                                case 127:
                                                    var v94 float32 = (v19 - v19)
                                                    _ = v94
                                                case -730:
                                                    v20 = (v20 || v35)
                                                }
                                                _ = v93
                                            }
                                            var v95 uint32 = v18
                                            _ = v95
                                        }
                                    }
                                    if (-465 >= (-(v23))) {
                                        switch "foo bar" {
                                        case "":
                                            v2 -= (-(v2))
                                        case "hello":
                                            v22 = v23
                                        case "a":
                                            v96 := GenS8{F0: -981, F1: 42, F2: -2147483648}
                                            _ = v96
                                        }
                                    } else {
                                        var v97 int32 = -351
                                        if v29 {
                                            v3 = append(v3, "\n\t")
                                            v97 = (v97 - v97)
                                        } else {
                                            if false {
                                                var v98 int64 = v4
                                                switch (fi34 << (((v1 % (((v21 & (v1 * (v1 / (((v1 + v21)) | 1))))) | 1))) & 63)) {
                                                case 2147483647:
                                                    v20 = (32767 < 312)
                                                }
                                                _ = v98
                                            } else {
                                                v99 := GenS13{F0: (-(v30)), F1: (v31 << ((1) & 63)), F2: 6}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v99
                                            }
                                            v29 = (!(!true))
                                            if v29 {
                                                var v100 int8 = int8(fi33)
                                                _ = v100
                                            }
                                        }
                                        v101 := GenS13{F0: 487, F1: (^(v31)), F2: 1}
                                        _ = v101
                                        _ = v97
                                    }
                                }
                            }
                        }
                        _ = v36
                    } else {
                        switch fi33 {
                        case -9223372036854775808:
                            v102 := GenS4{F0: v31, F1: 42}
                            var v103 int8 = -128
                            _ = v102
                            _ = v103
                        }
                        var v104 bool = v29
                        v105 := GenS14{F0: -94, F1: 0, F2: (v18 + v18)}
                        _ = v104
                        _ = v105
                    }
                    fmt.Println(v20, v29)
                    v20 = v35
                }
                var v106 uint = v21
                var v107 uint8 = 127
                _ = v106
                _ = v107
            }
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
