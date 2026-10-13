package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 int; F1 string; F2 uint }
type GenS10 struct { F0 bool; F1 bool; F2 bool }
type GenS11 struct { F0 bool }
func genFn0(p0 uint, p1 bool) uint8 {
    switch (((("foo bar" + ("a" + "")) + (("a" + "hello") + "a")) + (("a" + ("hello" + (("hello" + "hello") + ("\n\t" + "hello")))) + "")) + "a") {
    case "":
        var v6 uint32 = 1
        _ = v6
    default:
        var v7 int8 = 11
        _ = v7
    }
    switch 2147483647 {
    case 1:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{}
            fmt.Println(fi8, fi8)
            _ = fi8
            _ = v9
        }
    case 9223372036854775807:
        var v10 int = -9223372036854775808
        _ = v10
    case 9:
        var v11 int64 = 9223372036854775807
        _ = v11
    default:
        v12 := []int{32768, -9223372036854775808, 0, -9223372036854775808}
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = ("hello" + "\n\t")
            p1 = ((p1 && p1) || (!(((!(!(((((1e10 - 1e10) * float32(float64(0.0))) * 0.1) + ((((-2.25 / 3.14159) + (3.14159 / 1.5)) + 3.14159) - ((((float32(((((-2.25 + -2.25) * float64(3.14159)) + 0.1) * 1.5)) + 3.14159) - (1.5 - (1e10 / (3.14159 / -2.25)))) + (-2.25 - 0.0)) * (0.0 + (0.1 / (0.1 + 3.14159)))))) < (-2.25 - float32(1e10))))) || true) || p1)))
            fmt.Println(p0, v15)
            p1 = true
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 0, F1: false, F2: 2147483647}
        if p1 {
            v17 := GenS2{F0: 42}
            _ = v17
        }
        _ = fi13
        _ = v16
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []string{"", (((("\n\t" + "") + "foo bar") + "a") + ((("a" + "") + "") + ("a" + ("" + "")))), "", "a"}
    v2 := []int{0, 2147483647, 1}
    var v3 uint = 980
    v4 := make(map[uint32]string)
    var v5 uint16 = 42
    var v18 uint8 = 68
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 float64 = (0.0 * (1.5 * 1.5))
        var v21 uint8 = (v18 >> (((v3 >> (((v3 % ((2) | 1))) & 63))) & 63))
        for fi22 := 0; fi22 < 11; fi22++ {
            switch fi19 {
            case -9223372036854775808:
                delete(v4, 2)
            case 397:
                v23 := GenS1{F0: -973, F1: true, F2: -832}
                _ = v23
            default:
                v18 = (^(v21))
            }
            var v24 uint8 = genFn0((^(v3)), ((-(fi22)) <= fi19))
            var v25 int8 = -74
            _ = fi22
            _ = v24
            _ = v25
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!true) {
            var v27 float32 = 1e10
            _ = v27
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(fi28)
        v29 := GenS2{F0: 149}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v3, v3, v5)
        _ = fi30
    }
    var v31 uint8 = (v18 >> ((v3) & 63))
    fmt.Println(v31, v31)
    switch (("foo bar" + "") + "\n\t") {
    case "":
        var v32 int32 = -594
        _ = v32
    case "hello":
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS3{F0: 2147483647, F1: 0, F2: -1}
            v3 = (v3 << ((v3) & 63))
            var v35 int32 = 0
            _ = fi33
            _ = v34
            _ = v35
        }
    default:
        if ((v18 << (((v3 / ((v3) | 1))) & 63)) <= (^(v18))) {
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := GenS4{F0: -1, F1: -9223372036854775808}
                _ = fi36
                _ = v37
            }
            if (false && false) {
                v38 := GenS3{F0: 42, F1: 2, F2: -1}
                var v39 uint16 = 32767
                switch -2147483648 {
                case 127:
                    var v40 uint = (v3 << (((^(v3))) & 63))
                    _ = v40
                default:
                    _ = len(v2)
                }
                _ = v38
                _ = v39
            } else {
                v5 = v5
                v41 := v4[2147483648]
                _ = v41
            }
            v5 += uint16(v3)
        } else {
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 uint = (v3 ^ (v3 % (((v3 & v3)) | 1)))
                for fi44 := 0; fi44 < 9; fi44++ {
                    delete(v4, 42)
                    v18 -= 1
                    if (-32768 > -529) {
                        v45 := GenS4{F0: 0, F1: 42}
                        var v46 int = fi44
                        var v47 int16 = 42
                        if ("a" <= "a") {
                            var v48 int64 = 9223372036854775807
                            _ = v48
                        } else {
                            v49 := GenS5{F0: fi44, F1: 100}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                _ = len(v4)
                                v47 += (v47 % (((v47 << (((v3 - uint(v18))) & 63))) | 1))
                                v50 := GenS6{F0: v31, F1: 208, F2: 0}
                                _ = v50
                            }
                            _ = v49
                        }
                        if (255 < v43) {
                            v31 -= (^(v31))
                            v5 = (v5 & v5)
                        } else {
                            v18 += genFn0(255, false)
                        }
                        switch (fi42 * fi42) {
                        case -2147483648:
                            if true {
                                v47 += 42
                                var v51 uint64 = 1
                                v18 -= uint8(v5)
                                _ = v51
                            }
                        case 32768:
                            var v52 string = "a"
                            _ = v52
                        case 4294967295:
                            v53 := GenS7{F0: (fi42 & 1)}
                            _ = v53
                        }
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    switch (fi42 >> ((872) & 63)) {
                    case -58:
                        if false {
                            fmt.Println(fi42)
                        } else {
                            var v54 uint64 = 1
                            v4[uint32(v5)] = "a"
                            v5 = 1
                            _ = v54
                        }
                    case -510:
                        var v55 string = "a"
                        _ = v55
                    default:
                        var v56 float32 = 3.14159
                        _ = v56
                    }
                    var v57 uint64 = 5
                    _ = fi44
                    _ = v57
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi42
                _ = v43
            }
            v58 := GenS8{F0: 0, F1: 2147483648}
            _ = v58
        }
    }
    switch "" {
    case "foo bar":
        if false {
            if (42 < 906) {
                if (!(int(v3) >= 140)) {
                    v5 += 2
                } else {
                    v18 -= genFn0((^(v3)), ((((0.1 + ((0.1 + -2.25) / 0.1)) - 3.14159) - 1.5) >= ((((1.5 / 3.14159) + 3.14159) + (float32(float64(0.1)) * 0.1)) * 1e10)))
                    var v59 float32 = ((float32((((-2.25 / 1.5) * 0.1) / 1e10)) - (float32(0.1) - 1.5)) - -2.25)
                    switch "\n\t" {
                    case "\n\t":
                        var v60 uint16 = v5
                        _ = v60
                    default:
                        if (18446744073709551615 != 319) {
                            var v61 string = (("\n\t" + "") + "a")
                            v3 -= (v3 << ((v3) & 63))
                            v62 := GenS7{F0: 0}
                            _ = v61
                            _ = v62
                        } else {
                            v4[67] = "hello"
                        }
                    }
                    _ = v59
                }
                _ = len(v2)
                if (42 <= 1) {
                    for fi63 := 0; fi63 < 14; fi63++ {
                        fmt.Println(v31, v31)
                        v64 := GenS9{F0: 463, F1: ("hello" + ("" + ("foo bar" + "\n\t"))), F2: v3}
                        _ = fi63
                        _ = v64
                    }
                    v65 := GenS7{F0: -169}
                    _ = v65
                } else {
                    var v66 int = -9223372036854775808
                    fmt.Println(v5)
                    v67 := GenS10{F0: false, F1: (-1 < -32768), F2: ("foo bar" <= ("foo bar" + ("\n\t" + "")))}
                    _ = v66
                    _ = v67
                }
            } else {
                v1 = append(v1, "")
                switch ("hello" + ("a" + "\n\t")) {
                case "a":
                    fmt.Println(v5, v5, v18)
                case "\n\t":
                    v18 = (v31 + v31)
                }
                v4[4294967295] = "hello"
            }
        }
    case "hello":
        v68 := GenS11{F0: false}
        _ = v68
    case "":
        v18 = 255
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
