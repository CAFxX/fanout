package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint32; F1 int32; F2 int16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int32; F1 int64; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 int64; F2 bool }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int8; F1 uint64 }
func genFn0(p0 string, p1 uint16, p2 int) int64 {
    v6 := []int{p2}
    fmt.Println(p1, p0, p2)
    if false {
        v7 := GenS1{F0: 4294967295, F1: (p1 | (p1 ^ p1)), F2: -32768}
        _ = v7
    } else {
        if len(v6) > 0 {
            v6[0] = p2
        }
        p1 = (p1 + 0)
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 646
        _ = fi8
        _ = v9
    }
    p2 = (p2 & 266)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 826
}
func genFn1(p0 int16) bool {
    for fi10 := 0; fi10 < 2; fi10++ {
        if true {
            switch "a" {
            case "hello":
                p0 += p0
            case "\n\t":
                fmt.Println(fi10, fi10)
            case "a":
                var v11 float64 = 3.14159
                _ = v11
            }
            v12 := GenS2{F0: 2, F1: (fi10 | fi10)}
            _ = v12
        }
        var v13 uint16 = 2
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint32 = 42
            var v16 uint64 = 4294967295
            v16 = (v16 >> ((255) & 63))
            v17 := []int{(^(fi10)), (fi10 << ((uint(v16)) & 63)), -2147483648}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi10
        _ = v13
    }
    v18 := GenS3{F0: 602, F1: 42}
    _ = p0
    _ = v18
    return false
}

