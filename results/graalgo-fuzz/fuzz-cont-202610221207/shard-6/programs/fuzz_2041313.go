package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int64; F1 uint16; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{972, 42}
    var v3 uint = 4294967295
    var v4 int8 = 100
    _ = len(v2)
    if true {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint64 = 4294967295
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 42, F1: 42}
                if true {
                    fmt.Println(v3)
                    var v9 int8 = v4
                    switch 32767 {
                    case -2147483648:
                        var v10 uint16 = uint16(v3)
                        _ = v10
                    case -94:
                        fmt.Println(v6, fi7)
                    case 127:
                        var v11 uint32 = 4294967295
                        _ = v11
                    }
                    _ = v9
                } else {
                    var v12 uint64 = 1
                    if false {
                        v3 -= (^(v3))
                        switch "hello" {
                        case "a":
                            v3 += uint(v12)
                            switch ("foo bar" + ("a" + "hello")) {
                            case "\n\t":
                                if (!(366 != -109)) {
                                    _ = len(v1)
                                }
                            case "foo bar":
                                var v13 uint32 = 761
                                _ = v13
                            case "a":
                                v14 := GenS1{F0: 1, F1: (v4 ^ (-(v4)))}
                                _ = v14
                            default:
                                v1[""] = (-(fi7))
                                var v15 uint = v3
                                v2 = append(v2, 2)
                                _ = v15
                            }
                            var v16 float32 = 1e10
                            switch "" {
                            case "hello":
                                if false {
                                    var v17 bool = false
                                    _ = v17
                                } else {
                                    var v18 uint16 = 65535
                                    var v19 string = "foo bar"
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                fmt.Println(fi7)
                            }
                            _ = v16
                        case "":
                            var v20 uint16 = 609
                            _ = v20
                        case "hello":
                            if true {
                                var v21 uint16 = 406
                                v22 := GenS1{F0: 191, F1: v4}
                                v23 := GenS2{F0: (^(v3)), F1: -2147483648}
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                fmt.Println(fi5)
                            }
                        default:
                            var v24 int64 = 9223372036854775807
                            _ = v24
                        }
                    }
                    _ = v12
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi7
                _ = v8
            }
            if false {
                v25 := GenS3{F0: -719, F1: 65535, F2: "hello"}
                _ = v25
            }
            _ = fi5
            _ = v6
        }
        _ = len(v2)
        v3 = v3
    } else {
        v26 := GenS3{F0: 9223372036854775807, F1: 318, F2: "foo bar"}
        v3 -= v3
        _ = v26
    }
    v3 = 514
    v3 = (^(v3))
    v4 -= (v4 & v4)
    if true {
        v3 += (v3 / ((v3) | 1))
    }
    if true {
        var v27 uint = (^(v3))
        v27 = 988
        _ = v27
    } else {
        var v28 uint32 = 115
        var v29 float64 = float64(1.5)
        switch 9223372036854775807 {
        case -2147483648:
            for fi30 := 0; fi30 < 7; fi30++ {
                v3 += (^(v3))
                var v31 int8 = v4
                _ = fi30
                _ = v31
            }
        default:
            if (-736 != -2147483648) {
                switch "foo bar" {
                case "foo bar":
                    v29 = v29
                case "a":
                    v28 -= v28
                case "hello":
                    v32 := GenS3{F0: 9223372036854775807, F1: 2, F2: ""}
                    _ = v32
                }
                v33 := GenS2{F0: (^(v3)), F1: 2147483647}
                _ = v33
            } else {
                v28 -= (v28 + 4294967295)
            }
        }
        _ = v28
        _ = v29
    }
    switch "" {
    case "":
        if (((1e10 == 3.14159) || (!(false || (-20 != (v4 >> (((v3 >> ((4294967295) & 63))) & 63)))))) && ((true || (!true)) && false)) {
            fmt.Println(v4, v4)
        } else {
            var v34 uint16 = 0
            _ = len(v2)
            fmt.Println(v3, v3)
            _ = v34
        }
    case "hello":
        v1["a"] = 32767
    case "a":
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 int = fi35
            _ = fi35
            _ = v36
        }
    default:
        if (0.1 < float32(0.0)) {
            if (!((2 < 42) && (95 > 915))) {
                if (-29 != 32767) {
                    v37 := GenS1{F0: 255, F1: v4}
                    v38 := GenS1{F0: 127, F1: v4}
                    _ = v37
                    _ = v38
                } else {
                    fmt.Println(v4, v3)
                    var v39 uint32 = 2147483647
                    _ = v39
                }
                for fi40 := 0; fi40 < 3; fi40++ {
                    v4 = v4
                    var v41 bool = false
                    _ = fi40
                    _ = v41
                }
            }
            v4 += (v4 | (v4 % (((v4 * v4)) | 1)))
        }
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        v43 := GenS2{F0: v3, F1: 98}
        _ = fi42
        _ = v43
    }
    var v44 float32 = float32(((((-2.25 * (1.5 - ((0.1 * (-2.25 * -2.25)) - 1.5))) - float64(-2.25)) / 1.5) / 1.0))
    for fi45 := 0; fi45 < 12; fi45++ {
        v4 += (-(v4))
        var v46 uint8 = 255
        _ = fi45
        _ = v46
    }
    var v47 float64 = float64(0.1)
    var v48 int = 9223372036854775807
    var v49 uint16 = 255
    switch 128 {
    case -473:
        for fi50 := 0; fi50 < 10; fi50++ {
            v47 = -2.25
            v1[(("a" + "foo bar") + "\n\t")] = (v48 ^ (v48 % ((0) | 1)))
            switch fi50 {
            case -161:
                if (3.14159 != (v47 * v47)) {
                    for fi51 := 0; fi51 < 18; fi51++ {
                        var v52 float32 = v44
                        v3 = (v3 << (((^(v3))) & 63))
                        var v53 float32 = (((v44 + 0.0) / v52) * ((((v44 - 0.0) + -2.25) / v44) / v52))
                        fmt.Println(v53, v3, v3)
                        switch (fi51 * 240) {
                        case 667:
                            fmt.Println(fi50, v3, v53)
                            if ((("a" == ("\n\t" + "\n\t")) || (((v49 < (v49 * (v49 * (v49 << ((v3) & 63))))) && (!false)) && (!false))) && (32767 > 0)) {
                                var v54 float32 = v52
                                var v55 bool = false
                                fmt.Println(v48, v55)
                                _ = v54
                                _ = v55
                            } else {
                                switch ((("\n\t" + "") + ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + ("hello" + ("hello" + ("hello" + "a"))))) + "") {
                                case "a":
                                    if false {
                                        v52 = ((3.14159 * 0.1) - 3.14159)
                                        if (1 != -440) {
                                            var v56 string = "hello"
                                            if (!((true && false) || (v47 <= v47))) {
                                                if (-1 == 648) {
                                                    v3 += (v3 ^ v3)
                                                }
                                                v57 := GenS4{F0: (v4 >> ((v3) & 63))}
                                                v44 -= float32((v47 * (v47 + v47)))
                                                _ = v57
                                            }
                                            var v58 int8 = v4
                                            _ = v56
                                            _ = v58
                                        }
                                    } else {
                                        v59 := GenS4{F0: v4}
                                        v48 = fi50
                                        v53 = v52
                                        v53 += (float32(3.14159) + -2.25)
                                        switch ("" + "foo bar") {
                                        case "hello":
                                            fmt.Println(v47, v4, v44)
                                        }
                                        var v60 uint = v3
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                if false {
                                    if ("hello" >= "a") {
                                        var v61 int16 = -780
                                        v49 += 65535
                                        v62 := GenS5{F0: 1, F1: 1}
                                        _ = v61
                                        _ = v62
                                    }
                                    v63 := GenS4{F0: (-(v4))}
                                    _ = v63
                                } else {
                                    v64 := GenS4{F0: 1}
                                    if true {
                                        v65 := GenS3{F0: 9223372036854775807, F1: (v49 * (v49 + (^(v49)))), F2: ""}
                                        fmt.Println(fi50)
                                        v66 := GenS6{F0: -813, F1: true}
                                        var v67 uint64 = uint64(v3)
                                        var v68 uint16 = (v49 >> ((v3) & 63))
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                    v69 := GenS4{F0: 0}
                                    _ = v64
                                    _ = v69
                                }
                                v1["hello"] = (v48 >> ((700) & 63))
                                switch fi50 {
                                case 2147483647:
                                    v53 = 0.0
                                default:
                                    v70 := GenS1{F0: 142, F1: v4}
                                    switch "hello" {
                                    case "hello":
                                        v48 = v48
                                    case "":
                                        v71 := GenS7{F0: "hello"}
                                        _ = v71
                                    }
                                    _ = v70
                                }
                                var v72 int8 = 127
                                _ = v72
                            }
                            fmt.Println(v53, v52)
                            _ = len(v2)
                            v73 := v1["\n\t"]
                            _ = v73
                        case 128:
                            var v74 int32 = 42
                            _ = v74
                        case -1:
                            if false {
                                switch "hello" {
                                case "hello":
                                    _ = len(v2)
                                case "":
                                    v3 -= v3
                                default:
                                    var v75 int = 4294967295
                                    _ = v75
                                }
                                fmt.Println(v52, fi51, v47)
                            }
                            v76 := GenS1{F0: 127, F1: v4}
                            _ = v76
                        }
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                }
            case 674:
                fmt.Println(v49, v44)
            case 1:
                v44 = 1e10
            default:
                var v77 float64 = float64(1.5)
                _ = v77
            }
            _ = fi50
        }
    default:
        var v78 float64 = ((float64(-2.25) + (float64(((v44 / v44) + v44)) - 1.5)) * 0.0)
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v49
}
