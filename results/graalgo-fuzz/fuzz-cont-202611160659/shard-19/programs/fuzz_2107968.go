package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 string }
func genFn0() uint64 {
    var v7 float64 = 0.0
    v8 := GenS1{F0: 4294967295}
    v7 = v7
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1() uint8 {
    var v9 uint = 4294967295
    v9 -= v9
    _ = v9
    return 109
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{}
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    var v6 int8 = 0
    if (!(!(true || (128 >= 1)))) {
        for fi10 := 0; fi10 < 5; fi10++ {
            delete(v5, 42)
            _ = fi10
        }
    } else {
        var v11 bool = ((v6 >> ((42) & 63)) != (v6 | 0))
        switch "\n\t" {
        case "foo bar":
            switch 127 {
            case -1:
                switch (("foo bar" + "a") + "hello") {
                case "\n\t":
                    v2 = append(v2, 7)
                case "hello":
                    if v11 {
                        var v12 int = -97
                        v6 = v6
                        _ = v12
                    }
                case "foo bar":
                    fmt.Println(v11, v1)
                default:
                    fmt.Println(v1)
                }
            case -2147483648:
                if (543 > 127) {
                    v11 = ((v6 & v6) == (v6 / ((v6) | 1)))
                }
            default:
                for fi13 := 0; fi13 < 2; fi13++ {
                    v11 = false
                    v14 := GenS1{F0: 2147483647}
                    v15 := GenS2{F0: 0, F1: -1}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        case "a":
            v11 = v11
        case "\n\t":
            var v16 int8 = (-(v6))
            _ = v16
        }
        _ = v11
    }
    v6 = 74
    for fi17 := 0; fi17 < 17; fi17++ {
        if (((!true) && false) && true) {
            switch fi17 {
            case 21:
                var v18 uint = 18446744073709551615
                _ = v18
            case 2147483647:
                v1 += (0.0 - (1.5 / v1))
            case -240:
                _ = len(v3)
            }
            if len(v3) > 0 {
                v3[0] = (("foo bar" + ("" + ("a" + ("" + ("foo bar" + ("\n\t" + (("a" + "") + ("foo bar" + ("a" + "a"))))))))) + "hello")
            }
            switch "hello" {
            case "hello":
                _ = len(v3)
            case "foo bar":
                v1 = v1
            default:
                v3 = append(v3, "a")
            }
        } else {
            v19 := GenS1{F0: 65535}
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            v6 += (^(v6))
            _ = fi20
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if true {
            for fi22 := 0; fi22 < 9; fi22++ {
                v5[42] = (!true)
                _ = fi22
            }
            var v23 uint8 = 2
            _ = v23
        }
        v2 = append(v2, (fi21 + (fi21 % (((fi21 - (fi21 | fi21))) | 1))))
        _ = fi21
    }
    var v24 float32 = (3.14159 - 0.0)
    if (!true) {
        var v25 uint16 = 42
        v26 := GenS3{F0: 1, F1: 1, F2: 40}
        var v27 uint8 = 127
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS4{F0: 0}
    v1 -= (v1 + 3.14159)
    if (true && (((true && true) || (42 <= 0)) || true)) {
        if false {
            if true {
                for fi29 := 0; fi29 < 1; fi29++ {
                    v24 = (float32((float64(v24) - float64(((1.5 + v24) / (v24 / v24))))) + v24)
                    _ = fi29
                }
                fmt.Println(v1)
                for fi30 := 0; fi30 < 6; fi30++ {
                    switch 42 {
                    case 65:
                        v6 += (v6 >> ((1) & 63))
                    case 922:
                        fmt.Println(v6, v24)
                    default:
                        v24 -= 0.0
                    }
                    v24 = (((0.1 - ((v24 + (0.1 - 0.1)) - ((float32(v1) - v24) - (0.0 / 1.0)))) / v24) / v24)
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v24 -= ((v24 - (v24 - 0.0)) * v24)
                        v6 = 100
                        v1 = ((-2.25 + ((1.5 * -2.25) - v1)) + (float64(v24) + v1))
                        _ = fi31
                    }
                    _ = fi30
                }
            } else {
                var v32 uint32 = 1
                var v33 float32 = 1e10
                _ = v32
                _ = v33
            }
        }
        switch -209 {
        case 613:
            v24 = v24
        case 585:
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 bool = (true && (!(float64(float32(v1)) >= 3.14159)))
                _ = fi34
                _ = v35
            }
        case -174:
            for fi36 := 0; fi36 < 6; fi36++ {
                if false {
                    for fi37 := 0; fi37 < 19; fi37++ {
                        v24 -= float32(float64(v24))
                        _ = fi37
                    }
                    var v38 string = "foo bar"
                    _ = v38
                } else {
                    v1 = (((1e10 * 0.1) * v1) * v1)
                }
                fmt.Println(fi36)
                switch "foo bar" {
                case "":
                    v1 = 1.5
                case "\n\t":
                    var v39 string = ("" + "hello")
                    _ = v39
                case "a":
                    v40 := GenS4{F0: -32768}
                    _ = v40
                }
                _ = fi36
            }
        default:
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
        }
    } else {
        if true {
            v41 := GenS2{F0: 1, F1: -800}
            _ = v41
        } else {
            var v42 int8 = v6
            v43 := GenS4{F0: 869}
            v42 = v42
            _ = v42
            _ = v43
        }
        if (4294967295 != 32768) {
            switch -9223372036854775808 {
            case 903:
                v44 := GenS5{F0: (v6 << ((0) & 63)), F1: false, F2: ("foo bar" + (("\n\t" + "") + ("hello" + "a")))}
                _ = v44
            case -2147483648:
                v24 = (-2.25 * v24)
            case 432:
                v6 += v6
            default:
                var v45 string = ("hello" + ("foo bar" + "foo bar"))
                _ = v45
            }
        }
        var v46 uint = 25
        _ = v46
    }
    var v47 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
