package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64; F1 int; F2 uint64 }
type GenS6 struct { F0 uint; F1 int; F2 string }
func genFn0() bool {
    v4 := GenS1{F0: 255, F1: 3, F2: 127}
    switch "foo bar" {
    case "hello":
        v5 := []int{199, 2147483647, -9223372036854775808, 0}
        _ = v5
    case "\n\t":
        var v6 float32 = (-2.25 / 0.1)
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                if ((208 >= 472) && (!(("foo bar" > ("a" + "hello")) && (!(!((1.5 > (((-2.25 + 1e10) / 0.1) - 0.0)) && (!false))))))) {
                    if false {
                        v9 := GenS1{F0: 766, F1: 397, F2: 227}
                        var v10 bool = true
                        switch fi7 {
                        case 2147483647:
                            v10 = (-2147483648 <= 642)
                        default:
                            switch (fi7 * (fi8 & 792)) {
                            case 0:
                                v11 := []int{fi7}
                                _ = v11
                            default:
                                var v12 int = fi8
                                _ = v12
                            }
                        }
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(fi7, fi8)
                }
                _ = fi8
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := []int{fi7, fi7, fi7, 1}
                var v15 int16 = -489
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int8 = -1
            _ = fi7
            _ = v16
        }
    default:
        if (2 != 796) {
            var v17 int8 = -128
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS2{F0: 42, F1: 2}
        v20 := []int{(fi18 | fi18), fi18, (fi18 ^ (fi18 >> ((628) & 63))), 0}
        _ = fi18
        _ = v19
        _ = v20
    }
    switch 1 {
    case -502:
        for fi21 := 0; fi21 < 16; fi21++ {
            if (499 != 2147483648) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 uint16 = 216
                    v24 := []int{32767, (fi21 * (fi21 ^ -1)), 0, (-(fi21))}
                    var v25 int16 = -1
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            } else {
                switch ("a" + (("hello" + ("\n\t" + ("" + ("" + (("hello" + "hello") + "foo bar"))))) + ((((("" + ("" + (("a" + ((("a" + "\n\t") + "\n\t") + "hello")) + ""))) + ("hello" + "\n\t")) + ("a" + "hello")) + "foo bar") + "\n\t"))) {
                case "foo bar":
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 int8 = 42
                        _ = fi26
                        _ = v27
                    }
                }
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 uint32 = 0
                _ = fi28
                _ = v29
            }
            _ = fi21
        }
    case -1:
        var v30 bool = true
        _ = v30
    case 32768:
        var v31 float64 = 1.5
        _ = v31
    default:
        var v32 int64 = -928
        _ = v32
    }
    var v33 uint16 = 42
    _ = v33
    _ = v4
    return (65535 != (^(v33)))
}

