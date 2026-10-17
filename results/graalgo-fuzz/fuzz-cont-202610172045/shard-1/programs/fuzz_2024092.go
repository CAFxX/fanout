package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int16; F1 int; F2 string }
type GenS7 struct { F0 bool; F1 uint16; F2 int64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int8, p1 int) uint16 {
    p1 = (-(p1))
    fmt.Println(p1)
    v5 := []int{p1}
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() bool {
    var v6 int8 = 42
    var v7 float64 = ((1.5 / 3.14159) / 1e10)
    switch 32768 {
    case -195:
        v8 := GenS1{F0: (v6 ^ v6), F1: 65535, F2: 2}
        _ = v8
    case -320:
        var v9 int16 = -32768
        _ = v9
    case -893:
        for fi10 := 0; fi10 < 1; fi10++ {
            if (!true) {
                v11 := GenS1{F0: v6, F1: 911, F2: 42}
                var v12 uint8 = 42
                v7 += v7
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 1, F1: 1, F2: 715}
            if false {
                for fi14 := 0; fi14 < 11; fi14++ {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            v15 := GenS2{F0: 4294967295}
                            var v16 string = (("" + "a") + "")
                            _ = v15
                            _ = v16
                        }
                    case "foo bar":
                        var v17 int64 = 1
                        _ = v17
                    default:
                        v7 = (3.14159 + ((v7 - v7) - 0.0))
                    }
                    _ = fi14
                }
                v18 := GenS2{F0: 708}
                v7 = ((v7 * float64(((0.1 * 1e10) * (3.14159 / 1.5)))) + v7)
                _ = v18
            } else {
                v6 = v6
            }
            _ = fi10
            _ = v13
        }
    default:
        v7 += v7
    }
    if (-451 >= 42) {
        var v19 uint16 = 47
        _ = v19
    } else {
        v7 = 1.5
    }
    v20 := []int{2147483647, -878, 33, 1}
    _ = v20
    _ = v6
    _ = v7
    return (v6 > (v6 & 0))
}

