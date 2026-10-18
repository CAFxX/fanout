package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint; F1 uint64 }
func genFn0(p0 int64, p1 uint8) bool {
    for fi4 := 0; fi4 < 16; fi4++ {
        p1 += (p1 % (((^(p1))) | 1))
        _ = fi4
    }
    p1 += (^(p1))
    p0 = (-(p0))
    if (4294967295 <= -2147483648) {
        p1 = (^(p1))
        v5 := GenS1{F0: (!(false || ((false && true) && true))), F1: 573, F2: 814}
        v6 := []int{32767, 657, -9223372036854775808}
        _ = v5
        _ = v6
    } else {
        var v7 uint64 = uint64(p1)
        var v8 int64 = (p0 >> ((2) & 63))
        v9 := []int{0}
        _ = v7
        _ = v8
        _ = v9
    }
    if true {
        p0 -= (p0 | p0)
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 string, p1 uint) uint64 {
    switch (("foo bar" + ((p0 + (p0 + "")) + (p0 + ("a" + p0)))) + p0) {
    case "a":
        var v10 uint32 = 42
        _ = v10
    case "hello":
        var v11 int64 = 0
        _ = v11
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: true, F1: (p1 * (p1 ^ (p1 | (p1 & 998)))), F2: 1}
            p1 = (p1 & p1)
            p0 = ("" + p0)
            _ = fi12
            _ = v13
        }
    }
    p1 += (p1 ^ (p1 * p1))
    var v14 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v14
    return v14
}

