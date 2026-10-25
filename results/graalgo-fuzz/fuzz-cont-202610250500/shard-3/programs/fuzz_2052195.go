package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string; F1 bool; F2 bool }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"", "a", "\n\t"}
    var v3 int64 = 42
    var v4 int8 = 42
    fmt.Println(v3, v4)
    switch -2147483648 {
    case 98:
        switch int(v3) {
        case 42:
            switch 32768 {
            case 332:
                var v5 uint8 = 2
                _ = v5
            case 128:
                for fi6 := 0; fi6 < 3; fi6++ {
                    switch fi6 {
                    case 32768:
                        if false {
                            fmt.Println(fi6)
                            if (!(!false)) {
                                var v7 int64 = v3
                                v4 = (^(v4))
                                var v8 uint8 = 127
                                _ = v7
                                _ = v8
                            }
                        } else {
                            if false {
                                var v9 string = "foo bar"
                                if ((v3 << ((756) & 63)) >= v3) {
                                    v1 -= (v1 + v1)
                                    var v10 bool = true
                                    _ = v10
                                } else {
                                    if (!(!true)) {
                                        if ((!true) || false) {
                                            var v11 uint32 = 530
                                            _ = v11
                                        }
                                    }
                                    _ = len(v2)
                                    var v12 float32 = (0.1 - ((3.14159 - (0.1 / (0.1 + 0.1))) * 0.1))
                                    _ = v12
                                }
                                _ = v9
                            } else {
                                v13 := GenS1{F0: (v1 | (v1 | (v1 << ((4294967295) & 63))))}
                                switch 1 {
                                case -9223372036854775808:
                                    for fi14 := 0; fi14 < 20; fi14++ {
                                        var v15 float64 = 1.5
                                        if (0 > 989) {
                                            v3 -= v3
                                        } else {
                                            if (127 <= (fi6 >> ((2) & 63))) {
                                                var v16 uint8 = 127
                                                v17 := GenS1{F0: (v1 - 12)}
                                                _ = v16
                                                _ = v17
                                            }
                                            var v18 bool = (-464 >= 32767)
                                            if (!v18) {
                                                var v19 uint16 = 32767
                                                if (v3 != v3) {
                                                    v20 := GenS1{F0: v1}
                                                    _ = v20
                                                } else {
                                                    if (0 <= 0) {
                                                        v1 += v1
                                                        if true {
                                                            var v21 uint8 = 161
                                                            switch "\n\t" {
                                                            case "a":
                                                                _ = len(v2)
                                                            case "":
                                                                fmt.Println(v4)
                                                            case "\n\t":
                                                                v15 -= 0.0
                                                            default:
                                                                v3 = 9223372036854775807
                                                                v22 := GenS2{F0: true, F1: v4}
                                                                _ = v22
                                                            }
                                                            var v23 uint = 0
                                                            _ = v21
                                                            _ = v23
                                                        }
                                                        var v24 int16 = 42
                                                        _ = v24
                                                    }
                                                    v18 = v18
                                                    var v25 int64 = (-(v3))
                                                    switch fi6 {
                                                    case -2147483648:
                                                        if false {
                                                            switch 0 {
                                                            case 1:
                                                                var v26 uint32 = 65535
                                                                _ = v26
                                                            case 9223372036854775807:
                                                                var v27 uint = 4294967295
                                                                _ = v27
                                                            case 42:
                                                                v28 := GenS3{F0: ("" + "foo bar"), F1: true, F2: (0 < 1)}
                                                                _ = v28
                                                            }
                                                            if v18 {
                                                                v29 := GenS2{F0: (false && v18), F1: v4}
                                                                if v18 {
                                                                    fmt.Println(fi6, v18, v25)
                                                                }
                                                                if ((v1 != (^(v1))) || ((255 < 423) || v18)) {
                                                                    v19 = (v19 | (v19 << ((4294967295) & 63)))
                                                                    v19 = 65535
                                                                    fmt.Println(v25, fi14)
                                                                }
                                                                _ = v29
                                                            } else {
                                                                var v30 int64 = (v25 % ((-1) | 1))
                                                                _ = v30
                                                            }
                                                        }
                                                    }
                                                    _ = v25
                                                }
                                                if ((v4 << ((65535) & 63)) > v4) {
                                                    fmt.Println(v1, v3, v18)
                                                } else {
                                                    if v18 {
                                                        fmt.Println(fi14)
                                                        var v31 string = (("a" + "\n\t") + "")
                                                        switch fi14 {
                                                        case 1:
                                                            if len(v2) > 0 {
                                                                v2[0] = "hello"
                                                            }
                                                        case 9223372036854775807:
                                                            var v32 bool = false
                                                            _ = v32
                                                        case -481:
                                                            if true {
                                                                v4 = -110
                                                                if ((v18 || ((128 >= 0) && (!v18))) || true) {
                                                                    if (!false) {
                                                                        v18 = (((v4 / (((^(v4))) | 1)) != (v4 | (v4 & 42))) && v18)
                                                                        v15 = v15
                                                                        var v33 float32 = (float32((-2.25 / (v15 - 3.14159))) / 1e10)
                                                                        _ = v33
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                    v34 := GenS4{F0: 4294967295, F1: v4}
                                                                    _ = v34
                                                                }
                                                            }
                                                        }
                                                        switch v31 {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = v31
                                                            }
                                                        case "foo bar":
                                                            v2 = append(v2, v31)
                                                        case "a":
                                                            var v35 float64 = (v15 * 0.0)
                                                            _ = v35
                                                        default:
                                                            v4 = v4
                                                        }
                                                        _ = v31
                                                    }
                                                }
                                                _ = v19
                                            }
                                            _ = v18
                                        }
                                        v36 := GenS5{F0: "a"}
                                        var v37 int = fi6
                                        v3 = (v3 % (((v3 ^ v3)) | 1))
                                        var v38 uint64 = v1
                                        if true {
                                            v37 = (-(fi14))
                                            switch (("\n\t" + ("" + "\n\t")) + "foo bar") {
                                            case "foo bar":
                                                v39 := GenS2{F0: false, F1: 25}
                                                _ = v39
                                            case "a":
                                                if true {
                                                    v2 = append(v2, ("" + "\n\t"))
                                                    var v40 uint32 = 274
                                                    var v41 uint64 = (v1 + (v1 % ((v1) | 1)))
                                                    _ = v40
                                                    _ = v41
                                                }
                                            default:
                                                v1 = (^(v38))
                                            }
                                            v1 = v1
                                        } else {
                                            var v42 float64 = (v15 / (-2.25 * float64(3.14159)))
                                            if (true && (534 > 4294967295)) {
                                                if false {
                                                    v43 := GenS1{F0: v1}
                                                    var v44 uint8 = 28
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    var v45 bool = false
                                                    v46 := GenS4{F0: 716, F1: (v4 ^ v4)}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 uint8 = 140
                                                fmt.Println(v15, v38)
                                                _ = v47
                                            } else {
                                                if (2 > 1) {
                                                    var v48 uint8 = 2
                                                    v38 = (^(v1))
                                                    v15 += 0.1
                                                    _ = v48
                                                }
                                            }
                                            var v49 uint16 = 32767
                                            _ = v42
                                            _ = v49
                                        }
                                        _ = fi14
                                        _ = v15
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                fmt.Println(fi6, v4, v1)
                                _ = v13
                            }
                            v4 -= v4
                            var v50 int8 = 40
                            _ = v50
                        }
                    case 1:
                        var v51 bool = false
                        _ = v51
                    default:
                        fmt.Println(v3)
                    }
                    v3 += -9223372036854775808
                    fmt.Println(v1, v3)
                    _ = fi6
                }
            default:
                v3 = v3
            }
        }
    case 42:
        var v52 uint8 = 44
        _ = v52
    case 38:
        var v53 string = ""
        _ = v53
    default:
        var v54 int64 = (v3 * v3)
        _ = v54
    }
    fmt.Println(v3, v4, v4)
    fmt.Println(v1, v3, v1)
    if false {
        v2 = append(v2, (("a" + ("" + "foo bar")) + "hello"))
        for fi55 := 0; fi55 < 2; fi55++ {
            for fi56 := 0; fi56 < 7; fi56++ {
                _ = len(v2)
                v3 += (v3 & v3)
                v57 := GenS6{F0: "foo bar"}
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
        v2 = append(v2, "a")
    } else {
        if (true || ((!true) || (!(true || (true && true))))) {
            fmt.Println(v3, v4)
            if false {
                fmt.Println(v1, v3)
                if (true || (!false)) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        v1 += 2
                    case 271:
                        var v58 uint = 18446744073709551615
                        _ = v58
                    case -464:
                        var v59 bool = (!(42 != 127))
                        _ = v59
                    }
                    fmt.Println(v3, v4)
                } else {
                    if (!(false || (v1 >= 42))) {
                        switch 4294967295 {
                        case -9223372036854775808:
                            if (!(true && ("\n\t" <= ""))) {
                                switch "\n\t" {
                                case "":
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        switch ("hello" + "foo bar") {
                                        case "a":
                                            v60 := GenS1{F0: (v1 * v1)}
                                            _ = v60
                                        default:
                                            var v61 uint16 = 42
                                            _ = v61
                                        }
                                    case "a":
                                        for fi62 := 0; fi62 < 7; fi62++ {
                                            v2 = append(v2, ("foo bar" + "hello"))
                                            if (!true) {
                                                var v63 float32 = 3.14159
                                                _ = v63
                                            }
                                            fmt.Println(v4)
                                            _ = fi62
                                        }
                                    }
                                case "foo bar":
                                    var v64 bool = false
                                    _ = v64
                                case "\n\t":
                                    var v65 uint64 = v1
                                    _ = v65
                                default:
                                    var v66 uint64 = 18446744073709551615
                                    _ = v66
                                }
                                var v67 bool = true
                                _ = v67
                            } else {
                                for fi68 := 0; fi68 < 3; fi68++ {
                                    v3 = v3
                                    _ = fi68
                                }
                            }
                        case 0:
                            switch "a" {
                            case "hello":
                                fmt.Println(v4)
                            default:
                                v1 = 0
                            }
                        case -2147483648:
                            if false {
                                v4 += (v4 | v4)
                            }
                        default:
                            for fi69 := 0; fi69 < 20; fi69++ {
                                v3 = v3
                                _ = fi69
                            }
                        }
                        if false {
                            v4 -= (v4 / (((v4 / ((v4) | 1))) | 1))
                            v3 -= (v3 & 253)
                        } else {
                            for fi70 := 0; fi70 < 6; fi70++ {
                                var v71 uint32 = 2147483647
                                _ = fi70
                                _ = v71
                            }
                            switch ("" + ("foo bar" + "hello")) {
                            case "a":
                                var v72 int = 42
                                _ = v72
                            default:
                                for fi73 := 0; fi73 < 3; fi73++ {
                                    _ = len(v2)
                                    fmt.Println(v4, v4)
                                    _ = fi73
                                }
                            }
                        }
                        v74 := GenS5{F0: "foo bar"}
                        _ = v74
                    }
                    fmt.Println(v1)
                    var v75 int8 = v4
                    _ = v75
                }
                v1 = (^(v1))
            }
        } else {
            v4 += 42
            for fi76 := 0; fi76 < 18; fi76++ {
                v1 = (v1 + (v1 % ((v1) | 1)))
                _ = fi76
            }
        }
        if (!true) {
            for fi77 := 0; fi77 < 11; fi77++ {
                var v78 uint = 4294967295
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi77
                _ = v78
            }
            var v79 int8 = 42
            _ = v79
        }
        v80 := GenS6{F0: ("\n\t" + "a")}
        _ = v80
    }
    switch "foo bar" {
    case "":
        switch 147 {
        case 2147483647:
            v1 = v1
        }
    case "\n\t":
        for fi81 := 0; fi81 < 2; fi81++ {
            v4 += (v4 + (v4 / ((v4) | 1)))
            v1 = 732
            _ = fi81
        }
    case "hello":
        v82 := GenS5{F0: "hello"}
        _ = v82
    }
    if false {
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v83 int16 = 0
    var v84 uint8 = 128
    v84 -= (^(v84))
    v84 = 2
    if (true || (true && true)) {
        if false {
            for fi85 := 0; fi85 < 3; fi85++ {
                var v86 bool = ((fi85 ^ (fi85 - 32767)) >= 4294967295)
                _ = fi85
                _ = v86
            }
            v83 = (^(v83))
        }
        var v87 string = ("\n\t" + "hello")
        _ = v87
    }
    v84 -= (v84 + (^(v84)))
    switch "hello" {
    case "":
        v84 += v84
    }
    for fi88 := 0; fi88 < 15; fi88++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
        _ = fi88
    }
    v3 -= (v3 | -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
}
