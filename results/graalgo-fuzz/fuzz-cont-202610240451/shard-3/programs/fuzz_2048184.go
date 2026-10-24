package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint16; F2 int32 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32, p1 int16, p2 string) uint16 {
    var v7 uint8 = 0
    v8 := []int{282, 4294967295, -695, 684}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool) string {
    fmt.Println(p0)
    p0 = p0
    p0 = (!false)
    switch (("" + ("a" + "\n\t")) + ((("hello" + "\n\t") + "foo bar") + "")) {
    case "":
        switch "\n\t" {
        case "foo bar":
            for fi9 := 0; fi9 < 5; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch (fi9 ^ fi9) {
                    case -9223372036854775808:
                        p0 = (p0 || p0)
                    case 32767:
                        fmt.Println(fi9, fi10, fi10)
                    case 127:
                        var v11 int32 = 42
                        _ = v11
                    default:
                        p0 = p0
                    }
                    _ = fi10
                }
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch 782 {
                    case -2147483648:
                        var v13 float32 = 3.14159
                        v14 := GenS1{F0: 42, F1: "\n\t"}
                        _ = v13
                        _ = v14
                    case -828:
                        if (fi9 > -744) {
                            if (fi9 != 128) {
                                v15 := GenS2{F0: -9223372036854775808, F1: 2}
                                p0 = (1.5 != 1.5)
                                _ = v15
                            }
                            p0 = true
                            var v16 float64 = float64(((1e10 + ((0.1 * 3.14159) - 1e10)) + 3.14159))
                            _ = v16
                        } else {
                            var v17 int8 = -70
                            var v18 uint64 = 2
                            _ = v17
                            _ = v18
                        }
                    default:
                        if (1 > 0) {
                            p0 = (!p0)
                        } else {
                            switch fi9 {
                            case 533:
                                var v19 bool = p0
                                v20 := GenS3{F0: 227, F1: (fi9 >> ((988) & 63)), F2: 11}
                                _ = v19
                                _ = v20
                            default:
                                var v21 uint64 = 1
                                v21 = v21
                                p0 = true
                                v22 := []int{(-(fi12))}
                                _ = v21
                                _ = v22
                            }
                        }
                    }
                    v23 := GenS2{F0: 9223372036854775807, F1: 42}
                    fmt.Println(fi12, p0, p0)
                    _ = fi12
                    _ = v23
                }
                fmt.Println(p0)
                _ = fi9
            }
        case "":
            if p0 {
                p0 = p0
                switch "hello" {
                case "hello":
                    v24 := GenS4{F0: "\n\t", F1: 128}
                    _ = v24
                case "\n\t":
                    var v25 int32 = 2147483647
                    _ = v25
                }
                var v26 uint16 = 42
                _ = v26
            }
        default:
            fmt.Println(p0)
        }
    default:
        fmt.Println(p0)
    }
    v27 := []int{4294967295, -965, 2147483647}
    _ = p0
    _ = v27
    return ("foo bar" + "")
}

