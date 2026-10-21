package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint64; F2 int }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 uint = 1
    var v4 int16 = -1
    v1[-427] = 2
    v2 += 42
    if false {
        if (true || true) {
            v2 -= v2
            fmt.Println(v4, v2, v3)
        } else {
            if (v4 < v4) {
                if false {
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        v5 := GenS1{F0: 1, F1: ("\n\t" + (("hello" + "a") + ""))}
                        _ = v5
                    case "":
                        v4 = v4
                    }
                }
            }
        }
        v6 := GenS1{F0: 956, F1: "foo bar"}
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS2{F0: (true && (true && (false || (true || (2 > 2))))), F1: (-(v2)), F2: 1}
            v9 := GenS3{F0: ((fi7 >> ((v3) & 63)) <= -2147483648), F1: 4294967295, F2: -730}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        v10 := GenS2{F0: (!false), F1: -912, F2: 218}
        _ = v10
    }
    if ((v4 % ((70) | 1)) < v4) {
        var v11 bool = (float32(((1e10 * ((0.0 / 1.5) - -2.25)) / float64((1.5 / 1.0)))) != (1.5 / (((3.14159 - (0.1 * (0.1 / 3.14159))) + 0.1) * 1e10)))
        var v12 uint16 = 42
        v4 = (^(v4))
        _ = v11
        _ = v12
    } else {
        switch -2147483648 {
        case 32767:
            if ((1 == 712) && ((!(true && (false || (true && (!false))))) || (false && true))) {
                v2 -= (v2 - int32(v4))
                var v13 int32 = v2
                _ = v13
            } else {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v2)
                case "foo bar":
                    v14 := []int{int(v2)}
                    _ = v14
                }
                if false {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 uint64 = 42
                        _ = fi15
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 int64 = -9223372036854775808
                        var v19 int32 = v2
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v18 += (v18 ^ -9223372036854775808)
                            var v21 string = "a"
                            _ = len(v1)
                            _ = fi20
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if true {
                        delete(v1, 42)
                        var v22 bool = true
                        v3 = 42
                        _ = v22
                    }
                    v3 -= v3
                }
            }
        default:
            var v23 float32 = float32(-2.25)
            _ = v23
        }
    }
    v4 = 32767
    var v24 uint = (v3 | 0)
    switch 708 {
    case -560:
        v2 -= 1
    case -9223372036854775808:
        if (((false || false) || true) && (42 <= 255)) {
            var v25 int16 = v4
            v26 := []int{128, 9223372036854775807, 138}
            _ = v25
            _ = v26
        }
    case -950:
        v4 = (v4 * (v4 << ((v3) & 63)))
    }
    var v27 int64 = 1
    v28 := GenS4{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
