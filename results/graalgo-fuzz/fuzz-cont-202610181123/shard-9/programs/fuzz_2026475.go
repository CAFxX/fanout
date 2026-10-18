package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 bool; F1 int8; F2 uint8 }

func main() {
    v1 := []string{"a", "", "a"}
    var v2 uint64 = 2
    var v3 uint16 = uint16(v2)
    v4 := make(map[int]int)
    v2 -= 795
    for fi5 := 0; fi5 < 17; fi5++ {
        delete(v4, (^(fi5)))
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint32 = 65535
            var v8 uint8 = 107
            _ = len(v4)
            var v9 float64 = 3.14159
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        switch fi5 {
        case 127:
            if false {
                fmt.Println(fi5)
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 float64 = -2.25
                    _ = fi10
                    _ = v11
                }
                if true {
                    var v12 uint64 = 18446744073709551615
                    fmt.Println(fi5, fi5, v2)
                    _ = v12
                }
            }
        case 543:
            for fi13 := 0; fi13 < 14; fi13++ {
                v3 = v3
                if len(v1) > 0 {
                    v1[0] = ("hello" + "hello")
                }
                _ = fi13
            }
        }
        _ = fi5
    }
    v4[4294967295] = -220
    v14 := GenS1{F0: false, F1: -845}
    if true {
        if false {
            var v15 float32 = 1.5
            v15 = (v15 + v15)
            v15 += v15
            _ = v15
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: (v3 | (v3 % ((v3) | 1)))}
                if (true || (1 != 342)) {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if (-2147483648 != fi18) {
                                if (27 != 152) {
                                    fmt.Println(v2, fi16)
                                    v19 := GenS3{F0: (!false), F1: 42, F2: -32768}
                                    var v20 uint32 = 43
                                    _ = v19
                                    _ = v20
                                }
                                _ = len(v4)
                            } else {
                                if true {
                                    fmt.Println(v3)
                                    switch "a" {
                                    case "a":
                                        v3 = v3
                                    case "foo bar":
                                        v21 := GenS1{F0: false, F1: int32(v3)}
                                        _ = v21
                                    case "":
                                        if false {
                                            var v22 float32 = (1.5 - 1e10)
                                            _ = v22
                                        } else {
                                            var v23 int = (fi18 * -304)
                                            v24 := GenS1{F0: false, F1: 42}
                                            _ = v23
                                            _ = v24
                                        }
                                    default:
                                        var v25 uint8 = 255
                                        _ = v25
                                    }
                                } else {
                                    switch (("a" + "a") + ("" + "a")) {
                                    case "hello":
                                        var v26 uint16 = 0
                                        _ = v26
                                    }
                                }
                                var v27 uint = 2
                                switch fi16 {
                                case 32767:
                                    v3 = v3
                                default:
                                    if (false || (-208 >= -183)) {
                                        var v28 float32 = 3.14159
                                        var v29 float32 = 1e10
                                        switch -2147483648 {
                                        case 128:
                                            var v30 uint8 = 255
                                            _ = v30
                                        case 905:
                                            var v31 bool = (!(99 == 170))
                                            _ = v31
                                        default:
                                            v32 := GenS4{F0: -103, F1: -473}
                                            _ = v32
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                _ = v27
                            }
                            _ = fi18
                        }
                    case "foo bar":
                        var v33 int32 = 0
                        _ = v33
                    case "\n\t":
                        _ = len(v1)
                    }
                    v3 = 0
                    if ((!false) && (1 != -128)) {
                        v3 -= (^(v3))
                        var v34 uint32 = 1
                        _ = v34
                    } else {
                        if (!((1.5 * ((((float64(1.5) / 0.1) / 1.0) / (1.5 / -2.25)) - (-2.25 - 0.0))) != (3.14159 / 0.1))) {
                            fmt.Println(fi16, v2)
                        } else {
                            switch "hello" {
                            case "":
                                var v35 int8 = -128
                                _ = v35
                            default:
                                switch fi16 {
                                case -1:
                                    v36 := v4[(fi16 + fi16)]
                                    _ = v36
                                }
                            }
                            v37 := GenS2{F0: v3}
                            _ = v37
                        }
                        v38 := GenS5{F0: (v2 ^ (v2 - (v2 - (^(v2))))), F1: 69, F2: 721}
                        v39 := GenS6{F0: -128}
                        _ = v38
                        _ = v39
                    }
                } else {
                    if false {
                        var v40 uint16 = (^(v3))
                        v2 = v2
                        _ = v40
                    }
                    var v41 float32 = (1e10 - 1e10)
                    _ = v41
                }
                v2 = (v2 / ((v2) | 1))
                _ = fi16
                _ = v17
            }
            if false {
                v42 := GenS4{F0: 127, F1: 42}
                for fi43 := 0; fi43 < 19; fi43++ {
                    fmt.Println(v3, v2, v3)
                    switch (("hello" + (("hello" + "\n\t") + "")) + "foo bar") {
                    case "hello":
                        v3 = (^(v3))
                    }
                    _ = fi43
                }
                var v44 uint = uint(v2)
                _ = v42
                _ = v44
            }
            if true {
                if (!true) {
                    switch 433 {
                    case 0:
                        _ = len(v4)
                    case -234:
                        for fi45 := 0; fi45 < 6; fi45++ {
                            v46 := GenS7{F0: 9223372036854775807, F1: 201}
                            v3 -= (v3 & 215)
                            _ = fi45
                            _ = v46
                        }
                    case 782:
                        for fi47 := 0; fi47 < 10; fi47++ {
                            var v48 uint32 = 945
                            v48 = uint32(v2)
                            _ = fi47
                            _ = v48
                        }
                    }
                    switch ((("\n\t" + "hello") + "") + "") {
                    case "":
                        var v49 int16 = 1
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 9; fi50++ {
                        if false {
                            switch (((("\n\t" + (("" + "a") + "")) + (((("\n\t" + "") + "a") + "a") + ("a" + ("a" + (("hello" + "hello") + (("\n\t" + "foo bar") + "a")))))) + "hello") + "") {
                            case "":
                                delete(v4, (fi50 & (fi50 - (^(fi50)))))
                            }
                            var v51 uint64 = 42
                            _ = v51
                        }
                        fmt.Println(v2)
                        fmt.Println(fi50)
                        _ = fi50
                    }
                } else {
                    v2 = v2
                    if (26 < 128) {
                        var v52 uint8 = 96
                        v3 = (v3 << ((62) & 63))
                        _ = v52
                    }
                    v53 := v4[32768]
                    _ = v53
                }
                if false {
                    v3 += 83
                    v54 := GenS3{F0: false, F1: 1, F2: 42}
                    if (int16(v2) == 52) {
                        v55 := GenS2{F0: 700}
                        switch -9223372036854775808 {
                        case -1:
                            var v56 uint = 18446744073709551615
                            _ = v56
                        case -507:
                            if false {
                                if ((false && true) || (128 != 127)) {
                                    if (808 >= 282) {
                                        switch (("a" + ("a" + "hello")) + "a") {
                                        case "a":
                                            var v57 int8 = 42
                                            _ = v57
                                        default:
                                            for fi58 := 0; fi58 < 9; fi58++ {
                                                v59 := GenS7{F0: 9223372036854775807, F1: 0}
                                                v2 = v2
                                                _ = fi58
                                                _ = v59
                                            }
                                        }
                                        var v60 uint = 65535
                                        if (-123 < 1) {
                                            v3 = (v3 << ((4294967295) & 63))
                                            var v61 float32 = (0.1 + (3.14159 + 1e10))
                                            _ = v61
                                        }
                                        _ = v60
                                    } else {
                                        var v62 bool = (162 > 16)
                                        for fi63 := 0; fi63 < 14; fi63++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = fi63
                                        }
                                        var v64 uint16 = v3
                                        _ = v62
                                        _ = v64
                                    }
                                    fmt.Println(v2, v2)
                                }
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    fmt.Println(fi65, v2)
                                    v2 = (^(v2))
                                    var v66 uint64 = (^(v2))
                                    _ = fi65
                                    _ = v66
                                }
                                if (!true) {
                                    v1 = append(v1, ((("hello" + ("a" + (("foo bar" + "\n\t") + ""))) + "\n\t") + "hello"))
                                } else {
                                    v3 = v3
                                    v67 := GenS4{F0: 76, F1: 707}
                                    _ = v67
                                }
                            } else {
                                if false {
                                    switch ("hello" + "a") {
                                    case "hello":
                                        if false {
                                            v2 -= v2
                                            for fi68 := 0; fi68 < 15; fi68++ {
                                                if (-911 != 0) {
                                                    if (!(!false)) {
                                                        v3 = 1
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(v2, v2, fi68)
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v69 := GenS7{F0: -1, F1: 0}
                                                                _ = v69
                                                            }
                                                        default:
                                                            v70 := GenS2{F0: (v3 - (v3 >> ((18446744073709551615) & 63)))}
                                                            _ = v70
                                                        }
                                                    }
                                                    switch fi68 {
                                                    case 753:
                                                        if ((138 < 255) && ("\n\t" == (("\n\t" + "a") + "a"))) {
                                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                                var v72 float32 = 0.1
                                                                if ((false || ((v2 + 4294967295) != 18446744073709551615)) || false) {
                                                                    fmt.Println(v2)
                                                                }
                                                                _ = fi71
                                                                _ = v72
                                                            }
                                                        }
                                                    case -1:
                                                        if false {
                                                            if true {
                                                                var v73 uint = 1
                                                                _ = v73
                                                            } else {
                                                                fmt.Println(fi68)
                                                            }
                                                            var v74 bool = true
                                                            _ = v74
                                                        } else {
                                                            _ = len(v1)
                                                        }
                                                    case 373:
                                                        v75 := GenS8{F0: 42, F1: 126}
                                                        _ = v75
                                                    }
                                                }
                                                var v76 uint16 = (v3 ^ v3)
                                                v76 = 480
                                                _ = fi68
                                                _ = v76
                                            }
                                        }
                                    default:
                                        if false {
                                            v4[458] = -555
                                        } else {
                                            var v77 uint32 = 42
                                            _ = v77
                                        }
                                    }
                                    _ = len(v1)
                                    fmt.Println(v3)
                                }
                                v3 += v3
                            }
                        case 1:
                            for fi78 := 0; fi78 < 5; fi78++ {
                                var v79 float64 = 0.0
                                v3 = 659
                                _ = fi78
                                _ = v79
                            }
                        default:
                            _ = len(v1)
                        }
                        _ = v55
                    } else {
                        v4[2147483647] = 9223372036854775807
                    }
                    _ = v54
                }
            }
        }
        if true {
            for fi80 := 0; fi80 < 6; fi80++ {
                for fi81 := 0; fi81 < 17; fi81++ {
                    if ((true || true) && false) {
                        var v82 float64 = 3.14159
                        v3 += (v3 | (^(v3)))
                        _ = v82
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    var v83 float32 = 0.0
                    var v84 uint32 = 1
                    _ = fi81
                    _ = v83
                    _ = v84
                }
                if (("" <= "foo bar") || (!false)) {
                    var v85 float64 = (((1.5 / float64(((float32((float64(0.1) * float64(0.1))) + (3.14159 * 0.0)) - 1.5))) * 0.0) * 0.1)
                    v1 = append(v1, ("hello" + ""))
                    _ = v85
                }
                v86 := GenS3{F0: (!false), F1: 0, F2: 0}
                _ = fi80
                _ = v86
            }
        } else {
            v87 := GenS9{F0: (((true || ((-356 == 0) || true)) && true) || (-1 >= 1)), F1: 97, F2: 0}
            var v88 int64 = 0
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = v87
            _ = v88
        }
        fmt.Println(v2, v3)
    }
    var v89 int32 = -191
    if (2147483648 > 353) {
        v90 := v4[0]
        switch (v90 / ((4294967295) | 1)) {
        case 165:
            var v91 float64 = (0.1 - (((0.0 + 0.1) - ((-2.25 + 1.5) * 1e10)) * (1e10 + 3.14159)))
            _ = v91
        default:
            for fi92 := 0; fi92 < 7; fi92++ {
                var v93 int16 = 1000
                v3 = (v3 - v3)
                v3 = v3
                _ = fi92
                _ = v93
            }
        }
        _ = v90
    } else {
        for fi94 := 0; fi94 < 4; fi94++ {
            v95 := GenS3{F0: (((v89 * 42) == v89) && (false || false)), F1: -128, F2: -1}
            _ = fi94
            _ = v95
        }
        switch 42 {
        case -2147483648:
            var v96 uint16 = 42
            _ = v96
        case -1:
            fmt.Println(v89, v89)
        case 2147483647:
            var v97 uint16 = (^(v3))
            _ = v97
        default:
            v3 -= (^(v3))
        }
        delete(v4, 32767)
    }
    for fi98 := 0; fi98 < 9; fi98++ {
        v1 = append(v1, (("hello" + ("\n\t" + ("" + (("\n\t" + ("\n\t" + "")) + "\n\t")))) + ("\n\t" + ("" + "hello"))))
        _ = fi98
    }
    var v99 uint8 = 42
    v99 -= (^(v99))
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v89:", v89)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v89
    _ = v99
}