func main() {
    var v1 uint32 = 773
    var v2 float64 = (float64((((1e10 + 0.0) + -2.25) + 3.14159)) * (float64(0.0) / ((0.0 * float64(1.5)) + 1e10)))
    var v3 int32 = -2147483648
    v4 := []uint32{v1}
    v5 := []int{-932}
    var v6 int16 = -1
    var v28 uint8 = 61
    fmt.Println(v2, v1)
    var v29 int32 = -981
    v30 := GenS5{F0: (v28 % (((^(v28))) | 1)), F1: 1, F2: 148}
    if (!false) {
        if ((true && true) && true) {
            if false {
                v28 -= 127
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v31 float64 = (v2 / (v2 - (1.5 / float64(-2.25))))
                    _ = v31
                } else {
                    var v32 bool = ((^(v1)) > 2147483647)
                    _ = v32
                }
                var v33 float64 = (0.1 * float64((3.14159 + -2.25)))
                _ = v33
            } else {
                v1 += (v1 << ((4294967295) & 63))
            }
            if (!(true || true)) {
                switch 3 {
                case 0:
                    var v34 uint16 = 255
                    _ = v34
                case 32768:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 553:
                    if false {
                        if false {
                            if true {
                                var v35 uint16 = 220
                                _ = v35
                            }
                            if ((!false) || true) {
                                fmt.Println(v28, v6, v3)
                            } else {
                                if (!((!true) && false)) {
                                    fmt.Println(v3, v6, v29)
                                }
                                var v36 int16 = v6
                                _ = v36
                            }
                            var v37 bool = false
                            _ = v37
                        } else {
                            v38 := GenS2{F0: 42, F1: 2147483647}
                            if false {
                                v39 := GenS6{F0: v1, F1: -9223372036854775808}
                                v28 = 6
                                _ = v39
                            }
                            v2 = float64(float32((-2.25 / (v2 - (v2 - v2)))))
                            _ = v38
                        }
                        var v40 int32 = v3
                        _ = v40
                    }
                default:
                    fmt.Println(v2)
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    v2 = v2
                    v42 := GenS1{F0: 220, F1: "\n\t"}
                    v2 += 0.1
                    _ = fi41
                    _ = v42
                }
                if false {
                    switch 412 {
                    case 42:
                        switch "foo bar" {
                        case "hello":
                            var v43 int64 = 331
                            _ = v43
                        case "foo bar":
                            if true {
                                var v44 string = "a"
                                if false {
                                    if true {
                                        if false {
                                            var v45 int16 = (-(v6))
                                            switch 877 {
                                            case 1:
                                                v1 = 1
                                            case 42:
                                                if len(v5) > 0 {
                                                    v5[0] = -9223372036854775808
                                                }
                                            case 0:
                                                v3 += 81
                                            }
                                            switch (v44 + (v44 + "")) {
                                            case "hello":
                                                var v46 uint64 = 0
                                                _ = v46
                                            case "foo bar":
                                                var v47 uint16 = 612
                                                _ = v47
                                            case "\n\t":
                                                var v48 float64 = v2
                                                _ = v48
                                            }
                                            _ = v45
                                        } else {
                                            v4 = append(v4, (v1 + (v1 ^ v1)))
                                            var v49 uint8 = v28
                                            switch "" {
                                            case "\n\t":
                                                for fi50 := 0; fi50 < 5; fi50++ {
                                                    for fi51 := 0; fi51 < 14; fi51++ {
                                                        v49 += (^(v28))
                                                        var v52 int = fi50
                                                        _ = fi51
                                                        _ = v52
                                                    }
                                                    var v53 float64 = v2
                                                    v44 = v44
                                                    _ = fi50
                                                    _ = v53
                                                }
                                            default:
                                                var v54 uint32 = v1
                                                _ = v54
                                            }
                                            _ = v49
                                        }
                                        for fi55 := 0; fi55 < 13; fi55++ {
                                            switch "a" {
                                            case "a":
                                                v56 := GenS6{F0: v1, F1: int64(v6)}
                                                _ = v56
                                            }
                                            if false {
                                                if true {
                                                    var v57 int16 = v6
                                                    _ = v57
                                                } else {
                                                    switch (^(fi55)) {
                                                    case 32767:
                                                        fmt.Println(fi55)
                                                    default:
                                                        for fi58 := 0; fi58 < 12; fi58++ {
                                                            v59 := GenS3{F0: 255, F1: 2147483647, F2: 948}
                                                            var v60 string = v44
                                                            var v61 int64 = 574
                                                            _ = fi58
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    }
                                                }
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v29 -= (v3 & 315)
                                            }
                                            _ = fi55
                                        }
                                        fmt.Println(v1, v1, v28)
                                    } else {
                                        if false {
                                            v62 := GenS1{F0: 4294967295, F1: v44}
                                            for fi63 := 0; fi63 < 19; fi63++ {
                                                v5 = append(v5, (fi63 | (fi63 / ((fi63) | 1))))
                                                fmt.Println(v2, v3)
                                                _ = fi63
                                            }
                                            if true {
                                                if (v1 >= (v1 << ((18446744073709551615) & 63))) {
                                                    for fi64 := 0; fi64 < 16; fi64++ {
                                                        switch fi64 {
                                                        case -2147483648:
                                                            v65 := GenS7{F0: (!((-43 < 0) || (v44 > v44))), F1: (^(v1))}
                                                            _ = v65
                                                        case 1:
                                                            fmt.Println(v44)
                                                        case 32768:
                                                            v66 := GenS4{F0: v44, F1: 0}
                                                            _ = v66
                                                        default:
                                                            var v67 int16 = int16(v1)
                                                            _ = v67
                                                        }
                                                        var v68 int64 = 42
                                                        if true {
                                                            v69 := GenS4{F0: "hello", F1: (v28 ^ (v28 / (((v28 << ((2) & 63))) | 1)))}
                                                            var v70 float64 = v2
                                                            var v71 int = int(v6)
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            v72 := GenS8{F0: -1}
                                                            _ = v72
                                                        }
                                                        _ = fi64
                                                        _ = v68
                                                    }
                                                    v6 = v6
                                                } else {
                                                    var v73 int32 = v3
                                                    v44 = v44
                                                    for fi74 := 0; fi74 < 7; fi74++ {
                                                        var v75 string = ("foo bar" + "\n\t")
                                                        v76 := GenS9{F0: (0 < 65535), F1: uint64(v28), F2: 1}
                                                        _ = fi74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    _ = v73
                                                }
                                            }
                                            _ = v62
                                        }
                                    }
                                    v2 = 3.14159
                                    v28 += (v28 | v28)
                                }
                                var v77 uint8 = (v28 ^ 191)
                                _ = v44
                                _ = v77
                            } else {
                                var v78 uint32 = (^(v1))
                                for fi79 := 0; fi79 < 13; fi79++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi79
                                }
                                v6 += (v6 | (v6 - v6))
                                _ = v78
                            }
                        case "a":
                            var v80 uint = 342
                            _ = v80
                        default:
                            fmt.Println(v1, v3)
                        }
                    case 0:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            var v81 int = 127
            _ = v81
        }
    }
    if true {
        var v82 bool = true
        _ = len(v5)
        if ((v3 / (((v3 ^ (v3 << ((42) & 63)))) | 1)) != int32(v28)) {
            v1 -= v1
            fmt.Println(v82)
            for fi83 := 0; fi83 < 5; fi83++ {
                v84 := GenS10{F0: 9223372036854775807, F1: 151, F2: int32(v28)}
                if len(v5) > 0 {
                    v5[0] = fi83
                }
                var v85 float32 = 0.0
                _ = fi83
                _ = v84
                _ = v85
            }
        }
        _ = v82
    }
    if true {
        v86 := GenS11{F0: -1, F1: v1}
        _ = v86
    }
    fmt.Println(v6, v2)
    for fi87 := 0; fi87 < 1; fi87++ {
        v88 := GenS12{F0: -348, F1: -765}
        _ = fi87
        _ = v88
    }
    var v89 string = ""
    if true {
        switch -9223372036854775808 {
        case 9223372036854775807:
            fmt.Println(v29)
        case 665:
            v2 = 0.1
        case 32767:
            if len(v4) > 0 {
                v4[0] = (v1 * v1)
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if true {
            if (false && (1e10 == 0.0)) {
                v1 = v1
                for fi90 := 0; fi90 < 8; fi90++ {
                    var v91 bool = (false && true)
                    if v91 {
                        var v92 string = v89
                        var v93 uint = 246
                        var v94 bool = (609 != 996)
                        _ = v92
                        _ = v93
                        _ = v94
                    }
                    fmt.Println(fi90)
                    _ = fi90
                    _ = v91
                }
            } else {
                v95 := GenS8{F0: -92}
                fmt.Println(v6)
                v29 -= -1
                _ = v95
            }
            v1 = v1
        }
        v89 = v89
    }
    if false {
        if true {
            fmt.Println(v6, v29)
            var v96 int16 = -711
            _ = v96
        } else {
            _ = len(v4)
            if len(v4) > 0 {
                v4[0] = (v1 * 4294967295)
            }
            for fi97 := 0; fi97 < 14; fi97++ {
                _ = len(v4)
                v98 := GenS13{F0: 42, F1: -1}
                _ = fi97
                _ = v98
            }
        }
        switch v89 {
        case "foo bar":
            var v99 int32 = -47
            _ = v99
        case "a":
            for fi100 := 0; fi100 < 14; fi100++ {
                if len(v5) > 0 {
                    v5[0] = int(v6)
                }
                for fi101 := 0; fi101 < 17; fi101++ {
                    var v102 int8 = int8(v6)
                    _ = fi101
                    _ = v102
                }
                var v103 int16 = (-(v6))
                _ = fi100
                _ = v103
            }
        case "\n\t":
            v6 = v6
        }
        var v104 uint8 = (v28 | (v28 & uint8(v1)))
        _ = v104
    } else {
        var v105 int = -2147483648
        _ = v105
    }
    v3 = (^(v3))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6, v6, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v89
}
