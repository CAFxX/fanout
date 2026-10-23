package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 string; F1 uint64 }
func genFn0(p0 int, p1 int16) int64 {
    if false {
        v7 := GenS1{F0: 0, F1: 9223372036854775807, F2: 2}
        var v8 uint = 42
        _ = v7
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (-361 > -1) {
                p1 = 654
            } else {
                p1 = -1
                p1 = 32767
                p1 += p1
            }
            _ = fi9
        }
        p1 += p1
        var v10 bool = (2147483648 > 2147483647)
        _ = v10
    } else {
        p0 = (p0 | -149)
        var v11 int64 = -9223372036854775808
        v11 = (v11 | v11)
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := []int{767, (-(p0)), fi12}
        p1 = p1
        v14 := GenS2{F0: 18446744073709551615, F1: p1}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: true}
    if (263 <= 42) {
        v16 := []int{(p0 + 2147483647), 2147483647}
        var v17 uint16 = 65535
        var v18 uint8 = 199
        _ = v16
        _ = v17
        _ = v18
    } else {
        switch (^(p0)) {
        case 999:
            var v19 int64 = 775
            _ = v19
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v15
    return -135
}
func genFn1(p0 bool, p1 int) bool {
    fmt.Println(p0)
    var v20 uint64 = 92
    _ = p0
    _ = p1
    _ = v20
    return ((p0 || p0) || p0)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 bool = (false || true)
    var v5 uint32 = 481
    var v6 float32 = 1e10
    var v21 float32 = v6
    fmt.Println(v4, v6)
    var v22 int16 = -1
    v6 = ((1e10 * 3.14159) + float32(3.14159))
    if (!((255 != 64) || (39 >= (v22 - 1000)))) {
        for fi23 := 0; fi23 < 20; fi23++ {
            if true {
                var v24 int32 = -2147483648
                v22 += (v22 % (((v22 >> ((42) & 63))) | 1))
                _ = v24
            } else {
                var v25 uint64 = 930
                _ = v25
            }
            var v26 uint32 = v5
            _ = fi23
            _ = v26
        }
        var v27 uint = 42
        v22 = (v22 | -1)
        _ = v27
    } else {
        var v28 uint8 = 14
        fmt.Println(v6, v21)
        _ = v28
    }
    v2[128] = 127
    if v4 {
        if false {
            var v29 float32 = v6
            switch ("" + "\n\t") {
            case "":
                if v4 {
                    if ((v29 + (((0.1 / (((v21 / 1e10) / 3.14159) * -2.25)) + (((v21 + (v6 * v6)) / v21) * -2.25)) / v29)) > 1.5) {
                        var v30 uint = 65535
                        delete(v3, ("hello" + ""))
                        var v31 int = 935
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(v29)
                    for fi32 := 0; fi32 < 13; fi32++ {
                        fmt.Println(v21)
                        fmt.Println(v5)
                        _ = fi32
                    }
                }
            case "a":
                var v33 string = ("foo bar" + "hello")
                _ = v33
            case "\n\t":
                v34 := v2[4294967295]
                _ = v34
            default:
                if (false || ((-108 < 27) || v4)) {
                    v6 -= v29
                } else {
                    v21 = v21
                    for fi35 := 0; fi35 < 6; fi35++ {
                        fmt.Println(v4)
                        v36 := GenS1{F0: (^(v5)), F1: -9223372036854775808, F2: 2}
                        switch fi35 {
                        case -45:
                            if v4 {
                                var v37 int32 = 654
                                _ = v37
                            } else {
                                if (!true) {
                                    fmt.Println(fi35)
                                    v6 = 0.1
                                    v38 := GenS3{F0: v4}
                                    _ = v38
                                }
                                var v39 uint32 = (v5 - (v5 << ((255) & 63)))
                                for fi40 := 0; fi40 < 16; fi40++ {
                                    switch -1 {
                                    case 0:
                                        var v41 int = 1
                                        _ = v41
                                    default:
                                        if (!true) {
                                            v29 -= float32((float64(((1.5 * 3.14159) / (v6 - 1e10))) - ((-2.25 / (0.0 - 1.5)) * 1e10)))
                                            v21 = v21
                                        } else {
                                            var v42 uint64 = 42
                                            _ = v42
                                        }
                                    }
                                    var v43 uint = 65535
                                    if v4 {
                                        v44 := GenS4{F0: -1, F1: 32767}
                                        var v45 bool = (1 == 149)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v46 := GenS5{F0: (((("\n\t" + "") + "foo bar") + "\n\t") + ("foo bar" + ("a" + ("foo bar" + "")))), F1: 549}
                                        _ = v46
                                    }
                                    _ = fi40
                                    _ = v43
                                }
                                _ = v39
                            }
                        case 42:
                            var v47 float32 = 1e10
                            _ = v47
                        case -2147483648:
                            for fi48 := 0; fi48 < 5; fi48++ {
                                var v49 int32 = 555
                                _ = fi48
                                _ = v49
                            }
                        default:
                            var v50 uint = 0
                            _ = v50
                        }
                        _ = fi35
                        _ = v36
                    }
                }
            }
            _ = v29
        } else {
            var v51 float32 = 1e10
            var v52 uint = 163
            v52 += (v52 - (v52 & 65535))
            _ = v51
            _ = v52
        }
        var v53 float32 = v21
        _ = v53
    } else {
        for fi54 := 0; fi54 < 8; fi54++ {
            v22 += (^(v22))
            var v55 uint8 = 2
            _ = fi54
            _ = v55
        }
    }
    var v56 int16 = v22
    if true {
        for fi57 := 0; fi57 < 18; fi57++ {
            if ((fi57 != (^(fi57))) && (3.14159 < 0.0)) {
                if (-2147483648 <= -1) {
                    v58 := []int{fi57, fi57}
                    v21 = v6
                    _ = v58
                }
                if ((fi57 >> ((2) & 63)) > fi57) {
                    fmt.Println(v56)
                }
            } else {
                for fi59 := 0; fi59 < 2; fi59++ {
                    v56 -= v56
                    if (v4 && (2 >= 149)) {
                        v22 -= (v56 & v56)
                        v22 = 1
                        var v60 uint = 2
                        _ = v60
                    } else {
                        if true {
                            if (genFn0(fi59, 42) != -555) {
                                v56 += (^(v22))
                                v61 := v2[fi59]
                                _ = v61
                            } else {
                                v56 = v22
                                var v62 uint16 = 32767
                                _ = v62
                            }
                            v4 = v4
                        } else {
                            var v63 int = (^(fi57))
                            v3[((("a" + ("a" + ("\n\t" + "foo bar"))) + "a") + "")] = ""
                            _ = v63
                        }
                    }
                    _ = fi59
                }
                for fi64 := 0; fi64 < 15; fi64++ {
                    v22 = -1
                    _ = fi64
                }
                v4 = ((7 != 1) || (!(v4 || v4)))
            }
            v65 := []int{(fi57 / ((32767) | 1)), (-(fi57))}
            _ = fi57
            _ = v65
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
