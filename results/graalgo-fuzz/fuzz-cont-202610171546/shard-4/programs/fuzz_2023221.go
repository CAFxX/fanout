package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 int64 }
func genFn0(p0 int16, p1 uint8) uint32 {
    v5 := []int{485, 127, 0}
    p1 = (^(p1))
    switch 237 {
    case 2147483647:
        for fi6 := 0; fi6 < 20; fi6++ {
            p1 -= 233
            var v7 uint = 42
            var v8 int64 = 1
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    if false {
        var v9 int16 = p0
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v9, v9, p0)
        }
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 0, F1: 1}
        var v12 string = "a"
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    return 65535
}
func genFn1(p0 uint64) int8 {
    v13 := []int{}
    if (32767 < 0) {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: -122, F1: 2}
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 0
        _ = v16
    } else {
        var v17 uint64 = (p0 - p0)
        if (42 != 127) {
            var v18 uint16 = 0
            var v19 int32 = 0
            if ((-823 >= 127) || true) {
                switch -15 {
                case 128:
                    switch -9223372036854775808 {
                    case 32768:
                        if (!true) {
                            if len(v13) > 0 {
                                v13[0] = 4294967295
                            }
                            v17 = 1
                            for fi20 := 0; fi20 < 9; fi20++ {
                                v21 := GenS1{F0: 100, F1: 255}
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    if (fi20 <= fi22) {
                                        var v23 float32 = -2.25
                                        switch fi22 {
                                        case -347:
                                            var v24 int16 = -955
                                            _ = v24
                                        }
                                        v17 = p0
                                        _ = v23
                                    } else {
                                        if true {
                                            v13 = append(v13, 4294967295)
                                            var v25 float64 = 3.14159
                                            fmt.Println(fi22)
                                            _ = v25
                                        } else {
                                            switch (-(fi22)) {
                                            case 32768:
                                                if true {
                                                    fmt.Println(v19, v18, v17)
                                                }
                                            case 32767:
                                                var v26 float64 = ((((0.1 * (((1e10 + ((-2.25 - 3.14159) * (3.14159 + -2.25))) + 1e10) + -2.25)) + 0.1) - 0.1) / ((float64(-2.25) / 0.1) - float64((0.0 / (0.1 + float32((float64(0.0) + (1e10 + float64(((3.14159 - -2.25) + (1.5 * 3.14159)))))))))))
                                                _ = v26
                                            }
                                            fmt.Println(p0, v17)
                                        }
                                        var v27 int32 = (-(v19))
                                        _ = v27
                                    }
                                    p0 += (p0 + 644)
                                    fmt.Println(v19, fi22, p0)
                                    _ = fi22
                                }
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            v28 := GenS1{F0: 0, F1: 429}
                            v29 := GenS2{F0: int(v19), F1: -1, F2: (p0 & v17)}
                            _ = v28
                            _ = v29
                        }
                    case -1:
                        v19 += (^(v19))
                    case -395:
                        if true {
                            v30 := GenS3{F0: ("\n\t" + "a")}
                            var v31 uint = 755
                            v32 := GenS2{F0: -164, F1: -128, F2: 0}
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            if len(v13) > 0 {
                                v13[0] = 578
                            }
                        }
                    }
                default:
                    var v33 int8 = 0
                    _ = v33
                }
                var v34 uint64 = (v17 >> ((4294967295) & 63))
                _ = v34
            }
            _ = v18
            _ = v19
        } else {
            v13 = append(v13, 32767)
            for fi35 := 0; fi35 < 17; fi35++ {
                v17 -= (^(v17))
                if true {
                    switch fi35 {
                    case 37:
                        for fi36 := 0; fi36 < 2; fi36++ {
                            switch (fi36 + (fi35 - fi36)) {
                            case -858:
                                v13 = append(v13, (fi35 - (-(fi36))))
                            case -9223372036854775808:
                                p0 = p0
                            default:
                                var v37 uint8 = 68
                                _ = v37
                            }
                            var v38 uint = 4294967295
                            v39 := GenS4{F0: 0, F1: 1, F2: 43}
                            _ = fi36
                            _ = v38
                            _ = v39
                        }
                    case 0:
                        fmt.Println(v17, fi35)
                    case -952:
                        var v40 uint = 4294967295
                        _ = v40
                    }
                }
                _ = fi35
            }
            var v41 int32 = -706
            _ = v41
        }
        v17 -= (p0 | (p0 & (p0 - (^(v17)))))
        _ = v17
    }
    p0 += p0
    var v42 int = 32768
    _ = p0
    _ = v13
    _ = v42
    return -128
}

