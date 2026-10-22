package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 int }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint16; F1 int; F2 uint16 }
type GenS7 struct { F0 bool; F1 uint8; F2 int64 }
type GenS8 struct { F0 uint64; F1 string }
type GenS9 struct { F0 int; F1 int16; F2 bool }

func main() {
    var v1 uint8 = 255
    var v2 int8 = 42
    var v3 int16 = 42
    var v4 float64 = 0.1
    var v5 uint32 = 0
    var v6 uint = 18446744073709551615
    fmt.Println(v3, v5)
    for fi7 := 0; fi7 < 4; fi7++ {
        switch fi7 {
        case 42:
            v8 := GenS1{F0: (v6 - 2), F1: 980, F2: ((v6 == 42) && (-99 != int32(v3)))}
            _ = v8
        case 621:
            v9 := GenS1{F0: uint(v5), F1: uint64(v6), F2: (true && false)}
            _ = v9
        case 128:
            v10 := []int{-9223372036854775808, fi7, (fi7 + 32767), (fi7 / ((fi7) | 1))}
            _ = v10
        default:
            v11 := []int{}
            _ = v11
        }
        v12 := []int{}
        var v13 uint8 = v1
        _ = fi7
        _ = v12
        _ = v13
    }
    switch 127 {
    case 127:
        if ((true || true) && (v4 < (1e10 * v4))) {
            var v14 int8 = (v2 - v2)
            var v15 float64 = v4
            if (((!(!true)) || (false && false)) && true) {
                v16 := GenS1{F0: v6, F1: 296, F2: true}
                _ = v16
            } else {
                var v17 float64 = v15
                _ = v17
            }
            _ = v14
            _ = v15
        } else {
            switch 127 {
            case 2147483647:
                v3 = (^(v3))
            case 0:
                switch 4294967295 {
                case 2147483647:
                    var v18 int8 = (v2 - (v2 + -37))
                    _ = v18
                case 9223372036854775807:
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v3)
                    }
                default:
                    switch -433 {
                    case 501:
                        var v19 int8 = 100
                        _ = v19
                    default:
                        v20 := GenS2{F0: (v3 ^ v3), F1: -589, F2: "a"}
                        _ = v20
                    }
                }
            case 127:
                v3 += v3
            default:
                fmt.Println(v1, v3)
            }
            var v21 string = "a"
            switch 243 {
            case -1:
                v5 = (v5 * v5)
            case 4294967295:
                var v22 uint64 = 18446744073709551615
                _ = v22
            }
            _ = v21
        }
    case -617:
        v3 = v3
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v2 = v2
        v24 := GenS3{F0: (v5 | 4294967295)}
        if false {
            var v25 int = fi23
            fmt.Println(v6, fi23)
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    v2 += -1
    v1 = v1
    for fi26 := 0; fi26 < 4; fi26++ {
        v2 = v2
        switch (fi26 & 1) {
        case 42:
            var v27 uint8 = (v1 >> ((v6) & 63))
            _ = v27
        case 0:
            v2 = (v2 * (v2 ^ -1))
        default:
            var v28 uint64 = 1
            _ = v28
        }
        _ = fi26
    }
    if (true && false) {
        for fi29 := 0; fi29 < 14; fi29++ {
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := []int{fi30, fi29, (fi30 / ((fi29) | 1))}
                v31 = append(v31, (fi30 >> (((v6 - 0)) & 63)))
                if false {
                    v2 += 100
                }
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        v3 = int16(v1)
        if (((float64(0.0) / float64((((0.0 - (-2.25 / (-2.25 * (1.5 - 0.1)))) * 0.0) + 0.1))) / 0.1) >= 3.14159) {
            fmt.Println(v4, v5)
            v2 = (v2 / ((v2) | 1))
            v32 := GenS2{F0: 42, F1: 127, F2: "hello"}
            _ = v32
        } else {
            v33 := GenS4{F0: 0, F1: -1, F2: -2147483648}
            v34 := []int{-497, -828}
            if false {
                if true {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        for fi36 := 0; fi36 < 3; fi36++ {
                            var v37 uint32 = v5
                            v37 = v5
                            fmt.Println(fi36, v37)
                            _ = fi36
                            _ = v37
                        }
                        v38 := GenS5{F0: 790, F1: 302}
                        _ = fi35
                        _ = v38
                    }
                    v39 := GenS1{F0: (v6 >> ((18446744073709551615) & 63)), F1: 18446744073709551615, F2: false}
                    var v40 uint16 = 42
                    _ = v39
                    _ = v40
                }
            } else {
                v34 = append(v34, -631)
                for fi41 := 0; fi41 < 17; fi41++ {
                    if true {
                        v42 := GenS6{F0: 0, F1: fi41, F2: 675}
                        _ = v42
                    } else {
                        if (!(-2147483648 >= 454)) {
                            switch -437 {
                            case 968:
                                for fi43 := 0; fi43 < 1; fi43++ {
                                    _ = len(v34)
                                    if (true && true) {
                                        switch (("" + ("" + "foo bar")) + ("hello" + "\n\t")) {
                                        case "a":
                                            var v44 uint64 = 2
                                            _ = v44
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "foo bar":
                                                v45 := GenS6{F0: 943, F1: fi43, F2: 255}
                                                _ = v45
                                            case "\n\t":
                                                fmt.Println(v6, v5, fi43)
                                            default:
                                                var v46 int32 = 0
                                                _ = v46
                                            }
                                        case "\n\t":
                                            v47 := GenS4{F0: 9223372036854775807, F1: 0, F2: (fi41 | (fi41 << (((v6 + v6)) & 63)))}
                                            _ = v47
                                        default:
                                            var v48 int = fi41
                                            _ = v48
                                        }
                                    }
                                    v49 := GenS3{F0: (v5 / ((v5) | 1))}
                                    _ = fi43
                                    _ = v49
                                }
                            case -636:
                                var v50 string = ""
                                _ = v50
                            default:
                                var v51 uint64 = 792
                                _ = v51
                            }
                            v52 := GenS7{F0: true, F1: v1, F2: 9223372036854775807}
                            if len(v34) > 0 {
                                v34[0] = 42
                            }
                            _ = v52
                        } else {
                            if len(v34) > 0 {
                                v34[0] = 127
                            }
                            switch ((("" + "\n\t") + "foo bar") + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                _ = len(v34)
                            case "a":
                                fmt.Println(fi41, fi41)
                            case "foo bar":
                                v2 = (^(v2))
                            default:
                                switch -811 {
                                case 32768:
                                    switch ("foo bar" + ("foo bar" + "\n\t")) {
                                    case "\n\t":
                                        var v53 bool = ((-170 > 1) || ((("" + "") + (("foo bar" + "\n\t") + ("\n\t" + ("a" + "hello")))) >= "hello"))
                                        _ = v53
                                    case "a":
                                        for fi54 := 0; fi54 < 18; fi54++ {
                                            _ = len(v34)
                                            var v55 int64 = -1
                                            if len(v34) >= 2 {
                                                v34 = v34[1:len(v34)]
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                    case "hello":
                                        var v56 bool = true
                                        _ = v56
                                    default:
                                        for fi57 := 0; fi57 < 14; fi57++ {
                                            var v58 int = (fi57 % ((fi57) | 1))
                                            v58 -= (-(fi41))
                                            _ = fi57
                                            _ = v58
                                        }
                                    }
                                case 2147483647:
                                    if ((v5 % (((v5 & v5)) | 1)) <= v5) {
                                        for fi59 := 0; fi59 < 20; fi59++ {
                                            v60 := GenS3{F0: v5}
                                            switch "hello" {
                                            case "hello":
                                                v34 = append(v34, -1)
                                            case "a":
                                                v61 := GenS4{F0: 1, F1: 42, F2: -943}
                                                _ = v61
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                        var v62 float32 = 1.5
                                        fmt.Println(v1, v1, v3)
                                        _ = v62
                                    }
                                case 42:
                                    var v63 uint16 = 623
                                    _ = v63
                                default:
                                    if ((((-(v3)) == v3) && (!true)) && ("\n\t" >= "\n\t")) {
                                        var v64 int16 = v3
                                        for fi65 := 0; fi65 < 6; fi65++ {
                                            v1 -= v1
                                            _ = fi65
                                        }
                                        _ = v64
                                    }
                                }
                            }
                            v66 := GenS8{F0: 801, F1: "hello"}
                            _ = v66
                        }
                    }
                    for fi67 := 0; fi67 < 19; fi67++ {
                        if ((3.14159 == float32((0.0 / 3.14159))) && false) {
                            switch "foo bar" {
                            case "hello":
                                v3 = -32768
                            case "foo bar":
                                v3 = v3
                            }
                        }
                        _ = fi67
                    }
                    v3 += 32767
                    _ = fi41
                }
                var v68 int16 = v3
                _ = v68
            }
            _ = v33
            _ = v34
        }
    }
    var v69 int8 = (v2 * (v2 * (v2 | 86)))
    v70 := GenS9{F0: -646, F1: 42, F2: ((v3 ^ (v3 << ((v6) & 63))) >= 42)}
    v71 := []int{}
    if (v5 >= v5) {
        v4 = v4
        v72 := GenS9{F0: 9223372036854775807, F1: (v3 * (v3 ^ -32768)), F2: ((float64(1.5) / 1.5) <= v4)}
        fmt.Println(v69, v1, v5)
        _ = v72
    } else {
        switch 165 {
        case 0:
            if (v69 >= 1) {
                var v73 bool = true
                _ = v73
            }
        case 4294967295:
            v71 = append(v71, int(v1))
        case 2147483647:
            if len(v71) > 0 {
                v71[0] = 127
            }
        default:
            v2 = v69
        }
        if (1 > 42) {
            for fi74 := 0; fi74 < 16; fi74++ {
                var v75 uint16 = 802
                v75 = (^(v75))
                v76 := GenS9{F0: fi74, F1: v3, F2: true}
                _ = fi74
                _ = v75
                _ = v76
            }
            fmt.Println(v3, v2)
            var v77 int64 = 316
            _ = v77
        } else {
            for fi78 := 0; fi78 < 5; fi78++ {
                fmt.Println(v5, v5)
                v79 := GenS9{F0: 9223372036854775807, F1: (v3 & (^(v3))), F2: (true || true)}
                _ = fi78
                _ = v79
            }
            fmt.Println(v1, v6, v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
}
