package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string; F2 uint64 }
type GenS6 struct { F0 string; F1 int32; F2 int64 }
type GenS7 struct { F0 int16; F1 uint8; F2 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int; F1 bool; F2 uint8 }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 int64; F1 uint32 }
type GenS18 struct { F0 int32; F1 int64; F2 int }
type GenS19 struct { F0 int; F1 uint8 }
type GenS20 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0() int32 {
    if true {
        switch "hello" {
        case "a":
            var v6 bool = true
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 9223372036854775807}
        v8 := []int{0, 114, 2147483647, 1}
        _ = v7
        _ = v8
    }
    var v9 uint32 = 65535
    _ = v9
    return 0
}
func genFn1() uint64 {
    v10 := GenS2{F0: 2, F1: -2147483648, F2: 18446744073709551615}
    var v11 uint32 = 4294967295
    if true {
        v12 := GenS1{F0: 2147483647}
        _ = v12
    }
    var v13 int = 32767
    if (!(18446744073709551615 > 293)) {
        if (0.1 >= (1e10 - float32((3.14159 - 1e10)))) {
            for fi14 := 0; fi14 < 11; fi14++ {
                v11 += v11
                if false {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v11 -= v11
                        _ = fi15
                    }
                }
                var v16 int32 = 2147483647
                _ = fi14
                _ = v16
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                switch ("foo bar" + ("hello" + "")) {
                case "\n\t":
                    if (!(18446744073709551615 <= 312)) {
                        var v18 uint8 = 2
                        v19 := []int{}
                        _ = v18
                        _ = v19
                    }
                case "hello":
                    fmt.Println(fi17, fi17, fi17)
                }
                _ = fi17
            }
        } else {
            var v20 int64 = 29
            _ = v20
        }
    } else {
        v21 := []int{v13, 128}
        var v22 uint8 = 2
        _ = v21
        _ = v22
    }
    _ = v10
    _ = v11
    _ = v13
    return 0
}

