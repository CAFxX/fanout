package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 uint8; F2 uint8 }

func main() {
    var v1 uint8 = 91
    var v2 float64 = (-2.25 + (3.14159 / 1.5))
    var v3 float64 = ((v2 - v2) + v2)
    if false {
        v4 := []int{}
        _ = v4
    } else {
        var v5 int16 = 652
        v6 := []int{42}
        v7 := GenS1{F0: 255, F1: 1, F2: 4294967295}
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v1 = (v1 + (^(v1)))
        v9 := []int{fi8, fi8, fi8}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = (v1 - (v1 << ((4294967295) & 63)))
    fmt.Println(v2, v10, v10)
    switch ("\n\t" + "hello") {
    case "foo bar":
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                switch "foo bar" {
                case "":
                    var v12 uint32 = 398
                    _ = v12
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    v3 += ((float64(-2.25) / v3) / v2)
                    if (-2147483648 <= 0) {
                        var v14 int = (^(fi11))
                        v1 += (v10 << ((48) & 63))
                        v1 = (v10 << ((uint(v1)) & 63))
                        _ = v14
                    }
                    var v15 uint32 = 42
                    _ = fi13
                    _ = v15
                }
                var v16 uint8 = (v1 | (v1 * (v1 + (^(v10)))))
                _ = fi11
                _ = v16
            }
        }
    case "a":
        v17 := GenS1{F0: 623, F1: 127, F2: 0}
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 8; fi18++ {
            v10 = v1
            switch "" {
            case "\n\t":
                fmt.Println(v10, v2)
            case "":
                v1 = v1
            case "hello":
                v19 := GenS1{F0: 290, F1: 127, F2: 0}
                _ = v19
            }
            _ = fi18
        }
    default:
        switch ("a" + "") {
        case "":
            var v20 uint32 = 65535
            _ = v20
        case "hello":
            var v21 int = -2147483648
            _ = v21
        default:
            v22 := GenS1{F0: 2, F1: int8(v1), F2: 1}
            _ = v22
        }
    }
    fmt.Println(v2, v1, v3)
    v23 := GenS2{F0: "a", F1: -9223372036854775808}
    v24 := GenS1{F0: 32767, F1: 97, F2: 18446744073709551615}
    if false {
        if (v10 <= 128) {
            for fi25 := 0; fi25 < 12; fi25++ {
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(v10)
                    v27 := GenS3{F0: ("foo bar" + "a")}
                    if (1 == 0) {
                        v2 += v3
                    } else {
                        v28 := GenS1{F0: 1, F1: -128, F2: 2}
                        v29 := GenS3{F0: ""}
                        _ = v28
                        _ = v29
                    }
                    if (42 < 127) {
                        var v30 string = ((("hello" + "\n\t") + "foo bar") + "")
                        switch 4294967295 {
                        case 4294967295:
                            v1 += (v1 | (v1 >> ((693) & 63)))
                        default:
                            var v31 int8 = 42
                            _ = v31
                        }
                        _ = v30
                    } else {
                        switch "hello" {
                        case "hello":
                            var v32 int16 = -32768
                            _ = v32
                        case "\n\t":
                            var v33 float64 = v3
                            _ = v33
                        case "":
                            v34 := GenS1{F0: 32767, F1: -128, F2: 511}
                            _ = v34
                        }
                        if false {
                            if false {
                                if false {
                                    v10 -= v10
                                    var v35 uint64 = 755
                                    v2 -= v3
                                    _ = v35
                                }
                                v3 = 3.14159
                                var v36 uint64 = 42
                                v10 = v10
                                v37 := GenS2{F0: "\n\t", F1: 1}
                                if (!(!(!(false || true)))) {
                                    v38 := []int{(fi25 % ((fi26) | 1)), -771}
                                    var v39 uint8 = (^(v10))
                                    var v40 uint16 = 461
                                    v41 := GenS1{F0: 2, F1: 14, F2: (v36 | 42)}
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                                _ = v36
                                _ = v37
                            } else {
                                var v42 string = (("hello" + "a") + "\n\t")
                                v1 = v10
                                _ = v42
                            }
                            v10 -= (v10 & (v1 >> ((18446744073709551615) & 63)))
                            v43 := GenS3{F0: ""}
                            v3 = (v2 / float64(1.5))
                            v2 -= v2
                            _ = v43
                        } else {
                            v44 := GenS4{F0: "a", F1: 2147483647, F2: -290}
                            if false {
                                if false {
                                    fmt.Println(v10, v1)
                                    v45 := []int{}
                                    _ = v45
                                }
                                var v46 int = (fi26 | (-(fi25)))
                                _ = v46
                            } else {
                                fmt.Println(v10, fi26)
                                v47 := []int{(-(fi26))}
                                if false {
                                    if (fi26 < (fi26 * 0)) {
                                        var v48 uint8 = v1
                                        var v49 int16 = 0
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v3 = v2
                                    }
                                    if len(v47) > 0 {
                                        v47[0] = fi26
                                    }
                                    v47 = append(v47, -2147483648)
                                    v50 := GenS5{F0: 42, F1: true}
                                    var v51 int16 = 32767
                                    if (((1 < -377) && (("a" != (("foo bar" + ("foo bar" + "")) + ("" + "\n\t"))) || (!((v51 > (v51 >> ((408) & 63))) || true)))) && true) {
                                        v47 = append(v47, 417)
                                    } else {
                                        if true {
                                            switch (fi26 / ((fi25) | 1)) {
                                            case -653:
                                                var v52 uint32 = 1
                                                _ = v52
                                            case 32768:
                                                if true {
                                                    if len(v47) >= 2 {
                                                        v47 = v47[1:len(v47)]
                                                    }
                                                    v2 -= (v3 / (0.1 + 0.0))
                                                    var v53 int32 = 42
                                                    _ = v53
                                                } else {
                                                    var v54 uint32 = 42
                                                    v2 -= (float64((1.5 - ((((((0.1 - (float32(-2.25) - 3.14159)) / (0.0 - (-2.25 * (0.1 / (0.1 / -2.25))))) + -2.25) * 1.5) * ((-2.25 * -2.25) + (1.5 - ((0.1 - float32(float64(1e10))) + -2.25)))) + float32((v3 * v3))))) - ((((0.0 + (((3.14159 * (0.0 + 3.14159)) + 0.0) - float64(0.0))) + float64((0.0 * 0.1))) * 1.5) - 0.0))
                                                    switch 0 {
                                                    case -352:
                                                        v55 := GenS6{F0: (v54 + 198), F1: (fi26 - (fi26 / (((-(fi25))) | 1)))}
                                                        _ = v55
                                                    }
                                                    _ = v54
                                                }
                                            case 2147483647:
                                                v56 := GenS3{F0: ("hello" + "a")}
                                                _ = v56
                                            }
                                            if (!(("hello" + "hello") <= ("a" + "foo bar"))) {
                                                v47 = append(v47, 42)
                                                if false {
                                                    _ = len(v47)
                                                    var v57 int64 = -9223372036854775808
                                                    v57 -= v57
                                                    _ = v57
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v58 int32 = 1
                                                    _ = v58
                                                }
                                            } else {
                                                var v59 uint64 = 358
                                                _ = v59
                                            }
                                        } else {
                                            if ((v51 <= (v51 % (((v51 + (v51 / ((v51) | 1)))) | 1))) && (1 <= 1)) {
                                                v47 = append(v47, fi25)
                                            }
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                } else {
                                    v60 := GenS3{F0: "\n\t"}
                                    _ = v60
                                }
                                _ = v47
                            }
                            _ = v44
                        }
                        v61 := []int{}
                        var v62 uint16 = 42
                        if (!false) {
                            v63 := GenS6{F0: 2147483647, F1: 32767}
                            _ = len(v61)
                            _ = v63
                        }
                        v3 -= v2
                        _ = v61
                        _ = v62
                    }
                    var v64 uint64 = 18446744073709551615
                    var v65 uint32 = uint32(v64)
                    v66 := GenS2{F0: "a", F1: -790}
                    _ = fi26
                    _ = v27
                    _ = v64
                    _ = v65
                    _ = v66
                }
                _ = fi25
            }
            if (42 <= 322) {
                v3 = v3
                v2 = 0.1
            }
            v67 := []int{127}
            _ = v67
        } else {
            if true {
                if true {
                    var v68 float32 = float32(1.5)
                    v68 -= v68
                    _ = v68
                } else {
                    v3 = (v2 * v2)
                    v69 := GenS6{F0: 2, F1: 1}
                    _ = v69
                }
                for fi70 := 0; fi70 < 14; fi70++ {
                    fmt.Println(v3, fi70, v2)
                    _ = fi70
                }
                var v71 int64 = -1
                _ = v71
            }
            switch 0 {
            case 32767:
                for fi72 := 0; fi72 < 14; fi72++ {
                    var v73 float64 = float64(0.0)
                    var v74 uint64 = uint64(v1)
                    v75 := GenS6{F0: uint32(v74), F1: fi72}
                    _ = fi72
                    _ = v73
                    _ = v74
                    _ = v75
                }
            case -2147483648:
                for fi76 := 0; fi76 < 4; fi76++ {
                    var v77 uint8 = (^(v10))
                    if ((v10 | 127) == (v10 & (v10 % (((^(v10))) | 1)))) {
                        for fi78 := 0; fi78 < 15; fi78++ {
                            var v79 int64 = 9223372036854775807
                            var v80 int = fi76
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                        for fi81 := 0; fi81 < 2; fi81++ {
                            v3 += ((0.1 - 1.5) - ((((0.1 - ((v2 / 1e10) + ((v2 / ((float64(0.0) / (v3 / 1.5)) + v3)) / v3))) - 3.14159) * (((v3 * v3) / ((v2 - v3) * v3)) - -2.25)) + float64(-2.25)))
                            v82 := []int{fi76, -241}
                            switch (-(fi76)) {
                            case 0:
                                var v83 int32 = 42
                                _ = v83
                            default:
                                v84 := GenS3{F0: ("" + "foo bar")}
                                v82 = append(v82, -2147483648)
                                _ = v84
                            }
                            _ = fi81
                            _ = v82
                        }
                        if true {
                            switch ("" + "foo bar") {
                            case "\n\t":
                                v85 := GenS2{F0: ("a" + ("foo bar" + ("foo bar" + "a"))), F1: 42}
                                _ = v85
                            case "foo bar":
                                switch -1 {
                                case 32768:
                                    fmt.Println(v1)
                                case 9223372036854775807:
                                    var v86 uint32 = 118
                                    _ = v86
                                }
                            default:
                                var v87 uint16 = 255
                                _ = v87
                            }
                            fmt.Println(v1, v1, v77)
                            var v88 int8 = -128
                            _ = v88
                        }
                    } else {
                        for fi89 := 0; fi89 < 11; fi89++ {
                            var v90 uint64 = 425
                            v91 := []int{1}
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    }
                    var v92 uint = 891
                    _ = fi76
                    _ = v77
                    _ = v92
                }
            case -752:
                v3 -= float64((0.1 + ((0.1 * (-2.25 + (0.1 + -2.25))) * 0.0)))
            }
            v10 = 255
        }
        var v93 float32 = (0.0 + -2.25)
        _ = v93
    }
    for fi94 := 0; fi94 < 11; fi94++ {
        for fi95 := 0; fi95 < 18; fi95++ {
            v96 := GenS3{F0: ""}
            _ = fi95
            _ = v96
        }
        _ = fi94
    }
    v97 := GenS6{F0: 1, F1: 2147483647}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v98 := GenS7{F0: 2}
        _ = v98
    case "a":
        for fi99 := 0; fi99 < 19; fi99++ {
            for fi100 := 0; fi100 < 6; fi100++ {
                var v101 uint = 303
                if true {
                    v102 := []int{(fi99 / ((32768) | 1))}
                    if true {
                        v2 -= v2
                        if len(v102) > 0 {
                            v102[0] = (fi99 >> ((v101) & 63))
                        }
                        fmt.Println(fi100, v2)
                        if true {
                            if (32767 != -1) {
                                var v103 bool = false
                                v102 = append(v102, -9223372036854775808)
                                v1 = 127
                                switch -1 {
                                case 381:
                                    var v104 string = ("hello" + ("" + "a"))
                                    _ = v104
                                case -536:
                                    v103 = v103
                                }
                                _ = v103
                            }
                        }
                    }
                    var v105 string = "\n\t"
                    _ = v102
                    _ = v105
                }
                var v106 string = "\n\t"
                _ = fi100
                _ = v101
                _ = v106
            }
            v10 += 18
            for fi107 := 0; fi107 < 13; fi107++ {
                switch (fi99 >> ((818) & 63)) {
                case 543:
                    v1 = (v1 - 1)
                case -996:
                    var v108 float32 = ((((3.14159 * ((1e10 / (-2.25 - 3.14159)) - 0.0)) + ((-2.25 / (0.0 + -2.25)) - (-2.25 - -2.25))) * 1.5) * ((1e10 - 1.5) * 0.1))
                    var v109 int32 = -1
                    var v110 uint16 = 42
                    _ = v108
                    _ = v109
                    _ = v110
                case -2147483648:
                    var v111 bool = false
                    _ = v111
                }
                v2 = (v2 - v3)
                if (!(32767 <= 714)) {
                    var v112 uint16 = 255
                    _ = v112
                }
                _ = fi107
            }
            _ = fi99
        }
    case "hello":
        var v113 uint32 = 417
        _ = v113
    }
    var v114 float64 = ((((float64(0.0) + 1.5) * ((0.1 * ((((v2 * 1.5) / (1.5 * ((v2 / v3) - float64((0.1 * 3.14159))))) - 1e10) - 3.14159)) / v3)) - v2) + 0.0)
    switch "\n\t" {
    case "foo bar":
        for fi115 := 0; fi115 < 12; fi115++ {
            v116 := GenS8{F0: (fi115 / ((fi115) | 1)), F1: 1, F2: 128}
            v3 = (float64(1e10) - (v2 * v114))
            _ = fi115
            _ = v116
        }
    case "a":
        var v117 int = -44
        _ = v117
    }
    var v118 uint16 = 2
    v119 := GenS3{F0: ("" + ("\n\t" + "foo bar"))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v114:", v114)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v114
    _ = v118
    _ = v119
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v97
}
