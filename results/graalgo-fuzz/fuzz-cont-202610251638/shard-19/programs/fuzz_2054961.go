package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 string; F2 int64 }
type GenS7 struct { F0 uint }
func genFn0(p0 int8) uint8 {
    switch ("a" + ("foo bar" + "hello")) {
    case "hello":
        p0 = (p0 * p0)
    case "\n\t":
        var v5 int64 = -1
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float32 = ((((-2.25 * 1e10) + 1e10) * 3.14159) * 3.14159)
        _ = fi6
        _ = v7
    }
    var v8 float32 = (0.0 - (0.0 + 3.14159))
    var v9 int = 725
    _ = p0
    _ = v8
    _ = v9
    return 31
}
func genFn1(p0 uint, p1 int16, p2 int16) bool {
    var v10 float64 = -2.25
    if true {
        switch ("a" + "foo bar") {
        case "a":
            for fi11 := 0; fi11 < 17; fi11++ {
                for fi12 := 0; fi12 < 11; fi12++ {
                    if ((false && (42 != -1)) || ((363 > 934) && true)) {
                        v13 := []int{1, (^(fi11))}
                        _ = v13
                    } else {
                        v14 := []int{fi11, (fi11 + (^(fi11))), (^(fi12))}
                        _ = v14
                    }
                    _ = fi12
                }
                v15 := GenS1{F0: 876, F1: (fi11 + (fi11 >> ((p0) & 63)))}
                _ = fi11
                _ = v15
            }
        case "":
            if (((p0 % (((p0 % (((p0 >> (((p0 * p0)) & 63))) | 1))) | 1)) >= 919) || true) {
                v10 += (-2.25 + v10)
                if (42 <= 1) {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := []int{fi16, (fi16 - fi16), (fi16 % ((689) | 1)), (fi16 + fi16)}
                        _ = fi16
                        _ = v17
                    }
                    switch 127 {
                    case -454:
                        v18 := []int{2147483647}
                        _ = v18
                    }
                } else {
                    p1 = p2
                    p1 -= -590
                    var v19 uint8 = 128
                    _ = v19
                }
                p1 -= 0
            } else {
                v10 -= v10
                v10 += (v10 + v10)
                if (65535 <= 2147483648) {
                    v20 := GenS2{F0: "\n\t", F1: "a", F2: 758}
                    v21 := GenS3{F0: false, F1: 0}
                    var v22 int = 4294967295
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v23 := []int{}
                    for fi24 := 0; fi24 < 11; fi24++ {
                        p0 -= (p0 - 255)
                        _ = fi24
                    }
                    _ = v23
                }
            }
        case "hello":
            v10 -= v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return false
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]string)
    var v3 uint64 = (v1 - 1)
    var v4 uint8 = 0
    var v25 bool = true
    if true {
        v4 = (v4 << ((769) & 63))
    } else {
        v26 := []int{-1, int(v3), -920, -1}
        _ = v26
    }
    v3 = uint64(v4)
    v27 := GenS4{F0: (v4 % ((255) | 1)), F1: 2, F2: 2147483648}
    for fi28 := 0; fi28 < 6; fi28++ {
        if ((v3 & (v1 * v3)) < (v1 ^ (v3 | (v3 & v1)))) {
            var v29 int16 = 473
            if false {
                var v30 uint = 2
                _ = v30
            }
            if (42 >= -277) {
                fmt.Println(v29, fi28)
                switch "" {
                case "hello":
                    if (2147483647 > 2147483647) {
                        v29 -= v29
                        switch "a" {
                        case "\n\t":
                            v31 := GenS1{F0: 0, F1: -2147483648}
                            _ = v31
                        case "foo bar":
                            v32 := GenS2{F0: (("foo bar" + "") + "hello"), F1: "", F2: 0}
                            _ = v32
                        default:
                            if (!((3.14159 + 3.14159) > (-2.25 - 1e10))) {
                                delete(v2, 128)
                                for fi33 := 0; fi33 < 9; fi33++ {
                                    v1 = 0
                                    var v34 uint = 0
                                    _ = fi33
                                    _ = v34
                                }
                            }
                        }
                    }
                default:
                    for fi35 := 0; fi35 < 17; fi35++ {
                        v1 = (v3 - (v1 ^ v1))
                        _ = fi35
                    }
                }
                v25 = (42 > 0)
            }
            _ = v29
        }
        _ = fi28
    }
    switch (("\n\t" + "foo bar") + "") {
    case "foo bar":
        if (v25 && (255 == 129)) {
            v3 -= v3
            v36 := []int{-2147483648}
            _ = v36
        }
    }
    v4 = v4
    if false {
        v37 := []int{127}
        var v38 int64 = -189
        _ = v37
        _ = v38
    } else {
        switch 1 {
        case 9223372036854775807:
            v39 := GenS1{F0: 65535, F1: -933}
            _ = v39
        case 4294967295:
            if v25 {
                v3 = (v1 % ((v3) | 1))
                var v40 uint32 = 2147483648
                _ = v40
            } else {
                v41 := GenS4{F0: 2, F1: 0, F2: 2147483647}
                switch 4294967295 {
                case -9223372036854775808:
                    v4 -= (v4 | (v4 >> ((2) & 63)))
                case 425:
                    v2[127] = (("\n\t" + ("foo bar" + "foo bar")) + ("hello" + (("hello" + "") + "a")))
                default:
                    v42 := GenS4{F0: (v4 + 128), F1: 0, F2: 65535}
                    _ = v42
                }
                var v43 uint16 = 1
                _ = v41
                _ = v43
            }
        case 32767:
            fmt.Println(v3, v3)
        default:
            v25 = (0 <= 42)
        }
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 uint32 = 65535
        _ = fi44
        _ = v45
    }
    if ((v3 | v3) > v3) {
        switch 127 {
        case 336:
            fmt.Println(v1)
        case 9223372036854775807:
            var v46 uint32 = 904
            _ = v46
        default:
            var v47 uint32 = 214
            _ = v47
        }
        if (143 < (v4 << ((42) & 63))) {
            _ = len(v2)
            v48 := GenS4{F0: genFn0(-116), F1: (^(v1)), F2: 247}
            _ = v48
        }
        switch 42 {
        case 4294967295:
            if ((v1 * 896) < v1) {
                v49 := GenS5{F0: -836, F1: true, F2: 255}
                _ = v49
            }
        case 42:
            for fi50 := 0; fi50 < 14; fi50++ {
                for fi51 := 0; fi51 < 9; fi51++ {
                    v25 = true
                    _ = fi51
                }
                if ((127 != 20) || (215 == 2)) {
                    v52 := GenS6{F0: 0, F1: "a", F2: -9223372036854775808}
                    _ = v52
                } else {
                    v1 = (v3 / (((v3 ^ (v3 + v1))) | 1))
                    switch "a" {
                    case "a":
                        var v53 int16 = -1
                        _ = v53
                    case "":
                        for fi54 := 0; fi54 < 12; fi54++ {
                            if v25 {
                                if (-1 < 1) {
                                    v55 := []int{(-(fi54))}
                                    switch 299 {
                                    case 2147483647:
                                        v56 := GenS7{F0: 0}
                                        _ = v56
                                    case -637:
                                        switch -886 {
                                        case 0:
                                            v3 += (v1 - 1)
                                        }
                                    }
                                    _ = v55
                                }
                                var v57 uint64 = v3
                                _ = v57
                            } else {
                                v58 := []int{4294967295, (fi50 >> ((0) & 63)), (fi50 | (^(fi54))), fi54}
                                var v59 int32 = 42
                                _ = v58
                                _ = v59
                            }
                            if (int(v4) <= -777) {
                                if (((!((-32768 < 0) && v25)) && (!false)) || true) {
                                    switch ("\n\t" + ("" + ("\n\t" + (("foo bar" + "") + ("hello" + ("a" + "hello")))))) {
                                    case "foo bar":
                                        v3 = (v1 * (v3 / ((835) | 1)))
                                    case "hello":
                                        switch 380 {
                                        case 127:
                                            v60 := []int{-1, (fi54 * (fi54 + fi50)), 9223372036854775807, (fi50 | fi54)}
                                            _ = v60
                                        case 2147483647:
                                            v4 -= v4
                                        case 128:
                                            var v61 string = ""
                                            _ = v61
                                        default:
                                            v25 = false
                                        }
                                    case "\n\t":
                                        var v62 int16 = 1
                                        _ = v62
                                    default:
                                        v63 := GenS3{F0: (!true), F1: 0}
                                        _ = v63
                                    }
                                } else {
                                    v1 = 496
                                    var v64 int32 = -1
                                    _ = v64
                                }
                                if (v25 || (1 > -161)) {
                                    v25 = false
                                    var v65 string = ""
                                    _ = v65
                                }
                                var v66 int = (fi54 | (fi50 ^ 4294967295))
                                _ = v66
                            } else {
                                v4 += 1
                                fmt.Println(v4)
                                var v67 string = ("hello" + (("" + (("\n\t" + "hello") + "a")) + "foo bar"))
                                _ = v67
                            }
                            _ = fi54
                        }
                    case "foo bar":
                        var v68 int32 = -654
                        _ = v68
                    default:
                        v69 := []int{(fi50 << ((1) & 63)), -926}
                        _ = v69
                    }
                    var v70 uint32 = uint32(v4)
                    _ = v70
                }
                if (!v25) {
                    v1 = 274
                }
                _ = fi50
            }
        default:
            fmt.Println(v1, v3, v25)
        }
    } else {
        v71 := []int{int(v3), 321, 4294967295, 127}
        _ = v71
    }
    v25 = true
    fmt.Println(v25, v4)
    fmt.Println(v25, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
