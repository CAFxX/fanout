package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint32; F1 int16; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS9 struct { F0 bool; F1 int64; F2 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64; F1 int32 }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint64; F1 bool }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS17 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS18 struct { F0 int32; F1 int8; F2 bool }
type GenS19 struct { F0 int64; F1 uint16 }
func genFn0() uint64 {
    var v4 uint32 = 1
    if true {
        fmt.Println(v4)
        v4 += (v4 & v4)
    } else {
        v5 := GenS1{F0: 822, F1: 255, F2: -34}
        var v6 int8 = 0
        _ = v5
        _ = v6
    }
    if (true || false) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint16 = uint16(v4)
            v9 := GenS2{F0: 32767, F1: 233, F2: (fi7 % (((^(fi7))) | 1))}
            fmt.Println(v8, fi7)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = v4
    return 138
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{1, 969}
    var v3 uint32 = 42
    switch -9223372036854775808 {
    case 1:
        if false {
            if (v3 >= 1) {
                if true {
                    v1 += 0
                }
                fmt.Println(v1)
            } else {
                v2 = append(v2, (v3 & (v3 ^ v3)))
            }
            var v10 float64 = -2.25
            _ = v10
        }
    case 0:
        var v11 uint32 = v3
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    v3 -= v3
    if false {
        v12 := GenS1{F0: 2147483647, F1: 63, F2: (v1 * v1)}
        _ = v12
    } else {
        if true {
            fmt.Println(v3, v1)
            switch "foo bar" {
            case "hello":
                v1 += v1
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v13 int = 32767
            v1 = -1
            _ = v13
        }
        var v14 int64 = -1
        _ = v14
    }
    if true {
        v1 += (v1 << ((0) & 63))
        v3 = (v3 >> ((4294967295) & 63))
        if ((v3 % ((v3) | 1)) != (v3 % (((v3 ^ (v3 + v3))) | 1))) {
            var v15 float32 = float32((((1e10 * -2.25) * 0.1) / (1.5 + (0.0 + (1e10 / -2.25)))))
            _ = v15
        } else {
            v16 := GenS1{F0: 127, F1: 934, F2: v1}
            _ = v16
        }
    }
    v17 := GenS2{F0: v1, F1: 42, F2: int(v1)}
    switch 0 {
    case 128:
        if false {
            fmt.Println(v3, v3)
            var v18 uint = 159
            _ = v18
        } else {
            var v19 uint32 = (v3 | 65535)
            if len(v2) > 0 {
                v2[0] = 968
            }
            _ = v19
        }
    }
    v20 := GenS3{F0: -9223372036854775808, F1: "a", F2: 380}
    var v21 string = (("a" + (("foo bar" + (("" + "hello") + "hello")) + ("foo bar" + "a"))) + "")
    for fi22 := 0; fi22 < 10; fi22++ {
        v1 = v1
        _ = fi22
    }
    if (!(!false)) {
        for fi23 := 0; fi23 < 7; fi23++ {
            v3 = (^(v3))
            _ = fi23
        }
    }
    switch 127 {
    case 0:
        v21 = ("\n\t" + (v21 + v21))
    case -623:
        if ((v1 << ((1) & 63)) > 148) {
            v24 := GenS4{F0: -75}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                switch "hello" {
                case "foo bar":
                    var v26 uint64 = 42
                    _ = v26
                case "hello":
                    var v27 float64 = 3.14159
                    _ = v27
                }
                v28 := GenS4{F0: 0}
                _ = fi25
                _ = v28
            }
        }
    case 740:
        if (255 != 2) {
            switch v21 {
            case "":
                for fi29 := 0; fi29 < 19; fi29++ {
                    if true {
                        fmt.Println(v1, v1, fi29)
                        if (("foo bar" != ("a" + v21)) && (42 > 65535)) {
                            var v30 int8 = -128
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v30
                        } else {
                            var v31 float64 = 1e10
                            if false {
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    v21 = ("\n\t" + "foo bar")
                                    _ = fi32
                                }
                                if false {
                                    var v33 int8 = int8(v3)
                                    var v34 uint16 = 0
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 int32 = 1
                                    _ = v35
                                }
                            } else {
                                var v36 string = "\n\t"
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    var v38 int = fi37
                                    _ = fi37
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 17; fi39++ {
                                    if true {
                                        if false {
                                            var v40 uint16 = 255
                                            if ((v1 - v1) != (v1 - v1)) {
                                                v31 += v31
                                                fmt.Println(v3, v3)
                                            } else {
                                                var v41 uint64 = 1
                                                if false {
                                                    if false {
                                                        if (!((!true) || ((v40 != v40) && false))) {
                                                            v42 := GenS3{F0: 1, F1: ("" + v36), F2: (v40 ^ v40)}
                                                            _ = v42
                                                        }
                                                        switch fi39 {
                                                        case 884:
                                                            if (0.1 == -2.25) {
                                                                switch fi39 {
                                                                case 618:
                                                                    if (v3 >= (^(v3))) {
                                                                        var v43 int32 = -2147483648
                                                                        if true {
                                                                            v21 = ("" + ("" + ((((v21 + v21) + v21) + "a") + "")))
                                                                            v2 = append(v2, (v3 / ((v3) | 1)))
                                                                            if (false && (v43 == v43)) {
                                                                                switch ((((v36 + "hello") + v21) + (v21 + "a")) + v36) {
                                                                                case "hello":
                                                                                    if (v43 == v43) {
                                                                                        fmt.Println(v40, fi29)
                                                                                        var v44 int64 = -973
                                                                                        fmt.Println(v44)
                                                                                        var v45 int8 = int8(v3)
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                    } else {
                                                                                        v2 = append(v2, 0)
                                                                                        if (true && true) {
                                                                                            if (1 < -128) {
                                                                                                v46 := GenS5{F0: (v41 * 18446744073709551615), F1: -166}
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 4294967295
                                                                                                }
                                                                                                fmt.Println(fi29)
                                                                                                if true {
                                                                                                    if (false && (!(-1 < -126))) {
                                                                                                        var v47 int32 = 1
                                                                                                        var v48 uint64 = genFn0()
                                                                                                        fmt.Println(v21, fi39)
                                                                                                        var v49 float32 = 0.0
                                                                                                        _ = v47
                                                                                                        _ = v48
                                                                                                        _ = v49
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v40 = (v40 / (((^(v40))) | 1))
                                                                                                        } else {
                                                                                                            v3 = v3
                                                                                                            if (42 > 42) {
                                                                                                                switch ((v36 + "hello") + ((v36 + "") + "")) {
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(v3, v40, fi29)
                                                                                                                case "":
                                                                                                                    v50 := GenS4{F0: -9223372036854775808}
                                                                                                                    _ = v50
                                                                                                                case "\n\t":
                                                                                                                    v51 := GenS3{F0: 1, F1: "a", F2: (^(v40))}
                                                                                                                    _ = v51
                                                                                                                }
                                                                                                                v52 := GenS6{F0: 895, F1: 42, F2: (v1 ^ (v1 << ((473) & 63)))}
                                                                                                                _ = v52
                                                                                                            }
                                                                                                            switch v21 {
                                                                                                            case "\n\t":
                                                                                                                switch fi39 {
                                                                                                                case -2147483648:
                                                                                                                    if false {
                                                                                                                        if true {
                                                                                                                            v53 := GenS7{F0: (v3 * 4294967295), F1: (-(v1)), F2: v1}
                                                                                                                            _ = v53
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 522:
                                                                                                                    fmt.Println(v36, v21, v41)
                                                                                                                default:
                                                                                                                    v54 := GenS8{F0: (^(v40)), F1: 18446744073709551615, F2: (v40 << ((1) & 63))}
                                                                                                                    _ = v54
                                                                                                                }
                                                                                                            case "":
                                                                                                                if false {
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = (v3 >> ((1) & 63))
                                                                                                                    }
                                                                                                                    v41 += (v41 << ((42) & 63))
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                var v55 uint32 = v3
                                                                                                                _ = v55
                                                                                                            }
                                                                                                            switch "foo bar" {
                                                                                                            case "foo bar":
                                                                                                                if false {
                                                                                                                    _ = len(v2)
                                                                                                                    v56 := GenS3{F0: 9223372036854775807, F1: "a", F2: v40}
                                                                                                                    if ((genFn0() != (v41 ^ 2)) && true) {
                                                                                                                        var v57 float32 = (0.0 - 0.0)
                                                                                                                        _ = v57
                                                                                                                    }
                                                                                                                    _ = v56
                                                                                                                } else {
                                                                                                                    switch (fi39 >> ((65535) & 63)) {
                                                                                                                    case -9223372036854775808:
                                                                                                                        if true {
                                                                                                                            switch (v36 + "") {
                                                                                                                            case "hello":
                                                                                                                                v40 = (v40 * 1)
                                                                                                                            case "a":
                                                                                                                                var v58 int32 = -2147483648
                                                                                                                                _ = v58
                                                                                                                            default:
                                                                                                                                if true {
                                                                                                                                    switch 2147483647 {
                                                                                                                                    case 0:
                                                                                                                                        v59 := GenS3{F0: 0, F1: v36, F2: (^(v40))}
                                                                                                                                        _ = v59
                                                                                                                                    default:
                                                                                                                                        v60 := GenS4{F0: 1}
                                                                                                                                        _ = v60
                                                                                                                                    }
                                                                                                                                    if len(v2) > 0 {
                                                                                                                                        v2[0] = (v3 - uint32(v40))
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    v43 = v43
                                                                                                                                    if false {
                                                                                                                                        var v61 uint16 = 42
                                                                                                                                        if false {
                                                                                                                                            var v62 uint64 = (^(v41))
                                                                                                                                            v41 = v41
                                                                                                                                            _ = v62
                                                                                                                                        }
                                                                                                                                        _ = v61
                                                                                                                                    } else {
                                                                                                                                        switch fi29 {
                                                                                                                                        case 9223372036854775807:
                                                                                                                                            v63 := GenS2{F0: v1, F1: 255, F2: -2147483648}
                                                                                                                                            _ = v63
                                                                                                                                        }
                                                                                                                                        fmt.Println(v41, fi39, fi39)
                                                                                                                                        if len(v2) > 0 {
                                                                                                                                            v2[0] = 716
                                                                                                                                        }
                                                                                                                                        v2 = append(v2, 1)
                                                                                                                                    }
                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v64 int8 = -88
                                                                                                                            var v65 uint8 = 29
                                                                                                                            _ = v64
                                                                                                                            _ = v65
                                                                                                                        } else {
                                                                                                                            if true {
                                                                                                                                if (-58 == (fi29 >> ((18446744073709551615) & 63))) {
                                                                                                                                    var v66 float64 = v31
                                                                                                                                    _ = v66
                                                                                                                                } else {
                                                                                                                                    v41 -= genFn0()
                                                                                                                                }
                                                                                                                                switch ("a" + "hello") {
                                                                                                                                case "foo bar":
                                                                                                                                    _ = len(v2)
                                                                                                                                case "":
                                                                                                                                    var v67 uint16 = uint16(v3)
                                                                                                                                    _ = v67
                                                                                                                                case "a":
                                                                                                                                    v36 = v21
                                                                                                                                default:
                                                                                                                                    v43 -= v43
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v40 = (v40 % ((v40) | 1))
                                                                                                                            v36 = ("a" + "\n\t")
                                                                                                                            fmt.Println(v36, v41)
                                                                                                                            v40 += 448
                                                                                                                        }
                                                                                                                    case 391:
                                                                                                                        var v68 uint64 = v41
                                                                                                                        _ = v68
                                                                                                                    default:
                                                                                                                        v2 = append(v2, 2147483647)
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(fi39, v36, fi39)
                                                                                                            case "":
                                                                                                                var v69 int = (fi39 / (((fi39 % ((fi29) | 1))) | 1))
                                                                                                                _ = v69
                                                                                                            case "a":
                                                                                                                var v70 bool = (true || (517 <= 2))
                                                                                                                _ = v70
                                                                                                            default:
                                                                                                                var v71 uint32 = v3
                                                                                                                _ = v71
                                                                                                            }
                                                                                                            var v72 uint8 = 153
                                                                                                            v41 += v41
                                                                                                            switch (v36 + v36) {
                                                                                                            case "":
                                                                                                                fmt.Println(v1, v43, v43)
                                                                                                            default:
                                                                                                                v73 := GenS9{F0: (!(((!false) && (!(uint64(v72) > (v41 & 0)))) || false)), F1: 42, F2: 100}
                                                                                                                v40 += (^(v40))
                                                                                                                v2 = append(v2, 285)
                                                                                                                fmt.Println(v31)
                                                                                                                v74 := GenS10{F0: 2147483647}
                                                                                                                _ = v73
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            switch ((v21 + "\n\t") + ("foo bar" + ("foo bar" + v21))) {
                                                                                                            case "foo bar":
                                                                                                                switch fi29 {
                                                                                                                case 1:
                                                                                                                    v72 -= uint8(v40)
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                if false {
                                                                                                                    var v75 uint = 18446744073709551615
                                                                                                                    v76 := GenS6{F0: v75, F1: (v72 + 127), F2: v1}
                                                                                                                    _ = v75
                                                                                                                    _ = v76
                                                                                                                } else {
                                                                                                                    v77 := GenS11{F0: (^(v41)), F1: 2147483647}
                                                                                                                    v2 = append(v2, 0)
                                                                                                                    if ((v41 * v41) > (^(v41))) {
                                                                                                                        v78 := GenS2{F0: -156, F1: uint8(v3), F2: fi29}
                                                                                                                        var v79 string = (v36 + (v36 + ""))
                                                                                                                        _ = v78
                                                                                                                        _ = v79
                                                                                                                    } else {
                                                                                                                        switch fi39 {
                                                                                                                        case 127:
                                                                                                                            var v80 bool = ((v3 % ((2147483647) | 1)) >= (v3 - 65535))
                                                                                                                            _ = v80
                                                                                                                        case 4294967295:
                                                                                                                            v21 = v21
                                                                                                                        case 32767:
                                                                                                                            v31 -= ((((-2.25 + 1.5) - v31) + ((1e10 * float64((0.1 + -2.25))) - 0.1)) * (v31 * (-2.25 * (v31 + 3.14159))))
                                                                                                                        default:
                                                                                                                            v81 := GenS6{F0: 0, F1: v72, F2: -635}
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        v21 = "hello"
                                                                                                                        var v82 bool = false
                                                                                                                        var v83 string = "foo bar"
                                                                                                                        _ = v82
                                                                                                                        _ = v83
                                                                                                                    }
                                                                                                                    _ = v77
                                                                                                                }
                                                                                                            }
                                                                                                            switch fi29 {
                                                                                                            case 42:
                                                                                                                switch v36 {
                                                                                                                case "\n\t":
                                                                                                                    switch fi39 {
                                                                                                                    case 159:
                                                                                                                        var v84 uint = 42
                                                                                                                        _ = v84
                                                                                                                    case -9223372036854775808:
                                                                                                                        if (false && ((!true) && false)) {
                                                                                                                            var v85 uint = 0
                                                                                                                            _ = v85
                                                                                                                        } else {
                                                                                                                            fmt.Println(v31, fi29, v1)
                                                                                                                            v21 = (v36 + v21)
                                                                                                                            if (v43 >= v43) {
                                                                                                                                var v86 float32 = 0.0
                                                                                                                                fmt.Println(v1, v36, v43)
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v41, v3, v72)
                                                                                                                                    switch fi29 {
                                                                                                                                    case 4294967295:
                                                                                                                                        fmt.Println(fi39, v21, v43)
                                                                                                                                    default:
                                                                                                                                        switch (fi29 << ((879) & 63)) {
                                                                                                                                        case -144:
                                                                                                                                            if len(v2) > 0 {
                                                                                                                                                v2[0] = (v3 / ((2147483648) | 1))
                                                                                                                                            }
                                                                                                                                            if (629 < int64(v72)) {
                                                                                                                                                var v87 bool = (float32(float64(1.5)) > (v86 - v86))
                                                                                                                                                var v88 uint64 = v41
                                                                                                                                                _ = len(v2)
                                                                                                                                                v2 = append(v2, 2147483647)
                                                                                                                                                _ = v87
                                                                                                                                                _ = v88
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(v40, v40, v40)
                                                                                                                                                if false {
                                                                                                                                                    var v89 uint16 = v40
                                                                                                                                                    v90 := GenS9{F0: true, F1: 839, F2: 1}
                                                                                                                                                    var v91 float64 = v31
                                                                                                                                                    _ = v89
                                                                                                                                                    _ = v90
                                                                                                                                                    _ = v91
                                                                                                                                                } else {
                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                    }
                                                                                                                                                    v92 := GenS12{F0: 42, F1: "a"}
                                                                                                                                                    if (false && (v1 < (v1 - v1))) {
                                                                                                                                                        if (v86 != 0.1) {
                                                                                                                                                            fmt.Println(fi29, v31)
                                                                                                                                                            switch (-(fi29)) {
                                                                                                                                                            case -9223372036854775808:
                                                                                                                                                                v43 += (v43 & -2147483648)
                                                                                                                                                            default:
                                                                                                                                                                v43 -= v43
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        v93 := GenS9{F0: false, F1: 498, F2: 42}
                                                                                                                                                        if (fi39 == fi39) {
                                                                                                                                                            switch ((v21 + "a") + "hello") {
                                                                                                                                                            case "hello":
                                                                                                                                                                fmt.Println(fi39, v43)
                                                                                                                                                            case "foo bar":
                                                                                                                                                                var v94 int8 = 1
                                                                                                                                                                _ = v94
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v93
                                                                                                                                                    }
                                                                                                                                                    _ = v92
                                                                                                                                                }
                                                                                                                                                switch fi29 {
                                                                                                                                                case -731:
                                                                                                                                                    v95 := GenS13{F0: 4294967295, F1: true}
                                                                                                                                                    _ = v95
                                                                                                                                                case 797:
                                                                                                                                                    fmt.Println(v31)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v3 -= 1
                                                                                                                                        case 2147483647:
                                                                                                                                            fmt.Println(v1, v21, v43)
                                                                                                                                        case 1:
                                                                                                                                            v72 -= v72
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v31, fi39)
                                                                                                                                }
                                                                                                                                _ = v86
                                                                                                                            } else {
                                                                                                                                fmt.Println(v31, v21, v3)
                                                                                                                                v96 := GenS4{F0: -1}
                                                                                                                                _ = v96
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch (fi39 ^ (-(fi39))) {
                                                                                                                        case 1:
                                                                                                                            var v97 uint32 = 65535
                                                                                                                            _ = v97
                                                                                                                        case -1:
                                                                                                                            if true {
                                                                                                                                switch (-(fi29)) {
                                                                                                                                case 1:
                                                                                                                                    v40 -= 1
                                                                                                                                case 4294967295:
                                                                                                                                    v98 := GenS7{F0: (v3 + v3), F1: v1, F2: (v1 | (v1 | v1))}
                                                                                                                                    _ = v98
                                                                                                                                }
                                                                                                                                fmt.Println(v3, fi29, v36)
                                                                                                                                if false {
                                                                                                                                    v99 := GenS7{F0: (v3 >> ((65535) & 63)), F1: int16(v41), F2: v1}
                                                                                                                                    _ = v99
                                                                                                                                } else {
                                                                                                                                    var v100 uint32 = (v3 % ((v3) | 1))
                                                                                                                                    _ = v100
                                                                                                                                }
                                                                                                                                v36 = v36
                                                                                                                                var v101 uint16 = (v40 - uint16(v72))
                                                                                                                                if false {
                                                                                                                                    v40 = (v101 >> ((18446744073709551615) & 63))
                                                                                                                                }
                                                                                                                                _ = v101
                                                                                                                            }
                                                                                                                            v102 := GenS8{F0: 42, F1: (v41 | (v41 ^ uint64(v3))), F2: (^(v40))}
                                                                                                                            _ = v102
                                                                                                                        case 0:
                                                                                                                            var v103 int16 = v1
                                                                                                                            _ = v103
                                                                                                                        default:
                                                                                                                            if (42 >= 127) {
                                                                                                                                if len(v2) > 0 {
                                                                                                                                    v2[0] = (v3 / ((v3) | 1))
                                                                                                                                }
                                                                                                                                v40 = v40
                                                                                                                                v104 := GenS1{F0: (fi39 % (((-(fi29))) | 1)), F1: 255, F2: (v1 * int16(v41))}
                                                                                                                                _ = v104
                                                                                                                            } else {
                                                                                                                                v1 -= (v1 + -982)
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    switch 0 {
                                                                                                                    case 4294967295:
                                                                                                                        if true {
                                                                                                                            if (v43 == (v43 * -2147483648)) {
                                                                                                                                if len(v2) >= 2 {
                                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                                }
                                                                                                                                v41 = (^(v41))
                                                                                                                            }
                                                                                                                            var v105 int64 = -1
                                                                                                                            fmt.Println(v40, v72, v31)
                                                                                                                            _ = v105
                                                                                                                        }
                                                                                                                    case 42:
                                                                                                                        if false {
                                                                                                                            v106 := GenS7{F0: 1, F1: v1, F2: (v1 | (v1 / (((v1 / (((v1 * 725)) | 1))) | 1)))}
                                                                                                                            v2 = append(v2, (^(v3)))
                                                                                                                            _ = v106
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(v1, fi39, v36)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v72
                                                                                                        }
                                                                                                    }
                                                                                                    v31 = float64(0.0)
                                                                                                    if (((!(fi29 <= (fi39 | (-(fi39))))) && false) || ((!(-321 == -1)) || true)) {
                                                                                                        var v107 float64 = 1e10
                                                                                                        var v108 int64 = -349
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    }
                                                                                                }
                                                                                                v36 = v36
                                                                                                _ = v46
                                                                                            } else {
                                                                                                switch v36 {
                                                                                                case "a":
                                                                                                    v43 = v43
                                                                                                case "hello":
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                }
                                                                                                var v109 uint = 255
                                                                                                v1 += (v1 ^ (-(v1)))
                                                                                                var v110 int32 = -1
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            }
                                                                                            v31 += v31
                                                                                            v43 = v43
                                                                                        } else {
                                                                                            fmt.Println(v41, v3)
                                                                                            var v111 uint64 = v41
                                                                                            _ = v111
                                                                                        }
                                                                                        fmt.Println(v41, v36, v31)
                                                                                        if true {
                                                                                            var v112 int16 = (v1 % (((v1 % (((v1 - v1)) | 1))) | 1))
                                                                                            v1 -= (v1 - -32768)
                                                                                            if (-9223372036854775808 <= (fi29 / ((fi39) | 1))) {
                                                                                                if (false && true) {
                                                                                                    v112 -= v1
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            if (209 == 2) {
                                                                                                                var v113 int16 = (v1 % (((v1 / (((^(v112))) | 1))) | 1))
                                                                                                                v113 = v1
                                                                                                                if false {
                                                                                                                    v1 = (v112 | int16(v43))
                                                                                                                }
                                                                                                                _ = v113
                                                                                                            }
                                                                                                            v31 = v31
                                                                                                        }
                                                                                                        v2 = append(v2, (v3 + uint32(v40)))
                                                                                                        fmt.Println(v3)
                                                                                                    } else {
                                                                                                        v114 := GenS8{F0: v40, F1: 261, F2: v40}
                                                                                                        _ = v114
                                                                                                    }
                                                                                                } else {
                                                                                                    var v115 int32 = v43
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                    _ = len(v2)
                                                                                                    _ = v115
                                                                                                }
                                                                                                var v116 string = (v36 + v21)
                                                                                                v117 := GenS14{F0: (v43 * (v43 >> ((232) & 63))), F1: v3}
                                                                                                _ = v116
                                                                                                _ = v117
                                                                                            }
                                                                                            _ = v112
                                                                                        } else {
                                                                                            v21 = ("" + v36)
                                                                                            var v118 uint8 = 28
                                                                                            v2 = append(v2, v3)
                                                                                            _ = v118
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v119 int16 = (v1 & v1)
                                                                                    _ = v119
                                                                                default:
                                                                                    if (v21 != ((((v36 + "\n\t") + v36) + (v21 + v21)) + v21)) {
                                                                                        if (3.14159 <= 1e10) {
                                                                                            fmt.Println(fi39, v1, v1)
                                                                                        } else {
                                                                                            v36 = (v21 + "hello")
                                                                                        }
                                                                                        var v120 uint64 = genFn0()
                                                                                        v121 := GenS11{F0: (v41 / (((v120 | 18446744073709551615)) | 1)), F1: 603}
                                                                                        v120 = (v41 - v41)
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    }
                                                                                }
                                                                                if (((!true) && true) && (true || false)) {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    fmt.Println(v43)
                                                                                    fmt.Println(v21, v40, v41)
                                                                                } else {
                                                                                    v41 = (v41 / ((1) | 1))
                                                                                    if (v1 >= v1) {
                                                                                        v122 := GenS15{F0: (v40 * v40)}
                                                                                        _ = v122
                                                                                    } else {
                                                                                        fmt.Println(v21, v21, v21)
                                                                                        var v123 bool = true
                                                                                        _ = len(v2)
                                                                                        _ = v123
                                                                                    }
                                                                                    var v124 uint8 = 64
                                                                                    _ = v124
                                                                                }
                                                                                if (!true) {
                                                                                    if (!(3.14159 < 3.14159)) {
                                                                                        fmt.Println(v41, v3)
                                                                                    } else {
                                                                                        fmt.Println(v41, fi39, fi29)
                                                                                    }
                                                                                    if (false && (!false)) {
                                                                                        if (0.0 == (-2.25 / 0.1)) {
                                                                                            switch v36 {
                                                                                            case "a":
                                                                                                v40 -= (v40 ^ v40)
                                                                                            }
                                                                                        }
                                                                                        v40 = 2
                                                                                        if (((v21 + v36) + v21) == "") {
                                                                                            _ = len(v2)
                                                                                            switch (v21 + v21) {
                                                                                            case "hello":
                                                                                                fmt.Println(v43)
                                                                                            default:
                                                                                                v31 = (0.0 * 0.0)
                                                                                            }
                                                                                            v36 = (v36 + "foo bar")
                                                                                            v31 += 0.0
                                                                                        } else {
                                                                                            if true {
                                                                                                switch fi29 {
                                                                                                case 2147483647:
                                                                                                    var v125 uint64 = (v41 & (^(v41)))
                                                                                                    _ = v125
                                                                                                case 524:
                                                                                                    v1 = (v1 * 1)
                                                                                                    v43 = (v43 % (((-(v43))) | 1))
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = v3
                                                                                                    }
                                                                                                }
                                                                                                v3 = v3
                                                                                                fmt.Println(v3, v21)
                                                                                            } else {
                                                                                                v126 := GenS11{F0: uint64(v3), F1: 810}
                                                                                                var v127 int32 = 1
                                                                                                var v128 float32 = 1e10
                                                                                                if (true && false) {
                                                                                                    v129 := GenS16{F0: (v3 % ((2147483647) | 1)), F1: v40, F2: v1}
                                                                                                    var v130 uint32 = v3
                                                                                                    _ = v129
                                                                                                    _ = v130
                                                                                                }
                                                                                                v36 = "foo bar"
                                                                                                _ = v126
                                                                                                _ = v127
                                                                                                _ = v128
                                                                                            }
                                                                                            v36 = (v36 + ("" + "foo bar"))
                                                                                            var v131 float64 = 1e10
                                                                                            _ = v131
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v40, fi29, v40)
                                                                                    v43 = int32(v41)
                                                                                }
                                                                            }
                                                                            v31 -= v31
                                                                            if (!((v1 % (((v1 >> ((18446744073709551615) & 63))) | 1)) >= v1)) {
                                                                                v43 = v43
                                                                                v132 := GenS10{F0: 770}
                                                                                _ = v132
                                                                            } else {
                                                                                v21 = ("" + "")
                                                                            }
                                                                            v133 := GenS11{F0: v41, F1: (v43 >> ((18446744073709551615) & 63))}
                                                                            _ = v133
                                                                        } else {
                                                                            if true {
                                                                                var v134 int64 = int64(v43)
                                                                                v31 -= float64(1e10)
                                                                                v135 := GenS2{F0: v1, F1: 1, F2: (fi39 & fi29)}
                                                                                _ = v134
                                                                                _ = v135
                                                                            } else {
                                                                                v21 = (v36 + (("a" + v21) + ("foo bar" + (v21 + "\n\t"))))
                                                                                v136 := GenS6{F0: 4294967295, F1: 0, F2: (v1 >> ((1) & 63))}
                                                                                _ = v136
                                                                            }
                                                                        }
                                                                        _ = v43
                                                                    } else {
                                                                        switch ("\n\t" + "") {
                                                                        case "":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v3
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(v40, v21)
                                                                        default:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v3
                                                                            }
                                                                        }
                                                                        var v137 string = ((v36 + "") + v36)
                                                                        var v138 string = (v137 + "foo bar")
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                case 305:
                                                                    v21 = "hello"
                                                                }
                                                            } else {
                                                                switch fi29 {
                                                                case 2147483647:
                                                                    v1 = 42
                                                                }
                                                                fmt.Println(v21)
                                                            }
                                                        case 32767:
                                                            fmt.Println(v1, v3, v31)
                                                        default:
                                                            var v139 uint16 = (v40 % (((v40 + v40)) | 1))
                                                            _ = v139
                                                        }
                                                    }
                                                }
                                                switch 127 {
                                                case -93:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 9223372036854775807:
                                                    v31 = 1.5
                                                }
                                                _ = v41
                                            }
                                            _ = v40
                                        }
                                        if true {
                                            v140 := GenS2{F0: v1, F1: 42, F2: (fi29 ^ int(v3))}
                                            switch v36 {
                                            case "foo bar":
                                                v141 := GenS13{F0: 584, F1: false}
                                                _ = v141
                                            default:
                                                if true {
                                                    var v142 int64 = -1
                                                    switch (v36 + v36) {
                                                    case "":
                                                        v31 -= 0.0
                                                    case "foo bar":
                                                        if ((277 <= 1) || false) {
                                                            v21 = (((v21 + v36) + v21) + (v21 + v21))
                                                            if len(v2) > 0 {
                                                                v2[0] = 400
                                                            }
                                                        }
                                                    case "a":
                                                        if (("\n\t" < "foo bar") && false) {
                                                            var v143 float32 = 1.5
                                                            switch (v36 + v21) {
                                                            case "foo bar":
                                                                v31 += 0.0
                                                                if true {
                                                                    var v144 float64 = (float64(v143) * (1e10 * float64(float32((v31 + 3.14159)))))
                                                                    var v145 uint16 = 958
                                                                    _ = v144
                                                                    _ = v145
                                                                } else {
                                                                    fmt.Println(v21, v142)
                                                                    v36 = "hello"
                                                                    var v146 float32 = float32((v31 * v31))
                                                                    _ = v146
                                                                }
                                                                if (!(false && (9223372036854775807 >= (v142 >> ((18446744073709551615) & 63))))) {
                                                                    var v147 int16 = 1
                                                                    _ = v147
                                                                } else {
                                                                    v148 := GenS5{F0: 42, F1: (v142 & -228)}
                                                                    v31 -= v31
                                                                    switch -120 {
                                                                    case 4294967295:
                                                                        _ = len(v2)
                                                                    case -9223372036854775808:
                                                                        var v149 bool = false
                                                                        _ = v149
                                                                    case 354:
                                                                        fmt.Println(v21)
                                                                    }
                                                                    _ = v148
                                                                }
                                                            case "hello":
                                                                if (65535 == v3) {
                                                                    if (!(v143 > float32((float64(((0.0 / v143) - 3.14159)) / -2.25)))) {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v3
                                                                        }
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    v143 = 0.0
                                                                } else {
                                                                    if (v1 <= (v1 / (((v1 << ((42) & 63))) | 1))) {
                                                                        var v150 uint64 = uint64(v3)
                                                                        var v151 int16 = 231
                                                                        v152 := GenS13{F0: (v150 - (v150 << ((457) & 63))), F1: true}
                                                                        _ = v150
                                                                        _ = v151
                                                                        _ = v152
                                                                    } else {
                                                                        v143 = (0.1 * ((-2.25 * (v143 / 1.0)) * v143))
                                                                        var v153 int = (fi29 - 1)
                                                                        _ = v153
                                                                    }
                                                                    if false {
                                                                        v154 := GenS16{F0: (^(v3)), F1: 42, F2: v1}
                                                                        _ = v154
                                                                    } else {
                                                                        var v155 uint64 = 4294967295
                                                                        var v156 int64 = 165
                                                                        v157 := GenS17{F0: 32767, F1: v1, F2: (^(v3))}
                                                                        _ = v155
                                                                        _ = v156
                                                                        _ = v157
                                                                    }
                                                                }
                                                            case "":
                                                                if len(v2) > 0 {
                                                                    v2[0] = (^(v3))
                                                                }
                                                            }
                                                            v158 := GenS10{F0: (v3 & v3)}
                                                            _ = v143
                                                            _ = v158
                                                        } else {
                                                            v36 = v36
                                                        }
                                                    }
                                                    _ = v142
                                                }
                                            }
                                            _ = v140
                                        } else {
                                            fmt.Println(v21, v21)
                                            v159 := GenS17{F0: 373, F1: 657, F2: v3}
                                            _ = v159
                                        }
                                        v160 := GenS4{F0: 42}
                                        _ = v160
                                    }
                                    v2 = append(v2, (v3 * v3))
                                    _ = fi39
                                }
                                _ = v36
                            }
                            _ = v31
                        }
                    }
                    for fi161 := 0; fi161 < 12; fi161++ {
                        v2 = append(v2, v3)
                        v162 := GenS18{F0: 42, F1: -1, F2: true}
                        var v163 bool = (!false)
                        _ = fi161
                        _ = v162
                        _ = v163
                    }
                    _ = fi29
                }
            case "hello":
                var v164 uint64 = 1
                _ = v164
            case "foo bar":
                for fi165 := 0; fi165 < 17; fi165++ {
                    if len(v2) > 0 {
                        v2[0] = (v3 << ((4294967295) & 63))
                    }
                    if (1.5 > (0.0 + -2.25)) {
                        if false {
                            if (42 != 663) {
                                v166 := GenS19{F0: 9223372036854775807, F1: 90}
                                fmt.Println(fi165, v1, fi165)
                                v167 := GenS5{F0: 18446744073709551615, F1: 172}
                                _ = v166
                                _ = v167
                            } else {
                                for fi168 := 0; fi168 < 14; fi168++ {
                                    v169 := GenS4{F0: 42}
                                    _ = fi168
                                    _ = v169
                                }
                            }
                            v21 = v21
                            v2 = append(v2, (^(v3)))
                        } else {
                            var v170 uint = 42
                            switch 386 {
                            case -1:
                                var v171 int = (fi165 % ((32768) | 1))
                                _ = v171
                            default:
                                for fi172 := 0; fi172 < 3; fi172++ {
                                    switch fi172 {
                                    case -2147483648:
                                        var v173 float64 = -2.25
                                        _ = v173
                                    case 32768:
                                        var v174 int32 = 0
                                        _ = v174
                                    default:
                                        var v175 bool = (1.5 >= (0.0 + 1.5))
                                        _ = v175
                                    }
                                    v1 += v1
                                    switch (v21 + v21) {
                                    case "hello":
                                        var v176 int16 = 942
                                        _ = v176
                                    }
                                    var v177 int = (fi165 << ((4294967295) & 63))
                                    _ = fi172
                                    _ = v177
                                }
                            }
                            fmt.Println(v21, v3, v1)
                            _ = v170
                        }
                        fmt.Println(v1)
                    } else {
                        _ = len(v2)
                        v178 := GenS2{F0: 42, F1: 1, F2: 32768}
                        _ = v178
                    }
                    v2 = append(v2, (v3 + 4294967295))
                    _ = fi165
                }
            default:
                for fi179 := 0; fi179 < 6; fi179++ {
                    _ = len(v2)
                    _ = fi179
                }
            }
            fmt.Println(v21)
            v21 = "foo bar"
        }
    }
    if len(v2) > 0 {
        v2[0] = 368
    }
    v2 = append(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
