package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int16; F1 int16; F2 uint64 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 uint = 255
    v3 := make(map[int]bool)
    v4 := []uint32{2}
    v5 := []uint32{65535, 588}
    v6 := GenS1{F0: -1, F1: 0}
    v1 = v1
    var v7 bool = true
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 1
        for fi10 := 0; fi10 < 9; fi10++ {
            if (-98 > (v9 * 42)) {
                switch fi8 {
                case -1:
                    v4 = append(v4, 540)
                    if v7 {
                        var v11 uint = 925
                        if (!false) {
                            switch "a" {
                            case "a":
                                var v12 int32 = -2147483648
                                _ = v12
                            case "foo bar":
                                v7 = ((!false) || v7)
                            default:
                                v1 = 32767
                            }
                            v13 := GenS1{F0: fi8, F1: 42}
                            _ = v13
                        }
                        _ = v11
                    } else {
                        v14 := GenS1{F0: fi8, F1: 763}
                        v1 += v1
                        _ = len(v3)
                        _ = v14
                    }
                    if (-9223372036854775808 < 439) {
                        v5 = append(v5, 859)
                    }
                default:
                    var v15 int16 = 32767
                    _ = v15
                }
                var v16 uint8 = 255
                _ = v16
            } else {
                switch -2147483648 {
                case -941:
                    v17 := GenS2{F0: 2147483647, F1: fi8}
                    _ = len(v3)
                    _ = v17
                case -146:
                    var v18 float64 = (float64(0.1) + 0.0)
                    _ = v18
                default:
                    v9 -= v9
                }
                if (!v7) {
                    v9 = v9
                } else {
                    fmt.Println(v9)
                    if ((float64((1.5 + (((-2.25 * (((float32(((1.5 + ((1.5 - float64(0.1)) * 3.14159)) - (0.0 * float64((-2.25 + 0.0))))) + (float32(1e10) + 0.0)) - -2.25) - float32(1e10))) - 3.14159) * 1.5))) + 0.1) < 3.14159) {
                        var v19 int16 = 42
                        _ = v19
                    }
                    var v20 string = "hello"
                    var v21 uint16 = 32767
                    v22 := GenS2{F0: 982, F1: (fi10 - fi10)}
                    fmt.Println(v9, v2)
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            v23 := GenS1{F0: (fi8 * -2147483648), F1: 87}
            _ = fi10
            _ = v23
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int32 = 42
            _ = fi24
            _ = v25
        }
        _ = fi8
        _ = v9
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v26 int = 827
    var v27 uint16 = uint16(v2)
    var v28 int = v26
    fmt.Println(v2, v27, v27)
    for fi29 := 0; fi29 < 12; fi29++ {
        v28 -= v26
        v7 = false
        v30 := GenS2{F0: 4294967295, F1: (fi29 - (v28 / ((fi29) | 1)))}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: v26, F1: 149}
    switch (("a" + ("" + ("\n\t" + ""))) + "foo bar") {
    case "a":
        if v7 {
            switch ("hello" + "a") {
            case "hello":
                v32 := GenS2{F0: 1, F1: v26}
                _ = v32
            case "\n\t":
                if v7 {
                    if (true && v7) {
                        var v33 int32 = 42
                        _ = v33
                    }
                    if v7 {
                        v27 += v1
                        var v34 uint64 = 428
                        v7 = v7
                        _ = v34
                    } else {
                        if false {
                            if v7 {
                                var v35 uint = v2
                                _ = v35
                            }
                            fmt.Println(v27)
                        }
                        if true {
                            if false {
                                if ((0.1 * 0.1) != 1e10) {
                                    if v7 {
                                        v36 := GenS3{F0: 138, F1: -433}
                                        fmt.Println(v26)
                                        _ = v36
                                    }
                                    var v37 int32 = -625
                                    _ = len(v4)
                                    _ = v37
                                } else {
                                    var v38 uint32 = 633
                                    if ((!v7) || (((v27 == v27) || ((1e10 * 3.14159) <= -2.25)) || (546 >= -2147483648))) {
                                        v39 := GenS4{F0: 42, F1: false}
                                        v40 := GenS4{F0: 0, F1: true}
                                        if (!((3.14159 * 1e10) == 1.5)) {
                                            for fi41 := 0; fi41 < 7; fi41++ {
                                                v42 := GenS5{F0: 255}
                                                switch v26 {
                                                case 0:
                                                    v26 = fi41
                                                }
                                                _ = fi41
                                                _ = v42
                                            }
                                            v43 := GenS5{F0: 42}
                                            switch (-(v28)) {
                                            case 128:
                                                for fi44 := 0; fi44 < 1; fi44++ {
                                                    v1 = v1
                                                    _ = fi44
                                                }
                                            case -349:
                                                if (!v7) {
                                                    var v45 float32 = (3.14159 / 1.0)
                                                    for fi46 := 0; fi46 < 15; fi46++ {
                                                        v47 := GenS6{F0: 1, F1: 42}
                                                        if ((122 < 122) || v7) {
                                                            if ((42 < (v2 & v2)) && true) {
                                                                if v7 {
                                                                    fmt.Println(v7)
                                                                }
                                                                switch (v28 << (((v2 - (v2 >> (((^(v2))) & 63)))) & 63)) {
                                                                case 9223372036854775807:
                                                                    for fi48 := 0; fi48 < 9; fi48++ {
                                                                        fmt.Println(v2)
                                                                        if v7 {
                                                                            v1 -= (^(v27))
                                                                            v49 := GenS3{F0: 1, F1: 0}
                                                                            fmt.Println(fi46, fi48)
                                                                            _ = v49
                                                                        } else {
                                                                            if v7 {
                                                                                v50 := GenS2{F0: 758, F1: fi46}
                                                                                v51 := GenS3{F0: 255, F1: -1}
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                v27 = 190
                                                                            }
                                                                            var v52 uint = (v2 * v2)
                                                                            _ = v52
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            if (v27 <= (^(v1))) {
                                                                                var v53 uint8 = 128
                                                                                if (!v7) {
                                                                                    var v54 int = (fi48 - fi46)
                                                                                    v7 = (802 < 303)
                                                                                    if (!(!(v38 < v38))) {
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v38
                                                                                        }
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            switch (v26 * (fi48 % ((v28) | 1))) {
                                                                                            case -40:
                                                                                                if v7 {
                                                                                                    if ((65535 >= (^(v38))) || v7) {
                                                                                                        var v55 int64 = -1
                                                                                                        v56 := GenS6{F0: -1, F1: -1}
                                                                                                        if v7 {
                                                                                                            fmt.Println(v1)
                                                                                                            v28 -= -1
                                                                                                        } else {
                                                                                                            v57 := GenS5{F0: 127}
                                                                                                            if (1 >= 42) {
                                                                                                                v58 := GenS2{F0: (^(v38)), F1: 4294967295}
                                                                                                                v59 := GenS2{F0: (v38 << (((v2 - (v2 - 732))) & 63)), F1: fi48}
                                                                                                                v60 := GenS4{F0: int32(v55), F1: v7}
                                                                                                                _ = v58
                                                                                                                _ = v59
                                                                                                                _ = v60
                                                                                                            } else {
                                                                                                                v38 += v38
                                                                                                                if v7 {
                                                                                                                    var v61 uint32 = v38
                                                                                                                    _ = v61
                                                                                                                }
                                                                                                            }
                                                                                                            v5 = append(v5, (v38 * v38))
                                                                                                            _ = v57
                                                                                                        }
                                                                                                        _ = v55
                                                                                                        _ = v56
                                                                                                    } else {
                                                                                                        v62 := GenS7{F0: 869, F1: 0}
                                                                                                        v45 += v45
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    v63 := GenS5{F0: v53}
                                                                                                    var v64 float64 = 0.0
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v65 int8 = 100
                                                                                        v66 := GenS3{F0: 128, F1: 656}
                                                                                        v1 = (^(v27))
                                                                                        if v7 {
                                                                                            switch "foo bar" {
                                                                                            case "a":
                                                                                                if (v7 && (-806 == int(v65))) {
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                }
                                                                                            case "hello":
                                                                                                var v67 int64 = -9223372036854775808
                                                                                                _ = v67
                                                                                            case "":
                                                                                                v38 -= (v38 * (v38 ^ v38))
                                                                                            }
                                                                                            var v68 uint = (v2 ^ v2)
                                                                                            var v69 int = (v54 ^ 32767)
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        } else {
                                                                                            var v70 int32 = 1
                                                                                            _ = v70
                                                                                        }
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                    _ = v54
                                                                                } else {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = uint32(v2)
                                                                                    }
                                                                                    v1 = (v1 & 32767)
                                                                                }
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = v38
                                                                                }
                                                                                _ = v53
                                                                            } else {
                                                                                v71 := v3[127]
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                        _ = fi48
                                                                    }
                                                                case 2147483647:
                                                                    var v72 int16 = 42
                                                                    _ = v72
                                                                case 271:
                                                                    v73 := GenS8{F0: -32768, F1: 828, F2: 342}
                                                                    _ = v73
                                                                default:
                                                                    var v74 uint8 = 93
                                                                    _ = v74
                                                                }
                                                            } else {
                                                                fmt.Println(v7, v38, v2)
                                                                var v75 bool = v7
                                                                if (((100 <= 0) || (-1 >= 1)) || (float64((((v45 * v45) + 1.5) - v45)) != 1.5)) {
                                                                    var v76 uint64 = 336
                                                                    for fi77 := 0; fi77 < 14; fi77++ {
                                                                        var v78 uint8 = 128
                                                                        v79 := GenS5{F0: (v78 / (((v78 | v78)) | 1))}
                                                                        _ = fi77
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                    fmt.Println(v1, fi46)
                                                                    _ = v76
                                                                } else {
                                                                    if true {
                                                                        if (357 > (v2 + v2)) {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (v38 << ((0) & 63))
                                                                            }
                                                                            var v80 uint64 = 18446744073709551615
                                                                            if ((127 >= 1) && v7) {
                                                                                v38 = (v38 % (((v38 + v38)) | 1))
                                                                            }
                                                                            _ = v80
                                                                        } else {
                                                                            for fi81 := 0; fi81 < 17; fi81++ {
                                                                                v26 = v28
                                                                                v82 := GenS4{F0: -1, F1: v75}
                                                                                _ = fi81
                                                                                _ = v82
                                                                            }
                                                                            var v83 string = "a"
                                                                            _ = v83
                                                                        }
                                                                        switch fi46 {
                                                                        case 20:
                                                                            _ = len(v3)
                                                                        case -1:
                                                                            v28 = 1
                                                                        case 32768:
                                                                            for fi84 := 0; fi84 < 19; fi84++ {
                                                                                var v85 int64 = 9223372036854775807
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    delete(v3, fi84)
                                                                                case "a":
                                                                                    var v86 bool = (1 == 2)
                                                                                    _ = v86
                                                                                case "":
                                                                                    v1 = (v1 << ((65535) & 63))
                                                                                }
                                                                                _ = fi84
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        v27 -= v27
                                                                    } else {
                                                                        var v87 uint = (v2 & 1)
                                                                        v88 := GenS7{F0: 18446744073709551615, F1: 42}
                                                                        if false {
                                                                            var v89 int = 388
                                                                            var v90 int = fi46
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                    var v91 float64 = (1.5 * 0.0)
                                                                    v92 := GenS9{F0: -9223372036854775808, F1: v27}
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                _ = v75
                                                            }
                                                        }
                                                        _ = fi46
                                                        _ = v47
                                                    }
                                                    switch ("foo bar" + "") {
                                                    case "foo bar":
                                                        for fi93 := 0; fi93 < 2; fi93++ {
                                                            if (!(0.1 == 1.5)) {
                                                                var v94 float32 = (0.0 - (-2.25 / -2.25))
                                                                if true {
                                                                    v95 := GenS10{F0: -48}
                                                                    _ = v95
                                                                }
                                                                v96 := GenS8{F0: 0, F1: -1, F2: 18446744073709551615}
                                                                _ = v94
                                                                _ = v96
                                                            }
                                                            _ = fi93
                                                        }
                                                    case "hello":
                                                        for fi97 := 0; fi97 < 7; fi97++ {
                                                            var v98 int16 = 42
                                                            fmt.Println(fi97)
                                                            _ = fi97
                                                            _ = v98
                                                        }
                                                    case "a":
                                                        var v99 bool = v7
                                                        _ = v99
                                                    default:
                                                        if false {
                                                            v2 = v2
                                                            for fi100 := 0; fi100 < 11; fi100++ {
                                                                fmt.Println(v1, v1, v7)
                                                                _ = fi100
                                                            }
                                                            var v101 int16 = 725
                                                            _ = v101
                                                        }
                                                    }
                                                    _ = v45
                                                }
                                            case -638:
                                                var v102 uint = (v2 >> ((4294967295) & 63))
                                                _ = v102
                                            }
                                            _ = v43
                                        } else {
                                            v103 := GenS4{F0: -2147483648, F1: (!v7)}
                                            _ = v103
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                    var v104 uint64 = 42
                                    _ = v104
                                    _ = v38
                                }
                                v105 := GenS8{F0: -510, F1: -1, F2: 42}
                                v106 := v3[(v28 & v28)]
                                _ = v105
                                _ = v106
                            } else {
                                v1 += 42
                            }
                        }
                        v7 = v7
                    }
                }
            case "a":
                v3[v28] = v7
            }
        } else {
            v107 := GenS10{F0: -82}
            delete(v3, (v26 + (-(v28))))
            v108 := GenS3{F0: 13, F1: 1}
            _ = v107
            _ = v108
        }
    default:
        fmt.Println(v28, v27, v7)
    }
    switch "\n\t" {
    case "":
        var v109 int64 = 0
        _ = v109
    }
    var v110 int8 = 1
    v111 := GenS4{F0: -2147483648, F1: true}
    fmt.Println(v7, v110, v110)
    delete(v3, -206)
    var v112 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
