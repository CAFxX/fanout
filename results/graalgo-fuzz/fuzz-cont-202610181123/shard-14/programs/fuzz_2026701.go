package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 int }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 bool; F2 int16 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8; F1 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 bool; F1 bool }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 string; F1 uint64; F2 uint16 }
type GenS18 struct { F0 uint64; F1 string; F2 int64 }
func genFn0() int8 {
    var v6 bool = (true || (255 > 65535))
    if true {
        var v7 int64 = -915
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int64 = v7
            for fi10 := 0; fi10 < 2; fi10++ {
                switch fi8 {
                case 32768:
                    v9 -= (v7 / (((v7 % ((v9) | 1))) | 1))
                    v11 := []int{(fi8 << ((582) & 63)), 720, -1, (fi8 | fi10)}
                    _ = v11
                case 127:
                    var v12 int64 = v7
                    _ = v12
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v7, v6)
        _ = v7
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v6 = (0.0 >= (0.0 + (float64(float32(((-2.25 / ((float64(1e10) * float64(0.0)) + (1.5 - 0.1))) - float64(1e10)))) - 1e10)))
        _ = fi13
    }
    v14 := []int{-9223372036854775808, -2147483648, -9223372036854775808, -1}
    if (-9223372036854775808 < -943) {
        v6 = (-68 != 1)
        if v6 {
            for fi15 := 0; fi15 < 2; fi15++ {
                v6 = true
                var v16 int = (fi15 | fi15)
                for fi17 := 0; fi17 < 18; fi17++ {
                    if len(v14) > 0 {
                        v14[0] = v16
                    }
                    v18 := GenS1{F0: 42, F1: 18446744073709551615}
                    fmt.Println(fi15)
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
        var v19 int16 = 0
        _ = v19
    } else {
        if true {
            v20 := GenS1{F0: 0, F1: 4294967295}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v20
        } else {
            switch 32768 {
            case 127:
                if len(v14) > 0 {
                    v14[0] = 1
                }
            case -9223372036854775808:
                var v21 int64 = 42
                _ = v21
            default:
                for fi22 := 0; fi22 < 17; fi22++ {
                    v6 = v6
                    var v23 int8 = -20
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    _ = v14
    _ = v6
    return 100
}
func genFn1(p0 int16, p1 int32) bool {
    v24 := GenS1{F0: 352, F1: 2}
    if (852 != 61) {
        for fi25 := 0; fi25 < 16; fi25++ {
            for fi26 := 0; fi26 < 5; fi26++ {
                if (!false) {
                    p0 += p0
                }
                if (true || true) {
                    v27 := []int{(-(fi26)), (fi25 | -2147483648), (fi26 >> ((462) & 63)), fi25}
                    switch (-(fi26)) {
                    case -23:
                        v28 := GenS1{F0: 703, F1: 1}
                        _ = v28
                    case -10:
                        if ((2 < 4294967295) || false) {
                            p1 = p1
                            v29 := GenS2{F0: (p1 & (-(p1))), F1: 4294967295, F2: 341}
                            v30 := GenS1{F0: 2, F1: 0}
                            _ = v29
                            _ = v30
                        } else {
                            switch fi25 {
                            case -571:
                                if false {
                                    var v31 float32 = 0.1
                                    var v32 int16 = -1
                                    switch "a" {
                                    case "foo bar":
                                        if (0 == 4294967295) {
                                            v27 = append(v27, fi26)
                                            if (false && false) {
                                                var v33 float32 = 0.0
                                                var v34 string = "foo bar"
                                                if false {
                                                    if true {
                                                        p0 = v32
                                                    } else {
                                                        if true {
                                                            fmt.Println(v31)
                                                            var v35 uint32 = 65535
                                                            fmt.Println(v31, p0, v32)
                                                            _ = v35
                                                        }
                                                        if false {
                                                            p0 = (-(v32))
                                                            if (!(0 >= -1)) {
                                                                p1 += p1
                                                            }
                                                            v36 := GenS3{F0: 18446744073709551615, F1: 1, F2: 1}
                                                            _ = v36
                                                        } else {
                                                            var v37 int = 1
                                                            var v38 uint8 = 128
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        v39 := GenS4{F0: 2147483647, F1: (p0 % (((v32 * 42)) | 1)), F2: 0}
                                                        _ = v39
                                                    }
                                                    v27 = append(v27, fi26)
                                                    var v40 uint = 0
                                                    v41 := GenS5{F0: false, F1: v40}
                                                    fmt.Println(v32, p0, p0)
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    var v42 float32 = v31
                                                    p0 = 16
                                                    _ = v42
                                                }
                                                if (!false) {
                                                    v27 = append(v27, (^(fi25)))
                                                }
                                                v34 = "hello"
                                                _ = v33
                                                _ = v34
                                            } else {
                                                fmt.Println(fi25, v31)
                                            }
                                            var v43 uint16 = 441
                                            _ = v43
                                        } else {
                                            v31 = (v31 - v31)
                                        }
                                    default:
                                        v44 := GenS4{F0: 426, F1: p0, F2: 0}
                                        _ = v44
                                    }
                                    _ = v31
                                    _ = v32
                                } else {
                                    fmt.Println(fi25, p0)
                                    var v45 int64 = 9223372036854775807
                                    var v46 int64 = 42
                                    _ = v45
                                    _ = v46
                                }
                            }
                            var v47 uint16 = 255
                            fmt.Println(p0, v47, p1)
                            if true {
                                fmt.Println(p1, v47)
                                if (93 == 116) {
                                    v27 = append(v27, -591)
                                    v48 := GenS6{F0: false}
                                    v47 += 235
                                    _ = v48
                                }
                                var v49 string = ("a" + "foo bar")
                                _ = v49
                            } else {
                                v27 = append(v27, fi25)
                                var v50 int8 = -128
                                _ = v50
                            }
                            _ = v47
                        }
                    default:
                        if len(v27) > 0 {
                            v27[0] = (fi25 * (fi25 ^ (fi25 << ((42) & 63))))
                        }
                    }
                    fmt.Println(fi26, fi26)
                    _ = v27
                } else {
                    var v51 uint16 = 42
                    _ = v51
                }
                _ = fi26
            }
            p0 = p0
            if false {
                fmt.Println(p1, fi25)
                v52 := []int{(^(fi25)), (fi25 & 32767)}
                fmt.Println(fi25, p1, p1)
                _ = v52
            } else {
                v53 := GenS7{F0: "\n\t"}
                if (-1 <= genFn0()) {
                    var v54 float32 = 1.5
                    v55 := GenS8{F0: 255}
                    var v56 int = fi25
                    _ = v54
                    _ = v55
                    _ = v56
                }
                for fi57 := 0; fi57 < 11; fi57++ {
                    if (102 < 2) {
                        var v58 int = fi57
                        var v59 uint8 = 2
                        v60 := []int{v58, (fi25 * fi25), int(p0), (v58 / ((fi25) | 1))}
                        if (!false) {
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                fmt.Println(fi57, v59, v58)
                            }
                            if len(v60) > 0 {
                                v60[0] = 32767
                            }
                        }
                        switch fi57 {
                        case 1:
                            switch (("foo bar" + ("\n\t" + "a")) + "hello") {
                            case "a":
                                fmt.Println(v58, fi57, v59)
                                v61 := GenS9{F0: (v59 + v59), F1: (false || (!false)), F2: (p0 & -169)}
                                v58 -= fi25
                                _ = v61
                            case "":
                                if (128 < (^(v59))) {
                                    p1 += -336
                                    fmt.Println(fi57, v58)
                                    var v62 bool = false
                                    _ = v62
                                }
                            case "\n\t":
                                fmt.Println(fi25)
                            }
                        case 2147483647:
                            switch (fi57 | (fi25 / (((fi57 ^ fi57)) | 1))) {
                            case -619:
                                v58 += (v58 - fi57)
                            }
                        default:
                            v63 := GenS9{F0: v59, F1: (!false), F2: (p0 + p0)}
                            v64 := GenS9{F0: v59, F1: (935 > 946), F2: p0}
                            _ = len(v60)
                            switch "" {
                            case "a":
                                if (p1 == p1) {
                                    p1 -= (p1 & p1)
                                    var v65 string = ("foo bar" + "")
                                    var v66 int32 = 2147483647
                                    fmt.Println(fi57, v59)
                                    v58 = (v58 / (((fi57 << ((4294967295) & 63))) | 1))
                                    var v67 uint16 = 255
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                } else {
                                    v60 = append(v60, fi57)
                                    v58 = (fi57 - (fi57 + v58))
                                    v58 = (v58 - fi57)
                                }
                            case "foo bar":
                                v68 := GenS10{F0: genFn0(), F1: true}
                                _ = v68
                            }
                            _ = v63
                            _ = v64
                        }
                        _ = v58
                        _ = v59
                        _ = v60
                    } else {
                        p0 += p0
                        var v69 string = "hello"
                        var v70 int64 = 42
                        _ = v69
                        _ = v70
                    }
                    var v71 uint32 = 207
                    _ = fi57
                    _ = v71
                }
                _ = v53
            }
            _ = fi25
        }
    } else {
        v72 := GenS11{F0: 127}
        v73 := GenS8{F0: 188}
        for fi74 := 0; fi74 < 9; fi74++ {
            for fi75 := 0; fi75 < 16; fi75++ {
                var v76 uint64 = 1
                v77 := []int{(fi74 / ((fi75) | 1)), -551, (fi75 & 4294967295), fi74}
                v78 := GenS12{F0: 42, F1: false}
                v77 = append(v77, fi75)
                _ = fi75
                _ = v76
                _ = v77
                _ = v78
            }
            _ = fi74
        }
        _ = v72
        _ = v73
    }
    p0 = (p0 & 1000)
    var v79 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v24
    _ = v79
    return (-173 >= -335)
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{}
    var v3 int32 = -583
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v80 int8 = genFn0()
    for fi81 := 0; fi81 < 17; fi81++ {
        v3 = -973
        switch (fi81 & (fi81 | 32767)) {
        case 662:
            v82 := v5[4294967295]
            _ = v82
        case 128:
            v83 := GenS13{F0: 2147483648}
            _ = v83
        }
        v84 := GenS14{F0: "hello"}
        _ = fi81
        _ = v84
    }
    var v85 uint64 = 804
    if true {
        if (("\n\t" + "") < "hello") {
            var v86 float64 = 0.0
            switch -908 {
            case 1:
                v87 := GenS15{F0: ((!true) || true), F1: (v80 < v80)}
                _ = v87
            case 2147483647:
                switch 2147483647 {
                case -207:
                    if false {
                        v88 := GenS16{F0: (v85 << ((65535) & 63))}
                        _ = v88
                    }
                case -1:
                    switch "a" {
                    case "a":
                        if ((v3 + 42) == (v3 | (v3 | (v3 - int32(v80))))) {
                            switch ((((("a" + "a") + "\n\t") + ("a" + ("" + "\n\t"))) + "\n\t") + ("foo bar" + "")) {
                            case "foo bar":
                                var v89 int8 = int8(v85)
                                _ = v89
                            case "a":
                                v2 = append(v2, "")
                            case "":
                                v90 := GenS11{F0: -2147483648}
                                _ = v90
                            }
                            fmt.Println(v80)
                        }
                    case "foo bar":
                        v86 = (float64(-2.25) - 1e10)
                    case "\n\t":
                        for fi91 := 0; fi91 < 1; fi91++ {
                            var v92 int64 = (v1 & -723)
                            v93 := GenS14{F0: "\n\t"}
                            if (42 >= 479) {
                                _ = len(v2)
                                for fi94 := 0; fi94 < 14; fi94++ {
                                    v86 = v86
                                    v3 += (v3 * v3)
                                    fmt.Println(fi94, v80, v85)
                                    _ = fi94
                                }
                                switch "" {
                                case "hello":
                                    switch ("hello" + ("hello" + "a")) {
                                    case "foo bar":
                                        var v95 int = int(v92)
                                        _ = v95
                                    case "hello":
                                        fmt.Println(v86)
                                    case "\n\t":
                                        v92 += (v92 ^ v1)
                                    default:
                                        for fi96 := 0; fi96 < 11; fi96++ {
                                            v1 += v1
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                            _ = fi96
                                        }
                                    }
                                case "a":
                                    fmt.Println(v92, v1, v1)
                                case "foo bar":
                                    if false {
                                        v1 = (^(v92))
                                        fmt.Println(v92, v3, v92)
                                    } else {
                                        var v97 string = "foo bar"
                                        _ = v97
                                    }
                                }
                            }
                            _ = fi91
                            _ = v92
                            _ = v93
                        }
                    default:
                        for fi98 := 0; fi98 < 20; fi98++ {
                            fmt.Println(v3, v86, v3)
                            v1 += 9223372036854775807
                            var v99 float32 = ((0.0 + -2.25) / 3.14159)
                            _ = fi98
                            _ = v99
                        }
                    }
                case 1:
                    fmt.Println(v3, v80)
                }
            }
            _ = v86
        }
        for fi100 := 0; fi100 < 6; fi100++ {
            v101 := GenS15{F0: false, F1: true}
            if true {
                var v102 uint = 18446744073709551615
                _ = v102
            } else {
                var v103 float32 = 3.14159
                v3 -= int32(fi100)
                var v104 string = ("a" + "foo bar")
                _ = v103
                _ = v104
            }
            _ = fi100
            _ = v101
        }
    } else {
        var v105 uint32 = 674
        _ = len(v2)
        _ = v105
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    delete(v5, 2)
    v106 := GenS8{F0: 1}
    v3 -= v3
    if true {
        v3 -= 2147483647
        var v107 int16 = -384
        if ((v80 & (-(v80))) > (^(v80))) {
            var v108 float64 = float64(1.5)
            v5[100] = 4294967295
            switch 9223372036854775807 {
            case -2147483648:
                v109 := GenS17{F0: "foo bar", F1: (v85 - v85), F2: 0}
                _ = v109
            default:
                var v110 uint32 = uint32(v85)
                _ = v110
            }
            _ = v108
        }
        _ = v107
    }
    if false {
        for fi111 := 0; fi111 < 12; fi111++ {
            _ = len(v2)
            _ = fi111
        }
        delete(v4, -9223372036854775808)
        switch -480 {
        case -1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        var v112 int8 = 1
        var v113 uint32 = 2147483647
        v114 := GenS18{F0: (v85 % (((v85 | (^(v85)))) | 1)), F1: "hello", F2: 42}
        _ = v112
        _ = v113
        _ = v114
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v85
}
