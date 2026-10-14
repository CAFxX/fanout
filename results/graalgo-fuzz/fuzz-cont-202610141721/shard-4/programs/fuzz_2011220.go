package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint8; F1 int32 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        if (1 <= 1000) {
            v5 := GenS1{F0: ("\n\t" + ""), F1: "a", F2: 0}
            _ = v5
        } else {
            var v6 uint32 = 0
            _ = v6
        }
        var v7 float64 = 1e10
        if (1 > 65535) {
            if (65535 == 1) {
                v8 := []int{(fi4 | (fi4 & 128))}
                v9 := GenS2{F0: -1, F1: true, F2: 32767}
                if true {
                    fmt.Println(v7)
                }
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint8 = 127
                var v12 int = fi10
                v7 += 0.0
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v7 += ((float64(1.5) - v7) - float64(1e10))
            v7 += ((3.14159 - v7) + (1.5 / (((v7 + (1.5 - ((1e10 * (0.1 - v7)) + v7))) / ((float64((0.1 / -2.25)) + -2.25) + (v7 * v7))) + v7)))
            if (0 >= 1) {
                var v13 int = 1
                v14 := []int{1}
                fmt.Println(v7, v13)
                _ = v13
                _ = v14
            } else {
                fmt.Println(fi4, fi4)
                var v15 int64 = 511
                _ = v15
            }
        }
        _ = fi4
        _ = v7
    }
    switch "hello" {
    case "\n\t":
        v16 := []int{42}
        _ = v16
    default:
        var v17 bool = false
        _ = v17
    }
    var v18 bool = true
    if (!v18) {
        v19 := GenS1{F0: "", F1: (("\n\t" + "foo bar") + "a"), F2: -1}
        if true {
            v20 := []int{-447}
            _ = v20
        } else {
            v21 := GenS3{F0: 4294967295}
            v22 := GenS3{F0: 557}
            if (false && v18) {
                v18 = (18446744073709551615 < 42)
            }
            _ = v21
            _ = v22
        }
        var v23 float32 = float32((float64(3.14159) + (1.5 * (0.0 + (0.0 - 0.1)))))
        _ = v19
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            switch "a" {
            case "\n\t":
                for fi25 := 0; fi25 < 6; fi25++ {
                    v18 = v18
                    if false {
                        var v26 bool = v18
                        _ = v26
                    } else {
                        fmt.Println(v18, fi25, v18)
                        var v27 int = fi24
                        _ = v27
                    }
                    _ = fi25
                }
            case "hello":
                switch -518 {
                case -114:
                    v28 := []int{}
                    _ = v28
                }
            default:
                v29 := GenS2{F0: -9223372036854775808, F1: (0.0 <= float64(((float32((((0.1 * 1e10) * -2.25) / float64(float32((1.5 - 1e10))))) + 3.14159) - (0.1 / 3.14159)))), F2: 42}
                _ = v29
            }
            v30 := GenS4{F0: 127}
            for fi31 := 0; fi31 < 3; fi31++ {
                fmt.Println(v18)
                if (18446744073709551615 > 1) {
                    var v32 uint32 = 42
                    _ = v32
                }
                switch "a" {
                case "a":
                    switch (("\n\t" + "") + "a") {
                    case "\n\t":
                        switch (fi24 * 0) {
                        case 42:
                            v18 = v18
                        default:
                            v18 = v18
                            v18 = v18
                            if (2 != 128) {
                                if v18 {
                                    v33 := GenS4{F0: 100}
                                    var v34 int32 = 450
                                    var v35 int8 = -128
                                    fmt.Println(fi24)
                                    v35 = (v35 + v35)
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    fmt.Println(fi31, fi24, fi31)
                                    var v36 int64 = 1
                                    v37 := []int{}
                                    _ = v36
                                    _ = v37
                                }
                                fmt.Println(v18, fi24, fi31)
                            }
                        }
                    case "":
                        if ((!((-2.25 - 0.1) != (float32(-2.25) - (((((((1.5 - (0.0 + 1.5)) * 1.5) / (0.0 - (3.14159 / 1.0))) * -2.25) / 1.5) * 0.1) + (0.1 / 1.0))))) && v18) {
                            v38 := GenS3{F0: 621}
                            _ = v38
                        }
                    default:
                        v39 := []int{(fi24 << ((255) & 63))}
                        _ = v39
                    }
                case "hello":
                    switch (fi31 + fi31) {
                    case -2147483648:
                        v40 := []int{32767, fi24, fi31, fi31}
                        _ = v40
                    default:
                        var v41 int8 = 1
                        _ = v41
                    }
                case "foo bar":
                    if (!(137 < 1)) {
                        v42 := []int{(fi31 % (((fi24 & (fi24 - fi31))) | 1)), (fi31 % ((-516) | 1))}
                        _ = v42
                    }
                }
                _ = fi31
            }
            _ = fi24
            _ = v30
        }
        fmt.Println(v18)
        var v43 float64 = ((0.0 * ((((0.0 * float64(1.5)) * 1e10) * 0.0) + 0.0)) / 3.14159)
        _ = v43
    }
    _ = v18
    return 324
}
func genFn1(p0 uint) string {
    if true {
        v44 := []int{242}
        v44 = append(v44, -1)
        _ = v44
    } else {
        var v45 int32 = 0
        fmt.Println(p0)
        _ = v45
    }
    var v46 float64 = (-2.25 - 0.1)
    _ = p0
    _ = v46
    return ""
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 374
    var v3 uint64 = 1
    if true {
        v47 := []int{9223372036854775807, 2147483647, 867}
        _ = v47
    } else {
        v48 := GenS2{F0: -9223372036854775808, F1: (!false), F2: 2}
        for fi49 := 0; fi49 < 3; fi49++ {
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := []int{}
                var v52 int32 = 0
                v1 = v1
                _ = fi50
                _ = v51
                _ = v52
            }
            var v53 float32 = -2.25
            _ = fi49
            _ = v53
        }
        for fi54 := 0; fi54 < 3; fi54++ {
            fmt.Println(v1, fi54)
            _ = fi54
        }
        _ = v48
    }
    switch 2147483647 {
    case 128:
        if false {
            for fi55 := 0; fi55 < 18; fi55++ {
                if (int32(v1) <= 718) {
                    for fi56 := 0; fi56 < 13; fi56++ {
                        v57 := []int{(fi55 & fi56), 2147483647}
                        _ = fi56
                        _ = v57
                    }
                } else {
                    v3 -= v3
                    var v58 uint64 = v3
                    _ = v58
                }
                for fi59 := 0; fi59 < 11; fi59++ {
                    var v60 string = ""
                    v61 := []int{fi59, (fi55 + (fi59 - fi55)), (fi55 | 40)}
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                for fi62 := 0; fi62 < 16; fi62++ {
                    if true {
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    if true {
                                        v1 += (v1 / (((^(v1))) | 1))
                                    } else {
                                        switch 9223372036854775807 {
                                        case 0:
                                            fmt.Println(v1, v1, v2)
                                        case 1:
                                            v63 := []int{(fi62 << (((v2 << ((39) & 63))) & 63)), 0}
                                            v64 := GenS5{F0: 6}
                                            var v65 uint16 = v1
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        case -9223372036854775808:
                                            if (("" + ("hello" + "foo bar")) > "hello") {
                                                if true {
                                                    switch "hello" {
                                                    case "a":
                                                        v66 := []int{-2147483648, 9223372036854775807, (fi62 + 2147483647)}
                                                        _ = v66
                                                    case "hello":
                                                        v67 := GenS3{F0: (^(v3))}
                                                        _ = v67
                                                    }
                                                }
                                                v68 := GenS1{F0: "", F1: ("foo bar" + "hello"), F2: -9}
                                                _ = v68
                                            } else {
                                                v69 := GenS5{F0: 35}
                                                v3 += v3
                                                _ = v69
                                            }
                                        }
                                        v2 -= v2
                                    }
                                    var v70 float64 = 1e10
                                    _ = v70
                                } else {
                                    var v71 uint16 = v1
                                    _ = v71
                                }
                            case "":
                                var v72 int64 = 0
                                _ = v72
                            case "foo bar":
                                var v73 uint8 = 127
                                _ = v73
                            }
                            v3 += 18446744073709551615
                        } else {
                            var v74 int16 = 440
                            _ = v74
                        }
                    } else {
                        v75 := []int{357}
                        _ = v75
                    }
                    if true {
                        v76 := []int{fi55, (fi55 * fi62), fi62, fi55}
                        var v77 int = 0
                        _ = v76
                        _ = v77
                    } else {
                        if false {
                            v78 := GenS6{F0: 92, F1: 2147483647}
                            v79 := []int{}
                            _ = v78
                            _ = v79
                        }
                        v3 += 657
                        var v80 uint8 = uint8(v1)
                        _ = v80
                    }
                    v81 := GenS6{F0: (fi55 | (fi55 % ((357) | 1))), F1: 2147483647}
                    _ = fi62
                    _ = v81
                }
                _ = fi55
            }
        }
    case 739:
        var v82 int32 = -2147483648
        _ = v82
    case 4294967295:
        fmt.Println(v2)
    default:
        v83 := GenS3{F0: (v3 | 971)}
        _ = v83
    }
    if (true || true) {
        v84 := GenS7{F0: 42, F1: 1}
        for fi85 := 0; fi85 < 9; fi85++ {
            v2 -= 42
            v1 -= v1
            _ = fi85
        }
        var v86 int16 = -1
        _ = v84
        _ = v86
    } else {
        v87 := GenS4{F0: 100}
        v2 -= v2
        switch 4294967295 {
        case 32767:
            if true {
                v3 = v3
            }
        }
        _ = v87
    }
    v88 := GenS5{F0: 60}
    fmt.Println(v1, v3)
    var v89 float32 = 0.0
    if true {
        v90 := []int{-1, -186, 42, -457}
        _ = len(v90)
        _ = v90
    }
    v3 = 42
    switch -1 {
    case -483:
        for fi91 := 0; fi91 < 13; fi91++ {
            v92 := GenS2{F0: int64(v3), F1: false, F2: -1}
            if (true || false) {
                var v93 float64 = float64(float32((0.0 + (0.1 * 1e10))))
                _ = v93
            } else {
                v3 += 4294967295
                var v94 uint = v2
                if true {
                    v1 -= uint16(v2)
                    v95 := []int{-1, -500}
                    v94 -= v2
                    _ = v95
                }
                _ = v94
            }
            _ = fi91
            _ = v92
        }
    case 628:
        v2 = (v2 % (((^(v2))) | 1))
    }
    for fi96 := 0; fi96 < 13; fi96++ {
        var v97 float32 = (3.14159 * (v89 / v89))
        for fi98 := 0; fi98 < 20; fi98++ {
            var v99 uint = (v2 >> ((v2) & 63))
            _ = fi98
            _ = v99
        }
        var v100 int = fi96
        _ = fi96
        _ = v100
        _ = v97
    }
    var v101 float32 = v89
    var v102 string = ("\n\t" + "")
    v103 := []int{632, 587, 9223372036854775807}
    var v104 int8 = 19
    for fi105 := 0; fi105 < 11; fi105++ {
        var v106 uint = v2
        var v107 uint64 = (v3 / (((v3 * v3)) | 1))
        for fi108 := 0; fi108 < 1; fi108++ {
            v106 = v2
            var v109 int = 4294967295
            switch (fi105 >> ((v2) & 63)) {
            case -86:
                switch (fi105 ^ fi105) {
                case -404:
                    var v110 float32 = (float32(0.0) + (3.14159 * -2.25))
                    _ = v110
                case 32768:
                    v101 = float32(1.5)
                case 0:
                    fmt.Println(v3, v107, v3)
                }
            case 127:
                v104 = (-(v104))
                var v111 uint8 = 255
                _ = v111
            case 4294967295:
                var v112 uint = (v2 * 1)
                _ = v112
            default:
                v101 = 0.1
            }
            _ = fi108
            _ = v109
        }
        _ = fi105
        _ = v106
        _ = v107
    }
    var v113 float64 = float64(0.0)
    v114 := GenS2{F0: 42, F1: true, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v104:", v104)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v88
    _ = v89
}