func main() {
    var v1 uint8 = 2
    var v2 int = -9223372036854775808
    v3 := make(map[string]uint32)
    v1 = v1
    for fi34 := 0; fi34 < 19; fi34++ {
        v1 += 236
        var v35 string = ("foo bar" + "foo bar")
        _ = fi34
        _ = v35
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        var v36 uint = 0
        _ = v36
    }
    var v37 bool = (115 <= 2)
    v38 := GenS3{F0: 1000}
    switch "\n\t" {
    case "":
        var v39 int16 = 1000
        _ = v39
    case "a":
        v1 = v1
    case "hello":
        if v37 {
            v40 := []int{(v2 << ((4294967295) & 63))}
            _ = v40
        } else {
            if v37 {
                fmt.Println(v37)
                if ((100 > 51) && (0 == 689)) {
                    if v37 {
                        var v41 uint32 = 2
                        if (-92 >= -46) {
                            v42 := GenS4{F0: 400, F1: 127, F2: 9223372036854775807}
                            _ = v42
                        } else {
                            v43 := []int{(v2 - v2), v2, 1}
                            fmt.Println(v37)
                            var v44 uint64 = 1
                            _ = v43
                            _ = v44
                        }
                        if false {
                            v2 = (v2 - v2)
                        } else {
                            for fi45 := 0; fi45 < 3; fi45++ {
                                var v46 int16 = -1
                                if (1 < (v1 << ((4294967295) & 63))) {
                                    v47 := GenS4{F0: 0, F1: uint8(v41), F2: -9223372036854775808}
                                    _ = v47
                                } else {
                                    var v48 int = 0
                                    _ = v48
                                }
                                fmt.Println(v46)
                                _ = fi45
                                _ = v46
                            }
                            var v49 uint = 255
                            v50 := []int{v2, (v2 + (v2 >> (((v49 % (((v49 / (((v49 * (v49 / ((437) | 1)))) | 1))) | 1))) & 63))), (v2 ^ (v2 & v2))}
                            _ = v49
                            _ = v50
                        }
                        _ = v41
                    }
                } else {
                    if v37 {
                        v51 := GenS4{F0: 42, F1: v1, F2: -188}
                        for fi52 := 0; fi52 < 3; fi52++ {
                            v2 += (fi52 & (fi52 | v2))
                            for fi53 := 0; fi53 < 13; fi53++ {
                                v54 := GenS2{F0: 2147483647, F1: 32767}
                                v2 += (fi53 ^ fi52)
                                _ = fi53
                                _ = v54
                            }
                            _ = fi52
                        }
                        _ = v51
                    } else {
                        v2 = (v2 % ((v2) | 1))
                        var v55 int64 = 9223372036854775807
                        v56 := GenS5{F0: v55, F1: v2, F2: 2}
                        _ = v55
                        _ = v56
                    }
                    v37 = v37
                    var v57 int16 = 1
                    _ = v57
                }
                for fi58 := 0; fi58 < 11; fi58++ {
                    v59 := GenS3{F0: -1}
                    _ = fi58
                    _ = v59
                }
            }
            if (false && v37) {
                switch (v2 + v2) {
                case -1:
                    fmt.Println(v37, v37)
                case -798:
                    var v60 bool = true
                    _ = v60
                }
            } else {
                if ((!(1 != 2)) && v37) {
                    v61 := []int{(v2 ^ -1), v2}
                    _ = v61
                } else {
                    v3["\n\t"] = 42
                    v37 = v37
                }
                for fi62 := 0; fi62 < 12; fi62++ {
                    v63 := v3[("a" + "a")]
                    for fi64 := 0; fi64 < 6; fi64++ {
                        var v65 float64 = (3.14159 + -2.25)
                        var v66 uint = 4294967295
                        v63 = v63
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    _ = fi62
                    _ = v63
                }
            }
            if false {
                if ((314 <= 42) && ("\n\t" != ((("a" + ("foo bar" + "hello")) + "a") + (("a" + "a") + "\n\t")))) {
                    v1 = 128
                    var v67 uint16 = 65535
                    _ = v67
                } else {
                    v3[("hello" + "a")] = 181
                    v2 = v2
                    for fi68 := 0; fi68 < 20; fi68++ {
                        fmt.Println(v37, fi68, fi68)
                        _ = fi68
                    }
                }
                v69 := []int{32767, 56, 805}
                var v70 float64 = 1e10
                _ = v69
                _ = v70
            }
        }
    }
    switch "" {
    case "\n\t":
        if (false || v37) {
            v37 = (!(!false))
        }
    case "hello":
        fmt.Println(v37, v37, v37)
    case "foo bar":
        v2 -= v2
    default:
        v2 -= (v2 / (((v2 >> ((42) & 63))) | 1))
    }
    var v71 int = v2
    v2 = (v2 << ((18446744073709551615) & 63))
    v1 = v1
    v1 = v1
    v3[("hello" + "hello")] = 2
    var v72 bool = (565 >= 276)
    for fi73 := 0; fi73 < 12; fi73++ {
        if ((-134 != (v2 << ((20) & 63))) || ((-1 < 1) && v37)) {
            v74 := []int{(fi73 & 4294967295)}
            v72 = true
            if true {
                var v75 uint16 = 42
                _ = v75
            } else {
                for fi76 := 0; fi76 < 8; fi76++ {
                    v1 = 0
                    switch fi76 {
                    case -2147483648:
                        v71 -= fi76
                    case 4294967295:
                        if (v1 < 0) {
                            v1 = (v1 + v1)
                            switch (-(v71)) {
                            case -1:
                                v77 := GenS6{F0: 2, F1: (-(v2)), F2: ("hello" + "")}
                                _ = v77
                            case 9223372036854775807:
                                if v72 {
                                    var v78 uint32 = 2147483648
                                    var v79 float32 = 0.0
                                    _ = v78
                                    _ = v79
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            var v80 uint8 = 255
                                            switch (fi73 * (^(fi76))) {
                                            case 337:
                                                var v81 uint = 1
                                                _ = v81
                                            case -900:
                                                fmt.Println(v1, v1)
                                            default:
                                                if v37 {
                                                    v82 := GenS6{F0: 1, F1: v71, F2: "\n\t"}
                                                    _ = v82
                                                }
                                                var v83 int8 = 0
                                                _ = v83
                                            }
                                            _ = v80
                                        } else {
                                            fmt.Println(v71)
                                        }
                                    case "":
                                        var v84 string = ""
                                        _ = v84
                                    }
                                    fmt.Println(v71, v72, v2)
                                    switch v71 {
                                    case -942:
                                        var v85 uint = 427
                                        _ = v85
                                    case 127:
                                        var v86 uint32 = 2147483647
                                        _ = v86
                                    }
                                    _ = len(v3)
                                }
                            default:
                                v37 = true
                            }
                        }
                    case 0:
                        var v87 float32 = 1e10
                        _ = v87
                    }
                    v3["foo bar"] = 490
                    _ = fi76
                }
                var v88 bool = (float64(-2.25) == float64(1.5))
                _ = v88
            }
            _ = v74
        }
        _ = fi73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v71
    _ = v72
}
