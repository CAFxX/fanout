package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 int16) int64 {
    if (0 > (^(p0))) {
        p0 = 267
        v7 := GenS1{F0: 2147483648}
        p0 -= (p0 % ((p0) | 1))
        _ = v7
    } else {
        var v8 uint16 = 255
        _ = v8
    }
    p0 = (p0 / ((0) | 1))
    p0 = 624
    switch "" {
    case "\n\t":
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 += (p0 * -721)
            if false {
                switch (fi9 & 127) {
                case -686:
                    var v10 bool = (128 == 2)
                    _ = v10
                case 585:
                    var v11 float32 = 3.14159
                    _ = v11
                }
                v12 := []int{1, fi9, -233}
                for fi13 := 0; fi13 < 18; fi13++ {
                    p0 += p0
                    _ = fi13
                }
                _ = v12
            } else {
                v14 := []int{429, 4294967295, (fi9 * fi9)}
                var v15 int32 = -2147483648
                fmt.Println(v15)
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                if (!(1e10 < 0.0)) {
                    var v17 string = "foo bar"
                    v18 := []int{}
                    v17 = v17
                    _ = v17
                    _ = v18
                }
                var v19 string = "a"
                _ = fi16
                _ = v19
            }
            _ = fi9
        }
    }
    p0 = p0
    _ = p0
    return -1
}
func genFn1(p0 uint16) uint64 {
    fmt.Println(p0)
    v20 := []int{1, 128, -287, 128}
    v21 := GenS2{F0: (p0 & (p0 | (^(p0)))), F1: 49}
    fmt.Println(p0)
    _ = p0
    _ = v20
    _ = v21
    return 42
}

