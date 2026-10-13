package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string; F1 int32; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 bool; F2 int8 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    var v3 bool = true
    var v4 int8 = -53
    var v5 int = 32768
    var v6 uint32 = 2147483648
    switch -2147483648 {
    case 127:
        v7 := GenS1{F0: 42, F1: 106}
        _ = v7
    }
    var v8 int16 = 197
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: v6, F1: 127}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint64 = 42
        v4 += v4
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v5 -= (v5 + 242)
        _ = fi13
    }
    v14 := GenS2{F0: 131, F1: (v4 % (((v4 % ((v4) | 1))) | 1))}
    if v3 {
        switch v5 {
        case 42:
            if false {
                var v15 float64 = (0.1 * (-2.25 - 1e10))
                for fi16 := 0; fi16 < 1; fi16++ {
                    switch (v5 - -586) {
                    case 127:
                        v3 = ((18446744073709551615 == 2) || (!v3))
                    }
                    v17 := []int{(^(fi16)), 0, (v5 / ((-9223372036854775808) | 1))}
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        }
        var v18 float32 = 0.0
        _ = v18
    }
    v19 := GenS2{F0: (-(v8)), F1: 0}
    if (!((^(v2)) >= v1)) {
        var v20 int = int(v1)
        _ = v20
    }
    v3 = v3
    if (-2147483648 != v2) {
        if false {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 string = "foo bar"
                if (v3 || v3) {
                    v23 := GenS2{F0: -32768, F1: v4}
                    v24 := GenS2{F0: (v8 / ((int16(v4)) | 1)), F1: 101}
                    var v25 int16 = 32767
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint64 = 284
                    v27 := []int{fi21, (-(v5)), fi21, v5}
                    v1 -= 42
                    _ = v26
                    _ = v27
                }
                var v28 float64 = (3.14159 + (((1e10 - 0.0) - 0.1) + 1.5))
                _ = fi21
                _ = v22
                _ = v28
            }
            if (4294967295 < 0) {
                v29 := []int{212, (v5 ^ v5), (v5 ^ (-(v5))), 5}
                _ = v29
            } else {
                var v30 int = (v5 << ((0) & 63))
                v8 = -1
                if (v6 == v6) {
                    if v3 {
                        v4 += -83
                        var v31 int32 = 1
                        _ = v31
                    }
                }
                _ = v30
            }
        } else {
            for fi32 := 0; fi32 < 8; fi32++ {
                if (v3 && v3) {
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v34 := []int{-944, int(v4), int(v4)}
                        switch (("" + "hello") + "foo bar") {
                        case "\n\t":
                            v4 += (v4 << ((504) & 63))
                        case "foo bar":
                            v4 = 48
                        default:
                            var v35 int16 = v8
                            _ = v35
                        }
                        _ = fi33
                        _ = v34
                    }
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 uint8 = 255
                        if v3 {
                            v38 := GenS2{F0: -32768, F1: (v4 ^ (v4 - 19))}
                            _ = v38
                        }
                        if (70 < -9223372036854775808) {
                            if false {
                                var v39 int32 = (v2 | (v1 / (((v2 % ((int32(v37)) | 1))) | 1)))
                                v4 = (v4 % ((v4) | 1))
                                var v40 uint8 = (v37 + v37)
                                _ = v39
                                _ = v40
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                    switch ("foo bar" + ("hello" + "a")) {
                    case "hello":
                        v6 = 42
                    case "\n\t":
                        v3 = true
                    case "foo bar":
                        switch -895 {
                        case 82:
                            v41 := []int{-314, v5, fi32}
                            _ = v41
                        case 9223372036854775807:
                            switch (("\n\t" + ((("\n\t" + ("" + "\n\t")) + "") + "\n\t")) + "") {
                            case "":
                                fmt.Println(v6)
                            case "\n\t":
                                var v42 int64 = -9223372036854775808
                                _ = v42
                            case "hello":
                                for fi43 := 0; fi43 < 18; fi43++ {
                                    fmt.Println(v8, v1)
                                    switch (("\n\t" + "hello") + "") {
                                    case "\n\t":
                                        v3 = (42 == 144)
                                    default:
                                        var v44 uint16 = 771
                                        _ = v44
                                    }
                                    _ = fi43
                                }
                            default:
                                var v45 string = ""
                                _ = v45
                            }
                        }
                    }
                } else {
                    var v46 uint8 = 175
                    var v47 int = (-(v5))
                    _ = v46
                    _ = v47
                }
                if v3 {
                    v6 = (v6 + (v6 - (v6 >> ((18446744073709551615) & 63))))
                } else {
                    var v48 uint32 = (^(v6))
                    var v49 int64 = 9223372036854775807
                    _ = v48
                    _ = v49
                }
                v4 -= v4
                _ = fi32
            }
        }
        switch v5 {
        case 128:
            v3 = false
        case -2147483648:
            v50 := GenS1{F0: 930, F1: 255}
            _ = v50
        case -152:
            var v51 int16 = v8
            _ = v51
        default:
            fmt.Println(v6, v4, v3)
        }
        v8 += (v8 ^ v8)
    }
    switch 2147483647 {
    case 17:
        if (42 != 1) {
            for fi52 := 0; fi52 < 14; fi52++ {
                var v53 uint8 = 127
                v53 = v53
                for fi54 := 0; fi54 < 10; fi54++ {
                    if v3 {
                        fmt.Println(v5, v53, v3)
                        v8 -= v8
                    }
                    var v55 uint64 = 669
                    _ = fi54
                    _ = v55
                }
                _ = fi52
                _ = v53
            }
        }
    case -662:
        var v56 string = ""
        _ = v56
    case 127:
        if false {
            var v57 float64 = (-2.25 / (3.14159 / -2.25))
            v3 = false
            _ = v57
        } else {
            switch (v5 / ((v5) | 1)) {
            case 128:
                v58 := GenS3{F0: "foo bar", F1: -2147483648, F2: (-(v8))}
                _ = v58
            case 600:
                switch ("\n\t" + ((("\n\t" + "\n\t") + "a") + "foo bar")) {
                case "\n\t":
                    if ((true && v3) || (((1.5 / 1.0) < 0.0) || (!v3))) {
                        for fi59 := 0; fi59 < 8; fi59++ {
                            if v3 {
                                for fi60 := 0; fi60 < 4; fi60++ {
                                    if (!(float32(1.5) != 0.1)) {
                                        var v61 string = "a"
                                        v62 := GenS3{F0: "", F1: v1, F2: -32768}
                                        v8 += (v8 & v8)
                                        if v3 {
                                            if (((v61 + v61) > "hello") || ((v4 & -128) >= (v4 / ((v4) | 1)))) {
                                                if v3 {
                                                    var v63 uint16 = 2
                                                    var v64 uint32 = (v6 * (^(v6)))
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 int32 = int32(v4)
                                                _ = v65
                                            }
                                        } else {
                                            v4 = (v4 & (v4 / ((v4) | 1)))
                                            v66 := []int{v5, (v5 >> ((65535) & 63))}
                                            switch v61 {
                                            case "hello":
                                                fmt.Println(v2)
                                            case "a":
                                                switch int(v4) {
                                                case 128:
                                                    v8 = v8
                                                default:
                                                    var v67 uint16 = 255
                                                    _ = v67
                                                }
                                            default:
                                                switch v61 {
                                                case "\n\t":
                                                    fmt.Println(v4)
                                                }
                                            }
                                            _ = v66
                                        }
                                        _ = v61
                                        _ = v62
                                    }
                                    v68 := []int{}
                                    _ = fi60
                                    _ = v68
                                }
                            }
                            _ = fi59
                        }
                        fmt.Println(v8, v4, v5)
                        fmt.Println(v6, v4, v1)
                    } else {
                        if (false && false) {
                            for fi69 := 0; fi69 < 3; fi69++ {
                                v70 := GenS4{F0: (^(v1))}
                                fmt.Println(v2)
                                _ = fi69
                                _ = v70
                            }
                        } else {
                            v6 += (v6 * (v6 / ((2) | 1)))
                            v3 = (!(v1 != (v1 + (^(v2)))))
                            if (((!false) || (false || (v8 >= v8))) || (v3 && v3)) {
                                var v71 uint64 = 79
                                _ = v71
                            }
                        }
                        v4 = (v4 / ((int8(v5)) | 1))
                    }
                case "":
                    v1 = (-(v1))
                case "foo bar":
                    v72 := GenS5{F0: (v5 % ((v5) | 1)), F1: (!v3), F2: 2}
                    _ = v72
                }
            default:
                v1 = 2147483647
            }
            for fi73 := 0; fi73 < 15; fi73++ {
                v4 += (v4 >> ((2) & 63))
                switch 4294967295 {
                case 128:
                    var v74 float64 = (1.5 + (3.14159 * 0.1))
                    _ = v74
                case 32767:
                    for fi75 := 0; fi75 < 13; fi75++ {
                        if v3 {
                            v5 += fi75
                            v6 = (v6 * 1)
                        }
                        v76 := GenS6{F0: v8}
                        _ = fi75
                        _ = v76
                    }
                case 9223372036854775807:
                    switch "foo bar" {
                    case "":
                        v5 = v5
                    default:
                        for fi77 := 0; fi77 < 1; fi77++ {
                            if (9223372036854775807 == -1) {
                                switch "a" {
                                case "":
                                    var v78 uint16 = 42
                                    _ = v78
                                case "a":
                                    v1 = v2
                                default:
                                    v8 = -32768
                                }
                                fmt.Println(v8, v8)
                                v79 := GenS6{F0: (v8 << ((0) & 63))}
                                _ = v79
                            }
                            var v80 uint16 = 2
                            var v81 string = "hello"
                            _ = fi77
                            _ = v80
                            _ = v81
                        }
                    }
                }
                _ = fi73
            }
            var v82 string = ""
            _ = v82
        }
    default:
        v83 := GenS6{F0: (v8 >> ((42) & 63))}
        _ = v83
    }
    v5 -= v5
    for fi84 := 0; fi84 < 9; fi84++ {
        if (v2 > v2) {
            v2 += (v2 >> ((835) & 63))
            var v85 bool = (1 > 4294967295)
            _ = v85
        } else {
            for fi86 := 0; fi86 < 16; fi86++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v8, v4, v8)
                case "a":
                    v87 := []int{(fi84 / ((fi84) | 1)), (^(v5))}
                    _ = v87
                case "foo bar":
                    if v3 {
                        v5 -= fi86
                        var v88 uint = 0
                        _ = v88
                    }
                default:
                    v1 += (v2 % (((v2 * int32(v5))) | 1))
                }
                if (v3 && (v5 > 42)) {
                    fmt.Println(v4, v6)
                    fmt.Println(v6, v2)
                    var v89 string = "hello"
                    _ = v89
                } else {
                    var v90 uint8 = 42
                    switch v5 {
                    case -970:
                        var v91 float32 = 1e10
                        _ = v91
                    case 42:
                        v92 := []int{}
                        _ = v92
                    case -1:
                        var v93 uint16 = uint16(v6)
                        _ = v93
                    }
                    var v94 uint16 = 65535
                    _ = v90
                    _ = v94
                }
                if true {
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v95 := []int{fi86, (^(v5))}
                        _ = v95
                    }
                    v96 := []int{fi84, -792, 2147483647}
                    if (!false) {
                        fmt.Println(v1)
                        if (v3 || v3) {
                            v1 -= (v1 ^ (-(v2)))
                            v4 = (v4 % ((v4) | 1))
                        }
                        v8 += (v8 % (((v8 * int16(v5))) | 1))
                    }
                    v4 -= 42
                    _ = v96
                } else {
                    if v3 {
                        v5 = -1
                    } else {
                        v97 := []int{(fi84 / ((fi84) | 1))}
                        _ = v97
                    }
                }
                _ = fi86
            }
            v98 := []int{2147483647, (v5 >> ((42) & 63))}
            _ = v98
        }
        v99 := []int{-280, v5, (^(fi84))}
        _ = fi84
        _ = v99
    }
    v100 := []int{1, (^(v5)), v5, v5}
    var v101 uint = 42
    v101 -= (^(v101))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v100
    _ = v101
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
