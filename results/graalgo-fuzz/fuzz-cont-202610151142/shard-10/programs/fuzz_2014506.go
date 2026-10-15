package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int; F2 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 int8) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        p0 = p0
        v5 := GenS1{F0: 42, F1: 255, F2: 131}
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    var v6 uint16 = 229
    _ = p0
    _ = v6
    return 788
}
func genFn1(p0 int32, p1 uint32, p2 int32) uint64 {
    v7 := GenS1{F0: genFn0(1), F1: 210, F2: p2}
    var v8 string = (("a" + "\n\t") + "")
    switch v8 {
    case "a":
        var v9 int64 = 0
        _ = v9
    case "":
        v10 := []int{}
        _ = v10
    default:
        if (1 == 43) {
            if false {
                var v11 uint = uint(p1)
                if true {
                    v12 := []int{32768, -2147483648, 42}
                    p2 = (p2 & (p2 << (((v11 << ((v11) & 63))) & 63)))
                    _ = v12
                } else {
                    if (!true) {
                        if true {
                            v11 -= (v11 + v11)
                        } else {
                            v13 := GenS2{F0: 0, F1: 4294967295}
                            v8 = (v8 + (v8 + v8))
                            _ = v13
                        }
                    } else {
                        p2 += p0
                    }
                    if (true || false) {
                        if false {
                            p1 = 0
                        }
                        if false {
                            p0 += (p0 / (((^(p0))) | 1))
                        }
                    } else {
                        v14 := GenS3{F0: p1, F1: 100}
                        var v15 float32 = 0.1
                        _ = v14
                        _ = v15
                    }
                }
                if false {
                    v16 := []int{-61, 1}
                    _ = v16
                }
                _ = v11
            }
            switch 908 {
            case 229:
                v17 := GenS2{F0: 32767, F1: uint64(p1)}
                _ = v17
            case -1:
                v18 := []int{}
                _ = v18
            case 32768:
                if true {
                    var v19 int16 = 32767
                    if ((v8 > "\n\t") || true) {
                        v8 = "foo bar"
                    }
                    if true {
                        var v20 int64 = 788
                        p2 = p2
                        _ = v20
                    } else {
                        fmt.Println(v8, p0)
                        switch 128 {
                        case 117:
                            var v21 uint16 = 255
                            _ = v21
                        }
                    }
                    _ = v19
                } else {
                    if ((((float64((3.14159 - -2.25)) * (float64(3.14159) + float64(3.14159))) * float64(float32(-2.25))) != float64(1.5)) && (!false)) {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            for fi23 := 0; fi23 < 4; fi23++ {
                                var v24 uint8 = 134
                                _ = fi23
                                _ = v24
                            }
                            _ = fi22
                        }
                        var v25 int8 = 42
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v27 := []int{}
                            v25 = v25
                            var v28 uint64 = 4294967295
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    } else {
                        v29 := GenS1{F0: 255, F1: 65535, F2: p0}
                        v30 := []int{-2, 389}
                        p0 = 2147483647
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS3{F0: p1, F1: 75}
                    var v32 uint8 = 2
                    _ = v31
                    _ = v32
                }
            default:
                for fi33 := 0; fi33 < 2; fi33++ {
                    if true {
                        for fi34 := 0; fi34 < 2; fi34++ {
                            v35 := []int{1, -2147483648}
                            if false {
                                var v36 bool = true
                                _ = len(v35)
                                if false {
                                    var v37 int32 = p2
                                    v38 := GenS2{F0: 718, F1: 0}
                                    p0 -= p0
                                    _ = v37
                                    _ = v38
                                }
                                var v39 float32 = ((float32((1e10 / 3.14159)) - (1e10 * 1e10)) * 3.14159)
                                v39 -= v39
                                _ = v36
                                _ = v39
                            } else {
                                var v40 int32 = p0
                                if (413 < 101) {
                                    fmt.Println(fi34, p2, p1)
                                    fmt.Println(p1, fi33, v8)
                                } else {
                                    v40 += (v40 % ((-2147483648) | 1))
                                }
                                var v41 uint32 = 42
                                if false {
                                    var v42 int32 = v40
                                    fmt.Println(v40, fi34, p0)
                                    _ = v42
                                }
                                var v43 uint16 = 255
                                _ = v40
                                _ = v41
                                _ = v43
                            }
                            v44 := GenS2{F0: 2, F1: 2}
                            _ = fi34
                            _ = v35
                            _ = v44
                        }
                    }
                    _ = fi33
                }
            }
        }
    }
    v45 := GenS4{F0: 155}
    _ = p0
    _ = p1
    _ = p2
    _ = v45
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = float64((((1.5 - 1e10) - float32((0.0 + 1e10))) - -2.25))
    v2 := make(map[string]int)
    var v3 uint16 = 847
    switch -472 {
    case -9223372036854775808:
        if (4294967295 != 2147483647) {
            if (!true) {
                switch "a" {
                case "a":
                    for fi46 := 0; fi46 < 19; fi46++ {
                        fmt.Println(v3, fi46, v1)
                        _ = fi46
                    }
                default:
                    v3 += v3
                }
            }
            fmt.Println(v1)
        } else {
            var v47 uint32 = 0
            for fi48 := 0; fi48 < 20; fi48++ {
                switch ("foo bar" + (("foo bar" + (("" + "a") + "foo bar")) + "hello")) {
                case "hello":
                    for fi49 := 0; fi49 < 12; fi49++ {
                        if (!true) {
                            if (1 > 9223372036854775807) {
                                switch "foo bar" {
                                case "foo bar":
                                    v3 += 255
                                    v50 := GenS5{F0: (false || (false || false)), F1: 796}
                                    _ = v50
                                case "\n\t":
                                    v51 := v2["foo bar"]
                                    _ = v51
                                case "":
                                    fmt.Println(v3)
                                }
                                v47 += (^(v47))
                                v52 := []int{-15, fi48, 128, fi48}
                                _ = v52
                            }
                            if false {
                                switch 4294967295 {
                                case -328:
                                    v53 := GenS3{F0: (v47 * (v47 | (v47 + v47))), F1: -128}
                                    _ = v53
                                case 32767:
                                    if false {
                                        var v54 bool = true
                                        _ = v54
                                    } else {
                                        v2[("" + "foo bar")] = (-(fi48))
                                        if false {
                                            v3 = 1
                                            var v55 int = (fi49 & (^(fi48)))
                                            v56 := GenS3{F0: 363, F1: 106}
                                            _ = v55
                                            _ = v56
                                        }
                                    }
                                case 607:
                                    fmt.Println(v47, fi49, v47)
                                }
                                if (true && true) {
                                    v1 = (v1 + (0.0 - v1))
                                    v57 := GenS6{F0: uint32(v3), F1: fi48, F2: 65535}
                                    _ = v57
                                }
                                v58 := []int{fi49, (fi48 << ((1) & 63)), (fi49 >> ((4294967295) & 63)), -320}
                                _ = v58
                            }
                        }
                        v47 = v47
                        var v59 int8 = 42
                        _ = fi49
                        _ = v59
                    }
                }
                _ = fi48
            }
            _ = v47
        }
    case -2147483648:
        fmt.Println(v3, v1)
    }
    v60 := GenS7{F0: true}
    if false {
        fmt.Println(v1, v1)
        if false {
            v61 := v2["a"]
            switch (v61 | v61) {
            case -173:
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 uint8 = 2
                    _ = fi62
                    _ = v63
                }
            case -849:
                v61 -= v61
            default:
                v64 := GenS4{F0: v3}
                _ = v64
            }
            _ = v61
        }
    } else {
        for fi65 := 0; fi65 < 3; fi65++ {
            fmt.Println(v3, v3, v1)
            var v66 uint16 = 32767
            _ = fi65
            _ = v66
        }
        var v67 uint64 = 42
        _ = v67
    }
    fmt.Println(v1)
    v68 := []int{278, -2147483648}
    switch "hello" {
    case "hello":
        v1 = (float64(1.5) - (float64(1e10) - -2.25))
    case "foo bar":
        v1 = v1
    case "\n\t":
        fmt.Println(v3, v3)
    default:
        if ((4294967295 <= 18446744073709551615) || (true || (!((false || (42 < 2147483647)) && false)))) {
            var v69 float32 = 1.5
            _ = v69
        }
    }
    v70 := GenS2{F0: 2, F1: 4294967295}
    var v71 float64 = 1.5
    var v72 int16 = 42
    var v73 int64 = -9223372036854775808
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v68
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
