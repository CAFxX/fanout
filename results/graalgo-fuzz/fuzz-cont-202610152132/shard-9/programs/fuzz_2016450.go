package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int64; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 int16; F1 uint16; F2 int64 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 uint64; F1 uint32 }
type GenS15 struct { F0 uint8; F1 uint }
type GenS16 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 uint16, p1 int32, p2 int8) uint8 {
    fmt.Println(p0, p1)
    p0 = p0
    fmt.Println(p2)
    v5 := GenS1{F0: 9223372036854775807, F1: 65535, F2: (p1 + -937)}
    v6 := GenS1{F0: 0, F1: 42, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1() int8 {
    v7 := GenS2{F0: 0, F1: 0}
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 18; fi9++ {
        if (18446744073709551615 > v8) {
            switch "" {
            case "hello":
                if true {
                    var v10 int32 = 1
                    v11 := []int{(-(fi9)), 9223372036854775807, fi9}
                    v12 := GenS2{F0: (v10 ^ v10), F1: 1}
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case "":
                fmt.Println(v8)
            }
            v8 -= v8
        } else {
            v13 := []int{fi9, fi9, fi9}
            v8 = v8
            for fi14 := 0; fi14 < 12; fi14++ {
                v8 += v8
                _ = fi14
            }
            _ = v13
        }
        _ = fi9
    }
    var v15 int16 = 0
    var v16 int32 = 580
    _ = v15
    _ = v16
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    fmt.Println(v1)
    v17 := []int{-660, int(v1), -911}
    fmt.Println(v1)
    var v18 uint32 = 46
    v18 -= (^(v18))
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS3{F0: 377, F1: 692, F2: "hello"}
        _ = fi19
        _ = v20
    }
    var v21 float64 = (0.0 / 1.5)
    v22 := GenS2{F0: 1, F1: 42}
    if (125 <= 18446744073709551615) {
        if (!true) {
            fmt.Println(v18, v21, v18)
        } else {
            v17 = append(v17, 0)
            v23 := GenS1{F0: 0, F1: (v18 & (v18 & v18)), F2: -744}
            _ = v23
        }
        _ = len(v17)
        var v24 uint = 253
        _ = v24
    }
    fmt.Println(v18, v1)
    var v25 int64 = -9223372036854775808
    if true {
        if true {
            fmt.Println(v1, v25, v1)
            delete(v2, (^(v18)))
            v26 := GenS4{F0: -111, F1: 32768}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 18; fi27++ {
                fmt.Println(v1, v1, v21)
                switch (("foo bar" + (("foo bar" + "") + "foo bar")) + "a") {
                case "foo bar":
                    v28 := GenS3{F0: 255, F1: 2, F2: "a"}
                    _ = v28
                case "\n\t":
                    switch fi27 {
                    case 9223372036854775807:
                        fmt.Println(fi27)
                    case 127:
                        if (127 > 0) {
                            var v29 float32 = 3.14159
                            if (266 >= 1) {
                                var v30 int64 = v25
                                var v31 uint32 = (v18 << ((0) & 63))
                                var v32 uint16 = 534
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                if false {
                                    v33 := GenS3{F0: 65535, F1: 42, F2: "\n\t"}
                                    for fi34 := 0; fi34 < 3; fi34++ {
                                        v35 := GenS3{F0: 520, F1: 533, F2: "a"}
                                        if ((v25 * v25) <= int64(fi34)) {
                                            fmt.Println(v29, fi34, fi27)
                                        } else {
                                            v29 = (float32(v21) * v29)
                                        }
                                        v1 = (v1 | (v1 >> ((18446744073709551615) & 63)))
                                        _ = fi34
                                        _ = v35
                                    }
                                    if false {
                                        if false {
                                            switch 0 {
                                            case 2147483647:
                                                v25 += (v25 + (v25 - (v25 & -9223372036854775808)))
                                            }
                                        }
                                        for fi36 := 0; fi36 < 18; fi36++ {
                                            var v37 int32 = (-(v1))
                                            _ = fi36
                                            _ = v37
                                        }
                                        var v38 int8 = -113
                                        _ = v38
                                    }
                                    _ = v33
                                }
                                var v39 int32 = v1
                                _ = v39
                            }
                            var v40 uint = 18446744073709551615
                            _ = v29
                            _ = v40
                        } else {
                            fmt.Println(fi27, v25)
                            for fi41 := 0; fi41 < 13; fi41++ {
                                switch "hello" {
                                case "\n\t":
                                    v42 := v2[2147483648]
                                    if true {
                                        _ = len(v17)
                                    } else {
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        delete(v4, (fi41 | (^(fi41))))
                                    }
                                    if true {
                                        v18 -= 2147483648
                                    }
                                    _ = v42
                                case "hello":
                                    v3["hello"] = fi41
                                    v18 = (v18 % ((0) | 1))
                                    if (fi41 > 32768) {
                                        var v43 int8 = 74
                                        v44 := GenS3{F0: 1, F1: 2, F2: "hello"}
                                        v45 := v4[(fi27 * fi41)]
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if true {
                                            v25 = v25
                                            v46 := v2[(v18 >> ((2) & 63))]
                                            _ = v46
                                        } else {
                                            if (((("\n\t" + "a") + "hello") <= ("hello" + "a")) || ((1e10 > (0.1 * 1e10)) && true)) {
                                                v47 := GenS5{F0: 2}
                                                if false {
                                                    var v48 int32 = (v1 & -185)
                                                    var v49 float64 = 0.1
                                                    v25 = (v25 * (^(v25)))
                                                    v18 = (v18 | 4294967295)
                                                    var v50 int64 = (v25 * -225)
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    var v51 uint16 = 2
                                                    fmt.Println(v25, v21, fi41)
                                                    fmt.Println(fi41)
                                                    _ = v51
                                                }
                                                v52 := v4[(-(fi27))]
                                                _ = v47
                                                _ = v52
                                            } else {
                                                v1 = (v1 | v1)
                                            }
                                            delete(v3, "")
                                            v1 = v1
                                        }
                                    }
                                default:
                                    var v53 int32 = 0
                                    _ = v53
                                }
                                var v54 int32 = v1
                                _ = fi41
                                _ = v54
                            }
                        }
                    case 2147483647:
                        var v55 int16 = 0
                        _ = v55
                    default:
                        v56 := GenS6{F0: 2147483647, F1: v25, F2: fi27}
                        _ = v56
                    }
                case "a":
                    v57 := GenS1{F0: (fi27 + 4294967295), F1: (v18 % (((v18 / ((v18) | 1))) | 1)), F2: (v1 << ((65535) & 63))}
                    _ = v57
                default:
                    if (!true) {
                        if false {
                            switch 2147483647 {
                            case 1:
                                for fi58 := 0; fi58 < 1; fi58++ {
                                    var v59 int16 = 834
                                    _ = fi58
                                    _ = v59
                                }
                            }
                            fmt.Println(v1, v1)
                            if false {
                                var v60 int8 = 100
                                var v61 uint64 = 42
                                _ = v60
                                _ = v61
                            } else {
                                var v62 uint16 = 0
                                v25 = v25
                                _ = v62
                            }
                        } else {
                            v25 -= (-(v25))
                            v18 += v18
                            if true {
                                if false {
                                    for fi63 := 0; fi63 < 2; fi63++ {
                                        switch 4294967295 {
                                        case 128:
                                            var v64 uint16 = 834
                                            if (65535 <= v18) {
                                                var v65 bool = true
                                                if (false && v65) {
                                                    var v66 int8 = 1
                                                    _ = v66
                                                } else {
                                                    var v67 float32 = (3.14159 - (0.0 * -2.25))
                                                    v68 := GenS7{F0: 235}
                                                    _ = v67
                                                    _ = v68
                                                }
                                                var v69 string = ""
                                                _ = v65
                                                _ = v69
                                            }
                                            v70 := v3["hello"]
                                            _ = v64
                                            _ = v70
                                        case -689:
                                            switch fi63 {
                                            case -1:
                                                switch "\n\t" {
                                                case "":
                                                    var v71 int32 = v1
                                                    _ = v71
                                                case "foo bar":
                                                    v3["hello"] = fi27
                                                default:
                                                    if false {
                                                        v72 := GenS8{F0: 18446744073709551615}
                                                        var v73 int32 = (v1 + 42)
                                                        var v74 int8 = genFn1()
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        switch fi63 {
                                                        case 128:
                                                            v75 := GenS9{F0: 411, F1: "hello"}
                                                            _ = v75
                                                        case 32768:
                                                            if (!true) {
                                                                v18 += v18
                                                            }
                                                        }
                                                    }
                                                }
                                            case 891:
                                                switch "foo bar" {
                                                case "hello":
                                                    v76 := GenS7{F0: 1}
                                                    _ = v76
                                                case "foo bar":
                                                    var v77 bool = true
                                                    _ = v77
                                                case "":
                                                    var v78 int64 = (-(v25))
                                                    _ = v78
                                                }
                                            case -487:
                                                if true {
                                                    v79 := GenS10{F0: 32767, F1: (fi27 + fi63)}
                                                    v1 += 0
                                                    v2[v18] = 537
                                                    _ = v79
                                                }
                                            default:
                                                var v80 int64 = (v25 / ((int64(v1)) | 1))
                                                v80 -= (v80 + v80)
                                                v81 := GenS11{F0: 0, F1: 42, F2: v80}
                                                _ = v80
                                                _ = v81
                                            }
                                        default:
                                            v18 -= 42
                                        }
                                        fmt.Println(fi63, v25, v1)
                                        _ = fi63
                                    }
                                    for fi82 := 0; fi82 < 11; fi82++ {
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        _ = fi82
                                    }
                                } else {
                                    var v83 int = fi27
                                    var v84 uint = 42
                                    _ = v83
                                    _ = v84
                                }
                            } else {
                                for fi85 := 0; fi85 < 8; fi85++ {
                                    if (-2.25 > ((0.1 + 3.14159) / 0.1)) {
                                        v86 := GenS12{F0: 255, F1: 2}
                                        v87 := GenS10{F0: 665, F1: (fi27 % (((fi27 + 32767)) | 1))}
                                        v1 = (v1 & v1)
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v1 -= v1
                                    }
                                    v25 -= (v25 >> ((255) & 63))
                                    switch fi85 {
                                    case 0:
                                        var v88 float32 = 3.14159
                                        var v89 bool = ((!((v21 * 1e10) <= (((v21 - v21) + (float64(1.5) * float64((((3.14159 * v88) * ((3.14159 * (float32(float64((float32(1e10) - v88))) * ((-2.25 * (1e10 / 1e10)) - float32(float64(float32(v21)))))) / v88)) / v88)))) - v21))) && ((!(4294967295 >= 2)) && (fi85 < fi27)))
                                        _ = v88
                                        _ = v89
                                    case 753:
                                        v18 = v18
                                    case -1:
                                        if ((fi85 % ((fi85) | 1)) == fi85) {
                                            v25 += v25
                                            fmt.Println(fi27, fi27, v21)
                                            v90 := v2[(v18 / ((v18) | 1))]
                                            _ = v90
                                        }
                                    }
                                    _ = fi85
                                }
                                for fi91 := 0; fi91 < 8; fi91++ {
                                    v18 += (v18 & v18)
                                    _ = fi91
                                }
                                v1 = (v1 << ((840) & 63))
                            }
                        }
                        v1 = v1
                    } else {
                        if (1 < -780) {
                            v92 := GenS13{F0: (^(fi27)), F1: 128}
                            if (float32(v21) == 1e10) {
                                v3["\n\t"] = fi27
                            } else {
                                switch ("a" + "a") {
                                case "\n\t":
                                    if ("a" > "a") {
                                        var v93 float32 = (((float32(v21) - 0.0) - 1.5) + ((((float32(1e10) - float32(1e10)) + (3.14159 + float32(v21))) - (1e10 * float32(v21))) + ((1e10 - 0.1) - (-2.25 + ((((1.5 * ((float32(v21) + ((0.0 + -2.25) - (float32(0.1) / (0.1 + float32(v21))))) / 1.0)) + ((3.14159 / 3.14159) / 1e10)) - 1.5) * 1e10)))))
                                        var v94 int16 = 32767
                                        _ = v93
                                        _ = v94
                                    } else {
                                        if false {
                                            switch "a" {
                                            case "a":
                                                switch "a" {
                                                case "a":
                                                    if (true && true) {
                                                        v95 := GenS9{F0: 157, F1: ("\n\t" + "")}
                                                        if (!false) {
                                                            if len(v17) > 0 {
                                                                v17[0] = fi27
                                                            }
                                                            var v96 int8 = -1
                                                            _ = v96
                                                        }
                                                        _ = v95
                                                    }
                                                case "":
                                                    fmt.Println(v25, v25, v25)
                                                case "\n\t":
                                                    v97 := GenS14{F0: 119, F1: 2}
                                                    _ = v97
                                                default:
                                                    fmt.Println(v25, v25)
                                                }
                                            case "":
                                                fmt.Println(fi27, v18)
                                            default:
                                                fmt.Println(v18, v18, v25)
                                            }
                                        } else {
                                            switch fi27 {
                                            case 127:
                                                for fi98 := 0; fi98 < 20; fi98++ {
                                                    if false {
                                                        var v99 int = (fi98 + fi98)
                                                        v99 += (fi27 ^ fi98)
                                                        v25 += v25
                                                        if false {
                                                            if len(v17) > 0 {
                                                                v17[0] = (fi27 | 9223372036854775807)
                                                            }
                                                            v100 := GenS15{F0: 1, F1: 4294967295}
                                                            v21 = v21
                                                            _ = v100
                                                        } else {
                                                            var v101 uint8 = uint8(v18)
                                                            _ = v101
                                                        }
                                                        if len(v17) >= 2 {
                                                            v17 = v17[1:len(v17)]
                                                        }
                                                        _ = v99
                                                    }
                                                    v18 += 65535
                                                    _ = fi98
                                                }
                                            }
                                            fmt.Println(v25)
                                            v21 += 0.1
                                        }
                                    }
                                case "":
                                    fmt.Println(v25, v21)
                                case "hello":
                                    for fi102 := 0; fi102 < 5; fi102++ {
                                        var v103 int32 = 0
                                        var v104 uint32 = v18
                                        if true {
                                            v1 = v103
                                            switch (-(fi102)) {
                                            case 169:
                                                v105 := GenS10{F0: 32767, F1: (fi102 >> ((255) & 63))}
                                                _ = v105
                                            }
                                            var v106 int64 = (v25 & 614)
                                            _ = v106
                                        }
                                        v3[("a" + "foo bar")] = 32767
                                        _ = fi102
                                        _ = v103
                                        _ = v104
                                    }
                                default:
                                    var v107 int32 = v1
                                    _ = v107
                                }
                                v108 := GenS4{F0: 1, F1: fi27}
                                _ = v108
                            }
                            _ = v92
                        } else {
                            var v109 float64 = (-2.25 + (1e10 * 0.1))
                            _ = v109
                        }
                    }
                }
                _ = fi27
            }
        }
        if true {
            for fi110 := 0; fi110 < 11; fi110++ {
                v21 -= (v21 / float64(1.5))
                _ = fi110
            }
            v111 := GenS16{F0: (!(false || true)), F1: 42, F2: 128}
            v112 := GenS14{F0: 572, F1: (^(v18))}
            _ = v111
            _ = v112
        }
    } else {
        v1 += v1
    }
    v1 += (v1 - (v1 | v1))
    fmt.Println(v18, v25, v1)
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
