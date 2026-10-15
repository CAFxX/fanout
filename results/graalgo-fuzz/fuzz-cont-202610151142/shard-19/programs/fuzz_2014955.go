package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 int8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 uint64; F1 int; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 uint; F1 uint16; F2 int8 }
type GenS13 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0(p0 int64) int {
    v4 := GenS1{F0: 65535, F1: 1, F2: 4294967295}
    p0 += p0
    _ = p0
    _ = v4
    return 668
}
func genFn1(p0 int) uint16 {
    p0 += 128
    if true {
        v5 := GenS2{F0: 128}
        _ = v5
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 42
    var v3 float32 = 1.5
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 19; fi7++ {
        v3 = v3
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi8 {
            case 742:
                v9 := GenS3{F0: 65535, F1: (false && false), F2: 47}
                _ = v9
            }
            switch "" {
            case "foo bar":
                fmt.Println(v6)
            default:
                fmt.Println(fi7, fi8, v6)
            }
            var v10 int16 = -562
            _ = fi8
            _ = v10
        }
        switch (-(fi7)) {
        case -34:
            switch "foo bar" {
            case "":
                if (411 <= 4294967295) {
                    fmt.Println(v6, v3, fi7)
                    for fi11 := 0; fi11 < 11; fi11++ {
                        v6 = v6
                        if true {
                            v12 := GenS1{F0: 65535, F1: 1, F2: 127}
                            _ = v12
                        }
                        _ = fi11
                    }
                    fmt.Println(v6)
                } else {
                    fmt.Println(v6)
                    v6 = (^(v6))
                    v13 := v1[2]
                    _ = v13
                }
            case "foo bar":
                var v14 uint16 = 213
                _ = v14
            case "a":
                fmt.Println(v2)
            default:
                for fi15 := 0; fi15 < 10; fi15++ {
                    v6 = 934
                    fmt.Println(fi7, fi7, fi15)
                    var v16 uint16 = 1
                    _ = fi15
                    _ = v16
                }
            }
        case 1:
            v3 += float32((float64(float32(0.0)) * 3.14159))
        case 4294967295:
            if (31 < 42) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    if (!true) {
                        var v18 int16 = 32767
                        v6 -= (^(v6))
                        fmt.Println(v6, v3)
                        if (uint32(v6) < 0) {
                            switch (^(fi7)) {
                            case -389:
                                v18 = v18
                            case 0:
                                var v19 float32 = (0.1 / (v3 * 0.0))
                                _ = v19
                            }
                            var v20 int16 = -1
                            if false {
                                v21 := GenS3{F0: 907, F1: true, F2: 127}
                                _ = v21
                            } else {
                                fmt.Println(v18, fi7)
                                if false {
                                    if ((0.0 >= 0.1) || (v2 <= (-(v2)))) {
                                        v22 := GenS4{F0: (v2 / (((v2 | int32(v20))) | 1)), F1: false}
                                        var v23 int64 = -1
                                        v24 := GenS1{F0: 65535, F1: 289, F2: (fi17 * fi17)}
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                } else {
                                    if (((-128 < -128) || true) || (-745 <= 9223372036854775807)) {
                                        v20 += -1
                                        switch "a" {
                                        case "":
                                            var v25 int16 = (v18 >> (((v6 >> ((v6) & 63))) & 63))
                                            _ = v25
                                        case "foo bar":
                                            var v26 bool = (0.0 <= 0.1)
                                            _ = v26
                                        case "hello":
                                            _ = len(v1)
                                        default:
                                            if (false || (!true)) {
                                                fmt.Println(v6, v18)
                                            }
                                        }
                                        var v27 uint8 = 0
                                        _ = v27
                                    }
                                    v20 += v20
                                    fmt.Println(v18)
                                }
                            }
                            _ = v20
                        }
                        var v28 bool = false
                        if false {
                            v29 := GenS5{F0: (v18 / ((v18) | 1)), F1: 234}
                            v30 := GenS5{F0: (v18 / ((1000) | 1)), F1: 1}
                            var v31 string = (("foo bar" + "a") + "a")
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            v18 = 680
                            _ = len(v1)
                            v6 -= v6
                        }
                        _ = v18
                        _ = v28
                    }
                    if true {
                        v6 = v6
                    }
                    _ = fi17
                }
                if (2 != 2) {
                    v3 += v3
                }
                if (false || (false || true)) {
                    var v32 float64 = 3.14159
                    _ = v32
                }
            }
        default:
            var v33 string = ("\n\t" + "hello")
            _ = v33
        }
        _ = fi7
    }
    v34 := GenS5{F0: -220, F1: 18446744073709551615}
    var v35 string = ("" + ("" + ("hello" + ("\n\t" + "a"))))
    for fi36 := 0; fi36 < 2; fi36++ {
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 uint8 = 90
            switch v35 {
            case "a":
                v39 := GenS5{F0: 191, F1: 626}
                fmt.Println(fi36)
                _ = v39
            case "":
                v38 += (v38 / ((v38) | 1))
            case "foo bar":
                if true {
                    var v40 uint64 = 912
                    var v41 int = fi36
                    _ = len(v1)
                    _ = v40
                    _ = v41
                } else {
                    v42 := GenS6{F0: fi37}
                    _ = v42
                }
            default:
                fmt.Println(v35, fi36, fi37)
            }
            _ = fi37
            _ = v38
        }
        v43 := GenS7{F0: fi36}
        for fi44 := 0; fi44 < 16; fi44++ {
            if true {
                v2 = v2
                if (!false) {
                    var v45 float32 = v3
                    if (!false) {
                        var v46 int32 = (v2 | v2)
                        if (v6 >= (v6 ^ (v6 | (v6 % ((v6) | 1))))) {
                            if true {
                                switch v35 {
                                case "hello":
                                    v47 := GenS1{F0: 65535, F1: 4294967295, F2: -353}
                                    if false {
                                        v48 := []int{-9223372036854775808}
                                        _ = v48
                                    }
                                    if ((0.0 / 1.0) > (-2.25 * 3.14159)) {
                                        var v49 int8 = 100
                                        if (false || false) {
                                            var v50 uint8 = 134
                                            _ = v50
                                        } else {
                                            if (!false) {
                                                var v51 string = "a"
                                                fmt.Println(v35, v49, v35)
                                                if (255 >= 255) {
                                                    v52 := GenS3{F0: 2, F1: false, F2: v49}
                                                    fmt.Println(fi36, v45)
                                                    _ = v52
                                                }
                                                _ = v51
                                            }
                                            v1[2] = uint32(v6)
                                        }
                                        if true {
                                            v49 = v49
                                            var v53 bool = false
                                            _ = v53
                                        }
                                        var v54 float32 = ((float32((((0.0 - (-2.25 + 1e10)) * float64((v3 * (v3 + float32((0.0 * 1.5)))))) + 1e10)) - 3.14159) / ((float32(float64(v45)) / v3) - 0.0))
                                        _ = v49
                                        _ = v54
                                    }
                                    if false {
                                        v55 := []int{int(v6), fi44}
                                        if len(v55) > 0 {
                                            v55[0] = (fi36 % (((fi44 - fi36)) | 1))
                                        }
                                        v56 := GenS8{F0: (^(fi36)), F1: (v2 ^ v46)}
                                        if len(v55) > 0 {
                                            v55[0] = genFn0(42)
                                        }
                                        v46 -= v46
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v47
                                case "a":
                                    var v57 int64 = 9223372036854775807
                                    _ = v57
                                }
                            }
                            if (42 < -232) {
                                var v58 uint8 = 127
                                v59 := GenS9{F0: 309, F1: (fi44 << ((255) & 63)), F2: 2}
                                _ = v58
                                _ = v59
                            } else {
                                var v60 uint = v6
                                switch v35 {
                                case "":
                                    switch -987 {
                                    case -1:
                                        fmt.Println(fi44)
                                    }
                                case "\n\t":
                                    v61 := GenS2{F0: 127}
                                    _ = v61
                                }
                                fmt.Println(fi36)
                                _ = v60
                            }
                        } else {
                            var v62 uint64 = 460
                            _ = v62
                        }
                        _ = v46
                    } else {
                        v45 -= v3
                        var v63 uint = 255
                        _ = v63
                    }
                    var v64 int16 = 715
                    v3 = v45
                    _ = v45
                    _ = v64
                }
                v2 -= v2
            }
            var v65 float32 = (v3 + v3)
            if true {
                v66 := GenS7{F0: genFn0(-1)}
                v35 = v35
                _ = v66
            } else {
                if (32767 >= genFn1(fi44)) {
                    var v67 string = "hello"
                    _ = v67
                }
                v35 = ("" + v35)
                var v68 uint = v6
                _ = v68
            }
            _ = fi44
            _ = v65
        }
        _ = fi36
        _ = v43
    }
    for fi69 := 0; fi69 < 2; fi69++ {
        var v70 int = 123
        fmt.Println(v3)
        v71 := GenS3{F0: 4294967295, F1: false, F2: 1}
        _ = fi69
        _ = v70
        _ = v71
    }
    switch genFn0(-380) {
    case -2147483648:
        if true {
            var v72 uint64 = 2
            _ = v72
        } else {
            fmt.Println(v3)
            fmt.Println(v2, v35)
        }
    }
    switch genFn0(361) {
    case 32768:
        for fi73 := 0; fi73 < 19; fi73++ {
            fmt.Println(v6)
            for fi74 := 0; fi74 < 3; fi74++ {
                v35 = "\n\t"
                _ = fi74
            }
            _ = fi73
        }
    case -9223372036854775808:
        if (-2147483648 <= 4294967295) {
            var v75 uint8 = 255
            fmt.Println(v75, v35)
            _ = v75
        }
    default:
        for fi76 := 0; fi76 < 20; fi76++ {
            var v77 bool = true
            _ = fi76
            _ = v77
        }
    }
    var v78 int16 = 1
    var v79 uint64 = 18446744073709551615
    v80 := GenS10{F0: false}
    v81 := GenS4{F0: 1, F1: ((v6 ^ 255) != v6)}
    v2 -= v2
    if (452 <= v6) {
        v82 := []int{127, 42, 32768}
        _ = v82
    } else {
        for fi83 := 0; fi83 < 1; fi83++ {
            var v84 uint8 = 55
            var v85 float32 = v3
            _ = fi83
            _ = v84
            _ = v85
        }
        delete(v1, 708)
        for fi86 := 0; fi86 < 9; fi86++ {
            v87 := []int{}
            for fi88 := 0; fi88 < 16; fi88++ {
                v89 := GenS1{F0: 32767, F1: v79, F2: (fi86 | (fi86 & fi88))}
                if (0 > -128) {
                    v90 := GenS11{F0: 412, F1: int32(v78)}
                    var v91 float64 = float64((1e10 * 3.14159))
                    _ = len(v87)
                    _ = v90
                    _ = v91
                } else {
                    if (v6 != 579) {
                        switch fi86 {
                        case 32767:
                            switch (((((v35 + v35) + (v35 + "foo bar")) + v35) + v35) + "a") {
                            case "foo bar":
                                var v92 int16 = (v78 ^ v78)
                                _ = v92
                            case "a":
                                fmt.Println(v6, fi86)
                            case "\n\t":
                                var v93 uint16 = 414
                                _ = v93
                            default:
                                fmt.Println(v79)
                            }
                        case 0:
                            if false {
                                v6 = (v6 % ((563) | 1))
                                if ((v35 < "hello") || (!false)) {
                                    fmt.Println(v78, v2)
                                } else {
                                    v94 := GenS1{F0: 200, F1: v79, F2: -1}
                                    if len(v87) >= 2 {
                                        v87 = v87[1:len(v87)]
                                    }
                                    var v95 int = genFn0(0)
                                    _ = v94
                                    _ = v95
                                }
                            } else {
                                var v96 uint8 = 1
                                switch v35 {
                                case "foo bar":
                                    if ((false && false) || false) {
                                        fmt.Println(v96, v6)
                                        v97 := GenS12{F0: 42, F1: 42, F2: 0}
                                        _ = v97
                                    } else {
                                        var v98 int = genFn0(762)
                                        switch fi86 {
                                        case 127:
                                            v6 -= (v6 ^ v6)
                                        case 554:
                                            fmt.Println(v6, fi86)
                                            _ = len(v1)
                                        case 680:
                                            v78 -= v78
                                        }
                                        var v99 float32 = float32(1e10)
                                        var v100 bool = true
                                        var v101 uint8 = (v96 / (((v96 << ((v6) & 63))) | 1))
                                        _ = v100
                                        _ = v101
                                        _ = v98
                                        _ = v99
                                    }
                                    if len(v87) > 0 {
                                        v87[0] = -175
                                    }
                                    var v102 int16 = 0
                                    _ = v102
                                case "hello":
                                    var v103 int16 = v78
                                    v1[65535] = 1
                                    _ = v103
                                case "\n\t":
                                    var v104 int32 = -869
                                    _ = v104
                                }
                                _ = v96
                            }
                        }
                    } else {
                        switch (v35 + "hello") {
                        case "hello":
                            var v105 int8 = 127
                            _ = v105
                        }
                        v106 := GenS13{F0: fi88, F1: 2, F2: 2}
                        switch (fi86 % ((fi88) | 1)) {
                        case 1:
                            fmt.Println(v3, v2)
                        case 0:
                            var v107 int32 = v2
                            _ = v107
                        }
                        _ = v106
                    }
                    var v108 uint64 = (v79 >> (((v6 ^ 0)) & 63))
                    _ = v108
                }
                switch "hello" {
                case "a":
                    v78 = (v78 * v78)
                }
                _ = fi88
                _ = v89
            }
            _ = fi86
            _ = v87
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
