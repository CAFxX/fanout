package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
func genFn0() int8 {
    v6 := []int{}
    for fi7 := 0; fi7 < 14; fi7++ {
        if len(v6) > 0 {
            v6[0] = fi7
        }
        var v8 uint8 = 233
        v9 := GenS1{F0: 42, F1: 586, F2: ""}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ("\n\t" + "a") {
    case "a":
        var v10 uint = 18446744073709551615
        _ = v10
    case "foo bar":
        var v11 bool = true
        _ = v11
    default:
        var v12 uint8 = 211
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            v14 := GenS2{F0: 276, F1: -1, F2: 0}
            _ = v14
        } else {
            v15 := GenS3{F0: 0}
            var v16 int8 = 0
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    _ = v6
    return -1
}
func genFn1(p0 int, p1 int) string {
    p1 = 4294967295
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 bool = (-1 == genFn0())
            _ = fi18
            _ = v19
        }
        fmt.Println(p1)
        _ = fi17
    }
    p0 -= p1
    for fi20 := 0; fi20 < 6; fi20++ {
        if true {
            for fi21 := 0; fi21 < 14; fi21++ {
                if true {
                    if (912 != 1) {
                        if true {
                            p0 = (p0 | (p1 / ((p1) | 1)))
                            fmt.Println(p0, p0, p0)
                        } else {
                            v22 := []int{p0}
                            p1 += (p0 / (((fi20 * 2147483647)) | 1))
                            _ = v22
                        }
                    } else {
                        v23 := GenS2{F0: 696, F1: 1, F2: -1}
                        var v24 uint64 = 78
                        switch (fi21 << ((65535) & 63)) {
                        case -419:
                            p1 = (fi21 * (p0 | (p0 % (((fi20 - 127)) | 1))))
                        default:
                            switch p1 {
                            case -948:
                                v25 := GenS2{F0: 65535, F1: 0, F2: 489}
                                _ = v25
                            case 4294967295:
                                v26 := GenS2{F0: 1, F1: 1, F2: -1}
                                var v27 string = ""
                                fmt.Println(fi20, v24, p1)
                                _ = v26
                                _ = v27
                            default:
                                var v28 uint16 = 764
                                _ = v28
                            }
                        }
                        v29 := []int{}
                        var v30 int16 = 0
                        _ = len(v29)
                        if len(v29) > 0 {
                            v29[0] = 9223372036854775807
                        }
                        v30 -= v30
                        _ = len(v29)
                        switch (fi20 >> ((65535) & 63)) {
                        case 1:
                            v30 += v30
                        case 0:
                            if (1 < 1) {
                                v29 = append(v29, (p0 ^ (p0 * fi21)))
                                p0 = (^(fi20))
                                fmt.Println(v24)
                            }
                        case 42:
                            v24 += 939
                        default:
                            fmt.Println(fi20, p0, v30)
                        }
                        _ = v23
                        _ = v24
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi21
            }
            v31 := []int{fi20, fi20}
            _ = v31
        } else {
            p0 = p1
            for fi32 := 0; fi32 < 19; fi32++ {
                p1 = fi32
                var v33 int64 = 1
                v33 += v33
                _ = fi32
                _ = v33
            }
            for fi34 := 0; fi34 < 20; fi34++ {
                fmt.Println(p0, p1, fi34)
                v35 := []int{4294967295}
                if true {
                    v36 := GenS3{F0: -900}
                    var v37 int8 = genFn0()
                    var v38 int16 = 1
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi34
                _ = v35
            }
        }
        _ = fi20
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    v1 := []int{0, 9223372036854775807, 685, 4294967295}
    var v2 int = 42
    var v3 uint16 = 617
    var v4 uint8 = 42
    v5 := make(map[uint32]bool)
    if (false || false) {
        var v39 bool = true
        if false {
            for fi40 := 0; fi40 < 14; fi40++ {
                if v39 {
                    _ = len(v1)
                } else {
                    if (34 > 903) {
                        for fi41 := 0; fi41 < 5; fi41++ {
                            switch ("foo bar" + ("a" + "hello")) {
                            case "a":
                                fmt.Println(fi41)
                            case "foo bar":
                                v42 := GenS4{F0: 2147483648}
                                _ = v42
                            default:
                                var v43 uint = 255
                                _ = v43
                            }
                            _ = fi41
                        }
                    }
                }
                var v44 float32 = ((1e10 * 1e10) / 3.14159)
                if (515 != 517) {
                    v39 = (-109 < -243)
                    for fi45 := 0; fi45 < 1; fi45++ {
                        var v46 float32 = v44
                        var v47 string = "a"
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    var v48 float32 = 1e10
                    _ = v48
                } else {
                    _ = len(v1)
                }
                _ = fi40
                _ = v44
            }
        } else {
            v39 = (!true)
            var v49 int32 = -1
            _ = v49
        }
        _ = v39
    } else {
        var v50 float64 = 1.5
        if ((v2 & (v2 / ((v2) | 1))) != v2) {
            v51 := GenS3{F0: 123}
            switch v2 {
            case -2147483648:
                v2 += v2
            case -765:
                var v52 uint64 = 42
                _ = v52
            case -543:
                fmt.Println(v4, v2, v3)
            default:
                v53 := GenS5{F0: 2147483647}
                _ = v53
            }
            if true {
                _ = len(v5)
                v2 += v2
                fmt.Println(v50, v50)
            } else {
                v4 += (v4 - v4)
            }
            _ = v51
        }
        v5[222] = true
        _ = v50
    }
    fmt.Println(v4, v4, v3)
    fmt.Println(v4, v3, v3)
    switch (-(v2)) {
    case 62:
        v2 = (^(v2))
    case 811:
        v54 := GenS2{F0: 4294967295, F1: 691, F2: 0}
        _ = v54
    }
    if (1 < (^(v3))) {
        var v55 uint = 2
        _ = v55
    }
    switch v2 {
    case -1:
        if (true || false) {
            v5[2147483648] = (32767 != -32768)
        } else {
            v3 = 180
            var v56 uint = 255
            var v57 float64 = 3.14159
            _ = v56
            _ = v57
        }
    }
    v4 = (^(v4))
    _ = len(v5)
    switch ("a" + "") {
    case "hello":
        var v58 float32 = float32(-2.25)
        _ = v58
    default:
        var v59 bool = true
        _ = v59
    }
    v60 := GenS2{F0: 2147483647, F1: -2147483648, F2: -409}
    for fi61 := 0; fi61 < 19; fi61++ {
        if false {
            v2 += 4294967295
            _ = len(v5)
        }
        switch (-(fi61)) {
        case -813:
            if (uint64(v3) <= 42) {
                if false {
                    v62 := GenS5{F0: -1}
                    for fi63 := 0; fi63 < 2; fi63++ {
                        var v64 int8 = genFn0()
                        v4 = (v4 - 1)
                        v2 -= (fi61 % (((v2 ^ -1)) | 1))
                        _ = fi63
                        _ = v64
                    }
                    if (246 != 1) {
                        switch (v2 | (v2 | (fi61 - (v2 % (((fi61 / (((^(v2))) | 1))) | 1))))) {
                        case -9223372036854775808:
                            fmt.Println(v4, v4)
                        case 0:
                            v4 = (v4 | 128)
                        case 42:
                            if (!(!((42 <= -128) || false))) {
                                var v65 int = (fi61 - (fi61 * 2147483647))
                                for fi66 := 0; fi66 < 19; fi66++ {
                                    switch (fi66 & (^(fi66))) {
                                    case 143:
                                        v4 -= (^(v4))
                                    }
                                    _ = fi66
                                }
                                v4 += (v4 ^ (v4 + 114))
                                _ = v65
                            }
                        }
                    } else {
                        var v67 uint32 = 772
                        v68 := GenS2{F0: (^(v67)), F1: -2147483648, F2: 0}
                        _ = v67
                        _ = v68
                    }
                    _ = v62
                }
            } else {
                v69 := GenS2{F0: 210, F1: 2147483647, F2: 9223372036854775807}
                fmt.Println(v4, v3)
                _ = v69
            }
        case 0:
            v3 = 255
        case 128:
            v70 := GenS4{F0: 4294967295}
            _ = v70
        default:
            v71 := v5[856]
            _ = v71
        }
        switch ("" + (("a" + ("hello" + "foo bar")) + "a")) {
        case "hello":
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 string = ("" + ("a" + "\n\t"))
                _ = fi72
                _ = v73
            }
        default:
            v3 = 1
        }
        _ = fi61
    }
    for fi74 := 0; fi74 < 20; fi74++ {
        _ = len(v5)
        if (true || false) {
            v2 += (v2 ^ v2)
        }
        _ = fi74
    }
    for fi75 := 0; fi75 < 4; fi75++ {
        fmt.Println(v2)
        if true {
            v1 = append(v1, 1000)
        } else {
            delete(v5, 0)
            var v76 int32 = -1
            _ = v76
        }
        _ = fi75
    }
    for fi77 := 0; fi77 < 10; fi77++ {
        var v78 int8 = 1
        _ = fi77
        _ = v78
    }
    v3 -= (v3 ^ v3)
    var v79 float64 = float64(1e10)
    v80 := GenS5{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v79
    _ = v80
}
