package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 int; F1 int; F2 bool }
type GenS11 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint8) bool {
    fmt.Println(p0, p1)
    switch 32767 {
    case 128:
        if (4294967295 < 42) {
            var v4 uint8 = uint8(p0)
            v4 -= p1
            var v5 uint8 = p1
            _ = v4
            _ = v5
        } else {
            var v6 float32 = float32(float64((((3.14159 / 3.14159) * (3.14159 / 1.5)) + (0.0 + ((((1e10 - 1.5) + float32(-2.25)) * ((-2.25 / 0.1) * 0.0)) / -2.25)))))
            _ = v6
        }
    case 1:
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{fi7, fi7, (fi7 % ((0) | 1))}
            p0 = 2
            _ = fi7
            _ = v8
        }
    case 500:
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 42}
            v11 := []int{fi9, 1, (fi9 % ((-9223372036854775808) | 1))}
            var v12 int16 = -188
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        var v13 uint16 = p0
        _ = v13
    }
    v14 := GenS1{F0: 30}
    _ = p0
    _ = p1
    _ = v14
    return false
}
func genFn1(p0 uint32, p1 uint32, p2 int8) int64 {
    if false {
        v15 := []int{2147483647, 4294967295}
        p0 -= (p1 * 77)
        if len(v15) > 0 {
            v15[0] = -639
        }
        _ = v15
    } else {
        var v16 uint8 = 2
        _ = v16
    }
    switch "hello" {
    case "a":
        var v17 int32 = -2147483648
        _ = v17
    case "":
        for fi18 := 0; fi18 < 15; fi18++ {
            p2 = p2
            for fi19 := 0; fi19 < 16; fi19++ {
                p1 = p0
                v20 := GenS1{F0: 100}
                v21 := GenS2{F0: p2, F1: 1, F2: 1}
                var v22 int = (fi19 >> ((65535) & 63))
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 uint8 = 127
            _ = fi18
            _ = v23
        }
    }
    var v24 int32 = 42
    var v25 int16 = 906
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 0
}

func main() {
    var v1 uint = 559
    var v2 int16 = 415
    var v3 int64 = 42
    switch 32768 {
    case 32767:
        v26 := GenS2{F0: 100, F1: 887, F2: -2147483648}
        _ = v26
    default:
        fmt.Println(v1, v3, v3)
    }
    switch -2147483648 {
    case 9223372036854775807:
        if ((3.14159 == (0.1 + 1.5)) && (true || true)) {
            v27 := GenS1{F0: 0}
            var v28 uint = v1
            _ = v27
            _ = v28
        } else {
            switch "foo bar" {
            case "\n\t":
                switch -9223372036854775808 {
                case 128:
                    v1 += v1
                case 32768:
                    var v29 int = 32767
                    _ = v29
                case -2147483648:
                    v30 := GenS3{F0: "\n\t"}
                    _ = v30
                default:
                    var v31 uint32 = 2
                    _ = v31
                }
            }
            fmt.Println(v3)
            v32 := []int{4294967295, 9223372036854775807, 575}
            _ = v32
        }
    case 127:
        if (!(1 < 0)) {
            var v33 int64 = (v3 + 42)
            _ = v33
        }
    case 513:
        if false {
            if true {
                switch "hello" {
                case "foo bar":
                    if (!false) {
                        var v34 string = "foo bar"
                        v35 := GenS3{F0: v34}
                        if false {
                            v34 = (v34 + v34)
                        } else {
                            for fi36 := 0; fi36 < 9; fi36++ {
                                v37 := GenS4{F0: ((v34 + "hello") + ("foo bar" + v34)), F1: -680, F2: true}
                                var v38 int = (fi36 << (((v1 - 4294967295)) & 63))
                                var v39 int64 = genFn1(42, 0, 127)
                                _ = fi36
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v40 := []int{405}
                            _ = v40
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v41 := GenS3{F0: "a"}
                        var v42 string = ""
                        _ = v41
                        _ = v42
                    }
                }
            } else {
                if true {
                    var v43 int = 9223372036854775807
                    if (false && true) {
                        v44 := GenS4{F0: "\n\t", F1: -2147483648, F2: (!(v3 < v3))}
                        _ = v44
                    }
                    _ = v43
                } else {
                    v2 += 32767
                    v2 -= -1
                    v45 := GenS5{F0: (v1 * (v1 ^ 1))}
                    _ = v45
                }
                v46 := GenS2{F0: 1, F1: 32767, F2: 0}
                _ = v46
            }
            fmt.Println(v1)
        } else {
            v47 := []int{128, -2147483648, int(v2), -717}
            _ = v47
        }
    }
    var v48 int = -447
    var v49 uint16 = 261
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        if (1e10 != ((float64(0.0) - float64(0.0)) + float64(0.1))) {
            for fi50 := 0; fi50 < 8; fi50++ {
                v2 = v2
                v51 := GenS6{F0: (v1 << (((v1 * v1)) & 63)), F1: 1}
                v52 := GenS7{F0: v1, F1: v48}
                _ = fi50
                _ = v51
                _ = v52
            }
            if (65535 <= 0) {
                v53 := GenS5{F0: (v1 - 1)}
                _ = v53
            } else {
                switch "" {
                case "foo bar":
                    v54 := []int{-9223372036854775808, -691, (v48 & 32768)}
                    _ = v54
                case "":
                    switch "" {
                    case "hello":
                        fmt.Println(v2, v1)
                    default:
                        for fi55 := 0; fi55 < 20; fi55++ {
                            v49 += 42
                            fmt.Println(v1)
                            _ = fi55
                        }
                    }
                }
                var v56 int8 = 127
                fmt.Println(v48)
                _ = v56
            }
            switch 128 {
            case 4294967295:
                v57 := GenS8{F0: v2, F1: uint8(v49), F2: 127}
                _ = v57
            case 42:
                v3 -= (^(v3))
            }
        } else {
            if true {
                if (!(!true)) {
                    switch v48 {
                    case 2147483647:
                        v58 := []int{127, (v48 / (((v48 >> (((^(v1))) & 63))) | 1)), (-(v48)), -1}
                        _ = v58
                    }
                    var v59 bool = (!false)
                    switch v48 {
                    case -1:
                        v60 := GenS4{F0: ("" + ("foo bar" + "")), F1: -290, F2: true}
                        _ = v60
                    }
                    _ = v59
                } else {
                    var v61 int8 = -128
                    v62 := GenS9{F0: (v1 / ((42) | 1)), F1: 1}
                    switch "foo bar" {
                    case "hello":
                        v3 = v3
                    case "foo bar":
                        fmt.Println(v49, v3, v61)
                    case "\n\t":
                        v63 := []int{127}
                        _ = v63
                    }
                    _ = v61
                    _ = v62
                }
            } else {
                v2 = v2
                if false {
                    var v64 int8 = 0
                    fmt.Println(v1, v49)
                    _ = v64
                } else {
                    var v65 int8 = 31
                    _ = v65
                }
                if ((v2 << (((v1 + (v1 + 604))) & 63)) <= 426) {
                    if ((v2 * (^(v2))) < v2) {
                        for fi66 := 0; fi66 < 17; fi66++ {
                            var v67 int = v48
                            _ = fi66
                            _ = v67
                        }
                        v49 = v49
                        v3 -= (v3 - v3)
                    }
                } else {
                    for fi68 := 0; fi68 < 13; fi68++ {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v48, v3, v1)
                        case "a":
                            v2 = int16(v1)
                        }
                        if (0 < int32(v3)) {
                            var v69 int32 = -2147483648
                            v70 := GenS6{F0: 2, F1: 42}
                            if false {
                                if false {
                                    switch ("foo bar" + "hello") {
                                    case "hello":
                                        v71 := []int{180, fi68, (v48 * (fi68 << ((v1) & 63))), fi68}
                                        _ = v71
                                    case "\n\t":
                                        fmt.Println(v3)
                                    }
                                } else {
                                    var v72 uint64 = 18446744073709551615
                                    for fi73 := 0; fi73 < 13; fi73++ {
                                        v74 := GenS10{F0: v48, F1: 32768, F2: (1 != v1)}
                                        _ = fi73
                                        _ = v74
                                    }
                                    switch v48 {
                                    case 0:
                                        var v75 uint32 = 65535
                                        _ = v75
                                    case 9223372036854775807:
                                        var v76 int = -2147483648
                                        _ = v76
                                    }
                                    _ = v72
                                }
                            }
                            _ = v69
                            _ = v70
                        } else {
                            switch ("hello" + ("a" + "hello")) {
                            case "foo bar":
                                switch "a" {
                                case "a":
                                    var v77 uint = v1
                                    _ = v77
                                default:
                                    for fi78 := 0; fi78 < 9; fi78++ {
                                        if true {
                                            var v79 int = (^(fi68))
                                            fmt.Println(fi78, v48, v49)
                                            _ = v79
                                        } else {
                                            v48 = int(v3)
                                        }
                                        _ = fi78
                                    }
                                }
                            case "a":
                                v48 += -282
                            default:
                                switch 475 {
                                case 128:
                                    v1 = (^(v1))
                                case 174:
                                    var v80 bool = ((false || false) || (100 == 100))
                                    _ = v80
                                }
                            }
                            v81 := []int{}
                            _ = v81
                        }
                        _ = fi68
                    }
                    fmt.Println(v48)
                    v82 := []int{v48, (-(v48)), (v48 / (((^(v48))) | 1)), (v48 & (-(v48)))}
                    _ = v82
                }
            }
            switch 4294967295 {
            case -2147483648:
                for fi83 := 0; fi83 < 2; fi83++ {
                    if (!false) {
                        var v84 int8 = -2
                        v85 := []int{fi83, fi83}
                        fmt.Println(fi83)
                        _ = v84
                        _ = v85
                    } else {
                        v86 := []int{fi83, fi83, int(v3), (fi83 >> (((v1 % ((346) | 1))) & 63))}
                        v86 = append(v86, fi83)
                        v49 = v49
                        _ = v86
                    }
                    v49 -= (^(v49))
                    _ = fi83
                }
            case 32768:
                v87 := []int{v48, v48, v48, (v48 % ((42) | 1))}
                _ = v87
            default:
                var v88 int16 = v2
                _ = v88
            }
            v89 := GenS11{F0: 18446744073709551615}
            _ = v89
        }
    }
    if true {
        for fi90 := 0; fi90 < 18; fi90++ {
            v91 := []int{2147483647}
            v2 = -804
            if true {
                for fi92 := 0; fi92 < 3; fi92++ {
                    v93 := GenS5{F0: v1}
                    _ = fi92
                    _ = v93
                }
            }
            _ = fi90
            _ = v91
        }
    }
    var v94 int = v48
    v48 -= (v48 & 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v94
}
