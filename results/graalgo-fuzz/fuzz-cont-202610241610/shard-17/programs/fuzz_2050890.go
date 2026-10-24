package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int64 }
func genFn0(p0 uint8, p1 uint, p2 int32) uint16 {
    p2 += p2
    v7 := GenS1{F0: 42, F1: (p0 != (^(p0))), F2: 935}
    p0 = 112
    v8 := []int{-9223372036854775808, -124, 9223372036854775807, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int32) bool {
    var v9 float32 = 0.0
    p0 += (p0 & p0)
    _ = p0
    _ = v9
    return (!false)
}

func main() {
    var v1 int16 = 0
    var v2 string = "\n\t"
    var v3 uint32 = 2147483647
    var v4 int = 2147483647
    var v5 uint64 = 748
    var v6 bool = false
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int = (fi10 - -2147483648)
        v1 -= (v1 * v1)
        if v6 {
            v12 := []int{v4, (^(v11))}
            _ = v12
        } else {
            var v13 uint = 65535
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v3 = (v3 % ((1) | 1))
        fmt.Println(v6)
        v6 = v6
        _ = fi14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            v6 = true
            v17 := []int{(fi15 << ((42) & 63)), 229, (-(fi15)), 32768}
            if (((v2 + ("hello" + "\n\t")) + "") == v2) {
                var v18 uint64 = v5
                if len(v17) > 0 {
                    v17[0] = (fi15 >> ((1) & 63))
                }
                switch (v2 + v2) {
                case "\n\t":
                    if (1.5 == ((-2.25 * -2.25) + 1e10)) {
                        if ((255 > 127) && false) {
                            var v19 uint32 = v3
                            fmt.Println(v19, v1)
                            if (v6 && v6) {
                                if v6 {
                                    switch (fi15 | 199) {
                                    case -595:
                                        switch "a" {
                                        case "foo bar":
                                            v20 := GenS2{F0: 42}
                                            _ = v20
                                        }
                                    }
                                    fmt.Println(fi15, v2)
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    fmt.Println(v4, v4)
                                }
                                v21 := GenS1{F0: 42, F1: false, F2: v4}
                                _ = v21
                            }
                            _ = v19
                        }
                    }
                }
                _ = v18
            } else {
                if (!false) {
                    v22 := GenS3{F0: 100, F1: v2, F2: v3}
                    v6 = v6
                    if ((int32(v4) == 0) || true) {
                        var v23 uint = 1
                        _ = v23
                    } else {
                        fmt.Println(v3)
                    }
                    _ = v22
                } else {
                    var v24 uint32 = v3
                    v17 = append(v17, v4)
                    var v25 int = fi16
                    _ = v24
                    _ = v25
                }
                v4 += fi15
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v26 := []int{-398, (v4 << ((18446744073709551615) & 63))}
    v4 = 129
    fmt.Println(v3)
    var v27 int8 = -128
    var v28 string = ("" + v2)
    switch (("a" + v28) + v2) {
    case "\n\t":
        switch 127 {
        case 0:
            v26 = append(v26, (^(v4)))
        case -899:
            v29 := GenS3{F0: -128, F1: "", F2: v3}
            _ = v29
        case 9223372036854775807:
            switch 0 {
            case -2147483648:
                if (42 > 42) {
                    v30 := GenS4{F0: v3, F1: (v5 >> ((0) & 63))}
                    fmt.Println(v3, v3, v2)
                    for fi31 := 0; fi31 < 9; fi31++ {
                        v5 -= (v5 + 43)
                        _ = fi31
                    }
                    _ = v30
                } else {
                    v1 = v1
                }
            case -968:
                v2 = "hello"
            case 0:
                var v32 uint64 = v5
                _ = v32
            }
        }
    case "foo bar":
        if (!(-1 >= -1)) {
            v2 = v28
            switch v4 {
            case 0:
                v2 = "a"
            case 9223372036854775807:
                v33 := GenS4{F0: v3, F1: (v5 - (v5 & (v5 | 0)))}
                _ = v33
            case -9223372036854775808:
                var v34 float32 = ((3.14159 + 1e10) - 1.5)
                _ = v34
            default:
                fmt.Println(v27, v5)
            }
        } else {
            if v6 {
                if (1e10 == 3.14159) {
                    v35 := GenS3{F0: (v27 & v27), F1: "\n\t", F2: (^(v3))}
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        v28 = "\n\t"
                        var v37 uint8 = 255
                        var v38 int = -1
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    var v39 uint32 = 1
                    if true {
                        var v40 int32 = 2147483647
                        _ = v40
                    } else {
                        if v6 {
                            var v41 float32 = 1.5
                            var v42 uint = 0
                            _ = v41
                            _ = v42
                        }
                    }
                    _ = v39
                }
                var v43 int64 = 42
                _ = v43
            } else {
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 float64 = ((3.14159 / 1e10) / 1e10)
                    for fi46 := 0; fi46 < 7; fi46++ {
                        switch int(v1) {
                        case 32768:
                            if v6 {
                                v28 = v28
                                if (fi44 > (fi46 - (fi44 + fi46))) {
                                    var v47 uint8 = 128
                                    _ = v47
                                } else {
                                    switch fi46 {
                                    case 42:
                                        v27 -= (v27 % ((v27) | 1))
                                    case 2147483647:
                                        v26 = append(v26, fi44)
                                    default:
                                        fmt.Println(v4, v27)
                                    }
                                }
                                var v48 uint32 = (v3 << ((uint(v3)) & 63))
                                _ = v48
                            }
                        case 42:
                            var v49 uint = 659
                            _ = v49
                        case 128:
                            v50 := GenS5{F0: true}
                            _ = v50
                        }
                        switch (("foo bar" + (("hello" + (v28 + v28)) + "a")) + ((("" + v28) + "foo bar") + v2)) {
                        case "foo bar":
                            v2 = "hello"
                        case "hello":
                            v26 = append(v26, fi44)
                        case "a":
                            var v51 uint32 = (v3 | v3)
                            fmt.Println(v28, v27)
                            fmt.Println(v28, v28, v5)
                            switch 1 {
                            case 4294967295:
                                switch v2 {
                                case "foo bar":
                                    v1 = 0
                                case "hello":
                                    _ = len(v26)
                                }
                            case 697:
                                if ((!(983 != 1)) && true) {
                                    if v6 {
                                        switch (("a" + v2) + v2) {
                                        case "a":
                                            if (v5 == 2) {
                                                switch fi44 {
                                                case 32768:
                                                    if v6 {
                                                        var v52 int8 = (v27 ^ v27)
                                                        if v6 {
                                                            if v6 {
                                                                _ = len(v26)
                                                                v3 -= v51
                                                                var v53 uint8 = 2
                                                                _ = v53
                                                            }
                                                            var v54 int = (fi44 + (v4 % ((-1) | 1)))
                                                            v26 = append(v26, 42)
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    }
                                                case 128:
                                                    v55 := GenS2{F0: 65535}
                                                    if (0 >= 128) {
                                                        v6 = v6
                                                    } else {
                                                        fmt.Println(v1)
                                                        v56 := GenS5{F0: v6}
                                                        _ = v56
                                                    }
                                                    var v57 uint8 = 145
                                                    _ = v55
                                                    _ = v57
                                                case 9:
                                                    if v6 {
                                                        var v58 bool = true
                                                        if (0.0 <= v45) {
                                                            v26 = append(v26, (fi44 ^ 0))
                                                        }
                                                        v27 = (v27 * (-(v27)))
                                                        var v59 uint32 = 0
                                                        if len(v26) > 0 {
                                                            v26[0] = (fi44 / (((fi44 % ((4294967295) | 1))) | 1))
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                }
                                                var v60 int16 = (v1 & v1)
                                                _ = v60
                                            }
                                        }
                                        switch (("\n\t" + v2) + (v2 + (v2 + "\n\t"))) {
                                        case "hello":
                                            fmt.Println(v27)
                                        }
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                    } else {
                                        fmt.Println(v1)
                                        v6 = (-1 > 42)
                                    }
                                }
                            default:
                                var v61 int16 = (v1 - v1)
                                _ = v61
                            }
                            var v62 uint32 = (v51 + (v51 - v51))
                            _ = v51
                            _ = v62
                        }
                        _ = fi46
                    }
                    _ = fi44
                    _ = v45
                }
            }
        }
    case "hello":
        for fi63 := 0; fi63 < 15; fi63++ {
            v2 = "foo bar"
            _ = fi63
        }
    default:
        if ((int(v27) <= v4) || (false && v6)) {
            var v64 bool = false
            _ = v64
        }
    }
    fmt.Println(v4, v6, v1)
    switch v28 {
    case "a":
        v6 = (-2.25 == 0.1)
    default:
        v65 := GenS5{F0: ((!true) || true)}
        _ = v65
    }
    v27 -= -21
    var v66 uint64 = v5
    v67 := GenS6{F0: (((876 != 1) || (v6 && (uint64(v3) < 2))) || ((int8(v66) != v27) && (false && v6))), F1: 207}
    var v68 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
}
