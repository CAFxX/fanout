package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.1
    var v3 string = ""
    v4 := GenS1{F0: 685}
    for fi5 := 0; fi5 < 13; fi5++ {
        v3 = v3
        _ = fi5
    }
    v2 = v2
    var v6 string = v3
    fmt.Println(v2, v3)
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!(-1 > -479)) {
            switch "" {
            case "":
                var v8 float32 = v2
                _ = v8
            case "a":
                switch v3 {
                case "\n\t":
                    switch fi7 {
                    case 706:
                        v2 += v2
                    case 42:
                        switch (v6 + v6) {
                        case "":
                            var v9 int64 = -9223372036854775808
                            _ = v9
                        case "a":
                            for fi10 := 0; fi10 < 14; fi10++ {
                                fmt.Println(v6, v3)
                                switch (v6 + "hello") {
                                case "\n\t":
                                    v1[v6] = (fi7 != 42)
                                case "":
                                    var v11 uint8 = 29
                                    _ = v11
                                default:
                                    var v12 float64 = -2.25
                                    _ = v12
                                }
                                var v13 uint8 = 149
                                _ = fi10
                                _ = v13
                            }
                        default:
                            for fi14 := 0; fi14 < 2; fi14++ {
                                v15 := GenS1{F0: 0}
                                v16 := v1[v6]
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        }
                    case 32768:
                        v17 := GenS1{F0: 32767}
                        _ = v17
                    }
                case "foo bar":
                    v18 := GenS2{F0: fi7}
                    _ = v18
                case "hello":
                    v19 := GenS3{F0: -376}
                    _ = v19
                }
            }
            var v20 uint16 = 32767
            _ = v20
        }
        _ = fi7
    }
    if ((!((32767 > 32767) && false)) && false) {
        switch "hello" {
        case "":
            if (true && true) {
                if false {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v6, v6, v3)
                    }
                }
            } else {
                for fi21 := 0; fi21 < 6; fi21++ {
                    switch fi21 {
                    case 32768:
                        v3 = v3
                    case 386:
                        v22 := GenS2{F0: 42}
                        _ = v22
                    }
                    _ = len(v1)
                    var v23 uint = 65535
                    _ = fi21
                    _ = v23
                }
            }
        case "hello":
            switch (v3 + "hello") {
            case "hello":
                for fi24 := 0; fi24 < 14; fi24++ {
                    if (false && (255 == 1)) {
                        delete(v1, (("foo bar" + v3) + ("a" + v3)))
                        v6 = "foo bar"
                        for fi25 := 0; fi25 < 3; fi25++ {
                            var v26 string = ("foo bar" + v3)
                            switch v26 {
                            case "":
                                var v27 uint32 = 42
                                _ = v27
                            case "foo bar":
                                v2 = v2
                            case "a":
                                var v28 uint8 = 1
                                _ = v28
                            default:
                                var v29 int = 128
                                _ = v29
                            }
                            if false {
                                if ((!true) || false) {
                                    var v30 uint64 = 2
                                    fmt.Println(v6, v2)
                                    _ = v30
                                } else {
                                    v3 = ""
                                    v31 := GenS4{F0: 18446744073709551615, F1: 255}
                                    var v32 int32 = -488
                                    _ = v31
                                    _ = v32
                                }
                                fmt.Println(v6, v2, v6)
                                fmt.Println(v3)
                            }
                            delete(v1, (v6 + (("" + v3) + "a")))
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        v33 := GenS4{F0: 0, F1: 2}
                        _ = v33
                    }
                    var v34 bool = (!(true && true))
                    v35 := GenS3{F0: 716}
                    _ = fi24
                    _ = v34
                    _ = v35
                }
            case "\n\t":
                if true {
                    v36 := []int{42, -292, 1}
                    _ = v36
                }
            case "":
                var v37 int8 = -28
                _ = v37
            default:
                v38 := GenS4{F0: 2, F1: 2}
                _ = v38
            }
        case "a":
            var v39 uint = 139
            _ = v39
        default:
            v40 := GenS1{F0: 0}
            _ = v40
        }
        if (((!false) && true) || (-1 == 1)) {
            v41 := GenS2{F0: 956}
            v42 := []int{2147483647, 128, -1, 492}
            switch (v6 + v6) {
            case "foo bar":
                if true {
                    var v43 int = 4294967295
                    var v44 uint32 = 42
                    if false {
                        v44 = 1
                        if false {
                            switch "a" {
                            case "a":
                                var v45 bool = true
                                _ = v45
                            case "\n\t":
                                v46 := GenS5{F0: 102, F1: 255}
                                _ = v46
                            }
                        }
                        var v47 bool = true
                        _ = v47
                    }
                    _ = v43
                    _ = v44
                }
            }
            _ = v41
            _ = v42
        } else {
            v6 = v6
            v2 = 0.0
        }
        v48 := []int{0, -1}
        _ = v48
    } else {
        v49 := GenS6{F0: 79, F1: 923, F2: 2147483647}
        _ = v49
    }
    if (0 > 467) {
        var v50 uint64 = 284
        _ = v50
    }
    if true {
        switch 9223372036854775807 {
        case 1:
            v51 := GenS7{F0: 0, F1: -1}
            _ = v51
        case -1:
            var v52 string = "hello"
            _ = v52
        default:
            var v53 uint64 = 792
            _ = v53
        }
        v2 += v2
        for fi54 := 0; fi54 < 1; fi54++ {
            for fi55 := 0; fi55 < 8; fi55++ {
                fmt.Println(v6, v2)
                v6 = v6
                _ = fi55
            }
            _ = fi54
        }
    } else {
        v1[((v3 + ((v3 + (((v3 + (((v6 + "hello") + "") + "hello")) + "a") + (v6 + ("foo bar" + ("\n\t" + "hello"))))) + "foo bar")) + "foo bar")] = false
    }
    fmt.Println(v6, v2)
    fmt.Println(v3, v2, v3)
    for fi56 := 0; fi56 < 3; fi56++ {
        v57 := GenS3{F0: -32768}
        _ = fi56
        _ = v57
    }
    v58 := GenS8{F0: 42, F1: 0}
    v6 = "\n\t"
    if (127 <= 255) {
        v59 := []int{}
        switch -123 {
        case -35:
            if (!(!(812 == 870))) {
                _ = len(v1)
                switch (("a" + "hello") + v6) {
                case "foo bar":
                    if true {
                        for fi60 := 0; fi60 < 8; fi60++ {
                            v3 = "\n\t"
                            _ = fi60
                        }
                        if true {
                            switch 1 {
                            case 0:
                                for fi61 := 0; fi61 < 2; fi61++ {
                                    for fi62 := 0; fi62 < 18; fi62++ {
                                        v63 := GenS2{F0: fi62}
                                        v59 = append(v59, (^(fi62)))
                                        fmt.Println(v2)
                                        _ = fi62
                                        _ = v63
                                    }
                                    var v64 uint = 65535
                                    _ = fi61
                                    _ = v64
                                }
                            case 9223372036854775807:
                                for fi65 := 0; fi65 < 6; fi65++ {
                                    if true {
                                        switch 4294967295 {
                                        case -2147483648:
                                            v66 := GenS9{F0: -9223372036854775808}
                                            _ = v66
                                        case 42:
                                            var v67 int64 = 1
                                            _ = v67
                                        case 127:
                                            var v68 uint32 = 0
                                            _ = v68
                                        }
                                        v3 = "hello"
                                        v69 := GenS6{F0: 1, F1: 65535, F2: 169}
                                        _ = v69
                                    }
                                    if false {
                                        v2 = v2
                                        fmt.Println(v3, v2)
                                        v3 = "\n\t"
                                    }
                                    _ = fi65
                                }
                            case -2147483648:
                                fmt.Println(v2)
                            default:
                                _ = len(v59)
                            }
                            v2 = (((1e10 - v2) - v2) / float32(-2.25))
                            var v70 float64 = 0.0
                            _ = v70
                        } else {
                            v6 = ""
                        }
                        fmt.Println(v3, v3, v3)
                    }
                case "hello":
                    if true {
                        for fi71 := 0; fi71 < 10; fi71++ {
                            if false {
                                var v72 int32 = -1
                                _ = v72
                            } else {
                                v2 = v2
                                if true {
                                    v6 = v6
                                } else {
                                    v6 = v6
                                    for fi73 := 0; fi73 < 2; fi73++ {
                                        fmt.Println(fi73)
                                        _ = fi73
                                    }
                                }
                            }
                            _ = fi71
                        }
                    } else {
                        var v74 int32 = 1
                        v74 += (v74 * 357)
                        _ = v74
                    }
                }
                for fi75 := 0; fi75 < 4; fi75++ {
                    v2 += v2
                    for fi76 := 0; fi76 < 15; fi76++ {
                        fmt.Println(fi76)
                        _ = fi76
                    }
                    v59 = append(v59, (fi75 << ((0) & 63)))
                    _ = fi75
                }
            } else {
                for fi77 := 0; fi77 < 16; fi77++ {
                    switch "\n\t" {
                    case "a":
                        for fi78 := 0; fi78 < 2; fi78++ {
                            if true {
                                v2 = 1e10
                                v1[(v3 + v3)] = true
                            } else {
                                if len(v59) > 0 {
                                    v59[0] = fi78
                                }
                            }
                            _ = fi78
                        }
                    case "\n\t":
                        switch v6 {
                        case "foo bar":
                            v79 := GenS8{F0: 2147483647, F1: 42}
                            _ = v79
                        case "a":
                            for fi80 := 0; fi80 < 7; fi80++ {
                                v3 = "a"
                                _ = fi80
                            }
                        }
                    default:
                        v1[((("a" + v6) + ((v3 + ("\n\t" + v6)) + v6)) + "hello")] = (672 < 263)
                    }
                    v3 = v3
                    v81 := GenS3{F0: 236}
                    _ = fi77
                    _ = v81
                }
                if true {
                    switch 1 {
                    case 32767:
                        v82 := v1[(v6 + ("\n\t" + ""))]
                        _ = v82
                    default:
                        if (true || (!(!(false || (((("foo bar" + v3) + (v3 + ("" + (((v3 + v3) + v6) + "foo bar")))) + v3) != "\n\t"))))) {
                            for fi83 := 0; fi83 < 10; fi83++ {
                                fmt.Println(v6)
                                v84 := GenS5{F0: 0, F1: 42}
                                fmt.Println(v6, v3, fi83)
                                _ = fi83
                                _ = v84
                            }
                        }
                    }
                    v85 := GenS10{F0: 349}
                    delete(v1, "a")
                    _ = v85
                }
            }
        case 42:
            var v86 int = 32767
            _ = v86
        case 2147483647:
            var v87 float64 = 1.5
            _ = v87
        default:
            fmt.Println(v6, v3, v6)
        }
        _ = v59
    }
    v3 = ("hello" + v3)
    v88 := GenS9{F0: 0}
    for fi89 := 0; fi89 < 11; fi89++ {
        if (32767 == 1000) {
            var v90 uint32 = 4294967295
            _ = v90
        }
        v6 = "\n\t"
        v3 = v6
        _ = fi89
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v6
    _ = v88
}
