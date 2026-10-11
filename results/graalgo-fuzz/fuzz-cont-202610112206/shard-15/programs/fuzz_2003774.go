package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 int64; F2 int16 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS11 struct { F0 string; F1 uint; F2 int64 }
type GenS12 struct { F0 int8 }
func genFn0() uint8 {
    switch ("" + "") {
    case "foo bar":
        var v6 uint64 = 1
        _ = v6
    case "\n\t":
        var v7 uint8 = 255
        _ = v7
    }
    v8 := GenS1{F0: 127, F1: 960}
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 7; fi10++ {
        v9 = v9
        if (!true) {
            var v11 string = v9
            _ = v11
        } else {
            if true {
                v12 := GenS1{F0: 0, F1: 4294967295}
                if ((!(true && (!(1e10 == 0.1)))) || (506 < 0)) {
                    fmt.Println(fi10)
                } else {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 uint32 = 65535
                        v14 = v14
                        _ = fi13
                        _ = v14
                    }
                    if true {
                        switch (-(fi10)) {
                        case -9223372036854775808:
                            var v15 float32 = 1e10
                            _ = v15
                        case -1:
                            if ((18446744073709551615 >= 4294967295) && false) {
                                var v16 int32 = -1
                                var v17 uint32 = 2147483648
                                _ = v16
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    v19 := []int{fi10}
                                    if (!false) {
                                        fmt.Println(v9, fi10, fi18)
                                    } else {
                                        v19 = append(v19, (fi10 ^ (fi18 + (-(fi18)))))
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                            }
                        }
                        v20 := GenS2{F0: 0}
                        _ = v20
                    } else {
                        v9 = v9
                        if ((!(false && false)) || false) {
                            v9 = ("hello" + v9)
                            if false {
                                v21 := GenS1{F0: int8(fi10), F1: 2147483648}
                                _ = v21
                            } else {
                                v9 = (v9 + v9)
                            }
                        }
                        fmt.Println(fi10, v9)
                    }
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 uint8 = 2
                        if (true && (42 >= 171)) {
                            switch v9 {
                            case "\n\t":
                                switch (fi10 / ((fi10) | 1)) {
                                case 661:
                                    v9 = (v9 + (v9 + "hello"))
                                case 128:
                                    v9 = "\n\t"
                                    v24 := []int{fi22}
                                    var v25 string = (((v9 + "a") + "a") + (((v9 + (v9 + v9)) + "foo bar") + v9))
                                    _ = v24
                                    _ = v25
                                case 42:
                                    var v26 bool = false
                                    _ = v26
                                }
                            case "a":
                                var v27 int = fi22
                                _ = v27
                            case "foo bar":
                                var v28 uint16 = 293
                                _ = v28
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                var v29 uint16 = 0
                _ = v12
                _ = v29
            } else {
                if ((fi10 | (fi10 - fi10)) == (fi10 - (fi10 ^ 32768))) {
                    var v30 uint32 = 0
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v32 := GenS3{F0: 65535}
                        v33 := []int{}
                        var v34 int16 = -32768
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    if (false && true) {
                        switch (-(fi10)) {
                        case 32767:
                            v35 := GenS3{F0: 536}
                            _ = v35
                        default:
                            v36 := GenS1{F0: -1, F1: 1}
                            _ = v36
                        }
                    } else {
                        switch (fi10 ^ fi10) {
                        case 737:
                            v9 = ""
                        case -2147483648:
                            switch (^(fi10)) {
                            case -1:
                                v37 := GenS2{F0: 778}
                                _ = v37
                            }
                        default:
                            v9 = "a"
                        }
                        v38 := GenS4{F0: 65535, F1: 1, F2: ("foo bar" + (v9 + v9))}
                        if (!false) {
                            switch "a" {
                            case "\n\t":
                                var v39 int16 = 32767
                                _ = v39
                            case "":
                                v9 = ""
                            case "hello":
                                v40 := GenS4{F0: 0, F1: 4294967295, F2: v9}
                                _ = v40
                            }
                            fmt.Println(fi10, v9)
                            var v41 string = "\n\t"
                            _ = v41
                        }
                        _ = v38
                    }
                }
            }
            v9 = ""
        }
        _ = fi10
    }
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 float32 = 1e10
    var v5 bool = (4294967295 >= 4294967295)
    if (v5 || (42 == 40)) {
        v1 = v1
        if true {
            if (v5 || true) {
                fmt.Println(v4, v4)
                v2 -= v2
                v4 = v4
            }
            v5 = (true && v5)
            var v42 int32 = -675
            _ = v42
        }
    }
    var v43 int32 = 42
    if true {
        var v44 int16 = 1
        _ = v44
    } else {
        if (!v5) {
            for fi45 := 0; fi45 < 11; fi45++ {
                if v5 {
                    switch 127 {
                    case 0:
                        var v46 bool = true
                        _ = v46
                    case -15:
                        var v47 bool = v5
                        _ = v47
                    default:
                        var v48 int8 = 42
                        _ = v48
                    }
                } else {
                    if false {
                        if (!v5) {
                            if (2147483648 == 1) {
                                v49 := GenS4{F0: 0, F1: 1, F2: (v1 + ("foo bar" + "foo bar"))}
                                _ = len(v3)
                                v4 += (-2.25 + float32((float64((v4 + (v4 + v4))) * 1.5)))
                                _ = v49
                            }
                            v50 := GenS3{F0: 1}
                            _ = v50
                        } else {
                            v51 := []int{}
                            fmt.Println(v4, v4, fi45)
                            var v52 int64 = -1
                            _ = v51
                            _ = v52
                        }
                        if true {
                            v53 := GenS2{F0: 184}
                            _ = v53
                        }
                    } else {
                        var v54 int32 = v43
                        if (0 < 0) {
                            if (!true) {
                                if v5 {
                                    switch (v2 * v2) {
                                    case 127:
                                        delete(v3, 2147483647)
                                    }
                                    v55 := GenS1{F0: 13, F1: 42}
                                    _ = v55
                                } else {
                                    if (v1 >= ((("a" + v1) + v1) + "")) {
                                        v54 -= -1
                                        v2 += v2
                                    } else {
                                        var v56 int8 = 42
                                        _ = v56
                                    }
                                    var v57 float64 = 0.1
                                    for fi58 := 0; fi58 < 10; fi58++ {
                                        v59 := GenS4{F0: 1, F1: 18446744073709551615, F2: "foo bar"}
                                        _ = fi58
                                        _ = v59
                                    }
                                    _ = v57
                                }
                                var v60 uint32 = 2
                                switch v2 {
                                case 127:
                                    var v61 int8 = -66
                                    _ = v61
                                case -449:
                                    if v5 {
                                        v54 -= v54
                                        if v5 {
                                            delete(v3, 65535)
                                            if true {
                                                fmt.Println(v43)
                                                _ = len(v3)
                                            } else {
                                                v62 := GenS5{F0: ((!(-128 != -50)) || v5)}
                                                v63 := GenS5{F0: (1 != -1)}
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                        v64 := []int{}
                                        _ = v64
                                    } else {
                                        v60 -= v60
                                        if (v60 <= 545) {
                                            if ("\n\t" != v1) {
                                                switch v1 {
                                                case "hello":
                                                    for fi65 := 0; fi65 < 2; fi65++ {
                                                        fmt.Println(v60, v1, v54)
                                                        _ = len(v3)
                                                        _ = fi65
                                                    }
                                                default:
                                                    for fi66 := 0; fi66 < 11; fi66++ {
                                                        var v67 uint = 296
                                                        var v68 int32 = v54
                                                        _ = fi66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            } else {
                                                v69 := []int{fi45, fi45, fi45, 128}
                                                _ = v69
                                            }
                                        } else {
                                            if (1 >= 2) {
                                                for fi70 := 0; fi70 < 17; fi70++ {
                                                    v71 := v3[(v60 % ((664) | 1))]
                                                    v3[(v60 << ((4294967295) & 63))] = v2
                                                    if v5 {
                                                        v1 = "foo bar"
                                                        v60 -= 2
                                                        if (v5 || true) {
                                                            if (!(v5 || false)) {
                                                                v1 = v1
                                                                v72 := GenS3{F0: 877}
                                                                if (1 > -1) {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v73 := []int{4294967295}
                                                                        _ = v73
                                                                    case "":
                                                                        fmt.Println(v71)
                                                                        v5 = false
                                                                        v74 := v3[42]
                                                                        _ = v74
                                                                    case "a":
                                                                        v75 := []int{(v71 << ((2) & 63)), 127, v2, (v2 * 9223372036854775807)}
                                                                        _ = v75
                                                                    default:
                                                                        v76 := GenS2{F0: 199}
                                                                        _ = v76
                                                                    }
                                                                }
                                                                _ = v72
                                                            }
                                                        } else {
                                                            switch v1 {
                                                            case "":
                                                                var v77 int8 = -128
                                                                _ = v77
                                                            case "\n\t":
                                                                var v78 bool = true
                                                                _ = v78
                                                            default:
                                                                switch "a" {
                                                                case "":
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        var v79 int32 = 42
                                                                        _ = v79
                                                                    case "\n\t":
                                                                        v80 := GenS6{F0: 65535, F1: -9223372036854775808, F2: 1000}
                                                                        _ = v80
                                                                    }
                                                                }
                                                            }
                                                            if v5 {
                                                                v81 := GenS7{F0: fi70, F1: 32767}
                                                                fmt.Println(v5, v71, v5)
                                                                _ = v81
                                                            }
                                                            v3[(v60 - v60)] = (-(fi45))
                                                            v71 += (v71 + (^(fi45)))
                                                        }
                                                    } else {
                                                        if v5 {
                                                            switch (fi45 * -2147483648) {
                                                            case -9223372036854775808:
                                                                var v82 int64 = 533
                                                                v3[v60] = -737
                                                                switch (v1 + v1) {
                                                                case "":
                                                                    var v83 uint = 18446744073709551615
                                                                    _ = v83
                                                                case "foo bar":
                                                                    var v84 int = fi45
                                                                    _ = v84
                                                                case "\n\t":
                                                                    if (v54 > v54) {
                                                                        var v85 float64 = (1e10 * 1e10)
                                                                        var v86 bool = v5
                                                                        v4 = 0.0
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                }
                                                                var v87 uint32 = v60
                                                                _ = v82
                                                                _ = v87
                                                            }
                                                            var v88 string = ""
                                                            _ = v88
                                                        }
                                                    }
                                                    v89 := GenS1{F0: -61, F1: (v60 | v60)}
                                                    _ = fi70
                                                    _ = v71
                                                    _ = v89
                                                }
                                                var v90 int32 = v43
                                                _ = v90
                                            } else {
                                                v91 := GenS5{F0: true}
                                                _ = v91
                                            }
                                            delete(v3, (v60 / ((265) | 1)))
                                        }
                                        v92 := GenS3{F0: 2147483648}
                                        _ = v92
                                    }
                                case 128:
                                    switch v1 {
                                    case "\n\t":
                                        var v93 uint64 = 613
                                        _ = v93
                                    case "a":
                                        v1 = "\n\t"
                                    default:
                                        v43 = v54
                                    }
                                }
                                _ = v60
                            } else {
                                for fi94 := 0; fi94 < 15; fi94++ {
                                    v3[2147483647] = fi45
                                    var v95 float64 = (((float64(float32(float64(-2.25))) / float64(v4)) - (1e10 / 1.0)) - 3.14159)
                                    if ((v5 || (2147483648 < 40)) || (128 == 0)) {
                                        if (!v5) {
                                            v4 -= 1.5
                                        } else {
                                            v96 := GenS3{F0: 65535}
                                            fmt.Println(fi94)
                                            _ = v96
                                        }
                                    }
                                    _ = fi94
                                    _ = v95
                                }
                            }
                        } else {
                            switch 42 {
                            case 792:
                                if (-976 <= 1000) {
                                    var v97 uint64 = 0
                                    v98 := GenS2{F0: -9223372036854775808}
                                    v99 := GenS5{F0: v5}
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                } else {
                                    for fi100 := 0; fi100 < 17; fi100++ {
                                        v101 := GenS8{F0: 1000, F1: 1, F2: 128}
                                        if false {
                                            v102 := GenS9{F0: 32767}
                                            var v103 uint64 = 484
                                            _ = v102
                                            _ = v103
                                        } else {
                                            switch (fi45 >> ((2) & 63)) {
                                            case -96:
                                                if (4294967295 >= 873) {
                                                    v104 := GenS6{F0: 2, F1: -1, F2: -1}
                                                    _ = v104
                                                } else {
                                                    var v105 bool = (!true)
                                                    _ = v105
                                                }
                                            case 128:
                                                var v106 int8 = 100
                                                _ = v106
                                            case 32768:
                                                switch v1 {
                                                case "hello":
                                                    v54 += (v43 << ((2) & 63))
                                                case "foo bar":
                                                    v107 := GenS10{F0: 2, F1: 701, F2: 17}
                                                    _ = v107
                                                case "\n\t":
                                                    v43 -= -911
                                                default:
                                                    fmt.Println(v1, v5, v1)
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            v54 = (v43 - -2147483648)
                                        }
                                        v5 = false
                                        _ = fi100
                                        _ = v101
                                    }
                                }
                            default:
                                v108 := GenS11{F0: (v1 + v1), F1: 65535, F2: -9223372036854775808}
                                _ = v108
                            }
                            v5 = v5
                        }
                        v43 = v54
                        _ = v54
                    }
                }
                _ = fi45
            }
            if (!((1 != 2) || (18 <= 3))) {
                fmt.Println(v4, v4)
                v109 := []int{(v2 >> ((255) & 63)), v2}
                v3[42] = (v2 % ((v2) | 1))
                _ = v109
            }
        } else {
            v4 = v4
            v110 := GenS8{F0: -695, F1: 32767, F2: 128}
            _ = v110
        }
        _ = len(v3)
        fmt.Println(v5, v2)
    }
    v2 = 27
    fmt.Println(v1)
    fmt.Println(v1, v5)
    for fi111 := 0; fi111 < 13; fi111++ {
        _ = len(v3)
        v112 := GenS10{F0: 2, F1: 1, F2: 42}
        _ = len(v3)
        _ = fi111
        _ = v112
    }
    delete(v3, 434)
    var v113 float64 = 0.0
    var v114 bool = v5
    var v115 int16 = 0
    for fi116 := 0; fi116 < 14; fi116++ {
        v117 := GenS12{F0: 127}
        v4 -= v4
        if v5 {
            v118 := GenS10{F0: 741, F1: 231, F2: 964}
            _ = len(v3)
            _ = v118
        }
        _ = fi116
        _ = v117
    }
    var v119 uint64 = 42
    var v120 uint32 = 1
    v2 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