func main() {
    v1 := []string{}
    var v2 uint8 = 127
    var v3 int32 = 479
    var v4 uint16 = 251
    var v5 string = "foo bar"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "\n\t" {
    case "a":
        if len(v1) > 0 {
            v1[0] = ""
        }
    default:
        if false {
            switch v5 {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    fmt.Println(v4)
                case "":
                    var v23 int8 = int8(v2)
                    _ = v23
                }
            case "a":
                v24 := GenS3{F0: 255, F1: (v4 / (((v4 - v4)) | 1))}
                _ = v24
            default:
                v3 = (v3 * (-(v3)))
            }
        } else {
            var v25 uint = 65535
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if ((1 == 42) && (v5 != "\n\t")) {
            var v27 float32 = (1.5 / (-2.25 + 0.1))
            _ = v27
        }
        switch v5 {
        case "\n\t":
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int8 = 42
                _ = fi28
                _ = v29
            }
        case "hello":
            v2 = (v2 ^ v2)
        default:
            var v30 uint64 = 2
            _ = v30
        }
        _ = fi26
    }
    v31 := GenS4{F0: -1}
    v32 := GenS2{F0: 18446744073709551615, F1: 32768, F2: genFn1()}
    if (208 != genFn1()) {
        fmt.Println(v5)
    } else {
        var v33 float64 = (((float64(3.14159) + -2.25) * (1.5 - 1.5)) - 3.14159)
        var v34 uint = 4294967295
        var v35 int16 = -1
        _ = v33
        _ = v34
        _ = v35
    }
    v2 += (v2 / ((v2) | 1))
    switch -579 {
    case -782:
        var v36 int8 = -72
        _ = v36
    case -126:
        v37 := GenS5{F0: -32768, F1: v5, F2: 18446744073709551615}
        _ = v37
    default:
        if true {
            if false {
                v5 = v5
                if true {
                    v2 += v2
                } else {
                    switch 32767 {
                    case -9223372036854775808:
                        fmt.Println(v4)
                    case 2147483647:
                        for fi38 := 0; fi38 < 2; fi38++ {
                            if false {
                                v3 = genFn0()
                                v3 += (v3 << ((1) & 63))
                                fmt.Println(v2, v4)
                            } else {
                                var v39 int64 = -9223372036854775808
                                var v40 uint32 = 0
                                if true {
                                    switch (fi38 * (fi38 ^ -9223372036854775808)) {
                                    case 128:
                                        var v41 int64 = v39
                                        _ = v41
                                    case 32768:
                                        if true {
                                            var v42 uint = 18446744073709551615
                                            for fi43 := 0; fi43 < 1; fi43++ {
                                                v44 := GenS6{F0: (v5 + v5), F1: v3, F2: (v39 >> (((v42 % (((v42 >> (((v42 ^ 4294967295)) & 63))) | 1))) & 63))}
                                                v45 := GenS7{F0: 1000, F1: v2, F2: 1}
                                                v5 = ("" + (v5 + v5))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                            }
                                            var v46 int32 = v3
                                            _ = v42
                                            _ = v46
                                        }
                                    }
                                    for fi47 := 0; fi47 < 4; fi47++ {
                                        if (((false && (-783 < 127)) && false) && (v39 == (v39 - 9223372036854775807))) {
                                            fmt.Println(fi38)
                                        } else {
                                            switch fi47 {
                                            case 0:
                                                v2 = (v2 >> ((472) & 63))
                                                if true {
                                                    v48 := GenS6{F0: v5, F1: -839, F2: v39}
                                                    var v49 float32 = (0.1 / 1e10)
                                                    if true {
                                                        fmt.Println(v49, fi47)
                                                        v5 = v5
                                                        v49 -= 1.5
                                                    } else {
                                                        v49 = ((v49 + v49) + float32(0.0))
                                                        v5 = "a"
                                                        v50 := GenS2{F0: 2, F1: -414, F2: 0}
                                                        _ = v50
                                                    }
                                                    fmt.Println(v49, v2, v40)
                                                    v40 -= 2
                                                    var v51 uint64 = 272
                                                    _ = v48
                                                    _ = v49
                                                    _ = v51
                                                } else {
                                                    v52 := GenS8{F0: (fi47 + (fi38 << ((42) & 63)))}
                                                    var v53 uint8 = v2
                                                    var v54 int32 = 2147483647
                                                    var v55 int32 = -1
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                                switch (((v5 + "a") + v5) + (v5 + (("hello" + v5) + "\n\t"))) {
                                                case "hello":
                                                    var v56 uint64 = 451
                                                    _ = v56
                                                default:
                                                    v57 := GenS5{F0: -32768, F1: (v5 + "hello"), F2: 2}
                                                    _ = v57
                                                }
                                            case 32767:
                                                if ((3.14159 / 1.5) >= 3.14159) {
                                                    fmt.Println(v4)
                                                    if (false && true) {
                                                        if true {
                                                            var v58 int = 42
                                                            if false {
                                                                _ = len(v1)
                                                                var v59 float64 = ((-2.25 - (0.0 - 0.1)) / (1e10 * (0.0 - 0.1)))
                                                                switch 127 {
                                                                case 127:
                                                                    v60 := GenS8{F0: (fi38 ^ (fi47 ^ v58))}
                                                                    _ = v60
                                                                case 128:
                                                                    var v61 int64 = v39
                                                                    _ = v61
                                                                case -2147483648:
                                                                    v2 = 83
                                                                }
                                                                _ = v59
                                                            }
                                                            switch fi47 {
                                                            case 9223372036854775807:
                                                                v62 := GenS9{F0: 65535, F1: (!false)}
                                                                var v63 int32 = (v3 / ((genFn0()) | 1))
                                                                switch v5 {
                                                                case "hello":
                                                                    var v64 float32 = 0.0
                                                                    _ = v64
                                                                }
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            _ = v58
                                                        } else {
                                                            fmt.Println(v3, v40)
                                                            v4 += 0
                                                            v65 := GenS10{F0: 32767}
                                                            _ = v65
                                                        }
                                                    }
                                                }
                                            case 32768:
                                                var v66 float64 = 1.5
                                                _ = v66
                                            }
                                            v67 := GenS11{F0: (v39 ^ (v39 >> ((633) & 63)))}
                                            _ = v67
                                        }
                                        var v68 uint16 = (v4 << ((uint(v40)) & 63))
                                        fmt.Println(v3, v39, fi38)
                                        _ = fi47
                                        _ = v68
                                    }
                                } else {
                                    var v69 uint8 = uint8(v4)
                                    if (0.1 != 1.5) {
                                        switch v5 {
                                        case "":
                                            _ = len(v1)
                                        default:
                                            for fi70 := 0; fi70 < 12; fi70++ {
                                                v71 := GenS12{F0: v40}
                                                v72 := GenS5{F0: 1, F1: (("foo bar" + v5) + "hello"), F2: 0}
                                                var v73 float64 = 3.14159
                                                _ = fi70
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            var v74 string = v5
                                            _ = v74
                                        default:
                                            v75 := GenS10{F0: (v4 / (((^(v4))) | 1))}
                                            _ = v75
                                        }
                                        if (false && ((false && ((42 > -860) || true)) || true)) {
                                            for fi76 := 0; fi76 < 12; fi76++ {
                                                v40 = 65535
                                                v4 += (v4 >> ((4294967295) & 63))
                                                switch (fi38 | (fi38 | fi38)) {
                                                case -9223372036854775808:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case -702:
                                                    var v77 uint16 = (v4 >> ((18446744073709551615) & 63))
                                                    var v78 float64 = (-2.25 + 1.5)
                                                    var v79 uint64 = 4294967295
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                case 1:
                                                    _ = len(v1)
                                                default:
                                                    fmt.Println(v39, v3)
                                                }
                                                if true {
                                                    var v80 int8 = 100
                                                    var v81 string = ("a" + v5)
                                                    _ = v80
                                                    _ = v81
                                                }
                                                if (1 >= 1) {
                                                    if (false || (false && (85 <= 819))) {
                                                        if len(v1) > 0 {
                                                            v1[0] = "hello"
                                                        }
                                                        v2 -= v2
                                                        _ = len(v1)
                                                        _ = len(v1)
                                                    }
                                                }
                                                _ = fi76
                                            }
                                            if true {
                                                v2 = uint8(v40)
                                                fmt.Println(v5, v69, v69)
                                            } else {
                                                if (fi38 < (fi38 >> ((218) & 63))) {
                                                    var v82 uint32 = v40
                                                    _ = len(v1)
                                                    if true {
                                                        v83 := GenS3{F0: (v2 * 128), F1: v4}
                                                        v82 = v40
                                                        _ = v83
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v82 = 674
                                                    }
                                                    _ = v82
                                                } else {
                                                    v3 = (v3 << ((42) & 63))
                                                }
                                                var v84 uint = 42
                                                if true {
                                                    var v85 uint16 = v4
                                                    var v86 uint32 = (^(v40))
                                                    _ = v85
                                                    _ = v86
                                                }
                                                _ = v84
                                            }
                                        } else {
                                            fmt.Println(fi38)
                                            v87 := GenS13{F0: -1, F1: false, F2: (^(v2))}
                                            v88 := GenS14{F0: 398, F1: (fi38 | (^(fi38)))}
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                    v89 := GenS15{F0: 100}
                                    _ = v69
                                    _ = v89
                                }
                                _ = v39
                                _ = v40
                            }
                            _ = fi38
                        }
                    case 32768:
                        if (3.14159 <= 3.14159) {
                            var v90 uint8 = (^(v2))
                            var v91 uint64 = 355
                            if false {
                                switch -9223372036854775808 {
                                case 844:
                                    v2 = (v90 ^ (v90 / (((v2 ^ 2)) | 1)))
                                case -1:
                                    for fi92 := 0; fi92 < 19; fi92++ {
                                        fmt.Println(v2)
                                        v91 += v91
                                        var v93 uint16 = 918
                                        _ = fi92
                                        _ = v93
                                    }
                                default:
                                    if true {
                                        for fi94 := 0; fi94 < 14; fi94++ {
                                            var v95 int32 = (v3 ^ v3)
                                            switch fi94 {
                                            case -36:
                                                var v96 int64 = 797
                                                _ = v96
                                            case 128:
                                                for fi97 := 0; fi97 < 8; fi97++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = (v5 + v5)
                                                    }
                                                    v90 -= uint8(v91)
                                                    var v98 uint8 = (v90 << ((4294967295) & 63))
                                                    _ = fi97
                                                    _ = v98
                                                }
                                            case -908:
                                                for fi99 := 0; fi99 < 11; fi99++ {
                                                    if (v5 == "hello") {
                                                        fmt.Println(fi99)
                                                        if false {
                                                            if (266 >= int64(fi99)) {
                                                                switch (("\n\t" + ("a" + (((v5 + (v5 + (("" + "foo bar") + v5))) + ((v5 + "") + (v5 + v5))) + v5))) + "foo bar") {
                                                                case "foo bar":
                                                                    v91 += (v91 % ((1) | 1))
                                                                case "\n\t":
                                                                    if (-2.25 != (1.5 - -2.25)) {
                                                                        v100 := GenS16{F0: uint64(v4), F1: 255}
                                                                        v101 := GenS1{F0: (fi94 << ((255) & 63))}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v91, v91)
                                                            var v102 bool = true
                                                            _ = v102
                                                        }
                                                    } else {
                                                        var v103 int8 = 0
                                                        _ = v103
                                                    }
                                                    var v104 int16 = -1
                                                    if len(v1) > 0 {
                                                        v1[0] = v5
                                                    }
                                                    _ = fi99
                                                    _ = v104
                                                }
                                            default:
                                                fmt.Println(v5)
                                            }
                                            v105 := GenS17{F0: -1, F1: 550}
                                            _ = fi94
                                            _ = v105
                                            _ = v95
                                        }
                                    }
                                }
                                switch ("" + "foo bar") {
                                case "":
                                    var v106 int64 = 1
                                    _ = v106
                                default:
                                    var v107 uint64 = (v91 + v91)
                                    _ = v107
                                }
                            } else {
                                v4 = 255
                                if true {
                                    v2 -= uint8(v91)
                                } else {
                                    v108 := GenS18{F0: int32(v4), F1: 968, F2: -329}
                                    if len(v1) > 0 {
                                        v1[0] = (v5 + (v5 + ""))
                                    }
                                    _ = v108
                                }
                                v109 := GenS19{F0: 488, F1: (^(v2))}
                                _ = v109
                            }
                            _ = v90
                            _ = v91
                        }
                    }
                    var v110 bool = (127 <= int8(v3))
                    v3 = (v3 / ((genFn0()) | 1))
                    _ = v110
                }
                for fi111 := 0; fi111 < 2; fi111++ {
                    var v112 uint64 = genFn1()
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = fi111
                    _ = v112
                }
            } else {
                v113 := GenS20{F0: (v4 ^ v4), F1: true, F2: 0}
                if ((!false) && true) {
                    v114 := GenS14{F0: 328, F1: 1}
                    var v115 string = "foo bar"
                    _ = v114
                    _ = v115
                }
                if true {
                    var v116 int16 = 1000
                    _ = v116
                }
                _ = v113
            }
        }
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
