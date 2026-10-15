package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 int; F2 uint }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    v1 := []string{"hello"}
    var v2 int64 = -1
    var v3 int32 = 42
    v4 := []int{-575, -2147483648, 9223372036854775807}
    var v5 bool = (!((786 < -764) && true))
    v6 := make(map[int]uint32)
    switch "hello" {
    case "hello":
        var v7 int8 = 127
        _ = v7
    case "a":
        v3 -= (v3 % (((v3 - v3)) | 1))
    default:
        v8 := GenS1{F0: 4294967295, F1: -852, F2: 368}
        _ = v8
    }
    var v9 int32 = (v3 * 42)
    if (v3 < v3) {
        if ("a" <= "foo bar") {
            delete(v6, 2147483647)
        }
        if v5 {
            var v10 int8 = -128
            switch 4294967295 {
            case -9223372036854775808:
                v11 := GenS2{F0: ("a" + ("foo bar" + "")), F1: -1, F2: 256}
                _ = v11
            case 2147483647:
                var v12 uint64 = 4294967295
                _ = v12
            case 125:
                v13 := v6[0]
                _ = v13
            default:
                v3 -= v3
            }
            _ = v10
        } else {
            _ = len(v6)
            if v5 {
                v9 += v3
                for fi14 := 0; fi14 < 4; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v9 = (v9 / (((-(v3))) | 1))
                        var v16 uint32 = 766
                        if false {
                            v9 = -90
                        }
                        _ = len(v6)
                        _ = fi15
                        _ = v16
                    }
                    v9 += 0
                    for fi17 := 0; fi17 < 9; fi17++ {
                        switch "" {
                        case "\n\t":
                            var v18 uint64 = 1
                            _ = v18
                        case "":
                            if (v5 && (-220 <= v9)) {
                                switch (fi14 ^ fi17) {
                                case 32767:
                                    if (3.14159 >= 3.14159) {
                                        var v19 int64 = (v2 << ((0) & 63))
                                        delete(v6, -1)
                                        v3 -= v9
                                        var v20 uint16 = 514
                                        switch "" {
                                        case "hello":
                                            var v21 uint8 = 127
                                            _ = v21
                                        }
                                        switch fi17 {
                                        case 4294967295:
                                            v19 = (v19 ^ (v19 / (((v2 % ((9223372036854775807) | 1))) | 1)))
                                        default:
                                            fmt.Println(v5, v9)
                                            switch "" {
                                            case "hello":
                                                if (!v5) {
                                                    var v22 uint16 = 65535
                                                    var v23 uint64 = 0
                                                    var v24 float32 = (1e10 * 1.5)
                                                    v1 = append(v1, "")
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                }
                                                var v25 int32 = 528
                                                _ = v25
                                            default:
                                                var v26 int32 = v9
                                                _ = v26
                                            }
                                        }
                                        v27 := GenS1{F0: 1, F1: -32768, F2: (v20 | 469)}
                                        if (((!v5) || (-88 == int8(fi14))) || (!false)) {
                                            var v28 uint16 = (v20 >> ((42) & 63))
                                            var v29 int32 = (v3 << ((4294967295) & 63))
                                            var v30 uint16 = (^(v20))
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        v3 += (^(v3))
                                        _ = v19
                                        _ = v20
                                        _ = v27
                                    }
                                case 32768:
                                    fmt.Println(v5)
                                default:
                                    var v31 int8 = int8(v9)
                                    _ = v31
                                }
                                var v32 uint8 = 127
                                if false {
                                    var v33 int16 = 1
                                    _ = v33
                                } else {
                                    if true {
                                        v9 -= v3
                                        v34 := GenS2{F0: (("foo bar" + "foo bar") + ("foo bar" + "hello")), F1: (-(fi17)), F2: 4294967295}
                                        fmt.Println(v32)
                                        _ = v34
                                    }
                                    if true {
                                        var v35 float32 = float32(1e10)
                                        var v36 int = fi17
                                        var v37 uint32 = 82
                                        switch ("a" + "foo bar") {
                                        case "hello":
                                            _ = len(v1)
                                        case "":
                                            v38 := GenS1{F0: 767, F1: 42, F2: 427}
                                            _ = v38
                                        default:
                                            if true {
                                                v37 = v37
                                                var v39 int16 = 42
                                                _ = v39
                                            } else {
                                                if v5 {
                                                    v40 := GenS1{F0: 0, F1: 42, F2: 65535}
                                                    if false {
                                                        v35 += v35
                                                    }
                                                    _ = v40
                                                } else {
                                                    v41 := GenS2{F0: ("" + "foo bar"), F1: (fi17 * fi17), F2: 267}
                                                    v42 := GenS3{F0: -58, F1: v2, F2: (v3 & (v3 & v3))}
                                                    v3 -= (v3 & (v9 ^ (v9 % ((v9) | 1))))
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                        }
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    v2 -= -182
                                }
                                _ = v32
                            }
                        case "a":
                            if v5 {
                                var v43 uint16 = 832
                                _ = v43
                            } else {
                                var v44 float64 = 1e10
                                _ = v44
                            }
                            v5 = (127 != 127)
                            var v45 int = (fi17 ^ (fi17 / ((fi17) | 1)))
                            v2 = v2
                            var v46 bool = ((((fi14 << ((1) & 63)) != -9223372036854775808) && false) || (!v5))
                            _ = v45
                            _ = v46
                        }
                        _ = len(v6)
                        var v47 string = "\n\t"
                        _ = fi17
                        _ = v47
                    }
                    _ = fi14
                }
            } else {
                if (462 == 571) {
                    _ = len(v6)
                    for fi48 := 0; fi48 < 5; fi48++ {
                        v9 += (v9 / ((1) | 1))
                        if ((v2 & 1) == v2) {
                            v49 := GenS4{F0: 255, F1: -128}
                            if v5 {
                                switch (fi48 + -1) {
                                case -214:
                                    var v50 int = fi48
                                    _ = v50
                                case 2147483647:
                                    v2 = (^(v2))
                                case 4294967295:
                                    v6[-2147483648] = 889
                                default:
                                    v51 := GenS5{F0: 42}
                                    _ = v51
                                }
                            } else {
                                v52 := GenS2{F0: "hello", F1: fi48, F2: 0}
                                _ = v52
                            }
                            _ = v49
                        } else {
                            fmt.Println(v3, v2)
                            var v53 uint64 = 2
                            _ = v53
                        }
                        _ = fi48
                    }
                    for fi54 := 0; fi54 < 3; fi54++ {
                        for fi55 := 0; fi55 < 9; fi55++ {
                            v56 := GenS6{F0: v3, F1: 255, F2: 2}
                            if v5 {
                                var v57 uint64 = 2
                                _ = len(v4)
                                _ = v57
                            } else {
                                _ = len(v6)
                                v3 = (v9 << ((65535) & 63))
                            }
                            _ = fi55
                            _ = v56
                        }
                        _ = fi54
                    }
                }
                v58 := GenS5{F0: 2}
                _ = v58
            }
        }
    } else {
        var v59 int16 = -1
        _ = v59
    }
    fmt.Println(v2, v3, v5)
    var v60 uint32 = 310
    _ = len(v6)
    v60 -= (v60 % (((v60 / (((v60 & 2147483647)) | 1))) | 1))
    fmt.Println(v60)
    v9 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v9
}
