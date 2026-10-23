package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int64; F1 string; F2 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 0
    v3 := []string{"a", "", ("hello" + "hello")}
    v4 := make(map[string]int)
    var v5 uint8 = 0
    var v6 string = ""
    switch v6 {
    case "":
        switch ("hello" + ((v6 + ("foo bar" + v6)) + v6)) {
        case "":
            fmt.Println(v6, v5, v2)
        case "a":
            v2 = v2
        case "hello":
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        }
    case "a":
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v8 := GenS1{F0: 0, F1: (v5 == (v5 | (v5 * (^(v5))))), F2: 1}
                _ = v8
            }
            switch "a" {
            case "hello":
                if (true || (535 <= 797)) {
                    fmt.Println(fi7, v2, v2)
                    if (2147483647 > 378) {
                        delete(v4, v6)
                        var v9 float64 = ((-2.25 * 0.1) / 3.14159)
                        fmt.Println(v6)
                        _ = v9
                    } else {
                        if ((127 <= 42) || (1 > 2147483647)) {
                            v10 := GenS2{F0: 129, F1: 2, F2: true}
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 15; fi11++ {
                                var v12 uint32 = 2
                                _ = fi11
                                _ = v12
                            }
                            for fi13 := 0; fi13 < 2; fi13++ {
                                if ((!(true && (false && (4294967295 != 0)))) && false) {
                                    v14 := GenS1{F0: v2, F1: (!(0.0 >= 1e10)), F2: 1}
                                    if true {
                                        if (false || ((^(fi7)) > (fi13 + fi13))) {
                                            v5 = v5
                                            var v15 bool = (2 != 65535)
                                            var v16 float32 = ((0.1 + (float32((((((1.5 - 0.1) - 1e10) * 0.0) - -2.25) + 0.1)) * 0.1)) + 1e10)
                                            _ = v15
                                            _ = v16
                                        }
                                        delete(v4, v6)
                                    } else {
                                        v2 = v2
                                    }
                                    var v17 float32 = 1.5
                                    _ = v14
                                    _ = v17
                                }
                                _ = fi13
                            }
                            v18 := GenS1{F0: v2, F1: false, F2: -1}
                            _ = v18
                        }
                        v5 = v5
                        var v19 uint64 = 42
                        _ = v19
                    }
                }
            case "\n\t":
                var v20 uint16 = 776
                _ = v20
            case "":
                var v21 int64 = v2
                _ = v21
            }
            _ = len(v3)
            _ = fi7
        }
    default:
        fmt.Println(v6, v6)
    }
    v22 := GenS1{F0: 0, F1: (false || false), F2: 25}
    for fi23 := 0; fi23 < 8; fi23++ {
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = ""
        }
        switch "a" {
        case "a":
            var v24 uint32 = 0
            _ = v24
        case "\n\t":
            v25 := GenS2{F0: 32767, F1: 4, F2: (-975 != -516)}
            _ = v25
        }
        _ = fi23
    }
    delete(v1, "a")
    for fi26 := 0; fi26 < 12; fi26++ {
        fmt.Println(v5)
        v2 = v2
        v27 := v4[v6]
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v5 = v5
        if false {
            switch (v6 + "\n\t") {
            case "hello":
                switch "\n\t" {
                case "hello":
                    fmt.Println(v6, v2, v6)
                }
            default:
                var v29 int64 = int64(v5)
                _ = v29
            }
        } else {
            v5 -= v5
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            v5 -= 2
            fmt.Println(v6, fi28)
            v5 += v5
            _ = fi30
        }
        _ = fi28
    }
    v5 = 127
    if true {
        v31 := v1[v6]
        _ = v31
    } else {
        switch ("foo bar" + v6) {
        case "a":
            v32 := GenS2{F0: 32767, F1: 2147483648, F2: true}
            _ = v32
        case "foo bar":
            v33 := GenS2{F0: 797, F1: 4294967295, F2: (111 != -1)}
            _ = v33
        }
    }
    _ = len(v3)
    if (true && false) {
        delete(v1, "hello")
    }
    if (42 == 32767) {
        v4["a"] = -303
        switch v6 {
        case "a":
            if true {
                for fi34 := 0; fi34 < 8; fi34++ {
                    v6 = (v6 + v6)
                    v35 := GenS2{F0: 0, F1: 484, F2: false}
                    v5 += v5
                    _ = fi34
                    _ = v35
                }
            } else {
                if true {
                    var v36 uint64 = 0
                    _ = v36
                } else {
                    if true {
                        delete(v4, ("foo bar" + "a"))
                    } else {
                        var v37 int64 = v2
                        v2 = (v37 - (v37 | v2))
                        _ = v37
                    }
                }
                if true {
                    if (-86 == 0) {
                        v38 := GenS2{F0: 32767, F1: 13, F2: (v6 == ("foo bar" + ("a" + ("foo bar" + v6))))}
                        switch v6 {
                        case "hello":
                            switch ("a" + v6) {
                            case "a":
                                switch -1 {
                                case 42:
                                    switch v6 {
                                    case "\n\t":
                                        fmt.Println(v5)
                                    case "a":
                                        if (962 == 32768) {
                                            if (!true) {
                                                var v39 uint8 = v5
                                                _ = v39
                                            }
                                            switch int(v5) {
                                            case 4294967295:
                                                v6 = ("" + v6)
                                            case 644:
                                                var v40 int32 = 2147483647
                                                _ = v40
                                            case -1:
                                                if (uint64(v5) >= 2) {
                                                    switch v6 {
                                                    case "":
                                                        switch 32768 {
                                                        case 128:
                                                            fmt.Println(v6, v2)
                                                        case 32768:
                                                            switch v6 {
                                                            case "a":
                                                                fmt.Println(v5)
                                                            case "\n\t":
                                                                v5 += (^(v5))
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                        case -9223372036854775808:
                                                            var v41 int = 685
                                                            _ = v41
                                                        }
                                                    case "hello":
                                                        for fi42 := 0; fi42 < 2; fi42++ {
                                                            switch v6 {
                                                            case "":
                                                                v5 = (v5 ^ v5)
                                                            case "a":
                                                                switch 4294967295 {
                                                                case 42:
                                                                    v5 = (v5 | 255)
                                                                case -9223372036854775808:
                                                                    v6 = "a"
                                                                case 32768:
                                                                    var v43 int8 = 0
                                                                    _ = v43
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v6)
                                                            }
                                                            _ = fi42
                                                        }
                                                    case "a":
                                                        var v44 float32 = -2.25
                                                        _ = v44
                                                    }
                                                    var v45 string = ((((v6 + v6) + v6) + "") + "a")
                                                    _ = v45
                                                }
                                            }
                                            v6 = "\n\t"
                                        }
                                    case "hello":
                                        v46 := GenS3{F0: true, F1: 2, F2: 2147483648}
                                        _ = v46
                                    default:
                                        for fi47 := 0; fi47 < 16; fi47++ {
                                            delete(v4, (v6 + v6))
                                            v6 = v6
                                            _ = fi47
                                        }
                                    }
                                case 9223372036854775807:
                                    v5 = 128
                                case 127:
                                    v48 := GenS3{F0: (false && true), F1: 332, F2: 514}
                                    _ = v48
                                }
                            }
                        case "":
                            if (false || true) {
                                switch 1 {
                                case -398:
                                    _ = len(v4)
                                }
                                var v49 int64 = v2
                                v50 := GenS3{F0: (-353 == 0), F1: 42, F2: 2}
                                _ = v49
                                _ = v50
                            }
                        default:
                            for fi51 := 0; fi51 < 7; fi51++ {
                                if (v5 > 2) {
                                    v52 := GenS2{F0: 0, F1: 836, F2: true}
                                    v6 = ("foo bar" + v6)
                                    v53 := GenS2{F0: 1000, F1: 1, F2: (1 < -1)}
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 uint16 = 42
                                    _ = v54
                                }
                                _ = fi51
                            }
                        }
                        _ = v38
                    }
                    switch "hello" {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    default:
                        var v55 int32 = 1
                        _ = v55
                    }
                }
                v6 = ""
            }
        default:
            if false {
                v56 := v1[((("hello" + (v6 + (("hello" + ((v6 + (v6 + ((v6 + "foo bar") + (v6 + ((v6 + v6) + v6))))) + "foo bar")) + "foo bar"))) + (v6 + "")) + (v6 + v6))]
                if false {
                    for fi57 := 0; fi57 < 14; fi57++ {
                        v6 = "\n\t"
                        var v58 uint8 = 86
                        _ = fi57
                        _ = v58
                    }
                    var v59 uint8 = 42
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 7; fi60++ {
                        var v61 int32 = 0
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = fi60
                        _ = v61
                    }
                }
                delete(v4, "")
                _ = v56
            }
        }
    } else {
        fmt.Println(v6)
    }
    for fi62 := 0; fi62 < 1; fi62++ {
        v5 += (v5 >> ((4294967295) & 63))
        v3 = append(v3, (v6 + v6))
        _ = fi62
    }
    var v63 int = 32767
    if false {
        v63 += v63
    } else {
        for fi64 := 0; fi64 < 10; fi64++ {
            if false {
                switch (v6 + "") {
                case "a":
                    fmt.Println(fi64, v5)
                default:
                    if (v5 >= (v5 << ((65535) & 63))) {
                        var v65 int8 = -38
                        if (true && (3.14159 < 0.1)) {
                            v63 = fi64
                            v66 := GenS4{F0: -126, F1: 0, F2: 82}
                            _ = v66
                        } else {
                            for fi67 := 0; fi67 < 14; fi67++ {
                                delete(v1, v6)
                                _ = fi67
                            }
                            if ((!(v63 < v63)) && false) {
                                v68 := GenS3{F0: false, F1: 1, F2: 2147483648}
                                var v69 uint = 0
                                _ = v68
                                _ = v69
                            } else {
                                var v70 string = (("foo bar" + v6) + v6)
                                var v71 int32 = 0
                                _ = v70
                                _ = v71
                            }
                            var v72 int = (fi64 + (^(v63)))
                            _ = v72
                        }
                        switch int(v2) {
                        case -281:
                            for fi73 := 0; fi73 < 6; fi73++ {
                                fmt.Println(v5, v65, v2)
                                _ = fi73
                            }
                        default:
                            v65 = v65
                        }
                        _ = v65
                    }
                }
                switch ((v6 + v6) + v6) {
                case "foo bar":
                    if false {
                        v6 = (v6 + ("a" + "foo bar"))
                    }
                case "":
                    v6 = "\n\t"
                }
                v74 := GenS5{F0: -38, F1: ("\n\t" + v6), F2: -2147483648}
                _ = v74
            }
            _ = fi64
        }
        v5 -= 128
        v75 := GenS3{F0: true, F1: 142, F2: 908}
        _ = v75
    }
    v63 += v63
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
