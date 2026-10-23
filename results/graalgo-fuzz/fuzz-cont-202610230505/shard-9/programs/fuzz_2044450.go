package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
type GenS4 struct { F0 uint; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 int64; F1 uint64; F2 uint }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 int16; F1 int16; F2 int }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 bool; F1 uint }
type GenS15 struct { F0 uint64; F1 int8; F2 int32 }
type GenS16 struct { F0 int64; F1 int16; F2 int16 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int32; F1 uint64 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint; F1 int16 }
type GenS24 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS25 struct { F0 int8; F1 uint16; F2 bool }
type GenS26 struct { F0 int16; F1 int8 }
type GenS27 struct { F0 int64; F1 int }

func main() {
    var v1 int8 = -1
    var v2 int8 = -128
    var v3 uint8 = 255
    var v4 int8 = (-(v2))
    var v5 bool = (!true)
    var v6 float64 = (1.5 + 1.5)
    switch 32767 {
    case 4294967295:
        switch "hello" {
        case "\n\t":
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(v4, v3)
                var v8 float32 = 1.5
                _ = fi7
                _ = v8
            }
        case "":
            switch ("foo bar" + ("a" + "foo bar")) {
            case "a":
                if true {
                    v3 = (v3 - (v3 * v3))
                } else {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            v11 := []int{(fi9 & fi9), fi9, fi9, (fi9 % ((fi9) | 1))}
                            _ = fi10
                            _ = v11
                        }
                        if false {
                            v1 = (v1 + v1)
                            var v12 int = (-(fi9))
                            _ = v12
                        }
                        switch 1 {
                        case 127:
                            if ((v2 / (((v1 & (v4 & (v4 / ((v2) | 1))))) | 1)) != v2) {
                                if (0.0 < v6) {
                                    v13 := []int{}
                                    var v14 int = fi9
                                    _ = v13
                                    _ = v14
                                } else {
                                    var v15 float32 = -2.25
                                    if v5 {
                                        v16 := []int{}
                                        _ = v16
                                    }
                                    v6 -= v6
                                    _ = v15
                                }
                                var v17 int64 = 380
                                _ = v17
                            } else {
                                fmt.Println(v1)
                                v18 := GenS1{F0: 18446744073709551615, F1: (!v5)}
                                _ = v18
                            }
                        case -9223372036854775808:
                            v19 := []int{fi9, fi9, fi9}
                            _ = v19
                        case -932:
                            v20 := GenS1{F0: 662, F1: (550 == 2147483648)}
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 4; fi21++ {
                                var v22 int = 600
                                fmt.Println(fi9)
                                v22 -= (fi21 >> ((18446744073709551615) & 63))
                                _ = fi21
                                _ = v22
                            }
                        }
                        _ = fi9
                    }
                    switch "hello" {
                    case "\n\t":
                        if ((-2.25 <= (-2.25 + 3.14159)) || v5) {
                            for fi23 := 0; fi23 < 5; fi23++ {
                                if false {
                                    v6 = float64(1.5)
                                    v24 := []int{fi23, fi23, (fi23 + (fi23 + fi23))}
                                    v25 := GenS1{F0: 0, F1: (830 < 363)}
                                    _ = v24
                                    _ = v25
                                }
                                _ = fi23
                            }
                            var v26 uint64 = 1
                            _ = v26
                        }
                    }
                }
            default:
                if (v5 && false) {
                    if (42 < -41) {
                        v27 := GenS2{F0: true}
                        _ = v27
                    } else {
                        v28 := GenS1{F0: 2, F1: (v5 || (!(0 <= 1)))}
                        for fi29 := 0; fi29 < 17; fi29++ {
                            for fi30 := 0; fi30 < 15; fi30++ {
                                var v31 int16 = -1
                                if v5 {
                                    var v32 float64 = v6
                                    _ = v32
                                } else {
                                    v33 := []int{fi30, fi29, fi30}
                                    var v34 string = (("a" + "foo bar") + "foo bar")
                                    var v35 int32 = int32(v2)
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                var v36 float32 = 0.0
                                _ = fi30
                                _ = v31
                                _ = v36
                            }
                            _ = fi29
                        }
                        if v5 {
                            v2 = -41
                            if v5 {
                                v1 += 127
                                v6 += float64(((-2.25 - (1.5 * float32(0.1))) / float32(1.5)))
                                if v5 {
                                    if false {
                                        switch -304 {
                                        case 405:
                                            switch 573 {
                                            case 1:
                                                fmt.Println(v6, v4, v3)
                                            case -334:
                                                v5 = (0 != 65535)
                                            case 971:
                                                if true {
                                                    var v37 uint8 = 127
                                                    if v5 {
                                                        for fi38 := 0; fi38 < 20; fi38++ {
                                                            v39 := GenS1{F0: 717, F1: v5}
                                                            if (!((fi38 ^ fi38) >= (fi38 + -9223372036854775808))) {
                                                                switch (fi38 << ((42) & 63)) {
                                                                case 9223372036854775807:
                                                                    v40 := []int{(fi38 << ((65535) & 63)), fi38}
                                                                    _ = v40
                                                                case 470:
                                                                    v41 := []int{fi38, (fi38 | fi38), (fi38 << ((197) & 63)), fi38}
                                                                    _ = v41
                                                                default:
                                                                    fmt.Println(fi38)
                                                                }
                                                                if v5 {
                                                                    v42 := GenS3{F0: v2, F1: v4, F2: 482}
                                                                    for fi43 := 0; fi43 < 2; fi43++ {
                                                                        v44 := []int{-9223372036854775808, (-(fi43)), (^(fi43))}
                                                                        if v5 {
                                                                            if len(v44) >= 2 {
                                                                                v44 = v44[1:len(v44)]
                                                                            }
                                                                        } else {
                                                                            v45 := GenS4{F0: 65535, F1: -73, F2: 18446744073709551615}
                                                                            var v46 int = fi43
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                        if len(v44) > 0 {
                                                                            v44[0] = fi43
                                                                        }
                                                                        if len(v44) >= 2 {
                                                                            v44 = v44[1:len(v44)]
                                                                        }
                                                                        fmt.Println(v37, v3)
                                                                        _ = fi43
                                                                        _ = v44
                                                                    }
                                                                    _ = v42
                                                                }
                                                                switch (((((("\n\t" + "") + "a") + "hello") + "foo bar") + ((("foo bar" + ("foo bar" + "foo bar")) + (("\n\t" + "foo bar") + ("hello" + ("foo bar" + "hello")))) + "foo bar")) + "hello") {
                                                                case "a":
                                                                    if v5 {
                                                                        var v47 bool = (!v5)
                                                                        v3 = v37
                                                                        var v48 int32 = 2147483647
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        v1 += 95
                                                                        v49 := []int{}
                                                                        v50 := GenS5{F0: 150, F1: (^(v3)), F2: false}
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                case "":
                                                                    for fi51 := 0; fi51 < 1; fi51++ {
                                                                        v52 := GenS6{F0: (v37 & (v37 * v3))}
                                                                        if (-1 >= 0) {
                                                                            v6 = v6
                                                                            v53 := []int{fi38}
                                                                            _ = v53
                                                                        } else {
                                                                            var v54 uint64 = 4294967295
                                                                            fmt.Println(v5)
                                                                            _ = v54
                                                                        }
                                                                        if ((v5 && v5) || ((455 != 4294967295) || false)) {
                                                                            fmt.Println(fi38, fi51, v4)
                                                                            v3 = (v37 * (v37 - (v3 + (v37 * 243))))
                                                                            v55 := GenS7{F0: fi38, F1: 0}
                                                                            fmt.Println(v5, fi51, v3)
                                                                            if v5 {
                                                                                var v56 uint16 = 42
                                                                                switch (("hello" + "hello") + ("foo bar" + "hello")) {
                                                                                case "a":
                                                                                    var v57 int16 = -32768
                                                                                    _ = v57
                                                                                case "\n\t":
                                                                                    var v58 int16 = 42
                                                                                    _ = v58
                                                                                case "":
                                                                                    var v59 int64 = 42
                                                                                    _ = v59
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        v37 -= (v3 * 37)
                                                                                    case "a":
                                                                                        v60 := []int{}
                                                                                        _ = v60
                                                                                    default:
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v61 := GenS2{F0: (((0 == 255) && v5) && v5)}
                                                                                            _ = v61
                                                                                        case "foo bar":
                                                                                            v62 := GenS5{F0: 200, F1: 128, F2: (18446744073709551615 <= 2)}
                                                                                            _ = v62
                                                                                        case "a":
                                                                                            var v63 int16 = 0
                                                                                            switch (-(fi38)) {
                                                                                            case 42:
                                                                                                v64 := GenS8{F0: 181}
                                                                                                _ = v64
                                                                                            case 128:
                                                                                                v3 += v37
                                                                                            default:
                                                                                                var v65 uint = 42
                                                                                                _ = v65
                                                                                            }
                                                                                            _ = v63
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v56
                                                                            } else {
                                                                                if true {
                                                                                    if v5 {
                                                                                        v66 := []int{(^(fi51)), (-(fi38))}
                                                                                        if (652 >= 2) {
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                if v5 {
                                                                                                    v67 := GenS1{F0: 46, F1: false}
                                                                                                    if v5 {
                                                                                                        fmt.Println(v3)
                                                                                                        var v68 float64 = 1.5
                                                                                                        v6 = -2.25
                                                                                                        _ = v68
                                                                                                    } else {
                                                                                                        v2 -= v4
                                                                                                    }
                                                                                                    _ = v67
                                                                                                } else {
                                                                                                    if v5 {
                                                                                                        fmt.Println(v37, v3, v4)
                                                                                                    } else {
                                                                                                        v69 := GenS9{F0: 0, F1: ((("a" + "hello") + "\n\t") + "\n\t")}
                                                                                                        v66 = append(v66, 128)
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    v70 := GenS10{F0: -9223372036854775808, F1: 42, F2: 255}
                                                                                                    var v71 float32 = 0.0
                                                                                                    if (-32768 == 0) {
                                                                                                        v72 := GenS11{F0: 42, F1: v1}
                                                                                                        v73 := GenS12{F0: 32767, F1: -635, F2: fi38}
                                                                                                        _ = v72
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        switch ("foo bar" + "foo bar") {
                                                                                                        case "":
                                                                                                            v5 = v5
                                                                                                        case "hello":
                                                                                                            if len(v66) > 0 {
                                                                                                                v66[0] = fi38
                                                                                                            }
                                                                                                        default:
                                                                                                            _ = len(v66)
                                                                                                        }
                                                                                                        switch fi38 {
                                                                                                        case 127:
                                                                                                            if true {
                                                                                                                v74 := GenS11{F0: 42, F1: 42}
                                                                                                                v37 = (v37 | v3)
                                                                                                                _ = v74
                                                                                                            } else {
                                                                                                                if len(v66) > 0 {
                                                                                                                    v66[0] = 497
                                                                                                                }
                                                                                                                var v75 string = ("\n\t" + "foo bar")
                                                                                                                _ = v75
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            v76 := GenS11{F0: -1, F1: (v2 | (v1 * (v2 << ((0) & 63))))}
                                                                                                            _ = v76
                                                                                                        case 335:
                                                                                                            if len(v66) > 0 {
                                                                                                                v66[0] = fi38
                                                                                                            }
                                                                                                        }
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            if v5 {
                                                                                                                if true {
                                                                                                                    v6 -= v6
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 62:
                                                                                                                        var v77 uint = 4294967295
                                                                                                                        v78 := GenS13{F0: 1, F1: 384}
                                                                                                                        _ = v77
                                                                                                                        _ = v78
                                                                                                                    case 1:
                                                                                                                        if v5 {
                                                                                                                            switch fi38 {
                                                                                                                            case 1:
                                                                                                                                fmt.Println(v37, v4)
                                                                                                                            default:
                                                                                                                                switch "a" {
                                                                                                                                case "a":
                                                                                                                                    v79 := GenS12{F0: 213, F1: -32768, F2: (fi38 - fi38)}
                                                                                                                                    switch "" {
                                                                                                                                    case "":
                                                                                                                                        var v80 int16 = -32768
                                                                                                                                        _ = v80
                                                                                                                                    }
                                                                                                                                    var v81 uint8 = (v37 + 209)
                                                                                                                                    if len(v66) > 0 {
                                                                                                                                        v66[0] = fi51
                                                                                                                                    }
                                                                                                                                    if (v5 || (("" + "a") < "foo bar")) {
                                                                                                                                        if (true || ((42 <= 385) && v5)) {
                                                                                                                                            var v82 uint32 = 2147483648
                                                                                                                                            v83 := GenS14{F0: v5, F1: 18446744073709551615}
                                                                                                                                            _ = v82
                                                                                                                                            _ = v83
                                                                                                                                        } else {
                                                                                                                                            var v84 int16 = -358
                                                                                                                                            _ = len(v66)
                                                                                                                                            var v85 int64 = 42
                                                                                                                                            _ = v84
                                                                                                                                            _ = v85
                                                                                                                                        }
                                                                                                                                        var v86 float64 = -2.25
                                                                                                                                        var v87 bool = (true || (fi51 != (fi38 % (((fi51 % ((fi51) | 1))) | 1))))
                                                                                                                                        if false {
                                                                                                                                            fmt.Println(v5)
                                                                                                                                            v88 := GenS15{F0: 4294967295, F1: v4, F2: 1}
                                                                                                                                            _ = v88
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(fi51, v5)
                                                                                                                                        }
                                                                                                                                        fmt.Println(v2)
                                                                                                                                        _ = v86
                                                                                                                                        _ = v87
                                                                                                                                    }
                                                                                                                                    _ = v79
                                                                                                                                    _ = v81
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if v5 {
                                                                                                                                switch "\n\t" {
                                                                                                                                case "\n\t":
                                                                                                                                    v89 := GenS1{F0: 2, F1: (v5 && ("" > "hello"))}
                                                                                                                                    _ = v89
                                                                                                                                case "hello":
                                                                                                                                    var v90 int = fi38
                                                                                                                                    _ = v90
                                                                                                                                default:
                                                                                                                                    v66 = append(v66, (fi38 & (fi51 ^ (^(fi38)))))
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if true {
                                                                                                                                    var v91 uint8 = (v37 * 134)
                                                                                                                                    _ = v91
                                                                                                                                } else {
                                                                                                                                    switch (-(fi51)) {
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        if v5 {
                                                                                                                                            if v5 {
                                                                                                                                                switch ("" + ("" + ("" + "hello"))) {
                                                                                                                                                case "hello":
                                                                                                                                                    v37 = v3
                                                                                                                                                default:
                                                                                                                                                    var v92 uint = 4294967295
                                                                                                                                                    var v93 uint64 = 18446744073709551615
                                                                                                                                                    var v94 int64 = 0
                                                                                                                                                    _ = v92
                                                                                                                                                    _ = v93
                                                                                                                                                    _ = v94
                                                                                                                                                }
                                                                                                                                                if v5 {
                                                                                                                                                    v95 := GenS14{F0: true, F1: 65535}
                                                                                                                                                    _ = v95
                                                                                                                                                }
                                                                                                                                                v96 := GenS16{F0: -9223372036854775808, F1: 1000, F2: -1}
                                                                                                                                                var v97 uint8 = (^(v3))
                                                                                                                                                _ = v96
                                                                                                                                                _ = v97
                                                                                                                                            }
                                                                                                                                            fmt.Println(v1)
                                                                                                                                            var v98 int = (^(fi51))
                                                                                                                                            var v99 int64 = 17
                                                                                                                                            _ = v98
                                                                                                                                            _ = v99
                                                                                                                                        }
                                                                                                                                    case 127:
                                                                                                                                        switch ((("foo bar" + "hello") + "a") + "a") {
                                                                                                                                        case "a":
                                                                                                                                            if len(v66) >= 2 {
                                                                                                                                                v66 = v66[1:len(v66)]
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        fmt.Println(v37, fi51)
                                                                                                                                        var v100 uint16 = 989
                                                                                                                                        _ = v100
                                                                                                                                    }
                                                                                                                                    switch "foo bar" {
                                                                                                                                    case "foo bar":
                                                                                                                                        var v101 uint = 395
                                                                                                                                        _ = v101
                                                                                                                                    case "a":
                                                                                                                                        v102 := GenS16{F0: -9223372036854775808, F1: 1000, F2: 1}
                                                                                                                                        _ = v102
                                                                                                                                    default:
                                                                                                                                        if false {
                                                                                                                                            fmt.Println(v37, v3, v71)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v103 := GenS12{F0: -1, F1: 32767, F2: fi38}
                                                                                                                                _ = v103
                                                                                                                            }
                                                                                                                            v66 = append(v66, (fi51 % ((-9223372036854775808) | 1)))
                                                                                                                            if (!false) {
                                                                                                                                v104 := GenS17{F0: -786}
                                                                                                                                _ = v104
                                                                                                                            } else {
                                                                                                                                fmt.Println(v5, v37)
                                                                                                                            }
                                                                                                                            var v105 uint16 = 255
                                                                                                                            v106 := GenS18{F0: 0}
                                                                                                                            _ = v105
                                                                                                                            _ = v106
                                                                                                                        }
                                                                                                                        _ = len(v66)
                                                                                                                    case 637:
                                                                                                                        v66 = append(v66, (fi38 + int(v1)))
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        case "a":
                                                                                                            if v5 {
                                                                                                                if (!v5) {
                                                                                                                    if (float64((v71 / v71)) > 1e10) {
                                                                                                                        v3 -= v3
                                                                                                                        if true {
                                                                                                                            var v107 int32 = -857
                                                                                                                            _ = v107
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v71 += 0.1
                                                                                                                }
                                                                                                                v108 := GenS12{F0: -32768, F1: 42, F2: (fi51 % ((fi51) | 1))}
                                                                                                                _ = v108
                                                                                                            }
                                                                                                        case "":
                                                                                                            if len(v66) >= 2 {
                                                                                                                v66 = v66[1:len(v66)]
                                                                                                            }
                                                                                                        default:
                                                                                                            v109 := GenS19{F0: 328}
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v6)
                                                                                                    v4 -= v1
                                                                                                    var v110 int = fi38
                                                                                                    _ = v110
                                                                                                    _ = v70
                                                                                                    _ = v71
                                                                                                }
                                                                                            case "foo bar":
                                                                                                if len(v66) > 0 {
                                                                                                    v66[0] = (fi38 / ((fi51) | 1))
                                                                                                }
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    var v111 int64 = 9223372036854775807
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        var v112 int8 = (-(v4))
                                                                                                        _ = v112
                                                                                                    case "a":
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            fmt.Println(v5)
                                                                                                        case "hello":
                                                                                                            fmt.Println(v5)
                                                                                                        case "foo bar":
                                                                                                            var v113 uint64 = 4294967295
                                                                                                            _ = v113
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        v4 = (v4 + 100)
                                                                                                    }
                                                                                                    if (!v5) {
                                                                                                        if v5 {
                                                                                                            v114 := GenS20{F0: -1}
                                                                                                            _ = v114
                                                                                                        } else {
                                                                                                            if len(v66) >= 2 {
                                                                                                                v66 = v66[1:len(v66)]
                                                                                                            }
                                                                                                            var v115 string = ""
                                                                                                            _ = v115
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v111
                                                                                                }
                                                                                            }
                                                                                            v6 -= float64(3.14159)
                                                                                        }
                                                                                        fmt.Println(v6, v3)
                                                                                        v116 := GenS14{F0: (!((false || ((v4 < v2) || false)) && true)), F1: 1}
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            var v117 uint32 = 606
                                                                                            _ = v117
                                                                                        case "":
                                                                                            var v118 uint = 206
                                                                                            _ = v118
                                                                                        case "hello":
                                                                                            v119 := GenS3{F0: (-(v2)), F1: (v4 - (v2 - 1)), F2: 229}
                                                                                            _ = v119
                                                                                        default:
                                                                                            v120 := GenS21{F0: -2147483648, F1: 0}
                                                                                            _ = v120
                                                                                        }
                                                                                        v121 := GenS22{F0: 972}
                                                                                        _ = v116
                                                                                        _ = v121
                                                                                        _ = v66
                                                                                    }
                                                                                    var v122 uint = 255
                                                                                    var v123 uint64 = 458
                                                                                    var v124 string = "foo bar"
                                                                                    v125 := GenS23{F0: v122, F1: 32767}
                                                                                    var v126 int = (fi51 % (((fi51 - fi38)) | 1))
                                                                                    switch v124 {
                                                                                    case "hello":
                                                                                        v127 := []int{903}
                                                                                        v128 := GenS24{F0: 42, F1: 128, F2: (^(v4))}
                                                                                        _ = len(v127)
                                                                                        _ = v127
                                                                                        _ = v128
                                                                                    case "\n\t":
                                                                                        fmt.Println(v124, v124, v6)
                                                                                    default:
                                                                                        var v129 string = v124
                                                                                        _ = v129
                                                                                    }
                                                                                    _ = v122
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                } else {
                                                                                    if ((^(fi38)) > (fi51 - 4294967295)) {
                                                                                        var v130 int16 = 42
                                                                                        var v131 uint8 = 217
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            var v132 int8 = 100
                                                                                            _ = v132
                                                                                        default:
                                                                                            v1 = v4
                                                                                        }
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                    } else {
                                                                                        v133 := GenS21{F0: -1, F1: 0}
                                                                                        _ = v133
                                                                                    }
                                                                                    v1 = 110
                                                                                    if true {
                                                                                        var v134 bool = (v5 || (18446744073709551615 != 378))
                                                                                        switch "" {
                                                                                        case "":
                                                                                            fmt.Println(v5)
                                                                                            v6 += v6
                                                                                            var v135 bool = false
                                                                                            if true {
                                                                                                v134 = ((-84 <= v2) && v135)
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    fmt.Println(v2, fi51)
                                                                                                default:
                                                                                                    var v136 uint8 = (v3 % (((v3 & 26)) | 1))
                                                                                                    _ = v136
                                                                                                }
                                                                                            }
                                                                                            v137 := GenS9{F0: 669, F1: "a"}
                                                                                            _ = v135
                                                                                            _ = v137
                                                                                        default:
                                                                                            var v138 float32 = ((3.14159 / ((1.5 - float32(v6)) - -2.25)) + (-2.25 * float32(((1.5 + 3.14159) / v6))))
                                                                                            _ = v138
                                                                                        }
                                                                                        if ((!(0 <= 42)) || (false || true)) {
                                                                                            if (true && (2 < 2)) {
                                                                                                var v139 int = 2147483647
                                                                                                var v140 float64 = 0.1
                                                                                                _ = v139
                                                                                                _ = v140
                                                                                            }
                                                                                        } else {
                                                                                            v134 = v134
                                                                                        }
                                                                                        _ = v134
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi51, v5)
                                                                            }
                                                                            var v141 int8 = 42
                                                                            _ = v141
                                                                            _ = v55
                                                                        } else {
                                                                            var v142 uint8 = (v3 % ((1) | 1))
                                                                            var v143 uint = 42
                                                                            _ = v142
                                                                            _ = v143
                                                                        }
                                                                        _ = fi51
                                                                        _ = v52
                                                                    }
                                                                case "foo bar":
                                                                    var v144 uint = 18446744073709551615
                                                                    _ = v144
                                                                }
                                                            } else {
                                                                var v145 uint8 = (v3 << ((18446744073709551615) & 63))
                                                                for fi146 := 0; fi146 < 17; fi146++ {
                                                                    if (427 < -2147483648) {
                                                                        v147 := GenS3{F0: (v4 + (v1 * v4)), F1: (v2 / ((v4) | 1)), F2: 0}
                                                                        if (3.14159 != 1.5) {
                                                                            v148 := []int{}
                                                                            _ = v148
                                                                        } else {
                                                                            switch -1 {
                                                                            case 42:
                                                                                var v149 int16 = 1000
                                                                                _ = v149
                                                                            case -713:
                                                                                fmt.Println(v5, v4, v5)
                                                                            default:
                                                                                var v150 uint32 = 984
                                                                                _ = v150
                                                                            }
                                                                            v3 = (v3 / ((v37) | 1))
                                                                            v4 -= v1
                                                                        }
                                                                        _ = v147
                                                                    }
                                                                    _ = fi146
                                                                }
                                                                _ = v145
                                                            }
                                                            _ = fi38
                                                            _ = v39
                                                        }
                                                        for fi151 := 0; fi151 < 17; fi151++ {
                                                            v152 := []int{-60, fi151, fi151, fi151}
                                                            _ = fi151
                                                            _ = v152
                                                        }
                                                        v37 += (v37 & 146)
                                                    } else {
                                                        if (v3 <= (v37 << ((42) & 63))) {
                                                            v153 := GenS25{F0: v4, F1: 65535, F2: true}
                                                            switch 850 {
                                                            case 4294967295:
                                                                var v154 string = ""
                                                                _ = v154
                                                            case 270:
                                                                var v155 float32 = 1e10
                                                                _ = v155
                                                            default:
                                                                v6 = 0.1
                                                            }
                                                            _ = v153
                                                        }
                                                    }
                                                    var v156 float64 = v6
                                                    _ = v156
                                                    _ = v37
                                                }
                                            }
                                        case 0:
                                            v4 = 42
                                        case -835:
                                            var v157 int32 = 1
                                            _ = v157
                                        }
                                        v5 = (v5 && (v5 && v5))
                                    } else {
                                        v158 := []int{32768, 42, 4294967295, -304}
                                        _ = len(v158)
                                        v158 = append(v158, 32767)
                                        _ = v158
                                    }
                                }
                            } else {
                                v159 := []int{-2147483648, 9223372036854775807}
                                if (!false) {
                                    if v5 {
                                        switch "" {
                                        case "a":
                                            if true {
                                                switch -1 {
                                                case 0:
                                                    v6 = 1e10
                                                case -2147483648:
                                                    v1 = v2
                                                default:
                                                    var v160 bool = false
                                                    _ = v160
                                                }
                                                for fi161 := 0; fi161 < 9; fi161++ {
                                                    var v162 float64 = v6
                                                    v163 := GenS26{F0: 543, F1: 100}
                                                    fmt.Println(v162)
                                                    _ = fi161
                                                    _ = v162
                                                    _ = v163
                                                }
                                                v2 = (v1 / ((v1) | 1))
                                            }
                                        case "foo bar":
                                            v5 = false
                                        case "":
                                            var v164 int64 = 760
                                            _ = v164
                                        }
                                        var v165 int32 = -1
                                        _ = v165
                                    }
                                    switch 1 {
                                    case 128:
                                        fmt.Println(v1, v4, v6)
                                    case 32767:
                                        if len(v159) >= 2 {
                                            v159 = v159[1:len(v159)]
                                        }
                                    case -1:
                                        fmt.Println(v4, v2, v6)
                                    default:
                                        fmt.Println(v5, v4, v4)
                                    }
                                    for fi166 := 0; fi166 < 9; fi166++ {
                                        if len(v159) >= 2 {
                                            v159 = v159[1:len(v159)]
                                        }
                                        var v167 int = (fi166 - (fi166 >> ((172) & 63)))
                                        v1 = v4
                                        _ = fi166
                                        _ = v167
                                    }
                                } else {
                                    v168 := GenS9{F0: 0, F1: ""}
                                    if v5 {
                                        if (0 != 2) {
                                            switch "foo bar" {
                                            case "":
                                                v4 += v4
                                            default:
                                                var v169 uint32 = 2147483647
                                                _ = v169
                                            }
                                        }
                                    } else {
                                        if v5 {
                                            if (0 > 189) {
                                                var v170 uint8 = 61
                                                for fi171 := 0; fi171 < 4; fi171++ {
                                                    if (!false) {
                                                        v172 := GenS25{F0: v1, F1: 1, F2: (32767 != 0)}
                                                        for fi173 := 0; fi173 < 1; fi173++ {
                                                            var v174 uint16 = 1
                                                            v174 += (v174 + v174)
                                                            var v175 uint8 = v170
                                                            _ = fi173
                                                            _ = v174
                                                            _ = v175
                                                        }
                                                        _ = v172
                                                    }
                                                    _ = fi171
                                                }
                                                fmt.Println(v2, v5)
                                                _ = v170
                                            }
                                            v1 -= v2
                                            for fi176 := 0; fi176 < 14; fi176++ {
                                                for fi177 := 0; fi177 < 13; fi177++ {
                                                    v178 := GenS27{F0: 9223372036854775807, F1: fi176}
                                                    _ = fi177
                                                    _ = v178
                                                }
                                                if len(v159) > 0 {
                                                    v159[0] = fi176
                                                }
                                                fmt.Println(v1)
                                                _ = fi176
                                            }
                                        } else {
                                            v179 := GenS3{F0: (v4 >> ((uint(v3)) & 63)), F1: (^(v4)), F2: 65535}
                                            switch ("a" + (("a" + "") + "")) {
                                            case "\n\t":
                                                fmt.Println(v4)
                                            case "":
                                                v1 += v2
                                            case "hello":
                                                if (!v5) {
                                                    v2 -= 127
                                                    v180 := GenS4{F0: 73, F1: v2, F2: 42}
                                                    v2 -= v4
                                                    _ = v180
                                                } else {
                                                    if (32767 >= 1000) {
                                                        switch 1 {
                                                        case 32768:
                                                            var v181 int = 374
                                                            _ = v181
                                                        }
                                                    }
                                                }
                                            default:
                                                var v182 uint32 = 42
                                                _ = v182
                                            }
                                            var v183 int = -158
                                            _ = v179
                                            _ = v183
                                        }
                                        if v5 {
                                            switch "hello" {
                                            case "hello":
                                                v3 = (v3 + 255)
                                            case "":
                                                v159 = append(v159, 127)
                                            }
                                            for fi184 := 0; fi184 < 14; fi184++ {
                                                var v185 uint64 = 0
                                                _ = fi184
                                                _ = v185
                                            }
                                            _ = len(v159)
                                        } else {
                                            if len(v159) > 0 {
                                                v159[0] = 2147483647
                                            }
                                            if (v5 || v5) {
                                                v1 += (v2 ^ v4)
                                            } else {
                                                if v5 {
                                                    if ((^(v3)) == 127) {
                                                        fmt.Println(v5, v1)
                                                        var v186 int8 = (-(v1))
                                                        _ = v186
                                                    } else {
                                                        for fi187 := 0; fi187 < 14; fi187++ {
                                                            switch "a" {
                                                            case "":
                                                                v5 = (0 == 0)
                                                            default:
                                                                switch (fi187 % ((128) | 1)) {
                                                                case 0:
                                                                    v188 := GenS21{F0: 0, F1: 352}
                                                                    _ = v188
                                                                case 1:
                                                                    var v189 string = "foo bar"
                                                                    _ = v189
                                                                }
                                                            }
                                                            if v5 {
                                                                for fi190 := 0; fi190 < 4; fi190++ {
                                                                    v2 -= v2
                                                                    if true {
                                                                        if len(v159) > 0 {
                                                                            v159[0] = 4294967295
                                                                        }
                                                                        v191 := GenS11{F0: 0, F1: (v4 + int8(v3))}
                                                                        fmt.Println(v5, v4)
                                                                        var v192 uint = 2
                                                                        _ = v191
                                                                        _ = v192
                                                                    } else {
                                                                        var v193 float64 = v6
                                                                        var v194 int32 = -2147483648
                                                                        fmt.Println(v5, v3, v2)
                                                                        _ = v193
                                                                        _ = v194
                                                                    }
                                                                    var v195 int32 = -2147483648
                                                                    _ = fi190
                                                                    _ = v195
                                                                }
                                                            } else {
                                                                if v5 {
                                                                    var v196 uint32 = 1
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v197 int64 = -332
                                                                        _ = v197
                                                                    case "hello":
                                                                        v159 = append(v159, (fi187 | (-(fi187))))
                                                                    default:
                                                                        v4 += v4
                                                                    }
                                                                    v198 := GenS21{F0: 205, F1: 18446744073709551615}
                                                                    _ = v196
                                                                    _ = v198
                                                                } else {
                                                                    var v199 uint = 205
                                                                    _ = v199
                                                                }
                                                                for fi200 := 0; fi200 < 20; fi200++ {
                                                                    v4 = 42
                                                                    v1 -= (^(v4))
                                                                    v1 = v1
                                                                    _ = fi200
                                                                }
                                                            }
                                                            var v201 int8 = (v2 & (v1 >> ((42) & 63)))
                                                            _ = fi187
                                                            _ = v201
                                                        }
                                                    }
                                                } else {
                                                    for fi202 := 0; fi202 < 7; fi202++ {
                                                        for fi203 := 0; fi203 < 8; fi203++ {
                                                            var v204 float64 = 3.14159
                                                            _ = fi203
                                                            _ = v204
                                                        }
                                                        _ = fi202
                                                    }
                                                }
                                            }
                                            fmt.Println(v2, v3)
                                        }
                                    }
                                    var v205 int8 = 100
                                    _ = v168
                                    _ = v205
                                }
                                v5 = v5
                                _ = v159
                            }
                            fmt.Println(v5)
                        } else {
                            var v206 float64 = v6
                            _ = v206
                        }
                        _ = v28
                    }
                    v3 += (^(v3))
                } else {
                    v1 = v4
                    var v207 int8 = v4
                    _ = v207
                }
            }
        }
    case 9223372036854775807:
        fmt.Println(v2)
    default:
        v4 = (v2 | v1)
    }
    v6 = v6
    v208 := GenS13{F0: 18446744073709551615, F1: 40}
    switch "a" {
    case "a":
        v209 := []int{0}
        _ = v209
    case "hello":
        v210 := GenS15{F0: 42, F1: (v4 / ((v1) | 1)), F2: -624}
        _ = v210
    case "":
        var v211 int32 = -2147483648
        _ = v211
    }
    var v212 int64 = 9223372036854775807
    v4 = (v2 * v4)
    v213 := []int{}
    fmt.Println(v212)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v212:", v212)
    fmt.Println("slicelen v213:", len(v213), v213)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v208
    _ = v212
    _ = v213
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
