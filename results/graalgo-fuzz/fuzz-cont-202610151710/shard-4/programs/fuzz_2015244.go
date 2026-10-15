package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 bool }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    var v3 uint16 = 65535
    var v4 int8 = 127
    var v5 int = 9223372036854775807
    fmt.Println(v4)
    v6 := []int{(v5 >> ((33) & 63)), 32768, (v5 * v5)}
    v2 += (v2 >> ((255) & 63))
    v5 = (v5 << ((65535) & 63))
    fmt.Println(v5)
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: -907, F1: 4294967295, F2: (v2 + (^(v2)))}
        var v9 int32 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    switch v5 {
    case 4294967295:
        for fi10 := 0; fi10 < 13; fi10++ {
            if (v5 > 127) {
                var v11 int = 1
                for fi12 := 0; fi12 < 18; fi12++ {
                    switch "\n\t" {
                    case "":
                        var v13 uint16 = v3
                        _ = v13
                    case "hello":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        v14 := GenS1{F0: fi10, F1: 614, F2: 147}
                        _ = v14
                    }
                    _ = fi12
                }
                var v15 string = "a"
                _ = v11
                _ = v15
            }
            switch 0 {
            case 128:
                fmt.Println(v4, v5, fi10)
            case -1:
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: 127, F1: 0, F2: (^(v2))}
                    _ = fi16
                    _ = v17
                }
            case 611:
                switch -742 {
                case -777:
                    v1[507] = true
                case 2147483647:
                    v2 -= v2
                case 32767:
                    if false {
                        switch (fi10 - 0) {
                        case -2147483648:
                            var v18 uint32 = 640
                            _ = v18
                        case 853:
                            if (!(((true && ((v3 % (((^(v3))) | 1)) <= v3)) || false) || (!true))) {
                                v19 := GenS1{F0: (v5 | (v5 & fi10)), F1: 1, F2: v2}
                                switch 4294967295 {
                                case 42:
                                    v6 = append(v6, v5)
                                case 9223372036854775807:
                                    v20 := GenS2{F0: 511, F1: -40, F2: false}
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    v22 := GenS2{F0: 1000, F1: 32767, F2: true}
                                    _ = fi21
                                    _ = v22
                                }
                                _ = v19
                            }
                        }
                    } else {
                        v23 := GenS3{F0: (-(v4)), F1: 128}
                        delete(v1, 544)
                        v1[4294967295] = (false && true)
                        _ = v23
                    }
                }
            }
            _ = fi10
        }
    case 443:
        v3 -= v3
    case 32767:
        v4 += 1
    }
    if (!true) {
        if false {
            if true {
                v2 += v2
                switch (v5 + (v5 + v5)) {
                case 78:
                    var v24 uint64 = 42
                    _ = v24
                case 9223372036854775807:
                    v2 += v2
                }
            } else {
                switch (v5 + v5) {
                case 32768:
                    switch "a" {
                    case "a":
                        v25 := GenS4{F0: -465}
                        _ = v25
                    default:
                        var v26 int8 = (v4 / (((v4 % (((v4 & -128)) | 1))) | 1))
                        _ = v26
                    }
                case -2147483648:
                    v27 := GenS5{F0: 1}
                    _ = v27
                default:
                    v1[2147483647] = true
                }
                v5 += int(v3)
            }
        }
        switch ("" + "") {
        case "foo bar":
            if (128 >= (v2 % ((v2) | 1))) {
                switch "a" {
                case "a":
                    v4 += (v4 - (-(v4)))
                case "foo bar":
                    fmt.Println(v2, v5)
                default:
                    var v28 int32 = 811
                    _ = v28
                }
            }
        case "":
            v2 -= v2
        }
    }
    var v29 uint64 = 18446744073709551615
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v6 = append(v6, fi31)
            if len(v6) > 0 {
                v6[0] = fi30
            }
            _ = fi31
        }
        _ = fi30
    }
    v32 := GenS2{F0: -32768, F1: -932, F2: false}
    v29 += v29
    switch v5 {
    case 774:
        v5 = v5
    case 4294967295:
        v3 += (v3 << ((916) & 63))
    case 0:
        var v33 uint32 = 30
        _ = v33
    }
    if false {
        v34 := GenS1{F0: v5, F1: 18446744073709551615, F2: (^(v2))}
        if false {
            var v35 int64 = -208
            if (!(int64(v5) < v35)) {
                if (1 >= 1) {
                    var v36 float64 = (0.1 - 1.5)
                    v6 = append(v6, (v5 / ((32768) | 1)))
                    _ = v36
                } else {
                    v37 := GenS5{F0: 2147483648}
                    var v38 int64 = v35
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if true {
                            var v40 uint8 = (v2 - 172)
                            for fi41 := 0; fi41 < 3; fi41++ {
                                v42 := GenS6{F0: v29, F1: (((("a" + "hello") + "foo bar") + ((("foo bar" + ("foo bar" + "")) + ((("\n\t" + "hello") + "foo bar") + ("a" + "foo bar"))) + ("\n\t" + (("\n\t" + ("" + "foo bar")) + "hello")))) < "foo bar")}
                                var v43 uint = 65535
                                fmt.Println(v43, v35)
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        }
                        v44 := GenS7{F0: (v3 & v3)}
                        var v45 bool = true
                        _ = fi39
                        _ = v44
                        _ = v45
                    }
                    _ = v37
                    _ = v38
                }
                for fi46 := 0; fi46 < 7; fi46++ {
                    _ = len(v1)
                    v35 = 0
                    switch (v5 >> ((18446744073709551615) & 63)) {
                    case 2147483647:
                        switch "\n\t" {
                        case "hello":
                            for fi47 := 0; fi47 < 2; fi47++ {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(fi46)
                                case "a":
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        if (!(!false)) {
                                            v35 += -1
                                            if false {
                                                v48 := GenS5{F0: 616}
                                                _ = v48
                                            }
                                        } else {
                                            if true {
                                                v49 := GenS7{F0: (^(v3))}
                                                v29 = 738
                                                _ = v49
                                            }
                                            if (-32768 <= -32768) {
                                                _ = len(v6)
                                                v6 = append(v6, fi46)
                                                v50 := GenS8{F0: 2147483647}
                                                _ = v50
                                            } else {
                                                var v51 uint8 = v2
                                                v52 := GenS9{F0: v29, F1: v3}
                                                if (!(v3 > 313)) {
                                                    _ = len(v1)
                                                    if false {
                                                        var v53 uint32 = 985
                                                        var v54 uint16 = (v3 >> ((18446744073709551615) & 63))
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    fmt.Println(v3, fi46)
                                                    if (v4 <= 0) {
                                                        v1[2147483648] = (682 < 4294967295)
                                                        v55 := GenS10{F0: 2147483647, F1: 2}
                                                        v56 := GenS10{F0: 2, F1: 1}
                                                        var v57 int16 = 669
                                                        fmt.Println(fi47)
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    v2 = v51
                                                }
                                                switch -990 {
                                                case 4294967295:
                                                    v58 := GenS2{F0: -156, F1: fi46, F2: true}
                                                    _ = v58
                                                default:
                                                    v6 = append(v6, fi47)
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v59 uint8 = v51
                                                        _ = v59
                                                    case "":
                                                        var v60 int32 = 301
                                                        _ = v60
                                                    default:
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v61 int = fi47
                                                            _ = v61
                                                        case "":
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v62 int8 = 61
                                                                _ = v62
                                                            case "\n\t":
                                                                delete(v1, 260)
                                                            case "a":
                                                                var v63 uint8 = (v51 >> ((65535) & 63))
                                                                _ = v63
                                                            default:
                                                                if len(v6) > 0 {
                                                                    v6[0] = int(v51)
                                                                }
                                                                switch "foo bar" {
                                                                case "a":
                                                                    fmt.Println(v4, fi46)
                                                                case "foo bar":
                                                                    if false {
                                                                        fmt.Println(v5)
                                                                    }
                                                                case "":
                                                                    v64 := GenS10{F0: 2147483647, F1: 65535}
                                                                    _ = v64
                                                                }
                                                            }
                                                        default:
                                                            var v65 uint = 574
                                                            _ = v65
                                                        }
                                                    }
                                                    if (v35 != 42) {
                                                        v5 = (v5 | 9223372036854775807)
                                                    } else {
                                                        if ((^(v35)) > (v35 & (^(v35)))) {
                                                            v66 := GenS11{F0: 127, F1: (v51 * v51), F2: v3}
                                                            v67 := GenS12{F0: 2}
                                                            fmt.Println(v3, v29, v5)
                                                            var v68 float32 = (0.0 + 3.14159)
                                                            var v69 uint = 18446744073709551615
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            if ((-(v35)) == (-(v35))) {
                                                                fmt.Println(v51)
                                                                var v70 int32 = -2147483648
                                                                var v71 float64 = -2.25
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            var v72 uint16 = 32767
                                                            var v73 float32 = 3.14159
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        v5 = (-(fi46))
                                                        v29 -= 18446744073709551615
                                                    }
                                                }
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                    }
                                }
                                var v74 int16 = 32767
                                if true {
                                    v4 = -128
                                    var v75 bool = true
                                    v74 = v74
                                    _ = v75
                                }
                                _ = fi47
                                _ = v74
                            }
                        case "\n\t":
                            v2 -= v2
                        }
                    case 128:
                        if (3.14159 >= 3.14159) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            var v76 bool = true
                            _ = v76
                        }
                    case -1:
                        v2 += 127
                    }
                    _ = fi46
                }
                v77 := GenS5{F0: 42}
                _ = v77
            } else {
                _ = len(v6)
                v35 = v35
            }
            _ = v35
        }
        switch ("foo bar" + "foo bar") {
        case "":
            v78 := GenS13{F0: (("\n\t" + ("foo bar" + "foo bar")) + ("\n\t" + "foo bar"))}
            _ = v78
        case "foo bar":
            v4 = 100
        case "\n\t":
            for fi79 := 0; fi79 < 20; fi79++ {
                var v80 uint64 = 169
                for fi81 := 0; fi81 < 11; fi81++ {
                    fmt.Println(v2, fi81)
                    if true {
                        v82 := GenS11{F0: (v4 >> ((65535) & 63)), F1: 35, F2: (^(v3))}
                        _ = v82
                    }
                    _ = len(v6)
                    _ = fi81
                }
                _ = fi79
                _ = v80
            }
        }
        _ = v34
    }
    var v83 uint32 = 339
    fmt.Println(v83, v5)
    v5 += (v5 >> ((18446744073709551615) & 63))
    var v84 uint64 = (v29 >> ((888) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
}
