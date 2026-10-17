package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 int64 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16; F1 int32 }
type GenS20 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint) int {
    fmt.Println(p0)
    p0 -= 255
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint32 = 1
        _ = fi4
        _ = v5
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 uint = 776
    var v2 int32 = 1
    v3 := []string{"a", "", "a"}
    v1 = (^(v1))
    var v6 int16 = 0
    if true {
        v2 -= (v2 ^ (-(v2)))
    } else {
        var v7 uint8 = 2
        _ = v7
    }
    v2 = v2
    var v8 float32 = 0.1
    switch "hello" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch "foo bar" {
        case "a":
            fmt.Println(v6, v2)
        case "foo bar":
            var v9 uint64 = 2
            _ = v9
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v2 -= (v2 * (v2 & -539))
    if false {
        v3 = append(v3, "a")
        var v10 uint64 = 4294967295
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    switch (fi11 >> (((^(v1))) & 63)) {
                    case -1:
                        if false {
                            if (("foo bar" > ("a" + ("hello" + "hello"))) && true) {
                                switch "" {
                                case "\n\t":
                                    if (100 <= 127) {
                                        var v12 int8 = -1
                                        _ = v12
                                    } else {
                                        v1 = (^(v1))
                                        for fi13 := 0; fi13 < 14; fi13++ {
                                            if true {
                                                var v14 uint = v1
                                                switch (fi11 & genFn0((v14 & v1))) {
                                                case -2147483648:
                                                    var v15 uint = 2
                                                    var v16 uint16 = 422
                                                    _ = v15
                                                    _ = v16
                                                case 752:
                                                    v17 := GenS1{F0: 24, F1: ""}
                                                    _ = v17
                                                }
                                                _ = v14
                                            } else {
                                                v10 = uint64(v1)
                                                if ((-1 > int64(v6)) && false) {
                                                    switch (fi11 % ((fi11) | 1)) {
                                                    case 100:
                                                        fmt.Println(v2, v1, v6)
                                                    case 2147483647:
                                                        v18 := GenS1{F0: 0, F1: ("hello" + "a")}
                                                        _ = v18
                                                    default:
                                                        v2 += (-(v2))
                                                    }
                                                    var v19 uint = 519
                                                    if true {
                                                        var v20 bool = (false || (!true))
                                                        switch (fi11 & fi11) {
                                                        case 32768:
                                                            var v21 uint64 = (v10 << (((^(v1))) & 63))
                                                            _ = v21
                                                        case -2147483648:
                                                            var v22 bool = (65535 == 0)
                                                            _ = v22
                                                        case 4294967295:
                                                            if len(v3) > 0 {
                                                                v3[0] = "a"
                                                            }
                                                        }
                                                        v23 := GenS1{F0: -1, F1: ""}
                                                        fmt.Println(v19, v2)
                                                        fmt.Println(v2)
                                                        fmt.Println(v8)
                                                        fmt.Println(v2, v8)
                                                        _ = v20
                                                        _ = v23
                                                    } else {
                                                        if false {
                                                            switch ("a" + ("a" + "")) {
                                                            case "a":
                                                                v24 := GenS2{F0: 2, F1: 102}
                                                                _ = v24
                                                            case "":
                                                                v19 += 255
                                                            default:
                                                                switch ("\n\t" + ("hello" + "")) {
                                                                case "\n\t":
                                                                    v10 += v10
                                                                case "foo bar":
                                                                    if false {
                                                                        var v25 int8 = 42
                                                                        v26 := GenS3{F0: 255, F1: 0, F2: (v2 * 0)}
                                                                        _ = v25
                                                                        _ = v26
                                                                    }
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "":
                                                                var v27 uint32 = 0
                                                                v28 := GenS3{F0: 53, F1: 0, F2: (^(v2))}
                                                                _ = v27
                                                                _ = v28
                                                            case "a":
                                                                var v29 int64 = 198
                                                                v6 -= (v6 + v6)
                                                                _ = len(v3)
                                                                var v30 int = (fi13 % (((fi11 ^ fi11)) | 1))
                                                                var v31 int32 = (v2 - v2)
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            case "\n\t":
                                                                fmt.Println(v2, v19, fi13)
                                                            }
                                                        }
                                                        v6 = (v6 * v6)
                                                    }
                                                    _ = v19
                                                } else {
                                                    if (((^(v1)) < (v1 % (((^(v1))) | 1))) || true) {
                                                        if (0.1 <= v8) {
                                                            var v32 uint = v1
                                                            if true {
                                                                fmt.Println(v8, v8)
                                                                if (-916 != 0) {
                                                                    v10 += v10
                                                                    fmt.Println(v1, v6, v32)
                                                                    switch ("hello" + ("" + "a")) {
                                                                    case "\n\t":
                                                                        if false {
                                                                            if (184 < -9223372036854775808) {
                                                                                switch (^(fi13)) {
                                                                                case -2147483648:
                                                                                    if (65 < 61) {
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            if (-32768 != v6) {
                                                                                                v1 -= v1
                                                                                            } else {
                                                                                                v6 += (v6 / ((v6) | 1))
                                                                                                v3 = append(v3, "\n\t")
                                                                                                var v33 int16 = v6
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    var v34 uint32 = uint32(v10)
                                                                                                    _ = v34
                                                                                                }
                                                                                                var v35 int64 = 233
                                                                                                if (true || true) {
                                                                                                    v3 = append(v3, "a")
                                                                                                    v35 = 0
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            v10 = (v10 % (((v10 | (^(v10)))) | 1))
                                                                                                            var v36 float32 = (v8 - (v8 / v8))
                                                                                                            _ = v36
                                                                                                        } else {
                                                                                                            v37 := GenS3{F0: 111, F1: 42, F2: v2}
                                                                                                            if true {
                                                                                                                v38 := GenS2{F0: 35, F1: 2}
                                                                                                                if false {
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = "\n\t"
                                                                                                                    }
                                                                                                                    var v39 bool = ("hello" >= "")
                                                                                                                    _ = v39
                                                                                                                } else {
                                                                                                                    _ = len(v3)
                                                                                                                    switch int(v1) {
                                                                                                                    case 32768:
                                                                                                                        fmt.Println(v1, fi13)
                                                                                                                    case -44:
                                                                                                                        v33 -= 1000
                                                                                                                    default:
                                                                                                                        v40 := GenS4{F0: (fi11 + 9223372036854775807), F1: v10}
                                                                                                                        _ = v40
                                                                                                                    }
                                                                                                                }
                                                                                                                v35 = v35
                                                                                                                fmt.Println(fi13)
                                                                                                                var v41 int32 = (^(v2))
                                                                                                                v10 += v10
                                                                                                                _ = v38
                                                                                                                _ = v41
                                                                                                            } else {
                                                                                                                v42 := GenS5{F0: 2}
                                                                                                                _ = v42
                                                                                                            }
                                                                                                            v6 += v6
                                                                                                            _ = v37
                                                                                                        }
                                                                                                    } else {
                                                                                                        v8 -= v8
                                                                                                    }
                                                                                                } else {
                                                                                                    var v43 int64 = (^(v35))
                                                                                                    _ = v43
                                                                                                }
                                                                                                _ = v33
                                                                                                _ = v35
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                if false {
                                                                                                    var v44 string = "foo bar"
                                                                                                    _ = v44
                                                                                                }
                                                                                                var v45 uint32 = 170
                                                                                                fmt.Println(v10, v10)
                                                                                                if (!false) {
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = "\n\t"
                                                                                                    }
                                                                                                    v2 -= (-(v2))
                                                                                                }
                                                                                                switch genFn0((v32 << ((v32) & 63))) {
                                                                                                case 2147483647:
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = ""
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v2, v6, fi11)
                                                                                                }
                                                                                                _ = v45
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v8, v6)
                                                                                        }
                                                                                        if true {
                                                                                            if true {
                                                                                                v3 = append(v3, ("" + "a"))
                                                                                            } else {
                                                                                                var v46 uint32 = 1
                                                                                                _ = v46
                                                                                            }
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (("a" + "a") + ("\n\t" + ("hello" + "hello")))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch (("hello" + ("a" + ("foo bar" + "foo bar"))) + "\n\t") {
                                                                                case "hello":
                                                                                    v10 += (v10 & (^(v10)))
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "\n\t"
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(fi13)
                                                                                    }
                                                                                }
                                                                                v1 = v1
                                                                            }
                                                                            v1 = (v32 + (v32 - v1))
                                                                            v47 := GenS6{F0: (v2 % ((v2) | 1))}
                                                                            _ = v47
                                                                        } else {
                                                                            if ("" != ((("hello" + ("hello" + ("\n\t" + ""))) + "") + "\n\t")) {
                                                                                _ = len(v3)
                                                                                v8 += v8
                                                                                v48 := GenS6{F0: (v2 / ((v2) | 1))}
                                                                                _ = v48
                                                                            }
                                                                            _ = len(v3)
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v49 uint64 = v10
                                                                            _ = v49
                                                                        } else {
                                                                            fmt.Println(v2, v1, v2)
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(fi13, fi11, v10)
                                                                var v50 float32 = 1.5
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v51 int16 = v6
                                                                    _ = v51
                                                                case "hello":
                                                                    switch "" {
                                                                    case "":
                                                                        fmt.Println(fi11, fi11, v2)
                                                                        var v52 bool = ((false || (-645 == int64(fi13))) && true)
                                                                        _ = v52
                                                                    }
                                                                case "\n\t":
                                                                    if ((v6 - 32767) == (v6 % ((v6) | 1))) {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            switch ("hello" + "a") {
                                                                            case "a":
                                                                                v8 -= (v50 * float32((1.5 + (1.5 / 0.1))))
                                                                            case "":
                                                                                v8 -= (1e10 * ((v8 * (v8 / v50)) * float32(1.5)))
                                                                            case "\n\t":
                                                                                v53 := GenS6{F0: v2}
                                                                                _ = v53
                                                                            }
                                                                            var v54 uint32 = uint32(v10)
                                                                            switch 0 {
                                                                            case -1:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                var v55 float64 = float64((0.1 * v8))
                                                                                _ = v55
                                                                            }
                                                                            if (true && false) {
                                                                                fmt.Println(fi13, v32)
                                                                                v6 += v6
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v3 = append(v3, "a")
                                                                                case "a":
                                                                                    switch (fi11 | (fi13 >> (((v1 << (((v32 | v32)) & 63))) & 63))) {
                                                                                    case 32767:
                                                                                        var v56 int16 = (-(v6))
                                                                                        _ = v56
                                                                                    case 445:
                                                                                        v57 := GenS7{F0: (fi11 / ((genFn0(v1)) | 1)), F1: v2}
                                                                                        _ = v57
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v54
                                                                        case "foo bar":
                                                                            v2 -= (v2 * v2)
                                                                        case "a":
                                                                            v3 = append(v3, "foo bar")
                                                                        }
                                                                    }
                                                                default:
                                                                    var v58 uint64 = (v10 >> ((v32) & 63))
                                                                    _ = v58
                                                                }
                                                                if true {
                                                                    v59 := GenS4{F0: fi13, F1: v10}
                                                                    v60 := GenS6{F0: (^(v2))}
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                                var v61 int = (fi13 + (fi13 << ((v32) & 63)))
                                                                _ = v50
                                                                _ = v61
                                                            }
                                                            v3 = append(v3, ("\n\t" + ("hello" + "a")))
                                                            if true {
                                                                var v62 uint8 = 161
                                                                _ = v62
                                                            }
                                                            var v63 float32 = float32((0.0 * (1.5 + float64(1e10))))
                                                            _ = v32
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(fi13, v2)
                                                            if true {
                                                                if false {
                                                                    fmt.Println(v6, v10, fi11)
                                                                    v1 += uint(v10)
                                                                } else {
                                                                    _ = len(v3)
                                                                }
                                                                v1 -= (v1 * (v1 - 18446744073709551615))
                                                            }
                                                            v2 = int32(fi13)
                                                        }
                                                    }
                                                    _ = len(v3)
                                                }
                                                if ("foo bar" <= ("\n\t" + "")) {
                                                    var v64 int8 = -86
                                                    switch 0 {
                                                    case 0:
                                                        v65 := GenS4{F0: 9223372036854775807, F1: uint64(v1)}
                                                        _ = v65
                                                    case 32767:
                                                        v3 = append(v3, "hello")
                                                    case -2147483648:
                                                        fmt.Println(v1)
                                                    }
                                                    v66 := GenS4{F0: (fi11 >> ((255) & 63)), F1: 18446744073709551615}
                                                    _ = v64
                                                    _ = v66
                                                }
                                            }
                                            var v67 float32 = 3.14159
                                            _ = fi13
                                            _ = v67
                                        }
                                        v10 = v10
                                    }
                                case "a":
                                    _ = len(v3)
                                }
                                if false {
                                    var v68 int8 = 42
                                    _ = v68
                                }
                            }
                            v1 += 187
                        }
                    }
                }
                _ = len(v3)
                _ = fi11
            }
        }
        _ = v10
    } else {
        if true {
            fmt.Println(v6, v6, v8)
            if true {
                _ = len(v3)
                v6 -= v6
                for fi69 := 0; fi69 < 19; fi69++ {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = fi69
                }
            } else {
                if false {
                    if true {
                        if true {
                            if (false || (-32768 >= (v6 ^ v6))) {
                                var v70 uint8 = 208
                                for fi71 := 0; fi71 < 15; fi71++ {
                                    if (true || true) {
                                        var v72 float64 = 1e10
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                        for fi73 := 0; fi73 < 17; fi73++ {
                                            v70 -= 164
                                            _ = fi73
                                        }
                                        _ = v72
                                    } else {
                                        var v74 int64 = -798
                                        _ = v74
                                    }
                                    if true {
                                        switch (("" + "hello") + "") {
                                        case "hello":
                                            var v75 float32 = v8
                                            _ = v75
                                        }
                                    } else {
                                        var v76 int64 = -1
                                        var v77 int16 = v6
                                        if false {
                                            if (false && false) {
                                                v70 += (v70 << ((v1) & 63))
                                                var v78 uint16 = uint16(v1)
                                                fmt.Println(v78, v76)
                                                _ = v78
                                            }
                                        }
                                        _ = v76
                                        _ = v77
                                    }
                                    var v79 int64 = 42
                                    _ = fi71
                                    _ = v79
                                }
                                fmt.Println(v6)
                                _ = v70
                            }
                            if (114 < 42) {
                                v80 := GenS7{F0: genFn0(42), F1: (v2 & v2)}
                                var v81 string = ""
                                if len(v3) > 0 {
                                    v3[0] = v81
                                }
                                _ = v80
                                _ = v81
                            } else {
                                fmt.Println(v6, v1)
                            }
                        }
                        v8 += 0.1
                        if (!true) {
                            if false {
                                for fi82 := 0; fi82 < 20; fi82++ {
                                    fmt.Println(v6, v6, v8)
                                    fmt.Println(fi82, v6)
                                    _ = fi82
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                if (1 < -9223372036854775808) {
                                    if false {
                                        switch ("\n\t" + ("" + "hello")) {
                                        case "hello":
                                            v3 = append(v3, "a")
                                        case "foo bar":
                                            for fi83 := 0; fi83 < 12; fi83++ {
                                                v2 -= v2
                                                for fi84 := 0; fi84 < 8; fi84++ {
                                                    if len(v3) > 0 {
                                                        v3[0] = ("hello" + "a")
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = ""
                                                    }
                                                    _ = fi84
                                                }
                                                _ = fi83
                                            }
                                        case "":
                                            if (true && true) {
                                                var v85 int32 = int32(v6)
                                                var v86 bool = false
                                                _ = v85
                                                _ = v86
                                            }
                                        }
                                    } else {
                                        v6 = 282
                                    }
                                    _ = len(v3)
                                    v1 = v1
                                }
                                v87 := GenS3{F0: 161, F1: 0, F2: 42}
                                _ = v87
                            }
                            if (!true) {
                                var v88 int64 = int64(v2)
                                if ((false && true) && true) {
                                    v8 += v8
                                } else {
                                    v89 := GenS1{F0: 111, F1: ("\n\t" + ("hello" + ("hello" + "hello")))}
                                    if ("hello" != "foo bar") {
                                        v1 += (v1 >> ((v1) & 63))
                                    } else {
                                        for fi90 := 0; fi90 < 13; fi90++ {
                                            if true {
                                                v88 = v88
                                            }
                                            var v91 float32 = v8
                                            _ = fi90
                                            _ = v91
                                        }
                                        v8 = v8
                                    }
                                    _ = v89
                                }
                                _ = v88
                            } else {
                                if ((!true) && false) {
                                    if (!(false || ("a" <= ("foo bar" + (("" + "hello") + ("\n\t" + ("a" + "foo bar"))))))) {
                                        switch int(v2) {
                                        case -121:
                                            fmt.Println(v2)
                                        case 980:
                                            var v92 uint64 = 42
                                            _ = v92
                                        default:
                                            var v93 int = 825
                                            _ = v93
                                        }
                                    }
                                }
                                var v94 int16 = (v6 | 1)
                                _ = v94
                            }
                            v95 := GenS8{F0: v6, F1: 0}
                            _ = v95
                        }
                    } else {
                        for fi96 := 0; fi96 < 8; fi96++ {
                            for fi97 := 0; fi97 < 19; fi97++ {
                                fmt.Println(v6, v2)
                                var v98 string = "a"
                                _ = fi97
                                _ = v98
                            }
                            fmt.Println(v1, fi96)
                            fmt.Println(v2, v8, fi96)
                            _ = fi96
                        }
                        v8 += v8
                    }
                    if (true && (((1e10 * 1.5) + float64((0.0 * ((-2.25 * v8) / float32(-2.25))))) == (-2.25 - 0.1))) {
                        v99 := GenS8{F0: (^(v6)), F1: 32768}
                        v6 += (v6 + (^(v6)))
                        _ = v99
                    }
                } else {
                    v8 -= v8
                    if true {
                        var v100 int = -840
                        _ = v100
                    } else {
                        for fi101 := 0; fi101 < 13; fi101++ {
                            v102 := GenS4{F0: -665, F1: 1}
                            for fi103 := 0; fi103 < 10; fi103++ {
                                if (-2.25 >= ((0.0 * (1e10 - (1e10 + 0.0))) + 3.14159)) {
                                    fmt.Println(v1)
                                    switch (fi101 >> ((v1) & 63)) {
                                    case 32767:
                                        if ((v8 - v8) != 3.14159) {
                                            if len(v3) > 0 {
                                                v3[0] = "a"
                                            }
                                        } else {
                                            v1 = (v1 / ((16) | 1))
                                        }
                                    default:
                                        var v104 uint = (v1 + (v1 + v1))
                                        _ = v104
                                    }
                                } else {
                                    var v105 uint32 = 4294967295
                                    var v106 bool = (0.1 != v8)
                                    v105 -= 1
                                    switch "foo bar" {
                                    case "":
                                        switch ("" + "a") {
                                        case "hello":
                                            var v107 int32 = (v2 / (((^(v2))) | 1))
                                            _ = v107
                                        default:
                                            if (v106 && v106) {
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                                var v108 int = fi101
                                                var v109 int = 892
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    case "foo bar":
                                        v110 := GenS6{F0: v2}
                                        _ = v110
                                    }
                                    _ = v105
                                    _ = v106
                                }
                                v6 = v6
                                var v111 bool = (-128 > 42)
                                _ = fi103
                                _ = v111
                            }
                            switch "a" {
                            case "":
                                if (v6 != v6) {
                                    _ = len(v3)
                                    switch int(v2) {
                                    case -38:
                                        for fi112 := 0; fi112 < 13; fi112++ {
                                            v113 := GenS7{F0: genFn0(v1), F1: 2147483647}
                                            v114 := GenS1{F0: 100, F1: ("\n\t" + "foo bar")}
                                            if (-9223372036854775808 == -1) {
                                                v115 := GenS4{F0: 0, F1: 507}
                                                v8 = v8
                                                var v116 int32 = (v2 << (((v1 & (v1 << ((v1) & 63)))) & 63))
                                                _ = v115
                                                _ = v116
                                            }
                                            _ = fi112
                                            _ = v113
                                            _ = v114
                                        }
                                    default:
                                        var v117 uint64 = 4294967295
                                        _ = v117
                                    }
                                    v6 += -1
                                } else {
                                    for fi118 := 0; fi118 < 14; fi118++ {
                                        switch (fi101 << ((261) & 63)) {
                                        case 0:
                                            v8 += (((v8 / (v8 * (3.14159 / v8))) - 1.5) - float32(((1.5 / 3.14159) + 1e10)))
                                        default:
                                            v119 := GenS6{F0: (v2 + v2)}
                                            _ = v119
                                        }
                                        _ = fi118
                                    }
                                    switch fi101 {
                                    case -1:
                                        v120 := GenS6{F0: v2}
                                        _ = v120
                                    }
                                }
                            case "hello":
                                v1 = 4294967295
                            case "a":
                                v121 := GenS7{F0: 32768, F1: v2}
                                _ = v121
                            default:
                                var v122 uint = v1
                                _ = v122
                            }
                            _ = fi101
                            _ = v102
                        }
                        v8 -= ((v8 - v8) - float32(((0.1 - 0.1) - float64(v8))))
                    }
                    var v123 bool = ((!true) && (int16(v2) < v6))
                    _ = v123
                }
                if false {
                    v2 -= 2147483647
                    v8 += v8
                    v8 -= (((1e10 + (v8 / v8)) * v8) + v8)
                } else {
                    if false {
                        var v124 float64 = (0.0 * 0.0)
                        switch ("foo bar" + "") {
                        case "hello":
                            _ = len(v3)
                        case "":
                            switch -434 {
                            case 180:
                                switch "foo bar" {
                                case "a":
                                    if (42 != 42) {
                                        v125 := GenS9{F0: v2, F1: 542, F2: 0}
                                        v2 = (v2 & v2)
                                        v2 = (v2 + v2)
                                        _ = v125
                                    }
                                case "hello":
                                    v126 := GenS4{F0: 128, F1: 472}
                                    _ = v126
                                default:
                                    v1 -= (v1 - v1)
                                }
                            case 42:
                                var v127 int = 127
                                _ = v127
                            default:
                                switch "hello" {
                                case "a":
                                    var v128 int = -9223372036854775808
                                    _ = v128
                                case "foo bar":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "":
                                    if (v2 > (v2 ^ 42)) {
                                        if true {
                                            if false {
                                                v129 := GenS10{F0: 2147483647}
                                                _ = v129
                                            }
                                            var v130 int16 = (v6 ^ (^(v6)))
                                            _ = len(v3)
                                            _ = v130
                                        }
                                        switch 4294967295 {
                                        case -578:
                                            var v131 int32 = 42
                                            _ = v131
                                        }
                                    } else {
                                        v2 -= (v2 | v2)
                                        v132 := GenS1{F0: -1, F1: "\n\t"}
                                        for fi133 := 0; fi133 < 5; fi133++ {
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                            _ = fi133
                                        }
                                        _ = v132
                                    }
                                default:
                                    if (false || true) {
                                        fmt.Println(v8)
                                    } else {
                                        var v134 int = genFn0(72)
                                        v6 = v6
                                        v1 -= 65535
                                        _ = v134
                                    }
                                }
                            }
                        }
                        fmt.Println(v6)
                        _ = v124
                    }
                }
            }
            if true {
                v6 += (v6 ^ (v6 + 1))
                for fi135 := 0; fi135 < 20; fi135++ {
                    if false {
                        fmt.Println(fi135, v8, v8)
                    } else {
                        if false {
                            var v136 int8 = 1
                            _ = v136
                        } else {
                            for fi137 := 0; fi137 < 15; fi137++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi137
                            }
                            for fi138 := 0; fi138 < 1; fi138++ {
                                switch fi138 {
                                case 2147483647:
                                    if (v8 >= float32(1e10)) {
                                        if false {
                                            v1 = v1
                                        } else {
                                            v139 := GenS11{F0: uint8(v1)}
                                            if (!(!false)) {
                                                var v140 int16 = v6
                                                v141 := GenS11{F0: 182}
                                                _ = v140
                                                _ = v141
                                            } else {
                                                var v142 uint16 = 32767
                                                v142 -= v142
                                                switch "\n\t" {
                                                case "hello":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "":
                                                    fmt.Println(v1, fi135)
                                                case "\n\t":
                                                    var v143 int64 = 0
                                                    _ = v143
                                                }
                                                if true {
                                                    v3 = append(v3, "")
                                                    if false {
                                                        fmt.Println(v6)
                                                        if false {
                                                            var v144 float64 = 1.5
                                                            v6 -= v6
                                                            _ = v144
                                                        }
                                                    }
                                                    if (!(15 >= 42)) {
                                                        if (false || (1 <= 135)) {
                                                            var v145 int8 = -128
                                                            var v146 bool = (!false)
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                    }
                                                } else {
                                                    var v147 uint = (v1 - (^(v1)))
                                                    if ((v142 != v142) && (((-(fi135)) >= (fi135 << ((v1) & 63))) && (4294967295 <= 4294967295))) {
                                                        switch fi135 {
                                                        case -2147483648:
                                                            v148 := GenS1{F0: 100, F1: "\n\t"}
                                                            _ = v148
                                                        case 128:
                                                            fmt.Println(v1, v147)
                                                        default:
                                                            fmt.Println(v8, fi135, v2)
                                                        }
                                                    }
                                                    _ = v147
                                                }
                                                _ = v142
                                            }
                                            v149 := GenS12{F0: "foo bar"}
                                            _ = v139
                                            _ = v149
                                        }
                                        if (244 <= v6) {
                                            v2 = v2
                                        } else {
                                            var v150 bool = (1 >= 1)
                                            var v151 int64 = 111
                                            var v152 int = (fi135 % ((42) | 1))
                                            _ = v150
                                            _ = v151
                                            _ = v152
                                        }
                                    }
                                }
                                _ = fi138
                            }
                        }
                    }
                    _ = fi135
                }
            } else {
                for fi153 := 0; fi153 < 9; fi153++ {
                    for fi154 := 0; fi154 < 12; fi154++ {
                        var v155 uint = (^(v1))
                        switch "\n\t" {
                        case "hello":
                            var v156 uint64 = 2
                            _ = v156
                        case "a":
                            var v157 int16 = (v6 | (v6 * v6))
                            _ = v157
                        case "foo bar":
                            v3 = append(v3, "")
                        }
                        _ = fi154
                        _ = v155
                    }
                    _ = len(v3)
                    for fi158 := 0; fi158 < 13; fi158++ {
                        fmt.Println(v1)
                        v8 = v8
                        _ = fi158
                    }
                    _ = fi153
                }
                if (!(((((v8 / v8) / v8) + v8) == 0.0) || ((717 >= 2147483647) || (9223372036854775807 < 830)))) {
                    v2 = v2
                    for fi159 := 0; fi159 < 8; fi159++ {
                        switch (fi159 ^ (fi159 / ((fi159) | 1))) {
                        case -980:
                            for fi160 := 0; fi160 < 18; fi160++ {
                                if (true || false) {
                                    v2 = (v2 >> ((v1) & 63))
                                    if (false && (-76 >= -80)) {
                                        v3 = append(v3, "")
                                        var v161 uint16 = 65535
                                        switch "foo bar" {
                                        case "a":
                                            if false {
                                                v162 := GenS9{F0: (-(v2)), F1: 2, F2: 400}
                                                v1 = (v1 / ((v1) | 1))
                                                _ = v162
                                            } else {
                                                v163 := GenS13{F0: -685, F1: (^(v1))}
                                                var v164 int64 = 806
                                                _ = v163
                                                _ = v164
                                            }
                                        case "hello":
                                            if true {
                                                v8 += ((1e10 + -2.25) * 0.1)
                                                var v165 uint64 = 523
                                                v166 := GenS9{F0: (v2 % ((v2) | 1)), F1: (v165 + (v165 / (((v165 >> (((v1 << ((42) & 63))) & 63))) | 1))), F2: 822}
                                                _ = v165
                                                _ = v166
                                            } else {
                                                var v167 int8 = int8(v1)
                                                var v168 int = (fi160 & (fi159 & genFn0(v1)))
                                                _ = v167
                                                _ = v168
                                            }
                                        default:
                                            if (-2147483648 <= (v2 ^ 2147483647)) {
                                                if false {
                                                    v169 := GenS8{F0: 1, F1: (fi159 * (fi160 >> (((v1 << ((v1) & 63))) & 63)))}
                                                    v6 -= v6
                                                    if false {
                                                        v161 = 42
                                                        var v170 int = (^(fi160))
                                                        v6 += 682
                                                        _ = v170
                                                    } else {
                                                        var v171 string = ("hello" + (("foo bar" + ("\n\t" + "hello")) + ((("\n\t" + "a") + (("foo bar" + "foo bar") + "a")) + ("foo bar" + "foo bar"))))
                                                        _ = v171
                                                    }
                                                    if true {
                                                        v172 := GenS14{F0: v161, F1: -1, F2: 121}
                                                        _ = v172
                                                    } else {
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            fmt.Println(v1)
                                                        case "\n\t":
                                                            v1 -= (v1 >> ((v1) & 63))
                                                        default:
                                                            if false {
                                                                var v173 uint32 = 319
                                                                var v174 uint16 = v161
                                                                v3 = append(v3, ("hello" + (("" + "a") + (("\n\t" + "a") + "a"))))
                                                                _ = v173
                                                                _ = v174
                                                            }
                                                            v161 = 255
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v1 = v1
                                                        }
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v175 float32 = v8
                                                            var v176 uint32 = 0
                                                            v177 := GenS15{F0: "a"}
                                                            _ = v175
                                                            _ = v176
                                                            _ = v177
                                                        case "foo bar":
                                                            var v178 uint32 = 65535
                                                            _ = v178
                                                        case "":
                                                            var v179 float64 = -2.25
                                                            _ = v179
                                                        }
                                                        v161 = 606
                                                    }
                                                    v180 := GenS9{F0: v2, F1: 1, F2: 42}
                                                    _ = v169
                                                    _ = v180
                                                }
                                                if false {
                                                    fmt.Println(fi160)
                                                    if true {
                                                        if false {
                                                            var v181 string = "\n\t"
                                                            var v182 bool = (true && ((v2 ^ v2) != (v2 << (((^(v1))) & 63))))
                                                            _ = v181
                                                            _ = v182
                                                        } else {
                                                            v1 = uint(v161)
                                                            switch 1 {
                                                            case -738:
                                                                var v183 int64 = 423
                                                                _ = v183
                                                            case 109:
                                                                var v184 uint32 = 2147483647
                                                                _ = v184
                                                            case -128:
                                                                if true {
                                                                    if ((1 != 42) || true) {
                                                                        var v185 float64 = -2.25
                                                                        _ = v185
                                                                    }
                                                                }
                                                            default:
                                                                switch "" {
                                                                case "foo bar":
                                                                    switch -1 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v161)
                                                                    default:
                                                                        var v186 string = "foo bar"
                                                                        _ = v186
                                                                    }
                                                                case "a":
                                                                    v187 := GenS9{F0: (v2 >> ((v1) & 63)), F1: 639, F2: 381}
                                                                    _ = v187
                                                                }
                                                            }
                                                        }
                                                        v161 -= 609
                                                        v8 = v8
                                                        if ("" != "\n\t") {
                                                            v188 := GenS14{F0: v161, F1: -1, F2: 127}
                                                            fmt.Println(v1, v6)
                                                            _ = v188
                                                        }
                                                    }
                                                    v161 -= 32767
                                                } else {
                                                    v3 = append(v3, "")
                                                }
                                            }
                                        }
                                        _ = v161
                                    }
                                    _ = len(v3)
                                } else {
                                    fmt.Println(v6, v8)
                                    var v189 int64 = 750
                                    _ = v189
                                }
                                fmt.Println(v1, v2, v2)
                                _ = fi160
                            }
                        case 674:
                            switch "\n\t" {
                            case "foo bar":
                                if ((!((v6 != 1000) && false)) || false) {
                                    v2 = v2
                                    v190 := GenS13{F0: 9223372036854775807, F1: (v1 & v1)}
                                    if (1 <= 1) {
                                        var v191 uint64 = 1
                                        _ = v191
                                    }
                                    _ = v190
                                }
                            }
                        }
                        if ((((v6 + v6) > v6) || false) && (!true)) {
                            fmt.Println(v2)
                        }
                        _ = fi159
                    }
                }
                v2 += -2147483648
            }
        }
    }
    for fi192 := 0; fi192 < 3; fi192++ {
        v193 := GenS16{F0: 127, F1: -9223372036854775808}
        _ = fi192
        _ = v193
    }
    if (0.0 >= (v8 + v8)) {
        _ = len(v3)
    } else {
        switch "foo bar" {
        case "":
            switch 934 {
            case 42:
                if ((v6 + v6) >= v6) {
                    if false {
                        v6 = (v6 % ((v6) | 1))
                        if (false && true) {
                            v6 = v6
                            for fi194 := 0; fi194 < 11; fi194++ {
                                fmt.Println(v2, v8, fi194)
                                if true {
                                    var v195 int64 = 42
                                    v2 += 1
                                    _ = v195
                                }
                                _ = fi194
                            }
                        }
                    }
                } else {
                    v196 := GenS17{F0: (-(v6))}
                    v197 := GenS8{F0: v6, F1: 2147483647}
                    if false {
                        v8 = 3.14159
                        v3 = append(v3, (("\n\t" + "") + "a"))
                        _ = len(v3)
                    }
                    _ = v196
                    _ = v197
                }
            case 4294967295:
                for fi198 := 0; fi198 < 13; fi198++ {
                    v6 -= v6
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    _ = fi198
                }
            case 0:
                v3 = append(v3, ("hello" + ("hello" + ("foo bar" + (("" + ("hello" + "")) + "foo bar")))))
            default:
                for fi199 := 0; fi199 < 8; fi199++ {
                    switch fi199 {
                    case 2147483647:
                        for fi200 := 0; fi200 < 13; fi200++ {
                            v1 += (^(v1))
                            _ = fi200
                        }
                    case -1:
                        v201 := GenS18{F0: (-128 == -128)}
                        _ = v201
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v202 int32 = (v2 * (v2 | 0))
                    _ = fi199
                    _ = v202
                }
            }
        case "\n\t":
            v203 := GenS19{F0: (v6 - (v6 % ((1000) | 1))), F1: -535}
            _ = v203
        case "foo bar":
            var v204 uint16 = 255
            _ = v204
        default:
            v205 := GenS4{F0: 32767, F1: 1}
            _ = v205
        }
        v206 := GenS20{F0: v6, F1: 43}
        _ = v206
    }
    _ = len(v3)
    v8 += (float32(((((3.14159 - 1e10) / ((1e10 / (1e10 / 1.0)) - float64(v8))) * 0.0) / 3.14159)) - (v8 / 0.1))
    if true {
        v2 = v2
        var v207 bool = false
        var v208 uint16 = 1
        _ = v207
        _ = v208
    }
    fmt.Println(v8, v6)
    var v209 int32 = (v2 << ((v1) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v209:", v209)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v209
    _ = v3
    _ = v6
    _ = v8
}
