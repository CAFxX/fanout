package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 1
    var v3 float64 = (3.14159 + (0.0 + (-2.25 / ((-2.25 + (-2.25 + 3.14159)) - 0.0))))
    var v4 uint = 42
    var v5 int32 = 136
    v6 := []int{9223372036854775807}
    var v7 int8 = 1
    _ = len(v6)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    if ((-2.25 * 0.0) >= float32((1e10 + 3.14159))) {
        v1 -= (v1 * v1)
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS1{F0: 127}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    if true {
        var v13 uint64 = 524
        if false {
            _ = len(v6)
            v14 := GenS2{F0: "foo bar", F1: v2}
            _ = v14
        }
        var v15 int64 = v2
        _ = v13
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint8 = 0
            if (true || ((0.1 >= 0.1) || (false && true))) {
                v18 := GenS3{F0: ("" + "\n\t"), F1: 0, F2: (v2 >> (((v4 ^ v4)) & 63))}
                switch ("a" + "foo bar") {
                case "hello":
                    var v19 bool = true
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if false {
                            fmt.Println(v4, fi20, v3)
                            v21 := GenS1{F0: (^(v17))}
                            _ = v21
                        } else {
                            if len(v6) > 0 {
                                v6[0] = -9223372036854775808
                            }
                        }
                        v17 = (v17 / ((uint8(v1)) | 1))
                        v7 = -128
                        _ = fi20
                    }
                case "a":
                    switch -1 {
                    case 2147483647:
                        fmt.Println(v4)
                    case -977:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            switch "" {
                            case "foo bar":
                                v17 = 128
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch fi22 {
                            case 9223372036854775807:
                                v23 := GenS1{F0: v17}
                                _ = v23
                            default:
                                fmt.Println(v5, v17, fi22)
                            }
                            _ = fi22
                        }
                    case -9223372036854775808:
                        v24 := GenS4{F0: -172, F1: (false && (3.14159 > (float64(float32(v3)) / (float64(float32(float64(0.0))) / ((v3 / (v3 / ((v3 + 0.1) - 1e10))) * v3))))), F2: 1}
                        _ = v24
                    }
                }
                _ = v18
            } else {
                var v25 int = fi16
                _ = v25
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint16 = 65535
                if ((v1 & (^(v1))) < v1) {
                    if false {
                        var v28 int = fi26
                        v28 = (fi26 * (fi16 >> ((4294967295) & 63)))
                        fmt.Println(fi26, v17, fi16)
                        _ = v28
                    } else {
                        if true {
                            var v29 string = (("a" + ("\n\t" + "foo bar")) + (("\n\t" + "a") + "\n\t"))
                            v29 = ("\n\t" + "hello")
                            _ = v29
                        } else {
                            var v30 uint64 = uint64(v4)
                            fmt.Println(v4, v4, fi26)
                            fmt.Println(v2, v7, v7)
                            var v31 uint = 470
                            fmt.Println(fi26, fi26)
                            _ = v30
                            _ = v31
                        }
                    }
                    v17 += (^(v17))
                    v32 := GenS1{F0: (v17 % (((v17 | (v17 - 128))) | 1))}
                    _ = v32
                }
                v33 := GenS3{F0: (("hello" + ("foo bar" + ("a" + "\n\t"))) + ("foo bar" + ("hello" + ""))), F1: 42, F2: 0}
                _ = fi26
                _ = v27
                _ = v33
            }
            _ = fi16
            _ = v17
        }
        switch 32767 {
        case -9223372036854775808:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 860:
            if false {
                for fi34 := 0; fi34 < 2; fi34++ {
                    switch fi34 {
                    case 659:
                        var v35 int8 = v7
                        _ = v35
                    case 1:
                        for fi36 := 0; fi36 < 18; fi36++ {
                            if false {
                                v37 := GenS3{F0: "\n\t", F1: uint64(v4), F2: 0}
                                v1 += (v1 / (((v1 << ((v4) & 63))) | 1))
                                _ = v37
                            }
                            if ((v2 < (v2 | (v2 >> ((0) & 63)))) && (715 == v1)) {
                                v6 = append(v6, int(v7))
                                var v38 int = (fi36 << (((v4 >> ((v4) & 63))) & 63))
                                _ = v38
                            } else {
                                if (3.14159 >= float64(1e10)) {
                                    if (v5 < (-(v5))) {
                                        if false {
                                            if false {
                                                v39 := GenS5{F0: 1}
                                                _ = v39
                                            }
                                            var v40 uint8 = 1
                                            _ = v40
                                        }
                                        switch "" {
                                        case "":
                                            v41 := GenS2{F0: "\n\t", F1: v2}
                                            _ = v41
                                        default:
                                            v7 = v7
                                        }
                                        var v42 uint8 = 42
                                        _ = v42
                                    }
                                } else {
                                    v43 := GenS3{F0: ("a" + "a"), F1: 18446744073709551615, F2: (v2 + (v2 * -9223372036854775808))}
                                    v4 = v4
                                    _ = v43
                                }
                                v6 = append(v6, (fi34 + fi34))
                            }
                            _ = fi36
                        }
                    default:
                        var v44 int64 = v2
                        _ = v44
                    }
                    _ = len(v6)
                    v1 = (v1 % (((v1 + (v1 >> (((v4 / (((v4 - v4)) | 1))) & 63)))) | 1))
                    _ = fi34
                }
            } else {
                var v45 string = "hello"
                var v46 float64 = v3
                var v47 int16 = 0
                _ = v45
                _ = v46
                _ = v47
            }
        case 127:
            fmt.Println(v3, v1, v4)
        }
        _ = len(v6)
    } else {
        var v48 int = -828
        v6 = append(v6, (v48 % ((2147483647) | 1)))
        _ = len(v6)
        _ = v48
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if false {
        if len(v6) > 0 {
            v6[0] = 32767
        }
    } else {
        v3 = 3.14159
    }
    var v49 int16 = 0
    var v50 float32 = 1.5
    var v51 uint = (v4 >> ((v4) & 63))
    v52 := GenS6{F0: "foo bar"}
    fmt.Println(v7, v1)
    for fi53 := 0; fi53 < 19; fi53++ {
        fmt.Println(v5)
        if false {
            var v54 uint32 = 42
            var v55 uint16 = 255
            for fi56 := 0; fi56 < 6; fi56++ {
                v7 = (v7 * 42)
                _ = fi56
            }
            _ = v54
            _ = v55
        }
        _ = fi53
    }
    v57 := GenS3{F0: "foo bar", F1: 4294967295, F2: (v2 - 1)}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v57
    _ = v6
    _ = v7
}
