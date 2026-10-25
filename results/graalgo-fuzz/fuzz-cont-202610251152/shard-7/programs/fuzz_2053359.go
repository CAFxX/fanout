package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int32) int {
    for fi7 := 0; fi7 < 2; fi7++ {
        p0 = (^(p0))
        _ = fi7
    }
    v8 := []int{-552}
    if false {
        v8 = append(v8, 0)
    } else {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            p0 = 0
        }
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    if (!true) {
        if false {
            if true {
                v8 = append(v8, 179)
            }
        }
        v10 := GenS1{F0: 32767, F1: "a"}
        fmt.Println(p0)
        _ = v10
    }
    if (!(!((true || true) && true))) {
        var v11 int8 = -111
        switch "" {
        case "":
            if true {
                if ("" == (("hello" + "a") + "\n\t")) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 float32 = float32((float64(1e10) + 1.5))
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(p0)
                    v11 += v11
                } else {
                    p0 = p0
                }
            } else {
                for fi14 := 0; fi14 < 9; fi14++ {
                    p0 += (p0 >> ((4294967295) & 63))
                    _ = fi14
                }
                fmt.Println(p0)
                var v15 string = "\n\t"
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 9; fi16++ {
                switch (("\n\t" + (("a" + ("foo bar" + "foo bar")) + ("hello" + ""))) + (("" + "a") + "a")) {
                case "":
                    v11 -= (-(v11))
                default:
                    var v17 int = (fi16 & fi16)
                    _ = v17
                }
                var v18 uint64 = 4294967295
                var v19 int64 = -9223372036854775808
                _ = fi16
                _ = v18
                _ = v19
            }
        }
        p0 += -710
        _ = v11
    }
    _ = p0
    _ = v8
    return 608
}

