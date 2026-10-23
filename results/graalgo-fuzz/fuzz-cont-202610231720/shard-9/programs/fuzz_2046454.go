package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS4 struct { F0 int; F1 int; F2 string }
type GenS5 struct { F0 uint8; F1 bool; F2 uint64 }

func main() {
    var v1 uint16 = 1
    v2 := []uint32{}
    var v3 int32 = 1
    v4 := make(map[int]int)
    var v5 float64 = (-2.25 * (((float64(0.0) - 1.5) + 3.14159) / 1e10))
    var v6 uint = 154
    var v7 float32 = (-2.25 - -2.25)
    fmt.Println(v5, v1)
    if (0 == 2) {
        _ = len(v2)
    } else {
        if true {
            v8 := GenS1{F0: 568, F1: 42}
            _ = v8
        }
        if (-2.25 <= 3.14159) {
            if false {
                switch "foo bar" {
                case "foo bar":
                    for fi9 := 0; fi9 < 5; fi9++ {
                        if false {
                            v5 += float64(3.14159)
                            var v10 uint = 4294967295
                            if false {
                                switch fi9 {
                                case 251:
                                    v1 = 65535
                                case -2147483648:
                                    v11 := v4[32768]
                                    _ = v11
                                case -26:
                                    _ = len(v2)
                                }
                                v10 += 496
                            }
                            _ = v10
                        }
                        for fi12 := 0; fi12 < 2; fi12++ {
                            var v13 int64 = -907
                            switch "\n\t" {
                            case "a":
                                v14 := GenS2{F0: "foo bar", F1: v1, F2: "foo bar"}
                                _ = v14
                            default:
                                var v15 uint32 = 790
                                _ = v15
                            }
                            var v16 float64 = v5
                            _ = fi12
                            _ = v13
                            _ = v16
                        }
                        _ = fi9
                    }
                default:
                    var v17 bool = (2147483647 <= 2)
                    _ = v17
                }
                var v18 uint32 = 2147483648
                _ = v18
            }
            if (-66 >= -48) {
                if false {
                    var v19 int = 704
                    _ = v19
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    if false {
                        if true {
                            v7 = 0.1
                            fmt.Println(v6, v6, fi20)
                        }
                        v2 = append(v2, uint32(v6))
                    } else {
                        if true {
                            switch 127 {
                            case -853:
                                v1 = v1
                            case 32768:
                                var v21 float64 = v5
                                _ = v21
                            case 9223372036854775807:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    var v23 uint = v6
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        } else {
                            fmt.Println(v5, v1)
                            v24 := v4[469]
                            var v25 int = fi20
                            _ = v24
                            _ = v25
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        v26 := GenS2{F0: "", F1: (^(v1)), F2: "a"}
                        _ = v26
                    }
                    _ = fi20
                }
                v1 -= (v1 << ((v6) & 63))
            }
        } else {
            v5 = 3.14159
        }
        v7 = 1.5
    }
    if true {
        if ("foo bar" < "a") {
            v27 := GenS1{F0: 9223372036854775807, F1: v6}
            switch 2147483647 {
            case -1:
                v28 := GenS2{F0: ("hello" + ""), F1: (v1 << (((^(v6))) & 63)), F2: ("\n\t" + "")}
                _ = v28
            case -740:
                if false {
                    v29 := GenS2{F0: "foo bar", F1: (v1 / (((v1 << ((v6) & 63))) | 1)), F2: ("hello" + "a")}
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 uint32 = 65535
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if true {
                                v3 = v3
                                var v33 float64 = float64((((v7 * v7) - v7) - 1.5))
                                _ = v33
                            } else {
                                if false {
                                    v34 := GenS1{F0: -564, F1: uint(v1)}
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: (v31 | v31), F1: v31, F2: -602}
                                    _ = v35
                                }
                                v31 = 42
                                var v36 uint64 = 1
                                _ = v36
                            }
                            switch fi32 {
                            case -2147483648:
                                switch "hello" {
                                case "":
                                    if false {
                                        var v37 int = 1
                                        _ = v37
                                    } else {
                                        v38 := GenS4{F0: fi30, F1: 9223372036854775807, F2: "\n\t"}
                                        if true {
                                            if false {
                                                v3 = 2147483647
                                                var v39 int32 = -632
                                                _ = v39
                                            }
                                            v6 = 966
                                            if (!true) {
                                                v40 := GenS2{F0: ("\n\t" + "a"), F1: (v1 ^ (^(v1))), F2: ("foo bar" + ("\n\t" + "\n\t"))}
                                                var v41 uint = (v6 + (v6 - 348))
                                                fmt.Println(fi32)
                                                _ = v40
                                                _ = v41
                                            }
                                        } else {
                                            if false {
                                                _ = len(v2)
                                                v42 := v4[fi32]
                                                _ = v42
                                            }
                                            switch -9223372036854775808 {
                                            case -877:
                                                v5 += v5
                                            default:
                                                v43 := GenS5{F0: 2, F1: (true || ("" == ("hello" + "\n\t"))), F2: 4294967295}
                                                _ = v43
                                            }
                                            var v44 string = ("" + "hello")
                                            _ = v44
                                        }
                                        var v45 uint8 = 1
                                        _ = v38
                                        _ = v45
                                    }
                                default:
                                    if true {
                                        v2 = append(v2, v31)
                                    }
                                }
                            }
                            _ = fi32
                        }
                        _ = fi30
                        _ = v31
                    }
                    v5 += (((-2.25 - 1.5) + (v5 / (v5 * float64((v7 / float32(((float64((v7 + 0.0)) / 0.1) * (((v5 / v5) / v5) / 3.14159)))))))) / float64(v7))
                    switch "hello" {
                    case "":
                        v1 -= (v1 % ((65535) | 1))
                    default:
                        v46 := GenS3{F0: 4294967295, F1: 2147483648, F2: 0}
                        _ = v46
                    }
                }
            }
            _ = v27
        } else {
            if (1000 < -628) {
                var v47 uint32 = 919
                _ = v47
            } else {
                _ = len(v2)
                switch ("\n\t" + "hello") {
                case "":
                    v6 = (^(v6))
                case "foo bar":
                    switch -2147483648 {
                    case -2147483648:
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                    case -1:
                        for fi48 := 0; fi48 < 4; fi48++ {
                            v7 = 0.0
                            _ = fi48
                        }
                    case 127:
                        if (127 == 81) {
                            var v49 int32 = v3
                            _ = v49
                        } else {
                            v50 := GenS2{F0: "foo bar", F1: (v1 * 32767), F2: ("hello" + "\n\t")}
                            _ = v50
                        }
                    }
                case "\n\t":
                    v3 += v3
                }
                if false {
                    for fi51 := 0; fi51 < 3; fi51++ {
                        var v52 float64 = v5
                        _ = len(v2)
                        _ = fi51
                        _ = v52
                    }
                }
            }
            var v53 float64 = v5
            _ = v53
        }
        _ = len(v2)
    } else {
        v7 = 3.14159
        v3 = (v3 >> ((4294967295) & 63))
        if false {
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 int16 = 1
                _ = fi54
                _ = v55
            }
            var v56 uint64 = 0
            _ = v56
        }
    }
    fmt.Println(v1)
    var v57 int = 128
    var v58 int32 = v3
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = float64(float32(v5))
    } else {
        for fi59 := 0; fi59 < 6; fi59++ {
            var v60 uint8 = 127
            v7 = -2.25
            _ = fi59
            _ = v60
        }
        for fi61 := 0; fi61 < 20; fi61++ {
            _ = len(v4)
            if (false && (0 <= 1)) {
                switch "hello" {
                case "foo bar":
                    var v62 uint64 = 18446744073709551615
                    _ = v62
                }
            } else {
                delete(v4, (fi61 ^ -975))
                v7 = 0.1
                var v63 int64 = 42
                _ = v63
            }
            _ = len(v4)
            _ = fi61
        }
        var v64 float64 = v5
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v7
}
