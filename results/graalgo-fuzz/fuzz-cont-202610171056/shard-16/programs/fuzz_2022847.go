package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 int; F2 int64 }
type GenS6 struct { F0 uint }

func main() {
    var v1 int64 = 1
    var v2 int = 32768
    var v3 uint64 = 42
    var v4 int32 = -2147483648
    v5 := []uint32{0, 83}
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: (true || (((v2 & (v2 + v2)) > (v2 % ((fi6) | 1))) && false)), F1: false}
            var v8 bool = (!(((!(-2.25 >= float32(3.14159))) && true) && true))
            fmt.Println(v1)
            _ = fi6
            _ = v7
            _ = v8
        }
    case "a":
        for fi9 := 0; fi9 < 12; fi9++ {
            switch v2 {
            case 4294967295:
                v4 += 42
            case 32767:
                switch "" {
                case "hello":
                    var v10 uint32 = 199
                    _ = v10
                case "foo bar":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "\n\t":
                    v11 := GenS1{F0: false, F1: false}
                    _ = v11
                }
            default:
                v5 = append(v5, 241)
            }
            v4 = 0
            _ = fi9
        }
    case "hello":
        if ((false && false) && true) {
            v5 = append(v5, 42)
            var v12 uint8 = 125
            v13 := GenS2{F0: 0, F1: 0, F2: 42}
            _ = v12
            _ = v13
        } else {
            if ((-(v4)) == (-(v4))) {
                var v14 uint = 4294967295
                _ = v14
            }
            v2 = (v2 - (-(v2)))
        }
    default:
        if ((v2 | v2) != v2) {
            v4 = (v4 << ((4294967295) & 63))
            v2 = (v2 & v2)
        }
    }
    fmt.Println(v4, v3, v4)
    if true {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint64 = 2
            _ = fi15
            _ = v16
        }
        if true {
            switch (-(v2)) {
            case -9223372036854775808:
                var v17 float32 = float32(0.1)
                _ = v17
            }
            fmt.Println(v1, v2, v3)
        }
    }
    fmt.Println(v4, v1, v4)
    if true {
        _ = len(v5)
    } else {
        var v18 uint16 = 32767
        if ((-(v1)) > v1) {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 int64 = (v1 >> ((586) & 63))
                if (0.1 == 1e10) {
                    v21 := GenS1{F0: (65535 <= 65535), F1: true}
                    switch ((("a" + "\n\t") + ("\n\t" + (("hello" + (("" + "") + ("hello" + "\n\t"))) + "foo bar"))) + "a") {
                    case "":
                        for fi22 := 0; fi22 < 16; fi22++ {
                            switch ("a" + (("\n\t" + "foo bar") + ("\n\t" + ((((("a" + "") + ("a" + "foo bar")) + (("a" + "a") + ("a" + (("" + ("foo bar" + (("foo bar" + "hello") + ""))) + "a")))) + "foo bar") + "\n\t")))) {
                            case "a":
                                switch (("\n\t" + "foo bar") + "\n\t") {
                                case "hello":
                                    if true {
                                        var v23 string = ("a" + ("" + "a"))
                                        _ = v23
                                    } else {
                                        v1 = (v20 - v1)
                                        switch 2147483647 {
                                        case -1:
                                            if false {
                                                v24 := GenS2{F0: 100, F1: v18, F2: 65535}
                                                var v25 int8 = 127
                                                v26 := GenS1{F0: (((v3 >> ((159) & 63)) >= v3) && (((!(fi19 > v2)) || false) || (false || false))), F1: ("foo bar" >= "foo bar")}
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            }
                                        case 503:
                                            if false {
                                                switch (((("" + "foo bar") + "hello") + ("\n\t" + "hello")) + "foo bar") {
                                                case "hello":
                                                    v27 := GenS1{F0: false, F1: (((0.1 + ((-2.25 / (-2.25 + 1.5)) * 1e10)) < 1e10) && true)}
                                                    _ = v27
                                                }
                                            } else {
                                                v20 = 9223372036854775807
                                            }
                                        }
                                        var v28 int16 = 28
                                        if ((v20 >> ((42) & 63)) != (v1 + int64(v4))) {
                                            fmt.Println(fi19, v2)
                                            if ((-(fi19)) >= fi22) {
                                                if true {
                                                    var v29 uint = 18446744073709551615
                                                    switch int(v4) {
                                                    case 0:
                                                        fmt.Println(v18, v29)
                                                    case 128:
                                                        _ = len(v5)
                                                    }
                                                    v20 = (v1 / ((v1) | 1))
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        if len(v5) > 0 {
                                                            v5[0] = 402
                                                        }
                                                    case "a":
                                                        switch (-(fi19)) {
                                                        case 32767:
                                                            fmt.Println(v18)
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v2, fi19)
                                                    }
                                                    _ = v29
                                                }
                                                if false {
                                                    _ = len(v5)
                                                    var v30 bool = (v28 <= v28)
                                                    v31 := GenS2{F0: 53, F1: 255, F2: 373}
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    switch fi22 {
                                                    case -524:
                                                        fmt.Println(v4, v3)
                                                    case 0:
                                                        var v32 bool = (!(241 == 889))
                                                        _ = v32
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 0
                                                    }
                                                    var v33 float64 = 3.14159
                                                    v3 += v3
                                                    _ = v33
                                                }
                                            } else {
                                                v34 := GenS1{F0: (!((v28 << ((42) & 63)) >= 32767)), F1: true}
                                                if true {
                                                    fmt.Println(v28)
                                                }
                                                _ = v34
                                            }
                                        }
                                        _ = v28
                                    }
                                default:
                                    fmt.Println(v2, v3, fi22)
                                }
                            case "hello":
                                switch fi22 {
                                case -1:
                                    var v35 string = ""
                                    if true {
                                        switch (fi22 ^ (fi22 + v2)) {
                                        case 32768:
                                            if ((v3 * v3) == v3) {
                                                var v36 bool = (!(!(!(!false))))
                                                if len(v5) > 0 {
                                                    v5[0] = 4294967295
                                                }
                                                _ = v36
                                            }
                                        case 447:
                                            v37 := GenS3{F0: 127}
                                            _ = v37
                                        case -1:
                                            v5 = append(v5, 798)
                                            var v38 uint32 = 2147483647
                                            if len(v5) > 0 {
                                                v5[0] = v38
                                            }
                                            _ = v38
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = 65535
                                        }
                                        var v39 int16 = -130
                                        v3 += v3
                                        _ = v39
                                    } else {
                                        switch fi22 {
                                        case 127:
                                            var v40 uint64 = v3
                                            _ = v40
                                        case -1:
                                            var v41 float64 = 3.14159
                                            _ = v41
                                        }
                                        var v42 bool = false
                                        var v43 uint8 = 124
                                        _ = v42
                                        _ = v43
                                    }
                                    switch (-(fi19)) {
                                    case -34:
                                        v20 -= (v20 | -928)
                                    case 2147483647:
                                        if (!(42 > 493)) {
                                            if true {
                                                if len(v5) > 0 {
                                                    v5[0] = 42
                                                }
                                                var v44 string = ("" + ((v35 + (v35 + "a")) + ("\n\t" + "")))
                                                _ = v44
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    default:
                                        switch "a" {
                                        case "a":
                                            v2 += 64
                                        default:
                                            if ((v3 - (v3 - v3)) == 589) {
                                                if true {
                                                    if true {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v45 uint = 42
                                                        fmt.Println(v1, v1, fi19)
                                                        _ = v45
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = 65535
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            v2 += v2
                                                            v35 = ""
                                                            if ("a" != v35) {
                                                                var v46 uint = 255
                                                                _ = v46
                                                            } else {
                                                                v20 = v1
                                                                var v47 uint64 = (^(v3))
                                                                v48 := GenS3{F0: 255}
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                        case "":
                                                            v5 = append(v5, 2)
                                                        case "a":
                                                            fmt.Println(fi22, v4)
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case "a":
                                                        if false {
                                                            var v49 int8 = 95
                                                            v3 = (v3 * 0)
                                                            switch (fi22 ^ (fi19 | fi19)) {
                                                            case -296:
                                                                var v50 uint64 = v3
                                                                _ = v50
                                                            case 32767:
                                                                switch v35 {
                                                                case "":
                                                                    fmt.Println(fi19, fi22, v35)
                                                                default:
                                                                    var v51 string = (v35 + (((v35 + v35) + (v35 + "foo bar")) + (v35 + v35)))
                                                                    _ = v51
                                                                }
                                                            case -9223372036854775808:
                                                                var v52 float64 = 1e10
                                                                var v53 int32 = (v4 ^ (v4 | v4))
                                                                v5 = append(v5, 0)
                                                                _ = v52
                                                                _ = v53
                                                            default:
                                                                v35 = v35
                                                            }
                                                            _ = v49
                                                        } else {
                                                            v18 = v18
                                                            v5 = append(v5, 34)
                                                            v4 -= (v4 ^ -1)
                                                            v54 := GenS1{F0: true, F1: true}
                                                            _ = v54
                                                        }
                                                    case "foo bar":
                                                        var v55 uint32 = 612
                                                        _ = v55
                                                    default:
                                                        v56 := GenS3{F0: 255}
                                                        _ = v56
                                                    }
                                                } else {
                                                    fmt.Println(v1)
                                                    switch "" {
                                                    case "":
                                                        v1 -= v20
                                                    case "\n\t":
                                                        if (v4 == v4) {
                                                            v1 = -926
                                                        }
                                                    case "a":
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    fmt.Println(fi19, v3, v35)
                                                }
                                                var v57 uint16 = 2
                                                if (1e10 <= (float32(1.5) + 0.1)) {
                                                    switch v35 {
                                                    case "foo bar":
                                                        switch (v2 % ((v2) | 1)) {
                                                        case -9223372036854775808:
                                                            fmt.Println(fi22, v4)
                                                        case -1:
                                                            var v58 uint64 = (v3 | (v3 ^ (v3 | (^(v3)))))
                                                            _ = v58
                                                        default:
                                                            fmt.Println(v2, v35)
                                                        }
                                                    case "a":
                                                        var v59 string = ((v35 + ((v35 + (v35 + v35)) + v35)) + v35)
                                                        _ = v59
                                                    case "":
                                                        if false {
                                                            v60 := GenS4{F0: -99, F1: 1, F2: 42}
                                                            v61 := GenS2{F0: 1, F1: (^(v18)), F2: 18446744073709551615}
                                                            v57 = 2
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            v62 := GenS5{F0: 0, F1: 0, F2: 1}
                                                            _ = v62
                                                        }
                                                    default:
                                                        v63 := GenS5{F0: 32767, F1: (fi22 << ((363) & 63)), F2: v20}
                                                        _ = v63
                                                    }
                                                }
                                                switch v35 {
                                                case "foo bar":
                                                    if len(v5) > 0 {
                                                        v5[0] = 898
                                                    }
                                                default:
                                                    v1 += 1
                                                    var v64 float64 = (1.5 - -2.25)
                                                    var v65 int16 = 1000
                                                    switch (fi22 - (fi19 & (^(fi22)))) {
                                                    case -169:
                                                        fmt.Println(v2, v3, v64)
                                                    default:
                                                        fmt.Println(v35, v1)
                                                    }
                                                    _ = v64
                                                    _ = v65
                                                }
                                                switch v35 {
                                                case "\n\t":
                                                    _ = len(v5)
                                                }
                                                _ = v57
                                            }
                                        }
                                    }
                                    _ = v35
                                case 314:
                                    fmt.Println(v20, v18)
                                }
                            }
                            _ = fi22
                        }
                    default:
                        v66 := GenS4{F0: 100, F1: 42, F2: 237}
                        _ = v66
                    }
                    _ = v21
                } else {
                    v2 += fi19
                    var v67 uint64 = (v3 % (((^(v3))) | 1))
                    _ = v67
                }
                _ = fi19
                _ = v20
            }
            _ = len(v5)
        }
        _ = v18
    }
    v68 := GenS6{F0: 830}
    v69 := GenS5{F0: -892, F1: (-(v2)), F2: 9223372036854775807}
    v1 = v1
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
}
