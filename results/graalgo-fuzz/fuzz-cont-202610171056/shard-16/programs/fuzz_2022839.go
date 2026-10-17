package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 int32 }
type GenS5 struct { F0 string; F1 string; F2 int64 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int64; F1 int8; F2 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 bool; F1 int8 }
func genFn0() int {
    v5 := []int{9223372036854775807, 989}
    var v6 float64 = (float64(((-2.25 / 1e10) - float32(3.14159))) * (((0.1 / 3.14159) + ((float64(1e10) * 0.0) + 3.14159)) / 3.14159))
    v6 += -2.25
    v6 = v6
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int64) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 4294967295}
        _ = fi7
        _ = v8
    }
    switch 9223372036854775807 {
    case 186:
        if (-652 != 0) {
            var v9 int8 = -1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int8 = -127
                    if true {
                        v12 = (v12 / ((v12) | 1))
                    } else {
                        v13 := []int{(^(fi11)), (fi10 % ((fi10) | 1)), fi10}
                        v12 -= -128
                        _ = v13
                    }
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
    case -366:
        var v14 float32 = -2.25
        _ = v14
    case 127:
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := []int{-496}
            _ = fi15
            _ = v16
        }
    }
    var v17 bool = false
    _ = p0
    _ = v17
    return 255
}

