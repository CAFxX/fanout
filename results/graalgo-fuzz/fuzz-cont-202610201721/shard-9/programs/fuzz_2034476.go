package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 int }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int16; F1 uint8; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint64; F2 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int; F1 int32; F2 int32 }

func main() {
    var v1 uint8 = 128
    var v2 uint = 1
    var v3 bool = (2147483647 != 42)
    v4 := make(map[string]int)
    var v5 uint32 = 2147483647
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 16; fi7++ {
        if (((3.14159 * float32((1.5 + (1.5 - (0.0 - 1e10))))) <= 1e10) && true) {
            v8 := GenS1{F0: -1, F1: v2, F2: (v1 - v1)}
            if v3 {
                v2 = uint(v5)
                switch ("a" + ("hello" + "hello")) {
                case "":
                    switch "" {
                    case "\n\t":
                        if (v3 && (!v3)) {
                            v5 = v5
                        } else {
                            switch (fi7 - fi7) {
                            case 127:
                                var v9 int64 = -246
                                _ = v9
                            case 4294967295:
                                v10 := []int{fi7}
                                _ = v10
                            }
                            if false {
                                var v11 int8 = int8(v2)
                                switch -1 {
                                case 4294967295:
                                    v12 := GenS2{F0: ("hello" + (("hello" + "\n\t") + "hello")), F1: (v11 / (((v11 - -128)) | 1))}
                                    _ = v12
                                case -2147483648:
                                    if v3 {
                                        v13 := GenS2{F0: "", F1: v11}
                                        v1 += (v1 & (v1 >> (((v2 * 0)) & 63)))
                                        _ = v13
                                    }
                                case 144:
                                    v14 := []int{}
                                    _ = v14
                                default:
                                    var v15 uint8 = v1
                                    _ = v15
                                }
                                var v16 int64 = -9223372036854775808
                                _ = v11
                                _ = v16
                            }
                        }
                    }
                case "foo bar":
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if (494 >= 0) {
                            if v3 {
                                var v18 float32 = ((((1e10 * float32(1.5)) + 0.1) - -2.25) - (3.14159 + ((((0.1 + (1.5 / 3.14159)) - -2.25) * -2.25) / (1e10 - (((0.0 / (float32(0.1) - ((0.0 + 1e10) + (-2.25 * 0.1)))) * (0.0 + (0.0 / 3.14159))) + (-2.25 - (0.0 / 1e10)))))))
                                if (((-128 == 0) && v3) && true) {
                                    if ((uint32(v1) > (v5 | v5)) && ((("a" + "a") + "a") <= "foo bar")) {
                                        v19 := GenS1{F0: 347, F1: (v2 & v2), F2: (v1 + (v1 | v1))}
                                        if false {
                                            var v20 uint16 = 65535
                                            var v21 int64 = 0
                                            if (7 < int8(v20)) {
                                                v22 := GenS2{F0: "\n\t", F1: 42}
                                                _ = v22
                                            } else {
                                                v18 = float32(1e10)
                                            }
                                            v20 = 2
                                            v1 -= 127
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v23 := GenS1{F0: -408, F1: (v2 | 42), F2: (v1 / (((v1 % ((v1) | 1))) | 1))}
                                            _ = v23
                                        }
                                        var v24 int8 = 42
                                        _ = v19
                                        _ = v24
                                    } else {
                                        fmt.Println(fi17)
                                    }
                                    var v25 uint8 = (v1 << (((v2 & 42)) & 63))
                                    if v3 {
                                        v26 := []int{fi17}
                                        _ = v26
                                    }
                                    _ = v25
                                } else {
                                    v18 = v18
                                }
                                v27 := GenS2{F0: "", F1: 0}
                                _ = v18
                                _ = v27
                            } else {
                                v28 := GenS3{F0: true, F1: 255, F2: uint32(v2)}
                                switch fi17 {
                                case 32767:
                                    v29 := GenS2{F0: "hello", F1: 0}
                                    _ = v29
                                }
                                _ = v28
                            }
                            v3 = (42 >= 100)
                        }
                        delete(v4, (("\n\t" + "a") + "a"))
                        _ = fi17
                    }
                }
            }
            _ = v8
        }
        v30 := []int{128, fi7, (^(fi7)), (fi7 - fi7)}
        _ = fi7
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v3 = (v3 && (3.14159 == (((0.0 - (-2.25 / (0.1 - float32((float64(-2.25) * (float64(-2.25) - 1e10)))))) / ((float32(float64(((((1.5 - ((0.0 - (float32((3.14159 / -2.25)) / 1.0)) / float32(-2.25))) + 1e10) / (1e10 - 3.14159)) / 1.0))) - 0.1) - 1.5)) - 1.5)))
        if (-2147483648 == 369) {
            for fi32 := 0; fi32 < 8; fi32++ {
                var v33 uint16 = 411
                var v34 uint = (v2 + (v2 & 0))
                _ = fi32
                _ = v33
                _ = v34
            }
            var v35 float32 = 0.1
            var v36 int64 = -1
            _ = v35
            _ = v36
        }
        v5 += 468
        _ = fi31
    }
    switch ((("foo bar" + "") + "a") + "a") {
    case "\n\t":
        v37 := GenS4{F0: (0 > 2), F1: 0, F2: 42}
        _ = v37
    case "hello":
        if ("a" >= ("\n\t" + "")) {
            for fi38 := 0; fi38 < 5; fi38++ {
                var v39 uint32 = (v5 ^ (^(v5)))
                fmt.Println(v1)
                _ = fi38
                _ = v39
            }
            v3 = (585 < 0)
        }
    case "":
        v40 := GenS2{F0: "hello", F1: 42}
        _ = v40
    }
    switch (("a" + ("\n\t" + ("" + "a"))) + "foo bar") {
    case "foo bar":
        v41 := []int{-321}
        _ = v41
    case "\n\t":
        if ((v1 | v1) < v1) {
            switch 1 {
            case 128:
                v6[v5] = ("hello" + "")
            case 127:
                _ = len(v4)
            case 962:
                var v42 int8 = 1
                _ = v42
            }
        } else {
            v43 := GenS5{F0: -913, F1: (v2 << ((198) & 63)), F2: -1}
            _ = v43
        }
    case "hello":
        v44 := GenS5{F0: -61, F1: 812, F2: 0}
        _ = v44
    default:
        var v45 int = -127
        _ = v45
    }
    if false {
        var v46 int32 = -1
        _ = v46
    } else {
        var v47 uint64 = 2
        _ = v47
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        if (18446744073709551615 >= 1) {
            var v49 float32 = 3.14159
            if v3 {
                var v50 uint = (v2 / (((^(v2))) | 1))
                for fi51 := 0; fi51 < 3; fi51++ {
                    v49 = (0.0 + (1.5 * float32((1.5 / 1.5))))
                    _ = fi51
                }
                var v52 uint16 = 255
                _ = v50
                _ = v52
            } else {
                if (42 != 4294967295) {
                    switch -1 {
                    case 9223372036854775807:
                        for fi53 := 0; fi53 < 15; fi53++ {
                            var v54 float32 = v49
                            if true {
                                switch fi53 {
                                case 127:
                                    switch fi53 {
                                    case 32768:
                                        var v55 uint64 = 274
                                        _ = v55
                                    case 42:
                                        if true {
                                            var v56 int = (fi53 >> ((255) & 63))
                                            v5 = (v5 | v5)
                                            switch fi53 {
                                            case -2147483648:
                                                if v3 {
                                                    v57 := GenS5{F0: 42, F1: 65535, F2: int8(v1)}
                                                    switch ("a" + (("" + "\n\t") + "hello")) {
                                                    case "a":
                                                        v54 = (v54 / v54)
                                                    case "\n\t":
                                                        v54 = (1.5 / v54)
                                                    case "":
                                                        v49 += ((v49 + (v49 / ((1e10 * ((v49 / (v49 + (v49 + 3.14159))) - (((v54 / 1.5) * 0.1) / -2.25))) + v54))) * (3.14159 * (1e10 + v54)))
                                                    }
                                                    v5 = uint32(v2)
                                                    _ = v57
                                                } else {
                                                    if v3 {
                                                        v58 := GenS5{F0: 42, F1: (^(v2)), F2: -1}
                                                        v59 := GenS6{F0: 9223372036854775807, F1: "foo bar"}
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        var v60 int16 = 42
                                                        _ = v60
                                                    }
                                                    v2 = 1
                                                }
                                            default:
                                                var v61 uint64 = uint64(v5)
                                                _ = v61
                                            }
                                            _ = v56
                                        } else {
                                            v62 := GenS7{F0: 42, F1: int64(v1)}
                                            v63 := GenS3{F0: (true || v3), F1: 2, F2: (^(v5))}
                                            _ = v62
                                            _ = v63
                                        }
                                    case 2147483647:
                                        v3 = (42 > 1)
                                    }
                                case -9223372036854775808:
                                    delete(v6, 4294967295)
                                }
                            }
                            _ = fi53
                            _ = v54
                        }
                    }
                    for fi64 := 0; fi64 < 18; fi64++ {
                        v5 = v5
                        _ = fi64
                    }
                    v65 := GenS4{F0: (v3 || false), F1: 18446744073709551615, F2: 32767}
                    _ = v65
                } else {
                    switch fi48 {
                    case 67:
                        fmt.Println(fi48, v1)
                    case -363:
                        v2 += (v2 - v2)
                    case -9223372036854775808:
                        if (42 >= -1) {
                            for fi66 := 0; fi66 < 18; fi66++ {
                                var v67 uint32 = uint32(v2)
                                _ = fi66
                                _ = v67
                            }
                        } else {
                            fmt.Println(fi48, fi48, v3)
                            for fi68 := 0; fi68 < 10; fi68++ {
                                v1 += 127
                                fmt.Println(v49, v3)
                                _ = fi68
                            }
                        }
                    }
                    v69 := []int{}
                    _ = v69
                }
                if v3 {
                    v2 = 100
                    if (0.0 >= (0.0 * (-2.25 * float64((v49 + v49))))) {
                        switch fi48 {
                        case 127:
                            fmt.Println(v2, fi48)
                        case 944:
                            var v70 int8 = 83
                            _ = v70
                        }
                    } else {
                        for fi71 := 0; fi71 < 13; fi71++ {
                            v5 -= uint32(v2)
                            if false {
                                v72 := GenS6{F0: 0, F1: ("" + "foo bar")}
                                v49 = (v49 / 1.0)
                                _ = v72
                            } else {
                                v2 += 783
                                var v73 uint64 = 42
                                _ = v73
                            }
                            var v74 int64 = -706
                            _ = fi71
                            _ = v74
                        }
                    }
                } else {
                    var v75 bool = (-1 < 42)
                    delete(v6, (v5 % ((0) | 1)))
                    _ = v75
                }
            }
            _ = v49
        }
        switch fi48 {
        case 369:
            var v76 uint16 = 255
            _ = v76
        case -692:
            v77 := GenS8{F0: 580, F1: (^(v1)), F2: 6}
            _ = v77
        }
        _ = fi48
    }
    v78 := GenS9{F0: 2, F1: 4294967295, F2: (v3 && v3)}
    if (v1 < (v1 * v1)) {
        var v79 int = -800
        switch 32768 {
        case 32767:
            for fi80 := 0; fi80 < 14; fi80++ {
                v81 := GenS10{F0: 41}
                v82 := []int{586, 32768}
                for fi83 := 0; fi83 < 5; fi83++ {
                    switch (("hello" + "\n\t") + (("foo bar" + "") + "\n\t")) {
                    case "\n\t":
                        if (100 != -1) {
                            switch v79 {
                            case 128:
                                var v84 int32 = 0
                                _ = v84
                            case -926:
                                if len(v82) > 0 {
                                    v82[0] = (v79 + (v79 >> ((v2) & 63)))
                                }
                            }
                        }
                        var v85 float64 = ((1e10 * 1.5) + 3.14159)
                        v86 := GenS11{F0: (fi83 | (fi80 & int(v5))), F1: -530, F2: 1}
                        var v87 uint32 = v5
                        if false {
                            fmt.Println(fi83)
                            var v88 int8 = 127
                            var v89 uint16 = 1
                            _ = v88
                            _ = v89
                        } else {
                            fmt.Println(v5, fi83)
                        }
                        _ = v85
                        _ = v86
                        _ = v87
                    case "hello":
                        fmt.Println(v5, v1)
                    case "a":
                        var v90 int32 = 0
                        _ = v90
                    default:
                        v5 = v5
                    }
                    fmt.Println(fi83)
                    var v91 uint = (v2 << (((^(v2))) & 63))
                    _ = fi83
                    _ = v91
                }
                _ = fi80
                _ = v81
                _ = v82
            }
        case 0:
            var v92 uint8 = 1
            _ = v92
        case 32768:
            v1 = v1
        default:
            v93 := []int{(^(v79)), int(v1), 904, v79}
            _ = v93
        }
        if (32767 > 756) {
            fmt.Println(v2, v1)
            v94 := GenS6{F0: 0, F1: ("foo bar" + "foo bar")}
            v5 -= 642
            _ = v94
        }
        _ = v79
    }
    v6[v5] = ""
    if false {
        _ = len(v4)
        for fi95 := 0; fi95 < 15; fi95++ {
            delete(v6, v5)
            v5 = (v5 + v5)
            _ = fi95
        }
    } else {
        v96 := GenS2{F0: "foo bar", F1: -1}
        if true {
            v97 := GenS4{F0: v3, F1: 165, F2: -702}
            for fi98 := 0; fi98 < 3; fi98++ {
                var v99 float64 = 0.1
                var v100 int = 127
                var v101 int = fi98
                _ = fi98
                _ = v100
                _ = v101
                _ = v99
            }
            _ = v97
        } else {
            v102 := GenS7{F0: 933, F1: -9223372036854775808}
            v1 -= (^(v1))
            _ = v102
        }
        v1 -= v1
        _ = v96
    }
    var v103 float32 = -2.25
    var v104 uint = (v2 | v2)
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
