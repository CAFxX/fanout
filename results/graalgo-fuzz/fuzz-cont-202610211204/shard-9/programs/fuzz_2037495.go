package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 string; F1 int64; F2 uint }
type GenS7 struct { F0 uint8 }
func genFn0() int {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 string = "foo bar"
        v6 := GenS1{F0: 0, F1: -1, F2: 19}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v7)
        _ = fi8
    }
    _ = v7
    return 4294967295
}
func genFn1(p0 int64) uint32 {
    var v9 int16 = 1
    var v10 uint = 42
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint = v10
        var v14 uint32 = 4294967295
        v14 -= (^(v14))
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 402
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    var v3 uint32 = 197
    v15 := GenS2{F0: "\n\t"}
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v3)
        switch (fi16 % ((-489) | 1)) {
        case 128:
            v17 := GenS1{F0: 247, F1: -519, F2: (v3 * v3)}
            _ = v17
        case -9223372036854775808:
            v1 = (v1 >> ((v1) & 63))
        default:
            switch "" {
            case "\n\t":
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v3, fi18)
                    if ((^(v1)) == (v1 * (v1 & (v1 >> ((v1) & 63))))) {
                        var v19 uint16 = 2
                        _ = v19
                    }
                    _ = fi18
                }
            case "a":
                if true {
                    if true {
                        if false {
                            var v20 float32 = 0.1
                            _ = v20
                        }
                        v3 -= (v3 % (((v3 | v3)) | 1))
                        v3 = v3
                    } else {
                        v3 -= genFn1(-9223372036854775808)
                        v1 += v1
                        if (!false) {
                            v21 := GenS1{F0: -234, F1: -2147483648, F2: v3}
                            var v22 int = genFn0()
                            _ = v21
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v3 = (v3 / (((v3 + 2)) | 1))
                                _ = fi23
                            }
                            var v24 int = (-(fi16))
                            if true {
                                var v25 int = 32767
                                v26 := GenS1{F0: -395, F1: -1, F2: (v3 << (((v1 << (((v1 ^ 926)) & 63))) & 63))}
                                _ = v25
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    var v28 int8 = 127
                                    _ = fi27
                                    _ = v28
                                }
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    if (false || (121 == 180)) {
                                        switch genFn0() {
                                        case 32767:
                                            var v30 int16 = 1
                                            _ = v30
                                        default:
                                            var v31 int64 = 42
                                            _ = v31
                                        }
                                        v32 := GenS1{F0: 42, F1: 866, F2: 930}
                                        _ = v32
                                    } else {
                                        v33 := []int{fi29}
                                        v34 := GenS1{F0: 1, F1: 42, F2: (^(v3))}
                                        _ = v33
                                        _ = v34
                                    }
                                    switch "hello" {
                                    case "hello":
                                        if (609 < 1) {
                                            var v35 string = "a"
                                            v36 := []int{(v24 | (fi29 + 0)), fi29}
                                            if true {
                                                if true {
                                                    if true {
                                                        fmt.Println(fi16, v3)
                                                        var v37 uint64 = 2
                                                        _ = v37
                                                    } else {
                                                        v1 += 42
                                                        v38 := GenS1{F0: -59, F1: -1, F2: (v3 >> (((^(v1))) & 63))}
                                                        _ = v38
                                                    }
                                                    var v39 int8 = -128
                                                    v3 = 2147483648
                                                    _ = v39
                                                }
                                                v1 = (v1 ^ (^(v1)))
                                                fmt.Println(v3, v3, fi29)
                                            } else {
                                                v3 = v3
                                                if false {
                                                    var v40 int16 = 0
                                                    _ = v40
                                                } else {
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                }
                                                var v41 uint64 = 18446744073709551615
                                                _ = v41
                                            }
                                            _ = v35
                                            _ = v36
                                        }
                                    case "foo bar":
                                        if (uint32(v1) == v3) {
                                            v1 = v1
                                        }
                                    case "a":
                                        switch (((("\n\t" + "\n\t") + "") + ("foo bar" + "\n\t")) + "foo bar") {
                                        case "hello":
                                            v3 = v3
                                        case "foo bar":
                                            v42 := GenS3{F0: 0, F1: 1, F2: 212}
                                            _ = v42
                                        case "":
                                            switch (fi29 * fi16) {
                                            case 1:
                                                var v43 int16 = -71
                                                _ = v43
                                            case 42:
                                                fmt.Println(v24, fi16)
                                            case 413:
                                                if (0 >= 42) {
                                                    v44 := GenS4{F0: "\n\t"}
                                                    _ = v44
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        switch 128 {
                                                        case 4294967295:
                                                            v45 := GenS5{F0: ("foo bar" + "a"), F1: 1, F2: (true && false)}
                                                            _ = v45
                                                        case -9223372036854775808:
                                                            if true {
                                                                var v46 float64 = (((0.1 / float64(1e10)) / 3.14159) - float64((((((1.5 * 0.0) - 3.14159) + -2.25) + 0.1) + 0.0)))
                                                                var v47 string = (((("a" + ("a" + ("\n\t" + "\n\t"))) + "\n\t") + "\n\t") + "foo bar")
                                                                var v48 int16 = 633
                                                                switch (fi29 >> (((v1 * v1)) & 63)) {
                                                                case 32767:
                                                                    v48 = (v48 - -220)
                                                                case 1:
                                                                    fmt.Println(v3, v47, v46)
                                                                case 9223372036854775807:
                                                                    v49 := GenS6{F0: ("foo bar" + v47), F1: -9223372036854775808, F2: (v1 % ((255) | 1))}
                                                                    _ = v49
                                                                }
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                var v50 int = genFn0()
                                                                v51 := GenS4{F0: ""}
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case 2147483647:
                                                            if ((((1e10 / float64(1.5)) / ((1.5 + 1e10) + 1.5)) >= -2.25) && (18446744073709551615 > 42)) {
                                                                v52 := GenS6{F0: ("" + ""), F1: 9223372036854775807, F2: v1}
                                                                v53 := []int{}
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                        default:
                                                            var v54 int8 = -124
                                                            _ = v54
                                                        }
                                                    }
                                                    v1 -= v1
                                                    if (!(127 < 11)) {
                                                        v55 := GenS3{F0: 4294967295, F1: 0, F2: 1}
                                                        _ = v55
                                                    } else {
                                                        var v56 int = (-(fi16))
                                                        _ = v56
                                                    }
                                                }
                                            default:
                                                v1 -= uint(v3)
                                            }
                                        default:
                                            var v57 uint64 = 1
                                            _ = v57
                                        }
                                    }
                                    _ = fi29
                                }
                            }
                            _ = v24
                        }
                    }
                    delete(v2, fi16)
                }
            case "foo bar":
                if true {
                    v58 := GenS7{F0: 127}
                    var v59 uint8 = 128
                    v1 = (v1 * uint(v59))
                    _ = v58
                    _ = v59
                }
            }
        }
        _ = fi16
    }
    var v60 bool = (!false)
    v61 := []int{4294967295, 128}
    fmt.Println(v3, v60, v60)
    for fi62 := 0; fi62 < 5; fi62++ {
        fmt.Println(fi62)
        v3 += genFn1(-9223372036854775808)
        _ = fi62
    }
    v60 = true
    v60 = (("hello" + "") == "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v60
    _ = v61
}
