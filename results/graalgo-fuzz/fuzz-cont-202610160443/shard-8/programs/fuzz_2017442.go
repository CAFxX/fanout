package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int; F1 uint64; F2 int }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 bool; F1 bool; F2 int32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint8; F1 bool; F2 int16 }
type GenS16 struct { F0 uint64; F1 uint8 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 int64, p1 uint32, p2 uint32) int64 {
    p2 = p2
    switch "foo bar" {
    case "\n\t":
        if false {
            var v4 uint16 = 255
            _ = v4
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1(p0 int16, p1 int64) uint64 {
    v5 := GenS1{F0: p0}
    v6 := GenS2{F0: true, F1: 1}
    var v7 int16 = 291
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = p0
            if false {
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 int64 = (p1 & p1)
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := []int{-194, 2147483647, fi8}
                _ = v12
            }
            var v13 float64 = 1e10
            _ = fi8
            _ = v13
            _ = v9
        }
    default:
        var v14 uint16 = 255
        _ = v14
    }
    var v15 uint8 = 8
    _ = p0
    _ = p1
    _ = v15
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint = 594
    var v2 string = ("a" + "")
    var v3 float64 = (0.0 / 3.14159)
    if (42 < 1) {
        fmt.Println(v1)
    } else {
        var v16 float32 = 3.14159
        _ = v16
    }
    v2 = "a"
    v17 := []int{}
    var v18 int16 = -712
    for fi19 := 0; fi19 < 18; fi19++ {
        if ((!false) && (v2 < v2)) {
            v1 = 0
            var v20 string = ("\n\t" + ((v2 + ("foo bar" + v2)) + v2))
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                if true {
                    v22 := GenS1{F0: (v18 / ((v18) | 1))}
                    v17 = append(v17, (fi21 & (-(fi19))))
                    switch v2 {
                    case "foo bar":
                        fmt.Println(v1, v1)
                    case "a":
                        if false {
                            var v23 uint16 = 2
                            switch (v2 + v2) {
                            case "":
                                v24 := GenS3{F0: fi21, F1: 65535, F2: -532}
                                _ = v24
                            case "foo bar":
                                var v25 int16 = (v18 - v18)
                                _ = v25
                            case "hello":
                                if false {
                                    switch fi21 {
                                    case 9223372036854775807:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    case 2147483647:
                                        v1 = (v1 - (v1 << ((666) & 63)))
                                        fmt.Println(fi21, v23)
                                    default:
                                        if false {
                                            if true {
                                                fmt.Println(v18)
                                                if false {
                                                    v3 = float64(1e10)
                                                } else {
                                                    fmt.Println(v3, fi19, fi21)
                                                    var v26 uint32 = 65535
                                                    _ = v26
                                                }
                                                if true {
                                                    v3 = (((v3 + float64(0.0)) / v3) + v3)
                                                    v27 := GenS4{F0: 1}
                                                    v28 := GenS1{F0: 1000}
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    v29 := GenS3{F0: (-(fi19)), F1: 4294967295, F2: -606}
                                                    _ = v29
                                                }
                                            }
                                            v30 := GenS5{F0: -288}
                                            if false {
                                                if true {
                                                    v31 := GenS2{F0: false, F1: 2}
                                                    v18 = v18
                                                    v3 = v3
                                                    _ = v31
                                                }
                                                if (42 != -9223372036854775808) {
                                                    v18 = (-(v18))
                                                    fmt.Println(fi19)
                                                    fmt.Println(v18)
                                                } else {
                                                    var v32 uint8 = 255
                                                    v33 := GenS6{F0: true, F1: "hello"}
                                                    _ = v32
                                                    _ = v33
                                                }
                                                var v34 float32 = ((1e10 + (3.14159 * (1e10 * -2.25))) / (0.0 - (1e10 / 1.0)))
                                                switch "foo bar" {
                                                case "":
                                                    v35 := GenS7{F0: genFn0(9223372036854775807, 578, 2)}
                                                    _ = v35
                                                case "foo bar":
                                                    if true {
                                                        var v36 uint32 = 42
                                                        v37 := GenS8{F0: (v23 << (((v1 % ((632) | 1))) & 63))}
                                                        fmt.Println(v3, v1, v18)
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                default:
                                                    var v38 uint64 = 18446744073709551615
                                                    _ = v38
                                                }
                                                v39 := GenS9{F0: int16(v23), F1: (fi21 / (((fi21 & -1)) | 1))}
                                                _ = v34
                                                _ = v39
                                            }
                                            v1 = (v1 + v1)
                                            _ = v30
                                        } else {
                                            var v40 uint8 = 127
                                            _ = v40
                                        }
                                    }
                                } else {
                                    v2 = ""
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    v41 := GenS10{F0: 0}
                                    _ = v41
                                }
                            default:
                                v2 = (v2 + v2)
                            }
                            if (42 >= -155) {
                                switch 127 {
                                case -2147483648:
                                    _ = len(v17)
                                case -1:
                                    var v42 string = v2
                                    _ = v42
                                case 199:
                                    var v43 bool = ((682 == (v1 >> ((255) & 63))) && (true && (!false)))
                                    _ = v43
                                }
                                if true {
                                    switch (fi21 ^ (fi19 & (^(fi19)))) {
                                    case 517:
                                        v23 -= v23
                                    default:
                                        var v44 int16 = 32767
                                        _ = v44
                                    }
                                } else {
                                    v45 := GenS11{F0: fi21, F1: 1, F2: 32767}
                                    v3 -= float64(1.5)
                                    v18 = v18
                                    var v46 int64 = -9223372036854775808
                                    _ = v45
                                    _ = v46
                                }
                                var v47 int16 = int16(fi21)
                                _ = v47
                            }
                            switch fi19 {
                            case -508:
                                v3 -= v3
                            }
                            v23 = v23
                            var v48 float64 = v3
                            v23 = (^(v23))
                            _ = v23
                            _ = v48
                        }
                    case "":
                        if true {
                            v17 = append(v17, (fi21 + fi21))
                        }
                        v3 -= v3
                        v2 = v2
                        if true {
                            switch (fi21 % (((fi21 / (((fi19 >> (((v1 % ((v1) | 1))) & 63))) | 1))) | 1)) {
                            case -1:
                                _ = len(v17)
                            }
                        }
                        if len(v17) > 0 {
                            v17[0] = fi19
                        }
                    default:
                        switch v2 {
                        case "a":
                            var v49 float32 = 0.0
                            _ = v49
                        default:
                            fmt.Println(v2, v2, v1)
                        }
                    }
                    if (!false) {
                        switch v2 {
                        case "a":
                            if (1 == 0) {
                                v18 = (v18 & (v18 >> ((v1) & 63)))
                                var v50 int16 = (v18 / ((896) | 1))
                                if (!(false && false)) {
                                    v1 -= (v1 * (^(v1)))
                                    if false {
                                        if true {
                                            var v51 uint8 = 255
                                            if (0.1 >= float64(1e10)) {
                                                if (!(1 > 2147483647)) {
                                                    if len(v17) > 0 {
                                                        v17[0] = fi21
                                                    }
                                                    var v52 string = v2
                                                    _ = v52
                                                } else {
                                                    var v53 uint = (v1 >> (((^(v1))) & 63))
                                                    fmt.Println(v51, v53, fi19)
                                                    if false {
                                                        if (v2 == v2) {
                                                            v54 := GenS12{F0: 697, F1: (v51 << (((v53 << ((42) & 63))) & 63))}
                                                            v55 := GenS11{F0: int(v51), F1: 18446744073709551615, F2: fi19}
                                                            var v56 int16 = -548
                                                            fmt.Println(v3, v53, v53)
                                                            if true {
                                                                fmt.Println(v53, v50)
                                                                var v57 float64 = (v3 + 1e10)
                                                                v2 = ((v2 + v2) + "a")
                                                                if len(v17) >= 2 {
                                                                    v17 = v17[1:len(v17)]
                                                                }
                                                                _ = v57
                                                            } else {
                                                                var v58 float64 = 0.0
                                                                _ = v58
                                                            }
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        switch -9223372036854775808 {
                                                        case -247:
                                                            _ = len(v17)
                                                        }
                                                        v51 = v51
                                                    } else {
                                                        v59 := GenS13{F0: ((!false) || false), F1: (((fi19 * (fi19 ^ 42)) >= (-(fi21))) || false), F2: 462}
                                                        if true {
                                                            _ = len(v17)
                                                            var v60 bool = (-161 == v18)
                                                            if false {
                                                                v60 = v60
                                                                switch v2 {
                                                                case "foo bar":
                                                                    var v61 int32 = 42
                                                                    _ = v61
                                                                }
                                                                v17 = append(v17, (fi19 % (((fi21 / (((fi21 / ((fi21) | 1))) | 1))) | 1)))
                                                            } else {
                                                                v18 = (v18 / ((v50) | 1))
                                                                var v62 int = (fi19 % ((int(v53)) | 1))
                                                                if (558 >= 1) {
                                                                    var v63 float64 = 0.0
                                                                    switch "" {
                                                                    case "hello":
                                                                        v18 -= (^(v18))
                                                                    }
                                                                    if (!(!(!v60))) {
                                                                        if (1 > 2147483647) {
                                                                            switch fi19 {
                                                                            case 669:
                                                                                var v64 uint32 = 740
                                                                                _ = v64
                                                                            case 2147483647:
                                                                                v50 = v50
                                                                            case 800:
                                                                                if v60 {
                                                                                    v60 = (float64((1e10 - ((1e10 + 1e10) * float32(v3)))) < ((-2.25 + v63) + (0.1 + v3)))
                                                                                    v63 = 1e10
                                                                                    if len(v17) > 0 {
                                                                                        v17[0] = fi21
                                                                                    }
                                                                                    fmt.Println(fi19, fi19, v60)
                                                                                }
                                                                                var v65 int8 = -56
                                                                                _ = v65
                                                                            }
                                                                            if len(v17) >= 2 {
                                                                                v17 = v17[1:len(v17)]
                                                                            }
                                                                            v51 -= v51
                                                                        }
                                                                        if len(v17) > 0 {
                                                                            v17[0] = 127
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v53, v53, v63)
                                                                            var v66 uint32 = 4294967295
                                                                            fmt.Println(v18, v63, v1)
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    v67 := GenS14{F0: v53}
                                                                    _ = v63
                                                                    _ = v67
                                                                } else {
                                                                    var v68 uint64 = 481
                                                                    _ = v68
                                                                }
                                                                _ = v62
                                                            }
                                                            _ = v60
                                                        }
                                                        var v69 uint = (v1 ^ v53)
                                                        _ = v59
                                                        _ = v69
                                                    }
                                                    _ = v53
                                                }
                                                fmt.Println(v2, v51)
                                                v70 := GenS12{F0: 2147483647, F1: v51}
                                                _ = v70
                                            } else {
                                                if false {
                                                    var v71 int = fi21
                                                    _ = v71
                                                } else {
                                                    v72 := GenS5{F0: 2147483647}
                                                    fmt.Println(v3, v3, v3)
                                                    _ = v72
                                                }
                                            }
                                            var v73 uint = 506
                                            _ = v51
                                            _ = v73
                                        }
                                    }
                                } else {
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    fmt.Println(v3)
                                }
                                _ = v50
                            }
                            if true {
                                if (true || (!false)) {
                                    if ((uint16(v1) < 0) || (v3 == 1.5)) {
                                        v17 = append(v17, (fi19 - fi21))
                                        if (!(false || (fi21 == 4294967295))) {
                                            var v74 bool = false
                                            switch (fi19 - fi21) {
                                            case 9223372036854775807:
                                                if false {
                                                    if false {
                                                        v75 := GenS15{F0: 127, F1: v74, F2: -1}
                                                        var v76 uint32 = 3
                                                        v77 := GenS7{F0: 1}
                                                        _ = len(v17)
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    v78 := GenS4{F0: -2147483648}
                                                    v74 = (!((v2 + v2) <= (v2 + (("\n\t" + v2) + (((("a" + v2) + (v2 + v2)) + v2) + "hello")))))
                                                    _ = v78
                                                }
                                            case 581:
                                                v74 = ((-779 >= -1) && false)
                                            }
                                            _ = v74
                                        }
                                        var v79 float64 = -2.25
                                        _ = v79
                                    } else {
                                        if (false || true) {
                                            v80 := GenS7{F0: -1}
                                            v1 = v1
                                            _ = v80
                                        } else {
                                            _ = len(v17)
                                        }
                                        if true {
                                            v81 := GenS4{F0: 0}
                                            _ = len(v17)
                                            switch fi19 {
                                            case -415:
                                                var v82 uint8 = 1
                                                _ = v82
                                            case 127:
                                                fmt.Println(v1, fi19)
                                            }
                                            _ = v81
                                        } else {
                                            var v83 int64 = 9223372036854775807
                                            var v84 uint64 = 308
                                            fmt.Println(v83)
                                            _ = v83
                                            _ = v84
                                        }
                                        var v85 float32 = 1e10
                                        _ = v85
                                    }
                                    var v86 int8 = 127
                                    var v87 uint8 = 133
                                    if len(v17) > 0 {
                                        v17[0] = (fi19 + (fi21 - fi21))
                                    }
                                    if false {
                                        v1 -= v1
                                    }
                                    _ = v86
                                    _ = v87
                                } else {
                                    if false {
                                        v17 = append(v17, fi21)
                                        v88 := GenS13{F0: (206 == 948), F1: (921 > 2147483648), F2: 42}
                                        _ = v88
                                    } else {
                                        var v89 uint16 = 255
                                        v2 = v2
                                        _ = len(v17)
                                        _ = v89
                                    }
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                }
                                v90 := GenS15{F0: 255, F1: true, F2: 0}
                                fmt.Println(v2, fi21, v18)
                                if len(v17) > 0 {
                                    v17[0] = 1
                                }
                                _ = v90
                            } else {
                                var v91 float32 = -2.25
                                var v92 uint32 = 923
                                _ = v91
                                _ = v92
                            }
                            v3 -= v3
                            if len(v17) > 0 {
                                v17[0] = fi19
                            }
                        case "\n\t":
                            v18 = (v18 % ((42) | 1))
                        }
                        if len(v17) > 0 {
                            v17[0] = fi21
                        }
                    } else {
                        fmt.Println(fi19, v18, fi19)
                        if (522 > (fi21 << (((v1 - (^(v1)))) & 63))) {
                            v93 := GenS16{F0: 0, F1: 1}
                            v94 := GenS4{F0: 660}
                            v1 -= v1
                            _ = v93
                            _ = v94
                        }
                        var v95 uint32 = 2147483648
                        v95 = 2147483647
                        if len(v17) > 0 {
                            v17[0] = fi21
                        }
                        v17 = append(v17, fi19)
                        _ = v95
                    }
                    _ = v22
                }
                v2 = "a"
                switch v2 {
                case "hello":
                    v3 += 3.14159
                case "a":
                    fmt.Println(v3)
                case "\n\t":
                    if len(v17) > 0 {
                        v17[0] = (fi21 / (((fi21 & fi19)) | 1))
                    }
                    if (2 != 128) {
                        v96 := GenS17{F0: v1}
                        if len(v17) > 0 {
                            v17[0] = (fi19 ^ (fi21 | (fi21 & fi19)))
                        }
                        _ = v96
                    }
                    v18 = (v18 * (v18 / ((v18) | 1)))
                    var v97 int32 = 0
                    fmt.Println(v1, v3)
                    v2 = "hello"
                    var v98 string = ((((v2 + v2) + v2) + "foo bar") + "a")
                    _ = v97
                    _ = v98
                default:
                    if false {
                        if false {
                            v99 := GenS15{F0: 127, F1: true, F2: v18}
                            var v100 uint16 = 0
                            v101 := GenS13{F0: false, F1: true, F2: 0}
                            _ = v100
                            _ = v101
                            _ = v99
                        }
                    }
                    fmt.Println(fi19)
                }
                var v102 string = v2
                var v103 uint = v1
                var v104 bool = true
                if v104 {
                    v3 = v3
                    var v105 float32 = 1.5
                    _ = v105
                } else {
                    var v106 bool = (1 < 2)
                    v1 += 65535
                    _ = v106
                }
                if len(v17) > 0 {
                    v17[0] = 0
                }
                _ = fi21
                _ = v102
                _ = v103
                _ = v104
            }
            var v107 bool = (v2 < "")
            var v108 int = fi19
            _ = v107
            _ = v108
        }
        _ = fi19
    }
    for fi109 := 0; fi109 < 13; fi109++ {
        var v110 string = ("foo bar" + v2)
        _ = fi109
        _ = v110
    }
    v3 = (v3 + 0.0)
    v17 = append(v17, -113)
    var v111 uint16 = 1
    for fi112 := 0; fi112 < 1; fi112++ {
        v113 := GenS18{F0: 2147483647, F1: (!false), F2: uint64(v111)}
        if ((^(v18)) >= v18) {
            v17 = append(v17, 128)
            switch (v2 + "foo bar") {
            case "":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case "hello":
                _ = len(v17)
            }
        } else {
            v1 -= v1
            switch fi112 {
            case -1:
                v111 += 32767
            case 246:
                if true {
                    v111 += v111
                }
            }
        }
        switch v2 {
        case "":
            if len(v17) > 0 {
                v17[0] = fi112
            }
        }
        _ = fi112
        _ = v113
    }
    var v114 uint8 = 0
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v111
    _ = v114
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
