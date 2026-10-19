package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint; F1 int32; F2 string }
type GenS9 struct { F0 int8; F1 bool; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS11 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS12 struct { F0 uint16; F1 uint }
type GenS13 struct { F0 uint64; F1 string; F2 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint32, p1 int16) int8 {
    p1 = (p1 & -1)
    var v7 int8 = -128
    v8 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 832
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int = 32768
    v5 := []uint32{}
    v6 := make(map[int]string)
    v4 += v4
    for fi9 := 0; fi9 < 19; fi9++ {
        switch -716 {
        case 128:
            v10 := GenS1{F0: (fi9 <= 127)}
            _ = v10
        case 1:
            var v11 int8 = -128
            _ = v11
        case -937:
            var v12 bool = (true || ("\n\t" != "\n\t"))
            _ = v12
        default:
            switch 127 {
            case 2147483647:
                switch "" {
                case "":
                    if false {
                        if (true && true) {
                            var v13 uint16 = 255
                            fmt.Println(v1, v4)
                            _ = v13
                        } else {
                            if (false || false) {
                                for fi14 := 0; fi14 < 18; fi14++ {
                                    if true {
                                        var v15 uint64 = v1
                                        _ = v15
                                    } else {
                                        var v16 uint64 = v1
                                        _ = v16
                                    }
                                    if false {
                                        if ((!true) && false) {
                                            switch (fi9 ^ (fi14 + (-(fi14)))) {
                                            case 345:
                                                if (((1000 <= 42) && true) && false) {
                                                    switch "a" {
                                                    case "":
                                                        v3[814] = 4294967295
                                                    case "\n\t":
                                                        if (100 != -73) {
                                                            v1 -= (v1 - (v1 % (((v1 & (v1 ^ v1))) | 1)))
                                                            var v17 uint64 = (v1 + (v1 ^ 1))
                                                            _ = v17
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            if true {
                                                                v18 := GenS2{F0: "foo bar", F1: -32768, F2: false}
                                                                delete(v2, fi9)
                                                                _ = v18
                                                            }
                                                            v5 = append(v5, 597)
                                                        } else {
                                                            if len(v5) > 0 {
                                                                v5[0] = 423
                                                            }
                                                            v1 += (v1 - 909)
                                                        }
                                                    }
                                                    fmt.Println(v1, v4, v4)
                                                }
                                            }
                                            v4 += fi14
                                        } else {
                                            if true {
                                                var v19 float32 = (float32((3.14159 - 1e10)) * 0.0)
                                                _ = v19
                                            } else {
                                                if true {
                                                    if (!(212 >= 128)) {
                                                        if true {
                                                            v20 := GenS3{F0: 28, F1: false, F2: 0}
                                                            if true {
                                                                var v21 int16 = 746
                                                                switch ("a" + (("foo bar" + "foo bar") + "hello")) {
                                                                case "":
                                                                    fmt.Println(v4)
                                                                default:
                                                                    if ((float32(0.0) - ((0.1 - 1.5) - 0.0)) == float32(1.5)) {
                                                                        v22 := GenS4{F0: 2147483647}
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v21 -= 1000
                                                                        case "":
                                                                            fmt.Println(v1)
                                                                        case "a":
                                                                            v1 = 383
                                                                        }
                                                                        _ = v22
                                                                    } else {
                                                                        v4 = (fi9 ^ (fi9 / (((fi14 | (-(v4)))) | 1)))
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v4 = v4
                                                                        default:
                                                                            v23 := GenS4{F0: -2147483648}
                                                                            _ = v23
                                                                        }
                                                                    }
                                                                    var v24 int8 = int8(v4)
                                                                    _ = v24
                                                                }
                                                                _ = v21
                                                            }
                                                            _ = v20
                                                        }
                                                    } else {
                                                        if false {
                                                            var v25 float64 = 0.1
                                                            v4 -= fi14
                                                            if false {
                                                                if len(v5) > 0 {
                                                                    v5[0] = 54
                                                                }
                                                            }
                                                            _ = v25
                                                        } else {
                                                            fmt.Println(fi14, v1)
                                                            v4 = fi9
                                                            if false {
                                                                var v26 int64 = 678
                                                                _ = v26
                                                            }
                                                        }
                                                        v27 := GenS2{F0: ((("" + "hello") + "\n\t") + (("foo bar" + "a") + "foo bar")), F1: 854, F2: true}
                                                        _ = v27
                                                    }
                                                    fmt.Println(v1)
                                                    var v28 bool = true
                                                    switch "foo bar" {
                                                    case "":
                                                        delete(v3, 978)
                                                    case "\n\t":
                                                        delete(v6, 219)
                                                    default:
                                                        v28 = (!v28)
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 139
                                                    }
                                                    _ = v28
                                                }
                                                fmt.Println(v1, fi9)
                                            }
                                            if ((42 < 32767) && (42 < 0)) {
                                                if (true && true) {
                                                    var v29 int = fi14
                                                    _ = v29
                                                } else {
                                                    var v30 int64 = 1
                                                    _ = v30
                                                }
                                                v31 := GenS5{F0: 915, F1: (!(1 >= 817))}
                                                v32 := GenS6{F0: 2, F1: v1, F2: 42}
                                                if false {
                                                    var v33 int16 = 1
                                                    if len(v5) > 0 {
                                                        v5[0] = 262
                                                    }
                                                    v33 += (v33 ^ 0)
                                                    _ = v33
                                                } else {
                                                    var v34 int32 = 0
                                                    _ = v34
                                                }
                                                if (false || true) {
                                                    if false {
                                                        if len(v5) > 0 {
                                                            v5[0] = 779
                                                        }
                                                        var v35 int = fi14
                                                        _ = v35
                                                    }
                                                    if true {
                                                        v1 += (v1 & v1)
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if false {
                                                                var v36 string = ""
                                                                v1 = v1
                                                                _ = v36
                                                            } else {
                                                                v1 -= (v1 & (v1 / ((v1) | 1)))
                                                                var v37 int64 = -353
                                                                _ = v37
                                                            }
                                                        case "\n\t":
                                                            if (((0.0 * (float32(0.0) * (float32(1.5) - 1e10))) + 1.5) < float32(((0.0 * 3.14159) + 0.0))) {
                                                                v38 := v2[(v4 / ((v4) | 1))]
                                                                _ = v38
                                                            }
                                                            v39 := GenS7{F0: "\n\t"}
                                                            _ = v39
                                                        case "":
                                                            v4 = (v4 & (-(fi14)))
                                                        default:
                                                            v40 := GenS8{F0: 900, F1: -2147483648, F2: (("foo bar" + "hello") + (("foo bar" + "a") + "a"))}
                                                            _ = v40
                                                        }
                                                        if false {
                                                            v41 := GenS6{F0: 781, F1: v1, F2: 0}
                                                            var v42 int = v4
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    } else {
                                                        switch v4 {
                                                        case -9223372036854775808:
                                                            v43 := GenS9{F0: 0, F1: true, F2: 254}
                                                            _ = v43
                                                        case -2147483648:
                                                            v44 := GenS5{F0: 187, F1: true}
                                                            _ = v44
                                                        case 9223372036854775807:
                                                            v45 := GenS8{F0: 433, F1: 2147483647, F2: "hello"}
                                                            var v46 uint8 = 0
                                                            if false {
                                                                v4 -= fi14
                                                            } else {
                                                                var v47 uint64 = v1
                                                                var v48 uint16 = 32767
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            _ = v45
                                                            _ = v46
                                                        default:
                                                            v49 := GenS9{F0: 39, F1: (true || true), F2: 32767}
                                                            _ = v49
                                                        }
                                                    }
                                                    v5 = append(v5, 2147483647)
                                                } else {
                                                    v1 = (v1 << ((4294967295) & 63))
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                        v4 += fi14
                                    } else {
                                        fmt.Println(v1, v4)
                                        fmt.Println(v1, fi9)
                                        var v50 float64 = float64((3.14159 * (-2.25 * ((1.5 - 1.5) + ((((1e10 * (1e10 / (-2.25 - 1.5))) + 3.14159) - 0.0) - ((((-2.25 - (0.0 + float32(((0.0 + -2.25) / float64(1.5))))) + 0.0) / (0.0 + (float32(((((0.0 + 3.14159) + ((((-2.25 - (((float64(-2.25) + 3.14159) * float64(0.1)) - (3.14159 - (0.1 - 1.5)))) * 1.5) * 3.14159) + (0.0 - -2.25))) * 1e10) / -2.25)) + 0.1))) / (1e10 / 3.14159)))))))
                                        if (false || false) {
                                            if (false || ((65535 >= 255) || (!true))) {
                                                switch (fi14 ^ 1) {
                                                case 32768:
                                                    var v51 uint8 = 1
                                                    _ = v51
                                                case -698:
                                                    switch (fi9 / ((v4) | 1)) {
                                                    case -443:
                                                        _ = len(v5)
                                                    }
                                                case 779:
                                                    v4 -= v4
                                                }
                                                v3[4294967295] = 65535
                                            } else {
                                                v4 = (fi14 << ((65535) & 63))
                                                if false {
                                                    fmt.Println(fi9, v1)
                                                    if len(v5) > 0 {
                                                        v5[0] = 675
                                                    }
                                                }
                                                if true {
                                                    v52 := GenS10{F0: 65535, F1: 929, F2: 1}
                                                    _ = v52
                                                }
                                            }
                                            var v53 uint8 = 42
                                            _ = v53
                                        }
                                        _ = v50
                                    }
                                    v54 := GenS11{F0: (v1 >> ((42) & 63)), F1: 0, F2: 1}
                                    _ = fi14
                                    _ = v54
                                }
                                v5 = append(v5, 785)
                                var v55 uint16 = 2
                                _ = v55
                            }
                            v1 = v1
                        }
                        for fi56 := 0; fi56 < 5; fi56++ {
                            switch fi56 {
                            case -336:
                                if false {
                                    var v57 int16 = -53
                                    if true {
                                        fmt.Println(fi56, v4, v57)
                                        v58 := GenS10{F0: 65535, F1: 973, F2: 42}
                                        _ = v58
                                    } else {
                                        v59 := GenS10{F0: 322, F1: 1, F2: 127}
                                        _ = v59
                                    }
                                    _ = v57
                                }
                            case -616:
                                if false {
                                    var v60 float64 = -2.25
                                    _ = v60
                                } else {
                                    fmt.Println(v4)
                                    var v61 bool = (((true && false) && true) || (("hello" <= "") && false))
                                    _ = v61
                                }
                            }
                            var v62 int16 = -32768
                            _ = fi56
                            _ = v62
                        }
                    } else {
                        if true {
                            if (139 <= 9223372036854775807) {
                                fmt.Println(fi9, v4, fi9)
                                if (4294967295 <= 255) {
                                    if false {
                                        var v63 uint64 = v1
                                        v63 += (v1 * 618)
                                        v4 = fi9
                                        _ = v63
                                    } else {
                                        fmt.Println(v4, fi9)
                                        v1 = (v1 & v1)
                                    }
                                    var v64 int64 = -1
                                    _ = v64
                                }
                            }
                            if (1 < -1) {
                                if true {
                                    v4 -= 127
                                    v65 := v2[1]
                                    fmt.Println(v1, v65, v1)
                                    _ = v65
                                }
                                fmt.Println(fi9)
                                switch (fi9 << ((255) & 63)) {
                                case 767:
                                    v66 := GenS12{F0: 634, F1: 4294967295}
                                    _ = v66
                                case 2147483647:
                                    v1 += v1
                                case -2147483648:
                                    v67 := v3[1]
                                    _ = v67
                                }
                            }
                        } else {
                            v1 = v1
                        }
                        if false {
                            var v68 int32 = 42
                            v6[(v4 % ((v4) | 1))] = "\n\t"
                            for fi69 := 0; fi69 < 14; fi69++ {
                                v70 := GenS7{F0: ("" + ("hello" + "hello"))}
                                var v71 uint32 = 926
                                switch "" {
                                case "":
                                    if ((127 >= 135) && (3.14159 < 0.0)) {
                                        v72 := GenS12{F0: 0, F1: 4294967295}
                                        v6[fi69] = ("hello" + "")
                                        _ = len(v3)
                                        _ = v72
                                    } else {
                                        var v73 float64 = 1e10
                                        if false {
                                            v74 := GenS13{F0: v1, F1: "", F2: uint32(v1)}
                                            if true {
                                                var v75 int64 = -9223372036854775808
                                                delete(v2, fi69)
                                                _ = len(v5)
                                                if false {
                                                    switch (v4 % ((fi9) | 1)) {
                                                    case 1:
                                                        if (123 >= 127) {
                                                            var v76 int32 = (v68 * int32(v71))
                                                            v4 += fi69
                                                            _ = v76
                                                        } else {
                                                            var v77 float32 = 1e10
                                                            _ = v77
                                                        }
                                                    case -662:
                                                        var v78 int8 = 127
                                                        _ = v78
                                                    }
                                                    v79 := GenS7{F0: ((("hello" + "\n\t") + (("hello" + "\n\t") + ("" + ("\n\t" + ("foo bar" + "foo bar"))))) + "")}
                                                    _ = v79
                                                } else {
                                                    var v80 uint8 = 32
                                                    _ = v80
                                                }
                                                _ = v75
                                            } else {
                                                v68 = v68
                                            }
                                            _ = v74
                                        }
                                        v81 := GenS14{F0: 502}
                                        _ = v73
                                        _ = v81
                                    }
                                case "foo bar":
                                    if false {
                                        v71 = (v71 - (^(v71)))
                                    } else {
                                        var v82 int8 = -116
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(fi69, v68, fi9)
                                        case "foo bar":
                                            v68 = v68
                                        case "\n\t":
                                            v83 := GenS10{F0: 1, F1: 255, F2: 127}
                                            _ = v83
                                        default:
                                            v4 -= (fi9 / (((fi69 + int(v68))) | 1))
                                        }
                                        _ = v82
                                    }
                                case "hello":
                                    var v84 int16 = 32767
                                    _ = v84
                                }
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                            _ = v68
                        }
                        fmt.Println(fi9, v4, v4)
                    }
                default:
                    v85 := GenS15{F0: int16(v1)}
                    _ = v85
                }
            case -1:
                v86 := GenS5{F0: 792, F1: (true || (255 > 563))}
                _ = v86
            case 650:
                v87 := GenS16{F0: 255, F1: 4294967295}
                _ = v87
            default:
                if (-2.25 >= (3.14159 + -2.25)) {
                    v1 = v1
                } else {
                    if len(v5) > 0 {
                        v5[0] = 247
                    }
                }
            }
        }
        fmt.Println(fi9, v1, fi9)
        _ = fi9
    }
    v1 = (^(v1))
    v4 = v4
    var v88 bool = (!((3.14159 + -2.25) == -2.25))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v88 = (false || v88)
    var v89 int16 = -162
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v89
}
