package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint16; F1 uint; F2 bool }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 uint8) bool {
    v4 := GenS1{F0: 0, F1: 2147483647}
    v5 := GenS1{F0: -2147483648, F1: -67}
    v6 := []int{128, -634, int(p0)}
    if false {
        p0 -= (p0 - (^(p0)))
        switch "foo bar" {
        case "foo bar":
            p0 = 255
        }
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 float64 = 3.14159
                p0 = p0
                switch ("foo bar" + "foo bar") {
                case "a":
                    p0 -= (p0 ^ p0)
                case "hello":
                    var v10 uint16 = 42
                    _ = v10
                default:
                    var v11 int32 = -200
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if (1 <= 2) {
                var v13 uint64 = 4294967295
                switch ("hello" + "hello") {
                case "foo bar":
                    if (!(48 != -1)) {
                        switch "\n\t" {
                        case "":
                            var v14 int16 = -1
                            _ = v14
                        case "\n\t":
                            p0 = (p0 / ((p0) | 1))
                        }
                        if ("hello" >= "\n\t") {
                            if true {
                                v15 := GenS1{F0: -1, F1: (fi12 / (((fi12 >> ((0) & 63))) | 1))}
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                                _ = v15
                            } else {
                                v13 = v13
                                p0 += uint8(v13)
                            }
                            var v16 int8 = 0
                            var v17 uint = 255
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint64 = (^(v13))
                            _ = v18
                        }
                        var v19 string = "\n\t"
                        _ = v19
                    } else {
                        switch (("hello" + "hello") + ("a" + "")) {
                        case "hello":
                            if (1 > 42) {
                                var v20 int = fi12
                                v21 := GenS1{F0: -1, F1: v20}
                                _ = v20
                                _ = v21
                            } else {
                                p0 = (p0 & (p0 * p0))
                                v13 = v13
                                v22 := GenS1{F0: 2147483647, F1: fi12}
                                _ = v22
                            }
                        default:
                            v13 += v13
                        }
                    }
                case "":
                    v13 = (v13 % (((v13 | 4294967295)) | 1))
                case "\n\t":
                    fmt.Println(v13, fi12, v13)
                default:
                    var v23 int32 = -2147483648
                    _ = v23
                }
                if len(v6) > 0 {
                    v6[0] = (fi12 | 32767)
                }
                _ = v13
            }
            _ = fi12
        }
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int64, p1 uint64) uint16 {
    v24 := []int{4294967295, 770}
    _ = len(v24)
    _ = p0
    _ = p1
    _ = v24
    return 65535
}

func main() {
    var v1 bool = (32767 >= 255)
    v2 := []int{202, 863, 598, 128}
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    v25 := v3[""]
    delete(v3, ("" + "hello"))
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 bool = (v25 == (v25 ^ (v25 & v25)))
        var v28 uint8 = 132
        for fi29 := 0; fi29 < 16; fi29++ {
            v27 = v27
            _ = fi29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    var v30 uint64 = 42
    switch 0 {
    case 0:
        _ = len(v2)
    default:
        if false {
            var v31 int64 = -611
            if (!(65535 < 65535)) {
                switch 1 {
                case 2147483647:
                    fmt.Println(v1, v1)
                case 791:
                    fmt.Println(v31)
                case 0:
                    v31 -= 9223372036854775807
                }
                v32 := GenS1{F0: -1, F1: 1}
                v3["a"] = 2147483647
                _ = v32
            } else {
                fmt.Println(v30)
                var v33 uint = 18446744073709551615
                for fi34 := 0; fi34 < 10; fi34++ {
                    if v1 {
                        var v35 bool = v1
                        switch int(v33) {
                        case 728:
                            if false {
                                var v36 bool = (724 <= 18446744073709551615)
                                var v37 int16 = -1
                                var v38 bool = (0 <= -2147483648)
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                v35 = v35
                            }
                        default:
                            if false {
                                switch -121 {
                                case 4294967295:
                                    v35 = false
                                case -191:
                                    v33 += v33
                                default:
                                    v2 = append(v2, fi34)
                                }
                                for fi39 := 0; fi39 < 11; fi39++ {
                                    fmt.Println(v33, v31)
                                    delete(v3, "foo bar")
                                    _ = fi39
                                }
                                var v40 uint32 = v25
                                _ = v40
                            }
                        }
                        _ = v35
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi34 * (fi34 >> ((v33) & 63)))
                    }
                    _ = fi34
                }
                _ = v33
            }
            v2 = append(v2, 127)
            _ = v31
        } else {
            var v41 int32 = -2147483648
            switch "a" {
            case "\n\t":
                if (!(!true)) {
                    switch "" {
                    case "\n\t":
                        if v1 {
                            v30 = (v30 ^ (v30 % (((v30 >> ((1) & 63))) | 1)))
                            var v42 uint8 = 127
                            _ = v42
                        }
                    case "hello":
                        fmt.Println(v41, v1, v25)
                    case "a":
                        if (v1 || v1) {
                            switch -2147483648 {
                            case 0:
                                var v43 uint16 = 255
                                _ = v43
                            case 42:
                                if ((v1 && v1) && ((!false) || v1)) {
                                    v44 := GenS2{F0: 174}
                                    v41 = v41
                                    _ = v44
                                }
                            case -1:
                                v3["foo bar"] = v25
                            }
                            v45 := GenS3{F0: 0, F1: 0}
                            _ = v45
                        } else {
                            if (v1 || false) {
                                fmt.Println(v41, v1, v25)
                                v25 = (v25 | v25)
                                fmt.Println(v25, v25, v41)
                            } else {
                                if true {
                                    v46 := GenS3{F0: 1, F1: 300}
                                    _ = v46
                                } else {
                                    fmt.Println(v25)
                                    if ((v30 ^ 1) == 2) {
                                        var v47 uint64 = v30
                                        v47 -= v30
                                        _ = v47
                                    } else {
                                        v48 := v3[""]
                                        var v49 uint8 = 43
                                        v30 -= 42
                                        _ = v48
                                        _ = v49
                                    }
                                    switch "hello" {
                                    case "":
                                        var v50 int8 = -128
                                        _ = v50
                                    case "a":
                                        delete(v3, "a")
                                    case "foo bar":
                                        v51 := GenS3{F0: 4294967295, F1: int64(v41)}
                                        _ = v51
                                    default:
                                        v25 -= (^(v25))
                                    }
                                }
                                for fi52 := 0; fi52 < 17; fi52++ {
                                    v30 -= 4294967295
                                    v53 := GenS1{F0: 2147483647, F1: (fi52 - fi52)}
                                    var v54 int = fi52
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                            var v55 int8 = 0
                            _ = v55
                        }
                    }
                    if v1 {
                        v41 += -1
                    } else {
                        fmt.Println(v41, v1)
                    }
                }
            case "a":
                v30 = (v30 >> ((18446744073709551615) & 63))
            case "":
                if (!true) {
                    if ((v25 >> ((18446744073709551615) & 63)) < (^(v25))) {
                        fmt.Println(v1)
                        switch (("a" + "foo bar") + "hello") {
                        case "\n\t":
                            v1 = (v1 && true)
                        default:
                            v1 = (int(v41) >= 2147483647)
                        }
                    }
                }
            }
            var v56 uint = 0
            _ = v41
            _ = v56
        }
    }
    var v57 int64 = -9223372036854775808
    if false {
        if len(v2) > 0 {
            v2[0] = -812
        }
        v58 := v3[(((("a" + "") + ((("hello" + "foo bar") + "foo bar") + "foo bar")) + "foo bar") + "hello")]
        _ = v58
    } else {
        switch 482 {
        case 128:
            for fi59 := 0; fi59 < 2; fi59++ {
                if v1 {
                    delete(v3, "a")
                    if (0 <= 42) {
                        var v60 int = fi59
                        v25 = v25
                        v61 := GenS2{F0: 2}
                        _ = v60
                        _ = v61
                    } else {
                        var v62 int = (fi59 + fi59)
                        var v63 int64 = v57
                        _ = v62
                        _ = v63
                    }
                    if (-943 == fi59) {
                        for fi64 := 0; fi64 < 14; fi64++ {
                            if (!v1) {
                                v65 := GenS3{F0: 1, F1: (v57 ^ int64(v30))}
                                v25 += (v25 ^ v25)
                                switch "hello" {
                                case "a":
                                    v66 := GenS1{F0: 42, F1: -344}
                                    _ = v66
                                default:
                                    var v67 int16 = -1
                                    _ = v67
                                }
                                _ = v65
                            }
                            _ = fi64
                        }
                        v3[(("" + "\n\t") + "hello")] = (^(v25))
                        _ = len(v3)
                    }
                }
                if true {
                    switch 32767 {
                    case 127:
                        if (!(3.14159 != 1e10)) {
                            v3["\n\t"] = (v25 | v25)
                        } else {
                            v25 += v25
                        }
                    case 0:
                        var v68 int64 = v57
                        _ = v68
                    case 128:
                        v1 = (39 != 127)
                    default:
                        for fi69 := 0; fi69 < 5; fi69++ {
                            fmt.Println(v1, v25)
                            v1 = v1
                            _ = fi69
                        }
                    }
                    v70 := v3[""]
                    _ = v70
                }
                var v71 bool = v1
                _ = fi59
                _ = v71
            }
        case 4294967295:
            switch -1 {
            case 319:
                for fi72 := 0; fi72 < 10; fi72++ {
                    v73 := GenS2{F0: 77}
                    if true {
                        fmt.Println(v57, fi72, v57)
                    }
                    v25 = 830
                    _ = fi72
                    _ = v73
                }
            case 32768:
                var v74 int16 = 0
                _ = v74
            case 0:
                var v75 uint = 42
                _ = v75
            default:
                v25 = v25
            }
        }
    }
    if false {
        v30 = (v30 + (v30 * (v30 ^ 42)))
        v76 := GenS3{F0: 18446744073709551615, F1: (v57 & v57)}
        fmt.Println(v57, v30)
        _ = v76
    } else {
        fmt.Println(v25, v25, v30)
    }
    fmt.Println(v1)
    for fi77 := 0; fi77 < 10; fi77++ {
        if len(v2) > 0 {
            v2[0] = (fi77 | (fi77 << ((4294967295) & 63)))
        }
        _ = fi77
    }
    switch -306 {
    case 32768:
        switch 32767 {
        case 32767:
            if false {
                v1 = (true && (!v1))
            } else {
                var v78 int64 = v57
                switch 4294967295 {
                case 9223372036854775807:
                    v79 := GenS4{F0: 1000, F1: ("hello" + "")}
                    _ = v79
                case -918:
                    v25 += 2147483647
                case 151:
                    if v1 {
                        v3[""] = v25
                        switch "hello" {
                        case "":
                            if (42 == 239) {
                                v80 := GenS4{F0: 42, F1: "hello"}
                                fmt.Println(v1)
                                fmt.Println(v25, v57)
                                _ = v80
                            } else {
                                for fi81 := 0; fi81 < 9; fi81++ {
                                    var v82 float32 = float32((1e10 - 0.1))
                                    _ = fi81
                                    _ = v82
                                }
                            }
                        case "a":
                            for fi83 := 0; fi83 < 15; fi83++ {
                                var v84 uint32 = 732
                                _ = fi83
                                _ = v84
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                switch 127 {
                                case 128:
                                    v30 = v30
                                case 1:
                                    v25 -= v25
                                case 2147483647:
                                    var v85 uint = 715
                                    _ = v85
                                }
                            default:
                                for fi86 := 0; fi86 < 13; fi86++ {
                                    var v87 uint32 = v25
                                    var v88 uint32 = v87
                                    v89 := GenS1{F0: -71, F1: fi86}
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                }
                            }
                        }
                    } else {
                        if true {
                            switch "foo bar" {
                            case "a":
                                var v90 uint32 = v25
                                _ = v90
                            case "hello":
                                var v91 int64 = v57
                                _ = v91
                            case "\n\t":
                                if v1 {
                                    if len(v2) > 0 {
                                        v2[0] = 19
                                    }
                                    var v92 uint = 18446744073709551615
                                    _ = v92
                                }
                            default:
                                if (!true) {
                                    for fi93 := 0; fi93 < 11; fi93++ {
                                        v94 := GenS3{F0: 0, F1: (v78 ^ v57)}
                                        v95 := v3["\n\t"]
                                        _ = fi93
                                        _ = v94
                                        _ = v95
                                    }
                                    var v96 int = 128
                                    for fi97 := 0; fi97 < 16; fi97++ {
                                        v96 += (fi97 % ((v96) | 1))
                                        _ = fi97
                                    }
                                    _ = v96
                                }
                            }
                            var v98 int64 = v78
                            _ = v98
                        } else {
                            var v99 int8 = 100
                            _ = v99
                        }
                        for fi100 := 0; fi100 < 14; fi100++ {
                            v101 := GenS5{F0: 0, F1: 837, F2: v1}
                            var v102 bool = (-32768 == 1000)
                            var v103 int16 = 0
                            _ = fi100
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                        var v104 uint8 = uint8(v30)
                        _ = v104
                    }
                }
                var v105 int16 = 42
                _ = v105
                _ = v78
            }
        default:
            fmt.Println(v25, v57)
        }
    }
    switch "a" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            if v1 {
                var v106 uint16 = 1
                for fi107 := 0; fi107 < 12; fi107++ {
                    for fi108 := 0; fi108 < 4; fi108++ {
                        v109 := GenS2{F0: 0}
                        v110 := GenS2{F0: 0}
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                    var v111 float32 = 1e10
                    _ = len(v2)
                    _ = fi107
                    _ = v111
                }
                _ = v106
            } else {
                fmt.Println(v30, v1)
                if ((false && v1) || v1) {
                    if v1 {
                        var v112 string = ""
                        _ = v112
                    }
                    var v113 int64 = 180
                    _ = v113
                }
                if (0.0 < -2.25) {
                    fmt.Println(v25, v1, v30)
                }
            }
        default:
            v57 += (v57 | (^(v57)))
        }
    case "\n\t":
        var v114 uint16 = 1
        _ = v114
    default:
        v115 := GenS1{F0: 589, F1: 128}
        _ = v115
    }
    var v116 int64 = 42
    if (3.14159 <= ((-2.25 - (0.1 * 1.5)) * float64(-2.25))) {
        v30 += 844
        var v117 bool = v1
        _ = v117
    } else {
        v30 -= (v30 / ((v30) | 1))
        v118 := GenS6{F0: 42, F1: v116}
        _ = v118
    }
    v25 = v25
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v116
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v57
}
