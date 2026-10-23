package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int16; F2 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32 }
func genFn0() uint16 {
    if (!true) {
        var v5 uint16 = 2
        _ = v5
    } else {
        switch -9223372036854775808 {
        case 42:
            switch "foo bar" {
            case "a":
                var v6 uint64 = 160
                _ = v6
            }
        }
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        var v8 uint64 = 4294967295
        var v9 string = ("a" + ("\n\t" + ("hello" + "foo bar")))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = 127
    if (2147483647 < 298) {
        fmt.Println(v10)
        var v11 float32 = 0.1
        switch ("foo bar" + (("foo bar" + "foo bar") + "\n\t")) {
        case "foo bar":
            v12 := []int{(^(v10))}
            _ = v12
        default:
            v13 := []int{v10, v10}
            _ = v13
        }
        _ = v11
    }
    _ = v10
    return 982
}
func genFn1(p0 uint16, p1 int32, p2 int) int {
    v14 := GenS1{F0: 127, F1: true, F2: 18446744073709551615}
    v15 := GenS1{F0: 1, F1: false, F2: 34}
    var v16 float32 = -2.25
    var v17 int = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return p2
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 4294967295
    var v3 uint32 = 994
    var v4 uint64 = (v2 - (v2 - v2))
    fmt.Println(v2)
    for fi18 := 0; fi18 < 2; fi18++ {
        switch (fi18 + fi18) {
        case 42:
            fmt.Println(v4, v3, v4)
        case 9223372036854775807:
            for fi19 := 0; fi19 < 4; fi19++ {
                if true {
                    v20 := GenS1{F0: 1, F1: (false || true), F2: (v2 / ((v2) | 1))}
                    v4 = (v2 >> ((2) & 63))
                    v4 = v2
                    _ = v20
                }
                v4 = (v2 << ((698) & 63))
                _ = fi19
            }
        default:
            v3 += 1
        }
        delete(v1, (fi18 * (fi18 / ((2147483647) | 1))))
        _ = fi18
    }
    if (("hello" + "foo bar") < ("foo bar" + "\n\t")) {
        delete(v1, genFn1(273, 1, 9223372036854775807))
    }
    v21 := GenS1{F0: 42, F1: (42 == 65535), F2: (v4 >> ((0) & 63))}
    if true {
        for fi22 := 0; fi22 < 15; fi22++ {
            switch ("foo bar" + "") {
            case "foo bar":
                var v23 int = (fi22 & (fi22 >> ((65535) & 63)))
                _ = v23
            default:
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 string = "\n\t"
                    if false {
                        if (!(42 != 42)) {
                            switch fi24 {
                            case -472:
                                if false {
                                    v25 = v25
                                    if false {
                                        v26 := GenS1{F0: 42, F1: (2 <= (v3 / ((2147483647) | 1))), F2: v4}
                                        if true {
                                            var v27 int16 = int16(fi22)
                                            v28 := GenS1{F0: 100, F1: (true || true), F2: (v2 ^ (v4 - v2))}
                                            v29 := GenS1{F0: -23, F1: (v25 > "a"), F2: (v4 & v4)}
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            _ = len(v1)
                                            v3 -= 4294967295
                                            switch "\n\t" {
                                            case "hello":
                                                v25 = (v25 + v25)
                                            }
                                            v3 += v3
                                        }
                                        _ = v26
                                    }
                                    fmt.Println(v3, v2)
                                    v30 := GenS1{F0: -1, F1: (!true), F2: v4}
                                    fmt.Println(v25, v3, fi24)
                                    if (((((v25 + v25) + ("\n\t" + "a")) + (((v25 + v25) + v25) + "a")) >= v25) || true) {
                                        var v31 int16 = 32767
                                        _ = v31
                                    } else {
                                        v3 = v3
                                        switch "" {
                                        case "":
                                            fmt.Println(v2, fi24, v25)
                                        case "\n\t":
                                            switch fi24 {
                                            case -1:
                                                v32 := []int{(fi22 / (((fi24 >> ((59) & 63))) | 1))}
                                                _ = v32
                                            case 256:
                                                switch v25 {
                                                case "hello":
                                                    switch v25 {
                                                    case "\n\t":
                                                        var v33 uint8 = 61
                                                        _ = v33
                                                    case "":
                                                        if (1e10 < 0.0) {
                                                            var v34 uint64 = uint64(v3)
                                                            _ = v34
                                                        }
                                                    case "a":
                                                        fmt.Println(fi24, v4, v2)
                                                        v35 := GenS2{F0: 0, F1: true}
                                                        fmt.Println(v4, v25)
                                                        v4 += 735
                                                        v36 := GenS2{F0: 42, F1: false}
                                                        _ = v35
                                                        _ = v36
                                                    default:
                                                        var v37 float32 = ((-2.25 * ((0.0 + float32(0.0)) + 3.14159)) + (1e10 / (-2.25 / 1.5)))
                                                        _ = v37
                                                    }
                                                case "foo bar":
                                                    var v38 uint64 = (^(v2))
                                                    _ = v38
                                                }
                                            case 127:
                                                v2 = v2
                                            }
                                            if true {
                                                var v39 bool = (2147483647 < -2147483648)
                                                if (2 != 4294967295) {
                                                    if v39 {
                                                        fmt.Println(fi24, fi22)
                                                        var v40 bool = true
                                                        delete(v1, (fi22 >> ((501) & 63)))
                                                        _ = v40
                                                    }
                                                    if v39 {
                                                        var v41 bool = true
                                                        v3 += 42
                                                        v42 := GenS1{F0: 124, F1: v39, F2: v2}
                                                        var v43 uint32 = v3
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    if v39 {
                                                        v25 = ("hello" + v25)
                                                    }
                                                    v44 := GenS1{F0: 127, F1: (!(!(!v39))), F2: v4}
                                                    _ = v44
                                                } else {
                                                    v2 += (^(v2))
                                                }
                                                v45 := []int{}
                                                _ = v39
                                                _ = v45
                                            } else {
                                                var v46 float32 = 3.14159
                                                _ = v46
                                            }
                                            if false {
                                                var v47 uint64 = 18446744073709551615
                                                v48 := GenS3{F0: (127 != 100), F1: -458}
                                                fmt.Println(fi24, v47, v25)
                                                _ = v47
                                                _ = v48
                                            }
                                            if (v3 == v3) {
                                                v49 := GenS2{F0: -9223372036854775808, F1: (9223372036854775807 <= fi22)}
                                                _ = v49
                                            }
                                            var v50 string = "foo bar"
                                            _ = v50
                                        case "hello":
                                            var v51 int64 = 42
                                            _ = v51
                                        }
                                    }
                                    if false {
                                        switch genFn1(65535, 42, 623) {
                                        case 52:
                                            v25 = (v25 + (v25 + (v25 + (v25 + (v25 + (v25 + v25))))))
                                        case 1:
                                            v52 := []int{fi24, fi24, genFn1(65535, -808, fi24), genFn1(42, 2147483647, 127)}
                                            _ = v52
                                        default:
                                            fmt.Println(v4)
                                        }
                                    }
                                    _ = v30
                                } else {
                                    v53 := GenS2{F0: int64(v3), F1: false}
                                    _ = v53
                                }
                            case 0:
                                if false {
                                    if true {
                                        v54 := GenS4{F0: (v3 / ((4294967295) | 1)), F1: 65535}
                                        v1[fi22] = (fi24 % ((66) | 1))
                                        if (!(((-1 < -1) || false) || true)) {
                                            v3 = 451
                                        }
                                        _ = v54
                                    } else {
                                        v3 += v3
                                        v3 += 2
                                        v2 -= (v4 ^ (v2 & (^(v4))))
                                        v3 = 883
                                    }
                                    var v55 uint32 = v3
                                    v3 -= uint32(v4)
                                    _ = v55
                                } else {
                                    if (81 >= 1) {
                                        if ((1.5 - -2.25) <= 0.0) {
                                            fmt.Println(v25, v4)
                                            fmt.Println(v4)
                                            var v56 uint8 = 75
                                            _ = v56
                                        }
                                    }
                                }
                            case 282:
                                var v57 string = "a"
                                _ = v57
                            default:
                                _ = len(v1)
                            }
                            if true {
                                v58 := GenS5{F0: genFn1(0, 493, (fi22 << ((42) & 63))), F1: 42}
                                _ = v58
                            }
                            v59 := GenS2{F0: 42, F1: false}
                            if (!(genFn1(32767, 1, fi22) >= fi24)) {
                                v1[870] = (fi24 - fi22)
                                var v60 float64 = 0.0
                                v61 := GenS6{F0: fi24, F1: 252}
                                _ = v60
                                _ = v61
                            }
                            _ = v59
                        }
                        v62 := GenS4{F0: (v3 ^ (v3 % ((v3) | 1))), F1: 65535}
                        var v63 int16 = int16(v2)
                        switch (fi22 << ((4294967295) & 63)) {
                        case 1:
                            if false {
                                v63 = 32767
                                if false {
                                    v64 := GenS3{F0: ((v4 & v2) >= 4294967295), F1: v63}
                                    var v65 uint32 = 0
                                    v66 := []int{}
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                } else {
                                    v4 += 4294967295
                                    switch 32768 {
                                    case 32767:
                                        var v67 uint = 42
                                        _ = v67
                                    }
                                }
                                switch (v25 + v25) {
                                case "a":
                                    v68 := []int{fi24, fi22, fi22}
                                    _ = v68
                                default:
                                    var v69 int64 = 42
                                    var v70 uint16 = 0
                                    _ = v69
                                    _ = v70
                                }
                                var v71 int8 = -62
                                var v72 int = (fi22 - (fi24 | (fi24 | 327)))
                                _ = v71
                                _ = v72
                            }
                        case -2147483648:
                            _ = len(v1)
                        }
                        switch fi22 {
                        case 128:
                            var v73 uint16 = 477
                            v3 = (v3 - (v3 | v3))
                            if ((0.0 + float32(-2.25)) == 1e10) {
                                v74 := []int{fi24}
                                var v75 uint = 2
                                _ = v74
                                _ = v75
                            } else {
                                var v76 float64 = (0.1 + float64(3.14159))
                                _ = v76
                            }
                            _ = v73
                        case -9223372036854775808:
                            v2 = (^(v2))
                        case 42:
                            v3 = v3
                        default:
                            var v77 int16 = v63
                            _ = v77
                        }
                        _ = v62
                        _ = v63
                    } else {
                        _ = len(v1)
                        v25 = "foo bar"
                        fmt.Println(v4)
                    }
                    _ = fi24
                    _ = v25
                }
            }
            if false {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        v78 := GenS2{F0: 42, F1: false}
                        _ = v78
                    } else {
                        v4 = (v2 ^ 677)
                    }
                }
                if (1 != -336) {
                    switch "" {
                    case "a":
                        v79 := GenS4{F0: (^(v3)), F1: 723}
                        _ = v79
                    case "hello":
                        if true {
                            if ((false && (0 < 9223372036854775807)) && (-1 >= 0)) {
                                var v80 uint16 = 263
                                for fi81 := 0; fi81 < 1; fi81++ {
                                    fmt.Println(v2, v80, v3)
                                    v82 := GenS7{F0: 65535}
                                    var v83 float32 = (((-2.25 + 0.0) + 0.0) / -2.25)
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                }
                                fmt.Println(v2, v4)
                                _ = v80
                            }
                            v2 = uint64(v3)
                            v2 -= 1
                        } else {
                            var v84 uint32 = (^(v3))
                            switch "foo bar" {
                            case "hello":
                                if (!((128 == 128) || (((!(445 != -1)) && (v4 == 1)) || (false && false)))) {
                                    v85 := []int{}
                                    if (v2 != v2) {
                                        v4 -= uint64(v3)
                                    } else {
                                        if false {
                                            v3 += v84
                                            if ((fi22 > fi22) && (!true)) {
                                                v86 := GenS4{F0: (v84 + (v84 << ((42) & 63))), F1: 65535}
                                                v3 += (v3 * (v3 % ((v3) | 1)))
                                                _ = v86
                                            } else {
                                                for fi87 := 0; fi87 < 17; fi87++ {
                                                    if len(v85) > 0 {
                                                        v85[0] = (fi87 / (((fi87 ^ 4294967295)) | 1))
                                                    }
                                                    _ = fi87
                                                }
                                            }
                                        }
                                        fmt.Println(v2, v2, v2)
                                    }
                                    _ = v85
                                } else {
                                    var v88 uint8 = 42
                                    switch fi22 {
                                    case 42:
                                        v88 = 1
                                    case 135:
                                        v3 = (^(v84))
                                    case -9223372036854775808:
                                        if (!((!false) && (1 >= 0))) {
                                            v84 = (v84 << ((685) & 63))
                                            v4 += (v2 >> ((136) & 63))
                                        }
                                    default:
                                        if ((-1 < -117) && ((float32((-2.25 - 1e10)) + -2.25) == 0.1)) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v84 -= 4294967295
                                            }
                                            v84 = v84
                                            v89 := v1[(fi22 / ((genFn1(283, 42, (fi22 >> ((18446744073709551615) & 63)))) | 1))]
                                            _ = v89
                                        }
                                    }
                                    v4 = v4
                                    _ = v88
                                }
                            case "foo bar":
                                v90 := []int{908, genFn1(202, 1, (fi22 | -988)), (^(fi22)), (fi22 / ((fi22) | 1))}
                                _ = v90
                            }
                            v84 -= v3
                            _ = v84
                        }
                    default:
                        _ = len(v1)
                    }
                }
                if true {
                    if false {
                        var v91 float64 = 3.14159
                        _ = v91
                    }
                    fmt.Println(v2, fi22)
                    if true {
                        v92 := []int{fi22}
                        delete(v1, (fi22 % ((genFn1(uint16(v4), 1, (fi22 * genFn1(73, 454, -571)))) | 1)))
                        for fi93 := 0; fi93 < 14; fi93++ {
                            v94 := GenS8{F0: "hello", F1: 769, F2: 0}
                            v4 = 806
                            v95 := v1[int(v2)]
                            _ = fi93
                            _ = v94
                            _ = v95
                        }
                        _ = v92
                    }
                } else {
                    var v96 bool = false
                    _ = v96
                }
            }
            var v97 uint = 18446744073709551615
            _ = fi22
            _ = v97
        }
        for fi98 := 0; fi98 < 20; fi98++ {
            fmt.Println(v3, v2)
            _ = fi98
        }
    } else {
        v99 := GenS9{F0: 1}
        var v100 float32 = (1.5 - (float32(0.1) * (float32(1e10) / 1e10)))
        v100 -= float32(3.14159)
        _ = v100
        _ = v99
    }
    if true {
        for fi101 := 0; fi101 < 2; fi101++ {
            v2 = v2
            _ = fi101
        }
        if false {
            v102 := GenS10{F0: v3}
            _ = v102
        }
        var v103 int32 = -1
        _ = v103
    } else {
        v4 -= v4
    }
    if (((v3 >> ((uint(v2)) & 63)) < (^(v3))) || false) {
        fmt.Println(v2, v2)
    }
    var v104 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v104
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
