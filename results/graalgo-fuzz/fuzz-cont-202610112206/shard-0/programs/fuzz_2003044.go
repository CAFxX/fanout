package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16) int16 {
    v5 := GenS1{F0: p0, F1: (^(p1)), F2: 18446744073709551615}
    if ((-463 > 1) || (!true)) {
        v6 := []int{}
        var v7 int = 4
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: (("foo bar" + ("a" + (p0 + ""))) + ((p0 + "foo bar") + (p0 + p0))), F1: p1, F2: 18446744073709551615}
        if (!(-1 > 0)) {
            var v9 int16 = 0
            v10 := GenS1{F0: "", F1: (p1 / ((997) | 1)), F2: 561}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: -7, F1: 1, F2: 237}
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                p1 = p1
                _ = fi13
            }
            p1 = p1
            _ = fi12
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 bool = true
    v1 = 1
    var v14 uint32 = 794
    var v15 float32 = 1.5
    var v16 string = "a"
    if true {
        v14 = 4294967295
        switch 10 {
        case -2147483648:
            var v17 int8 = -23
            _ = v17
        case -1:
            v18 := GenS1{F0: "a", F1: 32767, F2: 584}
            _ = v18
        case -593:
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 string = (v16 + "hello")
                v21 := []int{(fi19 & (fi19 % ((-9223372036854775808) | 1)))}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        if v4 {
            v16 = v16
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int = fi22
                v14 += 2147483648
                if v4 {
                    var v24 uint = 0
                    var v25 uint64 = 18446744073709551615
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v27 := []int{(v23 % ((-1) | 1)), (v23 % ((fi26) | 1)), int(v24)}
                        v27 = append(v27, v23)
                        delete(v2, (v16 + ("hello" + ("foo bar" + ((v16 + (v16 + v16)) + (v16 + v16))))))
                        v28 := GenS2{F0: 127, F1: (v1 >> ((v24) & 63)), F2: -2147483648}
                        var v29 int64 = 9223372036854775807
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v24
                    _ = v25
                } else {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 int8 = 48
                        v16 = v16
                        _ = fi30
                        _ = v31
                    }
                    if true {
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 int16 = -533
                            _ = fi32
                            _ = v33
                        }
                    }
                    switch -2147483648 {
                    case 128:
                        var v34 int = fi22
                        _ = v34
                    case -2147483648:
                        switch (^(v23)) {
                        case -613:
                            v16 = (v16 + ((("" + (((v16 + "a") + ("" + v16)) + v16)) + v16) + "a"))
                        case 32768:
                            if (42 < 42) {
                                switch (v16 + v16) {
                                case "a":
                                    var v35 int = v23
                                    _ = v35
                                case "hello":
                                    if ((v23 & (v23 * -2147483648)) > fi22) {
                                        var v36 string = v16
                                        var v37 bool = (false && true)
                                        switch (v36 + ((v36 + v16) + "hello")) {
                                        case "hello":
                                            fmt.Println(v4)
                                        case "":
                                            var v38 int = (-(v23))
                                            _ = v38
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    v2[v16] = 2147483648
                                    _ = fi39
                                }
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    v41 := []int{(fi22 >> ((4294967295) & 63)), (v23 / ((fi22) | 1))}
                                    _ = fi40
                                    _ = v41
                                }
                            }
                        default:
                            v15 -= -2.25
                        }
                    case 475:
                        if (!true) {
                            var v42 int16 = -444
                            switch "foo bar" {
                            case "hello":
                                v43 := GenS2{F0: -46, F1: v1, F2: (fi22 | (-(v23)))}
                                _ = v43
                            default:
                                var v44 uint64 = 42
                                _ = v44
                            }
                            _ = v42
                        } else {
                            v45 := GenS2{F0: 42, F1: (^(v1)), F2: (fi22 & fi22)}
                            var v46 int8 = 100
                            _ = v45
                            _ = v46
                        }
                    }
                }
                _ = fi22
                _ = v23
            }
            if true {
                var v47 int = 32768
                _ = v47
            } else {
                var v48 float32 = v15
                switch -9223372036854775808 {
                case 800:
                    delete(v3, v16)
                case 25:
                    var v49 uint64 = 607
                    _ = v49
                case -1:
                    v16 = (v16 + (v16 + (v16 + "foo bar")))
                default:
                    v50 := GenS3{F0: (^(v14)), F1: 4294967295, F2: (0 >= 0)}
                    _ = v50
                }
                for fi51 := 0; fi51 < 12; fi51++ {
                    for fi52 := 0; fi52 < 17; fi52++ {
                        v1 += v1
                        var v53 float32 = (0.1 * 3.14159)
                        delete(v2, v16)
                        _ = fi52
                        _ = v53
                    }
                    _ = fi51
                }
                _ = v48
            }
        } else {
            for fi54 := 0; fi54 < 13; fi54++ {
                var v55 int = (-(fi54))
                for fi56 := 0; fi56 < 10; fi56++ {
                    var v57 uint16 = 65535
                    var v58 uint32 = (v14 & v14)
                    var v59 uint64 = 707
                    _ = fi56
                    _ = v57
                    _ = v58
                    _ = v59
                }
                v16 = v16
                _ = fi54
                _ = v55
            }
            if ((float64(1e10) * 1.5) < ((1.5 + (1.5 * 0.1)) / float64(v15))) {
                v60 := []int{-976, -461, -857}
                switch v16 {
                case "\n\t":
                    delete(v3, ("hello" + (v16 + v16)))
                case "foo bar":
                    var v61 int = 127
                    _ = v61
                case "a":
                    for fi62 := 0; fi62 < 2; fi62++ {
                        _ = len(v2)
                        switch (fi62 ^ (fi62 * 4294967295)) {
                        case 128:
                            var v63 int64 = -9223372036854775808
                            _ = v63
                        case -670:
                            v60 = append(v60, (fi62 * (^(fi62))))
                        case 0:
                            v4 = (42 <= 777)
                        }
                        _ = fi62
                    }
                default:
                    v14 = 0
                }
                var v64 int64 = -70
                _ = v60
                _ = v64
            } else {
                if v4 {
                    _ = len(v3)
                } else {
                    switch "" {
                    case "\n\t":
                        v4 = v4
                    case "foo bar":
                        fmt.Println(v15, v1, v15)
                    case "":
                        fmt.Println(v4, v16, v14)
                    default:
                        switch v16 {
                        case "foo bar":
                            var v65 int8 = 1
                            _ = v65
                        case "\n\t":
                            _ = len(v3)
                        case "":
                            v66 := []int{-62, -751, 737}
                            _ = v66
                        }
                    }
                    v67 := GenS4{F0: -2147483648}
                    _ = v67
                }
            }
        }
    } else {
        switch v16 {
        case "hello":
            fmt.Println(v1, v14)
        case "foo bar":
            var v68 float32 = v15
            _ = v68
        }
        var v69 float64 = 3.14159
        var v70 string = v16
        _ = v69
        _ = v70
    }
    for fi71 := 0; fi71 < 9; fi71++ {
        fmt.Println(v1, v4)
        if v4 {
            for fi72 := 0; fi72 < 1; fi72++ {
                v73 := GenS4{F0: 829}
                _ = fi72
                _ = v73
            }
        } else {
            switch v16 {
            case "":
                var v74 uint = 65535
                _ = v74
            case "a":
                delete(v3, (v16 + v16))
            default:
                var v75 int16 = 42
                _ = v75
            }
            v1 += (v1 - v1)
            v76 := GenS5{F0: (v16 + "foo bar")}
            _ = v76
        }
        _ = fi71
    }
    fmt.Println(v14, v16, v16)
    if ((v14 - (v14 << ((0) & 63))) < (v14 % (((v14 / ((v14) | 1))) | 1))) {
        for fi77 := 0; fi77 < 13; fi77++ {
            var v78 uint32 = (^(v14))
            v79 := []int{-977}
            _ = fi77
            _ = v78
            _ = v79
        }
    }
    if (!v4) {
        v80 := GenS6{F0: uint(v14)}
        switch -754 {
        case 839:
            var v81 string = v16
            _ = v81
        case 32768:
            for fi82 := 0; fi82 < 20; fi82++ {
                var v83 int64 = -1
                _ = fi82
                _ = v83
            }
        case -374:
            v1 = (v1 ^ (v1 >> ((0) & 63)))
        default:
            if (!false) {
                v14 = 1
                if v4 {
                    v14 = 2147483648
                    var v84 int8 = 127
                    _ = v84
                } else {
                    var v85 uint16 = 0
                    _ = v85
                }
                v86 := GenS2{F0: -52, F1: 42, F2: -1}
                _ = v86
            } else {
                if v4 {
                    v14 += 42
                }
                v16 = "hello"
                switch 9223372036854775807 {
                case 32768:
                    if (236 <= 255) {
                        v87 := GenS7{F0: 871}
                        if (v1 >= 42) {
                            v15 = v15
                            if v4 {
                                var v88 int16 = -1
                                var v89 int = 4294967295
                                _ = v88
                                _ = v89
                            }
                            if v4 {
                                var v90 uint16 = 324
                                var v91 int8 = 33
                                _ = v90
                                _ = v91
                            }
                        }
                        _ = v87
                    }
                case -667:
                    for fi92 := 0; fi92 < 20; fi92++ {
                        var v93 uint8 = 42
                        var v94 int8 = -79
                        _ = fi92
                        _ = v93
                        _ = v94
                    }
                }
            }
        }
        delete(v3, (v16 + v16))
        _ = v80
    }
    var v95 int8 = -64
    var v96 int = 2147483647
    v4 = v4
    for fi97 := 0; fi97 < 5; fi97++ {
        for fi98 := 0; fi98 < 11; fi98++ {
            var v99 int16 = -32768
            _ = fi98
            _ = v99
        }
        for fi100 := 0; fi100 < 18; fi100++ {
            v101 := []int{v96, -9223372036854775808, (-(fi100)), 1}
            _ = fi100
            _ = v101
        }
        _ = fi97
    }
    v15 -= (v15 + v15)
    if true {
        v102 := GenS4{F0: v96}
        _ = v102
    }
    var v103 int64 = 595
    if v4 {
        if v4 {
            v104 := v3[v16]
            _ = v104
        } else {
            if v4 {
                fmt.Println(v103, v14, v1)
            } else {
                var v105 uint16 = 244
                var v106 uint32 = v14
                _ = v105
                _ = v106
            }
            switch v96 {
            case -2147483648:
                v96 -= 9223372036854775807
            default:
                var v107 float32 = v15
                _ = v107
            }
        }
        var v108 uint = 0
        if (9223372036854775807 != (v103 % ((-9223372036854775808) | 1))) {
            for fi109 := 0; fi109 < 5; fi109++ {
                var v110 int16 = 1000
                fmt.Println(v1)
                switch (((v16 + ("hello" + "hello")) + "") + "\n\t") {
                case "\n\t":
                    var v111 int8 = v95
                    _ = v111
                default:
                    var v112 string = ""
                    _ = v112
                }
                _ = fi109
                _ = v110
            }
            var v113 uint32 = (v14 ^ (v14 % ((v14) | 1)))
            fmt.Println(v96, v96)
            _ = v113
        } else {
            v1 = -473
            v108 = v108
        }
        _ = v108
    } else {
        v15 -= float32(1.5)
    }
    var v114 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v114:", v114)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v103
    _ = v114
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v96
}
