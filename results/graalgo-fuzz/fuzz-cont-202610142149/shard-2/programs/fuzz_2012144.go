package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint8; F1 int8; F2 int32 }
type GenS9 struct { F0 int8; F1 uint16; F2 uint }
type GenS10 struct { F0 uint; F1 string; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    v5 := v4[v3]
    switch ("" + (v3 + v3)) {
    case "":
        switch v3 {
        case "":
            v3 = v3
        case "\n\t":
            v6 := []int{32768, 9223372036854775807}
            _ = v6
        case "hello":
            var v7 float32 = (1e10 + float32(-2.25))
            _ = v7
        default:
            for fi8 := 0; fi8 < 10; fi8++ {
                if (!(true || false)) {
                    v3 = (v3 + "\n\t")
                    fmt.Println(v2, fi8, v2)
                }
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch 42 {
                    case 1:
                        fmt.Println(v3)
                    case 42:
                        var v10 bool = true
                        _ = v10
                    default:
                        v5 += 971
                    }
                    if true {
                        v2 -= 2147483647
                        switch v3 {
                        case "foo bar":
                            v11 := []int{882}
                            _ = v11
                        case "a":
                            if false {
                                var v12 string = ("a" + v3)
                                if true {
                                    v13 := GenS1{F0: 158}
                                    fmt.Println(fi8, v3)
                                    _ = v13
                                } else {
                                    v5 += v5
                                    v14 := GenS2{F0: 0}
                                    fmt.Println(v5, v5)
                                    _ = v14
                                }
                                var v15 uint64 = 18
                                _ = v12
                                _ = v15
                            }
                        default:
                            if false {
                                var v16 int = fi9
                                var v17 int8 = 127
                                v3 = v3
                                _ = v16
                                _ = v17
                            }
                        }
                        if (42 == 743) {
                            v5 = (v5 + (v5 & v5))
                            v2 -= -748
                        } else {
                            v2 = v2
                        }
                        switch -561 {
                        case 9223372036854775807:
                            v18 := GenS3{F0: 150}
                            _ = v18
                        case -723:
                            var v19 int64 = 42
                            if false {
                                delete(v4, v3)
                            }
                            _ = v19
                        }
                    } else {
                        fmt.Println(v5, v3, v5)
                        var v20 string = v3
                        _ = len(v4)
                        _ = v20
                    }
                    _ = fi9
                }
                var v21 uint64 = 0
                _ = fi8
                _ = v21
            }
        }
    case "a":
        v3 = v3
    }
    fmt.Println(v5, v2)
    var v22 int8 = -1
    fmt.Println(v5)
    delete(v4, ((v3 + v3) + (v3 + v3)))
    if false {
        v5 = 0
        v23 := []int{436, 4294967295}
        var v24 int64 = -178
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                if (0 < 283) {
                    var v27 int32 = -1
                    switch fi26 {
                    case 2147483647:
                        if ((^(v22)) > (-(v22))) {
                            v5 = (v5 & v5)
                        }
                    case -1:
                        fmt.Println(v3, v2, v2)
                        var v28 uint16 = 2
                        _ = v28
                    case -2147483648:
                        if true {
                            switch fi26 {
                            case -1:
                                if false {
                                    v3 = v3
                                    v27 = 914
                                } else {
                                    v5 = v5
                                    v2 += v2
                                    if true {
                                        var v29 int64 = -1
                                        _ = v29
                                    } else {
                                        var v30 uint32 = (v5 + 0)
                                        _ = v30
                                    }
                                }
                            default:
                                var v31 uint64 = 792
                                _ = v31
                            }
                            _ = len(v4)
                        }
                    }
                    switch (fi25 >> ((2) & 63)) {
                    case -67:
                        v32 := GenS4{F0: 218}
                        _ = v32
                    default:
                        fmt.Println(v22)
                    }
                    _ = v27
                } else {
                    switch "foo bar" {
                    case "":
                        var v33 float64 = 1e10
                        _ = v33
                    }
                }
                _ = fi26
            }
            _ = fi25
        }
        var v34 uint8 = 255
        v35 := []int{}
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        if (!(-885 >= 0)) {
            v37 := GenS5{F0: 1, F1: -1, F2: 42}
            _ = v37
        } else {
            switch fi36 {
            case 2147483647:
                if false {
                    switch fi36 {
                    case 32767:
                        v38 := []int{fi36}
                        _ = v38
                    case 9223372036854775807:
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v3 = v3
                            delete(v1, v5)
                            v40 := GenS5{F0: 496, F1: v2, F2: 0}
                            _ = fi39
                            _ = v40
                        }
                    default:
                        if (-2.25 >= float64(3.14159)) {
                            v22 = (^(v22))
                            v41 := GenS2{F0: 171}
                            var v42 float32 = 0.1
                            _ = v41
                            _ = v42
                        }
                    }
                } else {
                    switch v3 {
                    case "":
                        if false {
                            fmt.Println(v2)
                            for fi43 := 0; fi43 < 3; fi43++ {
                                var v44 uint8 = 2
                                _ = fi43
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 20; fi45++ {
                                var v46 uint64 = 4294967295
                                v47 := GenS5{F0: 2, F1: 1, F2: 42}
                                v22 = (-(v22))
                                var v48 float32 = 1.5
                                _ = fi45
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                        } else {
                            var v49 float32 = ((-2.25 - 1.5) * (1.5 / (1.5 - 0.1)))
                            v50 := GenS3{F0: 675}
                            for fi51 := 0; fi51 < 11; fi51++ {
                                if (0.0 >= (0.1 + (-2.25 + float64(v49)))) {
                                    switch (fi36 + 128) {
                                    case 32767:
                                        fmt.Println(v49)
                                        v22 = v22
                                    case 1:
                                        v4[(v3 + v3)] = (^(v5))
                                    case 42:
                                        v5 -= (v5 + (v5 << ((335) & 63)))
                                    }
                                } else {
                                    v49 -= v49
                                    v22 += v22
                                }
                                if (4294967295 > (fi36 >> ((1) & 63))) {
                                    switch 2147483647 {
                                    case 0:
                                        var v52 uint64 = 774
                                        _ = v52
                                    }
                                    v53 := GenS1{F0: uint8(v5)}
                                    v22 += (v22 << ((255) & 63))
                                    _ = v53
                                }
                                v54 := GenS6{F0: 14, F1: (v22 | (v22 - (v22 - (v22 & v22)))), F2: (v3 + ((v3 + ((v3 + v3) + ((v3 + v3) + "a"))) + ((v3 + v3) + ("" + v3))))}
                                _ = fi51
                                _ = v54
                            }
                            _ = v49
                            _ = v50
                        }
                    case "\n\t":
                        for fi55 := 0; fi55 < 5; fi55++ {
                            var v56 int16 = 1000
                            _ = len(v1)
                            fmt.Println(v3)
                            _ = fi55
                            _ = v56
                        }
                    }
                    v57 := []int{-1, fi36}
                    for fi58 := 0; fi58 < 15; fi58++ {
                        if (35 >= -1) {
                            v3 = "a"
                            var v59 int = (fi36 % (((fi58 + (fi36 >> ((42) & 63)))) | 1))
                            _ = v59
                        } else {
                            v2 += v2
                            switch (v3 + v3) {
                            case "foo bar":
                                if (!false) {
                                    var v60 uint8 = 90
                                    fmt.Println(fi58, fi36, v2)
                                    _ = v60
                                } else {
                                    v61 := v4[""]
                                    v62 := GenS7{F0: -1, F1: 42}
                                    fmt.Println(v22)
                                    v57 = append(v57, 128)
                                    _ = v61
                                    _ = v62
                                }
                            case "a":
                                var v63 int8 = (v22 & (v22 - (v22 & (v22 << ((184) & 63)))))
                                _ = v63
                            }
                            v2 += v2
                        }
                        v5 = v5
                        switch v3 {
                        case "\n\t":
                            switch fi36 {
                            case 32767:
                                if false {
                                    fmt.Println(v5, v5, v3)
                                } else {
                                    if ((v2 >> ((65535) & 63)) <= (v2 * -1)) {
                                        v5 -= (v5 << ((598) & 63))
                                        if (true || false) {
                                            if true {
                                                switch fi58 {
                                                case -9223372036854775808:
                                                    var v64 float64 = 0.1
                                                    _ = v64
                                                case -762:
                                                    if (true || ((3.14159 / float32(1.5)) != (((float32(3.14159) - 1e10) * 0.1) / ((1.5 + 3.14159) - 0.0)))) {
                                                        if len(v57) > 0 {
                                                            v57[0] = fi36
                                                        }
                                                        v5 -= 0
                                                    }
                                                }
                                                v65 := GenS8{F0: 1, F1: v22, F2: (v2 >> ((58) & 63))}
                                                var v66 string = "hello"
                                                _ = v65
                                                _ = v66
                                            } else {
                                                v67 := GenS9{F0: (v22 & (v22 >> ((842) & 63))), F1: 65535, F2: 255}
                                                fmt.Println(v22)
                                                _ = v67
                                            }
                                            if (true || (64 > 179)) {
                                                switch "hello" {
                                                case "a":
                                                    var v68 uint = 42
                                                    _ = v68
                                                case "\n\t":
                                                    fmt.Println(v22)
                                                case "foo bar":
                                                    if (659 == 1) {
                                                        var v69 int32 = v2
                                                        fmt.Println(fi36)
                                                        _ = v69
                                                    }
                                                }
                                                var v70 float32 = (3.14159 * (3.14159 + 0.0))
                                                _ = v70
                                            }
                                            v22 = v22
                                        } else {
                                            if len(v57) >= 2 {
                                                v57 = v57[1:len(v57)]
                                            }
                                        }
                                    } else {
                                        var v71 uint32 = v5
                                        _ = v71
                                    }
                                    v22 += v22
                                }
                            case 9223372036854775807:
                                if true {
                                    var v72 bool = (2 >= 262)
                                    if v72 {
                                        v2 -= (v2 >> ((336) & 63))
                                    } else {
                                        fmt.Println(v22, v5, v3)
                                        if (42 > 42) {
                                            if false {
                                                fmt.Println(fi58, v22)
                                                var v73 uint16 = 475
                                                _ = v73
                                            } else {
                                                var v74 int8 = -123
                                                v72 = v72
                                                v22 -= (v22 * (v22 | (v22 % (((v22 ^ (v74 & int8(v5)))) | 1))))
                                                fmt.Println(v74, v72, fi58)
                                                v75 := GenS8{F0: 2, F1: (v22 ^ 0), F2: (v2 * v2)}
                                                _ = v74
                                                _ = v75
                                            }
                                            if (v72 || ((66 >= v22) || v72)) {
                                                v76 := GenS2{F0: 127}
                                                delete(v1, (^(v5)))
                                                if len(v57) > 0 {
                                                    v57[0] = fi36
                                                }
                                                _ = v76
                                            }
                                            switch v3 {
                                            case "":
                                                if v72 {
                                                    v3 = ("\n\t" + ("foo bar" + "\n\t"))
                                                }
                                                var v77 uint16 = 42
                                                _ = v77
                                            case "\n\t":
                                                v2 += v2
                                            }
                                            v78 := GenS6{F0: 188, F1: -1, F2: (v3 + v3)}
                                            _ = v78
                                        }
                                        v79 := GenS1{F0: 0}
                                        _ = v79
                                    }
                                    _ = v72
                                } else {
                                    if (!(42 > 4294967295)) {
                                        switch (fi58 | fi36) {
                                        case 32767:
                                            fmt.Println(v3, v2)
                                        case -946:
                                            v80 := GenS7{F0: -9223372036854775808, F1: v2}
                                            _ = v80
                                        default:
                                            v5 += (v5 >> ((65535) & 63))
                                        }
                                    } else {
                                        fmt.Println(v22)
                                        if false {
                                            var v81 float64 = 0.1
                                            _ = v81
                                        } else {
                                            fmt.Println(v3, fi36)
                                        }
                                        var v82 int8 = 100
                                        _ = v82
                                    }
                                    fmt.Println(v3)
                                }
                                v4[""] = v5
                                var v83 bool = true
                                _ = v83
                            default:
                                v84 := v4["foo bar"]
                                _ = v84
                            }
                        case "a":
                            var v85 int32 = v2
                            _ = v85
                        case "hello":
                            fmt.Println(v2)
                        default:
                            if ((false || true) && (v3 < v3)) {
                                v22 = (v22 << ((2) & 63))
                                v86 := GenS3{F0: 719}
                                _ = v86
                            } else {
                                v5 = v5
                                var v87 uint64 = 4294967295
                                _ = v87
                            }
                        }
                        var v88 bool = true
                        v5 = 2147483648
                        var v89 int32 = (v2 / (((v2 >> ((734) & 63))) | 1))
                        v4[v3] = 740
                        v90 := GenS1{F0: 2}
                        _ = fi58
                        _ = v88
                        _ = v89
                        _ = v90
                    }
                    _ = v57
                }
            case 32767:
                if (v3 != v3) {
                    if (1 >= 0) {
                        v91 := GenS4{F0: 0}
                        _ = v91
                    } else {
                        var v92 int16 = 765
                        if true {
                            var v93 int8 = v22
                            _ = v93
                        } else {
                            var v94 float64 = 3.14159
                            if (!((v22 << ((905) & 63)) == v22)) {
                                v2 = v2
                                var v95 uint16 = 255
                                _ = v95
                            }
                            _ = v94
                        }
                        _ = v92
                    }
                    var v96 uint32 = v5
                    if false {
                        for fi97 := 0; fi97 < 19; fi97++ {
                            if true {
                                switch v3 {
                                case "\n\t":
                                    switch "hello" {
                                    case "\n\t":
                                        var v98 bool = false
                                        _ = v98
                                    case "hello":
                                        if false {
                                            v22 += (v22 >> ((18446744073709551615) & 63))
                                            v3 = ""
                                            v4[v3] = 410
                                        }
                                        var v99 int32 = (-(v2))
                                        switch "a" {
                                        case "":
                                            if true {
                                                if false {
                                                    var v100 string = v3
                                                    var v101 float64 = (-2.25 - 1e10)
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    fmt.Println(v96)
                                                }
                                            } else {
                                                fmt.Println(v22)
                                                if (1000 <= 943) {
                                                    if (((1.5 * 0.1) - float32((3.14159 * 3.14159))) == (float32(3.14159) * 0.1)) {
                                                        v99 += v99
                                                    }
                                                    var v102 int8 = 42
                                                    _ = v102
                                                } else {
                                                    v99 = (v2 << ((493) & 63))
                                                    var v103 int = (fi36 - 9223372036854775807)
                                                    _ = v103
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v2)
                                        }
                                        v104 := GenS10{F0: 4294967295, F1: "a", F2: -1}
                                        _ = v104
                                        _ = v99
                                    }
                                case "foo bar":
                                    if true {
                                        v3 = "\n\t"
                                        v105 := GenS6{F0: 128, F1: v22, F2: (("foo bar" + ("" + "hello")) + v3)}
                                        _ = v105
                                    } else {
                                        var v106 int8 = (v22 * (^(v22)))
                                        var v107 uint16 = 32767
                                        if false {
                                            v22 += (v106 | -91)
                                            var v108 int16 = 32767
                                            _ = v108
                                        }
                                        _ = v106
                                        _ = v107
                                    }
                                default:
                                    var v109 int32 = 42
                                    _ = v109
                                }
                            } else {
                                var v110 int8 = v22
                                fmt.Println(v2, v110, fi36)
                                var v111 uint64 = 0
                                _ = v110
                                _ = v111
                            }
                            _ = fi97
                        }
                    } else {
                        var v112 uint = 4294967295
                        _ = v112
                    }
                    _ = v96
                }
            default:
                var v113 uint64 = 2
                _ = v113
            }
            fmt.Println(v3, v5)
            v2 -= v2
        }
        v114 := GenS5{F0: 1, F1: 2147483647, F2: 9223372036854775807}
        _ = fi36
        _ = v114
    }
    for fi115 := 0; fi115 < 19; fi115++ {
        var v116 int16 = 1000
        _ = fi115
        _ = v116
    }
    for fi117 := 0; fi117 < 3; fi117++ {
        v4[v3] = v5
        v22 += (v22 / ((v22) | 1))
        for fi118 := 0; fi118 < 16; fi118++ {
            v119 := []int{(fi118 * (fi118 / ((fi117) | 1))), (^(fi118)), (fi118 | fi118)}
            _ = fi118
            _ = v119
        }
        _ = fi117
    }
    var v120 uint32 = v5
    fmt.Println(v120, v2, v2)
    var v121 int64 = -58
    v122 := []int{1, -1, 0}
    v123 := GenS2{F0: 128}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v120
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