func main() {
    var v1 bool = ("" <= "")
    var v2 bool = (-2.25 <= float64(0.1))
    v3 := []uint32{}
    var v4 uint16 = 32767
    v1 = (168 > 1)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 int16 = 202
        var v23 uint32 = 420
        v2 = (639 > 42)
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    if v2 {
        switch "" {
        case "a":
            for fi25 := 0; fi25 < 11; fi25++ {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v2 = ((!(!v24)) && v1)
                    fmt.Println(v4, v24)
                    var v27 int64 = -1
                    _ = len(v3)
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        default:
            if v24 {
                for fi28 := 0; fi28 < 10; fi28++ {
                    v4 = (v4 ^ 507)
                    var v29 float32 = 1.5
                    _ = fi28
                    _ = v29
                }
            } else {
                for fi30 := 0; fi30 < 9; fi30++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v31 int = -2147483648
                    _ = fi30
                    _ = v31
                }
                v2 = (!(!(!(!(255 < 536)))))
            }
        }
    } else {
        if v24 {
            if false {
                v32 := GenS3{F0: ((18446744073709551615 < 18446744073709551615) && v2)}
                if (!v2) {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        switch fi33 {
                        case 4294967295:
                            _ = len(v3)
                        default:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                switch fi33 {
                                case -2147483648:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 283:
                                    v35 := GenS2{F0: 709}
                                    _ = v35
                                case 539:
                                    v36 := GenS4{F0: v24, F1: 42}
                                    _ = v36
                                default:
                                    v4 += (v4 / ((genFn0(1, fi33)) | 1))
                                }
                                var v37 uint16 = v4
                                _ = fi34
                                _ = v37
                            }
                        }
                        _ = fi33
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = v32
            } else {
                fmt.Println(v2, v24, v1)
                if v24 {
                    v38 := GenS5{F0: 2147483647, F1: 348}
                    for fi39 := 0; fi39 < 15; fi39++ {
                        v40 := GenS6{F0: 1, F1: (fi39 << ((18446744073709551615) & 63)), F2: "foo bar"}
                        var v41 uint8 = 186
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    var v42 float64 = -2.25
                    _ = v38
                    _ = v42
                }
            }
            v4 -= 476
        }
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        fmt.Println(v24, v24, v24)
        for fi44 := 0; fi44 < 14; fi44++ {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v45 := GenS1{F0: -74, F1: 663, F2: 42}
                _ = v45
            } else {
                if v2 {
                    if v2 {
                        v46 := GenS4{F0: v2, F1: 1}
                        _ = v46
                    }
                    if (!(!(!(v4 >= (v4 - v4))))) {
                        if v2 {
                            v47 := GenS7{F0: v24, F1: v4, F2: -1}
                            v48 := GenS2{F0: 0}
                            var v49 bool = (v24 && (872 != 18446744073709551615))
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            v4 = (^(v4))
                            var v50 int8 = 1
                            v51 := GenS8{F0: "a", F1: -9223372036854775808}
                            _ = v50
                            _ = v51
                        }
                    } else {
                        switch "a" {
                        case "\n\t":
                            v4 = (v4 >> ((819) & 63))
                        }
                        if true {
                            v24 = v24
                            switch fi44 {
                            case -1:
                                var v52 string = "foo bar"
                                _ = v52
                            }
                            switch "hello" {
                            case "":
                                v1 = (42 != 32767)
                            case "a":
                                v2 = v1
                            default:
                                v1 = v24
                            }
                        }
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 += (v4 & v4)
                    switch "hello" {
                    case "hello":
                        switch "foo bar" {
                        case "foo bar":
                            v53 := GenS9{F0: ""}
                            _ = v53
                        }
                    default:
                        v4 -= v4
                    }
                }
                v4 += 65535
            }
            _ = fi44
        }
        _ = fi43
    }
    switch "\n\t" {
    case "hello":
        var v54 float32 = 0.0
        _ = v54
    }
    fmt.Println(v24, v1, v24)
    v55 := GenS6{F0: 989, F1: 32767, F2: ""}
    if (-1 > 42) {
        if (0.1 == 0.0) {
            v1 = true
            switch "foo bar" {
            case "hello":
                v56 := GenS8{F0: "\n\t", F1: 972}
                _ = v56
            case "a":
                v57 := GenS10{F0: 0}
                _ = v57
            case "\n\t":
                var v58 float32 = (float32((0.1 + (((1.5 * (0.1 * (1.5 * (1.5 - (1.5 - 3.14159))))) + -2.25) / 1.0))) / (float32(((-2.25 / ((1.5 * -2.25) * -2.25)) * ((0.1 - 0.0) + float64(float32(1.5))))) * (0.1 / (0.0 - float32(((1e10 * (((3.14159 * (0.0 - 3.14159)) * -2.25) - (((0.0 * float64(3.14159)) + (float64(((0.0 + -2.25) * 0.1)) + 1.5)) * ((0.1 - 1e10) / (1.5 / 1e10))))) - float64(((1e10 + -2.25) - (3.14159 + 3.14159)))))))))
                _ = v58
            default:
                v59 := GenS8{F0: ("foo bar" + "foo bar"), F1: 0}
                _ = v59
            }
            v4 = 537
        }
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        if true {
            var v61 string = "foo bar"
            if v2 {
                switch v61 {
                case "hello":
                    v1 = (1 > -512)
                case "a":
                    var v62 bool = false
                    _ = v62
                }
                v2 = (131 == 127)
                var v63 int8 = -128
                _ = v63
            }
            _ = v61
        } else {
            var v64 int8 = 1
            v1 = v2
            _ = v64
        }
        _ = fi60
    }
    v65 := GenS3{F0: ((-1 == 835) && (true && v24))}
    v24 = v1
    switch "" {
    case "hello":
        v66 := GenS11{F0: 171}
        _ = v66
    default:
        v24 = (0 > 0)
    }
    v24 = (-128 >= 0)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v67 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v55
    _ = v65
    _ = v67
}
