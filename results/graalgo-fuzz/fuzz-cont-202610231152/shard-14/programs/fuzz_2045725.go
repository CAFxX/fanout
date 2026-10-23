package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 int16; F1 string; F2 int16 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint }
type GenS7 struct { F0 int32; F1 int16 }

func main() {
    var v1 float32 = float32((0.1 * 0.1))
    v2 := make(map[int]bool)
    var v3 int16 = -591
    var v4 int16 = v3
    var v5 float32 = (((float32(((-2.25 + ((0.0 - (-2.25 - 3.14159)) / (1.5 - float64((-2.25 + v1))))) * 0.0)) + (((v1 - 0.1) / 1.0) / -2.25)) - v1) * 0.0)
    var v6 float32 = float32(float64((v5 / 1e10)))
    if (42 == 0) {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
    }
    v5 += v5
    v9 := GenS1{F0: 2147483648}
    var v10 float32 = 3.14159
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 65535}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if (32767 <= 912) {
                v6 = 0.0
                v15 := []int{(fi13 ^ fi14)}
                _ = v15
            }
            var v16 uint64 = 356
            v16 = (v16 | v16)
            _ = fi14
            _ = v16
        }
        _ = fi13
    }
    switch ("foo bar" + "hello") {
    case "":
        fmt.Println(v1, v6, v4)
    default:
        v17 := []int{128, 32767, 32767}
        _ = v17
    }
    v2[9223372036854775807] = false
    v10 = (v5 + (((1e10 + (v1 - v10)) - (float32(0.1) / 3.14159)) / float32(float64((0.1 * 3.14159)))))
    v1 -= (v6 / 1.5)
    if (-9223372036854775808 > -1) {
        var v18 float32 = (0.0 + (1.5 * ((1e10 - 0.0) - 1.5)))
        var v19 uint16 = 255
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int64 = -1
            var v22 uint16 = (v19 & 637)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v18
        _ = v19
    }
    v23 := GenS2{F0: 776, F1: 255, F2: 0}
    switch -2147483648 {
    case 2147483647:
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := []int{(-(fi24)), fi24, (fi24 % (((fi24 + (fi24 / ((127) | 1)))) | 1))}
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 float64 = float64(v1)
                fmt.Println(v4, v1)
                v28 := GenS3{F0: 0, F1: 42, F2: 2}
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 int = fi24
            _ = fi24
            _ = v25
            _ = v29
        }
    case 1:
        if false {
            if false {
                v30 := GenS4{F0: (18446744073709551615 > 0), F1: ((v1 / v10) <= v10), F2: -1}
                switch 32768 {
                case 869:
                    v3 -= (v4 | (v4 << ((4294967295) & 63)))
                case 242:
                    var v31 float32 = 1e10
                    _ = v31
                case 127:
                    fmt.Println(v6)
                default:
                    v32 := GenS4{F0: (!(float64((1.5 - 1.5)) > float64(v6))), F1: true, F2: 42}
                    _ = v32
                }
                v1 = v10
                _ = v30
            } else {
                v33 := []int{32768}
                _ = v33
            }
            switch ("\n\t" + ((("a" + ("foo bar" + "")) + ((("" + "") + (((("a" + ((("" + ("foo bar" + "foo bar")) + "\n\t") + "a")) + (("a" + (((((("" + ("" + "\n\t")) + "a") + ("a" + "foo bar")) + ("a" + ("foo bar" + "hello"))) + "") + "")) + "a")) + "hello") + "hello")) + "\n\t")) + "a")) {
            case "a":
                if true {
                    v34 := GenS5{F0: (v3 + v3), F1: "", F2: v4}
                    var v35 uint8 = 255
                    _ = v34
                    _ = v35
                } else {
                    var v36 float32 = 1e10
                    _ = v36
                }
            default:
                for fi37 := 0; fi37 < 10; fi37++ {
                    switch (("a" + "foo bar") + ("" + ("foo bar" + "\n\t"))) {
                    case "\n\t":
                        if true {
                            var v38 int = (fi37 + int(v3))
                            _ = v38
                        } else {
                            fmt.Println(v6)
                            var v39 int32 = 0
                            _ = v39
                        }
                    case "a":
                        v10 = v10
                    }
                    _ = fi37
                }
            }
            v40 := GenS6{F0: 0, F1: -1, F2: 255}
            _ = v40
        }
    case 9223372036854775807:
        if true {
            if true {
                switch "a" {
                case "hello":
                    v6 = (float32((-2.25 / 1e10)) * v1)
                case "\n\t":
                    if ("" < ("a" + "")) {
                        v41 := GenS3{F0: 22, F1: 43, F2: 650}
                        _ = v41
                    } else {
                        fmt.Println(v6, v1)
                    }
                case "foo bar":
                    fmt.Println(v6)
                default:
                    v42 := GenS4{F0: (42 < 59), F1: false, F2: -1}
                    _ = v42
                }
            } else {
                var v43 uint = 1
                if (!false) {
                    for fi44 := 0; fi44 < 17; fi44++ {
                        v45 := []int{(fi44 ^ fi44)}
                        v46 := GenS7{F0: 42, F1: int16(fi44)}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                }
                var v47 uint32 = 0
                _ = v43
                _ = v47
            }
            if ((-942 <= -9223372036854775808) && true) {
                v48 := GenS3{F0: 4294967295, F1: -1, F2: 1}
                v2[9223372036854775807] = (!(false || (v3 >= 401)))
                _ = v48
            }
            v10 += v10
        } else {
            v2[1] = (785 < 255)
        }
    }
    if true {
        v49 := []int{32768}
        _ = v49
    } else {
        for fi50 := 0; fi50 < 15; fi50++ {
            v1 -= ((v6 / 3.14159) * v5)
            _ = fi50
        }
    }
    v10 = (v6 * v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
