package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint32; F2 int8 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint16; F1 uint; F2 int64 }
type GenS10 struct { F0 uint; F1 int16; F2 int }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 string; F1 uint64 }
type GenS13 struct { F0 uint8; F1 bool; F2 string }
type GenS14 struct { F0 int32; F1 uint; F2 int64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int32; F1 uint8; F2 int8 }
type GenS17 struct { F0 int; F1 uint16 }
func genFn0() uint16 {
    var v5 int32 = 668
    v6 := GenS1{F0: 42}
    switch 32767 {
    case 127:
        var v7 uint16 = 389
        _ = v7
    case 0:
        if false {
            v5 = (v5 / ((v5) | 1))
            fmt.Println(v5)
        } else {
            switch "" {
            case "a":
                var v8 uint64 = 645
                _ = v8
            case "hello":
                var v9 int8 = -128
                _ = v9
            }
            fmt.Println(v5)
        }
    }
    v5 += (v5 % ((v5) | 1))
    var v10 int64 = 627
    _ = v10
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []string{"hello", "foo bar"}
    v2 := []string{"hello", ("\n\t" + ""), "foo bar"}
    var v3 string = "hello"
    var v4 uint = 255
    _ = len(v2)
    if (true && false) {
        switch -2147483648 {
        case 127:
            var v11 float32 = 1e10
            _ = v11
        case -101:
            v4 += v4
        case 694:
            fmt.Println(v3, v4)
        }
        if false {
            if true {
                if true {
                    v4 = v4
                    for fi12 := 0; fi12 < 8; fi12++ {
                        if true {
                            v13 := GenS1{F0: 42}
                            _ = v13
                        } else {
                            fmt.Println(fi12, v3)
                            _ = len(v2)
                        }
                        _ = fi12
                    }
                } else {
                    v4 = (v4 << (((v4 >> ((1) & 63))) & 63))
                }
                v14 := GenS1{F0: 42}
                v15 := GenS2{F0: 32767}
                _ = v14
                _ = v15
            } else {
                var v16 uint = 18446744073709551615
                _ = v16
            }
        } else {
            v3 = ("a" + (v3 + "a"))
            if ((true && (!true)) && (-758 < 42)) {
                fmt.Println(v4, v3)
                var v17 float64 = 1.5
                _ = v17
            }
            for fi18 := 0; fi18 < 15; fi18++ {
                switch 42 {
                case 4294967295:
                    switch fi18 {
                    case 32767:
                        v4 -= 1
                    }
                }
                _ = fi18
            }
        }
    }
    if (!true) {
        if (613 > 1000) {
            v19 := GenS3{F0: 0, F1: 65535}
            _ = v19
        }
    } else {
        switch v3 {
        case "a":
            var v20 uint16 = 2
            _ = v20
        case "":
            v21 := GenS3{F0: -73, F1: 1}
            _ = v21
        }
        switch v3 {
        case "":
            if false {
                for fi22 := 0; fi22 < 6; fi22++ {
                    fmt.Println(fi22, v3)
                    for fi23 := 0; fi23 < 7; fi23++ {
                        _ = len(v2)
                        var v24 int16 = 200
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 8; fi25++ {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch v3 {
                        case "a":
                            v26 := GenS4{F0: 65535, F1: -2147483648}
                            _ = v26
                        case "":
                            v27 := GenS1{F0: 2147483648}
                            _ = v27
                        case "\n\t":
                            fmt.Println(v3)
                        }
                        var v28 string = v3
                        _ = fi25
                        _ = v28
                    }
                    _ = fi22
                }
                if (9223372036854775807 == 4294967295) {
                    v3 = "\n\t"
                    switch 42 {
                    case 128:
                        for fi29 := 0; fi29 < 8; fi29++ {
                            if true {
                                if true {
                                    if (false || (false && false)) {
                                        for fi30 := 0; fi30 < 2; fi30++ {
                                            switch (fi30 ^ 128) {
                                            case 127:
                                                fmt.Println(fi30)
                                            case -1:
                                                v3 = v3
                                            default:
                                                v31 := GenS5{F0: 1, F1: 42, F2: 1}
                                                _ = v31
                                            }
                                            var v32 uint8 = 1
                                            _ = fi30
                                            _ = v32
                                        }
                                    }
                                    if true {
                                        var v33 float64 = 0.0
                                        for fi34 := 0; fi34 < 5; fi34++ {
                                            v35 := GenS3{F0: int8(fi29), F1: 4294967295}
                                            v36 := GenS6{F0: (394 >= 65535), F1: 65535, F2: 1}
                                            if false {
                                                v33 -= v33
                                                fmt.Println(v33)
                                            }
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                        }
                                        v4 = (v4 * v4)
                                        _ = v33
                                    } else {
                                        v4 -= (v4 >> ((42) & 63))
                                        v3 = v3
                                    }
                                } else {
                                    switch ("foo bar" + v3) {
                                    case "a":
                                        v4 -= v4
                                    case "\n\t":
                                        for fi37 := 0; fi37 < 16; fi37++ {
                                            v3 = v3
                                            _ = fi37
                                        }
                                    case "hello":
                                        if false {
                                            var v38 int32 = 0
                                            _ = v38
                                        } else {
                                            var v39 float64 = (3.14159 / 3.14159)
                                            v3 = (((v3 + v3) + ("hello" + v3)) + "\n\t")
                                            _ = v39
                                        }
                                    default:
                                        for fi40 := 0; fi40 < 20; fi40++ {
                                            var v41 uint64 = 42
                                            _ = fi40
                                            _ = v41
                                        }
                                    }
                                }
                            }
                            fmt.Println(v4, v4, v4)
                            if (!false) {
                                var v42 int = -786
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    if (-2.25 == (0.0 - -2.25)) {
                                        var v44 int = v42
                                        v4 = v4
                                        var v45 uint32 = 4294967295
                                        _ = v44
                                        _ = v45
                                    }
                                    fmt.Println(fi43)
                                    _ = fi43
                                }
                                _ = v42
                            } else {
                                v4 = (v4 | v4)
                                var v46 uint8 = 127
                                _ = v46
                            }
                            _ = fi29
                        }
                    case -378:
                        v3 = ("" + v3)
                    }
                    v3 = (v3 + v3)
                } else {
                    for fi47 := 0; fi47 < 18; fi47++ {
                        if (2 == 1) {
                            if false {
                                if (!false) {
                                    switch "\n\t" {
                                    case "":
                                        if (!true) {
                                            if true {
                                                v2 = append(v2, v3)
                                                var v48 int16 = 32767
                                                _ = v48
                                            }
                                            v1 = append(v1, v3)
                                        } else {
                                            var v49 uint64 = 123
                                            if false {
                                                v49 += (v49 + v49)
                                            } else {
                                                for fi50 := 0; fi50 < 1; fi50++ {
                                                    switch fi47 {
                                                    case 638:
                                                        v51 := GenS4{F0: 2147483647, F1: 2147483647}
                                                        _ = v51
                                                    case 440:
                                                        if len(v2) > 0 {
                                                            v2[0] = (v3 + "a")
                                                        }
                                                    }
                                                    _ = fi50
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v52 float32 = (1.5 * (3.14159 + (float32((0.1 - 1e10)) + 0.0)))
                                                _ = v52
                                            }
                                            if false {
                                                v53 := GenS7{F0: 128, F1: (^(fi47))}
                                                v54 := GenS8{F0: 0, F1: v49, F2: 381}
                                                _ = v53
                                                _ = v54
                                            }
                                            _ = v49
                                        }
                                    case "foo bar":
                                        var v55 uint32 = 65535
                                        _ = v55
                                    default:
                                        v2 = append(v2, v3)
                                    }
                                    for fi56 := 0; fi56 < 4; fi56++ {
                                        var v57 int = (^(fi56))
                                        if (2 == 581) {
                                            var v58 bool = true
                                            if (true || ((4294967295 == 42) || (42 != 2))) {
                                                _ = len(v2)
                                                v1 = append(v1, v3)
                                                v59 := GenS9{F0: 2, F1: v4, F2: 9223372036854775807}
                                                _ = v59
                                            } else {
                                                var v60 int16 = -32768
                                                if (!true) {
                                                    var v61 uint64 = 0
                                                    var v62 float32 = float32((1e10 * float64(((((1.5 / (1e10 + float32(1e10))) / 1.5) - float32(1e10)) + (0.1 / 1.0)))))
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    var v63 int32 = -12
                                                    switch 9223372036854775807 {
                                                    case 353:
                                                        v60 -= -32768
                                                        switch (fi56 & (v57 | (fi56 + 128))) {
                                                        case 2147483647:
                                                            fmt.Println(fi56, v58, v63)
                                                        case 127:
                                                            if len(v2) > 0 {
                                                                v2[0] = "hello"
                                                            }
                                                        }
                                                    case -1:
                                                        v1 = append(v1, "\n\t")
                                                    case 9223372036854775807:
                                                        v3 = (v3 + "hello")
                                                    }
                                                    v64 := GenS10{F0: (v4 | 2), F1: (v60 + (v60 % (((v60 & (v60 & v60))) | 1))), F2: 32768}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                _ = v60
                                            }
                                            if false {
                                                if false {
                                                    v65 := GenS2{F0: fi47}
                                                    _ = v65
                                                }
                                                var v66 uint16 = 65535
                                                _ = v66
                                            }
                                            var v67 string = ("\n\t" + v3)
                                            if (6 <= 1) {
                                                switch v67 {
                                                case "foo bar":
                                                    var v68 int32 = 2147483647
                                                    _ = v68
                                                case "\n\t":
                                                    v1 = append(v1, (v67 + v67))
                                                case "a":
                                                    var v69 uint32 = 0
                                                    _ = v69
                                                }
                                            } else {
                                                var v70 float64 = (0.1 / 1.0)
                                                if (0.0 < (-2.25 - 0.1)) {
                                                    switch fi47 {
                                                    case -173:
                                                        switch v67 {
                                                        case "\n\t":
                                                            v71 := GenS11{F0: 0, F1: 65535}
                                                            _ = v71
                                                        case "foo bar":
                                                            _ = len(v1)
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        default:
                                                            switch (fi47 - 420) {
                                                            case 4294967295:
                                                                v72 := GenS12{F0: v3, F1: 646}
                                                                _ = v72
                                                            default:
                                                                if false {
                                                                    var v73 int16 = 1000
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v3
                                                                    }
                                                                    _ = v73
                                                                } else {
                                                                    v4 -= 4294967295
                                                                    switch (v67 + v67) {
                                                                    case "hello":
                                                                        var v74 int64 = 651
                                                                        _ = v74
                                                                    case "":
                                                                        if ((993 == -71) || v58) {
                                                                            if (-1 <= 894) {
                                                                                var v75 uint8 = 2
                                                                                if v58 {
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v76 string = v3
                                                                                        v77 := GenS9{F0: 866, F1: v4, F2: -1}
                                                                                        v4 = (v4 & v4)
                                                                                        var v78 string = v76
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    case "foo bar":
                                                                                        v75 = uint8(v4)
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    v3 = "a"
                                                                                    fmt.Println(v57)
                                                                                } else {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (v67 + v3)
                                                                                    }
                                                                                    var v79 int = (v57 % ((v57) | 1))
                                                                                    _ = v79
                                                                                }
                                                                                v1 = append(v1, v3)
                                                                                _ = v75
                                                                            } else {
                                                                                v1 = append(v1, "foo bar")
                                                                                var v80 uint8 = 255
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v4, v57, v70)
                                                                    default:
                                                                        var v81 bool = (true || true)
                                                                        _ = v81
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v2 = append(v2, v3)
                                                    v70 += 1e10
                                                }
                                                if true {
                                                    v70 = v70
                                                    v57 -= fi47
                                                } else {
                                                    if false {
                                                        var v82 uint = 4294967295
                                                        _ = v82
                                                    }
                                                }
                                                _ = v70
                                            }
                                            var v83 string = ((("\n\t" + v3) + v3) + (((v3 + (v3 + "foo bar")) + v67) + "\n\t"))
                                            v4 += v4
                                            fmt.Println(v3, v57)
                                            if (42 != 42) {
                                                v84 := GenS13{F0: 251, F1: (-1 >= 42), F2: "foo bar"}
                                                fmt.Println(fi47, v83, v4)
                                                var v85 string = ("hello" + ("a" + "hello"))
                                                _ = v84
                                                _ = v85
                                            } else {
                                                if (65535 != 42) {
                                                    var v86 int16 = int16(v57)
                                                    v87 := GenS1{F0: 4294967295}
                                                    _ = v86
                                                    _ = v87
                                                }
                                                switch v67 {
                                                case "a":
                                                    fmt.Println(fi47, v4)
                                                case "hello":
                                                    v88 := GenS4{F0: 2147483647, F1: -2147483648}
                                                    _ = v88
                                                }
                                                v2 = append(v2, v83)
                                            }
                                            _ = v58
                                            _ = v67
                                            _ = v83
                                        }
                                        _ = fi56
                                        _ = v57
                                    }
                                    _ = len(v1)
                                } else {
                                    var v89 uint16 = 32767
                                    if true {
                                        var v90 float32 = 3.14159
                                        _ = v90
                                    } else {
                                        if false {
                                            var v91 string = v3
                                            switch (fi47 % ((32767) | 1)) {
                                            case 4294967295:
                                                v2 = append(v2, ("" + "\n\t"))
                                            case 42:
                                                v89 -= (^(v89))
                                            default:
                                                _ = len(v2)
                                            }
                                            _ = v91
                                        }
                                        var v92 int = 32768
                                        _ = v92
                                    }
                                    _ = v89
                                }
                                switch ("foo bar" + "hello") {
                                case "a":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "":
                                    if (0 >= int16(fi47)) {
                                        v4 -= (^(v4))
                                        var v93 int32 = 2147483647
                                        switch 127 {
                                        case -1:
                                            var v94 int8 = 0
                                            _ = v94
                                        case 32768:
                                            switch "" {
                                            case "a":
                                                fmt.Println(v4)
                                            case "foo bar":
                                                v95 := GenS7{F0: fi47, F1: (fi47 / ((fi47) | 1))}
                                                _ = v95
                                            case "\n\t":
                                                v96 := GenS14{F0: v93, F1: (v4 * v4), F2: -868}
                                                _ = v96
                                            }
                                        }
                                        _ = v93
                                    } else {
                                        if false {
                                            v3 = "a"
                                        }
                                        v3 = (v3 + ((v3 + ("" + ("hello" + v3))) + (v3 + "")))
                                    }
                                case "hello":
                                    v3 = "\n\t"
                                }
                                v2 = append(v2, v3)
                            }
                            v97 := GenS9{F0: 2, F1: (v4 ^ 20), F2: -1}
                            var v98 uint8 = 255
                            _ = v97
                            _ = v98
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi99 := 0; fi99 < 8; fi99++ {
                            var v100 bool = true
                            var v101 float32 = ((1e10 * 1.5) + (1e10 + (1e10 - 3.14159)))
                            v102 := GenS15{F0: 1}
                            _ = fi99
                            _ = v100
                            _ = v101
                            _ = v102
                        }
                        _ = fi47
                    }
                }
                for fi103 := 0; fi103 < 10; fi103++ {
                    var v104 uint = (v4 ^ v4)
                    fmt.Println(fi103, v4, v104)
                    _ = fi103
                    _ = v104
                }
            }
        case "\n\t":
            v3 = v3
        }
    }
    var v105 float32 = (0.1 / 1e10)
    var v106 bool = (42 < -2147483648)
    var v107 uint32 = 2
    if v106 {
        v105 -= (v105 / v105)
    }
    fmt.Println(v106)
    v108 := GenS16{F0: 42, F1: 2, F2: 1}
    if v106 {
        if true {
            v105 -= v105
        }
    }
    switch -2147483648 {
    case 32768:
        var v109 int16 = 0
        _ = v109
    case 9223372036854775807:
        v110 := GenS17{F0: 9223372036854775807, F1: 65535}
        _ = v110
    case -922:
        var v111 uint = (v4 | (v4 / ((uint(v107)) | 1)))
        _ = v111
    default:
        var v112 int64 = -899
        _ = v112
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
}
