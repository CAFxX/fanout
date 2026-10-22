package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS9 struct { F0 int; F1 int8; F2 uint64 }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{776, 129}
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    var v5 uint16 = 143
    var v6 float64 = (3.14159 * 3.14159)
    var v7 int64 = -850
    if (127 > 108) {
        fmt.Println(v6, v7)
    }
    v8 := GenS1{F0: 1, F1: 434}
    v6 -= float64(1e10)
    switch v1 {
    case "hello":
        if false {
            v1 = v1
            fmt.Println(v1, v6, v6)
        } else {
            var v9 uint32 = 2147483648
            v9 += v9
            _ = v9
        }
    case "\n\t":
        if (((1e10 / -2.25) + (-2.25 / 1.0)) >= (0.1 - 3.14159)) {
            var v10 uint64 = 2
            var v11 int8 = 75
            v1 = v1
            _ = v10
            _ = v11
        } else {
            if true {
                var v12 float64 = v6
                switch (v1 + v1) {
                case "a":
                    v12 += (v12 * v12)
                default:
                    v13 := GenS2{F0: 1}
                    _ = v13
                }
                _ = v12
            }
        }
    case "foo bar":
        for fi14 := 0; fi14 < 5; fi14++ {
            v7 += v7
            v5 = 32767
            v15 := GenS1{F0: 1000, F1: 338}
            _ = fi14
            _ = v15
        }
    default:
        v16 := GenS1{F0: 0, F1: 204}
        _ = v16
    }
    v17 := GenS3{F0: false}
    var v18 int = -1
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v19 uint64 = 4294967295
    _ = len(v3)
    v18 += (v18 % (((v18 | (v18 << ((1) & 63)))) | 1))
    for fi20 := 0; fi20 < 2; fi20++ {
        switch 32768 {
        case -9223372036854775808:
            var v21 int8 = int8(v5)
            _ = v21
        case 32767:
            _ = len(v4)
        case -890:
            fmt.Println(v6, v19)
        }
        v22 := GenS1{F0: 42, F1: 1}
        _ = fi20
        _ = v22
    }
    var v23 uint = 4294967295
    switch -9223372036854775808 {
    case -154:
        var v24 uint64 = (v19 & v19)
        _ = v24
    case -1:
        if true {
            switch v18 {
            case -818:
                switch v18 {
                case 954:
                    for fi25 := 0; fi25 < 20; fi25++ {
                        for fi26 := 0; fi26 < 8; fi26++ {
                            _ = len(v2)
                            var v27 int64 = v7
                            if true {
                                v5 += v5
                            } else {
                                v28 := GenS1{F0: 0, F1: 1}
                                _ = v28
                            }
                            _ = fi26
                            _ = v27
                        }
                        _ = fi25
                    }
                case -9223372036854775808:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        switch "\n\t" {
                        case "foo bar":
                            if ((v7 < v7) || false) {
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    v1 = v1
                                    _ = fi30
                                }
                                v5 = v5
                                _ = len(v2)
                            } else {
                                v7 -= (v7 / (((v7 % (((v7 + v7)) | 1))) | 1))
                                if (!false) {
                                    v2 = append(v2, 2147483648)
                                }
                            }
                        case "\n\t":
                            switch v1 {
                            case "hello":
                                var v31 float32 = 1e10
                                _ = v31
                            case "":
                                if (1.5 > (0.0 - 0.0)) {
                                    v6 -= float64(-2.25)
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        var v33 int16 = 32767
                                        _ = fi32
                                        _ = v33
                                    }
                                    _ = len(v2)
                                }
                            case "\n\t":
                                v34 := GenS4{F0: (v5 >> ((111) & 63))}
                                _ = v34
                            }
                        case "hello":
                            var v35 int = v18
                            _ = v35
                        default:
                            var v36 int = int(v5)
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 11; fi37++ {
                            if false {
                                switch (v1 + v1) {
                                case "foo bar":
                                    fmt.Println(v19)
                                case "hello":
                                    var v38 uint64 = v19
                                    _ = v38
                                case "\n\t":
                                    switch fi29 {
                                    case 9223372036854775807:
                                        v6 -= float64((-2.25 - (((1.5 * 1.5) + ((float32(0.0) * -2.25) * 0.0)) - ((1.5 + 0.0) + (1.5 - 1e10)))))
                                    default:
                                        if (2 == v19) {
                                            v1 = v1
                                            v23 = v23
                                            v6 = ((v6 / v6) - v6)
                                            fmt.Println(v19, v23)
                                            fmt.Println(fi29)
                                        }
                                    }
                                default:
                                    switch v18 {
                                    case -180:
                                        var v39 string = ""
                                        _ = v39
                                    case 9223372036854775807:
                                        v40 := GenS4{F0: v5}
                                        _ = v40
                                    default:
                                        var v41 int8 = -90
                                        _ = v41
                                    }
                                    if true {
                                        var v42 int = (^(fi29))
                                        _ = v42
                                    } else {
                                        var v43 uint32 = 970
                                        if false {
                                            if (v1 >= "foo bar") {
                                                if len(v2) > 0 {
                                                    v2[0] = v43
                                                }
                                                v43 += (v43 & v43)
                                                v44 := GenS5{F0: -32768, F1: 1}
                                                _ = v44
                                            } else {
                                                v45 := GenS2{F0: (v23 & (^(v23)))}
                                                _ = v45
                                            }
                                            v1 = v1
                                        }
                                        v1 = (v1 + v1)
                                        _ = v43
                                    }
                                    var v46 float64 = (v6 / (3.14159 * 3.14159))
                                    _ = v46
                                }
                            }
                            _ = fi37
                        }
                        _ = fi29
                    }
                }
            case -165:
                switch (v18 % (((v18 - int(v5))) | 1)) {
                case 127:
                    v4[int(v19)] = v1
                case 646:
                    v47 := v3[(v1 + v1)]
                    _ = v47
                case -9223372036854775808:
                    v5 -= v5
                default:
                    var v48 float32 = 3.14159
                    _ = v48
                }
            case -2147483648:
                var v49 bool = true
                _ = v49
            }
            v18 -= v18
            for fi50 := 0; fi50 < 14; fi50++ {
                if true {
                    var v51 int8 = 33
                    switch (v18 * v18) {
                    case -1:
                        v51 -= int8(v7)
                    default:
                        v52 := GenS6{F0: (fi50 % (((^(fi50))) | 1)), F1: (v51 & 32)}
                        _ = v52
                    }
                    _ = v51
                } else {
                    for fi53 := 0; fi53 < 10; fi53++ {
                        v54 := GenS7{F0: -649, F1: 2147483647}
                        v23 += v23
                        v55 := GenS1{F0: 1, F1: uint(v5)}
                        v19 = v19
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    v56 := GenS8{F0: 42, F1: 1, F2: v19}
                    v5 += 42
                    _ = v56
                }
                _ = fi50
            }
        } else {
            switch 4294967295 {
            case 9223372036854775807:
                v3[(((v1 + "a") + v1) + v1)] = (v1 + ((v1 + "hello") + v1))
            case -2147483648:
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 string = "hello"
                    v59 := GenS7{F0: -32768, F1: 42}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            case -9223372036854775808:
                v19 += (v19 << ((0) & 63))
            }
            if true {
                v19 -= (v19 & v19)
                if false {
                    v60 := GenS9{F0: v18, F1: -106, F2: (v19 | 18446744073709551615)}
                    switch -570 {
                    case -9223372036854775808:
                        for fi61 := 0; fi61 < 1; fi61++ {
                            v62 := GenS6{F0: 42, F1: -128}
                            _ = fi61
                            _ = v62
                        }
                    case -1:
                        for fi63 := 0; fi63 < 3; fi63++ {
                            v23 += v23
                            if (fi63 >= (v18 ^ fi63)) {
                                delete(v3, "\n\t")
                                if true {
                                    if (v1 > v1) {
                                        switch fi63 {
                                        case -2147483648:
                                            var v64 int = (^(fi63))
                                            _ = v64
                                        case 4294967295:
                                            v7 = v7
                                        default:
                                            for fi65 := 0; fi65 < 13; fi65++ {
                                                switch 9223372036854775807 {
                                                case 268:
                                                    if (true || false) {
                                                        var v66 uint64 = v19
                                                        _ = v66
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 56
                                                        }
                                                        _ = len(v4)
                                                    }
                                                case 42:
                                                    switch 128 {
                                                    case -1:
                                                        if (false && (0 < 94)) {
                                                            fmt.Println(v7, v19)
                                                            if false {
                                                                v67 := GenS10{F0: v19, F1: (v7 / (((v7 / ((v7) | 1))) | 1))}
                                                                var v68 int64 = v7
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                fmt.Println(v7, v7)
                                                                if false {
                                                                    if (!((143 <= 127) && (!(v1 == "foo bar")))) {
                                                                        var v69 string = (v1 + (v1 + ("a" + ((v1 + v1) + v1))))
                                                                        var v70 int8 = 42
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        v71 := GenS3{F0: false}
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            fmt.Println(v5, v7)
                                                                        }
                                                                        v2 = append(v2, 517)
                                                                        var v72 int = 0
                                                                        var v73 int32 = 1
                                                                        if ((0.1 >= (v6 / ((float64(0.0) * v6) / float64(-2.25)))) || (!((2 <= (v19 / (((v19 * 42)) | 1))) && true))) {
                                                                            var v74 int16 = 42
                                                                            var v75 int64 = 168
                                                                            _ = v74
                                                                            _ = v75
                                                                        } else {
                                                                            v7 += v7
                                                                            var v76 uint32 = 4294967295
                                                                            var v77 uint32 = 4294967295
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    var v78 int32 = -1
                                                                    _ = v78
                                                                } else {
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        v79 := GenS11{F0: 98}
                                                                        _ = v79
                                                                    }
                                                                    var v80 string = v1
                                                                    if (("\n\t" + "foo bar") == (("hello" + (v1 + "\n\t")) + (("hello" + v80) + v1))) {
                                                                        if false {
                                                                            v6 = v6
                                                                        } else {
                                                                            delete(v4, v18)
                                                                            v81 := GenS6{F0: (fi63 % (((v18 + (^(fi65)))) | 1)), F1: 127}
                                                                            fmt.Println(v1, fi65)
                                                                            fmt.Println(v5, v18, v19)
                                                                            var v82 int64 = v7
                                                                            var v83 bool = true
                                                                            _ = v81
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                    _ = v80
                                                                }
                                                                fmt.Println(v7)
                                                            }
                                                        }
                                                    case 2147483647:
                                                        switch (fi63 >> ((42) & 63)) {
                                                        case 1:
                                                            fmt.Println(fi63, v6, v1)
                                                        }
                                                    }
                                                case -604:
                                                    fmt.Println(fi63, v23, v6)
                                                default:
                                                    var v84 int8 = 100
                                                    _ = v84
                                                }
                                                var v85 uint16 = v5
                                                _ = len(v2)
                                                _ = fi65
                                                _ = v85
                                            }
                                        }
                                        var v86 string = v1
                                        _ = v86
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            fmt.Println(v1)
                                        default:
                                            var v87 float32 = 0.1
                                            _ = v87
                                        }
                                    }
                                    v88 := GenS4{F0: v5}
                                    _ = v88
                                }
                            } else {
                                v3[v1] = ((v1 + (v1 + "")) + "")
                            }
                            _ = fi63
                        }
                    }
                    _ = v60
                }
            }
            for fi89 := 0; fi89 < 7; fi89++ {
                fmt.Println(v5, v1, v19)
                _ = fi89
            }
        }
    default:
        v18 = (v18 << ((v23) & 63))
    }
    v19 -= v19
    if false {
        var v90 uint16 = (v5 >> ((v23) & 63))
        _ = v90
    }
    if (2147483647 != 2) {
        for fi91 := 0; fi91 < 17; fi91++ {
            var v92 uint8 = 114
            for fi93 := 0; fi93 < 6; fi93++ {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v94 := GenS12{F0: (^(v5)), F1: (v92 >> (((v23 / ((v23) | 1))) & 63))}
                    fmt.Println(v7, v6, v6)
                    v95 := GenS2{F0: (^(v23))}
                    _ = v94
                    _ = v95
                } else {
                    if (42 >= 0) {
                        v23 += (v23 & (v23 & v23))
                    }
                    v1 = v1
                }
                _ = fi93
            }
            _ = fi91
            _ = v92
        }
    }
    var v96 bool = (323 >= 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v96
}
