package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int16; F2 string }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 uint32; F2 string }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 string; F1 uint8 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 string; F1 string }
type GenS20 struct { F0 string; F1 string; F2 int32 }
type GenS21 struct { F0 uint32; F1 bool }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 uint8; F1 int32 }
type GenS25 struct { F0 int8; F1 int8 }
type GenS26 struct { F0 uint8; F1 string; F2 uint32 }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 uint }
type GenS29 struct { F0 int16; F1 int16; F2 uint64 }
type GenS30 struct { F0 bool }
type GenS31 struct { F0 int64; F1 uint8 }
type GenS32 struct { F0 uint32 }
type GenS33 struct { F0 string }
type GenS34 struct { F0 int64; F1 uint; F2 string }
type GenS35 struct { F0 int8; F1 uint }
type GenS36 struct { F0 uint16; F1 string }
type GenS37 struct { F0 int8; F1 int8; F2 uint8 }
type GenS38 struct { F0 uint32; F1 uint; F2 int16 }
type GenS39 struct { F0 int64; F1 uint16; F2 uint }
type GenS40 struct { F0 uint64 }
type GenS41 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS42 struct { F0 uint64 }
type GenS43 struct { F0 uint32; F1 uint64 }
type GenS44 struct { F0 uint8; F1 int16; F2 int8 }
type GenS45 struct { F0 uint32; F1 int; F2 uint }
type GenS46 struct { F0 uint16; F1 int16 }
type GenS47 struct { F0 string; F1 int; F2 uint }
type GenS48 struct { F0 int32 }
type GenS49 struct { F0 string; F1 uint16; F2 uint64 }
type GenS50 struct { F0 bool }
type GenS51 struct { F0 uint32; F1 string }
type GenS52 struct { F0 int16; F1 uint16 }
type GenS53 struct { F0 int8; F1 int32 }
type GenS54 struct { F0 int8 }
func genFn0() int64 {
    var v5 int8 = 1
    fmt.Println(v5)
    _ = v5
    return -408
}
func genFn1(p0 int8) string {
    var v6 uint = 349
    v7 := GenS1{F0: 42, F1: ("\n\t" + ("" + (("" + ("" + "")) + "\n\t"))), F2: genFn0()}
    _ = p0
    _ = v6
    _ = v7
    return ""
}

