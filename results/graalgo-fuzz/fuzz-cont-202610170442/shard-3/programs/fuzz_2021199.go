package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int; F1 bool; F2 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 int; F2 int32 }
type GenS12 struct { F0 uint; F1 int64; F2 string }
type GenS13 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint16; F1 int32; F2 uint32 }

func main() {
    v1 := []string{"foo bar", ("a" + (("foo bar" + "\n\t") + (("\n\t" + ("a" + "\n\t")) + ("" + "\n\t")))), "hello", "foo bar"}
    var v2 int8 = 59
    var v3 string = "a"
    v4 := make(map[int]int)
    v5 := GenS1{F0: "foo bar"}
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int8 = v2
        var v8 float32 = (3.14159 + 3.14159)
        if true {
            v8 += 1.5
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 2
        if false {
            if true {
                var v11 uint16 = 42
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 int32 = -2147483648
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            v3 = (v3 + (v3 + v3))
        } else {
            v14 := GenS1{F0: "hello"}
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint8 = 0
                if false {
                    switch fi15 {
                    case -955:
                        var v17 int8 = (v2 >> ((42) & 63))
                        _ = v17
                    default:
                        switch ("hello" + (v3 + v3)) {
                        case "hello":
                            v18 := GenS1{F0: "hello"}
                            v19 := GenS1{F0: "a"}
                            switch (v3 + v3) {
                            case "\n\t":
                                v10 = v10
                                if (-489 != -1) {
                                    if true {
                                        fmt.Println(v16, v3, v10)
                                    }
                                    if (false || false) {
                                        _ = len(v4)
                                        if (1000 == 32767) {
                                            if true {
                                                fmt.Println(v2)
                                                fmt.Println(fi15)
                                            }
                                            if ("" != "") {
                                                var v20 float64 = (1e10 - float64(0.0))
                                                _ = v20
                                            }
                                            _ = len(v1)
                                            v21 := GenS2{F0: true, F1: 42}
                                            _ = v21
                                        } else {
                                            switch v3 {
                                            case "":
                                                fmt.Println(v10, v16, fi15)
                                                fmt.Println(fi9)
                                                if (v3 <= (v3 + (("\n\t" + v3) + (v3 + ("" + (("a" + v3) + v3)))))) {
                                                    var v22 bool = true
                                                    _ = v22
                                                } else {
                                                    v10 += 449
                                                    switch 128 {
                                                    case 127:
                                                        v23 := GenS1{F0: "foo bar"}
                                                        _ = v23
                                                    case 32768:
                                                        fmt.Println(v2)
                                                    case -71:
                                                        var v24 int8 = (v2 << ((65535) & 63))
                                                        _ = v24
                                                    }
                                                }
                                            }
                                        }
                                        v25 := GenS3{F0: v16}
                                        v26 := v4[fi9]
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v2 += (v2 & -67)
                                        v2 = v2
                                    }
                                    switch v3 {
                                    case "hello":
                                        var v27 uint32 = 1
                                        _ = v27
                                    case "\n\t":
                                        v16 += 42
                                        var v28 int16 = 0
                                        if false {
                                            fmt.Println(v2)
                                            _ = len(v1)
                                            fmt.Println(v3)
                                            _ = len(v4)
                                        } else {
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v29 int32 = -2147483648
                                                v30 := GenS2{F0: (fi9 != (fi9 * 487)), F1: 42}
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        _ = v28
                                    }
                                } else {
                                    fmt.Println(fi15, v16, v2)
                                    var v31 uint16 = 169
                                    if false {
                                        v32 := GenS4{F0: 32767, F1: 2147483648, F2: 2147483647}
                                        var v33 string = "a"
                                        if ((^(v31)) > v31) {
                                            v34 := v4[fi15]
                                            var v35 int8 = (v2 >> ((913) & 63))
                                            if (true || ((fi9 * (-(fi9))) == (-(fi9)))) {
                                                fmt.Println(v3, v31, v31)
                                            } else {
                                                v3 = (v3 + "hello")
                                            }
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if (0 > v31) {
                                                fmt.Println(v16, v2, v3)
                                                _ = len(v1)
                                                delete(v4, fi9)
                                                fmt.Println(v33, v16)
                                            }
                                        }
                                        _ = v32
                                        _ = v33
                                    }
                                    if true {
                                        var v36 int32 = 0
                                        _ = v36
                                    }
                                    if true {
                                        v2 -= (-(v2))
                                        switch (fi15 ^ 0) {
                                        case -1:
                                            v31 += v31
                                        case 1:
                                            v37 := GenS4{F0: 42, F1: 2147483648, F2: 4294967295}
                                            _ = v37
                                        }
                                    } else {
                                        fmt.Println(v2, fi9, v10)
                                    }
                                    v2 = (v2 % ((-33) | 1))
                                    if (!true) {
                                        var v38 uint64 = v10
                                        var v39 float32 = 0.0
                                        var v40 string = v3
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int = 32767
                                        v42 := GenS5{F0: (^(v31)), F1: 127, F2: 65535}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v31
                                }
                                var v43 int32 = int32(v10)
                                var v44 uint32 = 42
                                _ = v43
                                _ = v44
                            case "a":
                                switch (fi9 | fi9) {
                                case -936:
                                    var v45 int16 = -32768
                                    _ = v45
                                default:
                                    v16 = 255
                                }
                            }
                            _ = v18
                            _ = v19
                        case "\n\t":
                            v46 := GenS6{F0: 0}
                            _ = v46
                        case "":
                            if (v3 >= "hello") {
                                v3 = "hello"
                            }
                        default:
                            v16 -= (v16 << ((65535) & 63))
                        }
                    }
                    switch v3 {
                    case "":
                        var v47 int16 = 31
                        _ = v47
                    case "a":
                        v10 += v10
                    case "foo bar":
                        v3 = v3
                    }
                    v16 += (v16 + v16)
                } else {
                    var v48 uint16 = uint16(v16)
                    _ = v48
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v49 uint16 = 65535
        _ = fi9
        _ = v10
        _ = v49
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 9223372036854775807 {
    case 32767:
        v2 = -1
    case -277:
        var v50 int = -373
        _ = v50
    case -363:
        _ = len(v1)
    default:
        if false {
            var v51 float64 = 1e10
            v4[128] = 0
            for fi52 := 0; fi52 < 17; fi52++ {
                v2 += (-(v2))
                var v53 int16 = -312
                var v54 uint = 65535
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = v51
        }
    }
    v55 := GenS3{F0: 22}
    switch (v3 + ((v3 + ((v3 + "foo bar") + (v3 + ""))) + "hello")) {
    case "foo bar":
        switch 444 {
        case -2147483648:
            v2 += -128
        case 32768:
            var v56 string = v3
            _ = v56
        case 127:
            switch 42 {
            case 736:
                v3 = v3
            }
        default:
            switch -285 {
            case -994:
                v2 -= (v2 + v2)
            case 239:
                var v57 int = 127
                _ = v57
            case 9223372036854775807:
                var v58 int8 = (v2 - v2)
                _ = v58
            }
        }
    case "a":
        v2 = (v2 << ((40) & 63))
    case "\n\t":
        var v59 int16 = 41
        _ = v59
    default:
        if true {
            fmt.Println(v3, v3)
            var v60 string = (v3 + (v3 + v3))
            switch -9223372036854775808 {
            case 179:
                if (!true) {
                    v3 = v60
                    var v61 uint = 255
                    _ = v61
                } else {
                    if false {
                        v62 := GenS7{F0: 2147483647, F1: 129}
                        _ = v62
                    }
                    if true {
                        v60 = v60
                    } else {
                        switch v3 {
                        case "\n\t":
                            var v63 float32 = (3.14159 / ((0.1 / float32(float64((0.1 - 1.5)))) - (1.5 - 1e10)))
                            _ = v63
                        default:
                            if len(v1) > 0 {
                                v1[0] = v60
                            }
                        }
                        fmt.Println(v2, v2)
                    }
                }
            default:
                v64 := v4[4294967295]
                _ = v64
            }
            _ = v60
        }
    }
    var v65 uint64 = 4294967295
    for fi66 := 0; fi66 < 1; fi66++ {
        var v67 bool = (!(2 != 65535))
        v68 := GenS8{F0: (!(-2.25 != 1.5)), F1: 2}
        _ = fi66
        _ = v67
        _ = v68
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "":
        v4[-9223372036854775808] = 32767
    case "hello":
        _ = len(v4)
    case "a":
        if false {
            fmt.Println(v65, v65)
            v69 := GenS9{F0: 4294967295, F1: false, F2: -614}
            _ = v69
        }
    default:
        var v70 int8 = (v2 + (v2 + v2))
        _ = v70
    }
    switch v3 {
    case "a":
        switch 32768 {
        case 350:
            if false {
                var v71 uint64 = (v65 >> ((181) & 63))
                switch 127 {
                case 128:
                    fmt.Println(v2)
                case -2147483648:
                    var v72 uint32 = 285
                    _ = v72
                case 305:
                    switch (v3 + (((v3 + ("\n\t" + "foo bar")) + v3) + "\n\t")) {
                    case "\n\t":
                        if (float64(-2.25) != -2.25) {
                            switch 745 {
                            case 9223372036854775807:
                                v2 -= v2
                            case 4294967295:
                                if (v3 < v3) {
                                    for fi73 := 0; fi73 < 20; fi73++ {
                                        for fi74 := 0; fi74 < 19; fi74++ {
                                            v75 := GenS2{F0: false, F1: 2147483647}
                                            _ = fi74
                                            _ = v75
                                        }
                                        var v76 float32 = 0.0
                                        _ = fi73
                                        _ = v76
                                    }
                                    switch ("" + "a") {
                                    case "hello":
                                        v77 := GenS4{F0: -265, F1: 868, F2: 2}
                                        _ = v77
                                    case "":
                                        var v78 float64 = (0.1 + 0.0)
                                        _ = v78
                                    default:
                                        fmt.Println(v65, v71, v71)
                                    }
                                }
                            case -9223372036854775808:
                                for fi79 := 0; fi79 < 12; fi79++ {
                                    v65 += (v65 / ((731) | 1))
                                    fmt.Println(v71)
                                    _ = fi79
                                }
                            default:
                                v65 += (v71 - v65)
                            }
                        } else {
                            v65 = 42
                            fmt.Println(v3, v2)
                        }
                    case "":
                        if (2 != 23) {
                            var v80 float64 = 1.5
                            v80 = float64(((0.1 - (((((3.14159 * float32(v80)) + 0.0) / 3.14159) + -2.25) + 0.1)) * 1e10))
                            _ = v80
                        } else {
                            if false {
                                v81 := GenS4{F0: 0, F1: 2, F2: 65535}
                                if true {
                                    v82 := GenS10{F0: 211}
                                    if (v2 <= (v2 * 1)) {
                                        v3 = ("hello" + ((("\n\t" + v3) + v3) + v3))
                                    } else {
                                        v83 := GenS11{F0: (v2 * (v2 | v2)), F1: 9223372036854775807, F2: -20}
                                        for fi84 := 0; fi84 < 18; fi84++ {
                                            if false {
                                                for fi85 := 0; fi85 < 6; fi85++ {
                                                    _ = len(v1)
                                                    fmt.Println(v3, v65)
                                                    _ = fi85
                                                }
                                                var v86 float64 = float64((0.0 / 1.0))
                                                _ = v86
                                            } else {
                                                if true {
                                                    if true {
                                                        if (!false) {
                                                            var v87 int64 = 0
                                                            v88 := GenS9{F0: fi84, F1: (1.5 > 0.0), F2: 832}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        if false {
                                                            v89 := GenS5{F0: 94, F1: 127, F2: 407}
                                                            v2 = v2
                                                            _ = v89
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = v3
                                                            }
                                                            v3 = (((v3 + v3) + "\n\t") + v3)
                                                        }
                                                    } else {
                                                        v90 := GenS12{F0: 1, F1: 0, F2: ((v3 + v3) + "a")}
                                                        v3 = v3
                                                        if true {
                                                            if ((!(float64(3.14159) != -2.25)) || (false || (!false))) {
                                                                v91 := GenS13{F0: -343, F1: 221, F2: 255}
                                                                _ = v91
                                                            } else {
                                                                for fi92 := 0; fi92 < 9; fi92++ {
                                                                    var v93 float32 = float32(0.1)
                                                                    switch (fi92 / (((fi84 / (((fi92 - fi84)) | 1))) | 1)) {
                                                                    case 243:
                                                                        v94 := GenS14{F0: 0}
                                                                        _ = v94
                                                                    }
                                                                    if (true || false) {
                                                                        v71 += 1
                                                                        fmt.Println(fi92)
                                                                    } else {
                                                                        v65 = (v65 + v65)
                                                                        fmt.Println(fi92)
                                                                        v95 := GenS7{F0: fi84, F1: 2147483648}
                                                                        _ = v95
                                                                    }
                                                                    _ = fi92
                                                                    _ = v93
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "\n\t":
                                                                fmt.Println(v71, v65)
                                                            case "":
                                                                var v96 uint = 756
                                                                _ = v96
                                                            case "foo bar":
                                                                v2 -= -128
                                                            }
                                                            var v97 string = "foo bar"
                                                            _ = v97
                                                        }
                                                        _ = v90
                                                    }
                                                } else {
                                                    var v98 uint = 2
                                                    _ = v98
                                                }
                                            }
                                            v99 := GenS15{F0: 665}
                                            v100 := v4[-807]
                                            _ = fi84
                                            _ = v100
                                            _ = v99
                                        }
                                        _ = v83
                                    }
                                    for fi101 := 0; fi101 < 2; fi101++ {
                                        v102 := GenS9{F0: 127, F1: true, F2: -1}
                                        v103 := GenS16{F0: 95, F1: -277, F2: 2147483647}
                                        _ = fi101
                                        _ = v102
                                        _ = v103
                                    }
                                    _ = v82
                                } else {
                                    v71 = (v71 << ((0) & 63))
                                }
                                _ = v81
                            } else {
                                if false {
                                    v104 := GenS11{F0: v2, F1: -183, F2: 818}
                                    _ = v104
                                } else {
                                    v105 := GenS1{F0: v3}
                                    var v106 uint = 18446744073709551615
                                    v65 += 492
                                    _ = v105
                                    _ = v106
                                }
                                delete(v4, 973)
                            }
                            v3 = (v3 + (v3 + ""))
                        }
                    }
                }
                v107 := GenS12{F0: 1, F1: int64(v2), F2: v3}
                _ = v107
                _ = v71
            }
        default:
            var v108 float32 = (1e10 + (1.5 + 1.5))
            _ = v108
        }
    default:
        for fi109 := 0; fi109 < 4; fi109++ {
            v4[(fi109 & (-(fi109)))] = -692
            if (fi109 >= (fi109 + (fi109 & fi109))) {
                var v110 int16 = 1000
                delete(v4, (fi109 | fi109))
                fmt.Println(fi109, v3)
                _ = v110
            }
            _ = fi109
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v65
}
