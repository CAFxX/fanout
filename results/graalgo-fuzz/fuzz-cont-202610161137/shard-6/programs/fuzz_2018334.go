package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 bool; F1 int64; F2 int64 }

func main() {
    var v1 int8 = 1
    var v2 int16 = 1000
    var v3 uint8 = 255
    var v4 int32 = 0
    var v5 int32 = (v4 + (^(v4)))
    v5 += (v4 >> ((580) & 63))
    v1 = (v1 % ((127) | 1))
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint16 = 830
        var v8 float32 = float32(-2.25)
        var v9 string = (("hello" + "a") + ("a" + ""))
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{-776, 42}
    var v11 int = 128
    var v12 int = (v11 / (((v11 | v11)) | 1))
    fmt.Println(v1, v5, v5)
    var v13 int32 = v4
    v12 += (v11 / ((-2147483648) | 1))
    if len(v10) > 0 {
        v10[0] = (v12 >> ((255) & 63))
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            if (80 != 1) {
                v12 = fi14
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            fmt.Println(v12, fi15)
            if true {
                fmt.Println(v2, v13)
                var v16 int64 = 513
                fmt.Println(v2, v4)
                if true {
                    v16 += v16
                } else {
                    switch "a" {
                    case "\n\t":
                        var v17 int32 = v13
                        _ = v17
                    case "hello":
                        var v18 int8 = (v1 >> ((42) & 63))
                        _ = v18
                    default:
                        if true {
                            if (!true) {
                                fmt.Println(fi15, fi15)
                                v5 = v13
                            }
                            v10 = append(v10, fi14)
                        }
                    }
                    var v19 uint64 = 12
                    v3 += v3
                    _ = v19
                }
                _ = v16
            } else {
                v20 := GenS1{F0: (^(v3))}
                if (1000 == (v2 + v2)) {
                    v3 = v3
                    if (!false) {
                        if true {
                            v2 = 32767
                        } else {
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        fmt.Println(v13, v12, v2)
                        if true {
                            var v22 bool = false
                            v13 = v4
                            if (262 < 42) {
                                if ((1 > 0) || (false || (v22 && (18446744073709551615 >= 208)))) {
                                    v3 = 1
                                } else {
                                    v10 = append(v10, v11)
                                }
                                var v23 int32 = 1
                                _ = v23
                            } else {
                                fmt.Println(v5)
                                v2 = (v2 - 1)
                            }
                            _ = v22
                        }
                        switch (v12 ^ 1) {
                        case 32767:
                            fmt.Println(v2)
                        case 1:
                            fmt.Println(v4, fi15, v1)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        case 128:
                            v10 = append(v10, (^(fi15)))
                        }
                        fmt.Println(fi15, fi15)
                        var v24 bool = (1e10 < 1.5)
                        if (1 <= 2) {
                            switch "" {
                            case "a":
                                fmt.Println(fi15, v13)
                                switch ("" + (("\n\t" + "") + "foo bar")) {
                                case "hello":
                                    var v25 uint16 = 65535
                                    _ = v25
                                default:
                                    v26 := GenS1{F0: (^(v3))}
                                    _ = v26
                                }
                                v27 := GenS2{F0: 1, F1: -1}
                                _ = v27
                            case "hello":
                                var v28 int32 = 0
                                _ = v28
                            }
                        }
                        v11 -= v11
                        switch ("a" + "foo bar") {
                        case "a":
                            switch (^(fi14)) {
                            case 32767:
                                if (18446744073709551615 <= 2) {
                                    v5 -= v5
                                    if false {
                                        switch ("foo bar" + "a") {
                                        case "foo bar":
                                            v24 = true
                                        default:
                                            v29 := GenS1{F0: v3}
                                            _ = v29
                                        }
                                        v10 = append(v10, 127)
                                    } else {
                                        var v30 float32 = 1.5
                                        if false {
                                            v31 := GenS3{F0: 519, F1: 4294967295}
                                            _ = v31
                                        } else {
                                            switch v11 {
                                            case -509:
                                                if true {
                                                    var v32 float32 = v30
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v33 float32 = 0.0
                                                    _ = v32
                                                    _ = v33
                                                }
                                            case -2147483648:
                                                var v34 uint = 692
                                                _ = v34
                                            case 127:
                                                var v35 uint = 1
                                                _ = v35
                                            }
                                            switch 127 {
                                            case 226:
                                                v36 := GenS4{F0: 609, F1: 255}
                                                _ = v36
                                            }
                                        }
                                        _ = v30
                                    }
                                    v13 = 2147483647
                                }
                            default:
                                if (!(1e10 >= 1e10)) {
                                    fmt.Println(v2, v12, v2)
                                    if v24 {
                                        v5 = v5
                                    } else {
                                        if false {
                                            var v37 int8 = (v1 << ((42) & 63))
                                            if true {
                                                switch (v12 | -346) {
                                                case -9223372036854775808:
                                                    var v38 int16 = (v2 | -417)
                                                    _ = v38
                                                case 9223372036854775807:
                                                    switch ((("foo bar" + ("a" + ("" + ""))) + (("foo bar" + ("\n\t" + "a")) + "")) + "") {
                                                    case "hello":
                                                        v5 += v13
                                                    case "foo bar":
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v12, v3, v5)
                                                }
                                                v13 = v4
                                            } else {
                                                var v39 uint32 = 695
                                                if (v2 != 42) {
                                                    if (!(0 < v39)) {
                                                        var v40 int = 128
                                                        v11 -= v12
                                                        v10 = append(v10, 127)
                                                        _ = v40
                                                    }
                                                } else {
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v41 float64 = 1.5
                                                    _ = len(v10)
                                                    _ = v41
                                                }
                                                if v24 {
                                                    switch ("hello" + "hello") {
                                                    case "a":
                                                        v42 := GenS5{F0: (v3 ^ (v3 * v3))}
                                                        _ = v42
                                                    case "":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(fi14, fi15)
                                                        }
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "":
                                                            v37 = (v1 + (v37 & (v1 % ((v37) | 1))))
                                                        case "foo bar":
                                                            if (18446744073709551615 >= 2) {
                                                                var v43 string = "foo bar"
                                                                var v44 int32 = v13
                                                                v39 = (v39 >> ((18446744073709551615) & 63))
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        case "a":
                                                            v10 = append(v10, (v11 + 0))
                                                        default:
                                                            var v45 uint = 255
                                                            _ = v45
                                                        }
                                                    }
                                                    v37 = int8(v2)
                                                } else {
                                                    fmt.Println(v37, v12, fi15)
                                                    var v46 int64 = 0
                                                    v47 := GenS6{F0: 0, F1: (("" + "hello") > ("a" + "hello")), F2: 601}
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 uint64 = 1
                                                var v49 int = -332
                                                _ = v39
                                                _ = v48
                                                _ = v49
                                            }
                                            v11 = 52
                                            _ = v37
                                        } else {
                                            v24 = (v3 <= (v3 | 255))
                                        }
                                        fmt.Println(fi15)
                                    }
                                    switch v11 {
                                    case -2147483648:
                                        if (v13 < (v13 >> ((42) & 63))) {
                                            var v50 uint16 = 0
                                            switch ("foo bar" + "a") {
                                            case "":
                                                switch ("foo bar" + "") {
                                                case "foo bar":
                                                    if v24 {
                                                        var v51 float32 = 1e10
                                                        _ = v51
                                                    }
                                                default:
                                                    if v24 {
                                                        switch -752 {
                                                        case -951:
                                                            switch "" {
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if (!v24) {
                                                                        v10 = append(v10, 32767)
                                                                        v10 = append(v10, (fi14 << ((968) & 63)))
                                                                        if (v24 || v24) {
                                                                            var v52 uint32 = 65535
                                                                            fmt.Println(v11)
                                                                            if v24 {
                                                                                v13 += (v13 | v4)
                                                                                var v53 float32 = (0.0 + 1e10)
                                                                                _ = v53
                                                                            }
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    var v54 int32 = v5
                                                                    _ = v54
                                                                }
                                                            }
                                                        }
                                                        switch fi14 {
                                                        case 4294967295:
                                                            var v55 float32 = 1e10
                                                            fmt.Println(v4, v11, v13)
                                                            if (v50 <= (^(v50))) {
                                                                var v56 int8 = 100
                                                                _ = v56
                                                            }
                                                            v5 -= (v4 & v13)
                                                            _ = v55
                                                        case 1:
                                                            fmt.Println(v3, v13, v2)
                                                        }
                                                        v1 += (v1 ^ v1)
                                                    } else {
                                                        fmt.Println(v24, v3)
                                                        if ((v1 ^ (v1 >> ((4294967295) & 63))) > (v1 * 100)) {
                                                            if v24 {
                                                                if v24 {
                                                                    var v57 int32 = v5
                                                                    var v58 uint16 = 65535
                                                                    v13 -= 0
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                fmt.Println(v1, v12)
                                                            } else {
                                                                fmt.Println(v13)
                                                                var v59 uint64 = 2
                                                                var v60 float64 = ((-2.25 - (float64(0.0) / 1.0)) / 0.1)
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            var v61 int16 = v2
                                                            v62 := GenS3{F0: 4294967295, F1: 4294967295}
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v63 uint16 = v50
                                                                _ = v63
                                                            default:
                                                                var v64 int64 = 1
                                                                _ = v64
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        v2 += 32767
                                                        switch "\n\t" {
                                                        case "a":
                                                            v50 -= v50
                                                        case "":
                                                            var v65 float32 = 3.14159
                                                            _ = v65
                                                        case "\n\t":
                                                            v4 -= v5
                                                        }
                                                        switch (v11 ^ -2147483648) {
                                                        case -2147483648:
                                                            var v66 float32 = 3.14159
                                                            _ = v66
                                                        case 624:
                                                            v67 := GenS5{F0: (v3 | v3)}
                                                            _ = v67
                                                        }
                                                        fmt.Println(v3)
                                                    }
                                                }
                                            }
                                            v24 = (((v24 && (((v4 >> ((124) & 63)) < int32(v2)) || true)) || (v2 != (-(v2)))) || false)
                                            _ = v50
                                        } else {
                                            v4 += v13
                                            v68 := GenS3{F0: 1, F1: 65535}
                                            _ = v68
                                        }
                                    default:
                                        if true {
                                            var v69 uint8 = v3
                                            var v70 uint16 = 2
                                            var v71 uint32 = 42
                                            v12 = fi15
                                            if len(v10) > 0 {
                                                v10[0] = v12
                                            }
                                            v72 := GenS7{F0: 255, F1: v24, F2: v71}
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        } else {
                                            var v73 int16 = v2
                                            _ = len(v10)
                                            _ = v73
                                        }
                                    }
                                }
                            }
                            v74 := GenS8{F0: (v11 >> ((2) & 63)), F1: int8(fi15)}
                            _ = v74
                        default:
                            v75 := GenS9{F0: (v3 == (v3 | v3)), F1: 98, F2: -9223372036854775808}
                            _ = v75
                        }
                        _ = v24
                    }
                }
                v2 = (v2 + -32768)
                var v76 uint = uint(v3)
                _ = v20
                _ = v76
            }
            _ = fi15
        }
        v2 = (v2 * v2)
        _ = fi14
    }
    v10 = append(v10, 2147483647)
    var v77 uint32 = 0
    fmt.Println(v11, v11, v2)
    for fi78 := 0; fi78 < 18; fi78++ {
        if false {
            if (((740 < -9223372036854775808) && (127 < v11)) && (false && (2147483648 != (v77 ^ (v77 >> ((0) & 63)))))) {
                var v79 float32 = float32(0.1)
                _ = v79
            }
            for fi80 := 0; fi80 < 12; fi80++ {
                fmt.Println(v3, v3)
                _ = fi80
            }
            var v81 int64 = 0
            _ = v81
        }
        v12 -= 769
        _ = fi78
    }
    v77 += v77
    if ((v3 + (^(v3))) >= 0) {
        var v82 int8 = (v1 << ((4294967295) & 63))
        v83 := GenS5{F0: v3}
        v10 = append(v10, -1)
        _ = v82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
