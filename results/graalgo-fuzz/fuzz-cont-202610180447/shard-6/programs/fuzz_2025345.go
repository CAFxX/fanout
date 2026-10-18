package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS4 struct { F0 int8; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 int32 }
type GenS6 struct { F0 uint; F1 int32; F2 string }
type GenS7 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 uint64, p1 int64, p2 int16) bool {
    var v6 bool = (42 > 2)
    switch "hello" {
    case "a":
        p1 += (p1 + (p1 - int64(p0)))
    default:
        p2 = (p2 | (^(p2)))
    }
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (!v6)
}
func genFn1(p0 int, p1 int32) uint8 {
    if (3.14159 < (1.5 * 0.0)) {
        v7 := GenS1{F0: 686, F1: "\n\t", F2: -985}
        _ = v7
    } else {
        p1 = 1
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            p0 -= fi9
            _ = fi9
        }
        _ = fi8
    }
    v10 := GenS2{F0: 0, F1: 42}
    for fi11 := 0; fi11 < 16; fi11++ {
        if (-128 <= 100) {
            var v12 int32 = 0
            _ = v12
        } else {
            if true {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v14 := []int{(fi13 << ((0) & 63)), 127, (^(fi11)), -1}
                    var v15 uint16 = 0
                    _ = len(v14)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint8 = 42
                _ = v16
            }
        }
        _ = fi11
    }
    v17 := []int{(p0 ^ -554), p0}
    _ = p0
    _ = p1
    _ = v10
    _ = v17
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    v3 := []int{42}
    var v4 uint64 = 0
    var v5 string = "foo bar"
    var v18 float32 = 0.1
    if (true || true) {
        var v19 int8 = 1
        _ = v19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 bool = (-637 != 9223372036854775807)
    fmt.Println(v4, v20)
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if ((v4 % ((uint64(v2)) | 1)) <= v4) {
                switch (fi22 >> ((633) & 63)) {
                case -662:
                    var v23 uint8 = 246
                    _ = v23
                case 9223372036854775807:
                    switch "hello" {
                    case "foo bar":
                        if false {
                            switch (fi22 + fi22) {
                            case -255:
                                switch fi21 {
                                case 32767:
                                    v24 := GenS2{F0: 1, F1: 1}
                                    _ = v24
                                default:
                                    v18 = (float32(1.5) - v18)
                                }
                            case 4294967295:
                                if (!(int16(fi22) <= 42)) {
                                    fmt.Println(v5, fi21)
                                }
                            case 2147483647:
                                v25 := GenS3{F0: (^(v4)), F1: (v4 * (v4 ^ (^(v4)))), F2: -803}
                                _ = v25
                            }
                            if true {
                                fmt.Println(v5, v2)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                switch (fi21 & (^(fi21))) {
                                case 9223372036854775807:
                                    var v26 int = -9223372036854775808
                                    _ = v26
                                case 0:
                                    v3 = append(v3, fi21)
                                default:
                                    fmt.Println(fi21, v2, v5)
                                }
                            }
                            v27 := GenS1{F0: 0, F1: ((("\n\t" + ("foo bar" + "")) + v5) + v5), F2: -894}
                            fmt.Println(v5)
                            _ = v27
                        } else {
                            if (836 > 2147483648) {
                                v28 := GenS4{F0: 100, F1: 255, F2: -397}
                                if v20 {
                                    var v29 uint8 = 203
                                    _ = v29
                                } else {
                                    v30 := GenS4{F0: 0, F1: v2, F2: 32767}
                                    _ = v30
                                }
                                var v31 int8 = -1
                                _ = v28
                                _ = v31
                            }
                        }
                    }
                case 0:
                    fmt.Println(fi22, v20)
                }
                _ = len(v1)
                v20 = ((^(v4)) > 128)
            } else {
                v32 := GenS5{F0: (-1 <= 2147483647), F1: fi21, F2: -441}
                if (!v20) {
                    if (v2 != v2) {
                        switch (("\n\t" + "foo bar") + (v5 + v5)) {
                        case "\n\t":
                            var v33 uint = 456
                            _ = v33
                        case "a":
                            var v34 uint = 1
                            _ = v34
                        case "hello":
                            if true {
                                v2 -= v2
                                switch v5 {
                                case "foo bar":
                                    if v20 {
                                        if v20 {
                                            var v35 int64 = 0
                                            _ = v35
                                        }
                                        if ((3.14159 + (-2.25 - float64(0.1))) == 1e10) {
                                            if v20 {
                                                var v36 float64 = ((0.0 / -2.25) / 1.0)
                                                v37 := GenS5{F0: (!(!v20)), F1: fi22, F2: 367}
                                                var v38 uint32 = 1
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            if (false || v20) {
                                                switch fi22 {
                                                case 42:
                                                    if (!false) {
                                                        fmt.Println(fi21, fi21)
                                                        var v39 uint64 = v4
                                                        _ = v39
                                                    }
                                                case 9223372036854775807:
                                                    fmt.Println(v20)
                                                default:
                                                    if v20 {
                                                        switch 1 {
                                                        case 32767:
                                                            if (0 >= 42) {
                                                                var v40 uint16 = v2
                                                                var v41 float32 = v18
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        case 9223372036854775807:
                                                            v42 := GenS3{F0: v4, F1: 42, F2: 622}
                                                            _ = v42
                                                        case -2147483648:
                                                            var v43 uint64 = 838
                                                            _ = v43
                                                        }
                                                    }
                                                }
                                                v3 = append(v3, fi22)
                                                v4 = 4294967295
                                                v3 = append(v3, (fi21 | fi22))
                                                if v20 {
                                                    v44 := GenS4{F0: 42, F1: 727, F2: int16(v4)}
                                                    var v45 float32 = v18
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    v46 := v1[(v5 + "\n\t")]
                                                    v18 += v18
                                                    _ = v46
                                                }
                                            } else {
                                                v20 = (-716 > 9223372036854775807)
                                                var v47 uint16 = (v2 | (^(v2)))
                                                v48 := GenS4{F0: -1, F1: (v47 * (v2 * (v2 % ((1) | 1)))), F2: 0}
                                                var v49 uint = 1
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        } else {
                                            v50 := GenS6{F0: 2, F1: -1, F2: ("foo bar" + v5)}
                                            _ = v50
                                        }
                                    } else {
                                        _ = len(v3)
                                        v51 := GenS4{F0: int8(v4), F1: v2, F2: 1}
                                        _ = v51
                                    }
                                case "a":
                                    var v52 int64 = 9223372036854775807
                                    _ = v52
                                default:
                                    if ((^(v4)) <= v4) {
                                        switch (("foo bar" + ("\n\t" + (v5 + (v5 + v5)))) + (v5 + (v5 + v5))) {
                                        case "\n\t":
                                            v53 := GenS6{F0: 2, F1: 322, F2: v5}
                                            _ = v53
                                        }
                                        var v54 uint16 = (^(v2))
                                        _ = v54
                                    }
                                }
                                var v55 float32 = 0.0
                                _ = v55
                            }
                        default:
                            switch (v5 + v5) {
                            case "foo bar":
                                var v56 float64 = (1e10 - 3.14159)
                                _ = v56
                            case "hello":
                                var v57 bool = false
                                _ = v57
                            case "a":
                                v2 -= 0
                            default:
                                var v58 uint8 = 0
                                _ = v58
                            }
                        }
                        v20 = ((!v20) || (2147483647 == 0))
                        var v59 float64 = 1.5
                        _ = v59
                    } else {
                        v18 -= (3.14159 * (((float32(0.1) + v18) + (v18 - (((v18 * -2.25) + float32(-2.25)) / (v18 / v18)))) * float32(((((-2.25 / 0.1) / (-2.25 / 3.14159)) - (((float64(0.1) - (-2.25 - 3.14159)) / (((1e10 * (((3.14159 * 0.0) * (-2.25 - 0.0)) - ((float64((0.1 - v18)) * 3.14159) - 1e10))) * 0.1) / 0.1)) * (-2.25 / ((3.14159 / (0.0 * float64((float32(float64(v18)) - (v18 + v18))))) - -2.25)))) * 3.14159))))
                    }
                    var v60 int32 = -1
                    switch (fi21 + -918) {
                    case -2147483648:
                        fmt.Println(v2)
                    }
                    _ = v60
                }
                if v20 {
                    v61 := GenS7{F0: v5, F1: (v4 << ((uint(v2)) & 63)), F2: v2}
                    var v62 bool = true
                    if (!v20) {
                        _ = len(v3)
                    }
                    fmt.Println(fi22, v20)
                    var v63 string = v5
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v61
                    _ = v62
                    _ = v63
                } else {
                    fmt.Println(v5)
                    v18 = v18
                }
                _ = v32
            }
            _ = fi22
        }
        _ = fi21
    }
    v2 -= (v2 % ((uint16(v4)) | 1))
    v64 := GenS5{F0: v20, F1: -45, F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v64
}
