package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 int8; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 string; F1 string }
func genFn0(p0 int64) uint64 {
    switch "" {
    case "\n\t":
        v7 := GenS1{F0: 1}
        _ = v7
    case "":
        fmt.Println(p0)
    }
    v8 := GenS1{F0: 1000}
    p0 = (p0 ^ p0)
    p0 -= 1
    _ = p0
    _ = v8
    return 42
}
func genFn1() int {
    v9 := GenS1{F0: -1}
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int8 = -128
        _ = fi10
        _ = v11
    }
    if ((((1e10 + (((3.14159 + -2.25) + 3.14159) / float64(((float32((float64(1.5) / (((1.5 / 1.0) + (-2.25 * 0.0)) - ((3.14159 - (((3.14159 / -2.25) / ((3.14159 / 1.5) - -2.25)) - -2.25)) * 3.14159)))) - float32(float64(0.0))) - 1e10)))) - 1.5) * (1e10 + float64(0.1))) <= 3.14159) {
        if (false && true) {
            if true {
                var v12 uint8 = 128
                v12 -= v12
                v13 := GenS1{F0: 42}
                _ = v12
                _ = v13
            }
            var v14 int = 646
            v14 = v14
            _ = v14
        }
    }
    var v15 uint64 = 1
    _ = v15
    _ = v9
    return 32768
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768, 1}
    v3 := make(map[int]int)
    var v4 bool = false
    var v5 int64 = int64(v1)
    v6 := make(map[string]string)
    v16 := v3[-545]
    v16 = (v16 % ((-1) | 1))
    delete(v3, (v16 >> ((1) & 63)))
    var v17 int64 = v5
    v4 = ((638 < -310) && (-537 != 0))
    v17 -= 625
    if v4 {
        switch (^(v16)) {
        case -429:
            if v4 {
                if (((("foo bar" + ("" + "\n\t")) + ("foo bar" + "a")) + "a") == "foo bar") {
                    v18 := v6["foo bar"]
                    switch v16 {
                    case -1:
                        var v19 int = int(v1)
                        _ = v19
                    case 76:
                        v1 -= (^(v1))
                    }
                    _ = v18
                }
                switch (v16 % ((42) | 1)) {
                case 32767:
                    fmt.Println(v16, v1, v1)
                case -9223372036854775808:
                    switch ("a" + "foo bar") {
                    case "":
                        switch ("\n\t" + "a") {
                        case "":
                            if v4 {
                                if v4 {
                                    for fi20 := 0; fi20 < 8; fi20++ {
                                        if (!false) {
                                            fmt.Println(v17, v17)
                                        }
                                        _ = fi20
                                    }
                                    switch 32768 {
                                    case 77:
                                        var v21 bool = (!v4)
                                        _ = v21
                                    default:
                                        fmt.Println(v16, v17, v1)
                                    }
                                    v1 = v1
                                }
                            } else {
                                v22 := GenS2{F0: 656, F1: (((!((v17 << ((355) & 63)) <= v5)) && true) || v4), F2: (^(v1))}
                                _ = v22
                            }
                        }
                    }
                default:
                    if v4 {
                        if true {
                            if (0 != 39) {
                                if false {
                                    fmt.Println(v4)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v23 uint64 = 748
                                    _ = v23
                                }
                                var v24 uint64 = 18446744073709551615
                                _ = v24
                            }
                            v6["foo bar"] = "\n\t"
                            v25 := GenS3{F0: 255, F1: 1, F2: -945}
                            _ = v25
                        }
                    }
                }
            } else {
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(v1, v5)
                    var v27 float32 = 1e10
                    _ = fi26
                    _ = v27
                }
            }
        case 374:
            v3[v16] = (^(v16))
        case 208:
            v16 += v16
        }
    } else {
        if (0 == 127) {
            var v28 uint = 592
            if (((v4 || v4) || v4) || (1 > 32767)) {
                var v29 uint8 = 1
                _ = v29
            }
            _ = v28
        }
        switch "\n\t" {
        case "hello":
            for fi30 := 0; fi30 < 5; fi30++ {
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 uint32 = v1
                    if (!(81 > 1)) {
                        v4 = v4
                        v33 := GenS2{F0: 1, F1: true, F2: v32}
                        var v34 string = "a"
                        switch (v16 ^ (^(fi30))) {
                        case 9223372036854775807:
                            var v35 uint16 = 1
                            _ = v35
                        case 0:
                            v16 = fi30
                        }
                        if (0 < uint(v1)) {
                            v4 = v4
                            switch genFn1() {
                            case 2147483647:
                                _ = len(v6)
                            case 127:
                                v3[(fi31 * (fi30 << ((65535) & 63)))] = (v16 << ((689) & 63))
                            }
                        } else {
                            _ = len(v6)
                        }
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                _ = fi30
            }
        default:
            v36 := GenS1{F0: -178}
            _ = v36
        }
        fmt.Println(v1)
    }
    var v37 string = "hello"
    if true {
        switch v16 {
        case 1:
            var v38 float32 = 1.5
            _ = v38
        default:
            if (-2.25 > ((0.0 - 0.0) / (1.5 + (3.14159 / 1.5)))) {
                for fi39 := 0; fi39 < 1; fi39++ {
                    v1 += v1
                    switch genFn1() {
                    case 2147483647:
                        v40 := GenS4{F0: (fi39 + -31), F1: 100, F2: 42}
                        _ = v40
                    case 42:
                        var v41 int32 = 1
                        _ = v41
                    case -774:
                        delete(v6, "foo bar")
                    default:
                        var v42 uint16 = 161
                        _ = v42
                    }
                    v43 := GenS4{F0: fi39, F1: -1, F2: 32767}
                    _ = fi39
                    _ = v43
                }
                var v44 uint = 0
                _ = v44
            } else {
                v45 := GenS2{F0: 1, F1: (311 >= v16), F2: v1}
                var v46 uint32 = v1
                _ = v45
                _ = v46
            }
        }
        for fi47 := 0; fi47 < 3; fi47++ {
            v48 := GenS2{F0: 465, F1: (v1 > (v1 / ((v1) | 1))), F2: (^(v1))}
            _ = fi47
            _ = v48
        }
    }
    var v49 uint = 211
    v50 := GenS4{F0: (v16 >> (((v49 + 683)) & 63)), F1: -121, F2: 533}
    if (!(!v4)) {
        if (!(!false)) {
            if (v1 <= 630) {
                var v51 int32 = -1
                _ = v51
            } else {
                switch (v16 + genFn1()) {
                case 1:
                    fmt.Println(v37)
                case 802:
                    delete(v6, "hello")
                default:
                    delete(v6, (v37 + ("foo bar" + v37)))
                }
                v37 = ""
            }
            var v52 int = int(v49)
            switch "hello" {
            case "foo bar":
                switch v52 {
                case 32768:
                    fmt.Println(v37, v49, v52)
                }
            case "hello":
                fmt.Println(v37, v1)
            case "":
                if (int8(v5) > 63) {
                    if true {
                        for fi53 := 0; fi53 < 9; fi53++ {
                            v4 = true
                            v1 = (v1 * (v1 & 112))
                            v54 := GenS5{F0: 0}
                            _ = fi53
                            _ = v54
                        }
                        v52 = (v16 ^ -2147483648)
                    }
                    if true {
                        v55 := GenS3{F0: 128, F1: 4294967295, F2: (v16 * genFn1())}
                        _ = v55
                    }
                    fmt.Println(v4, v52, v5)
                } else {
                    var v56 int32 = -590
                    _ = v56
                }
            default:
                v57 := GenS2{F0: 0, F1: false, F2: v1}
                _ = v57
            }
            _ = v52
        }
        v58 := GenS4{F0: 833, F1: 109, F2: -588}
        v59 := GenS5{F0: (v1 + 803)}
        _ = v58
        _ = v59
    } else {
        if ((2147483647 <= 2147483647) || (1e10 != (-2.25 + (0.0 - ((((3.14159 - 3.14159) + float32(0.1)) / 1.5) / (-2.25 / 3.14159)))))) {
            v17 = (v17 << ((v49) & 63))
            v60 := GenS2{F0: 1, F1: v4, F2: v1}
            _ = v60
        } else {
            switch v37 {
            case "foo bar":
                fmt.Println(v37, v1, v49)
            case "hello":
                switch v16 {
                case 4294967295:
                    v61 := v3[v16]
                    _ = v61
                case -2147483648:
                    v17 = (^(v5))
                case 300:
                    if len(v2) > 0 {
                        v2[0] = (v16 ^ -2147483648)
                    }
                default:
                    v2 = append(v2, (v16 >> ((v49) & 63)))
                }
            }
        }
        if (v4 || (32767 == 255)) {
            if ((^(v1)) == v1) {
                v62 := v6["a"]
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 18; fi63++ {
            if true {
                v64 := GenS6{F0: 32767, F1: 0}
                v16 += (^(v16))
                _ = v64
            } else {
                v37 = v37
                if (42 < 732) {
                    fmt.Println(v49, v37, v4)
                    v16 += int(v49)
                    for fi65 := 0; fi65 < 11; fi65++ {
                        v66 := GenS3{F0: 42, F1: 4294967295, F2: (fi65 << ((uint(v1)) & 63))}
                        v1 -= 4294967295
                        var v67 uint = (v49 ^ 0)
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                } else {
                    v1 -= v1
                    v5 = (-(v5))
                    delete(v3, genFn1())
                }
                var v68 uint64 = genFn0((^(v5)))
                _ = v68
            }
            var v69 string = v37
            switch (-(v16)) {
            case 127:
                v70 := GenS5{F0: v1}
                _ = v70
            case -95:
                var v71 int16 = 1000
                _ = v71
            default:
                v72 := GenS1{F0: -1}
                _ = v72
            }
            _ = fi63
            _ = v69
        }
    }
    var v73 float32 = ((0.0 + 1.5) / 1.5)
    v74 := GenS7{F0: v37, F1: v37}
    var v75 string = v37
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v73
    _ = v74
    _ = v75
}
