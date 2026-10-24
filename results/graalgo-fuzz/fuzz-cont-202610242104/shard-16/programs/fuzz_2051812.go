package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int; F1 string; F2 uint8 }
type GenS8 struct { F0 bool; F1 int64 }
func genFn0() int32 {
    if true {
        var v6 uint = 4294967295
        v6 = (^(v6))
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int = 9223372036854775807
            if (true && false) {
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 bool = (int8(fi9) > 0)
                    _ = fi9
                    _ = v10
                }
                v8 -= (fi7 | 127)
            }
            switch "hello" {
            case "a":
                v11 := []int{v8, (v8 | (v8 - v8)), (v8 * (-(v8))), v8}
                _ = v11
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v12 := GenS1{F0: -1}
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint32 = 2147483648
        _ = fi13
        _ = v14
    }
    if true {
        var v15 float32 = float32(float64((1e10 / 1.0)))
        fmt.Println(v15)
        _ = v15
    }
    var v16 uint = 65535
    _ = v16
    return 0
}
func genFn1(p0 uint64) bool {
    v17 := GenS1{F0: genFn0()}
    if (1e10 >= -2.25) {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: 2147483647}
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int64 = 42
                v22 := []int{fi18, (fi20 ^ 128)}
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi18
            _ = v19
        }
        fmt.Println(p0)
        fmt.Println(p0)
    }
    var v23 uint64 = (^(p0))
    _ = p0
    _ = v17
    _ = v23
    return (-478 != 1)
}

