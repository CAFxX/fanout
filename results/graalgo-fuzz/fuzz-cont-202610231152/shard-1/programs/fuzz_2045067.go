package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 string }
type GenS5 struct { F0 int32; F1 int; F2 uint }
type GenS6 struct { F0 uint16; F1 uint; F2 uint }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint }
type GenS9 struct { F0 string }

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 725
    v3 := make(map[uint32]string)
    if true {
        v2 -= v2
    }
    fmt.Println(v1)
    v4 := GenS1{F0: -1, F1: 100}
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v2)
        switch "hello" {
        case "\n\t":
            if false {
                switch fi5 {
                case -640:
                    var v6 int16 = 1000
                    _ = v6
                case 0:
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 uint32 = 2147483648
                        v9 := GenS1{F0: 1, F1: 127}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                default:
                    var v10 uint64 = 4294967295
                    _ = v10
                }
                if true {
                    var v11 string = ""
                    var v12 uint = 833
                    _ = v11
                    _ = v12
                }
            } else {
                v1 -= v1
            }
        case "":
            var v13 uint32 = 400
            _ = v13
        default:
            var v14 int16 = 1000
            _ = v14
        }
        _ = fi5
    }
    delete(v3, 65535)
    _ = len(v3)
    v15 := GenS1{F0: -2147483648, F1: 108}
    v2 += v2
    switch "hello" {
    case "":
        switch 0 {
        case -297:
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS1{F0: 42, F1: -128}
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := []int{(^(fi18))}
                    if true {
                        var v20 string = ""
                        v21 := v3[uint32(v2)]
                        _ = v20
                        _ = v21
                    } else {
                        _ = len(v3)
                        v2 = 0
                        v22 := GenS2{F0: fi16, F1: 1}
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
        default:
            v1 -= (^(v1))
        }
    case "a":
        fmt.Println(v1, v2)
    default:
        if true {
            if (1 > 564) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS3{F0: 42, F1: -111}
                    var v25 bool = ((3.14159 < (1.5 - float32(0.1))) || false)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                switch ((("foo bar" + "foo bar") + ("a" + "")) + "a") {
                case "foo bar":
                    if true {
                        v2 = v2
                    } else {
                        switch 168 {
                        case -383:
                            v26 := []int{}
                            _ = v26
                        case 1:
                            var v27 uint8 = 2
                            _ = v27
                        case -420:
                            if (false && ((v2 | (v2 / ((32767) | 1))) != (v2 * (v2 % ((v2) | 1))))) {
                                var v28 uint = 18446744073709551615
                                fmt.Println(v1, v1)
                                _ = v28
                            } else {
                                delete(v3, 65535)
                            }
                        }
                    }
                case "a":
                    if true {
                        var v29 string = "a"
                        v30 := []int{32768}
                        _ = v29
                        _ = v30
                    } else {
                        if (1.5 <= 1e10) {
                            v2 += (v2 + 32767)
                        }
                    }
                }
            } else {
                fmt.Println(v1)
                for fi31 := 0; fi31 < 8; fi31++ {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        var v33 bool = false
                        _ = fi32
                        _ = v33
                    }
                    _ = fi31
                }
                var v34 uint8 = (v1 >> ((2) & 63))
                _ = v34
            }
            v35 := GenS2{F0: 32767, F1: 384}
            _ = v35
        }
    }
    v2 = v2
    if true {
        if true {
            var v36 int16 = -32768
            _ = v36
        }
    } else {
        if true {
            v2 = (v2 + (v2 / (((v2 ^ v2)) | 1)))
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            for fi38 := 0; fi38 < 20; fi38++ {
                _ = len(v3)
                if ((4294967295 == 153) && true) {
                    var v39 uint = 4294967295
                    fmt.Println(v2)
                    switch "a" {
                    case "\n\t":
                        var v40 uint16 = 42
                        _ = v40
                    default:
                        v39 = (v39 % (((v39 << ((v39) & 63))) | 1))
                    }
                    _ = v39
                }
                if (true || false) {
                    if true {
                        var v41 int = 9223372036854775807
                        _ = v41
                    } else {
                        var v42 uint8 = v1
                        _ = v42
                    }
                    fmt.Println(v1)
                }
                _ = fi38
            }
            v43 := GenS1{F0: int32(fi37), F1: 42}
            v44 := v3[2147483647]
            _ = fi37
            _ = v43
            _ = v44
        }
        delete(v3, 2)
    }
    if ((v1 >> ((4294967295) & 63)) == v1) {
        var v45 uint16 = uint16(v1)
        for fi46 := 0; fi46 < 19; fi46++ {
            if (!(-2.25 < -2.25)) {
                v47 := v3[4294967295]
                _ = v47
            } else {
                v2 -= v45
            }
            switch 989 {
            case -285:
                for fi48 := 0; fi48 < 20; fi48++ {
                    v1 = (^(v1))
                    v49 := GenS1{F0: -1, F1: 1}
                    if true {
                        switch "a" {
                        case "":
                            v50 := GenS3{F0: -9223372036854775808, F1: 91}
                            _ = v50
                        case "foo bar":
                            if false {
                                if (9223372036854775807 >= 9223372036854775807) {
                                    var v51 uint8 = 127
                                    if (2 > 18446744073709551615) {
                                        fmt.Println(v1)
                                        if false {
                                            v52 := []int{918, fi48}
                                            v53 := GenS2{F0: (fi48 - fi46), F1: 980}
                                            _ = v52
                                            _ = v53
                                        } else {
                                            fmt.Println(v1, v1, v1)
                                            v45 = 42
                                            if (false && (0.1 == (-2.25 + 0.1))) {
                                                v54 := GenS3{F0: -1, F1: 852}
                                                v51 = v1
                                                if false {
                                                    var v55 int64 = 139
                                                    v56 := []int{(-(fi46)), -286}
                                                    var v57 int32 = -1
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                v58 := GenS4{F0: ("hello" + ("" + (("a" + "a") + "\n\t"))), F1: 0, F2: ""}
                                                v59 := GenS4{F0: "hello", F1: 226, F2: "a"}
                                                _ = v54
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v1 += 250
                                            }
                                            var v60 uint = 1
                                            _ = v60
                                        }
                                        var v61 int8 = 127
                                        _ = v61
                                    } else {
                                        v3[0] = "a"
                                        fmt.Println(v1)
                                    }
                                    v62 := GenS4{F0: (("foo bar" + "a") + ("" + "")), F1: 0, F2: "hello"}
                                    _ = v51
                                    _ = v62
                                }
                                if (((1e10 <= (1.5 / (3.14159 - (0.1 - 1e10)))) || true) || ((0.0 - (3.14159 - float64((0.0 + ((0.1 / (float32(((((-2.25 / 1.0) - -2.25) + 0.1) / 1.0)) + (-2.25 - (((1.5 / 1.0) + -2.25) / 0.1)))) - float32(1.5)))))) >= (((1e10 * (1.5 * -2.25)) + ((-2.25 / 1.0) / -2.25)) / (((1e10 - 3.14159) * 0.1) + 0.0)))) {
                                    switch (fi48 * 9223372036854775807) {
                                    case 0:
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(fi48)
                                        case "hello":
                                            v2 += 1
                                        case "foo bar":
                                            v63 := []int{}
                                            _ = v63
                                        }
                                    case 32767:
                                        v64 := []int{fi48, (fi46 ^ fi46), (fi48 % ((fi46) | 1))}
                                        _ = v64
                                    }
                                    var v65 uint16 = (v45 & (v2 ^ v2))
                                    var v66 int32 = 594
                                    _ = v65
                                    _ = v66
                                }
                                v1 += v1
                                v45 -= v45
                                if (313 != -32768) {
                                    if true {
                                        switch fi46 {
                                        case -1:
                                            v67 := GenS4{F0: "a", F1: (v2 & v2), F2: ""}
                                            _ = v67
                                        }
                                        delete(v3, 2147483647)
                                        v68 := []int{-1, fi46}
                                        _ = v68
                                    } else {
                                        fmt.Println(v45, fi48, fi48)
                                        if ((0.0 - 1e10) >= float32(1e10)) {
                                            v2 -= 1
                                            var v69 uint8 = (v1 * v1)
                                            _ = v69
                                        } else {
                                            v1 += (^(v1))
                                        }
                                    }
                                    var v70 uint = 18446744073709551615
                                    var v71 int32 = 1
                                    _ = v70
                                    _ = v71
                                }
                                fmt.Println(fi46)
                            } else {
                                switch ("" + ("" + (("foo bar" + ((("" + "a") + "foo bar") + "")) + "foo bar"))) {
                                case "\n\t":
                                    switch (fi46 << ((413) & 63)) {
                                    case -729:
                                        v1 -= (^(v1))
                                    case -498:
                                        var v72 uint = 0
                                        _ = v72
                                    case 127:
                                        v1 -= v1
                                    default:
                                        var v73 int64 = 261
                                        _ = v73
                                    }
                                default:
                                    v74 := GenS2{F0: -9223372036854775808, F1: 255}
                                    _ = v74
                                }
                            }
                        case "hello":
                            fmt.Println(fi46, fi48, v1)
                        default:
                            var v75 uint32 = 2147483648
                            switch fi46 {
                            case -9223372036854775808:
                                v76 := []int{(-(fi48)), fi46}
                                fmt.Println(v2, fi48)
                                var v77 bool = true
                                v2 = (^(v2))
                                _ = v76
                                _ = v77
                            default:
                                if (65535 < 255) {
                                    v75 += v75
                                    v78 := v3[(^(v75))]
                                    _ = v78
                                }
                            }
                            _ = v75
                        }
                        var v79 uint16 = (^(v45))
                        _ = v79
                    }
                    var v80 string = "\n\t"
                    _ = fi48
                    _ = v49
                    _ = v80
                }
            default:
                v1 += 128
            }
            _ = fi46
        }
        if true {
            fmt.Println(v1, v1, v2)
        } else {
            v81 := GenS5{F0: 0, F1: 444, F2: 2}
            switch 822 {
            case 9223372036854775807:
                v45 = (v45 * 889)
            case -2147483648:
                v82 := GenS6{F0: v45, F1: 4294967295, F2: 65535}
                _ = v82
            default:
                switch -773 {
                case 633:
                    for fi83 := 0; fi83 < 9; fi83++ {
                        var v84 bool = (fi83 < -569)
                        _ = fi83
                        _ = v84
                    }
                case 4294967295:
                    v45 -= (^(v45))
                case -287:
                    fmt.Println(v1)
                default:
                    var v85 float32 = (float32(-2.25) * 1.5)
                    _ = v85
                }
            }
            v86 := GenS7{F0: 42, F1: 42}
            _ = v81
            _ = v86
        }
        _ = v45
    } else {
        var v87 int32 = -26
        v88 := []int{713, 2147483647}
        if (2147483647 <= 564) {
            for fi89 := 0; fi89 < 20; fi89++ {
                if (1 < 18446744073709551615) {
                    fmt.Println(fi89)
                    delete(v3, 333)
                } else {
                    var v90 int = 0
                    _ = v90
                }
                _ = fi89
            }
        }
        _ = v87
        _ = v88
    }
    var v91 float64 = -2.25
    var v92 uint32 = 2
    var v93 uint64 = uint64(v92)
    if true {
        v1 = uint8(v2)
        for fi94 := 0; fi94 < 16; fi94++ {
            v91 += v91
            v1 = (^(v1))
            if (fi94 != (fi94 << ((255) & 63))) {
                delete(v3, v92)
                if true {
                    var v95 string = (("hello" + (((("\n\t" + "a") + "hello") + "foo bar") + "foo bar")) + ("\n\t" + ("" + "hello")))
                    v91 = v91
                    v1 -= v1
                    _ = v95
                } else {
                    fmt.Println(v92, v2, fi94)
                    v2 += (v2 % ((v2) | 1))
                    for fi96 := 0; fi96 < 5; fi96++ {
                        v97 := GenS5{F0: 317, F1: fi94, F2: 2}
                        _ = len(v3)
                        var v98 float64 = 0.1
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                }
            } else {
                for fi99 := 0; fi99 < 8; fi99++ {
                    var v100 float64 = v91
                    var v101 float32 = 0.1
                    _ = fi99
                    _ = v100
                    _ = v101
                }
            }
            _ = fi94
        }
        if true {
            var v102 uint64 = (v93 >> ((0) & 63))
            switch 32768 {
            case 4294967295:
                v103 := GenS6{F0: v2, F1: 819, F2: uint(v92)}
                _ = v103
            default:
                v104 := GenS8{F0: -720, F1: (v2 * v2), F2: 592}
                _ = v104
            }
            for fi105 := 0; fi105 < 18; fi105++ {
                v106 := GenS7{F0: (v1 / (((v1 / (((v1 - v1)) | 1))) | 1)), F1: -1}
                var v107 int = (fi105 + (fi105 ^ (fi105 % ((-2147483648) | 1))))
                var v108 string = "foo bar"
                _ = fi105
                _ = v106
                _ = v107
                _ = v108
            }
            _ = v102
        } else {
            if true {
                v109 := GenS2{F0: int(v1), F1: 1}
                _ = v109
            }
            if false {
                v93 += uint64(v1)
            } else {
                var v110 int32 = -1
                var v111 uint64 = 277
                _ = v110
                _ = v111
            }
            v92 = v92
        }
    } else {
        v2 += 2
    }
    for fi112 := 0; fi112 < 11; fi112++ {
        for fi113 := 0; fi113 < 7; fi113++ {
            var v114 int16 = 995
            _ = fi113
            _ = v114
        }
        v115 := GenS9{F0: "foo bar"}
        for fi116 := 0; fi116 < 1; fi116++ {
            v91 -= v91
            if true {
                v1 += 255
            }
            _ = fi116
        }
        _ = fi112
        _ = v115
    }
    var v117 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v117
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v91
    _ = v92
    _ = v93
}
