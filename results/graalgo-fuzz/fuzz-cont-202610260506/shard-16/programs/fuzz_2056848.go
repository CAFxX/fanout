package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 bool; F1 string }
func genFn0() int8 {
    var v6 uint8 = 68
    v7 := []int{9223372036854775807, 32767}
    if true {
        switch 2147483647 {
        case 852:
            v8 := GenS1{F0: -244, F1: -1}
            _ = v8
        }
        v7 = append(v7, 0)
    }
    var v9 uint64 = 728
    v6 -= (v6 << ((2) & 63))
    _ = v6
    _ = v7
    _ = v9
    return int8(v6)
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = 56
    var v3 string = "a"
    var v4 int = 0
    var v5 int32 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: (1 > int32(v4))}
        _ = fi10
        _ = v11
    }
    var v12 int16 = -32768
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            if true {
                for fi14 := 0; fi14 < 18; fi14++ {
                    fmt.Println(v4, fi14, fi13)
                    _ = fi14
                }
                v12 = 32767
                v4 -= v4
            }
            for fi15 := 0; fi15 < 3; fi15++ {
                switch (fi15 | fi13) {
                case 32767:
                    v3 = ""
                    v1 = 127
                    var v16 uint8 = 128
                    _ = v16
                default:
                    var v17 bool = true
                    _ = v17
                }
                _ = fi15
            }
            _ = fi13
        }
        if false {
            v18 := []int{(v4 / ((4294967295) | 1))}
            switch v4 {
            case -1:
                if (!true) {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        fmt.Println(fi19, v1)
                        var v20 uint = 2
                        var v21 int = 32767
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS2{F0: (0 < 2)}
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if false {
                            v3 = "foo bar"
                            if (42 == 42) {
                                fmt.Println(v5)
                                var v24 uint16 = 32767
                                var v25 int64 = 0
                                _ = v24
                                _ = v25
                            }
                        } else {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                var v27 uint8 = v2
                                v28 := GenS1{F0: int(v12), F1: 42}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v4 += (v4 & fi23)
                        }
                        v12 = (v12 | v12)
                        var v29 uint32 = 406
                        _ = fi23
                        _ = v29
                    }
                    _ = v22
                }
            case 32767:
                if (false || (((!false) && (((v2 % (((v1 & v1)) | 1)) != v2) && (false && (!false)))) && true)) {
                    var v30 string = "foo bar"
                    var v31 int = 9223372036854775807
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint = 4294967295
                    _ = v32
                }
            case 1:
                switch v4 {
                case -9223372036854775808:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        if (!(65535 <= 717)) {
                            if (!(320 != 42)) {
                                var v34 uint8 = (^(v1))
                                var v35 uint64 = 18446744073709551615
                                _ = v34
                                _ = v35
                            } else {
                                if false {
                                    if (255 < 40) {
                                        var v36 float64 = float64((((((1.5 + float32((0.1 / 0.1))) * 1e10) * (0.0 - 1.5)) * (0.1 - 1e10)) / 1.5))
                                        v5 = v5
                                        _ = v36
                                    } else {
                                        fmt.Println(v1)
                                    }
                                    var v37 uint32 = 261
                                    v1 = (v1 + (v1 - 151))
                                    _ = v37
                                }
                                _ = len(v18)
                                v3 = v3
                            }
                            fmt.Println(v3, fi33, v1)
                            v12 = v12
                        }
                        if true {
                            v18 = append(v18, (fi33 / ((fi33) | 1)))
                            if true {
                                v5 = v5
                                v5 = v5
                                var v38 uint32 = 65535
                                _ = v38
                            }
                        }
                        _ = fi33
                    }
                case -533:
                    var v39 uint32 = 1
                    _ = v39
                }
            }
            _ = v18
        } else {
            fmt.Println(v5, v12)
        }
    }
    v3 = "a"
    v3 = ("foo bar" + v3)
    switch "foo bar" {
    case "foo bar":
        v2 += 255
    }
    fmt.Println(v12, v12)
    v40 := GenS3{F0: genFn0()}
    var v41 int = (v4 & v4)
    v42 := []int{(v4 / ((725) | 1)), (v4 / ((v41) | 1)), v41, v4}
    switch ((((("" + "foo bar") + v3) + v3) + "\n\t") + v3) {
    case "hello":
        fmt.Println(v12)
    case "":
        v4 += (v4 << ((1) & 63))
    case "a":
        var v43 int16 = v12
        _ = v43
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        if false {
            for fi45 := 0; fi45 < 1; fi45++ {
                v4 += -814
                if false {
                    v46 := GenS3{F0: 100}
                    v3 = "foo bar"
                    v5 = (-(v5))
                    v3 = v3
                    _ = v46
                } else {
                    switch v4 {
                    case -9223372036854775808:
                        if false {
                            _ = len(v42)
                        }
                    case -2147483648:
                        if false {
                            if true {
                                v5 += (^(v5))
                                var v47 int64 = -9223372036854775808
                                v1 = v1
                                _ = v47
                            }
                            v4 = (v4 + (v4 & (-(v4))))
                            if (!true) {
                                if (!(true && (1 != 1))) {
                                    if (((1 <= 65535) || (((^(v2)) != 255) && false)) || false) {
                                        _ = len(v42)
                                        v3 = "hello"
                                        if true {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v48 := GenS3{F0: 127}
                                                _ = v48
                                            case "\n\t":
                                                if true {
                                                    var v49 int16 = (v12 / ((1) | 1))
                                                    v50 := GenS4{F0: true, F1: (v12 + (v12 % ((42) | 1)))}
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                    var v51 float64 = (0.1 * 1e10)
                                                    _ = v51
                                                }
                                            default:
                                                var v52 string = (v3 + "foo bar")
                                                _ = v52
                                            }
                                            var v53 int64 = 9223372036854775807
                                            _ = v53
                                        }
                                    } else {
                                        var v54 uint8 = 216
                                        var v55 uint16 = 32767
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 string = (v3 + "\n\t")
                                    fmt.Println(v5)
                                    _ = v56
                                } else {
                                    switch (fi45 + fi44) {
                                    case 127:
                                        var v57 float64 = (float64(0.1) + 3.14159)
                                        _ = v57
                                    case -292:
                                        if ((^(v12)) < 0) {
                                            _ = len(v42)
                                            v2 = 1
                                            var v58 string = ((((v3 + v3) + v3) + v3) + "\n\t")
                                            fmt.Println(v12, fi45)
                                            var v59 string = v58
                                            _ = v58
                                            _ = v59
                                        }
                                    case 2147483647:
                                        var v60 bool = false
                                        _ = v60
                                    }
                                    switch "" {
                                    case "\n\t":
                                        if len(v42) > 0 {
                                            v42[0] = (v41 >> ((18) & 63))
                                        }
                                        if (42 <= -128) {
                                            var v61 uint64 = 2
                                            var v62 uint64 = v61
                                            _ = v61
                                            _ = v62
                                        }
                                    case "a":
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                    case "hello":
                                        switch (fi45 & v41) {
                                        case 4294967295:
                                            if false {
                                                fmt.Println(v3)
                                            } else {
                                                if ((!true) && false) {
                                                    var v63 uint = 2
                                                    var v64 int32 = (-(v5))
                                                    var v65 int8 = -128
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                                v66 := GenS4{F0: (false || false), F1: 0}
                                                v5 += (v5 ^ v5)
                                                _ = v66
                                            }
                                        default:
                                            var v67 int16 = (-(v12))
                                            _ = v67
                                        }
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            if true {
                                                var v68 int = (fi44 ^ (v41 | fi44))
                                                _ = v68
                                            } else {
                                                if ((fi44 + 768) < (-(v41))) {
                                                    var v69 int16 = v12
                                                    _ = v69
                                                }
                                                var v70 int32 = v5
                                                _ = v70
                                            }
                                            v71 := GenS3{F0: 127}
                                            _ = v71
                                        } else {
                                            v41 += (v41 & (^(v41)))
                                            var v72 float64 = (-2.25 / (1.5 / ((float64(1.5) * 3.14159) / (1.5 + 0.1))))
                                            var v73 uint16 = 0
                                            _ = v72
                                            _ = v73
                                        }
                                    default:
                                        var v74 int = (v4 >> ((813) & 63))
                                        _ = v74
                                    }
                                    if (32767 >= 1) {
                                        var v75 int64 = -9223372036854775808
                                        if (v75 < (v75 * v75)) {
                                            switch "\n\t" {
                                            case "":
                                                var v76 int32 = -2147483648
                                                _ = v76
                                            case "foo bar":
                                                var v77 int = fi44
                                                _ = v77
                                            default:
                                                v12 = (v12 ^ v12)
                                            }
                                            _ = len(v42)
                                            var v78 string = "hello"
                                            var v79 uint8 = (v2 & (v2 / ((v2) | 1)))
                                            _ = v78
                                            _ = v79
                                        } else {
                                            v80 := GenS3{F0: genFn0()}
                                            _ = len(v42)
                                            switch v3 {
                                            case "a":
                                                v4 = 0
                                            case "\n\t":
                                                var v81 float64 = (0.0 / 1.5)
                                                _ = v81
                                            case "hello":
                                                var v82 int16 = (v12 * v12)
                                                _ = v82
                                            }
                                            _ = v80
                                        }
                                        v42 = append(v42, fi44)
                                        _ = v75
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if (665 < 0) {
                                            if ((((0.0 + float64(1e10)) - -2.25) / float64(1e10)) != 1.5) {
                                                var v83 int32 = -1
                                                _ = v83
                                            } else {
                                                v84 := GenS5{F0: 32768, F1: 745, F2: 328}
                                                _ = v84
                                            }
                                            if true {
                                                v42 = append(v42, (v41 + fi44))
                                                var v85 int = -180
                                                v85 += int(v12)
                                                _ = v85
                                            } else {
                                                v42 = append(v42, (-(v41)))
                                            }
                                        } else {
                                            var v86 bool = (-2.25 == (1e10 + 3.14159))
                                            v2 -= v2
                                            _ = v86
                                        }
                                    case "hello":
                                        switch fi44 {
                                        case -9223372036854775808:
                                            v87 := GenS5{F0: (fi45 * 261), F1: 699, F2: 1}
                                            _ = v87
                                        case 42:
                                            switch 700 {
                                            case 42:
                                                var v88 uint = 630
                                                _ = v88
                                            case 9223372036854775807:
                                                fmt.Println(v1, v12)
                                                if len(v42) >= 2 {
                                                    v42 = v42[1:len(v42)]
                                                }
                                                v5 += 2147483647
                                            case 526:
                                                v4 = v4
                                            default:
                                                var v89 string = v3
                                                _ = v89
                                            }
                                        }
                                    default:
                                        v12 += (v12 % ((v12) | 1))
                                    }
                                    v2 += v2
                                    if (!(true && (!true))) {
                                        v90 := GenS6{F0: v5}
                                        fmt.Println(v1, v2)
                                        _ = v90
                                    } else {
                                        v91 := GenS7{F0: v3}
                                        _ = v91
                                    }
                                }
                                if len(v42) > 0 {
                                    v42[0] = (^(v41))
                                }
                                if false {
                                    fmt.Println(v4, v5, v3)
                                } else {
                                    var v92 bool = (false && false)
                                    var v93 uint16 = 42
                                    v94 := GenS8{F0: true, F1: 1}
                                    _ = v92
                                    _ = v93
                                    _ = v94
                                }
                            } else {
                                if (233 < 4294967295) {
                                    var v95 int32 = 1
                                    _ = v95
                                } else {
                                    v4 = (-(v41))
                                }
                            }
                            v41 = (fi44 / ((-797) | 1))
                            switch v3 {
                            case "foo bar":
                                var v96 string = "hello"
                                _ = v96
                            case "":
                                fmt.Println(v1, v1, v41)
                            case "a":
                                v12 += 8
                            default:
                                var v97 float64 = (1e10 * 0.1)
                                if (((4294967295 > 18446744073709551615) && (int(v12) >= (v4 >> ((uint(v1)) & 63)))) && true) {
                                    if (v5 != v5) {
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                        v2 = 1
                                    } else {
                                        var v98 float64 = 1.5
                                        v99 := GenS9{F0: 255, F1: 127}
                                        var v100 uint64 = 18446744073709551615
                                        v5 -= (-(v5))
                                        _ = v100
                                        _ = v98
                                        _ = v99
                                    }
                                    fmt.Println(v12, v3, v41)
                                }
                                v4 = v4
                                _ = v97
                            }
                        } else {
                            var v101 uint64 = 18446744073709551615
                            _ = v101
                        }
                    case 182:
                        switch v3 {
                        case "hello":
                            v3 = v3
                        }
                    }
                    if false {
                        v102 := GenS10{F0: ((42 != -61) || (false && false)), F1: v3}
                        if false {
                            v12 -= (v12 - 770)
                            fmt.Println(v41, v41)
                        } else {
                            var v103 uint64 = 0
                            _ = v103
                        }
                        var v104 uint16 = 78
                        _ = v102
                        _ = v104
                    } else {
                        var v105 uint8 = 127
                        switch 127 {
                        case 127:
                            switch v3 {
                            case "hello":
                                var v106 uint8 = (v1 / (((^(v105))) | 1))
                                _ = v106
                            case "":
                                v107 := GenS8{F0: false, F1: 2}
                                _ = v107
                            }
                        default:
                            v4 = -1
                        }
                        if len(v42) > 0 {
                            v42[0] = 0
                        }
                        _ = v105
                    }
                }
                _ = fi45
            }
            fmt.Println(v4, fi44, v12)
            fmt.Println(v5)
        } else {
            v42 = append(v42, fi44)
            v108 := GenS6{F0: v5}
            _ = v108
        }
        fmt.Println(v4, v12)
        var v109 string = v3
        _ = fi44
        _ = v109
    }
    switch ("" + "\n\t") {
    case "":
        var v110 float64 = float64(0.0)
        _ = v110
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
