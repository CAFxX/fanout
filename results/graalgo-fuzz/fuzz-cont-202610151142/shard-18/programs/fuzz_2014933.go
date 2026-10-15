package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int32; F2 int64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int16 }

func main() {
    var v1 int16 = 500
    var v2 uint16 = 32767
    var v3 uint16 = v2
    var v4 int16 = (v1 & -98)
    var v5 int = 335
    v6 := make(map[string]string)
    v1 += (v1 >> ((42) & 63))
    switch ("foo bar" + "a") {
    case "foo bar":
        switch 218 {
        case -9223372036854775808:
            if (255 <= 42) {
                v6["\n\t"] = "hello"
                if true {
                    switch 2147483647 {
                    case 279:
                        fmt.Println(v1, v1, v3)
                    case 101:
                        if true {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v2, v3)
                            case "hello":
                                if (!(127 <= 32767)) {
                                    for fi7 := 0; fi7 < 4; fi7++ {
                                        for fi8 := 0; fi8 < 12; fi8++ {
                                            if (-2147483648 > 0) {
                                                var v9 int64 = 685
                                                var v10 int16 = v1
                                                _ = v10
                                                _ = v9
                                            }
                                            _ = fi8
                                        }
                                        v11 := GenS1{F0: 0, F1: 21, F2: ("" + "")}
                                        v5 += (v5 & v5)
                                        _ = fi7
                                        _ = v11
                                    }
                                    switch (v5 | 0) {
                                    case 158:
                                        switch (("hello" + ("foo bar" + "")) + "hello") {
                                        case "hello":
                                            var v12 uint8 = 42
                                            _ = v12
                                        case "a":
                                            var v13 int32 = -1
                                            _ = v13
                                        case "":
                                            v4 -= v1
                                        default:
                                            v14 := GenS1{F0: 497, F1: 0, F2: "foo bar"}
                                            _ = v14
                                        }
                                    case 1:
                                        var v15 uint64 = 123
                                        _ = v15
                                    case 497:
                                        fmt.Println(v4, v1)
                                    }
                                    v16 := GenS2{F0: 9223372036854775807}
                                    _ = v16
                                } else {
                                    fmt.Println(v2, v3)
                                    v17 := []int{v5, 9223372036854775807, int(v1), (-(v5))}
                                    if false {
                                        for fi18 := 0; fi18 < 15; fi18++ {
                                            v4 = v1
                                            v4 = v4
                                            for fi19 := 0; fi19 < 7; fi19++ {
                                                var v20 int64 = 42
                                                switch fi18 {
                                                case -159:
                                                    switch (v5 % ((32767) | 1)) {
                                                    case 0:
                                                        v17 = append(v17, (fi19 / (((-(fi18))) | 1)))
                                                    case 30:
                                                        var v21 uint8 = 0
                                                        v20 = v20
                                                        v6["a"] = "a"
                                                        switch ("hello" + "hello") {
                                                        case "foo bar":
                                                            var v22 bool = (!false)
                                                            _ = v22
                                                        case "a":
                                                            var v23 uint = 1
                                                            _ = v23
                                                        default:
                                                            var v24 float32 = float32(1e10)
                                                            if false {
                                                                v4 = (-(v1))
                                                            } else {
                                                                var v25 uint64 = 4294967295
                                                                v24 = ((float32((0.1 + -2.25)) + v24) - (float32((0.1 - 1.5)) - 3.14159))
                                                                _ = v25
                                                            }
                                                            var v26 int64 = v20
                                                            if true {
                                                                v6[""] = "\n\t"
                                                                fmt.Println(v21, v3, v2)
                                                                var v27 float32 = 0.0
                                                                _ = v27
                                                            }
                                                            var v28 int = 42
                                                            _ = v24
                                                            _ = v26
                                                            _ = v28
                                                        }
                                                        _ = v21
                                                    case 127:
                                                        var v29 uint32 = 2
                                                        _ = v29
                                                    }
                                                case 32767:
                                                    var v30 int16 = 0
                                                    _ = v30
                                                case -9223372036854775808:
                                                    v31 := GenS2{F0: (v20 >> ((uint(v3)) & 63))}
                                                    _ = v31
                                                }
                                                var v32 int64 = v20
                                                _ = fi19
                                                _ = v20
                                                _ = v32
                                            }
                                            _ = fi18
                                        }
                                        delete(v6, "a")
                                    } else {
                                        if (4294967295 == 0) {
                                            var v33 uint64 = 930
                                            _ = v33
                                        } else {
                                            fmt.Println(v1, v1)
                                        }
                                        var v34 int = (v5 - (v5 + v5))
                                        _ = v34
                                    }
                                    _ = v17
                                }
                            default:
                                v35 := GenS2{F0: 9223372036854775807}
                                _ = v35
                            }
                            v36 := []int{128, v5, v5, v5}
                            _ = v36
                        }
                    case -9223372036854775808:
                        delete(v6, ("hello" + ""))
                    }
                    var v37 bool = (((false && false) || (573 > 2)) && true)
                    v38 := GenS2{F0: -212}
                    _ = v37
                    _ = v38
                }
                v39 := GenS2{F0: -453}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 19; fi40++ {
                    switch -960 {
                    case 465:
                        for fi41 := 0; fi41 < 20; fi41++ {
                            v42 := GenS3{F0: 0}
                            _ = fi41
                            _ = v42
                        }
                    case 4294967295:
                        for fi43 := 0; fi43 < 11; fi43++ {
                            var v44 int16 = -753
                            _ = fi43
                            _ = v44
                        }
                    }
                    _ = fi40
                }
                for fi45 := 0; fi45 < 5; fi45++ {
                    var v46 string = "\n\t"
                    for fi47 := 0; fi47 < 4; fi47++ {
                        var v48 string = (v46 + v46)
                        v46 = v48
                        _ = fi47
                        _ = v48
                    }
                    if ((-1 < 42) || ((1 > 4294967295) && false)) {
                        switch (v5 / ((fi45) | 1)) {
                        case 458:
                            var v49 int32 = 803
                            _ = v49
                        case 712:
                            var v50 int = (fi45 * fi45)
                            _ = v50
                        case -2147483648:
                            var v51 uint = 255
                            _ = v51
                        default:
                            v4 = 0
                        }
                        v52 := GenS4{F0: 2, F1: int32(v1), F2: 723}
                        _ = v52
                    } else {
                        switch v5 {
                        case 9223372036854775807:
                            var v53 uint8 = 255
                            _ = v53
                        case -776:
                            switch fi45 {
                            case 513:
                                var v54 int32 = -269
                                _ = v54
                            case 4294967295:
                                if (2147483647 >= 913) {
                                    v55 := GenS5{F0: -625, F1: 18446744073709551615}
                                    v5 = fi45
                                    var v56 string = ""
                                    _ = v55
                                    _ = v56
                                } else {
                                    for fi57 := 0; fi57 < 7; fi57++ {
                                        v3 += v3
                                        fmt.Println(fi57, v5)
                                        _ = fi57
                                    }
                                    for fi58 := 0; fi58 < 15; fi58++ {
                                        v6["\n\t"] = ("hello" + v46)
                                        if ((253 <= 255) || true) {
                                            var v59 int = fi45
                                            var v60 uint64 = 18446744073709551615
                                            _ = v59
                                            _ = v60
                                        }
                                        fmt.Println(v5, v46, fi58)
                                        _ = fi58
                                    }
                                }
                            case -560:
                                v61 := []int{(fi45 - -2147483648), (v5 ^ (v5 * fi45))}
                                _ = v61
                            }
                        default:
                            delete(v6, v46)
                        }
                        switch -430 {
                        case -402:
                            if (true && false) {
                                for fi62 := 0; fi62 < 19; fi62++ {
                                    var v63 int32 = 1
                                    _ = fi62
                                    _ = v63
                                }
                            }
                        case 32767:
                            v2 = v3
                        case 9223372036854775807:
                            var v64 int64 = -38
                            _ = v64
                        default:
                            for fi65 := 0; fi65 < 3; fi65++ {
                                v4 = v1
                                _ = fi65
                            }
                        }
                        v46 = v46
                    }
                    _ = fi45
                    _ = v46
                }
                var v66 uint8 = 128
                _ = v66
            }
        case 32768:
            if true {
                delete(v6, "\n\t")
            } else {
                v67 := GenS6{F0: -1, F1: 65535}
                var v68 uint64 = 18446744073709551615
                var v69 int32 = -535
                _ = v67
                _ = v68
                _ = v69
            }
        }
    case "hello":
        var v70 uint16 = v2
        _ = v70
    default:
        v4 = 1000
    }
    var v71 uint64 = 234
    var v72 bool = (196 == (^(v71)))
    v6["\n\t"] = (("a" + (("" + (("a" + ("" + "foo bar")) + ("\n\t" + "\n\t"))) + "a")) + "hello")
    v3 += v2
    fmt.Println(v72, v3)
    v1 = v4
    var v73 uint64 = v71
    fmt.Println(v5, v2)
    var v74 int16 = (v4 & -1)
    switch "a" {
    case "a":
        for fi75 := 0; fi75 < 9; fi75++ {
            for fi76 := 0; fi76 < 5; fi76++ {
                v77 := GenS1{F0: -1, F1: 95, F2: ("foo bar" + "hello")}
                v71 = v71
                v78 := GenS2{F0: 42}
                _ = fi76
                _ = v77
                _ = v78
            }
            _ = fi75
        }
    case "foo bar":
        switch (v5 & v5) {
        case 32767:
            switch "a" {
            case "a":
                for fi79 := 0; fi79 < 15; fi79++ {
                    fmt.Println(v2, v5)
                    v80 := GenS2{F0: 42}
                    _ = fi79
                    _ = v80
                }
            case "":
                if v72 {
                    for fi81 := 0; fi81 < 2; fi81++ {
                        if ((4294967295 != v73) && (!true)) {
                            v6[("\n\t" + ("hello" + ""))] = ("hello" + "\n\t")
                        } else {
                            var v82 uint64 = (v71 + (v73 / (((^(v71))) | 1)))
                            _ = v82
                        }
                        var v83 int32 = 42
                        _ = fi81
                        _ = v83
                    }
                    for fi84 := 0; fi84 < 9; fi84++ {
                        var v85 uint32 = 378
                        for fi86 := 0; fi86 < 9; fi86++ {
                            var v87 int64 = 0
                            v5 -= (fi86 ^ fi86)
                            var v88 int16 = v74
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                        _ = fi84
                        _ = v85
                    }
                }
            case "foo bar":
                for fi89 := 0; fi89 < 14; fi89++ {
                    var v90 int64 = 0
                    var v91 int64 = 0
                    v92 := GenS7{F0: 0}
                    _ = fi89
                    _ = v90
                    _ = v91
                    _ = v92
                }
            default:
                v71 += (v71 - 2)
            }
        case 32768:
            if v72 {
                if (uint(v2) >= 797) {
                    v4 -= (v74 * (v74 & (v1 | (-(v1)))))
                    v93 := GenS5{F0: -1, F1: 432}
                    _ = v93
                } else {
                    v94 := GenS8{F0: true}
                    if v72 {
                        v1 = (v74 | (^(v1)))
                        v4 = v74
                    }
                    var v95 int32 = int32(v5)
                    _ = v94
                    _ = v95
                }
                switch "" {
                case "":
                    for fi96 := 0; fi96 < 8; fi96++ {
                        var v97 int = fi96
                        _ = fi96
                        _ = v97
                    }
                default:
                    var v98 int16 = v74
                    _ = v98
                }
                var v99 string = ("a" + "hello")
                _ = v99
            }
        case 4294967295:
            for fi100 := 0; fi100 < 7; fi100++ {
                delete(v6, "hello")
                var v101 uint16 = (v2 | v3)
                _ = fi100
                _ = v101
            }
        default:
            if (true && (uint(v3) > 92)) {
                var v102 uint32 = 870
                v103 := GenS5{F0: int64(v5), F1: 1}
                _ = v102
                _ = v103
            } else {
                fmt.Println(v3, v2, v1)
            }
        }
    }
    var v104 string = "foo bar"
    var v105 uint8 = 2
    if v72 {
        switch "foo bar" {
        case "\n\t":
            var v106 string = "foo bar"
            _ = v106
        case "a":
            for fi107 := 0; fi107 < 4; fi107++ {
                v5 -= (fi107 / ((v5) | 1))
                v104 = ((((v104 + "foo bar") + v104) + v104) + v104)
                v108 := v6[(v104 + v104)]
                _ = fi107
                _ = v108
            }
        case "foo bar":
            if (2147483647 != -2147483648) {
                if (v72 || v72) {
                    var v109 string = v104
                    v110 := GenS9{F0: v74, F1: 65535}
                    _ = v109
                    _ = v110
                }
            }
        }
        var v111 string = ((v104 + (v104 + "hello")) + v104)
        _ = v111
    }
    switch (v5 % ((int(v2)) | 1)) {
    case 42:
        fmt.Println(v73, v1)
    case 621:
        fmt.Println(v2, v104)
    }
    v112 := GenS10{F0: (v74 | (-(v74)))}
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v104
    _ = v105
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
