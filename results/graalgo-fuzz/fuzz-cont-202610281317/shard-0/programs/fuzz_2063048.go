package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 bool }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 string; F2 uint16 }

func main() {
    v1 := []uint32{}
    v2 := []string{"hello", ""}
    v3 := []string{"", ("foo bar" + ("\n\t" + "foo bar")), "hello"}
    var v4 int8 = 38
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]bool)
    if true {
        fmt.Println(v4)
        fmt.Println(v4)
    }
    v4 -= -128
    switch 32768 {
    case 32768:
        if true {
            for fi7 := 0; fi7 < 8; fi7++ {
                v4 -= -118
                v8 := v5[187]
                _ = fi7
                _ = v8
            }
            v4 -= 127
            switch "" {
            case "hello":
                v9 := GenS1{F0: 250, F1: 197}
                _ = v9
            case "\n\t":
                var v10 int32 = -1
                _ = v10
            }
        } else {
            if (724 >= 4294967295) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v4 = (v4 + 1)
                    v12 := GenS2{F0: 1}
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v4 = v4
                    _ = fi13
                }
            } else {
                v4 = (v4 + (v4 % (((v4 % (((-(v4))) | 1))) | 1)))
                switch 801 {
                case 9223372036854775807:
                    var v14 string = "hello"
                    _ = v14
                }
                fmt.Println(v4)
            }
        }
    case 1:
        for fi15 := 0; fi15 < 2; fi15++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
        }
    case 128:
        var v16 int16 = 1000
        _ = v16
    }
    var v17 int64 = -264
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 string = "a"
        v19 = (v19 + v19)
        if (!false) {
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(v17, v19, v17)
                v21 := GenS3{F0: (((v19 + ("" + (v19 + "a"))) + (v19 + "foo bar")) + "\n\t")}
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int = fi22
                var v24 uint8 = 2
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            switch fi18 {
            case 32768:
                v6[758] = true
            case 0:
                v6[65535] = true
            case 1:
                v4 = -128
            default:
                v25 := GenS2{F0: v17}
                _ = v25
            }
            v17 -= v17
            var v26 int = fi18
            _ = v26
        }
        _ = fi18
        _ = v19
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint64 = 422
        switch (fi27 & fi27) {
        case -955:
            var v29 int = fi27
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    v30 := GenS4{F0: -895}
    var v31 uint32 = 0
    var v32 int64 = (^(v17))
    switch 292 {
    case -951:
        var v33 int8 = (v4 ^ v4)
        _ = v33
    case 0:
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 uint8 = 2
            v32 -= (v17 & v32)
            _ = fi34
            _ = v35
        }
    }
    fmt.Println(v32, v17)
    v36 := v5[(v31 >> ((18446744073709551615) & 63))]
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS5{F0: 811, F1: 682, F2: false}
        _ = fi37
        _ = v38
    }
    v39 := GenS6{F0: -383, F1: 95, F2: (0.1 <= float32((1e10 + 0.0)))}
    v17 += 175
    if v36 {
        var v40 string = "\n\t"
        v4 += v4
        if (85 <= 42) {
            if (!true) {
                switch 0 {
                case 942:
                    for fi41 := 0; fi41 < 14; fi41++ {
                        delete(v6, (v31 | 202))
                        switch int(v4) {
                        case 42:
                            switch v40 {
                            case "":
                                if (!(!(!(210 > 0)))) {
                                    var v42 int = fi41
                                    if v36 {
                                        v31 = (v31 + 2147483648)
                                        v4 = v4
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v42
                                }
                            case "\n\t":
                                var v43 string = v40
                                _ = v43
                            case "hello":
                                switch (fi41 << ((18446744073709551615) & 63)) {
                                case 127:
                                    var v44 int8 = -1
                                    _ = v44
                                case 9223372036854775807:
                                    v45 := GenS7{F0: 0}
                                    _ = v45
                                }
                            }
                        case 4294967295:
                            v17 += -845
                        case 9223372036854775807:
                            var v46 string = ("foo bar" + "")
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 17; fi47++ {
                            v48 := GenS6{F0: int(v31), F1: 0, F2: v36}
                            var v49 int8 = -1
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        _ = fi41
                    }
                case -41:
                    var v50 uint16 = 32767
                    _ = v50
                case -706:
                    fmt.Println(v32)
                default:
                    switch v40 {
                    case "":
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 int = (fi51 ^ fi51)
                            if false {
                                if (!((fi51 ^ -811) <= v52)) {
                                    v53 := GenS8{F0: false, F1: -9223372036854775808}
                                    v52 = fi51
                                    _ = v53
                                }
                                fmt.Println(fi51)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                switch (v40 + (v40 + "hello")) {
                                case "hello":
                                    switch fi51 {
                                    case 407:
                                        v36 = ((v36 && (float64(1.5) <= 3.14159)) || v36)
                                    case -1:
                                        for fi54 := 0; fi54 < 5; fi54++ {
                                            v36 = (v32 == -1)
                                            v31 -= v31
                                            _ = fi54
                                        }
                                    case 490:
                                        var v55 uint16 = 529
                                        _ = v55
                                    }
                                case "":
                                    switch fi51 {
                                    case -2147483648:
                                        if (219 >= 833) {
                                            v56 := GenS9{F0: 0, F1: 493, F2: 42}
                                            if (((v31 / ((53) | 1)) <= v31) && ((-(fi51)) == (^(fi51)))) {
                                                for fi57 := 0; fi57 < 11; fi57++ {
                                                    if (1e10 < 1e10) {
                                                        if (float64(3.14159) > (0.1 + -2.25)) {
                                                            if (!true) {
                                                                if false {
                                                                    v58 := GenS10{F0: -32768}
                                                                    _ = v58
                                                                }
                                                                switch v40 {
                                                                case "\n\t":
                                                                    switch (v52 * -577) {
                                                                    case 9223372036854775807:
                                                                        v59 := GenS1{F0: 2, F1: v31}
                                                                        _ = v59
                                                                    default:
                                                                        v52 -= v52
                                                                    }
                                                                }
                                                                if true {
                                                                    v31 -= v31
                                                                    fmt.Println(v40, v32, v52)
                                                                }
                                                            }
                                                            if (0.1 == 0.0) {
                                                                var v60 int16 = 1
                                                                var v61 bool = v36
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            if v36 {
                                                                if (3.14159 < (1.5 / (float32(0.0) + -2.25))) {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    v62 := GenS8{F0: (v36 && v36), F1: v17}
                                                                    _ = v62
                                                                }
                                                            } else {
                                                                if v36 {
                                                                    fmt.Println(v52)
                                                                    if ((v4 & (v4 + -128)) != (v4 | v4)) {
                                                                        if v36 {
                                                                            v1 = append(v1, (v31 | 42))
                                                                            if (!false) {
                                                                                if (!true) {
                                                                                    if v36 {
                                                                                        fmt.Println(v31, v40, fi51)
                                                                                    } else {
                                                                                        _ = len(v5)
                                                                                        var v63 uint16 = 386
                                                                                        if (-2147483648 >= 123) {
                                                                                            fmt.Println(v4)
                                                                                            v64 := GenS8{F0: v36, F1: (v32 / ((v32) | 1))}
                                                                                            var v65 float64 = 3.14159
                                                                                            _ = v64
                                                                                            _ = v65
                                                                                        }
                                                                                        _ = v63
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    v52 = -9223372036854775808
                                                                                    if (((!(true || v36)) && v36) || (v4 > 100)) {
                                                                                        v66 := GenS7{F0: 127}
                                                                                        var v67 int64 = 0
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                } else {
                                                                                    var v68 float32 = 1e10
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (^(v31))
                                                                                    }
                                                                                    var v69 int8 = 1
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                                if (0 >= 1) {
                                                                                    var v70 uint16 = 0
                                                                                    _ = v70
                                                                                }
                                                                            } else {
                                                                                var v71 float32 = 0.0
                                                                                _ = v71
                                                                            }
                                                                        } else {
                                                                            v72 := GenS6{F0: fi51, F1: 255, F2: (!(v52 <= -9223372036854775808))}
                                                                            _ = v72
                                                                        }
                                                                    } else {
                                                                        var v73 int = fi57
                                                                        _ = len(v5)
                                                                        fmt.Println(fi51, v36, v40)
                                                                        _ = v73
                                                                    }
                                                                    fmt.Println(fi51)
                                                                    fmt.Println(v32, v40)
                                                                    v4 += v4
                                                                }
                                                                v17 = (v17 / (((v17 << ((255) & 63))) | 1))
                                                                v74 := GenS2{F0: 785}
                                                                _ = v74
                                                            }
                                                        }
                                                        v52 += (fi57 ^ fi57)
                                                        fmt.Println(v40, v17)
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                    _ = fi57
                                                }
                                                v75 := GenS9{F0: 255, F1: 2, F2: 1}
                                                _ = v75
                                            } else {
                                                if v36 {
                                                    v76 := GenS10{F0: -32768}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v76
                                                }
                                                v77 := GenS3{F0: v40}
                                                _ = v77
                                            }
                                            for fi78 := 0; fi78 < 1; fi78++ {
                                                var v79 uint16 = 1
                                                var v80 float32 = 0.0
                                                switch (v40 + (v40 + "foo bar")) {
                                                case "foo bar":
                                                    _ = len(v2)
                                                case "":
                                                    if (42 < 2) {
                                                        switch (-(fi51)) {
                                                        case -1:
                                                            v31 += (^(v31))
                                                        }
                                                        v79 = (v79 / ((v79) | 1))
                                                    } else {
                                                        fmt.Println(v31, v4)
                                                        switch ((v40 + v40) + "\n\t") {
                                                        case "\n\t":
                                                            switch fi51 {
                                                            case 919:
                                                                if v36 {
                                                                    v4 = 100
                                                                    _ = len(v6)
                                                                }
                                                            case 0:
                                                                var v81 float32 = (0.1 * 0.0)
                                                                _ = v81
                                                            case 128:
                                                                var v82 int32 = -2147483648
                                                                _ = v82
                                                            }
                                                        }
                                                        if (128 < 128) {
                                                            fmt.Println(v79, v36, v52)
                                                            if v36 {
                                                                if (18446744073709551615 > 87) {
                                                                    if true {
                                                                        var v83 uint = 0
                                                                        fmt.Println(v83)
                                                                        if v36 {
                                                                            _ = len(v6)
                                                                        } else {
                                                                            if ((v83 - (v83 / (((v83 | v83)) | 1))) != 255) {
                                                                                v17 = v17
                                                                            }
                                                                        }
                                                                        _ = v83
                                                                    }
                                                                    if true {
                                                                        v79 = 984
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v40 = v40
                                                                        case "\n\t":
                                                                            v3 = append(v3, v40)
                                                                        }
                                                                        v17 += (v17 / (((v17 - (v32 ^ v32))) | 1))
                                                                        if v36 {
                                                                            var v84 int64 = 42
                                                                            _ = len(v2)
                                                                            var v85 uint8 = 0
                                                                            _ = v84
                                                                            _ = v85
                                                                        } else {
                                                                            fmt.Println(v17, v4, v52)
                                                                        }
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            } else {
                                                                v2 = append(v2, "a")
                                                                v86 := v6[v31]
                                                                _ = v86
                                                            }
                                                        } else {
                                                            v32 = int64(v31)
                                                            var v87 int32 = 0
                                                            v88 := GenS4{F0: (v87 / (((v87 / (((v87 % (((v87 | v87)) | 1))) | 1))) | 1))}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    }
                                                default:
                                                    v17 += (v17 * v32)
                                                }
                                                var v89 uint64 = 18446744073709551615
                                                _ = fi78
                                                _ = v79
                                                _ = v80
                                                _ = v89
                                            }
                                            _ = v56
                                        } else {
                                            v36 = (v36 && true)
                                            v90 := GenS9{F0: 0, F1: 255, F2: 32}
                                            var v91 int32 = int32(v31)
                                            _ = v90
                                            _ = v91
                                        }
                                    case -716:
                                        var v92 int = fi51
                                        _ = v92
                                    default:
                                        for fi93 := 0; fi93 < 7; fi93++ {
                                            v6[(v31 | (v31 ^ (v31 ^ v31)))] = (2 <= 0)
                                            var v94 uint16 = 32767
                                            v3 = append(v3, v40)
                                            _ = fi93
                                            _ = v94
                                        }
                                    }
                                }
                            }
                            v95 := GenS4{F0: 792}
                            _ = fi51
                            _ = v52
                            _ = v95
                        }
                    case "foo bar":
                        switch 128 {
                        case -1:
                            v96 := GenS7{F0: 1}
                            _ = v96
                        }
                    default:
                        v97 := v6[v31]
                        _ = v97
                    }
                }
            }
            v17 = 9223372036854775807
            var v98 int64 = (v17 | -703)
            _ = v98
        }
        _ = v40
    } else {
        var v99 string = "\n\t"
        if len(v1) > 0 {
            v1[0] = 65535
        }
        if false {
            v100 := GenS11{F0: (v17 ^ (v32 << ((0) & 63))), F1: (v99 + v99), F2: 255}
            _ = v100
        }
        _ = v99
    }
    v101 := GenS5{F0: 2, F1: 0, F2: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
