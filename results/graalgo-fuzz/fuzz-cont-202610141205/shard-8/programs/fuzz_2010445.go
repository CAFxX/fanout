package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint8; F1 string; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 uint64; F2 uint32 }
func genFn0(p0 int8, p1 int) int8 {
    if true {
        if true {
            p0 = -1
            if false {
                v6 := GenS1{F0: -532}
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 uint64 = 1
                    if (!true) {
                        v9 := []int{(fi7 - (p1 / (((p1 | fi7)) | 1))), (fi7 * p1), p1, p1}
                        _ = v9
                    }
                    var v10 int32 = -313
                    _ = fi7
                    _ = v10
                    _ = v8
                }
                _ = v6
            }
        }
        if (!false) {
            v11 := []int{p1, p1, (p1 / ((p1) | 1)), (p1 - (-(p1)))}
            if true {
                var v12 uint64 = 2
                _ = v12
            } else {
                p1 = p1
            }
            _ = v11
        } else {
            p0 -= -20
            p0 = p0
            if ((!true) || false) {
                p1 -= p1
                p0 -= p0
            }
        }
    }
    if (!false) {
        var v13 int32 = -250
        _ = v13
    }
    v14 := GenS2{F0: 2}
    var v15 uint = 18446744073709551615
    v16 := []int{}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 62
}
func genFn1(p0 int16, p1 string, p2 int16) uint16 {
    var v17 float32 = 1.5
    v18 := GenS2{F0: 4294967295}
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 float32 = 0.1
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 491}
    p0 = (p2 ^ 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v21
    return 65535
}

