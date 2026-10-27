package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int8; F2 int }
type GenS9 struct { F0 int16; F1 int8; F2 int64 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int; F1 int; F2 uint16 }

func main() {
    var v1 int16 = 0
    var v2 string = (("a" + "foo bar") + "foo bar")
    v3 := []string{}
    var v4 uint64 = 1
    v1 += (v1 | -1)
    var v5 uint32 = 1
    var v6 uint32 = v5
    switch "a" {
    case "a":
        var v7 int64 = 1
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: (2147483647 > -122)}
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint8 = 0
        var v11 uint64 = v4
        var v12 uint32 = 42
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = uint(v5)
        fmt.Println(v4)
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 612}
    if true {
        var v16 float32 = 1.5
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 16; fi17++ {
            fmt.Println(v2, v5, v4)
            _ = fi17
        }
        if len(v3) > 0 {
            v3[0] = v2
        }
    } else {
        fmt.Println(v2)
        if true {
            fmt.Println(v2)
            if ((v6 / (((v6 | v6)) | 1)) < (v6 % ((v5) | 1))) {
                if (!true) {
                    v18 := GenS2{F0: 255}
                    _ = v18
                } else {
                    if false {
                        var v19 int16 = (v1 ^ v1)
                        switch 4294967295 {
                        case -2147483648:
                            v6 = v6
                        default:
                            fmt.Println(v19)
                        }
                        if true {
                            v4 += v4
                        }
                        _ = v19
                    }
                }
            }
            v20 := GenS2{F0: 4294967295}
            _ = v20
        } else {
            v21 := GenS2{F0: 621}
            v6 += (v6 & v5)
            _ = v21
        }
    }
    if (127 >= 42) {
        var v22 int16 = v1
        switch v2 {
        case "a":
            _ = len(v3)
        }
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        if (false && false) {
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v2)
                _ = fi24
            }
            v4 -= (v4 % ((388) | 1))
            _ = len(v3)
        } else {
            var v25 bool = false
            for fi26 := 0; fi26 < 12; fi26++ {
                v25 = (!(!v25))
                v1 = (v1 << ((167) & 63))
                switch 0 {
                case 0:
                    v25 = (3.14159 > 1.5)
                default:
                    v1 = int16(v4)
                }
                _ = fi26
            }
            var v27 float32 = 1.5
            _ = v25
            _ = v27
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS1{F0: (((^(v6)) <= v5) || (1 > 9))}
            if (!false) {
                v30 := GenS2{F0: 525}
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v31 := GenS3{F0: 0, F1: (v4 * 0), F2: (^(v4))}
                _ = v30
                _ = v31
            } else {
                var v32 int16 = 1000
                if true {
                    _ = len(v3)
                    fmt.Println(fi28, v1, v1)
                    if false {
                        var v33 int16 = 1
                        if (!(v2 == (((v2 + v2) + v2) + (v2 + v2)))) {
                            var v34 uint32 = v6
                            _ = v34
                        }
                        _ = v33
                    }
                }
                _ = v32
            }
            _ = len(v3)
            _ = fi28
            _ = v29
        }
        fmt.Println(v2, v4, v1)
        _ = fi23
    }
    v35 := GenS3{F0: 0, F1: 0, F2: v4}
    if false {
        switch 127 {
        case 0:
            for fi36 := 0; fi36 < 11; fi36++ {
                if (!false) {
                    v5 -= (v5 >> ((2) & 63))
                    var v37 uint = 255
                    var v38 int32 = 1
                    _ = v37
                    _ = v38
                } else {
                    var v39 int64 = 652
                    _ = v39
                }
                _ = fi36
            }
        case 4294967295:
            for fi40 := 0; fi40 < 19; fi40++ {
                var v41 int32 = -751
                if true {
                    fmt.Println(v5, v1, v4)
                    if false {
                        var v42 int64 = 0
                        if (255 <= 102) {
                            for fi43 := 0; fi43 < 12; fi43++ {
                                switch v2 {
                                case "hello":
                                    switch fi43 {
                                    case -9223372036854775808:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 2147483647:
                                        fmt.Println(v1, v41)
                                    case 0:
                                        if true {
                                            if false {
                                                v44 := GenS4{F0: (^(v5)), F1: v4, F2: v4}
                                                _ = v44
                                            }
                                            var v45 uint = 65535
                                            fmt.Println(v4, fi40, fi43)
                                            _ = v45
                                        } else {
                                            v41 += (^(v41))
                                        }
                                    }
                                default:
                                    switch fi40 {
                                    case -180:
                                        v46 := GenS5{F0: 0, F1: -2147483648}
                                        _ = v46
                                    case -522:
                                        var v47 string = "\n\t"
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v6 = 2
                                        switch v47 {
                                        case "hello":
                                            switch (fi43 ^ (fi43 - (fi43 >> ((42) & 63)))) {
                                            case 9223372036854775807:
                                                switch fi40 {
                                                case 127:
                                                    switch ((v2 + (v2 + v2)) + v2) {
                                                    case "foo bar":
                                                        v48 := GenS3{F0: v42, F1: (v4 ^ v4), F2: (v4 & v4)}
                                                        _ = v48
                                                    case "a":
                                                        v1 = -32768
                                                    }
                                                case 4294967295:
                                                    var v49 uint = 2
                                                    _ = v49
                                                default:
                                                    v3 = append(v3, "foo bar")
                                                }
                                            case 1:
                                                switch fi43 {
                                                case 0:
                                                    v42 -= 175
                                                case 9223372036854775807:
                                                    var v50 uint16 = 49
                                                    _ = v50
                                                default:
                                                    v51 := GenS1{F0: false}
                                                    _ = v51
                                                }
                                            default:
                                                if (float64(1.5) < 3.14159) {
                                                    if (false && (v2 >= ("" + "\n\t"))) {
                                                        fmt.Println(v5, v47)
                                                    } else {
                                                        switch v2 {
                                                        case "a":
                                                            if len(v3) > 0 {
                                                                v3[0] = ("hello" + v47)
                                                            }
                                                        }
                                                        if (3.14159 == float64(((1.5 * ((-2.25 + 1e10) / 1.0)) - 0.1))) {
                                                            v41 = (-(v41))
                                                            v1 = v1
                                                        } else {
                                                            var v52 uint16 = 42
                                                            _ = v52
                                                        }
                                                        v41 += -570
                                                    }
                                                } else {
                                                    if false {
                                                        var v53 uint8 = 42
                                                        var v54 uint8 = v53
                                                        v55 := GenS6{F0: 73, F1: 65535}
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    v47 = v2
                                                    if (!(!(!((((3.14159 + -2.25) / 1.0) * ((float32(float64(1e10)) / -2.25) - (((((-2.25 - 0.1) + 1.5) - (((1.5 / -2.25) * (0.0 * 1e10)) + (((1e10 + 0.0) * 1.5) - 0.0))) + (-2.25 * (1.5 - 0.1))) / (-2.25 - 0.0)))) >= ((0.0 * (0.1 / (float32(1.5) - (0.0 + 3.14159)))) * (float32(float64(1.5)) - 1e10)))))) {
                                                        var v56 uint8 = 255
                                                        var v57 uint8 = v56
                                                        v3 = append(v3, "\n\t")
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        v42 = v42
                                                        var v58 uint = 104
                                                        _ = v58
                                                    }
                                                }
                                            }
                                        default:
                                            var v59 uint8 = 0
                                            _ = v59
                                        }
                                        v6 += (^(v5))
                                        v4 = 83
                                        if false {
                                            var v60 int16 = v1
                                            var v61 int16 = (v60 * (v60 & -833))
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = v47
                                    default:
                                        fmt.Println(fi43, v5)
                                    }
                                }
                                fmt.Println(fi43, v2)
                                switch ("\n\t" + v2) {
                                case "a":
                                    var v62 uint = 2
                                    _ = v62
                                default:
                                    switch (v2 + "") {
                                    case "a":
                                        fmt.Println(v1, v42, v5)
                                    case "foo bar":
                                        var v63 string = "a"
                                        var v64 int = (-(fi40))
                                        if (0 <= int64(v1)) {
                                            fmt.Println(v41, v4)
                                            var v65 int = -9223372036854775808
                                            var v66 float32 = 1.5
                                            v67 := GenS5{F0: v65, F1: v41}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v63
                                        _ = v64
                                    case "\n\t":
                                        var v68 uint64 = 695
                                        _ = v68
                                    }
                                }
                                _ = fi43
                            }
                            fmt.Println(fi40, v41)
                            v5 = v5
                        }
                        fmt.Println(v2, v42)
                        _ = v42
                    }
                    var v69 uint16 = 732
                    _ = v69
                } else {
                    var v70 float64 = 1.5
                    for fi71 := 0; fi71 < 10; fi71++ {
                        v3 = append(v3, ((v2 + v2) + v2))
                        _ = fi71
                    }
                    _ = v70
                }
                var v72 int16 = (v1 ^ (-(v1)))
                _ = fi40
                _ = v41
                _ = v72
            }
        default:
            v5 = (v6 << ((1) & 63))
        }
        v73 := GenS1{F0: ((((true || ((v2 != "hello") && true)) || false) || false) || (-2147483648 < 0))}
        switch 288 {
        case -2147483648:
            var v74 uint32 = (v6 * 439)
            _ = v74
        case 32767:
            for fi75 := 0; fi75 < 4; fi75++ {
                v5 += uint32(v4)
                _ = fi75
            }
        case 404:
            for fi76 := 0; fi76 < 10; fi76++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi77 := 0; fi77 < 3; fi77++ {
                    v6 = (v5 * 607)
                    switch v2 {
                    case "a":
                        switch fi76 {
                        case 771:
                            if (!true) {
                                var v78 float64 = 1e10
                                var v79 uint16 = 42
                                _ = v78
                                _ = v79
                            } else {
                                v80 := GenS7{F0: (v6 & v5)}
                                v3 = append(v3, "hello")
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v81 uint8 = uint8(v4)
                                _ = v80
                                _ = v81
                            }
                        default:
                            switch ((v2 + "a") + v2) {
                            case "hello":
                                if (1e10 > float64((1e10 - 1.5))) {
                                    switch (-(fi76)) {
                                    case -767:
                                        v3 = append(v3, "hello")
                                    case -1:
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v6, v1)
                                        case "foo bar":
                                            fmt.Println(fi77, fi76, v4)
                                        case "":
                                            var v82 float64 = (-2.25 * 3.14159)
                                            _ = v82
                                        default:
                                            var v83 uint16 = 32767
                                            _ = v83
                                        }
                                    case 0:
                                        if (false || false) {
                                            _ = len(v3)
                                            v84 := GenS8{F0: 0, F1: 127, F2: fi76}
                                            var v85 int16 = v1
                                            _ = v84
                                            _ = v85
                                        }
                                    default:
                                        v3 = append(v3, v2)
                                    }
                                    var v86 uint = 65535
                                    _ = v86
                                }
                            case "foo bar":
                                var v87 uint = 1
                                _ = v87
                            case "":
                                if false {
                                    if false {
                                        var v88 int32 = -964
                                        v89 := GenS9{F0: -107, F1: 0, F2: 1}
                                        v3 = append(v3, v2)
                                        _ = v88
                                        _ = v89
                                    } else {
                                        v1 += v1
                                        v4 += 0
                                    }
                                    v4 = v4
                                }
                            default:
                                if false {
                                    if (fi76 >= (^(fi76))) {
                                        var v90 bool = (!(((1.5 - 0.0) + 3.14159) == ((1e10 * 0.1) * 1e10)))
                                        v3 = append(v3, v2)
                                        _ = v90
                                    }
                                } else {
                                    fmt.Println(fi76, v2)
                                    var v91 uint = 879
                                    _ = v91
                                }
                            }
                        }
                    case "\n\t":
                        if (false || false) {
                            var v92 int = 4294967295
                            if false {
                                v6 = v6
                                if (!((!true) || true)) {
                                    switch (v92 | fi76) {
                                    case -26:
                                        fmt.Println(fi77, fi76, fi76)
                                    default:
                                        switch fi76 {
                                        case 128:
                                            var v93 int64 = -248
                                            _ = v93
                                        case -600:
                                            if len(v3) > 0 {
                                                v3[0] = "a"
                                            }
                                        case -853:
                                            if (v5 == v5) {
                                                v94 := GenS10{F0: 1, F1: (v4 - 1)}
                                                var v95 bool = false
                                                if (uint(v4) >= uint(v5)) {
                                                    switch 70 {
                                                    case -167:
                                                        v92 = (fi77 & (fi77 - -880))
                                                    case 2147483647:
                                                        v95 = v95
                                                    }
                                                    var v96 uint64 = v4
                                                    v95 = (127 != 42)
                                                    _ = v96
                                                }
                                                v1 += v1
                                                if len(v3) > 0 {
                                                    v3[0] = (v2 + "a")
                                                }
                                                if false {
                                                    v97 := GenS5{F0: fi77, F1: 2147483647}
                                                    if v95 {
                                                        v6 -= 4294967295
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = v2
                                                    }
                                                    var v98 uint8 = 77
                                                    fmt.Println(v98, v1, fi76)
                                                    switch (fi77 * (v92 << ((4294967295) & 63))) {
                                                    case -2147483648:
                                                        switch fi76 {
                                                        case 42:
                                                            var v99 int16 = v1
                                                            _ = v99
                                                        default:
                                                            var v100 int64 = -9223372036854775808
                                                            _ = v100
                                                        }
                                                    case 255:
                                                        var v101 uint64 = (v4 << ((255) & 63))
                                                        _ = v101
                                                    default:
                                                        v4 -= (v4 * 0)
                                                    }
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    v6 = (v5 >> ((255) & 63))
                                                    fmt.Println(v2)
                                                    var v102 float32 = 0.0
                                                    _ = v102
                                                }
                                                if (!true) {
                                                    if ((53 == uint16(v5)) && true) {
                                                        var v103 int32 = 472
                                                        if len(v3) > 0 {
                                                            v3[0] = v2
                                                        }
                                                        v6 += (v5 | (^(v5)))
                                                        var v104 float64 = float64(float32(3.14159))
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    var v105 int16 = v1
                                                    var v106 int8 = 22
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    var v107 uint64 = 101
                                                    if (v2 >= v2) {
                                                        v107 = v4
                                                        if v95 {
                                                            if v95 {
                                                                v3 = append(v3, v2)
                                                            } else {
                                                                var v108 uint8 = 1
                                                                var v109 uint = 874
                                                                v2 = v2
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                        } else {
                                                            v95 = false
                                                        }
                                                    }
                                                    _ = v107
                                                }
                                                _ = v94
                                                _ = v95
                                            }
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v1 += v1
                                }
                                v6 = 151
                            }
                            if (true && (3.14159 <= (-2.25 + (3.14159 + 1e10)))) {
                                v4 -= v4
                            }
                            _ = v92
                        }
                    case "":
                        v3 = append(v3, v2)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi77
                }
                var v110 int64 = -1
                _ = fi76
                _ = v110
            }
        }
        _ = v73
    }
    v2 = ("" + (v2 + v2))
    v111 := GenS11{F0: int(v1), F1: -899, F2: 42}
    v4 -= 42
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
