package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 string; F1 int32; F2 string }
type GenS9 struct { F0 bool; F1 bool; F2 int16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int8, p1 int32) uint16 {
    var v4 uint8 = 42
    var v5 uint8 = v4
    switch "" {
    case "":
        var v6 string = "a"
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        if true {
            v4 -= v4
        } else {
            switch (fi7 >> ((312) & 63)) {
            case 4294967295:
                v8 := []int{fi7, fi7, (fi7 << ((255) & 63)), fi7}
                _ = v8
            default:
                if ((p1 + (p1 | 0)) < p1) {
                    if true {
                        p0 -= 59
                        v9 := []int{(fi7 >> ((4294967295) & 63)), fi7}
                        if true {
                            var v10 uint = 42
                            if true {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            v11 := GenS1{F0: 646, F1: 960, F2: 42}
                            _ = v10
                            _ = v11
                        } else {
                            if false {
                                var v12 uint8 = v4
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(fi7, v4, p0)
                                }
                                fmt.Println(v5, fi7)
                                _ = v12
                            }
                            v4 = 0
                            for fi13 := 0; fi13 < 13; fi13++ {
                                var v14 string = (("\n\t" + "a") + "\n\t")
                                fmt.Println(v4, v4)
                                _ = fi13
                                _ = v14
                            }
                        }
                        _ = v9
                    } else {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := []int{398, 42}
                            var v17 int64 = 1
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(v5, v4)
                    }
                    var v18 int16 = -32768
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        switch (fi7 + (fi7 ^ (fi19 / ((32767) | 1)))) {
                        case 1:
                            v20 := []int{(-(fi19)), (fi19 & fi19)}
                            _ = v20
                        case -1:
                            v21 := GenS1{F0: 4294967295, F1: 42, F2: 548}
                            _ = v21
                        case 32767:
                            v22 := GenS1{F0: 18446744073709551615, F1: 713, F2: 1}
                            _ = v22
                        default:
                            v23 := GenS2{F0: 0, F1: v5}
                            _ = v23
                        }
                        fmt.Println(p1, p0)
                        _ = fi19
                    }
                    fmt.Println(p0, v5)
                    switch 0 {
                    case 901:
                        if ((^(v4)) != 1) {
                            var v24 float32 = (3.14159 * 0.0)
                            var v25 int8 = 100
                            _ = v24
                            _ = v25
                        }
                    case 0:
                        var v26 int64 = -9223372036854775808
                        _ = v26
                    case -1:
                        for fi27 := 0; fi27 < 5; fi27++ {
                            fmt.Println(fi7)
                            v28 := []int{}
                            _ = fi27
                            _ = v28
                        }
                    default:
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v5 = 1
                            var v30 float64 = (3.14159 * 3.14159)
                            _ = fi29
                            _ = v30
                        }
                    }
                }
            }
            if (false || (3.14159 >= ((float32(0.1) * (float32(1e10) + 1e10)) * (0.1 + 0.1)))) {
                p0 += (p0 % (((p0 / ((int8(v5)) | 1))) | 1))
            }
            v31 := []int{}
            _ = v31
        }
        var v32 int32 = p1
        var v33 int = fi7
        _ = fi7
        _ = v32
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint16 {
    v34 := GenS2{F0: 0, F1: 127}
    if true {
        switch 148 {
        case 0:
            v35 := []int{-2147483648}
            _ = v35
        case 9223372036854775807:
            var v36 bool = (255 == 0)
            _ = v36
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            var v38 string = ""
            v39 := []int{(-(fi37)), (fi37 ^ (fi37 + 113)), fi37}
            for fi40 := 0; fi40 < 13; fi40++ {
                if false {
                    var v41 float32 = 0.0
                    v42 := GenS3{F0: -2147483648}
                    _ = v41
                    _ = v42
                } else {
                    if true {
                        if len(v39) > 0 {
                            v39[0] = (fi40 ^ (^(fi40)))
                        }
                    }
                    v43 := GenS3{F0: fi37}
                    _ = v43
                }
                var v44 uint16 = genFn0(28, -2147483648)
                _ = fi40
                _ = v44
            }
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    _ = v34
    return 326
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    v3 := make(map[string]uint32)
    for fi45 := 0; fi45 < 14; fi45++ {
        if true {
            if (2147483647 < 765) {
                for fi46 := 0; fi46 < 19; fi46++ {
                    if false {
                        var v47 uint = 0
                        _ = v47
                    }
                    v2 -= v2
                    _ = fi46
                }
            } else {
                var v48 int = -983
                _ = v48
            }
            v49 := []int{}
            v50 := GenS4{F0: 65535}
            _ = v49
            _ = v50
        } else {
            for fi51 := 0; fi51 < 14; fi51++ {
                v52 := []int{42, (fi51 - 9223372036854775807), fi51}
                v2 += 1
                v53 := GenS1{F0: 18446744073709551615, F1: 457, F2: 65535}
                _ = fi51
                _ = v52
                _ = v53
            }
            v2 = (v2 << ((255) & 63))
        }
        v54 := v3["\n\t"]
        _ = fi45
        _ = v54
    }
    if (128 >= 128) {
        var v55 bool = (3.14159 == (0.1 / (((1.5 / (float64((0.0 - -2.25)) * -2.25)) + 1.5) + 3.14159)))
        v56 := GenS5{F0: 42, F1: 166}
        _ = v55
        _ = v56
    }
    v2 += (v2 / ((83) | 1))
    var v57 int = 127
    v57 -= (v57 + v57)
    var v58 int16 = 446
    if ((v57 == v57) || false) {
        var v59 int16 = v58
        _ = v59
    } else {
        v60 := []int{(v57 - v57), v57}
        var v61 int8 = v2
        for fi62 := 0; fi62 < 7; fi62++ {
            if (false || (-2.25 >= 1.5)) {
                fmt.Println(v2, v61, v57)
            } else {
                var v63 int = v57
                fmt.Println(v61, v61)
                v64 := GenS6{F0: 255}
                _ = v63
                _ = v64
            }
            var v65 uint = 560
            v65 += (v65 | 402)
            _ = fi62
            _ = v65
        }
        _ = v60
        _ = v61
    }
    var v66 float32 = float32((1e10 + (((1.5 * float64((0.0 + (0.1 / 3.14159)))) * 3.14159) + -2.25)))
    if (0.1 != (v66 + ((v66 * v66) / v66))) {
        v2 += (v2 >> ((267) & 63))
    }
    switch v57 {
    case 4294967295:
        if false {
            if (42 != 2147483648) {
                v58 = -1
            }
        }
    default:
        if (!false) {
            switch "hello" {
            case "foo bar":
                v67 := GenS5{F0: v58, F1: 1}
                _ = v67
            case "\n\t":
                if ("\n\t" != "foo bar") {
                    var v68 int = v57
                    if (!false) {
                        v69 := []int{v68, (v57 % ((v57) | 1)), 320, v57}
                        fmt.Println(v57, v57)
                        _ = v69
                    }
                    v66 = (v66 - 1.5)
                    _ = v68
                } else {
                    switch (v57 << ((42) & 63)) {
                    case 128:
                        for fi70 := 0; fi70 < 15; fi70++ {
                            v71 := GenS7{F0: 333, F1: (((("" + "") + "foo bar") + "foo bar") + "hello")}
                            for fi72 := 0; fi72 < 7; fi72++ {
                                v73 := []int{(fi70 >> ((214) & 63)), v57, int(v58)}
                                _ = len(v3)
                                v74 := GenS8{F0: ("a" + "foo bar"), F1: 2147483647, F2: ("" + (("hello" + "\n\t") + "foo bar"))}
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            v75 := GenS3{F0: (v57 - (fi70 | v57))}
                            _ = fi70
                            _ = v71
                            _ = v75
                        }
                    }
                    v66 += 1e10
                }
            case "":
                v2 += 89
            }
        } else {
            var v76 float32 = -2.25
            for fi77 := 0; fi77 < 14; fi77++ {
                v57 = (v57 / ((fi77) | 1))
                v3[(("" + "a") + "foo bar")] = 1
                v58 += 1
                _ = fi77
            }
            _ = v76
        }
    }
    for fi78 := 0; fi78 < 6; fi78++ {
        v57 -= (v57 - (v57 ^ (fi78 * (fi78 - 42))))
        v79 := []int{(v57 * (fi78 * v57)), v57}
        for fi80 := 0; fi80 < 5; fi80++ {
            v58 -= 191
            if (!(127 < 44)) {
                if true {
                    var v81 uint8 = 255
                    v66 = 0.0
                    _ = v81
                }
                fmt.Println(fi78, fi78)
            } else {
                var v82 float64 = (1e10 * ((-2.25 / 1e10) * 1e10))
                switch (("a" + (("hello" + "foo bar") + "foo bar")) + "hello") {
                case "":
                    switch (v57 % (((fi78 ^ (fi80 * (fi80 - fi80)))) | 1)) {
                    case 477:
                        fmt.Println(v82)
                    default:
                        _ = len(v1)
                        if (!false) {
                            var v83 uint16 = 0
                            v83 = v83
                            switch "" {
                            case "\n\t":
                                v2 = v2
                            case "a":
                                v84 := GenS9{F0: (2 >= 42), F1: (v66 < (v66 - v66)), F2: (v58 % ((v58) | 1))}
                                if (v83 >= v83) {
                                    if len(v79) >= 2 {
                                        v79 = v79[1:len(v79)]
                                    }
                                    var v85 bool = true
                                    v82 -= v82
                                    var v86 uint32 = 1
                                    switch "hello" {
                                    case "hello":
                                        var v87 int64 = 0
                                        _ = v87
                                    case "\n\t":
                                        if len(v79) > 0 {
                                            v79[0] = (v57 >> ((1) & 63))
                                        }
                                        var v88 float32 = (0.0 + v66)
                                        _ = v88
                                    case "foo bar":
                                        var v89 uint8 = uint8(v83)
                                        _ = v89
                                    }
                                    _ = v85
                                    _ = v86
                                }
                                _ = v84
                            }
                            fmt.Println(v57, fi78)
                            var v90 int = (fi80 & -9223372036854775808)
                            _ = v83
                            _ = v90
                        } else {
                            if false {
                                v66 = (3.14159 * float32((3.14159 + (float64(float32(-2.25)) + v82))))
                                fmt.Println(v2, v2, fi80)
                            }
                            if len(v79) > 0 {
                                v79[0] = 32768
                            }
                        }
                        if false {
                            _ = len(v79)
                        }
                        v91 := GenS10{F0: 0}
                        var v92 int = (fi80 | -2147483648)
                        var v93 uint16 = 32767
                        _ = v91
                        _ = v92
                        _ = v93
                    }
                }
                if len(v79) >= 2 {
                    v79 = v79[1:len(v79)]
                }
                _ = v82
            }
            fmt.Println(v2)
            v58 += 0
            v57 -= (fi80 * 4294967295)
            _ = fi80
        }
        _ = fi78
        _ = v79
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            var v94 int8 = (v2 % ((v2) | 1))
            v95 := GenS11{F0: 77}
            v96 := GenS11{F0: 255}
            _ = v94
            _ = v95
            _ = v96
        } else {
            v97 := GenS6{F0: 2}
            switch "foo bar" {
            case "hello":
                v98 := GenS6{F0: 791}
                _ = v98
            case "a":
                v99 := []int{}
                _ = v99
            default:
                v2 = v2
            }
            _ = v97
        }
    case "\n\t":
        var v100 uint8 = 42
        _ = v100
    default:
        if (!true) {
            var v101 uint32 = 1
            _ = v101
        }
    }
    if (4294967295 <= 456) {
        var v102 bool = true
        if (!v102) {
            v102 = v102
        } else {
            v2 = v2
        }
        switch "a" {
        case "\n\t":
            for fi103 := 0; fi103 < 7; fi103++ {
                fmt.Println(v57)
                fmt.Println(fi103, v57, v58)
                _ = fi103
            }
        }
        _ = v102
    }
    v104 := []int{v57}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v66
}
