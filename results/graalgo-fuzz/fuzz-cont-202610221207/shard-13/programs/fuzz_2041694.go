package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 uint64, p1 int32, p2 bool) uint32 {
    p2 = p2
    p2 = p2
    for fi6 := 0; fi6 < 7; fi6++ {
        if (fi6 != 390) {
            v7 := GenS1{F0: 248, F1: 2}
            var v8 int16 = -32768
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int16 = v8
                v11 := GenS1{F0: 113, F1: 704}
                v12 := GenS2{F0: v10, F1: 42}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v7
            _ = v8
        } else {
            if ((2 < 0) && p2) {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS3{F0: (^(p0)), F1: 9223372036854775807, F2: 42}
                    v15 := []int{fi13, (^(fi13))}
                    p0 -= (^(p0))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                fmt.Println(fi6, fi6, p1)
                v16 := []int{(-(fi6))}
                _ = v16
            } else {
                p1 -= p1
                fmt.Println(p1)
                p0 = (p0 - p0)
            }
            fmt.Println(p2, p2)
            v17 := GenS1{F0: 255, F1: 1}
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int32 = (p1 ^ p1)
            v20 := []int{fi18}
            var v21 int16 = -112
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        if p2 {
            p1 -= p1
            var v22 int16 = 28
            if (0 != p0) {
                if p2 {
                    p1 -= 88
                    v23 := []int{}
                    p0 -= p0
                    _ = v23
                } else {
                    var v24 string = "\n\t"
                    fmt.Println(p2, fi6, p0)
                    v25 := []int{(fi6 & 0), fi6, fi6, (fi6 & (fi6 & fi6))}
                    _ = v24
                    _ = v25
                }
            }
            _ = v22
        } else {
            fmt.Println(fi6, p2, p2)
            switch (-(fi6)) {
            case 9223372036854775807:
                var v26 int8 = 100
                _ = v26
            case 32767:
                p0 = 36
            default:
                switch "hello" {
                case "hello":
                    if ((9223372036854775807 <= (fi6 << ((42) & 63))) && p2) {
                        var v27 float64 = float64((1e10 * ((-2.25 / (1e10 + (1.5 - (3.14159 / -2.25)))) / 1.0)))
                        p2 = p2
                        _ = v27
                    } else {
                        var v28 float64 = (((1e10 + 3.14159) / -2.25) * float64((1e10 * 3.14159)))
                        _ = v28
                    }
                case "\n\t":
                    for fi29 := 0; fi29 < 4; fi29++ {
                        p0 += p0
                        _ = fi29
                    }
                }
            }
            if ((((!("foo bar" >= "a")) && false) && false) || (9223372036854775807 == 943)) {
                var v30 uint64 = (p0 ^ (p0 - 42))
                var v31 float32 = 0.0
                _ = v30
                _ = v31
            }
        }
        _ = fi6
    }
    var v32 uint8 = 1
    fmt.Println(v32, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return 256
}
func genFn1(p0 string, p1 int) uint {
    switch "\n\t" {
    case "foo bar":
        v33 := GenS3{F0: 0, F1: -11, F2: 756}
        _ = v33
    case "\n\t":
        v34 := []int{}
        _ = v34
    case "":
        switch "foo bar" {
        case "hello":
            fmt.Println(p0)
        case "a":
            p1 += (p1 & 1)
        }
    }
    var v35 float32 = 1.5
    _ = p0
    _ = p1
    _ = v35
    return 757
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    v5 := []uint32{426}
    if (1 == -9223372036854775808) {
        var v36 int64 = -1
        _ = v36
    } else {
        var v37 int64 = 661
        v38 := GenS2{F0: 1, F1: (v37 | (v37 + v37))}
        v3 = (v3 + v3)
        _ = v37
        _ = v38
    }
    var v39 uint64 = 765
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 uint8 = 128
        _ = fi40
        _ = v41
    }
    var v42 float64 = 0.1
    if ((v39 | v39) == v39) {
        v43 := GenS4{F0: 0, F1: 240, F2: -1}
        _ = v43
    } else {
        var v44 int32 = 42
        fmt.Println(v42)
        v45 := GenS3{F0: v39, F1: 9223372036854775807, F2: 65535}
        _ = v44
        _ = v45
    }
    if (885 < 971) {
        var v46 bool = false
        _ = v46
    }
    var v47 int8 = 100
    if false {
        var v48 uint64 = v39
        v49 := GenS2{F0: 1, F1: -205}
        _ = v48
        _ = v49
    } else {
        var v50 uint8 = 255
        v47 += v47
        if true {
            v5 = append(v5, 487)
            v39 += (v39 & v39)
            if (true && true) {
                switch 374 {
                case 127:
                    v51 := GenS1{F0: v50, F1: genFn0((v39 - (v39 << ((434) & 63))), -2147483648, (-214 == -1))}
                    _ = v51
                }
                for fi52 := 0; fi52 < 1; fi52++ {
                    var v53 string = (("foo bar" + "foo bar") + "foo bar")
                    _ = fi52
                    _ = v53
                }
            } else {
                if false {
                    v54 := GenS1{F0: 0, F1: 2}
                    _ = v54
                }
                if true {
                    fmt.Println(v50, v39)
                }
                if false {
                    for fi55 := 0; fi55 < 2; fi55++ {
                        var v56 uint = 42
                        v57 := GenS5{F0: (^(v39))}
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    v58 := v4["\n\t"]
                    for fi59 := 0; fi59 < 4; fi59++ {
                        v58 -= -9223372036854775808
                        for fi60 := 0; fi60 < 19; fi60++ {
                            var v61 int64 = -708
                            delete(v4, ("hello" + v3))
                            _ = fi60
                            _ = v61
                        }
                        _ = fi59
                    }
                    _ = v58
                }
            }
        } else {
            switch ("foo bar" + v3) {
            case "a":
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 uint32 = 604
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi62
                    _ = v63
                }
            }
            var v64 uint = 65535
            fmt.Println(v3, v47, v3)
            _ = v64
        }
        _ = v50
    }
    v42 = v42
    if (!false) {
        for fi65 := 0; fi65 < 11; fi65++ {
            for fi66 := 0; fi66 < 16; fi66++ {
                v39 -= v39
                v39 = (v39 / ((18446744073709551615) | 1))
                v3 = v3
                _ = fi66
            }
            fmt.Println(v42, v3)
            _ = fi65
        }
        for fi67 := 0; fi67 < 13; fi67++ {
            if true {
                for fi68 := 0; fi68 < 18; fi68++ {
                    if (0.1 >= v42) {
                        fmt.Println(v42)
                        fmt.Println(v39, v3)
                    } else {
                        var v69 uint8 = 42
                        v70 := GenS6{F0: (v47 / ((v47) | 1)), F1: "a"}
                        _ = v69
                        _ = v70
                    }
                    v71 := GenS6{F0: (^(v47)), F1: (("foo bar" + v3) + (v3 + ("a" + (("hello" + v3) + v3))))}
                    _ = fi68
                    _ = v71
                }
                switch 1 {
                case 211:
                    var v72 int32 = -1
                    _ = v72
                case 32768:
                    v3 = v3
                }
                v2[208] = (v3 + "a")
            }
            var v73 float32 = (float32(v42) / ((((1e10 - (float32((float64((1.5 + -2.25)) * (v42 * v42))) - float32((v42 - -2.25)))) - 1e10) - float32((float64(((0.0 / 1e10) / (1.5 - 3.14159))) / 1e10))) - 1e10))
            var v74 int8 = (v47 | v47)
            _ = fi67
            _ = v73
            _ = v74
        }
        var v75 uint8 = 42
        _ = v75
    } else {
        delete(v4, "")
        v42 += v42
    }
    var v76 int32 = -677
    v77 := GenS7{F0: "", F1: 1, F2: v47}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v76
    _ = v77
}
