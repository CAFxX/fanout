package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 string; F2 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int8; F1 uint64; F2 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    v4 := []string{"\n\t", "", "foo bar", ((("" + (("a" + "hello") + "foo bar")) + "foo bar") + "hello")}
    var v5 string = "hello"
    if false {
        var v6 float64 = 1.5
        var v7 uint8 = 252
        switch (v5 + (v5 + "")) {
        case "foo bar":
            v6 = (v6 + (float64(float32(v6)) - v6))
        case "":
            v8 := GenS1{F0: 100, F1: 32767}
            _ = v8
        case "a":
            if (4294967295 < 4294967295) {
                v5 = v5
                switch "\n\t" {
                case "foo bar":
                    switch ((v5 + (("foo bar" + (v5 + ("hello" + ("\n\t" + ("" + v5))))) + ((v5 + (v5 + "hello")) + "\n\t"))) + ("a" + "hello")) {
                    case "foo bar":
                        v9 := GenS2{F0: 42}
                        _ = v9
                    }
                case "hello":
                    v5 = "hello"
                default:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if (-2.25 == 0.1) {
                            fmt.Println(v5, fi10, v7)
                        }
                        _ = fi10
                    }
                }
                v11 := GenS3{F0: 546, F1: ((v7 ^ v7) >= (v7 / (((^(v7))) | 1))), F2: 65535}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 4; fi12++ {
                    for fi13 := 0; fi13 < 6; fi13++ {
                        fmt.Println(fi13, v7, fi13)
                        if false {
                            v5 = "hello"
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
            }
        }
        _ = v6
        _ = v7
    }
    var v14 float32 = (3.14159 * ((0.0 - 1.5) + (1e10 - (3.14159 - 0.0))))
    v5 = (("\n\t" + ("hello" + ((("" + v5) + ("" + "hello")) + v5))) + ("hello" + v5))
    for fi15 := 0; fi15 < 18; fi15++ {
        v14 += (v14 + v14)
        if (v14 >= v14) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (1 > 158) {
                if true {
                    if ((((v5 + "foo bar") + ((((v5 + v5) + "foo bar") + v5) + "a")) + v5) < v5) {
                        v14 -= (1e10 + v14)
                        for fi16 := 0; fi16 < 15; fi16++ {
                            fmt.Println(fi15, v5, fi16)
                            _ = fi16
                        }
                    }
                    v3[(fi15 >> ((255) & 63))] = fi15
                    v5 = (v5 + v5)
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v3[1] = fi15
                }
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint32 = 0
                    _ = fi17
                    _ = v18
                }
            } else {
                var v19 float64 = ((1e10 / -2.25) / 1.5)
                _ = v19
            }
        }
        if (fi15 == (fi15 - (-(fi15)))) {
            switch (fi15 << ((101) & 63)) {
            case 127:
                if (false && true) {
                    if false {
                        if true {
                            v5 = v5
                            v20 := GenS4{F0: 677, F1: 4294967295}
                            _ = v20
                        } else {
                            var v21 int16 = 32767
                            var v22 int = fi15
                            _ = v21
                            _ = v22
                        }
                    } else {
                        if false {
                            fmt.Println(v5, v14, v5)
                            var v23 float64 = -2.25
                            _ = v23
                        }
                    }
                    v24 := GenS3{F0: -32768, F1: (-2147483648 > -2147483648), F2: 769}
                    delete(v2, "foo bar")
                    _ = v24
                }
            case 4294967295:
                var v25 float64 = float64((((v14 - v14) * float32(0.1)) - -2.25))
                _ = v25
            }
        }
        _ = fi15
    }
    if (1 < -424) {
        switch v5 {
        case "hello":
            v14 = (3.14159 + v14)
        }
    } else {
        var v26 uint = 551
        if false {
            if true {
                if (false || (2147483648 <= 65535)) {
                    var v27 string = "foo bar"
                    _ = v27
                } else {
                    if ((v26 << ((928) & 63)) <= (v26 * v26)) {
                        if true {
                            v28 := GenS5{F0: -702, F1: 42, F2: false}
                            if (-32768 >= -32768) {
                                if true {
                                    v29 := GenS6{F0: 250}
                                    v5 = "hello"
                                    _ = v29
                                } else {
                                    if true {
                                        v4 = append(v4, v5)
                                    }
                                }
                                delete(v1, v5)
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    v26 = v26
                                    if true {
                                        fmt.Println(v26, fi30)
                                        v14 += 1.5
                                        v14 = v14
                                    }
                                    _ = fi30
                                }
                            } else {
                                var v31 uint16 = 0
                                _ = v31
                            }
                            _ = v28
                        } else {
                            for fi32 := 0; fi32 < 4; fi32++ {
                                switch fi32 {
                                case 762:
                                    fmt.Println(v5, fi32, v26)
                                case 684:
                                    switch fi32 {
                                    case 1:
                                        if false {
                                            if (65535 <= 1) {
                                                if (345 > 2) {
                                                    if false {
                                                        var v33 uint8 = 74
                                                        _ = v33
                                                    }
                                                    var v34 string = v5
                                                    _ = v34
                                                } else {
                                                    v35 := GenS2{F0: 1}
                                                    _ = v35
                                                }
                                            }
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    case 32767:
                                        v36 := GenS1{F0: 4, F1: -32768}
                                        _ = v36
                                    case -449:
                                        v26 -= v26
                                    default:
                                        switch v5 {
                                        case "hello":
                                            v37 := v2[(("\n\t" + v5) + "a")]
                                            _ = v37
                                        }
                                    }
                                case 418:
                                    v14 = 1e10
                                }
                                switch 0 {
                                case 1:
                                    v14 = (v14 - -2.25)
                                case 2147483647:
                                    if (true || false) {
                                        if (!(234 != 127)) {
                                            var v38 uint64 = 0
                                            for fi39 := 0; fi39 < 7; fi39++ {
                                                v5 = v5
                                                v40 := GenS7{F0: -2147483648}
                                                v26 += (v26 & uint(v38))
                                                switch v5 {
                                                case "":
                                                    v38 += v38
                                                case "foo bar":
                                                    v41 := GenS2{F0: -337}
                                                    _ = v41
                                                }
                                                _ = fi39
                                                _ = v40
                                            }
                                            _ = v38
                                        }
                                        v42 := GenS8{F0: (fi32 * (fi32 ^ 42)), F1: (v5 + ("foo bar" + "a")), F2: "a"}
                                        var v43 int32 = -1
                                        _ = v42
                                        _ = v43
                                    } else {
                                        switch v5 {
                                        case "hello":
                                            for fi44 := 0; fi44 < 15; fi44++ {
                                                var v45 int16 = -592
                                                switch fi32 {
                                                case 9223372036854775807:
                                                    v26 -= 255
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = ("\n\t" + (v5 + ""))
                                                    }
                                                    v26 = 65535
                                                    v46 := GenS9{F0: (fi44 & 2147483647)}
                                                    _ = v46
                                                }
                                                _ = fi44
                                                _ = v45
                                            }
                                        case "":
                                            switch fi32 {
                                            case 819:
                                                v5 = (v5 + v5)
                                            }
                                        case "foo bar":
                                            var v47 uint32 = 2147483648
                                            _ = v47
                                        }
                                    }
                                default:
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        v14 = 0.1
                                        v49 := v1["a"]
                                        fmt.Println(v49, v14, fi48)
                                        _ = fi48
                                        _ = v49
                                    }
                                }
                                _ = fi32
                            }
                        }
                        v50 := GenS10{F0: false, F1: -9223372036854775808}
                        v51 := v1[v5]
                        _ = v50
                        _ = v51
                    } else {
                        if false {
                            for fi52 := 0; fi52 < 17; fi52++ {
                                switch (fi52 >> ((v26) & 63)) {
                                case -1:
                                    delete(v3, (fi52 >> ((119) & 63)))
                                case 292:
                                    if true {
                                        v53 := GenS5{F0: int32(v26), F1: -1, F2: (0 > -1)}
                                        v14 = 0.1
                                        _ = v53
                                    }
                                }
                                _ = fi52
                            }
                        } else {
                            if true {
                                if (18446744073709551615 == 42) {
                                    if true {
                                        for fi54 := 0; fi54 < 9; fi54++ {
                                            if len(v4) > 0 {
                                                v4[0] = v5
                                            }
                                            if false {
                                                switch -453 {
                                                case -139:
                                                    if false {
                                                        v14 += v14
                                                    }
                                                default:
                                                    v4 = append(v4, ((v5 + ((v5 + "\n\t") + (((v5 + ("a" + v5)) + v5) + v5))) + v5))
                                                }
                                                v14 = v14
                                                switch v5 {
                                                case "a":
                                                    var v55 int8 = 19
                                                    _ = v55
                                                case "hello":
                                                    delete(v2, "")
                                                default:
                                                    v14 -= float32(0.1)
                                                }
                                            }
                                            v56 := GenS4{F0: 952, F1: 42}
                                            _ = fi54
                                            _ = v56
                                        }
                                        v26 = v26
                                        delete(v1, v5)
                                    }
                                    for fi57 := 0; fi57 < 1; fi57++ {
                                        var v58 int16 = -32768
                                        _ = fi57
                                        _ = v58
                                    }
                                    if (!false) {
                                        for fi59 := 0; fi59 < 3; fi59++ {
                                            v60 := v3[(fi59 + (^(fi59)))]
                                            for fi61 := 0; fi61 < 4; fi61++ {
                                                v26 += (v26 >> ((v26) & 63))
                                                _ = fi61
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                        v4 = append(v4, ("" + v5))
                                        v5 = v5
                                    } else {
                                        for fi62 := 0; fi62 < 19; fi62++ {
                                            v63 := GenS9{F0: (^(fi62))}
                                            v64 := GenS11{F0: 42, F1: 103}
                                            if ((^(v26)) > 2) {
                                                var v65 uint = v26
                                                _ = v65
                                            }
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                        }
                                        var v66 uint = (v26 >> (((^(v26))) & 63))
                                        _ = v66
                                    }
                                }
                                v67 := GenS12{F0: -128, F1: 639, F2: ((v5 + v5) + ("\n\t" + "a"))}
                                v14 = v14
                                _ = v67
                            }
                        }
                        var v68 uint8 = 1
                        _ = v68
                    }
                    if (65 >= 0) {
                        v5 = v5
                        fmt.Println(v26, v5, v5)
                    }
                    v14 += float32(3.14159)
                }
                v26 = (v26 << ((v26) & 63))
                switch ((v5 + (v5 + (v5 + "foo bar"))) + "") {
                case "\n\t":
                    var v69 float32 = 0.1
                    _ = v69
                }
            } else {
                var v70 bool = (((-2.25 / ((3.14159 * 3.14159) - (0.0 - ((1.5 / 3.14159) / (1.5 / 1e10))))) != float64(0.1)) && false)
                _ = v70
            }
            switch "foo bar" {
            case "\n\t":
                v26 = (v26 ^ v26)
            }
        } else {
            var v71 int8 = 75
            _ = v71
        }
        _ = v26
    }
    switch (v5 + "foo bar") {
    case "hello":
        v5 = v5
    default:
        _ = len(v4)
    }
    delete(v1, v5)
    delete(v2, v5)
    var v72 int = 32768
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v73 uint8 = 128
    var v74 int32 = 567
    var v75 bool = false
    var v76 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
