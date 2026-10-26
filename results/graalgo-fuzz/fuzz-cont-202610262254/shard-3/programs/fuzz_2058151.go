package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }
type GenS5 struct { F0 int64 }
func genFn0(p0 int64) uint32 {
    var v6 int = 128
    v7 := []int{(-(v6)), 0}
    p0 = (p0 | p0)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int8 = -1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1() int16 {
    var v10 float64 = ((1.5 + 3.14159) / 0.1)
    for fi11 := 0; fi11 < 3; fi11++ {
        v10 = -2.25
        switch fi11 {
        case -1:
            if true {
                v12 := []int{(-(fi11)), (fi11 - fi11)}
                fmt.Println(v10, fi11)
                if true {
                    v10 -= (v10 / v10)
                } else {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        switch (fi13 | 128) {
                        case 32768:
                            var v14 int = 1
                            _ = v14
                        default:
                            v10 = (1e10 * v10)
                        }
                        var v15 int64 = 0
                        _ = fi13
                        _ = v15
                    }
                    _ = len(v12)
                    v10 = (1e10 - 0.0)
                }
                _ = v12
            }
        case 787:
            switch (fi11 / ((2147483647) | 1)) {
            case 2147483647:
                switch (fi11 ^ fi11) {
                case 773:
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v10 += float64(1e10)
                        v10 -= 0.1
                        v10 += 0.1
                        _ = fi16
                    }
                default:
                    var v17 int16 = 89
                    _ = v17
                }
            case 32767:
                v18 := GenS1{F0: 1}
                _ = v18
            }
        default:
            v10 -= v10
        }
        v19 := GenS1{F0: 86}
        _ = fi11
        _ = v19
    }
    if (42 != 2147483648) {
        for fi20 := 0; fi20 < 4; fi20++ {
            v10 += -2.25
            var v21 float32 = ((float32(float64(3.14159)) / ((1e10 / ((float32(v10) + -2.25) + 1e10)) - -2.25)) * (float32(1e10) / ((1.5 - 0.1) * ((3.14159 - 3.14159) + (0.1 * ((((0.1 / 1.0) * -2.25) * 1.5) / float32(v10)))))))
            _ = fi20
            _ = v21
        }
        if false {
            v22 := []int{-777, 4294967295, 2147483647, 32768}
            v23 := GenS2{F0: 1, F1: 7, F2: 42}
            var v24 float64 = v10
            _ = v22
            _ = v23
            _ = v24
        }
        if false {
            v25 := GenS2{F0: 18446744073709551615, F1: 255, F2: 2}
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 string = "foo bar"
                switch fi26 {
                case 744:
                    v10 += (v10 - v10)
                case 42:
                    v27 = ("foo bar" + (("foo bar" + ("hello" + v27)) + v27))
                default:
                    var v28 int8 = -120
                    _ = v28
                }
                v27 = "\n\t"
                _ = fi26
                _ = v27
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                if true {
                    if (v10 >= (1e10 / v10)) {
                        v10 += (float64((1.5 + (3.14159 + (1e10 * 1.5)))) - (((((v10 + v10) * (((0.0 / 1e10) - float64(float32(v10))) + v10)) * v10) * v10) - (v10 + 0.1)))
                        if ((v10 / 1.0) != float64(-2.25)) {
                            v30 := GenS3{F0: (fi29 ^ 81), F1: 42}
                            _ = v30
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    v31 := []int{32767}
                                    _ = v31
                                }
                            case "foo bar":
                                v32 := []int{291, (fi29 / ((fi29) | 1)), (fi29 + 32768)}
                                _ = v32
                            case "":
                                v10 += 0.0
                            default:
                                if (true || true) {
                                    fmt.Println(v10)
                                    v33 := []int{2147483647, fi29}
                                    for fi34 := 0; fi34 < 13; fi34++ {
                                        _ = len(v33)
                                        v10 = (v10 + (1.5 - (0.1 - float64(3.14159))))
                                        var v35 float64 = 0.1
                                        _ = fi34
                                        _ = v35
                                    }
                                    _ = v33
                                }
                            }
                            var v36 float32 = 1e10
                            if (false || true) {
                                var v37 int32 = 42
                                if (((!(!(!false))) || false) || (96 < 18446744073709551615)) {
                                    var v38 uint8 = 1
                                    var v39 uint16 = 240
                                    _ = v38
                                    _ = v39
                                }
                                for fi40 := 0; fi40 < 20; fi40++ {
                                    v41 := []int{}
                                    var v42 string = "foo bar"
                                    if (127 > 202) {
                                        switch v42 {
                                        case "":
                                            var v43 float64 = v10
                                            _ = v43
                                        case "hello":
                                            v37 = v37
                                        }
                                    } else {
                                        fmt.Println(fi40, fi40, fi29)
                                    }
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                                _ = v37
                            }
                            _ = v36
                        }
                    } else {
                        fmt.Println(v10, fi29)
                        var v44 float64 = ((v10 * v10) / -2.25)
                        v10 = v10
                        _ = v44
                    }
                } else {
                    v45 := GenS2{F0: 4294967295, F1: 0, F2: 48}
                    for fi46 := 0; fi46 < 6; fi46++ {
                        fmt.Println(v10, fi46, fi29)
                        _ = fi46
                    }
                    var v47 float64 = (v10 * float64(float32(float64(0.0))))
                    _ = v45
                    _ = v47
                }
                _ = fi29
            }
            _ = v25
        }
    }
    v48 := []int{-2147483648, -1, -524, 131}
    v48 = append(v48, -161)
    _ = v10
    _ = v48
    return 42
}

