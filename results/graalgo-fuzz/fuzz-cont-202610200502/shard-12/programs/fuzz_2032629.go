package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int8; F1 bool; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint64; F2 int8 }

func main() {
    var v1 uint = 1
    v2 := []string{"\n\t", "", "hello", "hello"}
    v3 := make(map[string]uint32)
    if (-2147483648 < 290) {
        _ = len(v2)
        v1 = (v1 ^ 0)
        v1 += 255
    } else {
        var v4 uint16 = uint16(v1)
        _ = v4
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if true {
        _ = len(v2)
        v1 += (v1 ^ (v1 * (v1 / ((v1) | 1))))
    }
    v3["\n\t"] = 837
    if (255 > 515) {
        v5 := GenS1{F0: 448, F1: 65535, F2: v1}
        _ = v5
    }
    switch -757 {
    case 32767:
        if true {
            switch 165 {
            case -990:
                if (!true) {
                    _ = len(v3)
                }
            }
            if (true || (!(false || (939 < -445)))) {
                var v6 int32 = 1
                _ = v6
            } else {
                v1 = v1
                var v7 uint16 = 546
                _ = v7
            }
            _ = len(v3)
        }
    case 32768:
        v1 = 594
    case -476:
        var v8 bool = (1.5 >= 1.5)
        _ = v8
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            v1 = (v1 - 0)
            fmt.Println(v1, fi9)
            _ = fi9
        }
    }
    switch 624 {
    case 2147483647:
        v1 -= 4294967295
    default:
        v1 += v1
    }
    v10 := GenS1{F0: 42, F1: 0, F2: v1}
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            fmt.Println(fi11, v1)
            var v12 int64 = -1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v12
        } else {
            if (494 >= 405) {
                v13 := GenS2{F0: 4294967295, F1: 4294967295, F2: 0}
                v1 = v1
                _ = v13
            }
        }
        _ = fi11
    }
    if (!false) {
        v14 := GenS2{F0: 42, F1: v1, F2: -1}
        fmt.Println(v1)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                for fi16 := 0; fi16 < 17; fi16++ {
                    if len(v2) > 0 {
                        v2[0] = (((("" + "\n\t") + "\n\t") + "\n\t") + "a")
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (-1 >= 557) {
                        switch "a" {
                        case "\n\t":
                            var v17 uint = 18446744073709551615
                            _ = v17
                        case "a":
                            if false {
                                if (false && ((fi15 | fi16) >= fi16)) {
                                    var v18 string = ("foo bar" + "")
                                    switch "" {
                                    case "hello":
                                        if (!false) {
                                            v2 = append(v2, v18)
                                            if false {
                                                v18 = v18
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v1 = (v1 + v1)
                                                    v19 := GenS3{F0: 297, F1: ((v18 + v18) + ("\n\t" + ("a" + "")))}
                                                    _ = v19
                                                }
                                            }
                                            var v20 float32 = ((float32(3.14159) + (-2.25 + float32(1e10))) * (float32((1.5 / 0.1)) / (float32(0.0) + (3.14159 / 1.0))))
                                            _ = len(v2)
                                            v20 -= v20
                                            _ = v20
                                        } else {
                                            fmt.Println(fi16)
                                            v3[(v18 + ("hello" + "hello"))] = 561
                                            _ = len(v2)
                                        }
                                    case "a":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "foo bar":
                                        v1 += 2
                                    default:
                                        v21 := GenS4{F0: 286, F1: false, F2: 42}
                                        _ = v21
                                    }
                                    if ((v1 + v1) >= v1) {
                                        if true {
                                            delete(v3, v18)
                                        } else {
                                            v1 = (v1 % ((65535) | 1))
                                            var v22 int = 0
                                            v22 += 42
                                            _ = v22
                                        }
                                        if (!(0 < 239)) {
                                            if (!(1 != 42)) {
                                                if (!((0 <= 0) && false)) {
                                                    v23 := GenS5{F0: 32768, F1: 18446744073709551615}
                                                    var v24 uint32 = 595
                                                    v1 = 42
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                        } else {
                                            v25 := GenS4{F0: 0, F1: (true && ((^(v1)) > 930)), F2: 0}
                                            var v26 float32 = (0.1 * (0.1 + float32(1e10)))
                                            _ = v25
                                            _ = v26
                                        }
                                    } else {
                                        switch v18 {
                                        case "foo bar":
                                            switch v18 {
                                            case "hello":
                                                v27 := v3[((v18 + v18) + "a")]
                                                _ = v27
                                            default:
                                                var v28 string = "a"
                                                fmt.Println(fi16)
                                                _ = v28
                                            }
                                        case "hello":
                                            var v29 float64 = float64(1.5)
                                            _ = v29
                                        }
                                    }
                                    _ = v18
                                } else {
                                    v30 := GenS2{F0: 42, F1: 4294967295, F2: 2147483647}
                                    switch (fi16 ^ fi15) {
                                    case 2147483647:
                                        fmt.Println(fi16, fi16, fi15)
                                    case -1:
                                        v1 = 1
                                    case 1:
                                        var v31 string = ("foo bar" + (("foo bar" + ("foo bar" + "foo bar")) + "\n\t"))
                                        _ = v31
                                    default:
                                        v2 = append(v2, "foo bar")
                                    }
                                    _ = v30
                                }
                            }
                        case "foo bar":
                            if (-1 < -626) {
                                v2 = append(v2, "\n\t")
                                var v32 int = fi15
                                _ = v32
                            } else {
                                v33 := GenS4{F0: 2, F1: true, F2: -2147483648}
                                v1 = (v1 | v1)
                                v1 += v1
                                if ((!(333 != -9223372036854775808)) || false) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v3["hello"] = 42
                                    case "hello":
                                        fmt.Println(v1)
                                    case "\n\t":
                                        switch -971 {
                                        case 32767:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case -2147483648:
                                            v2 = append(v2, ("" + ("a" + ("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + "a")))))
                                        }
                                        if false {
                                            v1 -= (v1 ^ (v1 - v1))
                                            v1 += (v1 ^ v1)
                                            if true {
                                                var v34 int64 = -938
                                                _ = len(v2)
                                                fmt.Println(v34, v34)
                                                _ = v34
                                            }
                                            fmt.Println(fi15, fi16)
                                        } else {
                                            delete(v3, "\n\t")
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch fi15 {
                                            case 0:
                                                switch "a" {
                                                case "hello":
                                                    var v35 string = "a"
                                                    _ = v35
                                                case "a":
                                                    var v36 string = "\n\t"
                                                    _ = v36
                                                case "\n\t":
                                                    switch (fi15 + (fi16 >> ((v1) & 63))) {
                                                    case 327:
                                                        fmt.Println(fi16)
                                                    case 9223372036854775807:
                                                        var v37 float64 = float64(0.1)
                                                        _ = v37
                                                    case -2147483648:
                                                        v2 = append(v2, "foo bar")
                                                    }
                                                default:
                                                    fmt.Println(v1, v1, fi16)
                                                }
                                            case 32768:
                                                if (1 >= 864) {
                                                    delete(v3, "a")
                                                    var v38 bool = true
                                                    v1 = (v1 | v1)
                                                    fmt.Println(v1, v1)
                                                    v39 := v3[("" + "\n\t")]
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            if (1e10 <= (0.0 / 3.14159)) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v40 uint8 = 124
                                                var v41 uint8 = (v40 * (v40 - (v40 - (v40 & v40))))
                                                _ = v40
                                                _ = v41
                                            } else {
                                                v42 := v3[""]
                                                fmt.Println(v1, fi16, v42)
                                                _ = v42
                                            }
                                            v2 = append(v2, "hello")
                                        }
                                        v1 -= (v1 + v1)
                                    default:
                                        var v43 uint32 = uint32(v1)
                                        _ = v43
                                    }
                                    var v44 float32 = 0.0
                                    _ = len(v2)
                                    _ = v44
                                } else {
                                    fmt.Println(fi16, fi15, fi16)
                                    var v45 bool = (754 > 905)
                                    _ = v45
                                }
                                v1 = 548
                                _ = v33
                            }
                        }
                        switch (fi16 & fi16) {
                        case 1:
                            if true {
                                _ = len(v2)
                            }
                        case 127:
                            var v46 bool = (!(!true))
                            _ = v46
                        case 829:
                            switch ("hello" + ("hello" + "")) {
                            case "":
                                v1 += (^(v1))
                            case "a":
                                v47 := GenS5{F0: 32767, F1: 171}
                                _ = v47
                            default:
                                v48 := v3["\n\t"]
                                _ = v48
                            }
                        default:
                            v49 := GenS6{F0: 2, F1: 2147483647}
                            _ = v49
                        }
                        if false {
                            switch ("hello" + "") {
                            case "foo bar":
                                switch 9223372036854775807 {
                                case 4294967295:
                                    fmt.Println(fi16, fi15, fi16)
                                case -1:
                                    switch (fi15 & -9223372036854775808) {
                                    case 4294967295:
                                        var v50 float64 = (3.14159 * (-2.25 / (0.1 * 0.1)))
                                        _ = v50
                                    case -830:
                                        v51 := GenS6{F0: 4294967295, F1: 37}
                                        if (!false) {
                                            if (-697 <= 2147483647) {
                                                v1 -= v1
                                            } else {
                                                fmt.Println(fi16)
                                            }
                                            if (0 == 1) {
                                                var v52 float32 = 0.0
                                                v52 -= 1.5
                                                var v53 int = -2147483648
                                                if ((((("foo bar" + "foo bar") + "a") + (("" + "\n\t") + "\n\t")) + "") <= "a") {
                                                    if (v1 < (v1 ^ v1)) {
                                                        var v54 int = fi15
                                                        v53 = fi15
                                                        var v55 uint64 = 4294967295
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                } else {
                                                    fmt.Println(v1, v52)
                                                    var v56 uint8 = 42
                                                    _ = v56
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                        _ = v51
                                    case -122:
                                        var v57 int16 = 1000
                                        _ = v57
                                    default:
                                        if (980 == 4294967295) {
                                            fmt.Println(v1, v1, v1)
                                            v1 = v1
                                        }
                                        switch -1 {
                                        case 1:
                                            fmt.Println(fi16, fi15, fi16)
                                        case -9223372036854775808:
                                            v1 = (v1 << ((v1) & 63))
                                        case -1:
                                            if true {
                                                var v58 int8 = 100
                                                if true {
                                                    var v59 uint16 = 599
                                                    var v60 float64 = 3.14159
                                                    if false {
                                                        if ("hello" <= "hello") {
                                                            v61 := GenS7{F0: -1, F1: true, F2: v58}
                                                            switch -967 {
                                                            case 1:
                                                                fmt.Println(v58, v58)
                                                            case 42:
                                                                var v62 float64 = v60
                                                                _ = v62
                                                            }
                                                            var v63 string = "a"
                                                            var v64 uint = (v1 - v1)
                                                            var v65 string = v63
                                                            _ = v61
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        fmt.Println(v58)
                                                        v66 := GenS6{F0: 4294967295, F1: 1}
                                                        _ = v66
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            v1 += 0
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        fmt.Println(fi16, v59)
                                                    }
                                                    var v67 int8 = v58
                                                    _ = v59
                                                    _ = v60
                                                    _ = v67
                                                } else {
                                                    v3[(("foo bar" + ("foo bar" + ("" + "a"))) + "foo bar")] = 2147483648
                                                }
                                                _ = v58
                                            }
                                        default:
                                            v68 := GenS8{F0: ("" + ("a" + "a"))}
                                            _ = v68
                                        }
                                        fmt.Println(fi16, fi16, v1)
                                    }
                                }
                            case "a":
                                var v69 uint32 = 677
                                v1 += (v1 ^ v1)
                                _ = v69
                            case "hello":
                                var v70 int16 = 32767
                                _ = v70
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            fmt.Println(fi16, fi16)
                            var v71 int32 = -1
                            _ = v71
                        }
                        if false {
                            v1 += 0
                        }
                        var v72 float64 = (float64(3.14159) * float64(0.0))
                        _ = v72
                    } else {
                        v1 += (v1 << ((4294967295) & 63))
                        switch "hello" {
                        case "hello":
                            fmt.Println(v1)
                        case "foo bar":
                            var v73 float32 = float32(1e10)
                            _ = v73
                        default:
                            var v74 uint32 = 4294967295
                            _ = v74
                        }
                        switch "foo bar" {
                        case "foo bar":
                            v75 := GenS8{F0: "foo bar"}
                            _ = v75
                        default:
                            if len(v2) > 0 {
                                v2[0] = ("" + "hello")
                            }
                        }
                        v76 := GenS9{F0: (32767 == 65535), F1: 4294967295, F2: 42}
                        v1 += v1
                        _ = v76
                    }
                    _ = fi16
                }
            } else {
                var v77 float64 = (((0.1 + (3.14159 / (0.0 + (3.14159 - ((3.14159 + 1.5) + (1e10 / -2.25)))))) * 0.1) - 1e10)
                v77 = 0.0
                var v78 uint64 = 4294967295
                _ = v77
                _ = v78
            }
            v79 := GenS4{F0: 0, F1: true, F2: 0}
            v1 -= (v1 * (^(v1)))
            _ = fi15
            _ = v79
        }
        v1 = 418
        _ = len(v3)
    }
    for fi80 := 0; fi80 < 1; fi80++ {
        switch 32767 {
        case 9223372036854775807:
            for fi81 := 0; fi81 < 2; fi81++ {
                if false {
                    v3[((("foo bar" + "hello") + "foo bar") + "a")] = uint32(v1)
                    v1 -= (v1 ^ (v1 & v1))
                    var v82 uint64 = 1
                    _ = v82
                } else {
                    var v83 int64 = 306
                    fmt.Println(v1, v83)
                    fmt.Println(fi81, fi80, fi81)
                    _ = v83
                }
                var v84 int16 = 32767
                v84 += v84
                _ = fi81
                _ = v84
            }
        default:
            var v85 int32 = -1
            _ = v85
        }
        for fi86 := 0; fi86 < 9; fi86++ {
            v1 = (v1 ^ v1)
            v3["\n\t"] = 209
            var v87 string = ("\n\t" + "hello")
            _ = fi86
            _ = v87
        }
        _ = fi80
    }
    if true {
        if false {
            fmt.Println(v1)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v88 := GenS3{F0: 1, F1: ("foo bar" + ((("\n\t" + "hello") + "hello") + "\n\t"))}
            _ = v88
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
