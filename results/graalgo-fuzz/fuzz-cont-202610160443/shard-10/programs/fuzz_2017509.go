package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 string; F2 int8 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 int; F2 uint64 }
type GenS11 struct { F0 int32; F1 int64; F2 int }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 uint16; F1 int }
type GenS14 struct { F0 uint8; F1 bool; F2 bool }
func genFn0() uint8 {
    var v6 float32 = (1e10 - 0.0)
    if true {
        var v7 uint16 = 1
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(v6, v6)
            switch fi8 {
            case 4294967295:
                v9 := []int{}
                _ = v9
            }
            var v10 bool = (1.5 >= v6)
            _ = fi8
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 string = "foo bar"
            switch fi11 {
            case 128:
                var v13 string = v12
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = v7
    }
    v14 := []int{32767, -23}
    v6 = 1.5
    _ = v14
    _ = v6
    return 1
}
func genFn1() uint16 {
    v15 := []int{127, -1, 18}
    if (65535 < 42) {
        var v16 uint8 = 2
        _ = v16
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int32 = 42
            _ = fi17
            _ = v18
        }
        if len(v15) > 0 {
            v15[0] = -989
        }
    }
    var v19 bool = (1e10 <= -2.25)
    v15 = append(v15, 32767)
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS1{F0: 322, F1: 65535}
        _ = fi20
        _ = v21
    }
    _ = v15
    _ = v19
    return 1
}