func main() {
    var v1 uint64 = 867
    var v2 int8 = 66
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    var v5 uint64 = (^(v1))
    v5 -= v1
    var v49 uint16 = 1000
    v50 := GenS4{F0: 577, F1: 2, F2: uint(v1)}
    v4[("foo bar" + "")] = (1 < -1)
    var v51 float32 = (((1e10 - (-2.25 * (0.1 - float32(((0.1 - 0.1) / 1.0))))) + float32(3.14159)) * 0.0)
    v1 = v1
    v52 := GenS5{F0: -1}
    switch 143 {
    case -2147483648:
        if true {
            v4[("foo bar" + "\n\t")] = (42 < 2147483647)
            for fi53 := 0; fi53 < 5; fi53++ {
                v54 := v3["foo bar"]
                v55 := GenS5{F0: 42}
                _ = fi53
                _ = v54
                _ = v55
            }
            var v56 string = ""
            _ = v56
        } else {
            if true {
                var v57 int32 = 2147483647
                _ = v57
            }
            var v58 uint = 693
            var v59 float32 = (1e10 * (v51 / ((0.0 * v51) - 0.0)))
            _ = v58
            _ = v59
        }
    case -9223372036854775808:
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            switch 9223372036854775807 {
            case 0:
                var v60 uint8 = 2
                _ = v60
            default:
                for fi61 := 0; fi61 < 19; fi61++ {
                    var v62 uint = 530
                    var v63 int8 = 93
                    _ = fi61
                    _ = v62
                    _ = v63
                }
            }
        case "a":
            v64 := []int{}
            _ = v64
        case "\n\t":
            var v65 float32 = (1e10 + 1e10)
            _ = v65
        }
    case 9223372036854775807:
        v51 = 0.0
    }
    v49 = v49
    v66 := GenS1{F0: (v5 | 821)}
    v67 := GenS3{F0: 860, F1: (^(v5))}
    fmt.Println(v49, v2, v51)
    var v68 int8 = (v2 & (v2 & (v2 ^ int8(v49))))
    for fi69 := 0; fi69 < 17; fi69++ {
        delete(v4, "\n\t")
        _ = fi69
    }
    var v70 float32 = v51
    switch -675 {
    case 313:
        v68 += -1
    case 42:
        v49 = (v49 / ((0) | 1))
    case -835:
        v71 := []int{32767, -851, 935, -1}
        _ = v71
    }
    var v72 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v66
    _ = v67
    _ = v68
    _ = v70
    _ = v72
}