func main() {
    var v1 bool = false
    var v2 int8 = 100
    var v3 uint16 = 42
    var v4 uint16 = v3
    v5 := make(map[int]string)
    var v19 uint = 25
    v20 := GenS4{F0: (^(v3)), F1: 2147483647}
    if (0.0 > (1.5 - 3.14159)) {
        switch "" {
        case "hello":
            if true {
                v21 := GenS2{F0: 1, F1: 2147483647}
                delete(v5, 585)
                v22 := GenS3{F0: v19, F1: 42}
                _ = v21
                _ = v22
            } else {
                v23 := GenS5{F0: (v19 * (v19 + (^(v19)))), F1: ""}
                var v24 uint64 = 2
                v4 = v3
                _ = v23
                _ = v24
            }
        case "foo bar":
            var v25 int = 0
            _ = v25
        }
    }
    v26 := []int{32768}
    v27 := GenS6{F0: 2147483647, F1: 281, F2: 819}
    var v28 string = "hello"
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    fmt.Println(v2)
    switch v28 {
    case "":
        v29 := GenS3{F0: v19, F1: genFn0(("foo bar" + "hello"), 308, 32767)}
        _ = v29
    case "\n\t":
        v19 -= (v19 + v19)
    case "hello":
        v4 = 42
    default:
        var v30 int8 = (v2 >> (((^(v19))) & 63))
        _ = v30
    }
    v2 = v2
    if false {
        v31 := GenS7{F0: 18446744073709551615, F1: -811}
        _ = v31
    } else {
        if ((((1.5 - ((0.1 - ((0.0 / 1.0) + (((0.1 + float64(3.14159)) - 1.5) + float64(3.14159)))) + float64(0.1))) + -2.25) - 0.0) == (0.0 / 3.14159)) {
            for fi32 := 0; fi32 < 14; fi32++ {
                v1 = true
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 bool = ((!true) || v1)
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    _ = fi33
                    _ = v34
                }
                switch -1 {
                case -2147483648:
                    var v35 int64 = 0
                    _ = v35
                case 128:
                    for fi36 := 0; fi36 < 11; fi36++ {
                        fmt.Println(v3, v19)
                        v37 := GenS8{F0: 1, F1: genFn0(((v28 + "\n\t") + v28), v4, fi36), F2: v4}
                        _ = fi36
                        _ = v37
                    }
                case -66:
                    v38 := v5[(fi32 | fi32)]
                    _ = v38
                default:
                    switch fi32 {
                    case 631:
                        v19 = 0
                    default:
                        if (9223372036854775807 <= -698) {
                            v3 = v3
                        }
                    }
                }
                _ = fi32
            }
            fmt.Println(v19)
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            if v1 {
                var v40 string = "\n\t"
                switch fi39 {
                case 279:
                    for fi41 := 0; fi41 < 14; fi41++ {
                        fmt.Println(v4, v2)
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = fi41
                    }
                case 0:
                    v42 := GenS9{F0: -314}
                    _ = v42
                }
                _ = v40
            }
            _ = fi39
        }
    }
    if ((false && true) && (-9223372036854775808 != -340)) {
        v2 -= (v2 >> ((v19) & 63))
        if v1 {
            var v43 int32 = -363
            _ = v43
        } else {
            var v44 int8 = v2
            var v45 uint = (^(v19))
            _ = v44
            _ = v45
        }
        if v1 {
            if (0.1 <= -2.25) {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                v4 = v4
                switch v28 {
                case "hello":
                    for fi46 := 0; fi46 < 15; fi46++ {
                        for fi47 := 0; fi47 < 1; fi47++ {
                            v48 := GenS4{F0: 0, F1: 671}
                            if (v2 >= -89) {
                                v19 = v19
                                var v49 int64 = -987
                                v50 := GenS7{F0: 18446744073709551615, F1: (v49 * (-(v49)))}
                                _ = v49
                                _ = v50
                            }
                            switch fi47 {
                            case -991:
                                if false {
                                    if false {
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        var v51 uint32 = 42
                                        _ = v51
                                    } else {
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        if ((956 <= 1000) && v1) {
                                            fmt.Println(fi46)
                                            switch (v28 + v28) {
                                            case "\n\t":
                                                v28 = v28
                                            case "hello":
                                                v52 := GenS2{F0: 4294967295, F1: fi47}
                                                _ = v52
                                            default:
                                                var v53 int32 = 42
                                                _ = v53
                                            }
                                            switch "a" {
                                            case "hello":
                                                switch (fi47 & -117) {
                                                case 1:
                                                    fmt.Println(fi46, v1, v19)
                                                }
                                                if false {
                                                    v28 = (v28 + v28)
                                                } else {
                                                    switch fi47 {
                                                    case 9223372036854775807:
                                                        v54 := GenS9{F0: int16(fi46)}
                                                        var v55 string = ""
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    switch v28 {
                                                    case "\n\t":
                                                        var v56 int32 = -928
                                                        _ = v56
                                                    default:
                                                        var v57 int8 = (v2 / ((0) | 1))
                                                        _ = v57
                                                    }
                                                    if (2147483647 < 9223372036854775807) {
                                                        fmt.Println(v1, v1)
                                                        v58 := GenS5{F0: (v19 * v19), F1: "foo bar"}
                                                        _ = v58
                                                    }
                                                }
                                                v59 := GenS4{F0: (v4 / (((v4 | v4)) | 1)), F1: 2147483647}
                                                if v1 {
                                                    v60 := GenS1{F0: (v19 - (v19 << (((v19 / ((2) | 1))) & 63))), F1: 2, F2: 562}
                                                    _ = v60
                                                }
                                                switch fi47 {
                                                case 55:
                                                    v3 += 32767
                                                default:
                                                    fmt.Println(fi47)
                                                }
                                                if (42 != 0) {
                                                    v28 = (v28 + v28)
                                                } else {
                                                    var v61 int8 = 127
                                                    if ((649 <= 4294967295) && (v1 || v1)) {
                                                        fmt.Println(v28)
                                                        v62 := GenS6{F0: 0, F1: 42, F2: -32768}
                                                        _ = v62
                                                    }
                                                    var v63 int8 = v2
                                                    _ = v61
                                                    _ = v63
                                                }
                                                switch (("" + v28) + v28) {
                                                case "hello":
                                                    v28 = "hello"
                                                case "a":
                                                    if true {
                                                        v26 = append(v26, (fi46 & (fi47 - fi46)))
                                                        var v64 uint16 = v3
                                                        var v65 int16 = 42
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        v26 = append(v26, (fi46 | fi47))
                                                        v28 = v28
                                                    }
                                                }
                                                _ = v59
                                            case "":
                                                if (fi46 != fi47) {
                                                    switch "" {
                                                    case "foo bar":
                                                        switch v28 {
                                                        case "foo bar":
                                                            var v66 float32 = 1e10
                                                            _ = v66
                                                        case "":
                                                            v2 = v2
                                                        default:
                                                            switch "\n\t" {
                                                            case "":
                                                                v19 += (v19 >> (((v19 & (v19 % (((v19 << (((v19 | 713)) & 63))) | 1)))) & 63))
                                                            case "foo bar":
                                                                var v67 int32 = int32(v4)
                                                                _ = v67
                                                            }
                                                        }
                                                    case "a":
                                                        var v68 float32 = 1.5
                                                        _ = v68
                                                    case "\n\t":
                                                        delete(v5, -620)
                                                    }
                                                    v1 = v1
                                                } else {
                                                    var v69 int16 = 1000
                                                    _ = v69
                                                }
                                            default:
                                                if (-2.25 < (3.14159 + -2.25)) {
                                                    var v70 uint32 = 2147483647
                                                    _ = v70
                                                }
                                                if (919 > 0) {
                                                    if (126 > v19) {
                                                        v4 = v4
                                                        var v71 uint = 920
                                                        _ = v71
                                                    } else {
                                                        v5[(fi46 & fi47)] = (v28 + v28)
                                                    }
                                                    v72 := GenS10{F0: 42, F1: 9223372036854775807, F2: (65535 <= (^(v19)))}
                                                    var v73 uint = (v19 ^ (v19 & 1))
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    var v74 uint64 = 18446744073709551615
                                                    if v1 {
                                                        fmt.Println(v19, v4, fi46)
                                                        if ((fi46 % ((4294967295) | 1)) >= 4294967295) {
                                                            var v75 uint = (v19 | 0)
                                                            fmt.Println(v3, v2, v74)
                                                            var v76 bool = true
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            fmt.Println(v4, v4)
                                                            var v77 bool = v1
                                                            _ = v77
                                                        }
                                                    }
                                                    var v78 bool = v1
                                                    _ = v74
                                                    _ = v78
                                                }
                                                v2 = v2
                                            }
                                        }
                                        v2 = v2
                                        v2 = (v2 << ((18446744073709551615) & 63))
                                        if v1 {
                                            switch (((v28 + v28) + v28) + v28) {
                                            case "a":
                                                v79 := GenS11{F0: 42, F1: 4294967295}
                                                _ = v79
                                            case "foo bar":
                                                if (!(v1 && (159 > 8))) {
                                                    if v1 {
                                                        fmt.Println(fi47, v2, v19)
                                                        v26 = append(v26, (fi47 * (-(fi47))))
                                                        if v1 {
                                                            if v1 {
                                                                switch v28 {
                                                                case "":
                                                                    fmt.Println(fi47)
                                                                default:
                                                                    var v80 uint64 = 1
                                                                    _ = v80
                                                                }
                                                            }
                                                            var v81 uint64 = 42
                                                            _ = v81
                                                        } else {
                                                            fmt.Println(v4, v4)
                                                            fmt.Println(v1)
                                                        }
                                                        var v82 bool = false
                                                        _ = v82
                                                    } else {
                                                        v83 := GenS12{F0: 42}
                                                        v19 += (v19 - (v19 % ((v19) | 1)))
                                                        v2 = 100
                                                        _ = v83
                                                    }
                                                }
                                            }
                                            v1 = v1
                                        }
                                    }
                                } else {
                                    var v84 uint64 = 42
                                    v1 = (-269 >= -1)
                                    fmt.Println(v4)
                                    var v85 float32 = 0.0
                                    if (v1 || (105 <= 43)) {
                                        v28 = v28
                                    } else {
                                        v1 = (2147483647 > 795)
                                        if v1 {
                                            v85 += (1e10 - v85)
                                        } else {
                                            v19 += 42
                                            var v86 uint16 = (v3 | (v4 >> ((v19) & 63)))
                                            _ = v86
                                        }
                                        _ = len(v26)
                                    }
                                    _ = v84
                                    _ = v85
                                }
                            }
                            _ = fi47
                            _ = v48
                        }
                        switch (fi46 + (fi46 ^ 2147483647)) {
                        case -9223372036854775808:
                            var v87 int = fi46
                            _ = v87
                        case -1:
                            fmt.Println(v4, v19)
                        case 32768:
                            fmt.Println(v1, fi46, v28)
                        }
                        v88 := GenS13{F0: (v2 / ((v2) | 1)), F1: 0}
                        _ = fi46
                        _ = v88
                    }
                default:
                    if (!v1) {
                        v19 = 615
                        for fi89 := 0; fi89 < 10; fi89++ {
                            var v90 uint8 = 2
                            _ = fi89
                            _ = v90
                        }
                        v28 = ""
                    } else {
                        var v91 bool = true
                        var v92 int32 = -1
                        for fi93 := 0; fi93 < 2; fi93++ {
                            _ = len(v5)
                            _ = fi93
                        }
                        _ = v91
                        _ = v92
                    }
                }
            } else {
                v2 = (-(v2))
                switch 2147483647 {
                case -277:
                    delete(v5, 937)
                case 42:
                    v3 = (v4 - (v3 | (v4 / ((v4) | 1))))
                case 32767:
                    var v94 uint64 = 171
                    _ = v94
                }
                var v95 uint32 = 65535
                _ = v95
            }
            v4 += (^(v4))
            fmt.Println(v19)
        }
    } else {
        var v96 float32 = 0.0
        _ = v96
    }
    if (!(182 != 0)) {
        v3 = 42
        var v97 uint16 = 248
        _ = v97
    }
    fmt.Println(v19, v3)
    var v98 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v98:", v98)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
