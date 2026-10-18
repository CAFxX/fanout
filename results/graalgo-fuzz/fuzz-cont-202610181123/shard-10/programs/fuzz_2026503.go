package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32; F1 int16; F2 uint }
type GenS10 struct { F0 uint32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -104
    var v3 uint8 = 239
    v4 := GenS1{F0: v3, F1: 65535}
    v3 += 2
    var v5 int32 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: v3, F1: 342}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: v3, F1: 0}
    switch "foo bar" {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS2{F0: v3, F1: ("\n\t" + "foo bar")}
            var v11 uint8 = 42
            switch fi9 {
            case -9223372036854775808:
                if (true || false) {
                    v12 := []int{-9223372036854775808, (fi9 * fi9)}
                    _ = v12
                } else {
                    fmt.Println(v2)
                    switch "a" {
                    case "foo bar":
                        v13 := GenS3{F0: 255, F1: (v5 | v5)}
                        _ = v13
                    case "\n\t":
                        if true {
                            if (!false) {
                                v5 -= (v5 >> ((0) & 63))
                                var v14 uint = 1
                                v15 := []int{4294967295, 32768, -2147483648, (fi9 << (((v14 / (((v14 * v14)) | 1))) & 63))}
                                _ = v14
                                _ = v15
                            } else {
                                v3 += (v11 % ((v11) | 1))
                            }
                            if (!(!true)) {
                                switch 2147483647 {
                                case 270:
                                    var v16 uint32 = 2147483648
                                    _ = v16
                                case -1:
                                    var v17 uint64 = 42
                                    _ = v17
                                }
                            } else {
                                fmt.Println(v2)
                            }
                        } else {
                            v18 := GenS2{F0: (v3 / ((127) | 1)), F1: "foo bar"}
                            var v19 uint32 = 2
                            v3 -= v11
                            _ = v18
                            _ = v19
                        }
                    default:
                        var v20 int = (v1 >> ((255) & 63))
                        _ = v20
                    }
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case "\n\t":
        var v21 uint64 = 18446744073709551615
        _ = v21
    }
    if true {
        var v22 uint8 = v3
        var v23 int32 = v5
        v24 := []int{-898}
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 int8 = 100
        var v26 bool = false
        _ = v25
        _ = v26
    }
    v27 := []int{v1}
    if (true || true) {
        if false {
            for fi28 := 0; fi28 < 9; fi28++ {
                v5 -= (v5 << ((692) & 63))
                switch "\n\t" {
                case "hello":
                    _ = len(v27)
                default:
                    v2 = v2
                }
                _ = fi28
            }
            fmt.Println(v2, v1)
            switch "a" {
            case "hello":
                switch "a" {
                case "hello":
                    if true {
                        for fi29 := 0; fi29 < 13; fi29++ {
                            v30 := GenS3{F0: v3, F1: (v5 | (v5 & (v5 >> ((4294967295) & 63))))}
                            if false {
                                if true {
                                    var v31 uint8 = (^(v3))
                                    var v32 uint16 = 1
                                    _ = v31
                                    _ = v32
                                } else {
                                    if false {
                                        fmt.Println(v2, v3, v1)
                                        var v33 int16 = 0
                                        switch 32767 {
                                        case 872:
                                            for fi34 := 0; fi34 < 12; fi34++ {
                                                v35 := GenS1{F0: (^(v3)), F1: 42}
                                                switch "foo bar" {
                                                case "":
                                                    switch ("\n\t" + "\n\t") {
                                                    case "a":
                                                        v33 += v33
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if (1.5 < 3.14159) {
                                                                fmt.Println(v33, v3, fi34)
                                                                if true {
                                                                    v36 := GenS2{F0: v3, F1: "a"}
                                                                    _ = v36
                                                                }
                                                                v2 = v2
                                                            }
                                                        }
                                                    default:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v2, v3)
                                                        case "a":
                                                            if len(v27) >= 2 {
                                                                v27 = v27[1:len(v27)]
                                                            }
                                                        default:
                                                            v37 := GenS3{F0: (v3 - 158), F1: (v5 / (((v5 % (((-(v5))) | 1))) | 1))}
                                                            _ = v37
                                                        }
                                                    }
                                                }
                                                _ = fi34
                                                _ = v35
                                            }
                                        case -9223372036854775808:
                                            v38 := GenS4{F0: 2, F1: 9223372036854775807}
                                            _ = v38
                                        default:
                                            var v39 int8 = (v2 ^ 1)
                                            _ = v39
                                        }
                                        _ = v33
                                    }
                                    v2 = (v2 & 100)
                                }
                                _ = len(v27)
                                v5 += 1
                            } else {
                                v2 += (v2 - -128)
                                if (18446744073709551615 == 1) {
                                    if true {
                                        v40 := GenS1{F0: v3, F1: 365}
                                        for fi41 := 0; fi41 < 18; fi41++ {
                                            if true {
                                                if true {
                                                    var v42 float32 = 1e10
                                                    if (1 < -1) {
                                                        _ = len(v27)
                                                        if (false || true) {
                                                            v3 = v3
                                                            var v43 uint64 = 452
                                                            var v44 int = int(v43)
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            v27 = append(v27, 405)
                                                            switch (fi29 * fi29) {
                                                            case 0:
                                                                if len(v27) > 0 {
                                                                    v27[0] = -1
                                                                }
                                                            case 127:
                                                                v27 = append(v27, (fi41 / ((fi41) | 1)))
                                                            case -1:
                                                                v45 := GenS5{F0: 94, F1: -1}
                                                                _ = v45
                                                            }
                                                            var v46 int16 = 1
                                                            _ = v46
                                                        }
                                                        var v47 float64 = 3.14159
                                                        _ = v47
                                                    }
                                                    _ = v42
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(v3, v5)
                                                case "foo bar":
                                                    var v48 uint8 = v3
                                                    _ = v48
                                                }
                                                v1 = (-(fi41))
                                            }
                                            var v49 int16 = 1000
                                            var v50 uint32 = 588
                                            _ = fi41
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v40
                                    }
                                } else {
                                    v1 = v1
                                }
                            }
                            if ((1e10 + 3.14159) > (3.14159 / 3.14159)) {
                                var v51 uint8 = (v3 & (v3 << ((4294967295) & 63)))
                                if len(v27) > 0 {
                                    v27[0] = 995
                                }
                                _ = v51
                            } else {
                                for fi52 := 0; fi52 < 13; fi52++ {
                                    var v53 int32 = v5
                                    var v54 int32 = 1
                                    v54 += (v5 / ((v54) | 1))
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v55 uint = 1
                        _ = v55
                    } else {
                        v3 -= v3
                        for fi56 := 0; fi56 < 6; fi56++ {
                            if (1 == -6) {
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            }
                            _ = fi56
                        }
                    }
                case "a":
                    fmt.Println(v1, v1, v3)
                }
            default:
                switch "foo bar" {
                case "hello":
                    for fi57 := 0; fi57 < 20; fi57++ {
                        fmt.Println(v5)
                        var v58 string = ("\n\t" + "")
                        v59 := GenS6{F0: 754}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                case "foo bar":
                    if true {
                        v2 += v2
                        var v60 int8 = v2
                        switch v1 {
                        case 32767:
                            v61 := GenS1{F0: (v3 * v3), F1: 42}
                            _ = v61
                        case -821:
                            _ = len(v27)
                        case 139:
                            if len(v27) > 0 {
                                v27[0] = -2147483648
                            }
                        }
                        _ = v60
                    } else {
                        v62 := GenS2{F0: (v3 >> ((455) & 63)), F1: "a"}
                        var v63 uint = 65535
                        v64 := GenS7{F0: v2, F1: 0}
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                case "\n\t":
                    _ = len(v27)
                default:
                    v65 := GenS8{F0: 2}
                    _ = v65
                }
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v66 uint = 255
        _ = v66
    case "a":
        v1 = v1
    case "\n\t":
        if false {
            if len(v27) > 0 {
                v27[0] = -9223372036854775808
            }
            v2 -= v2
        } else {
            if (1 != 1) {
                v67 := GenS9{F0: 1, F1: 676, F2: 98}
                for fi68 := 0; fi68 < 13; fi68++ {
                    if false {
                        v3 -= 0
                        v2 = (-(v2))
                        for fi69 := 0; fi69 < 17; fi69++ {
                            var v70 uint8 = v3
                            var v71 uint64 = 18446744073709551615
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                    }
                    v27 = append(v27, 32768)
                    if len(v27) > 0 {
                        v27[0] = (v1 & fi68)
                    }
                    _ = fi68
                }
                switch v1 {
                case 2147483647:
                    if false {
                        v3 += 1
                        switch "" {
                        case "a":
                            fmt.Println(v5)
                        default:
                            for fi72 := 0; fi72 < 3; fi72++ {
                                for fi73 := 0; fi73 < 8; fi73++ {
                                    var v74 float32 = 0.1
                                    v5 -= v5
                                    _ = fi73
                                    _ = v74
                                }
                                _ = fi72
                            }
                        }
                        var v75 string = ("foo bar" + ("" + "\n\t"))
                        _ = v75
                    } else {
                        v5 = (v5 * v5)
                        var v76 uint64 = 0
                        _ = v76
                    }
                case 658:
                    if ((v2 ^ v2) == (-(v2))) {
                        for fi77 := 0; fi77 < 6; fi77++ {
                            v27 = append(v27, (^(fi77)))
                            if true {
                                fmt.Println(v3, fi77, v2)
                            }
                            var v78 float64 = (1e10 - 0.1)
                            _ = fi77
                            _ = v78
                        }
                        fmt.Println(v1)
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                    } else {
                        var v79 int8 = (v2 & v2)
                        if len(v27) > 0 {
                            v27[0] = v1
                        }
                        var v80 int = -1
                        _ = v79
                        _ = v80
                    }
                default:
                    var v81 float64 = 0.1
                    _ = v81
                }
                _ = v67
            } else {
                for fi82 := 0; fi82 < 11; fi82++ {
                    v27 = append(v27, (-(v1)))
                    for fi83 := 0; fi83 < 12; fi83++ {
                        var v84 bool = ("a" != "")
                        v2 += (^(v2))
                        _ = fi83
                        _ = v84
                    }
                    _ = fi82
                }
                switch v1 {
                case 90:
                    v27 = append(v27, (v1 % ((v1) | 1)))
                case -9223372036854775808:
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                case 259:
                    v85 := GenS10{F0: 0}
                    _ = v85
                }
                var v86 uint8 = (v3 << ((120) & 63))
                _ = v86
            }
            v87 := GenS1{F0: 158, F1: 65535}
            for fi88 := 0; fi88 < 16; fi88++ {
                var v89 float64 = -2.25
                _ = fi88
                _ = v89
            }
            _ = v87
        }
    }
    fmt.Println(v1)
    if false {
        var v90 float32 = -2.25
        if (false && (1 == -9223372036854775808)) {
            var v91 int16 = -32768
            _ = v91
        } else {
            v1 -= int(v2)
            var v92 uint = 1
            v5 = int32(v1)
            _ = v92
        }
        var v93 int64 = 42
        _ = v90
        _ = v93
    }
    var v94 int32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v94
}
