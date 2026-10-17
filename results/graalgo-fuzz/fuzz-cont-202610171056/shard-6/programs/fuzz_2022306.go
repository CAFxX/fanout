package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 bool = true
    var v2 int32 = 2147483647
    v3 := []uint32{489, 2, 352}
    var v4 int = -2147483648
    v4 += 9223372036854775807
    if (699 < 867) {
        v5 := GenS1{F0: "hello"}
        var v6 uint32 = 42
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(v1, v4, v4)
            var v8 uint64 = 18446744073709551615
            switch fi7 {
            case 2147483647:
                if (683 > uint16(v8)) {
                    v9 := GenS1{F0: "\n\t"}
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if (false && v1) {
                            v4 += -1
                            _ = len(v3)
                            switch v4 {
                            case -568:
                                var v11 int16 = 32767
                                _ = v11
                            case 128:
                                v3 = append(v3, 1)
                                if (65535 < uint16(v8)) {
                                    fmt.Println(v1)
                                    switch (v4 * -2147483648) {
                                    case -2147483648:
                                        v8 = v8
                                    case 128:
                                        var v12 bool = true
                                        _ = v12
                                    case 1:
                                        var v13 int = (^(v4))
                                        _ = v13
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 835
                                        }
                                    }
                                } else {
                                    var v14 float64 = 3.14159
                                    v4 = 765
                                    _ = v14
                                }
                            case 2147483647:
                                fmt.Println(v1, v1, fi10)
                            default:
                                v8 = 950
                            }
                            v2 += v2
                            var v15 float64 = 3.14159
                            _ = v15
                        } else {
                            fmt.Println(fi7)
                        }
                        v1 = v1
                        if (!(-1 > -591)) {
                            var v16 int = (fi10 * v4)
                            var v17 uint32 = 207
                            var v18 int16 = 425
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        _ = fi10
                    }
                    v4 += (fi7 >> ((65535) & 63))
                    _ = v9
                } else {
                    var v19 uint32 = 65535
                    if ((^(v2)) < (v2 * -180)) {
                        if false {
                            v19 = (v19 - (v19 * v19))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch ("" + "a") {
                        case "\n\t":
                            v2 = v2
                        case "":
                            switch -9223372036854775808 {
                            case 32768:
                                fmt.Println(v4, v2, v2)
                            default:
                                var v20 uint32 = 2147483648
                                _ = v20
                            }
                        default:
                            switch v4 {
                            case 4294967295:
                                if (-2.25 >= -2.25) {
                                    if v1 {
                                        var v21 float32 = 1e10
                                        v22 := GenS2{F0: -1}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v23 := GenS2{F0: -9223372036854775808}
                                        _ = v23
                                    }
                                    var v24 uint = 984
                                    _ = v24
                                }
                            case 1:
                                v25 := GenS3{F0: 128, F1: 198, F2: 1}
                                _ = v25
                            }
                        }
                    } else {
                        if v1 {
                            if true {
                                _ = len(v3)
                            } else {
                                _ = len(v3)
                                if v1 {
                                    v8 += v8
                                }
                                switch fi7 {
                                case -808:
                                    if len(v3) > 0 {
                                        v3[0] = v19
                                    }
                                case -903:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        for fi26 := 0; fi26 < 8; fi26++ {
                                            if false {
                                                if ((v19 & v19) >= v19) {
                                                    v3 = append(v3, 2147483647)
                                                    if false {
                                                        if (-9223372036854775808 < -1) {
                                                            v27 := GenS1{F0: ""}
                                                            v8 = v8
                                                            _ = v27
                                                        }
                                                    }
                                                    var v28 uint8 = 127
                                                    _ = v28
                                                }
                                            } else {
                                                v29 := GenS2{F0: 0}
                                                fmt.Println(v1, v8)
                                                _ = v29
                                            }
                                            if false {
                                                fmt.Println(v4, v2, fi7)
                                                v30 := GenS4{F0: 242, F1: 42, F2: v2}
                                                _ = v30
                                            }
                                            _ = fi26
                                        }
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "a":
                                            var v31 float32 = 0.1
                                            _ = v31
                                        case "":
                                            for fi32 := 0; fi32 < 5; fi32++ {
                                                v2 = v2
                                                fmt.Println(v19)
                                                _ = fi32
                                            }
                                        case "\n\t":
                                            var v33 float32 = 1.5
                                            _ = v33
                                        default:
                                            for fi34 := 0; fi34 < 2; fi34++ {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = fi34
                                            }
                                        }
                                    case "a":
                                        v8 = v8
                                    }
                                }
                            }
                            for fi35 := 0; fi35 < 1; fi35++ {
                                var v36 int8 = -63
                                v4 += v4
                                _ = fi35
                                _ = v36
                            }
                        } else {
                            var v37 int16 = -1
                            var v38 uint = 1
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v2, v8)
                            default:
                                var v39 float64 = float64(1.5)
                                _ = v39
                            }
                            _ = v37
                            _ = v38
                        }
                        for fi40 := 0; fi40 < 8; fi40++ {
                            if len(v3) > 0 {
                                v3[0] = v19
                            }
                            switch int(v19) {
                            case 32768:
                                var v41 int64 = 9223372036854775807
                                _ = v41
                            default:
                                v4 -= fi7
                            }
                            var v42 uint64 = v8
                            _ = fi40
                            _ = v42
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v19
                }
            case -14:
                var v43 string = "hello"
                _ = v43
            case 42:
                var v44 uint = 0
                _ = v44
            default:
                switch ("\n\t" + "foo bar") {
                case "":
                    switch fi7 {
                    case 128:
                        v1 = true
                    case 32767:
                        var v45 uint32 = 2
                        _ = v45
                    case -2147483648:
                        if false {
                            if (((0.0 * 1.5) < (float32(1.5) + 1e10)) && false) {
                                switch "foo bar" {
                                case "a":
                                    if (!(39 <= 0)) {
                                        _ = len(v3)
                                        v46 := GenS2{F0: 42}
                                        if ((!v1) || (v1 || (fi7 > fi7))) {
                                            if (-1 == 1000) {
                                                for fi47 := 0; fi47 < 7; fi47++ {
                                                    v48 := GenS3{F0: 2, F1: 42, F2: (v4 ^ fi47)}
                                                    v49 := GenS4{F0: 42, F1: 0, F2: (v2 / ((v2) | 1))}
                                                    var v50 float32 = 0.0
                                                    _ = fi47
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                                var v51 float64 = ((0.0 + ((0.0 / 1.0) * float64(1.5))) * 0.0)
                                                v2 += v2
                                                _ = v51
                                            }
                                            switch ("foo bar" + "foo bar") {
                                            case "\n\t":
                                                switch -309 {
                                                case 2147483647:
                                                    v2 = v2
                                                case 609:
                                                    switch "hello" {
                                                    case "a":
                                                        v52 := GenS2{F0: 1}
                                                        _ = v52
                                                    }
                                                default:
                                                    v4 = 4294967295
                                                }
                                            case "a":
                                                v53 := GenS5{F0: 1, F1: 255, F2: 42}
                                                _ = v53
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            }
                                        } else {
                                            if (v1 || ((42 >= 578) && (!v1))) {
                                                if ((1 >= 85) && v1) {
                                                    for fi54 := 0; fi54 < 11; fi54++ {
                                                        v8 += (v8 >> ((2) & 63))
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483647
                                                        }
                                                        var v55 float32 = ((-2.25 * 1e10) * 1.5)
                                                        v4 += 2147483647
                                                        _ = fi54
                                                        _ = v55
                                                    }
                                                    fmt.Println(v1, v2, fi7)
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    fmt.Println(v1)
                                                }
                                                v1 = false
                                                for fi56 := 0; fi56 < 4; fi56++ {
                                                    fmt.Println(v1, fi7, fi7)
                                                    _ = fi56
                                                }
                                            } else {
                                                v1 = v1
                                                for fi57 := 0; fi57 < 12; fi57++ {
                                                    v58 := GenS5{F0: 127, F1: 949, F2: 783}
                                                    if true {
                                                        var v59 uint64 = (v8 << ((2) & 63))
                                                        switch "hello" {
                                                        case "":
                                                            v3 = append(v3, 948)
                                                        case "hello":
                                                            v4 -= 32767
                                                        }
                                                        _ = v59
                                                    }
                                                    v4 += 32767
                                                    _ = fi57
                                                    _ = v58
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    if len(v3) > 0 {
                                                        v3[0] = 2147483648
                                                    }
                                                default:
                                                    for fi60 := 0; fi60 < 5; fi60++ {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = fi60
                                                    }
                                                }
                                            }
                                            if v1 {
                                                if (-121 <= -1) {
                                                    v61 := GenS6{F0: 866, F1: 65535}
                                                    v62 := GenS7{F0: v2}
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 bool = (4294967295 >= 42)
                                                _ = v63
                                            } else {
                                                fmt.Println(v4)
                                            }
                                            v64 := GenS6{F0: 255, F1: 1}
                                            _ = v64
                                        }
                                        _ = v46
                                    } else {
                                        for fi65 := 0; fi65 < 13; fi65++ {
                                            var v66 int16 = 1000
                                            _ = fi65
                                            _ = v66
                                        }
                                    }
                                }
                                v1 = (-9223372036854775808 == -529)
                            }
                            v1 = (false && (v1 || true))
                            for fi67 := 0; fi67 < 20; fi67++ {
                                v8 = 385
                                _ = fi67
                            }
                        }
                    default:
                        v68 := GenS6{F0: 391, F1: 767}
                        _ = v68
                    }
                default:
                    if ((!(65535 <= 2)) && (42 > 782)) {
                        var v69 int16 = 93
                        switch ("foo bar" + ("a" + "\n\t")) {
                        case "":
                            v1 = true
                        case "\n\t":
                            fmt.Println(v69, v2)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v69
                    }
                }
            }
            _ = fi7
            _ = v8
        }
        var v70 uint = 192
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = v70
    }
    if v1 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v71 float64 = (((0.1 - (((float64(-2.25) * float64(-2.25)) * (0.1 / 1.0)) * (0.1 - ((-2.25 - 0.0) / 0.1)))) * float64(-2.25)) + (((0.1 + ((float64(0.1) * 1.5) * (0.0 * (1e10 * -2.25)))) + ((1.5 * 1.5) - 3.14159)) + 0.1))
        _ = v71
    }
    var v72 float64 = ((((1e10 + (-2.25 + ((1.5 - 0.0) + ((1.5 / 1.5) - 0.0)))) - ((-2.25 + (float64((1e10 * 1.5)) + 0.0)) * 0.0)) / ((1.5 - (3.14159 + 0.0)) + 3.14159)) / float64((3.14159 + 3.14159)))
    if false {
        var v73 int32 = -1
        v2 -= (v2 << ((65535) & 63))
        _ = v73
    }
    var v74 bool = v1
    for fi75 := 0; fi75 < 19; fi75++ {
        var v76 string = ((("a" + "foo bar") + "") + "hello")
        _ = fi75
        _ = v76
    }
    fmt.Println(v2, v74)
    for fi77 := 0; fi77 < 4; fi77++ {
        v4 -= -251
        v78 := GenS8{F0: 18446744073709551615}
        _ = fi77
        _ = v78
    }
    var v79 uint = 577
    var v80 int32 = -967
    v3 = append(v3, 1)
    if len(v3) > 0 {
        v3[0] = 735
    }
    if v1 {
        switch "a" {
        case "hello":
            var v81 uint16 = 65535
            _ = v81
        case "\n\t":
            for fi82 := 0; fi82 < 3; fi82++ {
                v4 -= 32767
                for fi83 := 0; fi83 < 18; fi83++ {
                    v84 := GenS1{F0: ""}
                    fmt.Println(v72, v1)
                    _ = fi83
                    _ = v84
                }
                _ = fi82
            }
        case "foo bar":
            fmt.Println(v1, v74)
        default:
            fmt.Println(v80)
        }
        fmt.Println(v80, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v74
    _ = v79
    _ = v80
}
