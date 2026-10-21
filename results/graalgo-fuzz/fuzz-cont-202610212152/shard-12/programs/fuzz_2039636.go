package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 string; F2 string }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int16 = 882
    var v2 float32 = 3.14159
    v3 := []uint32{}
    v4 := make(map[string]bool)
    fmt.Println(v2)
    if false {
        v1 = -32768
        var v5 uint = 670
        _ = v5
    } else {
        v6 := GenS1{F0: 127}
        _ = v6
    }
    fmt.Println(v1, v2)
    for fi7 := 0; fi7 < 9; fi7++ {
        v2 = v2
        _ = fi7
    }
    v4["a"] = false
    var v8 uint16 = 32767
    if true {
        if true {
            if (317 != 65535) {
                switch ("a" + ("\n\t" + ("hello" + "foo bar"))) {
                case "\n\t":
                    v9 := GenS1{F0: -1}
                    _ = v9
                case "":
                    switch "\n\t" {
                    case "a":
                        v1 -= int16(v8)
                    case "foo bar":
                        if (!(0 != 103)) {
                            if ((true || ((((float32(1e10) / 1.5) * v2) / (v2 * 1e10)) > v2)) || (v8 != v8)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch -423 {
                                case 32767:
                                    var v10 uint = 255
                                    _ = v10
                                case 127:
                                    if true {
                                        var v11 uint8 = 42
                                        _ = v11
                                    } else {
                                        var v12 uint16 = 255
                                        v4["\n\t"] = (!(false && (!true)))
                                        switch "" {
                                        case "foo bar":
                                            for fi13 := 0; fi13 < 18; fi13++ {
                                                if true {
                                                    v1 += v1
                                                } else {
                                                    for fi14 := 0; fi14 < 5; fi14++ {
                                                        switch "\n\t" {
                                                        case "":
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "hello":
                                                                var v15 float32 = float32(float64(v2))
                                                                _ = v15
                                                            }
                                                        }
                                                        _ = fi14
                                                    }
                                                }
                                                var v16 uint64 = 1
                                                _ = fi13
                                                _ = v16
                                            }
                                        case "\n\t":
                                            v1 -= v1
                                        case "a":
                                            for fi17 := 0; fi17 < 14; fi17++ {
                                                v8 = v8
                                                var v18 uint16 = 508
                                                for fi19 := 0; fi19 < 15; fi19++ {
                                                    if (-9223372036854775808 > int64(v8)) {
                                                        v1 += v1
                                                        v1 -= v1
                                                    }
                                                    var v20 uint16 = v18
                                                    _ = fi19
                                                    _ = v20
                                                }
                                                _ = fi17
                                                _ = v18
                                            }
                                        default:
                                            var v21 bool = true
                                            _ = v21
                                        }
                                        _ = v12
                                    }
                                case -628:
                                    if (!true) {
                                        v22 := v4["\n\t"]
                                        switch "foo bar" {
                                        case "a":
                                            if (v22 && v22) {
                                                v8 = 129
                                                switch "hello" {
                                                case "foo bar":
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                case "\n\t":
                                                    switch 658 {
                                                    case 2147483647:
                                                        var v23 uint32 = 342
                                                        _ = v23
                                                    }
                                                }
                                            }
                                        case "hello":
                                            var v24 int32 = 0
                                            _ = v24
                                        case "foo bar":
                                            switch 24 {
                                            case 886:
                                                switch -2147483648 {
                                                case 0:
                                                    if (((1 >= int32(v8)) || v22) && (42 < -128)) {
                                                        for fi25 := 0; fi25 < 3; fi25++ {
                                                            v8 = (^(v8))
                                                            _ = fi25
                                                        }
                                                    } else {
                                                        v26 := GenS1{F0: 82}
                                                        _ = v26
                                                    }
                                                case 42:
                                                    v8 -= (v8 ^ 736)
                                                case -1:
                                                    for fi27 := 0; fi27 < 20; fi27++ {
                                                        var v28 uint64 = 1
                                                        _ = fi27
                                                        _ = v28
                                                    }
                                                default:
                                                    var v29 bool = v22
                                                    _ = v29
                                                }
                                            case -2147483648:
                                                var v30 uint64 = 0
                                                _ = v30
                                            default:
                                                _ = len(v3)
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        v4[""] = (2 < 4294967295)
                                        v31 := GenS2{F0: 509, F1: 0}
                                        _ = v31
                                    }
                                default:
                                    switch -9223372036854775808 {
                                    case 997:
                                        if false {
                                            v32 := GenS1{F0: -128}
                                            fmt.Println(v8, v2)
                                            var v33 int64 = 878
                                            _ = v32
                                            _ = v33
                                        } else {
                                            fmt.Println(v8, v8)
                                            var v34 uint32 = 790
                                            _ = v34
                                        }
                                    case -743:
                                        var v35 uint = 2
                                        _ = v35
                                    case 0:
                                        v36 := GenS1{F0: -1}
                                        _ = v36
                                    }
                                }
                            }
                            for fi37 := 0; fi37 < 10; fi37++ {
                                v2 = 3.14159
                                _ = fi37
                            }
                        }
                    }
                case "a":
                    v38 := GenS2{F0: 0, F1: 2}
                    _ = v38
                }
                fmt.Println(v8)
            }
        }
        if len(v3) > 0 {
            v3[0] = 608
        }
    }
    switch "hello" {
    case "foo bar":
        if (333 != 374) {
            var v39 uint64 = 379
            v1 -= (v1 % ((-32768) | 1))
            switch 32768 {
            case -2147483648:
                var v40 uint = 1
                _ = v40
            case -815:
                v8 = (v8 & v8)
            default:
                var v41 uint = 2
                _ = v41
            }
            _ = v39
        } else {
            v1 = v1
        }
    case "a":
        switch "\n\t" {
        case "a":
            fmt.Println(v8)
        case "\n\t":
            if false {
                v1 = v1
                v1 += 234
                switch ("foo bar" + ("foo bar" + "")) {
                case "a":
                    v42 := GenS3{F0: 148, F1: "hello", F2: ""}
                    _ = v42
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = 524
                    }
                case "\n\t":
                    var v43 int = 829
                    _ = v43
                }
            } else {
                var v44 string = "hello"
                v2 = -2.25
                v8 += v8
                _ = v44
            }
        default:
            var v45 uint32 = 42
            _ = v45
        }
    default:
        _ = len(v3)
    }
    var v46 int = 1
    v1 += 32767
    v47 := GenS3{F0: 18446744073709551615, F1: "hello", F2: ""}
    var v48 uint = 65535
    fmt.Println(v2, v8, v46)
    if ((3.14159 * 3.14159) < 0.0) {
        v49 := GenS2{F0: 42, F1: v48}
        _ = v49
    } else {
        v8 = (v8 & 1)
        v50 := GenS2{F0: 36, F1: (^(v48))}
        fmt.Println(v8, v48, v2)
        _ = v50
    }
    if ((!false) || true) {
        if (886 != 0) {
            if true {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                var v51 int = (v46 ^ 4294967295)
                _ = v51
            }
            v52 := GenS1{F0: 100}
            if ("\n\t" < "") {
                v2 = 0.1
            }
            _ = v52
        } else {
            switch ("hello" + "") {
            case "a":
                if true {
                    v53 := GenS2{F0: 77, F1: 18446744073709551615}
                    _ = v53
                } else {
                    switch -9223372036854775808 {
                    case 42:
                        var v54 int8 = 0
                        _ = v54
                    case 32767:
                        if (-128 == 1) {
                            v1 = 1000
                        }
                    }
                    fmt.Println(v1, v2, v48)
                    v46 -= (v46 + (v46 - (-(v46))))
                }
            default:
                v55 := GenS4{F0: "foo bar", F1: (v46 % ((v46) | 1)), F2: 2}
                _ = v55
            }
            if true {
                v4["hello"] = false
                if true {
                    v2 = v2
                } else {
                    switch v46 {
                    case -500:
                        v48 = (v48 % ((1) | 1))
                    case 128:
                        for fi56 := 0; fi56 < 14; fi56++ {
                            var v57 uint = v48
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                if (!((3.14159 / 0.1) < 1e10)) {
                                    fmt.Println(v8, fi56)
                                    switch 834 {
                                    case -9223372036854775808:
                                        if true {
                                            var v58 int = -1
                                            for fi59 := 0; fi59 < 17; fi59++ {
                                                v2 = v2
                                                var v60 uint8 = 229
                                                switch (v58 - v58) {
                                                case 2147483647:
                                                    if (!(!true)) {
                                                        var v61 int64 = 467
                                                        _ = v61
                                                    } else {
                                                        var v62 uint64 = 4294967295
                                                        _ = v62
                                                    }
                                                case 1:
                                                    if (false && (true || false)) {
                                                        switch 635 {
                                                        case -1:
                                                            v63 := v4[""]
                                                            _ = v63
                                                        default:
                                                            v46 = (v46 ^ fi59)
                                                        }
                                                        var v64 uint64 = 158
                                                        if (float32(3.14159) < (v2 / ((((v2 / (((0.0 / (v2 + (3.14159 - v2))) - v2) * v2)) / ((((v2 * (0.1 / 1.5)) - (v2 / v2)) * v2) + 1.5)) - (v2 + v2)) / (v2 * v2)))) {
                                                            v65 := v4["hello"]
                                                            var v66 bool = (v65 && v65)
                                                            _ = v65
                                                            _ = v66
                                                        } else {
                                                            var v67 uint64 = (v64 << ((v48) & 63))
                                                            v68 := GenS5{F0: 65535, F1: 65535, F2: 300}
                                                            var v69 uint64 = v67
                                                            if true {
                                                                v58 += (-(v46))
                                                                if false {
                                                                    if true {
                                                                        v70 := GenS6{F0: v1}
                                                                        v71 := GenS2{F0: uint64(v8), F1: 65535}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 670
                                                                        }
                                                                        var v72 uint32 = 243
                                                                        _ = v70
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    _ = len(v4)
                                                                    fmt.Println(v60, v48, v1)
                                                                }
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 321
                                                                }
                                                            }
                                                            v60 = (v60 << (((v48 & 57)) & 63))
                                                            var v73 float32 = (1e10 * 0.0)
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                            _ = v73
                                                        }
                                                        _ = v64
                                                    } else {
                                                        if true {
                                                            var v74 uint32 = 2147483647
                                                            _ = v74
                                                        }
                                                        v3 = append(v3, 2)
                                                        var v75 int8 = 54
                                                        _ = v75
                                                    }
                                                }
                                                _ = fi59
                                                _ = v60
                                            }
                                            v58 += (v46 + v46)
                                            _ = v58
                                        }
                                    default:
                                        v76 := v4["a"]
                                        _ = v76
                                    }
                                } else {
                                    v46 = 127
                                    v1 = v1
                                    if true {
                                        _ = len(v3)
                                        fmt.Println(v48, v57, v2)
                                    } else {
                                        var v77 uint32 = uint32(v48)
                                        var v78 uint = v48
                                        if false {
                                            if len(v3) > 0 {
                                                v3[0] = 2147483647
                                            }
                                            var v79 string = ""
                                            _ = v79
                                        } else {
                                            var v80 int8 = 98
                                            for fi81 := 0; fi81 < 1; fi81++ {
                                                v82 := GenS6{F0: (v1 % (((v1 / ((42) | 1))) | 1))}
                                                v77 = v77
                                                switch 0 {
                                                case 448:
                                                    var v83 uint32 = 2147483648
                                                    _ = v83
                                                }
                                                _ = fi81
                                                _ = v82
                                            }
                                            var v84 bool = ((42 != 2147483647) || ((v80 & (-(v80))) != (v80 / (((v80 % (((-(v80))) | 1))) | 1))))
                                            _ = v80
                                            _ = v84
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            }
                            _ = fi56
                            _ = v57
                        }
                    case 0:
                        if false {
                            var v85 int16 = (-(v1))
                            fmt.Println(v46, v48, v1)
                            for fi86 := 0; fi86 < 13; fi86++ {
                                for fi87 := 0; fi87 < 6; fi87++ {
                                    v88 := GenS6{F0: -32768}
                                    _ = fi87
                                    _ = v88
                                }
                                _ = fi86
                            }
                            _ = v85
                        }
                    default:
                        v1 -= 32767
                    }
                    var v89 string = "foo bar"
                    fmt.Println(v1, v1)
                    _ = v89
                }
            }
        }
    } else {
        for fi90 := 0; fi90 < 17; fi90++ {
            for fi91 := 0; fi91 < 8; fi91++ {
                v92 := GenS3{F0: 4294967295, F1: "hello", F2: "hello"}
                _ = len(v3)
                var v93 uint = (v48 % (((v48 % (((v48 * 2)) | 1))) | 1))
                _ = fi91
                _ = v92
                _ = v93
            }
            _ = fi90
        }
        if false {
            for fi94 := 0; fi94 < 7; fi94++ {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if (0 <= 0) {
                    v3 = append(v3, 635)
                } else {
                    var v95 uint8 = 186
                    for fi96 := 0; fi96 < 15; fi96++ {
                        var v97 bool = (113 > v95)
                        var v98 int32 = 2147483647
                        v99 := GenS1{F0: 87}
                        var v100 uint8 = (v95 + v95)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v101 string = "foo bar"
                        _ = fi96
                        _ = v100
                        _ = v101
                        _ = v97
                        _ = v98
                        _ = v99
                    }
                    v1 -= (v1 + (v1 - (^(v1))))
                    _ = v95
                }
                var v102 int16 = v1
                _ = fi94
                _ = v102
            }
            var v103 float64 = float64(v2)
            _ = v103
        }
        var v104 uint8 = 2
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v8
}
