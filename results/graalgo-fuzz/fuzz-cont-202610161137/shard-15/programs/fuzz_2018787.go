package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 string; F1 int64; F2 int64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int32, p1 int32, p2 string) int {
    var v4 uint = 2
    if false {
        p0 += (p1 % ((-2147483648) | 1))
    } else {
        v5 := GenS1{F0: 97, F1: 42, F2: 255}
        switch p2 {
        case "hello":
            if (float32(0.1) <= 0.1) {
                v6 := []int{310}
                for fi7 := 0; fi7 < 12; fi7++ {
                    fmt.Println(p1)
                    v4 = 65535
                    for fi8 := 0; fi8 < 6; fi8++ {
                        p0 = 1
                        v9 := GenS1{F0: (fi7 << (((v4 + (v4 ^ (v4 >> ((v4) & 63))))) & 63)), F1: p0, F2: 128}
                        _ = fi8
                        _ = v9
                    }
                    _ = fi7
                }
                for fi10 := 0; fi10 < 20; fi10++ {
                    if ((fi10 & (fi10 - -1)) <= fi10) {
                        v11 := GenS2{F0: 117}
                        p2 = p2
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            if (42 < -9223372036854775808) {
                                p0 = 125
                            } else {
                                v6 = append(v6, (fi12 << ((780) & 63)))
                            }
                            v6 = append(v6, (fi10 ^ fi12))
                            v13 := GenS1{F0: -1, F1: p1, F2: 71}
                            _ = fi12
                            _ = v13
                        }
                        v14 := GenS1{F0: fi10, F1: p1, F2: 128}
                        p1 -= (^(p0))
                        _ = v14
                    }
                    if (1 <= 912) {
                        var v15 int = 9223372036854775807
                        _ = v15
                    }
                    _ = fi10
                }
                _ = v6
            } else {
                v16 := GenS3{F0: 58, F1: 255}
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := []int{860}
                    _ = len(v18)
                    switch fi17 {
                    case 128:
                        switch fi17 {
                        case 128:
                            var v19 uint = v4
                            _ = v19
                        }
                    case 2147483647:
                        var v20 int = (fi17 / ((int(v4)) | 1))
                        _ = v20
                    default:
                        if false {
                            if len(v18) > 0 {
                                v18[0] = 127
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                switch 897 {
                case 127:
                    v21 := []int{}
                    _ = v21
                default:
                    fmt.Println(p1, p1)
                }
                _ = v16
            }
        default:
            v4 = v4
        }
        _ = v5
    }
    v22 := GenS2{F0: 127}
    p0 += (p1 ^ int32(v4))
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v4
    return 2147483647
}
func genFn1(p0 uint32, p1 uint8, p2 int8) uint8 {
    v23 := GenS3{F0: 22, F1: 255}
    if (("foo bar" <= ("" + "foo bar")) && false) {
        v24 := GenS3{F0: (p1 * p1), F1: (^(p1))}
        var v25 int32 = -1
        v25 -= (v25 % (((v25 | v25)) | 1))
        _ = v24
        _ = v25
    }
    v26 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v26
    return (p1 >> ((993) & 63))
}

func main() {
    var v1 uint = 473
    var v2 uint64 = 18446744073709551615
    v3 := []int{-2147483648, 9223372036854775807}
    fmt.Println(v2, v1)
    v27 := GenS4{F0: ("" + ""), F1: 9223372036854775807, F2: -9223372036854775808}
    v28 := GenS5{F0: 492, F1: 4}
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 int16 = 1
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 int32 = -97
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v33 := GenS3{F0: genFn1(285, 255, 22), F1: 141}
    var v34 float64 = (0.0 - (-2.25 * float64(1e10)))
    if len(v3) > 0 {
        v3[0] = 32768
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        for fi36 := 0; fi36 < 13; fi36++ {
            fmt.Println(v2)
            if true {
                fmt.Println(v1, v2, v2)
            } else {
                switch ("" + "foo bar") {
                case "foo bar":
                    var v37 uint8 = 42
                    _ = v37
                case "\n\t":
                    fmt.Println(v2, fi36, fi35)
                }
                var v38 uint32 = 2
                v34 += float64(1.5)
                _ = v38
            }
            var v39 string = "\n\t"
            var v40 float32 = (-2.25 * 1e10)
            _ = fi36
            _ = v39
            _ = v40
        }
        v41 := GenS6{F0: -150}
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS7{F0: 447, F1: 42, F2: 42}
            v44 := GenS4{F0: "foo bar", F1: 9223372036854775807, F2: -9223372036854775808}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi35
        _ = v41
    }
    v45 := GenS6{F0: -939}
    if len(v3) > 0 {
        v3[0] = -762
    }
    var v46 uint16 = 1
    v47 := GenS7{F0: -9223372036854775808, F1: -9223372036854775808, F2: 1}
    var v48 int = 0
    var v49 uint64 = 554
    switch "foo bar" {
    case "a":
        var v50 uint32 = 57
        _ = v50
    default:
        fmt.Println(v49)
    }
    if true {
        fmt.Println(v48, v1, v48)
        if true {
            v3 = append(v3, genFn0(-912, 26, "hello"))
            var v51 uint64 = v49
            _ = v51
        } else {
            for fi52 := 0; fi52 < 2; fi52++ {
                var v53 uint64 = v2
                for fi54 := 0; fi54 < 4; fi54++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (false || true) {
                        if true {
                            _ = len(v3)
                            if true {
                                v49 = (v2 >> ((v1) & 63))
                                v34 += 1.5
                                var v55 uint32 = 1
                                v56 := GenS7{F0: -412, F1: 0, F2: genFn1((v55 - v55), 70, 14)}
                                v57 := GenS5{F0: 0, F1: 0}
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        }
                    }
                    if false {
                        v48 += fi52
                        v46 += 2
                        v1 = 637
                    } else {
                        var v58 string = "a"
                        switch v58 {
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "\n\t":
                            v59 := GenS4{F0: ((v58 + "\n\t") + v58), F1: 90, F2: 346}
                            _ = v59
                        default:
                            var v60 int8 = -128
                            _ = v60
                        }
                        _ = v58
                    }
                    _ = fi54
                }
                var v61 uint8 = 69
                _ = fi52
                _ = v53
                _ = v61
            }
        }
        if false {
            switch (v48 | (v48 + (v48 << ((v1) & 63)))) {
            case -2147483648:
                fmt.Println(v48, v49, v34)
            case 2147483647:
                for fi62 := 0; fi62 < 16; fi62++ {
                    v63 := GenS8{F0: 2147483648}
                    _ = fi62
                    _ = v63
                }
            case 537:
                v2 = v2
            default:
                var v64 int8 = 1
                _ = v64
            }
        } else {
            fmt.Println(v48, v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
