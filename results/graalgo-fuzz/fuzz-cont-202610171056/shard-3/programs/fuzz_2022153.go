package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int8; F2 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 0.1
    v3 := []uint32{338, 2147483647, 2147483647}
    v4 := make(map[string]int)
    switch 0 {
    case 128:
        var v5 float32 = 0.0
        _ = v5
    case -38:
        v6 := GenS1{F0: (!true), F1: false}
        _ = v6
    case 283:
        switch 10 {
        case -65:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            fmt.Println(v2)
        }
    default:
        if true {
            switch 128 {
            case 127:
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := GenS1{F0: (938 == 0), F1: true}
                    v9 := GenS1{F0: false, F1: false}
                    var v10 int64 = 0
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            if (true && (!false)) {
                fmt.Println(v2)
                v11 := GenS1{F0: (!(false && false)), F1: (469 >= 32767)}
                _ = v11
            } else {
                var v12 uint = 0
                _ = v12
            }
            var v13 int64 = -1
            _ = v13
        } else {
            if (127 != 255) {
                switch (("" + "foo bar") + "foo bar") {
                case "foo bar":
                    var v14 uint32 = 1
                    _ = v14
                }
                if (42 != 255) {
                    var v15 int8 = 100
                    if true {
                        v16 := GenS2{F0: 42, F1: 18446744073709551615}
                        switch 127 {
                        case 42:
                            if false {
                                for fi17 := 0; fi17 < 6; fi17++ {
                                    if true {
                                        if true {
                                            var v18 int8 = v15
                                            v15 -= (^(v18))
                                            _ = v18
                                        }
                                    }
                                    v19 := GenS1{F0: ((0 <= 1) && (0 <= 42)), F1: true}
                                    _ = fi17
                                    _ = v19
                                }
                                v20 := v1[("foo bar" + "hello")]
                                var v21 int8 = (v15 | v15)
                                _ = v20
                                _ = v21
                            } else {
                                switch -765 {
                                case 209:
                                    v3 = append(v3, 2147483648)
                                case -499:
                                    if (17 >= 2147483648) {
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v22 uint = 204
                                                _ = v22
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            default:
                                                var v23 float32 = (1e10 - -2.25)
                                                _ = v23
                                            }
                                        case "\n\t":
                                            if (-464 > 687) {
                                                if true {
                                                    for fi24 := 0; fi24 < 19; fi24++ {
                                                        v15 -= v15
                                                        _ = fi24
                                                    }
                                                    switch 32767 {
                                                    case -9223372036854775808:
                                                        v15 = -1
                                                    case 0:
                                                        switch -9223372036854775808 {
                                                        case 4294967295:
                                                            var v25 int8 = v15
                                                            _ = v25
                                                        case 2147483647:
                                                            if (!false) {
                                                                v15 += (v15 >> ((60) & 63))
                                                                fmt.Println(v15, v2)
                                                                for fi26 := 0; fi26 < 11; fi26++ {
                                                                    var v27 int = 128
                                                                    fmt.Println(v2, v27, v2)
                                                                    _ = fi26
                                                                    _ = v27
                                                                }
                                                            } else {
                                                                if true {
                                                                    if false {
                                                                        v28 := GenS3{F0: -600}
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v28
                                                                    } else {
                                                                        for fi29 := 0; fi29 < 16; fi29++ {
                                                                            var v30 uint16 = 1
                                                                            v31 := GenS4{F0: -2147483648, F1: 216}
                                                                            if (true || ("hello" > "")) {
                                                                                if (false || (false && true)) {
                                                                                    v32 := GenS5{F0: -1}
                                                                                    v33 := GenS6{F0: 407, F1: -128, F2: fi29}
                                                                                    var v34 uint8 = 42
                                                                                    _ = v32
                                                                                    _ = v33
                                                                                    _ = v34
                                                                                } else {
                                                                                    fmt.Println(v15)
                                                                                    var v35 bool = (4294967295 != 909)
                                                                                    _ = v35
                                                                                }
                                                                                if true {
                                                                                    v36 := v4["hello"]
                                                                                    if ("hello" < "") {
                                                                                        v36 += (v36 << ((674) & 63))
                                                                                        var v37 uint64 = 2
                                                                                        _ = v37
                                                                                    } else {
                                                                                        _ = len(v3)
                                                                                        var v38 int8 = v15
                                                                                        _ = v38
                                                                                    }
                                                                                    var v39 uint64 = 1
                                                                                    _ = v36
                                                                                    _ = v39
                                                                                }
                                                                                v15 = (v15 * int8(v30))
                                                                            } else {
                                                                                if true {
                                                                                    var v40 bool = (float32(v2) <= 3.14159)
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        var v41 uint = 245
                                                                                        _ = v41
                                                                                    case "":
                                                                                        var v42 int32 = 0
                                                                                        _ = v42
                                                                                    default:
                                                                                        if (0 <= 2) {
                                                                                            for fi43 := 0; fi43 < 13; fi43++ {
                                                                                                if v40 {
                                                                                                    v30 = 61
                                                                                                    switch fi29 {
                                                                                                    case 4294967295:
                                                                                                        var v44 int8 = v15
                                                                                                        _ = v44
                                                                                                    case 864:
                                                                                                        var v45 uint64 = 701
                                                                                                        _ = v45
                                                                                                    case 32768:
                                                                                                        switch "\n\t" {
                                                                                                        case "foo bar":
                                                                                                            if (v40 && (-495 > -2147483648)) {
                                                                                                                if v40 {
                                                                                                                    var v46 uint64 = 0
                                                                                                                    if false {
                                                                                                                        v30 = (v30 >> ((2) & 63))
                                                                                                                    }
                                                                                                                    v1["foo bar"] = v40
                                                                                                                    _ = v46
                                                                                                                }
                                                                                                                switch (fi29 * fi43) {
                                                                                                                case -2147483648:
                                                                                                                    v47 := GenS7{F0: 1}
                                                                                                                    _ = v47
                                                                                                                case 610:
                                                                                                                    v48 := GenS4{F0: 2147483647, F1: 128}
                                                                                                                    _ = v48
                                                                                                                }
                                                                                                                fmt.Println(fi29, fi43)
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if true {
                                                                                                                if (0.1 > float32(float64(0.1))) {
                                                                                                                    v40 = (1 > 2)
                                                                                                                    v49 := GenS8{F0: ((((fi43 * fi29) != fi43) || v40) && (2147483647 < 282)), F1: 0}
                                                                                                                    v2 += float64((1.5 * float32((v2 * 1.5))))
                                                                                                                    _ = v49
                                                                                                                } else {
                                                                                                                    v40 = (!v40)
                                                                                                                    var v50 int64 = 980
                                                                                                                    _ = v50
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v30, v40, v40)
                                                                                                        }
                                                                                                    default:
                                                                                                        v51 := v4["a"]
                                                                                                        var v52 uint32 = 65535
                                                                                                        v52 = v52
                                                                                                        if v40 {
                                                                                                            switch fi29 {
                                                                                                            case 50:
                                                                                                                v53 := v1[""]
                                                                                                                _ = v53
                                                                                                            case -2147483648:
                                                                                                                var v54 int = fi29
                                                                                                                _ = v54
                                                                                                            case 473:
                                                                                                                var v55 float64 = ((-2.25 / (((v2 + 0.0) + v2) + v2)) + 1e10)
                                                                                                                _ = v55
                                                                                                            default:
                                                                                                                var v56 int16 = 458
                                                                                                                _ = v56
                                                                                                            }
                                                                                                            fmt.Println(fi43, v30)
                                                                                                            if (v40 && (!true)) {
                                                                                                                _ = len(v3)
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = (v52 >> ((642) & 63))
                                                                                                                }
                                                                                                                fmt.Println(v30, v40)
                                                                                                                var v57 int64 = -9223372036854775808
                                                                                                                _ = v57
                                                                                                            } else {
                                                                                                                var v58 uint64 = 42
                                                                                                                _ = v58
                                                                                                            }
                                                                                                            _ = len(v4)
                                                                                                        } else {
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 2147483647
                                                                                                            }
                                                                                                            var v59 uint32 = 528
                                                                                                            var v60 uint64 = 689
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                v2 = ((v2 + v2) - 3.14159)
                                                                                                            default:
                                                                                                                delete(v4, (("hello" + "") + ""))
                                                                                                            }
                                                                                                            v15 -= v15
                                                                                                            _ = v59
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        v15 -= (v15 << ((2) & 63))
                                                                                                        _ = v51
                                                                                                        _ = v52
                                                                                                    }
                                                                                                    v30 = (v30 * 65535)
                                                                                                }
                                                                                                fmt.Println(v30)
                                                                                                _ = fi43
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if ("hello" <= "") {
                                                                                        if (0 >= -9223372036854775808) {
                                                                                            if (0.0 != -2.25) {
                                                                                                v61 := GenS8{F0: v40, F1: 30}
                                                                                                _ = v61
                                                                                            } else {
                                                                                                for fi62 := 0; fi62 < 8; fi62++ {
                                                                                                    fmt.Println(v30, v40, fi29)
                                                                                                    _ = fi62
                                                                                                }
                                                                                                v30 = v30
                                                                                                _ = len(v3)
                                                                                            }
                                                                                        }
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    _ = v40
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 32
                                                                                }
                                                                            }
                                                                            _ = fi29
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                        v2 = v2
                                                                        v15 += v15
                                                                    }
                                                                    if (true && (-853 > 42)) {
                                                                        var v63 int8 = 96
                                                                        for fi64 := 0; fi64 < 6; fi64++ {
                                                                            var v65 int8 = v15
                                                                            fmt.Println(v15, v2)
                                                                            _ = fi64
                                                                            _ = v65
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            v15 = (v63 >> ((0) & 63))
                                                                        case "a":
                                                                            v66 := GenS9{F0: 0, F1: 878}
                                                                            _ = v66
                                                                        case "hello":
                                                                            fmt.Println(v2, v15)
                                                                        default:
                                                                            v67 := GenS6{F0: -9223372036854775808, F1: (v63 - 127), F2: 127}
                                                                            _ = v67
                                                                        }
                                                                        _ = v63
                                                                    } else {
                                                                        var v68 float64 = v2
                                                                        _ = v68
                                                                    }
                                                                } else {
                                                                    var v69 int = 682
                                                                    _ = v69
                                                                }
                                                            }
                                                        }
                                                    case 32768:
                                                        for fi70 := 0; fi70 < 14; fi70++ {
                                                            var v71 uint64 = 18446744073709551615
                                                            _ = fi70
                                                            _ = v71
                                                        }
                                                    default:
                                                        v15 -= v15
                                                    }
                                                    var v72 string = "\n\t"
                                                    _ = v72
                                                }
                                                v73 := GenS4{F0: 1, F1: 48}
                                                _ = v73
                                            }
                                        case "hello":
                                            if false {
                                                v74 := GenS10{F0: 862, F1: "hello"}
                                                for fi75 := 0; fi75 < 13; fi75++ {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v2 -= v2
                                                    case "\n\t":
                                                        if (((-(v15)) <= (v15 << ((255) & 63))) || true) {
                                                            for fi76 := 0; fi76 < 5; fi76++ {
                                                                var v77 int = (fi76 ^ fi76)
                                                                _ = fi76
                                                                _ = v77
                                                            }
                                                        } else {
                                                            v15 += v15
                                                            v2 += v2
                                                        }
                                                    case "foo bar":
                                                        var v78 float32 = 1e10
                                                        _ = v78
                                                    default:
                                                        var v79 int16 = 0
                                                        _ = v79
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = fi75
                                                }
                                                v3 = append(v3, 136)
                                                _ = v74
                                            } else {
                                                v3 = append(v3, 871)
                                                v3 = append(v3, 0)
                                                for fi80 := 0; fi80 < 10; fi80++ {
                                                    v15 += v15
                                                    v2 -= float64((1.5 * (0.1 - (float32(((0.1 - 3.14159) - v2)) + 0.1))))
                                                    var v81 int = (^(fi80))
                                                    _ = fi80
                                                    _ = v81
                                                }
                                            }
                                        default:
                                            if (1 == -2147483648) {
                                                var v82 int64 = 712
                                                var v83 int8 = (v15 - 106)
                                                v82 = v82
                                                _ = v82
                                                _ = v83
                                            }
                                        }
                                        delete(v1, "foo bar")
                                        var v84 string = ("a" + "\n\t")
                                        _ = v84
                                    }
                                }
                            }
                        case -9223372036854775808:
                            v15 += (v15 ^ (v15 ^ (v15 * v15)))
                        default:
                            fmt.Println(v15)
                        }
                        _ = v16
                    } else {
                        var v85 bool = false
                        switch "hello" {
                        case "hello":
                            var v86 uint16 = 255
                            _ = v86
                        case "foo bar":
                            fmt.Println(v2, v15)
                        }
                        _ = v85
                    }
                    v15 = -123
                    _ = v15
                } else {
                    v2 = float64(((0.1 - float32(v2)) - 1e10))
                }
                v87 := v1["\n\t"]
                _ = v87
            }
            v2 += v2
            v2 -= v2
        }
    }
    switch "\n\t" {
    case "a":
        if (!true) {
            var v88 float32 = (1e10 + (-2.25 / (1.5 + (3.14159 / 1.0))))
            var v89 uint64 = 42
            _ = len(v1)
            _ = v88
            _ = v89
        }
    case "":
        v2 = 1.5
    case "foo bar":
        _ = len(v3)
    }
    for fi90 := 0; fi90 < 16; fi90++ {
        v91 := GenS11{F0: "hello"}
        var v92 float32 = 1.5
        _ = fi90
        _ = v91
        _ = v92
    }
    v2 = float64(1e10)
    switch "a" {
    case "":
        var v93 uint32 = 65535
        _ = v93
    }
    var v94 string = ""
    switch (v94 + ((v94 + "") + v94)) {
    case "\n\t":
        delete(v4, ("foo bar" + v94))
    case "hello":
        var v95 float32 = 1e10
        _ = v95
    default:
        _ = len(v3)
    }
    v94 = ("\n\t" + "hello")
    var v96 float64 = (((float64(0.1) * -2.25) / -2.25) + (((v2 * (v2 * (0.1 - (3.14159 - (0.1 - 0.0))))) + v2) * float64(float32(1e10))))
    fmt.Println(v2, v94, v96)
    var v97 uint64 = 635
    v2 = v2
    for fi98 := 0; fi98 < 15; fi98++ {
        v97 = 18446744073709551615
        _ = fi98
    }
    switch (v94 + v94) {
    case "a":
        if true {
            v94 = ("\n\t" + (("foo bar" + v94) + v94))
            v96 = v2
        }
    default:
        var v99 uint32 = 2147483647
        _ = v99
    }
    v96 += v96
    var v100 string = (v94 + v94)
    switch "foo bar" {
    case "hello":
        delete(v1, (v100 + ("" + "")))
    case "a":
        if false {
            _ = len(v3)
            v94 = "hello"
            var v101 string = ((v94 + v94) + v94)
            _ = v101
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v94
    _ = v96
    _ = v97
}
