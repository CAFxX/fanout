package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 bool; F1 int; F2 int }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 bool; F1 uint8; F2 uint }
type GenS9 struct { F0 uint64; F1 string; F2 uint64 }
type GenS10 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]bool)
    var v3 uint = 255
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            v1 = 287
            var v6 uint32 = 198
            fmt.Println(v1, fi4, fi4)
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint8 = 33
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: false, F1: 42}
            var v11 bool = true
            var v12 int32 = -490
            var v13 float32 = (0.1 + 0.0)
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v13)
            case "hello":
                fmt.Println(v1, v13)
            }
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi4
    }
    if false {
        v3 -= (v3 >> (((v3 % (((^(v3))) | 1))) & 63))
    }
    var v14 uint32 = 4294967295
    var v15 float64 = 1.5
    for fi16 := 0; fi16 < 5; fi16++ {
        if (uint32(v3) == v14) {
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS1{F0: true, F1: (^(v3))}
                v3 = (^(v3))
                var v19 bool = true
                fmt.Println(v19, fi17)
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            if (-1 >= -707) {
                var v20 float64 = (v15 * v15)
                v2["foo bar"] = (((fi16 << ((v3) & 63)) == fi16) || (3.14159 >= ((1e10 + 0.1) - ((1e10 + 1e10) / 1e10))))
                var v21 int8 = 42
                _ = v20
                _ = v21
            } else {
                v1 -= (^(v1))
                _ = len(v2)
            }
        }
        v22 := v2["hello"]
        _ = fi16
        _ = v22
    }
    v23 := v2[((("" + "a") + ("" + "a")) + ("" + "foo bar"))]
    v24 := GenS2{F0: 4294967295, F1: v1, F2: uint64(v3)}
    if false {
        v23 = true
    } else {
        v25 := []int{-405, 795, -646}
        _ = v25
    }
    v3 -= v3
    var v26 uint = v3
    v26 = 65535
    if (!false) {
        if v23 {
            v14 = v14
            switch 127 {
            case 128:
                switch "" {
                case "hello":
                    switch 2147483647 {
                    case 127:
                        switch ("foo bar" + "") {
                        case "a":
                            if (-154 >= 284) {
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    switch fi27 {
                                    case -2147483648:
                                        switch 9223372036854775807 {
                                        case 373:
                                            if (253 > 48) {
                                                v14 += v14
                                                if ((v14 ^ v14) != v14) {
                                                    var v28 uint64 = 2
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        for fi29 := 0; fi29 < 8; fi29++ {
                                                            v26 = (^(v26))
                                                            _ = fi29
                                                        }
                                                    case "a":
                                                        switch fi27 {
                                                        case 32767:
                                                            v30 := GenS1{F0: ((v26 >> ((v3) & 63)) > (v3 | 65535)), F1: (^(v26))}
                                                            _ = v30
                                                        case 42:
                                                            for fi31 := 0; fi31 < 15; fi31++ {
                                                                v26 -= (^(v3))
                                                                v23 = (!false)
                                                                if ((1 >= int(v1)) || (0 <= -39)) {
                                                                    _ = len(v2)
                                                                    v14 -= (^(v14))
                                                                } else {
                                                                    var v32 int64 = 9223372036854775807
                                                                    _ = v32
                                                                }
                                                                if (fi31 != fi31) {
                                                                    fmt.Println(v23, fi27)
                                                                    fmt.Println(v1, v23, v28)
                                                                    var v33 uint64 = v28
                                                                    _ = v33
                                                                }
                                                                _ = fi31
                                                            }
                                                        }
                                                    case "hello":
                                                        var v34 bool = (!(!((!(true || (-1 >= 0))) || true)))
                                                        _ = v34
                                                    }
                                                    var v35 uint32 = v14
                                                    _ = v28
                                                    _ = v35
                                                }
                                            } else {
                                                v15 -= 0.0
                                                for fi36 := 0; fi36 < 2; fi36++ {
                                                    if (!((4294967295 < 1) && (!("" >= "\n\t")))) {
                                                        var v37 int = fi36
                                                        if true {
                                                            v38 := []int{42, v37, (fi27 - (^(v37)))}
                                                            v39 := GenS1{F0: false, F1: (v3 / ((v26) | 1))}
                                                            var v40 uint64 = 42
                                                            _ = len(v38)
                                                            switch v37 {
                                                            case -670:
                                                                v41 := GenS3{F0: (v1 - v1), F1: ("\n\t" + (("a" + ("foo bar" + "\n\t")) + ""))}
                                                                _ = v41
                                                            case 534:
                                                                v37 -= fi27
                                                            case 1:
                                                                var v42 int16 = 1
                                                                _ = v42
                                                            }
                                                            var v43 uint16 = (v1 << (((v26 + v26)) & 63))
                                                            _ = v38
                                                            _ = v39
                                                            _ = v40
                                                            _ = v43
                                                        }
                                                        _ = v37
                                                    }
                                                    _ = len(v2)
                                                    v44 := GenS1{F0: v23, F1: v26}
                                                    v45 := []int{fi27, (fi27 / ((fi36) | 1)), (^(fi36))}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v15, v3, fi36)
                                                        v23 = false
                                                    default:
                                                        v45 = append(v45, (-(fi36)))
                                                    }
                                                    v46 := GenS4{F0: v23, F1: 374, F2: (v26 - 42)}
                                                    _ = fi36
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                }
                                                switch fi27 {
                                                case -697:
                                                    if true {
                                                        var v47 uint64 = 0
                                                        var v48 int8 = -1
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        for fi49 := 0; fi49 < 9; fi49++ {
                                                            var v50 int8 = 0
                                                            _ = fi49
                                                            _ = v50
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    if v23 {
                                                        v51 := []int{128, (fi27 ^ -482), (fi27 & 4294967295)}
                                                        _ = v51
                                                    } else {
                                                        v14 -= v14
                                                    }
                                                case -748:
                                                    if v23 {
                                                        v52 := v2[("a" + "")]
                                                        _ = len(v2)
                                                        if (18446744073709551615 >= uint64(v26)) {
                                                            var v53 uint = v26
                                                            _ = v53
                                                        }
                                                        _ = v52
                                                    } else {
                                                        if v23 {
                                                            fmt.Println(v15, v26)
                                                        } else {
                                                            if (!v23) {
                                                                v54 := GenS5{F0: (2 != v26), F1: (fi27 * -2147483648), F2: (fi27 ^ 265)}
                                                                v55 := GenS6{F0: (^(v26)), F1: 746, F2: (v1 + (v1 / ((745) | 1)))}
                                                                switch (fi27 & (fi27 >> ((uint(v1)) & 63))) {
                                                                case -1:
                                                                    if (fi27 < fi27) {
                                                                        v14 += (v14 + v14)
                                                                        if (v1 <= v1) {
                                                                            v56 := []int{(fi27 % (((fi27 << (((v26 / ((0) | 1))) & 63))) | 1)), (fi27 >> (((v26 & 2)) & 63)), (fi27 * (fi27 >> ((v26) & 63))), fi27}
                                                                            for fi57 := 0; fi57 < 10; fi57++ {
                                                                                var v58 uint16 = 937
                                                                                _ = fi57
                                                                                _ = v58
                                                                            }
                                                                            _ = v56
                                                                        } else {
                                                                            for fi59 := 0; fi59 < 15; fi59++ {
                                                                                switch (("hello" + (("hello" + "a") + "")) + ("hello" + ("" + "hello"))) {
                                                                                case "\n\t":
                                                                                    v2[""] = true
                                                                                default:
                                                                                    var v60 uint32 = (v14 | v14)
                                                                                    _ = v60
                                                                                }
                                                                                v61 := []int{fi27, fi59, (fi27 ^ (fi27 - (^(fi59))))}
                                                                                if false {
                                                                                    switch fi27 {
                                                                                    case 32768:
                                                                                        delete(v2, "")
                                                                                    case -2147483648:
                                                                                        v62 := GenS3{F0: (v1 ^ (v1 * (v1 & 2))), F1: "\n\t"}
                                                                                        _ = v62
                                                                                    case 259:
                                                                                        v1 = (v1 << ((18446744073709551615) & 63))
                                                                                    default:
                                                                                        var v63 bool = false
                                                                                        _ = v63
                                                                                    }
                                                                                }
                                                                                _ = fi59
                                                                                _ = v61
                                                                            }
                                                                            if v23 {
                                                                                v26 += v3
                                                                            }
                                                                        }
                                                                    }
                                                                case 32768:
                                                                    v15 = v15
                                                                case 4294967295:
                                                                    var v64 uint32 = (v14 - v14)
                                                                    _ = v64
                                                                default:
                                                                    if v23 {
                                                                        if v23 {
                                                                            for fi65 := 0; fi65 < 10; fi65++ {
                                                                                var v66 int64 = 42
                                                                                v2[(("a" + "hello") + ("a" + ("hello" + "\n\t")))] = v23
                                                                                if v23 {
                                                                                    var v67 int = (fi65 & fi65)
                                                                                    v68 := GenS7{F0: v66, F1: (v1 % (((v1 - uint16(v14))) | 1))}
                                                                                    switch (fi65 - fi65) {
                                                                                    case -2147483648:
                                                                                        var v69 string = "hello"
                                                                                        _ = v69
                                                                                    default:
                                                                                        if false {
                                                                                            v1 = v1
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    }
                                                                                    v70 := GenS8{F0: ((v1 << (((v3 + (v3 - (v26 >> ((859) & 63))))) & 63)) >= 42), F1: 42, F2: v3}
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                    _ = v70
                                                                                }
                                                                                switch fi65 {
                                                                                case 128:
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        var v71 float32 = 1.5
                                                                                        _ = v71
                                                                                    case "":
                                                                                        v2["a"] = (!v23)
                                                                                    case "hello":
                                                                                        var v72 int16 = -1
                                                                                        _ = v72
                                                                                    }
                                                                                }
                                                                                _ = fi65
                                                                                _ = v66
                                                                            }
                                                                            fmt.Println(v15, v23, v1)
                                                                            v23 = v23
                                                                        } else {
                                                                            v15 += 0.0
                                                                            v15 = v15
                                                                            if (!(!((fi27 % ((fi27) | 1)) <= fi27))) {
                                                                                var v73 uint16 = (v1 ^ 32767)
                                                                                v74 := GenS9{F0: 1, F1: "", F2: 18446744073709551615}
                                                                                for fi75 := 0; fi75 < 7; fi75++ {
                                                                                    v76 := []int{fi27, fi27, (fi75 / ((42) | 1))}
                                                                                    _ = fi75
                                                                                    _ = v76
                                                                                }
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        v77 := GenS10{F0: (^(v14)), F1: 0}
                                                                        for fi78 := 0; fi78 < 4; fi78++ {
                                                                            fmt.Println(fi78)
                                                                            if (v15 == v15) {
                                                                                var v79 uint = v26
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    fmt.Println(v1, v79)
                                                                                default:
                                                                                    var v80 uint64 = 429
                                                                                    _ = v80
                                                                                }
                                                                                _ = v79
                                                                            }
                                                                            var v81 uint16 = 255
                                                                            _ = fi78
                                                                            _ = v81
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                }
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            var v82 bool = (false && (1 != 196))
                                                            _ = v82
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            v23 = true
                                        }
                                    case 4294967295:
                                        v83 := GenS3{F0: uint16(v14), F1: "a"}
                                        _ = v83
                                    default:
                                        var v84 int64 = -9223372036854775808
                                        _ = v84
                                    }
                                    fmt.Println(v1)
                                    v14 += (v14 - (^(v14)))
                                    _ = fi27
                                }
                            } else {
                                v15 -= (((v15 / v15) - 0.1) * ((0.1 - 3.14159) / v15))
                                fmt.Println(v14, v15, v14)
                            }
                        case "":
                            _ = len(v2)
                        case "foo bar":
                            var v85 uint32 = (v14 % ((v14) | 1))
                            _ = v85
                        }
                    case -872:
                        var v86 bool = false
                        _ = v86
                    }
                case "":
                    var v87 uint64 = 42
                    _ = v87
                }
            case 32768:
                fmt.Println(v1, v1, v3)
            case -509:
                var v88 bool = ((v23 || v23) || true)
                _ = v88
            }
        }
        v89 := GenS1{F0: (235 > 77), F1: v26}
        for fi90 := 0; fi90 < 10; fi90++ {
            v91 := GenS8{F0: true, F1: 42, F2: (v26 % ((uint(v14)) | 1))}
            for fi92 := 0; fi92 < 2; fi92++ {
                if v23 {
                    var v93 string = "foo bar"
                    fmt.Println(v93)
                    _ = v93
                } else {
                    v94 := []int{(-(fi90))}
                    _ = v94
                }
                _ = fi92
            }
            _ = fi90
            _ = v91
        }
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
}
