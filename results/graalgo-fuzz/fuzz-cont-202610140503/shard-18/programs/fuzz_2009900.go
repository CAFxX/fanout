package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
type GenS7 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 string, p1 int, p2 bool) uint64 {
    var v4 int32 = 674
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 string = "\n\t"
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p2, fi7)
        _ = fi7
    }
    if p2 {
        switch ("a" + p0) {
        case "a":
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 float64 = 3.14159
                if (!p2) {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v11 := []int{128, (p1 + 42), (fi10 & (-(fi8)))}
                        _ = fi10
                        _ = v11
                    }
                    v12 := []int{p1}
                    if p2 {
                        v12 = append(v12, int(v4))
                    } else {
                        var v13 float32 = 0.1
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v15 := GenS1{F0: -682, F1: (!p2)}
                            v4 = v4
                            if len(v12) > 0 {
                                v12[0] = (p1 - -9223372036854775808)
                            }
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 5; fi16++ {
                            fmt.Println(v9, v4, p1)
                            fmt.Println(fi8)
                            _ = fi16
                        }
                        _ = v13
                    }
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
        case "foo bar":
            var v17 string = "foo bar"
            _ = v17
        case "\n\t":
            v18 := GenS2{F0: 78, F1: (v4 ^ v4), F2: -1}
            _ = v18
        default:
            fmt.Println(p2)
        }
        fmt.Println(v4, p1)
        if true {
            var v19 float64 = (3.14159 - 3.14159)
            fmt.Println(v19)
            _ = v19
        } else {
            v4 -= v4
            switch -332 {
            case -2147483648:
                v20 := []int{}
                _ = v20
            case 542:
                p0 = ("\n\t" + p0)
            case -1000:
                for fi21 := 0; fi21 < 15; fi21++ {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        p0 = ("\n\t" + p0)
                        switch p0 {
                        case "":
                            switch (-(p1)) {
                            case 32768:
                                v23 := []int{(fi21 ^ p1), (^(fi21)), fi22}
                                _ = v23
                            case 2147483647:
                                if (589 == -344) {
                                    v4 += v4
                                    p0 = ("foo bar" + p0)
                                }
                            case 4294967295:
                                switch p0 {
                                case "\n\t":
                                    v24 := []int{(^(p1)), fi21}
                                    _ = v24
                                case "":
                                    var v25 uint64 = 42
                                    _ = v25
                                case "foo bar":
                                    v26 := []int{(p1 << ((2) & 63))}
                                    _ = v26
                                }
                            default:
                                var v27 string = p0
                                _ = v27
                            }
                        case "\n\t":
                            switch p1 {
                            case 117:
                                v28 := []int{fi22, 127}
                                switch (p1 / ((p1) | 1)) {
                                case 128:
                                    switch (p0 + ((p0 + "a") + "")) {
                                    case "hello":
                                        v29 := GenS3{F0: 1}
                                        _ = v29
                                    }
                                default:
                                    p0 = ((p0 + p0) + (("a" + p0) + ""))
                                }
                                _ = v28
                            default:
                                switch (p1 * fi21) {
                                case 0:
                                    v30 := []int{fi22, 42, fi21, fi21}
                                    _ = v30
                                case -9223372036854775808:
                                    v31 := []int{(^(fi22)), (fi21 >> ((784) & 63)), -388, p1}
                                    _ = v31
                                case -281:
                                    v32 := GenS1{F0: 42, F1: p2}
                                    _ = v32
                                }
                            }
                            v33 := []int{fi22, (fi21 * fi21), fi22, 361}
                            _ = v33
                        default:
                            v34 := GenS2{F0: 4294967295, F1: (v4 + v4), F2: 0}
                            _ = v34
                        }
                        switch 42 {
                        case 9223372036854775807:
                            v35 := []int{-1, (^(p1)), p1}
                            _ = v35
                        case 101:
                            v36 := GenS1{F0: 0, F1: (((3.14159 - 3.14159) / 3.14159) < 1e10)}
                            _ = v36
                        case 0:
                            if true {
                                var v37 int32 = (v4 << ((42) & 63))
                                if (0 <= 827) {
                                    v38 := []int{(-(p1)), 1, p1}
                                    switch (p0 + ("foo bar" + p0)) {
                                    case "a":
                                        switch (fi21 - (fi22 + (fi21 | (^(p1))))) {
                                        case -2147483648:
                                            var v39 uint = 42
                                            _ = v39
                                        }
                                    case "\n\t":
                                        if true {
                                            v40 := GenS3{F0: 2}
                                            var v41 float64 = 3.14159
                                            switch 128 {
                                            case 9223372036854775807:
                                                fmt.Println(v4, v37, fi22)
                                            }
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        var v42 float64 = -2.25
                                        _ = v42
                                    }
                                    fmt.Println(p2, fi21)
                                    if false {
                                        fmt.Println(fi22, fi22, v4)
                                    } else {
                                        v37 += (^(v37))
                                    }
                                    _ = v38
                                } else {
                                    switch fi21 {
                                    case 1:
                                        if false {
                                            if ("a" <= (p0 + "\n\t")) {
                                                switch p1 {
                                                case 127:
                                                    var v43 string = p0
                                                    _ = v43
                                                }
                                                var v44 int = fi21
                                                v4 = (-(v4))
                                                _ = v44
                                            } else {
                                                v45 := GenS2{F0: 818, F1: 487, F2: 9}
                                                v46 := []int{}
                                                if (138 != 101) {
                                                    v47 := GenS2{F0: 13, F1: (v37 | (v37 & (v4 * (v37 ^ 63)))), F2: -128}
                                                    fmt.Println(v37, fi22, v37)
                                                    _ = v47
                                                }
                                                _ = v45
                                                _ = v46
                                            }
                                        } else {
                                            var v48 uint16 = 966
                                            switch (p1 % ((fi22) | 1)) {
                                            case -2147483648:
                                                var v49 float32 = 0.0
                                                _ = v49
                                            case 4294967295:
                                                var v50 uint16 = v48
                                                v51 := []int{(p1 >> ((4294967295) & 63))}
                                                switch 128 {
                                                case -983:
                                                    if (62 < 127) {
                                                        v52 := GenS1{F0: 42, F1: (!false)}
                                                        v53 := GenS1{F0: 166, F1: false}
                                                        fmt.Println(fi21, p2)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                case -204:
                                                    v54 := GenS3{F0: 1}
                                                    _ = v54
                                                case -9223372036854775808:
                                                    fmt.Println(p0, fi22, p0)
                                                default:
                                                    var v55 float32 = 3.14159
                                                    _ = v55
                                                }
                                                _ = v50
                                                _ = v51
                                            case 104:
                                                v56 := []int{(fi21 - p1), (fi22 % (((fi21 << ((944) & 63))) | 1)), fi21, (p1 - (-(fi21)))}
                                                _ = v56
                                            default:
                                                var v57 int8 = 127
                                                _ = v57
                                            }
                                            _ = v48
                                        }
                                    }
                                    var v58 int16 = 32767
                                    _ = v58
                                }
                                _ = v37
                            }
                        }
                        _ = fi22
                    }
                    fmt.Println(p1)
                    _ = fi21
                }
            default:
                v59 := GenS1{F0: -222, F1: false}
                _ = v59
            }
            fmt.Println(p1, p0, v4)
        }
    }
    v60 := []int{p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v60
    return 0
}
func genFn1() int64 {
    var v61 uint16 = 255
    if true {
        var v62 uint8 = 2
        var v63 int16 = -32768
        _ = v62
        _ = v63
    }
    var v64 float64 = 1.5
    v61 += 320
    v65 := GenS4{F0: -9223372036854775808, F1: true, F2: (2147483648 != 2)}
    _ = v61
    _ = v64
    _ = v65
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("" + ""), ""}
    var v3 uint64 = 1
    for fi66 := 0; fi66 < 10; fi66++ {
        fmt.Println(fi66)
        fmt.Println(fi66, fi66)
        if false {
            v1[127] = (fi66 + 1)
            _ = len(v2)
            v67 := v1[(fi66 / (((fi66 << ((42) & 63))) | 1))]
            _ = v67
        }
        _ = fi66
    }
    switch 563 {
    case -1:
        fmt.Println(v3)
    case 32767:
        if (-32768 > -622) {
            var v68 int64 = -1
            switch 1 {
            case 0:
                switch 1 {
                case -9223372036854775808:
                    var v69 uint16 = 32767
                    _ = v69
                }
            case -2147483648:
                var v70 float64 = (3.14159 * (0.0 + 0.0))
                _ = v70
            }
            var v71 string = ("hello" + "")
            _ = v68
            _ = v71
        }
    }
    v3 += (v3 + (v3 - v3))
    if true {
        fmt.Println(v3)
        v1[32767] = -1
        if (true || (!true)) {
            fmt.Println(v3)
        } else {
            v3 -= v3
            var v72 int32 = 0
            v72 -= v72
            _ = v72
        }
    }
    v73 := GenS5{F0: 0, F1: -1}
    for fi74 := 0; fi74 < 14; fi74++ {
        if ((float64(1e10) / 3.14159) == 3.14159) {
            var v75 bool = false
            for fi76 := 0; fi76 < 12; fi76++ {
                v77 := GenS2{F0: 0, F1: 1, F2: 0}
                _ = fi76
                _ = v77
            }
            _ = v75
        } else {
            var v78 float32 = -2.25
            if false {
                v3 = (^(v3))
                if ((!(!(!true))) && true) {
                    v79 := GenS5{F0: -26, F1: 407}
                    _ = v79
                } else {
                    if true {
                        fmt.Println(fi74, fi74)
                        fmt.Println(v78, fi74)
                    }
                    if (1 > 790) {
                        for fi80 := 0; fi80 < 4; fi80++ {
                            var v81 float32 = (v78 * v78)
                            _ = fi80
                            _ = v81
                        }
                    }
                    if false {
                        for fi82 := 0; fi82 < 4; fi82++ {
                            fmt.Println(v3)
                            var v83 int16 = 1
                            _ = fi82
                            _ = v83
                        }
                        var v84 float32 = v78
                        _ = v84
                    }
                }
            }
            _ = v78
        }
        if (1e10 != 1e10) {
            v1[-1] = fi74
        } else {
            for fi85 := 0; fi85 < 5; fi85++ {
                switch "" {
                case "hello":
                    var v86 uint32 = 98
                    _ = v86
                default:
                    if (((65535 != 188) && true) || (0.0 != (0.1 * (3.14159 + -2.25)))) {
                        if true {
                            var v87 int = fi85
                            if len(v2) > 0 {
                                v2[0] = ("a" + "foo bar")
                            }
                            v88 := GenS4{F0: 42, F1: (65535 >= 415), F2: true}
                            _ = v87
                            _ = v88
                        }
                    } else {
                        v89 := GenS5{F0: -106, F1: 908}
                        var v90 int16 = 0
                        var v91 string = "foo bar"
                        _ = v89
                        _ = v90
                        _ = v91
                    }
                }
                var v92 int = fi85
                _ = fi85
                _ = v92
            }
        }
        if (-1 >= 722) {
            v3 = v3
            var v93 uint8 = 42
            var v94 float32 = -2.25
            _ = v93
            _ = v94
        } else {
            for fi95 := 0; fi95 < 12; fi95++ {
                v96 := GenS1{F0: -32768, F1: true}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v97 := GenS6{F0: "hello", F1: 42, F2: v3}
                v3 = (v3 | (v3 | (^(v3))))
                _ = fi95
                _ = v96
                _ = v97
            }
            switch "" {
            case "a":
                v1[0] = (fi74 & (fi74 * fi74))
            case "hello":
                var v98 int64 = -1
                _ = v98
            case "\n\t":
                fmt.Println(v3)
            }
        }
        _ = fi74
    }
    v3 += (v3 << ((2) & 63))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    switch 2147483647 {
    case 42:
        _ = len(v1)
    case -1:
        v99 := GenS7{F0: 32768, F1: "", F2: 42}
        _ = v99
    default:
        fmt.Println(v3)
    }
    v3 -= (v3 % (((^(v3))) | 1))
    if (42 == 42) {
        v3 = 852
        v100 := GenS4{F0: 277, F1: false, F2: true}
        v101 := GenS5{F0: 42, F1: -1}
        _ = v100
        _ = v101
    } else {
        v102 := v1[-658]
        v103 := v1[v102]
        _ = v102
        _ = v103
    }
    var v104 float32 = 1e10
    var v105 int32 = -2147483648
    v106 := GenS5{F0: -1, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v73
}
