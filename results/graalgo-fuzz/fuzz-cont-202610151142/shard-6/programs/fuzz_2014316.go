package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 uint32, p2 int64) int32 {
    v7 := GenS1{F0: 18446744073709551615, F1: 255}
    p1 = (p1 ^ 1)
    switch 42 {
    case 1:
        switch "a" {
        case "":
            v8 := []int{4294967295, 42, -332, 530}
            _ = v8
        default:
            if (!(-1 > 0)) {
                v9 := []int{127, 539, 32767}
                var v10 int8 = (p0 << ((255) & 63))
                _ = v10
                _ = v9
            } else {
                var v11 int = 32767
                p2 = (^(p2))
                _ = v11
            }
        }
    case 127:
        for fi12 := 0; fi12 < 14; fi12++ {
            p2 = p2
            fmt.Println(fi12, p2, p1)
            _ = fi12
        }
    case 128:
        if (false || true) {
            fmt.Println(p1, p2, p0)
            p0 = (p0 >> ((0) & 63))
        } else {
            switch -2147483648 {
            case -1:
                fmt.Println(p0, p2, p1)
            case 42:
                v13 := GenS2{F0: 128, F1: 1, F2: -406}
                _ = v13
            }
        }
    }
    var v14 uint64 = 591
    v15 := GenS1{F0: 65535, F1: 76}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v7
    return -1
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = (42 != 1)
    v3 := make(map[int]string)
    var v4 uint32 = 1
    var v5 uint = 65535
    var v6 string = ""
    fmt.Println(v6, v4)
    var v16 uint8 = 42
    for fi17 := 0; fi17 < 9; fi17++ {
        switch 42 {
        case -312:
            var v18 int8 = 5
            _ = v18
        case 9223372036854775807:
            v19 := GenS2{F0: (v16 + (v16 & 42)), F1: 42, F2: fi17}
            _ = v19
        case 128:
            v20 := []int{(fi17 % ((fi17) | 1)), fi17, (fi17 ^ -1)}
            _ = v20
        }
        _ = fi17
    }
    if (1 <= 0) {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int16 = -838
            switch (v6 + ("\n\t" + ((v6 + "\n\t") + "a"))) {
            case "hello":
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (!(1 > (v22 * v22))) {
                        _ = len(v3)
                        v4 += 2
                        var v24 float32 = float32((1e10 / 1.0))
                        _ = v24
                    } else {
                        fmt.Println(v1, v6, v16)
                    }
                    _ = fi23
                }
            }
            if true {
                v25 := []int{fi21}
                if v2 {
                    var v26 uint = uint(v4)
                    _ = len(v3)
                    v2 = (v2 && (!v2))
                    _ = v26
                } else {
                    v6 = v6
                }
                if true {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v28 := GenS1{F0: v5, F1: v16}
                        _ = fi27
                        _ = v28
                    }
                    var v29 float32 = float32(1e10)
                    var v30 uint16 = uint16(v16)
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        if ("foo bar" != v6) {
                            v16 += (v16 >> (((v5 - 425)) & 63))
                        }
                        _ = fi31
                    }
                }
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        v32 := GenS2{F0: (v16 % (((v16 >> ((v5) & 63))) | 1)), F1: 32767, F2: 127}
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint8 = (^(v16))
        switch fi33 {
        case 32768:
            v35 := GenS1{F0: v5, F1: v34}
            _ = v35
        case -2147483648:
            delete(v3, (fi33 % (((fi33 - (fi33 / (((fi33 << ((v5) & 63))) | 1)))) | 1)))
        case 2147483647:
            switch v6 {
            case "":
                v1 = 1.5
            case "foo bar":
                var v36 uint = v5
                _ = v36
            }
        }
        v37 := GenS1{F0: (^(v5)), F1: (v34 << ((4294967295) & 63))}
        _ = fi33
        _ = v34
        _ = v37
    }
    fmt.Println(v6, v2, v2)
    var v38 uint = (v5 % (((v5 & (^(v5)))) | 1))
    for fi39 := 0; fi39 < 8; fi39++ {
        v5 += v5
        _ = fi39
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 uint32 = (v4 / ((v4) | 1))
        switch ((((v6 + (v6 + ("a" + "foo bar"))) + "") + (v6 + "\n\t")) + "foo bar") {
        case "\n\t":
            for fi42 := 0; fi42 < 11; fi42++ {
                switch (v6 + (v6 + "a")) {
                case "foo bar":
                    var v43 int64 = 42
                    _ = v43
                case "hello":
                    var v44 int8 = 42
                    _ = v44
                default:
                    v4 = v4
                }
                switch fi42 {
                case -1:
                    switch "\n\t" {
                    case "a":
                        v5 += (v5 & (^(v38)))
                    case "hello":
                        var v45 string = v6
                        _ = v45
                    case "foo bar":
                        if ((598 != 1) || ((v5 * 4294967295) <= 64)) {
                            fmt.Println(v1, v1)
                            if (!(v4 >= v4)) {
                                var v46 int16 = 0
                                v6 = v6
                                v4 -= (v41 | v41)
                                _ = v46
                            } else {
                                v41 = uint32(v16)
                                if (!false) {
                                    var v47 int64 = -419
                                    var v48 bool = true
                                    _ = v47
                                    _ = v48
                                } else {
                                    v49 := GenS3{F0: uint16(v41), F1: (^(v38))}
                                    if ((^(v16)) <= v16) {
                                        v50 := []int{(-(fi42)), (fi40 << ((42) & 63)), (fi40 | (fi40 % ((fi42) | 1))), (fi42 << ((641) & 63))}
                                        _ = v50
                                    } else {
                                        v51 := []int{fi40, fi40}
                                        _ = v51
                                    }
                                    var v52 uint32 = (v41 * (v4 & 2147483647))
                                    _ = v49
                                    _ = v52
                                }
                                v53 := []int{}
                                v54 := v3[fi40]
                                if len(v53) >= 2 {
                                    v53 = v53[1:len(v53)]
                                }
                                _ = v53
                                _ = v54
                            }
                            var v55 float64 = (1.5 - 1e10)
                            v56 := GenS3{F0: 942, F1: (^(v5))}
                            _ = v55
                            _ = v56
                        }
                    default:
                        v57 := GenS4{F0: v6, F1: v2}
                        _ = v57
                    }
                case -295:
                    var v58 bool = (!(((1.5 / float64(-2.25)) * ((0.1 - 1.5) / (0.1 + 1e10))) != ((1e10 / 1.5) - 0.1)))
                    v16 += (^(v16))
                    var v59 int64 = 1
                    _ = v58
                    _ = v59
                case 736:
                    v60 := GenS5{F0: (v6 + "")}
                    _ = v60
                }
                if v2 {
                    if v2 {
                        v6 = v6
                        v5 += 986
                        v61 := GenS1{F0: (v38 + (^(v38))), F1: (^(v16))}
                        _ = v61
                    }
                    v4 -= v41
                    v62 := GenS1{F0: 18446744073709551615, F1: v16}
                    _ = v62
                } else {
                    if v2 {
                        if v2 {
                            v16 = v16
                            v16 = (v16 * (v16 * v16))
                            var v63 int16 = 32767
                            _ = v63
                        }
                        var v64 int = (fi40 >> ((v5) & 63))
                        _ = v64
                    }
                }
                _ = fi42
            }
        default:
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 uint64 = 2
                v67 := GenS3{F0: 42, F1: v38}
                if v2 {
                    v66 -= uint64(v5)
                    if (!false) {
                        v41 -= (v41 ^ 512)
                    } else {
                        v68 := GenS4{F0: "foo bar", F1: v2}
                        v2 = v2
                        var v69 float32 = 0.1
                        _ = v68
                        _ = v69
                    }
                } else {
                    v16 += (v16 >> ((18446744073709551615) & 63))
                    v2 = true
                }
                _ = fi65
                _ = v66
                _ = v67
            }
        }
        _ = fi40
        _ = v41
    }
    var v70 int16 = 0
    fmt.Println(v16)
    if (v2 || (1 == 42)) {
        for fi71 := 0; fi71 < 19; fi71++ {
            v3[(fi71 ^ 32767)] = v6
            v16 -= v16
            var v72 bool = v2
            _ = fi71
            _ = v72
        }
    } else {
        if (1 <= -107) {
            fmt.Println(v70, v38, v70)
        }
    }
    fmt.Println(v2, v4)
    var v73 uint8 = (^(v16))
    v6 = v6
    v1 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v73
}
