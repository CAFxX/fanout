package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint; F2 bool }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8) int32 {
    var v5 bool = false
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{}
        var v9 float32 = 0.0
        v5 = v5
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1(p0 bool, p1 int, p2 uint64) int8 {
    for fi10 := 0; fi10 < 15; fi10++ {
        p2 += (p2 % ((p2) | 1))
        v11 := GenS1{F0: (p2 % (((p2 - (p2 + p2))) | 1)), F1: 2, F2: p0}
        _ = fi10
        _ = v11
    }
    var v12 int32 = 0
    for fi13 := 0; fi13 < 20; fi13++ {
        if p0 {
            var v14 int32 = genFn0(0)
            p0 = ("a" != ("" + ("hello" + ("foo bar" + (("hello" + "hello") + "hello")))))
            _ = v14
        }
        _ = fi13
    }
    if true {
        var v15 uint8 = 2
        var v16 float64 = ((float64(1.5) * 1.5) + -2.25)
        _ = v15
        _ = v16
    } else {
        fmt.Println(p0)
        var v17 uint32 = 978
        _ = v17
    }
    p0 = (42 == 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 100
}

func main() {
    var v1 uint = 784
    var v2 uint32 = 2
    var v3 int16 = -866
    var v4 bool = false
    if v4 {
        var v18 uint64 = 1
        var v19 int32 = -1
        switch 4294967295 {
        case 4294967295:
            v20 := []int{}
            _ = v20
        case 128:
            fmt.Println(v3)
        }
        _ = v18
        _ = v19
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v2 = 27
        _ = fi21
    }
    v22 := []int{1, 1, 218}
    v22 = append(v22, 4294967295)
    switch ("" + ("\n\t" + "")) {
    case "a":
        for fi23 := 0; fi23 < 18; fi23++ {
            v4 = v4
            v1 = v1
            _ = fi23
        }
    }
    v4 = v4
    for fi24 := 0; fi24 < 12; fi24++ {
        if len(v22) > 0 {
            v22[0] = (fi24 ^ (fi24 << ((v1) & 63)))
        }
        _ = fi24
    }
    switch "\n\t" {
    case "a":
        if ((345 < v3) || true) {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v25 uint32 = (v2 << (((v1 | (v1 ^ v1))) & 63))
                _ = v25
            case "a":
                var v26 float64 = 1.5
                _ = v26
            case "foo bar":
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int = (fi27 % ((int(v1)) | 1))
                    for fi29 := 0; fi29 < 7; fi29++ {
                        _ = len(v22)
                        if (65535 >= v2) {
                            v1 -= (v1 << (((v1 % ((v1) | 1))) & 63))
                            v28 = fi29
                            v1 = 18446744073709551615
                            v2 -= 2
                        } else {
                            if v4 {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            if false {
                                var v30 uint16 = 255
                                fmt.Println(v28)
                                _ = v30
                            } else {
                                fmt.Println(fi27, fi29, v28)
                            }
                        }
                        _ = fi29
                    }
                    if (v4 || (!(!false))) {
                        v28 += v28
                        var v31 int8 = genFn1(v4, -2147483648, 4294967295)
                        v32 := GenS2{F0: -1, F1: "\n\t", F2: 820}
                        _ = v31
                        _ = v32
                    }
                    _ = fi27
                    _ = v28
                }
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                v1 -= v1
                v4 = (42 != (^(v1)))
                _ = fi33
            }
            switch -1 {
            case 475:
                if false {
                    var v34 uint = v1
                    var v35 uint16 = 65535
                    v3 = (v3 ^ (-(v3)))
                    _ = v34
                    _ = v35
                }
            }
        }
    case "foo bar":
        var v36 uint16 = 0
        _ = v36
    case "":
        if v4 {
            var v37 float32 = 0.1
            v38 := GenS3{F0: -1}
            switch -97 {
            case 111:
                v39 := GenS4{F0: 328, F1: v2}
                _ = v39
            case 2147483647:
                switch 128 {
                case 32767:
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v41 := GenS2{F0: -2147483648, F1: ("hello" + ("\n\t" + "")), F2: 255}
                        _ = fi40
                        _ = v41
                    }
                case -1:
                    switch "a" {
                    case "a":
                        v4 = ("" != "foo bar")
                    case "\n\t":
                        if (138 == 972) {
                            var v42 bool = v4
                            if v42 {
                                v43 := GenS3{F0: int64(v1)}
                                v1 -= (v1 - (v1 ^ 18446744073709551615))
                                _ = v43
                            } else {
                                var v44 uint8 = 128
                                _ = v44
                            }
                            v45 := GenS5{F0: 0}
                            _ = v42
                            _ = v45
                        } else {
                            switch (("foo bar" + "\n\t") + (("a" + ("hello" + (("a" + "hello") + ""))) + "hello")) {
                            case "":
                                fmt.Println(v2)
                            case "foo bar":
                                if false {
                                    if (282 >= 623) {
                                        for fi46 := 0; fi46 < 3; fi46++ {
                                            if (fi46 >= fi46) {
                                                v2 = v2
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            } else {
                                                fmt.Println(v2, v2, v1)
                                            }
                                            var v47 uint32 = (^(v2))
                                            v47 = v2
                                            _ = fi46
                                            _ = v47
                                        }
                                        var v48 int = 32767
                                        if (((^(v2)) == 2147483647) || v4) {
                                            for fi49 := 0; fi49 < 6; fi49++ {
                                                if len(v22) > 0 {
                                                    v22[0] = (v48 >> (((v1 % ((868) | 1))) & 63))
                                                }
                                                var v50 uint = (v1 | 545)
                                                if (9223372036854775807 != 42) {
                                                    v37 += 3.14159
                                                    v51 := GenS2{F0: -2147483648, F1: (((((("a" + "\n\t") + "\n\t") + "") + "foo bar") + "\n\t") + "hello"), F2: 8}
                                                    _ = v51
                                                }
                                                _ = fi49
                                                _ = v50
                                            }
                                        } else {
                                            v52 := GenS6{F0: (v1 ^ 1), F1: v1, F2: v4}
                                            _ = v52
                                        }
                                        _ = v48
                                    } else {
                                        var v53 int = 4294967295
                                        var v54 int8 = 1
                                        _ = v53
                                        _ = v54
                                    }
                                    switch "a" {
                                    case "hello":
                                        switch 127 {
                                        case 4294967295:
                                            v22 = append(v22, 42)
                                        default:
                                            v55 := GenS4{F0: 4294967295, F1: v2}
                                            _ = v55
                                        }
                                    case "\n\t":
                                        if (true || ((-2147483648 <= -2147483648) || (v4 || (!((v3 % ((1000) | 1)) < 42))))) {
                                            if false {
                                                if len(v22) > 0 {
                                                    v22[0] = -460
                                                }
                                                v1 -= uint(v2)
                                                var v56 int = int(v1)
                                                _ = v56
                                            } else {
                                                if (42 > v1) {
                                                    if v4 {
                                                        var v57 string = "hello"
                                                        _ = v57
                                                    }
                                                } else {
                                                    v58 := GenS2{F0: 0, F1: "a", F2: 35}
                                                    _ = v58
                                                }
                                                for fi59 := 0; fi59 < 20; fi59++ {
                                                    v3 = 0
                                                    if false {
                                                        fmt.Println(v1, v1, v2)
                                                    } else {
                                                        switch 1 {
                                                        case -9223372036854775808:
                                                            var v60 int = fi59
                                                            _ = v60
                                                        case 128:
                                                            var v61 int = fi59
                                                            _ = v61
                                                        default:
                                                            for fi62 := 0; fi62 < 18; fi62++ {
                                                                var v63 float64 = float64(v37)
                                                                v2 = v2
                                                                _ = fi62
                                                                _ = v63
                                                            }
                                                        }
                                                    }
                                                    v3 = v3
                                                    _ = fi59
                                                }
                                                var v64 int = 1
                                                _ = v64
                                            }
                                        } else {
                                            switch ((("" + "a") + "hello") + ("" + ("foo bar" + "foo bar"))) {
                                            case "foo bar":
                                                v2 += (v2 ^ v2)
                                            case "a":
                                                if v4 {
                                                    _ = len(v22)
                                                    v65 := GenS4{F0: 809, F1: (^(v2))}
                                                    _ = v65
                                                } else {
                                                    var v66 int16 = (-(v3))
                                                    for fi67 := 0; fi67 < 9; fi67++ {
                                                        v66 -= (v66 ^ (v3 & v66))
                                                        _ = fi67
                                                    }
                                                    _ = v66
                                                }
                                            case "\n\t":
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v1, v2)
                                                default:
                                                    var v68 uint64 = 42
                                                    _ = v68
                                                }
                                            default:
                                                v4 = ((v4 || (v1 <= (v1 << ((0) & 63)))) || ((!v4) || v4))
                                            }
                                        }
                                    }
                                    fmt.Println(v37)
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v4, v1)
                                    case "":
                                        v4 = v4
                                    case "a":
                                        v37 = 3.14159
                                    }
                                    var v69 int = 813
                                    _ = v69
                                }
                            case "a":
                                v70 := GenS7{F0: 1, F1: v2}
                                _ = v70
                            }
                        }
                    default:
                        for fi71 := 0; fi71 < 18; fi71++ {
                            var v72 int = (fi71 % ((fi71) | 1))
                            for fi73 := 0; fi73 < 1; fi73++ {
                                v4 = true
                                _ = fi73
                            }
                            _ = fi71
                            _ = v72
                        }
                    }
                default:
                    v37 -= (v37 + -2.25)
                }
            default:
                v74 := GenS8{F0: 0, F1: 65535}
                _ = v74
            }
            _ = v37
            _ = v38
        }
    }
    _ = len(v22)
    v2 = v2
    for fi75 := 0; fi75 < 3; fi75++ {
        var v76 uint32 = (v2 | (^(v2)))
        _ = fi75
        _ = v76
    }
    if v4 {
        for fi77 := 0; fi77 < 7; fi77++ {
            if (((^(v1)) != v1) && (v3 >= (v3 >> (((v1 & v1)) & 63)))) {
                v78 := GenS4{F0: 111, F1: (v2 >> (((^(v1))) & 63))}
                fmt.Println(v3, v2)
                v2 = v2
                _ = v78
            }
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            v79 := GenS8{F0: -9223372036854775808, F1: 951}
            _ = fi77
            _ = v79
        }
    } else {
        v22 = append(v22, -9223372036854775808)
        v2 += 560
    }
    v2 -= v2
    if (1 < 42) {
        v80 := GenS4{F0: 386, F1: v2}
        if v4 {
            v4 = true
            v81 := GenS5{F0: 90}
            _ = v81
        } else {
            var v82 int16 = 532
            for fi83 := 0; fi83 < 5; fi83++ {
                for fi84 := 0; fi84 < 18; fi84++ {
                    v1 = (^(v1))
                    fmt.Println(fi84)
                    var v85 uint8 = 185
                    _ = fi84
                    _ = v85
                }
                var v86 float64 = -2.25
                switch 42 {
                case 42:
                    v1 = (v1 | v1)
                }
                _ = fi83
                _ = v86
            }
            v22 = append(v22, 9223372036854775807)
            _ = v82
        }
        _ = v80
    }
    v1 = v1
    v2 = v2
    v22 = append(v22, -960)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
