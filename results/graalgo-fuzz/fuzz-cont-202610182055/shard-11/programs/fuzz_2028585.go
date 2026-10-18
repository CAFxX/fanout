package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 bool; F2 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64; F1 int8 }
type GenS11 struct { F0 int8; F1 uint8 }
type GenS12 struct { F0 string; F1 int }

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 bool = (-727 > 42)
    v4 := []string{("" + "foo bar"), "a"}
    var v5 int64 = 135
    var v6 float64 = 0.1
    v7 := GenS1{F0: 99, F1: 885, F2: 43}
    fmt.Println(v6)
    fmt.Println(v6)
    switch 42 {
    case 32767:
        v8 := GenS1{F0: 243, F1: 255, F2: -128}
        _ = v8
    case 127:
        switch -510 {
        case -233:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 0:
            switch "a" {
            case "a":
                v9 := GenS2{F0: 439, F1: 42}
                _ = v9
            case "":
                fmt.Println(v3)
            }
        case 342:
            var v10 string = "\n\t"
            _ = v10
        default:
            var v11 int8 = -35
            _ = v11
        }
    case 559:
        v5 = 87
    default:
        fmt.Println(v3)
    }
    v12 := GenS3{F0: 65535, F1: 0, F2: 255}
    var v13 int = -1
    var v14 uint64 = 0
    v14 = v14
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS3{F0: 535, F1: 2, F2: 2}
        switch "" {
        case "foo bar":
            if (681 >= 0) {
                if (32767 < 1) {
                    delete(v2, "hello")
                    delete(v2, "hello")
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 float64 = float64((float32(float64(0.0)) + -2.25))
                        v19 := GenS1{F0: 109, F1: 101, F2: 0}
                        fmt.Println(v18)
                        v6 -= v18
                        fmt.Println(v14)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                } else {
                    v20 := GenS1{F0: 2, F1: 255, F2: -128}
                    _ = v20
                }
                v21 := GenS4{F0: true, F1: 868}
                for fi22 := 0; fi22 < 4; fi22++ {
                    switch "hello" {
                    case "":
                        v23 := GenS3{F0: 4294967295, F1: 2, F2: 2}
                        _ = v23
                    case "foo bar":
                        v4 = append(v4, "")
                    case "hello":
                        v6 -= ((3.14159 + float64((float32((float64(1e10) + v6)) - 0.1))) + (v6 / float64(3.14159)))
                    }
                    if (2 != 1) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi22
                }
                _ = v21
            }
        case "a":
            switch (("" + ("foo bar" + "a")) + "a") {
            case "a":
                v14 = v14
            case "foo bar":
                _ = len(v2)
            }
        }
        switch "a" {
        case "":
            if false {
                var v24 uint = 18446744073709551615
                _ = v24
            }
        case "foo bar":
            var v25 float64 = 1.5
            _ = v25
        case "a":
            switch "hello" {
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v26 string = ""
                    _ = v26
                case "\n\t":
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v5 -= v5
                        fmt.Println(v3)
                        if false {
                            var v28 string = ""
                            v29 := v2["a"]
                            if len(v1) > 0 {
                                v1[0] = v28
                            }
                            _ = v28
                            _ = v29
                        }
                        _ = fi27
                    }
                }
            case "hello":
                v30 := GenS3{F0: 4294967295, F1: 255, F2: 148}
                _ = v30
            }
        default:
            v4 = append(v4, "")
        }
        _ = fi15
        _ = v16
    }
    var v31 int8 = 119
    for fi32 := 0; fi32 < 19; fi32++ {
        if v3 {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v5)
            case "":
                fmt.Println(fi32, v14)
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v31 -= (v31 - int8(v14))
            }
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            switch (("a" + ("hello" + "hello")) + "hello") {
            case "hello":
                var v34 uint8 = uint8(v14)
                switch ("foo bar" + ("hello" + "hello")) {
                case "foo bar":
                    v35 := v2[""]
                    _ = v35
                }
                var v36 uint64 = v14
                _ = v34
                _ = v36
            default:
                if true {
                    if v3 {
                        switch ("" + "hello") {
                        case "foo bar":
                            var v37 int8 = v31
                            _ = v37
                        case "a":
                            v31 = v31
                        default:
                            v14 = (v14 | (^(v14)))
                            v38 := GenS3{F0: 2, F1: 1, F2: 2}
                            _ = v38
                        }
                        _ = len(v2)
                    } else {
                        v39 := GenS1{F0: 1, F1: 225, F2: v31}
                        v3 = false
                        _ = v39
                    }
                    var v40 int64 = v5
                    v41 := GenS2{F0: 0, F1: 0}
                    if v3 {
                        switch ("" + "\n\t") {
                        case "foo bar":
                            v13 = (v13 - -2147483648)
                        default:
                            var v42 float64 = 1.5
                            _ = v42
                        }
                        v43 := GenS3{F0: 723, F1: 255, F2: 1}
                        _ = v43
                    }
                    _ = v40
                    _ = v41
                } else {
                    var v44 bool = ((65535 != 32767) || (171 == 19))
                    var v45 int8 = v31
                    var v46 int16 = 0
                    _ = v44
                    _ = v45
                    _ = v46
                }
            }
            fmt.Println(fi33, v14)
            v31 = (v31 >> ((0) & 63))
            v5 += (v5 * (v5 * (v5 >> ((1) & 63))))
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v47 := GenS2{F0: 0, F1: 418}
            switch "foo bar" {
            case "foo bar":
                if v3 {
                    if ((v3 || (true && (!v3))) && true) {
                        fmt.Println(v13, v3)
                        if true {
                            fmt.Println(v3, fi32)
                        } else {
                            var v48 int32 = -797
                            _ = v48
                        }
                    } else {
                        switch int(v14) {
                        case 9223372036854775807:
                            if ((true || true) && v3) {
                                fmt.Println(v13)
                                var v49 uint8 = 0
                                _ = v49
                            }
                        }
                    }
                    switch "hello" {
                    case "\n\t":
                        v50 := GenS1{F0: 127, F1: 4294967295, F2: -1}
                        _ = v50
                    default:
                        v51 := GenS5{F0: (0 > 0)}
                        _ = v51
                    }
                    v3 = (true || true)
                }
                if (!(!false)) {
                    v52 := GenS1{F0: 1, F1: 18446744073709551615, F2: -1}
                    var v53 float64 = (v6 * v6)
                    fmt.Println(v6, v6)
                    _ = v52
                    _ = v53
                }
                v3 = (fi32 < (fi32 % ((-996) | 1)))
                if ((1 > -2147483648) || (759 < 1)) {
                    var v54 float64 = v6
                    v55 := GenS4{F0: v3, F1: 0}
                    v56 := GenS3{F0: 255, F1: 222, F2: 42}
                    _ = v54
                    _ = v55
                    _ = v56
                }
            case "":
                v57 := GenS3{F0: 42, F1: uint8(v14), F2: 128}
                _ = v57
            }
            _ = fi33
            _ = v47
        }
        _ = fi32
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        v58 := GenS6{F0: 1, F1: v3, F2: 1}
        _ = v58
    default:
        switch "" {
        case "":
            _ = len(v2)
        case "hello":
            for fi59 := 0; fi59 < 19; fi59++ {
                for fi60 := 0; fi60 < 2; fi60++ {
                    fmt.Println(v13)
                    v3 = true
                    v61 := v2[(("" + "hello") + "\n\t")]
                    v6 += float64((-2.25 / float32(v6)))
                    _ = fi60
                    _ = v61
                }
                if true {
                    for fi62 := 0; fi62 < 3; fi62++ {
                        var v63 uint16 = 1
                        if (0.1 >= v6) {
                            var v64 uint = 42
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v64
                        } else {
                            switch (fi62 >> ((65535) & 63)) {
                            case 419:
                                v65 := GenS5{F0: v3}
                                _ = v65
                            case -891:
                                v66 := v2[(("a" + "") + "a")]
                                _ = v66
                            case 223:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v3, v14, fi59)
                        }
                        _ = fi62
                        _ = v63
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v13 -= fi59
                } else {
                    switch (fi59 << ((42) & 63)) {
                    case 4294967295:
                        _ = len(v4)
                    case -9223372036854775808:
                        switch 128 {
                        case 2147483647:
                            if (18446744073709551615 <= 2) {
                                for fi67 := 0; fi67 < 5; fi67++ {
                                    var v68 int = (fi59 % ((-482) | 1))
                                    if false {
                                        _ = len(v2)
                                        var v69 uint32 = 912
                                        var v70 float32 = -2.25
                                        _ = v69
                                        _ = v70
                                    } else {
                                        var v71 uint32 = 478
                                        fmt.Println(v14, fi67, v6)
                                        var v72 string = ("" + ("" + "hello"))
                                        _ = v71
                                        _ = v72
                                    }
                                    var v73 uint64 = v14
                                    _ = fi67
                                    _ = v68
                                    _ = v73
                                }
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                fmt.Println(v31, fi59)
                            }
                        case 42:
                            v13 -= -1
                        }
                    }
                    for fi74 := 0; fi74 < 16; fi74++ {
                        switch v13 {
                        case 2147483647:
                            v3 = (!v3)
                        case 42:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v75 uint8 = 127
                        v76 := GenS7{F0: v13}
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                    switch (v13 * -2147483648) {
                    case -2147483648:
                        for fi77 := 0; fi77 < 6; fi77++ {
                            var v78 float64 = 1e10
                            if v3 {
                                if v3 {
                                    var v79 uint = 576
                                    var v80 uint64 = v14
                                    delete(v2, "\n\t")
                                    if false {
                                        if false {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            var v81 uint16 = 436
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v81 = (^(v81))
                                            _ = v81
                                        }
                                    } else {
                                        _ = len(v4)
                                        v82 := GenS8{F0: int16(v79), F1: int(v31)}
                                        v13 += fi77
                                        v13 = v13
                                        v83 := GenS9{F0: v79}
                                        delete(v2, "foo bar")
                                        v13 += fi59
                                        v2["a"] = (^(fi59))
                                        _ = v82
                                        _ = v83
                                    }
                                    v84 := GenS5{F0: v3}
                                    v85 := GenS10{F0: (^(v80)), F1: 26}
                                    var v86 uint64 = v14
                                    _ = v79
                                    _ = v80
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                } else {
                                    if true {
                                        switch fi77 {
                                        case -2147483648:
                                            var v87 uint64 = v14
                                            _ = v87
                                        case -698:
                                            v5 += v5
                                        }
                                        var v88 float32 = (float32(v78) - (float32(v78) * (3.14159 * 0.1)))
                                        _ = v88
                                    } else {
                                        v3 = v3
                                    }
                                    v6 -= (float64(1e10) + v6)
                                    var v89 string = "foo bar"
                                    _ = v89
                                }
                            } else {
                                v90 := GenS11{F0: int8(v5), F1: 42}
                                v78 += 3.14159
                                _ = v90
                            }
                            v2["foo bar"] = 2147483647
                            _ = fi77
                            _ = v78
                        }
                    default:
                        if false {
                            v6 = ((v6 - (v6 + (float64((1.5 + 1.5)) / v6))) - v6)
                            if ((v3 && v3) || v3) {
                                switch int(v14) {
                                case 2147483647:
                                    v4 = append(v4, "hello")
                                case 26:
                                    v14 = (v14 - v14)
                                case 127:
                                    if ((-2.25 / 3.14159) > 3.14159) {
                                        v91 := v2[""]
                                        _ = v91
                                    } else {
                                        var v92 bool = false
                                        v92 = v92
                                        _ = v92
                                    }
                                }
                            } else {
                                v2[("a" + "foo bar")] = fi59
                                for fi93 := 0; fi93 < 14; fi93++ {
                                    switch ("a" + "foo bar") {
                                    case "hello":
                                        var v94 uint = 42
                                        _ = v94
                                    }
                                    _ = fi93
                                }
                                var v95 bool = v3
                                _ = v95
                            }
                        } else {
                            var v96 int8 = v31
                            var v97 bool = (v6 < v6)
                            if false {
                                var v98 uint16 = 42
                                _ = v98
                            }
                            _ = v96
                            _ = v97
                        }
                    }
                }
                v99 := GenS12{F0: "a", F1: 953}
                _ = fi59
                _ = v99
            }
        default:
            v13 = -509
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