func main() {
    var v1 uint64 = 876
    v2 := []uint32{13, 1}
    var v3 int = 4294967295
    var v4 string = "foo bar"
    var v18 uint8 = 1
    v2 = append(v2, 2147483648)
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = 9223372036854775807
        for fi21 := 0; fi21 < 12; fi21++ {
            if (!true) {
                switch v4 {
                case "":
                    if (32767 >= 897) {
                        fmt.Println(v20)
                        if true {
                            switch (fi21 - fi21) {
                            case 0:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -380:
                                var v22 int64 = v20
                                _ = v22
                            case 2147483647:
                                v23 := GenS2{F0: -641}
                                _ = v23
                            default:
                                v24 := GenS1{F0: 4294967295}
                                _ = v24
                            }
                            if true {
                                var v25 bool = false
                                v3 = (v3 | (fi21 >> ((949) & 63)))
                                _ = v25
                            }
                            var v26 int64 = (^(v20))
                            _ = v26
                        }
                        if (!(255 >= 4294967295)) {
                            fmt.Println(v18, fi21)
                        } else {
                            var v27 int64 = (^(v20))
                            fmt.Println(v1, fi19)
                            _ = v27
                        }
                    } else {
                        var v28 bool = true
                        _ = v28
                    }
                case "foo bar":
                    _ = len(v2)
                case "a":
                    v2 = append(v2, 0)
                }
                v29 := GenS2{F0: (fi19 | fi19)}
                if (true || false) {
                    if (true && true) {
                        switch (fi19 & 0) {
                        case 0:
                            var v30 int8 = 87
                            _ = v30
                        case 4294967295:
                            switch -170 {
                            case -362:
                                fmt.Println(fi21, v3, fi21)
                            case -209:
                                v1 = v1
                            default:
                                var v31 uint64 = 1
                                _ = v31
                            }
                        default:
                            _ = len(v2)
                        }
                        if ((v4 + "a") != v4) {
                            switch v4 {
                            case "foo bar":
                                fmt.Println(v18, fi21)
                            case "hello":
                                v18 -= v18
                            }
                            fmt.Println(v20, v4, v1)
                            if true {
                                var v32 bool = true
                                _ = v32
                            } else {
                                fmt.Println(v1, fi21, v3)
                                v2 = append(v2, 739)
                                fmt.Println(v18)
                            }
                        } else {
                            var v33 uint64 = (v1 / (((v1 << ((255) & 63))) | 1))
                            if (genFn0() < v3) {
                                v1 = v33
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                            } else {
                                switch v4 {
                                case "hello":
                                    var v34 float32 = float32((-2.25 - 1.5))
                                    _ = v34
                                default:
                                    switch ("hello" + v4) {
                                    case "\n\t":
                                        var v35 bool = (("\n\t" + v4) <= v4)
                                        v35 = false
                                        if (((0 == -2147483648) || false) || false) {
                                            var v36 float32 = -2.25
                                            var v37 float64 = 0.1
                                            _ = v36
                                            _ = v37
                                        } else {
                                            var v38 float32 = float32(-2.25)
                                            _ = v38
                                        }
                                        _ = v35
                                    case "foo bar":
                                        if true {
                                            if (("a" <= ((((v4 + "foo bar") + ("hello" + (v4 + v4))) + v4) + "hello")) && ((v1 + (v33 >> ((609) & 63))) != v33)) {
                                                v1 += uint64(v18)
                                                if len(v2) > 0 {
                                                    v2[0] = 686
                                                }
                                                v3 = (^(fi21))
                                                v20 = (v20 - v20)
                                            } else {
                                                var v39 string = v4
                                                var v40 int8 = 100
                                                switch 929 {
                                                case 656:
                                                    if (!true) {
                                                        var v41 int32 = 2147483647
                                                        var v42 uint64 = 18446744073709551615
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                case -54:
                                                    var v43 int8 = v40
                                                    _ = v43
                                                case 128:
                                                    if ((^(v18)) > v18) {
                                                        var v44 bool = true
                                                        v45 := GenS1{F0: 4294967295}
                                                        v44 = (!((false || false) || (v44 || false)))
                                                        var v46 string = (v4 + "foo bar")
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                default:
                                                    v1 = 0
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            v20 = 525
                                            v4 = "hello"
                                        }
                                    case "a":
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v18)
                                        case "":
                                            if true {
                                                switch 2147483647 {
                                                case -2147483648:
                                                    var v47 uint64 = v1
                                                    _ = v47
                                                case -266:
                                                    switch ((v4 + "") + (((v4 + "a") + (v4 + (("foo bar" + v4) + ("a" + v4)))) + v4)) {
                                                    case "\n\t":
                                                        v1 = uint64(v18)
                                                    case "":
                                                        v20 = (-(v20))
                                                    }
                                                }
                                                var v48 float64 = 1.5
                                                var v49 int32 = 820
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch (("" + "foo bar") + "\n\t") {
                                                case "hello":
                                                    var v50 string = v4
                                                    _ = v50
                                                case "\n\t":
                                                    fmt.Println(v33, fi21, v20)
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case "foo bar":
                                            var v51 string = v4
                                            _ = v51
                                        default:
                                            var v52 int16 = 1
                                            _ = v52
                                        }
                                    }
                                    if true {
                                        var v53 uint8 = 107
                                        if true {
                                            switch (v3 - 32767) {
                                            case 555:
                                                var v54 uint = 42
                                                _ = v54
                                            }
                                            fmt.Println(v3)
                                            var v55 uint8 = (v53 | v18)
                                            _ = v55
                                        }
                                        v56 := GenS1{F0: 255}
                                        _ = v53
                                        _ = v56
                                    } else {
                                        var v57 uint16 = 65535
                                        v58 := GenS3{F0: (v57 % ((uint16(v1)) | 1)), F1: 0}
                                        if (v18 < (v18 + v18)) {
                                            v59 := GenS3{F0: v57, F1: 843}
                                            v33 = 4294967295
                                            _ = v59
                                        }
                                        _ = v57
                                        _ = v58
                                    }
                                    fmt.Println(v33, v20, v4)
                                }
                                var v60 uint = uint(v33)
                                v61 := GenS4{F0: (v60 % ((v60) | 1)), F1: (^(v18)), F2: 2147483647}
                                _ = v60
                                _ = v61
                            }
                            _ = v33
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v18 = 128
                        v20 = v20
                        v20 += (v20 << ((255) & 63))
                    }
                    v2 = append(v2, 1)
                    var v62 int8 = 42
                    _ = v62
                } else {
                    v2 = append(v2, 42)
                }
                _ = v29
            } else {
                var v63 string = "hello"
                if (false || (2 <= 2)) {
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    var v64 bool = ((v18 | v18) != 128)
                    if (!false) {
                        switch (fi19 * genFn0()) {
                        case 511:
                            fmt.Println(v63, v64)
                        default:
                            switch v4 {
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = 707
                                }
                            }
                        }
                    } else {
                        var v65 string = v63
                        v2 = append(v2, 2)
                        v66 := GenS5{F0: v4, F1: (("a" + v65) + (("hello" + v65) + v4)), F2: 1}
                        _ = v65
                        _ = v66
                    }
                    _ = v64
                } else {
                    var v67 int16 = 32767
                    switch ("" + (v4 + ((v4 + "\n\t") + "\n\t"))) {
                    case "":
                        v2 = append(v2, 647)
                    case "hello":
                        if ((279 == 956) || (0.1 != 1e10)) {
                            if true {
                                v3 -= (fi21 << ((0) & 63))
                            } else {
                                var v68 uint64 = v1
                                v69 := GenS6{F0: 338, F1: 1}
                                _ = v68
                                _ = v69
                            }
                            var v70 float32 = 0.0
                            _ = v70
                        }
                        v3 += 127
                        var v71 int8 = -128
                        _ = v71
                    }
                    v67 += v67
                    if true {
                        var v72 string = v4
                        var v73 bool = true
                        var v74 int16 = (^(v67))
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    fmt.Println(v1)
                    _ = v67
                }
                v1 = (v1 % ((2) | 1))
                _ = v63
            }
            fmt.Println(v20, v1)
            v20 = v20
            _ = fi21
        }
        v75 := GenS7{F0: -1, F1: ("hello" + v4)}
        _ = fi19
        _ = v20
        _ = v75
    }
    switch "foo bar" {
    case "":
        var v76 uint = 643
        _ = v76
    case "foo bar":
        if (!false) {
            v1 += v1
            _ = len(v2)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            for fi77 := 0; fi77 < 10; fi77++ {
                if (v4 <= (v4 + v4)) {
                    switch -398 {
                    case 497:
                        switch (^(fi77)) {
                        case 1:
                            var v78 int64 = -1
                            _ = v78
                        case -7:
                            switch ("\n\t" + v4) {
                            case "hello":
                                if true {
                                    var v79 uint64 = (v1 + v1)
                                    v18 = (^(v18))
                                    var v80 int = genFn0()
                                    _ = v79
                                    _ = v80
                                }
                            default:
                                switch (v3 ^ 1) {
                                case -1:
                                    fmt.Println(v18, v3)
                                case 128:
                                    for fi81 := 0; fi81 < 8; fi81++ {
                                        var v82 int8 = 62
                                        if len(v2) > 0 {
                                            v2[0] = 790
                                        }
                                        v1 -= v1
                                        _ = fi81
                                        _ = v82
                                    }
                                case -974:
                                    var v83 int16 = 42
                                    _ = v83
                                default:
                                    var v84 uint = 522
                                    _ = v84
                                }
                            }
                        default:
                            for fi85 := 0; fi85 < 8; fi85++ {
                                var v86 int64 = -1
                                v86 = v86
                                switch v3 {
                                case 101:
                                    v3 += (^(v3))
                                }
                                _ = fi85
                                _ = v86
                            }
                        }
                    case 0:
                        v18 += v18
                    default:
                        v4 = "foo bar"
                    }
                    v87 := GenS8{F0: 1, F1: -39, F2: (v18 / ((v18) | 1))}
                    var v88 int8 = 1
                    _ = v87
                    _ = v88
                } else {
                    v89 := GenS9{F0: -741}
                    if ("" > "\n\t") {
                        if false {
                            v3 = v3
                            v90 := GenS10{F0: 14, F1: v18, F2: 2147483647}
                            var v91 float64 = 3.14159
                            _ = v90
                            _ = v91
                        } else {
                            v92 := GenS11{F0: uint8(v1), F1: true}
                            _ = v92
                        }
                        var v93 int8 = 100
                        _ = v93
                    }
                    v3 += genFn0()
                    _ = v89
                }
                v94 := GenS12{F0: (!(2 != 240)), F1: 0}
                _ = fi77
                _ = v94
            }
            if (42 <= 0) {
                var v95 bool = (12 != 42)
                if (v95 && false) {
                    var v96 int = (v3 * genFn0())
                    var v97 bool = v95
                    for fi98 := 0; fi98 < 11; fi98++ {
                        v1 = 2
                        var v99 string = "hello"
                        _ = fi98
                        _ = v99
                    }
                    _ = v96
                    _ = v97
                } else {
                    fmt.Println(v95, v95)
                    fmt.Println(v1, v18, v3)
                }
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                _ = v95
            }
            var v100 float32 = (1e10 * 1e10)
            _ = v100
        }
    case "\n\t":
        _ = len(v2)
    default:
        var v101 float32 = 1e10
        _ = v101
    }
    fmt.Println(v1, v18, v1)
    var v102 int8 = 127
    var v103 uint8 = 2
    var v104 int = v3
    for fi105 := 0; fi105 < 9; fi105++ {
        fmt.Println(v1)
        _ = fi105
    }
    for fi106 := 0; fi106 < 5; fi106++ {
        if false {
            fmt.Println(v104, v104)
        }
        v107 := GenS7{F0: 0, F1: ""}
        _ = fi106
        _ = v107
    }
    var v108 uint16 = 65535
    for fi109 := 0; fi109 < 9; fi109++ {
        v1 = 2
        _ = fi109
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v110 := GenS4{F0: 452, F1: (v103 >> ((2) & 63)), F2: 0}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v108:", v108)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v108
    _ = v110
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
