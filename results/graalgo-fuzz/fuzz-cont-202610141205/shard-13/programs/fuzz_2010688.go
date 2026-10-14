package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 uint32, p1 string) string {
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 string = p1
        v7 := GenS1{F0: 1, F1: 42}
        v6 = v6
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int64 = -958
    v9 := GenS2{F0: 0, F1: 42}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{}
    var v3 uint64 = 4294967295
    v4 := []string{""}
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v10 int8 = 100
    var v11 int32 = -2147483648
    if (true && false) {
        var v12 float32 = (0.0 + 0.0)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            v3 -= (^(v3))
            var v14 uint = 255
            var v15 uint8 = 127
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint = 474
        v17 := GenS1{F0: 345, F1: 0}
        _ = v16
        _ = v17
    }
    fmt.Println(v11, v10, v10)
    v4 = append(v4, "a")
    var v18 float64 = ((float64(0.1) + v1) + (((1e10 - ((-2.25 / ((v1 * 1.5) * v1)) / float64(-2.25))) - 1e10) + v1))
    v3 = 545
    if (!false) {
        var v19 float64 = 1e10
        switch 891 {
        case 128:
            v10 += (v10 | v10)
        default:
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := GenS3{F0: 992, F1: 128, F2: false}
                if (27 == 128) {
                    fmt.Println(v10, fi20, v18)
                } else {
                    v22 := GenS4{F0: 128}
                    v10 = v10
                    var v23 int32 = v11
                    _ = v22
                    _ = v23
                }
                var v24 int = 737
                _ = fi20
                _ = v21
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            if len(v4) > 0 {
                v4[0] = ("" + "foo bar")
            }
            _ = fi25
        }
        _ = v19
    }
    switch "hello" {
    case "":
        if (-128 < (v10 << ((1) & 63))) {
            v1 = (v1 + v18)
            for fi26 := 0; fi26 < 9; fi26++ {
                if (!(!(v3 == v3))) {
                    switch (fi26 + 128) {
                    case 128:
                        v18 -= 0.0
                    }
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 int = fi26
                        _ = fi27
                        _ = v28
                    }
                    var v29 bool = true
                    _ = v29
                } else {
                    v10 += v10
                }
                v11 = (v11 << ((4294967295) & 63))
                fmt.Println(v10)
                _ = fi26
            }
            switch 2147483647 {
            case 4294967295:
                var v30 bool = (!(!false))
                _ = v30
            }
        }
    case "foo bar":
        if false {
            switch -1 {
            case -2147483648:
                switch ("hello" + "a") {
                case "":
                    v11 = (v11 ^ -1)
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        var v31 int32 = v11
                        _ = v31
                    default:
                        v32 := GenS3{F0: 9223372036854775807, F1: 219, F2: (42 < 0)}
                        _ = v32
                    }
                case "\n\t":
                    _ = len(v2)
                default:
                    var v33 bool = (65535 > 65535)
                    _ = v33
                }
            case 2147483647:
                v3 += 217
            case 32767:
                v34 := GenS2{F0: 1, F1: 65535}
                _ = v34
            }
            switch (("hello" + "\n\t") + "") {
            case "hello":
                if false {
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v11 = v11
                        _ = fi35
                    }
                    if (true || (3.14159 <= (v18 - (float64(0.1) * v1)))) {
                        for fi36 := 0; fi36 < 1; fi36++ {
                            if false {
                                var v37 int32 = v11
                                _ = v37
                            }
                            var v38 int = fi36
                            _ = fi36
                            _ = v38
                        }
                    }
                } else {
                    if (!(false && (!(-2147483648 < -9223372036854775808)))) {
                        switch ("" + ("" + "\n\t")) {
                        case "\n\t":
                            var v39 uint16 = 168
                            _ = v39
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            var v40 uint32 = 2147483647
                            _ = v40
                        }
                    }
                    _ = len(v2)
                }
            }
            v10 -= 0
        } else {
            for fi41 := 0; fi41 < 12; fi41++ {
                switch ("hello" + "\n\t") {
                case "hello":
                    switch ("foo bar" + "hello") {
                    case "a":
                        if (!false) {
                            var v42 uint64 = v3
                            if (true || false) {
                                fmt.Println(fi41)
                            }
                            var v43 uint32 = 434
                            _ = v42
                            _ = v43
                        } else {
                            for fi44 := 0; fi44 < 18; fi44++ {
                                if len(v4) > 0 {
                                    v4[0] = ("a" + ("foo bar" + ("" + "hello")))
                                }
                                _ = fi44
                            }
                            switch -9223372036854775808 {
                            case -731:
                                var v45 uint = 1
                                _ = v45
                            case -9223372036854775808:
                                switch (fi41 * fi41) {
                                case 32767:
                                    if true {
                                        v11 = (v11 / ((v11) | 1))
                                    }
                                case 127:
                                    if (!(v3 > v3)) {
                                        v46 := GenS5{F0: 65535}
                                        for fi47 := 0; fi47 < 8; fi47++ {
                                            v1 = 3.14159
                                            v4 = append(v4, ("\n\t" + "hello"))
                                            _ = fi47
                                        }
                                        var v48 uint8 = 200
                                        _ = v46
                                        _ = v48
                                    }
                                case 1:
                                    v49 := GenS3{F0: -9223372036854775808, F1: 128, F2: false}
                                    _ = v49
                                }
                            default:
                                v11 = (v11 - (v11 & v11))
                            }
                            v1 = v18
                        }
                    case "":
                        v1 = -2.25
                    default:
                        switch (("a" + (("a" + ((("a" + "hello") + (("hello" + ((("\n\t" + "hello") + (("a" + "hello") + "a")) + "\n\t")) + "hello")) + "hello")) + "\n\t")) + "hello") {
                        case "a":
                            switch (("hello" + "hello") + "hello") {
                            case "hello":
                                if (false || true) {
                                    var v50 int16 = 42
                                    v51 := GenS6{F0: v50, F1: int32(fi41)}
                                    _ = v50
                                    _ = v51
                                } else {
                                    if true {
                                        v11 = v11
                                    }
                                    switch "hello" {
                                    case "a":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "":
                                        fmt.Println(fi41, v1, v18)
                                    }
                                    switch fi41 {
                                    case 32767:
                                        if (0 <= 32767) {
                                            var v52 string = ("a" + "hello")
                                            var v53 bool = (false && true)
                                            var v54 int16 = 561
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        } else {
                                            if ((^(v11)) == v11) {
                                                if false {
                                                    v10 = int8(fi41)
                                                    for fi55 := 0; fi55 < 6; fi55++ {
                                                        var v56 float64 = (v18 - v18)
                                                        var v57 int64 = -1
                                                        _ = fi55
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    v10 = (v10 >> ((0) & 63))
                                                }
                                                switch fi41 {
                                                case -379:
                                                    v2 = append(v2, "hello")
                                                case 1:
                                                    var v58 uint64 = (v3 | v3)
                                                    _ = v58
                                                case 127:
                                                    v59 := GenS5{F0: 232}
                                                    _ = v59
                                                default:
                                                    if (255 <= 78) {
                                                        v60 := GenS7{F0: 42, F1: 1}
                                                        switch "" {
                                                        case "hello":
                                                            var v61 float32 = 1e10
                                                            _ = v61
                                                        case "":
                                                            v11 = (v11 + v11)
                                                        case "foo bar":
                                                            if true {
                                                                _ = len(v4)
                                                                for fi62 := 0; fi62 < 16; fi62++ {
                                                                    if false {
                                                                        fmt.Println(v18, v3)
                                                                        v63 := GenS7{F0: 0, F1: 2147483648}
                                                                        var v64 string = "a"
                                                                        _ = v63
                                                                        _ = v64
                                                                    } else {
                                                                        if (471 != 183) {
                                                                            var v65 int64 = 1
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    _ = fi62
                                                                }
                                                                fmt.Println(v1, v10)
                                                            } else {
                                                                var v66 int64 = 42
                                                                _ = v66
                                                            }
                                                        }
                                                        v67 := GenS3{F0: 0, F1: 127, F2: true}
                                                        _ = v60
                                                        _ = v67
                                                    }
                                                }
                                                if (0 <= 0) {
                                                    fmt.Println(v1, v1)
                                                } else {
                                                    fmt.Println(v18, v10)
                                                    v10 = v10
                                                }
                                            } else {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    fmt.Println(v10)
                                                case "\n\t":
                                                    var v68 bool = true
                                                    _ = v68
                                                case "hello":
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = "a"
                                                        }
                                                        v69 := GenS5{F0: 1}
                                                        v70 := GenS2{F0: 1000, F1: 32767}
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        v1 -= float64(3.14159)
                                                        v71 := GenS5{F0: 698}
                                                        _ = v71
                                                    }
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v72 int64 = 9223372036854775807
                                            _ = v72
                                        }
                                    case 70:
                                        v11 -= v11
                                    }
                                }
                            case "\n\t":
                                var v73 uint32 = 65535
                                _ = v73
                            case "":
                                for fi74 := 0; fi74 < 12; fi74++ {
                                    var v75 bool = true
                                    if v75 {
                                        v11 += v11
                                        var v76 int16 = 1000
                                        _ = v76
                                    } else {
                                        v18 = (float64(0.0) + (v1 + v18))
                                        var v77 int32 = v11
                                        v78 := GenS8{F0: 42, F1: 1}
                                        _ = v77
                                        _ = v78
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                            }
                        case "":
                            switch "hello" {
                            case "a":
                                v79 := GenS8{F0: 255, F1: -299}
                                _ = v79
                            case "foo bar":
                                for fi80 := 0; fi80 < 13; fi80++ {
                                    _ = len(v2)
                                    switch (-(fi80)) {
                                    case 32768:
                                        v18 = (v1 + v1)
                                    case -2147483648:
                                        if (false || (576 >= 255)) {
                                            v81 := GenS9{F0: (v10 >> ((2) & 63)), F1: 729}
                                            _ = v81
                                        }
                                    case 202:
                                        if (18446744073709551615 >= 0) {
                                            v1 -= v1
                                            fmt.Println(v3, v18, v1)
                                        } else {
                                            switch int(v10) {
                                            case 9223372036854775807:
                                                fmt.Println(v1, fi80)
                                            case -326:
                                                v82 := GenS9{F0: v10, F1: -1}
                                                _ = v82
                                            case 42:
                                                switch "\n\t" {
                                                case "a":
                                                    if true {
                                                        v2 = append(v2, ("hello" + "hello"))
                                                        var v83 int64 = -491
                                                        var v84 int = fi41
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                default:
                                                    var v85 uint64 = 4294967295
                                                    _ = v85
                                                }
                                            default:
                                                var v86 int64 = 742
                                                var v87 int16 = 1
                                                var v88 int32 = (-(v11))
                                                _ = v86
                                                _ = v87
                                                _ = v88
                                            }
                                            switch "" {
                                            case "hello":
                                                var v89 int16 = -165
                                                _ = v89
                                            case "":
                                                v90 := GenS4{F0: 155}
                                                _ = v90
                                            }
                                        }
                                    }
                                    _ = fi80
                                }
                            case "\n\t":
                                var v91 uint = 2
                                _ = v91
                            }
                        }
                    }
                default:
                    v92 := GenS10{F0: -154}
                    _ = v92
                }
                for fi93 := 0; fi93 < 18; fi93++ {
                    var v94 string = ""
                    _ = fi93
                    _ = v94
                }
                _ = len(v2)
                _ = fi41
            }
            for fi95 := 0; fi95 < 18; fi95++ {
                if ("hello" > "foo bar") {
                    v3 += v3
                    var v96 int32 = (v11 | (v11 * int32(v10)))
                    _ = v96
                }
                _ = fi95
            }
        }
    case "\n\t":
        v97 := GenS3{F0: 30, F1: 236, F2: (1 == 54)}
        _ = v97
    }
    v18 += v18
    if (false || true) {
        var v98 uint = 2
        if (!(0 != 9223372036854775807)) {
            for fi99 := 0; fi99 < 17; fi99++ {
                v11 = (v11 >> ((646) & 63))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v100 bool = true
                _ = fi99
                _ = v100
            }
        }
        _ = v98
    } else {
        v101 := GenS11{F0: "", F1: 2147483647, F2: 42}
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
