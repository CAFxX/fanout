package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int16; F1 string }
func genFn0(p0 uint8, p1 int32, p2 int32) string {
    p0 -= p0
    fmt.Println(p0, p1, p2)
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(p2, p1)
        p1 = p1
        fmt.Println(p2, fi6, p0)
        _ = fi6
    }
    var v7 uint = 300
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ""
}

func main() {
    v1 := []string{}
    var v2 int16 = -1
    var v3 int32 = -97
    v4 := []uint32{65535, 42}
    var v5 int32 = (v3 << ((255) & 63))
    if len(v4) > 0 {
        v4[0] = 708
    }
    fmt.Println(v3, v5)
    v2 -= v2
    if true {
        v8 := GenS1{F0: v5, F1: -761, F2: 127}
        _ = v8
    } else {
        switch (("\n\t" + "") + "foo bar") {
        case "":
            v2 += v2
        default:
            for fi9 := 0; fi9 < 6; fi9++ {
                v3 = v5
                var v10 int = 316
                _ = fi9
                _ = v10
            }
        }
    }
    v5 -= (v5 | v3)
    if (true && true) {
        if (!((((("" + ((((("a" + "a") + "a") + (("a" + "a") + ("a" + ""))) + "\n\t") + (("foo bar" + (("a" + "hello") + (("foo bar" + "") + "hello"))) + "foo bar"))) + ("\n\t" + ("hello" + "a"))) == "\n\t") && (831 != 0)) || (!false))) {
            v5 = (v3 >> ((2) & 63))
        }
        var v11 int64 = -347
        if true {
            fmt.Println(v5, v5, v3)
            v3 = (-(v5))
        }
        _ = v11
    } else {
        var v12 uint32 = 65535
        for fi13 := 0; fi13 < 8; fi13++ {
            v2 -= v2
            _ = fi13
        }
        v5 = (v5 * 2147483647)
        _ = v12
    }
    if ((-1 <= 32767) && false) {
        fmt.Println(v5, v5, v2)
        var v14 float32 = (float32((((0.0 - (1e10 + 1.5)) + float64(3.14159)) - 0.1)) / (3.14159 / (3.14159 * (((0.1 * ((3.14159 * 1e10) - 1e10)) * (3.14159 + 0.0)) + ((float32(1.5) * 0.1) * 0.1)))))
        v2 = 32767
        _ = v14
    }
    fmt.Println(v2, v2)
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 = v2
        _ = fi15
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v2 -= 1
        v2 += v2
        for fi17 := 0; fi17 < 4; fi17++ {
            switch (^(fi17)) {
            case -902:
                _ = len(v4)
            }
            _ = fi17
        }
        _ = fi16
    }
    v18 := GenS1{F0: (v5 & (v5 & v5)), F1: 787, F2: int(v2)}
    if false {
        v2 = v2
        v19 := GenS2{F0: 186, F1: v5, F2: 1}
        fmt.Println(v3)
        _ = v19
    } else {
        if ((("" + ("" + "foo bar")) < "hello") || (9223372036854775807 == -9223372036854775808)) {
            v4 = append(v4, 4294967295)
            v2 = v2
        } else {
            var v20 uint32 = 942
            for fi21 := 0; fi21 < 1; fi21++ {
                v22 := GenS1{F0: v3, F1: 534, F2: -1}
                fmt.Println(v3, v3, v20)
                _ = fi21
                _ = v22
            }
            var v23 int16 = (v2 + v2)
            _ = v20
            _ = v23
        }
        fmt.Println(v5, v5)
    }
    v4 = append(v4, 2147483648)
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: (v5 % (((v3 / (((v3 >> ((587) & 63))) | 1))) | 1)), F1: -810, F2: 42}
        var v26 int32 = v5
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: -1, F2: 2}
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 uint8 = 128
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (fi28 ^ -617) {
            case 439:
                v5 = v3
            default:
                fmt.Println(v3)
                switch ("" + "hello") {
                case "hello":
                    var v31 uint32 = 580
                    _ = v31
                case "\n\t":
                    v32 := GenS1{F0: 2147483647, F1: (-(fi29)), F2: -9223372036854775808}
                    _ = v32
                default:
                    if (389 < 42) {
                        v3 = v3
                        if (!(!true)) {
                            v33 := GenS3{F0: "", F1: (fi29 | (-(fi28))), F2: 889}
                            var v34 int = 68
                            _ = v33
                            _ = v34
                        }
                    }
                }
            }
            v3 -= v3
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v2, v5)
    default:
        if true {
            v4 = append(v4, 2147483648)
            if (true && true) {
                v5 += (v5 * v3)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    v35 := GenS3{F0: (("" + "foo bar") + "\n\t"), F1: 924, F2: 4294967295}
                    if false {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            if (false || (false && (!false))) {
                                v37 := GenS4{F0: "foo bar", F1: -85}
                                _ = v37
                            } else {
                                switch (fi36 ^ 1) {
                                case 800:
                                    v1 = append(v1, "\n\t")
                                case -922:
                                    var v38 float32 = ((float32(-2.25) - 0.0) + 3.14159)
                                    _ = v38
                                }
                            }
                            _ = fi36
                        }
                        v2 = v2
                    } else {
                        var v39 int32 = v5
                        switch -314 {
                        case 42:
                            v40 := GenS3{F0: "hello", F1: -9223372036854775808, F2: 36}
                            _ = v40
                        case 32768:
                            v39 += (v3 >> ((0) & 63))
                        case 0:
                            if len(v4) > 0 {
                                v4[0] = 475
                            }
                        }
                        switch "" {
                        case "":
                            if true {
                                _ = len(v1)
                            } else {
                                switch 9223372036854775807 {
                                case 885:
                                    var v41 float32 = 0.1
                                    _ = v41
                                case 32767:
                                    var v42 float32 = 0.0
                                    _ = v42
                                }
                                v1 = append(v1, "\n\t")
                                fmt.Println(v3)
                            }
                        case "foo bar":
                            v43 := GenS4{F0: "foo bar", F1: 0}
                            _ = v43
                        case "a":
                            v2 += (v2 & (v2 * (v2 & v2)))
                        default:
                            var v44 int = -2147483648
                            _ = v44
                        }
                        _ = v39
                    }
                    _ = v35
                }
            } else {
                v2 -= -32768
                fmt.Println(v3, v2, v2)
                for fi45 := 0; fi45 < 18; fi45++ {
                    var v46 uint16 = 235
                    v47 := GenS4{F0: ("\n\t" + ""), F1: 78}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                for fi48 := 0; fi48 < 19; fi48++ {
                    var v49 bool = (279 != 706)
                    var v50 int16 = v2
                    v3 = v3
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            case "hello":
                v51 := GenS5{F0: (-(v2)), F1: "a"}
                _ = v51
            default:
                for fi52 := 0; fi52 < 10; fi52++ {
                    switch "a" {
                    case "":
                        for fi53 := 0; fi53 < 15; fi53++ {
                            fmt.Println(v5, fi53, v3)
                            var v54 int64 = -9223372036854775808
                            _ = fi53
                            _ = v54
                        }
                    }
                    for fi55 := 0; fi55 < 16; fi55++ {
                        if ((v5 + 594) < (^(v5))) {
                            if false {
                                fmt.Println(fi52)
                            }
                            if (!false) {
                                var v56 uint16 = 1
                                v57 := GenS3{F0: ((("hello" + "\n\t") + "") + ("foo bar" + "hello")), F1: (fi55 * -2147483648), F2: 2}
                                v56 = 2
                                v58 := GenS1{F0: (v3 / ((v5) | 1)), F1: fi52, F2: fi52}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v59 := GenS3{F0: "hello", F1: fi52, F2: 2}
                                var v60 int = 32768
                                _ = v56
                                _ = v57
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                v61 := GenS2{F0: 0, F1: v3, F2: 4294967295}
                                if (0.0 <= (0.1 - (3.14159 / 1.0))) {
                                    var v62 uint = 18446744073709551615
                                    if (0.1 < (float32(3.14159) * -2.25)) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    var v63 int32 = 1
                                    var v64 string = ((("\n\t" + ("foo bar" + "")) + ("a" + "")) + "a")
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                } else {
                                    if true {
                                        if true {
                                            v3 = v5
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = ("" + (("hello" + (("" + (("\n\t" + ("hello" + "hello")) + (("a" + ("" + "a")) + ""))) + "")) + ("foo bar" + "a")))
                                            }
                                        }
                                    } else {
                                        switch "" {
                                        case "":
                                            v5 += (v5 % ((1) | 1))
                                        case "hello":
                                            fmt.Println(v2)
                                        case "\n\t":
                                            v65 := GenS2{F0: 1, F1: (v3 << ((653) & 63)), F2: 944}
                                            _ = v65
                                        default:
                                            var v66 int8 = 100
                                            _ = v66
                                        }
                                    }
                                    v5 = v3
                                }
                                _ = len(v4)
                                _ = v61
                            }
                            _ = len(v4)
                        } else {
                            var v67 bool = ((665 <= 2) && true)
                            _ = v67
                        }
                        var v68 float32 = (((0.1 * 3.14159) * 3.14159) / 1e10)
                        _ = fi55
                        _ = v68
                    }
                    _ = fi52
                }
            }
            for fi69 := 0; fi69 < 5; fi69++ {
                for fi70 := 0; fi70 < 4; fi70++ {
                    var v71 int8 = 1
                    v71 = (v71 + v71)
                    _ = fi70
                    _ = v71
                }
                _ = fi69
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
