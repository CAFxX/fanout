package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int64; F2 string }
type GenS7 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16, p1 int, p2 int8) int16 {
    v6 := []int{(p1 << ((1) & 63)), -311}
    v7 := GenS1{F0: (p1 * (p1 + (p1 % ((2147483647) | 1)))), F1: (-(p1))}
    var v8 float32 = (3.14159 / (0.1 / 0.1))
    v9 := GenS1{F0: p1, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -204
}
func genFn1() string {
    v10 := []int{-44, 32767, -608}
    switch ("foo bar" + "a") {
    case "hello":
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case "a":
        v11 := GenS1{F0: -932, F1: 474}
        _ = v11
    case "\n\t":
        var v12 uint = 18446744073709551615
        _ = v12
    }
    _ = v10
    return (("foo bar" + "\n\t") + "hello")
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -103
    v3 := []string{""}
    v4 := make(map[int]int)
    var v5 bool = true
    var v13 int8 = v2
    fmt.Println(v5)
    var v14 uint32 = 378
    if v5 {
        v2 += v13
        v15 := GenS1{F0: -1, F1: 1}
        v2 += v2
        _ = v15
    } else {
        var v16 uint8 = 0
        if false {
            if v5 {
                switch (("\n\t" + "foo bar") + "a") {
                case "foo bar":
                    if false {
                        if true {
                            if v5 {
                                var v17 uint8 = (v16 << ((65535) & 63))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v18 int = 9223372036854775807
                                _ = v17
                                _ = v18
                            } else {
                                var v19 bool = v5
                                _ = v19
                            }
                            if v5 {
                                if ((("" + "") < "hello") && false) {
                                    fmt.Println(v2)
                                } else {
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        for fi21 := 0; fi21 < 4; fi21++ {
                                            fmt.Println(v16, v13, v2)
                                            _ = fi21
                                        }
                                        fmt.Println(fi20, v2)
                                        _ = fi20
                                    }
                                }
                                for fi22 := 0; fi22 < 16; fi22++ {
                                    v5 = (!(-298 == 0))
                                    var v23 float32 = float32((((1e10 * (0.1 + 0.1)) / -2.25) + float64(3.14159)))
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            if v5 {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                } else {
                                    v24 := GenS1{F0: -9223372036854775808, F1: 127}
                                    _ = v24
                                }
                                _ = len(v1)
                            }
                        }
                        var v25 float32 = -2.25
                        _ = v25
                    } else {
                        v13 = v2
                    }
                case "a":
                    v26 := GenS2{F0: 0}
                    _ = v26
                }
                if v5 {
                    v5 = true
                } else {
                    v27 := GenS1{F0: -9223372036854775808, F1: -2147483648}
                    _ = v27
                }
            }
        }
        _ = v16
    }
    v14 = 42
    var v28 string = "a"
    v2 -= v2
    v28 = ("\n\t" + "hello")
    for fi29 := 0; fi29 < 9; fi29++ {
        if ((v14 ^ v14) <= v14) {
            if (1 > -2147483648) {
                v3 = append(v3, v28)
                v30 := GenS1{F0: 32768, F1: fi29}
                _ = v30
            } else {
                if v5 {
                    v4[(fi29 & (-(fi29)))] = fi29
                    v31 := v1[128]
                    for fi32 := 0; fi32 < 13; fi32++ {
                        fmt.Println(v2, v13)
                        _ = fi32
                    }
                    _ = v31
                } else {
                    v33 := GenS1{F0: -760, F1: int(v13)}
                    switch (fi29 | -889) {
                    case 127:
                        if (0.1 < ((((3.14159 * (0.0 / 1.5)) + float64(((1e10 - 3.14159) / 1e10))) - (0.0 * 1e10)) / (float64(1e10) * -2.25))) {
                            switch (fi29 >> ((2) & 63)) {
                            case -292:
                                if v5 {
                                    if v5 {
                                        for fi34 := 0; fi34 < 17; fi34++ {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi34
                                        }
                                        var v35 int64 = -1
                                        _ = v35
                                    } else {
                                        fmt.Println(v14, fi29, v2)
                                    }
                                    v36 := GenS3{F0: 1}
                                    v37 := v4[1]
                                    _ = v36
                                    _ = v37
                                }
                            default:
                                v38 := v4[fi29]
                                _ = v38
                            }
                            if (true && (v5 && (0 != 65535))) {
                                var v39 uint64 = 1
                                var v40 uint16 = 255
                                v41 := GenS2{F0: 875}
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                var v42 float64 = 0.1
                                var v43 uint32 = 0
                                _ = v42
                                _ = v43
                            }
                        }
                    default:
                        for fi44 := 0; fi44 < 5; fi44++ {
                            var v45 uint8 = 74
                            v46 := GenS4{F0: (^(v14)), F1: genFn0(1, 127, (v13 << ((65535) & 63)))}
                            _ = len(v4)
                            switch fi29 {
                            case -440:
                                if v5 {
                                    v47 := GenS3{F0: 222}
                                    switch (-(fi44)) {
                                    case -245:
                                        var v48 float32 = 0.0
                                        _ = v48
                                    case -48:
                                        var v49 bool = (((-2.25 * 3.14159) / 1.0) == 0.1)
                                        _ = v49
                                    case -9223372036854775808:
                                        var v50 int16 = -32768
                                        _ = v50
                                    default:
                                        v4[fi44] = (fi29 >> ((42) & 63))
                                    }
                                    _ = v47
                                } else {
                                    v51 := GenS3{F0: 901}
                                    _ = v51
                                }
                                var v52 uint64 = 42
                                v53 := GenS1{F0: int(v2), F1: fi29}
                                _ = v52
                                _ = v53
                            case 605:
                                v54 := GenS5{F0: "foo bar"}
                                _ = v54
                            }
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = v28
                    }
                    _ = v33
                }
                v14 = v14
                v55 := v1[fi29]
                _ = v55
            }
            var v56 uint = 0
            _ = v56
        }
        var v57 int8 = v13
        _ = fi29
        _ = v57
    }
    delete(v1, -529)
    v58 := GenS3{F0: 32767}
    v59 := GenS6{F0: -877, F1: 9223372036854775807, F2: v28}
    var v60 uint8 = 42
    if (810 != 4294967295) {
        if true {
            var v61 int16 = -657
            _ = len(v1)
            for fi62 := 0; fi62 < 17; fi62++ {
                _ = len(v4)
                var v63 float32 = 0.1
                _ = fi62
                _ = v63
            }
            _ = v61
        } else {
            if v5 {
                var v64 int8 = -128
                _ = v64
            }
        }
        switch "hello" {
        case "\n\t":
            for fi65 := 0; fi65 < 17; fi65++ {
                var v66 uint = 2
                for fi67 := 0; fi67 < 7; fi67++ {
                    var v68 int8 = (v2 % ((int8(v66)) | 1))
                    fmt.Println(v5, v13)
                    _ = fi67
                    _ = v68
                }
                for fi69 := 0; fi69 < 15; fi69++ {
                    fmt.Println(v13, v28)
                    _ = fi69
                }
                _ = fi65
                _ = v66
            }
        case "foo bar":
            v70 := GenS7{F0: 0, F1: -1}
            _ = v70
        case "hello":
            var v71 uint32 = uint32(v60)
            _ = v71
        }
        var v72 float32 = 1.5
        _ = v72
    } else {
        v73 := GenS4{F0: v14, F1: genFn0(32767, -1, v2)}
        for fi74 := 0; fi74 < 2; fi74++ {
            var v75 int16 = -458
            fmt.Println(v5, v75)
            _ = fi74
            _ = v75
        }
        _ = v73
    }
    delete(v1, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
}
