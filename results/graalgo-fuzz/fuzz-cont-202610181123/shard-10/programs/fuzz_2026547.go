package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int16; F2 int32 }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 int32; F1 int8; F2 int64 }
func genFn0(p0 string, p1 uint64, p2 uint16) bool {
    v5 := GenS1{F0: true}
    var v6 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (!(p1 != 1))
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{933}
    var v3 uint = 255
    var v4 bool = false
    v7 := GenS1{F0: ((false || v4) && false)}
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint32 = 0
            _ = fi9
            _ = v10
        }
        if v4 {
            for fi11 := 0; fi11 < 3; fi11++ {
                if false {
                    var v12 int32 = 504
                    _ = v12
                }
                if v4 {
                    v2 = append(v2, 2)
                    var v13 bool = (v3 == v3)
                    fmt.Println(v1, v4)
                    _ = v13
                } else {
                    if v4 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v3, fi8)
                    } else {
                        if v4 {
                            v3 = 1
                            var v14 uint8 = 2
                            if v4 {
                                if (0 >= 9223372036854775807) {
                                    switch "\n\t" {
                                    case "hello":
                                        v15 := GenS1{F0: v4}
                                        _ = v15
                                    case "":
                                        v16 := GenS1{F0: (42 != 0)}
                                        _ = v16
                                    }
                                    switch fi8 {
                                    case 32768:
                                        var v17 int32 = 0
                                        _ = v17
                                    case 0:
                                        switch "" {
                                        case "hello":
                                            switch ("a" + "hello") {
                                            case "a":
                                                v18 := GenS1{F0: v4}
                                                _ = v18
                                            default:
                                                v4 = false
                                            }
                                        case "":
                                            v3 = v3
                                        default:
                                            switch fi8 {
                                            case 9223372036854775807:
                                                var v19 string = "hello"
                                                _ = v19
                                            case -282:
                                                var v20 float64 = 3.14159
                                                _ = v20
                                            case -320:
                                                v21 := GenS2{F0: 134, F1: 2, F2: 538}
                                                fmt.Println(v3, fi11)
                                                _ = v21
                                            }
                                        }
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch ("a" + "hello") {
                                case "foo bar":
                                    fmt.Println(fi11, v1)
                                    if ("hello" <= (("" + ("foo bar" + ("" + "foo bar"))) + ("foo bar" + ("hello" + "a")))) {
                                        if false {
                                            v22 := GenS3{F0: (v14 + (v14 | (^(v14))))}
                                            if v4 {
                                                var v23 uint16 = v1
                                                _ = v23
                                            }
                                            fmt.Println(v4, fi11)
                                            _ = v22
                                        } else {
                                            var v24 float32 = float32((float64((1e10 - 1.5)) - 0.1))
                                            _ = v24
                                        }
                                        var v25 bool = v4
                                        _ = v25
                                    }
                                    switch (^(fi8)) {
                                    case 127:
                                        if v4 {
                                            v1 = 255
                                            var v26 int16 = 674
                                            _ = v26
                                        } else {
                                            _ = len(v2)
                                            if true {
                                                var v27 int16 = 0
                                                _ = v27
                                            }
                                            if (-887 <= 9223372036854775807) {
                                                switch "a" {
                                                case "foo bar":
                                                    v3 = (v3 - uint(v1))
                                                case "\n\t":
                                                    var v28 int16 = 384
                                                    _ = v28
                                                case "a":
                                                    v3 -= (v3 | v3)
                                                }
                                                var v29 string = ""
                                                fmt.Println(fi11, v29, v1)
                                                v30 := GenS4{F0: 28}
                                                fmt.Println(fi11, v4)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v31 := GenS4{F0: v14}
                                                    _ = v31
                                                }
                                                v32 := GenS3{F0: (v14 | 127)}
                                                var v33 int = fi11
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    case 42:
                                        var v34 int32 = 1
                                        _ = v34
                                    case 0:
                                        var v35 uint = 133
                                        _ = v35
                                    }
                                case "\n\t":
                                    var v36 uint64 = 340
                                    _ = v36
                                case "hello":
                                    v37 := GenS4{F0: v14}
                                    _ = v37
                                default:
                                    v3 += uint(v14)
                                }
                            }
                            var v38 uint32 = 0
                            var v39 int = (fi8 * fi8)
                            _ = v14
                            _ = v38
                            _ = v39
                        }
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                    }
                    _ = len(v2)
                    if v4 {
                        var v40 uint8 = 6
                        _ = v40
                    }
                }
                switch "\n\t" {
                case "hello":
                    v3 = v3
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                case "\n\t":
                    v41 := GenS4{F0: 255}
                    _ = v41
                }
                _ = fi11
            }
            v4 = v4
            var v42 bool = (!false)
            _ = v42
        } else {
            if (631 < (v3 & v3)) {
                fmt.Println(v3, v1)
                v1 = v1
                if v4 {
                    for fi43 := 0; fi43 < 19; fi43++ {
                        var v44 int16 = 353
                        var v45 float32 = 0.0
                        if v4 {
                            var v46 int16 = (v44 | v44)
                            switch "hello" {
                            case "foo bar":
                                _ = len(v2)
                            }
                            v44 = (^(v46))
                            _ = v46
                        } else {
                            var v47 uint = 1
                            if v4 {
                                if ("\n\t" > "hello") {
                                    fmt.Println(v44, v47)
                                    var v48 int32 = 2147483647
                                    _ = v48
                                } else {
                                    switch (fi8 + 2147483647) {
                                    case -2147483648:
                                        v3 = (v3 | v3)
                                    case 4294967295:
                                        var v49 int16 = 1
                                        _ = v49
                                    default:
                                        v4 = true
                                    }
                                }
                            }
                            v4 = true
                            _ = v47
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
            }
            v2 = append(v2, 2147483648)
            v4 = (-751 > 2147483647)
        }
        fmt.Println(v1)
        _ = fi8
    }
    fmt.Println(v3, v4, v4)
    for fi50 := 0; fi50 < 2; fi50++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi51 := 0; fi51 < 15; fi51++ {
            v52 := GenS3{F0: 88}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi51
            _ = v52
        }
        v2 = append(v2, 2147483648)
        _ = fi50
    }
    if v4 {
        v4 = false
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        v4 = v4
        v54 := GenS4{F0: 156}
        v55 := GenS2{F0: 469, F1: 1, F2: 790}
        _ = fi53
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 16; fi56++ {
        if v4 {
            v57 := GenS5{F0: v1, F1: 9223372036854775807}
            if v4 {
                var v58 bool = v4
                _ = v58
            } else {
                var v59 int64 = 0
                if v4 {
                    var v60 bool = v4
                    _ = v60
                } else {
                    fmt.Println(v59, fi56)
                    v61 := GenS3{F0: 127}
                    _ = v61
                }
                v62 := GenS6{F0: int8(v3)}
                _ = v59
                _ = v62
            }
            _ = v57
        } else {
            v63 := GenS2{F0: 0, F1: 42, F2: 991}
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = v63
        }
        v3 -= (v3 >> ((18446744073709551615) & 63))
        _ = fi56
    }
    v64 := GenS4{F0: 100}
    if ((v4 || (!v4)) || true) {
        for fi65 := 0; fi65 < 4; fi65++ {
            var v66 float32 = 0.1
            _ = fi65
            _ = v66
        }
    }
    v2 = append(v2, 4294967295)
    fmt.Println(v1)
    v67 := GenS7{F0: 2147483648, F1: -32768, F2: 1}
    for fi68 := 0; fi68 < 1; fi68++ {
        if (0.0 == (3.14159 - 1e10)) {
            for fi69 := 0; fi69 < 11; fi69++ {
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    v70 := GenS2{F0: 2147483648, F1: 0, F2: 4294967295}
                    _ = v70
                default:
                    fmt.Println(v3, v1)
                }
                var v71 int8 = 100
                if true {
                    v3 = v3
                    v72 := GenS7{F0: 1, F1: -600, F2: 1}
                    fmt.Println(v3)
                    _ = v72
                }
                if v4 {
                    v71 -= v71
                    fmt.Println(v3, v71, v4)
                } else {
                    v71 = v71
                    if (!v4) {
                        if v4 {
                            var v73 uint32 = 627
                            if v4 {
                                switch "foo bar" {
                                case "":
                                    _ = len(v2)
                                case "hello":
                                    v4 = (1 < -1)
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            var v74 int64 = -1
                            _ = v73
                            _ = v74
                        } else {
                            if (false || (v71 != (v71 % ((v71) | 1)))) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch -144 {
                                case -653:
                                    var v75 string = "a"
                                    _ = v75
                                }
                            }
                            var v76 uint32 = 935
                            switch "\n\t" {
                            case "hello":
                                var v77 uint16 = 473
                                _ = v77
                            default:
                                v76 = 154
                                var v78 uint32 = v76
                                if v4 {
                                    if v4 {
                                        if (!((!v4) || v4)) {
                                            if ((v3 % ((42) | 1)) != 2) {
                                                var v79 uint64 = 18446744073709551615
                                                _ = v79
                                            } else {
                                                v76 -= uint32(v1)
                                            }
                                        } else {
                                            v80 := GenS2{F0: v76, F1: uint8(v78), F2: 0}
                                            switch 128 {
                                            case 32768:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case 497:
                                                v81 := GenS3{F0: 127}
                                                _ = v81
                                            case -9223372036854775808:
                                                var v82 int32 = -159
                                                _ = v82
                                            default:
                                                switch 127 {
                                                case 764:
                                                    if true {
                                                        fmt.Println(v76)
                                                        switch 704 {
                                                        case -9223372036854775808:
                                                            v1 -= v1
                                                        case -1:
                                                            var v83 int8 = (v71 ^ v71)
                                                            _ = v83
                                                        case 128:
                                                            var v84 int = -2147483648
                                                            _ = v84
                                                        default:
                                                            var v85 uint = v3
                                                            _ = v85
                                                        }
                                                        v86 := GenS1{F0: true}
                                                        _ = v86
                                                    } else {
                                                        v71 += v71
                                                        var v87 int32 = 253
                                                        v3 += (v3 / (((v3 ^ v3)) | 1))
                                                        _ = v87
                                                    }
                                                case 631:
                                                    fmt.Println(v71, v76)
                                                case 42:
                                                    var v88 bool = v4
                                                    _ = v88
                                                default:
                                                    var v89 string = ("foo bar" + "foo bar")
                                                    _ = v89
                                                }
                                            }
                                            if (-2147483648 > 342) {
                                                if len(v2) > 0 {
                                                    v2[0] = uint32(v3)
                                                }
                                                var v90 string = ("a" + "\n\t")
                                                _ = v90
                                            }
                                            _ = v80
                                        }
                                        switch ("\n\t" + ((("a" + (("hello" + ("hello" + "\n\t")) + (((((("hello" + "a") + "hello") + "") + ("hello" + "hello")) + "\n\t") + (("\n\t" + "hello") + "foo bar")))) + "hello") + ("hello" + ("foo bar" + ((("hello" + ("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + "a")))) + "hello") + "\n\t"))))) {
                                        case "foo bar":
                                            switch "hello" {
                                            case "a":
                                                if true {
                                                    var v91 uint32 = v78
                                                    var v92 uint64 = 491
                                                    _ = v91
                                                    _ = v92
                                                }
                                            case "foo bar":
                                                v78 -= (v76 << ((v3) & 63))
                                            case "hello":
                                                v76 -= 65535
                                            }
                                        }
                                    }
                                } else {
                                    var v93 uint8 = 2
                                    _ = v93
                                }
                                fmt.Println(fi68)
                                var v94 uint32 = (v76 % ((v78) | 1))
                                _ = v78
                                _ = v94
                            }
                            var v95 int = fi69
                            _ = v76
                            _ = v95
                        }
                    }
                    var v96 uint16 = 65535
                    v97 := GenS8{F0: (v96 << ((1) & 63)), F1: 0}
                    v4 = v4
                    _ = v96
                    _ = v97
                }
                _ = fi69
                _ = v71
            }
            v4 = v4
            if (80 != 7) {
                var v98 uint8 = 42
                _ = v98
            } else {
                for fi99 := 0; fi99 < 6; fi99++ {
                    var v100 int = 42
                    if (42 >= -1) {
                        switch (-(fi99)) {
                        case 2147483647:
                            var v101 int = (v100 << (((v3 / ((v3) | 1))) & 63))
                            _ = v101
                        case 127:
                            v102 := GenS5{F0: v1, F1: 1}
                            _ = v102
                        default:
                            var v103 int = -162
                            _ = v103
                        }
                        var v104 bool = (!(!((true && v4) && (1 >= 2147483647))))
                        _ = len(v2)
                        _ = v104
                    }
                    fmt.Println(v1, v3, fi99)
                    _ = fi99
                    _ = v100
                }
            }
        }
        if (0.0 != (3.14159 * float64(((((1.5 * 0.0) * float32(1e10)) - -2.25) - 0.0)))) {
            for fi105 := 0; fi105 < 15; fi105++ {
                fmt.Println(v4)
                var v106 uint = v3
                _ = fi105
                _ = v106
            }
            v107 := GenS8{F0: 953, F1: 0}
            if ((179 <= uint8(v1)) && true) {
                v108 := GenS4{F0: uint8(v1)}
                var v109 int = (fi68 + fi68)
                v1 = (v1 * v1)
                _ = v108
                _ = v109
            }
            _ = v107
        } else {
            v110 := GenS2{F0: 4294967295, F1: uint8(v1), F2: 2}
            switch fi68 {
            case 1:
                if ((v1 / ((v1) | 1)) >= 2) {
                    switch "a" {
                    case "hello":
                        var v111 uint16 = v1
                        _ = v111
                    case "foo bar":
                        v4 = true
                    case "":
                        var v112 uint8 = 42
                        _ = v112
                    }
                } else {
                    switch fi68 {
                    case 594:
                        var v113 uint64 = 1
                        _ = v113
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v114 float64 = 1.5
                        _ = v114
                    case "a":
                        v3 = v3
                    case "":
                        v3 -= 1
                    default:
                        var v115 string = "foo bar"
                        _ = v115
                    }
                }
            case 128:
                v116 := GenS9{F0: (fi68 + fi68), F1: (!(98 < int32(v3)))}
                _ = v116
            default:
                var v117 uint64 = 4294967295
                _ = v117
            }
            _ = v110
        }
        v1 = (v1 & 42)
        _ = fi68
    }
    fmt.Println(v4)
    for fi118 := 0; fi118 < 13; fi118++ {
        v4 = (32767 != -119)
        switch fi118 {
        case -2147483648:
            v119 := GenS10{F0: 42, F1: 57, F2: 863}
            _ = v119
        default:
            v3 += v3
        }
        _ = fi118
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v67
    _ = v7
}
