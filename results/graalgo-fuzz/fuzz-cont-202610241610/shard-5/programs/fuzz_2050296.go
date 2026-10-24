package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }
type GenS4 struct { F0 int32 }
func genFn0() uint64 {
    var v6 uint64 = 2
    if (false || false) {
        var v7 int8 = -58
        v7 = (v7 ^ v7)
        _ = v7
    }
    v6 += (v6 ^ (v6 >> ((2) & 63)))
    fmt.Println(v6)
    _ = v6
    return (v6 >> ((255) & 63))
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint8 = 0
            if true {
                v11 := []int{fi8, 9223372036854775807, fi8, 0}
                _ = v11
            } else {
                switch "" {
                case "\n\t":
                    var v12 uint16 = 892
                    fmt.Println(fi9, fi8, v12)
                    var v13 bool = true
                    if v13 {
                        if v13 {
                            v13 = (!v13)
                        } else {
                            fmt.Println(v10, fi8, v13)
                        }
                    }
                    _ = v12
                    _ = v13
                default:
                    var v14 int32 = -463
                    _ = v14
                }
            }
            _ = fi9
            _ = v10
        }
        var v15 int64 = 1
        _ = fi8
        _ = v15
    }
    var v16 int64 = 1
    if (!(((1.5 - 1.5) / (0.0 - 0.1)) < (0.0 * (1e10 + float32(1e10))))) {
        if (!((0.1 / float64(1.5)) == ((3.14159 * 1.5) + ((3.14159 * 1e10) + ((0.1 * (-2.25 - -2.25)) + (0.0 * 0.1)))))) {
            var v17 int32 = 1
            var v18 int = 127
            v19 := GenS1{F0: 1}
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                if true {
                    var v21 uint64 = 555
                    v21 -= (v21 << ((463) & 63))
                    _ = v21
                } else {
                    var v22 uint = 115
                    if true {
                        var v23 uint8 = 1
                        if false {
                            fmt.Println(v16, fi20)
                            v24 := GenS1{F0: 1}
                            v23 = (^(v23))
                            _ = v24
                        }
                        _ = v23
                    }
                    fmt.Println(v22, v16, v22)
                    _ = v22
                }
                _ = fi20
            }
            v25 := GenS1{F0: 0}
            _ = v25
        }
        if (42 >= -978) {
            var v26 string = ("foo bar" + "")
            v26 = ("hello" + v26)
            for fi27 := 0; fi27 < 9; fi27++ {
                v26 = v26
                _ = fi27
            }
            _ = v26
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 int = fi28
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 5; fi30++ {
                v16 = v16
                v31 := []int{155, fi30, (fi30 | fi30)}
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                v33 := GenS2{F0: (!true), F1: true, F2: 18446744073709551615}
                v34 := []int{(fi32 / ((fi32) | 1)), -434, fi32, (fi32 << ((255) & 63))}
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    }
    _ = v16
    return 42
}

