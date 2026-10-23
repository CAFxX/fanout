package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 int64 = -432
    var v4 float64 = (3.14159 * (3.14159 + (1.5 * (0.1 - -2.25))))
    v5 := GenS1{F0: 0, F1: 42}
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            v3 = 54
            for fi7 := 0; fi7 < 19; fi7++ {
                v3 += (v3 ^ -9223372036854775808)
                _ = fi7
            }
        }
        if (!(1 == 1)) {
            if false {
                v8 := []int{128, (fi6 | 241), 128}
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            if (!(0 > 2)) {
                switch (fi9 << ((255) & 63)) {
                case 290:
                    var v10 uint8 = 42
                    _ = v10
                case 32767:
                    fmt.Println(fi6, fi9, fi6)
                case 4294967295:
                    var v11 uint64 = 0
                    _ = v11
                default:
                    switch fi6 {
                    case 2147483647:
                        v3 -= 1
                    case -1:
                        v3 = v3
                    case 1:
                        if (!(310 > 4294967295)) {
                            if (false || (!(131 > 1))) {
                                var v12 float64 = float64(1e10)
                                switch fi6 {
                                case 4294967295:
                                    v13 := []int{(-(fi9)), fi9, fi9, fi9}
                                    _ = v13
                                case -991:
                                    v3 -= (^(v3))
                                }
                                _ = v12
                            }
                            if (!(fi6 >= (fi6 / (((fi6 & fi9)) | 1)))) {
                                v3 += (v3 % (((v3 << ((0) & 63))) | 1))
                            }
                        } else {
                            v3 = (v3 - v3)
                            v4 = (v4 * 1.5)
                        }
                    }
                }
                var v14 float64 = (v4 - v4)
                var v15 int16 = -90
                _ = v14
                _ = v15
            } else {
                var v16 uint64 = 18446744073709551615
                switch ((("foo bar" + "hello") + ("\n\t" + "\n\t")) + ("foo bar" + "hello")) {
                case "hello":
                    v16 += 275
                case "":
                    v17 := []int{32768, (^(fi9))}
                    _ = v17
                default:
                    var v18 uint8 = 128
                    _ = v18
                }
                _ = v16
            }
            _ = fi9
        }
        _ = fi6
    }
    var v19 uint16 = 0
    var v20 uint64 = 0
    var v21 uint = 4294967295
    switch ("\n\t" + "") {
    case "foo bar":
        var v22 uint8 = 2
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if (42 >= -128) {
                var v25 int64 = v3
                fmt.Println(v21, v21)
                v25 = -1
                _ = v25
            } else {
                v20 = (^(v20))
            }
            _ = fi24
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int64 = (v3 * (v3 - v3))
            if true {
                var v28 uint32 = 659
                if (v20 > (^(v20))) {
                    fmt.Println(v28, fi26, v19)
                    if (2 == 103) {
                        v4 = v4
                        if (true && (!(100 >= -128))) {
                            fmt.Println(fi23, v21)
                            v29 := []int{fi23}
                            var v30 uint32 = v28
                            switch fi26 {
                            case 32767:
                                var v31 string = "hello"
                                _ = v31
                            case 403:
                                fmt.Println(v27, fi26, v27)
                            }
                            _ = v29
                            _ = v30
                        }
                        v32 := []int{fi26, fi23, (fi23 & fi26)}
                        var v33 uint64 = 42
                        _ = v32
                        _ = v33
                    }
                }
                _ = v28
            }
            _ = fi26
            _ = v27
        }
        _ = fi23
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 bool = (true && ((!true) || true))
        _ = fi34
        _ = v35
    }
    switch -9223372036854775808 {
    case -519:
        for fi36 := 0; fi36 < 13; fi36++ {
            v37 := GenS1{F0: 255, F1: (fi36 / (((fi36 >> ((65535) & 63))) | 1))}
            if false {
                var v38 float64 = 0.0
                switch "\n\t" {
                case "foo bar":
                    var v39 int32 = 42
                    _ = v39
                case "":
                    if (!(false || true)) {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            v41 := []int{(fi36 << ((v21) & 63)), (fi36 % ((32768) | 1)), fi36, fi40}
                            _ = fi40
                            _ = v41
                        }
                        switch (("foo bar" + "\n\t") + "hello") {
                        case "hello":
                            var v42 float64 = -2.25
                            _ = v42
                        case "\n\t":
                            v4 -= v38
                        case "a":
                            v43 := GenS2{F0: v20, F1: (255 <= 1)}
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 5; fi44++ {
                                v4 += v38
                                var v45 float64 = v4
                                if true {
                                    v20 += v20
                                    v46 := []int{}
                                    var v47 uint8 = 45
                                    v48 := GenS1{F0: (^(v47)), F1: (fi36 ^ fi44)}
                                    v49 := GenS1{F0: (v47 % ((v47) | 1)), F1: fi44}
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                } else {
                                    if false {
                                        var v50 int64 = -9223372036854775808
                                        _ = v50
                                    } else {
                                        if false {
                                            switch (fi36 >> ((v21) & 63)) {
                                            case 1:
                                                v3 -= -9223372036854775808
                                            case 127:
                                                fmt.Println(v3)
                                            case -9223372036854775808:
                                                var v51 uint = 814
                                                _ = v51
                                            default:
                                                v20 += v20
                                            }
                                            if (!(!(42 < 42))) {
                                                v45 -= (v45 - ((v45 * 1.5) / v38))
                                            }
                                        }
                                        var v52 int16 = 1
                                        _ = v52
                                    }
                                    if (!false) {
                                        switch (("hello" + "\n\t") + "a") {
                                        case "\n\t":
                                            var v53 float32 = (0.0 / 1.5)
                                            _ = v53
                                        case "":
                                            var v54 float32 = 1e10
                                            _ = v54
                                        }
                                        var v55 uint32 = 396
                                        var v56 int64 = (v3 + v3)
                                        _ = v55
                                        _ = v56
                                    } else {
                                        if false {
                                            v57 := GenS2{F0: (^(v20)), F1: (!((uint8(v19) == uint8(v19)) || ((v19 == v19) || false)))}
                                            v58 := GenS2{F0: 0, F1: true}
                                            v2["hello"] = true
                                            v59 := []int{-1, (fi44 - fi36)}
                                            var v60 int = 9223372036854775807
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                }
                                _ = fi44
                                _ = v45
                            }
                        }
                    } else {
                        fmt.Println(v38)
                        var v61 uint64 = 1
                        _ = v61
                    }
                case "\n\t":
                    var v62 int64 = v3
                    _ = v62
                default:
                    var v63 uint = 871
                    _ = v63
                }
                v64 := GenS3{F0: true}
                _ = v38
                _ = v64
            } else {
                if false {
                    v65 := GenS2{F0: v20, F1: (641 >= int32(v19))}
                    v3 += (v3 ^ -879)
                    _ = v65
                } else {
                    v66 := GenS4{F0: v21, F1: -496, F2: uint32(v19)}
                    _ = v66
                }
            }
            for fi67 := 0; fi67 < 3; fi67++ {
                v68 := []int{fi67, fi67, -2147483648}
                fmt.Println(v21)
                _ = fi67
                _ = v68
            }
            _ = fi36
            _ = v37
        }
    case -965:
        switch "hello" {
        case "foo bar":
            var v69 string = ""
            _ = v69
        }
    case -1:
        var v70 uint8 = uint8(v20)
        _ = v70
    }
    if true {
        if false {
            switch "\n\t" {
            case "a":
                v3 = -9223372036854775808
            case "foo bar":
                switch int(v19) {
                case 32768:
                    var v71 string = "foo bar"
                    _ = v71
                case -1:
                    v19 -= (v19 - 0)
                }
            case "\n\t":
                for fi72 := 0; fi72 < 13; fi72++ {
                    for fi73 := 0; fi73 < 9; fi73++ {
                        if (v20 <= v20) {
                            v20 -= (^(v20))
                        } else {
                            v74 := GenS4{F0: 1, F1: 42, F2: 466}
                            _ = v74
                        }
                        var v75 int = -663
                        var v76 int = fi72
                        _ = fi73
                        _ = v75
                        _ = v76
                    }
                    v77 := GenS3{F0: (792 < 1)}
                    _ = fi72
                    _ = v77
                }
            }
            var v78 uint = 18446744073709551615
            _ = v78
        } else {
            var v79 uint32 = 2147483647
            for fi80 := 0; fi80 < 8; fi80++ {
                fmt.Println(v19, v79, v20)
                for fi81 := 0; fi81 < 15; fi81++ {
                    v21 = (v21 - v21)
                    _ = fi81
                }
                v19 = (v19 % ((v19) | 1))
                _ = fi80
            }
            _ = v79
        }
        for fi82 := 0; fi82 < 14; fi82++ {
            var v83 int = (fi82 ^ fi82)
            _ = fi82
            _ = v83
        }
        v19 += 544
    }
    v20 -= (v20 * (v20 ^ v20))
    fmt.Println(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
