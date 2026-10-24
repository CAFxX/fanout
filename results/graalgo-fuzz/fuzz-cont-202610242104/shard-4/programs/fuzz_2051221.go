package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 int }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 int; F1 uint16; F2 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 bool; F1 int64 }

func main() {
    var v1 string = (("a" + ("\n\t" + "")) + "foo bar")
    var v2 int32 = -1
    var v3 uint32 = 0
    var v4 uint64 = 4294967295
    v3 = v3
    if (-2.25 > float32((-2.25 - (1.5 * 1e10)))) {
        if (false && (false || false)) {
            v5 := GenS1{F0: 683, F1: -32768, F2: 87}
            var v6 string = v1
            _ = v5
            _ = v6
        }
        if (18446744073709551615 <= 4294967295) {
            v1 = "a"
            v7 := GenS1{F0: 18446744073709551615, F1: -675, F2: 42}
            _ = v7
        } else {
            switch 42 {
            case 127:
                var v8 int16 = -32
                _ = v8
            default:
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS2{F0: true, F1: 1, F2: 227}
                    v4 = v4
                    _ = fi9
                    _ = v10
                }
            }
        }
    } else {
        v4 = v4
    }
    v1 = (("a" + "") + v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 bool = true
        if false {
            var v13 uint = 2
            switch (fi11 ^ (fi11 ^ fi11)) {
            case -9223372036854775808:
                var v14 int64 = 0
                _ = v14
            default:
                fmt.Println(v4, v13, fi11)
            }
            v15 := GenS3{F0: (^(v4)), F1: (v4 >> ((v13) & 63))}
            _ = v13
            _ = v15
        }
        v16 := GenS3{F0: (v4 ^ (^(v4))), F1: 383}
        _ = fi11
        _ = v12
        _ = v16
    }
    if (true && false) {
        v17 := GenS3{F0: (^(v4)), F1: (v4 + v4)}
        _ = v17
    } else {
        v18 := []int{-404, 0, 42, -556}
        v4 = v4
        _ = v18
    }
    switch (v1 + (v1 + v1)) {
    case "a":
        var v19 uint = 65535
        _ = v19
    }
    if false {
        switch 859 {
        case 2147483647:
            switch 818 {
            case -403:
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := GenS4{F0: 9223372036854775807}
                    if false {
                        if (!((^(v4)) != (v4 | v4))) {
                            if (!(0.1 >= 3.14159)) {
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    v23 := GenS5{F0: 159, F1: v1, F2: 128}
                                    _ = fi22
                                    _ = v23
                                }
                                var v24 bool = false
                                var v25 uint = 0
                                _ = v24
                                _ = v25
                            }
                        } else {
                            v26 := []int{}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 9; fi27++ {
                            if true {
                                var v28 int64 = -1
                                _ = v28
                            } else {
                                if ((!false) || ((1e10 * float64(0.1)) > (0.1 / (1.5 - ((1e10 - float64(0.0)) / (((0.1 - 3.14159) + 3.14159) - 0.1)))))) {
                                    if true {
                                        switch (v1 + "a") {
                                        case "hello":
                                            v29 := GenS5{F0: 1, F1: (v1 + v1), F2: 128}
                                            _ = v29
                                        case "":
                                            switch (fi27 + (fi20 - fi27)) {
                                            case -1:
                                                v1 = "a"
                                            }
                                        }
                                    }
                                    v30 := []int{(^(fi20)), (-(fi20)), fi27}
                                    v31 := GenS6{F0: 48, F1: 0}
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v32 int64 = 945
                                    if true {
                                        var v33 uint16 = 255
                                        v32 += (v32 | v32)
                                        _ = v33
                                    } else {
                                        v34 := GenS6{F0: 255, F1: 851}
                                        _ = v34
                                    }
                                    _ = v32
                                }
                                var v35 int16 = 32767
                                v2 += int32(fi20)
                                switch (fi27 / ((128) | 1)) {
                                case 127:
                                    v36 := []int{}
                                    _ = v36
                                default:
                                    v37 := GenS7{F0: 32767, F1: 32767, F2: fi27}
                                    v38 := GenS3{F0: (v4 / (((v4 | (v4 / (((v4 >> ((0) & 63))) | 1)))) | 1)), F1: 42}
                                    v2 += -574
                                    v39 := []int{fi20, fi27}
                                    if true {
                                        fmt.Println(v3, v4)
                                    } else {
                                        var v40 int64 = -1
                                        v41 := GenS3{F0: v4, F1: 2}
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                _ = v35
                            }
                            fmt.Println(v1, v4, v1)
                            v42 := []int{(fi27 + -512), (fi27 / ((fi20) | 1)), -206, fi27}
                            _ = fi27
                            _ = v42
                        }
                    }
                    switch fi20 {
                    case 127:
                        var v43 uint = 0
                        _ = v43
                    case -210:
                        v44 := GenS2{F0: ((("hello" + v1) + v1) <= v1), F1: -9223372036854775808, F2: 2}
                        _ = v44
                    case 32768:
                        var v45 float32 = (3.14159 - (0.0 - float32(0.0)))
                        _ = v45
                    default:
                        v2 -= 213
                    }
                    _ = fi20
                    _ = v21
                }
            default:
                v2 -= (v2 / (((^(v2))) | 1))
            }
        case 9223372036854775807:
            var v46 int = 32768
            _ = v46
        default:
            v47 := GenS6{F0: 127, F1: 18446744073709551615}
            _ = v47
        }
        for fi48 := 0; fi48 < 4; fi48++ {
            v49 := []int{fi48}
            if (((-2.25 - -2.25) / (0.0 - 3.14159)) != ((0.0 / (0.1 - 0.0)) - 1.5)) {
                for fi50 := 0; fi50 < 5; fi50++ {
                    var v51 int8 = -79
                    if ((true || true) && (float64((float32(3.14159) * 3.14159)) < 1e10)) {
                        switch fi50 {
                        case -268:
                            if true {
                                v1 = (("" + ("a" + v1)) + v1)
                            }
                        case -1:
                            var v52 int16 = int16(v3)
                            var v53 bool = false
                            _ = v52
                            _ = v53
                        }
                        fmt.Println(v2)
                    } else {
                        fmt.Println(fi50)
                        var v54 int16 = 42
                        v55 := GenS2{F0: false, F1: 42, F2: 0}
                        _ = v54
                        _ = v55
                    }
                    _ = fi50
                    _ = v51
                }
                _ = len(v49)
                _ = len(v49)
            }
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            _ = fi48
            _ = v49
        }
        v1 = v1
    }
    var v56 uint64 = v4
    v57 := []int{780, 4294967295, 646}
    var v58 float64 = float64(-2.25)
    switch 32768 {
    case -456:
        for fi59 := 0; fi59 < 6; fi59++ {
            if false {
                if false {
                    switch fi59 {
                    case 32768:
                        v60 := GenS5{F0: 187, F1: "a", F2: (-(fi59))}
                        _ = v60
                    case 2147483647:
                        v56 += v56
                    default:
                        switch fi59 {
                        case 1:
                            if (false && ("hello" >= v1)) {
                                v2 = (v2 ^ (^(v2)))
                                fmt.Println(v56)
                            } else {
                                v56 = (v4 - (v56 + 4294967295))
                            }
                        case 191:
                            var v61 uint32 = 1
                            _ = v61
                        }
                    }
                } else {
                    if (!(!true)) {
                        switch (fi59 ^ (fi59 + -630)) {
                        case 32767:
                            for fi62 := 0; fi62 < 6; fi62++ {
                                var v63 uint64 = (v4 >> ((163) & 63))
                                _ = fi62
                                _ = v63
                            }
                        }
                        switch fi59 {
                        case 2147483647:
                            if false {
                                var v64 uint16 = 0
                                var v65 uint = 348
                                if len(v57) > 0 {
                                    v57[0] = int(v65)
                                }
                                _ = v64
                                _ = v65
                            }
                        case 1:
                            v3 = v3
                        }
                    }
                    if true {
                        v4 += 42
                        v66 := GenS2{F0: false, F1: -722, F2: 255}
                        var v67 uint64 = (v4 / ((v56) | 1))
                        _ = v66
                        _ = v67
                    } else {
                        v68 := GenS8{F0: (v2 | (-(v2)))}
                        _ = v68
                    }
                }
            } else {
                if false {
                    if len(v57) > 0 {
                        v57[0] = (fi59 ^ 1)
                    }
                }
                for fi69 := 0; fi69 < 16; fi69++ {
                    if (v1 <= "foo bar") {
                        var v70 uint = 42
                        switch "hello" {
                        case "foo bar":
                            var v71 int = fi69
                            _ = v71
                        case "hello":
                            switch 128 {
                            case 0:
                                fmt.Println(v70)
                            case -473:
                                var v72 int8 = 127
                                _ = v72
                            default:
                                v4 -= v56
                            }
                        }
                        if false {
                            var v73 int32 = 0
                            var v74 string = ""
                            var v75 bool = false
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        _ = v70
                    }
                    v76 := GenS9{F0: 901, F1: (v3 * v3), F2: (v3 / ((42) | 1))}
                    _ = fi69
                    _ = v76
                }
                if true {
                    for fi77 := 0; fi77 < 4; fi77++ {
                        v1 = "foo bar"
                        if (!false) {
                            if (v2 != v2) {
                                v56 = v4
                                var v78 int64 = -9223372036854775808
                                v78 -= v78
                                var v79 float64 = float64(0.1)
                                _ = v78
                                _ = v79
                            }
                            var v80 int = fi59
                            if (118 < 2) {
                                v81 := GenS10{F0: -1, F1: true}
                                _ = v81
                            }
                            _ = v80
                        } else {
                            var v82 bool = true
                            _ = v82
                        }
                        v83 := GenS11{F0: ((3.14159 + 1.5) <= float32(v58)), F1: 0}
                        _ = fi77
                        _ = v83
                    }
                    switch (fi59 | 2147483647) {
                    case 637:
                        v84 := GenS10{F0: (v2 % ((-1) | 1)), F1: (-9223372036854775808 >= -898)}
                        _ = v84
                    case -365:
                        v85 := GenS5{F0: 65535, F1: (("a" + v1) + ("hello" + "\n\t")), F2: (^(fi59))}
                        _ = v85
                    }
                    var v86 int8 = 83
                    _ = v86
                } else {
                    if (false || ((!(true || true)) && false)) {
                        fmt.Println(v4)
                        v56 -= v4
                        if false {
                            var v87 int64 = 0
                            _ = v87
                        }
                    }
                    var v88 uint32 = 42
                    _ = v88
                }
            }
            if false {
                for fi89 := 0; fi89 < 5; fi89++ {
                    v90 := GenS9{F0: 803, F1: (v3 & v3), F2: (v3 % ((v3) | 1))}
                    if (((-2.25 - 1e10) >= (1.5 * 1e10)) && false) {
                        v3 -= v3
                        if (!false) {
                            v57 = append(v57, (fi89 >> ((0) & 63)))
                        } else {
                            v56 = v4
                            if (!true) {
                                switch ((v1 + (v1 + v1)) + v1) {
                                case "":
                                    if len(v57) >= 2 {
                                        v57 = v57[1:len(v57)]
                                    }
                                case "\n\t":
                                    v91 := GenS10{F0: 42, F1: true}
                                    _ = v91
                                case "a":
                                    v1 = v1
                                default:
                                    fmt.Println(fi89, v2, v58)
                                }
                                fmt.Println(fi89, v2)
                            } else {
                                fmt.Println(v2, fi59, fi59)
                            }
                            _ = len(v57)
                        }
                        switch v1 {
                        case "foo bar":
                            if ((42 != 1) && (false || (v4 == 42))) {
                                if (18446744073709551615 == 561) {
                                    if true {
                                        var v92 int32 = (v2 ^ v2)
                                        v4 = 594
                                        var v93 string = "hello"
                                        _ = v92
                                        _ = v93
                                    } else {
                                        if len(v57) >= 2 {
                                            v57 = v57[1:len(v57)]
                                        }
                                        v94 := GenS4{F0: 0}
                                        switch fi59 {
                                        case 127:
                                            if len(v57) >= 2 {
                                                v57 = v57[1:len(v57)]
                                            }
                                        case 0:
                                            fmt.Println(v58)
                                            if false {
                                                v1 = (v1 + "hello")
                                                var v95 string = ("" + "foo bar")
                                                v3 -= v3
                                                var v96 float64 = (3.14159 * 3.14159)
                                                _ = v95
                                                _ = v96
                                            } else {
                                                var v97 uint8 = 0
                                                if ((v3 * v3) < 21) {
                                                    if len(v57) > 0 {
                                                        v57[0] = 0
                                                    }
                                                    v1 = v1
                                                    if len(v57) >= 2 {
                                                        v57 = v57[1:len(v57)]
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v1, v2)
                                                    } else {
                                                        fmt.Println(v58)
                                                        var v98 int = fi59
                                                        if len(v57) > 0 {
                                                            v57[0] = (-(fi89))
                                                        }
                                                        _ = v98
                                                    }
                                                    var v99 uint64 = (v4 | v56)
                                                    _ = v99
                                                }
                                                _ = v97
                                            }
                                        case 32768:
                                            var v100 int = fi59
                                            _ = v100
                                        default:
                                            switch fi59 {
                                            case 9223372036854775807:
                                                v4 += (v4 + 208)
                                            }
                                        }
                                        _ = v94
                                    }
                                }
                                var v101 float64 = 0.1
                                var v102 int = (-(fi59))
                                v1 = ""
                                _ = v101
                                _ = v102
                            } else {
                                _ = len(v57)
                                v4 += v56
                            }
                        case "a":
                            v56 -= v4
                        default:
                            fmt.Println(fi89, v1)
                        }
                    } else {
                        var v103 uint8 = 42
                        v3 += 2147483647
                        v58 = v58
                        _ = v103
                    }
                    _ = fi89
                    _ = v90
                }
                v1 = v1
            }
            _ = fi59
        }
    case -720:
        var v104 bool = false
        _ = v104
    case 4294967295:
        _ = len(v57)
    }
    var v105 int64 = 42
    v4 -= 1
    var v106 int16 = -1
    v1 = ((v1 + ("\n\t" + v1)) + v1)
    fmt.Println(v58, v56)
    var v107 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
}
