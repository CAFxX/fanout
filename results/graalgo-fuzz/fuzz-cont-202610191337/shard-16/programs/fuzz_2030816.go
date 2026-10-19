package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS13 struct { F0 uint; F1 int }
func genFn0(p0 bool, p1 int) uint32 {
    p1 += (-(p1))
    if p0 {
        for fi4 := 0; fi4 < 13; fi4++ {
            p1 -= fi4
            _ = fi4
        }
        if false {
            fmt.Println(p0)
        } else {
            if p0 {
                fmt.Println(p1, p1)
                var v5 int64 = 9223372036854775807
                v6 := GenS1{F0: 107, F1: 1}
                _ = v5
                _ = v6
            }
            var v7 uint = 42
            _ = v7
        }
        var v8 uint16 = 1
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint) uint16 {
    switch 32768 {
    case -1:
        var v9 int64 = 9223372036854775807
        _ = v9
    case -334:
        var v10 int64 = -1
        _ = v10
    case 9223372036854775807:
        fmt.Println(p0)
    }
    p0 = (p0 - (p0 & (p0 / ((p0) | 1))))
    if (-52 > 123) {
        v11 := GenS2{F0: "", F1: 1}
        v12 := []int{9223372036854775807, 9223372036854775807, 32767}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v11
        _ = v12
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int8 = 1
    var v2 uint = 268
    v3 := []string{}
    var v13 float32 = 1e10
    v2 = (v2 / ((129) | 1))
    v14 := GenS1{F0: 9223372036854775807, F1: 32767}
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            switch "" {
            case "foo bar":
                v2 += (v2 << ((4294967295) & 63))
            case "\n\t":
                var v16 bool = (fi15 < (fi15 * -214))
                _ = v16
            case "":
                for fi17 := 0; fi17 < 11; fi17++ {
                    v1 = v1
                    var v18 string = "\n\t"
                    v19 := GenS2{F0: ((v18 + v18) + v18), F1: int8(fi15)}
                    if false {
                        var v20 bool = (v13 > v13)
                        _ = v20
                    } else {
                        switch 127 {
                        case 32768:
                            if false {
                                switch v18 {
                                case "hello":
                                    switch fi15 {
                                    case 4294967295:
                                        v21 := GenS3{F0: (fi17 - fi15), F1: 0}
                                        _ = v21
                                    case 0:
                                        var v22 int = fi17
                                        var v23 uint16 = 362
                                        _ = v22
                                        _ = v23
                                    }
                                    v24 := GenS4{F0: fi17, F1: 518}
                                    v25 := GenS1{F0: -464, F1: -1}
                                    _ = v24
                                    _ = v25
                                case "":
                                    var v26 int16 = 1000
                                    _ = v26
                                case "a":
                                    if (2 <= 2147483647) {
                                        v27 := GenS5{F0: 32767, F1: genFn1((v2 + (v2 >> ((v2) & 63))))}
                                        fmt.Println(v18, fi17)
                                        v28 := GenS6{F0: 935, F1: 42}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                v29 := GenS5{F0: 42, F1: 633}
                                fmt.Println(v13)
                                _ = v29
                            }
                            v30 := GenS7{F0: uint8(v2)}
                            var v31 bool = (!(250 > 42))
                            _ = v30
                            _ = v31
                        case 0:
                            if false {
                                v2 = 415
                                var v32 uint64 = 2
                                v33 := GenS8{F0: (fi15 ^ -681), F1: 65535, F2: "foo bar"}
                                var v34 int16 = -370
                                v35 := GenS1{F0: 964, F1: 852}
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        if (false || false) {
                            v13 = v13
                        } else {
                            if true {
                                v36 := GenS9{F0: v2}
                                var v37 uint64 = 152
                                _ = v36
                                _ = v37
                            }
                        }
                        switch v18 {
                        case "foo bar":
                            v38 := GenS10{F0: 50}
                            _ = v38
                        default:
                            switch "a" {
                            case "":
                                switch (fi15 / (((fi17 - (fi15 ^ -2147483648))) | 1)) {
                                case 1:
                                    v18 = "a"
                                case 140:
                                    switch fi15 {
                                    case 526:
                                        var v39 uint8 = 5
                                        _ = v39
                                    }
                                }
                            case "\n\t":
                                var v40 bool = false
                                _ = v40
                            }
                        }
                    }
                    v41 := GenS10{F0: 1000}
                    _ = len(v3)
                    v42 := GenS1{F0: 9223372036854775807, F1: -629}
                    if (42 != 51) {
                        switch "foo bar" {
                        case "hello":
                            v18 = v18
                        }
                    } else {
                        var v43 uint64 = 0
                        if len(v3) > 0 {
                            v3[0] = v18
                        }
                        var v44 uint32 = 42
                        _ = v43
                        _ = v44
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v41
                    _ = v42
                }
            default:
                if false {
                    for fi45 := 0; fi45 < 7; fi45++ {
                        _ = len(v3)
                        _ = fi45
                    }
                    if (!(-2147483648 <= -1)) {
                        v46 := GenS11{F0: 0, F1: 1}
                        var v47 uint = v2
                        _ = v46
                        _ = v47
                    } else {
                        if ((fi15 % ((4294967295) | 1)) != 1) {
                            v2 -= 65535
                        } else {
                            var v48 int = (fi15 % ((551) | 1))
                            _ = v48
                        }
                        if false {
                            fmt.Println(v13, fi15, v2)
                            for fi49 := 0; fi49 < 15; fi49++ {
                                var v50 int16 = 42
                                v50 = -32768
                                _ = fi49
                                _ = v50
                            }
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        } else {
                            var v51 int = (fi15 * 0)
                            if (false && true) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v3 = append(v3, ("hello" + (("\n\t" + "a") + "foo bar")))
                                if (2 != 4294967295) {
                                    if ((v51 < v51) || ((-2147483648 == 2147483647) || true)) {
                                        var v52 uint16 = 42
                                        _ = v52
                                    } else {
                                        var v53 int16 = 1000
                                        _ = v53
                                    }
                                } else {
                                    v13 = 1e10
                                    fmt.Println(v51, v1)
                                }
                                v51 = (fi15 % (((v51 & v51)) | 1))
                            }
                            _ = v51
                        }
                        var v54 int64 = 9223372036854775807
                        _ = v54
                    }
                    var v55 int16 = 95
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v1 -= v1
                        if false {
                            fmt.Println(v1, v1, fi15)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(fi56, fi56, v2)
                        }
                        _ = fi56
                    }
                }
            }
            v13 -= float32(3.14159)
        }
        _ = fi15
    }
    _ = len(v3)
    for fi57 := 0; fi57 < 14; fi57++ {
        if false {
            v1 -= (v1 | (v1 * 127))
        }
        fmt.Println(fi57, v1)
        switch fi57 {
        case 9223372036854775807:
            for fi58 := 0; fi58 < 1; fi58++ {
                var v59 float64 = -2.25
                v2 -= (v2 >> (((v2 >> (((v2 | (^(v2)))) & 63))) & 63))
                _ = fi58
                _ = v59
            }
        case 1:
            for fi60 := 0; fi60 < 15; fi60++ {
                fmt.Println(fi57)
                var v61 uint8 = 2
                if true {
                    var v62 int64 = 0
                    var v63 int32 = 1
                    _ = v62
                    _ = v63
                }
                fmt.Println(v61, v13, fi60)
                _ = fi60
                _ = v61
            }
        }
        _ = fi57
    }
    v64 := GenS6{F0: 0, F1: 2}
    switch ("" + "") {
    case "a":
        if false {
            for fi65 := 0; fi65 < 5; fi65++ {
                var v66 float64 = 1.5
                var v67 int8 = (v1 / ((100) | 1))
                _ = fi65
                _ = v66
                _ = v67
            }
            v1 = (v1 % ((v1) | 1))
            fmt.Println(v1)
        } else {
            var v68 float64 = 0.1
            for fi69 := 0; fi69 < 4; fi69++ {
                if false {
                    for fi70 := 0; fi70 < 4; fi70++ {
                        fmt.Println(fi70, v68, v1)
                        var v71 uint32 = 0
                        if (v2 < (v2 ^ (^(v2)))) {
                            var v72 int8 = v1
                            _ = v72
                        }
                        _ = fi70
                        _ = v71
                    }
                    var v73 string = ""
                    var v74 uint = 0
                    _ = v73
                    _ = v74
                }
                if (42 == 0) {
                    v68 = float64((0.1 + (0.1 * -2.25)))
                }
                v75 := GenS1{F0: 768, F1: -32768}
                _ = fi69
                _ = v75
            }
            _ = v68
        }
    default:
        switch "" {
        case "hello":
            if (true || false) {
                v3 = append(v3, (((("foo bar" + "foo bar") + "") + "hello") + (("\n\t" + ("a" + "\n\t")) + "hello")))
                v1 = v1
                var v76 int16 = 32767
                _ = v76
            }
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            switch 127 {
            case 32767:
                v77 := GenS3{F0: 208, F1: -1}
                _ = v77
            case 9223372036854775807:
                for fi78 := 0; fi78 < 16; fi78++ {
                    v79 := GenS12{F0: 255, F1: false, F2: 753}
                    for fi80 := 0; fi80 < 8; fi80++ {
                        v1 += (v1 + -128)
                        var v81 int8 = (-(v1))
                        v13 -= v13
                        _ = fi80
                        _ = v81
                    }
                    var v82 int64 = -9223372036854775808
                    _ = fi78
                    _ = v79
                    _ = v82
                }
            }
        default:
            fmt.Println(v13)
        }
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + (("" + "") + "a"))
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "\n\t")
    }
    var v83 uint8 = 0
    switch "\n\t" {
    case "hello":
        v84 := GenS13{F0: 1, F1: 9223372036854775807}
        _ = v84
    }
    for fi85 := 0; fi85 < 15; fi85++ {
        if (!false) {
            v3 = append(v3, "a")
            v86 := GenS7{F0: 128}
            _ = v86
        } else {
            v2 += (v2 | v2)
            if ((true && (false || true)) && ((v83 - (v83 / (((^(v83))) | 1))) < v83)) {
                for fi87 := 0; fi87 < 8; fi87++ {
                    if false {
                        var v88 uint64 = 116
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v89 uint8 = (v83 | v83)
                        _ = v88
                        _ = v89
                    }
                    v2 -= (v2 / (((v2 & (v2 & (v2 | 255)))) | 1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v2, fi87, fi85)
                    _ = fi87
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v13 += 1e10
        _ = fi85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v64
    _ = v83
}
