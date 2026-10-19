package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int16; F1 bool; F2 bool }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int8; F1 int; F2 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 string; F1 int8; F2 int }
type GenS12 struct { F0 int; F1 uint32; F2 string }
type GenS13 struct { F0 int64; F1 int64; F2 int64 }
type GenS14 struct { F0 int8; F1 uint16 }
type GenS15 struct { F0 int64; F1 uint; F2 int8 }
type GenS16 struct { F0 string; F1 int; F2 string }
type GenS17 struct { F0 int; F1 int64; F2 string }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint8; F1 uint }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint32; F1 int16 }
type GenS22 struct { F0 uint; F1 uint }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 string; F1 uint64; F2 int8 }
type GenS27 struct { F0 int32 }
type GenS28 struct { F0 int8 }
type GenS29 struct { F0 string }
type GenS30 struct { F0 uint32; F1 int64 }
type GenS31 struct { F0 int32; F1 int; F2 uint8 }
type GenS32 struct { F0 string }
type GenS33 struct { F0 uint8; F1 int8; F2 int64 }
type GenS34 struct { F0 uint }
type GenS35 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS36 struct { F0 string }
type GenS37 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS38 struct { F0 uint32 }
type GenS39 struct { F0 uint8 }
type GenS40 struct { F0 int; F1 uint; F2 string }
type GenS41 struct { F0 int; F1 bool }
func genFn0(p0 string, p1 uint) uint8 {
    p1 = (p1 % ((p1) | 1))
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int64 = 9223372036854775807
        p0 = ("foo bar" + p0)
        var v9 int32 = -122
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 uint64 = 676
    var v2 uint = uint(v1)
    var v3 uint32 = 34
    v4 := make(map[string]string)
    var v5 string = "foo bar"
    var v6 int64 = 1
    fmt.Println(v3, v1, v3)
    _ = len(v4)
    if (-179 >= 0) {
        for fi10 := 0; fi10 < 8; fi10++ {
            if (v3 <= (v3 >> ((uint(v3)) & 63))) {
                var v11 float32 = 0.1
                if false {
                    var v12 bool = (850 <= -329)
                    _ = v12
                } else {
                    v11 = float32(3.14159)
                    for fi13 := 0; fi13 < 2; fi13++ {
                        switch "a" {
                        case "foo bar":
                            if false {
                                if (true || (true || true)) {
                                    var v14 bool = (0 == -361)
                                    _ = v14
                                }
                                var v15 float32 = -2.25
                                _ = v15
                            } else {
                                delete(v4, "\n\t")
                                var v16 int = fi10
                                _ = v16
                            }
                        case "":
                            v17 := []int{(fi13 & fi13), fi10, (fi13 << ((4294967295) & 63)), fi13}
                            _ = v17
                        default:
                            if true {
                                var v18 bool = (2 <= 615)
                                if false {
                                    if false {
                                        v18 = v18
                                        v19 := v4[((v5 + (v5 + "foo bar")) + ((v5 + (v5 + v5)) + (v5 + v5)))]
                                        var v20 uint8 = uint8(v2)
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if true {
                                            var v21 bool = (-1 >= 52)
                                            _ = v21
                                        } else {
                                            v1 += (v1 & 42)
                                            v18 = (float32(1.5) <= v11)
                                        }
                                    }
                                    if v18 {
                                        var v22 float64 = (float64(((-2.25 + v11) / 0.1)) * float64(0.0))
                                        fmt.Println(v11)
                                        delete(v4, "foo bar")
                                        switch (fi10 | (fi10 / ((fi10) | 1))) {
                                        case 166:
                                            var v23 int16 = 1000
                                            _ = v23
                                        case 128:
                                            v4["a"] = v5
                                        default:
                                            if (!(!(42 == 32767))) {
                                                var v24 uint16 = 255
                                                var v25 int16 = -32768
                                                v3 += uint32(v1)
                                                _ = v24
                                                _ = v25
                                            } else {
                                                if (!(2 > 2)) {
                                                    if false {
                                                        if v18 {
                                                            var v26 uint32 = 933
                                                            v27 := GenS1{F0: v26, F1: int32(v1)}
                                                            var v28 float64 = 1.5
                                                            fmt.Println(v28)
                                                            _ = v26
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        v11 = float32(0.1)
                                                    } else {
                                                        v22 -= v22
                                                        if v18 {
                                                            var v29 int16 = 32767
                                                            _ = v29
                                                        } else {
                                                            v1 -= (v1 % (((v1 ^ 78)) | 1))
                                                        }
                                                        v5 = v5
                                                        v6 = v6
                                                        var v30 int8 = 66
                                                        v5 = ("" + v5)
                                                        fmt.Println(v11)
                                                        fmt.Println(v2, v11)
                                                        fmt.Println(v2, v1)
                                                        _ = v30
                                                    }
                                                } else {
                                                    var v31 int8 = -128
                                                    if (fi10 < (fi13 + (fi10 >> (((v2 % ((1) | 1))) & 63)))) {
                                                        fmt.Println(v11, v31, v5)
                                                    } else {
                                                        if (!(0 > 566)) {
                                                            fmt.Println(v2, v11)
                                                            v32 := []int{32768, (fi13 / ((fi10) | 1)), 0, (fi13 | fi10)}
                                                            _ = v32
                                                        } else {
                                                            v33 := v4["a"]
                                                            switch "hello" {
                                                            case "hello":
                                                                if (255 > 128) {
                                                                    var v34 int = (fi10 % (((fi13 & (fi10 & (fi10 - fi13)))) | 1))
                                                                    v35 := GenS1{F0: (^(v3)), F1: 0}
                                                                    if true {
                                                                        v1 += v1
                                                                        var v36 int64 = (-(v6))
                                                                        if (!(v18 && (v31 <= (v31 << ((v2) & 63))))) {
                                                                            fmt.Println(v11)
                                                                            v4[(((v33 + ("hello" + v5)) + ("a" + "a")) + "foo bar")] = v33
                                                                            if (42 == 42) {
                                                                                var v37 string = v33
                                                                                var v38 uint32 = 1
                                                                                if v18 {
                                                                                    v39 := GenS1{F0: 42, F1: 1}
                                                                                    v40 := GenS1{F0: v3, F1: -1}
                                                                                    v34 = (fi10 * (v34 & (v34 * 42)))
                                                                                    _ = v39
                                                                                    _ = v40
                                                                                }
                                                                                _ = v37
                                                                                _ = v38
                                                                            } else {
                                                                                if (v18 || ((v2 == (v2 << (((v2 | (^(v2)))) & 63))) && v18)) {
                                                                                    _ = len(v4)
                                                                                    switch (fi13 * (v34 + v34)) {
                                                                                    case -1:
                                                                                        if v18 {
                                                                                            v4[v5] = (v33 + v33)
                                                                                            var v41 int = 0
                                                                                            if v18 {
                                                                                                v3 = v3
                                                                                            }
                                                                                            _ = v41
                                                                                        } else {
                                                                                            var v42 uint = v2
                                                                                            _ = v42
                                                                                        }
                                                                                    default:
                                                                                        var v43 int = (-(fi10))
                                                                                        v44 := GenS2{F0: 32767, F1: (v1 + v1)}
                                                                                        switch v33 {
                                                                                        case "foo bar":
                                                                                            v1 += v1
                                                                                        }
                                                                                        _ = v43
                                                                                        _ = v44
                                                                                    }
                                                                                    if (((1e10 * v22) * v22) >= v22) {
                                                                                        v45 := []int{int(v1), 421}
                                                                                        fmt.Println(fi10, v3)
                                                                                        _ = v45
                                                                                    } else {
                                                                                        v1 -= (v1 << ((2) & 63))
                                                                                    }
                                                                                }
                                                                                if v18 {
                                                                                    if v18 {
                                                                                        v46 := []int{}
                                                                                        _ = v46
                                                                                    }
                                                                                    switch (v33 + ((((v5 + v5) + (v5 + ((("\n\t" + (v33 + ("\n\t" + v33))) + v5) + v5))) + ((v33 + "\n\t") + ("" + "a"))) + "hello")) {
                                                                                    case "hello":
                                                                                        switch v33 {
                                                                                        case "foo bar":
                                                                                            if ((v6 ^ v36) < -432) {
                                                                                                v47 := GenS1{F0: v3, F1: -2147483648}
                                                                                                _ = v47
                                                                                            } else {
                                                                                                var v48 int32 = -69
                                                                                                v22 += (v22 * (v22 * (v22 * (float64(v11) * float64((v11 / (1.5 / 3.14159)))))))
                                                                                                _ = v48
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch (v34 * (fi13 - fi10)) {
                                                                                case -56:
                                                                                    v49 := GenS3{F0: v2, F1: v6, F2: 558}
                                                                                    _ = v49
                                                                                case 32768:
                                                                                    fmt.Println(v36)
                                                                                default:
                                                                                    if true {
                                                                                        v50 := GenS4{F0: 93, F1: (v3 % (((v3 >> ((831) & 63))) | 1)), F2: 128}
                                                                                        _ = v50
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (!true) {
                                                                                var v51 int = 127
                                                                                _ = v51
                                                                            } else {
                                                                                if (42 < 69) {
                                                                                    var v52 uint64 = (v1 + v1)
                                                                                    switch (fi10 / ((fi13) | 1)) {
                                                                                    case -916:
                                                                                        fmt.Println(v33)
                                                                                    case -9223372036854775808:
                                                                                        var v53 uint8 = 128
                                                                                        _ = v53
                                                                                    default:
                                                                                        v54 := GenS5{F0: 1000, F1: v18, F2: (!((v31 + (v31 >> ((v2) & 63))) != v31))}
                                                                                        _ = v54
                                                                                    }
                                                                                    var v55 uint8 = 158
                                                                                    _ = v52
                                                                                    _ = v55
                                                                                } else {
                                                                                    v56 := []int{(^(fi10)), (fi10 / (((fi10 << (((^(v2))) & 63))) | 1)), (fi13 | fi13), -9223372036854775808}
                                                                                    _ = v56
                                                                                }
                                                                                var v57 int32 = 42
                                                                                _ = v57
                                                                            }
                                                                            fmt.Println(v2)
                                                                            switch 2147483647 {
                                                                            case 32768:
                                                                                if v18 {
                                                                                    if v18 {
                                                                                        v58 := GenS6{F0: (v34 * 32768), F1: v18}
                                                                                        switch fi13 {
                                                                                        case 0:
                                                                                            fmt.Println(v18)
                                                                                        case 2147483647:
                                                                                            v22 -= v22
                                                                                        default:
                                                                                            var v59 uint = 2
                                                                                            _ = v59
                                                                                        }
                                                                                        _ = v58
                                                                                    }
                                                                                    var v60 uint = (v2 * v2)
                                                                                    switch (v5 + ("" + "\n\t")) {
                                                                                    case "foo bar":
                                                                                        if (-9223372036854775808 > 127) {
                                                                                            v61 := GenS1{F0: uint32(v1), F1: -49}
                                                                                            _ = v61
                                                                                        }
                                                                                    case "a":
                                                                                        v3 += v3
                                                                                    case "":
                                                                                        if ((v60 >> ((42) & 63)) > (v60 * v60)) {
                                                                                            v11 -= (v11 * (float32((float64(0.1) * float64(((1e10 * ((v11 + v11) / v11)) - v11)))) * (v11 / 1e10)))
                                                                                        } else {
                                                                                            var v62 int64 = (v36 / (((v6 ^ (v6 >> ((v60) & 63)))) | 1))
                                                                                            v1 = v1
                                                                                            v2 += (v2 + v60)
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                    _ = v60
                                                                                } else {
                                                                                    if ((v3 ^ (v3 / ((v3) | 1))) < 1) {
                                                                                        var v63 uint8 = 42
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v64 uint16 = 65535
                                                                                        v65 := GenS5{F0: 734, F1: (1 != 127), F2: ((^(v3)) < (v3 / ((65535) | 1)))}
                                                                                        v66 := GenS7{F0: (v2 & 4294967295), F1: v31}
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v36
                                                                    } else {
                                                                        var v67 int = (fi13 * fi10)
                                                                        v68 := GenS7{F0: v2, F1: (v31 ^ (^(v31)))}
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            default:
                                                                v31 += 100
                                                            }
                                                            _ = v33
                                                        }
                                                        v69 := GenS4{F0: 127, F1: v3, F2: genFn0(v5, v2)}
                                                        v70 := v4["hello"]
                                                        if (-2.25 != float64((((((1e10 * float32(3.14159)) * 1.5) + 0.1) * 3.14159) + v11))) {
                                                            v71 := v4[((v5 + v5) + ("\n\t" + "hello"))]
                                                            _ = v71
                                                        }
                                                        fmt.Println(v18)
                                                        if v18 {
                                                            v72 := GenS8{F0: v31, F1: fi13, F2: (v31 >> ((v2) & 63))}
                                                            _ = v72
                                                        } else {
                                                            v73 := v4["foo bar"]
                                                            v3 = (v3 ^ 729)
                                                            var v74 uint64 = (v1 + (v1 & 2))
                                                            v18 = ((v31 * 1) <= v31)
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        if (v22 >= v22) {
                                                            fmt.Println(v22, v18)
                                                            if (!v18) {
                                                                if v18 {
                                                                    var v75 int8 = (v31 + v31)
                                                                    v76 := GenS7{F0: (v2 ^ v2), F1: int8(v3)}
                                                                    v31 -= (v31 / (((v31 | v75)) | 1))
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                                v77 := []int{(fi10 ^ fi10)}
                                                                _ = v77
                                                            }
                                                        } else {
                                                            v1 = v1
                                                            fmt.Println(v18, v31, v31)
                                                            var v78 uint16 = 42
                                                            _ = v78
                                                        }
                                                        if (v22 != v22) {
                                                            var v79 uint8 = genFn0(("hello" + v70), (v2 << (((^(v2))) & 63)))
                                                            v80 := []int{127, fi10}
                                                            v81 := GenS9{F0: 1}
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        if v18 {
                                                            v82 := GenS6{F0: int(v3), F1: ((-(fi10)) < fi13)}
                                                            var v83 float32 = (float32((v22 - 0.1)) / (((((v11 / (v11 + v11)) - (3.14159 - v11)) / v11) - v11) * v11))
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            var v84 int64 = v6
                                                            if v18 {
                                                                var v85 int32 = -1
                                                                _ = v85
                                                            } else {
                                                                var v86 string = (v70 + ((v5 + ((v70 + "foo bar") + ((v5 + v5) + "foo bar"))) + (("a" + v5) + v5)))
                                                                v22 += 3.14159
                                                                var v87 uint32 = v3
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                            var v88 float32 = (v11 + (v11 - (float32(v22) / float32(v22))))
                                                            if v18 {
                                                                var v89 uint64 = (v1 - 1)
                                                                if ((^(v89)) < (v89 + 644)) {
                                                                    v2 = (v2 ^ 18446744073709551615)
                                                                    v11 = v88
                                                                    v11 += 1.5
                                                                }
                                                                v22 = 1.5
                                                                _ = v89
                                                            }
                                                            _ = v84
                                                            _ = v88
                                                        }
                                                        v6 = 42
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    _ = v31
                                                }
                                                if (((!(162 == (^(v6)))) && (fi13 < -1)) || (v5 >= v5)) {
                                                    v6 = -446
                                                    var v90 int16 = 0
                                                    _ = v90
                                                } else {
                                                    if ((-(v6)) <= v6) {
                                                        v91 := GenS2{F0: 66, F1: 18446744073709551615}
                                                        v92 := GenS5{F0: int16(v2), F1: ((int32(v2) > 0) && false), F2: (1.5 <= v11)}
                                                        if v18 {
                                                            if false {
                                                                v5 = v5
                                                            }
                                                            v93 := []int{128, fi10, -275, (^(fi13))}
                                                            _ = v93
                                                        } else {
                                                            v18 = v18
                                                        }
                                                        v94 := GenS4{F0: 2, F1: (^(v3)), F2: 127}
                                                        v95 := []int{}
                                                        switch fi13 {
                                                        case 127:
                                                            v3 = (v3 % (((v3 / ((1) | 1))) | 1))
                                                            if ((fi10 % ((128) | 1)) == -9223372036854775808) {
                                                                switch fi10 {
                                                                case 0:
                                                                    if v18 {
                                                                        var v96 uint64 = v1
                                                                        if v18 {
                                                                            if (v6 == 9223372036854775807) {
                                                                                switch v5 {
                                                                                case "foo bar":
                                                                                    delete(v4, v5)
                                                                                }
                                                                                var v97 int16 = -537
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        switch v5 {
                                                                        case "":
                                                                            v98 := v4[v5]
                                                                            _ = v98
                                                                        case "hello":
                                                                            switch fi10 {
                                                                            case 127:
                                                                                if len(v95) > 0 {
                                                                                    v95[0] = fi13
                                                                                }
                                                                            default:
                                                                                v2 = v2
                                                                            }
                                                                        default:
                                                                            v22 = 0.1
                                                                        }
                                                                        v6 += -9223372036854775808
                                                                        _ = v96
                                                                    }
                                                                default:
                                                                    var v99 float64 = (3.14159 / v22)
                                                                    _ = v99
                                                                }
                                                                var v100 string = v5
                                                                _ = v100
                                                            } else {
                                                                if true {
                                                                    var v101 int64 = v6
                                                                    _ = v101
                                                                }
                                                                if v18 {
                                                                    if len(v95) >= 2 {
                                                                        v95 = v95[1:len(v95)]
                                                                    }
                                                                    v3 = (v3 % (((^(v3))) | 1))
                                                                }
                                                                var v102 uint64 = (v1 >> ((940) & 63))
                                                                _ = v102
                                                            }
                                                        case 42:
                                                            if ((v2 >> (((v2 >> ((1) & 63))) & 63)) < (v2 * v2)) {
                                                                var v103 uint32 = (v3 << ((2) & 63))
                                                                v3 = (v103 % ((42) | 1))
                                                                switch "a" {
                                                                case "a":
                                                                    var v104 float64 = float64(float32(0.0))
                                                                    _ = v104
                                                                case "\n\t":
                                                                    fmt.Println(v11)
                                                                }
                                                                _ = v103
                                                            }
                                                        default:
                                                            v105 := GenS10{F0: -961}
                                                            v6 = (v6 << ((89) & 63))
                                                            _ = v105
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        var v106 uint8 = genFn0((v5 + v5), (v2 * (v2 % ((v2) | 1))))
                                                        v107 := []int{-794}
                                                        if (!false) {
                                                            fmt.Println(fi10, v2)
                                                            v108 := GenS11{F0: v5, F1: 100, F2: fi10}
                                                            _ = v108
                                                        } else {
                                                            if ((v18 || (v18 || v18)) || (v18 || v18)) {
                                                                v6 += -198
                                                                var v109 float32 = (v11 + ((0.1 + (1e10 + v11)) / v11))
                                                                v18 = ((v18 && (!(fi13 == fi13))) || v18)
                                                                _ = v109
                                                            } else {
                                                                v110 := v4[((((v5 + (v5 + v5)) + "a") + (v5 + "\n\t")) + "a")]
                                                                var v111 int = fi13
                                                                var v112 uint16 = 32767
                                                                _ = v110
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            fmt.Println(fi13, v106)
                                                            switch (-(fi10)) {
                                                            case -286:
                                                                var v113 int32 = 0
                                                                _ = v113
                                                            case 32768:
                                                                if (!v18) {
                                                                    if true {
                                                                        switch (fi10 / (((fi13 + (fi13 ^ fi13))) | 1)) {
                                                                        case 9223372036854775807:
                                                                            v114 := GenS7{F0: v2, F1: 100}
                                                                            _ = v114
                                                                        case 4294967295:
                                                                            v6 = v6
                                                                        }
                                                                        fmt.Println(v11, fi13, v3)
                                                                        if v18 {
                                                                            v4[(((("a" + v5) + ((("hello" + v5) + "\n\t") + v5)) + (v5 + v5)) + "")] = v5
                                                                        } else {
                                                                            if (v18 && (((-1 >= -32768) && v18) || false)) {
                                                                                v115 := GenS5{F0: -32768, F1: ((v106 << ((42) & 63)) == v106), F2: false}
                                                                                if v18 {
                                                                                    if len(v107) >= 2 {
                                                                                        v107 = v107[1:len(v107)]
                                                                                    }
                                                                                }
                                                                                _ = v115
                                                                            } else {
                                                                                var v116 uint16 = 2
                                                                                var v117 int16 = 42
                                                                                var v118 bool = (true && v18)
                                                                                _ = len(v107)
                                                                                switch v5 {
                                                                                case "foo bar":
                                                                                    v2 += (v2 / (((v2 << (((v2 ^ (v2 % ((v2) | 1)))) & 63))) | 1))
                                                                                case "":
                                                                                    v119 := v4[((v5 + v5) + v5)]
                                                                                    _ = v119
                                                                                case "\n\t":
                                                                                    var v120 int64 = v6
                                                                                    _ = v120
                                                                                default:
                                                                                    if v18 {
                                                                                        if v118 {
                                                                                            v106 += v106
                                                                                            v121 := GenS7{F0: v2, F1: -14}
                                                                                            if ((v116 - 32767) == (^(v116))) {
                                                                                                fmt.Println(v106)
                                                                                            } else {
                                                                                                switch 32767 {
                                                                                                case -502:
                                                                                                    var v122 int16 = v117
                                                                                                    _ = v122
                                                                                                default:
                                                                                                    var v123 uint32 = 2147483648
                                                                                                    _ = v123
                                                                                                }
                                                                                                fmt.Println(v18, v18, v1)
                                                                                            }
                                                                                            _ = v121
                                                                                        } else {
                                                                                            fmt.Println(v106)
                                                                                            var v124 float32 = 3.14159
                                                                                            _ = v124
                                                                                        }
                                                                                        if true {
                                                                                            v5 = v5
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v116
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                            fmt.Println(v2, v22, v2)
                                                                        }
                                                                        var v125 string = "foo bar"
                                                                        _ = v125
                                                                    } else {
                                                                        if ((v106 * (v106 + (v106 - v106))) > 128) {
                                                                            if (!true) {
                                                                                if v18 {
                                                                                    var v126 int8 = -108
                                                                                    fmt.Println(v6)
                                                                                    _ = v126
                                                                                }
                                                                                if len(v107) > 0 {
                                                                                    v107[0] = fi13
                                                                                }
                                                                                fmt.Println(v5, fi10, fi13)
                                                                            } else {
                                                                                if v18 {
                                                                                    if v18 {
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            var v127 int = (fi10 % ((fi13) | 1))
                                                                                            _ = v127
                                                                                        case "hello":
                                                                                            v3 -= (v3 - uint32(v2))
                                                                                        }
                                                                                        if (!v18) {
                                                                                            fmt.Println(v5)
                                                                                            if (((v22 - v22) * v22) < (1e10 + v22)) {
                                                                                                if true {
                                                                                                    if true {
                                                                                                        if (!(v18 && v18)) {
                                                                                                            var v128 int64 = (v6 + 9223372036854775807)
                                                                                                            _ = v128
                                                                                                        } else {
                                                                                                            if (v11 == (v11 * v11)) {
                                                                                                                v22 += (1e10 * (v22 * 1.5))
                                                                                                            }
                                                                                                            fmt.Println(v22)
                                                                                                            if v18 {
                                                                                                                switch v5 {
                                                                                                                case "a":
                                                                                                                    v22 = v22
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v129 := GenS12{F0: 937, F1: (v3 % (((v3 & (^(v3)))) | 1)), F2: ("\n\t" + "hello")}
                                                                                                        if ((255 < 255) || true) {
                                                                                                            if v18 {
                                                                                                                fmt.Println(v106, v106, v3)
                                                                                                            } else {
                                                                                                                fmt.Println(v106, fi13)
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            if v18 {
                                                                                                                fmt.Println(v6)
                                                                                                                v130 := GenS12{F0: 4294967295, F1: uint32(v1), F2: v5}
                                                                                                                _ = v130
                                                                                                            }
                                                                                                            v131 := GenS13{F0: v6, F1: (^(v6)), F2: v6}
                                                                                                            _ = v131
                                                                                                        } else {
                                                                                                            var v132 bool = v18
                                                                                                            switch v5 {
                                                                                                            case "":
                                                                                                                switch "foo bar" {
                                                                                                                case "a":
                                                                                                                    v6 += (v6 / (((v6 % ((v6) | 1))) | 1))
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v106, v106, v22)
                                                                                                            }
                                                                                                            _ = v132
                                                                                                        }
                                                                                                        _ = v129
                                                                                                    }
                                                                                                    var v133 int16 = 0
                                                                                                    if true {
                                                                                                        v106 -= 0
                                                                                                        fmt.Println(fi13, v133)
                                                                                                    } else {
                                                                                                        var v134 uint8 = (^(v106))
                                                                                                        _ = v134
                                                                                                    }
                                                                                                    _ = v133
                                                                                                }
                                                                                                v3 -= v3
                                                                                            } else {
                                                                                                v135 := GenS14{F0: -1, F1: 255}
                                                                                                if v18 {
                                                                                                    v6 -= -86
                                                                                                    if v18 {
                                                                                                        v22 -= v22
                                                                                                        v2 = 65535
                                                                                                        if len(v107) > 0 {
                                                                                                            v107[0] = (fi10 >> ((v2) & 63))
                                                                                                        }
                                                                                                    }
                                                                                                    var v136 uint32 = (^(v3))
                                                                                                    _ = v136
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v6)
                                                                                                case "hello":
                                                                                                    var v137 uint8 = v106
                                                                                                    _ = v137
                                                                                                case "foo bar":
                                                                                                    switch (^(fi10)) {
                                                                                                    case 482:
                                                                                                        v107 = append(v107, fi10)
                                                                                                    case -9223372036854775808:
                                                                                                        if (2 >= 106) {
                                                                                                            var v138 int32 = -2147483648
                                                                                                            v3 = v3
                                                                                                            v106 = (v106 + (^(v106)))
                                                                                                            _ = v138
                                                                                                        }
                                                                                                    case 127:
                                                                                                        _ = len(v4)
                                                                                                    default:
                                                                                                        var v139 float64 = v22
                                                                                                        _ = v139
                                                                                                    }
                                                                                                default:
                                                                                                    v140 := GenS15{F0: v6, F1: 1, F2: -128}
                                                                                                    _ = v140
                                                                                                }
                                                                                                _ = v135
                                                                                            }
                                                                                        } else {
                                                                                            var v141 int64 = v6
                                                                                            v11 = float32((v22 + 3.14159))
                                                                                            _ = v141
                                                                                        }
                                                                                        v142 := v4[("foo bar" + v5)]
                                                                                        _ = v142
                                                                                    } else {
                                                                                        v143 := GenS12{F0: (fi10 + (-(fi10))), F1: v3, F2: "foo bar"}
                                                                                        switch (-(fi10)) {
                                                                                        case 42:
                                                                                            v144 := v4[v5]
                                                                                            _ = v144
                                                                                        case 381:
                                                                                            var v145 bool = (false || ((v18 || (true || v18)) || v18))
                                                                                            _ = v145
                                                                                        case 127:
                                                                                            v1 -= v1
                                                                                        default:
                                                                                            v6 = (v6 | v6)
                                                                                        }
                                                                                        if ((true || (v3 >= 42)) || false) {
                                                                                            if false {
                                                                                                if (!(v11 != ((v11 * 0.0) - v11))) {
                                                                                                    var v146 uint32 = 965
                                                                                                    var v147 uint32 = (v146 << ((v2) & 63))
                                                                                                    _ = v146
                                                                                                    _ = v147
                                                                                                }
                                                                                                var v148 int16 = -32768
                                                                                                _ = v148
                                                                                            }
                                                                                            if true {
                                                                                                v149 := GenS1{F0: uint32(v1), F1: -745}
                                                                                                _ = v149
                                                                                            }
                                                                                            v150 := GenS16{F0: v5, F1: fi13, F2: "foo bar"}
                                                                                            _ = v150
                                                                                        }
                                                                                        _ = v143
                                                                                    }
                                                                                    fmt.Println(v2)
                                                                                    v3 += 618
                                                                                    if ((v3 + v3) < (v3 | v3)) {
                                                                                        if (v18 && (!(!(v5 == v5)))) {
                                                                                            switch ("\n\t" + "a") {
                                                                                            case "hello":
                                                                                                fmt.Println(v6, v106, v22)
                                                                                            case "\n\t":
                                                                                                v151 := GenS14{F0: 1, F1: 212}
                                                                                                _ = v151
                                                                                            }
                                                                                            v152 := GenS17{F0: fi10, F1: (v6 * (-(v6))), F2: v5}
                                                                                            _ = v152
                                                                                        } else {
                                                                                            if (v22 != (v22 + -2.25)) {
                                                                                                var v153 float64 = (v22 / v22)
                                                                                                _ = v153
                                                                                            } else {
                                                                                                var v154 bool = v18
                                                                                                switch fi10 {
                                                                                                case -875:
                                                                                                    var v155 int = fi10
                                                                                                    _ = v155
                                                                                                case -1:
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        if (42 == 0) {
                                                                                                            var v156 string = "a"
                                                                                                            v157 := GenS4{F0: genFn0((v156 + "foo bar"), 4294967295), F1: 1, F2: v106}
                                                                                                            var v158 int32 = 0
                                                                                                            _ = v156
                                                                                                            _ = v157
                                                                                                            _ = v158
                                                                                                        }
                                                                                                    default:
                                                                                                        if v154 {
                                                                                                            v159 := GenS18{F0: v154}
                                                                                                            v160 := GenS19{F0: v106, F1: (v2 | (v2 % ((uint(v1)) | 1)))}
                                                                                                            fmt.Println(v11, v6, v106)
                                                                                                            if (!(1 != 115)) {
                                                                                                                fmt.Println(v1, v6, v2)
                                                                                                                v161 := GenS20{F0: (v3 | (v3 + v3))}
                                                                                                                _ = v161
                                                                                                            }
                                                                                                            _ = v159
                                                                                                            _ = v160
                                                                                                        } else {
                                                                                                            switch (fi10 / (((fi13 - (fi10 ^ int(v3)))) | 1)) {
                                                                                                            case -9223372036854775808:
                                                                                                                if false {
                                                                                                                    v162 := GenS13{F0: (v6 / (((v6 >> (((v2 << ((v2) & 63))) & 63))) | 1)), F1: v6, F2: v6}
                                                                                                                    v11 = float32(v22)
                                                                                                                    var v163 int8 = -1
                                                                                                                    v164 := GenS21{F0: 65535, F1: 42}
                                                                                                                    _ = v162
                                                                                                                    _ = v163
                                                                                                                    _ = v164
                                                                                                                } else {
                                                                                                                    switch v5 {
                                                                                                                    case "":
                                                                                                                        if v154 {
                                                                                                                            var v165 int = -255
                                                                                                                            _ = v165
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        v166 := GenS11{F0: ("hello" + ((v5 + ((v5 + (((v5 + "foo bar") + v5) + ((v5 + (v5 + ("\n\t" + v5))) + v5))) + v5)) + (v5 + ("foo bar" + v5)))), F1: -89, F2: fi13}
                                                                                                                        _ = v166
                                                                                                                    case "foo bar":
                                                                                                                        v6 -= (v6 & (^(v6)))
                                                                                                                    default:
                                                                                                                        if (!false) {
                                                                                                                            if ((v2 + v2) >= (v2 & 255)) {
                                                                                                                                v167 := GenS14{F0: 100, F1: 514}
                                                                                                                                var v168 string = v5
                                                                                                                                var v169 uint32 = (v3 ^ (v3 & v3))
                                                                                                                                _ = v167
                                                                                                                                _ = v168
                                                                                                                                _ = v169
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            switch "\n\t" {
                                                                                                                            case "foo bar":
                                                                                                                                v106 = uint8(v2)
                                                                                                                            case "hello":
                                                                                                                                v170 := v4[v5]
                                                                                                                                _ = v170
                                                                                                                            default:
                                                                                                                                var v171 float64 = ((0.0 * (float64(((v11 * -2.25) + v11)) * (-2.25 / (((-2.25 + (1e10 / (v22 + 1e10))) - (-2.25 * 1.5)) - 1.5)))) * (-2.25 - 1.5))
                                                                                                                                _ = v171
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v172 uint64 = v1
                                                                                                                    var v173 uint32 = (v3 / ((v3) | 1))
                                                                                                                    v1 += v1
                                                                                                                    switch 2147483647 {
                                                                                                                    case 32768:
                                                                                                                        v174 := v4[(v5 + ("a" + v5))]
                                                                                                                        _ = v174
                                                                                                                    case 2147483647:
                                                                                                                        v6 += v6
                                                                                                                    }
                                                                                                                    _ = v172
                                                                                                                    _ = v173
                                                                                                                }
                                                                                                            case 127:
                                                                                                                v106 += (v106 | (v106 >> ((917) & 63)))
                                                                                                            default:
                                                                                                                v2 = (v2 ^ (v2 ^ (v2 | (v2 % (((v2 | v2)) | 1)))))
                                                                                                            }
                                                                                                            if false {
                                                                                                                switch (fi13 >> ((4294967295) & 63)) {
                                                                                                                case 4294967295:
                                                                                                                    var v175 uint = 65535
                                                                                                                    _ = v175
                                                                                                                default:
                                                                                                                    var v176 int64 = 247
                                                                                                                    _ = v176
                                                                                                                }
                                                                                                                if (v2 > uint(v1)) {
                                                                                                                    if len(v107) >= 2 {
                                                                                                                        v107 = v107[1:len(v107)]
                                                                                                                    }
                                                                                                                    switch ((v5 + v5) + (((v5 + "foo bar") + v5) + v5)) {
                                                                                                                    case "":
                                                                                                                        v11 -= 1e10
                                                                                                                    case "foo bar":
                                                                                                                        var v177 uint = uint(v3)
                                                                                                                        _ = v177
                                                                                                                    case "a":
                                                                                                                        var v178 float64 = float64(((0.1 / ((v11 / ((float32(v22) * v11) / 1e10)) + (float32(v22) / v11))) / v11))
                                                                                                                        _ = v178
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v179 int32 = -129
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                            }
                                                                                                            var v180 int64 = v6
                                                                                                            _ = v180
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v3, v154)
                                                                                                    var v181 float32 = ((((v11 - 1e10) + v11) / v11) / (v11 - (1.5 - v11)))
                                                                                                    fmt.Println(v106)
                                                                                                    _ = v181
                                                                                                case 763:
                                                                                                    fmt.Println(v6, v2)
                                                                                                }
                                                                                                v1 -= v1
                                                                                                _ = v154
                                                                                            }
                                                                                            var v182 int16 = 42
                                                                                            if true {
                                                                                                var v183 int = (fi10 * -1)
                                                                                                var v184 int32 = 42
                                                                                                if v18 {
                                                                                                    var v185 string = ("\n\t" + v5)
                                                                                                    fmt.Println(v185)
                                                                                                    _ = v185
                                                                                                }
                                                                                                v182 += v182
                                                                                                _ = v183
                                                                                                _ = v184
                                                                                            } else {
                                                                                                var v186 int32 = 0
                                                                                                _ = v186
                                                                                            }
                                                                                            var v187 int8 = 0
                                                                                            _ = v182
                                                                                            _ = v187
                                                                                        }
                                                                                        fmt.Println(v18, v2, v1)
                                                                                        fmt.Println(v106)
                                                                                    } else {
                                                                                        switch v5 {
                                                                                        case "\n\t":
                                                                                            var v188 uint16 = 65535
                                                                                            _ = v188
                                                                                        case "foo bar":
                                                                                            if (((0.1 * 0.1) + (-2.25 * (v22 * float64(v11)))) < float64(v11)) {
                                                                                                v1 = 0
                                                                                                v18 = (!(v6 < (v6 * v6)))
                                                                                            }
                                                                                        case "":
                                                                                            var v189 string = (v5 + (v5 + v5))
                                                                                            _ = v189
                                                                                        }
                                                                                        v2 = 467
                                                                                    }
                                                                                    v190 := GenS22{F0: (v2 | (v2 % ((v2) | 1))), F1: (v2 << ((v2) & 63))}
                                                                                    _ = v190
                                                                                } else {
                                                                                    switch fi13 {
                                                                                    case -390:
                                                                                        fmt.Println(v1, v6)
                                                                                    case -332:
                                                                                        fmt.Println(v1)
                                                                                    case 32767:
                                                                                        v5 = ((v5 + (("\n\t" + v5) + "")) + "hello")
                                                                                    }
                                                                                    v18 = true
                                                                                    switch v5 {
                                                                                    case "foo bar":
                                                                                        if (3.14159 < (v11 + v11)) {
                                                                                            fmt.Println(v5, v22)
                                                                                            if (true || true) {
                                                                                                var v191 uint8 = (v106 / ((235) | 1))
                                                                                                delete(v4, v5)
                                                                                                _ = v191
                                                                                            } else {
                                                                                                var v192 bool = (false && v18)
                                                                                                _ = v192
                                                                                            }
                                                                                            if v18 {
                                                                                                fmt.Println(v22, v5, fi10)
                                                                                                var v193 float64 = (((1.5 - v22) / (float64((0.1 + -2.25)) - v22)) + (v22 * 1e10))
                                                                                                if (32767 == 1) {
                                                                                                    v1 = v1
                                                                                                }
                                                                                                _ = v193
                                                                                            }
                                                                                            if (((^(v106)) <= (^(v106))) || ((!v18) && (v18 || (!(!(!((v5 + "a") > v5))))))) {
                                                                                                v194 := GenS14{F0: -1, F1: 32767}
                                                                                                var v195 uint32 = (v3 - (v3 * (v3 << ((425) & 63))))
                                                                                                _ = v194
                                                                                                _ = v195
                                                                                            }
                                                                                            fmt.Println(v11, v5, fi13)
                                                                                        }
                                                                                        var v196 uint = v2
                                                                                        if ((v106 == 145) || (-2.25 >= -2.25)) {
                                                                                            if len(v107) >= 2 {
                                                                                                v107 = v107[1:len(v107)]
                                                                                            }
                                                                                            var v197 uint32 = 65535
                                                                                            v198 := GenS23{F0: v3}
                                                                                            _ = v197
                                                                                            _ = v198
                                                                                        } else {
                                                                                            v3 = v3
                                                                                            v1 -= (v1 / ((uint64(v106)) | 1))
                                                                                            fmt.Println(v196)
                                                                                            v199 := GenS24{F0: v3}
                                                                                            _ = v199
                                                                                        }
                                                                                        if (v18 && (126 >= -128)) {
                                                                                            var v200 bool = v18
                                                                                            v201 := GenS25{F0: 0}
                                                                                            _ = v200
                                                                                            _ = v201
                                                                                        } else {
                                                                                            fmt.Println(v2, fi13)
                                                                                            v11 = float32((float64(v11) + float64(0.0)))
                                                                                            v107 = append(v107, (-(fi13)))
                                                                                            if false {
                                                                                                if ((v6 >> ((uint(v1)) & 63)) > 0) {
                                                                                                    if (32768 != fi10) {
                                                                                                        var v202 float64 = (((-2.25 + float64(-2.25)) - 0.1) - v22)
                                                                                                        delete(v4, "a")
                                                                                                        _ = v202
                                                                                                    }
                                                                                                } else {
                                                                                                    v22 += (((v22 / v22) + v22) - v22)
                                                                                                }
                                                                                                v203 := GenS1{F0: 2147483648, F1: -1}
                                                                                                if len(v107) > 0 {
                                                                                                    v107[0] = int(v6)
                                                                                                }
                                                                                                _ = v203
                                                                                            }
                                                                                            v6 = (v6 >> (((v2 / (((^(v2))) | 1))) & 63))
                                                                                            v11 += v11
                                                                                            v204 := GenS22{F0: v2, F1: v2}
                                                                                            _ = v204
                                                                                        }
                                                                                        var v205 string = ("foo bar" + v5)
                                                                                        _ = v196
                                                                                        _ = v205
                                                                                    case "hello":
                                                                                        var v206 int8 = 0
                                                                                        _ = v206
                                                                                    default:
                                                                                        if v18 {
                                                                                            v207 := GenS19{F0: uint8(v2), F1: (v2 & v2)}
                                                                                            var v208 int32 = -2147483648
                                                                                            _ = v207
                                                                                            _ = v208
                                                                                        }
                                                                                    }
                                                                                    switch v5 {
                                                                                    case "hello":
                                                                                        var v209 uint = v2
                                                                                        _ = v209
                                                                                    case "foo bar":
                                                                                        if (-32768 == 1) {
                                                                                            v22 -= float64(0.0)
                                                                                            if false {
                                                                                                var v210 uint8 = v106
                                                                                                _ = v210
                                                                                            } else {
                                                                                                var v211 int64 = (v6 << ((v2) & 63))
                                                                                                v212 := GenS17{F0: 42, F1: v6, F2: "hello"}
                                                                                                _ = v211
                                                                                                _ = v212
                                                                                            }
                                                                                            var v213 int64 = (^(v6))
                                                                                            _ = v213
                                                                                        } else {
                                                                                            if (42 > (v1 + v1)) {
                                                                                                var v214 uint = (v2 * uint(v1))
                                                                                                _ = v214
                                                                                            } else {
                                                                                                var v215 int64 = (v6 >> ((v2) & 63))
                                                                                                _ = v215
                                                                                            }
                                                                                            if ((fi10 / (((fi13 / (((fi13 ^ fi10)) | 1))) | 1)) > (fi10 * 1)) {
                                                                                                v3 = v3
                                                                                            } else {
                                                                                                switch 42 {
                                                                                                case 9223372036854775807:
                                                                                                    v106 = uint8(v2)
                                                                                                case -1:
                                                                                                    var v216 bool = (v5 <= "\n\t")
                                                                                                    _ = v216
                                                                                                case 127:
                                                                                                    v106 -= uint8(v1)
                                                                                                    var v217 int = 2147483647
                                                                                                    switch (("" + v5) + (v5 + v5)) {
                                                                                                    case "a":
                                                                                                        v5 = v5
                                                                                                    case "hello":
                                                                                                        v218 := GenS15{F0: v6, F1: (v2 ^ (^(v2))), F2: 84}
                                                                                                        _ = v218
                                                                                                    case "":
                                                                                                        fmt.Println(v3, v11)
                                                                                                        v22 = (v22 / (float64(v11) / float64(v11)))
                                                                                                    }
                                                                                                    fmt.Println(v1)
                                                                                                    _ = v217
                                                                                                default:
                                                                                                    var v219 uint = uint(v1)
                                                                                                    _ = v219
                                                                                                }
                                                                                                v220 := GenS20{F0: (^(v3))}
                                                                                                _ = v220
                                                                                            }
                                                                                            var v221 uint64 = v1
                                                                                            _ = v221
                                                                                        }
                                                                                        v11 = float32(1.5)
                                                                                    default:
                                                                                        v5 = v5
                                                                                    }
                                                                                }
                                                                                if v18 {
                                                                                    v222 := GenS3{F0: 42, F1: 512, F2: -1}
                                                                                    var v223 float64 = v22
                                                                                    _ = v222
                                                                                    _ = v223
                                                                                }
                                                                                var v224 uint64 = (v1 >> (((^(v2))) & 63))
                                                                                _ = v224
                                                                            }
                                                                            if (!true) {
                                                                                v225 := GenS26{F0: ("hello" + v5), F1: 42, F2: 0}
                                                                                v226 := GenS9{F0: 237}
                                                                                v2 -= v2
                                                                                _ = v225
                                                                                _ = v226
                                                                            }
                                                                        } else {
                                                                            switch fi13 {
                                                                            case 1:
                                                                                var v227 uint32 = (v3 % (((v3 ^ (v3 / ((v3) | 1)))) | 1))
                                                                                fmt.Println(v5, v3, v3)
                                                                                _ = v227
                                                                            case 0:
                                                                                v4[v5] = v5
                                                                                if v18 {
                                                                                    v1 = v1
                                                                                } else {
                                                                                    v3 = 4294967295
                                                                                    fmt.Println(fi13, v2)
                                                                                    _ = len(v4)
                                                                                }
                                                                                fmt.Println(v2, v1, v2)
                                                                                if (v11 != 0.1) {
                                                                                    fmt.Println(fi13, v1, fi13)
                                                                                    fmt.Println(v6, v22)
                                                                                    switch (fi10 / ((fi10) | 1)) {
                                                                                    case 32768:
                                                                                        _ = len(v4)
                                                                                    case 0:
                                                                                        v2 = (^(v2))
                                                                                    case 42:
                                                                                        if v18 {
                                                                                            if v18 {
                                                                                                var v228 float32 = v11
                                                                                                if (v18 || true) {
                                                                                                    v229 := GenS14{F0: 127, F1: 230}
                                                                                                    switch v5 {
                                                                                                    case "hello":
                                                                                                        var v230 int16 = -642
                                                                                                        _ = v230
                                                                                                    case "":
                                                                                                        v6 += int64(fi13)
                                                                                                    default:
                                                                                                        fmt.Println(fi13)
                                                                                                    }
                                                                                                    fmt.Println(v228, v3, v106)
                                                                                                    _ = v229
                                                                                                } else {
                                                                                                    v231 := GenS24{F0: 2147483648}
                                                                                                    fmt.Println(v6, v5)
                                                                                                    v22 += 0.0
                                                                                                    switch (v5 + v5) {
                                                                                                    case "":
                                                                                                        if (!true) {
                                                                                                            fmt.Println(v5, v6)
                                                                                                        }
                                                                                                        v232 := GenS27{F0: 42}
                                                                                                        _ = v232
                                                                                                    default:
                                                                                                        fmt.Println(v6, v2, v18)
                                                                                                    }
                                                                                                    switch 9223372036854775807 {
                                                                                                    case -36:
                                                                                                        var v233 uint16 = 1
                                                                                                        _ = v233
                                                                                                    default:
                                                                                                        if v18 {
                                                                                                            switch (v5 + v5) {
                                                                                                            case "a":
                                                                                                                fmt.Println(fi13, fi10)
                                                                                                            case "\n\t":
                                                                                                                if v18 {
                                                                                                                    fmt.Println(v6)
                                                                                                                    if len(v107) > 0 {
                                                                                                                        v107[0] = -1
                                                                                                                    }
                                                                                                                    v234 := GenS28{F0: -67}
                                                                                                                    _ = v234
                                                                                                                }
                                                                                                                v235 := GenS15{F0: v6, F1: (v2 - (v2 | v2)), F2: 114}
                                                                                                                if (!(-1 >= -2147483648)) {
                                                                                                                    switch fi10 {
                                                                                                                    case 42:
                                                                                                                        _ = len(v4)
                                                                                                                        var v236 uint64 = (v1 ^ 914)
                                                                                                                        _ = v236
                                                                                                                    default:
                                                                                                                        v237 := GenS28{F0: 1}
                                                                                                                        _ = v237
                                                                                                                    }
                                                                                                                    _ = len(v4)
                                                                                                                    v2 = (v2 * (v2 ^ v2))
                                                                                                                    switch -766 {
                                                                                                                    case 32768:
                                                                                                                        v11 += float32(v22)
                                                                                                                        if true {
                                                                                                                            switch fi10 {
                                                                                                                            case -361:
                                                                                                                                var v238 float32 = v11
                                                                                                                                _ = v238
                                                                                                                            case 1:
                                                                                                                                if len(v107) >= 2 {
                                                                                                                                    v107 = v107[1:len(v107)]
                                                                                                                                }
                                                                                                                            case 42:
                                                                                                                                v4[v5] = v5
                                                                                                                            default:
                                                                                                                                v228 += 3.14159
                                                                                                                            }
                                                                                                                            fmt.Println(v106, v3)
                                                                                                                            v22 -= (float64(v11) - (1.5 + v22))
                                                                                                                        } else {
                                                                                                                            var v239 uint8 = genFn0("a", 18446744073709551615)
                                                                                                                            v107 = append(v107, (fi13 + (fi10 | int(v2))))
                                                                                                                            if (!v18) {
                                                                                                                                var v240 int16 = 42
                                                                                                                                var v241 int8 = 0
                                                                                                                                _ = v240
                                                                                                                                _ = v241
                                                                                                                            } else {
                                                                                                                                delete(v4, v5)
                                                                                                                            }
                                                                                                                            var v242 float32 = v11
                                                                                                                            v243 := GenS27{F0: -2147483648}
                                                                                                                            v242 -= v242
                                                                                                                            _ = v239
                                                                                                                            _ = v242
                                                                                                                            _ = v243
                                                                                                                        }
                                                                                                                    case -799:
                                                                                                                        if ((65535 == v2) && ("foo bar" == v5)) {
                                                                                                                            v244 := v4[v5]
                                                                                                                            _ = v244
                                                                                                                        } else {
                                                                                                                            var v245 int64 = 42
                                                                                                                            v18 = true
                                                                                                                            _ = v245
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch (fi13 & 2147483647) {
                                                                                                                        case -1:
                                                                                                                            var v246 int32 = 42
                                                                                                                            _ = v246
                                                                                                                        case 42:
                                                                                                                            fmt.Println(v1, fi10)
                                                                                                                        case 757:
                                                                                                                            v247 := GenS12{F0: fi13, F1: 505, F2: "a"}
                                                                                                                            _ = v247
                                                                                                                        default:
                                                                                                                            switch (v5 + v5) {
                                                                                                                            case "\n\t":
                                                                                                                                v248 := GenS29{F0: (("" + "foo bar") + v5)}
                                                                                                                                _ = v248
                                                                                                                            case "":
                                                                                                                                var v249 uint64 = v1
                                                                                                                                _ = v249
                                                                                                                            case "foo bar":
                                                                                                                                v22 -= float64(v228)
                                                                                                                            }
                                                                                                                            fmt.Println(v2, v5, v18)
                                                                                                                            var v250 int8 = 127
                                                                                                                            var v251 uint64 = uint64(v106)
                                                                                                                            v252 := v4[v5]
                                                                                                                            _ = v250
                                                                                                                            _ = v251
                                                                                                                            _ = v252
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v235
                                                                                                            case "":
                                                                                                                v106 += genFn0((v5 + ("" + ("foo bar" + ((v5 + ((v5 + "\n\t") + v5)) + v5)))), (v2 + (v2 - v2)))
                                                                                                            }
                                                                                                            v18 = v18
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v231
                                                                                                }
                                                                                                _ = v228
                                                                                            } else {
                                                                                                v1 -= 0
                                                                                                var v253 float64 = 3.14159
                                                                                                _ = v253
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v254 int64 = v6
                                                                                        _ = v254
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v107 = append(v107, fi10)
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v106, v1)
                                                            }
                                                        }
                                                        v107 = append(v107, fi10)
                                                        v255 := GenS25{F0: fi13}
                                                        _ = v106
                                                        _ = v107
                                                        _ = v255
                                                    }
                                                }
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        delete(v4, "hello")
                                        v256 := GenS11{F0: (("hello" + ((((v5 + "a") + (v5 + (((v5 + "\n\t") + "") + v5))) + "") + ("foo bar" + v5))) + (v5 + "a")), F1: -1, F2: fi13}
                                        var v257 string = v5
                                        fmt.Println(v2, v1)
                                        var v258 uint16 = 427
                                        _ = v256
                                        _ = v257
                                        _ = v258
                                    }
                                    var v259 string = v5
                                    _ = v259
                                } else {
                                    v260 := []int{fi13, 886, fi13}
                                    _ = v260
                                }
                                v4["foo bar"] = (v5 + "foo bar")
                                _ = v18
                            } else {
                                fmt.Println(fi13, v6)
                                switch fi10 {
                                case -474:
                                    v5 = v5
                                case 2147483647:
                                    v11 += 1e10
                                default:
                                    v261 := []int{61, 127, -515}
                                    _ = v261
                                }
                            }
                        }
                        v262 := []int{(-(fi10))}
                        _ = fi13
                        _ = v262
                    }
                }
                switch fi10 {
                case 9223372036854775807:
                    var v263 uint8 = 128
                    _ = v263
                case 127:
                    if false {
                        for fi264 := 0; fi264 < 14; fi264++ {
                            fmt.Println(v11, v5, v5)
                            var v265 uint16 = 1
                            _ = fi264
                            _ = v265
                        }
                        switch fi10 {
                        case 127:
                            v266 := []int{fi10, fi10, (fi10 + -765)}
                            _ = v266
                        }
                        for fi267 := 0; fi267 < 15; fi267++ {
                            if false {
                                switch fi10 {
                                case 9223372036854775807:
                                    v268 := GenS29{F0: "foo bar"}
                                    _ = v268
                                }
                                fmt.Println(fi267, v5)
                                switch v5 {
                                case "a":
                                    if (v3 >= v3) {
                                        v269 := GenS21{F0: v3, F1: 42}
                                        v270 := GenS30{F0: 4294967295, F1: -22}
                                        switch fi10 {
                                        case -938:
                                            v271 := []int{(-(fi10)), (fi267 >> ((uint(v3)) & 63)), (fi10 >> ((uint(v3)) & 63)), fi10}
                                            _ = v271
                                        case 0:
                                            fmt.Println(v1, v11)
                                        }
                                        _ = v269
                                        _ = v270
                                    }
                                case "":
                                    var v272 uint = 2
                                    _ = v272
                                case "hello":
                                    if (42 > 195) {
                                        var v273 string = v5
                                        switch (v5 + v273) {
                                        case "a":
                                            v274 := GenS31{F0: 1, F1: fi267, F2: 255}
                                            _ = v274
                                        case "foo bar":
                                            var v275 float64 = 1.5
                                            v4[("hello" + "a")] = v273
                                            switch (fi267 + (fi267 - fi267)) {
                                            case 42:
                                                if false {
                                                    if (!false) {
                                                        v1 += (v1 >> ((v2) & 63))
                                                    }
                                                }
                                                var v276 uint16 = 65535
                                                _ = v276
                                            case 2147483647:
                                                v2 = (v2 / (((v2 - v2)) | 1))
                                            default:
                                                v1 -= (v1 >> ((42) & 63))
                                            }
                                            _ = v275
                                        case "hello":
                                            v277 := GenS15{F0: 0, F1: v2, F2: -64}
                                            _ = v277
                                        default:
                                            var v278 float32 = ((float32(3.14159) - 1.5) + (v11 - (v11 - ((float32((0.0 + (((3.14159 / ((float64((1.5 + -2.25)) / 1e10) * float64(-2.25))) * 0.0) * 0.0))) + (((v11 / (v11 / -2.25)) * v11) - v11)) * -2.25))))
                                            _ = v278
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            switch (fi267 + (fi267 << ((v2) & 63))) {
                                            case 127:
                                                v279 := v4[v273]
                                                _ = v279
                                            case 2147483647:
                                                if false {
                                                    v3 = v3
                                                    v3 = 0
                                                } else {
                                                    delete(v4, ((v5 + "\n\t") + (v273 + ((v5 + "hello") + (((((v273 + v273) + "\n\t") + v5) + v5) + ("\n\t" + ""))))))
                                                    v280 := GenS32{F0: ("foo bar" + "hello")}
                                                    v2 += (v2 | 4294967295)
                                                    v2 -= (v2 & 12)
                                                    _ = v280
                                                }
                                            default:
                                                v281 := GenS33{F0: 184, F1: 0, F2: v6}
                                                _ = v281
                                            }
                                        case "a":
                                            v282 := v4[("foo bar" + "hello")]
                                            _ = v282
                                        case "foo bar":
                                            var v283 uint = (v2 % ((v2) | 1))
                                            _ = v283
                                        default:
                                            if (-2.25 > ((0.0 - 0.1) + 3.14159)) {
                                                switch v5 {
                                                case "hello":
                                                    switch fi267 {
                                                    case 128:
                                                        if false {
                                                            v3 = (v3 - v3)
                                                        } else {
                                                            var v284 int32 = 2147483647
                                                            switch -9223372036854775808 {
                                                            case -601:
                                                                _ = len(v4)
                                                            case 127:
                                                                if false {
                                                                    v285 := GenS34{F0: 756}
                                                                    if true {
                                                                        v11 += float32((3.14159 - 1e10))
                                                                        v286 := []int{(fi267 << ((268) & 63)), fi267, (fi10 + (fi267 | (fi267 << ((v2) & 63)))), (fi10 << (((v2 & v2)) & 63))}
                                                                        _ = v286
                                                                    } else {
                                                                        var v287 uint32 = v3
                                                                        if (-2.25 < v11) {
                                                                            v284 -= (v284 | (v284 ^ -1))
                                                                        } else {
                                                                            v273 = v5
                                                                            v288 := v4[v5]
                                                                            v273 = v273
                                                                            _ = v288
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v11 = v11
                                                                        case "\n\t":
                                                                            if false {
                                                                                var v289 int8 = 1
                                                                                var v290 uint32 = (^(v287))
                                                                                _ = v289
                                                                                _ = v290
                                                                            }
                                                                        case "hello":
                                                                            switch fi10 {
                                                                            case 4294967295:
                                                                                v291 := GenS26{F0: ("foo bar" + ((v5 + "foo bar") + ("a" + v5))), F1: 891, F2: -23}
                                                                                _ = v291
                                                                            case -2147483648:
                                                                                v2 = v2
                                                                            case 102:
                                                                                var v292 int64 = (v6 << ((2) & 63))
                                                                                v292 = (v292 + 0)
                                                                                _ = v292
                                                                            }
                                                                            v293 := GenS35{F0: (!true), F1: 4294967295, F2: v3}
                                                                            v11 -= v11
                                                                            _ = v293
                                                                        default:
                                                                            v3 += v3
                                                                        }
                                                                        _ = v287
                                                                    }
                                                                    _ = v285
                                                                }
                                                            }
                                                            _ = len(v4)
                                                            v294 := GenS18{F0: true}
                                                            var v295 uint = v2
                                                            _ = v284
                                                            _ = v294
                                                            _ = v295
                                                        }
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        v296 := GenS31{F0: 42, F1: fi10, F2: genFn0((v5 + v273), (^(v2)))}
                                                        _ = v296
                                                    } else {
                                                        v297 := []int{}
                                                        fmt.Println(v11, v3)
                                                        var v298 string = v5
                                                        _ = v297
                                                        _ = v298
                                                    }
                                                default:
                                                    v11 -= 0.1
                                                    switch (v273 + "a") {
                                                    case "foo bar":
                                                        delete(v4, v5)
                                                    case "hello":
                                                        if (fi10 > fi10) {
                                                            v299 := []int{}
                                                            v300 := GenS36{F0: "a"}
                                                            switch fi10 {
                                                            case -2147483648:
                                                                if len(v299) > 0 {
                                                                    v299[0] = fi10
                                                                }
                                                            case -1:
                                                                var v301 uint64 = (v1 & (v1 << ((v2) & 63)))
                                                                _ = v301
                                                            default:
                                                                var v302 uint32 = v3
                                                                _ = v302
                                                            }
                                                            _ = v299
                                                            _ = v300
                                                        }
                                                    case "a":
                                                        v2 = (v2 << (((^(v2))) & 63))
                                                    default:
                                                        var v303 float64 = ((1.5 * 1e10) * 3.14159)
                                                        _ = v303
                                                    }
                                                    v304 := GenS19{F0: 127, F1: (v2 ^ (v2 % ((936) | 1)))}
                                                    _ = v304
                                                }
                                            }
                                        }
                                        _ = v273
                                    } else {
                                        v2 = v2
                                        v305 := GenS11{F0: v5, F1: -124, F2: (fi10 / (((fi10 ^ (fi10 << ((42) & 63)))) | 1))}
                                        _ = v305
                                    }
                                }
                            } else {
                                if false {
                                    v306 := GenS18{F0: ((uint8(v2) >= 255) && (0 < genFn0("a", v2)))}
                                    var v307 int = fi10
                                    _ = v306
                                    _ = v307
                                }
                                if true {
                                    v3 -= 4294967295
                                    var v308 int = (fi267 / (((fi267 << (((^(v2))) & 63))) | 1))
                                    switch -9223372036854775808 {
                                    case 0:
                                        var v309 int8 = 1
                                        _ = v309
                                    }
                                    _ = v308
                                }
                                switch (v5 + ("hello" + "a")) {
                                case "\n\t":
                                    var v310 uint32 = (^(v3))
                                    _ = v310
                                case "a":
                                    if false {
                                        fmt.Println(v5, v1)
                                    } else {
                                        v311 := []int{(fi10 ^ fi267)}
                                        switch ("" + "") {
                                        case "a":
                                            switch (^(fi10)) {
                                            case -9223372036854775808:
                                                if (!(-2.25 > -2.25)) {
                                                    var v312 int64 = (v6 & v6)
                                                    var v313 uint = (v2 >> (((v2 << ((v2) & 63))) & 63))
                                                    v314 := GenS37{F0: 619, F1: 42, F2: 41}
                                                    _ = v312
                                                    _ = v313
                                                    _ = v314
                                                } else {
                                                    switch 32767 {
                                                    case 0:
                                                        switch fi10 {
                                                        case 128:
                                                            v315 := v4[(v5 + v5)]
                                                            _ = v315
                                                        }
                                                    case 128:
                                                        var v316 uint16 = uint16(v3)
                                                        _ = v316
                                                    case -1:
                                                        v6 = 0
                                                    default:
                                                        v1 += 2
                                                    }
                                                    v11 -= (v11 / float32(3.14159))
                                                    if len(v311) >= 2 {
                                                        v311 = v311[1:len(v311)]
                                                    }
                                                }
                                                v311 = append(v311, fi267)
                                                _ = len(v311)
                                            case -2147483648:
                                                v1 = 4294967295
                                            case 32767:
                                                if ((v2 / (((v2 >> (((v2 | uint(v3))) & 63))) | 1)) > v2) {
                                                    fmt.Println(v1, v3)
                                                }
                                            }
                                        case "hello":
                                            v317 := GenS38{F0: v3}
                                            _ = v317
                                        case "foo bar":
                                            if (true || false) {
                                                v318 := GenS39{F0: 255}
                                                if len(v311) >= 2 {
                                                    v311 = v311[1:len(v311)]
                                                }
                                                v5 = (v5 + "")
                                                _ = v318
                                            } else {
                                                v4[v5] = "a"
                                            }
                                        default:
                                            switch (fi267 | (fi267 * fi267)) {
                                            case 56:
                                                v311 = append(v311, fi10)
                                            default:
                                                switch (fi10 / ((433) | 1)) {
                                                case 32768:
                                                    delete(v4, (("hello" + "") + "\n\t"))
                                                default:
                                                    var v319 int64 = -1
                                                    _ = v319
                                                }
                                            }
                                        }
                                        if (((float32(-2.25) * -2.25) - v11) >= 1.5) {
                                            var v320 int = (fi267 << (((v2 / ((v2) | 1))) & 63))
                                            _ = v320
                                        }
                                        _ = v311
                                    }
                                default:
                                    v3 = (v3 ^ v3)
                                }
                            }
                            v6 = v6
                            var v321 float64 = -2.25
                            _ = fi267
                            _ = v321
                        }
                    } else {
                        v322 := GenS40{F0: 128, F1: v2, F2: v5}
                        if ((v1 / ((v1) | 1)) == v1) {
                            var v323 uint64 = v1
                            for fi324 := 0; fi324 < 20; fi324++ {
                                v325 := []int{(fi324 + fi10), (^(fi324)), 42, fi324}
                                var v326 bool = ((v2 + uint(v323)) != v2)
                                var v327 int64 = v6
                                v328 := GenS41{F0: int(v1), F1: (0 >= 0)}
                                _ = fi324
                                _ = v325
                                _ = v326
                                _ = v327
                                _ = v328
                            }
                            _ = v323
                        } else {
                            var v329 float32 = v11
                            v6 -= (v6 << ((864) & 63))
                            _ = v329
                        }
                        var v330 uint32 = 65535
                        _ = v322
                        _ = v330
                    }
                case 128:
                    if false {
                        var v331 int = int(v1)
                        if true {
                            for fi332 := 0; fi332 < 3; fi332++ {
                                var v333 float32 = float32(1.5)
                                _ = fi332
                                _ = v333
                            }
                            fmt.Println(v6)
                            var v334 string = v5
                            _ = v334
                        }
                        _ = v331
                    }
                default:
                    if false {
                        for fi335 := 0; fi335 < 2; fi335++ {
                            if true {
                                switch -233 {
                                case 2147483647:
                                    v6 = 709
                                case 32767:
                                    v11 += 0.1
                                }
                                v3 += (v3 + (v3 + v3))
                            } else {
                                v336 := GenS13{F0: (v6 - (^(v6))), F1: 1, F2: (v6 * (v6 - -985))}
                                var v337 uint16 = 2
                                var v338 int16 = 1
                                fmt.Println(v3, v338)
                                _ = v336
                                _ = v337
                                _ = v338
                            }
                            _ = fi335
                        }
                        if true {
                            v339 := GenS3{F0: (^(v2)), F1: v6, F2: 80}
                            v6 = v6
                            var v340 int64 = v6
                            _ = v339
                            _ = v340
                        }
                        v341 := []int{fi10, (fi10 % ((fi10) | 1)), (fi10 + fi10)}
                        _ = v341
                    } else {
                        v2 += uint(v1)
                    }
                }
                _ = v11
            }
            v342 := []int{(-(fi10))}
            v5 = v5
            _ = fi10
            _ = v342
        }
        v5 = (v5 + "\n\t")
        var v343 string = "a"
        _ = v343
    }
    v344 := v4[v5]
    v1 += (v1 & 114)
    var v345 int = 42
    switch v344 {
    case "a":
        switch 42 {
        case 0:
            var v346 float32 = 3.14159
            _ = v346
        case 32767:
            switch "a" {
            case "a":
                switch v5 {
                case "\n\t":
                    var v347 int8 = 1
                    _ = v347
                case "hello":
                    v345 -= v345
                }
            default:
                switch v345 {
                case 719:
                    if (!false) {
                        var v348 uint64 = (v1 + (v1 / ((18446744073709551615) | 1)))
                        var v349 int64 = v6
                        v345 = -6
                        _ = v348
                        _ = v349
                    }
                case 2147483647:
                    v350 := GenS21{F0: v3, F1: -1}
                    _ = v350
                case 128:
                    var v351 int16 = 0
                    _ = v351
                }
            }
        }
    case "hello":
        _ = len(v4)
    default:
        var v352 int64 = (v6 + -1)
        _ = v352
    }
    if true {
        var v353 uint64 = (v1 + v1)
        if (0 != -1) {
            fmt.Println(v344)
            v354 := []int{v345, v345}
            _ = v354
        }
        fmt.Println(v344, v6)
        _ = v353
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v344:", v344)
    fmt.Println("v345:", v345)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v344
    _ = v345
    _ = v4
    _ = v5
    _ = v6
}
