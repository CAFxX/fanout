package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32(-2.25)
    v3 := make(map[string]uint32)
    var v4 uint8 = 1
    if (42 <= 2) {
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(fi5, fi5)
                var v7 uint8 = v4
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{0, fi8}
            _ = fi8
            _ = v9
        }
        v1[2147483648] = -9223372036854775808
    } else {
        switch 480 {
        case 1:
            if (0 == -2147483648) {
                var v10 float64 = 0.1
                v10 -= 1.5
                _ = v10
            } else {
                fmt.Println(v4, v2)
                v2 = float32(3.14159)
                v11 := []int{0, 9223372036854775807}
                _ = v11
            }
        case 32768:
            if (!(0 < 793)) {
                fmt.Println(v4, v2)
                var v12 uint32 = 0
                _ = v12
            }
        case 0:
            var v13 int8 = 42
            _ = v13
        default:
            if true {
                delete(v3, "hello")
                var v14 int = 750
                _ = v14
            } else {
                if (!true) {
                    var v15 bool = true
                    v2 = (((float32((3.14159 - float64((float32(0.0) - (float32((1.5 * -2.25)) * -2.25))))) / ((3.14159 * (0.0 / (3.14159 * float32(1e10)))) * (v2 * v2))) * float32(0.1)) - (float32(1e10) + (v2 - v2)))
                    _ = v15
                } else {
                    v4 += v4
                    if (42 > -11) {
                        fmt.Println(v4, v2)
                        v2 += 0.1
                    } else {
                        var v16 int32 = -725
                        _ = len(v3)
                        var v17 bool = (v2 != ((float32(3.14159) + float32((1e10 / float64(v2)))) / 1.5))
                        _ = v16
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 18; fi18++ {
                    v19 := GenS1{F0: 128, F1: 802}
                    var v20 float32 = v2
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if (!(false && true)) {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        if true {
                            if (!true) {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    switch fi21 {
                                    case 633:
                                        v2 = v2
                                    }
                                    if (("hello" < ("a" + "\n\t")) || true) {
                                        v23 := []int{}
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: (fi21 & (fi21 & 852)), F1: 419}
                                        _ = v24
                                    }
                                    if (v2 > v2) {
                                        if false {
                                            var v25 float32 = v2
                                            v26 := []int{1, (fi21 | fi22), (fi22 - fi21), 32768}
                                            fmt.Println(v2)
                                            _ = v25
                                            _ = v26
                                        }
                                        v1[2147483647] = -1
                                    } else {
                                        fmt.Println(v4, v2, fi21)
                                        v27 := []int{fi22, 0}
                                        _ = v27
                                    }
                                    _ = fi22
                                }
                                if false {
                                    delete(v3, "foo bar")
                                    v4 = (v4 + v4)
                                    v2 = v2
                                }
                                v28 := []int{(-(fi21)), (-(fi21)), fi21}
                                _ = v28
                            } else {
                                v29 := GenS1{F0: -9223372036854775808, F1: 65535}
                                _ = v29
                            }
                            fmt.Println(v4, v4)
                        } else {
                            v30 := GenS1{F0: fi21, F1: 559}
                            fmt.Println(v2, v4)
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 8; fi31++ {
                            v2 = v2
                            switch ((("a" + "\n\t") + "") + "") {
                            case "\n\t":
                                var v32 uint8 = (^(v4))
                                _ = v32
                            case "":
                                switch fi21 {
                                case 963:
                                    v33 := []int{9223372036854775807, fi31, fi21, (fi31 - fi31)}
                                    _ = v33
                                case -225:
                                    v2 += 0.0
                                }
                            case "hello":
                                if (false || false) {
                                    var v34 int64 = 1
                                    var v35 int8 = 127
                                    _ = v34
                                    _ = v35
                                } else {
                                    v4 = v4
                                    var v36 int16 = -493
                                    _ = v36
                                }
                            default:
                                v37 := GenS1{F0: fi21, F1: 0}
                                _ = v37
                            }
                            _ = fi31
                        }
                        _ = len(v1)
                        _ = fi21
                    }
                    fmt.Println(v2)
                    if true {
                        for fi38 := 0; fi38 < 13; fi38++ {
                            if (((0 <= v4) && true) || true) {
                                v39 := GenS1{F0: 1, F1: 0}
                                delete(v1, 2)
                                v40 := GenS2{F0: 3, F1: (v4 << ((18446744073709551615) & 63))}
                                _ = v39
                                _ = v40
                            }
                            _ = fi38
                        }
                        var v41 float32 = (float32(1.5) / v2)
                        var v42 float64 = 1e10
                        _ = v41
                        _ = v42
                    }
                } else {
                    if false {
                        fmt.Println(v4)
                        v43 := []int{32767, -9223372036854775808, 4294967295, 32767}
                        _ = v43
                    } else {
                        v44 := []int{291, -257, -628}
                        _ = v44
                    }
                }
            }
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            v2 -= float32(0.1)
            for fi46 := 0; fi46 < 14; fi46++ {
                v47 := GenS1{F0: fi45, F1: 1}
                v2 = 0.1
                v48 := []int{}
                _ = fi46
                _ = v47
                _ = v48
            }
            v49 := GenS1{F0: 32767, F1: 42}
            _ = fi45
            _ = v49
        }
    }
    v4 = v4
    v4 += v4
    fmt.Println(v4)
    var v50 uint = 955
    v51 := GenS2{F0: v4, F1: 255}
    for fi52 := 0; fi52 < 13; fi52++ {
        if ((-2.25 * (1e10 + -2.25)) == 0.1) {
            v50 -= v50
            for fi53 := 0; fi53 < 14; fi53++ {
                switch -9223372036854775808 {
                case -2147483648:
                    var v54 uint8 = 230
                    _ = v54
                case -873:
                    v2 -= -2.25
                    if true {
                        var v55 float64 = (1.5 * -2.25)
                        switch fi52 {
                        case 2147483647:
                            var v56 bool = false
                            _ = v56
                        case -9223372036854775808:
                            if true {
                                switch ("hello" + ("hello" + (("hello" + "foo bar") + "a"))) {
                                case "a":
                                    var v57 uint64 = 42
                                    _ = v57
                                case "\n\t":
                                    v58 := GenS2{F0: v4, F1: 127}
                                    _ = v58
                                }
                                var v59 uint8 = v4
                                v60 := GenS1{F0: 42, F1: uint16(v59)}
                                _ = v59
                                _ = v60
                            }
                        case -1:
                            v4 = 255
                        }
                        _ = v55
                    }
                    var v61 string = "hello"
                    var v62 int32 = 2147483647
                    _ = v61
                    _ = v62
                }
                v50 = v50
                var v63 int = fi53
                if (!false) {
                    var v64 uint16 = uint16(v50)
                    _ = v64
                } else {
                    v65 := []int{fi52, -465, 9223372036854775807}
                    if ("\n\t" <= "a") {
                        if true {
                            v63 -= (fi53 << ((18446744073709551615) & 63))
                        }
                        var v66 uint64 = 2
                        fmt.Println(v4, v4, v50)
                        _ = v66
                    } else {
                        v4 = (v4 + 2)
                        var v67 int16 = 1
                        _ = v67
                    }
                    v63 += 127
                    _ = v65
                }
                _ = fi53
                _ = v63
            }
        }
        var v68 int = fi52
        for fi69 := 0; fi69 < 9; fi69++ {
            if ((2 != 2) && (4294967295 < 65535)) {
                switch v68 {
                case -9223372036854775808:
                    if ((true && (!false)) && (((2147483647 != 2) && false) && (((v68 / (((^(fi69))) | 1)) > 0) && false))) {
                        v3[(("foo bar" + (("\n\t" + "foo bar") + "\n\t")) + "\n\t")] = 602
                        v3["a"] = 2147483647
                    } else {
                        fmt.Println(v50, v4, v2)
                        switch "foo bar" {
                        case "hello":
                            v70 := GenS2{F0: uint8(v50), F1: (^(v4))}
                            _ = v70
                        case "foo bar":
                            if true {
                                switch v68 {
                                case 440:
                                    var v71 uint = v50
                                    _ = v71
                                default:
                                    if (2147483647 == 2) {
                                        var v72 int64 = -9223372036854775808
                                        fmt.Println(v4)
                                        v73 := GenS1{F0: v68, F1: 32767}
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            if false {
                                                v74 := v3[(("a" + ("a" + "a")) + "foo bar")]
                                                var v75 bool = true
                                                _ = v74
                                                _ = v75
                                            }
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v76 := GenS1{F0: 338, F1: 0}
                                        v77 := GenS2{F0: v4, F1: v4}
                                        _ = v76
                                        _ = v77
                                    }
                                }
                            }
                        case "":
                            if false {
                                var v78 int16 = 1
                                v79 := v3[("hello" + "foo bar")]
                                _ = v78
                                _ = v79
                            } else {
                                v68 = 561
                                switch "foo bar" {
                                case "\n\t":
                                    var v80 int64 = 685
                                    _ = v80
                                case "foo bar":
                                    if true {
                                        if (1 > -9223372036854775808) {
                                            v81 := GenS1{F0: (v68 >> (((v50 & 263)) & 63)), F1: 354}
                                            var v82 int8 = int8(v50)
                                            fmt.Println(v4, fi52)
                                            _ = v81
                                            _ = v82
                                        } else {
                                            fmt.Println(v68)
                                            if (!(!true)) {
                                                v83 := GenS2{F0: 255, F1: 217}
                                                _ = v83
                                            }
                                            if false {
                                                v2 += v2
                                            }
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            var v84 float32 = (v2 * (v2 + float32(0.0)))
                                            fmt.Println(v50, v68, v4)
                                            _ = v84
                                        case "":
                                            fmt.Println(v4, v2)
                                        case "\n\t":
                                            v85 := GenS1{F0: (fi69 | (^(v68))), F1: 32767}
                                            _ = v85
                                        default:
                                            v86 := GenS3{F0: 0, F1: 18446744073709551615, F2: -1}
                                            _ = v86
                                        }
                                    }
                                case "a":
                                    v87 := GenS4{F0: 0, F1: 1000, F2: true}
                                    _ = v87
                                }
                                v1[0] = (v68 << ((1) & 63))
                                var v88 string = ((("foo bar" + "\n\t") + "") + "")
                                v89 := GenS4{F0: 0, F1: 32767, F2: (!true)}
                                v90 := []int{195, -9223372036854775808, (fi69 * (-(v68))), 9223372036854775807}
                                var v91 uint16 = 2
                                _ = v88
                                _ = v89
                                _ = v90
                                _ = v91
                            }
                        default:
                            fmt.Println(v68)
                        }
                        v92 := GenS4{F0: 42, F1: -32768, F2: ((uint8(v50) != v4) && false)}
                        _ = v92
                    }
                case 336:
                    var v93 float32 = float32(float64(v2))
                    _ = v93
                default:
                    v94 := GenS1{F0: (fi69 & (v68 ^ -458)), F1: 432}
                    _ = v94
                }
            } else {
                switch "a" {
                case "hello":
                    var v95 uint8 = v4
                    _ = v95
                case "":
                    v96 := GenS1{F0: (v68 * v68), F1: 1}
                    _ = v96
                }
                v97 := GenS4{F0: 1, F1: int16(v4), F2: false}
                var v98 float64 = 0.1
                v99 := GenS5{F0: (v50 + (v50 ^ v50))}
                v100 := []int{628, 464, v68}
                _ = v100
                _ = v97
                _ = v98
                _ = v99
            }
            var v101 int64 = 9223372036854775807
            _ = fi69
            _ = v101
        }
        _ = fi52
        _ = v68
    }
    v102 := []int{4294967295, 32767, 1}
    var v103 int16 = 1
    if len(v102) >= 2 {
        v102 = v102[1:len(v102)]
    }
    v4 = (v4 + v4)
    switch (("hello" + ("foo bar" + "hello")) + "hello") {
    case "\n\t":
        var v104 uint = (v50 - (v50 - 18446744073709551615))
        _ = v104
    case "foo bar":
        if (!((460 < 986) && true)) {
            v4 -= v4
            var v105 int = 128
            _ = v105
        }
    }
    for fi106 := 0; fi106 < 13; fi106++ {
        v2 = ((v2 / (1.5 * ((float32(3.14159) * v2) / v2))) / ((-2.25 / ((-2.25 + v2) - (0.0 + v2))) / 0.1))
        var v107 int64 = 181
        for fi108 := 0; fi108 < 20; fi108++ {
            v109 := GenS6{F0: fi106}
            v4 = v4
            v3[("foo bar" + ("\n\t" + "a"))] = 42
            fmt.Println(fi108, v2, v107)
            _ = fi108
            _ = v109
        }
        _ = fi106
        _ = v107
    }
    v2 -= v2
    fmt.Println(v4)
    for fi110 := 0; fi110 < 11; fi110++ {
        fmt.Println(v4, v2, v4)
        _ = fi110
    }
    v103 += v103
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