func main() {
    var v1 int8 = 127
    v2 := []string{"", ("foo bar" + ""), "foo bar"}
    var v3 bool = true
    var v4 uint8 = 165
    var v5 int16 = 380
    if (255 >= 32767) {
        fmt.Println(v3, v4, v3)
        fmt.Println(v4, v3)
        v22 := GenS1{F0: 65535, F1: 2}
        _ = v22
    } else {
        switch 0 {
        case -930:
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS1{F0: 771, F1: 648}
                _ = fi23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            v5 = v5
            switch -676 {
            case 127:
                _ = len(v2)
            default:
                if false {
                    switch fi25 {
                    case -492:
                        v1 = v1
                    case -9223372036854775808:
                        var v26 int32 = 2147483647
                        _ = v26
                    }
                    var v27 int64 = 9223372036854775807
                    _ = v27
                }
            }
            v3 = (!true)
            _ = fi25
        }
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 int32 = -979
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: 65535}
    v31 := GenS2{F0: v3, F1: 188}
    if (("foo bar" + ("\n\t" + "")) <= ("\n\t" + "")) {
        v3 = false
        if (!((4294967295 > 18446744073709551615) || v3)) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            switch "\n\t" {
            case "":
                if (726 == 65535) {
                    var v32 uint32 = 42
                    fmt.Println(v1, v32, v4)
                    _ = v32
                } else {
                    fmt.Println(v3)
                    if (42 >= 65535) {
                        switch 540 {
                        case 0:
                            for fi33 := 0; fi33 < 13; fi33++ {
                                var v34 uint = 2
                                if ((fi33 <= fi33) && false) {
                                    var v35 uint32 = 619
                                    _ = v35
                                } else {
                                    var v36 int64 = 0
                                    if (!false) {
                                        v3 = false
                                        v37 := GenS3{F0: (v34 * 18446744073709551615), F1: fi33, F2: 32767}
                                        for fi38 := 0; fi38 < 4; fi38++ {
                                            v39 := GenS4{F0: 2147483647}
                                            v5 += -1
                                            v34 = (v34 * 336)
                                            _ = fi38
                                            _ = v39
                                        }
                                        _ = v37
                                    } else {
                                        if (!(false || (!v3))) {
                                            fmt.Println(v36, v34)
                                            if v3 {
                                                var v40 bool = (!(!true))
                                                _ = v40
                                            } else {
                                                switch fi33 {
                                                case 32767:
                                                    for fi41 := 0; fi41 < 9; fi41++ {
                                                        var v42 int32 = 42
                                                        _ = fi41
                                                        _ = v42
                                                    }
                                                case -804:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 4294967295:
                                                    switch (fi33 ^ 2) {
                                                    case -9223372036854775808:
                                                        var v43 uint8 = v4
                                                        _ = v43
                                                    case 4294967295:
                                                        for fi44 := 0; fi44 < 16; fi44++ {
                                                            v45 := GenS1{F0: 0, F1: 2}
                                                            _ = fi44
                                                            _ = v45
                                                        }
                                                    }
                                                default:
                                                    if (745 <= 65535) {
                                                        for fi46 := 0; fi46 < 2; fi46++ {
                                                            v5 = 950
                                                            var v47 int64 = v36
                                                            if (v3 && ((!v3) && false)) {
                                                                v48 := GenS5{F0: 887, F1: ((("hello" + "\n\t") + "foo bar") + "foo bar"), F2: v1}
                                                                v2 = append(v2, "foo bar")
                                                                _ = v48
                                                            } else {
                                                                _ = len(v2)
                                                                var v49 uint64 = 18446744073709551615
                                                                _ = v49
                                                            }
                                                            v50 := GenS3{F0: 255, F1: (fi46 | fi33), F2: 2}
                                                            v5 -= v5
                                                            _ = fi46
                                                            _ = v47
                                                            _ = v50
                                                        }
                                                    } else {
                                                        for fi51 := 0; fi51 < 3; fi51++ {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v3 = false
                                                            case "a":
                                                                if v3 {
                                                                    if ((-9223372036854775808 >= int(v36)) || (!v3)) {
                                                                        v52 := GenS5{F0: 2, F1: "foo bar", F2: (v1 | v1)}
                                                                        v5 -= v5
                                                                        _ = v52
                                                                    } else {
                                                                        if v3 {
                                                                            v3 = ((!(((0.1 + (((1.5 * 0.0) + ((1e10 / 1e10) - 3.14159)) * (((1.5 * 1e10) + 1e10) + ((float64(float32(float64(0.1))) - -2.25) * -2.25)))) - (1e10 + -2.25)) > (0.0 + -2.25))) && v3)
                                                                            var v53 int16 = v5
                                                                            v54 := GenS6{F0: "hello", F1: (-(v5))}
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                        switch ("hello" + "a") {
                                                                        case "":
                                                                            var v55 uint64 = 4294967295
                                                                            _ = v55
                                                                        default:
                                                                            var v56 uint8 = (v4 | genFn0())
                                                                            _ = v56
                                                                        }
                                                                        var v57 float64 = 1.5
                                                                        _ = v57
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v58 := GenS7{F0: ""}
                                                                    _ = v58
                                                                }
                                                            }
                                                            if true {
                                                                v59 := GenS7{F0: ("foo bar" + "foo bar")}
                                                                _ = v59
                                                            } else {
                                                                v3 = ("hello" != (("\n\t" + "\n\t") + "foo bar"))
                                                                v60 := GenS8{F0: 42, F1: 42}
                                                                var v61 int64 = (v36 >> (((^(v34))) & 63))
                                                                fmt.Println(v61)
                                                                if ((-2.25 > (0.0 - -2.25)) || (2 > 42)) {
                                                                    var v62 float32 = 0.0
                                                                    _ = v62
                                                                }
                                                                fmt.Println(v34, v36)
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = "a"
                                                            }
                                                            _ = fi51
                                                        }
                                                        for fi63 := 0; fi63 < 20; fi63++ {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch (fi33 - (fi33 << (((v34 ^ (^(v34)))) & 63))) {
                                                            case 986:
                                                                v36 += (v36 >> ((635) & 63))
                                                            default:
                                                                v3 = v3
                                                                var v64 float64 = ((3.14159 / 1.0) + (float64(3.14159) + 0.1))
                                                                v65 := GenS9{F0: (v4 | 42)}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            v66 := GenS6{F0: ((("foo bar" + "a") + "") + (("hello" + "foo bar") + "foo bar")), F1: 1000}
                                                            var v67 string = "foo bar"
                                                            v3 = v3
                                                            v4 = v4
                                                            _ = fi63
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                                v68 := GenS10{F0: (fi33 & 2147483647), F1: (fi33 * -461), F2: 18446744073709551615}
                                                fmt.Println(v1, v1)
                                                _ = v68
                                            }
                                        } else {
                                            v34 += (v34 >> ((65535) & 63))
                                            switch fi33 {
                                            case 32768:
                                                switch ("\n\t" + "a") {
                                                case "":
                                                    _ = len(v2)
                                                }
                                            case 341:
                                                for fi69 := 0; fi69 < 4; fi69++ {
                                                    v34 = v34
                                                    _ = fi69
                                                }
                                            }
                                        }
                                    }
                                    _ = v36
                                }
                                _ = fi33
                                _ = v34
                            }
                        case 141:
                            _ = len(v2)
                        default:
                            v70 := GenS6{F0: "a", F1: (v5 - (v5 - 42))}
                            _ = v70
                        }
                        if false {
                            if (97 >= v1) {
                                v71 := GenS3{F0: 121, F1: 870, F2: 65535}
                                if (1 < 407) {
                                    var v72 int64 = 1
                                    _ = v72
                                } else {
                                    v4 = 42
                                    v73 := GenS1{F0: 793, F1: 1000}
                                    _ = v73
                                }
                                fmt.Println(v1, v5, v1)
                                _ = v71
                            }
                        }
                        _ = len(v2)
                    }
                    v1 = (^(v1))
                }
            case "hello":
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if v3 {
        var v74 uint64 = 4294967295
        v74 = (v74 / ((2) | 1))
        for fi75 := 0; fi75 < 19; fi75++ {
            for fi76 := 0; fi76 < 6; fi76++ {
                if v3 {
                    if (!(true || ("" != ""))) {
                        v77 := GenS11{F0: -302, F1: -9223372036854775808, F2: fi75}
                        var v78 uint16 = 426
                        _ = v77
                        _ = v78
                    } else {
                        v79 := GenS12{F0: (v5 >> ((747) & 63)), F1: false}
                        if (65535 != 1) {
                            switch "a" {
                            case "hello":
                                v1 += v1
                            default:
                                fmt.Println(v1, v3)
                            }
                            fmt.Println(fi76, v1)
                            var v80 uint = 65535
                            _ = v80
                        } else {
                            fmt.Println(v4, fi76, fi76)
                            if (!true) {
                                var v81 bool = ((true || true) && v3)
                                _ = v81
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v79
                    }
                } else {
                    switch (fi75 >> ((uint(v4)) & 63)) {
                    case 1:
                        var v82 float64 = 1e10
                        _ = v82
                    case 2147483647:
                        fmt.Println(fi76, fi75, v3)
                    case 105:
                        if (!false) {
                            if false {
                                v3 = v3
                                if v3 {
                                    v5 = v5
                                    fmt.Println(v1, v3)
                                    switch -9223372036854775808 {
                                    case 128:
                                        var v83 int32 = 806
                                        _ = v83
                                    case 127:
                                        if v3 {
                                            v74 += (v74 ^ (v74 | (v74 << ((255) & 63))))
                                            var v84 uint32 = 0
                                            if (fi75 >= fi75) {
                                                v85 := GenS11{F0: 0, F1: -574, F2: fi76}
                                                _ = v85
                                            }
                                            _ = v84
                                        }
                                    case 556:
                                        fmt.Println(v3, fi76)
                                    }
                                }
                            }
                        }
                    }
                }
                _ = fi76
            }
            v5 -= -1
            _ = fi75
        }
        _ = v74
    }
    v1 = v1
    fmt.Println(v4, v3)
    v86 := GenS12{F0: v5, F1: false}
    for fi87 := 0; fi87 < 5; fi87++ {
        switch "foo bar" {
        case "":
            v88 := GenS13{F0: 65535, F1: fi87}
            _ = v88
        case "foo bar":
            v89 := GenS14{F0: (v4 + 0), F1: true, F2: (42 < (v5 >> ((0) & 63)))}
            _ = v89
        }
        _ = fi87
    }
    for fi90 := 0; fi90 < 8; fi90++ {
        v4 -= genFn0()
        v91 := GenS14{F0: 64, F1: true, F2: false}
        v3 = ((v4 * 2) <= genFn0())
        _ = fi90
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v86
}
