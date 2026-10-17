package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int8; F2 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int; F1 uint; F2 int16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 uint8 }
type GenS15 struct { F0 uint; F1 uint8 }

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 float64 = -2.25
    var v3 int32 = 42
    for fi4 := 0; fi4 < 10; fi4++ {
        switch fi4 {
        case 2147483647:
            var v5 int8 = -111
            _ = v5
        }
        v6 := []int{}
        _ = fi4
        _ = v6
    }
    var v7 int = 2147483647
    var v8 int8 = 37
    var v9 float64 = v2
    v7 += v7
    v10 := GenS1{F0: v7, F1: 1, F2: 255}
    if (!(0 > 760)) {
        var v11 uint32 = 65535
        v12 := GenS1{F0: 128, F1: 2, F2: 128}
        v13 := GenS1{F0: (^(v7)), F1: 0, F2: 128}
        _ = v11
        _ = v12
        _ = v13
    } else {
        switch v7 {
        case -1:
            if (!true) {
                v7 -= (v7 | v7)
            } else {
                if false {
                    if (65535 == 42) {
                        fmt.Println(v9)
                        if true {
                            if true {
                                v1 = ((("hello" + ("hello" + "a")) + v1) + v1)
                                v14 := []int{v7, 127}
                                _ = v14
                            }
                            var v15 bool = false
                            _ = v15
                        } else {
                            switch v1 {
                            case "foo bar":
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 bool = true
                                    switch v1 {
                                    case "\n\t":
                                        for fi18 := 0; fi18 < 16; fi18++ {
                                            if false {
                                                v3 = -2147483648
                                            }
                                            switch v1 {
                                            case "foo bar":
                                                var v19 bool = (v17 && (!v17))
                                                _ = v19
                                            }
                                            switch (^(fi16)) {
                                            case 4294967295:
                                                switch 876 {
                                                case 1:
                                                    v20 := GenS2{F0: 127, F1: (v8 * v8), F2: 65535}
                                                    _ = v20
                                                case 42:
                                                    fmt.Println(v9)
                                                }
                                            case 2147483647:
                                                switch "hello" {
                                                case "":
                                                    var v21 int16 = 42
                                                    _ = v21
                                                case "foo bar":
                                                    v22 := []int{(fi16 % ((fi16) | 1))}
                                                    _ = v22
                                                case "a":
                                                    v1 = "hello"
                                                default:
                                                    v17 = v17
                                                }
                                            }
                                            v23 := GenS3{F0: 255}
                                            v24 := []int{(v7 << ((0) & 63)), v7, (v7 + fi18), (v7 / ((int(v8)) | 1))}
                                            _ = fi18
                                            _ = v23
                                            _ = v24
                                        }
                                    case "hello":
                                        for fi25 := 0; fi25 < 15; fi25++ {
                                            v26 := GenS4{F0: -90}
                                            if (2 <= 630) {
                                                fmt.Println(fi25)
                                                var v27 uint16 = 255
                                                fmt.Println(v17)
                                                _ = v27
                                            } else {
                                                if (v9 == (v9 - 0.0)) {
                                                    v28 := []int{v7}
                                                    _ = v28
                                                } else {
                                                    v29 := GenS1{F0: (v7 << ((368) & 63)), F1: 0, F2: 1}
                                                    switch "" {
                                                    case "foo bar":
                                                        v9 = v2
                                                    case "a":
                                                        var v30 uint16 = 32767
                                                        _ = v30
                                                    default:
                                                        v7 = (^(v7))
                                                    }
                                                    _ = v29
                                                }
                                                v8 += v8
                                            }
                                            v17 = false
                                            v7 = fi16
                                            if false {
                                                v31 := GenS2{F0: (v8 | v8), F1: (v8 | v8), F2: 255}
                                                if v17 {
                                                    v32 := GenS5{F0: 479, F1: 1, F2: 1000}
                                                    _ = v32
                                                } else {
                                                    v33 := GenS6{F0: 0}
                                                    switch (v1 + v1) {
                                                    case "foo bar":
                                                        var v34 int16 = -137
                                                        _ = v34
                                                    case "hello":
                                                        fmt.Println(v9)
                                                    }
                                                    _ = v33
                                                }
                                                var v35 int16 = -37
                                                _ = v31
                                                _ = v35
                                            }
                                            _ = fi25
                                            _ = v26
                                        }
                                    case "a":
                                        fmt.Println(v3)
                                    default:
                                        for fi36 := 0; fi36 < 10; fi36++ {
                                            if ((v3 + v3) < (-(v3))) {
                                                fmt.Println(v8, v7, v9)
                                                if v17 {
                                                    if (0.1 >= v9) {
                                                        var v37 uint64 = 2
                                                        v38 := []int{(fi36 & (fi36 + -1))}
                                                        v9 = v9
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        v39 := GenS3{F0: 221}
                                                        _ = v39
                                                    }
                                                    v40 := GenS7{F0: 9223372036854775807, F1: (^(v8)), F2: v1}
                                                    _ = v40
                                                } else {
                                                    v41 := []int{v7}
                                                    var v42 uint = 2
                                                    var v43 uint16 = 255
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            } else {
                                                v3 = (^(v3))
                                                v17 = v17
                                            }
                                            fmt.Println(fi36, v1)
                                            v2 = 0.0
                                            _ = fi36
                                        }
                                    }
                                    v44 := GenS8{F0: 0}
                                    _ = fi16
                                    _ = v17
                                    _ = v44
                                }
                            case "":
                                fmt.Println(v8, v1, v7)
                            default:
                                for fi45 := 0; fi45 < 15; fi45++ {
                                    if true {
                                        var v46 int = fi45
                                        if true {
                                            fmt.Println(v2)
                                        }
                                        _ = v46
                                    } else {
                                        v2 = float64(3.14159)
                                    }
                                    _ = fi45
                                }
                            }
                            v47 := []int{}
                            _ = v47
                        }
                        var v48 int16 = 163
                        _ = v48
                    }
                    fmt.Println(v8, v9)
                } else {
                    for fi49 := 0; fi49 < 1; fi49++ {
                        fmt.Println(v3, v7, v8)
                        if (fi49 == (fi49 + fi49)) {
                            var v50 uint64 = 4294967295
                            var v51 int8 = (v8 - 105)
                            _ = v50
                            _ = v51
                        }
                        _ = fi49
                    }
                    v52 := GenS9{F0: -1}
                    if (42 <= 2) {
                        for fi53 := 0; fi53 < 18; fi53++ {
                            switch "" {
                            case "":
                                fmt.Println(v7, v1, fi53)
                            case "a":
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    if true {
                                        v2 -= ((v2 - v2) * v2)
                                    }
                                    var v55 uint16 = 958
                                    _ = fi54
                                    _ = v55
                                }
                            default:
                                v56 := []int{32768, (-(fi53)), (v7 + v7), 1000}
                                _ = v56
                            }
                            var v57 int32 = (-(v3))
                            fmt.Println(v9)
                            _ = fi53
                            _ = v57
                        }
                    } else {
                        if (((^(v7)) < (v7 / ((2147483647) | 1))) && (!((v7 ^ (v7 & (^(v7)))) != v7))) {
                            switch (-(v7)) {
                            case 9223372036854775807:
                                var v58 int16 = 1000
                                _ = v58
                            case -534:
                                switch v1 {
                                case "\n\t":
                                    var v59 uint8 = 42
                                    _ = v59
                                }
                            default:
                                var v60 int32 = -867
                                _ = v60
                            }
                            fmt.Println(v1, v3, v7)
                        } else {
                            var v61 uint16 = 453
                            var v62 uint16 = (v61 | (v61 - (v61 ^ v61)))
                            _ = v61
                            _ = v62
                        }
                        fmt.Println(v8, v1, v8)
                    }
                    _ = v52
                }
                fmt.Println(v7, v2)
                if false {
                    fmt.Println(v8)
                } else {
                    var v63 uint32 = 953
                    if true {
                        v64 := GenS10{F0: v63}
                        v65 := []int{}
                        v1 = v1
                        _ = v64
                        _ = v65
                    } else {
                        var v66 uint = 0
                        switch ((v1 + v1) + "hello") {
                        case "":
                            for fi67 := 0; fi67 < 18; fi67++ {
                                var v68 int32 = -2147483648
                                if (!false) {
                                    var v69 uint32 = v63
                                    v70 := GenS11{F0: (fi67 % ((int(v69)) | 1)), F1: 282, F2: int16(v8)}
                                    _ = v69
                                    _ = v70
                                } else {
                                    v71 := GenS9{F0: 0}
                                    if true {
                                        fmt.Println(v3, v1)
                                    }
                                    switch (v1 + v1) {
                                    case "a":
                                        if true {
                                            for fi72 := 0; fi72 < 7; fi72++ {
                                                if (((true && true) || (!false)) || (223 <= 153)) {
                                                    switch v1 {
                                                    case "":
                                                        v73 := []int{v7, (fi67 + 0), 32768, fi67}
                                                        _ = v73
                                                    case "a":
                                                        v74 := GenS12{F0: 1}
                                                        _ = v74
                                                    default:
                                                        v75 := GenS13{F0: 65535}
                                                        _ = v75
                                                    }
                                                    var v76 uint64 = 42
                                                    var v77 int64 = -9223372036854775808
                                                    _ = v76
                                                    _ = v77
                                                }
                                                v1 = v1
                                                _ = fi72
                                            }
                                        }
                                    case "hello":
                                        fmt.Println(v3)
                                    case "":
                                        var v78 int = fi67
                                        _ = v78
                                    }
                                    _ = v71
                                }
                                for fi79 := 0; fi79 < 19; fi79++ {
                                    if true {
                                        var v80 uint64 = 1
                                        if true {
                                            var v81 int = (fi67 >> ((uint(v63)) & 63))
                                            if (false && true) {
                                                fmt.Println(v3)
                                                v82 := []int{fi67, (fi67 * fi67), fi79}
                                                _ = v82
                                            }
                                            _ = v81
                                        }
                                        _ = v80
                                    }
                                    _ = fi79
                                }
                                _ = fi67
                                _ = v68
                            }
                        default:
                            var v83 int = 1
                            _ = v83
                        }
                        for fi84 := 0; fi84 < 16; fi84++ {
                            switch (v7 & 828) {
                            case 32767:
                                var v85 float64 = (v2 + v9)
                                _ = v85
                            case 0:
                                if (42 <= 4294967295) {
                                    for fi86 := 0; fi86 < 1; fi86++ {
                                        v63 += v63
                                        _ = fi86
                                    }
                                    var v87 uint16 = 65535
                                    fmt.Println(v8)
                                    _ = v87
                                }
                            default:
                                if ((2 >= v66) || false) {
                                    v63 -= v63
                                }
                            }
                            for fi88 := 0; fi88 < 11; fi88++ {
                                fmt.Println(v66)
                                v89 := GenS6{F0: -242}
                                v2 = (v2 - 3.14159)
                                if false {
                                    v90 := GenS10{F0: (v63 % ((v63) | 1))}
                                    var v91 uint32 = 402
                                    v2 = 0.0
                                    _ = v90
                                    _ = v91
                                } else {
                                    var v92 int64 = -9223372036854775808
                                    switch v1 {
                                    case "":
                                        fmt.Println(v3, v3)
                                    case "a":
                                        var v93 int64 = (v92 >> ((v66) & 63))
                                        _ = v93
                                    case "hello":
                                        var v94 float64 = v9
                                        _ = v94
                                    default:
                                        if false {
                                            v63 -= 684
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                switch v7 {
                                                case -2147483648:
                                                    var v95 uint64 = 18446744073709551615
                                                    _ = v95
                                                case 32767:
                                                    v96 := []int{(fi88 % ((4294967295) | 1)), (v7 - 128), 42}
                                                    _ = v96
                                                case -393:
                                                    var v97 uint16 = 1
                                                    _ = v97
                                                default:
                                                    fmt.Println(v2, v7)
                                                    v98 := GenS14{F0: 0, F1: 91}
                                                    _ = v98
                                                }
                                            case "hello":
                                                v99 := GenS6{F0: 864}
                                                _ = v99
                                            }
                                        }
                                    }
                                    v1 = ("foo bar" + ((v1 + ("" + (v1 + "foo bar"))) + "foo bar"))
                                    v100 := GenS2{F0: 127, F1: 90, F2: 724}
                                    _ = v100
                                    _ = v92
                                }
                                _ = fi88
                                _ = v89
                            }
                            _ = fi84
                        }
                        _ = v66
                    }
                    for fi101 := 0; fi101 < 15; fi101++ {
                        var v102 uint16 = 376
                        v103 := []int{(v7 & fi101), (fi101 >> ((uint(v102)) & 63))}
                        v104 := GenS1{F0: fi101, F1: 4294967295, F2: 1}
                        _ = fi101
                        _ = v102
                        _ = v103
                        _ = v104
                    }
                    _ = v63
                }
            }
        case 2147483647:
            var v105 uint = 255
            _ = v105
        default:
            var v106 int = 32767
            _ = v106
        }
        for fi107 := 0; fi107 < 3; fi107++ {
            v108 := []int{v7, v7}
            _ = fi107
            _ = v108
        }
    }
    if (!false) {
        v109 := GenS15{F0: 42, F1: 136}
        _ = v109
    }
    v110 := []int{}
    v8 += -7
    for fi111 := 0; fi111 < 11; fi111++ {
        var v112 uint8 = 103
        if len(v110) >= 2 {
            v110 = v110[1:len(v110)]
        }
        for fi113 := 0; fi113 < 15; fi113++ {
            var v114 string = v1
            v110 = append(v110, (fi111 / ((127) | 1)))
            _ = fi113
            _ = v114
        }
        _ = fi111
        _ = v112
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v110
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
