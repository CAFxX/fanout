package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int8; F1 int8; F2 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32 }
func genFn0(p0 uint8) int32 {
    fmt.Println(p0)
    var v6 string = "a"
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7)
        p0 = p0
        _ = fi7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v6 = "a"
        v9 := []int{(fi8 >> ((0) & 63)), 9223372036854775807}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16) int {
    fmt.Println(p0)
    v10 := GenS1{F0: (true && ((float32(1e10) + 3.14159) >= 1e10)), F1: 42, F2: (p0 & p0)}
    _ = p0
    _ = v10
    return -654
}

func main() {
    var v1 uint = 700
    var v2 int32 = 0
    v3 := make(map[string]int)
    var v4 float32 = 1e10
    v5 := []uint32{2147483647, 65535, 4294967295, 350}
    var v11 float32 = v4
    if ((v1 >> (((v1 & (v1 + v1))) & 63)) > (v1 & 42)) {
        v4 += 3.14159
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v11 -= v4
                v1 += v1
                _ = fi13
            }
            v14 := GenS1{F0: true, F1: 1, F2: 2}
            if (0 == 127) {
                v4 = float32(1.5)
            }
            _ = fi12
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            _ = len(v3)
            _ = fi15
        }
        _ = len(v5)
        fmt.Println(v1)
    }
    if (728 < 0) {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 4; fi17++ {
                v4 = 0.1
                _ = fi17
            }
            if false {
                var v18 uint16 = 0
                for fi19 := 0; fi19 < 3; fi19++ {
                    delete(v3, "a")
                    _ = fi19
                }
                _ = v18
            }
            v20 := GenS2{F0: 4294967295, F1: 2}
            _ = fi16
            _ = v20
        }
    } else {
        fmt.Println(v2, v2)
    }
    if (!(-2.25 >= v4)) {
        v21 := GenS1{F0: true, F1: 21, F2: 0}
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 float32 = v4
                    var v25 uint16 = 0
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                fmt.Println(v1)
                _ = fi22
            }
        }
        _ = v21
    } else {
        fmt.Println(v11)
        var v26 int32 = 0
        var v27 int = -86
        _ = v26
        _ = v27
    }
    _ = len(v5)
    if (32767 == 1000) {
        switch "hello" {
        case "\n\t":
            if (false || (1e10 > 1e10)) {
                fmt.Println(v11, v4, v2)
                var v28 int8 = 42
                _ = v28
            } else {
                v2 += 42
                var v29 string = ("" + ("\n\t" + "hello"))
                _ = v29
            }
        }
        switch 4294967295 {
        case 32767:
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS3{F0: 0}
                fmt.Println(v4)
                _ = fi30
                _ = v31
            }
        }
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint16 = 32767
                if true {
                    fmt.Println(v4, v4, v11)
                    if (0 <= 1) {
                        var v35 float32 = v4
                        if false {
                            var v36 uint16 = 2
                            _ = v36
                        }
                        v11 += 3.14159
                        var v37 int64 = -9223372036854775808
                        fmt.Println(v37, v1)
                        _ = v35
                        _ = v37
                    } else {
                        v38 := GenS4{F0: genFn0(255), F1: 2, F2: (fi33 & (fi33 >> ((1) & 63)))}
                        fmt.Println(fi33)
                        v39 := GenS2{F0: 23, F1: 18446744073709551615}
                        if ((("foo bar" + "") == (("foo bar" + "foo bar") + "hello")) && true) {
                            v2 = int32(v34)
                            var v40 string = ("foo bar" + ("\n\t" + "hello"))
                            if true {
                                v2 = (v2 - (v2 << (((v1 & (v1 - v1))) & 63)))
                                v1 = v1
                            }
                            _ = v40
                        } else {
                            _ = len(v5)
                            v41 := GenS5{F0: fi32, F1: (("\n\t" + "hello") + "foo bar")}
                            _ = v41
                        }
                        switch (fi33 * genFn1(65535)) {
                        case -2147483648:
                            switch ("foo bar" + "hello") {
                            case "a":
                                v2 += (^(v2))
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        var v42 int16 = 1000
                                        switch -9223372036854775808 {
                                        case -343:
                                            v2 = (v2 % ((genFn0(2)) | 1))
                                        case 0:
                                            var v43 int8 = 1
                                            if (!(true && true)) {
                                                v44 := GenS5{F0: (fi32 >> ((v1) & 63)), F1: "\n\t"}
                                                if len(v5) > 0 {
                                                    v5[0] = 410
                                                }
                                                _ = v44
                                            } else {
                                                if true {
                                                    v45 := v3[""]
                                                    switch "" {
                                                    case "hello":
                                                        var v46 int64 = -1
                                                        _ = v46
                                                    case "":
                                                        v34 = v34
                                                    }
                                                    var v47 float32 = ((v11 * ((float32(0.0) - v11) + v11)) * v4)
                                                    _ = v45
                                                    _ = v47
                                                }
                                                var v48 float32 = (1e10 + float32(-2.25))
                                                _ = v48
                                            }
                                            _ = v43
                                        }
                                        var v49 uint32 = 65535
                                        if false {
                                            if ((0 >= 2) && false) {
                                                if true {
                                                    if (false || true) {
                                                        fmt.Println(v4)
                                                        if len(v5) > 0 {
                                                            v5[0] = (v49 & v49)
                                                        }
                                                    }
                                                    if false {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v50 int32 = genFn0(2)
                                                            _ = v50
                                                        case "foo bar":
                                                            v51 := GenS1{F0: (v4 >= ((3.14159 * 3.14159) - 1e10)), F1: 79, F2: v34}
                                                            _ = v51
                                                        case "":
                                                            if (!false) {
                                                                v4 -= float32(1.5)
                                                            }
                                                        default:
                                                            delete(v3, "")
                                                        }
                                                        v5 = append(v5, 634)
                                                        v52 := GenS6{F0: (v34 << ((881) & 63)), F1: ""}
                                                        _ = v52
                                                    }
                                                } else {
                                                    var v53 int8 = 127
                                                    v5 = append(v5, v49)
                                                    _ = v53
                                                }
                                                switch fi32 {
                                                case 2147483647:
                                                    var v54 uint32 = (v49 << (((v1 ^ v1)) & 63))
                                                    _ = v54
                                                case 42:
                                                    if (float32(1.5) >= (-2.25 / v11)) {
                                                        v55 := GenS5{F0: fi32, F1: ("a" + ("foo bar" + "a"))}
                                                        v1 -= v1
                                                        v49 = (v49 % ((v49) | 1))
                                                        _ = v55
                                                    } else {
                                                        var v56 float64 = 0.0
                                                        var v57 float32 = v4
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        fmt.Println(v42, fi32)
                                                        fmt.Println(v2)
                                                        var v58 int = -1
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                }
                                            }
                                            if true {
                                                var v59 int8 = 119
                                                _ = v59
                                            } else {
                                                _ = len(v5)
                                                v4 = float32((-2.25 * -2.25))
                                            }
                                            v49 -= v49
                                        }
                                        _ = v42
                                        _ = v49
                                    } else {
                                        if false {
                                            v60 := GenS3{F0: -778}
                                            switch fi32 {
                                            case 4294967295:
                                                var v61 uint8 = 1
                                                _ = v61
                                            case 42:
                                                v34 -= (v34 << (((^(v1))) & 63))
                                            }
                                            _ = v60
                                        }
                                        v2 = (^(v2))
                                    }
                                }
                            case "\n\t":
                                var v62 float64 = (float64((v11 / float32(((0.1 * (1e10 + -2.25)) + 0.1)))) + 0.0)
                                _ = v62
                            default:
                                v63 := GenS7{F0: 114, F1: -32, F2: -1}
                                _ = v63
                            }
                        case 2147483647:
                            switch fi33 {
                            case 32768:
                                v3[""] = (fi33 + genFn1(v34))
                            case 169:
                                if (!false) {
                                    v4 = v4
                                } else {
                                    var v64 uint16 = v34
                                    v34 = (v34 + v34)
                                    var v65 uint32 = 330
                                    fmt.Println(v34, v2, v64)
                                    var v66 uint64 = uint64(v65)
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            case 32:
                                var v67 uint64 = 18446744073709551615
                                _ = v67
                            default:
                                var v68 int8 = 100
                                _ = v68
                            }
                        case 4294967295:
                            v2 = genFn0(1)
                        default:
                            v34 += (v34 & v34)
                        }
                        _ = v38
                        _ = v39
                    }
                }
                switch "a" {
                case "foo bar":
                    v69 := v3["\n\t"]
                    _ = v69
                case "\n\t":
                    switch (fi32 ^ fi33) {
                    case 1:
                        v11 -= (v4 + v4)
                    case 32768:
                        switch (("a" + "") + "") {
                        case "a":
                            var v70 int8 = 42
                            _ = v70
                        case "":
                            if true {
                                v71 := GenS8{F0: false}
                                var v72 int32 = v2
                                _ = v71
                                _ = v72
                            }
                        case "hello":
                            v2 = genFn0(255)
                        default:
                            var v73 int64 = -9223372036854775808
                            _ = v73
                        }
                    case -1:
                        v2 = v2
                    default:
                        v74 := GenS9{F0: (v2 | (-(v2)))}
                        _ = v74
                    }
                }
                _ = fi33
                _ = v34
            }
            v75 := GenS2{F0: 1, F1: 703}
            _ = fi32
            _ = v75
        }
        for fi76 := 0; fi76 < 7; fi76++ {
            switch "hello" {
            case "foo bar":
                v3[("hello" + "a")] = fi76
            }
            if true {
                for fi77 := 0; fi77 < 17; fi77++ {
                    v4 = -2.25
                    _ = fi77
                }
                for fi78 := 0; fi78 < 1; fi78++ {
                    v5 = append(v5, 808)
                    _ = fi78
                }
                v3[(("a" + ("" + "\n\t")) + (("\n\t" + "") + "hello"))] = (fi76 % ((fi76) | 1))
            }
            _ = fi76
        }
    }
    fmt.Println(v4, v1, v11)
    for fi79 := 0; fi79 < 11; fi79++ {
        var v80 uint32 = 4294967295
        _ = fi79
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
