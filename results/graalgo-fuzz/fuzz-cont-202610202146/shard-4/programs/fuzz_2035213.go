package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int; F2 int64 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint8; F1 int; F2 int }
type GenS13 struct { F0 bool; F1 bool }
type GenS14 struct { F0 int64; F1 uint32 }
type GenS15 struct { F0 uint32 }

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{65535, 2147483647, 905}
    var v3 int32 = -973
    var v4 float64 = (0.0 * 1.5)
    var v5 int = 4294967295
    _ = len(v2)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int64 = -128
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 990, F1: 974, F2: 252}
            fmt.Println(v7)
            fmt.Println(fi6, v5)
            _ = fi8
            _ = v9
        }
        switch (("a" + "\n\t") + "") {
        case "\n\t":
            var v10 string = ""
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v3, v3)
    for fi11 := 0; fi11 < 6; fi11++ {
        switch (("\n\t" + "a") + ("hello" + (("" + "\n\t") + "hello"))) {
        case "a":
            if true {
                for fi12 := 0; fi12 < 2; fi12++ {
                    v3 -= v1
                    switch 32767 {
                    case 42:
                        var v13 uint32 = 65535
                        _ = v13
                    case 32767:
                        v14 := GenS2{F0: 65535, F1: -14, F2: 255}
                        _ = v14
                    }
                    if (!(!(true || false))) {
                        var v15 uint8 = 20
                        _ = v15
                    } else {
                        var v16 uint = 0
                        _ = v16
                    }
                    v17 := GenS3{F0: 2, F1: 504}
                    switch "\n\t" {
                    case "":
                        v2 = append(v2, 1)
                    default:
                        switch v5 {
                        case 197:
                            var v18 uint16 = 32767
                            _ = v18
                        default:
                            v5 -= fi12
                        }
                    }
                    var v19 uint32 = 42
                    _ = fi12
                    _ = v17
                    _ = v19
                }
                var v20 int16 = -293
                var v21 uint64 = 290
                _ = v20
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 string = ("" + ((((("foo bar" + "\n\t") + "\n\t") + "hello") + "hello") + "a"))
                if false {
                    if (42 >= -32768) {
                        fmt.Println(v3)
                        if (0 >= 1) {
                            v2 = append(v2, 244)
                            fmt.Println(v1, fi11, fi11)
                        } else {
                            if (2147483647 > 1) {
                                if true {
                                    var v24 uint = 18446744073709551615
                                    _ = v24
                                } else {
                                    var v25 int64 = -415
                                    _ = v25
                                }
                                if true {
                                    v4 = 1.5
                                    v4 = (float64(float32(v4)) - (float64(1.5) * v4))
                                } else {
                                    fmt.Println(fi11, fi22, v4)
                                    v26 := GenS2{F0: 42, F1: 127, F2: 2}
                                    var v27 int64 = -1
                                    _ = v26
                                    _ = v27
                                }
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v1 = -1
                                }
                            } else {
                                if (219 > 2) {
                                    var v28 uint64 = 1
                                    _ = v28
                                }
                                fmt.Println(v3)
                                var v29 uint = 65535
                                _ = v29
                            }
                            v5 -= v5
                            v30 := GenS3{F0: 893, F1: 0}
                            _ = v30
                        }
                        if true {
                            fmt.Println(v5)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = 325
                                }
                            } else {
                                v3 = v1
                                v3 -= v1
                                v5 = (-(fi11))
                                v23 = v23
                                fmt.Println(v5, v4)
                            }
                            v1 += (v3 * v3)
                        }
                    } else {
                        if true {
                            var v31 int16 = -32768
                            if (100 == 127) {
                                v1 += v1
                                v32 := GenS4{F0: 183}
                                _ = v32
                            }
                            if true {
                                if ("" <= v23) {
                                    v31 = (v31 - v31)
                                    v33 := GenS1{F0: (v31 % ((int16(v1)) | 1)), F1: 2147483648, F2: 0}
                                    if true {
                                        switch v23 {
                                        case "":
                                            _ = len(v2)
                                        case "\n\t":
                                            if (65535 >= 4294967295) {
                                                if (fi11 == 9223372036854775807) {
                                                    v3 = v3
                                                    v34 := GenS5{F0: -1, F1: (v31 * v31)}
                                                    v35 := GenS6{F0: 0, F1: 81}
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v4 = 0.1
                                                }
                                                var v36 uint16 = 2
                                                fmt.Println(fi11, v1)
                                                _ = v36
                                            }
                                        }
                                        var v37 int64 = -9223372036854775808
                                        _ = v37
                                    } else {
                                        v1 = (v3 ^ v1)
                                        var v38 int8 = -1
                                        v39 := GenS7{F0: (v1 + (v1 % (((v1 % ((v1) | 1))) | 1))), F1: (-(fi22))}
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v33
                                } else {
                                    v23 = (("hello" + "a") + v23)
                                    _ = len(v2)
                                }
                                v40 := GenS8{F0: 2, F1: 0}
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                    v41 := GenS9{F0: false, F1: 0}
                                    v42 := GenS9{F0: ((v3 >> ((1) & 63)) < (v1 | 42)), F1: 549}
                                    _ = v41
                                    _ = v42
                                } else {
                                    _ = len(v2)
                                }
                                _ = v40
                            } else {
                                var v43 string = ("hello" + v23)
                                _ = v43
                            }
                            _ = v31
                        } else {
                            fmt.Println(v5, v5)
                            v3 -= (v1 & v3)
                            v4 += v4
                            if (true || ((((v4 * (v4 * v4)) - v4) / v4) != v4)) {
                                var v44 int = v5
                                fmt.Println(v23)
                                fmt.Println(fi11)
                                v45 := GenS1{F0: 0, F1: 2147483647, F2: 35}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v46 int64 = -1
                                var v47 uint32 = 0
                                fmt.Println(v3, fi11)
                                switch v23 {
                                case "":
                                    var v48 int64 = (v46 - 9223372036854775807)
                                    _ = v48
                                default:
                                    var v49 int = (fi11 & 127)
                                    _ = v49
                                }
                                fmt.Println(v1, v44, v44)
                                v50 := GenS2{F0: 177, F1: 42, F2: 127}
                                _ = v44
                                _ = v45
                                _ = v46
                                _ = v47
                                _ = v50
                            }
                            var v51 int16 = 32767
                            _ = v51
                        }
                    }
                    var v52 int32 = v1
                    _ = v52
                } else {
                    if (1 == -128) {
                        if true {
                            v5 = (^(v5))
                            fmt.Println(fi11, v3, fi22)
                        } else {
                            if (2147483647 > 496) {
                                if (false || (42 == 0)) {
                                    v3 += v1
                                    var v53 int = 32767
                                    if true {
                                        fmt.Println(v1, v5)
                                    }
                                    var v54 float32 = -2.25
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                var v55 float64 = v4
                                var v56 int = -821
                                var v57 uint64 = 42
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        }
                        switch ("\n\t" + "a") {
                        case "hello":
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi11, v4, v23)
                                switch 210 {
                                case 424:
                                    v3 += v3
                                default:
                                    if (false && (1000 != -1)) {
                                        switch 4294967295 {
                                        case 968:
                                            if (1 == 4294967295) {
                                                if (117 > 192) {
                                                    var v58 int32 = -2147483648
                                                    var v59 int = 127
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    if true {
                                                        var v60 float64 = v4
                                                        fmt.Println(v1, v3, v23)
                                                        _ = v60
                                                    }
                                                    var v61 float64 = (v4 - 3.14159)
                                                    if true {
                                                        v62 := GenS2{F0: 42, F1: 0, F2: 0}
                                                        var v63 bool = true
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        fmt.Println(v1, v61)
                                                        var v64 bool = true
                                                        switch (-(fi22)) {
                                                        case 756:
                                                            var v65 bool = v64
                                                            _ = v65
                                                        default:
                                                            var v66 uint8 = 42
                                                            _ = v66
                                                        }
                                                        _ = v64
                                                    }
                                                    _ = v61
                                                }
                                                v67 := GenS5{F0: -128, F1: 32767}
                                                _ = v67
                                            } else {
                                                v68 := GenS10{F0: 2, F1: fi22, F2: -648}
                                                _ = v68
                                            }
                                        case -110:
                                            if false {
                                                var v69 int8 = 78
                                                _ = v69
                                            } else {
                                                fmt.Println(fi22)
                                            }
                                        case -374:
                                            switch fi11 {
                                            case 128:
                                                v70 := GenS11{F0: ((v1 > (v1 & v3)) || false), F1: 9223372036854775807}
                                                _ = v70
                                            case -9223372036854775808:
                                                switch v23 {
                                                case "":
                                                    var v71 int32 = v3
                                                    _ = v71
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                            }
                                            v72 := GenS12{F0: 2, F1: fi11, F2: (-(v5))}
                                            _ = v72
                                        }
                                        var v73 uint64 = 895
                                        if true {
                                            var v74 float32 = 0.0
                                            v5 = (v5 & 4294967295)
                                            _ = v74
                                        } else {
                                            v75 := GenS3{F0: (v73 ^ v73), F1: 930}
                                            fmt.Println(v73, v1, fi11)
                                            _ = v75
                                        }
                                        v3 -= (v1 % ((v1) | 1))
                                        _ = v73
                                    } else {
                                        v3 = v1
                                        v23 = (v23 + "a")
                                    }
                                }
                                v76 := GenS13{F0: false, F1: (0 > 0)}
                                _ = v76
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v77 uint = 263
                            _ = v77
                        }
                    } else {
                        v1 -= v1
                    }
                    if (1 < 996) {
                        switch fi22 {
                        case 128:
                            var v78 int32 = 1
                            _ = v78
                        default:
                            v5 -= (fi11 & (v5 & -2147483648))
                        }
                        v79 := GenS6{F0: 9, F1: -84}
                        _ = v79
                    } else {
                        v80 := GenS1{F0: -32768, F1: 2, F2: 165}
                        v81 := GenS9{F0: true, F1: 214}
                        switch "hello" {
                        case "a":
                            var v82 string = v23
                            _ = v82
                        }
                        _ = v80
                        _ = v81
                    }
                    v4 += 1.5
                }
                _ = fi22
                _ = v23
            }
        }
        var v83 uint32 = 2147483648
        _ = fi11
        _ = v83
    }
    if true {
        fmt.Println(v4, v4)
    }
    var v84 uint8 = 1
    v4 += 1.5
    if false {
        switch -287 {
        case 0:
            fmt.Println(v84)
        case 28:
            switch (("a" + "a") + ("a" + "\n\t")) {
            case "foo bar":
                if (68 < 4294967295) {
                    _ = len(v2)
                } else {
                    for fi85 := 0; fi85 < 3; fi85++ {
                        switch ("a" + "\n\t") {
                        case "":
                            var v86 uint32 = 393
                            _ = v86
                        case "a":
                            switch "hello" {
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                _ = len(v2)
                            case "":
                                var v87 bool = (2 <= 2)
                                _ = v87
                            default:
                                var v88 uint64 = 2
                                _ = v88
                            }
                        case "hello":
                            fmt.Println(v1)
                        }
                        if false {
                            v2 = append(v2, 0)
                            v4 += float64(float32(0.0))
                        } else {
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                if false {
                                    v2 = append(v2, 753)
                                    var v89 uint = 63
                                    v4 += v4
                                    _ = v89
                                }
                            case 9223372036854775807:
                                switch v5 {
                                case 1:
                                    v3 += (v3 << ((18446744073709551615) & 63))
                                case 4294967295:
                                    var v90 int8 = 42
                                    _ = v90
                                default:
                                    var v91 float32 = 3.14159
                                    _ = v91
                                }
                            }
                        }
                        v92 := GenS11{F0: false, F1: 563}
                        _ = fi85
                        _ = v92
                    }
                    var v93 bool = true
                    v94 := GenS14{F0: -89, F1: 1}
                    _ = v93
                    _ = v94
                }
            case "a":
                v95 := GenS10{F0: (^(v84)), F1: v5, F2: 428}
                _ = v95
            case "\n\t":
                v84 = (v84 + v84)
            }
        default:
            for fi96 := 0; fi96 < 6; fi96++ {
                v97 := GenS15{F0: 1}
                _ = fi96
                _ = v97
            }
        }
    } else {
        v4 -= ((v4 / v4) * float64(3.14159))
        for fi98 := 0; fi98 < 8; fi98++ {
            var v99 int64 = 1
            var v100 int32 = 1
            _ = fi98
            _ = v100
            _ = v99
        }
    }
    v101 := GenS9{F0: false, F1: 65535}
    v84 += (v84 - 253)
    if false {
        v84 = 128
    } else {
        for fi102 := 0; fi102 < 10; fi102++ {
            v2 = append(v2, 0)
            var v103 string = "\n\t"
            for fi104 := 0; fi104 < 18; fi104++ {
                var v105 uint16 = 0
                v84 += uint8(v105)
                _ = fi104
                _ = v105
            }
            _ = fi102
            _ = v103
        }
        var v106 uint8 = v84
        v3 -= (v1 % (((v3 ^ (v1 & (^(v1))))) | 1))
        _ = v106
    }
    var v107 uint = 4294967295
    fmt.Println(v5, v107)
    v5 = v5
    if false {
        var v108 int64 = 42
        _ = v108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v101
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
}
