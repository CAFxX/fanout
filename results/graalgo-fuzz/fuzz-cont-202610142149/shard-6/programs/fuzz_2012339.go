package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 int64; F2 int32 }

func main() {
    var v1 int64 = 290
    var v2 uint = 887
    var v3 int16 = 32767
    for fi4 := 0; fi4 < 15; fi4++ {
        fmt.Println(v1)
        var v5 uint32 = 172
        _ = fi4
        _ = v5
    }
    v2 = (^(v2))
    var v6 string = "foo bar"
    v7 := GenS1{F0: -128}
    v8 := []int{2147483647, 1, -9223372036854775808}
    fmt.Println(v3, v3, v1)
    v2 += v2
    var v9 bool = true
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: fi10, F1: 1, F2: (v3 & v3)}
        if (((v2 & (v2 / (((^(v2))) | 1))) > v2) && (42 < 127)) {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int16 = (v3 - v3)
                if (!false) {
                    fmt.Println(v3)
                    if v9 {
                        switch (fi12 * int(v13)) {
                        case -9223372036854775808:
                            v9 = (!v9)
                        default:
                            if v9 {
                                _ = len(v8)
                                var v14 uint = (v2 - v2)
                                var v15 int = int(v2)
                                if v9 {
                                    if v9 {
                                        if v9 {
                                            fmt.Println(v2, fi10)
                                            v16 := GenS3{F0: v15, F1: 42, F2: 819}
                                            v14 = (^(v14))
                                            _ = v16
                                        }
                                    } else {
                                        if (v1 >= (v1 + (v1 % (((v1 | (v1 * 0))) | 1)))) {
                                            v8 = append(v8, fi12)
                                            if v9 {
                                                if (float64((0.0 / 1.0)) < (1.5 / ((1e10 + 3.14159) + 0.1))) {
                                                    if v9 {
                                                        var v17 float64 = 1e10
                                                        _ = v17
                                                    }
                                                    v8 = append(v8, 4294967295)
                                                    fmt.Println(v6, v3, fi10)
                                                } else {
                                                    v18 := GenS1{F0: 0}
                                                    v8 = append(v8, (fi12 + (fi10 >> ((v2) & 63))))
                                                    v14 -= (v14 - (v2 ^ (v14 | v14)))
                                                    _ = v18
                                                }
                                                if v9 {
                                                    if false {
                                                        if v9 {
                                                            var v19 bool = (!(false || v9))
                                                            _ = v19
                                                        } else {
                                                            switch v6 {
                                                            case "\n\t":
                                                                v3 += (-(v13))
                                                            case "hello":
                                                                if ((18446744073709551615 <= 244) && true) {
                                                                    v20 := GenS4{F0: 2, F1: 128}
                                                                    var v21 bool = ((255 != 128) && (42 >= (v14 / (((v2 % ((v14) | 1))) | 1))))
                                                                    _ = v20
                                                                    _ = v21
                                                                } else {
                                                                    fmt.Println(v13)
                                                                }
                                                            }
                                                            if v9 {
                                                                if (!(!(1 == 4294967295))) {
                                                                    fmt.Println(fi10, v6)
                                                                    if v9 {
                                                                        switch (v15 * -9223372036854775808) {
                                                                        case 101:
                                                                            if true {
                                                                                var v22 uint = v14
                                                                                _ = v22
                                                                            } else {
                                                                                switch fi10 {
                                                                                case 2147483647:
                                                                                    fmt.Println(v13)
                                                                                    v15 += v15
                                                                                case 1:
                                                                                    v23 := GenS5{F0: "hello"}
                                                                                    _ = v23
                                                                                case -739:
                                                                                    v24 := GenS2{F0: fi12, F1: 4294967295, F2: v13}
                                                                                    switch v6 {
                                                                                    case "hello":
                                                                                        switch v6 {
                                                                                        case "foo bar":
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                v15 += fi12
                                                                                            case "foo bar":
                                                                                                if v9 {
                                                                                                    var v25 bool = v9
                                                                                                    fmt.Println(v1)
                                                                                                    if v9 {
                                                                                                        fmt.Println(v3, v2, v1)
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(fi10)
                                                                                                        case "\n\t":
                                                                                                            if len(v8) >= 2 {
                                                                                                                v8 = v8[1:len(v8)]
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            _ = len(v8)
                                                                                                        }
                                                                                                        v2 -= v2
                                                                                                    }
                                                                                                    if (v13 != 1) {
                                                                                                        var v26 bool = (false || true)
                                                                                                        _ = v26
                                                                                                    }
                                                                                                    _ = v25
                                                                                                } else {
                                                                                                    v13 = (v3 % ((v13) | 1))
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = 0
                                                                                                    }
                                                                                                    v6 = (((v6 + (v6 + "\n\t")) + v6) + "")
                                                                                                }
                                                                                            case "":
                                                                                                switch -716 {
                                                                                                case 1:
                                                                                                    var v27 int16 = (v13 * (^(v3)))
                                                                                                    v28 := GenS1{F0: -1}
                                                                                                    _ = v27
                                                                                                    _ = v28
                                                                                                case -1:
                                                                                                    if (float32(-2.25) == 3.14159) {
                                                                                                        v6 = v6
                                                                                                    } else {
                                                                                                        _ = len(v8)
                                                                                                        _ = len(v8)
                                                                                                        var v29 int8 = -72
                                                                                                        var v30 bool = (259 != 569)
                                                                                                        _ = v29
                                                                                                        _ = v30
                                                                                                    }
                                                                                                case -9223372036854775808:
                                                                                                    v31 := GenS6{F0: 2, F1: 1}
                                                                                                    _ = v31
                                                                                                default:
                                                                                                    fmt.Println(v2)
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            if (1 <= v13) {
                                                                                                switch v6 {
                                                                                                case "a":
                                                                                                    v32 := GenS7{F0: (v14 + v2), F1: 1}
                                                                                                    _ = v32
                                                                                                case "hello":
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = fi10
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    _ = len(v8)
                                                                                                }
                                                                                                var v33 int64 = v1
                                                                                                if len(v8) >= 2 {
                                                                                                    v8 = v8[1:len(v8)]
                                                                                                }
                                                                                                _ = v33
                                                                                            } else {
                                                                                                _ = len(v8)
                                                                                                if (1 < 255) {
                                                                                                    if false {
                                                                                                        v14 += (v14 << (((v14 % (((v2 & (^(v2)))) | 1))) & 63))
                                                                                                    } else {
                                                                                                        var v34 float64 = ((float64(float32(3.14159)) + (1e10 * 3.14159)) / float64(1e10))
                                                                                                        _ = v34
                                                                                                    }
                                                                                                    v9 = v9
                                                                                                    if ((0.0 / 1.0) != 1.5) {
                                                                                                        v35 := GenS5{F0: v6}
                                                                                                        fmt.Println(v14, v1, v9)
                                                                                                        if (!(147 < v2)) {
                                                                                                            var v36 uint8 = 127
                                                                                                            v37 := GenS1{F0: -11}
                                                                                                            var v38 int32 = -1
                                                                                                            _ = v36
                                                                                                            _ = v37
                                                                                                            _ = v38
                                                                                                        }
                                                                                                        _ = v35
                                                                                                    }
                                                                                                    fmt.Println(v1)
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = v15
                                                                                                    }
                                                                                                }
                                                                                                var v39 uint = 0
                                                                                                _ = v39
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v40 := GenS8{F0: -113}
                                                                                        v13 -= v3
                                                                                        if ((1e10 < float64((((0.0 / 3.14159) - (3.14159 + 3.14159)) - 0.1))) && (v9 && v9)) {
                                                                                            if (109 <= -1) {
                                                                                                v13 += (^(v13))
                                                                                            } else {
                                                                                                fmt.Println(v3)
                                                                                                v15 += fi12
                                                                                            }
                                                                                        }
                                                                                        _ = v40
                                                                                    default:
                                                                                        v1 = v1
                                                                                    }
                                                                                    fmt.Println(v14, v6, v1)
                                                                                    _ = v24
                                                                                }
                                                                            }
                                                                            if (((v14 * 42) < v2) && v9) {
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                                v9 = (!v9)
                                                                                v41 := GenS6{F0: 1, F1: 42}
                                                                                _ = v41
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if v9 {
                                                                            if v9 {
                                                                                var v42 uint16 = 32767
                                                                                _ = v42
                                                                            }
                                                                            var v43 uint = (^(v14))
                                                                            fmt.Println(v13)
                                                                            _ = v43
                                                                        } else {
                                                                            _ = len(v8)
                                                                            if false {
                                                                                v14 += (v2 + v2)
                                                                                fmt.Println(v9)
                                                                                v13 += int16(v1)
                                                                                var v44 float64 = float64(1.5)
                                                                                v9 = (4294967295 == 2)
                                                                                _ = v44
                                                                            }
                                                                            if ((-(v3)) <= (v3 * (v3 ^ 0))) {
                                                                                v2 = (v2 & v2)
                                                                                if v9 {
                                                                                    var v45 string = v6
                                                                                    if true {
                                                                                        v8 = append(v8, fi10)
                                                                                    }
                                                                                    fmt.Println(fi12, v13, fi12)
                                                                                    v2 += (v2 ^ v2)
                                                                                    _ = v45
                                                                                } else {
                                                                                    var v46 int64 = (v1 / ((v1) | 1))
                                                                                    if v9 {
                                                                                        v3 = 1
                                                                                        if v9 {
                                                                                            var v47 int64 = 9223372036854775807
                                                                                            v46 -= v46
                                                                                            _ = v47
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v6)
                                                                                    }
                                                                                    _ = v46
                                                                                }
                                                                            } else {
                                                                                var v48 bool = (!(!false))
                                                                                _ = v48
                                                                            }
                                                                            switch v15 {
                                                                            case -2147483648:
                                                                                var v49 int16 = (v3 >> ((v14) & 63))
                                                                                _ = v49
                                                                            case 42:
                                                                                v50 := GenS1{F0: -1}
                                                                                _ = v50
                                                                            }
                                                                        }
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                    }
                                                                    v14 = 1
                                                                } else {
                                                                    var v51 string = v6
                                                                    v6 = v51
                                                                    _ = v51
                                                                }
                                                                if false {
                                                                    var v52 int64 = v1
                                                                    _ = v52
                                                                }
                                                                if v9 {
                                                                    v53 := GenS5{F0: "\n\t"}
                                                                    v13 += (v3 - v13)
                                                                    v54 := GenS2{F0: (-(fi10)), F1: 42, F2: (v13 & (-(v13)))}
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                var v55 uint8 = 127
                                                                if false {
                                                                    v56 := GenS7{F0: (v2 | v2), F1: 945}
                                                                    v57 := GenS2{F0: int(v1), F1: uint64(v55), F2: (^(v13))}
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                switch v6 {
                                                                case "a":
                                                                    var v58 float64 = (1.5 - 3.14159)
                                                                    _ = v58
                                                                default:
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi10
                                                                    }
                                                                }
                                                                _ = v55
                                                            } else {
                                                                fmt.Println(v15)
                                                            }
                                                        }
                                                    } else {
                                                        if ((v14 << ((183) & 63)) > v14) {
                                                            v8 = append(v8, (v15 * (v15 % ((v15) | 1))))
                                                            var v59 float64 = 1e10
                                                            v60 := GenS5{F0: (v6 + (v6 + v6))}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        switch (fi12 * fi12) {
                                                        case -71:
                                                            var v61 int = (fi10 / (((fi12 << (((v2 * 18446744073709551615)) & 63))) | 1))
                                                            _ = v61
                                                        default:
                                                            v13 = v13
                                                        }
                                                        v9 = false
                                                    }
                                                    if len(v8) > 0 {
                                                        v8[0] = (v15 & (-(fi12)))
                                                    }
                                                    if (1e10 >= (1e10 + -2.25)) {
                                                        fmt.Println(v1, fi10)
                                                        _ = len(v8)
                                                    } else {
                                                        v1 -= (v1 | (-(v1)))
                                                        v62 := GenS5{F0: v6}
                                                        _ = v62
                                                    }
                                                }
                                                fmt.Println(v13, v1, v3)
                                            }
                                            var v63 int = 555
                                            if (!(v9 || v9)) {
                                                var v64 int64 = (v1 | (-(v1)))
                                                v65 := GenS7{F0: (v14 & v14), F1: 255}
                                                _ = v64
                                                _ = v65
                                            }
                                            if v9 {
                                                fmt.Println(fi10, v13)
                                                fmt.Println(v6)
                                                v63 -= fi10
                                                v66 := GenS3{F0: fi10, F1: v1, F2: 4294967295}
                                                _ = v66
                                            }
                                            v67 := GenS3{F0: (v15 & v15), F1: v1, F2: 0}
                                            _ = v63
                                            _ = v67
                                        }
                                    }
                                    v68 := GenS1{F0: 37}
                                    _ = v68
                                } else {
                                    v69 := GenS9{F0: 42}
                                    v15 = 9223372036854775807
                                    _ = v69
                                }
                                v15 += (-(fi12))
                                if false {
                                    var v70 bool = (164 != 1)
                                    var v71 string = v6
                                    fmt.Println(v3, v15, v2)
                                    _ = v70
                                    _ = v71
                                } else {
                                    switch (fi12 >> ((v2) & 63)) {
                                    case 230:
                                        v14 = 65535
                                    case -1:
                                        var v72 uint = v14
                                        _ = v72
                                    case 9223372036854775807:
                                        var v73 int16 = (-(v13))
                                        _ = v73
                                    default:
                                        var v74 int8 = 127
                                        _ = v74
                                    }
                                    v75 := GenS10{F0: 2147483648}
                                    v3 = v3
                                    _ = v75
                                }
                                _ = v14
                                _ = v15
                            }
                        }
                        v3 -= v3
                        fmt.Println(v3)
                        fmt.Println(v13)
                        v8 = append(v8, (fi10 | fi12))
                        v76 := GenS11{F0: (v1 | v1), F1: v1, F2: 42}
                        if true {
                            if false {
                                v8 = append(v8, fi10)
                            } else {
                                v77 := GenS8{F0: -2147483648}
                                _ = v77
                            }
                            v6 = (v6 + v6)
                        }
                        v2 -= (^(v2))
                        _ = v76
                    }
                }
                switch "" {
                case "hello":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "a":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    var v78 int = fi10
                    _ = v78
                }
                v2 -= (^(v2))
                fmt.Println(fi12)
                _ = fi12
                _ = v13
            }
            switch v6 {
            case "":
                v79 := GenS10{F0: 42}
                _ = v79
            case "foo bar":
                var v80 uint64 = 545
                _ = v80
            case "hello":
                if len(v8) > 0 {
                    v8[0] = 32767
                }
            default:
                if true {
                    var v81 uint16 = 643
                    _ = v81
                }
            }
            var v82 uint32 = 272
            _ = v82
        } else {
            v1 += v1
            for fi83 := 0; fi83 < 4; fi83++ {
                switch (fi83 * fi83) {
                case 42:
                    v84 := GenS11{F0: v1, F1: v1, F2: 42}
                    if (1 <= v2) {
                        v6 = "hello"
                        var v85 uint = 255
                        v1 += (v1 * v1)
                        if ((((1e10 / float64(1e10)) / 1.5) / 1e10) > (1.5 + 0.1)) {
                            var v86 int8 = 121
                            fmt.Println(v86, v9)
                            _ = v86
                        } else {
                            if (!v9) {
                                v85 = v85
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if false {
                                    fmt.Println(fi83, fi83)
                                } else {
                                    var v87 int32 = 1
                                    var v88 uint16 = 42
                                    if len(v8) > 0 {
                                        v8[0] = 42
                                    }
                                    _ = v87
                                    _ = v88
                                }
                            }
                            fmt.Println(fi10, v3)
                            v9 = v9
                        }
                        _ = v85
                    }
                    _ = v84
                case 0:
                    if v9 {
                        var v89 uint = 2
                        _ = v89
                    } else {
                        var v90 int = (^(fi10))
                        var v91 bool = (!(true || true))
                        _ = v90
                        _ = v91
                    }
                case 9223372036854775807:
                    fmt.Println(fi10, v6, v9)
                default:
                    switch "a" {
                    case "a":
                        var v92 float32 = ((float32((1.5 + 1e10)) - 1e10) * 1e10)
                        _ = v92
                    case "hello":
                        if true {
                            var v93 uint64 = 18446744073709551615
                            _ = v93
                        }
                    case "foo bar":
                        v2 = v2
                    }
                }
                v94 := GenS4{F0: (v2 | (v2 << ((v2) & 63))), F1: 255}
                _ = fi83
                _ = v94
            }
        }
        for fi95 := 0; fi95 < 6; fi95++ {
            fmt.Println(v6, v9, v2)
            v1 = 1
            _ = fi95
        }
        _ = fi10
        _ = v11
    }
    if false {
        v96 := GenS4{F0: v2, F1: 128}
        _ = v96
    } else {
        if (v6 < v6) {
            if v9 {
                for fi97 := 0; fi97 < 14; fi97++ {
                    v6 = v6
                    _ = fi97
                }
                fmt.Println(v3, v2)
                v3 += v3
            } else {
                switch 1 {
                case 900:
                    fmt.Println(v1, v6)
                case 127:
                    switch (v6 + v6) {
                    case "":
                        for fi98 := 0; fi98 < 20; fi98++ {
                            for fi99 := 0; fi99 < 17; fi99++ {
                                if v9 {
                                    _ = len(v8)
                                    var v100 int8 = -10
                                    _ = v100
                                } else {
                                    v9 = (-19 <= -1)
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = fi99
                            }
                            _ = fi98
                        }
                    }
                case -9223372036854775808:
                    fmt.Println(v2, v1, v2)
                default:
                    v3 = -254
                }
            }
            var v101 int = 4294967295
            var v102 bool = v9
            _ = v101
            _ = v102
        } else {
            for fi103 := 0; fi103 < 9; fi103++ {
                v6 = ("foo bar" + (v6 + (v6 + v6)))
                fmt.Println(v9, v6)
                v8 = append(v8, 4294967295)
                _ = fi103
            }
            if len(v8) > 0 {
                v8[0] = 4294967295
            }
            var v104 int64 = (v1 << ((4294967295) & 63))
            _ = v104
        }
    }
    var v105 uint = 2
    switch "" {
    case "":
        v106 := GenS1{F0: 1}
        _ = v106
    default:
        var v107 float64 = 0.1
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
