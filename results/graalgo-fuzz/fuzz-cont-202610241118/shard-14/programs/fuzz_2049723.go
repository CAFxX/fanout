package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS12 struct { F0 int }
func genFn0(p0 int16) int32 {
    v5 := []int{-259}
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = p0
        p0 -= p0
        _ = fi6
    }
    p0 = (p0 | p0)
    var v7 uint = 4294967295
    _ = p0
    _ = v5
    _ = v7
    return 168
}
func genFn1(p0 bool, p1 int32, p2 bool) uint16 {
    if (32767 != 568) {
        if (32767 > 65535) {
            var v8 uint = 665
            p0 = ((p1 * 2147483647) >= p1)
            _ = v8
        }
    } else {
        v9 := GenS1{F0: 42}
        if p2 {
            p0 = false
        } else {
            p1 -= (p1 ^ (p1 & -2147483648))
            switch 32768 {
            case 857:
                v10 := []int{1, 127}
                _ = v10
            case 32768:
                var v11 uint32 = 774
                _ = v11
            default:
                p1 += (p1 % ((p1) | 1))
            }
            if p0 {
                var v12 uint8 = 170
                if false {
                    v13 := []int{}
                    fmt.Println(p1, p1)
                    for fi14 := 0; fi14 < 5; fi14++ {
                        p2 = p0
                        _ = fi14
                    }
                    _ = v13
                } else {
                    v15 := []int{1}
                    _ = v15
                }
                if (390 != 2) {
                    v16 := []int{-390, 9223372036854775807}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 uint = 42
                        var v19 string = ""
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = v12
            }
        }
        switch 355 {
        case 795:
            fmt.Println(p1)
        case 0:
            v20 := []int{}
            _ = v20
        default:
            var v21 uint8 = 2
            _ = v21
        }
        _ = v9
    }
    var v22 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 532
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 1
    var v3 string = "hello"
    v4 := make(map[int]bool)
    _ = len(v1)
    v3 = v3
    var v23 int16 = 350
    var v24 uint32 = 2
    var v25 uint64 = 783
    for fi26 := 0; fi26 < 4; fi26++ {
        switch (v3 + "foo bar") {
        case "a":
            if (42 <= 0) {
                v3 = "a"
                v25 = 42
            } else {
                switch (fi26 - 2147483647) {
                case -791:
                    v27 := GenS2{F0: 158, F1: v2, F2: v2}
                    _ = v27
                case 32767:
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 int = -1
                        var v30 int64 = -9223372036854775808
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                v31 := GenS3{F0: (v24 | 1), F1: true, F2: 18446744073709551615}
                _ = v31
            }
        default:
            v24 += (v24 << ((26) & 63))
        }
        var v32 int64 = 42
        _ = fi26
        _ = v32
    }
    switch 127 {
    case 128:
        if true {
            for fi33 := 0; fi33 < 10; fi33++ {
                v24 = v24
                fmt.Println(v2)
                _ = fi33
            }
            v25 = uint64(v2)
        }
    case 9223372036854775807:
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 int64 = 1
            switch ((v3 + v3) + v3) {
            case "hello":
                v2 = genFn1((!(v25 < (v25 + v25))), 2147483647, (0 == 0))
            case "a":
                v24 -= 511
            case "foo bar":
                if true {
                    fmt.Println(v24, v2)
                } else {
                    for fi36 := 0; fi36 < 6; fi36++ {
                        fmt.Println(v25, v2, fi34)
                        _ = fi36
                    }
                    switch fi34 {
                    case 32768:
                        if (((1.5 - 0.0) - float32(1.5)) > -2.25) {
                            v37 := GenS4{F0: (fi34 | (fi34 * fi34))}
                            var v38 int16 = 39
                            _ = v37
                            _ = v38
                        } else {
                            v25 -= v25
                        }
                    case -1:
                        if false {
                            var v39 int64 = v35
                            v2 -= 255
                            for fi40 := 0; fi40 < 20; fi40++ {
                                if ((!false) || ((-2147483648 <= 42) || (-2147483648 >= genFn0((v23 | (v23 - (-(v23)))))))) {
                                    v41 := GenS4{F0: -1}
                                    if (float64(-2.25) <= 1.5) {
                                        fmt.Println(v35)
                                        fmt.Println(fi40, v2)
                                        fmt.Println(v25, fi34)
                                    } else {
                                        if (0 > 2147483647) {
                                            v42 := []int{}
                                            v43 := GenS4{F0: fi40}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v25 += (v25 - (v25 - (v25 | v25)))
                                            var v44 uint8 = 54
                                            _ = v44
                                        }
                                    }
                                    _ = v41
                                }
                                var v45 float32 = 1.5
                                v46 := GenS5{F0: v3}
                                delete(v1, (fi34 / (((fi34 | fi40)) | 1)))
                                _ = fi40
                                _ = v45
                                _ = v46
                            }
                            _ = v39
                        } else {
                            switch fi34 {
                            case 127:
                                var v47 uint16 = uint16(v25)
                                _ = v47
                            case -2147483648:
                                if true {
                                    switch fi34 {
                                    case 127:
                                        for fi48 := 0; fi48 < 10; fi48++ {
                                            v49 := GenS6{F0: 127, F1: 981}
                                            _ = fi48
                                            _ = v49
                                        }
                                    case 424:
                                        v50 := GenS7{F0: 127}
                                        _ = v50
                                    case 42:
                                        v51 := []int{}
                                        _ = v51
                                    default:
                                        var v52 uint = 793
                                        _ = v52
                                    }
                                    switch fi34 {
                                    case 4294967295:
                                        var v53 int64 = (v35 << ((0) & 63))
                                        _ = v53
                                    case 32768:
                                        for fi54 := 0; fi54 < 5; fi54++ {
                                            var v55 float64 = 0.0
                                            v35 -= (v35 << ((18446744073709551615) & 63))
                                            switch "" {
                                            case "":
                                                var v56 uint32 = 4294967295
                                                v57 := GenS3{F0: 4294967295, F1: true, F2: 1}
                                                _ = v56
                                                _ = v57
                                            case "foo bar":
                                                v58 := GenS8{F0: (-(v35))}
                                                _ = v58
                                            case "\n\t":
                                                v3 = (v3 + v3)
                                            default:
                                                var v59 float32 = 0.1
                                                _ = v59
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                    default:
                                        fmt.Println(v25, fi34, fi34)
                                    }
                                }
                            case 128:
                                fmt.Println(v24, v35)
                            }
                            delete(v1, fi34)
                            v60 := GenS3{F0: (v24 + (v24 + (v24 + v24))), F1: (v3 >= v3), F2: 42}
                            _ = v60
                        }
                    case 0:
                        switch (fi34 % ((fi34) | 1)) {
                        case 128:
                            delete(v4, 9223372036854775807)
                        case 127:
                            v61 := GenS9{F0: 79, F1: int(v25)}
                            _ = v61
                        default:
                            var v62 uint16 = (v2 >> ((18446744073709551615) & 63))
                            _ = v62
                        }
                    }
                }
            default:
                v63 := GenS3{F0: v24, F1: false, F2: 1}
                _ = v63
            }
            for fi64 := 0; fi64 < 10; fi64++ {
                var v65 string = v3
                _ = fi64
                _ = v65
            }
            _ = fi34
            _ = v35
        }
    case 994:
        switch v3 {
        case "a":
            var v66 float32 = (-2.25 * (-2.25 - (3.14159 - (3.14159 / 1e10))))
            _ = v66
        case "\n\t":
            switch (v3 + v3) {
            case "":
                if false {
                    if true {
                        for fi67 := 0; fi67 < 15; fi67++ {
                            if (v23 == v23) {
                                fmt.Println(v2)
                                v68 := []int{}
                                _ = v68
                            }
                            _ = fi67
                        }
                        for fi69 := 0; fi69 < 17; fi69++ {
                            v70 := GenS9{F0: (v24 * v24), F1: -492}
                            v71 := GenS8{F0: -1}
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                        switch -982 {
                        case -2147483648:
                            v23 = v23
                        case 9223372036854775807:
                            var v72 int = 32767
                            _ = v72
                        }
                    }
                    v25 = (v25 * (v25 + (v25 & (v25 + (v25 / ((v25) | 1))))))
                    if false {
                        v23 = (v23 >> ((1) & 63))
                        var v73 bool = false
                        _ = v73
                    }
                } else {
                    for fi74 := 0; fi74 < 3; fi74++ {
                        for fi75 := 0; fi75 < 20; fi75++ {
                            v76 := GenS4{F0: int(v23)}
                            if true {
                                var v77 uint32 = v24
                                v78 := GenS10{F0: 42}
                                if ((-2.25 + 0.0) < 0.0) {
                                    v25 = (^(v25))
                                    var v79 uint16 = 568
                                    v23 = v23
                                    _ = v79
                                }
                                _ = v77
                                _ = v78
                            } else {
                                fmt.Println(v3, v23)
                                v80 := []int{1, (fi74 / ((fi75) | 1))}
                                _ = v80
                            }
                            _ = fi75
                            _ = v76
                        }
                        _ = fi74
                    }
                    switch 40 {
                    case 42:
                        var v81 int32 = -2147483648
                        _ = v81
                    default:
                        v3 = v3
                    }
                    fmt.Println(v24, v23)
                }
            case "foo bar":
                fmt.Println(v24, v25)
            case "hello":
                _ = len(v4)
            default:
                for fi82 := 0; fi82 < 10; fi82++ {
                    switch int(v2) {
                    case 9223372036854775807:
                        _ = len(v4)
                    case 9:
                        _ = len(v4)
                    case 1:
                        var v83 int32 = 2147483647
                        _ = v83
                    }
                    var v84 int8 = -69
                    for fi85 := 0; fi85 < 15; fi85++ {
                        fmt.Println(v3, v84)
                        if true {
                            v86 := GenS4{F0: fi85}
                            v23 += -32768
                            switch (v3 + v3) {
                            case "a":
                                switch (v3 + ("" + "foo bar")) {
                                case "\n\t":
                                    v87 := GenS7{F0: 0}
                                    _ = v87
                                }
                            case "":
                                if (true && (v3 == v3)) {
                                    v88 := GenS11{F0: (v24 + (v24 ^ (^(v24)))), F1: -2147483648, F2: 1}
                                    _ = v88
                                }
                            case "hello":
                                switch v3 {
                                case "foo bar":
                                    v89 := GenS12{F0: fi82}
                                    _ = v89
                                case "hello":
                                    v90 := []int{}
                                    _ = v90
                                case "":
                                    switch (fi82 / ((fi82) | 1)) {
                                    case -923:
                                        var v91 uint32 = v24
                                        var v92 int = fi85
                                        var v93 int = fi82
                                        v94 := GenS2{F0: 954, F1: (v2 * (v2 * uint16(v91))), F2: 255}
                                        _ = v91
                                        _ = v92
                                        _ = v93
                                        _ = v94
                                    case 75:
                                        v2 -= uint16(v25)
                                    }
                                }
                            default:
                                var v95 int = fi82
                                _ = v95
                            }
                            _ = v86
                        } else {
                            var v96 int16 = int16(v24)
                            v97 := []int{int(v2), (fi85 - -329)}
                            _ = v96
                            _ = v97
                        }
                        var v98 float32 = 0.0
                        _ = fi85
                        _ = v98
                    }
                    _ = fi82
                    _ = v84
                }
            }
        }
    default:
        for fi99 := 0; fi99 < 13; fi99++ {
            v100 := GenS8{F0: 1}
            if true {
                for fi101 := 0; fi101 < 11; fi101++ {
                    var v102 uint32 = v24
                    var v103 uint64 = (v25 & 4294967295)
                    v103 = 272
                    _ = fi101
                    _ = v102
                    _ = v103
                }
                if (100 < 42) {
                    var v104 int64 = -47
                    _ = v104
                }
            } else {
                v105 := GenS10{F0: 2}
                for fi106 := 0; fi106 < 12; fi106++ {
                    var v107 int8 = 127
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v107)
                    }
                    var v108 float64 = (-2.25 - 0.0)
                    _ = fi106
                    _ = v107
                    _ = v108
                }
                _ = v105
            }
            _ = fi99
            _ = v100
        }
    }
    fmt.Println(v25, v3, v3)
    if true {
        var v109 int16 = (v23 >> ((1) & 63))
        fmt.Println(v25, v25)
        v23 += (v109 << ((18446744073709551615) & 63))
        _ = v109
    } else {
        var v110 float32 = (1.5 * 3.14159)
        v110 = v110
        _ = v110
    }
    fmt.Println(v25)
    v1[-732] = false
    var v111 int32 = 993
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v111
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
