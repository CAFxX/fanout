package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int32; F1 uint }
func genFn0(p0 uint16) int16 {
    v6 := []int{704, 4294967295}
    v7 := GenS1{F0: 1}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 uint32) int32 {
    v8 := []int{128}
    switch -1 {
    case 127:
        p0 = (p0 * (p0 + 255))
    case 9223372036854775807:
        fmt.Println(p0, p1)
    }
    var v9 int8 = -123
    if false {
        var v10 int64 = -9223372036854775808
        switch -898 {
        case -837:
            switch "foo bar" {
            case "hello":
                p0 += p0
            case "":
                if false {
                    switch "hello" {
                    case "foo bar":
                        var v11 uint = p0
                        _ = v11
                    case "":
                        p1 = 475
                    case "a":
                        v12 := GenS2{F0: "a"}
                        _ = v12
                    default:
                        var v13 int = -348
                        _ = v13
                    }
                } else {
                    fmt.Println(v10, v10)
                    var v14 string = ("foo bar" + "\n\t")
                    if (879 > 598) {
                        var v15 uint = p0
                        _ = v15
                    } else {
                        var v16 string = ("foo bar" + (((v14 + "a") + "") + (v14 + v14)))
                        _ = v16
                    }
                    _ = v14
                }
            default:
                fmt.Println(p0, p1, v9)
            }
        case 32768:
            v17 := GenS3{F0: 42, F1: 0}
            _ = v17
        }
        _ = v10
    } else {
        var v18 uint8 = 1
        if (-2.25 >= 0.1) {
            p0 = (p0 + p0)
            v19 := GenS2{F0: "a"}
            _ = v19
        }
        _ = v18
    }
    var v20 float32 = 1e10
    _ = p0
    _ = p1
    _ = v20
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 619
    v3 := make(map[int]bool)
    var v4 uint16 = 255
    v5 := make(map[int]int)
    if false {
        fmt.Println(v2)
        fmt.Println(v2)
    }
    v21 := GenS4{F0: v2}
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int8 = -47
        var v24 uint32 = 2147483648
        v25 := GenS4{F0: 478}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    switch ("" + ("\n\t" + "a")) {
    case "a":
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int = 736
            _ = fi26
            _ = v27
        }
    }
    var v28 int = -292
    for fi29 := 0; fi29 < 8; fi29++ {
        if true {
            v30 := []int{42, fi29, (v28 * (v28 + 127)), fi29}
            fmt.Println(v2)
            _ = v30
        }
        switch (((("foo bar" + "") + ((("foo bar" + ("a" + ("\n\t" + ("\n\t" + ("hello" + ("\n\t" + "a")))))) + (("\n\t" + "a") + "")) + ("\n\t" + "foo bar"))) + ("foo bar" + "")) + (("\n\t" + ("\n\t" + "\n\t")) + "a")) {
        case "a":
            for fi31 := 0; fi31 < 2; fi31++ {
                if (32767 < genFn0((^(v4)))) {
                    v28 = (^(fi29))
                }
                var v32 int16 = int16(v2)
                _ = fi31
                _ = v32
            }
        case "":
            fmt.Println(fi29)
        case "foo bar":
            if false {
                if true {
                    if false {
                        fmt.Println(v2, v4)
                        var v33 int32 = -2147483648
                        v28 = 127
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 1; fi34++ {
                        var v35 float64 = (0.1 - (((float64(float32(0.1)) - float64(v1)) * 1e10) * float64(v1)))
                        var v36 int8 = 1
                        if true {
                            v37 := []int{fi29}
                            _ = len(v37)
                            var v38 int32 = 0
                            _ = v37
                            _ = v38
                        } else {
                            v35 += v35
                            var v39 float32 = v1
                            v40 := GenS5{F0: 9223372036854775807, F1: (false && false)}
                            _ = v39
                            _ = v40
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v41 := []int{(fi29 / ((-193) | 1)), 42}
                    _ = v41
                }
                v4 = (v4 * v4)
            } else {
                v42 := GenS2{F0: "hello"}
                fmt.Println(v2)
                v1 = (v1 * (1.5 / v1))
                _ = v42
            }
        }
        if (true && false) {
            fmt.Println(v28, fi29, v1)
        } else {
            for fi43 := 0; fi43 < 4; fi43++ {
                fmt.Println(v2, v4)
                _ = fi43
            }
            if false {
                var v44 float32 = 1e10
                switch (^(fi29)) {
                case -752:
                    var v45 int = (fi29 % ((fi29) | 1))
                    _ = v45
                case 0:
                    var v46 uint32 = 2
                    _ = v46
                case -800:
                    var v47 string = "hello"
                    _ = v47
                }
                _ = v44
            }
            var v48 int = (v28 * fi29)
            _ = v48
        }
        _ = fi29
    }
    if true {
        v49 := []int{(v28 & (^(v28))), (^(v28))}
        for fi50 := 0; fi50 < 14; fi50++ {
            var v51 string = "hello"
            _ = fi50
            _ = v51
        }
        _ = v49
    }
    var v52 uint = 1
    if true {
        for fi53 := 0; fi53 < 16; fi53++ {
            if (-1 <= 825) {
                var v54 uint64 = 18446744073709551615
                v55 := GenS5{F0: -1, F1: false}
                _ = v54
                _ = v55
            } else {
                var v56 uint16 = v4
                switch "foo bar" {
                case "\n\t":
                    if true {
                        v1 = float32((((((1e10 - -2.25) + 3.14159) / ((0.1 * 1e10) - 3.14159)) * 3.14159) - float64(1.5)))
                        if false {
                            v57 := []int{(v28 + v28), (v28 * (^(v28)))}
                            if (false || true) {
                                v28 -= -668
                            } else {
                                v58 := GenS6{F0: -2147483648}
                                _ = v58
                            }
                            if true {
                                var v59 float32 = v1
                                _ = v59
                            }
                            _ = v57
                        }
                    }
                case "foo bar":
                    for fi60 := 0; fi60 < 13; fi60++ {
                        fmt.Println(v4)
                        _ = fi60
                    }
                }
                _ = v56
            }
            if true {
                fmt.Println(fi53, v4)
                v61 := GenS7{F0: 114, F1: 42, F2: v28}
                var v62 uint16 = (v4 % ((v4) | 1))
                _ = v61
                _ = v62
            } else {
                fmt.Println(v52)
                v63 := []int{}
                _ = v63
            }
            var v64 int = fi53
            _ = fi53
            _ = v64
        }
        for fi65 := 0; fi65 < 10; fi65++ {
            var v66 uint64 = v2
            for fi67 := 0; fi67 < 18; fi67++ {
                switch fi65 {
                case 1:
                    var v68 float64 = 3.14159
                    if true {
                        if false {
                            var v69 int8 = -1
                            fmt.Println(v66, fi67)
                            var v70 float32 = v1
                            switch "hello" {
                            case "":
                                var v71 float32 = (1e10 + (v1 - 1e10))
                                _ = v71
                            case "\n\t":
                                v72 := GenS2{F0: "a"}
                                _ = v72
                            case "foo bar":
                                fmt.Println(v1)
                            default:
                                var v73 int32 = genFn1(uint(v4), 65535)
                                _ = v73
                            }
                            _ = v69
                            _ = v70
                        }
                    }
                    _ = v68
                }
                _ = fi67
            }
            v74 := v5[(^(v28))]
            _ = fi65
            _ = v66
            _ = v74
        }
    } else {
        if true {
            for fi75 := 0; fi75 < 11; fi75++ {
                for fi76 := 0; fi76 < 5; fi76++ {
                    v77 := GenS8{F0: -1}
                    if true {
                        fmt.Println(v1, v1, v52)
                        v2 -= 18446744073709551615
                        var v78 int32 = -2147483648
                        if (0 >= -733) {
                            if ((v78 / (((v78 / ((-301) | 1))) | 1)) < v78) {
                                fmt.Println(v52, v52, v78)
                                switch v28 {
                                case 32768:
                                    v79 := []int{314}
                                    _ = v79
                                case 21:
                                    v80 := GenS8{F0: genFn1(120, 2147483648)}
                                    _ = v80
                                case 949:
                                    v28 = int(v2)
                                default:
                                    if (!false) {
                                        if false {
                                            if ("\n\t" < "\n\t") {
                                                var v81 uint32 = 282
                                                var v82 uint32 = (v81 * (v81 << ((v52) & 63)))
                                                _ = v81
                                                _ = v82
                                            }
                                            v83 := []int{}
                                            switch fi75 {
                                            case 814:
                                                fmt.Println(v4)
                                            case 891:
                                                v2 += (v2 << ((4294967295) & 63))
                                            default:
                                                v4 -= 255
                                            }
                                            _ = v83
                                        }
                                        v84 := GenS4{F0: v2}
                                        if (true || false) {
                                            var v85 bool = false
                                            _ = v85
                                        }
                                        _ = v84
                                    } else {
                                        v86 := GenS1{F0: (v2 / ((v2) | 1))}
                                        v87 := GenS4{F0: 4294967295}
                                        _ = v86
                                        _ = v87
                                    }
                                }
                                v78 -= -1
                                var v88 int16 = -190
                                _ = v88
                            }
                            v89 := GenS9{F0: (!(!true)), F1: v28}
                            if false {
                                v90 := GenS8{F0: 683}
                                _ = v90
                            } else {
                                v91 := GenS10{F0: 255, F1: 42}
                                switch (fi75 & fi75) {
                                case 9223372036854775807:
                                    fmt.Println(fi76, v52)
                                    if false {
                                        switch v28 {
                                        case -908:
                                            switch (("a" + "hello") + ("foo bar" + "hello")) {
                                            case "":
                                                v2 -= v2
                                            }
                                        case 1:
                                            v1 += (v1 + float32(1.5))
                                        case 32767:
                                            var v92 int32 = 42
                                            _ = v92
                                        }
                                    }
                                }
                                _ = v91
                            }
                            _ = v89
                        }
                        _ = v78
                    } else {
                        fmt.Println(fi75)
                        v93 := GenS11{F0: (v4 % ((734) | 1))}
                        v94 := GenS12{F0: 2, F1: 86, F2: 958}
                        fmt.Println(fi76, v2)
                        v95 := []int{}
                        _ = v93
                        _ = v94
                        _ = v95
                    }
                    fmt.Println(v1)
                    _ = fi76
                    _ = v77
                }
                var v96 uint16 = 255
                _ = fi75
                _ = v96
            }
            if (-186 >= 9223372036854775807) {
                for fi97 := 0; fi97 < 9; fi97++ {
                    switch "" {
                    case "\n\t":
                        v98 := []int{(v28 * (v28 ^ (fi97 & (v28 / (((fi97 | v28)) | 1)))))}
                        _ = v98
                    case "hello":
                        if true {
                            for fi99 := 0; fi99 < 13; fi99++ {
                                if false {
                                    if ((false || (-2.25 <= 1.5)) && false) {
                                        var v100 int8 = 100
                                        _ = v100
                                    }
                                    var v101 int = 0
                                    _ = v101
                                } else {
                                    var v102 float64 = 1.5
                                    _ = v102
                                }
                                var v103 float32 = ((v1 / (v1 * 1.5)) * v1)
                                _ = fi99
                                _ = v103
                            }
                            v104 := []int{(v28 ^ v28), v28}
                            for fi105 := 0; fi105 < 3; fi105++ {
                                v2 += 1
                                _ = fi105
                            }
                            _ = v104
                        }
                    default:
                        if false {
                            _ = len(v3)
                            v28 = (v28 * (v28 << (((v52 / ((1) | 1))) & 63)))
                        }
                    }
                    if (!true) {
                        var v106 float32 = 3.14159
                        fmt.Println(v106)
                        _ = v106
                    }
                    if (2 >= v52) {
                        var v107 uint32 = 2147483648
                        _ = v107
                    } else {
                        var v108 uint32 = 1
                        var v109 int = (v28 << (((v52 | 65535)) & 63))
                        var v110 uint32 = (v108 ^ (v108 / ((v108) | 1)))
                        _ = v108
                        _ = v109
                        _ = v110
                    }
                    _ = fi97
                }
            }
            v111 := []int{}
            _ = v111
        } else {
            var v112 uint = v52
            var v113 int = v28
            _ = v112
            _ = v113
        }
        var v114 uint = 0
        _ = v114
    }
    var v115 int = 32768
    if false {
        fmt.Println(v28, v4, v4)
    } else {
        switch (v115 + 128) {
        case 32767:
            var v116 uint32 = 42
            _ = v116
        case 9223372036854775807:
            if false {
                v4 -= v4
                v117 := GenS13{F0: v115}
                _ = v117
            }
        case 32768:
            for fi118 := 0; fi118 < 5; fi118++ {
                var v119 string = "foo bar"
                v120 := GenS14{F0: 2147483647, F1: 4294967295}
                _ = fi118
                _ = v119
                _ = v120
            }
        }
        v121 := GenS1{F0: (v2 * (^(v2)))}
        _ = v121
    }
    var v122 float32 = float32(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v115
    _ = v122
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
