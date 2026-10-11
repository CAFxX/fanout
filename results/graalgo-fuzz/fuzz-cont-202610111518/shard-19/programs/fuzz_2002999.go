package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int16; F1 int; F2 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint32; F1 int; F2 string }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint16; F1 string; F2 int }
type GenS15 struct { F0 uint8; F1 uint; F2 string }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint32; F1 int8 }
type GenS18 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS19 struct { F0 int; F1 uint }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "a"
    var v3 int32 = -519
    var v4 string = v2
    if false {
        v2 = "foo bar"
    } else {
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := GenS1{F0: 42}
            v4 = "foo bar"
            _ = fi5
            _ = v6
        }
        var v7 float32 = 3.14159
        var v8 uint = 2
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 18446744073709551615, F1: -1}
    fmt.Println(v4, v3)
    var v10 int = 32768
    if true {
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint16 = 1
                _ = fi12
                _ = v13
            }
            if false {
                v1[v10] = (255 == 0)
                v14 := GenS3{F0: -302}
                v15 := v1[fi11]
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 3; fi16++ {
                    if (((9223372036854775807 >= 128) && true) && (!(((v10 >> ((2) & 63)) <= (^(fi16))) || true))) {
                        v17 := GenS4{F0: 1, F1: (v3 + v3)}
                        v3 = v3
                        _ = v17
                    } else {
                        switch (fi16 << ((0) & 63)) {
                        case 593:
                            switch "\n\t" {
                            case "":
                                if (1 != 1) {
                                    v10 -= fi16
                                    switch 9223372036854775807 {
                                    case 128:
                                        v18 := GenS2{F0: 18446744073709551615, F1: 1}
                                        _ = v18
                                    default:
                                        var v19 uint8 = 42
                                        _ = v19
                                    }
                                    v2 = (v4 + (v2 + ("hello" + ((v2 + v4) + v2))))
                                    v4 = "\n\t"
                                    var v20 int = (fi16 / ((fi11) | 1))
                                    _ = v20
                                } else {
                                    var v21 bool = (0 <= 128)
                                    var v22 int32 = (v3 & v3)
                                    v23 := GenS1{F0: 0}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            case "hello":
                                switch (("\n\t" + ((v4 + v4) + v4)) + v2) {
                                case "foo bar":
                                    switch ("\n\t" + (v4 + (v2 + (("" + v4) + v2)))) {
                                    case "\n\t":
                                        switch 1 {
                                        case 42:
                                            v4 = (("foo bar" + "\n\t") + ((v2 + v4) + v4))
                                        case -9223372036854775808:
                                            var v24 int = 32767
                                            _ = v24
                                        default:
                                            if false {
                                                v10 = (-(v10))
                                                delete(v1, v10)
                                                v10 = v10
                                            }
                                        }
                                    case "a":
                                        switch v4 {
                                        case "\n\t":
                                            if false {
                                                var v25 int16 = -145
                                                if true {
                                                    switch v2 {
                                                    case "foo bar":
                                                        v26 := GenS5{F0: 255}
                                                        _ = v26
                                                    default:
                                                        v27 := GenS6{F0: 125, F1: 190, F2: false}
                                                        _ = v27
                                                    }
                                                    var v28 int = fi16
                                                    _ = v28
                                                }
                                                _ = v25
                                            } else {
                                                if (128 == fi16) {
                                                    fmt.Println(v2, v2, v3)
                                                    v4 = "a"
                                                }
                                            }
                                        default:
                                            if ((-2.25 / 0.1) > 3.14159) {
                                                v10 = fi16
                                                v2 = v4
                                            } else {
                                                v10 = 32768
                                                switch v2 {
                                                case "":
                                                    switch v10 {
                                                    case -367:
                                                        fmt.Println(fi11)
                                                    case 665:
                                                        delete(v1, (v10 << ((0) & 63)))
                                                    case 9223372036854775807:
                                                        delete(v1, fi16)
                                                    }
                                                }
                                                var v29 uint64 = 42
                                                if true {
                                                    var v30 float32 = -2.25
                                                    var v31 uint8 = 2
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    fmt.Println(v10, v4, v29)
                                                }
                                                switch (("\n\t" + (v4 + (((v4 + "foo bar") + "foo bar") + (v4 + "")))) + ((v2 + ("\n\t" + (v2 + ""))) + (v2 + ((v4 + v2) + ("foo bar" + v2))))) {
                                                case "a":
                                                    var v32 int32 = -156
                                                    _ = v32
                                                }
                                                _ = v29
                                            }
                                        }
                                    }
                                }
                            }
                        case 128:
                            v4 = v4
                        case 42:
                            fmt.Println(v3, v4, v3)
                        }
                        v4 = "foo bar"
                        if true {
                            if (39 < 100) {
                                var v33 float32 = (float32(((float64(1e10) / 1.5) * 3.14159)) + float32(-2.25))
                                v2 = v4
                                var v34 uint = 4294967295
                                v33 = v33
                                v3 -= v3
                                _ = v33
                                _ = v34
                            }
                            v10 += v10
                            switch 901 {
                            case -402:
                                v2 = ("" + "")
                            case 0:
                                v35 := GenS1{F0: -1}
                                if (30 < 136) {
                                    switch (fi11 & 127) {
                                    case -357:
                                        v36 := GenS7{F0: 2, F1: (2 < 2)}
                                        var v37 float64 = float64(3.14159)
                                        v1[fi16] = false
                                        _ = v36
                                        _ = v37
                                    case 32768:
                                        if (true || (65535 == 18446744073709551615)) {
                                            v10 = 0
                                        } else {
                                            v38 := GenS8{F0: -2147483648, F1: 2}
                                            v39 := GenS6{F0: 65, F1: 42, F2: true}
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                v2 = ""
                                                var v40 float32 = 3.14159
                                                v4 = (("" + (v4 + ((v4 + v4) + v2))) + v4)
                                                _ = v40
                                            default:
                                                v2 = ("\n\t" + ("" + ("" + ("" + "foo bar"))))
                                            }
                                            v4 = v4
                                            if ((v3 >= (v3 % ((0) | 1))) && (2 != 1)) {
                                                v41 := []int{2147483647, (fi11 << ((4294967295) & 63))}
                                                switch v4 {
                                                case "\n\t":
                                                    if (-1 <= 0) {
                                                        if len(v41) >= 2 {
                                                            v41 = v41[1:len(v41)]
                                                        }
                                                        var v42 float64 = float64((1e10 * float32(((0.0 - 0.0) * float64(((0.1 + ((3.14159 - float32(((1e10 + -2.25) + 1e10))) + 1e10)) + 3.14159))))))
                                                        fmt.Println(v3, v42)
                                                        _ = v42
                                                    } else {
                                                        if (true && (1 <= 42)) {
                                                            if false {
                                                                v43 := GenS9{F0: 276, F1: false}
                                                                _ = v43
                                                            }
                                                            switch (v4 + (((("a" + v2) + "") + v2) + "\n\t")) {
                                                            case "a":
                                                                v10 = (^(fi16))
                                                                switch 9223372036854775807 {
                                                                case 2147483647:
                                                                    var v44 bool = true
                                                                    _ = v44
                                                                case 32767:
                                                                    var v45 bool = true
                                                                    _ = v45
                                                                case -527:
                                                                    fmt.Println(fi16, v10, v4)
                                                                }
                                                                if (1 <= -5) {
                                                                    _ = len(v1)
                                                                } else {
                                                                    if ("hello" >= ((v2 + ("foo bar" + v2)) + "a")) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v46 uint8 = 128
                                                                            _ = v46
                                                                        case "hello":
                                                                            var v47 string = v4
                                                                            _ = v47
                                                                        case "a":
                                                                            var v48 uint8 = 1
                                                                            _ = v48
                                                                        }
                                                                        var v49 float64 = (1.5 - (1.5 * (-2.25 * 0.1)))
                                                                        if false {
                                                                            switch fi16 {
                                                                            case 127:
                                                                                v4 = "hello"
                                                                            case -326:
                                                                                v41 = append(v41, 0)
                                                                            case 4294967295:
                                                                                v10 += v10
                                                                                var v50 uint = 18446744073709551615
                                                                                _ = v50
                                                                            default:
                                                                                v10 = (^(v10))
                                                                            }
                                                                            var v51 uint16 = 65535
                                                                            _ = v51
                                                                        } else {
                                                                            var v52 int = (fi11 << ((0) & 63))
                                                                            fmt.Println(fi16, v3, v3)
                                                                            var v53 uint8 = 92
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        _ = v49
                                                                    } else {
                                                                        if false {
                                                                            var v54 bool = (-1 != 0)
                                                                            v54 = (!((3.14159 / (float64(1e10) / (((0.0 - 0.1) * 0.1) / 1.0))) == 0.1))
                                                                            if (!((1e10 + -2.25) >= 0.1)) {
                                                                                switch ("" + v2) {
                                                                                case "":
                                                                                    v55 := GenS5{F0: 1}
                                                                                    _ = v55
                                                                                case "hello":
                                                                                    v3 -= int32(v10)
                                                                                default:
                                                                                    delete(v1, (-(fi11)))
                                                                                }
                                                                                v2 = "hello"
                                                                            } else {
                                                                                if false {
                                                                                    var v56 float32 = 3.14159
                                                                                    _ = v56
                                                                                } else {
                                                                                    v10 = -2147483648
                                                                                    var v57 int32 = v3
                                                                                    switch -225 {
                                                                                    case -2147483648:
                                                                                        v10 -= fi11
                                                                                    case 1:
                                                                                        var v58 uint16 = 2
                                                                                        _ = v58
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                    _ = v57
                                                                                }
                                                                                v59 := GenS10{F0: -1, F1: v10, F2: v3}
                                                                                _ = v59
                                                                            }
                                                                            v2 = "hello"
                                                                            v60 := GenS11{F0: 1}
                                                                            v61 := GenS12{F0: 2147483647, F1: (fi11 * (^(v10))), F2: (v2 + "\n\t")}
                                                                            var v62 float64 = 1.5
                                                                            _ = v54
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        _ = len(v41)
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v1[(fi16 | v10)] = false
                                                            }
                                                            if len(v41) >= 2 {
                                                                v41 = v41[1:len(v41)]
                                                            }
                                                            if (-401 <= -32768) {
                                                                var v63 uint = 4294967295
                                                                v64 := GenS13{F0: 1}
                                                                v2 = v4
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                            switch v10 {
                                                            case -383:
                                                                var v65 uint8 = 128
                                                                v10 -= -638
                                                                v2 = v2
                                                                _ = v65
                                                            case 643:
                                                                _ = len(v41)
                                                            case 57:
                                                                v4 = v2
                                                            default:
                                                                v66 := GenS14{F0: 560, F1: (v2 + (("\n\t" + "foo bar") + "\n\t")), F2: 127}
                                                                _ = v66
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    switch v4 {
                                                    case "":
                                                        switch fi16 {
                                                        case 139:
                                                            switch fi11 {
                                                            case -639:
                                                                fmt.Println(v4)
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                        case 128:
                                                            var v67 int8 = 1
                                                            _ = v67
                                                        case 32768:
                                                            v68 := GenS10{F0: -573, F1: fi11, F2: v3}
                                                            _ = v68
                                                        }
                                                    case "a":
                                                        if false {
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                switch fi11 {
                                                                case 564:
                                                                    var v69 int32 = (v3 << ((1) & 63))
                                                                    _ = v69
                                                                case 9223372036854775807:
                                                                    fmt.Println(v10, v2, v10)
                                                                case 604:
                                                                    v3 -= v3
                                                                default:
                                                                    var v70 uint16 = 200
                                                                    if (v2 < "a") {
                                                                        if true {
                                                                            v10 -= (fi11 >> ((370) & 63))
                                                                            var v71 uint = 363
                                                                            if (v70 > v70) {
                                                                                if (0 >= -1) {
                                                                                    v72 := GenS12{F0: 42, F1: fi16, F2: ("\n\t" + "hello")}
                                                                                    _ = v72
                                                                                }
                                                                                if false {
                                                                                    var v73 bool = false
                                                                                    var v74 int64 = 1
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                } else {
                                                                                    delete(v1, 42)
                                                                                    switch (v4 + ("\n\t" + v4)) {
                                                                                    case "a":
                                                                                        if len(v41) >= 2 {
                                                                                            v41 = v41[1:len(v41)]
                                                                                        }
                                                                                    case "":
                                                                                        if true {
                                                                                            if false {
                                                                                                if (!true) {
                                                                                                    fmt.Println(fi16, v10)
                                                                                                    delete(v1, 127)
                                                                                                    v10 += v10
                                                                                                }
                                                                                                if true {
                                                                                                    fmt.Println(fi11, v71, v2)
                                                                                                    var v75 uint8 = 48
                                                                                                    _ = v75
                                                                                                }
                                                                                                if false {
                                                                                                    v4 = (v4 + v2)
                                                                                                    fmt.Println(v3, v4, v2)
                                                                                                    v41 = append(v41, fi16)
                                                                                                    v76 := GenS9{F0: 2, F1: false}
                                                                                                    var v77 int32 = (v3 / ((v3) | 1))
                                                                                                    _ = v76
                                                                                                    _ = v77
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v3)
                                                                                            }
                                                                                        } else {
                                                                                            switch int(v3) {
                                                                                            case 1:
                                                                                                var v78 int = v10
                                                                                                _ = v78
                                                                                            default:
                                                                                                v79 := GenS15{F0: 42, F1: (v71 ^ (^(v71))), F2: v2}
                                                                                                _ = v79
                                                                                            }
                                                                                        }
                                                                                    case "hello":
                                                                                        v10 -= -67
                                                                                    }
                                                                                    var v80 uint = v71
                                                                                    _ = v80
                                                                                }
                                                                            } else {
                                                                                switch 28 {
                                                                                case 128:
                                                                                    v81 := GenS8{F0: (fi11 * (fi11 % (((fi11 ^ (-(fi16)))) | 1))), F1: 42}
                                                                                    _ = v81
                                                                                case 423:
                                                                                    if (-2147483648 == (fi16 ^ -390)) {
                                                                                        v82 := GenS16{F0: false}
                                                                                        var v83 int16 = 261
                                                                                        v10 = (v10 | (fi11 % (((v10 | -9223372036854775808)) | 1)))
                                                                                        fmt.Println(v71)
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                    } else {
                                                                                        v41 = append(v41, (^(fi11)))
                                                                                        v71 += v71
                                                                                        v3 = 0
                                                                                    }
                                                                                default:
                                                                                    v1[fi11] = (0 == -441)
                                                                                }
                                                                            }
                                                                            _ = v71
                                                                        } else {
                                                                            switch "a" {
                                                                            case "a":
                                                                                v41 = append(v41, (v10 * (fi16 & fi11)))
                                                                                switch v4 {
                                                                                case "":
                                                                                    _ = len(v1)
                                                                                default:
                                                                                    delete(v1, (fi16 << ((788) & 63)))
                                                                                }
                                                                                var v84 int = fi16
                                                                                v4 = v2
                                                                                _ = len(v1)
                                                                                _ = v84
                                                                            default:
                                                                                var v85 int32 = v3
                                                                                _ = v85
                                                                            }
                                                                            var v86 int = (fi11 * (fi16 + 34))
                                                                            v87 := GenS13{F0: v70}
                                                                            fmt.Println(fi11, v3, fi11)
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                        v3 = v3
                                                                        if true {
                                                                            fmt.Println(v2, fi11)
                                                                            var v88 int32 = v3
                                                                            v89 := v1[127]
                                                                            if len(v41) >= 2 {
                                                                                v41 = v41[1:len(v41)]
                                                                            }
                                                                            v4 = ((v4 + ("\n\t" + v4)) + v2)
                                                                            fmt.Println(fi11, v70, v3)
                                                                            _ = v88
                                                                            _ = v89
                                                                        } else {
                                                                            v4 = "hello"
                                                                        }
                                                                    } else {
                                                                        var v90 uint64 = 1
                                                                        v91 := GenS8{F0: fi16, F1: (v90 - v90)}
                                                                        fmt.Println(v4)
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    _ = v70
                                                                }
                                                            case 2147483647:
                                                                v92 := GenS17{F0: 1, F1: -128}
                                                                _ = v92
                                                            case 32768:
                                                                v93 := GenS18{F0: 0, F1: v3, F2: 42}
                                                                _ = v93
                                                            }
                                                        } else {
                                                            v1[(v10 - v10)] = false
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v10, v4, v10)
                                                    }
                                                case "":
                                                    if (42 == 106) {
                                                        v94 := v1[(^(fi16))]
                                                        var v95 string = v4
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v96 := GenS2{F0: 4294967295, F1: 742}
                                                        _ = v96
                                                    }
                                                    var v97 uint64 = 485
                                                    _ = v97
                                                default:
                                                    if false {
                                                        v2 = ("" + ("hello" + v4))
                                                    } else {
                                                        switch (v10 - fi16) {
                                                        case 9223372036854775807:
                                                            v98 := GenS18{F0: 18446744073709551615, F1: v3, F2: 65535}
                                                            _ = v98
                                                        case -9223372036854775808:
                                                            if (fi11 < 1) {
                                                                var v99 float64 = -2.25
                                                                fmt.Println(fi16)
                                                                var v100 uint64 = 42
                                                                var v101 bool = (((((-2.25 + -2.25) * v99) - 3.14159) + 3.14159) <= v99)
                                                                _ = v100
                                                                _ = v101
                                                                _ = v99
                                                            } else {
                                                                v4 = v4
                                                                if true {
                                                                    var v102 int64 = 0
                                                                    if false {
                                                                        if true {
                                                                            var v103 float32 = ((0.0 + 1e10) - (0.0 * -2.25))
                                                                            v104 := GenS17{F0: 4294967295, F1: -128}
                                                                            _ = v103
                                                                            _ = v104
                                                                        } else {
                                                                            v4 = "\n\t"
                                                                        }
                                                                    }
                                                                    _ = v102
                                                                } else {
                                                                    var v105 uint64 = 748
                                                                    delete(v1, 32767)
                                                                    _ = v105
                                                                }
                                                                if false {
                                                                    v2 = "\n\t"
                                                                    var v106 uint16 = 65535
                                                                    _ = v106
                                                                } else {
                                                                    var v107 int64 = 42
                                                                    _ = v107
                                                                }
                                                                v10 = (fi11 | 127)
                                                                v108 := v1[(fi11 >> ((1) & 63))]
                                                                _ = v108
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v41
                                            } else {
                                                fmt.Println(v2, fi16, v4)
                                            }
                                            _ = v38
                                            _ = v39
                                        }
                                    default:
                                        v2 = (v2 + v4)
                                    }
                                    fmt.Println(fi11)
                                } else {
                                    var v109 uint8 = 176
                                    v2 = v4
                                    v110 := GenS12{F0: 65535, F1: fi16, F2: "foo bar"}
                                    _ = v109
                                    _ = v110
                                }
                                if false {
                                    var v111 int8 = 100
                                    _ = v111
                                } else {
                                    v3 = (v3 % (((v3 & v3)) | 1))
                                    v2 = "hello"
                                }
                                var v112 uint8 = 234
                                _ = v112
                                _ = v35
                            case -2147483648:
                                v113 := v1[(v10 & 42)]
                                _ = v113
                            default:
                                v114 := GenS9{F0: 4294967295, F1: false}
                                _ = v114
                            }
                        } else {
                            if (1000 > 32767) {
                                if (278 != 2147483648) {
                                    _ = len(v1)
                                    if (v3 > v3) {
                                        var v115 uint = 255
                                        v10 = 331
                                        _ = v115
                                    }
                                }
                                if false {
                                    var v116 bool = (!(-21 >= 100))
                                    fmt.Println(v2, v10, v116)
                                    _ = v116
                                }
                            }
                        }
                    }
                    _ = fi16
                }
                switch (((((v2 + ((v4 + v4) + v4)) + v4) + "foo bar") + "a") + v2) {
                case "\n\t":
                    v3 = v3
                case "a":
                    v117 := []int{v10}
                    _ = v117
                }
                for fi118 := 0; fi118 < 3; fi118++ {
                    var v119 bool = false
                    v119 = false
                    _ = fi118
                    _ = v119
                }
            }
            var v120 int32 = (v3 + (v3 * v3))
            _ = fi11
            _ = v120
        }
    } else {
        v121 := GenS19{F0: v10, F1: 0}
        fmt.Println(v10)
        v122 := GenS20{F0: 161}
        _ = v121
        _ = v122
    }
    v123 := GenS11{F0: 0}
    v4 = "hello"
    fmt.Println(v3, v10)
    if (false || true) {
        switch (v10 ^ v10) {
        case -945:
            var v124 int8 = -1
            _ = v124
        case 32767:
            v4 = "hello"
        default:
            v125 := []int{(v10 & v10), (v10 / (((v10 * v10)) | 1))}
            _ = v125
        }
        if false {
            v1[v10] = true
        }
        v2 = v4
    } else {
        if true {
            fmt.Println(v10, v10)
        } else {
            var v126 float64 = 3.14159
            fmt.Println(v4, v3)
            _ = v126
        }
        v3 += -1
    }
    v10 = v10
    v1[(v10 - 32768)] = false
    switch 0 {
    case 32767:
        for fi127 := 0; fi127 < 3; fi127++ {
            switch ("" + (v2 + v2)) {
            case "a":
                fmt.Println(fi127, fi127, fi127)
            default:
                v128 := GenS16{F0: false}
                _ = v128
            }
            v1[128] = true
            var v129 int16 = -1000
            _ = fi127
            _ = v129
        }
    case 9223372036854775807:
        v3 -= v3
    case 308:
        for fi130 := 0; fi130 < 7; fi130++ {
            if false {
                var v131 uint64 = 0
                _ = v131
            }
            fmt.Println(v3, v4, v2)
            _ = fi130
        }
    default:
        var v132 bool = (!false)
        _ = v132
    }
    for fi133 := 0; fi133 < 11; fi133++ {
        for fi134 := 0; fi134 < 3; fi134++ {
            var v135 uint32 = 2147483648
            fmt.Println(v3, fi133)
            _ = fi134
            _ = v135
        }
        v136 := GenS21{F0: v4, F1: 312}
        switch v4 {
        case "a":
            var v137 string = "a"
            _ = v137
        case "\n\t":
            var v138 float64 = ((0.1 * (0.0 + 0.1)) - -2.25)
            _ = v138
        default:
            switch (fi133 / (((-(v10))) | 1)) {
            case 135:
                for fi139 := 0; fi139 < 4; fi139++ {
                    v140 := []int{1}
                    _ = fi139
                    _ = v140
                }
            case -1:
                for fi141 := 0; fi141 < 2; fi141++ {
                    v142 := []int{}
                    v143 := GenS2{F0: 566, F1: fi133}
                    v144 := GenS3{F0: 116}
                    _ = fi141
                    _ = v142
                    _ = v143
                    _ = v144
                }
            }
        }
        _ = fi133
        _ = v136
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
