package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
type GenS4 struct { F0 bool }
func genFn0(p0 string, p1 uint8, p2 int16) int16 {
    v5 := GenS1{F0: ((p0 <= (p0 + "foo bar")) || false), F1: 2147483647}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = (!((^(p1)) != (p1 * (p1 * (p1 >> ((4294967295) & 63))))))
        v8 := GenS2{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -32768
}
func genFn1(p0 int16) uint64 {
    var v9 int = 1
    fmt.Println(v9, p0)
    p0 = p0
    p0 += p0
    var v10 int = (v9 + v9)
    _ = p0
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []uint32{430, 65535, 2, 4294967295}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 1
    var v11 int64 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 uint = 442
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            var v14 string = (("hello" + "a") + "hello")
            _ = v14
        }
        v12 = (v12 << ((uint(v4)) & 63))
        for fi15 := 0; fi15 < 8; fi15++ {
            v12 -= (v12 << ((v12) & 63))
            v16 := GenS2{F0: v12}
            v1 = append(v1, 985)
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    if true {
        switch "hello" {
        case "\n\t":
            v11 = (v11 - (-(v11)))
        case "":
            if true {
                for fi17 := 0; fi17 < 16; fi17++ {
                    fmt.Println(v4, v12, v4)
                    switch ("hello" + "") {
                    case "foo bar":
                        delete(v2, 4294967295)
                    case "a":
                        fmt.Println(fi17)
                    default:
                        var v18 bool = (!true)
                        _ = v18
                    }
                    _ = fi17
                }
            } else {
                switch 32767 {
                case 305:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 int = fi19
                        _ = fi19
                        _ = v20
                    }
                case 1:
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if false {
                            fmt.Println(v11)
                            var v22 int8 = 127
                            _ = v22
                        }
                        v23 := v3[(v4 | v4)]
                        v24 := GenS2{F0: (v12 % ((v12) | 1))}
                        _ = fi21
                        _ = v23
                        _ = v24
                    }
                }
                if ((3.14159 * 1e10) != 0.1) {
                    v4 = (^(v4))
                    v25 := GenS3{F0: (v4 >> ((65535) & 63)), F1: 42, F2: 9223372036854775807}
                    _ = v25
                } else {
                    if true {
                        var v26 bool = (4294967295 <= 1)
                        if v26 {
                            for fi27 := 0; fi27 < 4; fi27++ {
                                if (0 > 42) {
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        var v28 string = ((("a" + "\n\t") + "hello") + "\n\t")
                                        _ = v28
                                    case "a":
                                        if v26 {
                                            v12 = 1
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                _ = len(v1)
                                            default:
                                                var v29 uint = (^(v12))
                                                _ = v29
                                            }
                                            for fi30 := 0; fi30 < 7; fi30++ {
                                                fmt.Println(v4, fi27, v26)
                                                if v26 {
                                                    v26 = (!false)
                                                    if true {
                                                        if v26 {
                                                            if true {
                                                                var v31 string = "foo bar"
                                                                var v32 int32 = 1
                                                                switch int(v12) {
                                                                case -9223372036854775808:
                                                                    v32 += v32
                                                                case 32767:
                                                                    v26 = ((v11 > v11) && false)
                                                                case 127:
                                                                    v31 = v31
                                                                }
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            fmt.Println(fi27)
                                                        }
                                                        switch fi27 {
                                                        case 127:
                                                            fmt.Println(v4)
                                                        }
                                                    }
                                                }
                                                _ = fi30
                                            }
                                            var v33 float64 = 3.14159
                                            _ = v33
                                        }
                                    default:
                                        _ = len(v1)
                                    }
                                    var v34 uint8 = 255
                                    for fi35 := 0; fi35 < 7; fi35++ {
                                        v36 := GenS1{F0: ((v26 && (0 <= (v34 ^ v34))) || true), F1: int32(v12)}
                                        v11 += (^(v11))
                                        fmt.Println(v12, fi35)
                                        switch fi35 {
                                        case -1:
                                            switch (fi35 * 137) {
                                            case 148:
                                                switch fi27 {
                                                case 159:
                                                    if false {
                                                        v11 -= (-(v11))
                                                        switch fi27 {
                                                        case 4294967295:
                                                            delete(v3, 229)
                                                        case 127:
                                                            if len(v1) > 0 {
                                                                v1[0] = v4
                                                            }
                                                        default:
                                                            v37 := GenS2{F0: v12}
                                                            _ = v37
                                                        }
                                                        v12 += 42
                                                    } else {
                                                        delete(v3, v4)
                                                        var v38 int = (fi27 % ((825) | 1))
                                                        v12 = 804
                                                        _ = v38
                                                    }
                                                case 128:
                                                    _ = len(v1)
                                                case 4294967295:
                                                    if (false && v26) {
                                                        var v39 int32 = 508
                                                        v40 := GenS2{F0: 4294967295}
                                                        v26 = ("a" != "foo bar")
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                }
                                            default:
                                                v3[v4] = 102
                                            }
                                        case -2147483648:
                                            switch (fi27 % ((230) | 1)) {
                                            case 492:
                                                if (1000 >= -972) {
                                                    _ = len(v3)
                                                    fmt.Println(v26, v26, v34)
                                                    if ((v34 + (v34 - (v34 + v34))) > (v34 % ((127) | 1))) {
                                                        var v41 float32 = 1.5
                                                        var v42 uint32 = (v4 % ((v4) | 1))
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        var v43 uint = (v12 >> ((v12) & 63))
                                                        v26 = (138 < 42)
                                                        if v26 {
                                                            var v44 int = (-(fi35))
                                                            var v45 uint8 = v34
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        v34 = (v34 + v34)
                                                        var v46 float64 = 3.14159
                                                        if (56 == 1) {
                                                            fmt.Println(v11, fi35)
                                                        }
                                                        if v26 {
                                                            if (65535 <= 466) {
                                                                var v47 int16 = -343
                                                                var v48 int32 = -280
                                                                if ((158 == 4294967295) && v26) {
                                                                    v11 = (v11 + (v11 >> ((581) & 63)))
                                                                } else {
                                                                    switch fi35 {
                                                                    case 2147483647:
                                                                        var v49 uint = v12
                                                                        _ = v49
                                                                    case 917:
                                                                        var v50 uint64 = 42
                                                                        _ = v50
                                                                    case 687:
                                                                        var v51 uint32 = 627
                                                                        _ = v51
                                                                    }
                                                                    fmt.Println(v48, fi27, fi27)
                                                                    fmt.Println(v34)
                                                                }
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                        }
                                                        _ = v43
                                                        _ = v46
                                                    }
                                                }
                                            }
                                        case 4294967295:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            fmt.Println(v4)
                                            if (32767 > genFn0("hello", 127, 1000)) {
                                                if v26 {
                                                    if len(v1) > 0 {
                                                        v1[0] = (v4 % (((v4 << (((^(v12))) & 63))) | 1))
                                                    }
                                                }
                                                if v26 {
                                                    var v52 int64 = (v11 * (v11 << ((uint(v34)) & 63)))
                                                    if (18446744073709551615 > 2) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(v26, v26, fi27)
                                                    }
                                                    _ = v52
                                                } else {
                                                    v4 = 2
                                                }
                                                if v26 {
                                                    v3[v4] = v4
                                                    switch (fi27 >> ((438) & 63)) {
                                                    case 1:
                                                        v53 := GenS3{F0: v4, F1: -32768, F2: -451}
                                                        _ = v53
                                                    case 4294967295:
                                                        if len(v1) > 0 {
                                                            v1[0] = v4
                                                        }
                                                    case -223:
                                                        if (("hello" >= "\n\t") || (1 >= 129)) {
                                                            if true {
                                                                v54 := GenS1{F0: false, F1: 42}
                                                                v4 = v4
                                                                var v55 uint8 = v34
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        }
                                                    }
                                                    v34 = (v34 & v34)
                                                    var v56 int8 = -1
                                                    _ = v56
                                                }
                                            }
                                        }
                                        var v57 float32 = -2.25
                                        v4 = (v4 ^ 4294967295)
                                        v58 := GenS1{F0: ((1e10 - 0.1) < -2.25), F1: -2147483648}
                                        _ = fi35
                                        _ = v36
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = v34
                                } else {
                                    var v59 int8 = 39
                                    v59 = v59
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v59
                                }
                                _ = fi27
                            }
                        } else {
                            v26 = (!true)
                            if (!v26) {
                                v11 = v11
                                v12 -= v12
                                v60 := GenS1{F0: v26, F1: 42}
                                _ = v60
                            } else {
                                v26 = (2 > 483)
                                fmt.Println(v26, v11, v4)
                                _ = len(v1)
                            }
                        }
                        if v26 {
                            v12 -= 0
                        } else {
                            for fi61 := 0; fi61 < 16; fi61++ {
                                fmt.Println(v11, v4)
                                var v62 int32 = -2147483648
                                _ = fi61
                                _ = v62
                            }
                            v4 -= 2147483648
                            v12 = (v12 + v12)
                        }
                        _ = v26
                    } else {
                        v12 = (v12 >> (((v12 & v12)) & 63))
                        v63 := GenS1{F0: true, F1: 2147483647}
                        _ = v63
                    }
                    v12 += (^(v12))
                    v64 := GenS1{F0: ("hello" <= "a"), F1: 2147483647}
                    _ = v64
                }
                v65 := GenS2{F0: 18446744073709551615}
                _ = v65
            }
        case "hello":
            v4 -= v4
        default:
            v4 = (v4 & (v4 ^ v4))
        }
        if true {
            for fi66 := 0; fi66 < 19; fi66++ {
                if (!true) {
                    v67 := GenS3{F0: 974, F1: -283, F2: (fi66 * int(v11))}
                    fmt.Println(v11)
                    _ = v67
                } else {
                    v11 = (v11 % (((v11 & v11)) | 1))
                    var v68 int32 = 1
                    _ = v68
                }
                _ = fi66
            }
        }
    } else {
        switch ("hello" + "\n\t") {
        case "a":
            var v69 string = "foo bar"
            _ = v69
        case "\n\t":
            if true {
                v11 = 9223372036854775807
            }
        case "":
            for fi70 := 0; fi70 < 6; fi70++ {
                if (fi70 < (fi70 % ((32767) | 1))) {
                    if (v4 <= 2147483647) {
                        for fi71 := 0; fi71 < 2; fi71++ {
                            var v72 int32 = 42
                            _ = fi71
                            _ = v72
                        }
                        v4 -= (v4 + v4)
                    }
                    switch (((("" + "hello") + "") + "") + "foo bar") {
                    case "\n\t":
                        for fi73 := 0; fi73 < 11; fi73++ {
                            v1 = append(v1, (v4 * 42))
                            _ = len(v1)
                            _ = fi73
                        }
                    case "":
                        v74 := GenS4{F0: true}
                        _ = v74
                    default:
                        var v75 int16 = -657
                        _ = v75
                    }
                } else {
                    for fi76 := 0; fi76 < 5; fi76++ {
                        var v77 bool = true
                        _ = fi76
                        _ = v77
                    }
                    switch fi70 {
                    case -2147483648:
                        fmt.Println(v12)
                    default:
                        v1 = append(v1, v4)
                    }
                    fmt.Println(fi70, fi70)
                }
                for fi78 := 0; fi78 < 18; fi78++ {
                    v1 = append(v1, (v4 >> (((v12 << ((4294967295) & 63))) & 63)))
                    v79 := v2[-2147483648]
                    _ = fi78
                    _ = v79
                }
                var v80 uint32 = (v4 + (^(v4)))
                _ = fi70
                _ = v80
            }
        default:
            var v81 string = ("foo bar" + ("hello" + "\n\t"))
            _ = v81
        }
    }
    if false {
        fmt.Println(v11)
        delete(v3, (v4 ^ v4))
    }
    if len(v1) > 0 {
        v1[0] = (^(v4))
    }
    var v82 int32 = -361
    if (!true) {
        var v83 int64 = (^(v11))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v84 uint8 = 42
        _ = v83
        _ = v84
    }
    if (513 < 794) {
        fmt.Println(v82, v4, v4)
    }
    for fi85 := 0; fi85 < 6; fi85++ {
        v12 += v12
        _ = len(v1)
        _ = fi85
    }
    var v86 int8 = -32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v86
}
