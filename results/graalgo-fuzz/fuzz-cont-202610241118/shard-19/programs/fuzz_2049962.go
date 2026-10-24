package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 string; F2 uint16 }
type GenS12 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 uint, p1 uint8) int16 {
    v6 := GenS1{F0: "foo bar", F1: 1}
    p0 -= 989
    var v7 int32 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 591
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 42
    var v3 string = ("foo bar" + "foo bar")
    var v4 uint16 = 32767
    var v5 uint16 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = 42
        _ = fi8
        _ = v9
    }
    var v10 int16 = 0
    var v11 int16 = (v10 | v10)
    v12 := GenS2{F0: 1, F1: -9223372036854775808, F2: genFn0(39, 127)}
    switch ("a" + ((("hello" + (v3 + v3)) + ((v3 + (v3 + ("" + v3))) + v3)) + v3)) {
    case "a":
        if false {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS3{F0: 28, F1: 1, F2: 0}
                _ = fi13
                _ = v14
            }
            if false {
                if false {
                    v15 := []int{-545, int(v10)}
                    v16 := GenS4{F0: 42}
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{9223372036854775807}
                    v2 += (v2 & uint32(v4))
                    _ = v17
                }
            } else {
                v18 := []int{4294967295, 4294967295, 1, 1}
                _ = v18
            }
        }
    }
    v19 := []int{-245}
    if (65535 <= 18446744073709551615) {
        for fi20 := 0; fi20 < 1; fi20++ {
            switch "a" {
            case "":
                if true {
                    if ((true && false) || true) {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            fmt.Println(v5, v4)
                            var v22 float64 = v1
                            v2 -= 42
                            _ = fi21
                            _ = v22
                        }
                    }
                    if true {
                        var v23 float32 = 1.5
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 float64 = 0.0
                        _ = fi24
                        _ = v25
                    }
                } else {
                    switch (((v3 + v3) + v3) + "a") {
                    case "foo bar":
                        fmt.Println(v1, v11)
                    case "a":
                        v19 = append(v19, fi20)
                    }
                    v26 := GenS1{F0: v3, F1: (v2 & (v2 << ((18446744073709551615) & 63)))}
                    if false {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            v28 := GenS5{F0: -82, F1: fi20}
                            var v29 uint32 = v2
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v10 += (v11 >> ((0) & 63))
                    }
                    _ = v26
                }
            case "foo bar":
                if (858 > 4294967295) {
                    fmt.Println(fi20, v5)
                    fmt.Println(v4, v11)
                    v2 = (v2 >> ((uint(v2)) & 63))
                } else {
                    var v30 uint8 = 42
                    v11 -= (-(v11))
                    _ = v30
                }
            case "hello":
                for fi31 := 0; fi31 < 4; fi31++ {
                    var v32 uint16 = (v5 << ((1) & 63))
                    switch ((((v3 + v3) + v3) + "") + "a") {
                    case "a":
                        v33 := GenS6{F0: "foo bar", F1: "\n\t"}
                        _ = v33
                    case "":
                        v34 := GenS7{F0: 832, F1: (true && true)}
                        _ = v34
                    case "\n\t":
                        fmt.Println(v2, v2, v5)
                    default:
                        v3 = v3
                    }
                    if (0.0 > (1e10 / (v1 * 0.1))) {
                        v35 := GenS8{F0: 9223372036854775807, F1: fi20}
                        v36 := GenS1{F0: v3, F1: 189}
                        v11 -= v11
                        var v37 string = v3
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    switch -1 {
                    case 9223372036854775807:
                        var v38 float64 = 1e10
                        _ = v38
                    default:
                        fmt.Println(v11, v11, v11)
                    }
                    _ = fi31
                    _ = v32
                }
            default:
                if true {
                    var v39 float64 = v1
                    fmt.Println(v11, v2, v2)
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 9; fi40++ {
                        fmt.Println(v11, v4, v10)
                        v1 = v1
                        switch fi20 {
                        case 0:
                            switch fi20 {
                            case -612:
                                var v41 int64 = 1
                                _ = v41
                            case 4294967295:
                                var v42 uint64 = 2
                                _ = v42
                            case 128:
                                var v43 bool = true
                                _ = v43
                            default:
                                if (-95 < int8(v2)) {
                                    if (!true) {
                                        var v44 int64 = -828
                                        _ = v44
                                    } else {
                                        var v45 uint16 = v4
                                        v46 := GenS9{F0: 42}
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                fmt.Println(v11, v1)
                                fmt.Println(fi20)
                            }
                        default:
                            switch fi20 {
                            case -9223372036854775808:
                                if ((((v3 + (v3 + (v3 + (("hello" + v3) + "foo bar")))) + v3) >= (((v3 + "a") + v3) + v3)) || false) {
                                    switch fi40 {
                                    case 0:
                                        if false {
                                            v47 := GenS8{F0: 1, F1: 2147483647}
                                            var v48 int64 = -945
                                            v4 += 417
                                            _ = v47
                                            _ = v48
                                        } else {
                                            switch (v3 + v3) {
                                            case "\n\t":
                                                switch ((v3 + (v3 + v3)) + v3) {
                                                case "a":
                                                    if len(v19) > 0 {
                                                        v19[0] = 32767
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        v19 = append(v19, (fi40 & fi20))
                                                        if ("a" != v3) {
                                                            v49 := GenS10{F0: genFn0(18446744073709551615, 51)}
                                                            var v50 uint64 = 931
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v51 := GenS8{F0: int64(v2), F1: (-(fi40))}
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 uint64 = 2
                                                        _ = v52
                                                    }
                                                case "hello":
                                                    fmt.Println(v11, v1)
                                                default:
                                                    v11 = 235
                                                }
                                                v10 = 574
                                                var v53 uint = 0
                                                _ = v53
                                            case "foo bar":
                                                var v54 uint8 = 255
                                                _ = v54
                                            }
                                            v3 = v3
                                        }
                                    }
                                    v1 -= v1
                                    var v55 float32 = float32(v1)
                                    _ = v55
                                } else {
                                    var v56 bool = (!true)
                                    v10 = (v11 % (((-(v10))) | 1))
                                    _ = v56
                                }
                            default:
                                if (v3 < v3) {
                                    v10 = v11
                                }
                            }
                        }
                        _ = fi40
                    }
                }
            }
            _ = fi20
        }
        var v57 float32 = (3.14159 + -2.25)
        _ = v57
    } else {
        for fi58 := 0; fi58 < 14; fi58++ {
            var v59 uint32 = (v2 / (((v2 * (v2 ^ v2))) | 1))
            _ = fi58
            _ = v59
        }
        fmt.Println(v11, v4)
        v60 := GenS11{F0: 100, F1: "\n\t", F2: (v5 | (v4 + 42))}
        _ = v60
    }
    var v61 int32 = 2147483647
    v61 = v61
    var v62 string = "hello"
    fmt.Println(v1, v3)
    var v63 int16 = (v11 >> ((1) & 63))
    for fi64 := 0; fi64 < 9; fi64++ {
        for fi65 := 0; fi65 < 9; fi65++ {
            fmt.Println(fi65, v61, v63)
            _ = fi65
        }
        v66 := GenS12{F0: (^(v61)), F1: 1, F2: v4}
        switch fi64 {
        case -2147483648:
            var v67 int64 = 1
            _ = v67
        case 4294967295:
            if true {
                var v68 int8 = 16
                if (136 <= (v61 >> ((18446744073709551615) & 63))) {
                    var v69 float64 = (v1 * v1)
                    for fi70 := 0; fi70 < 6; fi70++ {
                        switch fi64 {
                        case 712:
                            v62 = ((("" + "\n\t") + v3) + v3)
                        case 487:
                            v11 = genFn0(162, 2)
                        case 4294967295:
                            var v71 int = 9223372036854775807
                            _ = v71
                        default:
                            v69 += (v1 / 1.0)
                        }
                        v3 = ("\n\t" + (("hello" + v62) + v3))
                        v19 = append(v19, fi70)
                        _ = fi70
                    }
                    var v72 string = v3
                    _ = v69
                    _ = v72
                } else {
                    for fi73 := 0; fi73 < 10; fi73++ {
                        v74 := GenS12{F0: 42, F1: uint64(v2), F2: (v4 + v4)}
                        var v75 uint32 = 0
                        var v76 int64 = 0
                        _ = fi73
                        _ = v74
                        _ = v75
                        _ = v76
                    }
                }
                switch ("a" + v3) {
                case "a":
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case "foo bar":
                    v77 := GenS2{F0: v61, F1: 0, F2: v63}
                    _ = v77
                default:
                    v61 = (v61 + 0)
                }
                _ = v68
            } else {
                var v78 float32 = (((float32(float64(-2.25)) / 1.0) - 1e10) * 0.0)
                var v79 uint32 = (v2 << ((34) & 63))
                _ = v78
                _ = v79
            }
        }
        _ = fi64
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
}
