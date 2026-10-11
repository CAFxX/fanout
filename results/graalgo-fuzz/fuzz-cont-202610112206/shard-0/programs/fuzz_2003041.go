package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint64; F2 int8 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0(p0 int16, p1 uint64, p2 uint) bool {
    p1 += 0
    v7 := []int{2147483647}
    v7 = append(v7, 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 uint32, p1 bool) int64 {
    v8 := []int{}
    if p1 {
        _ = len(v8)
        v9 := GenS1{F0: 255}
        var v10 int = -1
        _ = v10
        _ = v9
    }
    if p1 {
        p1 = (!p1)
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                var v12 string = ("hello" + ("hello" + ("foo bar" + "\n\t")))
                _ = v12
            }
            _ = fi11
        }
        switch "\n\t" {
        case "foo bar":
            v13 := GenS2{F0: 1, F1: 18446744073709551615}
            _ = v13
        case "hello":
            switch 32768 {
            case 32767:
                for fi14 := 0; fi14 < 18; fi14++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi14
                }
            default:
                p0 = (p0 % ((2147483648) | 1))
            }
        default:
            if (p1 || true) {
                for fi15 := 0; fi15 < 9; fi15++ {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS3{F0: 133}
                        p1 = (1 > -1)
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS3{F0: 2}
                    if len(v8) > 0 {
                        v8[0] = fi15
                    }
                    _ = fi15
                    _ = v18
                }
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
        }
    }
    var v19 int64 = 0
    _ = p0
    _ = p1
    _ = v19
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 int = 869
    var v5 int64 = -9223372036854775808
    v6 := []string{"\n\t", ""}
    v4 -= (v4 & (v4 / ((v4) | 1)))
    var v20 string = (("\n\t" + ((((("foo bar" + "foo bar") + "\n\t") + ("hello" + "foo bar")) + "a") + ("hello" + "\n\t"))) + "hello")
    _ = len(v2)
    if false {
        for fi21 := 0; fi21 < 17; fi21++ {
            if false {
                if (32767 == v4) {
                    var v22 int8 = -55
                    switch ((((v20 + ((v20 + ("\n\t" + "foo bar")) + v20)) + v20) + ((v20 + ((((v20 + (v20 + v20)) + v20) + "\n\t") + "\n\t")) + "hello")) + "hello") {
                    case "foo bar":
                        if (true || ((-2.25 > -2.25) || false)) {
                            switch (v4 << ((18446744073709551615) & 63)) {
                            case 32767:
                                var v23 uint16 = 766
                                _ = v23
                            }
                            var v24 int8 = (v22 - 127)
                            _ = v24
                        }
                    }
                    fmt.Println(fi21, v22)
                    _ = v22
                } else {
                    if len(v6) > 0 {
                        v6[0] = (v20 + "\n\t")
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if false {
                            switch -2147483648 {
                            case -7:
                                if ((^(v5)) >= genFn1(601, (86 > 450))) {
                                    if false {
                                        switch (-(fi25)) {
                                        case 869:
                                            v26 := GenS1{F0: 127}
                                            _ = v26
                                        default:
                                            fmt.Println(fi25)
                                        }
                                        v1 = v1
                                        var v27 int16 = 0
                                        _ = v27
                                    } else {
                                        var v28 uint8 = 104
                                        _ = v28
                                    }
                                } else {
                                    var v29 float64 = 3.14159
                                    _ = v29
                                }
                            case 1:
                                var v30 float32 = 0.0
                                _ = v30
                            case 58:
                                _ = len(v6)
                            default:
                                v31 := GenS4{F0: (v20 + (v20 + "a")), F1: 0, F2: -128}
                                _ = v31
                            }
                        }
                        var v32 int16 = -322
                        v1 = v1
                        switch 2147483647 {
                        case 32768:
                            switch v4 {
                            case 128:
                                var v33 float64 = 0.0
                                _ = v33
                            }
                        case -1:
                            if true {
                                v32 = (v32 << ((4294967295) & 63))
                                if len(v6) > 0 {
                                    v6[0] = "\n\t"
                                }
                                _ = len(v6)
                            } else {
                                v4 = 128
                                if false {
                                    if false {
                                        var v34 string = v20
                                        v6 = append(v6, v34)
                                        _ = v34
                                    } else {
                                        var v35 uint8 = 255
                                        v6 = append(v6, "a")
                                        var v36 int64 = v5
                                        _ = v35
                                        _ = v36
                                    }
                                    v5 += v5
                                }
                            }
                        case 842:
                            _ = len(v6)
                        default:
                            var v37 uint64 = 18446744073709551615
                            _ = v37
                        }
                        v20 = "\n\t"
                        _ = fi25
                        _ = v32
                    }
                }
            }
            _ = fi21
        }
    } else {
        switch (-(v4)) {
        case 2147483647:
            if false {
                v1 -= v1
                var v38 float32 = -2.25
                _ = v38
            } else {
                switch (v20 + "a") {
                case "\n\t":
                    var v39 bool = (32767 <= -619)
                    _ = v39
                default:
                    var v40 float32 = 0.1
                    _ = v40
                }
                v41 := GenS5{F0: 0, F1: uint64(v1)}
                if false {
                    v5 = v5
                }
                _ = v41
            }
        case -1:
            for fi42 := 0; fi42 < 1; fi42++ {
                if true {
                    v20 = "hello"
                }
                _ = fi42
            }
        case 0:
            fmt.Println(v20, v4)
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v43 := GenS2{F0: -1, F1: 2}
        fmt.Println(v1, v4, v5)
        _ = v43
    }
    v5 = 9223372036854775807
    if true {
        switch "\n\t" {
        case "hello":
            var v44 uint32 = 4294967295
            _ = v44
        default:
            var v45 int32 = 994
            _ = v45
        }
        v5 += 199
        _ = len(v6)
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        for fi47 := 0; fi47 < 17; fi47++ {
            v5 -= 42
            switch fi46 {
            case 0:
                switch (v20 + ((v20 + (v20 + "hello")) + (v20 + v20))) {
                case "\n\t":
                    switch fi47 {
                    case 1:
                        var v48 bool = true
                        if (((v5 % (((v5 & 115)) | 1)) != v5) && false) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v49 := GenS3{F0: v1}
                            v48 = v48
                            _ = v49
                        } else {
                            switch (fi47 * fi46) {
                            case 32768:
                                var v50 int64 = (v5 >> (((^(v1))) & 63))
                                _ = v50
                            default:
                                var v51 uint = v1
                                _ = v51
                            }
                        }
                        _ = v48
                    default:
                        var v52 float32 = float32((0.0 + float64((((float32(0.0) * 3.14159) / 1.5) - 1.5))))
                        _ = v52
                    }
                default:
                    var v53 float32 = 1.5
                    _ = v53
                }
            default:
                var v54 float64 = 0.1
                _ = v54
            }
            v5 = -63
            if true {
                v55 := GenS6{F0: genFn1(0, false)}
                if true {
                    if len(v6) > 0 {
                        v6[0] = "foo bar"
                    }
                    if true {
                        v56 := v3["a"]
                        _ = v56
                    } else {
                        _ = len(v3)
                    }
                    var v57 string = v20
                    _ = v57
                } else {
                    v5 -= v5
                    var v58 int = (fi46 + fi46)
                    v20 = "\n\t"
                    _ = v58
                }
                _ = v55
            } else {
                var v59 string = v20
                v60 := GenS4{F0: "\n\t", F1: 854, F2: 127}
                _ = v59
                _ = v60
            }
            if (!false) {
                var v61 uint = 4294967295
                if true {
                    v62 := GenS1{F0: 42}
                    if true {
                        delete(v3, v20)
                        switch 128 {
                        case -20:
                            v20 = "foo bar"
                            v63 := GenS6{F0: (v5 % ((v5) | 1))}
                            var v64 int64 = genFn1(uint32(v1), (1 >= -1))
                            _ = v63
                            _ = v64
                        }
                    } else {
                        var v65 uint64 = 2
                        v4 = v4
                        _ = v65
                    }
                    _ = v62
                } else {
                    v66 := GenS7{F0: v20, F1: 256}
                    _ = v66
                }
                _ = v61
            }
            _ = fi47
        }
        _ = fi46
    }
    v5 -= (v5 >> ((v1) & 63))
    v1 -= 18446744073709551615
    if true {
        switch v4 {
        case -1:
            for fi67 := 0; fi67 < 17; fi67++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v68 uint32 = 42
                _ = fi67
                _ = v68
            }
        default:
            v69 := GenS8{F0: 65535, F1: v20}
            _ = v69
        }
        var v70 int32 = 2147483647
        _ = v70
    } else {
        v71 := GenS3{F0: v1}
        v72 := GenS9{F0: v5, F1: 1, F2: 401}
        _ = v71
        _ = v72
    }
    fmt.Println(v20, v4, v4)
    if false {
        v1 = (^(v1))
    } else {
        v5 = (v5 % (((v5 & (v5 % (((v5 - (^(v5)))) | 1)))) | 1))
        v4 += 270
    }
    v73 := GenS2{F0: 0, F1: 509}
    var v74 uint8 = 1
    if ((v74 % ((v74) | 1)) > (v74 % (((v74 - v74)) | 1))) {
        var v75 uint32 = 2147483647
        var v76 uint64 = 20
        var v77 uint32 = 225
        _ = v75
        _ = v76
        _ = v77
    } else {
        _ = len(v6)
        for fi78 := 0; fi78 < 4; fi78++ {
            fmt.Println(v74, fi78, v4)
            var v79 int8 = -126
            for fi80 := 0; fi80 < 6; fi80++ {
                v74 = v74
                var v81 int = (^(fi80))
                var v82 float64 = 1e10
                var v83 int8 = v79
                _ = fi80
                _ = v81
                _ = v82
                _ = v83
            }
            _ = fi78
            _ = v79
        }
        v5 += 0
    }
    v84 := GenS9{F0: genFn1(2147483647, ((false && true) || false)), F1: 1, F2: 727}
    var v85 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v74:", v74)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v84
    _ = v85
}