func main() {
    var v1 uint16 = 136
    var v2 string = "hello"
    var v3 int8 = -13
    var v4 uint16 = (v1 + 1)
    v5 := []int{32768}
    v1 -= (v4 / ((v4) | 1))
    fmt.Println(v4, v4, v4)
    v3 -= v3
    for fi35 := 0; fi35 < 4; fi35++ {
        v2 = "hello"
        if (((v1 >> ((4294967295) & 63)) == (v1 + (v1 + 65535))) || true) {
            v4 = v4
        } else {
            if ((!false) && (-32768 < -949)) {
                for fi36 := 0; fi36 < 2; fi36++ {
                    v2 = ("hello" + (v2 + "hello"))
                    _ = fi36
                }
                v3 -= v3
            }
            switch -1 {
            case 32768:
                if ((!(true && false)) || true) {
                    if false {
                        if ((!((0.0 - (-2.25 / 1.0)) > 0.1)) && false) {
                            for fi37 := 0; fi37 < 14; fi37++ {
                                v2 = "foo bar"
                                if false {
                                    switch fi35 {
                                    case 32768:
                                        if true {
                                            fmt.Println(v1)
                                            v5 = append(v5, (fi37 >> ((2) & 63)))
                                        } else {
                                            switch v2 {
                                            case "a":
                                                v2 = (v2 + ((v2 + v2) + v2))
                                            case "":
                                                var v38 int16 = 42
                                                _ = v38
                                            }
                                            var v39 uint32 = 30
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v39
                                        }
                                    case 2147483647:
                                        v3 = -77
                                    case 1:
                                        v40 := GenS3{F0: 65535, F1: int64(v4), F2: 0}
                                        _ = v40
                                    }
                                    if false {
                                        if (!(49 > 127)) {
                                            var v41 float32 = ((0.0 * 1e10) / ((((3.14159 - ((0.0 / 1e10) * 0.1)) - 0.1) * 3.14159) * 1e10))
                                            _ = v41
                                        } else {
                                            v42 := GenS3{F0: 42, F1: 469, F2: genFn1()}
                                            _ = v42
                                        }
                                        v3 -= -128
                                        fmt.Println(v3)
                                        if false {
                                            if true {
                                                v43 := GenS1{F0: -523}
                                                if (!false) {
                                                    if (-845 == 1) {
                                                        v1 = v4
                                                        v1 -= (v4 + (v1 + 255))
                                                    } else {
                                                        fmt.Println(fi35, fi35)
                                                    }
                                                    var v44 int = (fi35 - fi35)
                                                    v44 += 1
                                                    _ = v44
                                                } else {
                                                    switch v2 {
                                                    case "\n\t":
                                                        var v45 uint16 = 319
                                                        _ = v45
                                                    default:
                                                        if len(v5) > 0 {
                                                            v5[0] = fi37
                                                        }
                                                    }
                                                    v46 := GenS2{F0: true, F1: (false || false), F2: 18446744073709551615}
                                                    _ = v46
                                                }
                                                _ = v43
                                            } else {
                                                switch "" {
                                                case "":
                                                    v4 -= (v4 | (v4 ^ (v1 * v1)))
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        v2 = ""
                                        switch 9223372036854775807 {
                                        case -2147483648:
                                            switch (v2 + (v2 + ("a" + (((v2 + v2) + v2) + "\n\t")))) {
                                            case "\n\t":
                                                var v47 int64 = -1
                                                _ = v47
                                            case "foo bar":
                                                fmt.Println(v3, fi37, v1)
                                            default:
                                                var v48 uint16 = v1
                                                _ = v48
                                            }
                                        case 42:
                                            var v49 int16 = 42
                                            _ = v49
                                        default:
                                            if ((v3 - (v3 | (v3 * 97))) == (v3 ^ v3)) {
                                                fmt.Println(fi35)
                                                if true {
                                                    var v50 bool = false
                                                    if len(v5) > 0 {
                                                        v5[0] = fi37
                                                    }
                                                    var v51 int64 = 9223372036854775807
                                                    switch fi35 {
                                                    case 0:
                                                        var v52 int64 = v51
                                                        _ = v52
                                                    case 42:
                                                        v50 = (0 > 94)
                                                    default:
                                                        v51 += v51
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                } else {
                                                    var v53 bool = (32767 == -1)
                                                    v2 = "\n\t"
                                                    _ = v53
                                                }
                                            } else {
                                                v2 = "foo bar"
                                            }
                                        }
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                fmt.Println(v1, fi37)
                                _ = fi37
                            }
                            var v54 float64 = 3.14159
                            _ = v54
                        }
                    } else {
                        switch (("\n\t" + (v2 + ("a" + "\n\t"))) + "foo bar") {
                        case "a":
                            v3 = v3
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            for fi55 := 0; fi55 < 10; fi55++ {
                                v2 = (("hello" + ((v2 + "\n\t") + v2)) + v2)
                                _ = fi55
                            }
                        }
                    }
                    var v56 int8 = (^(v3))
                    v5 = append(v5, (fi35 >> ((42) & 63)))
                    _ = v56
                }
            }
            var v57 float64 = (float64((3.14159 / 3.14159)) * 3.14159)
            _ = v57
        }
        v1 -= 42
        _ = fi35
    }
    v1 = (v1 >> ((0) & 63))
    if ((414 <= int16(v3)) && true) {
        switch 4294967295 {
        case -904:
            for fi58 := 0; fi58 < 14; fi58++ {
                switch fi58 {
                case 32768:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi58
            }
        case 1:
            fmt.Println(v2, v2)
        default:
            v2 = v2
        }
        fmt.Println(v2, v1)
        var v59 float64 = ((((0.1 + 0.0) * (0.0 / ((((-2.25 * 1e10) - (-2.25 / 1.0)) / float64(1e10)) * ((0.0 + -2.25) + 3.14159)))) / float64(-2.25)) + (1e10 + 0.1))
        _ = v59
    }
    if true {
        for fi60 := 0; fi60 < 11; fi60++ {
            if len(v5) > 0 {
                v5[0] = -2147483648
            }
            _ = fi60
        }
    }
    v3 += (v3 >> ((255) & 63))
    v61 := GenS1{F0: 1000}
    v3 = (v3 ^ v3)
    v3 = v3
    switch 4294967295 {
    case 657:
        fmt.Println(v4, v4)
    case 0:
        if len(v5) > 0 {
            v5[0] = 562
        }
    case 454:
        switch "" {
        case "foo bar":
            fmt.Println(v1, v4, v1)
        }
    }
    var v62 int64 = -9223372036854775808
    for fi63 := 0; fi63 < 9; fi63++ {
        var v64 int = (-(fi63))
        if true {
            v65 := GenS1{F0: 42}
            var v66 uint16 = 255
            fmt.Println(v62)
            _ = v65
            _ = v66
        } else {
            v67 := GenS4{F0: 42}
            _ = len(v5)
            _ = v67
        }
        _ = fi63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
}
