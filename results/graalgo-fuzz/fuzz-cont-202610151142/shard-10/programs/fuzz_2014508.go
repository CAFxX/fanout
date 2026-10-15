package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 int64; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 uint64) uint64 {
    v6 := []int{502, 2147483647}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: fi7, F1: -9223372036854775808, F2: 255}
        var v9 string = "\n\t"
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v10 := GenS2{F0: 244, F1: "hello", F2: 65535}
    p0 += (p0 >> ((777) & 63))
    _ = p0
    _ = v10
    _ = v6
    return 4294967295
}
func genFn1(p0 uint8, p1 uint, p2 int8) int8 {
    v11 := GenS2{F0: 339, F1: ("foo bar" + "hello"), F2: p1}
    for fi12 := 0; fi12 < 17; fi12++ {
        switch fi12 {
        case 32767:
            p2 += p2
        case 1:
            v13 := []int{fi12, fi12, fi12, fi12}
            _ = v13
        default:
            var v14 int8 = (p2 - 100)
            _ = v14
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            p2 -= 42
            p2 -= p2
            _ = fi15
        }
        v16 := GenS3{F0: genFn0(279), F1: false, F2: 42}
        _ = fi12
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 bool = (2 == 260)
    var v2 uint32 = 2
    v3 := []uint32{v2}
    v4 := []uint32{(v2 << ((4294967295) & 63))}
    v5 := make(map[uint32]bool)
    switch "hello" {
    case "hello":
        var v17 int8 = -128
        _ = v17
    case "a":
        var v18 int = 32767
        _ = v18
    case "foo bar":
        v19 := GenS4{F0: genFn0(432), F1: 42, F2: 0}
        _ = v19
    }
    if (649 < 1) {
        fmt.Println(v2, v2)
        v1 = (-1 <= 32767)
        v1 = (v1 && ((-2.25 != 1e10) && (32767 > 770)))
    } else {
        var v20 int8 = genFn1(14, 198, 127)
        _ = v20
    }
    fmt.Println(v2, v2)
    v2 = (v2 % ((2) | 1))
    var v21 int32 = 42
    if (865 <= -1) {
        if true {
            fmt.Println(v21)
            var v22 int8 = genFn1(1, 65535, 127)
            _ = v22
        } else {
            v23 := GenS4{F0: 0, F1: 1, F2: 569}
            if true {
                if (-52 > (v21 ^ (v21 % ((v21) | 1)))) {
                    v3 = append(v3, 4294967295)
                    fmt.Println(v2)
                } else {
                    var v24 uint16 = 0
                    var v25 uint = 1
                    v4 = append(v4, (v2 << ((18446744073709551615) & 63)))
                    _ = v24
                    _ = v25
                }
                v1 = ("foo bar" >= "a")
                for fi26 := 0; fi26 < 4; fi26++ {
                    v21 = (v21 >> ((862) & 63))
                    v27 := GenS4{F0: 18446744073709551615, F1: 120, F2: 4294967295}
                    _ = fi26
                    _ = v27
                }
            }
            if (127 != 42) {
                if (-9223372036854775808 == -9223372036854775808) {
                    v28 := GenS5{F0: 0}
                    if (1 < genFn0(4294967295)) {
                        switch 4294967295 {
                        case 42:
                            v1 = v1
                        case 4294967295:
                            v2 = v2
                        case -15:
                            v21 += (v21 / ((2147483647) | 1))
                        default:
                            v2 = (v2 ^ v2)
                        }
                        fmt.Println(v2, v21, v1)
                        if v1 {
                            v1 = (622 > 587)
                            v21 += v21
                        }
                    } else {
                        switch ("hello" + "a") {
                        case "hello":
                            var v29 int16 = 0
                            _ = v29
                        case "a":
                            fmt.Println(v2, v21)
                        default:
                            for fi30 := 0; fi30 < 18; fi30++ {
                                if v1 {
                                    switch (fi30 | 42) {
                                    case -1:
                                        if (159 != 158) {
                                            switch -41 {
                                            case 2147483647:
                                                fmt.Println(fi30, v2, v21)
                                            case -2147483648:
                                                var v31 bool = v1
                                                _ = v31
                                            case 887:
                                                fmt.Println(v1, v21)
                                            default:
                                                var v32 bool = (v1 && (1.5 <= 0.1))
                                                _ = v32
                                            }
                                        }
                                    default:
                                        switch fi30 {
                                        case 32768:
                                            if true {
                                                var v33 int32 = (v21 & (v21 + v21))
                                                _ = v33
                                            }
                                        case -1:
                                            v1 = (0.1 < float32((-2.25 * 0.0)))
                                        case -9223372036854775808:
                                            if (571 < 0) {
                                                var v34 string = ("\n\t" + (("\n\t" + "\n\t") + "\n\t"))
                                                _ = v34
                                            } else {
                                                if false {
                                                    var v35 uint8 = 0
                                                    _ = v35
                                                }
                                                v21 = v21
                                                for fi36 := 0; fi36 < 4; fi36++ {
                                                    var v37 int8 = -80
                                                    _ = fi36
                                                    _ = v37
                                                }
                                            }
                                        default:
                                            fmt.Println(v21, v21)
                                        }
                                    }
                                    switch (fi30 & (fi30 << ((453) & 63))) {
                                    case 1:
                                        v2 = 42
                                    case 9223372036854775807:
                                        var v38 float64 = (((3.14159 - (0.1 - (float64((((0.0 - (-2.25 * 1.5)) / 1.5) * -2.25)) - float64((1e10 * (3.14159 / 1.5)))))) / (1.5 * 0.1)) / 1.0)
                                        _ = v38
                                    default:
                                        v39 := GenS3{F0: genFn0(4294967295), F1: (false || (-269 > 0)), F2: 4294967295}
                                        _ = v39
                                    }
                                } else {
                                    v40 := GenS3{F0: 18446744073709551615, F1: true, F2: 2}
                                    v21 = v21
                                    for fi41 := 0; fi41 < 14; fi41++ {
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                v1 = true
                                                if len(v4) > 0 {
                                                    v4[0] = (v2 + (v2 | (v2 - 698)))
                                                }
                                            }
                                        case "foo bar":
                                            var v42 float32 = 1e10
                                            _ = v42
                                        case "":
                                            v1 = v1
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        _ = fi41
                                    }
                                    _ = v40
                                }
                                switch (fi30 | (fi30 * -1)) {
                                case 4294967295:
                                    v43 := GenS6{F0: -1}
                                    _ = v43
                                case 127:
                                    fmt.Println(v2, v1)
                                case 32768:
                                    if (0 < -1) {
                                        var v44 float64 = 1e10
                                        var v45 uint32 = (v2 << ((29) & 63))
                                        fmt.Println(v44, v1)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 uint64 = 236
                                        var v47 string = ("\n\t" + "")
                                        for fi48 := 0; fi48 < 20; fi48++ {
                                            v46 += (v46 % (((v46 >> ((42) & 63))) | 1))
                                            switch 9223372036854775807 {
                                            case 0:
                                                var v49 string = ("" + (v47 + v47))
                                                _ = v49
                                            case 32767:
                                                v50 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: 18446744073709551615}
                                                _ = v50
                                            case 515:
                                                if (10 >= 120) {
                                                    v1 = false
                                                    v51 := GenS7{F0: 1, F1: 9223372036854775807, F2: 102}
                                                    var v52 uint16 = 23
                                                    switch fi30 {
                                                    case -654:
                                                        var v53 int64 = 522
                                                        _ = v53
                                                    case -1:
                                                        if (v47 < "") {
                                                            if (v1 || (v52 == (v52 & v52))) {
                                                                v54 := GenS8{F0: v46, F1: v2}
                                                                v46 = 4294967295
                                                                var v55 bool = (42 > 69)
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        }
                                                    case 2147483647:
                                                        var v56 float64 = (3.14159 + 1e10)
                                                        _ = v56
                                                    default:
                                                        v57 := GenS2{F0: v21, F1: "\n\t", F2: 65535}
                                                        _ = v57
                                                    }
                                                    _ = len(v5)
                                                    if (921 >= 32767) {
                                                        v58 := GenS9{F0: 100, F1: v52}
                                                        switch fi48 {
                                                        case 32767:
                                                            fmt.Println(v1, v52, v21)
                                                        case 127:
                                                            v59 := GenS10{F0: 734, F1: 42, F2: (v21 % ((v21) | 1))}
                                                            _ = v59
                                                        }
                                                        _ = v58
                                                    }
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    var v60 uint16 = 0
                                                    v2 += v2
                                                    _ = v60
                                                }
                                            default:
                                                v5[v2] = (v1 && (v1 || (97 < genFn1(252, 18446744073709551615, -128))))
                                            }
                                            _ = fi48
                                        }
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                _ = fi30
                            }
                        }
                        fmt.Println(v1)
                    }
                    _ = v28
                }
            } else {
                var v61 uint64 = 2
                _ = v61
            }
            _ = v23
        }
        v3 = append(v3, 1)
    }
    fmt.Println(v1, v1)
    switch 800 {
    case 948:
        v2 += 2147483648
    case 0:
        v1 = (v1 && ((!(v2 >= (^(v2)))) || ((v1 || v1) || true)))
    }
    _ = len(v5)
    for fi62 := 0; fi62 < 12; fi62++ {
        var v63 int64 = 42
        _ = fi62
        _ = v63
    }
    var v64 float32 = -2.25
    if true {
        var v65 int64 = 0
        fmt.Println(v1, v1, v65)
        for fi66 := 0; fi66 < 16; fi66++ {
            fmt.Println(fi66, v64)
            var v67 int = (fi66 - fi66)
            _ = fi66
            _ = v67
        }
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v64
}
