package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 int }
type GenS7 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0(p0 int, p1 uint8, p2 int32) int16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        p1 -= (p1 / (((^(p1))) | 1))
        v8 := []int{(^(fi7)), (-(p0)), (fi7 >> ((65535) & 63)), fi7}
        _ = fi7
        _ = v8
    }
    var v9 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -32768
}

func main() {
    var v1 float64 = (0.1 + -2.25)
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 string = ""
    v5 := []string{}
    var v6 uint16 = 795
    delete(v3, v4)
    var v10 uint64 = 1
    v11 := GenS1{F0: 4294967295, F1: 2147483647, F2: -1}
    var v12 bool = false
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: 2, F1: -740, F2: 42}
            v16 := GenS2{F0: (v4 + v4), F1: v4, F2: 42}
            if false {
                v17 := GenS3{F0: v6}
                _ = v17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v18 uint = 42
        _ = fi13
        _ = v18
    }
    v19 := GenS2{F0: "hello", F1: ("\n\t" + v4), F2: 853}
    var v20 uint64 = v10
    var v21 int16 = -32768
    if v12 {
        if false {
            switch 127 {
            case -171:
                v3[((((v4 + v4) + (v4 + (("a" + v4) + "\n\t"))) + "foo bar") + v4)] = v4
            case 25:
                v22 := GenS2{F0: ((v4 + (v4 + "foo bar")) + ("\n\t" + v4)), F1: v4, F2: 2}
                _ = v22
            default:
                if (!(!v12)) {
                    v6 = (v6 | (v6 ^ v6))
                    if ("" <= "") {
                        if v12 {
                            var v23 int32 = 42
                            _ = v23
                        }
                    } else {
                        v4 = v4
                        var v24 bool = (65535 == 18446744073709551615)
                        v25 := GenS4{F0: genFn0(587, 255, int32(v20))}
                        _ = v24
                        _ = v25
                    }
                    var v26 int16 = (v21 - genFn0(int(v20), 1, -183))
                    _ = v26
                }
            }
            var v27 float64 = v1
            _ = v27
        }
        if v12 {
            if (127 > 255) {
                if (true || ((!true) && v12)) {
                    var v28 int = 127
                    var v29 int = (^(v28))
                    switch v29 {
                    case 32768:
                        var v30 float64 = 0.1
                        _ = v30
                    case 32767:
                        if (!(v12 && v12)) {
                            if v12 {
                                v29 += 535
                                if (1 >= -1) {
                                    v5 = append(v5, ((v4 + v4) + "hello"))
                                } else {
                                    if v12 {
                                        fmt.Println(v6, v21, v1)
                                        var v31 int16 = genFn0(v28, 230, 42)
                                        v31 -= genFn0(v29, 255, 0)
                                        _ = v31
                                    }
                                }
                                v32 := GenS5{F0: 900, F1: 1}
                                _ = v32
                            } else {
                                var v33 int16 = v21
                                switch 32768 {
                                case 122:
                                    if (false && v12) {
                                        v2[v4] = 4294967295
                                        switch (v4 + (("hello" + v4) + ("foo bar" + v4))) {
                                        case "a":
                                            switch 445 {
                                            case -306:
                                                v20 += (v10 / (((^(v10))) | 1))
                                            default:
                                                var v34 uint16 = 255
                                                _ = v34
                                            }
                                        default:
                                            var v35 uint64 = 468
                                            _ = v35
                                        }
                                    } else {
                                        for fi36 := 0; fi36 < 10; fi36++ {
                                            v29 = (v29 % ((fi36) | 1))
                                            _ = fi36
                                        }
                                        v12 = true
                                        var v37 uint16 = (v6 << ((65535) & 63))
                                        _ = v37
                                    }
                                }
                                if v12 {
                                    v38 := GenS2{F0: v4, F1: v4, F2: 4294967295}
                                    v29 = (v29 ^ (v28 * v29))
                                    v10 = (v20 - (^(v10)))
                                    _ = v38
                                } else {
                                    fmt.Println(v6, v20, v6)
                                    var v39 float64 = v1
                                    _ = v39
                                }
                                _ = v33
                            }
                        } else {
                            v40 := GenS5{F0: (v21 | (v21 * -32768)), F1: 1}
                            v21 = genFn0(1, 2, -506)
                            var v41 float32 = 1.5
                            _ = v40
                            _ = v41
                        }
                    default:
                        fmt.Println(v12)
                    }
                    _ = v28
                    _ = v29
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v6, v21, v6)
                }
            } else {
                v6 = v6
                var v42 float64 = v1
                _ = v42
            }
            v12 = ((32768 != -2147483648) && (2147483648 < 65535))
            v43 := v2[(v4 + "\n\t")]
            _ = v43
        } else {
            if true {
                if ("" <= v4) {
                    if (9223372036854775807 == 0) {
                        v1 = ((((-2.25 * 0.1) / ((0.1 + (v1 + v1)) / 1.5)) * v1) + -2.25)
                    } else {
                        switch v4 {
                        case "a":
                            for fi44 := 0; fi44 < 18; fi44++ {
                                var v45 float32 = float32((-2.25 * v1))
                                var v46 uint32 = 0
                                var v47 uint16 = v6
                                _ = fi44
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        default:
                            v1 -= v1
                        }
                        switch 4294967295 {
                        case -127:
                            var v48 uint16 = 240
                            _ = v48
                        case 128:
                            for fi49 := 0; fi49 < 13; fi49++ {
                                switch (-(fi49)) {
                                case 2147483647:
                                    fmt.Println(v21, v1)
                                case 127:
                                    switch fi49 {
                                    case -1:
                                        v50 := GenS6{F0: v10, F1: 0, F2: 127}
                                        _ = v50
                                    case -2147483648:
                                        if (3.14159 > 0.1) {
                                            var v51 int = fi49
                                            v52 := GenS4{F0: v21}
                                            _ = v51
                                            _ = v52
                                        } else {
                                            fmt.Println(v20)
                                            v1 += v1
                                            v21 += (v21 % (((-(v21))) | 1))
                                        }
                                    }
                                }
                                var v53 float32 = 1e10
                                v6 -= v6
                                _ = fi49
                                _ = v53
                            }
                        case 858:
                            v21 -= (v21 % ((v21) | 1))
                        default:
                            var v54 int16 = v21
                            _ = v54
                        }
                        v55 := v3[v4]
                        _ = v55
                    }
                    var v56 int64 = 1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v56
                }
                for fi57 := 0; fi57 < 16; fi57++ {
                    fmt.Println(v20, v12, v20)
                    _ = len(v5)
                    var v58 int = (fi57 | 128)
                    _ = fi57
                    _ = v58
                }
            }
            v10 += 18446744073709551615
            for fi59 := 0; fi59 < 19; fi59++ {
                fmt.Println(fi59, v12, v1)
                var v60 uint64 = (v20 % ((0) | 1))
                _ = fi59
                _ = v60
            }
        }
        for fi61 := 0; fi61 < 2; fi61++ {
            var v62 int64 = -1
            switch 121 {
            case -369:
                switch -1 {
                case 32768:
                    for fi63 := 0; fi63 < 17; fi63++ {
                        var v64 uint64 = (v10 % ((v10) | 1))
                        switch fi63 {
                        case 32768:
                            v65 := GenS4{F0: int16(v62)}
                            _ = v65
                        case 9223372036854775807:
                            var v66 int32 = 2147483647
                            _ = v66
                        case -1:
                            v67 := v3["a"]
                            _ = v67
                        default:
                            var v68 uint16 = v6
                            _ = v68
                        }
                        _ = fi63
                        _ = v64
                    }
                case 252:
                    for fi69 := 0; fi69 < 20; fi69++ {
                        switch (fi61 % ((fi61) | 1)) {
                        case -561:
                            if v12 {
                                v70 := v3["a"]
                                _ = v70
                            }
                            v4 = "a"
                            switch (fi69 << ((42) & 63)) {
                            case 42:
                                v71 := GenS7{F0: 127, F1: v20, F2: 1}
                                _ = v71
                            case -151:
                                v62 = -820
                                var v72 uint = 975
                                _ = v72
                            }
                        case -66:
                            v21 -= v21
                        }
                        _ = fi69
                    }
                default:
                    v1 += v1
                }
            case 0:
                v1 += 0.0
            }
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
