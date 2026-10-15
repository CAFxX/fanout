package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint32; F1 bool; F2 string }
type GenS6 struct { F0 uint; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 uint32; F1 int32; F2 bool }
type GenS10 struct { F0 int; F1 uint16; F2 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    if true {
        var v4 int = 32767
        fmt.Println(v4)
        _ = v4
    } else {
        var v5 int = 2147483647
        switch (-(v5)) {
        case -9223372036854775808:
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS1{F0: (-2.25 < 3.14159)}
                _ = fi6
                _ = v7
            }
        case -457:
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint = 4294967295
                if false {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v11 := []int{(v5 + fi8)}
                        _ = fi10
                        _ = v11
                    }
                }
                v12 := []int{v5, v5}
                _ = fi8
                _ = v12
                _ = v9
            }
        case -2147483648:
            var v13 int64 = -195
            _ = v13
        default:
            if false {
                var v14 uint16 = 2
                _ = v14
            }
        }
        _ = v5
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    switch 561 {
    case 128:
        switch "foo bar" {
        case "foo bar":
            v16 := []int{727, -587}
            _ = v16
        default:
            v17 := GenS2{F0: (32768 != 4294967295), F1: 2147483647, F2: 1}
            _ = v17
        }
    case 2147483647:
        for fi18 := 0; fi18 < 7; fi18++ {
            switch "a" {
            case "a":
                var v19 int = (fi18 * fi18)
                _ = v19
            case "\n\t":
                var v20 float32 = 1.5
                _ = v20
            default:
                var v21 int = fi18
                _ = v21
            }
            switch ("" + "foo bar") {
            case "hello":
                var v22 uint32 = 764
                _ = v22
            case "a":
                switch fi18 {
                case -819:
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if (453 < 2147483647) {
                            var v24 int8 = 0
                            var v25 string = "foo bar"
                            var v26 int16 = -914
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        var v27 uint16 = 936
                        v28 := GenS3{F0: v27, F1: fi18, F2: 1}
                        var v29 uint16 = v27
                        _ = fi23
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                case 4294967295:
                    v30 := GenS4{F0: 42, F1: 2}
                    _ = v30
                case 32767:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        if (-1 > -1) {
                            switch "a" {
                            case "foo bar":
                                v32 := GenS5{F0: 2, F1: false, F2: "hello"}
                                _ = v32
                            case "a":
                                v33 := []int{(-(fi31)), (fi18 / ((fi18) | 1)), -359}
                                _ = v33
                            case "hello":
                                v34 := []int{9223372036854775807, fi18, fi18, -490}
                                _ = v34
                            default:
                                if false {
                                    var v35 string = ""
                                    v36 := []int{fi18}
                                    _ = v35
                                    _ = v36
                                }
                            }
                            fmt.Println(fi18, fi18)
                            v37 := GenS5{F0: 841, F1: (2147483647 <= 559), F2: "hello"}
                            _ = v37
                        }
                        if (true || false) {
                            v38 := GenS2{F0: true, F1: 2147483647, F2: 127}
                            if true {
                                var v39 int16 = 0
                                _ = v39
                            }
                            _ = v38
                        } else {
                            if (734 == 2) {
                                var v40 int64 = 9223372036854775807
                                switch "hello" {
                                case "foo bar":
                                    v41 := []int{}
                                    _ = v41
                                }
                                v42 := GenS4{F0: 4294967295, F1: 919}
                                _ = v40
                                _ = v42
                            } else {
                                var v43 uint8 = 0
                                v44 := GenS3{F0: 620, F1: fi18, F2: 856}
                                v45 := GenS6{F0: 255, F1: -695, F2: v43}
                                v46 := []int{(fi31 & (fi18 ^ 1)), fi31, (fi18 / (((fi18 + (^(fi18)))) | 1)), (fi18 - fi31)}
                                _ = v43
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        }
                        fmt.Println(fi31)
                        _ = fi31
                    }
                }
            case "foo bar":
                for fi47 := 0; fi47 < 19; fi47++ {
                    if (!false) {
                        var v48 uint16 = 42
                        _ = v48
                    }
                    var v49 bool = true
                    _ = fi47
                    _ = v49
                }
            }
            _ = fi18
        }
    case -2147483648:
        var v50 int8 = 1
        _ = v50
    }
    v51 := GenS3{F0: 255, F1: 800, F2: 42}
    var v52 int32 = 42
    _ = v51
    _ = v52
    return 4294967295
}
func genFn1(p0 uint8) int8 {
    v53 := GenS7{F0: 0, F1: 42, F2: 1}
    p0 = p0
    if true {
        var v54 int64 = -9223372036854775808
        p0 += (p0 % ((p0) | 1))
        _ = v54
    }
    p0 += p0
    _ = p0
    _ = v53
    return 55
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 int8 = -41
    v55 := []int{584}
    for fi56 := 0; fi56 < 17; fi56++ {
        fmt.Println(fi56, fi56, v2)
        v3 = (v3 ^ (v3 / ((v3) | 1)))
        var v57 uint8 = 0
        _ = fi56
        _ = v57
    }
    fmt.Println(v2)
    if len(v55) >= 2 {
        v55 = v55[1:len(v55)]
    }
    for fi58 := 0; fi58 < 13; fi58++ {
        switch (("" + ("" + "")) + "a") {
        case "\n\t":
            v2 = true
        case "":
            if false {
                for fi59 := 0; fi59 < 2; fi59++ {
                    if false {
                        var v60 float32 = (0.0 - -2.25)
                        _ = v60
                    } else {
                        v1["a"] = 340
                        var v61 int8 = v3
                        fmt.Println(v2, fi58, fi59)
                        _ = v61
                    }
                    if v2 {
                        fmt.Println(fi58, v2)
                    }
                    var v62 uint8 = 1
                    _ = fi59
                    _ = v62
                }
            } else {
                var v63 int16 = int16(fi58)
                var v64 uint = 2
                _ = v63
                _ = v64
            }
        default:
            v65 := GenS8{F0: (v3 + (^(v3))), F1: 255}
            _ = v65
        }
        if (!true) {
            var v66 bool = false
            var v67 uint16 = 2
            var v68 uint = 255
            _ = v66
            _ = v67
            _ = v68
        }
        _ = fi58
    }
    v2 = (true && v2)
    switch 4294967295 {
    case -2147483648:
        var v69 float32 = 3.14159
        _ = v69
    case 9223372036854775807:
        v2 = false
    case -1:
        for fi70 := 0; fi70 < 2; fi70++ {
            for fi71 := 0; fi71 < 2; fi71++ {
                var v72 int = -9223372036854775808
                switch (-(fi70)) {
                case 9223372036854775807:
                    switch 127 {
                    case 42:
                        switch 218 {
                        case -1:
                            fmt.Println(v72, v72, fi71)
                        case 128:
                            var v73 int32 = 0
                            _ = v73
                        }
                    case 128:
                        if v2 {
                            var v74 uint64 = 531
                            _ = v74
                        } else {
                            var v75 uint16 = 32767
                            _ = v75
                        }
                    default:
                        v72 -= fi71
                    }
                case -382:
                    switch -9223372036854775808 {
                    case -692:
                        v76 := GenS8{F0: (v3 >> ((0) & 63)), F1: 255}
                        _ = v76
                    case -2147483648:
                        _ = len(v1)
                    case 128:
                        var v77 float32 = 1e10
                        _ = v77
                    }
                default:
                    fmt.Println(v3, v3)
                }
                _ = fi71
                _ = v72
            }
            v78 := GenS5{F0: 453, F1: false, F2: (("a" + ("foo bar" + "\n\t")) + ("hello" + (("foo bar" + "\n\t") + "a")))}
            _ = fi70
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 8; fi79++ {
        if true {
            for fi80 := 0; fi80 < 13; fi80++ {
                fmt.Println(v3, v3, fi79)
                if v2 {
                    var v81 int64 = 0
                    _ = v81
                }
                v55 = append(v55, (fi79 | -9223372036854775808))
                _ = fi80
            }
            _ = len(v1)
            if true {
                switch fi79 {
                case 128:
                    for fi82 := 0; fi82 < 20; fi82++ {
                        if len(v55) >= 2 {
                            v55 = v55[1:len(v55)]
                        }
                        v55 = append(v55, fi82)
                        fmt.Println(v2, fi79)
                        _ = fi82
                    }
                }
                v3 -= v3
            } else {
                _ = len(v1)
            }
        } else {
            v2 = (false && ("a" != ("\n\t" + ("" + "a"))))
        }
        for fi83 := 0; fi83 < 1; fi83++ {
            v84 := GenS5{F0: genFn0(), F1: v2, F2: "foo bar"}
            fmt.Println(fi79)
            if false {
                if (1 >= -1) {
                    switch (fi83 + fi79) {
                    case 9223372036854775807:
                        switch "a" {
                        case "\n\t":
                            if (0 == -2147483648) {
                                if len(v55) > 0 {
                                    v55[0] = (fi79 * fi83)
                                }
                            }
                        case "hello":
                            v3 = 0
                        }
                    case -157:
                        var v85 uint32 = 539
                        _ = v85
                    case 554:
                        if (65535 == 522) {
                            if (0 < 42) {
                                switch "hello" {
                                case "a":
                                    var v86 string = "\n\t"
                                    _ = v86
                                case "foo bar":
                                    v87 := GenS9{F0: 1, F1: 0, F2: false}
                                    _ = v87
                                case "":
                                    var v88 uint = 603
                                    switch 144 {
                                    case 170:
                                        var v89 int64 = 9223372036854775807
                                        _ = v89
                                    case -89:
                                        _ = len(v1)
                                    default:
                                        v3 -= (v3 >> ((v88) & 63))
                                    }
                                    if ((3.14159 - float32((3.14159 / 1.0))) != -2.25) {
                                        fmt.Println(fi83)
                                        v90 := GenS1{F0: false}
                                        _ = v90
                                    }
                                    _ = v88
                                }
                                var v91 uint32 = 42
                                _ = v91
                            }
                            var v92 int8 = genFn1(0)
                            if false {
                                v93 := GenS10{F0: fi79, F1: 32767, F2: 42}
                                v94 := GenS11{F0: fi83}
                                _ = v93
                                _ = v94
                            } else {
                                switch (((("\n\t" + "\n\t") + "\n\t") + ("foo bar" + "\n\t")) + (("hello" + ("a" + "")) + "")) {
                                case "foo bar":
                                    var v95 uint64 = 1
                                    _ = v95
                                case "":
                                    if (!true) {
                                        fmt.Println(v3, v2, v3)
                                        v2 = v2
                                    } else {
                                        if (3.14159 > 3.14159) {
                                            v96 := GenS12{F0: -2147483648, F1: 0, F2: genFn0()}
                                            if len(v55) > 0 {
                                                v55[0] = (-(fi79))
                                            }
                                            var v97 int8 = (v92 ^ (v3 - -32))
                                            _ = v96
                                            _ = v97
                                        } else {
                                            if v2 {
                                                v92 = (v3 + 0)
                                            } else {
                                                v3 += (v92 | (v92 - v3))
                                            }
                                        }
                                        v3 = (v3 * 1)
                                        if (!true) {
                                            v3 -= (v3 + 100)
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                var v98 float32 = 1e10
                                                _ = v98
                                            case "a":
                                                v99 := GenS9{F0: genFn0(), F1: 2147483647, F2: v2}
                                                _ = v99
                                            }
                                            v3 -= (v3 & -2)
                                        }
                                    }
                                default:
                                    switch (fi83 + (fi83 - (-(fi83)))) {
                                    case -627:
                                        var v100 uint16 = 1
                                        _ = v100
                                    case 561:
                                        var v101 uint = 65535
                                        _ = v101
                                    case 32768:
                                        v3 -= -128
                                    }
                                }
                            }
                            _ = v92
                        }
                    }
                    var v102 uint32 = 65535
                    _ = v102
                }
                v3 = (v3 / ((v3) | 1))
                v2 = (1 > 0)
            }
            _ = fi83
            _ = v84
        }
        if (true && ((v2 && v2) && v2)) {
            if false {
                if (false && (-850 <= -1)) {
                    for fi103 := 0; fi103 < 9; fi103++ {
                        var v104 float64 = (3.14159 + (1e10 - ((0.0 * float64(0.0)) / 3.14159)))
                        _ = fi103
                        _ = v104
                    }
                    _ = len(v55)
                    v2 = ((!(true && (1 <= 0))) && (!(0 != 2)))
                } else {
                    v3 = v3
                    var v105 uint = 18446744073709551615
                    if (!v2) {
                        switch -1 {
                        case -2147483648:
                            var v106 int8 = v3
                            _ = v106
                        }
                    }
                    _ = v105
                }
                v107 := GenS8{F0: genFn1(85), F1: 127}
                _ = v107
            }
            var v108 int = -178
            v109 := GenS3{F0: 0, F1: (-(fi79)), F2: 1}
            _ = v108
            _ = v109
        }
        _ = fi79
    }
    v3 = 1
    for fi110 := 0; fi110 < 8; fi110++ {
        v2 = (true && v2)
        v3 += 1
        _ = fi110
    }
    var v111 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v55
}