func main() {
    var v1 float32 = (1.5 * float32(float64((-2.25 - 3.14159))))
    v2 := []uint32{0}
    var v3 string = ("foo bar" + "a")
    v4 := make(map[int]uint32)
    if (-2147483648 > -2147483648) {
        if false {
            var v43 int64 = 9223372036854775807
            fmt.Println(v43, v3, v3)
            _ = v43
        }
        delete(v4, 831)
    }
    v3 = v3
    if false {
        for fi44 := 0; fi44 < 3; fi44++ {
            fmt.Println(v3)
            var v45 uint32 = genFn0(671, 0)
            if false {
                v46 := GenS4{F0: 0, F1: -2147483648, F2: 32767}
                for fi47 := 0; fi47 < 10; fi47++ {
                    var v48 int64 = -279
                    fmt.Println(fi44, v3)
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
            _ = fi44
            _ = v45
        }
    }
    v3 = v3
    v3 = v3
    for fi49 := 0; fi49 < 8; fi49++ {
        _ = len(v2)
        fmt.Println(v3, v3)
        v50 := GenS1{F0: 0, F1: 336}
        _ = fi49
        _ = v50
    }
    fmt.Println(v3)
    var v51 float64 = (-2.25 / float64((((3.14159 + v1) / 1.0) / v1)))
    var v52 uint = 726
    fmt.Println(v1, v51)
    for fi53 := 0; fi53 < 4; fi53++ {
        switch v3 {
        case "a":
            switch "hello" {
            case "":
                fmt.Println(v3, v1)
            case "foo bar":
                switch (fi53 << ((v52) & 63)) {
                case -325:
                    v54 := GenS5{F0: (("foo bar" + v3) + v3), F1: 208, F2: 9223372036854775807}
                    _ = v54
                case 4294967295:
                    if false {
                        switch 128 {
                        case -988:
                            if true {
                                for fi55 := 0; fi55 < 18; fi55++ {
                                    v1 = 0.1
                                    var v56 uint16 = 1
                                    v57 := GenS1{F0: -1, F1: v52}
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                            } else {
                                v58 := GenS5{F0: v3, F1: 247, F2: 9223372036854775807}
                                v51 -= (v51 / v51)
                                _ = v58
                            }
                        case 846:
                            for fi59 := 0; fi59 < 5; fi59++ {
                                var v60 int32 = 1
                                v51 = 3.14159
                                _ = fi59
                                _ = v60
                            }
                        }
                    }
                default:
                    v1 -= (-2.25 - 3.14159)
                }
            case "a":
                v61 := GenS2{F0: (fi53 >> (((v52 + (v52 * 52))) & 63)), F1: -47, F2: 18446744073709551615}
                _ = v61
            default:
                _ = len(v4)
            }
        default:
            for fi62 := 0; fi62 < 12; fi62++ {
                _ = len(v2)
                v3 = v3
                _ = fi62
            }
        }
        _ = fi53
    }
    _ = len(v2)
    switch "a" {
    case "a":
        var v63 int64 = -9223372036854775808
        _ = v63
    case "hello":
        if true {
            var v64 int64 = 9223372036854775807
            _ = v64
        }
    }
    v65 := GenS4{F0: 0, F1: 42, F2: 32767}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v52 = (v52 * (v52 << (((v52 & (v52 - 0))) & 63)))
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v65
}
