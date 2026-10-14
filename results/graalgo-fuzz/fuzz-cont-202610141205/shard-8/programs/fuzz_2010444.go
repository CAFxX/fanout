package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 int8; F2 int16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint64; F1 int }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint; F1 int16; F2 int }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int16; F1 int32 }
type GenS23 struct { F0 uint; F1 uint }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 int16; F1 uint; F2 int8 }
type GenS26 struct { F0 uint64 }
func genFn0(p0 int, p1 uint64) int32 {
    var v5 bool = (127 <= 127)
    if ((!v5) && (false && (v5 && v5))) {
        v6 := GenS1{F0: 2}
        _ = v6
    } else {
        v7 := GenS1{F0: 4294967295}
        v8 := GenS1{F0: 65535}
        _ = v7
        _ = v8
    }
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 2; fi10++ {
        p0 -= 127
        p0 = -2147483648
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: 1, F1: 110, F2: -9223372036854775808}
            p1 = (p1 % ((294) | 1))
            p1 = (p1 - 1)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 42
}
func genFn1() uint64 {
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            if false {
                v14 := []int{(^(fi13)), 32768, fi13, (-(fi13))}
                for fi15 := 0; fi15 < 6; fi15++ {
                    if len(v14) > 0 {
                        v14[0] = fi15
                    }
                    var v16 uint32 = 1
                    v16 -= (v16 << ((18446744073709551615) & 63))
                    var v17 float64 = ((1.5 / 1.0) / 1e10)
                    var v18 uint16 = 1
                    v16 += (v16 ^ (v16 | v16))
                    fmt.Println(v17, v17)
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
            var v19 int16 = -32768
            v19 = (v19 % ((v19) | 1))
            _ = v19
        }
        _ = fi13
    }
    v20 := []int{128, 0}
    _ = v20
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 int16 = -948
    var v3 uint32 = 4294967295
    var v4 int32 = 587
    v2 = 42
    if false {
        var v21 int16 = (v2 << ((869) & 63))
        switch -269 {
        case 9223372036854775807:
            var v22 float32 = (0.0 + (-2.25 - 1.5))
            _ = v22
        case 2147483647:
            v3 += 65535
        case -2147483648:
            switch "foo bar" {
            case "":
                if ((0 == 42) || true) {
                    var v23 uint16 = 2
                    _ = v23
                } else {
                    switch "\n\t" {
                    case "hello":
                        if ((v4 | v4) == (v4 * (v4 * (v4 - (v4 ^ v4))))) {
                            for fi24 := 0; fi24 < 3; fi24++ {
                                v25 := GenS3{F0: 65535, F1: -56}
                                if ((!false) && (v4 <= v4)) {
                                    if true {
                                        var v26 uint16 = 32767
                                        var v27 int = 42
                                        v28 := GenS4{F0: ((^(v26)) == 42), F1: (v4 + v4)}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 bool = ((3.14159 >= (3.14159 + 1e10)) || (1 == v3))
                                        _ = v29
                                    }
                                    switch fi24 {
                                    case -509:
                                        var v30 string = "a"
                                        _ = v30
                                    case -9223372036854775808:
                                        if (-2.25 < (1e10 / 1.0)) {
                                            switch "hello" {
                                            case "\n\t":
                                                switch fi24 {
                                                case -2147483648:
                                                    if (true && true) {
                                                        v31 := GenS5{F0: 1}
                                                        _ = v31
                                                    } else {
                                                        switch 9223372036854775807 {
                                                        case 2147483647:
                                                            var v32 string = ((("\n\t" + "a") + ("foo bar" + "foo bar")) + ("a" + "\n\t"))
                                                            _ = v32
                                                        }
                                                    }
                                                case -448:
                                                    var v33 int = fi24
                                                    _ = v33
                                                default:
                                                    for fi34 := 0; fi34 < 10; fi34++ {
                                                        v4 -= v4
                                                        _ = fi34
                                                    }
                                                }
                                            case "":
                                                v3 = v3
                                            case "foo bar":
                                                v35 := GenS5{F0: 1}
                                                _ = v35
                                            }
                                            for fi36 := 0; fi36 < 5; fi36++ {
                                                if (true || false) {
                                                    if false {
                                                        v1 = append(v1, ("\n\t" + ("foo bar" + ((("" + "foo bar") + "a") + ((("" + "\n\t") + "hello") + "hello")))))
                                                        if true {
                                                            fmt.Println(fi36, v21)
                                                            if (1 < 0) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                                fmt.Println(fi36, v21)
                                                            } else {
                                                                switch fi36 {
                                                                case 128:
                                                                    if (255 == 65535) {
                                                                        if true {
                                                                            if false {
                                                                                if (1e10 == (0.1 - (float64(1.5) - (1e10 * 3.14159)))) {
                                                                                    if false {
                                                                                        if true {
                                                                                            switch (fi36 << ((18446744073709551615) & 63)) {
                                                                                            case 228:
                                                                                                v37 := GenS6{F0: -1, F1: 65535}
                                                                                                if (-1 < 0) {
                                                                                                    v1 = append(v1, "hello")
                                                                                                } else {
                                                                                                    var v38 uint = 588
                                                                                                    _ = v38
                                                                                                }
                                                                                                switch ("hello" + ((((("hello" + "a") + ("hello" + ("" + "a"))) + "\n\t") + ("foo bar" + "foo bar")) + "\n\t")) {
                                                                                                case "hello":
                                                                                                    var v39 uint8 = 198
                                                                                                    _ = v39
                                                                                                default:
                                                                                                    if false {
                                                                                                        var v40 int8 = 127
                                                                                                        v1 = append(v1, "hello")
                                                                                                        var v41 uint32 = 65535
                                                                                                        _ = v40
                                                                                                        _ = v41
                                                                                                    } else {
                                                                                                        switch (-(fi36)) {
                                                                                                        case 127:
                                                                                                            switch fi24 {
                                                                                                            case 1:
                                                                                                                if true {
                                                                                                                    fmt.Println(fi24, fi24, v3)
                                                                                                                    v42 := GenS7{F0: 361}
                                                                                                                    _ = v42
                                                                                                                } else {
                                                                                                                    var v43 uint32 = v3
                                                                                                                    switch "foo bar" {
                                                                                                                    case "hello":
                                                                                                                        fmt.Println(fi36)
                                                                                                                    case "a":
                                                                                                                        var v44 string = "foo bar"
                                                                                                                        _ = v44
                                                                                                                    }
                                                                                                                    v43 = 4294967295
                                                                                                                    _ = v43
                                                                                                                }
                                                                                                            case 128:
                                                                                                                if false {
                                                                                                                    fmt.Println(fi36, v21)
                                                                                                                    v21 = (v21 >> ((2) & 63))
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            if false {
                                                                                                                v45 := GenS6{F0: 0, F1: 255}
                                                                                                                _ = v45
                                                                                                            } else {
                                                                                                                var v46 uint16 = 568
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = "a"
                                                                                                                }
                                                                                                                v2 += (v2 ^ v21)
                                                                                                                _ = v46
                                                                                                            }
                                                                                                        }
                                                                                                        var v47 uint64 = 4294967295
                                                                                                        v2 -= v2
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = ("a" + (("" + ("foo bar" + "\n\t")) + "a"))
                                                                                                        }
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        _ = v47
                                                                                                    }
                                                                                                }
                                                                                                var v48 int32 = v4
                                                                                                _ = v37
                                                                                                _ = v48
                                                                                            case 1:
                                                                                                v49 := GenS8{F0: genFn0((fi24 | (fi24 | fi24)), 124)}
                                                                                                _ = v49
                                                                                            }
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "a"
                                                                                            }
                                                                                            switch ("" + (((("a" + ("hello" + "foo bar")) + "hello") + "hello") + ("a" + "foo bar"))) {
                                                                                            case "\n\t":
                                                                                                v50 := GenS7{F0: 18446744073709551615}
                                                                                                _ = v50
                                                                                            case "foo bar":
                                                                                                v1 = append(v1, "")
                                                                                            case "a":
                                                                                                switch fi36 {
                                                                                                case 128:
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        v21 -= (v21 - (v2 >> ((4294967295) & 63)))
                                                                                                        fmt.Println(fi24, v2, v3)
                                                                                                        v2 = 42
                                                                                                    case "hello":
                                                                                                        var v51 uint16 = 42
                                                                                                        _ = v51
                                                                                                    case "a":
                                                                                                        var v52 string = ("" + "hello")
                                                                                                        _ = v52
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                var v53 uint32 = v3
                                                                                                _ = v53
                                                                                            }
                                                                                            if (7 < v3) {
                                                                                                v2 += (-(v2))
                                                                                                fmt.Println(v21, v21, fi24)
                                                                                            }
                                                                                        }
                                                                                        v54 := GenS9{F0: 65535}
                                                                                        _ = v54
                                                                                    }
                                                                                    switch (^(fi24)) {
                                                                                    case -1:
                                                                                        var v55 bool = false
                                                                                        _ = v55
                                                                                    case 325:
                                                                                        var v56 uint16 = 2
                                                                                        _ = v56
                                                                                    case 4294967295:
                                                                                        v21 = (v21 - -1)
                                                                                    default:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = "foo bar"
                                                                                        }
                                                                                    }
                                                                                    var v57 uint16 = 613
                                                                                    if (v3 > (v3 % ((v3) | 1))) {
                                                                                        _ = len(v1)
                                                                                        fmt.Println(v4, v4, v2)
                                                                                    } else {
                                                                                        fmt.Println(v3, v2)
                                                                                        v58 := GenS9{F0: v57}
                                                                                        var v59 int32 = 981
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                    _ = v57
                                                                                } else {
                                                                                    v2 -= -32768
                                                                                }
                                                                                if ((!false) && false) {
                                                                                    var v60 uint64 = genFn1()
                                                                                    var v61 bool = false
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                } else {
                                                                                    fmt.Println(v2, fi36, v2)
                                                                                    fmt.Println(v21, v3)
                                                                                }
                                                                            } else {
                                                                                v62 := GenS10{F0: (v4 | v4)}
                                                                                v21 += v2
                                                                                if (!true) {
                                                                                    v63 := GenS6{F0: 0, F1: 18446744073709551615}
                                                                                    _ = v63
                                                                                }
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                        var v64 uint64 = 42
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v64
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "foo bar"
                                                                        }
                                                                        v3 -= (^(v3))
                                                                        v1 = append(v1, ("" + "\n\t"))
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ("" + "\n\t")
                                                                        }
                                                                    }
                                                                case 42:
                                                                    var v65 int32 = (v4 * v4)
                                                                    _ = v65
                                                                case 4294967295:
                                                                    fmt.Println(v21)
                                                                }
                                                            }
                                                            if (true && true) {
                                                                var v66 uint64 = genFn1()
                                                                v66 = (v66 >> ((252) & 63))
                                                                _ = v66
                                                            }
                                                        }
                                                        var v67 string = "\n\t"
                                                        _ = v67
                                                    }
                                                    var v68 uint = 2
                                                    _ = v68
                                                }
                                                var v69 int16 = v2
                                                if true {
                                                    if false {
                                                        var v70 uint8 = 1
                                                        var v71 bool = ((v3 * (v3 + v3)) > 4294967295)
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        v3 = (v3 & (v3 % (((v3 / ((v3) | 1))) | 1)))
                                                        var v72 uint = 65535
                                                        if false {
                                                            var v73 uint32 = 2
                                                            v74 := GenS3{F0: 32767, F1: 1}
                                                            switch "a" {
                                                            case "a":
                                                                fmt.Println(v2)
                                                            case "foo bar":
                                                                var v75 int64 = 1
                                                                _ = v75
                                                            }
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        _ = v72
                                                    }
                                                    v21 -= (v2 * v69)
                                                    v4 = 2147483647
                                                } else {
                                                    if (v4 <= (v4 << ((497) & 63))) {
                                                        v2 += v69
                                                    } else {
                                                        var v76 int16 = (v21 - v69)
                                                        fmt.Println(v2, v21, fi36)
                                                        _ = v76
                                                    }
                                                }
                                                _ = fi36
                                                _ = v69
                                            }
                                            switch fi24 {
                                            case -639:
                                                switch (fi24 & fi24) {
                                                case -795:
                                                    v4 = (-(v4))
                                                default:
                                                    if (!((-(v21)) < -148)) {
                                                        v77 := GenS11{F0: 2, F1: 1, F2: v2}
                                                        _ = v77
                                                    }
                                                }
                                            case 32767:
                                                fmt.Println(v2, fi24, v3)
                                            default:
                                                var v78 int8 = 60
                                                _ = v78
                                            }
                                        } else {
                                            switch 4294967295 {
                                            case 32767:
                                                fmt.Println(v21)
                                            }
                                            v21 += 248
                                        }
                                    default:
                                        v2 = v21
                                    }
                                } else {
                                    v4 = v4
                                }
                                if (!false) {
                                    v4 += genFn0((fi24 + 2147483647), 2)
                                    var v79 uint8 = 0
                                    _ = v79
                                }
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            if (v4 == (v4 - v4)) {
                                _ = len(v1)
                                v1 = append(v1, "a")
                                var v80 string = "foo bar"
                                _ = v80
                            } else {
                                v3 = v3
                                switch 1 {
                                case 9223372036854775807:
                                    if (12 < 0) {
                                        fmt.Println(v21, v21, v3)
                                        fmt.Println(v21)
                                    } else {
                                        v2 -= 1000
                                        if (!(65535 == uint(v3))) {
                                            fmt.Println(v3)
                                            v21 += v2
                                        }
                                        var v81 uint16 = 272
                                        _ = v81
                                    }
                                case -9223372036854775808:
                                    if true {
                                        fmt.Println(v4)
                                        for fi82 := 0; fi82 < 14; fi82++ {
                                            if (((3.14159 - float32((3.14159 * float64(-2.25)))) + 3.14159) != float32(0.0)) {
                                                v3 -= v3
                                                var v83 bool = (true && false)
                                                _ = v83
                                            } else {
                                                var v84 uint = 255
                                                _ = v84
                                            }
                                            v85 := GenS12{F0: (v4 % ((42) | 1))}
                                            v4 = (v4 / (((v4 << ((42) & 63))) | 1))
                                            _ = fi82
                                            _ = v85
                                        }
                                    }
                                case -962:
                                    for fi86 := 0; fi86 < 6; fi86++ {
                                        v21 += (v2 * v2)
                                        var v87 uint64 = 1
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        _ = fi86
                                        _ = v87
                                    }
                                default:
                                    v88 := GenS3{F0: 0, F1: -1}
                                    _ = v88
                                }
                                var v89 int64 = -1
                                _ = v89
                            }
                        }
                    default:
                        v4 -= v4
                    }
                }
            case "a":
                if (!false) {
                    v4 -= v4
                } else {
                    if true {
                        if (!false) {
                            v21 = (v21 - (v21 * -204))
                            var v90 uint8 = 127
                            _ = v90
                        } else {
                            for fi91 := 0; fi91 < 5; fi91++ {
                                if len(v1) > 0 {
                                    v1[0] = (("" + ((("hello" + "") + "") + "foo bar")) + ("foo bar" + "foo bar"))
                                }
                                var v92 string = (("\n\t" + ("hello" + "foo bar")) + ("foo bar" + ("a" + "\n\t")))
                                _ = fi91
                                _ = v92
                            }
                            switch 32768 {
                            case -421:
                                var v93 uint32 = v3
                                _ = v93
                            case 2147483647:
                                var v94 uint16 = 42
                                _ = v94
                            case 127:
                                v95 := GenS5{F0: 721}
                                _ = v95
                            default:
                                var v96 bool = ((v2 * (^(v2))) > v2)
                                _ = v96
                            }
                            v97 := GenS13{F0: v21}
                            _ = v97
                        }
                        v1 = append(v1, (((((("a" + ("hello" + "foo bar")) + (("\n\t" + ("\n\t" + "foo bar")) + "hello")) + (("" + ("a" + "")) + "foo bar")) + "foo bar") + "a") + "hello"))
                        for fi98 := 0; fi98 < 13; fi98++ {
                            var v99 uint8 = 42
                            if (true && true) {
                                fmt.Println(v99)
                                var v100 int = fi98
                                _ = v100
                            }
                            _ = fi98
                            _ = v99
                        }
                    }
                    var v101 uint = 121
                    var v102 string = ("a" + ("hello" + "hello"))
                    _ = v101
                    _ = v102
                }
            }
        }
        var v103 float32 = (1.5 - 3.14159)
        _ = v103
        _ = v21
    } else {
        v104 := GenS1{F0: 42}
        v105 := GenS14{F0: uint(v3)}
        v4 = (v4 >> ((2) & 63))
        _ = v104
        _ = v105
    }
    var v106 float32 = float32((float64(0.1) * 0.1))
    switch ("foo bar" + "foo bar") {
    case "a":
        v4 += 0
    case "":
        for fi107 := 0; fi107 < 9; fi107++ {
            for fi108 := 0; fi108 < 15; fi108++ {
                v2 = v2
                switch "foo bar" {
                case "a":
                    fmt.Println(v2)
                case "":
                    v109 := GenS15{F0: 42, F1: "a"}
                    if len(v1) > 0 {
                        v1[0] = (("\n\t" + "") + "\n\t")
                    }
                    if false {
                        v106 = v106
                    } else {
                        v110 := GenS1{F0: 0}
                        _ = v110
                    }
                    _ = v109
                case "hello":
                    var v111 int16 = -1
                    _ = v111
                default:
                    v4 = v4
                }
                _ = fi108
            }
            if (v4 >= int32(fi107)) {
                for fi112 := 0; fi112 < 7; fi112++ {
                    if false {
                        switch 2147483647 {
                        case 10:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 32768:
                            var v113 uint32 = 65535
                            _ = v113
                        case 192:
                            v4 = -1
                        }
                        v114 := GenS16{F0: 1}
                        v115 := GenS7{F0: 65535}
                        _ = v114
                        _ = v115
                    }
                    switch ("\n\t" + "") {
                    case "a":
                        v1 = append(v1, "foo bar")
                    case "hello":
                        switch "foo bar" {
                        case "":
                            if (true || false) {
                                if (1 <= 1) {
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                }
                                var v116 int64 = 0
                                _ = v116
                            } else {
                                fmt.Println(v4, v4)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (true || (true || false)) {
                                    if (18446744073709551615 == 2) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch fi107 {
                                            case 1:
                                                v117 := GenS17{F0: 1, F1: (fi107 % (((-(fi112))) | 1))}
                                                _ = v117
                                            case 127:
                                                if (((!true) && false) && (0 != 0)) {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v118 uint16 = 283
                                                        _ = v118
                                                    case "":
                                                        var v119 int8 = 1
                                                        _ = v119
                                                    }
                                                    fmt.Println(v2)
                                                    var v120 uint64 = 42
                                                    v121 := GenS8{F0: v4}
                                                    _ = v120
                                                    _ = v121
                                                }
                                            case 32767:
                                                switch "" {
                                                case "foo bar":
                                                    v3 -= v3
                                                default:
                                                    v2 += (v2 << ((65535) & 63))
                                                }
                                            }
                                        case "":
                                            if ("foo bar" < "a") {
                                                if false {
                                                    var v122 float32 = v106
                                                    var v123 uint64 = 4294967295
                                                    v124 := GenS18{F0: v3}
                                                    _ = v122
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    fmt.Println(v2, v106)
                                                    v2 = v2
                                                    var v125 int = fi112
                                                    _ = v125
                                                }
                                                fmt.Println(v2, v3)
                                            }
                                        }
                                        switch ("hello" + "a") {
                                        case "a":
                                            var v126 uint64 = 2
                                            _ = v126
                                        case "\n\t":
                                            v2 = v2
                                            _ = len(v1)
                                            var v127 int = int(v2)
                                            _ = v127
                                        case "":
                                            fmt.Println(fi112, v2, fi112)
                                        }
                                        v1 = append(v1, "hello")
                                        if false {
                                            if (v106 < (v106 / 3.14159)) {
                                                switch "\n\t" {
                                                case "a":
                                                    v128 := GenS16{F0: 2}
                                                    _ = v128
                                                case "foo bar":
                                                    fmt.Println(v106, fi112)
                                                case "":
                                                    if (!true) {
                                                        var v129 float64 = -2.25
                                                        _ = len(v1)
                                                        _ = v129
                                                    }
                                                default:
                                                    switch ("a" + ("hello" + "foo bar")) {
                                                    case "":
                                                        if len(v1) > 0 {
                                                            v1[0] = ("hello" + "")
                                                        }
                                                    case "\n\t":
                                                        v2 -= v2
                                                        switch "a" {
                                                        case "":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "a":
                                                            v3 = v3
                                                        }
                                                        fmt.Println(fi112)
                                                        var v130 uint32 = (v3 % ((v3) | 1))
                                                        v106 -= v106
                                                        _ = v130
                                                    case "a":
                                                        v131 := GenS5{F0: 255}
                                                        _ = v131
                                                    }
                                                }
                                                var v132 float64 = 0.0
                                                if false {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v133 uint8 = 128
                                                    _ = v133
                                                }
                                                _ = v132
                                            } else {
                                                _ = len(v1)
                                                v134 := GenS13{F0: (v2 * -32768)}
                                                v135 := GenS11{F0: 32767, F1: 100, F2: (v2 >> ((2) & 63))}
                                                _ = v134
                                                _ = v135
                                            }
                                            var v136 uint32 = (v3 & v3)
                                            _ = v136
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v137 string = "\n\t"
                                            var v138 bool = (false || false)
                                            _ = v137
                                            _ = v138
                                        }
                                    }
                                    v2 -= v2
                                    if true {
                                        v139 := GenS3{F0: 1, F1: -1}
                                        _ = v139
                                    } else {
                                        fmt.Println(v106, v4)
                                        if (!true) {
                                            fmt.Println(fi112, v106, v3)
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v140 int32 = 1
                                            _ = v140
                                        }
                                        var v141 float32 = ((v106 + v106) * 3.14159)
                                        _ = v141
                                    }
                                } else {
                                    var v142 bool = ((v3 / (((v3 * v3)) | 1)) > v3)
                                    fmt.Println(fi107, fi107, v2)
                                    var v143 float32 = v106
                                    var v144 uint32 = 2
                                    _ = v142
                                    _ = v143
                                    _ = v144
                                }
                            }
                        case "foo bar":
                            var v145 uint = 0
                            _ = v145
                        case "a":
                            if true {
                                var v146 uint32 = v3
                                _ = v146
                            }
                        }
                    default:
                        if (304 == 233) {
                            switch (fi107 % ((42) | 1)) {
                            case -178:
                                switch -9223372036854775808 {
                                case 128:
                                    var v147 uint64 = genFn1()
                                    _ = v147
                                case 0:
                                    v4 += 68
                                case -348:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    if (false || true) {
                                        v148 := GenS9{F0: 621}
                                        v149 := GenS19{F0: 1, F1: v2, F2: (fi107 & -1)}
                                        v2 += v2
                                        v150 := GenS2{F0: 1, F1: 115, F2: 1}
                                        _ = v148
                                        _ = v149
                                        _ = v150
                                    }
                                }
                            case -1:
                                v4 = 1
                            case 1:
                                if true {
                                    if false {
                                        if (!(-1 != -1)) {
                                            switch "\n\t" {
                                            case "\n\t":
                                                if (float32(1e10) != 0.1) {
                                                    var v151 uint16 = 655
                                                    switch "" {
                                                    case "foo bar":
                                                        if true {
                                                            var v152 bool = ((v4 >> ((18446744073709551615) & 63)) != (v4 << ((42) & 63)))
                                                            switch (fi112 - (-(fi107))) {
                                                            case -2147483648:
                                                                var v153 int64 = -697
                                                                _ = v153
                                                            case 32768:
                                                                var v154 int32 = genFn0(fi112, 166)
                                                                _ = v154
                                                            case -9223372036854775808:
                                                                if (v152 && (!false)) {
                                                                    var v155 uint16 = v151
                                                                    if v152 {
                                                                        var v156 uint32 = (v3 * v3)
                                                                        _ = v156
                                                                    }
                                                                    var v157 bool = (!false)
                                                                    _ = v155
                                                                    _ = v157
                                                                }
                                                            }
                                                            _ = v152
                                                        }
                                                    case "hello":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "\n\t":
                                                        v158 := GenS15{F0: (v151 * 2), F1: "a"}
                                                        var v159 uint32 = (^(v3))
                                                        switch 1 {
                                                        case -983:
                                                            var v160 uint8 = 127
                                                            _ = v160
                                                        case 4294967295:
                                                            v161 := GenS20{F0: 2}
                                                            _ = v161
                                                        case 127:
                                                            v1 = append(v1, "foo bar")
                                                            var v162 int64 = 574
                                                            _ = v162
                                                        }
                                                        var v163 int = (fi112 & fi112)
                                                        _ = v158
                                                        _ = v159
                                                        _ = v163
                                                    }
                                                    var v164 uint8 = 255
                                                    _ = v151
                                                    _ = v164
                                                }
                                            }
                                            switch int(v3) {
                                            case 1:
                                                v2 = (v2 % ((42) | 1))
                                            case 32767:
                                                v165 := GenS12{F0: (v4 % ((v4) | 1))}
                                                _ = v165
                                            default:
                                                if true {
                                                    v166 := GenS14{F0: 4294967295}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v166
                                                }
                                            }
                                            var v167 string = (("foo bar" + "foo bar") + "\n\t")
                                            _ = v167
                                        } else {
                                            var v168 int64 = 1
                                            var v169 bool = true
                                            v3 += v3
                                            _ = v168
                                            _ = v169
                                        }
                                    } else {
                                        _ = len(v1)
                                        v1 = append(v1, "a")
                                        v170 := GenS21{F0: 86}
                                        _ = v170
                                    }
                                }
                            }
                            v2 -= v2
                            v171 := GenS17{F0: 42, F1: fi107}
                            v4 += genFn0((^(fi112)), 0)
                            v1 = append(v1, "hello")
                            _ = v171
                        }
                    }
                    _ = fi112
                }
                fmt.Println(fi107, v106, v2)
                var v172 int16 = (v2 - 0)
                _ = v172
            } else {
                fmt.Println(v4)
                var v173 string = "\n\t"
                v2 = (^(v2))
                _ = v173
            }
            var v174 int32 = v4
            _ = fi107
            _ = v174
        }
    }
    if (209 <= 109) {
        v106 = (v106 * 3.14159)
        var v175 int16 = (v2 ^ (v2 ^ 126))
        _ = v175
    }
    if true {
        switch 891 {
        case 9223372036854775807:
            v4 -= 1
        case 32768:
            for fi176 := 0; fi176 < 17; fi176++ {
                v177 := GenS22{F0: (v2 ^ v2), F1: (v4 | (v4 + 2147483647))}
                _ = fi176
                _ = v177
            }
        case 2147483647:
            fmt.Println(v2)
        }
    } else {
        switch 32767 {
        case -9223372036854775808:
            if (v4 == v4) {
                v178 := GenS23{F0: 112, F1: 430}
                var v179 string = (("\n\t" + "") + "")
                _ = v178
                _ = v179
            }
        case 32768:
            fmt.Println(v2, v3, v106)
        default:
            var v180 int8 = 127
            _ = v180
        }
    }
    if ((^(v3)) < 333) {
        v1 = append(v1, "foo bar")
        switch -1 {
        case 9223372036854775807:
            v4 = v4
        default:
            fmt.Println(v4, v2)
        }
        v106 = v106
    }
    if (true || false) {
        for fi181 := 0; fi181 < 1; fi181++ {
            v3 -= 1
            v1 = append(v1, ("\n\t" + "hello"))
            for fi182 := 0; fi182 < 7; fi182++ {
                var v183 int64 = 42
                v106 = (0.0 * (-2.25 / 0.1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi182
                _ = v183
            }
            _ = fi181
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v184 := GenS24{F0: true}
        v4 = (v4 - (-(v4)))
        _ = v184
    }
    var v185 uint32 = (v3 % ((0) | 1))
    v186 := GenS4{F0: ((float32((-2.25 - 0.0)) + v106) >= float32(0.0)), F1: v4}
    for fi187 := 0; fi187 < 12; fi187++ {
        v3 += (^(v185))
        if (!(!false)) {
            fmt.Println(v3, v106, v2)
        } else {
            if (v2 < -228) {
                var v188 float64 = 0.0
                v189 := GenS25{F0: int16(v4), F1: 65535, F2: 127}
                _ = v188
                _ = v189
            }
            v3 = 15
        }
        if false {
            fmt.Println(v4, v106)
            v190 := GenS26{F0: genFn1()}
            _ = v190
        }
        _ = fi187
    }
    fmt.Println(v2, v106, v2)
    var v191 uint = 429
    var v192 uint8 = 2
    if false {
        for fi193 := 0; fi193 < 15; fi193++ {
            switch "a" {
            case "foo bar":
                v194 := GenS14{F0: (^(v191))}
                _ = v194
            }
            var v195 int = (fi193 & (^(fi193)))
            _ = fi193
            _ = v195
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v185:", v185)
    fmt.Println("v191:", v191)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v185
    _ = v186
    _ = v191
    _ = v192
    _ = v2
    _ = v3
    _ = v4
}