func main() {
    var v1 int16 = 0
    var v2 int64 = -242
    var v3 bool = false
    var v4 int = int(v1)
    var v5 uint8 = 255
    fmt.Println(v1, v1)
    switch ("" + "") {
    case "foo bar":
        v22 := GenS3{F0: 42, F1: (42 == uint(v5))}
        _ = v22
    case "hello":
        fmt.Println(v2, v4, v4)
    case "":
        var v23 int16 = (v1 | v1)
        _ = v23
    }
    if false {
        var v24 uint64 = 651
        v24 = (v24 | 18446744073709551615)
        v24 = (v24 ^ 289)
        _ = v24
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint = 1
                v5 += (^(v5))
                _ = fi25
                _ = v26
            }
        default:
            v27 := GenS2{F0: 18446744073709551615}
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        v4 = (v4 | v4)
        _ = fi28
    }
    switch (-(v4)) {
    case 42:
        v29 := GenS4{F0: (v5 << ((1) & 63)), F1: ("hello" + ("" + "\n\t")), F2: 478}
        _ = v29
    case 2147483647:
        v2 = v2
    case 315:
        for fi30 := 0; fi30 < 8; fi30++ {
            v2 = (v2 & (v2 + v2))
            _ = fi30
        }
    default:
        var v31 int64 = v2
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := GenS1{F0: 42}
            fmt.Println(v4)
            if v3 {
                v3 = (0.0 != float64(3.14159))
                v5 += (^(v5))
                v4 += (v4 + fi33)
            }
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        if v3 {
            for fi36 := 0; fi36 < 8; fi36++ {
                fmt.Println(fi35)
                v1 += v1
                _ = fi36
            }
            v4 = v4
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 int32 = 2147483647
            v3 = v3
            _ = fi37
            _ = v38
        }
        _ = fi35
    }
    fmt.Println(v2, v3)
    if (42 > -1) {
        v2 -= v2
        switch "a" {
        case "foo bar":
            if v3 {
                var v39 uint32 = 0
                _ = v39
            } else {
                var v40 uint64 = 701
                fmt.Println(v1, v5, v4)
                _ = v40
            }
        case "\n\t":
            switch ("hello" + "a") {
            case "foo bar":
                v41 := []int{(-(v4))}
                _ = v41
            case "hello":
                switch -1 {
                case 1:
                    v1 -= -32768
                case 0:
                    if ((0 <= 1) || ((v1 % (((v1 % ((v1) | 1))) | 1)) <= 0)) {
                        v1 = v1
                        var v42 int16 = 1
                        switch 127 {
                        case 479:
                            fmt.Println(v5, v42, v4)
                        case -9223372036854775808:
                            var v43 int8 = -1
                            _ = v43
                        case 0:
                            var v44 string = "\n\t"
                            _ = v44
                        }
                        _ = v42
                    }
                }
            default:
                v4 -= -476
            }
        case "a":
            if false {
                if (!(954 > 1)) {
                    if (!(!v3)) {
                        if (v3 && (!((v2 & (v2 & v2)) <= v2))) {
                            for fi45 := 0; fi45 < 18; fi45++ {
                                v1 -= 652
                                if (2147483648 != 4294967295) {
                                    v46 := []int{(v4 * (v4 * (^(v4)))), (v4 + v4), v4, 128}
                                    v5 += 252
                                    _ = v46
                                } else {
                                    var v47 uint8 = 222
                                    var v48 uint = 18446744073709551615
                                    v1 += (v1 - 1000)
                                    _ = v47
                                    _ = v48
                                }
                                if (v3 || v3) {
                                    var v49 float64 = 0.0
                                    switch "a" {
                                    case "a":
                                        v50 := []int{fi45, v4, (fi45 + v4)}
                                        _ = v50
                                    case "hello":
                                        var v51 uint32 = 427
                                        _ = v51
                                    case "":
                                        v52 := []int{(^(fi45))}
                                        _ = v52
                                    default:
                                        fmt.Println(v4)
                                    }
                                    _ = v49
                                } else {
                                    v4 -= fi45
                                }
                                _ = fi45
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                switch ("hello" + "hello") {
                                case "a":
                                    for fi53 := 0; fi53 < 18; fi53++ {
                                        v3 = (!(-9223372036854775808 == (^(v2))))
                                        var v54 uint16 = uint16(v5)
                                        v55 := []int{1, v4, (v4 | (fi53 + (fi53 - -2147483648))), v4}
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            case "foo bar":
                                if v3 {
                                    if true {
                                        for fi56 := 0; fi56 < 2; fi56++ {
                                            var v57 float64 = (float64(((1.5 - 0.0) - (1e10 / 1.5))) * 1e10)
                                            for fi58 := 0; fi58 < 8; fi58++ {
                                                v4 += v4
                                                switch (fi56 | int(v5)) {
                                                case 32767:
                                                    v3 = (42 != 0)
                                                case 36:
                                                    var v59 int8 = genFn0(-125, (v4 / ((fi58) | 1)))
                                                    _ = v59
                                                default:
                                                    var v60 bool = (((((!(float32((v57 / v57)) >= (0.1 * ((0.0 + ((0.1 + ((0.1 * (((1.5 + -2.25) / 1.0) * 1e10)) * 0.0)) - (((3.14159 / 1.0) / ((1e10 + 1.5) * (0.1 + 0.1))) + 1e10))) * 0.0)))) || false) || (((v4 << ((255) & 63)) > (v4 % ((128) | 1))) || (!((-2.25 / (-2.25 * 0.1)) < 3.14159)))) && (!false)) || ((v4 < v4) && false))
                                                    _ = v60
                                                }
                                                _ = fi58
                                            }
                                            for fi61 := 0; fi61 < 8; fi61++ {
                                                if (663 != 514) {
                                                    fmt.Println(v3)
                                                } else {
                                                    var v62 int16 = (v1 << ((47) & 63))
                                                    v62 -= (v62 * v1)
                                                    _ = v62
                                                }
                                                switch fi56 {
                                                case 776:
                                                    fmt.Println(v4)
                                                case -950:
                                                    var v63 uint32 = 2147483648
                                                    _ = v63
                                                case 0:
                                                    v64 := GenS3{F0: 1, F1: false}
                                                    _ = v64
                                                }
                                                var v65 float32 = (0.1 / 1.0)
                                                _ = fi61
                                                _ = v65
                                            }
                                            _ = fi56
                                            _ = v57
                                        }
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    for fi66 := 0; fi66 < 12; fi66++ {
                                        v3 = ((v1 ^ v1) < v1)
                                        v5 = v5
                                        _ = fi66
                                    }
                                    var v67 int32 = 0
                                    _ = v67
                                } else {
                                    v68 := GenS4{F0: 42, F1: "hello", F2: 255}
                                    switch "hello" {
                                    case "a":
                                        if (!((v1 & v1) <= v1)) {
                                            if false {
                                                for fi69 := 0; fi69 < 5; fi69++ {
                                                    for fi70 := 0; fi70 < 11; fi70++ {
                                                        var v71 float32 = ((((0.1 + ((float32(((0.0 * 1e10) * 3.14159)) - ((float32((-2.25 + 0.1)) + 0.1) * -2.25)) + ((1e10 / 1.0) * 1.5))) / 1.5) / 3.14159) * (0.0 * (((0.0 + (-2.25 + 1.5)) / ((float32((0.0 - 1e10)) + ((3.14159 + ((0.1 - float32(1e10)) - ((0.0 + 0.1) * -2.25))) / 3.14159)) / -2.25)) + 0.1)))
                                                        _ = fi70
                                                        _ = v71
                                                    }
                                                    _ = fi69
                                                }
                                                switch "a" {
                                                case "":
                                                    if ("foo bar" > ("" + ("hello" + "hello"))) {
                                                        v72 := GenS5{F0: 42}
                                                        _ = v72
                                                    }
                                                default:
                                                    var v73 uint16 = 255
                                                    _ = v73
                                                }
                                                var v74 float32 = 0.0
                                                _ = v74
                                            }
                                            for fi75 := 0; fi75 < 13; fi75++ {
                                                v76 := []int{}
                                                _ = fi75
                                                _ = v76
                                            }
                                        }
                                    default:
                                        if (("a" < "foo bar") && v3) {
                                            v77 := GenS1{F0: (-(v1))}
                                            v78 := GenS6{F0: (^(v5))}
                                            fmt.Println(v1)
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    v79 := GenS7{F0: "a", F1: v5, F2: (v4 >> ((65535) & 63))}
                                    _ = v68
                                    _ = v79
                                }
                            default:
                                var v80 int32 = -41
                                _ = v80
                            }
                        }
                        if v3 {
                            fmt.Println(v1)
                            v3 = (!false)
                        }
                        for fi81 := 0; fi81 < 20; fi81++ {
                            var v82 string = "\n\t"
                            v83 := []int{}
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                    } else {
                        var v84 float32 = (1.5 - 0.0)
                        if (!v3) {
                            for fi85 := 0; fi85 < 11; fi85++ {
                                v86 := GenS4{F0: 2, F1: "foo bar", F2: 255}
                                _ = fi85
                                _ = v86
                            }
                            v87 := []int{v4, 2147483647, -9223372036854775808, (v4 % ((v4) | 1))}
                            _ = v87
                        }
                        _ = v84
                    }
                    if v3 {
                        var v88 float32 = 3.14159
                        v89 := []int{}
                        _ = v88
                        _ = v89
                    } else {
                        v90 := GenS8{F0: v5}
                        v91 := []int{v4, v4, (v4 + 547), 128}
                        _ = v90
                        _ = v91
                    }
                }
                var v92 int64 = v2
                v93 := []int{(-(v4)), 32767, (^(v4))}
                _ = v92
                _ = v93
            } else {
                v5 = (^(v5))
                v94 := GenS9{F0: int64(v1)}
                var v95 uint8 = v5
                _ = v94
                _ = v95
            }
        default:
            v96 := []int{int(v2), (v4 >> ((448) & 63)), 42}
            _ = v96
        }
    }
    switch (v4 & v4) {
    case 0:
        var v97 int = v4
        _ = v97
    default:
        if v3 {
            var v98 string = "\n\t"
            var v99 bool = ((v5 * (v5 - v5)) == v5)
            var v100 int64 = 1
            _ = v100
            _ = v98
            _ = v99
        } else {
            v2 -= (^(v2))
            for fi101 := 0; fi101 < 14; fi101++ {
                if v3 {
                    if (65535 <= 914) {
                        var v102 uint16 = 109
                        _ = v102
                    } else {
                        for fi103 := 0; fi103 < 12; fi103++ {
                            v104 := GenS10{F0: -601, F1: 481, F2: 2147483648}
                            _ = fi103
                            _ = v104
                        }
                        var v105 int32 = 0
                        if v3 {
                            v106 := []int{fi101, v4, -351, (fi101 - 1)}
                            fmt.Println(v4, fi101, v2)
                            _ = v106
                        }
                        _ = v105
                    }
                    var v107 int64 = (v2 << ((791) & 63))
                    v107 += v2
                    _ = v107
                } else {
                    v108 := GenS1{F0: (-(v1))}
                    if (v3 && (uint16(v5) > 453)) {
                        v4 = (v4 << ((42) & 63))
                        var v109 bool = (v3 && v3)
                        _ = v109
                    } else {
                        v110 := []int{fi101, (fi101 - 538)}
                        _ = v110
                    }
                    v3 = ((v2 << ((715) & 63)) == v2)
                    _ = v108
                }
                _ = fi101
            }
            v5 = 255
        }
    }
    fmt.Println(v2, v2)
    var v111 int8 = -1
    if ((1.5 - 0.1) == 3.14159) {
        var v112 int8 = (v111 * (v111 >> ((4294967295) & 63)))
        _ = v112
    }
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
