package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int32 = 2147483647
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int64 = -547
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch ("" + "foo bar") {
    case "\n\t":
        var v10 int64 = 0
        _ = v10
    default:
        var v11 int8 = 88
        _ = v11
    }
    if (false || false) {
        var v12 bool = (false && (!true))
        v12 = v12
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v12, v12)
            var v14 uint32 = 65535
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 uint = 65535
    _ = v15
    return 2147483647
}

func main() {
    var v1 uint16 = 60
    var v2 uint32 = 61
    var v3 int8 = 127
    var v4 bool = true
    var v5 uint16 = v1
    var v16 int8 = 100
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 float32 = float32((1e10 - (((0.0 - 1e10) + (((1.5 * (-2.25 - float64((1.5 - 1e10)))) + 1.5) - ((3.14159 / -2.25) * -2.25))) / (1.5 - (float64((0.1 * (0.0 * (3.14159 - ((0.0 * 1e10) - 1e10))))) + (0.1 - 1.5))))))
        for fi19 := 0; fi19 < 13; fi19++ {
            if (!((v2 << ((563) & 63)) != (^(v2)))) {
                v20 := []int{(fi17 / ((fi19) | 1))}
                _ = v20
            } else {
                if (!((v3 < (v3 ^ v3)) && (18446744073709551615 != 18446744073709551615))) {
                    v21 := []int{fi17, 985, 4294967295}
                    _ = v21
                } else {
                    v2 += (v2 ^ (v2 - (v2 | genFn0())))
                }
                if true {
                    var v22 string = "foo bar"
                    switch v22 {
                    case "\n\t":
                        var v23 string = ("foo bar" + v22)
                        _ = v23
                    case "a":
                        v24 := []int{127}
                        _ = v24
                    case "foo bar":
                        if (0 == 18446744073709551615) {
                            v25 := []int{}
                            v3 = v16
                            v26 := GenS1{F0: (v1 ^ v5), F1: 4294967295, F2: 255}
                            _ = v25
                            _ = v26
                        }
                    default:
                        v3 += -1
                    }
                    _ = v22
                }
                v27 := []int{fi17, 405, (fi19 << ((1) & 63))}
                _ = v27
            }
            var v28 float64 = 3.14159
            var v29 int16 = -32768
            _ = fi19
            _ = v28
            _ = v29
        }
        _ = fi17
        _ = v18
    }
    v3 += -1
    fmt.Println(v3)
    if v4 {
        var v30 uint16 = (v1 + (v1 >> ((18446744073709551615) & 63)))
        if v4 {
            if (!v4) {
                for fi31 := 0; fi31 < 10; fi31++ {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        switch fi31 {
                        case 42:
                            v33 := []int{}
                            _ = v33
                        case 1:
                            switch ("\n\t" + "\n\t") {
                            case "":
                                fmt.Println(v30)
                                if v4 {
                                    v34 := GenS1{F0: (^(v1)), F1: 18446744073709551615, F2: 1}
                                    _ = v34
                                } else {
                                    fmt.Println(v5)
                                }
                                v5 = v30
                            case "foo bar":
                                v1 += v1
                            case "\n\t":
                                switch "" {
                                case "":
                                    var v35 int32 = 2147483647
                                    _ = v35
                                }
                            }
                        }
                        var v36 uint64 = 2
                        _ = fi32
                        _ = v36
                    }
                    var v37 bool = (((3.14159 - 3.14159) - 1.5) != 1.5)
                    _ = fi31
                    _ = v37
                }
            }
        }
        _ = v30
    } else {
        for fi38 := 0; fi38 < 12; fi38++ {
            for fi39 := 0; fi39 < 3; fi39++ {
                switch (^(fi38)) {
                case 0:
                    var v40 int64 = 42
                    _ = v40
                case 32767:
                    var v41 string = ("" + "a")
                    _ = v41
                case -2147483648:
                    v42 := GenS1{F0: (v5 | (v5 << ((42) & 63))), F1: 42, F2: 797}
                    _ = v42
                default:
                    v1 = (v1 & v5)
                }
                _ = fi39
            }
            var v43 uint16 = 2
            v4 = (false || v4)
            _ = fi38
            _ = v43
        }
        for fi44 := 0; fi44 < 20; fi44++ {
            v5 += 746
            _ = fi44
        }
    }
    v45 := GenS1{F0: (v5 & v5), F1: 345, F2: 1}
    v46 := []int{9223372036854775807, 2147483647}
    v47 := GenS2{F0: v2}
    var v48 float32 = ((float32(0.0) * float32(0.0)) + 3.14159)
    for fi49 := 0; fi49 < 19; fi49++ {
        if (true && false) {
            var v50 uint8 = 42
            _ = v50
        }
        if (!(v48 <= v48)) {
            if false {
                if v4 {
                    v4 = ("" != ("\n\t" + ("hello" + "hello")))
                } else {
                    var v51 bool = false
                    v5 += (v1 + 0)
                    _ = v51
                }
                var v52 uint64 = 0
                _ = v52
            } else {
                var v53 bool = ((0 <= 102) || false)
                v2 += v2
                if (!(!(!(!v4)))) {
                    if len(v46) > 0 {
                        v46[0] = fi49
                    }
                    if (v53 || true) {
                        fmt.Println(v4, v2)
                        for fi54 := 0; fi54 < 7; fi54++ {
                            switch "a" {
                            case "\n\t":
                                var v55 int8 = v16
                                _ = v55
                            }
                            var v56 uint8 = 2
                            switch "foo bar" {
                            case "a":
                                if len(v46) >= 2 {
                                    v46 = v46[1:len(v46)]
                                }
                            case "hello":
                                switch "\n\t" {
                                case "\n\t":
                                    if false {
                                        _ = len(v46)
                                        switch ("" + (("foo bar" + (("foo bar" + ("\n\t" + "")) + "")) + "hello")) {
                                        case "hello":
                                            var v57 uint16 = v1
                                            _ = v57
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            v56 = v56
                                        case "a":
                                            var v58 float64 = float64(v48)
                                            _ = v58
                                        }
                                        fmt.Println(v4, v2)
                                        v59 := GenS1{F0: 1, F1: 42, F2: 150}
                                        _ = v59
                                    } else {
                                        var v60 int32 = -2147483648
                                        if (1.5 < 1e10) {
                                            v61 := GenS1{F0: (v5 >> ((4294967295) & 63)), F1: 42, F2: 18446744073709551615}
                                            _ = v61
                                        }
                                        v48 = v48
                                        _ = v60
                                    }
                                case "":
                                    var v62 float32 = 1e10
                                    _ = v62
                                default:
                                    var v63 uint64 = 1
                                    var v64 float64 = 0.0
                                    var v65 int64 = 1
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            case "foo bar":
                                v4 = false
                            }
                            v66 := GenS1{F0: 65535, F1: 654, F2: 255}
                            _ = fi54
                            _ = v56
                            _ = v66
                        }
                    } else {
                        var v67 uint = 4294967295
                        switch (-(fi49)) {
                        case 2147483647:
                            if (-2.25 <= v48) {
                                fmt.Println(v4, v48)
                                for fi68 := 0; fi68 < 9; fi68++ {
                                    if (((float64(1.5) / 1e10) / float64(v48)) >= float64(v48)) {
                                        v16 += -128
                                        if (v4 && true) {
                                            if true {
                                                if v4 {
                                                    var v69 uint16 = v1
                                                    _ = v69
                                                } else {
                                                    var v70 uint = v67
                                                    var v71 float64 = ((((0.0 / (1e10 * float64(-2.25))) - (3.14159 / -2.25)) * 3.14159) * 1e10)
                                                    _ = v70
                                                    _ = v71
                                                }
                                            } else {
                                                switch ("\n\t" + "foo bar") {
                                                case "hello":
                                                    var v72 uint16 = (v1 / ((v5) | 1))
                                                    _ = v72
                                                case "":
                                                    if (42 != 2147483647) {
                                                        var v73 uint16 = (^(v1))
                                                        fmt.Println(v5, v48)
                                                        _ = v73
                                                    }
                                                case "a":
                                                    if (!false) {
                                                        var v74 uint8 = 255
                                                        var v75 int16 = 861
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            var v76 int64 = 42
                                            v4 = v4
                                            var v77 int32 = -2147483648
                                            _ = v76
                                            _ = v77
                                        } else {
                                            var v78 uint32 = v2
                                            _ = v78
                                        }
                                        if v53 {
                                            v79 := GenS1{F0: (v1 << ((0) & 63)), F1: 447, F2: (v67 * (v67 + v67))}
                                            v2 = 2147483647
                                            var v80 bool = (42 > -317)
                                            _ = v79
                                            _ = v80
                                        }
                                    }
                                    switch fi49 {
                                    case 1:
                                        var v81 uint = (v67 * v67)
                                        _ = v81
                                    case -908:
                                        v82 := GenS1{F0: v1, F1: 0, F2: v67}
                                        _ = v82
                                    }
                                    if len(v46) > 0 {
                                        v46[0] = (fi68 / ((-9223372036854775808) | 1))
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        var v83 string = "hello"
                                        _ = v83
                                    case "":
                                        switch (("" + "foo bar") + (("" + "\n\t") + "")) {
                                        case "\n\t":
                                            var v84 bool = v53
                                            _ = v84
                                        case "hello":
                                            var v85 uint16 = 748
                                            switch "\n\t" {
                                            case "foo bar":
                                                v46 = append(v46, fi49)
                                            case "a":
                                                if (fi68 == (^(fi49))) {
                                                    v67 -= v67
                                                    var v86 uint64 = 1
                                                    var v87 int8 = 127
                                                    if v4 {
                                                        v3 -= 100
                                                    } else {
                                                        v16 = int8(v67)
                                                        v87 = (v3 & v16)
                                                        switch "" {
                                                        case "":
                                                            v87 += (v3 & 1)
                                                        case "\n\t":
                                                            fmt.Println(v2, v5, v3)
                                                        default:
                                                            switch (-(fi68)) {
                                                            case -9223372036854775808:
                                                                if (0.1 > 0.0) {
                                                                    v88 := GenS3{F0: (v67 & v67)}
                                                                    switch "" {
                                                                    case "hello":
                                                                        if v4 {
                                                                            fmt.Println(v86, fi49, v48)
                                                                        }
                                                                    case "foo bar":
                                                                        var v89 bool = v4
                                                                        _ = v89
                                                                    case "":
                                                                        _ = len(v46)
                                                                    }
                                                                    if (!(v87 > v16)) {
                                                                        var v90 uint16 = v1
                                                                        v5 -= v85
                                                                        _ = v90
                                                                    }
                                                                    _ = v88
                                                                }
                                                            case 1:
                                                                var v91 int8 = v87
                                                                var v92 uint16 = v1
                                                                if (698 < 1) {
                                                                    switch ("a" + "") {
                                                                    case "foo bar":
                                                                        v93 := GenS1{F0: v1, F1: (v86 + v86), F2: (v67 >> (((v67 % ((18446744073709551615) | 1))) & 63))}
                                                                        _ = v93
                                                                    case "":
                                                                        v67 -= v67
                                                                    case "a":
                                                                        var v94 bool = ((127 < int(v2)) || true)
                                                                        _ = v94
                                                                    default:
                                                                        v67 += (v67 % ((65535) | 1))
                                                                    }
                                                                    v53 = (uint8(v86) >= 1)
                                                                    var v95 int16 = -1
                                                                    _ = len(v46)
                                                                    var v96 float32 = ((-2.25 + float32(float64(v48))) * v48)
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                                _ = v91
                                                                _ = v92
                                                            default:
                                                                _ = len(v46)
                                                            }
                                                        }
                                                        v67 = (v67 / ((255) | 1))
                                                        var v97 uint8 = 128
                                                        switch "" {
                                                        case "":
                                                            v3 = (v87 & v87)
                                                        }
                                                        var v98 uint64 = v86
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    v1 = uint16(v2)
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    var v99 int64 = -9223372036854775808
                                                    _ = v99
                                                }
                                            case "hello":
                                                v100 := GenS1{F0: v5, F1: 118, F2: v67}
                                                _ = v100
                                            }
                                            _ = v85
                                        default:
                                            v4 = (v2 != (v2 ^ genFn0()))
                                        }
                                    case "\n\t":
                                        v3 = (v3 >> (((v67 * (^(v67)))) & 63))
                                    }
                                    _ = fi68
                                }
                                v53 = v4
                            }
                        case 4294967295:
                            var v101 int16 = 42
                            _ = v101
                        case 1:
                            var v102 uint16 = (v1 - v1)
                            _ = v102
                        default:
                            v16 = v3
                        }
                        _ = v67
                    }
                }
                _ = v53
            }
            var v103 float64 = 0.1
            v1 += 42
            _ = v103
        }
        _ = fi49
    }
    _ = len(v46)
    fmt.Println(v48, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
