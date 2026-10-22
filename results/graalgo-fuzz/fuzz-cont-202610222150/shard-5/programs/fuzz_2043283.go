package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 uint; F1 int32; F2 string }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 bool; F1 int; F2 int64 }
func genFn0() uint {
    v7 := GenS1{F0: 1}
    if (1 >= 42) {
        v8 := []int{4294967295, 1}
        switch "a" {
        case "foo bar":
            var v9 uint16 = 0
            _ = v9
        default:
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 int32 = 1
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        var v13 uint16 = 751
        _ = v13
        _ = v8
    } else {
        var v14 float64 = 0.1
        _ = v14
    }
    switch 116 {
    case -9223372036854775808:
        switch -978 {
        case 1:
            var v15 uint32 = 2147483648
            _ = v15
        case 0:
            if true {
                if true {
                    var v16 float32 = (3.14159 * 1.5)
                    _ = v16
                }
                var v17 string = "\n\t"
                _ = v17
            } else {
                var v18 uint8 = 42
                v19 := GenS2{F0: (4294967295 != 1)}
                _ = v18
                _ = v19
            }
        }
    case 1:
        var v20 uint8 = 255
        _ = v20
    }
    switch -12 {
    case 42:
        var v21 int32 = 1
        _ = v21
    case 32768:
        var v22 float64 = (float64(0.0) / 1.0)
        _ = v22
    case -164:
        var v23 uint16 = 49
        _ = v23
    default:
        v24 := GenS3{F0: 2, F1: 42, F2: false}
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int8 = 80
        _ = fi25
        _ = v26
    }
    _ = v7
    return 255
}
func genFn1(p0 int64, p1 uint, p2 int16) uint {
    p0 = (p0 + p0)
    var v27 string = (("a" + (("" + "foo bar") + "hello")) + "a")
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return (p1 + p1)
}

func main() {
    var v1 int32 = 1
    var v2 uint32 = 2147483647
    var v3 uint16 = 2
    var v4 float32 = (1e10 - 1e10)
    var v5 string = "hello"
    var v6 uint64 = 4294967295
    var v28 int16 = 32767
    v4 = v4
    if true {
        v28 += 42
    }
    v29 := []int{-1, -734}
    v30 := GenS4{F0: 863}
    v2 -= 2
    switch v5 {
    case "foo bar":
        if false {
            if false {
                v6 -= uint64(v2)
            } else {
                if (uint(v3) == 255) {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 uint16 = (v3 ^ v3)
                        v33 := GenS3{F0: 32767, F1: -16, F2: true}
                        var v34 uint64 = v6
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 14; fi35++ {
                        if (42 > 0) {
                            v36 := GenS2{F0: true}
                            for fi37 := 0; fi37 < 10; fi37++ {
                                var v38 int64 = 9223372036854775807
                                var v39 float32 = v4
                                if true {
                                    switch v5 {
                                    case "hello":
                                        if false {
                                            var v40 int = fi37
                                            fmt.Println(v38)
                                            _ = v40
                                        } else {
                                            fmt.Println(v4, v38, fi35)
                                            v5 = "\n\t"
                                            var v41 uint64 = 0
                                            _ = v41
                                        }
                                    case "foo bar":
                                        v5 = (v5 + v5)
                                    }
                                    v38 += v38
                                }
                                if true {
                                    v3 = (v3 >> ((1) & 63))
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        v28 += -32768
                                    case -9223372036854775808:
                                        var v42 int32 = -1
                                        _ = v42
                                    default:
                                        var v43 int64 = v38
                                        _ = v43
                                    }
                                    v44 := GenS1{F0: 550}
                                    _ = v44
                                }
                                v45 := GenS3{F0: v28, F1: 62, F2: ((fi35 * (fi37 + fi37)) == fi37)}
                                fmt.Println(v1, v38, v5)
                                v46 := GenS3{F0: 32767, F1: -25, F2: true}
                                v2 = v2
                                _ = fi37
                                _ = v38
                                _ = v39
                                _ = v45
                                _ = v46
                            }
                            _ = v36
                        }
                        var v47 int8 = -85
                        fmt.Println(v3)
                        _ = fi35
                        _ = v47
                    }
                    fmt.Println(v3, v6)
                }
                v48 := GenS5{F0: -709}
                _ = v48
            }
            if false {
                var v49 int = 127
                _ = v49
            } else {
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
            }
            v50 := GenS6{F0: 128}
            _ = v50
        }
    case "":
        v51 := GenS1{F0: -513}
        _ = v51
    case "\n\t":
        v52 := GenS7{F0: v28}
        _ = v52
    default:
        var v53 float64 = 3.14159
        _ = v53
    }
    v4 += (0.1 - float32(1.5))
    v1 = v1
    v29 = append(v29, -872)
    switch "foo bar" {
    case "":
        for fi54 := 0; fi54 < 4; fi54++ {
            if ((0.1 < 1e10) && (448 == int64(v28))) {
                fmt.Println(v28, v4, fi54)
                v55 := GenS8{F0: -1, F1: 1}
                v29 = append(v29, 127)
                _ = v55
            } else {
                fmt.Println(v5)
                if (!((v5 < "foo bar") || ((v4 > (0.1 * 3.14159)) && (!(-1 != 1))))) {
                    if false {
                        switch (fi54 % ((2147483647) | 1)) {
                        case 4294967295:
                            switch "" {
                            case "hello":
                                if len(v29) > 0 {
                                    v29[0] = fi54
                                }
                            case "":
                                var v56 string = v5
                                _ = v56
                            default:
                                v57 := GenS8{F0: 42, F1: 86}
                                _ = v57
                            }
                        }
                        if false {
                            v29 = append(v29, fi54)
                            var v58 uint16 = v3
                            v59 := GenS5{F0: 1}
                            _ = v58
                            _ = v59
                        } else {
                            var v60 float64 = 1.5
                            v61 := GenS3{F0: (^(v28)), F1: 127, F2: false}
                            _ = len(v29)
                            _ = v60
                            _ = v61
                        }
                    }
                }
            }
            _ = fi54
        }
    case "foo bar":
        var v62 int16 = v28
        _ = v62
    case "\n\t":
        var v63 string = v5
        _ = v63
    default:
        v64 := GenS7{F0: v28}
        _ = v64
    }
    if (!(uint32(v3) > (v2 ^ (v2 & (v2 / ((994) | 1)))))) {
        v1 += (v1 % (((v1 - -770)) | 1))
        var v65 int32 = -579
        if (v2 > (v2 - 2)) {
            _ = len(v29)
        } else {
            if (!(9223372036854775807 != 825)) {
                var v66 uint16 = (v3 | (v3 % ((42) | 1)))
                _ = v66
            }
            for fi67 := 0; fi67 < 20; fi67++ {
                v28 = (v28 ^ (v28 & 1000))
                fmt.Println(v3)
                _ = fi67
            }
            var v68 float64 = -2.25
            _ = v68
        }
        _ = v65
    }
    if false {
        switch 127 {
        case 897:
            var v69 float64 = (0.1 + (1e10 * 3.14159))
            _ = v69
        case -898:
            for fi70 := 0; fi70 < 19; fi70++ {
                v3 += (v3 - 153)
                _ = fi70
            }
        }
        v29 = append(v29, -228)
    } else {
        if (((32768 != 4294967295) || (((v28 << ((2) & 63)) >= v28) && (114 == 68))) || (uint8(v3) < 1)) {
            v71 := GenS9{F0: (v28 + v28), F1: 9223372036854775807}
            for fi72 := 0; fi72 < 13; fi72++ {
                if true {
                    _ = len(v29)
                    var v73 uint64 = v6
                    _ = v73
                }
                if len(v29) > 0 {
                    v29[0] = fi72
                }
                v74 := GenS7{F0: (v28 ^ (v28 % (((v28 << ((18446744073709551615) & 63))) | 1)))}
                _ = fi72
                _ = v74
            }
            _ = v71
        } else {
            _ = len(v29)
            v3 = v3
            for fi75 := 0; fi75 < 17; fi75++ {
                var v76 int16 = v28
                for fi77 := 0; fi77 < 1; fi77++ {
                    if true {
                        var v78 int = 4294967295
                        v3 = (v3 + v3)
                        _ = v78
                    }
                    v79 := GenS10{F0: 255, F1: v1, F2: v5}
                    switch (fi77 - fi75) {
                    case -1:
                        if (true || ((v6 & (v6 + (v6 & uint64(v3)))) > (^(v6)))) {
                            switch "a" {
                            case "foo bar":
                                v29 = append(v29, fi75)
                            case "":
                                switch 48 {
                                case 2147483647:
                                    if (-128 < 127) {
                                        v3 -= 412
                                        v80 := GenS7{F0: (-(v28))}
                                        if false {
                                            var v81 float32 = ((0.1 - (3.14159 * v4)) - v4)
                                            _ = v81
                                        } else {
                                            var v82 uint16 = 65535
                                            v1 = (v1 & (v1 ^ v1))
                                            _ = v82
                                        }
                                        if len(v29) > 0 {
                                            v29[0] = (fi77 ^ int(v6))
                                        }
                                        _ = v80
                                    } else {
                                        v28 -= (v76 | v76)
                                        v2 += v2
                                        fmt.Println(fi77)
                                    }
                                }
                                if true {
                                    var v83 uint64 = v6
                                    var v84 uint8 = 42
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 uint8 = uint8(v6)
                                    v3 = v3
                                    _ = v85
                                }
                            case "\n\t":
                                var v86 int32 = (v1 % ((v1) | 1))
                                _ = v86
                            }
                            var v87 uint = 337
                            _ = v87
                        }
                    case 391:
                        v88 := GenS6{F0: 109}
                        _ = v88
                    default:
                        v89 := GenS5{F0: 9223372036854775807}
                        _ = v89
                    }
                    _ = fi77
                    _ = v79
                }
                v90 := GenS11{F0: v2}
                _ = fi75
                _ = v76
                _ = v90
            }
        }
        if (-740 < 468) {
            fmt.Println(v2, v28, v5)
            if len(v29) > 0 {
                v29[0] = -312
            }
            for fi91 := 0; fi91 < 1; fi91++ {
                fmt.Println(v3, v28)
                _ = fi91
            }
        } else {
            v92 := GenS1{F0: (v28 * (v28 >> ((708) & 63)))}
            fmt.Println(v28, v28, v4)
            switch v5 {
            case "\n\t":
                for fi93 := 0; fi93 < 11; fi93++ {
                    for fi94 := 0; fi94 < 6; fi94++ {
                        v95 := GenS9{F0: (^(v28)), F1: (fi94 & int(v28))}
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            var v96 int32 = (v1 | v1)
                            _ = v96
                        case "hello":
                            var v97 string = v5
                            _ = v97
                        }
                        _ = fi94
                        _ = v95
                    }
                    v4 -= 0.1
                    if (!(v28 < v28)) {
                        switch v5 {
                        case "\n\t":
                            if false {
                                var v98 uint = genFn1(int64(v2), 411, v28)
                                _ = v98
                            }
                        default:
                            switch "hello" {
                            case "hello":
                                fmt.Println(fi93, v2, v3)
                            case "foo bar":
                                if (v6 <= v6) {
                                    v3 -= v3
                                }
                            default:
                                var v99 int = fi93
                                _ = v99
                            }
                        }
                        fmt.Println(v4, v4)
                    } else {
                        var v100 int8 = -4
                        v4 -= v4
                        v101 := GenS12{F0: (!(true && (v2 != (v2 + 334)))), F1: fi93, F2: -590}
                        _ = v100
                        _ = v101
                    }
                    _ = fi93
                }
            }
            _ = v92
        }
    }
    for fi102 := 0; fi102 < 17; fi102++ {
        fmt.Println(v2)
        var v103 int16 = (v28 / ((v28) | 1))
        v29 = append(v29, (^(fi102)))
        _ = fi102
        _ = v103
    }
    if len(v29) > 0 {
        v29[0] = 128
    }
    v4 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