func main() {
    v1 := []uint32{1}
    var v2 float64 = 3.14159
    var v3 int16 = 228
    v4 := []string{"", "hello"}
    var v5 uint64 = 1
    var v24 int64 = -739
    v24 += int64(v3)
    v25 := GenS2{F0: uint32(v5), F1: v24}
    var v26 bool = (((^(v5)) < v5) && false)
    if (0 == -818) {
        v24 = v24
    }
    if (v2 > (1.5 + v2)) {
        for fi27 := 0; fi27 < 11; fi27++ {
            v28 := GenS1{F0: -2147483648}
            fmt.Println(v26, v24)
            v29 := GenS1{F0: 994}
            _ = fi27
            _ = v28
            _ = v29
        }
        if (0 < 228) {
            for fi30 := 0; fi30 < 5; fi30++ {
                if (1.5 < 1e10) {
                    switch "foo bar" {
                    case "foo bar":
                        if v26 {
                            v5 = 4294967295
                        }
                    case "":
                        _ = len(v4)
                    case "hello":
                        v31 := GenS3{F0: v5}
                        _ = v31
                    }
                    fmt.Println(v3, v5, v26)
                } else {
                    switch "\n\t" {
                    case "hello":
                        v3 = (v3 << ((765) & 63))
                    case "":
                        var v32 uint16 = 2
                        _ = v32
                    }
                    var v33 uint = 18446744073709551615
                    _ = v33
                }
                _ = fi30
            }
            v5 += 55
            var v34 int32 = -2147483648
            _ = v34
        }
        v5 += v5
    }
    v26 = (2 == (v5 ^ v5))
    switch -268 {
    case 42:
        for fi35 := 0; fi35 < 17; fi35++ {
            for fi36 := 0; fi36 < 14; fi36++ {
                fmt.Println(fi35, v2, fi36)
                v3 = v3
                _ = fi36
            }
            if (!v26) {
                switch (fi35 << ((1) & 63)) {
                case 32768:
                    v37 := GenS1{F0: 1}
                    _ = v37
                default:
                    var v38 float64 = 1.5
                    _ = v38
                }
            }
            _ = fi35
        }
    default:
        for fi39 := 0; fi39 < 7; fi39++ {
            if ((v24 / (((v24 / ((v24) | 1))) | 1)) > 1) {
                var v40 uint32 = 4294967295
                if false {
                    var v41 bool = ((v5 | 2) >= 18446744073709551615)
                    if (18446744073709551615 >= 465) {
                        if (42 < 797) {
                            switch fi39 {
                            case -9223372036854775808:
                                fmt.Println(v24, v26, v41)
                            case 0:
                                v42 := GenS4{F0: "hello", F1: (v40 * (v40 | (v40 ^ (^(v40))))), F2: 2}
                                _ = v42
                            case 2147483647:
                                var v43 uint16 = 2
                                _ = v43
                            default:
                                switch "hello" {
                                case "\n\t":
                                    var v44 uint16 = 255
                                    _ = v44
                                case "":
                                    var v45 float64 = v2
                                    _ = v45
                                case "foo bar":
                                    var v46 float64 = 1e10
                                    _ = v46
                                }
                            }
                        }
                        v47 := GenS5{F0: 1}
                        _ = v47
                    }
                    v48 := GenS5{F0: 114}
                    _ = v41
                    _ = v48
                } else {
                    v26 = (1 != -2147483648)
                }
                _ = v40
            } else {
                switch (fi39 | fi39) {
                case 1:
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v50 := GenS6{F0: ("" + "\n\t"), F1: v5, F2: v5}
                        if v26 {
                            var v51 float64 = v2
                            v51 = v51
                            if (1 >= 514) {
                                v26 = (v51 <= v2)
                                if (243 != 2) {
                                    var v52 bool = (v24 == v24)
                                    v53 := GenS4{F0: "foo bar", F1: 206, F2: 300}
                                    _ = v52
                                    _ = v53
                                }
                                if v26 {
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v2)
                                    case "\n\t":
                                        var v54 int8 = 53
                                        _ = v54
                                    default:
                                        var v55 uint64 = v5
                                        _ = v55
                                    }
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 4294967295
                                    }
                                }
                            }
                            _ = v51
                        } else {
                            v2 = v2
                        }
                        _ = fi49
                        _ = v50
                    }
                case -208:
                    v56 := GenS6{F0: (((("" + "a") + "") + "\n\t") + ("foo bar" + "hello")), F1: 0, F2: (^(v5))}
                    _ = v56
                case 9223372036854775807:
                    v3 = v3
                }
            }
            if v26 {
                if v26 {
                    v5 += v5
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v58 := GenS7{F0: fi39, F1: "\n\t", F2: 255}
                        if v26 {
                            fmt.Println(v24, v2)
                            v26 = (0 >= 219)
                            v5 = (v5 | (^(v5)))
                            v1 = append(v1, 65535)
                            switch "\n\t" {
                            case "\n\t":
                                var v59 int32 = 1
                                _ = v59
                            case "a":
                                var v60 uint8 = 160
                                _ = v60
                            default:
                                v5 = v5
                            }
                        }
                        _ = fi57
                        _ = v58
                    }
                } else {
                    v3 -= (v3 - -922)
                    fmt.Println(v2, v2)
                    var v61 bool = false
                    _ = v61
                }
                _ = len(v1)
                for fi62 := 0; fi62 < 17; fi62++ {
                    var v63 uint8 = 209
                    v24 -= (v24 - (v24 / ((v24) | 1)))
                    _ = fi62
                    _ = v63
                }
            }
            _ = fi39
        }
    }
    v64 := GenS4{F0: "\n\t", F1: 407, F2: 2}
    switch 1 {
    case 1:
        if true {
            if len(v1) > 0 {
                v1[0] = 42
            }
            fmt.Println(v5, v24)
        } else {
            var v65 int8 = -128
            _ = v65
        }
    case 127:
        v2 += (-2.25 + v2)
    default:
        v5 = (v5 - (v5 * (v5 | (v5 ^ (v5 / ((2) | 1))))))
    }
    v5 = (^(v5))
    var v66 string = "foo bar"
    v67 := GenS8{F0: (1 > -2147483648), F1: (^(v24))}
    v1 = append(v1, 1)
    var v68 float32 = -2.25
    v26 = v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v66
    _ = v67
    _ = v68
}
