package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int32 }
func genFn0(p0 uint16) uint32 {
    var v5 int16 = -638
    switch -2147483648 {
    case 127:
        fmt.Println(v5, v5)
    case -9223372036854775808:
        var v6 string = "foo bar"
        _ = v6
    case -1:
        var v7 uint64 = 1
        _ = v7
    default:
        var v8 uint = 42
        _ = v8
    }
    v9 := []int{}
    if true {
        p0 -= (p0 << ((0) & 63))
    } else {
        v10 := GenS1{F0: p0}
        v5 = (v5 % (((v5 + (v5 ^ 739))) | 1))
        _ = v10
    }
    p0 += (^(p0))
    _ = p0
    _ = v5
    _ = v9
    return 65535
}
func genFn1(p0 uint32, p1 uint16) bool {
    v11 := []int{-49, -905}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v12 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (!true)
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 bool = (1 != -1)
    var v13 int8 = -70
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = 0
        _ = fi14
        _ = v15
    }
    if v4 {
        fmt.Println(v4, v1)
        v16 := GenS2{F0: 18446744073709551615}
        _ = v16
    } else {
        v4 = true
        for fi17 := 0; fi17 < 11; fi17++ {
            if ((!true) && false) {
                switch "\n\t" {
                case "foo bar":
                    for fi18 := 0; fi18 < 19; fi18++ {
                        if (fi17 < fi18) {
                            fmt.Println(fi18, fi17, fi17)
                            var v19 uint = 576
                            _ = v19
                        }
                        _ = fi18
                    }
                default:
                    v1 -= 3.14159
                }
                delete(v3, fi17)
            } else {
                switch fi17 {
                case 1:
                    switch "" {
                    case "":
                        var v20 int32 = -2147483648
                        _ = v20
                    case "foo bar":
                        if false {
                            v1 -= (v1 + v1)
                            var v21 bool = (266 > 907)
                            if (fi17 >= (fi17 / ((4294967295) | 1))) {
                                var v22 int = fi17
                                var v23 bool = v4
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                    case "a":
                        v13 -= -18
                    }
                case 32768:
                    v24 := GenS3{F0: 42}
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 7; fi25++ {
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            if (v4 || (v4 && v4)) {
                                v26 := GenS1{F0: 32767}
                                _ = v26
                            }
                        case -2147483648:
                            v27 := GenS1{F0: 902}
                            _ = v27
                        }
                        var v28 int = fi17
                        v29 := GenS2{F0: 18446744073709551615}
                        _ = fi25
                        _ = v28
                        _ = v29
                    }
                }
                if (1 == 134) {
                    if false {
                        if ((v4 || v4) && ((!v4) && ((0.0 > (1.5 * 0.0)) || (42 != 0)))) {
                            v13 = (v13 >> ((255) & 63))
                        } else {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                fmt.Println(v4, fi17, v1)
                                if v4 {
                                    v31 := GenS1{F0: 114}
                                    fmt.Println(v13, v4)
                                    _ = len(v2)
                                    _ = v31
                                }
                                v32 := GenS4{F0: 468}
                                switch "a" {
                                case "hello":
                                    if ((v1 - 0.1) == v1) {
                                        if (!(!(1e10 > ((0.0 - -2.25) / (1.5 * 1e10))))) {
                                            switch "a" {
                                            case "":
                                                v1 = 0.1
                                            case "foo bar":
                                                var v33 int8 = (v13 >> ((628) & 63))
                                                _ = v33
                                            default:
                                                if v4 {
                                                    var v34 uint8 = 2
                                                    _ = v34
                                                } else {
                                                    v13 -= v13
                                                    var v35 uint64 = 42
                                                    _ = v35
                                                }
                                            }
                                            if (!false) {
                                                v36 := GenS2{F0: 839}
                                                v37 := GenS3{F0: 334}
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v38 uint32 = 177
                                                _ = v38
                                            }
                                        }
                                    }
                                case "a":
                                    v39 := v3[(fi30 + fi30)]
                                    v4 = (genFn0(1) >= 0)
                                    v3[fi17] = fi17
                                    _ = len(v2)
                                    fmt.Println(v13, fi30)
                                    _ = v39
                                case "foo bar":
                                    switch fi30 {
                                    case 42:
                                        var v40 int16 = 1000
                                        _ = v40
                                    case 407:
                                        v2 = append(v2, 9223372036854775807)
                                    }
                                default:
                                    delete(v3, (fi30 & (-(fi30))))
                                }
                                var v41 float64 = float64((v1 * (v1 - -2.25)))
                                _ = fi30
                                _ = v32
                                _ = v41
                            }
                            v2 = append(v2, 2147483647)
                            fmt.Println(v4, v1, v1)
                        }
                    } else {
                        v42 := GenS1{F0: 65535}
                        v43 := GenS3{F0: 0}
                        fmt.Println(v1, v4, v1)
                        _ = v42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 16; fi44++ {
                        _ = len(v2)
                        _ = fi44
                    }
                } else {
                    v45 := GenS5{F0: 1, F1: (fi17 % ((887) | 1))}
                    v1 = -2.25
                    _ = v45
                }
                var v46 int8 = v13
                _ = v46
            }
            for fi47 := 0; fi47 < 16; fi47++ {
                switch -9223372036854775808 {
                case 4294967295:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v1, fi47)
                    delete(v3, fi47)
                case 127:
                    switch "" {
                    case "hello":
                        fmt.Println(v4, v1)
                    }
                case 22:
                    v48 := GenS1{F0: 0}
                    _ = v48
                default:
                    if ((v13 / ((v13) | 1)) == v13) {
                        v4 = (v4 || (18446744073709551615 >= 2))
                    } else {
                        if v4 {
                            fmt.Println(fi17, v13)
                        }
                        v49 := GenS4{F0: 255}
                        _ = v49
                    }
                }
                _ = fi47
            }
            fmt.Println(fi17, v13, fi17)
            _ = fi17
        }
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 bool = (!(!(42 != 1)))
            _ = fi50
            _ = v51
        }
    }
    if (!v4) {
        v1 += 1e10
        var v52 int = 124
        _ = v52
    }
    var v53 int64 = 1
    if true {
        if true {
            v54 := GenS2{F0: 732}
            v53 = (v53 ^ (v53 / ((0) | 1)))
            _ = v54
        } else {
            if v4 {
                for fi55 := 0; fi55 < 6; fi55++ {
                    for fi56 := 0; fi56 < 12; fi56++ {
                        v57 := GenS4{F0: 136}
                        _ = fi56
                        _ = v57
                    }
                    _ = fi55
                }
                v13 -= (v13 * 31)
                var v58 uint64 = 0
                _ = v58
            } else {
                switch (("\n\t" + (("" + ("foo bar" + "\n\t")) + "hello")) + "") {
                case "":
                    var v59 uint16 = 774
                    _ = v59
                case "\n\t":
                    var v60 bool = v4
                    _ = v60
                case "a":
                    for fi61 := 0; fi61 < 16; fi61++ {
                        if v4 {
                            var v62 float64 = (-2.25 + 0.0)
                            switch "hello" {
                            case "":
                                var v63 bool = (!(32767 <= 1))
                                _ = v63
                            case "\n\t":
                                if (v4 && ((v1 + v1) != (float32((((1.5 - v62) + v62) + v62)) - v1))) {
                                    if (551 >= 2147483647) {
                                        v64 := v3[42]
                                        v64 = 32768
                                        var v65 uint64 = 18446744073709551615
                                        _ = v64
                                        _ = v65
                                    }
                                } else {
                                    v53 -= (-(v53))
                                    v53 -= (^(v53))
                                }
                            case "a":
                                switch fi61 {
                                case 32767:
                                    v66 := GenS6{F0: (-670 != 32767), F1: 255}
                                    _ = v66
                                case 9223372036854775807:
                                    var v67 uint16 = 0
                                    _ = v67
                                default:
                                    var v68 string = "hello"
                                    _ = v68
                                }
                            default:
                                var v69 uint16 = 551
                                _ = v69
                            }
                            _ = v62
                        } else {
                            _ = len(v3)
                            var v70 int64 = (v53 >> ((255) & 63))
                            var v71 float64 = -2.25
                            _ = v70
                            _ = v71
                        }
                        switch ("a" + "foo bar") {
                        case "hello":
                            for fi72 := 0; fi72 < 14; fi72++ {
                                fmt.Println(v4)
                                fmt.Println(v53)
                                var v73 bool = v4
                                _ = fi72
                                _ = v73
                            }
                        }
                        _ = fi61
                    }
                default:
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v2)
                    }
                }
                switch -756 {
                case 105:
                    var v74 uint32 = 2147483647
                    _ = v74
                case 127:
                    v3[42] = 4294967295
                }
            }
        }
        if v4 {
            if true {
                fmt.Println(v4, v4, v13)
                v75 := GenS4{F0: 751}
                _ = v75
            }
            v1 = v1
            _ = len(v2)
        } else {
            for fi76 := 0; fi76 < 16; fi76++ {
                var v77 int32 = -2147483648
                _ = fi76
                _ = v77
            }
            for fi78 := 0; fi78 < 12; fi78++ {
                fmt.Println(fi78)
                if ((v4 && v4) || (!(false && ((-(v13)) > 0)))) {
                    v79 := v3[9223372036854775807]
                    for fi80 := 0; fi80 < 17; fi80++ {
                        v81 := GenS7{F0: 2, F1: v13, F2: 2}
                        v79 += (fi78 - 9223372036854775807)
                        _ = fi80
                        _ = v81
                    }
                    for fi82 := 0; fi82 < 15; fi82++ {
                        fmt.Println(v53)
                        _ = fi82
                    }
                    _ = v79
                } else {
                    v53 -= (v53 | (v53 & v53))
                    for fi83 := 0; fi83 < 11; fi83++ {
                        var v84 uint32 = 2
                        v85 := GenS2{F0: 4294967295}
                        var v86 int16 = int16(fi78)
                        _ = fi83
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                }
                _ = fi78
            }
            var v87 uint32 = 0
            _ = v87
        }
    } else {
        for fi88 := 0; fi88 < 17; fi88++ {
            var v89 uint = 655
            _ = fi88
            _ = v89
        }
        switch (("foo bar" + "\n\t") + "hello") {
        case "\n\t":
            v90 := GenS8{F0: (2 <= 42), F1: 2}
            _ = v90
        case "a":
            if v4 {
                if false {
                    v4 = (!v4)
                    fmt.Println(v1, v53)
                } else {
                    switch 32768 {
                    case 42:
                        v4 = (585 <= 491)
                    default:
                        var v91 uint16 = 757
                        _ = v91
                    }
                }
                if true {
                    if v4 {
                        fmt.Println(v53, v13, v4)
                        var v92 uint64 = 0
                        if false {
                            if ((v92 > (v92 - v92)) && v4) {
                                v4 = v4
                                if ((v53 + (v53 | (v53 - v53))) > -9223372036854775808) {
                                    v93 := GenS9{F0: -2147483648}
                                    v94 := GenS2{F0: (v92 / ((v92) | 1))}
                                    var v95 int64 = -9223372036854775808
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                }
                            }
                            v1 += float32((0.1 - 0.1))
                        } else {
                            for fi96 := 0; fi96 < 15; fi96++ {
                                var v97 uint64 = (v92 + v92)
                                v3[fi96] = (fi96 + -2147483648)
                                v13 -= (v13 - (v13 + 92))
                                _ = fi96
                                _ = v97
                            }
                        }
                        _ = v92
                    }
                    v98 := GenS4{F0: 255}
                    _ = v98
                }
                for fi99 := 0; fi99 < 20; fi99++ {
                    if v4 {
                        for fi100 := 0; fi100 < 8; fi100++ {
                            v101 := GenS4{F0: 65535}
                            if v4 {
                                switch (fi99 << ((708) & 63)) {
                                case 127:
                                    v53 = (v53 + v53)
                                case 32768:
                                    _ = len(v3)
                                default:
                                    v102 := GenS9{F0: -2147483648}
                                    _ = v102
                                }
                            }
                            var v103 uint32 = 338
                            _ = fi100
                            _ = v101
                            _ = v103
                        }
                        if (0 == 18446744073709551615) {
                            if len(v2) > 0 {
                                v2[0] = fi99
                            }
                            v13 -= -1
                        }
                        if len(v2) > 0 {
                            v2[0] = (fi99 & (^(fi99)))
                        }
                    }
                    v3[fi99] = fi99
                    v104 := GenS1{F0: 2}
                    _ = fi99
                    _ = v104
                }
            } else {
                v1 -= (v1 / (-2.25 + ((v1 / (float32(0.1) - ((((float32((1.5 + 1.5)) * 3.14159) + v1) + ((3.14159 * v1) + -2.25)) - v1))) / 1.5)))
            }
        case "foo bar":
            var v105 uint16 = 65535
            _ = v105
        default:
            for fi106 := 0; fi106 < 3; fi106++ {
                v107 := v3[(fi106 * fi106)]
                v108 := GenS4{F0: 255}
                v109 := GenS7{F0: 100, F1: 100, F2: 65535}
                _ = fi106
                _ = v107
                _ = v108
                _ = v109
            }
        }
    }
    if true {
        var v110 float64 = 0.0
        _ = v110
    } else {
        if (42 == 255) {
            fmt.Println(v4, v53)
        }
        if (!(1000 > -92)) {
            v2 = append(v2, 32768)
        }
    }
    v2 = append(v2, 664)
    var v111 uint8 = 128
    fmt.Println(v53)
    v111 = v111
    var v112 uint = 640
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v111
    _ = v112
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v53
}