func main() {
    var v1 uint16 = 42
    v2 := []string{}
    var v3 int = 32767
    var v4 int8 = 105
    var v5 uint = 4294967295
    v6 := []string{}
    v5 += (^(v5))
    v22 := GenS3{F0: ((v5 ^ v5) >= v5)}
    if true {
        v4 = (v4 ^ 100)
    } else {
        if false {
            v5 -= v5
            if (42 == 0) {
                v3 = (v3 & v3)
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v1 = 0
                var v23 int32 = 15
                _ = v23
            }
            v6 = append(v6, "hello")
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v24 uint16 = v1
        _ = v24
    }
    if true {
        _ = len(v6)
        if true {
            if ((-(v4)) < (v4 / ((v4) | 1))) {
                v4 += v4
                for fi25 := 0; fi25 < 17; fi25++ {
                    if ((v5 >> ((uint(v1)) & 63)) <= (v5 % (((v5 - (v5 % ((v5) | 1)))) | 1))) {
                        v26 := GenS4{F0: v5, F1: (v4 + (v4 / ((42) | 1))), F2: "a"}
                        v4 = v4
                        _ = v26
                    } else {
                        var v27 float32 = 1e10
                        _ = v27
                    }
                    v28 := GenS3{F0: (!(65535 < (v1 % ((433) | 1))))}
                    var v29 string = "a"
                    _ = fi25
                    _ = v28
                    _ = v29
                }
                if false {
                    if ((-420 <= -2147483648) || (!true)) {
                        if true {
                            _ = len(v6)
                            var v30 uint = 0
                            if false {
                                var v31 float32 = float32(3.14159)
                                if true {
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        if (1 != 2) {
                                            var v33 bool = true
                                            v34 := GenS4{F0: (v5 % ((v30) | 1)), F1: -1, F2: ("hello" + "foo bar")}
                                            if (!(!(42 != v4))) {
                                                v5 -= v30
                                                switch 797 {
                                                case 1:
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                case -9223372036854775808:
                                                    v31 = v31
                                                case -600:
                                                    v35 := GenS3{F0: true}
                                                    _ = v35
                                                }
                                            }
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v30 -= 4294967295
                                        }
                                        fmt.Println(v1)
                                        if len(v6) > 0 {
                                            v6[0] = "foo bar"
                                        }
                                        _ = fi32
                                    }
                                    v36 := GenS3{F0: true}
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 7; fi37++ {
                                    for fi38 := 0; fi38 < 12; fi38++ {
                                        var v39 int8 = v4
                                        v31 -= float32(1e10)
                                        var v40 bool = false
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                    v3 -= (-(v3))
                                    _ = fi37
                                }
                                _ = v31
                            }
                            _ = v30
                        } else {
                            var v41 bool = false
                            fmt.Println(v4, v3)
                            var v42 bool = v41
                            _ = v41
                            _ = v42
                        }
                        var v43 uint64 = 42
                        _ = v43
                    }
                    fmt.Println(v4, v1)
                    var v44 uint = (v5 | (v5 % (((v5 * (v5 / ((v5) | 1)))) | 1)))
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 15; fi45++ {
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v5 = 18446744073709551615
                            v47 := GenS3{F0: ((!(2147483647 != 2147483648)) || (("\n\t" + "hello") != "a"))}
                            fmt.Println(fi46, fi45, v3)
                            switch "\n\t" {
                            case "":
                                fmt.Println(v4, fi46, v4)
                                if false {
                                    if (false && false) {
                                        if (!false) {
                                            switch "a" {
                                            case "hello":
                                                v5 += (v5 + (^(v5)))
                                            case "":
                                                v48 := GenS3{F0: true}
                                                _ = v48
                                            case "\n\t":
                                                var v49 bool = false
                                                _ = v49
                                            default:
                                                v1 = (v1 & v1)
                                            }
                                            var v50 int64 = 42
                                            var v51 int64 = (-(v50))
                                            _ = v50
                                            _ = v51
                                        } else {
                                            fmt.Println(v5, v3, v5)
                                            v1 = (v1 * (v1 & v1))
                                        }
                                        if (float32(0.0) == (1e10 / float32((1e10 + float64(-2.25))))) {
                                            v3 += (fi45 / ((fi46) | 1))
                                            var v52 uint32 = uint32(v5)
                                            _ = v52
                                        } else {
                                            fmt.Println(fi45, v5, fi45)
                                            switch "\n\t" {
                                            case "a":
                                                var v53 int16 = 26
                                                _ = v53
                                            default:
                                                var v54 uint16 = (v1 << ((v5) & 63))
                                                v4 -= -20
                                                fmt.Println(v4, fi46)
                                                _ = v54
                                            }
                                            var v55 int8 = (v4 >> ((v5) & 63))
                                            _ = v55
                                        }
                                        v56 := GenS5{F0: 65535, F1: 1, F2: 972}
                                        _ = v56
                                    }
                                    v4 = (^(v4))
                                    switch (("a" + "\n\t") + "") {
                                    case "\n\t":
                                        v4 = 127
                                    case "foo bar":
                                        fmt.Println(v3, fi45)
                                    }
                                } else {
                                    v4 -= (v4 | v4)
                                    var v57 uint64 = 2
                                    _ = v57
                                }
                                var v58 uint32 = 57
                                _ = v58
                            case "foo bar":
                                v3 -= v3
                            }
                            v59 := GenS6{F0: v1, F1: -2147483648}
                            switch "" {
                            case "":
                                switch (fi46 / ((fi46) | 1)) {
                                case -9223372036854775808:
                                    var v60 float64 = 0.0
                                    _ = v60
                                case 723:
                                    v61 := GenS3{F0: false}
                                    _ = v61
                                case -336:
                                    v5 += 440
                                default:
                                    var v62 float32 = -2.25
                                    _ = v62
                                }
                            default:
                                var v63 int = fi45
                                _ = v63
                            }
                            _ = fi46
                            _ = v47
                            _ = v59
                        }
                        _ = fi45
                    }
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        v5 = v5
    }
    v4 = (v4 / (((v4 + v4)) | 1))
    v1 = 1
    var v64 float32 = 0.0
    if true {
        v4 += v4
        v4 += v4
        v65 := GenS4{F0: v5, F1: (v4 & (v4 - (v4 >> ((uint(v1)) & 63)))), F2: "a"}
        _ = v65
    }
    if true {
        v3 = -9223372036854775808
        v66 := GenS4{F0: v5, F1: 100, F2: ""}
        _ = v66
    }
    v2 = append(v2, "foo bar")
    if ("\n\t" < "hello") {
        var v67 int8 = v4
        var v68 int64 = 1
        v69 := GenS3{F0: (false || ((v68 + (v68 << ((v5) & 63))) <= int64(v4)))}
        _ = v67
        _ = v68
        _ = v69
    }
    v4 = v4
    switch "hello" {
    case "":
        v1 += (v1 + (v1 >> ((v5) & 63)))
    case "a":
        for fi70 := 0; fi70 < 7; fi70++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v6 = append(v6, "\n\t")
            _ = fi70
        }
    case "hello":
        v71 := GenS7{F0: "\n\t", F1: 100, F2: 1000}
        _ = v71
    default:
        v64 += ((float32(float64((1e10 - float32(-2.25)))) + (3.14159 + ((v64 - v64) * (v64 - float32((float64(v64) + 1.5)))))) / (3.14159 + float32((0.1 * ((float64(v64) / 1.0) * ((0.1 * (0.0 - -2.25)) / 1.0))))))
    }
    v4 -= 105
    for fi72 := 0; fi72 < 9; fi72++ {
        switch 32768 {
        case -9223372036854775808:
            v2 = append(v2, "")
        case 527:
            v73 := GenS7{F0: "\n\t", F1: v4, F2: 1000}
            _ = v73
        default:
            var v74 uint64 = 4294967295
            _ = v74
        }
        _ = fi72
    }
    v75 := GenS1{F0: 4294967295}
    v64 -= float32(1.5)
    for fi76 := 0; fi76 < 15; fi76++ {
        var v77 uint = v5
        if len(v2) > 0 {
            v2[0] = (("a" + ("foo bar" + "")) + "a")
        }
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v75
}
