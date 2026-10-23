package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }
type GenS7 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint16; F1 int8; F2 uint16 }

func main() {
    var v1 uint = 420
    var v2 uint8 = 42
    var v3 int64 = 365
    var v4 uint = 42
    var v5 uint8 = 2
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 = v5
        _ = fi7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            fmt.Println(v1, v3)
            v5 = (v5 + (v2 - (v5 << (((v4 & (v4 / ((18446744073709551615) | 1)))) & 63))))
        } else {
            v3 -= 611
        }
        v2 -= v5
        var v9 int8 = 42
        _ = fi8
        _ = v9
    }
    v10 := []int{4294967295}
    switch (("a" + (("\n\t" + "foo bar") + "a")) + "hello") {
    case "\n\t":
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(fi11, fi11, v4)
            fmt.Println(v3)
            _ = fi11
        }
    }
    switch -688 {
    case -9223372036854775808:
        v12 := GenS1{F0: 0}
        _ = v12
    case 790:
        if false {
            v13 := GenS2{F0: 0}
            v4 = v1
            _ = v13
        } else {
            v14 := GenS3{F0: v1, F1: 1, F2: -713}
            _ = v14
        }
    case 4294967295:
        v15 := GenS2{F0: 42}
        _ = v15
    }
    if false {
        v10 = append(v10, 9223372036854775807)
        for fi16 := 0; fi16 < 15; fi16++ {
            v1 = 4294967295
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi16
        }
        _ = len(v10)
    }
    var v17 int64 = (v3 + (v3 - -9223372036854775808))
    v5 = (^(v5))
    v18 := GenS4{F0: (^(v5)), F1: 0, F2: v4}
    switch "a" {
    case "\n\t":
        switch "hello" {
        case "":
            v19 := GenS1{F0: 702}
            _ = v19
        case "\n\t":
            var v20 float64 = float64(3.14159)
            _ = v20
        case "a":
            v21 := GenS2{F0: int8(v2)}
            _ = v21
        default:
            if false {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 float64 = 1.5
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v25 := GenS1{F0: 32767}
                        v1 = v1
                        if (1.5 > (1e10 - -2.25)) {
                            v23 = v23
                        } else {
                            if len(v10) > 0 {
                                v10[0] = fi24
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
            }
        }
    case "a":
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 float32 = 3.14159
            for fi28 := 0; fi28 < 3; fi28++ {
                if true {
                    var v29 float32 = (float32(0.0) * v27)
                    _ = v29
                }
                var v30 uint = 0
                var v31 int8 = -1
                var v32 uint64 = 42
                if len(v10) > 0 {
                    v10[0] = fi28
                }
                _ = fi28
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi26
            _ = v27
        }
    case "":
        switch ("hello" + ("\n\t" + "")) {
        case "foo bar":
            var v33 int16 = 1
            _ = v33
        case "\n\t":
            v10 = append(v10, 128)
        case "a":
            if (65535 <= 544) {
                switch 127 {
                case 127:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        var v35 int16 = 1
                        _ = fi34
                        _ = v35
                    }
                case -946:
                    fmt.Println(v3, v4)
                case 729:
                    fmt.Println(v2)
                default:
                    for fi36 := 0; fi36 < 2; fi36++ {
                        v37 := GenS3{F0: 1, F1: 997, F2: 42}
                        fmt.Println(v3, fi36)
                        _ = fi36
                        _ = v37
                    }
                }
                var v38 string = (("" + ("a" + "")) + "\n\t")
                _ = v38
            } else {
                for fi39 := 0; fi39 < 18; fi39++ {
                    for fi40 := 0; fi40 < 1; fi40++ {
                        v10 = append(v10, fi39)
                        fmt.Println(v1, v4)
                        _ = fi40
                    }
                    for fi41 := 0; fi41 < 9; fi41++ {
                        var v42 uint16 = 2
                        v4 += v1
                        var v43 int16 = 0
                        if (1e10 < 0.1) {
                            if (((("foo bar" + "") + "a") + "foo bar") < "hello") {
                                fmt.Println(v42, v43, v4)
                            } else {
                                fmt.Println(fi41, v1, fi41)
                            }
                            switch "" {
                            case "\n\t":
                                switch ("\n\t" + "foo bar") {
                                case "a":
                                    v42 = 32767
                                case "foo bar":
                                    v44 := GenS4{F0: (v2 << ((1) & 63)), F1: (-(v43)), F2: v1}
                                    _ = v44
                                case "hello":
                                    v3 += (^(v17))
                                default:
                                    v43 = (v43 | (v43 & 1000))
                                }
                            }
                            switch (fi41 | 32768) {
                            case 9223372036854775807:
                                _ = len(v10)
                            case 0:
                                if len(v10) > 0 {
                                    v10[0] = 42
                                }
                            }
                        } else {
                            var v45 int64 = v17
                            _ = v45
                        }
                        var v46 int64 = v3
                        switch "hello" {
                        case "":
                            if ((true && ((860 > 1) && (!false))) && (!true)) {
                                v4 += v4
                                v5 = 255
                                v43 = int16(v46)
                            } else {
                                v46 = (v17 ^ (^(v17)))
                                if (!(v3 >= (v3 % ((int64(v1)) | 1)))) {
                                    v47 := GenS5{F0: -2147483648, F1: 597}
                                    var v48 bool = true
                                    if v48 {
                                        v1 = v4
                                        v49 := GenS6{F0: int8(fi41), F1: (v3 ^ v46), F2: 2}
                                        _ = v49
                                    } else {
                                        var v50 uint8 = (v2 % (((v5 - (v5 / ((v2) | 1)))) | 1))
                                        if (!((fi41 >> (((v4 << ((v1) & 63))) & 63)) >= int(v3))) {
                                            v42 += (^(v42))
                                            switch fi41 {
                                            case 878:
                                                var v51 int8 = 127
                                                _ = v51
                                            }
                                        } else {
                                            v43 -= (v43 * (v43 & (-(v43))))
                                            fmt.Println(v42, v46, v2)
                                        }
                                        if true {
                                            if v48 {
                                                v52 := GenS6{F0: 0, F1: v46, F2: uint64(v50)}
                                                _ = v52
                                            }
                                            v53 := GenS7{F0: 0, F1: 75, F2: 1}
                                            _ = v53
                                        } else {
                                            if true {
                                                fmt.Println(v5)
                                                if (int(v43) >= fi41) {
                                                    v54 := GenS6{F0: 100, F1: v17, F2: 2}
                                                    var v55 string = "\n\t"
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if (!((0.0 - (1.5 / (-2.25 / -2.25))) < 1.5)) {
                                                    var v56 uint8 = 255
                                                    _ = v56
                                                }
                                            }
                                            var v57 string = ""
                                            v10 = append(v10, fi39)
                                            _ = v57
                                        }
                                        fmt.Println(v5, fi41)
                                        var v58 uint32 = 65535
                                        _ = v50
                                        _ = v58
                                    }
                                    if true {
                                        if v48 {
                                            var v59 uint8 = (v2 / (((v5 & (^(v5)))) | 1))
                                            _ = v59
                                        }
                                        var v60 uint = 0
                                        var v61 string = ("\n\t" + ("" + "foo bar"))
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v46 -= 1
                                        if v48 {
                                            var v62 int = (^(fi41))
                                            switch 1 {
                                            case -1:
                                                fmt.Println(fi41, v17, fi41)
                                            }
                                            switch (v62 % ((int(v42)) | 1)) {
                                            case 1:
                                                var v63 uint16 = v42
                                                _ = v63
                                            case 454:
                                                v10 = append(v10, -9223372036854775808)
                                            default:
                                                if v48 {
                                                    v64 := GenS8{F0: (fi39 % (((^(fi39))) | 1)), F1: 0}
                                                    _ = v64
                                                } else {
                                                    v1 += 37
                                                    v1 = (^(v4))
                                                }
                                            }
                                            _ = v62
                                        } else {
                                            var v65 int32 = 660
                                            _ = v65
                                        }
                                    }
                                    _ = v47
                                    _ = v48
                                } else {
                                    v42 = v42
                                    if (v43 > -1) {
                                        v17 = v46
                                        switch (fi41 << ((4294967295) & 63)) {
                                        case -95:
                                            var v66 uint8 = 2
                                            var v67 bool = true
                                            v68 := GenS7{F0: (v46 >> ((v4) & 63)), F1: 42, F2: v2}
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        default:
                                            fmt.Println(v4)
                                        }
                                        if true {
                                            var v69 uint8 = (v2 << ((266) & 63))
                                            _ = v69
                                        } else {
                                            _ = len(v10)
                                            var v70 float32 = ((1e10 - (-2.25 + 3.14159)) - 1e10)
                                            fmt.Println(v2)
                                            _ = v70
                                        }
                                        var v71 int32 = -100
                                        v72 := GenS9{F0: 42, F1: v43}
                                        _ = v71
                                        _ = v72
                                    } else {
                                        if (!("a" >= ((("hello" + ("a" + "\n\t")) + ("\n\t" + ("foo bar" + "\n\t"))) + ((("a" + "") + ("foo bar" + "a")) + ("hello" + ""))))) {
                                            if (502 <= 2147483647) {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                v3 = 9223372036854775807
                                            } else {
                                                v10 = append(v10, (fi41 + (fi41 & (^(fi39)))))
                                                if (fi41 < 127) {
                                                    if len(v10) > 0 {
                                                        v10[0] = fi39
                                                    }
                                                    var v73 float64 = (1.5 - float64(0.0))
                                                    v17 += -229
                                                    _ = v73
                                                }
                                                var v74 int64 = 1
                                                _ = v74
                                            }
                                        }
                                        if true {
                                            v4 += v1
                                            v3 -= v3
                                        }
                                        var v75 string = "\n\t"
                                        var v76 int = 2147483647
                                        if ((!false) && false) {
                                            fmt.Println(v17)
                                            v77 := GenS4{F0: (v5 * v5), F1: -341, F2: (v4 - v4)}
                                            v78 := GenS3{F0: (^(v1)), F1: 2147483648, F2: v43}
                                            _ = v77
                                            _ = v78
                                        } else {
                                            var v79 uint64 = 18446744073709551615
                                            _ = v79
                                        }
                                        _ = v75
                                        _ = v76
                                    }
                                }
                                var v80 float32 = -2.25
                                _ = v80
                            }
                            var v81 float32 = (float32((((0.0 - 1e10) * (-2.25 + 1.5)) / (3.14159 + -2.25))) * (-2.25 - -2.25))
                            _ = v81
                        case "\n\t":
                            if false {
                                v5 = (^(v2))
                                fmt.Println(v46, v1, v46)
                                var v82 uint16 = 633
                                _ = v82
                            }
                        }
                        v2 -= 42
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v46
                    }
                    _ = fi39
                }
                v17 = (v3 ^ (-(v17)))
                v2 = v2
            }
        }
    default:
        for fi83 := 0; fi83 < 19; fi83++ {
            v84 := GenS10{F0: 231, F1: 127, F2: 1}
            _ = fi83
            _ = v84
        }
    }
    var v85 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