func main() {
    var v1 int32 = -109
    var v2 string = ("a" + "a")
    var v3 uint8 = 153
    var v4 uint32 = 2
    var v5 uint64 = 42
    v6 := make(map[int]uint32)
    v5 = v5
    v1 = (v1 ^ v1)
    v5 += v5
    _ = len(v6)
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint16 = 0
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := []int{(fi22 * 127), fi22, 1}
        _ = fi22
        _ = v23
    }
    var v24 int64 = -9223372036854775808
    var v25 uint = 832
    switch v2 {
    case "":
        v26 := GenS1{F0: 0, F1: "foo bar"}
        _ = v26
    case "foo bar":
        v2 = "foo bar"
    case "hello":
        switch genFn0((v1 * v1)) {
        case 127:
            for fi27 := 0; fi27 < 15; fi27++ {
                switch v2 {
                case "hello":
                    if (v24 != (v24 ^ -464)) {
                        v28 := []int{(fi27 / ((fi27) | 1)), (fi27 - 32768), (fi27 ^ 1)}
                        var v29 float64 = float64(-2.25)
                        var v30 uint16 = 1
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                case "a":
                    v5 -= (v5 - v5)
                case "foo bar":
                    v31 := v6[(fi27 & genFn0((-(v1))))]
                    _ = v31
                default:
                    switch (fi27 ^ (fi27 % (((fi27 & (fi27 / (((fi27 & (fi27 & 165))) | 1)))) | 1))) {
                    case 127:
                        switch 32768 {
                        case -217:
                            v6[(fi27 - (fi27 / ((genFn0(int32(v24))) | 1)))] = 2147483647
                        case 1:
                            for fi32 := 0; fi32 < 14; fi32++ {
                                v33 := GenS2{F0: 0, F1: 255, F2: (fi27 << (((v25 & (v25 << ((v25) & 63)))) & 63))}
                                v34 := []int{fi32, (fi32 / (((fi32 + 396)) | 1)), (fi32 * genFn0(v1))}
                                v35 := GenS3{F0: v25, F1: v24}
                                _ = len(v6)
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        case -9223372036854775808:
                            v36 := GenS1{F0: 113, F1: ("" + v2)}
                            _ = v36
                        default:
                            if true {
                                v3 += (v3 % ((38) | 1))
                                switch genFn0(0) {
                                case -2147483648:
                                    v4 -= v4
                                case 127:
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        var v38 uint64 = v5
                                        _ = fi37
                                        _ = v38
                                    }
                                case 964:
                                    var v39 int16 = -1
                                    _ = v39
                                }
                            } else {
                                if ((false && (v3 <= (^(v3)))) || (0.1 > float64(-2.25))) {
                                    v40 := GenS1{F0: -472, F1: ""}
                                    v41 := GenS1{F0: 1, F1: v2}
                                    var v42 uint64 = (v5 ^ (^(v5)))
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                    default:
                        _ = len(v6)
                    }
                }
                if false {
                    if (!false) {
                        var v43 int16 = 536
                        v2 = (v2 + "foo bar")
                        var v44 float32 = 1e10
                        _ = v43
                        _ = v44
                    } else {
                        if true {
                            var v45 int32 = v1
                            _ = v45
                        }
                        v24 = (v24 % (((v24 / ((-1) | 1))) | 1))
                        switch fi27 {
                        case 494:
                            fmt.Println(v2, fi27)
                        case -74:
                            var v46 float32 = (3.14159 / (0.1 + -2.25))
                            _ = v46
                        default:
                            v25 -= (^(v25))
                        }
                    }
                    var v47 int32 = (v1 | 575)
                    var v48 uint = 98
                    _ = v47
                    _ = v48
                } else {
                    v5 = v5
                    for fi49 := 0; fi49 < 20; fi49++ {
                        switch "\n\t" {
                        case "":
                            var v50 int64 = v24
                            _ = v50
                        case "\n\t":
                            if (!false) {
                                v51 := []int{(-(fi27))}
                                _ = v51
                            } else {
                                switch v2 {
                                case "a":
                                    v2 = v2
                                default:
                                    var v52 uint32 = uint32(v5)
                                    _ = v52
                                }
                                switch (v2 + v2) {
                                case "hello":
                                    var v53 uint32 = v4
                                    _ = v53
                                case "":
                                    var v54 uint8 = (v3 - uint8(v25))
                                    _ = v54
                                default:
                                    v55 := GenS4{F0: (v24 & 854), F1: (v24 + (v24 | -786))}
                                    _ = v55
                                }
                                fmt.Println(v5)
                                v56 := v6[fi49]
                                _ = v56
                            }
                        }
                        v25 += 2
                        if false {
                            switch ("\n\t" + v2) {
                            case "hello":
                                fmt.Println(v3, v24, v2)
                            case "":
                                v4 -= 2
                            case "a":
                                if true {
                                    var v57 bool = true
                                    if v57 {
                                        v58 := []int{}
                                        if v57 {
                                            var v59 uint32 = (v4 - v4)
                                            if len(v58) >= 2 {
                                                v58 = v58[1:len(v58)]
                                            }
                                            _ = v59
                                        }
                                        if (fi27 < (fi49 & genFn0((v1 * 167)))) {
                                            var v60 uint32 = uint32(v5)
                                            var v61 string = (v2 + v2)
                                            v62 := GenS3{F0: (v25 + (v25 << ((v25) & 63))), F1: (v24 | v24)}
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                        _ = v58
                                    } else {
                                        v24 = 9223372036854775807
                                    }
                                    v63 := GenS3{F0: (v25 / (((v25 | v25)) | 1)), F1: v24}
                                    var v64 int64 = v24
                                    v65 := GenS1{F0: 514, F1: "hello"}
                                    _ = v57
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            }
                            var v66 float32 = 3.14159
                            _ = v66
                        }
                        v67 := GenS5{F0: (v25 >> ((uint(v3)) & 63))}
                        v68 := []int{}
                        _ = fi49
                        _ = v67
                        _ = v68
                    }
                }
                _ = fi27
            }
        case -1:
            switch genFn0(v1) {
            case -417:
                v69 := GenS6{F0: (v3 - v3)}
                _ = v69
            case 128:
                v5 -= 42
            case -1:
                v3 = v3
            default:
                v70 := GenS4{F0: v24, F1: (v24 ^ -1)}
                _ = v70
            }
        default:
            v2 = ""
        }
    }
    v71 := v6[0]
    if (!(!true)) {
        v72 := []int{127, genFn0(v1), 271}
        _ = v72
    }
    var v73 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v73
}
