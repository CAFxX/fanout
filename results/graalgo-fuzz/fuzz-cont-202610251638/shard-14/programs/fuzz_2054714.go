package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 bool }

func main() {
    var v1 int = -1
    var v2 uint16 = 956
    v3 := make(map[string]string)
    var v4 int32 = -2147483648
    v5 := []string{}
    for fi6 := 0; fi6 < 7; fi6++ {
        v4 += (v4 % ((v4) | 1))
        _ = fi6
    }
    v7 := GenS1{F0: true, F1: -1, F2: 32768}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int64 = -484
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: "a"}
    var v11 float32 = (float32((1.5 / ((0.1 + 3.14159) + float64(0.1)))) - 0.1)
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 = v11
        for fi13 := 0; fi13 < 12; fi13++ {
            if true {
                switch (v1 << ((42) & 63)) {
                case 228:
                    switch 1 {
                    case 32767:
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                    case 0:
                        _ = len(v5)
                    case -9223372036854775808:
                        switch ("" + "hello") {
                        case "foo bar":
                            v14 := GenS2{F0: 42, F1: "foo bar"}
                            _ = v14
                        case "\n\t":
                            v15 := GenS3{F0: 42, F1: 170, F2: 2}
                            _ = v15
                        }
                    }
                    fmt.Println(v11, v2, fi13)
                    if (false && true) {
                        fmt.Println(v4, fi13)
                        _ = len(v5)
                        switch (-(fi13)) {
                        case -9223372036854775808:
                            v4 = (^(v4))
                        case 989:
                            v1 += v1
                        default:
                            v16 := GenS3{F0: 781, F1: 128, F2: 4294967295}
                            if false {
                                if false {
                                    switch (fi12 - fi12) {
                                    case 2147483647:
                                        v4 += v4
                                    case 127:
                                        _ = len(v5)
                                        fmt.Println(v4)
                                    }
                                }
                            }
                            switch v1 {
                            case -2147483648:
                                var v17 int8 = -23
                                _ = v17
                            case -9223372036854775808:
                                fmt.Println(v11)
                            default:
                                v11 += float32(0.0)
                            }
                            _ = v16
                        }
                    }
                case 773:
                    fmt.Println(fi12, v11)
                case -8:
                    v11 += ((v11 / 1.5) + (v11 / v11))
                }
                v4 = v4
            }
            v4 += 0
            _ = fi13
        }
        _ = fi12
    }
    if (1 <= 1) {
        if (true || (!true)) {
            if (-73 <= 4) {
                var v18 int = 0
                _ = v18
            }
        } else {
            v3["\n\t"] = "hello"
            v3[""] = "hello"
            v19 := GenS2{F0: 445, F1: "hello"}
            _ = v19
        }
        delete(v3, ((("\n\t" + "foo bar") + ("foo bar" + "")) + "a"))
        fmt.Println(v2, v4)
    }
    v2 = v2
    var v20 uint = 613
    for fi21 := 0; fi21 < 10; fi21++ {
        _ = len(v5)
        v2 -= (v2 | 823)
        switch (^(v1)) {
        case 1:
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 bool = true
                if false {
                    var v24 int64 = 9223372036854775807
                    var v25 float32 = 0.0
                    _ = v24
                    _ = v25
                }
                switch "a" {
                case "":
                    fmt.Println(v20)
                case "\n\t":
                    if v23 {
                        fmt.Println(v4, fi21, v1)
                        var v26 int16 = 1000
                        _ = v26
                    }
                }
                _ = fi22
                _ = v23
            }
        case -567:
            for fi27 := 0; fi27 < 9; fi27++ {
                if false {
                    fmt.Println(v1, v2, fi27)
                    v2 = 32767
                } else {
                    var v28 string = "a"
                    v3[((("hello" + v28) + "foo bar") + "\n\t")] = "foo bar"
                    _ = v28
                }
                delete(v3, "hello")
                _ = fi27
            }
        }
        _ = fi21
    }
    _ = len(v5)
    var v29 uint64 = 0
    for fi30 := 0; fi30 < 15; fi30++ {
        switch "" {
        case "a":
            switch "foo bar" {
            case "foo bar":
                for fi31 := 0; fi31 < 9; fi31++ {
                    v20 = (v20 ^ 524)
                    v2 -= v2
                    _ = fi31
                }
            }
        default:
            var v32 float32 = 0.0
            _ = v32
        }
        v20 = v20
        _ = fi30
    }
    switch ("hello" + "") {
    case "":
        var v33 uint32 = 4294967295
        _ = v33
    case "hello":
        v34 := v3["hello"]
        _ = v34
    case "a":
        if (true || true) {
            v4 -= (v4 + 0)
            v1 += (v1 - (v1 << ((v20) & 63)))
            v35 := GenS4{F0: true}
            _ = v35
        }
    }
    if (!false) {
        var v36 bool = false
        switch "a" {
        case "":
            if ((v1 >> (((v20 << (((v20 & v20)) & 63))) & 63)) != (v1 & -656)) {
                var v37 int = 32767
                if ((^(v2)) < v2) {
                    fmt.Println(v2, v37, v36)
                } else {
                    var v38 float64 = 0.1
                    if (1 <= 32767) {
                        for fi39 := 0; fi39 < 8; fi39++ {
                            for fi40 := 0; fi40 < 2; fi40++ {
                                var v41 uint32 = 4294967295
                                switch "" {
                                case "a":
                                    v37 = 127
                                case "":
                                    switch (-(fi40)) {
                                    case 9223372036854775807:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case -1:
                                        v42 := GenS5{F0: int8(v41)}
                                        _ = v42
                                    case 42:
                                        fmt.Println(v37, v4)
                                        v36 = v36
                                    }
                                }
                                var v43 uint32 = (v41 * v41)
                                _ = fi40
                                _ = v41
                                _ = v43
                            }
                            _ = fi39
                        }
                    }
                    _ = v38
                }
                _ = v37
            } else {
                v44 := GenS6{F0: (v36 || false)}
                for fi45 := 0; fi45 < 10; fi45++ {
                    fmt.Println(v2, v1)
                    v46 := GenS5{F0: 100}
                    for fi47 := 0; fi47 < 4; fi47++ {
                        if v36 {
                            var v48 float32 = v11
                            _ = v48
                        } else {
                            v4 += v4
                        }
                        var v49 int8 = 1
                        _ = fi47
                        _ = v49
                    }
                    _ = fi45
                    _ = v46
                }
                _ = v44
            }
        case "a":
            if (1 <= -9223372036854775808) {
                if (!((v4 > -748) || ((v2 / ((0) | 1)) <= (v2 << (((^(v20))) & 63))))) {
                    switch v1 {
                    case -1:
                        v20 += 102
                    case 0:
                        v4 = (v4 - 2147483647)
                    }
                }
            }
        case "hello":
            v29 = 1
        }
        _ = v36
    } else {
        v50 := GenS5{F0: -128}
        var v51 int = v1
        _ = v50
        _ = v51
    }
    if ((v2 + (^(v2))) == (v2 >> (((v20 ^ v20)) & 63))) {
        switch (((("" + "foo bar") + ((("foo bar" + ("hello" + "hello")) + "foo bar") + ("\n\t" + "foo bar"))) + "foo bar") + ("\n\t" + ("a" + "a"))) {
        case "foo bar":
            for fi52 := 0; fi52 < 20; fi52++ {
                v20 += 18446744073709551615
                _ = fi52
            }
        case "hello":
            for fi53 := 0; fi53 < 3; fi53++ {
                var v54 uint16 = v2
                _ = fi53
                _ = v54
            }
        case "":
            v20 += (v20 ^ v20)
        default:
            for fi55 := 0; fi55 < 15; fi55++ {
                v56 := GenS3{F0: 1000, F1: 5, F2: 4294967295}
                _ = fi55
                _ = v56
            }
        }
        _ = len(v5)
        for fi57 := 0; fi57 < 5; fi57++ {
            v11 -= (v11 - (v11 / 1.0))
            v1 += v1
            if true {
                v11 += (float32(-2.25) * v11)
                v20 = (v20 % ((1) | 1))
                var v58 string = (("hello" + "hello") + "")
                _ = v58
            } else {
                if false {
                    if true {
                        switch "hello" {
                        case "":
                            _ = len(v5)
                        }
                        var v59 int32 = v4
                        _ = v59
                    } else {
                        var v60 uint32 = 2147483647
                        fmt.Println(v60, v4, v4)
                        for fi61 := 0; fi61 < 15; fi61++ {
                            if true {
                                _ = len(v3)
                                switch (v1 >> (((v20 % ((0) | 1))) & 63)) {
                                case 0:
                                    var v62 uint64 = v29
                                    _ = v62
                                case 4294967295:
                                    if (v29 <= 2) {
                                        v2 = v2
                                        v63 := GenS7{F0: 1}
                                        _ = len(v5)
                                        fmt.Println(v4, fi61)
                                        _ = v63
                                    }
                                default:
                                    v29 = 261
                                }
                                var v64 uint16 = v2
                                v2 = (v2 >> (((v20 + (v20 & v20))) & 63))
                                _ = v64
                            } else {
                                switch fi57 {
                                case 1:
                                    var v65 string = ""
                                    _ = v65
                                case 127:
                                    v66 := GenS8{F0: true}
                                    _ = v66
                                }
                                v67 := GenS5{F0: 95}
                                _ = v67
                            }
                            _ = fi61
                        }
                        _ = v60
                    }
                    v68 := GenS9{F0: 707, F1: (0 >= -9223372036854775808)}
                    var v69 uint16 = (v2 & v2)
                    _ = v68
                    _ = v69
                } else {
                    fmt.Println(v1, fi57)
                }
                for fi70 := 0; fi70 < 17; fi70++ {
                    v20 = (v20 - v20)
                    delete(v3, ("foo bar" + "a"))
                    _ = fi70
                }
            }
            _ = fi57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
