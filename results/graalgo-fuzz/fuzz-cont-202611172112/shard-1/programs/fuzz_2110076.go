package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int16; F1 uint8; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int64; F2 uint }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 int8; F1 int64; F2 int16 }
type GenS12 struct { F0 int32 }
func genFn0(p0 string, p1 int32) uint32 {
    var v6 string = "a"
    if (!false) {
        v7 := []int{2147483647, -389, -1, -460}
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                var v9 string = ((p0 + (("a" + "") + ("" + "\n\t"))) + p0)
                switch (fi8 | int(p1)) {
                case 128:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        v6 = (("foo bar" + p0) + p0)
                        v9 = p0
                        _ = fi10
                    }
                }
                _ = v9
            } else {
                var v11 int = fi8
                fmt.Println(fi8, p1, v6)
                _ = v11
            }
            p0 = p0
            if true {
                p0 = (v6 + "")
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v12 uint64 = 42
                _ = v12
            }
            _ = fi8
        }
        _ = v7
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int64 = 1
        var v15 int64 = 9223372036854775807
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint = 18446744073709551615
        if false {
            v18 := GenS1{F0: fi16, F1: 32767, F2: 142}
            var v19 int = fi16
            var v20 int = (fi16 ^ -209)
            _ = v18
            _ = v19
            _ = v20
        } else {
            v17 = v17
            fmt.Println(v17, v6, p0)
        }
        if true {
            switch 9223372036854775807 {
            case 1:
                var v21 float64 = 0.1
                _ = v21
            case 127:
                v22 := GenS2{F0: 175}
                _ = v22
            case 467:
                v17 = v17
            default:
                if (false && (1 == 1)) {
                    v17 += (v17 - v17)
                    var v23 float64 = 0.1
                    switch (^(fi16)) {
                    case 128:
                        var v24 uint = v17
                        _ = v24
                    }
                    _ = v23
                } else {
                    p1 = -2147483648
                }
            }
        } else {
            var v25 int64 = 20
            _ = v25
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 float32 = 0.1
    v4 := make(map[int]bool)
    var v5 int8 = 59
    if (!(42 > 1)) {
        fmt.Println(v3)
        if false {
            v26 := GenS3{F0: 65535, F1: 85}
            v27 := []int{32768, -1, 2147483647, -469}
            var v28 uint = 2
            _ = v26
            _ = v27
            _ = v28
        } else {
            delete(v4, -9223372036854775808)
            for fi29 := 0; fi29 < 13; fi29++ {
                if (false && true) {
                    v30 := GenS2{F0: 0}
                    if true {
                        v31 := GenS2{F0: 258}
                        var v32 int16 = -32768
                        var v33 uint32 = genFn0(v1, 0)
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v1 = v1
                    _ = v30
                } else {
                    var v34 bool = true
                    if v34 {
                        if false {
                            if v34 {
                                var v35 int8 = (v5 ^ 0)
                                v35 = (v5 * v5)
                                fmt.Println(fi29)
                                _ = v35
                            }
                            v36 := []int{(fi29 * 32768), (fi29 | fi29), 127}
                            _ = v36
                        }
                    }
                    v37 := GenS4{F0: 706, F1: 42, F2: 633}
                    _ = v34
                    _ = v37
                }
                switch (v1 + v1) {
                case "hello":
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v1 = v1
                        switch (-(fi29)) {
                        case 0:
                            if true {
                                v39 := []int{fi38}
                                _ = v39
                            } else {
                                var v40 bool = ((-32768 < -1) || (false || true))
                                if ((!(!(!(3.14159 == 3.14159)))) && v40) {
                                    v1 = v1
                                } else {
                                    v3 = 1e10
                                }
                                _ = v40
                            }
                        case 127:
                            v1 = ("" + v1)
                        case 367:
                            fmt.Println(v5)
                        default:
                            v41 := GenS5{F0: 4294967295}
                            _ = v41
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v3, v5, v1)
                        case "a":
                            fmt.Println(fi38)
                        case "foo bar":
                            v42 := []int{(fi38 % (((fi38 / ((fi38) | 1))) | 1)), (fi38 ^ fi38)}
                            _ = v42
                        }
                        v5 = (^(v5))
                        _ = fi38
                    }
                }
                if ((^(fi29)) >= 345) {
                    for fi43 := 0; fi43 < 16; fi43++ {
                        fmt.Println(fi29, v1)
                        var v44 bool = false
                        v5 -= (v5 - -1)
                        v45 := []int{fi29}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                } else {
                    var v46 uint32 = 1
                    _ = v46
                }
                _ = fi29
            }
            v2[32767] = false
        }
        v47 := GenS2{F0: 1}
        _ = v47
    } else {
        for fi48 := 0; fi48 < 14; fi48++ {
            if true {
                v49 := GenS6{F0: 2, F1: 2147483647, F2: 0}
                v50 := GenS5{F0: 43}
                if (!(false || true)) {
                    v5 = v5
                }
                _ = v49
                _ = v50
            } else {
                if false {
                    if (false && false) {
                        switch (v1 + (v1 + "hello")) {
                        case "a":
                            for fi51 := 0; fi51 < 16; fi51++ {
                                fmt.Println(v5, v1)
                                v52 := GenS7{F0: 42}
                                _ = fi51
                                _ = v52
                            }
                        }
                        var v53 int16 = 1
                        fmt.Println(v1)
                        _ = v53
                    } else {
                        if (630 >= 0) {
                            v54 := []int{(fi48 / ((fi48) | 1))}
                            _ = v54
                        }
                    }
                } else {
                    switch (v1 + ("hello" + v1)) {
                    case "foo bar":
                        if true {
                            v1 = v1
                            var v55 float64 = 1.5
                            _ = v55
                        }
                    default:
                        for fi56 := 0; fi56 < 5; fi56++ {
                            var v57 bool = (-673 >= 1000)
                            _ = fi56
                            _ = v57
                        }
                    }
                }
                fmt.Println(v3)
                if ((false || false) && false) {
                    var v58 uint = 1
                    _ = v58
                } else {
                    v5 = (^(v5))
                }
            }
            v59 := []int{fi48}
            _ = fi48
            _ = v59
        }
    }
    if true {
        for fi60 := 0; fi60 < 10; fi60++ {
            var v61 bool = (42 <= -310)
            for fi62 := 0; fi62 < 1; fi62++ {
                v63 := GenS1{F0: fi62, F1: 42, F2: 127}
                v1 = ((("foo bar" + (v1 + "")) + v1) + "")
                _ = len(v2)
                if v61 {
                    var v64 uint16 = 0
                    fmt.Println(v3)
                    var v65 bool = false
                    _ = v64
                    _ = v65
                } else {
                    v66 := GenS4{F0: 32767, F1: 255, F2: (fi60 | (fi62 | (^(fi60))))}
                    _ = v66
                }
                var v67 int64 = 788
                _ = fi62
                _ = v63
                _ = v67
            }
            _ = fi60
            _ = v61
        }
        switch (v1 + v1) {
        case "hello":
            v68 := []int{-2147483648, 9223372036854775807, -390, 127}
            _ = v68
        case "a":
            if (false && ((389 == 18446744073709551615) || (!false))) {
                v69 := GenS7{F0: 127}
                _ = v69
            }
        default:
            var v70 uint8 = 42
            _ = v70
        }
        for fi71 := 0; fi71 < 3; fi71++ {
            v3 = v3
            v72 := []int{fi71, (-(fi71)), (fi71 + 42)}
            v5 += v5
            _ = fi71
            _ = v72
        }
    }
    if true {
        v73 := GenS3{F0: 0, F1: 4294967295}
        _ = v73
    } else {
        v5 -= v5
    }
    var v74 bool = ((float64((v3 + v3)) * (0.1 / (float64((((float32(1e10) - float32(1.5)) - (((v3 * (1.5 + float32(1.5))) * 1e10) + 1e10)) - float32(0.1))) / 1e10))) >= 0.1)
    switch "foo bar" {
    case "":
        for fi75 := 0; fi75 < 13; fi75++ {
            for fi76 := 0; fi76 < 9; fi76++ {
                if v74 {
                    switch "a" {
                    case "":
                        v77 := GenS8{F0: 42}
                        _ = v77
                    case "hello":
                        v78 := []int{(fi75 + (fi75 << ((1) & 63))), -77}
                        fmt.Println(v1, fi76, v1)
                        _ = v78
                    }
                } else {
                    fmt.Println(fi76, v3, v3)
                    v1 = v1
                    v3 -= (v3 - (v3 / 1.0))
                    switch (((v1 + (((v1 + (("hello" + "a") + "")) + "a") + "hello")) + ("hello" + "foo bar")) + v1) {
                    case "":
                        v79 := GenS4{F0: -32768, F1: 127, F2: (fi76 >> ((65535) & 63))}
                        _ = v79
                    case "a":
                        if (1000 < 1) {
                            switch fi76 {
                            case 9223372036854775807:
                                var v80 uint16 = 2
                                _ = v80
                            case 4294967295:
                                fmt.Println(v74, fi76, v5)
                            default:
                                if false {
                                    if (127 != 1) {
                                        var v81 uint32 = 814
                                        if v74 {
                                            fmt.Println(fi76, v3, fi76)
                                            if v74 {
                                                v3 += (v3 - (0.0 / float32(1e10)))
                                                fmt.Println(fi75, v3, v74)
                                            }
                                            var v82 uint8 = 0
                                            _ = v82
                                        } else {
                                            v1 = "a"
                                        }
                                        _ = v81
                                    } else {
                                        v83 := GenS9{F0: v1, F1: -9223372036854775808, F2: 65535}
                                        v74 = true
                                        v5 = v5
                                        _ = v83
                                    }
                                    v74 = (938 != 2147483647)
                                    if v74 {
                                        var v84 int16 = 32767
                                        _ = v84
                                    } else {
                                        if v74 {
                                            var v85 uint = 65535
                                            v4[(fi76 / (((fi76 + 128)) | 1))] = true
                                            v86 := GenS3{F0: 256, F1: 128}
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                }
                            }
                            v87 := []int{(fi76 * (fi75 - (fi75 + (fi75 / ((fi76) | 1))))), fi75, fi76}
                            _ = v87
                        } else {
                            if v74 {
                                v88 := GenS3{F0: 42, F1: 32767}
                                switch v1 {
                                case "":
                                    var v89 uint16 = 591
                                    _ = v89
                                case "foo bar":
                                    var v90 int8 = 127
                                    _ = v90
                                case "\n\t":
                                    var v91 int16 = 698
                                    if v74 {
                                        switch v1 {
                                        case "hello":
                                            delete(v4, fi76)
                                        }
                                        v92 := GenS10{F0: 32767, F1: 255}
                                        _ = v92
                                    }
                                    switch v1 {
                                    case "a":
                                        v4[128] = (!v74)
                                    case "foo bar":
                                        fmt.Println(v5, fi75)
                                    case "hello":
                                        switch (fi75 << ((65535) & 63)) {
                                        case -9223372036854775808:
                                            v5 = v5
                                        }
                                    default:
                                        switch fi76 {
                                        case 32767:
                                            switch (fi75 >> ((364) & 63)) {
                                            case 128:
                                                v93 := GenS2{F0: 4294967295}
                                                v1 = v1
                                                v94 := GenS5{F0: 702}
                                                _ = v93
                                                _ = v94
                                            case 1:
                                                fmt.Println(fi76, v74, fi75)
                                            }
                                            v3 = 1e10
                                            v95 := []int{}
                                            _ = v95
                                        }
                                    }
                                    _ = v91
                                }
                                v96 := []int{943, 127, 1}
                                _ = v88
                                _ = v96
                            } else {
                                if v74 {
                                    _ = len(v2)
                                } else {
                                    var v97 uint64 = 831
                                    _ = v97
                                }
                                v5 = -20
                                v1 = "hello"
                                var v98 string = "hello"
                                fmt.Println(v98, v74)
                                _ = v98
                            }
                        }
                        v99 := GenS8{F0: 65535}
                        _ = v99
                    default:
                        fmt.Println(v74, v74, v74)
                    }
                }
                _ = fi76
            }
            _ = fi75
        }
    case "a":
        var v100 uint32 = 856
        _ = v100
    }
    v74 = (32767 <= 255)
    for fi101 := 0; fi101 < 5; fi101++ {
        if (false || (true || v74)) {
            switch v1 {
            case "\n\t":
                if (v1 > v1) {
                    v102 := GenS9{F0: (v1 + v1), F1: 42, F2: 42}
                    _ = v102
                } else {
                    fmt.Println(v3, v3)
                }
            }
            v103 := []int{127, 4294967295}
            var v104 uint32 = 965
            _ = v103
            _ = v104
        }
        _ = fi101
    }
    switch 42 {
    case 128:
        v105 := []int{42, 665, 0, 128}
        _ = v105
    case -998:
        for fi106 := 0; fi106 < 7; fi106++ {
            for fi107 := 0; fi107 < 17; fi107++ {
                var v108 float32 = float32(0.1)
                _ = fi107
                _ = v108
            }
            fmt.Println(fi106)
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                var v109 uint32 = genFn0((v1 + (((v1 + ("a" + v1)) + (("\n\t" + v1) + ((((("\n\t" + (v1 + ("foo bar" + ((v1 + "a") + (((v1 + v1) + v1) + (v1 + v1)))))) + v1) + "hello") + "") + "a"))) + "a")), -2147483648)
                _ = v109
            default:
                var v110 int8 = (v5 - (-(v5)))
                _ = v110
            }
            _ = fi106
        }
    case -430:
        v2[-9223372036854775808] = ((v5 / (((-(v5))) | 1)) < -67)
    }
    v74 = true
    v111 := GenS5{F0: 18446744073709551615}
    v112 := []int{-9223372036854775808, -2147483648, 42}
    v113 := GenS11{F0: (v5 ^ (v5 / (((v5 << ((18446744073709551615) & 63))) | 1))), F1: 0, F2: 42}
    var v114 string = v1
    if false {
        var v115 int8 = (v5 % ((-79) | 1))
        for fi116 := 0; fi116 < 12; fi116++ {
            if len(v112) >= 2 {
                v112 = v112[1:len(v112)]
            }
            for fi117 := 0; fi117 < 9; fi117++ {
                var v118 int16 = -32768
                v118 -= 704
                if (v114 <= (v114 + v1)) {
                    var v119 uint = 558
                    v120 := GenS7{F0: fi117}
                    _ = v119
                    _ = v120
                } else {
                    var v121 int64 = 1
                    var v122 float32 = v3
                    _ = v121
                    _ = v122
                }
                _ = fi117
                _ = v118
            }
            _ = fi116
        }
        v123 := GenS12{F0: 851}
        _ = v115
        _ = v123
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
}
