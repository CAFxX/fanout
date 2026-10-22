package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int16; F2 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS16 struct { F0 bool; F1 bool }
type GenS17 struct { F0 int16; F1 int32 }
type GenS18 struct { F0 uint; F1 uint64 }
type GenS19 struct { F0 uint16; F1 uint16 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int32; F1 int32 }
type GenS23 struct { F0 uint32 }

func main() {
    var v1 uint = 65535
    v2 := []uint32{2, 2147483648, 65535, 590}
    var v3 uint16 = 32767
    v1 -= (v1 / ((65535) | 1))
    var v4 int16 = 0
    var v5 int32 = -1
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v3 -= v3
        _ = fi6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 bool = (v4 > v4)
        v8 = (false && false)
        v2 = append(v2, 42)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v3 * (v3 + v3)), F1: 288}
    fmt.Println(v3)
    v10 := GenS2{F0: 0}
    var v11 bool = (-32768 < (v4 | (-(v4))))
    switch 4294967295 {
    case 42:
        v12 := GenS3{F0: 1, F1: -1, F2: 2}
        _ = v12
    case 604:
        v5 = v5
    case -866:
        v11 = true
    default:
        if true {
            if v11 {
                if v11 {
                    if len(v2) > 0 {
                        v2[0] = 672
                    }
                    var v13 uint64 = 0
                    var v14 float64 = 1e10
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS4{F0: (v1 ^ (v1 / (((v1 * (v1 ^ 18446744073709551615))) | 1)))}
                    switch 32768 {
                    case 2147483647:
                        v16 := GenS5{F0: 826, F1: (v3 + 111)}
                        _ = v16
                    case 32767:
                        v1 = (v1 ^ v1)
                    case -739:
                        if false {
                            switch ("hello" + "hello") {
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = 2147483648
                                }
                            case "hello":
                                var v17 int8 = -57
                                _ = v17
                            default:
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi18
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v19 int64 = 42
                            _ = v19
                        }
                    default:
                        switch ("\n\t" + "a") {
                        case "":
                            var v20 int32 = 42
                            _ = v20
                        default:
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        }
                    }
                    _ = v15
                }
            } else {
                for fi21 := 0; fi21 < 5; fi21++ {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        var v23 int32 = v5
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
                switch "hello" {
                case "hello":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if ((v11 && ((0 != 0) && v11)) || v11) {
                            for fi25 := 0; fi25 < 2; fi25++ {
                                var v26 int64 = 1
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            v4 += 1000
                            if v11 {
                                for fi27 := 0; fi27 < 11; fi27++ {
                                    switch "hello" {
                                    case "\n\t":
                                        switch fi24 {
                                        case -2147483648:
                                            v11 = (!((!(0 < -360)) && v11))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = len(v2)
                                        default:
                                            if (!((v1 <= v1) && v11)) {
                                                v4 = v4
                                                if v11 {
                                                    var v28 string = ("a" + ("foo bar" + "foo bar"))
                                                    var v29 uint64 = 0
                                                    if v11 {
                                                        fmt.Println(v4, v28, fi24)
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    v5 = v5
                                                case "hello":
                                                    var v30 int8 = -114
                                                    _ = v30
                                                case "foo bar":
                                                    if v11 {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                default:
                                                    v5 += -41
                                                }
                                                var v31 uint32 = 2147483648
                                                _ = v31
                                            }
                                        }
                                    case "a":
                                        switch fi24 {
                                        case 32767:
                                            var v32 int32 = v5
                                            _ = v32
                                        }
                                    case "hello":
                                        v33 := GenS1{F0: (v3 / ((v3) | 1)), F1: (fi27 - fi24)}
                                        _ = v33
                                    default:
                                        v5 += v5
                                    }
                                    v5 = v5
                                    v3 += (v3 - 1)
                                    if v11 {
                                        v5 = v5
                                        var v34 int32 = (v5 - (v5 / (((v5 % (((v5 % (((v5 | (v5 & (v5 + v5)))) | 1))) | 1))) | 1)))
                                        _ = v34
                                    } else {
                                        var v35 uint16 = 42
                                        _ = v35
                                    }
                                    var v36 uint = (v1 ^ v1)
                                    _ = fi27
                                    _ = v36
                                }
                                var v37 string = "a"
                                _ = v37
                            }
                        }
                        _ = fi24
                    }
                case "":
                    v38 := GenS4{F0: (v1 ^ 4294967295)}
                    _ = v38
                }
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v4 = (v4 << (((v1 ^ v1)) & 63))
                case 602:
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v3 = v3
                        if v11 {
                            v40 := GenS6{F0: -128, F1: (v4 + (v4 / (((v4 | 1)) | 1))), F2: ""}
                            _ = v40
                        } else {
                            if ((uint32(v1) != 0) || (!((v3 % (((^(v3))) | 1)) > 1))) {
                                v41 := GenS7{F0: "a"}
                                _ = v41
                            } else {
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    if false {
                                        if (v11 || (1.5 > (-2.25 + 1e10))) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = len(v2)
                                        } else {
                                            v42 := GenS8{F0: 4294967295}
                                            switch "hello" {
                                            case "hello":
                                                for fi43 := 0; fi43 < 10; fi43++ {
                                                    if ((fi43 != 0) && (128 < 255)) {
                                                        v44 := GenS9{F0: v5}
                                                        var v45 int16 = v4
                                                        var v46 uint64 = 18446744073709551615
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        switch fi39 {
                                                        case 150:
                                                            v47 := GenS4{F0: (v1 << (((v1 - (v1 << (((v1 & 300)) & 63)))) & 63))}
                                                            _ = v47
                                                        case 122:
                                                            _ = len(v2)
                                                            _ = len(v2)
                                                            _ = len(v2)
                                                        case -919:
                                                            if (255 > 2) {
                                                                var v48 int64 = 0
                                                                v49 := GenS10{F0: 704}
                                                                switch "a" {
                                                                case "a":
                                                                    var v50 int32 = 42
                                                                    _ = v50
                                                                case "":
                                                                    v51 := GenS11{F0: v1, F1: (^(v5))}
                                                                    _ = v51
                                                                case "\n\t":
                                                                    v2 = append(v2, 4294967295)
                                                                    v52 := GenS6{F0: -60, F1: v4, F2: ""}
                                                                    switch -964 {
                                                                    case -9223372036854775808:
                                                                        var v53 int16 = (v4 | v4)
                                                                        _ = v53
                                                                    case -377:
                                                                        var v54 uint = (v1 ^ (^(v1)))
                                                                        _ = v54
                                                                    default:
                                                                        v4 -= v4
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 2147483647
                                                                        }
                                                                    }
                                                                    _ = v52
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                v2 = append(v2, 661)
                                                                v4 = (^(v4))
                                                            }
                                                        default:
                                                            switch (fi43 + fi39) {
                                                            case 1:
                                                                v1 += (^(v1))
                                                            }
                                                        }
                                                        var v55 bool = (!(v11 && (1 != 66)))
                                                        if len(v2) > 0 {
                                                            v2[0] = 300
                                                        }
                                                        _ = v55
                                                    }
                                                    _ = fi43
                                                }
                                            case "":
                                                v56 := GenS2{F0: 51}
                                                _ = v56
                                            case "a":
                                                var v57 int32 = 42
                                                _ = v57
                                            default:
                                                for fi58 := 0; fi58 < 6; fi58++ {
                                                    var v59 int = (fi39 + (fi58 << ((v1) & 63)))
                                                    if v11 {
                                                        if (42 < (v3 + (v3 ^ v3))) {
                                                            v1 = (v1 ^ 255)
                                                            v60 := GenS4{F0: (^(v1))}
                                                            var v61 uint8 = 128
                                                            var v62 int32 = 42
                                                            var v63 string = ("" + "\n\t")
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            switch "" {
                                                            case "\n\t":
                                                                v1 = 2
                                                            case "hello":
                                                                var v64 int64 = 0
                                                                _ = v64
                                                            default:
                                                                fmt.Println(fi58)
                                                            }
                                                        }
                                                        v65 := GenS8{F0: 4294967295}
                                                        var v66 int8 = 1
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    switch ("hello" + "a") {
                                                    case "foo bar":
                                                        var v67 string = ("hello" + "foo bar")
                                                        _ = v67
                                                    case "\n\t":
                                                        switch ("\n\t" + "") {
                                                        case "hello":
                                                            if ((float64(3.14159) + float64(-2.25)) != 3.14159) {
                                                                switch ("foo bar" + ("foo bar" + "a")) {
                                                                case "foo bar":
                                                                    switch v59 {
                                                                    case -9223372036854775808:
                                                                        if (!(!v11)) {
                                                                            var v68 int16 = int16(fi39)
                                                                            var v69 uint16 = (v3 << ((v1) & 63))
                                                                            var v70 int16 = (v4 << (((v1 / ((v1) | 1))) & 63))
                                                                            _ = v68
                                                                            _ = v69
                                                                            _ = v70
                                                                        } else {
                                                                            fmt.Println(v1)
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if v11 {
                                                                            v71 := GenS6{F0: 127, F1: (v4 << ((0) & 63)), F2: ("a" + ("" + ""))}
                                                                            var v72 float32 = ((1e10 + ((3.14159 + (1.5 / ((3.14159 * 3.14159) * (float32(0.1) / 1.0)))) / (float32(0.1) + 1.5))) * 1e10)
                                                                            v59 -= 32767
                                                                            v73 := GenS12{F0: uint(v3)}
                                                                            if (4294967295 >= (v59 / (((fi39 ^ (fi39 << (((v1 & v1)) & 63)))) | 1))) {
                                                                                if (-1 >= (^(fi39))) {
                                                                                    _ = len(v2)
                                                                                    if ((!true) || v11) {
                                                                                        if ((!true) && (9223372036854775807 > 42)) {
                                                                                            v11 = false
                                                                                            v59 += (-(fi39))
                                                                                            v74 := GenS10{F0: 0}
                                                                                            v1 = v1
                                                                                            _ = v74
                                                                                        }
                                                                                        var v75 int64 = int64(fi58)
                                                                                        if ("\n\t" <= "") {
                                                                                            v1 = 0
                                                                                        } else {
                                                                                            var v76 int64 = (v75 & (^(v75)))
                                                                                            _ = len(v2)
                                                                                            _ = v76
                                                                                        }
                                                                                        _ = v75
                                                                                    }
                                                                                } else {
                                                                                    v59 = (v59 % (((fi39 / ((fi39) | 1))) | 1))
                                                                                }
                                                                                var v77 float64 = 0.1
                                                                                var v78 bool = true
                                                                                _ = v77
                                                                                _ = v78
                                                                            } else {
                                                                                v5 -= (v5 | (v5 | (v5 - 42)))
                                                                            }
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v59 += fi58
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 677
                                                                            }
                                                                            var v79 uint64 = 1
                                                                            if (3.14159 <= 1.5) {
                                                                                v80 := GenS5{F0: 0, F1: uint16(v1)}
                                                                                _ = v80
                                                                            } else {
                                                                                v2 = append(v2, 4294967295)
                                                                                fmt.Println(v79)
                                                                                var v81 int64 = -9223372036854775808
                                                                                var v82 int64 = (v81 >> (((v1 + v1)) & 63))
                                                                                switch 508 {
                                                                                case 62:
                                                                                    fmt.Println(v82, fi58)
                                                                                case 128:
                                                                                    if (!(false || v11)) {
                                                                                        v2 = append(v2, 0)
                                                                                        v83 := GenS10{F0: v79}
                                                                                        v81 -= 0
                                                                                        switch ("foo bar" + "") {
                                                                                        case "":
                                                                                            if ((!(v82 <= (v81 & (v82 << ((v1) & 63))))) || (!((!(111 >= 34)) || ((3.14159 * 1e10) > 1.5)))) {
                                                                                                var v84 bool = (316 < 1)
                                                                                                _ = v84
                                                                                            }
                                                                                        case "a":
                                                                                            v85 := GenS8{F0: (v79 ^ (^(v79)))}
                                                                                            _ = v85
                                                                                        case "hello":
                                                                                            v86 := GenS10{F0: 18446744073709551615}
                                                                                            _ = v86
                                                                                        default:
                                                                                            v59 += (fi39 + 0)
                                                                                        }
                                                                                        v87 := GenS11{F0: v1, F1: int32(fi39)}
                                                                                        _ = v83
                                                                                        _ = v87
                                                                                    } else {
                                                                                        if (!(v82 == (v82 & (v81 << ((v1) & 63))))) {
                                                                                            v82 = (v82 << ((v1) & 63))
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                            var v88 int16 = v4
                                                                                            _ = v88
                                                                                        }
                                                                                        v5 = -26
                                                                                    }
                                                                                default:
                                                                                    if v11 {
                                                                                        var v89 string = "hello"
                                                                                        switch fi39 {
                                                                                        case 2147483647:
                                                                                            if v11 {
                                                                                                v2 = append(v2, 220)
                                                                                            }
                                                                                        case 32768:
                                                                                            _ = len(v2)
                                                                                        default:
                                                                                            var v90 uint8 = 127
                                                                                            var v91 uint32 = 2
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        }
                                                                                        switch v89 {
                                                                                        case "a":
                                                                                            fmt.Println(v59)
                                                                                        case "hello":
                                                                                            v5 += (^(v5))
                                                                                        case "\n\t":
                                                                                            v2 = append(v2, 65535)
                                                                                        }
                                                                                        _ = v89
                                                                                    }
                                                                                }
                                                                                var v92 uint64 = v79
                                                                                _ = v81
                                                                                _ = v82
                                                                                _ = v92
                                                                            }
                                                                            var v93 int = int(v4)
                                                                            switch "a" {
                                                                            case "":
                                                                                var v94 uint8 = 236
                                                                                _ = v94
                                                                            }
                                                                            var v95 uint16 = (v3 / ((599) | 1))
                                                                            _ = v79
                                                                            _ = v93
                                                                            _ = v95
                                                                        }
                                                                    case -1:
                                                                        if v11 {
                                                                            fmt.Println(v5, v1)
                                                                            if v11 {
                                                                                if (v1 > v1) {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 2147483647
                                                                                    }
                                                                                    if ((v4 ^ v4) >= -1) {
                                                                                        var v96 float64 = 0.0
                                                                                        var v97 int = -9223372036854775808
                                                                                        var v98 uint32 = 1
                                                                                        if (0.0 < 1e10) {
                                                                                            v59 += fi58
                                                                                            if (v4 == v4) {
                                                                                                _ = len(v2)
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if false {
                                                                                                    v59 = fi39
                                                                                                    v99 := GenS13{F0: (v97 + fi58)}
                                                                                                    var v100 int32 = 42
                                                                                                    _ = v100
                                                                                                    _ = v99
                                                                                                }
                                                                                                var v101 uint = (^(v1))
                                                                                                _ = v101
                                                                                            }
                                                                                            if false {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 42
                                                                                                }
                                                                                                var v102 uint32 = (v98 - (v98 << ((18446744073709551615) & 63)))
                                                                                                v4 += v4
                                                                                                var v103 uint64 = 42
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                            } else {
                                                                                                v2 = append(v2, 2147483647)
                                                                                                if false {
                                                                                                    v98 = v98
                                                                                                } else {
                                                                                                    var v104 int8 = 0
                                                                                                    fmt.Println(v59, v104, fi58)
                                                                                                    var v105 int64 = -809
                                                                                                    _ = v104
                                                                                                    _ = v105
                                                                                                }
                                                                                                v106 := GenS7{F0: ((("" + "\n\t") + ("a" + ("\n\t" + "hello"))) + "\n\t")}
                                                                                                _ = v106
                                                                                            }
                                                                                        }
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                    if (false || (!v11)) {
                                                                                        if (2 >= 18446744073709551615) {
                                                                                            var v107 int = (fi39 ^ -869)
                                                                                            fmt.Println(v107, v1)
                                                                                            v108 := GenS14{F0: 42, F1: (v1 | 255)}
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                        if (!(-963 != 0)) {
                                                                                            if ((("\n\t" + "") + "") == "a") {
                                                                                                var v109 int64 = 42
                                                                                                var v110 int = 4294967295
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            }
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                v5 -= (v5 / ((v5) | 1))
                                                                                            }
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = 306
                                                                                            }
                                                                                        }
                                                                                        v11 = true
                                                                                    } else {
                                                                                        v111 := GenS3{F0: 2, F1: 1, F2: 42}
                                                                                        fmt.Println(fi39)
                                                                                        v112 := GenS15{F0: v3, F1: 255, F2: 1}
                                                                                        switch fi39 {
                                                                                        case 32767:
                                                                                            if v11 {
                                                                                                v59 += 127
                                                                                                switch 127 {
                                                                                                case 32768:
                                                                                                    fmt.Println(v11, v1)
                                                                                                case 9223372036854775807:
                                                                                                    if false {
                                                                                                        if (!(!(((((((1.5 * float32(((((1.5 + 3.14159) - 0.0) * 0.1) * 1.5))) - (0.1 * 0.0)) * 0.1) / (1e10 - -2.25)) * 1.5) * ((3.14159 - float32(0.1)) + 1e10)) > 1.5))) {
                                                                                                            v113 := GenS16{F0: ((v3 | v3) <= 65535), F1: (true || (v11 && ((^(v3)) > (v3 / (((v3 >> (((v1 & v1)) & 63))) | 1)))))}
                                                                                                            v11 = ((false || false) && (246 != 2))
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = 42
                                                                                                            }
                                                                                                            v114 := GenS4{F0: (v1 % ((v1) | 1))}
                                                                                                            if false {
                                                                                                                fmt.Println(v5, v3, v1)
                                                                                                                fmt.Println(v11)
                                                                                                            }
                                                                                                            _ = v113
                                                                                                            _ = v114
                                                                                                        }
                                                                                                        if v11 {
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                        }
                                                                                                        if v11 {
                                                                                                            switch (v59 >> (((v1 | 65535)) & 63)) {
                                                                                                            case 211:
                                                                                                                v59 = (-(fi39))
                                                                                                            }
                                                                                                            var v115 int = (fi58 * (v59 ^ (v59 % ((v59) | 1))))
                                                                                                            _ = v115
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v116 int8 = -1
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                        _ = v116
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v4, fi39)
                                                                                                if (v11 && (2 >= 1)) {
                                                                                                    var v117 uint32 = 803
                                                                                                    _ = v117
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v118 uint32 = 1
                                                                                            _ = v118
                                                                                        }
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    var v119 string = "hello"
                                                                                    var v120 uint = v1
                                                                                    fmt.Println(v4, v4)
                                                                                    switch 6 {
                                                                                    case -756:
                                                                                        var v121 int32 = (v5 << (((^(v1))) & 63))
                                                                                        _ = v121
                                                                                    default:
                                                                                        switch fi58 {
                                                                                        case 128:
                                                                                            var v122 string = ("" + v119)
                                                                                            _ = v122
                                                                                        }
                                                                                    }
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                } else {
                                                                                    v4 = v4
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (127 > -128) {
                                                                            v59 -= (v59 >> ((v1) & 63))
                                                                            v123 := GenS17{F0: (^(v4)), F1: -2147483648}
                                                                            _ = len(v2)
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v124 := GenS18{F0: v1, F1: 0}
                                                                        _ = v124
                                                                    case "":
                                                                        if v11 {
                                                                            v125 := GenS19{F0: v3, F1: v3}
                                                                            if v11 {
                                                                                var v126 uint = 0
                                                                                _ = v126
                                                                            }
                                                                            _ = v125
                                                                        } else {
                                                                            if (!((("a" + "foo bar") + "hello") >= "")) {
                                                                                if v11 {
                                                                                    if v11 {
                                                                                        fmt.Println(v59, v5)
                                                                                        v1 += (v1 + (v1 / ((2) | 1)))
                                                                                    } else {
                                                                                        var v127 int64 = -1
                                                                                        _ = v127
                                                                                    }
                                                                                    v59 = (fi39 / ((-1) | 1))
                                                                                    var v128 uint32 = uint32(v3)
                                                                                    _ = v128
                                                                                }
                                                                                v129 := GenS8{F0: 610}
                                                                                if v11 {
                                                                                    v1 -= (v1 % ((v1) | 1))
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    if v11 {
                                                                                        fmt.Println(fi39)
                                                                                    } else {
                                                                                        if (true || v11) {
                                                                                            if v11 {
                                                                                                var v130 float32 = 3.14159
                                                                                                _ = v130
                                                                                            } else {
                                                                                                var v131 int8 = 0
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 2147483648
                                                                                                }
                                                                                                fmt.Println(v4, v1)
                                                                                                var v132 bool = (0 == 42)
                                                                                                var v133 int16 = v4
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                                _ = v133
                                                                                            }
                                                                                            v4 = v4
                                                                                            v59 = 0
                                                                                        } else {
                                                                                            var v134 uint = v1
                                                                                            _ = v134
                                                                                        }
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        switch fi39 {
                                                                                        case -543:
                                                                                            v5 = (v5 >> (((v1 >> (((v1 | v1)) & 63))) & 63))
                                                                                        case 32767:
                                                                                            v59 = fi39
                                                                                        default:
                                                                                            v5 += v5
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v5 += v5
                                                                                    v135 := GenS19{F0: v3, F1: 340}
                                                                                    v4 = (^(v4))
                                                                                    _ = v135
                                                                                }
                                                                                _ = v129
                                                                            } else {
                                                                                var v136 int32 = int32(fi39)
                                                                                if ((127 <= -16) && v11) {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = uint32(v3)
                                                                                    }
                                                                                    if (v136 <= v136) {
                                                                                        v137 := GenS20{F0: (^(v3))}
                                                                                        _ = v137
                                                                                    }
                                                                                    v138 := GenS21{F0: v3}
                                                                                    v3 = 379
                                                                                    var v139 int8 = 0
                                                                                    _ = v138
                                                                                    _ = v139
                                                                                } else {
                                                                                    v11 = false
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        v136 -= (v136 % ((v136) | 1))
                                                                                    case "hello":
                                                                                        var v140 int16 = (v4 - -235)
                                                                                        _ = v140
                                                                                    case "":
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    default:
                                                                                        var v141 float32 = float32((((0.1 - 0.0) - (-2.25 + (1.5 / (float64(3.14159) / 1e10)))) * 1e10))
                                                                                        _ = v141
                                                                                    }
                                                                                    v3 += v3
                                                                                }
                                                                                var v142 float64 = (1.5 * (float64(float32((0.0 - ((-2.25 + (((((0.1 - 1e10) - ((float64(0.1) * 1e10) * -2.25)) - (3.14159 + 1e10)) - float64(-2.25)) * (0.0 - 1e10))) / 0.1)))) + ((float64(-2.25) - 0.0) * ((1.5 * 0.0) + float64(3.14159)))))
                                                                                _ = v136
                                                                                _ = v142
                                                                            }
                                                                            switch (v59 >> (((v1 / ((v1) | 1))) & 63)) {
                                                                            case 2147483647:
                                                                                var v143 uint16 = (v3 ^ 993)
                                                                                _ = v143
                                                                            }
                                                                            var v144 uint32 = 2147483648
                                                                            _ = v144
                                                                        }
                                                                    case "hello":
                                                                        v5 -= v5
                                                                    }
                                                                case "\n\t":
                                                                    v4 += 42
                                                                default:
                                                                    v5 = (v5 - v5)
                                                                }
                                                                var v145 uint64 = 2
                                                                v146 := GenS9{F0: (v5 << ((858) & 63))}
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                        case "foo bar":
                                                            v59 = 128
                                                        default:
                                                            v147 := GenS1{F0: (v3 ^ v3), F1: (fi39 << (((v1 * (v1 - (v1 | v1)))) & 63))}
                                                            _ = v147
                                                        }
                                                    }
                                                    _ = fi58
                                                    _ = v59
                                                }
                                            }
                                            _ = v42
                                        }
                                        v148 := GenS22{F0: (-(v5)), F1: int32(v1)}
                                        for fi149 := 0; fi149 < 15; fi149++ {
                                            if (127 >= 0) {
                                                v3 = 628
                                                v4 += v4
                                            }
                                            v3 += (v3 ^ 1)
                                            v4 += v4
                                            _ = fi149
                                        }
                                        _ = v148
                                    }
                                case "foo bar":
                                    switch 42 {
                                    case 410:
                                        v3 = v3
                                    case 9223372036854775807:
                                        switch (fi39 << ((v1) & 63)) {
                                        case 42:
                                            for fi150 := 0; fi150 < 16; fi150++ {
                                                _ = len(v2)
                                                v2 = append(v2, 0)
                                                _ = fi150
                                            }
                                        case 124:
                                            var v151 uint = (v1 ^ v1)
                                            _ = v151
                                        }
                                    }
                                }
                                if (-956 == v5) {
                                    if v11 {
                                        for fi152 := 0; fi152 < 9; fi152++ {
                                            var v153 int8 = 51
                                            _ = fi152
                                            _ = v153
                                        }
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 2147483648
                                        }
                                        var v154 bool = (((v3 & (^(v3))) < 542) && v11)
                                        _ = v154
                                    }
                                    var v155 int16 = (-(v4))
                                    if len(v2) > 0 {
                                        v2[0] = uint32(v3)
                                    }
                                    _ = v155
                                } else {
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                        v2 = append(v2, 65535)
                                    }
                                }
                                switch ("hello" + ("" + "foo bar")) {
                                case "hello":
                                    switch "a" {
                                    case "a":
                                        if true {
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = 42
                                                }
                                            } else {
                                                var v156 float64 = -2.25
                                                _ = v156
                                            }
                                        }
                                    }
                                }
                            }
                            if (v11 || false) {
                                var v157 float64 = (1.5 - (1.5 - 1.5))
                                if false {
                                    var v158 uint8 = 42
                                    _ = v158
                                }
                                for fi159 := 0; fi159 < 5; fi159++ {
                                    v3 = v3
                                    switch (((("\n\t" + "foo bar") + ("\n\t" + "a")) + "\n\t") + ("" + "")) {
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = uint32(v3)
                                        }
                                    case "":
                                        v3 -= (^(v3))
                                        v3 -= v3
                                    case "\n\t":
                                        v11 = false
                                    }
                                    v160 := GenS21{F0: (v3 << ((4294967295) & 63))}
                                    _ = fi159
                                    _ = v160
                                }
                                _ = v157
                            }
                            v2 = append(v2, 1)
                        }
                        _ = fi39
                    }
                case -2147483648:
                    for fi161 := 0; fi161 < 1; fi161++ {
                        v4 = (v4 - v4)
                        v162 := GenS23{F0: 598}
                        fmt.Println(v5)
                        _ = fi161
                        _ = v162
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = uint32(v3)
                    }
                }
            }
            var v163 uint = v1
            v4 = (v4 + 1)
            _ = v163
        } else {
            var v164 int32 = (v5 ^ -1)
            _ = v164
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
