package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 int; F2 int }
func genFn0(p0 uint8, p1 string) int16 {
    switch (p1 + ((p1 + "foo bar") + ((("\n\t" + (p1 + p1)) + (p1 + p1)) + p1))) {
    case "":
        fmt.Println(p0)
    }
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(p1)
        var v5 uint = 2
        fmt.Println(p0)
        _ = fi4
        _ = v5
    }
    v6 := []int{2147483647, -9223372036854775808, 0}
    p0 = p0
    switch 1 {
    case 1:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    return 1000
}
func genFn1() uint16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint32 = 2147483647
        _ = fi8
        _ = v9
    }
    if ((2147483647 > 139) && false) {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int16 = genFn0(2, ("" + (("hello" + ("a" + (("a" + ("foo bar" + "")) + ""))) + (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "hello"))))
            for fi12 := 0; fi12 < 5; fi12++ {
                v11 -= 1000
                v13 := GenS1{F0: true}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(v11, fi14)
                _ = fi14
            }
            _ = fi10
            _ = v11
        }
        v7 = append(v7, 32767)
        if (!true) {
            if ((0.1 + 0.1) == (((((((3.14159 + 1e10) * ((0.0 - 3.14159) + 0.1)) - 0.1) - float64(-2.25)) + (3.14159 / 0.1)) - 0.0) / ((1e10 - 0.1) / (3.14159 + (1.5 - 3.14159))))) {
                var v15 int8 = 100
                var v16 uint64 = 4294967295
                v17 := GenS2{F0: 702, F1: -927, F2: (v16 + v16)}
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 int64 = -328
                _ = v18
            }
        } else {
            switch ("hello" + "a") {
            case "foo bar":
                for fi19 := 0; fi19 < 8; fi19++ {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if (42 >= -45) {
                            if (!false) {
                                if (((3.14159 * 0.1) <= 1.5) && (18446744073709551615 == 4294967295)) {
                                    fmt.Println(fi20, fi19)
                                    v21 := GenS2{F0: 275, F1: 1, F2: 0}
                                    var v22 int32 = -362
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                if false {
                                    if (fi19 == fi20) {
                                        v7 = append(v7, fi19)
                                        switch "a" {
                                        case "":
                                            v23 := GenS2{F0: 65535, F1: genFn0(255, ("foo bar" + ("a" + "hello"))), F2: 1}
                                            var v24 float32 = ((((1e10 + (3.14159 / 1.0)) / 1.5) - 0.1) / -2.25)
                                            var v25 int = 2147483647
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        default:
                                            var v26 int32 = 1
                                            _ = v26
                                        }
                                    } else {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        _ = len(v7)
                                        v7 = append(v7, 42)
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        var v27 int32 = -338
                                        _ = v27
                                    case "foo bar":
                                        v7 = append(v7, fi19)
                                    case "":
                                        var v28 int16 = -559
                                        if true {
                                            v28 += v28
                                            v7 = append(v7, 4294967295)
                                            fmt.Println(fi19, fi19)
                                            var v29 float64 = 1.5
                                            _ = v29
                                        } else {
                                            switch ("" + ("hello" + ("\n\t" + "\n\t"))) {
                                            case "foo bar":
                                                v30 := GenS3{F0: -117, F1: true}
                                                _ = v30
                                            case "":
                                                var v31 int32 = 623
                                                _ = v31
                                            case "\n\t":
                                                v28 = v28
                                            }
                                        }
                                        _ = v28
                                    }
                                    _ = len(v7)
                                    if false {
                                        var v32 float64 = 1e10
                                        _ = v32
                                    }
                                    if (128 >= 173) {
                                        var v33 int64 = -9223372036854775808
                                        _ = v33
                                    }
                                    if false {
                                        v34 := GenS4{F0: 42, F1: -814}
                                        var v35 uint64 = 720
                                        var v36 int8 = 42
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                v37 := GenS4{F0: 130, F1: -32768}
                                _ = v37
                            }
                            v7 = append(v7, 2147483647)
                        } else {
                            var v38 string = ("hello" + "")
                            if true {
                                v38 = (v38 + ((v38 + (v38 + (v38 + "foo bar"))) + v38))
                                v38 = "a"
                            }
                            v39 := GenS5{F0: 2147483647, F1: v38}
                            _ = v38
                            _ = v39
                        }
                        var v40 uint32 = 335
                        _ = fi20
                        _ = v40
                    }
                    _ = fi19
                }
            case "":
                switch 1 {
                case 728:
                    switch "a" {
                    case "a":
                        if len(v7) > 0 {
                            v7[0] = -2147483648
                        }
                    case "\n\t":
                        for fi41 := 0; fi41 < 2; fi41++ {
                            if ((3.14159 - 0.0) <= 1e10) {
                                if true {
                                    v7 = append(v7, (fi41 >> ((255) & 63)))
                                } else {
                                    v7 = append(v7, (fi41 / (((fi41 | fi41)) | 1)))
                                    if ((0.0 * 3.14159) == -2.25) {
                                        if (0 < 2) {
                                            if (false && false) {
                                                var v42 int64 = 9223372036854775807
                                                if false {
                                                    for fi43 := 0; fi43 < 9; fi43++ {
                                                        v42 = (v42 * 62)
                                                        _ = fi43
                                                    }
                                                    var v44 bool = false
                                                    _ = v44
                                                }
                                                fmt.Println(v42, fi41)
                                                _ = v42
                                            }
                                            for fi45 := 0; fi45 < 17; fi45++ {
                                                if ((float32(1.5) * 1.5) < float32(((1.5 - 1.5) - 0.0))) {
                                                    var v46 int = (fi41 & (fi41 * fi45))
                                                    v46 = -119
                                                    if len(v7) > 0 {
                                                        v7[0] = 2147483647
                                                    }
                                                    _ = v46
                                                } else {
                                                    var v47 uint64 = 317
                                                    v48 := GenS1{F0: ("hello" != "hello")}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                if true {
                                                    v49 := GenS2{F0: 2, F1: 471, F2: 995}
                                                    _ = v49
                                                } else {
                                                    v50 := GenS1{F0: (0 >= 255)}
                                                    var v51 int64 = 972
                                                    _ = v50
                                                    _ = v51
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = fi45
                                                }
                                                _ = fi45
                                            }
                                        } else {
                                            v52 := GenS5{F0: -1, F1: "hello"}
                                            for fi53 := 0; fi53 < 13; fi53++ {
                                                fmt.Println(fi41, fi53)
                                                var v54 int = fi53
                                                if true {
                                                    v54 = (v54 % (((-(fi53))) | 1))
                                                    v55 := GenS6{F0: "foo bar", F1: 0}
                                                    v54 -= fi53
                                                    _ = v55
                                                }
                                                _ = fi53
                                                _ = v54
                                            }
                                            _ = v52
                                        }
                                    }
                                    _ = len(v7)
                                }
                                switch "hello" {
                                case "":
                                    fmt.Println(fi41)
                                case "a":
                                    var v56 int64 = 9223372036854775807
                                    _ = v56
                                case "foo bar":
                                    v57 := GenS4{F0: 1, F1: 1}
                                    _ = v57
                                default:
                                    v58 := GenS1{F0: ((float32(1e10) == 0.0) || (!(1.5 != float64(1.5))))}
                                    _ = v58
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            } else {
                                fmt.Println(fi41)
                                for fi59 := 0; fi59 < 8; fi59++ {
                                    var v60 uint = 0
                                    switch (("foo bar" + "a") + "") {
                                    case "\n\t":
                                        v60 = (^(v60))
                                    case "":
                                        v61 := GenS7{F0: 0, F1: (^(v60))}
                                        _ = v61
                                    }
                                    _ = fi59
                                    _ = v60
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            if (-9223372036854775808 == 980) {
                                var v62 int16 = -1
                                for fi63 := 0; fi63 < 5; fi63++ {
                                    var v64 uint8 = 255
                                    if len(v7) > 0 {
                                        v7[0] = (fi41 % ((42) | 1))
                                    }
                                    _ = fi63
                                    _ = v64
                                }
                                v62 = (v62 + v62)
                                _ = v62
                            }
                            if (-1 > -128) {
                                if (false || (!(1 <= 42))) {
                                    var v65 float64 = 0.0
                                    _ = v65
                                }
                                var v66 int64 = -9223372036854775808
                                var v67 uint32 = 2147483648
                                _ = v66
                                _ = v67
                            }
                            _ = fi41
                        }
                    }
                case 32767:
                    var v68 float64 = float64(0.0)
                    _ = v68
                case 0:
                    if (false || true) {
                        var v69 bool = (!true)
                        for fi70 := 0; fi70 < 1; fi70++ {
                            if len(v7) > 0 {
                                v7[0] = fi70
                            }
                            _ = fi70
                        }
                        _ = v69
                    }
                default:
                    if len(v7) > 0 {
                        v7[0] = 9223372036854775807
                    }
                }
            case "a":
                switch -558 {
                case 0:
                    if true {
                        var v71 float64 = (0.1 - (3.14159 - 0.1))
                        _ = v71
                    } else {
                        v72 := GenS5{F0: 59, F1: "hello"}
                        var v73 uint16 = 191
                        switch 446 {
                        case 459:
                            switch ("" + "hello") {
                            case "a":
                                fmt.Println(v73)
                            }
                        default:
                            v74 := GenS8{F0: (v73 + 476), F1: 42}
                            _ = v74
                        }
                        _ = v72
                        _ = v73
                    }
                }
            }
        }
    } else {
        if (1e10 < 3.14159) {
            var v75 uint32 = 90
            var v76 int8 = -1
            _ = v75
            _ = v76
        }
    }
    _ = v7
    return 931
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2147483647
    var v3 uint8 = 255
    var v77 uint64 = 298
    if (100 > -1) {
        fmt.Println(v3)
        var v78 bool = (0.0 > (((3.14159 + ((float32(1e10) * (0.0 - 0.1)) - 0.0)) - -2.25) * (float32(float64((float32(1.5) * (0.0 - 0.0)))) - 1.5)))
        var v79 string = "a"
        _ = v78
        _ = v79
    }
    if ((v77 & (v77 - (v77 * 1))) == (^(v77))) {
        var v80 string = ((("\n\t" + "a") + "") + ("hello" + ("foo bar" + "a")))
        v2 = (^(v2))
        _ = v80
    } else {
        v81 := GenS9{F0: 42, F1: (v77 >> ((42) & 63))}
        v82 := GenS10{F0: int32(v2)}
        var v83 uint = 350
        _ = v81
        _ = v82
        _ = v83
    }
    for fi84 := 0; fi84 < 6; fi84++ {
        var v85 uint32 = v2
        _ = fi84
        _ = v85
    }
    var v86 int16 = 1000
    for fi87 := 0; fi87 < 4; fi87++ {
        if false {
            var v88 int32 = -782
            _ = v88
        } else {
            var v89 bool = true
            _ = v89
        }
        if true {
            v90 := GenS11{F0: true}
            for fi91 := 0; fi91 < 17; fi91++ {
                switch (fi87 & -9223372036854775808) {
                case -788:
                    v92 := []int{fi91, fi87}
                    _ = v92
                }
                _ = fi91
            }
            var v93 int32 = 749
            _ = v90
            _ = v93
        }
        var v94 uint32 = 0
        _ = fi87
        _ = v94
    }
    for fi95 := 0; fi95 < 15; fi95++ {
        v3 -= (v3 * (v3 + (^(v3))))
        v77 = (v77 - 0)
        v96 := GenS12{F0: "hello", F1: 0}
        _ = fi95
        _ = v96
    }
    v2 = v2
    var v97 uint32 = uint32(v3)
    v98 := []int{}
    for fi99 := 0; fi99 < 10; fi99++ {
        var v100 uint16 = uint16(v77)
        _ = fi99
        _ = v100
    }
    if true {
        fmt.Println(v97, v97, v97)
        v101 := GenS7{F0: -779, F1: 18446744073709551615}
        v97 += (v97 % (((^(v97))) | 1))
        _ = v101
    } else {
        if (!false) {
            if true {
                v102 := v1[127]
                for fi103 := 0; fi103 < 9; fi103++ {
                    _ = len(v98)
                    _ = fi103
                }
                switch v102 {
                case "a":
                    var v104 float64 = 1e10
                    _ = v104
                case "":
                    var v105 int = 32768
                    _ = v105
                case "\n\t":
                    _ = len(v1)
                default:
                    switch v102 {
                    case "":
                        if false {
                            var v106 int64 = -152
                            _ = v106
                        } else {
                            var v107 int32 = 475
                            v102 = "foo bar"
                            fmt.Println(v107)
                            _ = v107
                        }
                    case "hello":
                        v77 = (^(v77))
                    case "a":
                        v1[127] = v102
                    }
                }
                _ = v102
            } else {
                v86 -= (-(v86))
                var v108 uint64 = v77
                var v109 int64 = 0
                _ = v108
                _ = v109
            }
            var v110 bool = (0 == 0)
            fmt.Println(v77, v3)
            _ = v110
        } else {
            var v111 string = "\n\t"
            var v112 uint16 = 2
            _ = v111
            _ = v112
        }
        if true {
            fmt.Println(v86, v3)
            v113 := GenS13{F0: 42}
            _ = v113
        } else {
            v114 := GenS14{F0: false, F1: -9223372036854775808, F2: 128}
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v97)
            case 800:
                if (((v97 % ((v97) | 1)) < 4294967295) || true) {
                    fmt.Println(v86)
                    v115 := GenS4{F0: 113, F1: (v86 | (v86 | 282))}
                    _ = v115
                }
            case -251:
                var v116 int16 = v86
                _ = v116
            }
            _ = v114
        }
        var v117 string = ""
        _ = v117
    }
    for fi118 := 0; fi118 < 3; fi118++ {
        for fi119 := 0; fi119 < 5; fi119++ {
            v98 = append(v98, 127)
            _ = fi119
        }
        _ = fi118
    }
    _ = len(v98)
    var v120 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v86:", v86)
    fmt.Println("v97:", v97)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v77
    _ = v86
    _ = v97
    _ = v98
}