func main() {
    var v1 uint8 = 202
    var v2 uint = 622
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32(1.5)
    v3[2147483647] = 810
    v2 = (v2 % ((18446744073709551615) | 1))
    v8 := []int{1, 32768, 127, 32767}
    var v9 int64 = genFn0()
    if true {
        switch 1 {
        case 423:
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v2 = 648
                    _ = fi11
                }
                if (2147483647 <= 128) {
                    if (false || false) {
                        for fi12 := 0; fi12 < 20; fi12++ {
                            if false {
                                v13 := v3[115]
                                _ = len(v3)
                                v4 = v4
                                v14 := GenS1{F0: -411, F1: ("" + ("a" + "a")), F2: v9}
                                var v15 string = "hello"
                                v8 = append(v8, (fi12 | (fi10 * (fi12 % ((612) | 1)))))
                                var v16 int8 = -109
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            _ = fi12
                        }
                    } else {
                        v17 := GenS1{F0: 2147483647, F1: "", F2: genFn0()}
                        _ = v17
                    }
                    if len(v8) > 0 {
                        v8[0] = int(v1)
                    }
                } else {
                    if false {
                        v4 = 0.1
                        var v18 uint8 = (^(v1))
                        v1 += 2
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 10; fi19++ {
                            var v20 uint64 = 0
                            _ = fi19
                            _ = v20
                        }
                        if (v2 >= v2) {
                            v1 = (^(v1))
                            if (!((^(fi10)) >= (fi10 * (fi10 + (fi10 << (((^(v2))) & 63)))))) {
                                fmt.Println(fi10, v4)
                                v21 := v3[296]
                                v22 := GenS1{F0: 91, F1: "foo bar", F2: v9}
                                _ = v21
                                _ = v22
                            }
                        } else {
                            if true {
                                if false {
                                    fmt.Println(v2)
                                } else {
                                    v8 = append(v8, (fi10 / ((fi10) | 1)))
                                    fmt.Println(v2)
                                    var v23 int8 = 42
                                    _ = v23
                                }
                            }
                        }
                    }
                }
                _ = fi10
            }
        case 847:
            var v24 uint16 = uint16(v2)
            _ = v24
        case 2147483647:
            v25 := v3[185]
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS2{F0: ((("\n\t" + "a") + (("\n\t" + ("\n\t" + "\n\t")) + "\n\t")) + "\n\t")}
            v1 = v1
            v2 = (v2 | (^(v2)))
            _ = fi26
            _ = v27
        }
        v28 := GenS3{F0: 32767}
        _ = v28
    }
    fmt.Println(v9, v9)
    var v29 bool = true
    if v29 {
        var v30 uint32 = 2
        v2 -= (v2 % ((v2) | 1))
        v8 = append(v8, 1)
        _ = v30
    } else {
        if (!(!(0 >= 2147483647))) {
            fmt.Println(v4, v9)
        }
        if v29 {
            _ = len(v8)
            var v31 string = "a"
            for fi32 := 0; fi32 < 11; fi32++ {
                v33 := GenS4{F0: 1}
                var v34 uint64 = 2
                for fi35 := 0; fi35 < 10; fi35++ {
                    v36 := GenS3{F0: -1}
                    if v29 {
                        var v37 string = v31
                        var v38 int16 = -32768
                        _ = len(v8)
                        _ = v37
                        _ = v38
                    } else {
                        var v39 bool = false
                        var v40 uint64 = v34
                        v2 += (v2 | 255)
                        _ = v39
                        _ = v40
                    }
                    _ = fi35
                    _ = v36
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = v31
        } else {
            var v41 uint64 = 18446744073709551615
            v29 = true
            if v29 {
                fmt.Println(v41, v9, v29)
                if ((v1 ^ 2) == v1) {
                    for fi42 := 0; fi42 < 19; fi42++ {
                        for fi43 := 0; fi43 < 14; fi43++ {
                            v4 = v4
                            fmt.Println(fi43, fi43)
                            if false {
                                if (!true) {
                                    fmt.Println(v9, v1, v29)
                                }
                                var v44 uint = 42
                                switch ((("" + ("foo bar" + ("foo bar" + "foo bar"))) + ("foo bar" + "foo bar")) + (("" + "\n\t") + "")) {
                                case "hello":
                                    if v29 {
                                        if ((false || (!false)) && (v29 || v29)) {
                                            v8 = append(v8, fi42)
                                            v3[65535] = 2147483647
                                        } else {
                                            if v29 {
                                                if ((fi43 & (fi42 % ((2147483647) | 1))) < -325) {
                                                    v45 := GenS4{F0: 127}
                                                    v4 += 0.0
                                                    _ = v45
                                                }
                                                if (v29 || (108 <= 702)) {
                                                    v29 = false
                                                    if (-32768 >= int16(v9)) {
                                                        switch "hello" {
                                                        case "hello":
                                                            if (!(v41 < (^(v41)))) {
                                                                v46 := GenS5{F0: (^(v44))}
                                                                var v47 uint64 = v41
                                                                _ = v46
                                                                _ = v47
                                                            } else {
                                                                var v48 uint8 = v1
                                                                _ = v48
                                                            }
                                                        case "a":
                                                            var v49 int16 = 1
                                                            _ = v49
                                                        case "":
                                                            var v50 int = fi43
                                                            _ = v50
                                                        }
                                                        v1 -= v1
                                                    }
                                                    if (839 >= 42) {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v41, v44, v9)
                                                v51 := GenS4{F0: -57}
                                                if v29 {
                                                    var v52 float64 = -2.25
                                                    var v53 uint8 = (v1 * v1)
                                                    _ = v52
                                                    _ = v53
                                                }
                                                _ = v51
                                            }
                                        }
                                        switch (fi43 >> ((656) & 63)) {
                                        case 9223372036854775807:
                                            v8 = append(v8, (fi43 / ((int(v9)) | 1)))
                                        case 623:
                                            if true {
                                                v4 = v4
                                            } else {
                                                if false {
                                                    fmt.Println(fi42, fi43, v29)
                                                } else {
                                                    v54 := GenS1{F0: 0, F1: ("a" + ("a" + ("" + "\n\t"))), F2: (v9 & v9)}
                                                    v41 -= v41
                                                    _ = v54
                                                }
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                                v29 = v29
                                            }
                                        default:
                                            var v55 uint16 = 2
                                            _ = v55
                                        }
                                        if (!v29) {
                                            v1 = (^(v1))
                                            fmt.Println(v2, v29, v44)
                                        }
                                    }
                                case "":
                                    v2 -= v2
                                case "a":
                                    if (false || (!v29)) {
                                        v56 := GenS6{F0: 2, F1: v41}
                                        v57 := GenS6{F0: v44, F1: (^(v41))}
                                        fmt.Println(v44, v41, v41)
                                        _ = v56
                                        _ = v57
                                    } else {
                                        switch (fi42 << (((^(v2))) & 63)) {
                                        case 128:
                                            v8 = append(v8, 4294967295)
                                        default:
                                            v58 := GenS7{F0: 0, F1: 42, F2: "hello"}
                                            _ = v58
                                        }
                                        if (885 <= -2147483648) {
                                            if ((v41 % (((v41 << ((uint(v41)) & 63))) | 1)) < v41) {
                                                var v59 float32 = float32((0.1 - 3.14159))
                                                var v60 int = (fi43 & 482)
                                                v44 = (v44 / ((uint(v1)) | 1))
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                        v41 += v41
                                        v61 := GenS8{F0: 18446744073709551615, F1: 0}
                                        if v29 {
                                            var v62 int = fi42
                                            if ((v1 & v1) != v1) {
                                                if ("foo bar" != "foo bar") {
                                                    fmt.Println(v1)
                                                    v4 += v4
                                                } else {
                                                    v29 = v29
                                                    if (v41 <= (v41 + v41)) {
                                                        if (0 <= (v9 ^ v9)) {
                                                            v29 = v29
                                                        }
                                                        v4 = v4
                                                    } else {
                                                        var v63 int8 = 42
                                                        if v29 {
                                                            var v64 float32 = v4
                                                            var v65 int64 = (v9 + 0)
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch v62 {
                                                                case 691:
                                                                    v3[2147483647] = 117
                                                                default:
                                                                    if ((!(!v29)) || (((v64 - (v64 + (0.0 + (float32((0.1 / 0.1)) - v4)))) + v4) != (v4 * v4))) {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v4 += v4
                                                                        case "\n\t":
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                        default:
                                                                            v66 := GenS9{F0: (v44 >> (((v2 / ((v2) | 1))) & 63)), F1: v1}
                                                                            _ = v66
                                                                        }
                                                                        fmt.Println(v62, v44)
                                                                        v67 := GenS6{F0: v44, F1: (v41 << (((v44 % ((v44) | 1))) & 63))}
                                                                        v68 := GenS3{F0: v62}
                                                                        var v69 bool = v29
                                                                        _ = v67
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                }
                                                                var v70 uint64 = v41
                                                                _ = v70
                                                            case "":
                                                                if (!(v29 || (false && ((v44 - v44) < v2)))) {
                                                                    v3[1] = 2147483647
                                                                    var v71 float32 = (v4 - v64)
                                                                    _ = v71
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(fi42)
                                                            default:
                                                                switch v62 {
                                                                case -180:
                                                                    var v72 float32 = 1e10
                                                                    _ = v72
                                                                case 32768:
                                                                    v73 := GenS4{F0: -13}
                                                                    _ = v73
                                                                case 42:
                                                                    var v74 int8 = v63
                                                                    _ = v74
                                                                }
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            v3[769] = 392
                                                            if (true && (((!true) || v29) || v29)) {
                                                                fmt.Println(v44)
                                                                var v75 int8 = (v63 - v63)
                                                                if (-32768 <= -958) {
                                                                    if v29 {
                                                                        var v76 uint = (v44 | (v2 & v44))
                                                                        if v29 {
                                                                            var v77 bool = true
                                                                            switch (fi42 % ((2147483647) | 1)) {
                                                                            case 1:
                                                                                v78 := GenS8{F0: 922, F1: (^(v41))}
                                                                                _ = v78
                                                                            default:
                                                                                var v79 uint16 = 489
                                                                                _ = v79
                                                                            }
                                                                            _ = v77
                                                                        }
                                                                        _ = v76
                                                                    }
                                                                } else {
                                                                    delete(v3, 2)
                                                                }
                                                                _ = v75
                                                            } else {
                                                                var v80 string = "hello"
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                if v29 {
                                                                    if ((!false) || (2147483647 != 42)) {
                                                                        v81 := GenS10{F0: 1}
                                                                        _ = len(v8)
                                                                        v2 -= v44
                                                                        _ = v81
                                                                    } else {
                                                                        if (0 != int64(v1)) {
                                                                            var v82 int32 = -1
                                                                            var v83 float64 = 1e10
                                                                            var v84 float64 = (v83 + 0.0)
                                                                            _ = v82
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            if v29 {
                                                                                v2 -= (^(v44))
                                                                                v85 := GenS11{F0: 65535, F1: 138}
                                                                                switch v62 {
                                                                                case 4294967295:
                                                                                    fmt.Println(fi42, v1, fi42)
                                                                                case 869:
                                                                                    fmt.Println(fi42, v4)
                                                                                }
                                                                                v86 := GenS12{F0: ((true && (v1 <= uint8(v41))) && (v29 || (0 == int32(v63)))), F1: 0, F2: (v1 / ((0) | 1))}
                                                                                v4 -= float32(3.14159)
                                                                                _ = v85
                                                                                _ = v86
                                                                            } else {
                                                                                _ = len(v8)
                                                                            }
                                                                            v87 := GenS2{F0: v80}
                                                                            _ = v87
                                                                        }
                                                                        var v88 int32 = 1
                                                                        var v89 uint32 = 803
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                _ = v80
                                                            }
                                                            v62 += fi42
                                                        }
                                                        fmt.Println(fi42)
                                                        _ = v63
                                                    }
                                                    if false {
                                                        if (!(42 >= genFn0())) {
                                                            v3[42] = 428
                                                            v1 = (v1 * (v1 * v1))
                                                            switch "a" {
                                                            case "a":
                                                                var v90 uint64 = 1
                                                                _ = v90
                                                            case "foo bar":
                                                                fmt.Println(fi42)
                                                            default:
                                                                if v29 {
                                                                    v91 := GenS3{F0: (v62 << ((42) & 63))}
                                                                    if (0 >= 2147483648) {
                                                                        var v92 int16 = 1000
                                                                        v9 = 42
                                                                        _ = v92
                                                                    }
                                                                    var v93 int32 = 711
                                                                    _ = v91
                                                                    _ = v93
                                                                } else {
                                                                    switch (fi43 + fi43) {
                                                                    case 32767:
                                                                        v94 := GenS2{F0: "hello"}
                                                                        _ = v94
                                                                    }
                                                                    v8 = append(v8, (^(fi43)))
                                                                }
                                                            }
                                                            v41 += (v41 / ((v41) | 1))
                                                            if (v29 && v29) {
                                                                if ((116 >= -1) && true) {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        var v95 int = (v62 - 85)
                                                                        _ = v95
                                                                    case "":
                                                                        v96 := GenS6{F0: (^(v44)), F1: v41}
                                                                        _ = v96
                                                                    case "hello":
                                                                        switch v62 {
                                                                        case -242:
                                                                            if len(v8) > 0 {
                                                                                v8[0] = -948
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch (fi42 | fi42) {
                                                                        case 1:
                                                                            v4 += v4
                                                                        case -93:
                                                                            var v97 int32 = 563
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                }
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v62)
                                                                    v2 += 1
                                                                    var v98 float32 = float32(((-2.25 * 1e10) - 0.0))
                                                                    _ = v98
                                                                }
                                                                v99 := GenS3{F0: 9223372036854775807}
                                                                _ = v99
                                                            }
                                                            var v100 uint64 = 0
                                                            v62 = (fi42 & (fi43 & int(v44)))
                                                            _ = v100
                                                        } else {
                                                            var v101 float32 = 1.5
                                                            switch (fi43 << ((uint(v41)) & 63)) {
                                                            case -93:
                                                                var v102 bool = v29
                                                                _ = v102
                                                            case 9223372036854775807:
                                                                var v103 uint8 = 10
                                                                _ = v103
                                                            default:
                                                                v104 := GenS5{F0: 0}
                                                                _ = v104
                                                            }
                                                            var v105 int8 = 42
                                                            v1 = 1
                                                            if v29 {
                                                                switch ("" + "a") {
                                                                case "foo bar":
                                                                    v106 := GenS4{F0: v105}
                                                                    _ = v106
                                                                case "hello":
                                                                    v101 = (((v4 - -2.25) * (0.0 - (v4 - (v4 + (0.1 - 1e10))))) / (v101 / float32(-2.25)))
                                                                case "":
                                                                    v2 = (v2 / (((v2 & v2)) | 1))
                                                                }
                                                                switch 2147483647 {
                                                                case -9223372036854775808:
                                                                    v62 = (^(fi43))
                                                                    var v107 uint8 = (v1 / ((v1) | 1))
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if true {
                                                                            v108 := GenS13{F0: (v2 >> (((v2 + v44)) & 63))}
                                                                            var v109 int32 = -147
                                                                            _ = v108
                                                                            _ = v109
                                                                        } else {
                                                                            switch ("" + "a") {
                                                                            case "\n\t":
                                                                                v110 := GenS14{F0: 978}
                                                                                _ = v110
                                                                            case "hello":
                                                                                switch 4294967295 {
                                                                                case -1:
                                                                                    v44 += 255
                                                                                }
                                                                            default:
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = fi43
                                                                                }
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(v29, v9, v1)
                                                                        var v111 int8 = 15
                                                                        _ = len(v3)
                                                                        _ = v111
                                                                    case "\n\t":
                                                                        v41 = 151
                                                                    }
                                                                    _ = v107
                                                                case -1:
                                                                    var v112 int8 = int8(v62)
                                                                    _ = v112
                                                                }
                                                            }
                                                            _ = v101
                                                            _ = v105
                                                        }
                                                        v113 := v3[2147483648]
                                                        if len(v8) > 0 {
                                                            v8[0] = 32767
                                                        }
                                                        _ = v113
                                                    }
                                                    var v114 uint16 = 32767
                                                    _ = v114
                                                }
                                                switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                case "foo bar":
                                                    var v115 uint = v44
                                                    fmt.Println(v29, v2, v9)
                                                    v116 := GenS15{F0: 49, F1: 842, F2: "a"}
                                                    _ = v115
                                                    _ = v116
                                                }
                                            }
                                            switch "a" {
                                            case "":
                                                v62 -= 127
                                            case "\n\t":
                                                switch (fi42 * (^(v62))) {
                                                case 1:
                                                    v117 := GenS4{F0: -1}
                                                    _ = v117
                                                case -129:
                                                    switch -601 {
                                                    case 929:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v1 = (^(v1))
                                                        case "foo bar":
                                                            if v29 {
                                                                v118 := GenS15{F0: -277, F1: 65535, F2: ""}
                                                                var v119 int = (^(fi42))
                                                                _ = v118
                                                                _ = v119
                                                            } else {
                                                                v41 -= (v41 & v41)
                                                                v120 := GenS16{F0: v41}
                                                                v29 = v29
                                                                _ = v120
                                                            }
                                                        case "a":
                                                            delete(v3, 2147483648)
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                var v121 uint16 = 65535
                                                _ = v121
                                            }
                                            _ = v62
                                        } else {
                                            if (v4 < v4) {
                                                fmt.Println(v4, v29)
                                                v4 -= (((v4 + v4) / v4) - -2.25)
                                            }
                                            var v122 bool = (uint32(v44) >= 65535)
                                            _ = v122
                                        }
                                        switch -623 {
                                        case 42:
                                            switch 0 {
                                            case 127:
                                                v123 := GenS17{F0: "foo bar", F1: (v1 >> ((1) & 63))}
                                                _ = v123
                                            default:
                                                v124 := v3[326]
                                                _ = v124
                                            }
                                        case 0:
                                            fmt.Println(fi42, fi42)
                                        case -9223372036854775808:
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v44, v29, fi42)
                                            case "foo bar":
                                                if (false && false) {
                                                    v41 += (^(v41))
                                                }
                                            default:
                                                if v29 {
                                                    switch "\n\t" {
                                                    case "":
                                                        var v125 uint = v44
                                                        v41 += 0
                                                        _ = v125
                                                    case "hello":
                                                        _ = len(v8)
                                                    case "\n\t":
                                                        var v126 float64 = ((0.1 - float64((0.1 / v4))) - (float64(v4) * ((1.5 * (0.1 / float64(v4))) + -2.25)))
                                                        _ = v126
                                                    default:
                                                        v127 := v3[2147483648]
                                                        _ = v127
                                                    }
                                                    v128 := GenS16{F0: 42}
                                                    switch ("" + "") {
                                                    case "\n\t":
                                                        v9 -= -9223372036854775808
                                                    case "a":
                                                        _ = len(v3)
                                                    case "":
                                                        fmt.Println(v29, v41, v4)
                                                    }
                                                    _ = v128
                                                } else {
                                                    if ((^(v9)) <= (^(v9))) {
                                                        var v129 uint = (v2 % (((v2 + v2)) | 1))
                                                        _ = v129
                                                    }
                                                    v130 := GenS18{F0: v2}
                                                    _ = v130
                                                }
                                            }
                                        default:
                                            if v29 {
                                                v2 += 4294967295
                                                v131 := GenS19{F0: "foo bar", F1: ""}
                                                if v29 {
                                                    var v132 int32 = -647
                                                    if true {
                                                        _ = len(v8)
                                                        if true {
                                                            v44 = (v2 & (^(v2)))
                                                            v1 = v1
                                                            delete(v3, 195)
                                                        }
                                                    } else {
                                                        var v133 uint32 = 2
                                                        _ = v133
                                                    }
                                                    _ = v132
                                                }
                                                _ = v131
                                            }
                                        }
                                        _ = v61
                                    }
                                default:
                                    v2 = v2
                                }
                                if v29 {
                                    v1 = v1
                                } else {
                                    fmt.Println(v1, v41, v2)
                                    if (!v29) {
                                        switch "" {
                                        case "hello":
                                            v134 := GenS20{F0: "", F1: (("a" + ("hello" + "foo bar")) + ""), F2: -2147483648}
                                            _ = v134
                                        case "a":
                                            if v29 {
                                                if ((-128 <= 1) && (!(!v29))) {
                                                    v135 := GenS3{F0: (fi42 / (((-(fi42))) | 1))}
                                                    v2 -= v44
                                                    var v136 string = "hello"
                                                    _ = v135
                                                    _ = v136
                                                } else {
                                                    v137 := GenS21{F0: 42, F1: v29}
                                                    var v138 int8 = -128
                                                    v3[2147483648] = 865
                                                    _ = v137
                                                    _ = v138
                                                }
                                                fmt.Println(v44, v44)
                                                v4 = v4
                                            } else {
                                                if ((((((("foo bar" + "hello") + ("a" + "hello")) + "\n\t") + "\n\t") + "a") + "") != ((("" + "a") + "foo bar") + ("\n\t" + "foo bar"))) {
                                                    switch (fi43 & fi43) {
                                                    case 4294967295:
                                                        var v139 int8 = 116
                                                        _ = v139
                                                    }
                                                    v4 += (float32(((0.1 + 0.1) / (0.1 / 1.0))) * -2.25)
                                                } else {
                                                    switch (fi43 >> ((0) & 63)) {
                                                    case 42:
                                                        if v29 {
                                                            if v29 {
                                                                fmt.Println(v29)
                                                                v140 := GenS9{F0: 1, F1: v1}
                                                                _ = v140
                                                            } else {
                                                                switch -9223372036854775808 {
                                                                case -9223372036854775808:
                                                                    var v141 uint32 = 4294967295
                                                                    _ = v141
                                                                case -625:
                                                                    if (v29 || (4294967295 > v41)) {
                                                                        var v142 uint64 = (^(v41))
                                                                        v143 := GenS2{F0: (("a" + "hello") + "a")}
                                                                        _ = v142
                                                                        _ = v143
                                                                    } else {
                                                                        v144 := GenS22{F0: v41}
                                                                        if true {
                                                                            var v145 int = fi43
                                                                            v146 := GenS19{F0: "foo bar", F1: (("foo bar" + "") + (((("a" + "foo bar") + "a") + "\n\t") + "a"))}
                                                                            v147 := GenS23{F0: (v44 >> (((v44 + (v2 ^ v44))) & 63))}
                                                                            _ = v145
                                                                            _ = v146
                                                                            _ = v147
                                                                        } else {
                                                                            v9 += v9
                                                                            v148 := GenS18{F0: v2}
                                                                            _ = v148
                                                                        }
                                                                        if true {
                                                                            var v149 float64 = 0.0
                                                                            v2 = (v44 << ((v2) & 63))
                                                                            v149 = 0.1
                                                                            _ = v149
                                                                        } else {
                                                                            v1 = 127
                                                                            delete(v3, 2147483647)
                                                                        }
                                                                        _ = v144
                                                                    }
                                                                case 486:
                                                                    var v150 uint16 = 0
                                                                    _ = v150
                                                                default:
                                                                    if false {
                                                                        v151 := GenS16{F0: 4294967295}
                                                                        fmt.Println(fi42, v1)
                                                                        var v152 int = (fi43 & fi43)
                                                                        _ = v151
                                                                        _ = v152
                                                                    }
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v2 = v2
                                                            default:
                                                                v2 = (v2 >> (((^(v44))) & 63))
                                                            }
                                                            v153 := GenS24{F0: v1, F1: -452}
                                                            v2 = v44
                                                            if v29 {
                                                                switch fi43 {
                                                                case 127:
                                                                    var v154 uint = 4294967295
                                                                    _ = v154
                                                                case 666:
                                                                    var v155 string = "\n\t"
                                                                    _ = v155
                                                                case 752:
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi43
                                                                    }
                                                                }
                                                                if false {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        delete(v3, 65535)
                                                                        var v156 uint16 = 255
                                                                        _ = v156
                                                                    default:
                                                                        var v157 uint8 = v1
                                                                        _ = v157
                                                                    }
                                                                    v158 := GenS25{F0: int8(v41), F1: 42}
                                                                    v159 := GenS26{F0: v1, F1: ("foo bar" + "a"), F2: 4294967295}
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    fmt.Println(fi43)
                                                                case "hello":
                                                                    v29 = (1000 <= -802)
                                                                default:
                                                                    if v29 {
                                                                        if v29 {
                                                                            v160 := GenS27{F0: 32767}
                                                                            var v161 int = (^(fi43))
                                                                            if false {
                                                                                fmt.Println(v9, v1)
                                                                                v4 = v4
                                                                                if v29 {
                                                                                    var v162 uint8 = v1
                                                                                    if (-1 < 0) {
                                                                                        var v163 int32 = 1
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            if ((uint32(v44) != 786) && v29) {
                                                                                                if (true && ((1 < uint16(v2)) || true)) {
                                                                                                    var v164 int32 = v163
                                                                                                    v2 -= (v2 % (((v2 / ((18446744073709551615) | 1))) | 1))
                                                                                                    _ = v164
                                                                                                } else {
                                                                                                    v165 := GenS7{F0: 2147483648, F1: -1, F2: "hello"}
                                                                                                    switch (("hello" + "a") + (("hello" + ("\n\t" + "a")) + "\n\t")) {
                                                                                                    case "a":
                                                                                                        var v166 int8 = -128
                                                                                                        _ = v166
                                                                                                    case "\n\t":
                                                                                                        v167 := GenS1{F0: (-(v163)), F1: "\n\t", F2: v9}
                                                                                                        _ = v167
                                                                                                    default:
                                                                                                        var v168 string = ""
                                                                                                        _ = v168
                                                                                                    }
                                                                                                    v2 -= (v2 / (((^(v44))) | 1))
                                                                                                    _ = v165
                                                                                                }
                                                                                                _ = len(v8)
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v9, v44, v4)
                                                                                        }
                                                                                        _ = v163
                                                                                    } else {
                                                                                        v169 := GenS19{F0: "foo bar", F1: ("\n\t" + "foo bar")}
                                                                                        if v29 {
                                                                                            v170 := GenS6{F0: 65535, F1: 42}
                                                                                            var v171 int64 = (v9 % ((930) | 1))
                                                                                            v29 = false
                                                                                            v172 := GenS28{F0: 65535}
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    if v29 {
                                                                                                        var v173 string = "a"
                                                                                                        var v174 uint64 = (v41 * v41)
                                                                                                        if true {
                                                                                                            v175 := GenS5{F0: 367}
                                                                                                            _ = v175
                                                                                                        }
                                                                                                        _ = v173
                                                                                                        _ = v174
                                                                                                    } else {
                                                                                                        fmt.Println(v44, v9, v41)
                                                                                                        v8 = append(v8, fi43)
                                                                                                        v176 := GenS29{F0: -521, F1: 0, F2: (^(v41))}
                                                                                                        v3[42] = 99
                                                                                                        _ = v176
                                                                                                    }
                                                                                                    var v177 uint16 = 42
                                                                                                    v3[4294967295] = 4294967295
                                                                                                    if len(v8) >= 2 {
                                                                                                        v8 = v8[1:len(v8)]
                                                                                                    }
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        v178 := GenS30{F0: true}
                                                                                                        _ = v178
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            if (!(!((fi43 << ((v44) & 63)) == 42))) {
                                                                                                                if v29 {
                                                                                                                    if false {
                                                                                                                        _ = len(v3)
                                                                                                                    } else {
                                                                                                                        fmt.Println(v177)
                                                                                                                        fmt.Println(v171, v162)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v179 int32 = 0
                                                                                                                    switch "hello" {
                                                                                                                    case "\n\t":
                                                                                                                        v180 := GenS23{F0: v2}
                                                                                                                        _ = v180
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v162)
                                                                                                                    case "hello":
                                                                                                                        fmt.Println(v162)
                                                                                                                    default:
                                                                                                                        if (0.0 > (float64(v4) / 1.5)) {
                                                                                                                            if false {
                                                                                                                                if true {
                                                                                                                                    v177 -= v177
                                                                                                                                    if ((v177 ^ v177) <= (v177 << ((v44) & 63))) {
                                                                                                                                        var v181 uint32 = 2
                                                                                                                                        _ = len(v3)
                                                                                                                                        switch 42 {
                                                                                                                                        case -818:
                                                                                                                                            var v182 int = fi42
                                                                                                                                            if len(v8) >= 2 {
                                                                                                                                                v8 = v8[1:len(v8)]
                                                                                                                                            }
                                                                                                                                            fmt.Println(v9)
                                                                                                                                            switch ("foo bar" + "") {
                                                                                                                                            case "\n\t":
                                                                                                                                                v1 = 127
                                                                                                                                                v1 = v1
                                                                                                                                                if false {
                                                                                                                                                    fmt.Println(v179, v4)
                                                                                                                                                }
                                                                                                                                                if (-969 >= -1) {
                                                                                                                                                    v44 -= v44
                                                                                                                                                    var v183 uint64 = 18446744073709551615
                                                                                                                                                    _ = v183
                                                                                                                                                } else {
                                                                                                                                                    if (!(!(!(7 == uint8(v41))))) {
                                                                                                                                                        var v184 bool = (float32(-2.25) > v4)
                                                                                                                                                        switch "a" {
                                                                                                                                                        case "foo bar":
                                                                                                                                                            v184 = ("hello" > "hello")
                                                                                                                                                        case "":
                                                                                                                                                            var v185 uint32 = v181
                                                                                                                                                            _ = v185
                                                                                                                                                        case "\n\t":
                                                                                                                                                            if false {
                                                                                                                                                                if v29 {
                                                                                                                                                                    fmt.Println(v9, v44)
                                                                                                                                                                }
                                                                                                                                                                fmt.Println(v41)
                                                                                                                                                                if v29 {
                                                                                                                                                                    if ((v9 ^ v171) == (v171 - v9)) {
                                                                                                                                                                        var v186 int16 = 32767
                                                                                                                                                                        v181 += v181
                                                                                                                                                                        if ((!false) || false) {
                                                                                                                                                                            if (v177 <= (^(v177))) {
                                                                                                                                                                                if (true && v184) {
                                                                                                                                                                                    if ((!(-1 > 42)) && ((v179 - v179) > (v179 - v179))) {
                                                                                                                                                                                        v179 += (v179 >> ((v44) & 63))
                                                                                                                                                                                    } else {
                                                                                                                                                                                        _ = len(v8)
                                                                                                                                                                                        v181 += v181
                                                                                                                                                                                        fmt.Println(v2)
                                                                                                                                                                                    }
                                                                                                                                                                                    switch v182 {
                                                                                                                                                                                    case -251:
                                                                                                                                                                                        if (((1.5 - (float64(float32((1.5 + 1.5))) - 1.5)) - (-2.25 + (float64(0.0) + 0.0))) < 3.14159) {
                                                                                                                                                                                            var v187 uint32 = (^(v181))
                                                                                                                                                                                            v171 -= int64(v181)
                                                                                                                                                                                            _ = v187
                                                                                                                                                                                        } else {
                                                                                                                                                                                            v188 := GenS31{F0: (v171 % ((-9223372036854775808) | 1)), F1: (^(v1))}
                                                                                                                                                                                            _ = v188
                                                                                                                                                                                        }
                                                                                                                                                                                    case 127:
                                                                                                                                                                                        v189 := GenS32{F0: 2147483648}
                                                                                                                                                                                        _ = v189
                                                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                                                        v3[uint32(v41)] = v181
                                                                                                                                                                                    default:
                                                                                                                                                                                        var v190 int8 = 0
                                                                                                                                                                                        _ = v190
                                                                                                                                                                                    }
                                                                                                                                                                                    var v191 int32 = v179
                                                                                                                                                                                    _ = v191
                                                                                                                                                                                }
                                                                                                                                                                                if (!(false && v184)) {
                                                                                                                                                                                    v162 -= v162
                                                                                                                                                                                    fmt.Println(v186)
                                                                                                                                                                                }
                                                                                                                                                                                v162 = (v1 + (v1 ^ 127))
                                                                                                                                                                                v192 := GenS33{F0: "a"}
                                                                                                                                                                                fmt.Println(v2, v186, v181)
                                                                                                                                                                                _ = v192
                                                                                                                                                                            }
                                                                                                                                                                            var v193 int8 = -69
                                                                                                                                                                            var v194 float32 = float32(-2.25)
                                                                                                                                                                            _ = v193
                                                                                                                                                                            _ = v194
                                                                                                                                                                        } else {
                                                                                                                                                                            fmt.Println(v161, v29, v162)
                                                                                                                                                                            v195 := v3[v181]
                                                                                                                                                                            v162 = 213
                                                                                                                                                                            _ = v195
                                                                                                                                                                        }
                                                                                                                                                                        _ = v186
                                                                                                                                                                    } else {
                                                                                                                                                                        v41 = v41
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v44, v44, v182)
                                                                                                                                                                    if false {
                                                                                                                                                                        v29 = v29
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    switch (v182 & 599) {
                                                                                                                                                                    case 2147483647:
                                                                                                                                                                        if v184 {
                                                                                                                                                                            var v196 int16 = -310
                                                                                                                                                                            if true {
                                                                                                                                                                                var v197 uint64 = 924
                                                                                                                                                                                if (-1 <= 2147483647) {
                                                                                                                                                                                    var v198 uint8 = v1
                                                                                                                                                                                    if len(v8) >= 2 {
                                                                                                                                                                                        v8 = v8[1:len(v8)]
                                                                                                                                                                                    }
                                                                                                                                                                                    v182 += (-(v161))
                                                                                                                                                                                    _ = v198
                                                                                                                                                                                }
                                                                                                                                                                                _ = len(v8)
                                                                                                                                                                                _ = v197
                                                                                                                                                                            } else {
                                                                                                                                                                                switch "" {
                                                                                                                                                                                case "":
                                                                                                                                                                                    fmt.Println(v1)
                                                                                                                                                                                case "foo bar":
                                                                                                                                                                                    if true {
                                                                                                                                                                                        var v199 uint8 = v1
                                                                                                                                                                                        _ = v199
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v44 += v44
                                                                                                                                                                                    }
                                                                                                                                                                                case "a":
                                                                                                                                                                                    v29 = v184
                                                                                                                                                                                default:
                                                                                                                                                                                    switch v161 {
                                                                                                                                                                                    case 127:
                                                                                                                                                                                        var v200 int8 = 0
                                                                                                                                                                                        _ = v200
                                                                                                                                                                                    default:
                                                                                                                                                                                        if (v41 != (^(v41))) {
                                                                                                                                                                                            v162 += v162
                                                                                                                                                                                        }
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                if (v29 || (-858 == -2147483648)) {
                                                                                                                                                                                    fmt.Println(v161)
                                                                                                                                                                                } else {
                                                                                                                                                                                    if v184 {
                                                                                                                                                                                        var v201 uint16 = (v177 >> ((v44) & 63))
                                                                                                                                                                                        fmt.Println(v179)
                                                                                                                                                                                        var v202 string = ""
                                                                                                                                                                                        _ = v201
                                                                                                                                                                                        _ = v202
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v203 := GenS34{F0: int64(v1), F1: (v44 / ((v44) | 1)), F2: (("\n\t" + "") + "foo bar")}
                                                                                                                                                                                        v204 := v3[v181]
                                                                                                                                                                                        v204 += v181
                                                                                                                                                                                        _ = v203
                                                                                                                                                                                        _ = v204
                                                                                                                                                                                    }
                                                                                                                                                                                    var v205 uint32 = 4294967295
                                                                                                                                                                                    _ = v205
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            v182 += (fi43 % ((fi42) | 1))
                                                                                                                                                                            _ = v196
                                                                                                                                                                        } else {
                                                                                                                                                                            var v206 bool = (v162 >= (^(v1)))
                                                                                                                                                                            _ = v206
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        v3[(^(v181))] = 697
                                                                                                                                                                    }
                                                                                                                                                                    if (!v29) {
                                                                                                                                                                        v207 := GenS26{F0: (v1 | 60), F1: ("" + (("foo bar" + ("" + ((("foo bar" + "") + "") + ("hello" + "hello")))) + (("hello" + ("hello" + ("a" + "\n\t"))) + "a"))), F2: v181}
                                                                                                                                                                        v2 = v2
                                                                                                                                                                        v29 = (!v184)
                                                                                                                                                                        _ = v207
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                v4 += v4
                                                                                                                                                                var v208 float64 = (3.14159 * -2.25)
                                                                                                                                                                _ = v208
                                                                                                                                                            }
                                                                                                                                                            v209 := GenS32{F0: (v181 << (((v44 - uint(v41))) & 63))}
                                                                                                                                                            v29 = (uint(v177) >= (v44 % (((v44 % (((v44 << ((0) & 63))) | 1))) | 1)))
                                                                                                                                                            _ = v209
                                                                                                                                                        }
                                                                                                                                                        var v210 int32 = (v179 % ((v179) | 1))
                                                                                                                                                        _ = v184
                                                                                                                                                        _ = v210
                                                                                                                                                    } else {
                                                                                                                                                        var v211 float32 = 1.5
                                                                                                                                                        _ = v211
                                                                                                                                                    }
                                                                                                                                                    v212 := GenS15{F0: -32768, F1: 54, F2: ((("hello" + ("a" + "")) + ("\n\t" + "")) + "\n\t")}
                                                                                                                                                    _ = v212
                                                                                                                                                }
                                                                                                                                                var v213 uint16 = 32767
                                                                                                                                                v214 := GenS31{F0: (v171 & v171), F1: 255}
                                                                                                                                                if (!true) {
                                                                                                                                                    if true {
                                                                                                                                                        switch ("" + "a") {
                                                                                                                                                        case "a":
                                                                                                                                                            var v215 int = (fi43 << ((v44) & 63))
                                                                                                                                                            _ = v215
                                                                                                                                                        case "\n\t":
                                                                                                                                                            var v216 int = int(v179)
                                                                                                                                                            _ = v216
                                                                                                                                                        default:
                                                                                                                                                            var v217 uint16 = (v177 << (((v44 << ((v44) & 63))) & 63))
                                                                                                                                                            _ = v217
                                                                                                                                                        }
                                                                                                                                                        var v218 float32 = v4
                                                                                                                                                        if (false || v29) {
                                                                                                                                                            var v219 uint64 = (v41 << ((18446744073709551615) & 63))
                                                                                                                                                            switch "a" {
                                                                                                                                                            case "":
                                                                                                                                                                var v220 uint16 = (v213 | v177)
                                                                                                                                                                _ = v220
                                                                                                                                                            case "\n\t":
                                                                                                                                                                fmt.Println(fi43)
                                                                                                                                                            default:
                                                                                                                                                                var v221 float64 = 1.5
                                                                                                                                                                _ = v221
                                                                                                                                                            }
                                                                                                                                                            var v222 float64 = 3.14159
                                                                                                                                                            v1 -= v1
                                                                                                                                                            var v223 int32 = v179
                                                                                                                                                            _ = v219
                                                                                                                                                            _ = v222
                                                                                                                                                            _ = v223
                                                                                                                                                        }
                                                                                                                                                        _ = v218
                                                                                                                                                    } else {
                                                                                                                                                        v224 := GenS4{F0: 41}
                                                                                                                                                        if v29 {
                                                                                                                                                            switch "a" {
                                                                                                                                                            case "foo bar":
                                                                                                                                                                if v29 {
                                                                                                                                                                    v225 := GenS25{F0: int8(v213), F1: 26}
                                                                                                                                                                    var v226 float32 = v4
                                                                                                                                                                    _ = v225
                                                                                                                                                                    _ = v226
                                                                                                                                                                } else {
                                                                                                                                                                    if v29 {
                                                                                                                                                                        switch -792 {
                                                                                                                                                                        case -763:
                                                                                                                                                                            fmt.Println(v179, v162, v9)
                                                                                                                                                                        case 2147483647:
                                                                                                                                                                            var v227 uint64 = 2
                                                                                                                                                                            _ = v227
                                                                                                                                                                        default:
                                                                                                                                                                            fmt.Println(v181)
                                                                                                                                                                        }
                                                                                                                                                                        var v228 uint16 = v177
                                                                                                                                                                        _ = v228
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                fmt.Println(v179, v182)
                                                                                                                                                            }
                                                                                                                                                            v229 := GenS3{F0: 42}
                                                                                                                                                            _ = v229
                                                                                                                                                        }
                                                                                                                                                        _ = v224
                                                                                                                                                    }
                                                                                                                                                    switch (fi43 % (((-(v161))) | 1)) {
                                                                                                                                                    case 127:
                                                                                                                                                        var v230 string = "a"
                                                                                                                                                        _ = v230
                                                                                                                                                    }
                                                                                                                                                    v181 = v181
                                                                                                                                                }
                                                                                                                                                v231 := GenS5{F0: v2}
                                                                                                                                                _ = v213
                                                                                                                                                _ = v214
                                                                                                                                                _ = v231
                                                                                                                                            case "foo bar":
                                                                                                                                                var v232 uint8 = 2
                                                                                                                                                _ = v232
                                                                                                                                            }
                                                                                                                                            fmt.Println(v1, v2)
                                                                                                                                            _ = v182
                                                                                                                                        }
                                                                                                                                        _ = v181
                                                                                                                                    }
                                                                                                                                    fmt.Println(v177)
                                                                                                                                } else {
                                                                                                                                    var v233 string = ((("foo bar" + "\n\t") + "hello") + "foo bar")
                                                                                                                                    v177 -= 0
                                                                                                                                    var v234 int8 = -1
                                                                                                                                    _ = v233
                                                                                                                                    _ = v234
                                                                                                                                }
                                                                                                                                if v29 {
                                                                                                                                    var v235 int8 = 1
                                                                                                                                    v235 += v235
                                                                                                                                    _ = v235
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                                v236 := GenS35{F0: 42, F1: (v44 ^ v2)}
                                                                                                                switch v161 {
                                                                                                                case 32768:
                                                                                                                    v237 := GenS36{F0: v177, F1: "a"}
                                                                                                                    _ = v237
                                                                                                                }
                                                                                                                if (-651 > -32768) {
                                                                                                                    v238 := GenS25{F0: 100, F1: 100}
                                                                                                                    var v239 float64 = -2.25
                                                                                                                    switch (fi43 ^ (fi42 * -747)) {
                                                                                                                    case 32768:
                                                                                                                        switch fi43 {
                                                                                                                        case 127:
                                                                                                                            var v240 int = (v161 << ((v44) & 63))
                                                                                                                            _ = v240
                                                                                                                        default:
                                                                                                                            v162 += (v162 % ((213) | 1))
                                                                                                                        }
                                                                                                                        var v241 uint16 = 1
                                                                                                                        _ = len(v8)
                                                                                                                        _ = v241
                                                                                                                    case 42:
                                                                                                                        delete(v3, 189)
                                                                                                                    case 2147483647:
                                                                                                                        v161 += v161
                                                                                                                    default:
                                                                                                                        v242 := GenS37{F0: 42, F1: 127, F2: (^(v1))}
                                                                                                                        var v243 int64 = (v9 / ((v171) | 1))
                                                                                                                        _ = v242
                                                                                                                        _ = v243
                                                                                                                    }
                                                                                                                    _ = v238
                                                                                                                    _ = v239
                                                                                                                }
                                                                                                                _ = v236
                                                                                                            }
                                                                                                            v244 := GenS6{F0: v2, F1: v41}
                                                                                                            var v245 int = -596
                                                                                                            delete(v3, 4294967295)
                                                                                                            v246 := GenS24{F0: v1, F1: 42}
                                                                                                            _ = v244
                                                                                                            _ = v245
                                                                                                            _ = v246
                                                                                                        } else {
                                                                                                            if len(v8) > 0 {
                                                                                                                v8[0] = fi43
                                                                                                            }
                                                                                                            v44 = v44
                                                                                                            if true {
                                                                                                                if (v29 || v29) {
                                                                                                                    v247 := GenS38{F0: 65535, F1: 42, F2: 64}
                                                                                                                    v171 = -1
                                                                                                                    _ = v247
                                                                                                                }
                                                                                                            } else {
                                                                                                                if v29 {
                                                                                                                    fmt.Println(v1)
                                                                                                                    switch 238 {
                                                                                                                    case -598:
                                                                                                                        switch ("hello" + "\n\t") {
                                                                                                                        case "":
                                                                                                                            var v248 int = (fi42 - (fi42 >> (((v44 * (v2 - 0))) & 63)))
                                                                                                                            _ = v248
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v249 bool = v29
                                                                                                                    _ = v249
                                                                                                                }
                                                                                                            }
                                                                                                            var v250 int = fi42
                                                                                                            switch fi43 {
                                                                                                            case 127:
                                                                                                                var v251 uint16 = (v177 + 255)
                                                                                                                _ = v251
                                                                                                            case -896:
                                                                                                                fmt.Println(v250)
                                                                                                            }
                                                                                                            _ = v250
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v3[2147483648] = 65535
                                                                                                    }
                                                                                                    _ = v177
                                                                                                case "hello":
                                                                                                    var v252 uint8 = v162
                                                                                                    _ = v252
                                                                                                case "a":
                                                                                                    v2 += 42
                                                                                                }
                                                                                                v29 = (fi42 < (-(v161)))
                                                                                            }
                                                                                            if (genFn0() <= v171) {
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = fi42
                                                                                                }
                                                                                            } else {
                                                                                                v1 = v1
                                                                                            }
                                                                                            _ = v170
                                                                                            _ = v171
                                                                                            _ = v172
                                                                                        } else {
                                                                                            var v253 uint8 = v1
                                                                                            var v254 uint64 = (v41 - (v41 + 152))
                                                                                            v162 = 1
                                                                                            fmt.Println(v2, v161, v9)
                                                                                            v255 := GenS35{F0: -17, F1: 18446744073709551615}
                                                                                            _ = v253
                                                                                            _ = v254
                                                                                            _ = v255
                                                                                        }
                                                                                        switch ("\n\t" + ("a" + ("a" + "a"))) {
                                                                                        case "foo bar":
                                                                                            if len(v8) > 0 {
                                                                                                v8[0] = (fi43 / ((-1) | 1))
                                                                                            }
                                                                                        case "":
                                                                                            v3[2147483648] = 0
                                                                                        case "hello":
                                                                                            var v256 float32 = v4
                                                                                            _ = v256
                                                                                        default:
                                                                                            fmt.Println(v9)
                                                                                        }
                                                                                        _ = v169
                                                                                    }
                                                                                    _ = v162
                                                                                } else {
                                                                                    if (!true) {
                                                                                        _ = len(v8)
                                                                                        if (v29 && (v29 && (!(502 >= 255)))) {
                                                                                            _ = len(v8)
                                                                                            switch 967 {
                                                                                            case 32767:
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = v161
                                                                                                }
                                                                                            case 738:
                                                                                                fmt.Println(v29)
                                                                                            default:
                                                                                                if false {
                                                                                                    var v257 int16 = 1
                                                                                                    _ = v257
                                                                                                } else {
                                                                                                    var v258 float64 = 1e10
                                                                                                    _ = v258
                                                                                                }
                                                                                            }
                                                                                            var v259 string = "a"
                                                                                            _ = v259
                                                                                        }
                                                                                    }
                                                                                    var v260 int32 = -1
                                                                                    _ = v260
                                                                                }
                                                                                if v29 {
                                                                                    fmt.Println(v41, v4, v41)
                                                                                    switch -9223372036854775808 {
                                                                                    case 125:
                                                                                        var v261 bool = v29
                                                                                        var v262 int32 = 577
                                                                                        _ = v261
                                                                                        _ = v262
                                                                                    case 4294967295:
                                                                                        if (422 != int32(v9)) {
                                                                                            var v263 uint8 = (^(v1))
                                                                                            _ = v263
                                                                                        } else {
                                                                                            if (!v29) {
                                                                                                if false {
                                                                                                    var v264 int64 = (v9 * v9)
                                                                                                    _ = len(v3)
                                                                                                    v1 -= (v1 * uint8(v44))
                                                                                                    _ = v264
                                                                                                }
                                                                                            } else {
                                                                                                if ((399 == uint(v1)) && ((fi43 << (((v44 >> ((v2) & 63))) & 63)) <= 405)) {
                                                                                                    var v265 string = ""
                                                                                                    v266 := GenS39{F0: v9, F1: 0, F2: (v2 >> ((v44) & 63))}
                                                                                                    var v267 uint = (v44 ^ 255)
                                                                                                    if v29 {
                                                                                                        if len(v8) > 0 {
                                                                                                            v8[0] = fi43
                                                                                                        }
                                                                                                        var v268 int16 = -1
                                                                                                        _ = v268
                                                                                                    }
                                                                                                    _ = v265
                                                                                                    _ = v266
                                                                                                    _ = v267
                                                                                                }
                                                                                                if false {
                                                                                                    var v269 string = ("" + "a")
                                                                                                    var v270 float32 = 0.1
                                                                                                    switch 32767 {
                                                                                                    case 9223372036854775807:
                                                                                                        if v29 {
                                                                                                            v271 := GenS40{F0: v41}
                                                                                                            delete(v3, 923)
                                                                                                            if ((0 != 42) && v29) {
                                                                                                                if v29 {
                                                                                                                    v44 += 65535
                                                                                                                    if v29 {
                                                                                                                        v41 = v41
                                                                                                                    } else {
                                                                                                                        if (((v269 + (v269 + "")) + (v269 + v269)) <= "") {
                                                                                                                            if true {
                                                                                                                                var v272 int32 = 2147483647
                                                                                                                                fmt.Println(v269)
                                                                                                                                _ = v272
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v273 := GenS25{F0: 127, F1: -1}
                                                                                                                        var v274 int = -1
                                                                                                                        v275 := GenS41{F0: 336, F1: 255, F2: genFn0()}
                                                                                                                        _ = v273
                                                                                                                        _ = v274
                                                                                                                        _ = v275
                                                                                                                    }
                                                                                                                    fmt.Println(v41, fi42, v1)
                                                                                                                }
                                                                                                                v276 := GenS42{F0: (^(v41))}
                                                                                                                _ = v276
                                                                                                            }
                                                                                                            var v277 float32 = 3.14159
                                                                                                            switch "foo bar" {
                                                                                                            case "a":
                                                                                                                if (4294967295 < 994) {
                                                                                                                    v278 := v3[60]
                                                                                                                    v279 := GenS5{F0: (v44 % ((uint(v41)) | 1))}
                                                                                                                    _ = v278
                                                                                                                    _ = v279
                                                                                                                } else {
                                                                                                                    var v280 float64 = (3.14159 / -2.25)
                                                                                                                    switch -204 {
                                                                                                                    case 244:
                                                                                                                        v281 := GenS43{F0: 65535, F1: (v41 + (v41 / ((1) | 1)))}
                                                                                                                        _ = v281
                                                                                                                    case -54:
                                                                                                                        v3[2] = 2
                                                                                                                    case -1:
                                                                                                                        if (v29 || (!false)) {
                                                                                                                            v282 := v3[0]
                                                                                                                            _ = v282
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v283 := GenS43{F0: 4294967295, F1: v41}
                                                                                                                        _ = v283
                                                                                                                    }
                                                                                                                    _ = v280
                                                                                                                }
                                                                                                            case "":
                                                                                                                v284 := GenS27{F0: 32767}
                                                                                                                switch "foo bar" {
                                                                                                                case "\n\t":
                                                                                                                    v44 += v2
                                                                                                                }
                                                                                                                var v285 int64 = v9
                                                                                                                _ = v284
                                                                                                                _ = v285
                                                                                                            default:
                                                                                                                v2 = 255
                                                                                                            }
                                                                                                            fmt.Println(v270)
                                                                                                            _ = v271
                                                                                                            _ = v277
                                                                                                        }
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v269, v9)
                                                                                                    case -136:
                                                                                                        v270 -= v4
                                                                                                    default:
                                                                                                        v4 += v4
                                                                                                    }
                                                                                                    _ = v269
                                                                                                    _ = v270
                                                                                                }
                                                                                            }
                                                                                            v4 = v4
                                                                                            v286 := v3[42]
                                                                                            _ = v286
                                                                                        }
                                                                                    case 867:
                                                                                        var v287 int32 = -2147483648
                                                                                        _ = v287
                                                                                    }
                                                                                    var v288 uint8 = v1
                                                                                    v44 += v2
                                                                                    _ = len(v3)
                                                                                    _ = v288
                                                                                }
                                                                            }
                                                                            fmt.Println(v2, v4, v161)
                                                                            switch "" {
                                                                            case "a":
                                                                                v289 := GenS35{F0: 1, F1: (v2 >> (((v2 / ((v44) | 1))) & 63))}
                                                                                _ = v289
                                                                            case "":
                                                                                v29 = (((v41 / ((v41) | 1)) == (v41 << ((v44) & 63))) && (!true))
                                                                            default:
                                                                                delete(v3, uint32(v2))
                                                                            }
                                                                            var v290 string = "\n\t"
                                                                            v291 := GenS21{F0: 591, F1: false}
                                                                            var v292 int = v161
                                                                            _ = v160
                                                                            _ = v161
                                                                            _ = v290
                                                                            _ = v291
                                                                            _ = v292
                                                                        }
                                                                    } else {
                                                                        v293 := GenS44{F0: v1, F1: -280, F2: 127}
                                                                        switch fi42 {
                                                                        case 0:
                                                                            var v294 int32 = 842
                                                                            _ = v294
                                                                        case 32768:
                                                                            var v295 uint32 = 4294967295
                                                                            _ = len(v8)
                                                                            _ = v295
                                                                        default:
                                                                            if (v4 > (((v4 - float32(1e10)) * (v4 + float32((1e10 + float64(v4))))) - v4)) {
                                                                                v296 := GenS34{F0: v9, F1: 2, F2: ""}
                                                                                switch (fi43 & fi42) {
                                                                                case 735:
                                                                                    var v297 string = ("a" + "")
                                                                                    var v298 string = ("a" + v297)
                                                                                    _ = len(v3)
                                                                                    _ = v297
                                                                                    _ = v298
                                                                                }
                                                                                _ = v296
                                                                            }
                                                                        }
                                                                        switch int(v9) {
                                                                        case -9223372036854775808:
                                                                            if true {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    if true {
                                                                                        if (("\n\t" >= "") && v29) {
                                                                                            v299 := GenS28{F0: v2}
                                                                                            v300 := GenS9{F0: v44, F1: (v1 >> ((255) & 63))}
                                                                                            v9 -= v9
                                                                                            v2 = (v44 % ((v2) | 1))
                                                                                            _ = v299
                                                                                            _ = v300
                                                                                        } else {
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                v41 = uint64(v1)
                                                                                            default:
                                                                                                v4 = v4
                                                                                            }
                                                                                        }
                                                                                        var v301 uint16 = 1
                                                                                        if v29 {
                                                                                            v29 = v29
                                                                                            if ((2147483647 < 0) && (v9 != v9)) {
                                                                                                var v302 uint32 = 454
                                                                                                v303 := GenS28{F0: v2}
                                                                                                _ = v302
                                                                                                _ = v303
                                                                                            } else {
                                                                                                if v29 {
                                                                                                    if v29 {
                                                                                                        var v304 int64 = v9
                                                                                                        _ = v304
                                                                                                    } else {
                                                                                                        v2 -= 951
                                                                                                    }
                                                                                                    v301 -= (v301 + v301)
                                                                                                } else {
                                                                                                    v3[4294967295] = 2147483648
                                                                                                    if len(v8) >= 2 {
                                                                                                        v8 = v8[1:len(v8)]
                                                                                                    }
                                                                                                    var v305 uint8 = 213
                                                                                                    _ = v305
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            delete(v3, 65535)
                                                                                            switch (^(fi42)) {
                                                                                            case 1:
                                                                                                switch fi43 {
                                                                                                case 42:
                                                                                                    if true {
                                                                                                        v306 := GenS13{F0: v2}
                                                                                                        v3[uint32(v1)] = 922
                                                                                                        var v307 int8 = 42
                                                                                                        _ = v306
                                                                                                        _ = v307
                                                                                                    }
                                                                                                case 504:
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        switch fi42 {
                                                                                                        case 127:
                                                                                                            if len(v8) > 0 {
                                                                                                                v8[0] = fi43
                                                                                                            }
                                                                                                            v41 -= v41
                                                                                                            fmt.Println(fi42)
                                                                                                        default:
                                                                                                            v4 = (v4 / v4)
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        if ((377 != uint32(v1)) || (int64(fi43) != 9223372036854775807)) {
                                                                                                            v4 += v4
                                                                                                        }
                                                                                                    case "a":
                                                                                                        switch 795 {
                                                                                                        case 32767:
                                                                                                            if ((fi43 % (((fi42 & 674)) | 1)) > fi43) {
                                                                                                                var v308 uint = v44
                                                                                                                var v309 uint8 = (v1 % ((v1) | 1))
                                                                                                                switch "\n\t" {
                                                                                                                case "foo bar":
                                                                                                                    var v310 float32 = (1.5 - v4)
                                                                                                                    _ = v310
                                                                                                                case "hello":
                                                                                                                    v311 := GenS8{F0: 1, F1: (v41 | (v41 - v41))}
                                                                                                                    _ = v311
                                                                                                                default:
                                                                                                                    switch int(v309) {
                                                                                                                    case 32768:
                                                                                                                        delete(v3, 277)
                                                                                                                    default:
                                                                                                                        v2 += (v44 << (((v44 & 18446744073709551615)) & 63))
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v308
                                                                                                                _ = v309
                                                                                                            } else {
                                                                                                                var v312 int32 = -2147483648
                                                                                                                _ = v312
                                                                                                            }
                                                                                                        case -557:
                                                                                                            v9 += genFn0()
                                                                                                        case 0:
                                                                                                            v3[uint32(v2)] = 157
                                                                                                        default:
                                                                                                            fmt.Println(fi42)
                                                                                                        }
                                                                                                    }
                                                                                                case 1:
                                                                                                    var v313 int16 = -32768
                                                                                                    _ = v313
                                                                                                }
                                                                                            case 128:
                                                                                                if ((v41 ^ (v41 >> ((2) & 63))) > v41) {
                                                                                                    if (!v29) {
                                                                                                        var v314 string = ("a" + "a")
                                                                                                        v315 := GenS45{F0: 152, F1: -137, F2: 4294967295}
                                                                                                        var v316 uint16 = v301
                                                                                                        _ = v314
                                                                                                        _ = v315
                                                                                                        _ = v316
                                                                                                    }
                                                                                                    v1 = 42
                                                                                                    if ((!(v29 || true)) || true) {
                                                                                                        v317 := GenS36{F0: 2, F1: ""}
                                                                                                        switch (("foo bar" + "\n\t") + "") {
                                                                                                        case "hello":
                                                                                                            v318 := GenS19{F0: "", F1: "\n\t"}
                                                                                                            _ = v318
                                                                                                        case "foo bar":
                                                                                                            v1 = (v1 << ((65535) & 63))
                                                                                                        case "a":
                                                                                                            if true {
                                                                                                                var v319 uint16 = v301
                                                                                                                if ((v41 % (((v41 | (^(v41)))) | 1)) < v41) {
                                                                                                                    switch 1 {
                                                                                                                    case -144:
                                                                                                                        v320 := GenS38{F0: 4294967295, F1: 65535, F2: -227}
                                                                                                                        _ = v320
                                                                                                                    default:
                                                                                                                        switch 950 {
                                                                                                                        case 4294967295:
                                                                                                                            if ((v301 << (((v2 + (v44 * (^(v44))))) & 63)) == uint16(v1)) {
                                                                                                                                v319 -= v319
                                                                                                                                var v321 int64 = (v9 + (v9 % (((v9 / (((v9 / (((^(v9))) | 1))) | 1))) | 1)))
                                                                                                                                if len(v8) >= 2 {
                                                                                                                                    v8 = v8[1:len(v8)]
                                                                                                                                }
                                                                                                                                _ = v321
                                                                                                                            } else {
                                                                                                                                if ((v41 > 4294967295) || v29) {
                                                                                                                                    _ = len(v3)
                                                                                                                                    v322 := GenS46{F0: v301, F1: 42}
                                                                                                                                    switch (-(fi42)) {
                                                                                                                                    case 32768:
                                                                                                                                        delete(v3, uint32(v44))
                                                                                                                                    case -1:
                                                                                                                                        v4 = 0.0
                                                                                                                                    default:
                                                                                                                                        if v29 {
                                                                                                                                            v41 += v41
                                                                                                                                            switch (fi42 + -2147483648) {
                                                                                                                                            case 4294967295:
                                                                                                                                                v8 = append(v8, 1)
                                                                                                                                            case 0:
                                                                                                                                                var v323 int64 = (v9 ^ v9)
                                                                                                                                                _ = v323
                                                                                                                                            }
                                                                                                                                            var v324 bool = (false && v29)
                                                                                                                                            _ = v324
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v322
                                                                                                                                }
                                                                                                                                v301 = (^(v301))
                                                                                                                                v325 := GenS47{F0: "foo bar", F1: fi43, F2: 0}
                                                                                                                                var v326 int64 = (v9 * v9)
                                                                                                                                var v327 uint = (v2 | v44)
                                                                                                                                if (uint16(v44) > 2) {
                                                                                                                                    if v29 {
                                                                                                                                        if v29 {
                                                                                                                                            v328 := GenS48{F0: 2147483647}
                                                                                                                                            var v329 int8 = 127
                                                                                                                                            v330 := GenS23{F0: v327}
                                                                                                                                            _ = v328
                                                                                                                                            _ = v329
                                                                                                                                            _ = v330
                                                                                                                                        }
                                                                                                                                        if (!((v301 % (((v319 >> ((18446744073709551615) & 63))) | 1)) != (^(v319)))) {
                                                                                                                                            if len(v8) > 0 {
                                                                                                                                                v8[0] = -9223372036854775808
                                                                                                                                            }
                                                                                                                                            if v29 {
                                                                                                                                                if len(v8) > 0 {
                                                                                                                                                    v8[0] = 32768
                                                                                                                                                }
                                                                                                                                                var v331 int = -1
                                                                                                                                                v331 += 32767
                                                                                                                                                v331 = fi43
                                                                                                                                                _ = v331
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = len(v3)
                                                                                                                                    } else {
                                                                                                                                        v332 := GenS49{F0: "", F1: 851, F2: v41}
                                                                                                                                        var v333 int16 = 273
                                                                                                                                        v9 = -9223372036854775808
                                                                                                                                        _ = v332
                                                                                                                                        _ = v333
                                                                                                                                    }
                                                                                                                                    v334 := GenS17{F0: "hello", F1: uint8(v301)}
                                                                                                                                    _ = v334
                                                                                                                                } else {
                                                                                                                                    var v335 int32 = -2147483648
                                                                                                                                    _ = v335
                                                                                                                                }
                                                                                                                                var v336 bool = true
                                                                                                                                _ = v325
                                                                                                                                _ = v326
                                                                                                                                _ = v327
                                                                                                                                _ = v336
                                                                                                                            }
                                                                                                                        case 436:
                                                                                                                            v2 = (v44 >> ((255) & 63))
                                                                                                                        case 127:
                                                                                                                            var v337 uint = (v2 + (v2 + (v2 + 18446744073709551615)))
                                                                                                                            _ = v337
                                                                                                                        default:
                                                                                                                            v29 = true
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v338 string = ("" + "foo bar")
                                                                                                                    _ = v338
                                                                                                                } else {
                                                                                                                    if v29 {
                                                                                                                        v1 += v1
                                                                                                                        var v339 bool = v29
                                                                                                                        _ = v339
                                                                                                                    } else {
                                                                                                                        var v340 uint16 = (v301 - uint16(v41))
                                                                                                                        v341 := v3[1]
                                                                                                                        _ = v340
                                                                                                                        _ = v341
                                                                                                                    }
                                                                                                                    delete(v3, 1)
                                                                                                                    if false {
                                                                                                                        var v342 uint = uint(v319)
                                                                                                                        _ = v342
                                                                                                                    } else {
                                                                                                                        var v343 int8 = 127
                                                                                                                        var v344 uint32 = uint32(v44)
                                                                                                                        v344 = v344
                                                                                                                        _ = v343
                                                                                                                        _ = v344
                                                                                                                    }
                                                                                                                    fmt.Println(v29, v41)
                                                                                                                }
                                                                                                                _ = v319
                                                                                                            } else {
                                                                                                                fmt.Println(v301)
                                                                                                                v9 -= v9
                                                                                                                switch ((("foo bar" + (("hello" + "foo bar") + ("" + "\n\t"))) + ("a" + "\n\t")) + "a") {
                                                                                                                case "a":
                                                                                                                    var v345 uint8 = (v1 - v1)
                                                                                                                    _ = v345
                                                                                                                case "":
                                                                                                                    v9 -= genFn0()
                                                                                                                default:
                                                                                                                    v1 += v1
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v317
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            var v346 int32 = 42
                                                                                                            v4 += (v4 / float32(0.1))
                                                                                                            _ = v346
                                                                                                        }
                                                                                                        fmt.Println(v9, v44, v2)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v347 int16 = -243
                                                                                            _ = v347
                                                                                        }
                                                                                        _ = v301
                                                                                    } else {
                                                                                        if (((42 > (fi42 / ((-1) | 1))) && (20 <= 0)) || (783 == 42)) {
                                                                                            var v348 uint = (v2 | v44)
                                                                                            v41 -= (v41 * 2)
                                                                                            _ = v348
                                                                                        }
                                                                                        var v349 string = ("foo bar" + "foo bar")
                                                                                        if (!(0 < 2)) {
                                                                                            var v350 int32 = -2147483648
                                                                                            var v351 uint64 = (v41 & (v41 + v41))
                                                                                            switch 1 {
                                                                                            case 127:
                                                                                                v352 := GenS43{F0: 1, F1: v351}
                                                                                                _ = v352
                                                                                            case -2147483648:
                                                                                                if ((v4 < ((((float32(((float64((v4 - 1e10)) * 1.5) * 0.0)) * v4) / float32((3.14159 * 0.1))) - v4) - float32((3.14159 / -2.25)))) && ("hello" >= v349)) {
                                                                                                    fmt.Println(v4, fi42, v29)
                                                                                                    v29 = (42 == v1)
                                                                                                }
                                                                                            }
                                                                                            _ = v350
                                                                                            _ = v351
                                                                                        } else {
                                                                                            var v353 bool = ((2147483647 <= -598) && true)
                                                                                            switch (v349 + (v349 + v349)) {
                                                                                            case "":
                                                                                                var v354 uint8 = v1
                                                                                                _ = v354
                                                                                            }
                                                                                            _ = v353
                                                                                        }
                                                                                        v355 := GenS2{F0: v349}
                                                                                        _ = v349
                                                                                        _ = v355
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -2147483648:
                                                                            v1 -= v1
                                                                        case 936:
                                                                            if (-2.25 >= v4) {
                                                                                v41 += 1
                                                                                v8 = append(v8, -9223372036854775808)
                                                                            } else {
                                                                                switch ("" + "\n\t") {
                                                                                case "\n\t":
                                                                                    if ((v29 || (((v41 >= (v41 << ((v2) & 63))) || true) || true)) || false) {
                                                                                        var v356 float32 = float32(-2.25)
                                                                                        _ = len(v3)
                                                                                        _ = v356
                                                                                    }
                                                                                case "":
                                                                                    v9 += 9223372036854775807
                                                                                case "hello":
                                                                                    v3[0] = 65535
                                                                                default:
                                                                                    v357 := GenS22{F0: (^(v41))}
                                                                                    _ = v357
                                                                                }
                                                                                var v358 string = (("foo bar" + "\n\t") + "")
                                                                                v359 := v3[uint32(v1)]
                                                                                v2 += 42
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = 2147483647
                                                                                }
                                                                                _ = v358
                                                                                _ = v359
                                                                            }
                                                                        default:
                                                                            var v360 uint64 = v41
                                                                            _ = v360
                                                                        }
                                                                        var v361 uint8 = (^(v1))
                                                                        _ = v293
                                                                        _ = v361
                                                                    }
                                                                }
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            }
                                                            var v362 uint64 = v41
                                                            _ = v153
                                                            _ = v362
                                                        } else {
                                                            var v363 int16 = 1000
                                                            _ = v363
                                                        }
                                                    case -2147483648:
                                                        v364 := GenS50{F0: (42 == (v41 * (v41 ^ (^(v41)))))}
                                                        _ = v364
                                                    case 339:
                                                        _ = len(v3)
                                                    }
                                                    if v29 {
                                                        fmt.Println(v2, v4, v44)
                                                        var v365 float32 = float32((1.5 + 0.0))
                                                        fmt.Println(v2)
                                                        _ = v365
                                                    }
                                                }
                                                var v366 float32 = -2.25
                                                _ = v366
                                            }
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(fi42, v9, fi42)
                                            }
                                        }
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        v367 := GenS32{F0: 625}
                                        _ = v367
                                    }
                                    var v368 bool = (!(!(255 >= 215)))
                                    v41 = 2
                                    var v369 uint32 = 0
                                    _ = v368
                                    _ = v369
                                }
                                _ = v44
                            }
                            v370 := v3[0]
                            fmt.Println(v9)
                            _ = fi43
                            _ = v370
                        }
                        for fi371 := 0; fi371 < 5; fi371++ {
                            v372 := GenS51{F0: 4294967295, F1: "a"}
                            fmt.Println(v2, fi42)
                            var v373 int16 = 42
                            _ = fi371
                            _ = v372
                            _ = v373
                        }
                        if true {
                            if v29 {
                                fmt.Println(v1, v1, v29)
                            } else {
                                v4 -= -2.25
                                _ = len(v8)
                            }
                            v374 := GenS33{F0: ("\n\t" + "")}
                            _ = v374
                        } else {
                            switch fi42 {
                            case 32768:
                                if v29 {
                                    v2 += 317
                                    var v375 int64 = int64(v2)
                                    _ = v375
                                }
                            case 42:
                                if (42 <= 1) {
                                    var v376 int = (fi42 % ((fi42) | 1))
                                    if len(v8) > 0 {
                                        v8[0] = (fi42 | (v376 * (v376 >> (((v2 >> (((^(v2))) & 63))) & 63))))
                                    }
                                    var v377 int32 = int32(fi42)
                                    _ = v376
                                    _ = v377
                                } else {
                                    for fi378 := 0; fi378 < 1; fi378++ {
                                        if false {
                                            if (false && false) {
                                                v1 -= v1
                                                fmt.Println(v2, fi42, v41)
                                            } else {
                                                var v379 float64 = 1.5
                                                v29 = v29
                                                _ = v379
                                            }
                                        } else {
                                            v1 += (v1 | v1)
                                            var v380 float64 = float64(v4)
                                            var v381 int16 = int16(v9)
                                            var v382 string = ("hello" + ("a" + "a"))
                                            _ = v380
                                            _ = v381
                                            _ = v382
                                        }
                                        _ = fi378
                                    }
                                    _ = len(v3)
                                }
                            default:
                                switch 0 {
                                case 128:
                                    v29 = (((v41 == (^(v41))) || false) && (v29 || (0 != 2)))
                                default:
                                    v1 = (^(v1))
                                }
                            }
                            v383 := GenS52{F0: 0, F1: 465}
                            switch "a" {
                            case "hello":
                                var v384 uint16 = 15
                                _ = v384
                            case "":
                                if false {
                                    v4 += v4
                                }
                            }
                            _ = v383
                        }
                        _ = fi42
                    }
                }
                v4 -= v4
            } else {
                var v385 int64 = v9
                v386 := v3[2]
                _ = v385
                _ = v386
            }
            _ = v41
        }
        var v387 uint64 = 18446744073709551615
        _ = v387
    }
    for fi388 := 0; fi388 < 4; fi388++ {
        for fi389 := 0; fi389 < 5; fi389++ {
            if len(v8) > 0 {
                v8[0] = fi389
            }
            _ = fi389
        }
        _ = fi388
    }
    v3[0] = 4294967295
    switch 32767 {
    case 1:
        if false {
            v390 := GenS53{F0: 100, F1: 0}
            _ = v390
        }
    default:
        if v29 {
            if (false && v29) {
                if (v29 && (((174 >= v9) && v29) && (v9 == genFn0()))) {
                    if false {
                        if (0.1 != float32((1e10 / 0.1))) {
                            v1 -= 202
                            v1 = (v1 >> ((65535) & 63))
                        }
                        v2 += 958
                        if v29 {
                            fmt.Println(v4)
                        } else {
                            var v391 uint = 1
                            switch 424 {
                            case -9223372036854775808:
                                if ("a" == "a") {
                                    for fi392 := 0; fi392 < 19; fi392++ {
                                        fmt.Println(v29, v2)
                                        var v393 float64 = (1e10 - 0.0)
                                        _ = fi392
                                        _ = v393
                                    }
                                    var v394 uint16 = 396
                                    _ = v394
                                } else {
                                    var v395 uint64 = 18446744073709551615
                                    _ = v395
                                }
                            }
                            v2 = 1
                            _ = v391
                        }
                    }
                } else {
                    v1 -= 111
                }
            }
        }
    }
    if ((42 != 2147483647) || false) {
        if v29 {
            v396 := v3[1]
            var v397 int64 = genFn0()
            _ = v396
            _ = v397
        } else {
            for fi398 := 0; fi398 < 20; fi398++ {
                for fi399 := 0; fi399 < 18; fi399++ {
                    v4 -= 1e10
                    v4 -= v4
                    switch (-(fi398)) {
                    case -819:
                        v400 := v3[1]
                        _ = v400
                    }
                    _ = fi399
                }
                v4 = ((v4 + (v4 * v4)) + 0.1)
                var v401 int = -897
                _ = fi398
                _ = v401
            }
            for fi402 := 0; fi402 < 15; fi402++ {
                var v403 int64 = v9
                for fi404 := 0; fi404 < 7; fi404++ {
                    if (("foo bar" + (("foo bar" + "a") + "foo bar")) <= "foo bar") {
                        switch "" {
                        case "a":
                            var v405 uint = (v2 << ((v2) & 63))
                            _ = v405
                        case "hello":
                            var v406 uint16 = 2
                            _ = v406
                        case "\n\t":
                            fmt.Println(v2)
                        }
                        if v29 {
                            switch fi402 {
                            case -1:
                                var v407 int8 = 127
                                _ = v407
                            case 128:
                                var v408 int8 = 33
                                _ = v408
                            default:
                                v409 := GenS54{F0: -70}
                                _ = v409
                            }
                        }
                        v410 := GenS17{F0: "a", F1: (^(v1))}
                        _ = v410
                    } else {
                        var v411 float32 = float32(-2.25)
                        v2 = v2
                        v9 -= v403
                        _ = v411
                    }
                    _ = fi404
                }
                _ = fi402
                _ = v403
            }
            switch 1 {
            case -2147483648:
                if v29 {
                    var v412 uint = (v2 / ((v2) | 1))
                    _ = v412
                }
            case 777:
                v413 := GenS48{F0: 2147483647}
                _ = v413
            }
        }
        v2 = v2
        v4 = (-2.25 + 1e10)
    } else {
        v29 = v29
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v414 := GenS22{F0: 2}
    v29 = (v1 > uint8(v2))
    var v415 float64 = 1.5
    v29 = true
    for fi416 := 0; fi416 < 20; fi416++ {
        v417 := v3[2147483648]
        if (!false) {
            if ("hello" == "\n\t") {
                var v418 int16 = 0
                _ = v418
            }
            if (!((true && true) && (((!(65535 < 1)) || v29) || (!v29)))) {
                for fi419 := 0; fi419 < 1; fi419++ {
                    _ = len(v8)
                    _ = fi419
                }
            } else {
                for fi420 := 0; fi420 < 12; fi420++ {
                    v421 := GenS11{F0: 435, F1: v1}
                    _ = len(v8)
                    v415 -= (-2.25 + (float64(0.1) * v415))
                    _ = fi420
                    _ = v421
                }
            }
            if false {
                var v422 uint16 = 42
                for fi423 := 0; fi423 < 16; fi423++ {
                    var v424 int32 = 188
                    var v425 uint8 = v1
                    v3[(v417 * v417)] = 65535
                    if v29 {
                        var v426 int64 = v9
                        _ = v426
                    } else {
                        fmt.Println(v422, v415)
                        v29 = ((fi423 * int(v2)) <= fi416)
                        var v427 uint64 = 18446744073709551615
                        v4 -= v4
                        _ = v427
                    }
                    var v428 float64 = v415
                    _ = fi423
                    _ = v424
                    _ = v425
                    _ = v428
                }
                _ = v422
            }
        }
        var v429 int = 9223372036854775807
        _ = fi416
        _ = v417
        _ = v429
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v415:", v415)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v414
    _ = v415
    _ = v8
    _ = v9
}
