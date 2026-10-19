package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string }
func genFn0() uint64 {
    var v6 uint8 = 127
    v6 = v6
    if true {
        v7 := GenS1{F0: 127, F1: 18446744073709551615, F2: 0}
        v6 = (v6 - (^(v6)))
        if (32767 > 42) {
            var v8 uint64 = 403
            _ = v8
        }
        _ = v7
    } else {
        v6 = v6
        v9 := GenS2{F0: 32767}
        if false {
            var v10 uint64 = 227
            var v11 uint8 = (^(v6))
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint32 = 0
                _ = fi12
                _ = v13
            }
            v14 := GenS2{F0: 255}
            _ = v14
        }
        _ = v9
    }
    _ = v6
    return 882
}
func genFn1() int64 {
    v15 := []int{1, -2147483648, 128}
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS1{F0: -128, F1: 2, F2: 2147483648}
            _ = fi17
            _ = v18
        }
        var v19 uint = 4294967295
        v19 = v19
        _ = fi16
        _ = v19
    }
    var v20 bool = true
    _ = v15
    _ = v20
    return 9223372036854775807
}

func main() {
    var v1 bool = (!(!(1 <= 1)))
    var v2 bool = false
    var v3 uint16 = 991
    var v4 uint8 = 0
    var v5 float64 = 1.5
    v5 -= v5
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS3{F0: ("a" + "a"), F1: 1}
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 bool = v2
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    switch -577 {
    case -947:
        v4 = (v4 / ((62) | 1))
    case 32768:
        v25 := GenS1{F0: 1, F1: 42, F2: 0}
        _ = v25
    case -383:
        v2 = false
    default:
        if false {
            if (!v2) {
                v26 := GenS3{F0: "\n\t", F1: 65535}
                v27 := GenS2{F0: (^(v3))}
                for fi28 := 0; fi28 < 1; fi28++ {
                    fmt.Println(v2, v1)
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if (!(65535 <= 30)) {
                            v5 -= 1.5
                            switch fi28 {
                            case 985:
                                fmt.Println(v4)
                            case 743:
                                var v30 string = "\n\t"
                                _ = v30
                            default:
                                v31 := GenS4{F0: "\n\t"}
                                _ = v31
                            }
                            var v32 float32 = 0.1
                            _ = v32
                        } else {
                            fmt.Println(fi29, v5)
                            fmt.Println(v3, v5)
                            v33 := GenS2{F0: 65535}
                            _ = v33
                        }
                        _ = fi29
                    }
                    if ((-9223372036854775808 == 0) && ((!(32767 <= (^(v3)))) && (51 <= 66))) {
                        v34 := GenS5{F0: 713, F1: 1}
                        v4 += 1
                        _ = v34
                    }
                    _ = fi28
                }
                _ = v26
                _ = v27
            } else {
                fmt.Println(v5, v5)
            }
            v35 := []int{-1}
            v36 := GenS6{F0: 2}
            _ = v35
            _ = v36
        }
    }
    v37 := GenS1{F0: 100, F1: 18446744073709551615, F2: 316}
    var v38 bool = ((^(v3)) == v3)
    if (v38 || ("hello" < ("\n\t" + ("hello" + "\n\t")))) {
        v39 := GenS2{F0: v3}
        _ = v39
    } else {
        v1 = v38
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 string = ("foo bar" + "foo bar")
            if (true && (!v1)) {
                var v42 int32 = 2147483647
                _ = v42
            } else {
                var v43 float32 = ((3.14159 * 1.5) + (3.14159 * float32((-2.25 / float64(((0.1 + 1.5) * 1.5))))))
                fmt.Println(v1, v1, v2)
                _ = v43
            }
            var v44 int = 32767
            _ = fi40
            _ = v41
            _ = v44
        }
    }
    v45 := GenS7{F0: "\n\t", F1: ("hello" + ((("a" + "\n\t") + ("hello" + ((("\n\t" + ("a" + ("hello" + ("\n\t" + "hello")))) + "\n\t") + "\n\t"))) + ""))}
    if true {
        if v1 {
            v5 = -2.25
            var v46 int64 = -9223372036854775808
            v47 := []int{2147483647, 9223372036854775807, 570, 4294967295}
            _ = v46
            _ = v47
        } else {
            v48 := []int{32767, -9223372036854775808, 42, -366}
            v1 = (!true)
            _ = v48
        }
    }
    if (-835 != -801) {
        fmt.Println(v4, v1)
    }
    if (!(!(((0.0 + v5) / 3.14159) <= ((0.1 / v5) / v5)))) {
        if (!true) {
            var v49 bool = v38
            switch 2147483647 {
            case 128:
                fmt.Println(v49)
            case 127:
                fmt.Println(v49, v2)
            default:
                fmt.Println(v3, v4)
            }
            _ = v49
        } else {
            var v50 uint = 18446744073709551615
            switch "\n\t" {
            case "foo bar":
                if (1 >= (v4 ^ 42)) {
                    if ((float64((0.0 * (0.1 - 1.5))) != ((((0.1 * v5) * 0.0) - float64(-2.25)) - 0.1)) && v2) {
                        if (v38 && true) {
                            for fi51 := 0; fi51 < 1; fi51++ {
                                var v52 int64 = 0
                                v53 := []int{(fi51 + fi51), 127, (fi51 / (((fi51 | (fi51 / (((fi51 >> (((v50 << (((v50 << (((^(v50))) & 63))) & 63))) & 63))) | 1)))) | 1)), (fi51 % ((fi51) | 1))}
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        } else {
                            fmt.Println(v50, v4, v38)
                        }
                    }
                    v54 := GenS5{F0: 849, F1: uint64(v3)}
                    _ = v54
                } else {
                    v38 = true
                    if true {
                        fmt.Println(v5, v5, v2)
                    }
                }
            case "a":
                var v55 int = -662
                _ = v55
            default:
                if ((29 != 127) || v1) {
                    if false {
                        if (42 == -9223372036854775808) {
                            v56 := []int{42, 4294967295, -9223372036854775808, 903}
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 8; fi57++ {
                            fmt.Println(v38, v3)
                            v58 := []int{fi57, -11, fi57}
                            _ = fi57
                            _ = v58
                        }
                        if v1 {
                            v59 := []int{int(v50), -1}
                            _ = v59
                        }
                    } else {
                        var v60 float32 = 0.1
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 13; fi61++ {
                        v62 := []int{}
                        v2 = false
                        if v1 {
                            v5 += v5
                        }
                        _ = fi61
                        _ = v62
                    }
                }
            }
            if false {
                var v63 uint = 255
                _ = v63
            }
            _ = v50
        }
        var v64 uint16 = (v3 % ((v3) | 1))
        v3 -= v64
        _ = v64
    } else {
        for fi65 := 0; fi65 < 2; fi65++ {
            v5 = v5
            switch "foo bar" {
            case "a":
                var v66 int16 = -746
                _ = v66
            case "foo bar":
                if (v3 != v3) {
                    if ("\n\t" < "foo bar") {
                        var v67 uint32 = 42
                        v68 := GenS5{F0: 18446744073709551615, F1: genFn0()}
                        _ = v67
                        _ = v68
                    }
                } else {
                    for fi69 := 0; fi69 < 3; fi69++ {
                        if true {
                            switch -2147483648 {
                            case 32768:
                                v3 = 65535
                            case 9223372036854775807:
                                v5 = v5
                            }
                            v1 = ((!(v3 != (v3 % (((v3 + (v3 * (v3 << ((255) & 63))))) | 1)))) && v1)
                            if false {
                                fmt.Println(fi69, v5)
                            }
                        }
                        if (591 > 42) {
                            var v70 string = "\n\t"
                            var v71 uint16 = v3
                            v4 -= v4
                            fmt.Println(v3, v5)
                            _ = v70
                            _ = v71
                        }
                        _ = fi69
                    }
                }
            }
            v72 := GenS6{F0: 950}
            _ = fi65
            _ = v72
        }
    }
    var v73 int = -1
    v2 = v38
    switch "a" {
    case "foo bar":
        var v74 uint16 = v3
        _ = v74
    case "a":
        v5 = 3.14159
    default:
        for fi75 := 0; fi75 < 4; fi75++ {
            v76 := []int{}
            _ = fi75
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 8; fi77++ {
        v78 := []int{9223372036854775807, (v73 | 1)}
        _ = fi77
        _ = v78
    }
    for fi79 := 0; fi79 < 14; fi79++ {
        if v38 {
            v80 := []int{}
            var v81 uint = 770
            _ = v80
            _ = v81
        } else {
            var v82 uint32 = 1
            v83 := []int{(fi79 << ((255) & 63)), 778}
            if len(v83) > 0 {
                v83[0] = v73
            }
            _ = v82
            _ = v83
        }
        var v84 int64 = -1
        _ = fi79
        _ = v84
    }
    fmt.Println(v1, v3, v38)
    for fi85 := 0; fi85 < 15; fi85++ {
        if ("" <= ("hello" + ("hello" + "\n\t"))) {
            for fi86 := 0; fi86 < 4; fi86++ {
                var v87 int32 = 0
                v88 := []int{int(v3), 252}
                _ = fi86
                _ = v87
                _ = v88
            }
        }
        _ = fi85
    }
    if (42 == 169) {
        switch "\n\t" {
        case "hello":
            v89 := []int{v73, (v73 ^ (v73 | (v73 & v73)))}
            _ = v89
        case "a":
            v1 = (((685 < 2147483647) || (127 >= (v73 & v73))) || (!((115 <= genFn0()) && true)))
        case "\n\t":
            if true {
                var v90 uint8 = 42
                var v91 uint32 = 2
                for fi92 := 0; fi92 < 8; fi92++ {
                    v93 := []int{(v73 % ((2147483647) | 1)), v73}
                    if len(v93) > 0 {
                        v93[0] = v73
                    }
                    _ = fi92
                    _ = v93
                }
                _ = v90
                _ = v91
            }
        default:
            if true {
                for fi94 := 0; fi94 < 3; fi94++ {
                    v95 := []int{32768, 4294967295}
                    _ = fi94
                    _ = v95
                }
            } else {
                switch "\n\t" {
                case "":
                    if (0.0 <= ((1e10 + float32((v5 * float64(1.5)))) + (3.14159 / (0.0 - (1e10 * 3.14159))))) {
                        for fi96 := 0; fi96 < 16; fi96++ {
                            v97 := GenS6{F0: 1}
                            v73 = v73
                            _ = fi96
                            _ = v97
                        }
                        fmt.Println(v3)
                    } else {
                        v3 += 350
                        if (false || true) {
                            v38 = v38
                            if (1 > genFn0()) {
                                var v98 float64 = v5
                                _ = v98
                            } else {
                                for fi99 := 0; fi99 < 19; fi99++ {
                                    var v100 float64 = -2.25
                                    for fi101 := 0; fi101 < 15; fi101++ {
                                        fmt.Println(v3, v100)
                                        _ = fi101
                                    }
                                    _ = fi99
                                    _ = v100
                                }
                            }
                        } else {
                            var v102 uint32 = 0
                            for fi103 := 0; fi103 < 12; fi103++ {
                                if false {
                                    v73 = v73
                                    var v104 int64 = 1
                                    _ = v104
                                }
                                var v105 uint64 = 458
                                _ = fi103
                                _ = v105
                            }
                            fmt.Println(v1)
                            _ = v102
                        }
                        var v106 uint32 = 344
                        _ = v106
                    }
                default:
                    var v107 int8 = 19
                    _ = v107
                }
                var v108 uint = 253
                switch ("hello" + "hello") {
                case "hello":
                    var v109 float32 = (1e10 * 1.5)
                    _ = v109
                case "a":
                    var v110 int16 = 1000
                    _ = v110
                }
                _ = v108
            }
        }
    } else {
        v4 = (^(v4))
        var v111 uint16 = v3
        switch "\n\t" {
        case "\n\t":
            v112 := GenS6{F0: 167}
            _ = v112
        case "foo bar":
            var v113 float32 = (-2.25 + -2.25)
            _ = v113
        case "a":
            var v114 int64 = -9223372036854775808
            _ = v114
        }
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
    _ = v73
}
