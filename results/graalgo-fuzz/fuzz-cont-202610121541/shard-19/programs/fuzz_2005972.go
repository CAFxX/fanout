package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint64; F1 int32; F2 int }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 bool; F2 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint }
func genFn0(p0 int64, p1 uint16) int {
    v5 := []int{1, 0, 4294967295}
    v6 := GenS1{F0: 42, F1: 83}
    if ((true && true) || false) {
        p1 = 240
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 bool, p1 int, p2 bool) uint16 {
    var v7 bool = true
    p0 = true
    if (-1 <= -317) {
        v8 := []int{p1, p1, p1, genFn0(-515, 1)}
        fmt.Println(p1)
        _ = v8
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 int = 9223372036854775807
    if (!true) {
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(v4)
            v10 := []int{v4}
            _ = fi9
            _ = v10
        }
        var v11 uint32 = (v1 << ((4294967295) & 63))
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        if false {
            var v13 string = "\n\t"
            v13 = "foo bar"
            _ = v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                switch "" {
                case "foo bar":
                    var v15 int32 = -2147483648
                    _ = v15
                case "":
                    v1 = (^(v1))
                }
                v4 = fi12
                fmt.Println(v4, v1)
                _ = fi14
            }
        }
        if true {
            for fi16 := 0; fi16 < 13; fi16++ {
                v1 -= (v1 >> ((65535) & 63))
                switch ("" + ((("" + (("foo bar" + "\n\t") + "\n\t")) + "a") + ((("\n\t" + "") + "hello") + "hello"))) {
                case "\n\t":
                    var v17 int64 = -9223372036854775808
                    _ = v17
                case "hello":
                    delete(v2, (^(fi12)))
                }
                var v18 bool = false
                _ = fi16
                _ = v18
            }
            if true {
                v4 = 1
                if (!true) {
                    switch "foo bar" {
                    case "hello":
                        v19 := GenS2{F0: genFn1((!false), -9223372036854775808, false), F1: 78, F2: 65535}
                        _ = v19
                    case "a":
                        var v20 bool = (-616 > v4)
                        _ = v20
                    default:
                        if (1e10 >= (0.1 + 1e10)) {
                            if true {
                                var v21 bool = (-17 != -1)
                                if ((^(v4)) >= v4) {
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        var v23 uint64 = 135
                                        v2[4294967295] = (fi22 + -110)
                                        _ = fi22
                                        _ = v23
                                    }
                                    var v24 float32 = 0.1
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v21, fi12, v24)
                                    case "foo bar":
                                        v21 = (true || (!v21))
                                    }
                                    _ = v24
                                } else {
                                    v21 = true
                                }
                                _ = v21
                            } else {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    switch "" {
                                    case "\n\t":
                                        v4 = (^(fi12))
                                    case "foo bar":
                                        var v26 uint32 = (v1 << ((87) & 63))
                                        _ = v26
                                    case "hello":
                                        var v27 uint8 = 255
                                        _ = v27
                                    }
                                    if true {
                                        _ = len(v2)
                                        if (true && (!true)) {
                                            v28 := []int{-2147483648, genFn0(63, 255)}
                                            v29 := GenS1{F0: 971, F1: 0}
                                            _ = v28
                                            _ = v29
                                        } else {
                                            delete(v3, v1)
                                        }
                                        v1 += 394
                                    }
                                    _ = fi25
                                }
                                v30 := []int{}
                                if false {
                                    var v31 string = "hello"
                                    v32 := v3[v1]
                                    var v33 float32 = 1e10
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                _ = v30
                            }
                        } else {
                            v34 := GenS3{F0: genFn1(false, v4, (!false)), F1: -587}
                            _ = v34
                        }
                    }
                    var v35 bool = true
                    v36 := GenS1{F0: 38, F1: 1}
                    _ = v35
                    _ = v36
                } else {
                    switch "\n\t" {
                    case "a":
                        for fi37 := 0; fi37 < 1; fi37++ {
                            _ = len(v3)
                            if true {
                                v1 = v1
                                var v38 int64 = -9223372036854775808
                                _ = v38
                            }
                            _ = fi37
                        }
                    }
                    delete(v3, (v1 | (^(v1))))
                    v39 := GenS3{F0: genFn1((fi12 < genFn0(1, 892)), genFn0(0, 2), false), F1: -94}
                    _ = v39
                }
                for fi40 := 0; fi40 < 15; fi40++ {
                    if (!false) {
                        var v41 uint32 = 2147483647
                        v41 = 42
                        _ = v41
                    } else {
                        v42 := []int{}
                        _ = v42
                    }
                    var v43 int16 = 1
                    _ = fi40
                    _ = v43
                }
            }
        } else {
            if ((23 != uint64(v1)) && (!(176 != 4294967295))) {
                v44 := GenS3{F0: 776, F1: 42}
                var v45 int64 = 0
                switch "" {
                case "a":
                    if (197 > 1) {
                        switch "foo bar" {
                        case "hello":
                            var v46 uint8 = 246
                            _ = v46
                        case "":
                            v45 += (v45 - v45)
                        case "\n\t":
                            switch "a" {
                            case "":
                                v47 := GenS4{F0: 42, F1: 42, F2: (v4 ^ v4)}
                                _ = v47
                            case "hello":
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            var v48 uint16 = 0
                                            _ = v48
                                        } else {
                                            var v49 uint = 2
                                            _ = v49
                                        }
                                    }
                                default:
                                    v45 = -9223372036854775808
                                }
                            default:
                                v50 := GenS5{F0: 100, F1: 32767}
                                _ = v50
                            }
                        default:
                            var v51 uint16 = 42
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 9; fi52++ {
                            var v53 float32 = 0.1
                            var v54 int8 = 127
                            if false {
                                if (0.1 > 3.14159) {
                                    var v55 float64 = ((0.1 / 0.1) + 0.1)
                                    _ = v55
                                } else {
                                    v45 = 1
                                    v56 := GenS6{F0: 32767}
                                    _ = v56
                                }
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        if (219 >= 2) {
                            v57 := []int{}
                            for fi58 := 0; fi58 < 4; fi58++ {
                                var v59 int8 = -69
                                v60 := GenS7{F0: 0, F1: (-134 != 46), F2: 1}
                                v61 := GenS1{F0: 0, F1: -813}
                                var v62 int8 = (v59 + 127)
                                v59 += (-(v59))
                                _ = fi58
                                _ = v59
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            var v63 string = "\n\t"
                            _ = v57
                            _ = v63
                        }
                    }
                case "hello":
                    for fi64 := 0; fi64 < 7; fi64++ {
                        v45 = (v45 % ((v45) | 1))
                        _ = fi64
                    }
                }
                _ = v44
                _ = v45
            } else {
                fmt.Println(v1, v1)
                if ((!false) && false) {
                    if (((^(fi12)) < v4) && true) {
                        switch (v4 >> ((262) & 63)) {
                        case -9223372036854775808:
                            switch ((("" + ("\n\t" + (("\n\t" + ("hello" + "")) + ("hello" + ("foo bar" + "\n\t"))))) + "foo bar") + "") {
                            case "a":
                                if false {
                                    v1 = (v1 << ((338) & 63))
                                } else {
                                    var v65 uint64 = 862
                                    v66 := GenS8{F0: 353}
                                    _ = v65
                                    _ = v66
                                }
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            delete(v2, (v4 + 128))
                        }
                        var v67 int8 = -3
                        _ = v67
                    }
                    var v68 bool = (3.14159 < 1.5)
                    _ = v68
                } else {
                    switch (("\n\t" + "") + "hello") {
                    case "foo bar":
                        for fi69 := 0; fi69 < 15; fi69++ {
                            v1 = (v1 + (v1 / ((2147483647) | 1)))
                            if true {
                                v70 := GenS9{F0: 65535, F1: 127, F2: 318}
                                _ = v70
                            } else {
                                v1 -= (v1 % ((v1) | 1))
                                var v71 int = v4
                                _ = v71
                            }
                            _ = fi69
                        }
                    case "\n\t":
                        var v72 uint16 = 255
                        _ = v72
                    default:
                        for fi73 := 0; fi73 < 17; fi73++ {
                            fmt.Println(fi73, fi73)
                            v74 := []int{(fi12 * (fi12 >> ((65535) & 63))), (v4 ^ genFn0(34, 0)), (fi12 ^ (v4 | fi12)), (fi73 * v4)}
                            v75 := v2[128]
                            _ = fi73
                            _ = v74
                            _ = v75
                        }
                    }
                    var v76 uint64 = 18446744073709551615
                    v77 := GenS3{F0: 64, F1: -1}
                    _ = v76
                    _ = v77
                }
                v78 := GenS6{F0: -1}
                _ = v78
            }
        }
        _ = fi12
    }
    v1 = v1
    var v79 uint32 = 0
    switch (v4 | 42) {
    case 127:
        v79 -= v1
    case 4294967295:
        v80 := GenS9{F0: 333, F1: 190, F2: 938}
        _ = v80
    }
    v79 = (v79 << ((682) & 63))
    if true {
        for fi81 := 0; fi81 < 14; fi81++ {
            _ = len(v3)
            switch (v4 * genFn0(0, 255)) {
            case 0:
                v82 := []int{519, (v4 & fi81), int(v79), (v4 & (^(fi81)))}
                _ = v82
            default:
                fmt.Println(fi81, v1)
            }
            _ = fi81
        }
        var v83 int8 = 127
        _ = v83
    }
    if true {
        if (uint16(v79) > 255) {
            var v84 int = (v4 << ((812) & 63))
            var v85 float32 = 0.0
            if (!(3.14159 <= v85)) {
                fmt.Println(v79)
            }
            _ = v84
            _ = v85
        }
        v4 -= 0
    } else {
        var v86 string = ((("" + "foo bar") + "hello") + "hello")
        delete(v3, (v1 | 0))
        var v87 uint = 1
        _ = v86
        _ = v87
    }
    v88 := GenS5{F0: 127, F1: 11}
    v89 := GenS10{F0: 151}
    if true {
        var v90 int16 = 639
        for fi91 := 0; fi91 < 12; fi91++ {
            v1 += v79
            _ = fi91
        }
        v92 := GenS11{F0: 255}
        _ = v90
        _ = v92
    } else {
        v79 = (v79 - 1)
        if (11 != 42) {
            var v93 int32 = 2147483647
            _ = v93
        } else {
            _ = len(v3)
            v4 += int(v79)
        }
        v94 := GenS3{F0: 255, F1: 1}
        _ = v94
    }
    v95 := GenS7{F0: 4294967295, F1: (((!true) && false) && true), F2: 418}
    for fi96 := 0; fi96 < 10; fi96++ {
        var v97 int64 = -83
        switch fi96 {
        case 128:
            var v98 string = (("a" + "hello") + "hello")
            _ = v98
        case -2147483648:
            v99 := GenS11{F0: 0}
            _ = v99
        case 127:
            switch "hello" {
            case "\n\t":
                v100 := v3[(v79 % ((65535) | 1))]
                _ = v100
            case "foo bar":
                fmt.Println(v4)
            case "a":
                switch fi96 {
                case 4294967295:
                    v101 := GenS12{F0: 1, F1: 42, F2: 65535}
                    _ = v101
                case -37:
                    v102 := GenS4{F0: 1, F1: 545, F2: fi96}
                    _ = v102
                }
            default:
                v1 = v79
            }
        default:
            v79 -= (v79 / ((4294967295) | 1))
        }
        var v103 float64 = 0.1
        _ = fi96
        _ = v103
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v79
    _ = v88
    _ = v89
    _ = v95
}