func main() {
    var v1 uint32 = 2
    var v2 int16 = 1000
    var v3 uint = 1
    if false {
        var v15 bool = (65535 <= uint16(v3))
        v15 = (!(!true))
        switch 2147483647 {
        case -9223372036854775808:
            fmt.Println(v3)
        case 127:
            if (1 < 0) {
                fmt.Println(v3)
                if (32767 == -1) {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v2 = (v2 >> ((v3) & 63))
                        _ = fi16
                    }
                    v17 := GenS2{F0: 37}
                    _ = v17
                }
                v18 := GenS1{F0: false, F1: (v3 ^ (v3 * 962)), F2: 2}
                _ = v18
            } else {
                fmt.Println(v1, v2)
                v19 := GenS2{F0: -41}
                var v20 int16 = 900
                _ = v19
                _ = v20
            }
        case 919:
            var v21 float64 = (1.5 / -2.25)
            _ = v21
        default:
            fmt.Println(v3, v3)
        }
        _ = v15
    }
    fmt.Println(v3)
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := []int{fi22}
        for fi24 := 0; fi24 < 11; fi24++ {
            if (0.1 >= (1.5 / 3.14159)) {
                var v25 string = "\n\t"
                var v26 uint64 = 2
                _ = v25
                _ = v26
            }
            fmt.Println(fi22, fi22)
            if len(v23) > 0 {
                v23[0] = (-(fi24))
            }
            v27 := GenS3{F0: -876, F1: (fi24 | fi22)}
            if (!false) {
                if false {
                    if true {
                        var v28 bool = true
                        _ = v28
                    }
                    var v29 uint8 = 2
                    var v30 float32 = (((0.1 + (1.5 - 1e10)) - -2.25) / 0.1)
                    _ = v29
                    _ = v30
                } else {
                    if (float64((3.14159 * 1e10)) == 1.5) {
                        switch "" {
                        case "\n\t":
                            var v31 uint8 = 157
                            _ = v31
                        case "foo bar":
                            v32 := GenS1{F0: false, F1: v3, F2: 2}
                            _ = v32
                        default:
                            if (true && true) {
                                v3 += v3
                            } else {
                                if true {
                                    v2 += v2
                                    _ = len(v23)
                                }
                                var v33 int8 = 64
                                v34 := GenS2{F0: (v33 << ((v3) & 63))}
                                _ = v33
                                _ = v34
                            }
                        }
                        v1 = v1
                        var v35 uint16 = 258
                        _ = v35
                    } else {
                        if (false || (!true)) {
                            var v36 int16 = (v2 + 1000)
                            if false {
                                var v37 bool = (!true)
                                _ = v37
                            }
                            _ = v36
                        } else {
                            v1 -= 583
                            if false {
                                v38 := GenS2{F0: int8(v1)}
                                v39 := GenS4{F0: false}
                                var v40 uint8 = 128
                                switch "hello" {
                                case "":
                                    switch (-(fi24)) {
                                    case 268:
                                        var v41 float32 = -2.25
                                        _ = v41
                                    default:
                                        var v42 uint32 = (v1 * (v1 * (^(v1))))
                                        _ = v42
                                    }
                                case "\n\t":
                                    _ = len(v23)
                                }
                                var v43 uint32 = (v1 ^ 2147483647)
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v43
                            }
                            v44 := GenS2{F0: 100}
                            var v45 bool = true
                            _ = v44
                            _ = v45
                        }
                    }
                    var v46 uint64 = 608
                    fmt.Println(v2, fi22)
                    _ = v46
                }
            }
            if (!((v1 / (((^(v1))) | 1)) < (v1 / (((v1 % (((v1 % (((v1 ^ (v1 / ((v1) | 1)))) | 1))) | 1))) | 1)))) {
                var v47 uint64 = 840
                if true {
                    var v48 uint16 = 1
                    v48 -= (v48 | v48)
                    switch (fi24 / (((^(fi22))) | 1)) {
                    case 2147483647:
                        if (((v47 ^ v47) >= v47) || false) {
                            v1 = 2
                        } else {
                            var v49 int = fi22
                            v49 = (-(fi22))
                            v2 -= (v2 + v2)
                            _ = v49
                        }
                    case -2147483648:
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                    case -678:
                        switch 128 {
                        case 989:
                            v47 = (v47 + (^(v47)))
                        case 2147483647:
                            v47 += 1
                        case 196:
                            var v50 int16 = v2
                            _ = v50
                        }
                    default:
                        var v51 int8 = int8(v1)
                        if true {
                            var v52 string = "hello"
                            var v53 uint8 = 254
                            v47 += (v47 | genFn1(v52, uint(v53)))
                            _ = v52
                            _ = v53
                        } else {
                            fmt.Println(v3, fi24, v1)
                            v23 = append(v23, -255)
                        }
                        v54 := GenS5{F0: fi22}
                        _ = v51
                        _ = v54
                    }
                    switch -9223372036854775808 {
                    case -1:
                        var v55 int64 = -9223372036854775808
                        _ = v55
                    case 782:
                        v1 = (^(v1))
                    case 32768:
                        var v56 string = ""
                        _ = v56
                    default:
                        fmt.Println(v48, v3)
                    }
                    _ = v48
                } else {
                    switch 204 {
                    case 2147483647:
                        if true {
                            var v57 bool = false
                            var v58 uint8 = 60
                            v57 = (!v57)
                            _ = v57
                            _ = v58
                        }
                    }
                    v59 := GenS6{F0: ("foo bar" + "hello"), F1: 128}
                    _ = v59
                }
                _ = v47
            }
            if false {
                switch fi22 {
                case 698:
                    v1 -= 2147483647
                case -1:
                    switch (fi22 | (^(fi22))) {
                    case -1:
                        v3 -= 42
                    case 9223372036854775807:
                        v1 += v1
                    case -9223372036854775808:
                        v1 += (v1 ^ v1)
                    default:
                        var v60 float64 = (3.14159 + 3.14159)
                        _ = v60
                    }
                case 279:
                    if len(v23) > 0 {
                        v23[0] = -2147483648
                    }
                }
                _ = len(v23)
                fmt.Println(fi24, fi22, v2)
            }
            _ = fi24
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    v61 := []int{1}
    var v62 int32 = -1
    v1 -= 2147483648
    for fi63 := 0; fi63 < 2; fi63++ {
        for fi64 := 0; fi64 < 2; fi64++ {
            fmt.Println(v3, fi63, v62)
            v2 += (v2 << ((0) & 63))
            _ = fi64
        }
        _ = fi63
    }
    var v65 uint64 = 871
    for fi66 := 0; fi66 < 5; fi66++ {
        v67 := GenS3{F0: v62, F1: (fi66 - -2147483648)}
        v68 := GenS3{F0: v62, F1: fi66}
        for fi69 := 0; fi69 < 14; fi69++ {
            switch "hello" {
            case "":
                v65 -= v65
            case "\n\t":
                if false {
                    var v70 uint16 = 42
                    switch ("" + (((("" + ((("" + "\n\t") + ("a" + (("foo bar" + ((((("\n\t" + "hello") + "hello") + ("foo bar" + ("a" + ("" + ("" + (("\n\t" + "foo bar") + "\n\t")))))) + ((("\n\t" + (("\n\t" + "hello") + "foo bar")) + "\n\t") + "a")) + ("hello" + "foo bar"))) + ("a" + ("" + "hello"))))) + ("" + "a"))) + "") + "foo bar") + "")) {
                    case "":
                        v71 := GenS2{F0: 1}
                        _ = v71
                    case "foo bar":
                        fmt.Println(v2, v65)
                    }
                    _ = v70
                }
            case "a":
                v1 = (v1 | v1)
            default:
                switch (fi69 ^ 775) {
                case -9223372036854775808:
                    v1 += 870
                case -2147483648:
                    switch fi69 {
                    case 1:
                        var v72 uint32 = v1
                        _ = v72
                    case 128:
                        var v73 bool = (!true)
                        _ = v73
                    case 592:
                        v62 -= (v62 / ((v62) | 1))
                    default:
                        var v74 uint16 = 1
                        _ = v74
                    }
                default:
                    var v75 float32 = 0.1
                    _ = v75
                }
            }
            v65 = (v65 << (((v3 * (v3 / (((v3 << (((v3 * (v3 ^ v3))) & 63))) | 1)))) & 63))
            _ = fi69
        }
        _ = fi66
        _ = v67
        _ = v68
    }
    for fi76 := 0; fi76 < 17; fi76++ {
        var v77 int = fi76
        _ = fi76
        _ = v77
    }
    for fi78 := 0; fi78 < 5; fi78++ {
        var v79 float32 = 1e10
        var v80 string = ("a" + "a")
        for fi81 := 0; fi81 < 4; fi81++ {
            if (true || (false || (false || (((1e10 + ((v79 / v79) * 0.1)) + (0.1 / float32(3.14159))) != v79)))) {
                v61 = append(v61, (fi81 << (((v3 >> (((v3 | v3)) & 63))) & 63)))
                var v82 int8 = int8(v3)
                v83 := GenS4{F0: (true || (int(v65) >= fi78))}
                _ = v82
                _ = v83
            }
            var v84 int = 127
            _ = fi81
            _ = v84
        }
        _ = fi78
        _ = v79
        _ = v80
    }
    var v85 int = 393
    _ = len(v61)
    if true {
        if true {
            var v86 uint = (v3 << ((v3) & 63))
            _ = v86
        }
    } else {
        v2 = 25
    }
    var v87 bool = (!(false && true))
    v65 = 42
    if false {
        fmt.Println(v3, v3)
    } else {
        for fi88 := 0; fi88 < 18; fi88++ {
            v89 := GenS7{F0: v3, F1: (v65 % (((v65 | 1)) | 1))}
            if len(v61) > 0 {
                v61[0] = -161
            }
            _ = fi88
            _ = v89
        }
        switch (("foo bar" + "\n\t") + ("a" + ((("a" + "a") + "") + "hello"))) {
        case "hello":
            var v90 float32 = 3.14159
            _ = v90
        default:
            switch 128 {
            case 127:
                switch (v85 ^ v85) {
                case -617:
                    v87 = v87
                case 42:
                    v91 := GenS1{F0: (v2 <= v2), F1: (v3 | (v3 & (^(v3)))), F2: 32767}
                    _ = v91
                case 4294967295:
                    var v92 int = int(v62)
                    _ = v92
                default:
                    for fi93 := 0; fi93 < 11; fi93++ {
                        v61 = append(v61, fi93)
                        _ = fi93
                    }
                }
            case 211:
                fmt.Println(v85, v85)
            }
        }
        if len(v61) >= 2 {
            v61 = v61[1:len(v61)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v65
    _ = v85
    _ = v87
}
