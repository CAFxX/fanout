package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32; F1 string; F2 uint64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 string; F2 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint32) string {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = 42
        var v9 uint8 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        if (2 >= 840) {
            switch "" {
            case "foo bar":
                p0 = 2
            default:
                v10 := []int{32767}
                _ = v10
            }
        }
        if (4294967295 > (p0 % ((391) | 1))) {
            var v11 int32 = -2147483648
            _ = v11
        }
        switch (("hello" + "a") + "foo bar") {
        case "":
            fmt.Println(p0)
        }
    }
    _ = p0
    return "foo bar"
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 uint16 = 2
    v5 := []int{-591}
    var v6 uint8 = 204
    v5 = append(v5, -474)
    v12 := v2[4294967295]
    v1 -= v1
    var v13 uint = 42
    switch "" {
    case "\n\t":
        var v14 int = -601
        _ = v14
    case "":
        fmt.Println(v13)
    default:
        var v15 uint32 = v12
        _ = v15
    }
    if false {
        v16 := GenS1{F0: (v4 / ((v4) | 1)), F1: (v12 - v12)}
        _ = v16
    } else {
        v17 := GenS1{F0: v4, F1: (v12 % ((uint32(v6)) | 1))}
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: -128}
    v12 -= v12
    for fi19 := 0; fi19 < 20; fi19++ {
        v6 = (v6 / ((v6) | 1))
        switch "hello" {
        case "":
            fmt.Println(v3, v4)
        }
        _ = fi19
    }
    var v20 uint64 = v1
    if (int32(v13) >= -724) {
        switch 127 {
        case 836:
            v1 += v20
        case -578:
            if (128 < -1) {
                fmt.Println(v20)
                fmt.Println(v12, v6, v4)
            }
        case -2147483648:
            v21 := GenS2{F0: 2147483647, F1: -1}
            _ = v21
        default:
            var v22 uint64 = 2
            _ = v22
        }
        v3 = (1 <= 100)
        v23 := GenS3{F0: 42, F1: "", F2: (v20 << (((v13 + v13)) & 63))}
        _ = v23
    } else {
        if (v3 && (!((1.5 + ((1e10 * ((0.0 / ((1.5 / float64(1.5)) + -2.25)) - (((0.1 + (0.0 - -2.25)) - 1.5) * 1e10))) + float64(-2.25))) <= 0.0))) {
            fmt.Println(v4, v4, v3)
            fmt.Println(v12)
        } else {
            switch -2147483648 {
            case 127:
                var v24 uint16 = uint16(v6)
                _ = v24
            case 4294967295:
                var v25 int16 = -32768
                _ = v25
            case 0:
                if v3 {
                    v2[65535] = (v12 / ((v12) | 1))
                    var v26 uint = (v13 ^ v13)
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                switch (("a" + "foo bar") + "foo bar") {
                case "\n\t":
                    v20 += 18446744073709551615
                default:
                    v4 = v4
                }
                v28 := v2[v12]
                switch fi27 {
                case -9223372036854775808:
                    var v29 int8 = -98
                    _ = v29
                case 42:
                    var v30 bool = v3
                    _ = v30
                case 9223372036854775807:
                    v12 -= 2
                }
                _ = fi27
                _ = v28
            }
            switch 9223372036854775807 {
            case -2147483648:
                v31 := GenS1{F0: v4, F1: 2147483647}
                _ = v31
            case 32767:
                for fi32 := 0; fi32 < 2; fi32++ {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v34 := GenS4{F0: v13, F1: 1}
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS2{F0: 2147483647, F1: 100}
                    switch fi32 {
                    case -2147483648:
                        if true {
                            switch "" {
                            case "hello":
                                if v3 {
                                    v36 := GenS3{F0: 143, F1: "\n\t", F2: v20}
                                    var v37 int64 = 0
                                    _ = v36
                                    _ = v37
                                } else {
                                    v20 -= v20
                                    v13 = (^(v13))
                                }
                            }
                            v38 := GenS5{F0: int64(v12), F1: 338, F2: v13}
                            for fi39 := 0; fi39 < 2; fi39++ {
                                v5 = append(v5, 32767)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi39
                            }
                            _ = v38
                        } else {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                var v41 uint32 = (^(v12))
                                _ = fi40
                                _ = v41
                            }
                        }
                    default:
                        if (!(0 == -1)) {
                            if (v3 || v3) {
                                var v42 uint16 = 255
                                v43 := v2[364]
                                _ = v42
                                _ = v43
                            } else {
                                v44 := GenS6{F0: (v20 * v1)}
                                _ = v44
                            }
                            if (!(!v3)) {
                                switch -693 {
                                case 0:
                                    switch ("" + ("\n\t" + ((("hello" + "\n\t") + (("hello" + "foo bar") + "\n\t")) + ""))) {
                                    case "a":
                                        v45 := GenS7{F0: (fi32 ^ 127), F1: "\n\t", F2: (v6 << (((v13 << ((v13) & 63))) & 63))}
                                        _ = v45
                                    case "hello":
                                        var v46 float64 = 1.5
                                        _ = v46
                                    case "":
                                        v4 += 723
                                    default:
                                        if false {
                                            var v47 int64 = 40
                                            v48 := v2[581]
                                            var v49 int64 = v47
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 float32 = 0.1
                                            v51 := GenS3{F0: 42, F1: ("" + ("a" + (((("foo bar" + (("\n\t" + "foo bar") + "a")) + ("a" + "\n\t")) + "") + ""))), F2: 42}
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                case 128:
                                    fmt.Println(v12, fi32, v3)
                                }
                                for fi52 := 0; fi52 < 15; fi52++ {
                                    switch (fi52 >> (((v13 | (^(v13)))) & 63)) {
                                    case 597:
                                        var v53 uint64 = (v1 & v20)
                                        _ = v53
                                    case -2147483648:
                                        if v3 {
                                            v3 = (42 > 0)
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                v54 := GenS8{F0: ((int8(v6) >= 0) && true)}
                                                _ = v54
                                            }
                                        }
                                    case 128:
                                        var v55 string = ("" + "\n\t")
                                        _ = v55
                                    default:
                                        if v3 {
                                            if v3 {
                                                fmt.Println(v4)
                                            }
                                        }
                                    }
                                    var v56 float64 = ((0.1 + float64(float32(3.14159))) / 3.14159)
                                    _ = fi52
                                    _ = v56
                                }
                            } else {
                                var v57 uint64 = v20
                                _ = v57
                            }
                        }
                    }
                    _ = fi32
                    _ = v35
                }
            }
        }
        var v58 int64 = 0
        _ = v58
    }
    var v59 bool = true
    if (uint(v6) > (v13 >> ((v13) & 63))) {
        v6 = (^(v6))
        var v60 int64 = 1
        v2[267] = 311
        _ = v60
    } else {
        v6 += (v6 & uint8(v13))
        var v61 float64 = 1.5
        if ((v4 - v4) > uint16(v6)) {
            switch 32768 {
            case 79:
                for fi62 := 0; fi62 < 18; fi62++ {
                    if v3 {
                        if v3 {
                            for fi63 := 0; fi63 < 12; fi63++ {
                                v64 := GenS6{F0: (v20 - v20)}
                                v65 := GenS7{F0: (fi63 | fi63), F1: "a", F2: (v6 << (((v13 + v13)) & 63))}
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            for fi66 := 0; fi66 < 5; fi66++ {
                                v67 := GenS8{F0: (!v3)}
                                _ = fi66
                                _ = v67
                            }
                        }
                    } else {
                        fmt.Println(fi62)
                    }
                    if v59 {
                        var v68 uint16 = v4
                        v69 := GenS8{F0: (1 > v6)}
                        _ = v68
                        _ = v69
                    }
                    _ = fi62
                }
            case 128:
                switch "foo bar" {
                case "hello":
                    v20 += 0
                case "a":
                    v59 = (int(v1) != 128)
                }
            case 1:
                var v70 uint = 18446744073709551615
                _ = v70
            default:
                if (v59 || false) {
                    fmt.Println(v13)
                    v71 := GenS9{F0: -9223372036854775808, F1: (v6 + (v6 ^ uint8(v12)))}
                    var v72 uint32 = 2
                    _ = v71
                    _ = v72
                } else {
                    v3 = (!(true || (false && (v3 || (v3 || v3)))))
                    v2[65535] = 920
                }
            }
            fmt.Println(v6, v3, v20)
            v2[(v12 ^ 42)] = 2147483647
        }
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
