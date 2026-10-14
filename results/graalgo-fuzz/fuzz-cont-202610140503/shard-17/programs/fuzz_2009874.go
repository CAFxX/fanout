package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int64 = 1
    var v2 int = 4294967295
    v3 := []int{32767, (v2 & v2), -126}
    var v4 float64 = 3.14159
    var v5 string = "hello"
    fmt.Println(v5, v4, v4)
    if len(v3) > 0 {
        v3[0] = 445
    }
    v2 = (v2 - v2)
    v6 := GenS1{F0: v1, F1: -128, F2: ((4294967295 == 492) || (195 >= 32767))}
    if (-1 == -1) {
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v4)
            for fi8 := 0; fi8 < 4; fi8++ {
                if true {
                    v4 += v4
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 9223372036854775807
                        }
                        v3 = append(v3, fi7)
                    } else {
                        var v9 int32 = 2147483647
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v9
                    }
                    if false {
                        if ((621 <= 564) && false) {
                            var v10 float64 = (v4 + ((3.14159 - ((v4 * (float64(1.5) / v4)) - 1.5)) - v4))
                            v5 = (v5 + v5)
                            v1 = -9223372036854775808
                            v10 -= (float64(0.0) / (float64(3.14159) - -2.25))
                            v11 := GenS1{F0: (-(v1)), F1: -128, F2: false}
                            v2 -= 128
                            var v12 int64 = (v1 | v1)
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                        var v13 uint = 255
                        var v14 int = fi8
                        v15 := GenS1{F0: v1, F1: -1, F2: true}
                        if false {
                            if true {
                                var v16 uint = v13
                                if (2147483647 <= 1) {
                                    var v17 uint8 = 200
                                    _ = v17
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v16
                            } else {
                                v18 := GenS1{F0: 42, F1: -128, F2: false}
                                _ = v18
                            }
                        }
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi8
            }
            v5 = ""
            _ = fi7
        }
        switch ((((v5 + v5) + v5) + (v5 + v5)) + "hello") {
        case "hello":
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 += (v1 & (^(v1)))
                _ = fi19
            }
        }
    }
    switch 486 {
    case 32768:
        for fi20 := 0; fi20 < 10; fi20++ {
            if false {
                if (!true) {
                    if true {
                        if true {
                            fmt.Println(fi20, v1, fi20)
                            var v21 uint16 = 65535
                            if (1 >= 1000) {
                                v1 += (v1 * 9223372036854775807)
                                var v22 uint = 2
                                var v23 int64 = v1
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if false {
                                _ = len(v3)
                                fmt.Println(fi24, v5, v1)
                                v5 = v5
                            }
                            switch 0 {
                            case 1:
                                v2 = fi20
                            case -1:
                                var v25 int8 = -116
                                _ = v25
                            default:
                                v26 := GenS1{F0: -528, F1: 100, F2: true}
                                _ = v26
                            }
                            v27 := GenS1{F0: v1, F1: -37, F2: true}
                            _ = fi24
                            _ = v27
                        }
                        switch 32767 {
                        case 455:
                            fmt.Println(v4, v2, fi20)
                        case 374:
                            _ = len(v3)
                        case 262:
                            v4 = (v4 + ((-2.25 - 1.5) + -2.25))
                        }
                    } else {
                        v2 += 4294967295
                        fmt.Println(fi20, v1, v2)
                        v1 -= (v1 | v1)
                    }
                    if false {
                        if true {
                            if ((v4 + v4) <= v4) {
                                fmt.Println(v2, v2, fi20)
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    if false {
                                        v29 := GenS2{F0: 89, F1: fi28, F2: 296}
                                        v5 = v5
                                        var v30 uint32 = 2147483648
                                        var v31 uint = 42
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = fi28
                                }
                                if true {
                                    _ = len(v3)
                                } else {
                                    var v32 int16 = 0
                                    if true {
                                        v33 := GenS3{F0: 0, F1: 42}
                                        _ = v33
                                    } else {
                                        var v34 int64 = 0
                                        switch v2 {
                                        case -9223372036854775808:
                                            if (2 < 65535) {
                                                for fi35 := 0; fi35 < 18; fi35++ {
                                                    var v36 string = ("foo bar" + "\n\t")
                                                    var v37 int32 = 117
                                                    switch ("foo bar" + (v5 + (((v36 + ((v36 + ((("hello" + ("a" + v36)) + v36) + (v36 + "hello"))) + (v36 + v36))) + ((((("foo bar" + v36) + "foo bar") + ((v5 + "a") + v36)) + "\n\t") + ((v36 + ("" + "foo bar")) + v5))) + ""))) {
                                                    case "\n\t":
                                                        v38 := GenS2{F0: -91, F1: -9223372036854775808, F2: 45}
                                                        _ = v38
                                                    case "":
                                                        v2 -= 32768
                                                    }
                                                    _ = fi35
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v39 int = (-(fi20))
                                                _ = v39
                                            }
                                        case 4294967295:
                                            v40 := GenS1{F0: (v1 ^ (v34 >> ((18446744073709551615) & 63))), F1: -63, F2: false}
                                            _ = v40
                                        default:
                                            v1 = v1
                                        }
                                        _ = v34
                                    }
                                    _ = v32
                                }
                            } else {
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    if false {
                                        var v42 uint16 = 1
                                        _ = v42
                                    }
                                    _ = fi41
                                }
                            }
                            v3 = append(v3, fi20)
                        } else {
                            switch ("hello" + ("" + v5)) {
                            case "foo bar":
                                v4 += v4
                            case "\n\t":
                                for fi43 := 0; fi43 < 10; fi43++ {
                                    v44 := GenS4{F0: 42, F1: ((v5 + v5) + v5)}
                                    var v45 int = -1
                                    var v46 float32 = (1.5 + 0.0)
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            case "hello":
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 int32 = int32(v2)
                                    v49 := GenS5{F0: 1}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            default:
                                for fi50 := 0; fi50 < 20; fi50++ {
                                    var v51 uint64 = 2
                                    fmt.Println(fi20, v1, fi20)
                                    var v52 int32 = -2147483648
                                    v53 := GenS3{F0: 65535, F1: (v51 ^ v51)}
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                            }
                            v2 -= v2
                            if false {
                                v54 := GenS3{F0: 24, F1: 42}
                                var v55 int8 = -1
                                _ = v54
                                _ = v55
                            }
                        }
                        for fi56 := 0; fi56 < 6; fi56++ {
                            v2 -= 0
                            _ = fi56
                        }
                    }
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v1 = (v1 & v1)
                        var v58 uint64 = 42
                        var v59 uint = 424
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = (v2 / (((-(v2))) | 1))
                }
                fmt.Println(v1, v5)
            }
            _ = fi20
        }
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v60 float64 = (v4 * (0.1 + v4))
    if len(v3) > 0 {
        v3[0] = v2
    }
    for fi61 := 0; fi61 < 1; fi61++ {
        if false {
            var v62 float64 = 0.1
            var v63 int64 = (-(v1))
            _ = v62
            _ = v63
        }
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                v1 = v1
                v5 = (("hello" + "") + v5)
            }
            v2 = fi61
        }
        _ = fi61
    }
    for fi64 := 0; fi64 < 13; fi64++ {
        var v65 int64 = v1
        v4 = (v60 + (0.0 * v4))
        _ = fi64
        _ = v65
    }
    var v66 uint8 = 128
    for fi67 := 0; fi67 < 5; fi67++ {
        v5 = ((v5 + (v5 + (v5 + ("hello" + v5)))) + ("hello" + "\n\t"))
        fmt.Println(v2, v1, v2)
        v4 += v60
        _ = fi67
    }
    v3 = append(v3, (v2 & v2))
    fmt.Println(v5, v4, v4)
    var v68 uint8 = 42
    v69 := GenS5{F0: -6}
    fmt.Println(v1)
    v60 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v66
    _ = v68
    _ = v69
}
