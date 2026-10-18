package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int8; F1 uint32; F2 int }
type GenS8 struct { F0 string; F1 int32; F2 uint32 }
type GenS9 struct { F0 int64; F1 bool; F2 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int; F1 uint32; F2 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string; F1 int; F2 int32 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int16; F1 int16; F2 int }
type GenS19 struct { F0 uint; F1 int8; F2 uint8 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 uint64; F1 uint32 }
type GenS22 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS23 struct { F0 int64; F1 int64; F2 uint8 }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 uint64; F1 uint8 }
type GenS26 struct { F0 int32; F1 uint8 }
type GenS27 struct { F0 bool }
type GenS28 struct { F0 uint8 }
type GenS29 struct { F0 uint16; F1 uint }
type GenS30 struct { F0 uint16 }
type GenS31 struct { F0 string; F1 int64 }
type GenS32 struct { F0 uint; F1 uint32 }
type GenS33 struct { F0 int32; F1 bool }
type GenS34 struct { F0 int32; F1 string }
type GenS35 struct { F0 uint64 }
type GenS36 struct { F0 int8; F1 int32; F2 int64 }
type GenS37 struct { F0 int32; F1 bool }
type GenS38 struct { F0 uint32 }
type GenS39 struct { F0 uint }
type GenS40 struct { F0 string; F1 int; F2 bool }
type GenS41 struct { F0 bool; F1 int32; F2 uint16 }
type GenS42 struct { F0 int64 }
type GenS43 struct { F0 int32; F1 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 int64 = -165
    var v4 int32 = v1
    v5 := make(map[string]bool)
    v3 -= v3
    var v6 int = 42
    v7 := GenS1{F0: v3}
    var v8 float64 = 1.5
    var v9 int = -9223372036854775808
    var v10 int32 = v4
    if false {
        for fi11 := 0; fi11 < 10; fi11++ {
            v1 = 42
            v10 -= v1
            switch "\n\t" {
            case "hello":
                v12 := GenS2{F0: 42, F1: 42}
                _ = v12
            default:
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := []int{(^(fi11)), (fi13 % ((v6) | 1))}
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi11
        }
        v4 = (v1 >> ((255) & 63))
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint16 = 2
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS1{F0: v3}
            switch "a" {
            case "\n\t":
                var v19 uint64 = 1
                _ = v19
            case "a":
                if false {
                    var v20 int32 = 1
                    fmt.Println(v20)
                    _ = v20
                } else {
                    if true {
                        v6 -= fi15
                        var v21 int = -1
                        switch fi15 {
                        case 127:
                            if (-1 >= v1) {
                                var v22 bool = (1 <= 349)
                                var v23 string = (("a" + "a") + (("a" + "foo bar") + ""))
                                _ = v22
                                _ = v23
                            }
                        case 1:
                            fmt.Println(v16, v1)
                        case 997:
                            var v24 int = (v9 ^ v6)
                            _ = v24
                        }
                        _ = v21
                    } else {
                        if (1000 != -32768) {
                            v4 = v10
                            var v25 uint64 = 0
                            _ = v25
                        }
                        var v26 string = ("foo bar" + "a")
                        var v27 int16 = -1
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v9, v1, v1)
                }
            case "hello":
                switch v6 {
                case 9223372036854775807:
                    if false {
                        v6 = (v9 - v6)
                        v9 = v9
                    } else {
                        switch ("a" + "a") {
                        case "foo bar":
                            var v28 int8 = 1
                            _ = v28
                        case "a":
                            fmt.Println(v6, v1)
                        default:
                            fmt.Println(v8, v9)
                        }
                    }
                case 42:
                    if false {
                        if true {
                            var v29 uint = 0
                            _ = len(v2)
                            switch "a" {
                            case "foo bar":
                                switch "\n\t" {
                                case "a":
                                    v30 := GenS2{F0: 2147483647, F1: 857}
                                    v31 := []int{32767, fi15, fi17, (fi17 % ((996) | 1))}
                                    v16 += (v16 ^ (v16 * (v16 / (((v16 ^ v16)) | 1))))
                                    _ = v30
                                    _ = v31
                                case "":
                                    switch v6 {
                                    case 9223372036854775807:
                                        delete(v5, ("hello" + ((((("foo bar" + "\n\t") + "") + ("hello" + "")) + (((("hello" + "a") + ("foo bar" + "")) + "hello") + "foo bar")) + ("\n\t" + ""))))
                                    }
                                case "\n\t":
                                    if ((-(v4)) < (v1 | (v10 ^ -1))) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v32 int16 = 32767
                                            _ = v32
                                        case "":
                                            var v33 string = "foo bar"
                                            _ = v33
                                        case "a":
                                            v8 += 1e10
                                        }
                                        v2[int(v3)] = true
                                        var v34 string = "\n\t"
                                        _ = v34
                                    } else {
                                        var v35 float32 = 1.5
                                        v10 = v1
                                        if false {
                                            v10 += (v1 - (^(v4)))
                                            v36 := []int{}
                                            _ = v36
                                        } else {
                                            v8 = (0.1 * v8)
                                        }
                                        _ = v35
                                    }
                                default:
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            if false {
                                                v37 := GenS1{F0: 968}
                                                fmt.Println(fi17, v9, fi17)
                                                _ = v37
                                            } else {
                                                var v38 uint = (v29 << (((v29 / ((2) | 1))) & 63))
                                                v6 += v6
                                                _ = len(v5)
                                                if false {
                                                    switch "" {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            switch (fi17 + int(v3)) {
                                                            case 32768:
                                                                if false {
                                                                    v39 := []int{}
                                                                    _ = v39
                                                                }
                                                            case -9223372036854775808:
                                                                var v40 int = v6
                                                                _ = v40
                                                            }
                                                        default:
                                                            v9 = v6
                                                        }
                                                    }
                                                    v41 := GenS2{F0: 688, F1: -352}
                                                    _ = v41
                                                } else {
                                                    var v42 float32 = ((((1.5 * (1.5 * 3.14159)) - 1.5) + 1.5) + -2.25)
                                                    var v43 float64 = (1e10 * 1e10)
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v38
                                            }
                                            var v44 int32 = 0
                                            var v45 uint = v29
                                            _ = v44
                                            _ = v45
                                        }
                                    case "\n\t":
                                        v4 += v1
                                    case "":
                                        fmt.Println(fi17, fi17, v29)
                                    }
                                }
                            case "a":
                                v8 += (v8 + (v8 * 0.0))
                            default:
                                v46 := []int{v9, (fi17 << ((v29) & 63)), (fi15 << ((v29) & 63)), 9223372036854775807}
                                _ = v46
                            }
                            v10 = 42
                            var v47 int16 = 1000
                            v48 := []int{(v6 * 672), (v6 & fi15), -1}
                            var v49 int64 = 1
                            if true {
                                if ((!(!(v16 >= (v16 % ((v16) | 1))))) && ("\n\t" <= "\n\t")) {
                                    var v50 uint = (v29 << ((v29) & 63))
                                    _ = v50
                                } else {
                                    if false {
                                        var v51 float64 = v8
                                        fmt.Println(v51, fi17, v29)
                                        v3 += v49
                                        _ = v51
                                    } else {
                                        v16 = 572
                                        switch "\n\t" {
                                        case "a":
                                            switch "foo bar" {
                                            case "a":
                                                if (fi15 > fi17) {
                                                    var v52 int16 = -1
                                                    if (((v1 + (^(v4))) < (v1 % (((v1 | v4)) | 1))) && true) {
                                                        v53 := GenS1{F0: -327}
                                                        v48 = append(v48, -1)
                                                        v9 = v9
                                                        _ = v53
                                                    }
                                                    if true {
                                                        v54 := GenS2{F0: 1, F1: (v47 % ((v47) | 1))}
                                                        _ = v54
                                                    } else {
                                                        _ = len(v48)
                                                    }
                                                    var v55 uint32 = 65535
                                                    v48 = append(v48, fi17)
                                                    var v56 uint8 = 42
                                                    v9 -= v9
                                                    _ = v52
                                                    _ = v55
                                                    _ = v56
                                                }
                                            case "":
                                                v8 += 1e10
                                            case "\n\t":
                                                var v57 float64 = (1e10 - (v8 + -2.25))
                                                _ = v57
                                            }
                                        case "hello":
                                            if true {
                                                if true {
                                                    v9 = (-(fi15))
                                                    fmt.Println(v29, v3, v10)
                                                    var v58 int32 = v10
                                                    var v59 int8 = 99
                                                    v10 += (v4 >> ((18446744073709551615) & 63))
                                                    _ = v58
                                                    _ = v59
                                                }
                                            } else {
                                                v9 += v9
                                                var v60 int8 = 127
                                                _ = v60
                                            }
                                        default:
                                            var v61 uint8 = 120
                                            _ = v61
                                        }
                                        var v62 bool = ((v47 * (v47 | v47)) > 1000)
                                        if (("foo bar" + "") != "\n\t") {
                                            if (v47 < v47) {
                                                var v63 int32 = v10
                                                _ = v63
                                            }
                                            var v64 uint64 = 42
                                            _ = v64
                                        }
                                        _ = v62
                                    }
                                }
                            } else {
                                _ = len(v48)
                                if false {
                                    if false {
                                        v4 = v10
                                        switch "a" {
                                        case "foo bar":
                                            var v65 uint = 2
                                            _ = v65
                                        case "hello":
                                            if ((v16 + (v16 << ((v29) & 63))) <= (v16 - 32767)) {
                                                switch 128 {
                                                case 2147483647:
                                                    var v66 int64 = (v49 * 0)
                                                    _ = v66
                                                case -1:
                                                    if false {
                                                        fmt.Println(v1, v4, v1)
                                                        var v67 int32 = v4
                                                        var v68 int16 = v47
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                case 32768:
                                                    delete(v2, (v6 / (((fi17 * (^(fi17)))) | 1)))
                                                default:
                                                    var v69 bool = (false && true)
                                                    _ = v69
                                                }
                                            } else {
                                                if false {
                                                    _ = len(v48)
                                                    if ((1.5 <= 0.1) || (v16 > (v16 - 757))) {
                                                        var v70 int8 = 84
                                                        if len(v48) > 0 {
                                                            v48[0] = 1
                                                        }
                                                        _ = v70
                                                    } else {
                                                        if true {
                                                            fmt.Println(v1, v47)
                                                            var v71 int64 = (v49 % ((v3) | 1))
                                                            var v72 int = (fi15 / (((^(v9))) | 1))
                                                            var v73 uint32 = uint32(v29)
                                                            v5["a"] = true
                                                            var v74 float32 = ((float32(((v8 / (v8 + 0.0)) - v8)) / 1.0) * -2.25)
                                                            if (((2 > 42) || (((float32(3.14159) * (v74 / ((v74 + v74) + 3.14159))) / 1.0) >= ((((3.14159 + 1.5) / (v74 + (v74 * v74))) + v74) + v74))) || (!((float32((-2.25 + 0.0)) - ((((v74 * (1e10 - v74)) + v74) / v74) - v74)) == (v74 / 1.5)))) {
                                                                var v75 uint16 = 377
                                                                v74 -= float32((float64((v74 - (v74 + ((0.1 / v74) + v74)))) / v8))
                                                                var v76 uint32 = (v73 | (v73 << ((v29) & 63)))
                                                                v74 -= v74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            v6 = fi17
                                                            var v77 int32 = (v1 | 0)
                                                            if (v47 < v47) {
                                                                fmt.Println(v49, fi15, v16)
                                                                var v78 int32 = (v1 * (v77 >> ((255) & 63)))
                                                                _ = len(v2)
                                                                if len(v48) > 0 {
                                                                    v48[0] = 32767
                                                                }
                                                                v79 := GenS3{F0: "a"}
                                                                if ((203 < 128) && (!true)) {
                                                                    _ = len(v2)
                                                                    v48 = append(v48, int(v4))
                                                                    v80 := v5[(((("foo bar" + "\n\t") + "\n\t") + ("" + "hello")) + "")]
                                                                    var v81 float32 = 1e10
                                                                    _ = v80
                                                                    _ = v81
                                                                }
                                                                _ = v78
                                                                _ = v79
                                                            } else {
                                                                var v82 uint8 = 42
                                                                _ = v82
                                                            }
                                                            v83 := GenS4{F0: 139, F1: 127}
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                            _ = v74
                                                            _ = v77
                                                            _ = v83
                                                        } else {
                                                            if ((v16 >> (((v29 % ((v29) | 1))) & 63)) <= 65535) {
                                                                var v84 int16 = v47
                                                                var v85 int64 = (v49 % ((42) | 1))
                                                                _ = len(v48)
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            fmt.Println(v10)
                                                            fmt.Println(v6, v6, v29)
                                                        }
                                                        switch 963 {
                                                        case 42:
                                                            if false {
                                                                fmt.Println(v10, v9, v3)
                                                                var v86 float32 = 0.0
                                                                _ = v86
                                                            } else {
                                                                _ = len(v48)
                                                                if false {
                                                                    if ((v8 > ((v8 - v8) * v8)) && true) {
                                                                        v87 := GenS5{F0: (v10 << ((v29) & 63)), F1: -9223372036854775808}
                                                                        _ = v87
                                                                    } else {
                                                                        v4 += (v4 ^ 332)
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            if (false && false) {
                                                                                v6 -= (-(v9))
                                                                                v88 := GenS4{F0: v29, F1: 127}
                                                                                fmt.Println(v10, v49)
                                                                                v16 += v16
                                                                                _ = v88
                                                                            }
                                                                        case "":
                                                                            var v89 uint16 = v16
                                                                            _ = v89
                                                                        case "hello":
                                                                            v29 = (^(v29))
                                                                        }
                                                                    }
                                                                    var v90 int16 = (v47 << ((1) & 63))
                                                                    _ = v90
                                                                } else {
                                                                    v8 -= v8
                                                                }
                                                                fmt.Println(fi15)
                                                            }
                                                            var v91 uint32 = 2147483648
                                                            _ = v91
                                                        default:
                                                            switch 1 {
                                                            case -967:
                                                                v92 := GenS6{F0: -128, F1: 1}
                                                                _ = v92
                                                            case 128:
                                                                v4 = v1
                                                            case 816:
                                                                _ = len(v48)
                                                            }
                                                        }
                                                        _ = len(v2)
                                                        switch v9 {
                                                        case 4294967295:
                                                            v16 -= 0
                                                        }
                                                    }
                                                    var v93 uint8 = 0
                                                    fmt.Println(v47, v49, v16)
                                                    switch "a" {
                                                    case "a":
                                                        var v94 int16 = v47
                                                        if false {
                                                            v95 := GenS4{F0: (v29 + v29), F1: v93}
                                                            _ = v95
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                fmt.Println(v29, fi15, fi17)
                                                            default:
                                                                var v96 int16 = 0
                                                                _ = v96
                                                            }
                                                            if false {
                                                                if true {
                                                                    v8 = ((float64(((0.1 * 1e10) * 3.14159)) - (v8 / (0.0 + 1e10))) + (-2.25 - v8))
                                                                    switch (fi17 % (((v9 / ((int(v16)) | 1))) | 1)) {
                                                                    case 2147483647:
                                                                        v8 = float64(1.5)
                                                                    case 42:
                                                                        var v97 int = v6
                                                                        v98 := GenS6{F0: 42, F1: (v93 >> ((v29) & 63))}
                                                                        v99 := v2[(v9 | (v6 << (((v29 * (^(v29)))) & 63)))]
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                }
                                                            } else {
                                                                v6 -= fi15
                                                                v100 := GenS4{F0: (v29 + (v29 * 538)), F1: v93}
                                                                var v101 int = (fi17 & (v9 * (fi15 * -9223372036854775808)))
                                                                _ = v100
                                                                _ = v101
                                                            }
                                                        }
                                                        var v102 int32 = v4
                                                        _ = v102
                                                        _ = v94
                                                    case "foo bar":
                                                        delete(v2, 32767)
                                                    }
                                                    if (1 != 4294967295) {
                                                        v5["a"] = (v93 == 2)
                                                        switch 32767 {
                                                        case 714:
                                                            v93 = (v93 | (^(v93)))
                                                        default:
                                                            if false {
                                                                switch v9 {
                                                                case 42:
                                                                    var v103 uint32 = 421
                                                                    _ = v103
                                                                case -2147483648:
                                                                    if true {
                                                                        v48 = append(v48, 504)
                                                                    }
                                                                }
                                                                if true {
                                                                    v104 := GenS7{F0: 100, F1: 2147483648, F2: 0}
                                                                    if (false || ((v47 != 1) && true)) {
                                                                        v8 -= 0.1
                                                                        switch 207 {
                                                                        case -1:
                                                                            delete(v2, (fi15 - (fi15 ^ (fi15 * (fi15 & fi17)))))
                                                                        case 4294967295:
                                                                            v105 := GenS3{F0: ""}
                                                                            _ = v105
                                                                        case 32768:
                                                                            v10 = v10
                                                                        }
                                                                        if false {
                                                                            if (4294967295 > 387) {
                                                                                v4 -= (v4 >> ((uint(v16)) & 63))
                                                                                v93 = 42
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if ((!true) || (v16 == v16)) {
                                                                            v29 -= 42
                                                                            if (true && (fi15 > (-(v6)))) {
                                                                                if (!false) {
                                                                                    var v106 string = "foo bar"
                                                                                    _ = v106
                                                                                }
                                                                                v9 = (-(v9))
                                                                                if ((fi17 | v6) == 32767) {
                                                                                    var v107 uint64 = 18446744073709551615
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 int = (fi17 % ((127) | 1))
                                                                                    var v109 uint64 = 18446744073709551615
                                                                                    if (!(((float64(1.5) - 0.0) - ((((v8 - v8) * (float64((float32(v8) + ((1e10 * 0.0) * (0.0 / 0.1)))) * 1.5)) + (v8 + (1.5 + v8))) + (0.0 - v8))) >= v8)) {
                                                                                        if len(v48) >= 2 {
                                                                                            v48 = v48[1:len(v48)]
                                                                                        }
                                                                                        var v110 uint = 255
                                                                                        _ = v110
                                                                                    } else {
                                                                                        fmt.Println(fi15, v10, v49)
                                                                                    }
                                                                                    switch (fi15 - (v108 & fi15)) {
                                                                                    case -685:
                                                                                        if false {
                                                                                            var v111 int64 = -1
                                                                                            v112 := GenS2{F0: 42, F1: (v47 - v47)}
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                    case 128:
                                                                                        if true {
                                                                                            v113 := GenS6{F0: 88, F1: (v93 >> (((^(v29))) & 63))}
                                                                                            fmt.Println(fi17)
                                                                                            v3 -= int64(v47)
                                                                                            _ = v113
                                                                                        } else {
                                                                                            _ = len(v5)
                                                                                            if len(v48) > 0 {
                                                                                                v48[0] = -1
                                                                                            }
                                                                                        }
                                                                                    case -2147483648:
                                                                                        _ = len(v48)
                                                                                    }
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                            v4 -= (v10 ^ v10)
                                                                            if (true || false) {
                                                                                v114 := GenS8{F0: "\n\t", F1: (v1 - 0), F2: 65535}
                                                                                switch -1 {
                                                                                case 32768:
                                                                                    if true {
                                                                                        var v115 int8 = 22
                                                                                        v47 += 214
                                                                                        _ = v115
                                                                                    }
                                                                                case 42:
                                                                                    var v116 float64 = v8
                                                                                    _ = v116
                                                                                default:
                                                                                    v6 += (^(v6))
                                                                                    var v117 bool = true
                                                                                    var v118 int32 = (v1 + v1)
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                }
                                                                                _ = v114
                                                                            } else {
                                                                                var v119 float64 = float64((3.14159 - 0.1))
                                                                                _ = len(v48)
                                                                                var v120 string = (("" + ("a" + "hello")) + "a")
                                                                                var v121 int8 = 4
                                                                                v121 += 0
                                                                                v122 := GenS1{F0: (v3 % ((1) | 1))}
                                                                                _ = v119
                                                                                _ = v120
                                                                                _ = v121
                                                                                _ = v122
                                                                            }
                                                                            switch (v6 ^ (fi15 * fi17)) {
                                                                            case 1:
                                                                                v123 := GenS5{F0: -2147483648, F1: (v6 / ((v6) | 1))}
                                                                                _ = v123
                                                                            case 0:
                                                                                var v124 int32 = 42
                                                                                _ = v124
                                                                            case 331:
                                                                                v29 += (v29 ^ v29)
                                                                            default:
                                                                                v4 = 1
                                                                            }
                                                                            v125 := GenS9{F0: 9223372036854775807, F1: (0 != -1), F2: false}
                                                                            v126 := GenS3{F0: "a"}
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                        v10 += v1
                                                                        v127 := GenS10{F0: false}
                                                                        _ = v127
                                                                    }
                                                                    _ = v104
                                                                }
                                                                var v128 uint8 = v93
                                                                _ = v128
                                                            } else {
                                                                v129 := GenS5{F0: (v4 * (-(v1))), F1: fi17}
                                                                _ = v129
                                                            }
                                                        }
                                                        var v130 int16 = (v47 & v47)
                                                        _ = v130
                                                    }
                                                    _ = v93
                                                }
                                            }
                                        case "":
                                            if false {
                                                v1 = v10
                                            } else {
                                                v29 += 255
                                            }
                                            fmt.Println(v6)
                                            v131 := GenS6{F0: -128, F1: 42}
                                            var v132 int = 861
                                            var v133 int8 = 1
                                            _ = v131
                                            _ = v132
                                            _ = v133
                                        }
                                        if (true && (89 == 42)) {
                                            v8 = (1.5 * -2.25)
                                            var v134 int32 = (^(v4))
                                            _ = v134
                                        }
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        v135 := GenS11{F0: v16}
                                        _ = v135
                                    }
                                    var v136 int64 = (v3 * (v3 | v49))
                                    _ = v136
                                }
                                switch ("a" + "a") {
                                case "":
                                    v6 -= (fi17 ^ v9)
                                case "hello":
                                    v137 := GenS12{F0: (v6 - -2147483648), F1: 745, F2: (v16 & v16)}
                                    _ = v137
                                }
                            }
                            v138 := v2[(v6 | 9223372036854775807)]
                            if v138 {
                                var v139 bool = (128 > 42)
                                v49 = v49
                                v9 = fi17
                                _ = v139
                            }
                            _ = v138
                            _ = v29
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            switch (fi15 + 435) {
                            case 9223372036854775807:
                                fmt.Println(v16, v4)
                            }
                        }
                        var v140 bool = false
                        var v141 int32 = (v1 + v1)
                        _ = v140
                        _ = v141
                    }
                case 380:
                    if false {
                        v4 = v4
                        v142 := GenS13{F0: 1}
                        var v143 uint64 = 42
                        _ = v142
                        _ = v143
                    }
                default:
                    var v144 uint64 = 18446744073709551615
                    _ = v144
                }
            }
            switch -2147483648 {
            case -1:
                v145 := []int{(v9 >> ((107) & 63))}
                _ = v145
            default:
                v4 -= v10
            }
            if ((0.0 + float32((0.1 + (-2.25 * v8)))) > ((3.14159 * 1e10) - 3.14159)) {
                v8 = ((v8 + 1e10) - 3.14159)
                if ((fi17 ^ (v9 % ((128) | 1))) > (v9 - v9)) {
                    if false {
                        if (true && true) {
                            var v146 uint64 = 2
                            var v147 string = "a"
                            if (!(v3 > (v3 << ((4294967295) & 63)))) {
                                var v148 uint32 = 4294967295
                                _ = v148
                            } else {
                                v149 := GenS8{F0: v147, F1: (v4 & v4), F2: uint32(v16)}
                                if true {
                                    var v150 uint32 = 407
                                    var v151 bool = ((^(v146)) >= (v146 / ((1) | 1)))
                                    _ = v150
                                    _ = v151
                                } else {
                                    var v152 uint16 = (^(v16))
                                    v8 -= (((v8 / -2.25) + v8) * 1.5)
                                    _ = v152
                                }
                                _ = v149
                            }
                            _ = v146
                            _ = v147
                        }
                    } else {
                        var v153 uint16 = (v16 * (^(v16)))
                        v154 := GenS14{F0: 0}
                        _ = v153
                        _ = v154
                    }
                } else {
                    v155 := GenS15{F0: ("a" + "foo bar"), F1: (fi15 ^ fi17), F2: 42}
                    v156 := []int{fi17, (fi17 << ((1) & 63)), fi15, v9}
                    v157 := GenS16{F0: true}
                    _ = v155
                    _ = v156
                    _ = v157
                }
                fmt.Println(v1)
                delete(v5, (("a" + ("\n\t" + "")) + ""))
            }
            _ = fi17
            _ = v18
        }
        v1 = 42
        _ = fi15
        _ = v16
    }
    switch "\n\t" {
    case "foo bar":
        var v158 uint8 = 216
        _ = v158
    case "hello":
        switch ("\n\t" + (((("" + ("\n\t" + ((("foo bar" + ("" + "foo bar")) + "\n\t") + ("foo bar" + "\n\t")))) + "") + ("hello" + "")) + "\n\t")) {
        case "":
            for fi159 := 0; fi159 < 15; fi159++ {
                v2[v9] = (2 < 0)
                var v160 uint32 = 42
                _ = fi159
                _ = v160
            }
        case "a":
            if true {
                fmt.Println(v4, v4)
                for fi161 := 0; fi161 < 12; fi161++ {
                    if true {
                        for fi162 := 0; fi162 < 20; fi162++ {
                            v1 = 0
                            switch "hello" {
                            case "\n\t":
                                if true {
                                    var v163 uint16 = 662
                                    if ((fi162 | v9) != v6) {
                                        var v164 uint64 = 2
                                        fmt.Println(v3)
                                        v8 += v8
                                        _ = v164
                                    } else {
                                        v165 := GenS13{F0: 944}
                                        _ = v165
                                    }
                                    _ = v163
                                } else {
                                    v3 = (v3 ^ (v3 ^ (v3 * v3)))
                                    v3 += v3
                                    var v166 uint32 = 65535
                                    v167 := GenS5{F0: 0, F1: 718}
                                    var v168 int32 = v1
                                    _ = v166
                                    _ = v167
                                    _ = v168
                                }
                            case "hello":
                                var v169 uint = 4294967295
                                _ = v169
                            case "foo bar":
                                v170 := GenS17{F0: 1}
                                _ = v170
                            default:
                                v171 := v2[(fi162 % ((v9) | 1))]
                                _ = v171
                            }
                            v8 = 1.5
                            _ = fi162
                        }
                        if true {
                            v172 := GenS7{F0: 42, F1: 892, F2: fi161}
                            switch 219 {
                            case 32767:
                                v1 += (^(v10))
                            case 0:
                                for fi173 := 0; fi173 < 5; fi173++ {
                                    switch (("\n\t" + ((("a" + ("foo bar" + (("a" + (("hello" + ("foo bar" + "a")) + "\n\t")) + "a"))) + "foo bar") + "hello")) + "foo bar") {
                                    case "\n\t":
                                        v4 = (v10 >> ((2) & 63))
                                    }
                                    var v174 uint = 0
                                    v175 := GenS18{F0: 408, F1: 685, F2: (fi161 ^ (fi161 * v6))}
                                    v1 += 42
                                    switch (v9 - (^(v9))) {
                                    case 32767:
                                        if false {
                                            v9 = v9
                                            if ((3.14159 * (((1e10 + 1e10) / -2.25) - (1e10 * -2.25))) == -2.25) {
                                                v8 = 1e10
                                                v8 += (((v8 - 3.14159) / float64((-2.25 * 3.14159))) / v8)
                                                switch ("\n\t" + "\n\t") {
                                                case "hello":
                                                    v5["\n\t"] = true
                                                default:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v176 uint8 = 255
                                                        _ = v176
                                                    default:
                                                        var v177 int64 = v3
                                                        _ = v177
                                                    }
                                                }
                                            }
                                            switch "hello" {
                                            case "hello":
                                                switch "" {
                                                case "":
                                                    v174 = v174
                                                }
                                                if true {
                                                    var v178 string = "foo bar"
                                                    var v179 int = (v9 & v6)
                                                    _ = v178
                                                    _ = v179
                                                }
                                                if false {
                                                    fmt.Println(fi161, fi173)
                                                }
                                                if ((v3 | (v3 ^ v3)) >= 0) {
                                                    fmt.Println(v3)
                                                } else {
                                                    fmt.Println(fi161, fi173, v8)
                                                }
                                                v180 := GenS18{F0: -32768, F1: 0, F2: fi161}
                                                _ = v180
                                            case "":
                                                v181 := GenS19{F0: (v174 | (v174 % ((v174) | 1))), F1: 114, F2: 56}
                                                _ = v181
                                            }
                                        } else {
                                            v182 := GenS2{F0: 4294967295, F1: 0}
                                            switch "" {
                                            case "":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (!false) {
                                                        fmt.Println(v1, v174)
                                                        v183 := GenS20{F0: 2}
                                                        _ = v183
                                                    }
                                                case "":
                                                    v174 = v174
                                                default:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if ((!false) || (1 > 581)) {
                                                            v184 := GenS19{F0: v174, F1: 10, F2: 2}
                                                            _ = v184
                                                        }
                                                    default:
                                                        fmt.Println(fi173, v3, fi173)
                                                    }
                                                }
                                            case "hello":
                                                v1 = v1
                                            default:
                                                v174 = 42
                                            }
                                            var v185 int64 = (v3 + 42)
                                            _ = v182
                                            _ = v185
                                        }
                                    case 9223372036854775807:
                                        v186 := v5["a"]
                                        _ = v186
                                    default:
                                        if (!true) {
                                            v187 := GenS21{F0: 2, F1: 65535}
                                            v1 = v10
                                            switch ((("" + "\n\t") + "a") + "foo bar") {
                                            case "foo bar":
                                                var v188 uint8 = 255
                                                _ = v188
                                            case "a":
                                                var v189 int = fi161
                                                _ = v189
                                            case "\n\t":
                                                v190 := []int{128, fi173, 127, -1}
                                                _ = v190
                                            default:
                                                if true {
                                                    v10 = v1
                                                    v1 = v1
                                                    switch (fi161 ^ 32768) {
                                                    case 4294967295:
                                                        v9 -= fi161
                                                    default:
                                                        v9 = 1
                                                        var v191 int = 525
                                                        if ("" < ("\n\t" + "\n\t")) {
                                                            var v192 int64 = (v3 ^ (v3 | (v3 - (v3 % ((v3) | 1)))))
                                                            _ = v192
                                                        }
                                                        switch v191 {
                                                        case -913:
                                                            fmt.Println(v1, v1, v10)
                                                        case 42:
                                                            var v193 int32 = (v4 - (v4 & (-(v10))))
                                                            _ = v193
                                                        case 127:
                                                            delete(v5, ("" + "foo bar"))
                                                        default:
                                                            if (!((0.1 + (((3.14159 - 0.0) - (-2.25 + (1e10 + (-2.25 / float32(1.5))))) + 0.1)) <= -2.25)) {
                                                                switch (fi173 << ((v174) & 63)) {
                                                                case -9223372036854775808:
                                                                    var v194 int32 = v4
                                                                    _ = v194
                                                                }
                                                            }
                                                        }
                                                        var v195 float64 = 0.0
                                                        _ = v191
                                                        _ = v195
                                                    }
                                                    switch (fi173 ^ (-(v9))) {
                                                    case 849:
                                                        if ((true || (((float64(float32((float64(((0.0 - 1.5) / (1e10 * (0.1 - -2.25)))) * (((0.1 + v8) * (((v8 / ((((float64((float32(3.14159) - 1.5)) + 1e10) - (v8 - (((-2.25 + 1e10) * v8) / (v8 / 1.5)))) / 3.14159) + v8)) * v8) - (0.1 * (v8 - float64(0.1))))) + 3.14159)))) - 0.1) - ((v8 / v8) + v8)) <= (0.1 + -2.25))) && false) {
                                                            v196 := []int{(v6 & 1), (v6 + fi161), -681, -1}
                                                            v197 := GenS22{F0: 255, F1: v10, F2: 4294967295}
                                                            fmt.Println(v9, v4)
                                                            _ = v196
                                                            _ = v197
                                                        }
                                                    case 811:
                                                        v9 = (fi161 + -578)
                                                    case -917:
                                                        if (0 > 103) {
                                                            v3 = int64(v6)
                                                            fmt.Println(v6, v4)
                                                            if true {
                                                                fmt.Println(v1)
                                                            }
                                                        } else {
                                                            var v198 uint64 = 42
                                                            v199 := []int{fi173, (v6 | fi173)}
                                                            _ = v198
                                                            _ = v199
                                                        }
                                                    default:
                                                        v200 := []int{v9, v9}
                                                        _ = v200
                                                    }
                                                }
                                            }
                                            _ = v187
                                        }
                                    }
                                    fmt.Println(fi173)
                                    v201 := GenS23{F0: v3, F1: (v3 & int64(fi173)), F2: 0}
                                    _ = fi173
                                    _ = v174
                                    _ = v175
                                    _ = v201
                                }
                            case 4294967295:
                                if ((0 < 42) || ("\n\t" > "\n\t")) {
                                    v202 := GenS24{F0: v3}
                                    var v203 uint16 = 578
                                    var v204 int32 = (v1 & -2147483648)
                                    _ = v202
                                    _ = v203
                                    _ = v204
                                } else {
                                    if (float32(v8) > float32(v8)) {
                                        for fi205 := 0; fi205 < 10; fi205++ {
                                            switch 32767 {
                                            case 127:
                                                switch v6 {
                                                case 0:
                                                    v206 := GenS25{F0: 501, F1: 210}
                                                    _ = v206
                                                case 85:
                                                    v207 := v2[(fi161 + v9)]
                                                    _ = v207
                                                case 213:
                                                    var v208 int8 = 100
                                                    _ = v208
                                                default:
                                                    v209 := GenS21{F0: 271, F1: 2147483648}
                                                    var v210 int64 = v3
                                                    _ = v209
                                                    _ = v210
                                                }
                                            case 2147483647:
                                                if (!(v4 < 2147483647)) {
                                                    delete(v2, (fi161 & (v9 % ((v6) | 1))))
                                                    v211 := GenS18{F0: 1, F1: -32768, F2: (^(v6))}
                                                    _ = v211
                                                }
                                            case -9223372036854775808:
                                                if (65535 != 2147483647) {
                                                    var v212 uint16 = 1
                                                    switch ("a" + "") {
                                                    case "foo bar":
                                                        v213 := GenS10{F0: false}
                                                        _ = v213
                                                    case "\n\t":
                                                        var v214 float64 = v8
                                                        _ = v214
                                                    case "":
                                                        switch 32768 {
                                                        case 322:
                                                            v4 = 1
                                                        case 1:
                                                            v8 = v8
                                                        case -2147483648:
                                                            v215 := GenS26{F0: (v10 << ((0) & 63)), F1: 128}
                                                            var v216 uint32 = 4294967295
                                                            _ = v215
                                                            _ = v216
                                                        }
                                                    }
                                                    _ = v212
                                                } else {
                                                    switch ("" + "\n\t") {
                                                    case "foo bar":
                                                        var v217 uint32 = 42
                                                        _ = v217
                                                    case "\n\t":
                                                        _ = len(v5)
                                                    case "hello":
                                                        var v218 uint16 = 300
                                                        _ = v218
                                                    default:
                                                        v6 += v9
                                                    }
                                                }
                                            }
                                            var v219 float32 = 1.5
                                            var v220 uint64 = 42
                                            _ = fi205
                                            _ = v219
                                            _ = v220
                                        }
                                    } else {
                                        if true {
                                            var v221 uint16 = 42
                                            var v222 int16 = 1
                                            _ = v221
                                            _ = v222
                                        } else {
                                            v8 -= (1e10 - v8)
                                        }
                                        for fi223 := 0; fi223 < 14; fi223++ {
                                            var v224 int16 = 42
                                            var v225 int8 = 127
                                            _ = fi223
                                            _ = v224
                                            _ = v225
                                        }
                                    }
                                }
                            }
                            _ = v172
                        } else {
                            for fi226 := 0; fi226 < 20; fi226++ {
                                v227 := GenS4{F0: 65535, F1: 128}
                                if (-32768 == -1) {
                                    v228 := GenS8{F0: "", F1: 0, F2: 4294967295}
                                    var v229 int64 = (v3 / (((v3 ^ 1)) | 1))
                                    var v230 uint8 = 128
                                    v1 = 0
                                    _ = v228
                                    _ = v229
                                    _ = v230
                                }
                                _ = fi226
                                _ = v227
                            }
                        }
                    }
                    v231 := []int{(fi161 | (fi161 & 1)), v9, 4294967295, v6}
                    if false {
                        switch -443 {
                        case -1:
                            fmt.Println(v4)
                        case 9223372036854775807:
                            v4 = (v4 * v10)
                        case -2147483648:
                            fmt.Println(v1, v8)
                        default:
                            if true {
                                v232 := GenS27{F0: (true || false)}
                                _ = v232
                            } else {
                                if false {
                                    v233 := GenS28{F0: 43}
                                    _ = v233
                                } else {
                                    for fi234 := 0; fi234 < 5; fi234++ {
                                        var v235 float64 = (1e10 / (((1.5 * ((v8 + 1.5) - v8)) + 3.14159) * 0.0))
                                        _ = fi234
                                        _ = v235
                                    }
                                    v1 = 0
                                }
                                if len(v231) >= 2 {
                                    v231 = v231[1:len(v231)]
                                }
                            }
                        }
                    } else {
                        v8 = v8
                        if len(v231) >= 2 {
                            v231 = v231[1:len(v231)]
                        }
                        for fi236 := 0; fi236 < 5; fi236++ {
                            var v237 int32 = v1
                            var v238 bool = false
                            if (-32768 > -564) {
                                switch 9223372036854775807 {
                                case 2147483647:
                                    v239 := GenS18{F0: -291, F1: 717, F2: 32767}
                                    _ = v239
                                case 32768:
                                    v8 = v8
                                default:
                                    if (fi236 >= (fi236 * (^(fi236)))) {
                                        v4 += v4
                                    } else {
                                        if v238 {
                                            v6 = fi236
                                        }
                                        var v240 int = fi161
                                        _ = v240
                                    }
                                }
                            }
                            v241 := GenS29{F0: 2, F1: 18446744073709551615}
                            var v242 uint16 = 2
                            _ = fi236
                            _ = v237
                            _ = v238
                            _ = v241
                            _ = v242
                        }
                    }
                    _ = fi161
                    _ = v231
                }
                v243 := []int{}
                _ = v243
            }
        default:
            var v244 string = "a"
            _ = v244
        }
    case "":
        if false {
            v9 -= 32768
            if ((255 == 394) || (0.0 != 3.14159)) {
                fmt.Println(v6, v8)
                var v245 int8 = 42
                v246 := GenS7{F0: v245, F1: 1, F2: 0}
                _ = v245
                _ = v246
            }
        }
    default:
        v247 := GenS30{F0: 255}
        _ = v247
    }
    if true {
        for fi248 := 0; fi248 < 17; fi248++ {
            switch 734 {
            case 97:
                switch v6 {
                case 0:
                    switch (("a" + ("\n\t" + "hello")) + "foo bar") {
                    case "a":
                        fmt.Println(v1, v3, v3)
                    case "":
                        fmt.Println(v8)
                    case "hello":
                        v249 := GenS31{F0: "foo bar", F1: (v3 * 0)}
                        _ = v249
                    }
                }
            default:
                v6 += fi248
            }
            if true {
                switch "hello" {
                case "\n\t":
                    switch fi248 {
                    case 9223372036854775807:
                        if (!(!false)) {
                            var v250 uint64 = 292
                            v251 := GenS32{F0: 42, F1: 65535}
                            _ = v250
                            _ = v251
                        } else {
                            v8 = v8
                            for fi252 := 0; fi252 < 6; fi252++ {
                                v2[v6] = false
                                v253 := GenS9{F0: (v3 / ((-9223372036854775808) | 1)), F1: (("\n\t" + (("foo bar" + ("foo bar" + "hello")) + "")) < "\n\t"), F2: (0.0 >= 0.0)}
                                _ = fi252
                                _ = v253
                            }
                            var v254 uint32 = 42
                            _ = v254
                        }
                    case -2147483648:
                        if true {
                            switch ("" + "") {
                            case "a":
                                v6 = v6
                            default:
                                if true {
                                    var v255 int64 = v3
                                    _ = v255
                                }
                            }
                        }
                    case 4294967295:
                        v6 += (fi248 * v6)
                    default:
                        if (((true || (0 < 50)) && false) || ((v3 + v3) != v3)) {
                            if (-1 != -1) {
                                switch v6 {
                                case 32768:
                                    v256 := []int{-9223372036854775808, -567}
                                    _ = v256
                                case 128:
                                    v1 -= (^(v10))
                                default:
                                    v257 := GenS6{F0: 127, F1: 127}
                                    _ = v257
                                }
                                var v258 uint = 794
                                v4 += 2147483647
                                _ = v258
                            }
                            fmt.Println(v4, v4, v3)
                        } else {
                            v259 := v5[("a" + "\n\t")]
                            v260 := []int{v6, v6, (v9 << ((892) & 63)), 9223372036854775807}
                            _ = v259
                            _ = v260
                        }
                    }
                case "foo bar":
                    for fi261 := 0; fi261 < 17; fi261++ {
                        v262 := v2[-9223372036854775808]
                        switch v9 {
                        case -611:
                            v4 = (v10 >> ((65535) & 63))
                        case 32768:
                            var v263 uint64 = 42
                            _ = v263
                        default:
                            v264 := []int{v9}
                            _ = v264
                        }
                        v262 = ((-575 < 32768) && ((v8 + ((1.5 * v8) - 3.14159)) > 3.14159))
                        _ = fi261
                        _ = v262
                    }
                case "hello":
                    for fi265 := 0; fi265 < 20; fi265++ {
                        switch fi265 {
                        case 243:
                            fmt.Println(v6, v9)
                        case 325:
                            var v266 int32 = (v4 ^ (v1 << ((65535) & 63)))
                            _ = v266
                        case -1:
                            v8 += (0.1 - (v8 - (((v8 * float64(0.1)) + ((0.1 + v8) + v8)) - float64(0.0))))
                        }
                        v3 = (v3 % (((v3 * (v3 + v3))) | 1))
                        _ = fi265
                    }
                default:
                    if true {
                        if false {
                            switch (v6 & 1) {
                            case -164:
                                var v267 uint8 = 42
                                _ = v267
                            case 9223372036854775807:
                                v6 -= (-(fi248))
                            case 423:
                                if (float64(-2.25) <= v8) {
                                    if true {
                                        switch "a" {
                                        case "hello":
                                            switch v6 {
                                            case 9223372036854775807:
                                                var v268 int8 = -1
                                                _ = v268
                                            }
                                        case "a":
                                            var v269 uint64 = 403
                                            _ = v269
                                        default:
                                            for fi270 := 0; fi270 < 3; fi270++ {
                                                if false {
                                                    var v271 string = "\n\t"
                                                    _ = v271
                                                } else {
                                                    if (0.0 >= (3.14159 * -2.25)) {
                                                        v3 = -1
                                                        v272 := GenS2{F0: 507, F1: 0}
                                                        v8 = 0.1
                                                        fmt.Println(v8)
                                                        _ = v272
                                                    } else {
                                                        v8 = 1e10
                                                        v273 := GenS33{F0: (v10 + (-(v4))), F1: false}
                                                        _ = v273
                                                    }
                                                    if (2 <= 42) {
                                                        if false {
                                                            switch "" {
                                                            case "\n\t":
                                                                if true {
                                                                    v10 = (v4 & v10)
                                                                    var v274 uint64 = 18446744073709551615
                                                                    var v275 int = (-(fi248))
                                                                    _ = v274
                                                                    _ = v275
                                                                }
                                                            case "a":
                                                                _ = len(v5)
                                                            case "hello":
                                                                v9 -= v6
                                                            }
                                                            v276 := []int{(^(v6)), v9, v6}
                                                            if (!(!(((!(((1.5 - 3.14159) > 1.5) || true)) && ((false || (260 < 0)) && (v4 == (v10 - v4)))) && false))) {
                                                                v277 := GenS26{F0: (v1 | 1), F1: 0}
                                                                v278 := GenS34{F0: v4, F1: ("a" + "a")}
                                                                switch fi270 {
                                                                case -1:
                                                                    v1 -= v4
                                                                case -2147483648:
                                                                    v279 := GenS35{F0: 757}
                                                                    if (4294967295 < 0) {
                                                                        switch v6 {
                                                                        case -9223372036854775808:
                                                                            if false {
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        if (v8 != v8) {
                                                                            if true {
                                                                                var v280 uint16 = 42
                                                                                var v281 float64 = v8
                                                                                var v282 int64 = v3
                                                                                _ = v280
                                                                                _ = v281
                                                                                _ = v282
                                                                            }
                                                                        } else {
                                                                            if len(v276) > 0 {
                                                                                v276[0] = v6
                                                                            }
                                                                            v283 := v2[(fi248 >> ((0) & 63))]
                                                                            _ = v283
                                                                        }
                                                                        var v284 string = (("hello" + "\n\t") + "")
                                                                        v6 = -454
                                                                        _ = v284
                                                                    } else {
                                                                        v9 = -2147483648
                                                                        if true {
                                                                            fmt.Println(fi248)
                                                                            if true {
                                                                                var v285 float32 = (((((3.14159 + (-2.25 - 1e10)) + 1e10) - (float32(0.1) / -2.25)) - (0.1 - (1.5 + -2.25))) - 1.5)
                                                                                fmt.Println(v3, v10, v8)
                                                                                _ = v285
                                                                            } else {
                                                                                v4 -= v10
                                                                                if (!(false && false)) {
                                                                                    switch -9223372036854775808 {
                                                                                    case 1:
                                                                                        v276 = append(v276, 4294967295)
                                                                                    }
                                                                                } else {
                                                                                    switch (v6 + (fi248 * (v9 / (((v9 >> ((255) & 63))) | 1)))) {
                                                                                    case 127:
                                                                                        v286 := GenS18{F0: -32768, F1: 1, F2: fi270}
                                                                                        _ = v286
                                                                                    case 1:
                                                                                        if ((true && (128 < 1)) && false) {
                                                                                            var v287 float64 = v8
                                                                                            fmt.Println(fi270, v3)
                                                                                            _ = v287
                                                                                        } else {
                                                                                            v1 -= (v1 / ((1) | 1))
                                                                                        }
                                                                                    default:
                                                                                        var v288 int = (v6 + 32767)
                                                                                        _ = v288
                                                                                    }
                                                                                    v10 = (v10 * 0)
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v6, v8)
                                                                                }
                                                                                v289 := GenS36{F0: 42, F1: (v1 ^ v10), F2: v3}
                                                                                var v290 uint8 = 40
                                                                                _ = v289
                                                                                _ = v290
                                                                            }
                                                                        }
                                                                        switch ("foo bar" + (((("foo bar" + "foo bar") + ("\n\t" + "hello")) + ("\n\t" + "hello")) + (("hello" + ("" + ("hello" + ""))) + "\n\t"))) {
                                                                        case "hello":
                                                                            v10 += v4
                                                                        case "foo bar":
                                                                            v6 = fi270
                                                                        default:
                                                                            switch v9 {
                                                                            case -1:
                                                                                v291 := GenS19{F0: 1, F1: 1, F2: 255}
                                                                                _ = v291
                                                                            case 32767:
                                                                                var v292 uint16 = 483
                                                                                _ = v292
                                                                            default:
                                                                                v293 := GenS37{F0: v10, F1: true}
                                                                                _ = v293
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v279
                                                                default:
                                                                    var v294 uint = 18446744073709551615
                                                                    _ = v294
                                                                }
                                                                _ = v277
                                                                _ = v278
                                                            }
                                                            _ = v276
                                                        } else {
                                                            if true {
                                                                var v295 uint16 = 255
                                                                _ = v295
                                                            }
                                                            var v296 int = (v6 | v6)
                                                            var v297 uint16 = 540
                                                            _ = v296
                                                            _ = v297
                                                        }
                                                        v3 += (-(v3))
                                                        var v298 int16 = 0
                                                        _ = v298
                                                    } else {
                                                        _ = len(v2)
                                                        v299 := []int{(fi270 - 4294967295)}
                                                        var v300 uint32 = 2147483648
                                                        _ = v299
                                                        _ = v300
                                                    }
                                                    var v301 int64 = v3
                                                    if ((!((0 < v6) || (171 > 42))) && true) {
                                                        v302 := []int{(v6 + int(v301))}
                                                        var v303 uint = 4294967295
                                                        v304 := GenS38{F0: 525}
                                                        _ = v302
                                                        _ = v303
                                                        _ = v304
                                                    } else {
                                                        v305 := []int{v6, 131}
                                                        var v306 int64 = v301
                                                        delete(v5, ("a" + (("\n\t" + ("hello" + "a")) + "\n\t")))
                                                        _ = v305
                                                        _ = v306
                                                    }
                                                    _ = v301
                                                }
                                                fmt.Println(v8)
                                                _ = fi270
                                            }
                                        }
                                        if false {
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v307 := GenS2{F0: 2, F1: 380}
                                                    _ = v307
                                                case "hello":
                                                    for fi308 := 0; fi308 < 18; fi308++ {
                                                        if true {
                                                            fmt.Println(v6, v4)
                                                            fmt.Println(v1, v3, v3)
                                                            switch (fi308 + 0) {
                                                            case 9223372036854775807:
                                                                delete(v2, (fi308 ^ (-(v6))))
                                                            case 4294967295:
                                                                switch (-(fi308)) {
                                                                case 9223372036854775807:
                                                                    v6 += v6
                                                                case 128:
                                                                    var v309 uint8 = 128
                                                                    _ = v309
                                                                case 0:
                                                                    v310 := GenS39{F0: 42}
                                                                    _ = v310
                                                                default:
                                                                    var v311 uint32 = 2147483648
                                                                    _ = v311
                                                                }
                                                            case 2147483647:
                                                                var v312 uint16 = 32767
                                                                _ = v312
                                                            }
                                                        } else {
                                                            delete(v2, fi248)
                                                            delete(v5, ("hello" + "\n\t"))
                                                        }
                                                        _ = fi308
                                                    }
                                                case "":
                                                    v3 -= 1
                                                default:
                                                    var v313 int = 128
                                                    _ = v313
                                                }
                                                for fi314 := 0; fi314 < 6; fi314++ {
                                                    v4 = v1
                                                    if (((652 > 18446744073709551615) || (true && (42 >= -183))) && false) {
                                                        var v315 uint = 65535
                                                        if false {
                                                            var v316 uint64 = 0
                                                            fmt.Println(v9, fi248, v10)
                                                            _ = v316
                                                        } else {
                                                            var v317 uint16 = 625
                                                            fmt.Println(v6, v6)
                                                            fmt.Println(fi314, v317)
                                                            _ = v317
                                                        }
                                                        switch (("a" + ("\n\t" + "")) + (("\n\t" + ("" + "")) + "hello")) {
                                                        case "\n\t":
                                                            v4 = 2147483647
                                                            v10 += 2147483647
                                                            var v318 uint8 = 255
                                                            v3 = v3
                                                            _ = v318
                                                        case "a":
                                                            v319 := GenS26{F0: (v4 % ((v1) | 1)), F1: 55}
                                                            var v320 uint16 = 32767
                                                            var v321 uint = 4294967295
                                                            _ = v319
                                                            _ = v320
                                                            _ = v321
                                                        default:
                                                            var v322 uint = v315
                                                            _ = v322
                                                        }
                                                        _ = v315
                                                    } else {
                                                        var v323 uint16 = 847
                                                        _ = v323
                                                    }
                                                    v6 += v6
                                                    _ = fi314
                                                }
                                                for fi324 := 0; fi324 < 20; fi324++ {
                                                    delete(v5, "foo bar")
                                                    _ = fi324
                                                }
                                            }
                                            var v325 int64 = -1
                                            _ = v325
                                        }
                                    } else {
                                        v326 := GenS40{F0: (("foo bar" + "hello") + "a"), F1: v6, F2: false}
                                        _ = v326
                                    }
                                    v327 := []int{(v6 | fi248)}
                                    if (-2147483648 <= 0) {
                                        v328 := GenS39{F0: 1}
                                        _ = v328
                                    }
                                    _ = v327
                                }
                            default:
                                for fi329 := 0; fi329 < 2; fi329++ {
                                    var v330 bool = ((0 > 1) || (42 < 65535))
                                    if (float32(((v8 / float64((0.1 + 0.0))) - (-2.25 - v8))) >= 1.5) {
                                        v6 = 4294967295
                                    } else {
                                        switch fi248 {
                                        case 128:
                                            if ((1.5 - 0.1) == (0.0 * float32(v8))) {
                                                v6 = v6
                                                if v330 {
                                                    fmt.Println(v8, v3, v1)
                                                    v331 := GenS24{F0: v3}
                                                    _ = v331
                                                }
                                            }
                                        case 126:
                                            var v332 uint16 = 1
                                            _ = v332
                                        case 1:
                                            var v333 int32 = (v10 ^ (v4 & v10))
                                            _ = v333
                                        }
                                        v334 := GenS41{F0: v330, F1: v4, F2: 32767}
                                        _ = v334
                                    }
                                    fmt.Println(v6)
                                    _ = fi329
                                    _ = v330
                                }
                            }
                            for fi335 := 0; fi335 < 17; fi335++ {
                                if false {
                                    v336 := []int{}
                                    _ = v336
                                }
                                var v337 int32 = (v1 * v1)
                                if (v3 >= (^(v3))) {
                                    v3 += (v3 ^ -1)
                                    if false {
                                        v338 := GenS10{F0: false}
                                        v339 := GenS33{F0: v4, F1: true}
                                        _ = v338
                                        _ = v339
                                    } else {
                                        fmt.Println(v4, v8)
                                        switch "hello" {
                                        case "\n\t":
                                            switch v9 {
                                            case 1:
                                                var v340 int = (v9 + v9)
                                                _ = v340
                                            default:
                                                fmt.Println(v10)
                                            }
                                        }
                                        var v341 int = (v6 * -460)
                                        _ = v341
                                    }
                                    var v342 int64 = (v3 % ((v3) | 1))
                                    _ = v342
                                }
                                _ = fi335
                                _ = v337
                            }
                            fmt.Println(v1, v8, v6)
                        } else {
                            if (!((!true) && (3.14159 == 1.5))) {
                                var v343 uint64 = 2
                                var v344 float32 = ((-2.25 - -2.25) * ((0.0 - (0.0 - float32(v8))) + -2.25))
                                _ = v343
                                _ = v344
                            }
                        }
                    } else {
                        v345 := []int{fi248, -367}
                        v346 := GenS42{F0: v3}
                        _ = v345
                        _ = v346
                    }
                }
                v347 := GenS43{F0: (v1 | v10), F1: v6}
                if (1 != 854) {
                    v348 := []int{0, (-(fi248)), v9}
                    _ = v348
                }
                _ = v347
            } else {
                v6 = 127
            }
            var v349 bool = false
            _ = fi248
            _ = v349
        }
        var v350 uint16 = 255
        for fi351 := 0; fi351 < 16; fi351++ {
            v3 += (v3 + 0)
            v352 := []int{(-(fi351)), v6, 127}
            _ = len(v352)
            _ = fi351
            _ = v352
        }
        _ = v350
    }
    v4 = (v4 ^ v10)
    fmt.Println(v10, v9, v10)
    var v353 int8 = 100
    v354 := []int{(v9 >> ((65535) & 63))}
    var v355 int32 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v353:", v353)
    fmt.Println("slicelen v354:", len(v354), v354)
    fmt.Println("v355:", v355)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v353
    _ = v354
    _ = v355
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
