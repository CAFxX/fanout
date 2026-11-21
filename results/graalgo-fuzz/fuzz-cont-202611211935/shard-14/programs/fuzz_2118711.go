package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int8; F1 bool; F2 bool }
type GenS8 struct { F0 int; F1 int32; F2 uint }
type GenS9 struct { F0 int32; F1 uint; F2 uint32 }
type GenS10 struct { F0 uint32; F1 int16; F2 uint8 }

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v4 := []string{"a"}
    var v5 float32 = (((0.0 - 1.5) - 1e10) + 0.1)
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v1, fi7)
        delete(v3, 27)
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                switch (fi7 >> ((2) & 63)) {
                case 2147483647:
                    v4 = append(v4, "")
                case 32768:
                    v9 := GenS1{F0: (fi7 | (fi7 % ((fi7) | 1)))}
                    _ = v9
                case 748:
                    v10 := GenS2{F0: -769}
                    _ = v10
                default:
                    var v11 int64 = 42
                    _ = v11
                }
                v1 += 0
                if false {
                    v1 += 52
                    switch (fi8 % ((fi8) | 1)) {
                    case 4294967295:
                        if true {
                            switch fi7 {
                            case 0:
                                fmt.Println(fi8, v5)
                            case -250:
                                switch -157 {
                                case 1:
                                    v12 := GenS1{F0: (fi8 >> ((4294967295) & 63))}
                                    _ = v12
                                case -376:
                                    v4 = append(v4, "hello")
                                case 4294967295:
                                    fmt.Println(fi8, v5, v1)
                                }
                            case 32767:
                                switch "\n\t" {
                                case "a":
                                    var v13 bool = ((3.14159 < 0.1) && true)
                                    _ = v13
                                case "":
                                    var v14 uint = 222
                                    _ = v14
                                case "\n\t":
                                    v1 = v1
                                default:
                                    var v15 uint64 = 0
                                    _ = v15
                                }
                            }
                        }
                    case -1:
                        if true {
                            fmt.Println(v1, v5)
                        } else {
                            switch fi7 {
                            case 1:
                                v16 := GenS2{F0: -1}
                                _ = v16
                            case -863:
                                v6[fi7] = ("hello" + "hello")
                            }
                        }
                    case 128:
                        if (float32(-2.25) < float32(0.0)) {
                            v17 := GenS3{F0: 675}
                            v18 := GenS2{F0: -815}
                            _ = v17
                            _ = v18
                        } else {
                            fmt.Println(fi7, v5, fi7)
                        }
                    default:
                        var v19 int = (fi7 - (fi8 / ((127) | 1)))
                        _ = v19
                    }
                } else {
                    var v20 uint32 = 2
                    switch fi8 {
                    case 9223372036854775807:
                        switch "foo bar" {
                        case "hello":
                            switch (fi7 - 1) {
                            case 0:
                                var v21 int32 = 0
                                _ = v21
                            case -462:
                                switch 2147483647 {
                                case 1:
                                    var v22 int = (-(fi7))
                                    _ = v22
                                default:
                                    var v23 uint8 = 255
                                    _ = v23
                                }
                            default:
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v24 int8 = (v1 << ((42) & 63))
                                    if true {
                                        v25 := GenS1{F0: fi8}
                                        _ = v25
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v26 uint64 = 1
                                    _ = v24
                                    _ = v26
                                default:
                                    if ("a" <= "foo bar") {
                                        v1 += 77
                                        v5 -= v5
                                        if (false || (!true)) {
                                            if false {
                                                var v27 float64 = 1.5
                                                fmt.Println(v27, v27)
                                                var v28 string = ("a" + "hello")
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v5 -= float32(0.0)
                                                var v29 string = ("foo bar" + "foo bar")
                                                var v30 uint16 = 353
                                                _ = v29
                                                _ = v30
                                            }
                                            var v31 int = fi7
                                            _ = v31
                                        }
                                        v5 += 0.0
                                    } else {
                                        if true {
                                            if (!false) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        } else {
                                            v4 = append(v4, ("" + "\n\t"))
                                            var v32 int64 = 1
                                            _ = v32
                                        }
                                        switch -9223372036854775808 {
                                        case 42:
                                            var v33 uint8 = 128
                                            _ = v33
                                        case 127:
                                            if false {
                                                v1 += 100
                                                fmt.Println(v20, fi8)
                                            }
                                        case -644:
                                            var v34 bool = true
                                            _ = v34
                                        default:
                                            var v35 int8 = -128
                                            _ = v35
                                        }
                                    }
                                }
                            }
                            var v36 uint8 = 42
                            _ = v36
                        }
                    case 50:
                        v5 = ((0.0 * (-2.25 * (v5 * v5))) * 3.14159)
                    default:
                        v20 -= 4294967295
                    }
                    _ = v20
                }
            }
            if (1 < 65535) {
                if false {
                    v5 -= v5
                    v1 += (v1 % ((23) | 1))
                }
                if (!true) {
                    if ((216 == 2) || false) {
                        if false {
                            var v37 int = (fi8 - fi8)
                            v37 += 919
                            _ = v37
                        } else {
                            v1 += v1
                            delete(v2, "hello")
                        }
                        if ((2 < 1) || false) {
                            if (!(!false)) {
                                fmt.Println(fi7, v1)
                                switch 4294967295 {
                                case 32768:
                                    v38 := GenS2{F0: -153}
                                    _ = v38
                                case 2147483647:
                                    v39 := GenS4{F0: 1}
                                    _ = v39
                                }
                            }
                        }
                        var v40 uint = 2
                        _ = v40
                    } else {
                        delete(v3, 137)
                        fmt.Println(v5)
                    }
                    var v41 float32 = v5
                    var v42 uint32 = 4294967295
                    _ = v41
                    _ = v42
                }
            } else {
                if false {
                    var v43 uint16 = 613
                    _ = v43
                }
                fmt.Println(v1)
                _ = len(v4)
            }
            fmt.Println(v5, v1, v5)
            _ = fi8
        }
        _ = fi7
    }
    var v44 uint8 = 128
    v5 = v5
    for fi45 := 0; fi45 < 5; fi45++ {
        v44 -= (v44 << ((485) & 63))
        _ = fi45
    }
    var v46 uint16 = 42
    var v47 int8 = (v1 >> ((uint(v44)) & 63))
    v44 = (v44 >> ((18446744073709551615) & 63))
    var v48 uint = 2
    var v49 int64 = 9223372036854775807
    for fi50 := 0; fi50 < 20; fi50++ {
        var v51 uint64 = 339
        for fi52 := 0; fi52 < 8; fi52++ {
            v53 := GenS5{F0: v49, F1: false}
            _ = fi52
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    delete(v2, "")
    v44 += 128
    for fi54 := 0; fi54 < 19; fi54++ {
        v49 += (v49 / (((v49 | 0)) | 1))
        _ = fi54
    }
    if true {
        v55 := GenS1{F0: 128}
        for fi56 := 0; fi56 < 1; fi56++ {
            v1 += v47
            for fi57 := 0; fi57 < 17; fi57++ {
                var v58 bool = true
                var v59 uint = v48
                _ = fi57
                _ = v58
                _ = v59
            }
            v48 += v48
            _ = fi56
        }
        _ = v55
    }
    v60 := GenS4{F0: (^(v44))}
    if false {
        v4 = append(v4, "hello")
        if (-32768 < 645) {
            if true {
                v1 = -76
                for fi61 := 0; fi61 < 5; fi61++ {
                    var v62 int32 = 0
                    if true {
                        fmt.Println(v46, v1, v62)
                    } else {
                        if (!(!true)) {
                            fmt.Println(v44, fi61, v48)
                        } else {
                            if true {
                                var v63 uint8 = (^(v44))
                                v48 = (v48 >> ((v48) & 63))
                                _ = v63
                            } else {
                                v5 += v5
                                var v64 uint16 = v46
                                _ = v64
                            }
                        }
                        if ((fi61 / ((fi61) | 1)) < -9223372036854775808) {
                            v65 := GenS6{F0: (^(v46)), F1: 521, F2: (v44 % ((v44) | 1))}
                            switch fi61 {
                            case -376:
                                v66 := GenS7{F0: (v1 % ((v47) | 1)), F1: (!false), F2: false}
                                _ = v66
                            case 38:
                                for fi67 := 0; fi67 < 7; fi67++ {
                                    if (((4294967295 <= 1) && true) && (false || true)) {
                                        if false {
                                            switch "hello" {
                                            case "\n\t":
                                                _ = len(v4)
                                            case "a":
                                                var v68 int32 = v62
                                                _ = v68
                                            default:
                                                var v69 int8 = (-(v47))
                                                _ = v69
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        v70 := GenS8{F0: -9223372036854775808, F1: v62, F2: (v48 % (((v48 % ((0) | 1))) | 1))}
                                        switch fi61 {
                                        case 2147483647:
                                            v47 += int8(v62)
                                        default:
                                            var v71 float32 = v5
                                            _ = v71
                                        }
                                        _ = v70
                                    }
                                    v44 -= (v44 << ((v48) & 63))
                                    if false {
                                        if (18446744073709551615 < 2) {
                                            var v72 int = int(v62)
                                            _ = v72
                                        }
                                        var v73 int = 9223372036854775807
                                        _ = v73
                                    }
                                    _ = fi67
                                }
                            case 0:
                                v6[(fi61 % (((fi61 ^ -65)) | 1))] = "a"
                            default:
                                var v74 uint8 = v44
                                _ = v74
                            }
                            _ = v65
                        }
                        if (265 <= 2) {
                            var v75 uint16 = v46
                            var v76 int32 = int32(v46)
                            _ = v75
                            _ = v76
                        } else {
                            for fi77 := 0; fi77 < 3; fi77++ {
                                var v78 uint64 = 42
                                switch fi77 {
                                case -418:
                                    delete(v2, (("foo bar" + "foo bar") + "foo bar"))
                                case -2147483648:
                                    fmt.Println(v5, fi61)
                                default:
                                    fmt.Println(v5, v1)
                                }
                                if true {
                                    v1 += v1
                                    if false {
                                        if (int8(v62) < (v1 / (((v1 / ((int8(v46)) | 1))) | 1))) {
                                            v62 += (v62 / ((v62) | 1))
                                            var v79 int64 = v49
                                            v49 = (v79 >> ((4294967295) & 63))
                                            _ = v79
                                        } else {
                                            var v80 float32 = v5
                                            _ = v80
                                        }
                                        if (!(false || false)) {
                                            var v81 uint16 = 1
                                            v5 += v5
                                            v49 = v49
                                            _ = v81
                                        } else {
                                            var v82 uint8 = v44
                                            v47 = (v47 + 127)
                                            v83 := GenS5{F0: (v49 + -695), F1: true}
                                            _ = v82
                                            _ = v83
                                        }
                                        fmt.Println(fi77, fi61)
                                    }
                                    v5 = v5
                                }
                                _ = fi77
                                _ = v78
                            }
                            v1 += v47
                        }
                    }
                    _ = fi61
                    _ = v62
                }
                if false {
                    if false {
                        if false {
                            v46 += (v46 & (v46 % ((v46) | 1)))
                            v44 = (v44 >> ((v48) & 63))
                            v84 := GenS9{F0: 1, F1: 65535, F2: 246}
                            _ = v84
                        } else {
                            _ = len(v6)
                            v48 -= v48
                            v1 = (-(v47))
                        }
                    }
                } else {
                    switch 531 {
                    case 127:
                        v47 -= (v47 & 127)
                    default:
                        _ = len(v2)
                    }
                }
            } else {
                for fi85 := 0; fi85 < 18; fi85++ {
                    for fi86 := 0; fi86 < 15; fi86++ {
                        if ((v49 != 9223372036854775807) || (1.5 >= 0.1)) {
                            if (786 > -647) {
                                v49 = (v49 % ((v49) | 1))
                                if (2147483647 <= 211) {
                                    switch ("a" + "hello") {
                                    case "foo bar":
                                        if (800 <= v48) {
                                            fmt.Println(v47, v47)
                                        }
                                    case "":
                                        v48 = 606
                                    case "a":
                                        v44 = 2
                                    default:
                                        var v87 uint = v48
                                        _ = v87
                                    }
                                }
                            } else {
                                v48 = v48
                                if true {
                                    if ((v46 ^ (v46 & (v46 << (((v48 / ((0) | 1))) & 63)))) != (v46 << (((v48 % ((0) | 1))) & 63))) {
                                        v4 = append(v4, ("foo bar" + "a"))
                                        fmt.Println(v1, v1)
                                        fmt.Println(v1, v49)
                                    }
                                } else {
                                    v47 += (v1 * (v47 * (v1 - (^(v47)))))
                                    var v88 int = fi86
                                    _ = v88
                                }
                            }
                        } else {
                            v2["\n\t"] = uint32(v48)
                            fmt.Println(fi86, fi86, v1)
                        }
                        v1 = (v1 / (((v1 << ((42) & 63))) | 1))
                        v48 = 393
                        _ = fi86
                    }
                    var v89 int64 = (v49 % ((v49) | 1))
                    _ = fi85
                    _ = v89
                }
            }
            switch -281 {
            case 1:
                v5 += 0.0
            }
            v3[2147483648] = false
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi90 := 0; fi90 < 10; fi90++ {
        switch ("hello" + "hello") {
        case "a":
            v91 := GenS2{F0: v49}
            _ = v91
        case "hello":
            var v92 float64 = 1e10
            _ = v92
        case "foo bar":
            var v93 uint16 = v46
            _ = v93
        }
        _ = fi90
    }
    if (!false) {
        fmt.Println(v49, v1)
    } else {
        fmt.Println(v49, v49)
        for fi94 := 0; fi94 < 15; fi94++ {
            for fi95 := 0; fi95 < 19; fi95++ {
                if ((fi94 % ((fi94) | 1)) < fi94) {
                    v5 += ((v5 * 0.0) - -2.25)
                    fmt.Println(v47)
                    v1 = -128
                }
                _ = fi95
            }
            _ = fi94
        }
        if false {
            var v96 uint8 = 42
            fmt.Println(v49, v44, v47)
            _ = v96
        } else {
            if false {
                switch "hello" {
                case "foo bar":
                    for fi97 := 0; fi97 < 6; fi97++ {
                        for fi98 := 0; fi98 < 6; fi98++ {
                            var v99 bool = false
                            var v100 bool = (!v99)
                            _ = fi98
                            _ = v100
                            _ = v99
                        }
                        for fi101 := 0; fi101 < 19; fi101++ {
                            v102 := GenS2{F0: int64(v1)}
                            v103 := GenS6{F0: 2, F1: 822, F2: 42}
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            _ = fi101
                            _ = v102
                            _ = v103
                        }
                        _ = fi97
                    }
                case "hello":
                    var v104 uint32 = 843
                    _ = v104
                default:
                    v48 -= uint(v46)
                }
                v105 := GenS10{F0: 42, F1: 32767, F2: 1}
                _ = v105
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v60
}
