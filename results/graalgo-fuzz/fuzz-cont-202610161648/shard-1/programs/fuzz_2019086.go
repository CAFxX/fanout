package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int; F1 uint8; F2 int64 }
type GenS7 struct { F0 int64; F1 int32; F2 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8; F1 uint; F2 bool }
type GenS12 struct { F0 string; F1 string; F2 int32 }
type GenS13 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int32; F1 int8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string; F1 bool; F2 uint16 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS20 struct { F0 uint32; F1 int32 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 int64; F1 string; F2 uint }
type GenS25 struct { F0 uint64; F1 int64; F2 int8 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 int; F1 uint16; F2 int8 }
type GenS28 struct { F0 uint }
type GenS29 struct { F0 bool; F1 int8; F2 uint64 }
type GenS30 struct { F0 uint8; F1 bool; F2 uint }
type GenS31 struct { F0 uint32; F1 int8 }
type GenS32 struct { F0 uint8; F1 int; F2 string }
type GenS33 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS34 struct { F0 uint8; F1 bool }
type GenS35 struct { F0 int32; F1 int32 }
type GenS36 struct { F0 int8; F1 uint8 }
func genFn0(p0 int32, p1 int32, p2 int64) string {
    v6 := GenS1{F0: 100}
    var v7 float32 = 0.1
    if true {
        for fi8 := 0; fi8 < 5; fi8++ {
            p0 += (p1 ^ 2147483647)
            v9 := []int{}
            p0 = p1
            _ = fi8
            _ = v9
        }
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint16 = 0
    v2 := []int{-9223372036854775808, -2147483648, 9223372036854775807}
    var v3 uint32 = 91
    var v4 string = "\n\t"
    v5 := []string{(((v4 + ("\n\t" + "\n\t")) + (v4 + "")) + "a"), "foo bar", "a"}
    for fi11 := 0; fi11 < 12; fi11++ {
        if (2 < 2) {
            var v12 string = ((v4 + v4) + (v4 + ((v4 + v4) + (v4 + "\n\t"))))
            var v13 int = 4294967295
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: -61}
    v15 := GenS1{F0: 65}
    for fi16 := 0; fi16 < 10; fi16++ {
        if (!((!false) || (!true))) {
            var v17 int64 = -1
            _ = v17
        } else {
            fmt.Println(v3, v1, v1)
            if false {
                switch "hello" {
                case "":
                    v4 = "a"
                case "\n\t":
                    if (false && (v4 != (((v4 + "\n\t") + (v4 + "\n\t")) + ((v4 + "a") + v4)))) {
                        var v18 int32 = -2147483648
                        for fi19 := 0; fi19 < 5; fi19++ {
                            switch fi16 {
                            case -1:
                                if (false && (0 >= 42)) {
                                    v5 = append(v5, v4)
                                    var v20 int16 = 1
                                    v4 = "\n\t"
                                    _ = v20
                                }
                            default:
                                v4 = "a"
                            }
                            if false {
                                fmt.Println(v18)
                                switch fi16 {
                                case 42:
                                    v3 = v3
                                case 128:
                                    if (32767 > v1) {
                                        if (0 < 4294967295) {
                                            v4 = "a"
                                        }
                                        v21 := GenS2{F0: 127, F1: "a"}
                                        _ = v21
                                    }
                                case 2147483647:
                                    switch v4 {
                                    case "\n\t":
                                        v22 := GenS2{F0: -1, F1: v4}
                                        _ = v22
                                    case "a":
                                        var v23 bool = ((0.1 + (3.14159 + (((1.5 - 1e10) - 1.5) / 3.14159))) != ((0.1 * 3.14159) - (3.14159 + 3.14159)))
                                        _ = v23
                                    }
                                }
                                if (!(18446744073709551615 <= 1)) {
                                    var v24 string = "\n\t"
                                    _ = v24
                                }
                            }
                            v25 := GenS1{F0: 1}
                            _ = fi19
                            _ = v25
                        }
                        if true {
                            switch ((v4 + v4) + v4) {
                            case "\n\t":
                                v4 = "a"
                            case "":
                                var v26 int = int(v1)
                                _ = v26
                            }
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    if (42 <= 2) {
                                        for fi27 := 0; fi27 < 7; fi27++ {
                                            if false {
                                                if (240 < 248) {
                                                    v1 = 7
                                                    var v28 float64 = 0.1
                                                    _ = v28
                                                }
                                                var v29 float32 = float32(1e10)
                                                if (int8(v18) < 22) {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v30 uint = 1
                                                        _ = v30
                                                    default:
                                                        var v31 int = (fi27 + -9223372036854775808)
                                                        _ = v31
                                                    }
                                                    _ = len(v2)
                                                    var v32 uint64 = 42
                                                    _ = v32
                                                }
                                                _ = v29
                                            } else {
                                                var v33 bool = false
                                                switch -1 {
                                                case -970:
                                                    var v34 string = ("a" + (v4 + v4))
                                                    _ = v34
                                                case -345:
                                                    var v35 string = ("a" + v4)
                                                    _ = v35
                                                case -2147483648:
                                                    switch v4 {
                                                    case "foo bar":
                                                        if v33 {
                                                            v36 := GenS3{F0: 704, F1: 42}
                                                            v5 = append(v5, ("hello" + (v4 + v4)))
                                                            v37 := GenS2{F0: 100, F1: v4}
                                                            var v38 uint = 18446744073709551615
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    }
                                                default:
                                                    v33 = (42 == 1)
                                                }
                                                _ = v33
                                            }
                                            v18 = v18
                                            _ = fi27
                                        }
                                        if (false && (147 >= 42)) {
                                            v39 := GenS4{F0: 61, F1: 283, F2: v3}
                                            _ = v39
                                        } else {
                                            for fi40 := 0; fi40 < 11; fi40++ {
                                                v3 = v3
                                                if len(v2) > 0 {
                                                    v2[0] = 4294967295
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = "hello"
                                                }
                                                _ = fi40
                                            }
                                        }
                                        if (-2.25 < 3.14159) {
                                            v4 = v4
                                        }
                                    }
                                default:
                                    v18 -= v18
                                }
                                for fi41 := 0; fi41 < 19; fi41++ {
                                    var v42 int16 = -564
                                    _ = fi41
                                    _ = v42
                                }
                                if false {
                                    v43 := GenS5{F0: -32768, F1: 49}
                                    _ = v43
                                }
                            } else {
                                switch v4 {
                                case "foo bar":
                                    var v44 int = (fi16 | fi16)
                                    _ = v44
                                default:
                                    for fi45 := 0; fi45 < 20; fi45++ {
                                        v2 = append(v2, fi45)
                                        if ((v4 + v4) >= ("\n\t" + v4)) {
                                            v18 = v18
                                            fmt.Println(fi16)
                                            v2 = append(v2, fi16)
                                            v4 = v4
                                        }
                                        fmt.Println(fi45)
                                        _ = fi45
                                    }
                                }
                            }
                        } else {
                            for fi46 := 0; fi46 < 16; fi46++ {
                                if (false || true) {
                                    if false {
                                        switch (fi16 / ((fi46) | 1)) {
                                        case 127:
                                            if ((2 >= 1) && (-195 >= v18)) {
                                                var v47 int8 = 0
                                                switch v4 {
                                                case "":
                                                    var v48 int16 = 32767
                                                    _ = v48
                                                default:
                                                    v49 := GenS2{F0: int8(fi46), F1: (("" + v4) + (((v4 + v4) + v4) + ("\n\t" + "a")))}
                                                    _ = v49
                                                }
                                                v1 -= (v1 / ((v1) | 1))
                                                v4 = v4
                                                var v50 uint = 1
                                                if true {
                                                    v51 := GenS1{F0: v47}
                                                    _ = v51
                                                } else {
                                                    v52 := GenS6{F0: fi16, F1: 56, F2: 9223372036854775807}
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi46 + fi46)
                                                    }
                                                    _ = v52
                                                }
                                                _ = v47
                                                _ = v50
                                            }
                                        default:
                                            fmt.Println(v1, fi16)
                                        }
                                        if ((int16(v1) < 764) || (uint64(v3) >= 871)) {
                                            var v53 int16 = 1000
                                            var v54 int8 = 100
                                            var v55 float32 = 1.5
                                            fmt.Println(fi46)
                                            if false {
                                                v3 = (v3 * v3)
                                                v53 = (v53 / ((v53) | 1))
                                                v4 = (v4 + v4)
                                                if (!(2 <= 0)) {
                                                    fmt.Println(v1)
                                                    switch ("\n\t" + v4) {
                                                    case "hello":
                                                        if false {
                                                            v53 = -1
                                                            if false {
                                                                if false {
                                                                    if true {
                                                                        if false {
                                                                            var v56 int8 = (v54 % ((v54) | 1))
                                                                            var v57 uint64 = 17
                                                                            var v58 uint32 = v3
                                                                            _ = v56
                                                                            _ = v57
                                                                            _ = v58
                                                                        } else {
                                                                            var v59 uint = 18446744073709551615
                                                                            _ = v59
                                                                        }
                                                                        switch (fi16 << ((447) & 63)) {
                                                                        case 2147483647:
                                                                            v60 := GenS3{F0: 1, F1: 774}
                                                                            _ = v60
                                                                        }
                                                                        v4 = v4
                                                                    }
                                                                } else {
                                                                    var v61 uint = 18446744073709551615
                                                                    v55 -= v55
                                                                    var v62 uint64 = 18446744073709551615
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                v54 = -97
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi16
                                                                }
                                                            } else {
                                                                if true {
                                                                    v55 = 1e10
                                                                    fmt.Println(v18)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = ""
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v1 -= (v1 | (v1 | v1))
                                                    case "\n\t":
                                                        var v63 float64 = 1.5
                                                        _ = v63
                                                    }
                                                } else {
                                                    v53 += 997
                                                    fmt.Println(v18, v1, fi16)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch (fi46 % ((2147483647) | 1)) {
                                                        case 816:
                                                            var v64 int64 = -9223372036854775808
                                                            _ = v64
                                                        case -469:
                                                            switch ("foo bar" + v4) {
                                                            case "\n\t":
                                                                v65 := GenS7{F0: -751, F1: (v18 + (v18 ^ (v18 - v18))), F2: (v54 >> ((1) & 63))}
                                                                _ = v65
                                                            default:
                                                                v4 = v4
                                                            }
                                                        default:
                                                            if true {
                                                                v1 -= (v1 * v1)
                                                                v55 = 0.0
                                                                fmt.Println(v3, v4, v53)
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v66 := GenS8{F0: 9}
                                            if ((-32768 > 1) || (fi16 > fi16)) {
                                                var v67 int32 = (-(v18))
                                                _ = v67
                                            }
                                            v68 := GenS7{F0: 9223372036854775807, F1: v18, F2: 95}
                                            _ = v66
                                            _ = v68
                                        }
                                        fmt.Println(v4)
                                        var v69 uint = 4294967295
                                        v3 = v3
                                        switch (fi16 % ((9223372036854775807) | 1)) {
                                        case 32767:
                                            v69 += uint(v3)
                                        case 128:
                                            fmt.Println(v69)
                                        default:
                                            if true {
                                                var v70 int32 = (v18 * v18)
                                                _ = v70
                                            }
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v71 := GenS8{F0: 223}
                                            _ = v71
                                        case "foo bar":
                                            var v72 int32 = -708
                                            _ = v72
                                        default:
                                            var v73 float64 = ((-2.25 * 0.0) + float64(0.0))
                                            _ = v73
                                        }
                                        _ = v69
                                    }
                                    var v74 uint16 = (v1 - 32767)
                                    v1 += (v1 / (((^(v1))) | 1))
                                    _ = v74
                                } else {
                                    fmt.Println(fi46, fi16, fi16)
                                    if (0.1 != 0.1) {
                                        if (795 == 42) {
                                            v4 = (v4 + (v4 + (v4 + v4)))
                                            _ = len(v5)
                                            _ = len(v2)
                                        }
                                        var v75 int8 = -1
                                        v1 += (^(v1))
                                        _ = v75
                                    }
                                }
                                _ = fi46
                            }
                        }
                        _ = v18
                    }
                }
            }
            v3 = v3
        }
        switch v4 {
        case "\n\t":
            fmt.Println(v4, v3)
        default:
            v3 += v3
        }
        v5 = append(v5, "")
        _ = fi16
    }
    if false {
        v1 = v1
        if false {
            for fi76 := 0; fi76 < 6; fi76++ {
                v4 = ((v4 + v4) + v4)
                _ = fi76
            }
            if (false || true) {
                for fi77 := 0; fi77 < 16; fi77++ {
                    if false {
                        for fi78 := 0; fi78 < 15; fi78++ {
                            var v79 bool = false
                            var v80 int8 = 127
                            if v79 {
                                switch -171 {
                                case -1:
                                    v3 = v3
                                    if v79 {
                                        var v81 int8 = v80
                                        _ = v81
                                    } else {
                                        switch 128 {
                                        case -980:
                                            v82 := GenS2{F0: v80, F1: (v4 + (v4 + v4))}
                                            _ = v82
                                        case 1:
                                            if (!true) {
                                                if false {
                                                    var v83 float64 = (1.5 - (3.14159 / 0.1))
                                                    v84 := GenS6{F0: (fi78 / ((fi77) | 1)), F1: 255, F2: 9223372036854775807}
                                                    if v79 {
                                                        fmt.Println(v4, fi78)
                                                    } else {
                                                        v85 := GenS3{F0: (v1 >> ((0) & 63)), F1: 0}
                                                        v79 = true
                                                        v86 := GenS9{F0: 0, F1: v1}
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                    fmt.Println(fi77)
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    if v79 {
                                                        v1 = v1
                                                        var v87 uint = 786
                                                        _ = v87
                                                    } else {
                                                        var v88 uint = 4294967295
                                                        fmt.Println(v88)
                                                        _ = v88
                                                    }
                                                    if ((0 <= -1) && (0 != 170)) {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        switch v4 {
                                                        case "\n\t":
                                                            switch (fi78 << ((267) & 63)) {
                                                            case 32768:
                                                                v5 = append(v5, v4)
                                                            case -2147483648:
                                                                v5 = append(v5, v4)
                                                            case -9223372036854775808:
                                                                switch v4 {
                                                                case "foo bar":
                                                                    var v89 int = (fi78 << ((2) & 63))
                                                                    _ = v89
                                                                default:
                                                                    if (fi78 == fi78) {
                                                                        if true {
                                                                            if ((v3 << ((60) & 63)) <= (v3 ^ v3)) {
                                                                                v79 = (42 <= -9223372036854775808)
                                                                                v90 := GenS10{F0: 18446744073709551615}
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    _ = len(v2)
                                                                                }
                                                                                v4 = ((((((v4 + v4) + ("hello" + v4)) + (("" + "a") + (v4 + v4))) + v4) + "\n\t") + ("hello" + "hello"))
                                                                                v1 += (v1 * v1)
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                            var v91 int32 = 42
                                                                            _ = v91
                                                                        }
                                                                        var v92 int8 = -115
                                                                        if true {
                                                                            switch (-(fi78)) {
                                                                            case 128:
                                                                                var v93 uint = 255
                                                                                _ = v93
                                                                            case -69:
                                                                                if (0.0 < -2.25) {
                                                                                    v79 = v79
                                                                                    var v94 uint8 = 127
                                                                                    _ = v94
                                                                                }
                                                                            case 42:
                                                                                var v95 uint = 0
                                                                                _ = v95
                                                                            }
                                                                        } else {
                                                                            var v96 uint = 65535
                                                                            _ = v96
                                                                        }
                                                                        v79 = v79
                                                                        _ = v92
                                                                    }
                                                                }
                                                            }
                                                            v79 = (((fi78 + fi78) <= fi77) || (!((-1 > v80) && v79)))
                                                        default:
                                                            if true {
                                                                if (v79 && (!(v79 && false))) {
                                                                    var v97 int8 = (v80 - (-(v80)))
                                                                    var v98 uint = 717
                                                                    var v99 int = fi77
                                                                    _ = v97
                                                                    _ = v98
                                                                    _ = v99
                                                                } else {
                                                                    var v100 int32 = -1
                                                                    _ = v100
                                                                }
                                                                v4 = v4
                                                                switch (^(fi78)) {
                                                                case -9223372036854775808:
                                                                    v80 += (v80 ^ 100)
                                                                    v79 = false
                                                                    var v101 uint16 = v1
                                                                    _ = v101
                                                                case 291:
                                                                    _ = len(v5)
                                                                case -62:
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = (v4 + v4)
                                                    }
                                                }
                                                var v102 int64 = -9223372036854775808
                                                var v103 int64 = v102
                                                _ = v102
                                                _ = v103
                                            }
                                        case 4294967295:
                                            if len(v2) > 0 {
                                                v2[0] = int(v1)
                                            }
                                        }
                                        v2 = append(v2, fi78)
                                    }
                                case -379:
                                    var v104 uint = 18446744073709551615
                                    _ = v104
                                default:
                                    fmt.Println(v1, v79)
                                }
                                if false {
                                    fmt.Println(v79, v80, v3)
                                    switch v4 {
                                    case "":
                                        v4 = v4
                                    case "a":
                                        if (!v79) {
                                            _ = len(v5)
                                        }
                                    case "hello":
                                        var v105 uint16 = 136
                                        v5 = append(v5, "hello")
                                        if (true && true) {
                                            v4 = ""
                                            v4 = v4
                                        } else {
                                            v105 = v1
                                            if ((true && v79) || (1 >= uint32(v1))) {
                                                v1 += v1
                                                v79 = (((v1 & (v1 % ((v105) | 1))) < (v105 << ((65535) & 63))) || v79)
                                            } else {
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if (!(v80 > v80)) {
                                                        switch ("" + v4) {
                                                        case "\n\t":
                                                            v106 := GenS10{F0: 18446744073709551615}
                                                            _ = v106
                                                        }
                                                    } else {
                                                        v4 = ((("hello" + v4) + v4) + (v4 + "hello"))
                                                    }
                                                } else {
                                                    switch int(v80) {
                                                    case 593:
                                                        v79 = false
                                                    case 42:
                                                        if v79 {
                                                            v1 = 32767
                                                            if (true && ((!(-2147483648 <= fi77)) && v79)) {
                                                                switch ((("a" + ("foo bar" + ("" + v4))) + (v4 + "a")) + v4) {
                                                                case "hello":
                                                                    if v79 {
                                                                        _ = len(v2)
                                                                        if false {
                                                                            var v107 int32 = int32(v80)
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            switch v4 {
                                                                            case "a":
                                                                                fmt.Println(v80, v79, v1)
                                                                            case "\n\t":
                                                                                v108 := GenS3{F0: (v1 & (v1 >> ((321) & 63))), F1: 42}
                                                                                _ = v108
                                                                            }
                                                                            _ = v107
                                                                        }
                                                                        var v109 float32 = 1.5
                                                                        var v110 uint = 65535
                                                                        _ = v109
                                                                        _ = v110
                                                                    }
                                                                case "\n\t":
                                                                    if (v79 || ((v80 / ((v80) | 1)) == (^(v80)))) {
                                                                        var v111 uint8 = uint8(v1)
                                                                        v80 = v80
                                                                        var v112 int32 = 0
                                                                        _ = v111
                                                                        _ = v112
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v79 = true
                                                                    }
                                                                default:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v4 + v4)
                                                                    }
                                                                }
                                                            } else {
                                                                v3 -= (v3 >> ((255) & 63))
                                                            }
                                                            switch (fi77 >> ((4294967295) & 63)) {
                                                            case 1:
                                                                if (87 > 128) {
                                                                    if v79 {
                                                                        if (1e10 <= 0.1) {
                                                                            if (((2 > 0) || false) && (3.14159 > -2.25)) {
                                                                                if (760 != 18446744073709551615) {
                                                                                    v79 = false
                                                                                    var v113 uint = 243
                                                                                    v114 := GenS11{F0: 1, F1: (v113 & v113), F2: (2147483647 != 0)}
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                } else {
                                                                                    v4 = v4
                                                                                }
                                                                                v79 = false
                                                                                v3 = v3
                                                                            }
                                                                        } else {
                                                                            if (!(-1 > 1)) {
                                                                                v105 -= v1
                                                                            } else {
                                                                                v115 := GenS8{F0: 127}
                                                                                switch v4 {
                                                                                case "foo bar":
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (v4 + (v4 + (((v4 + "\n\t") + v4) + v4)))
                                                                                    }
                                                                                }
                                                                                var v116 int64 = int64(v3)
                                                                                switch v4 {
                                                                                case "hello":
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = ""
                                                                                    }
                                                                                case "foo bar":
                                                                                    v3 += (^(v3))
                                                                                case "\n\t":
                                                                                    switch fi77 {
                                                                                    case 277:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v116 = -1
                                                                                var v117 uint16 = v105
                                                                                _ = v115
                                                                                _ = v116
                                                                                _ = v117
                                                                            }
                                                                            var v118 uint16 = 255
                                                                            if v79 {
                                                                                var v119 float64 = float64(-2.25)
                                                                                _ = v119
                                                                            }
                                                                            _ = v118
                                                                        }
                                                                        switch fi78 {
                                                                        case 0:
                                                                            v120 := GenS2{F0: (v80 * 100), F1: "foo bar"}
                                                                            _ = v120
                                                                        case -9223372036854775808:
                                                                            v121 := GenS4{F0: uint8(v3), F1: 4294967295, F2: v3}
                                                                            _ = v121
                                                                        case 1:
                                                                            var v122 int8 = v80
                                                                            _ = v122
                                                                        default:
                                                                            v123 := GenS6{F0: 329, F1: 1, F2: -9223372036854775808}
                                                                            _ = v123
                                                                        }
                                                                    } else {
                                                                        var v124 int8 = -9
                                                                        _ = v124
                                                                    }
                                                                    v125 := GenS3{F0: v1, F1: 792}
                                                                    switch (fi77 & (^(fi78))) {
                                                                    case 9223372036854775807:
                                                                        v80 += (v80 >> ((539) & 63))
                                                                    case 32768:
                                                                        switch (fi77 >> ((18446744073709551615) & 63)) {
                                                                        case 1:
                                                                            var v126 float32 = float32(float64((float32((-2.25 - ((0.1 + (((0.0 + (0.1 * 0.0)) - -2.25) / 1.5)) - -2.25))) * (((float32((-2.25 / 1.5)) + (float32(-2.25) / 1.5)) * 0.0) + 3.14159))))
                                                                            _ = v126
                                                                        default:
                                                                            fmt.Println(v4)
                                                                        }
                                                                        v79 = (0 <= -1)
                                                                        v127 := GenS9{F0: 564, F1: (v105 << ((65535) & 63))}
                                                                        _ = v127
                                                                    default:
                                                                        var v128 int32 = -334
                                                                        _ = v128
                                                                    }
                                                                    _ = v125
                                                                } else {
                                                                    var v129 uint = 255
                                                                    var v130 int = fi78
                                                                    v131 := GenS12{F0: "foo bar", F1: ("" + v4), F2: 494}
                                                                    if (v79 || v79) {
                                                                        var v132 int = 2147483647
                                                                        _ = v132
                                                                    } else {
                                                                        v129 = (v129 << ((v129) & 63))
                                                                        var v133 int8 = (-(v80))
                                                                        v134 := GenS12{F0: v4, F1: ((v4 + (v4 + (v4 + v4))) + (("" + v4) + v4)), F2: 42}
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                    _ = v129
                                                                    _ = v130
                                                                    _ = v131
                                                                }
                                                            default:
                                                                if v79 {
                                                                    v79 = v79
                                                                    var v135 int = (fi77 / ((32767) | 1))
                                                                    _ = v135
                                                                }
                                                            }
                                                            switch v4 {
                                                            case "foo bar":
                                                                switch v4 {
                                                                case "a":
                                                                    if v79 {
                                                                        _ = len(v2)
                                                                        if true {
                                                                            var v136 int8 = (v80 & (v80 / (((v80 + 42)) | 1)))
                                                                            v3 = v3
                                                                            if (570 < int16(v136)) {
                                                                                v2 = append(v2, -9223372036854775808)
                                                                            }
                                                                            v137 := GenS13{F0: 2, F1: 43, F2: (v1 >> ((65535) & 63))}
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                    }
                                                                case "":
                                                                    _ = len(v2)
                                                                default:
                                                                    if v79 {
                                                                        var v138 uint32 = uint32(v105)
                                                                        _ = v138
                                                                    }
                                                                }
                                                            case "":
                                                                var v139 int8 = v80
                                                                _ = v139
                                                            }
                                                            var v140 uint = 18446744073709551615
                                                            var v141 uint = v140
                                                            if v79 {
                                                                v142 := GenS14{F0: (v3 & (v3 * 2))}
                                                                switch fi78 {
                                                                case 32768:
                                                                    if (3.14159 <= ((3.14159 + (1e10 + 1.5)) - 0.0)) {
                                                                        v143 := GenS15{F0: -2147483648, F1: (v80 / (((v80 & 42)) | 1))}
                                                                        _ = len(v5)
                                                                        v141 += uint(v1)
                                                                        switch (fi77 >> ((v140) & 63)) {
                                                                        case 796:
                                                                            switch v4 {
                                                                            case "":
                                                                                v141 = 2
                                                                            case "a":
                                                                                if v79 {
                                                                                    if (!v79) {
                                                                                        fmt.Println(v141, fi77, v3)
                                                                                        var v144 string = v4
                                                                                        var v145 uint = v140
                                                                                        _ = v144
                                                                                        _ = v145
                                                                                    }
                                                                                    if true {
                                                                                        if (!(!(v79 && ((((v3 % (((^(v3))) | 1)) < v3) || (9223372036854775807 < -1)) && (v79 && v79))))) {
                                                                                            var v146 float64 = -2.25
                                                                                            _ = v146
                                                                                        } else {
                                                                                            var v147 float64 = 3.14159
                                                                                            _ = v147
                                                                                        }
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        var v148 uint8 = 115
                                                                                        var v149 uint8 = (v148 >> ((601) & 63))
                                                                                        _ = v148
                                                                                        _ = v149
                                                                                    } else {
                                                                                        v140 += v141
                                                                                    }
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = "hello"
                                                                                    }
                                                                                    var v150 float64 = ((1e10 * 1e10) + float64(3.14159))
                                                                                    if (884 == 245) {
                                                                                        switch 2147483647 {
                                                                                        case 4294967295:
                                                                                            fmt.Println(v79, v140, v105)
                                                                                        case -543:
                                                                                            var v151 int8 = -128
                                                                                            v4 = (v4 + "")
                                                                                            v141 -= (v141 + (v141 % (((^(v140))) | 1)))
                                                                                            var v152 int = fi78
                                                                                            _ = v151
                                                                                            _ = v152
                                                                                        case 32767:
                                                                                            var v153 float64 = float64(1.5)
                                                                                            if (!v79) {
                                                                                                var v154 uint64 = uint64(v141)
                                                                                                _ = v154
                                                                                            } else {
                                                                                                v141 = 4294967295
                                                                                                v2 = append(v2, fi77)
                                                                                                var v155 float32 = (float32(v153) + 1.5)
                                                                                                _ = v155
                                                                                            }
                                                                                            var v156 int32 = 1
                                                                                            _ = v153
                                                                                            _ = v156
                                                                                        default:
                                                                                            var v157 string = v4
                                                                                            _ = v157
                                                                                        }
                                                                                        var v158 int32 = -2147483648
                                                                                        var v159 int = 516
                                                                                        fmt.Println(fi77, fi78, v105)
                                                                                        var v160 int64 = 0
                                                                                        var v161 float32 = 3.14159
                                                                                        v1 += v105
                                                                                        v79 = false
                                                                                        var v162 string = v4
                                                                                        _ = v158
                                                                                        _ = v159
                                                                                        _ = v160
                                                                                        _ = v161
                                                                                        _ = v162
                                                                                    }
                                                                                    switch v4 {
                                                                                    case "a":
                                                                                        if v79 {
                                                                                            var v163 uint = v141
                                                                                            v79 = v79
                                                                                            _ = v163
                                                                                        } else {
                                                                                            v164 := GenS16{F0: -668}
                                                                                            switch v4 {
                                                                                            case "":
                                                                                                var v165 uint64 = 303
                                                                                                _ = v165
                                                                                            case "\n\t":
                                                                                                switch (fi78 + (-(fi77))) {
                                                                                                case -9223372036854775808:
                                                                                                    var v166 int64 = 1
                                                                                                    _ = v166
                                                                                                case 127:
                                                                                                    if true {
                                                                                                        v167 := GenS17{F0: v4, F1: ((((-2.25 + 0.1) + 0.0) / 1.5) == v150), F2: v105}
                                                                                                        _ = v167
                                                                                                    }
                                                                                                case 4294967295:
                                                                                                    switch (fi77 ^ fi78) {
                                                                                                    case 4294967295:
                                                                                                        v79 = false
                                                                                                    case -621:
                                                                                                        v80 = v80
                                                                                                    case -899:
                                                                                                        _ = len(v5)
                                                                                                    default:
                                                                                                        if true {
                                                                                                            v168 := GenS3{F0: 591, F1: 0}
                                                                                                            _ = v168
                                                                                                        }
                                                                                                        if (4294967295 >= v140) {
                                                                                                            fmt.Println(v1)
                                                                                                        } else {
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = v4
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                var v169 int8 = v80
                                                                                                                _ = v169
                                                                                                            case "hello":
                                                                                                                var v170 uint16 = v105
                                                                                                                _ = v170
                                                                                                            case "":
                                                                                                                switch "foo bar" {
                                                                                                                case "hello":
                                                                                                                    var v171 uint64 = 1
                                                                                                                    _ = v171
                                                                                                                case "a":
                                                                                                                    if len(v5) > 0 {
                                                                                                                        v5[0] = "\n\t"
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    switch 31 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        v3 = 42
                                                                                                                        switch fi77 {
                                                                                                                        case 32767:
                                                                                                                            v172 := GenS18{F0: -2147483648}
                                                                                                                            _ = v172
                                                                                                                        case 128:
                                                                                                                            if len(v2) > 0 {
                                                                                                                                v2[0] = fi78
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            v2 = append(v2, fi78)
                                                                                                                            var v173 int32 = 2147483647
                                                                                                                            var v174 uint64 = 389
                                                                                                                            _ = v173
                                                                                                                            _ = v174
                                                                                                                        } else {
                                                                                                                            fmt.Println(v150, fi78, v79)
                                                                                                                            v3 = v3
                                                                                                                            v1 = v1
                                                                                                                            v175 := GenS8{F0: 255}
                                                                                                                            _ = v175
                                                                                                                        }
                                                                                                                    case 474:
                                                                                                                        var v176 float64 = (float64(float32(1e10)) * v150)
                                                                                                                        _ = v176
                                                                                                                    case 9223372036854775807:
                                                                                                                        var v177 float64 = (v150 + (v150 / 1.0))
                                                                                                                        _ = v177
                                                                                                                    }
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        v105 -= (v105 | v105)
                                                                                                                    case "a":
                                                                                                                        var v178 uint32 = 1
                                                                                                                        _ = v178
                                                                                                                    }
                                                                                                                    fmt.Println(v4)
                                                                                                                default:
                                                                                                                    var v179 uint32 = uint32(v140)
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                            default:
                                                                                                                var v180 int = (fi77 - (fi78 | fi78))
                                                                                                                _ = v180
                                                                                                            }
                                                                                                            v181 := GenS8{F0: 2}
                                                                                                            v150 = 1e10
                                                                                                            _ = v181
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v164
                                                                                        }
                                                                                    }
                                                                                    _ = v150
                                                                                }
                                                                            case "\n\t":
                                                                                var v182 int64 = -9223372036854775808
                                                                                _ = v182
                                                                            }
                                                                        }
                                                                        var v183 int32 = 42
                                                                        v105 += uint16(v3)
                                                                        _ = v143
                                                                        _ = v183
                                                                    } else {
                                                                        v184 := GenS11{F0: 0, F1: (v141 ^ (v141 + (v141 ^ 11))), F2: v79}
                                                                        v185 := GenS17{F0: v4, F1: (0 >= int64(fi77)), F2: 32767}
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (fi78 & fi78)
                                                                        }
                                                                        _ = v184
                                                                        _ = v185
                                                                    }
                                                                default:
                                                                    if (v79 || ((0.1 == -2.25) && (605 == 1))) {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v142
                                                            }
                                                            var v186 uint = 471
                                                            var v187 uint8 = 168
                                                            var v188 float64 = 1e10
                                                            if v79 {
                                                                if len(v2) > 0 {
                                                                    v2[0] = int(v3)
                                                                }
                                                                v187 = 2
                                                                fmt.Println(v141)
                                                            }
                                                            if ((v4 + "\n\t") == v4) {
                                                                var v189 int = 128
                                                                v1 = (v1 >> ((v140) & 63))
                                                                v5 = append(v5, v4)
                                                                _ = v189
                                                            } else {
                                                                if v79 {
                                                                    if (!v79) {
                                                                        var v190 uint = (^(v186))
                                                                        v140 = (^(v140))
                                                                        var v191 float64 = (v188 + v188)
                                                                        _ = v190
                                                                        _ = v191
                                                                    } else {
                                                                        v187 -= 167
                                                                        v187 -= v187
                                                                        fmt.Println(v80, v3)
                                                                    }
                                                                    if v79 {
                                                                        v2 = append(v2, int(v187))
                                                                    } else {
                                                                        var v192 float64 = v188
                                                                        v4 = "foo bar"
                                                                        _ = v192
                                                                    }
                                                                    if true {
                                                                        if v79 {
                                                                            if false {
                                                                                fmt.Println(v187, v188)
                                                                                if v79 {
                                                                                    var v193 int16 = -715
                                                                                    var v194 string = v4
                                                                                    v2 = append(v2, (fi78 - 32767))
                                                                                    _ = v193
                                                                                    _ = v194
                                                                                } else {
                                                                                    if (v79 && false) {
                                                                                        fmt.Println(v187)
                                                                                        var v195 int32 = 601
                                                                                        var v196 bool = (0 < 171)
                                                                                        _ = v195
                                                                                        _ = v196
                                                                                    } else {
                                                                                        var v197 bool = (-412 != 9223372036854775807)
                                                                                        _ = v197
                                                                                    }
                                                                                    v198 := GenS12{F0: v4, F1: (v4 + (((v4 + (("foo bar" + "") + v4)) + "\n\t") + v4)), F2: 70}
                                                                                    v105 = 1
                                                                                    _ = v198
                                                                                }
                                                                                var v199 int16 = -430
                                                                                _ = v199
                                                                            } else {
                                                                                if v79 {
                                                                                    v4 = v4
                                                                                    var v200 int = (fi77 << ((v186) & 63))
                                                                                    switch (fi78 * (fi78 / ((v200) | 1))) {
                                                                                    case -9223372036854775808:
                                                                                        switch (fi78 / ((42) | 1)) {
                                                                                        case -93:
                                                                                            v3 += 0
                                                                                        case 32767:
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                    case 32767:
                                                                                        var v201 float32 = (-2.25 + 0.0)
                                                                                        _ = v201
                                                                                    case 2147483647:
                                                                                        if (v79 || (v141 != 2)) {
                                                                                            v141 -= v140
                                                                                        } else {
                                                                                            v202 := GenS5{F0: -564, F1: (v80 % (((^(v80))) | 1))}
                                                                                            fmt.Println(v4, v188)
                                                                                            if false {
                                                                                                var v203 int8 = v80
                                                                                                v2 = append(v2, 42)
                                                                                                _ = v203
                                                                                            }
                                                                                            _ = v202
                                                                                        }
                                                                                    }
                                                                                    _ = v200
                                                                                } else {
                                                                                    v204 := GenS19{F0: (v80 * int8(v3)), F1: (^(v105)), F2: 42}
                                                                                    v205 := GenS6{F0: (fi77 ^ 9223372036854775807), F1: (v187 << ((v141) & 63)), F2: 1}
                                                                                    v206 := GenS11{F0: (^(v187)), F1: (v140 + v141), F2: ((849 >= 42) || (!v79))}
                                                                                    _ = v204
                                                                                    _ = v205
                                                                                    _ = v206
                                                                                }
                                                                            }
                                                                            var v207 uint = uint(v3)
                                                                            _ = v207
                                                                        } else {
                                                                            v141 += uint(v187)
                                                                            v80 -= (v80 << ((710) & 63))
                                                                            v1 = 936
                                                                            if true {
                                                                                if v79 {
                                                                                    v208 := GenS9{F0: 1, F1: (v105 - (v105 + (v1 >> (((v140 >> (((v140 * v186)) & 63))) & 63))))}
                                                                                    switch v4 {
                                                                                    case "hello":
                                                                                        var v209 int16 = 32767
                                                                                        _ = v209
                                                                                    case "a":
                                                                                        v105 = 177
                                                                                    default:
                                                                                        v4 = v4
                                                                                    }
                                                                                    v187 -= 42
                                                                                    var v210 uint8 = 0
                                                                                    _ = v208
                                                                                    _ = v210
                                                                                }
                                                                            }
                                                                            v211 := GenS20{F0: v3, F1: -163}
                                                                            _ = v211
                                                                        }
                                                                        v212 := GenS17{F0: v4, F1: (!false), F2: 443}
                                                                        v105 -= 255
                                                                        _ = v212
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                if (v188 <= (1.5 - 0.0)) {
                                                                    var v213 float32 = 3.14159
                                                                    var v214 string = ""
                                                                    var v215 int64 = -566
                                                                    var v216 uint32 = 65535
                                                                    _ = v213
                                                                    _ = v214
                                                                    _ = v215
                                                                    _ = v216
                                                                }
                                                                switch fi78 {
                                                                case 724:
                                                                    switch (v4 + v4) {
                                                                    case "foo bar":
                                                                        var v217 int64 = -800
                                                                        _ = v217
                                                                    case "\n\t":
                                                                        var v218 bool = false
                                                                        v80 = (-(v80))
                                                                        _ = v218
                                                                    case "hello":
                                                                        var v219 uint16 = v1
                                                                        _ = v219
                                                                    default:
                                                                        var v220 uint64 = 18446744073709551615
                                                                        _ = v220
                                                                    }
                                                                }
                                                                var v221 float32 = 0.1
                                                                var v222 int16 = -358
                                                                _ = v221
                                                                _ = v222
                                                            }
                                                            _ = v140
                                                            _ = v141
                                                            _ = v186
                                                            _ = v187
                                                            _ = v188
                                                        } else {
                                                            v223 := GenS21{F0: 12}
                                                            _ = v223
                                                        }
                                                    }
                                                    switch v4 {
                                                    case "foo bar":
                                                        _ = len(v5)
                                                    default:
                                                        fmt.Println(v1, fi77, fi77)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v105
                                    default:
                                        var v224 int64 = 1
                                        _ = v224
                                    }
                                    v225 := GenS22{F0: 9223372036854775807}
                                    _ = v225
                                }
                            }
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                        var v226 uint8 = 113
                        for fi227 := 0; fi227 < 1; fi227++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi227
                        }
                        _ = v226
                    } else {
                        for fi228 := 0; fi228 < 3; fi228++ {
                            switch ("\n\t" + v4) {
                            case "hello":
                                fmt.Println(v3, v3)
                            case "":
                                v4 = "foo bar"
                            case "foo bar":
                                var v229 uint = 42
                                _ = v229
                            }
                            fmt.Println(fi228, v3, fi228)
                            _ = fi228
                        }
                        for fi230 := 0; fi230 < 16; fi230++ {
                            v231 := GenS23{F0: (v1 % ((65535) | 1))}
                            if (!false) {
                                if (0.1 != 0.0) {
                                    if (!(0 < 32767)) {
                                        var v232 int32 = -460
                                        var v233 int8 = 0
                                        switch 42 {
                                        case 127:
                                            if ("\n\t" > "foo bar") {
                                                v234 := GenS24{F0: -800, F1: "hello", F2: 0}
                                                switch (fi230 | fi77) {
                                                case 42:
                                                    if false {
                                                        fmt.Println(v4, v1)
                                                        v4 = v4
                                                    }
                                                case 9223372036854775807:
                                                    if false {
                                                        var v235 int16 = 32767
                                                        v233 = (v233 ^ v233)
                                                        if ((!true) || true) {
                                                            v4 = ((("foo bar" + (v4 + v4)) + v4) + ("a" + v4))
                                                        }
                                                        _ = v235
                                                    }
                                                    if (!(2147483647 < (^(fi77)))) {
                                                        v4 = (((v4 + "hello") + (("a" + v4) + v4)) + v4)
                                                        v236 := GenS25{F0: 0, F1: 0, F2: (v233 << ((2) & 63))}
                                                        v3 += (v3 * (^(v3)))
                                                        v1 -= 65535
                                                        _ = v236
                                                    }
                                                    v4 = (v4 + "a")
                                                default:
                                                    if (true || false) {
                                                        v237 := GenS10{F0: 845}
                                                        _ = v237
                                                    }
                                                }
                                                v4 = ""
                                                _ = v234
                                            } else {
                                                var v238 uint16 = (v1 * (v1 * 0))
                                                _ = v238
                                            }
                                        case 42:
                                            v239 := GenS11{F0: 153, F1: 776, F2: true}
                                            _ = v239
                                        default:
                                            if (v1 == (v1 + v1)) {
                                                var v240 int16 = 42
                                                v4 = "\n\t"
                                                switch (v4 + "hello") {
                                                case "a":
                                                    v233 += v233
                                                    v241 := GenS26{F0: int16(fi77)}
                                                    if len(v5) > 0 {
                                                        v5[0] = ""
                                                    }
                                                    if true {
                                                        v242 := GenS27{F0: (fi77 & (fi230 + fi77)), F1: v1, F2: (-(v233))}
                                                        _ = v242
                                                    }
                                                    _ = v241
                                                case "foo bar":
                                                    v240 = (v240 % ((v240) | 1))
                                                case "\n\t":
                                                    if len(v2) > 0 {
                                                        v2[0] = (-(fi230))
                                                    }
                                                default:
                                                    v240 = v240
                                                }
                                                _ = v240
                                            }
                                        }
                                        fmt.Println(v1, v3, v1)
                                        _ = v232
                                        _ = v233
                                    }
                                    if (fi230 > 128) {
                                        if true {
                                            v4 = ""
                                        } else {
                                            switch (fi77 >> ((591) & 63)) {
                                            case 640:
                                                if false {
                                                    v1 -= (v1 / ((v1) | 1))
                                                    var v243 uint8 = 127
                                                    _ = v243
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = fi230
                                                    }
                                                    v4 = "\n\t"
                                                }
                                            case -319:
                                                switch (fi77 ^ (fi77 - fi230)) {
                                                case 674:
                                                    if ((v1 << ((2) & 63)) <= (^(v1))) {
                                                        fmt.Println(fi230)
                                                        switch ("foo bar" + ("a" + "foo bar")) {
                                                        case "\n\t":
                                                            switch (v4 + ("a" + (("hello" + v4) + ((v4 + "\n\t") + "foo bar")))) {
                                                            case "hello":
                                                                if true {
                                                                    v244 := GenS2{F0: 0, F1: ""}
                                                                    _ = v244
                                                                }
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            if true {
                                                                if ((1e10 + 0.0) > 1e10) {
                                                                    v5 = append(v5, (("hello" + v4) + v4))
                                                                    _ = len(v5)
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v4 = "\n\t"
                                                                var v245 string = ("hello" + ((v4 + "a") + "foo bar"))
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi77
                                                                }
                                                                _ = v245
                                                            }
                                                            v2 = append(v2, (fi77 | 32768))
                                                        default:
                                                            v246 := GenS28{F0: 65535}
                                                            _ = v246
                                                        }
                                                        v247 := GenS28{F0: uint(v3)}
                                                        _ = v247
                                                    }
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            case 32767:
                                                v248 := GenS29{F0: false, F1: int8(v1), F2: 2}
                                                _ = v248
                                            default:
                                                v249 := GenS30{F0: 1, F1: ((1e10 <= 1.5) || (87 > (v1 + 573))), F2: 500}
                                                _ = v249
                                            }
                                            if true {
                                                var v250 float32 = float32(0.1)
                                                _ = v250
                                            }
                                            if true {
                                                v3 = v3
                                                switch v4 {
                                                case "hello":
                                                    v251 := GenS22{F0: 1}
                                                    _ = v251
                                                default:
                                                    v1 += 534
                                                }
                                                v4 = ((v4 + (v4 + "hello")) + v4)
                                                v252 := GenS18{F0: -1}
                                                _ = v252
                                            }
                                        }
                                        v3 -= 42
                                        if (!(0 < 608)) {
                                            v253 := GenS31{F0: uint32(v1), F1: 127}
                                            _ = v253
                                        } else {
                                            v4 = v4
                                        }
                                    }
                                    _ = len(v2)
                                }
                            } else {
                                var v254 bool = (v4 != "\n\t")
                                _ = v254
                            }
                            v255 := GenS32{F0: 1, F1: fi77, F2: v4}
                            _ = fi230
                            _ = v231
                            _ = v255
                        }
                    }
                    v4 = "a"
                    _ = fi77
                }
                _ = len(v5)
                v256 := GenS33{F0: v3, F1: 127, F2: 1}
                _ = v256
            }
        } else {
            var v257 uint8 = uint8(v3)
            _ = v257
        }
        v258 := GenS10{F0: 4294967295}
        _ = v258
    } else {
        for fi259 := 0; fi259 < 3; fi259++ {
            for fi260 := 0; fi260 < 16; fi260++ {
                v261 := GenS34{F0: 255, F1: (-32768 < 42)}
                var v262 uint8 = uint8(v3)
                _ = fi260
                _ = v261
                _ = v262
            }
            for fi263 := 0; fi263 < 3; fi263++ {
                fmt.Println(v4)
                _ = fi263
            }
            if false {
                var v264 float32 = 1.5
                v4 = "hello"
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                _ = v264
            }
            _ = fi259
        }
        var v265 uint32 = (v3 % ((v3) | 1))
        _ = v265
    }
    if false {
        var v266 uint64 = 681
        for fi267 := 0; fi267 < 9; fi267++ {
            v268 := GenS7{F0: 313, F1: 42, F2: 127}
            _ = fi267
            _ = v268
        }
        v269 := GenS35{F0: 1, F1: 1}
        _ = v266
        _ = v269
    }
    if (948 < 1) {
        for fi270 := 0; fi270 < 20; fi270++ {
            var v271 uint = 18446744073709551615
            fmt.Println(fi270, v271)
            var v272 int64 = -9223372036854775808
            _ = fi270
            _ = v271
            _ = v272
        }
        v1 = v1
        var v273 string = "foo bar"
        _ = v273
    }
    if (2 != 42) {
        if (!true) {
            fmt.Println(v3, v4)
            var v274 int = 1
            _ = v274
        } else {
            fmt.Println(v1, v4)
            var v275 uint64 = 18446744073709551615
            v276 := GenS30{F0: 84, F1: ((v4 + "hello") <= v4), F2: 255}
            _ = v275
            _ = v276
        }
    }
    switch (("hello" + "hello") + v4) {
    case "hello":
        v277 := GenS36{F0: 17, F1: 127}
        _ = v277
    case "foo bar":
        if true {
            for fi278 := 0; fi278 < 7; fi278++ {
                for fi279 := 0; fi279 < 8; fi279++ {
                    v3 -= uint32(v1)
                    var v280 uint8 = 42
                    var v281 string = "foo bar"
                    var v282 float64 = 0.1
                    _ = fi279
                    _ = v280
                    _ = v281
                    _ = v282
                }
                var v283 uint32 = (v3 + uint32(v1))
                var v284 float32 = -2.25
                _ = fi278
                _ = v283
                _ = v284
            }
        } else {
            fmt.Println(v3)
            var v285 int8 = -110
            v4 = v4
            _ = v285
        }
    default:
        var v286 uint16 = v1
        _ = v286
    }
    if false {
        var v287 uint64 = 1
        for fi288 := 0; fi288 < 17; fi288++ {
            switch "\n\t" {
            case "\n\t":
                for fi289 := 0; fi289 < 6; fi289++ {
                    var v290 float64 = -2.25
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi289
                    _ = v290
                }
            case "a":
                var v291 uint = 789
                _ = v291
            case "foo bar":
                var v292 float32 = 1.5
                _ = v292
            }
            _ = fi288
        }
        if false {
            v1 -= (v1 - (^(v1)))
            fmt.Println(v1, v4)
        }
        _ = v287
    } else {
        switch int(v1) {
        case 32767:
            var v293 int64 = -295
            _ = v293
        case 769:
            var v294 uint64 = 330
            _ = v294
        case 665:
            var v295 uint = 1
            _ = v295
        default:
            fmt.Println(v3, v4)
        }
        fmt.Println(v4)
        for fi296 := 0; fi296 < 11; fi296++ {
            if false {
                v297 := GenS29{F0: (526 > 2), F1: 100, F2: 18446744073709551615}
                for fi298 := 0; fi298 < 12; fi298++ {
                    v4 = v4
                    _ = fi298
                }
                var v299 uint32 = v3
                _ = v297
                _ = v299
            } else {
                fmt.Println(v1, v1)
            }
            if len(v5) > 0 {
                v5[0] = ("" + v4)
            }
            _ = fi296
        }
    }
    for fi300 := 0; fi300 < 20; fi300++ {
        if (!true) {
            v301 := GenS25{F0: 4294967295, F1: 42, F2: -11}
            _ = v301
        }
        if ((v1 | v1) != v1) {
            if false {
                v3 = (^(v3))
            }
            v1 = 737
        } else {
            v4 = v4
            if len(v2) > 0 {
                v2[0] = 609
            }
            v4 = ((v4 + v4) + v4)
        }
        v302 := GenS21{F0: 783}
        _ = fi300
        _ = v302
    }
    switch 678 {
    case 32768:
        for fi303 := 0; fi303 < 11; fi303++ {
            var v304 int8 = 11
            v305 := GenS19{F0: 127, F1: (v1 % ((65535) | 1)), F2: 0}
            _ = fi303
            _ = v304
            _ = v305
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
