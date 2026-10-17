package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint64, p2 bool) uint8 {
    var v7 uint32 = 2147483648
    p1 += (p1 ^ 4294967295)
    v7 = (v7 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 string, p2 int16) uint16 {
    switch 0 {
    case 128:
        p1 = (p1 + (("hello" + "\n\t") + p1))
    case -558:
        var v8 float32 = 0.1
        _ = v8
    case -474:
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int16 = p2
            p0 = (p0 + 606)
            var v11 int64 = p0
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ((p1 + (p1 + p1)) + (p1 + (((((("a" + "foo bar") + (("foo bar" + "") + (("foo bar" + (("" + (p1 + (p1 + p1))) + "foo bar")) + ("hello" + ((p1 + ((("" + p1) + ((p1 + p1) + "")) + "\n\t")) + (("" + p1) + ("foo bar" + ""))))))) + ((p1 + p1) + "foo bar")) + p1) + p1) + (("\n\t" + "a") + ("\n\t" + ((("\n\t" + ((p1 + (p1 + p1)) + "foo bar")) + ((("" + p1) + p1) + p1)) + "foo bar")))))) {
            case "a":
                p2 -= (p2 | p2)
            case "":
                v13 := GenS1{F0: 1, F1: (p2 % (((p2 / ((p2) | 1))) | 1))}
                _ = v13
            }
            if ((p1 != "foo bar") || false) {
                p2 += p2
            } else {
                v14 := []int{(fi12 / ((fi12) | 1))}
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 float64 = 0.1
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := []int{(fi17 << ((976) & 63)), (fi17 * (^(fi12)))}
                if len(v18) > 0 {
                    v18[0] = 936
                }
                switch (fi12 + (fi12 >> ((255) & 63))) {
                case 42:
                    if (1 <= 4294967295) {
                        v19 := GenS1{F0: 0, F1: p2}
                        switch (fi17 ^ (fi17 ^ fi17)) {
                        case 1:
                            v20 := GenS2{F0: p2}
                            _ = v20
                        case -654:
                            _ = len(v18)
                        default:
                            v21 := GenS2{F0: 0}
                            _ = v21
                        }
                        _ = v19
                    }
                case -1:
                    switch fi12 {
                    case -934:
                        if true {
                            if (2 > genFn0(2, 327, false)) {
                                p2 = p2
                                var v22 int = (fi17 ^ fi12)
                                switch fi12 {
                                case -155:
                                    v22 = fi12
                                case 42:
                                    var v23 float32 = -2.25
                                    _ = v23
                                }
                                _ = v22
                            } else {
                                if true {
                                    p2 -= 42
                                    fmt.Println(p1, p0, p2)
                                    var v24 float64 = 3.14159
                                    _ = v24
                                }
                                if len(v18) > 0 {
                                    v18[0] = fi12
                                }
                                fmt.Println(fi17, fi17, p0)
                            }
                            if ((!true) || true) {
                                var v25 int = fi17
                                var v26 uint8 = 25
                                _ = v25
                                _ = v26
                            }
                            var v27 uint8 = 1
                            _ = v27
                        } else {
                            v18 = append(v18, fi12)
                            v28 := GenS3{F0: (p0 + 9223372036854775807), F1: 255}
                            if (42 >= 255) {
                                v29 := GenS4{F0: -1}
                                fmt.Println(fi12)
                                _ = v29
                            } else {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            _ = v28
                        }
                    case 568:
                        if (!true) {
                            v30 := GenS1{F0: 42, F1: p2}
                            fmt.Println(p1, fi17)
                            var v31 uint = 386
                            v31 += v31
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint64 = 219
                            v33 := GenS4{F0: 0}
                            var v34 int8 = 0
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    default:
                        var v35 bool = true
                        _ = v35
                    }
                case -677:
                    var v36 uint = 18446744073709551615
                    _ = v36
                }
                _ = fi17
                _ = v18
            }
            _ = fi12
        }
    }
    p1 = ((p1 + "foo bar") + p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 644
    v3 := make(map[uint32]string)
    v4 := []string{"hello", "foo bar"}
    v5 := []uint32{}
    var v6 int = 4294967295
    v2 -= (v2 | v2)
    fmt.Println(v2, v2)
    for fi37 := 0; fi37 < 1; fi37++ {
        if len(v4) > 0 {
            v4[0] = ""
        }
        for fi38 := 0; fi38 < 2; fi38++ {
            if (!((v6 / (((fi37 ^ -2147483648)) | 1)) >= (fi37 + 32767))) {
                v1 -= (v1 ^ (v1 | v1))
                var v39 int = (-(fi38))
                switch v39 {
                case 32767:
                    var v40 int16 = 0
                    _ = v40
                }
                _ = v39
            } else {
                v41 := GenS5{F0: (v1 - v1)}
                _ = v41
            }
            v1 = 52
            _ = fi38
        }
        var v42 uint64 = 18446744073709551615
        _ = fi37
        _ = v42
    }
    fmt.Println(v6, v2, v1)
    var v43 bool = (1e10 > 0.1)
    for fi44 := 0; fi44 < 6; fi44++ {
        switch "a" {
        case "\n\t":
            if (1 <= -350) {
                fmt.Println(fi44, v2, v43)
                if v43 {
                    switch (v6 >> ((65535) & 63)) {
                    case 42:
                        for fi45 := 0; fi45 < 3; fi45++ {
                            switch "foo bar" {
                            case "hello":
                                v2 = (^(v2))
                            case "a":
                                _ = len(v3)
                            case "foo bar":
                                v3[v2] = ("foo bar" + "foo bar")
                            default:
                                if false {
                                    var v46 uint8 = 42
                                    fmt.Println(v1, v43, v6)
                                    var v47 uint32 = (^(v2))
                                    _ = v46
                                    _ = v47
                                }
                            }
                            _ = fi45
                        }
                    }
                    v3[(v2 ^ 65535)] = (("foo bar" + "foo bar") + "foo bar")
                }
            } else {
                var v48 uint16 = 2
                _ = v48
            }
        case "foo bar":
            for fi49 := 0; fi49 < 4; fi49++ {
                v1 += (v1 << ((409) & 63))
                var v50 int32 = -1
                var v51 string = "a"
                _ = fi49
                _ = v50
                _ = v51
            }
        case "a":
            v52 := GenS2{F0: 32767}
            _ = v52
        }
        v2 = (v2 & (^(v2)))
        _ = fi44
    }
    var v53 int64 = -550
    fmt.Println(v1)
    v54 := GenS6{F0: 1000, F1: 597, F2: 127}
    for fi55 := 0; fi55 < 1; fi55++ {
        for fi56 := 0; fi56 < 6; fi56++ {
            _ = len(v3)
            _ = fi56
        }
        v57 := GenS1{F0: (^(v2)), F1: 42}
        if (!false) {
            if (42 >= 0) {
                switch fi55 {
                case -9223372036854775808:
                    var v58 int8 = v1
                    _ = v58
                case 936:
                    switch int(v2) {
                    case 127:
                        var v59 float64 = 1.5
                        _ = v59
                    }
                case 32768:
                    switch (^(v6)) {
                    case 32768:
                        v60 := GenS7{F0: 4294967295}
                        _ = v60
                    case -467:
                        v53 = v53
                    case 1:
                        v61 := GenS8{F0: 255}
                        _ = v61
                    default:
                        for fi62 := 0; fi62 < 14; fi62++ {
                            var v63 bool = v43
                            v64 := GenS2{F0: 32767}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                }
                switch v6 {
                case -7:
                    if (18446744073709551615 == 42) {
                        var v65 uint16 = 2
                        _ = v65
                    } else {
                        if (65535 <= 2) {
                            switch "" {
                            case "hello":
                                v53 -= (^(v53))
                            case "\n\t":
                                v2 += (v2 << ((188) & 63))
                            default:
                                var v66 int = (^(v6))
                                _ = v66
                            }
                            for fi67 := 0; fi67 < 11; fi67++ {
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                                v6 = (^(fi55))
                                switch "" {
                                case "\n\t":
                                    v1 += (^(v1))
                                }
                                if v43 {
                                    v6 = fi55
                                }
                                _ = fi67
                            }
                        }
                        var v68 uint8 = 128
                        _ = v68
                    }
                }
            }
            var v69 uint64 = 2
            _ = v69
        }
        _ = fi55
        _ = v57
    }
    fmt.Println(v2, v6)
    var v70 float32 = -2.25
    v71 := GenS3{F0: v53, F1: 1}
    switch "a" {
    case "foo bar":
        for fi72 := 0; fi72 < 4; fi72++ {
            v3[(^(v2))] = "foo bar"
            var v73 int32 = -2147483648
            v74 := GenS1{F0: 2, F1: 0}
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    var v75 bool = v43
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v70
    _ = v71
    _ = v75
}
