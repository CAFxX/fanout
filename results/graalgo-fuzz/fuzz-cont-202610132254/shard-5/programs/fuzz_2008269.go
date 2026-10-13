package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint; F1 bool; F2 uint }
type GenS8 struct { F0 int32; F1 int; F2 uint }
type GenS9 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint64; F1 int32 }

func main() {
    v1 := []uint32{800, 65535}
    var v2 uint8 = 0
    v3 := make(map[string]uint32)
    v4 := []int{32768, 836, -404}
    v5 := make(map[string]int)
    var v6 int8 = 42
    switch 0 {
    case -2147483648:
        v6 = -80
    case -1:
        v7 := GenS1{F0: 32767, F1: -1, F2: 42}
        _ = v7
    default:
        fmt.Println(v6)
    }
    v2 += 1
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 724
    v11 := GenS1{F0: 1, F1: 42, F2: -2147483648}
    switch 42 {
    case 127:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float64 = float64(0.1)
            v14 := GenS1{F0: -1, F1: 997, F2: 1}
            v15 := GenS1{F0: 1, F1: -2147483648, F2: 633}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    case -975:
        var v16 int = 32767
        _ = v16
    default:
        v17 := GenS1{F0: 32767, F1: 2147483647, F2: int32(v10)}
        _ = v17
    }
    var v18 int32 = -1
    var v19 uint = 123
    if (!(!(-2.25 != 0.0))) {
        var v20 uint64 = 1
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS1{F0: 32767, F1: 42, F2: (v18 << (((v19 / ((v19) | 1))) & 63))}
            var v23 int = fi21
            _ = fi21
            _ = v22
            _ = v23
        }
        if true {
            if (true && true) {
                v2 += uint8(v10)
                if ((("hello" + "\n\t") != ("" + "hello")) && ((1.5 / 1.0) == ((float32((((0.0 + 0.0) + float64((((0.0 + (0.1 / 1e10)) / (float32(((1.5 / ((0.0 / 3.14159) + ((1e10 * (((-2.25 - -2.25) - (((float64(-2.25) + 0.1) * 1.5) / 1.0)) - (0.1 - ((1e10 * (((-2.25 - ((0.1 - (0.1 * 1e10)) * 0.1)) / 0.1) * 0.1)) / -2.25)))) - 0.0))) / 3.14159)) * (3.14159 + (1e10 / (3.14159 + 3.14159))))) - 1e10))) / 1.5)) / (-2.25 / 0.1)) + ((3.14159 + 3.14159) / 1e10)))) {
                    if ((-32768 != 0) || false) {
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v19 = (^(v19))
                            fmt.Println(v10, v2)
                            var v25 int32 = 0
                            _ = fi24
                            _ = v25
                        }
                        var v26 int32 = -1
                        v27 := GenS1{F0: 0, F1: v18, F2: v18}
                        _ = v26
                        _ = v27
                    } else {
                        switch -55 {
                        case 152:
                            v2 = 160
                        case -665:
                            if len(v1) > 0 {
                                v1[0] = v10
                            }
                        case -2147483648:
                            switch 9223372036854775807 {
                            case 4294967295:
                                v28 := GenS1{F0: 693, F1: 364, F2: v18}
                                _ = v28
                            default:
                                v18 += int32(v19)
                            }
                        default:
                            v10 -= (v10 / (((^(v10))) | 1))
                        }
                    }
                    v29 := GenS2{F0: 42, F1: (((^(v20)) == uint64(v2)) && (-137 != -923)), F2: -32768}
                    _ = v29
                }
                v30 := GenS3{F0: v20, F1: 127}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 15; fi31++ {
                    v19 = v19
                    v2 = (^(v2))
                    switch -936 {
                    case 2147483647:
                        v32 := GenS4{F0: v18, F1: v20}
                        _ = v32
                    case -461:
                        var v33 uint = (v19 / ((65535) | 1))
                        _ = v33
                    }
                    _ = fi31
                }
            }
            if (-834 != -392) {
                v34 := GenS5{F0: (v19 >> ((v19) & 63)), F1: 9223372036854775807}
                _ = v34
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                switch -913 {
                case 978:
                    v35 := v5["\n\t"]
                    _ = v35
                default:
                    v36 := GenS2{F0: 42, F1: (false && (true || (9223372036854775807 == 9223372036854775807))), F2: 42}
                    _ = v36
                }
            case "hello":
                _ = len(v3)
            case "foo bar":
                for fi37 := 0; fi37 < 16; fi37++ {
                    _ = len(v1)
                    v4 = append(v4, 372)
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi37, v19, v20)
                    case "a":
                        for fi38 := 0; fi38 < 18; fi38++ {
                            if true {
                                _ = len(v4)
                            }
                            v18 += (v18 ^ 2147483647)
                            if ("" >= "foo bar") {
                                switch "foo bar" {
                                case "\n\t":
                                    v6 += (v6 + 42)
                                case "a":
                                    if (1 < 32767) {
                                        v6 -= int8(v10)
                                        v39 := GenS6{F0: v19, F1: -1}
                                        v2 = 42
                                        v5["hello"] = fi37
                                        var v40 uint = v19
                                        if len(v1) > 0 {
                                            v1[0] = (v10 >> (((^(v19))) & 63))
                                        }
                                        fmt.Println(v2, fi37, fi37)
                                        _ = v39
                                        _ = v40
                                    } else {
                                        _ = len(v5)
                                        var v41 uint32 = (v10 << ((2) & 63))
                                        if false {
                                            switch "" {
                                            case "":
                                                _ = len(v1)
                                            case "foo bar":
                                                if (("hello" + "a") < ("hello" + "")) {
                                                    var v42 bool = false
                                                    var v43 uint64 = (v20 + v20)
                                                    fmt.Println(v10, v20)
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    v6 -= v6
                                                    v6 += (v6 & (^(v6)))
                                                    var v44 int16 = 32767
                                                    _ = v44
                                                }
                                            default:
                                                switch fi37 {
                                                case 388:
                                                    v45 := GenS3{F0: 42, F1: 118}
                                                    _ = v45
                                                case -897:
                                                    v41 = (v10 % ((4294967295) | 1))
                                                case -971:
                                                    switch "\n\t" {
                                                    case "":
                                                        v46 := GenS1{F0: 32767, F1: (v18 | (v18 / (((v18 >> (((v19 / ((42) | 1))) & 63))) | 1))), F2: 753}
                                                        _ = v46
                                                    case "\n\t":
                                                        if false {
                                                            if ((v18 % (((v18 & v18)) | 1)) > (v18 << ((159) & 63))) {
                                                                v19 -= 42
                                                                v41 = 2147483647
                                                                v47 := GenS7{F0: 834, F1: true, F2: (v19 - (v19 << ((4294967295) & 63)))}
                                                                v48 := v5["\n\t"]
                                                                v49 := GenS2{F0: 0, F1: true, F2: -32768}
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        }
                                                    case "hello":
                                                        v20 += (v20 << ((uint(v10)) & 63))
                                                    }
                                                default:
                                                    v1 = append(v1, v41)
                                                }
                                            }
                                            fmt.Println(v2)
                                            v10 = 2147483647
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = fi38
                                            }
                                        }
                                        var v50 uint32 = v10
                                        switch "hello" {
                                        case "":
                                            v3[""] = v10
                                        }
                                        var v51 uint = 2
                                        if (42 >= -1) {
                                            v52 := GenS5{F0: v51, F1: 9223372036854775807}
                                            v20 += (^(v20))
                                            v53 := GenS7{F0: v51, F1: true, F2: (^(v19))}
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 uint = 2
                                            if (!(!(!(true || true)))) {
                                                if (2147483647 <= -2147483648) {
                                                    var v55 int = -999
                                                    _ = v55
                                                } else {
                                                    if (!true) {
                                                        var v56 int = (fi38 + 577)
                                                        _ = v56
                                                    }
                                                }
                                                var v57 float64 = 0.0
                                                var v58 int64 = 1
                                                _ = v57
                                                _ = v58
                                            }
                                            v19 -= 133
                                            _ = v54
                                        }
                                        _ = v41
                                        _ = v50
                                        _ = v51
                                    }
                                case "":
                                    v59 := GenS8{F0: 0, F1: (fi38 * (^(fi37))), F2: 431}
                                    _ = v59
                                }
                            } else {
                                fmt.Println(v2, fi37, v2)
                                v60 := GenS7{F0: (^(v19)), F1: true, F2: (v19 | 255)}
                                v19 -= v19
                                v61 := GenS9{F0: 1, F1: v2, F2: v6}
                                if (!false) {
                                    v19 = v19
                                    v62 := GenS10{F0: 4294967295, F1: (v18 * (-(v18)))}
                                    _ = v62
                                }
                                _ = v60
                                _ = v61
                            }
                            _ = fi38
                        }
                    default:
                        v5["foo bar"] = (-(fi37))
                    }
                    _ = fi37
                }
            default:
                var v63 uint8 = v2
                _ = v63
            }
            for fi64 := 0; fi64 < 8; fi64++ {
                var v65 uint8 = v2
                v6 -= (v6 - (-(v6)))
                var v66 uint32 = 567
                _ = fi64
                _ = v65
                _ = v66
            }
        }
        _ = v20
    }
    if len(v1) > 0 {
        v1[0] = (v10 - 2147483648)
    }
    switch "hello" {
    case "hello":
        for fi67 := 0; fi67 < 11; fi67++ {
            _ = len(v1)
            v10 = v10
            switch (("\n\t" + "foo bar") + ((("a" + ("" + (("foo bar" + "a") + "hello"))) + "a") + (("\n\t" + "\n\t") + ("hello" + (("\n\t" + "foo bar") + "\n\t"))))) {
            case "hello":
                v68 := GenS4{F0: (-(v18)), F1: 384}
                _ = v68
            }
            _ = fi67
        }
    case "":
        var v69 int16 = 42
        _ = v69
    }
    if true {
        var v70 int8 = (v6 - (^(v6)))
        fmt.Println(v70, v10, v6)
        _ = v70
    } else {
        _ = len(v4)
        v6 = 100
    }
    _ = len(v1)
    if false {
        v6 = v6
    } else {
        var v71 bool = true
        v19 = (v19 % ((v19) | 1))
        _ = v71
    }
    switch 32768 {
    case -2147483648:
        var v72 uint64 = 377
        _ = v72
    }
    var v73 int32 = (v18 << ((101) & 63))
    for fi74 := 0; fi74 < 15; fi74++ {
        var v75 bool = true
        var v76 string = ""
        _ = fi74
        _ = v75
        _ = v76
    }
    if (!((("a" != "") && true) && (!(!true)))) {
        fmt.Println(v18, v6)
        for fi77 := 0; fi77 < 13; fi77++ {
            for fi78 := 0; fi78 < 12; fi78++ {
                v19 -= 403
                v79 := GenS4{F0: v18, F1: 0}
                _ = fi78
                _ = v79
            }
            var v80 int16 = -32768
            _ = fi77
            _ = v80
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
