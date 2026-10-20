package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
func genFn0(p0 uint8, p1 uint32) int32 {
    p0 = p0
    var v4 uint = 739
    _ = p0
    _ = p1
    _ = v4
    return -645
}
func genFn1(p0 int8) uint16 {
    if ((255 <= 744) && (!true)) {
        v5 := GenS1{F0: -2147483648}
        _ = v5
    }
    if (0 <= -530) {
        p0 += p0
    }
    v6 := GenS1{F0: 32768}
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := []uint32{}
    if len(v3) > 0 {
        v3[0] = 1
    }
    v7 := GenS2{F0: (v1 - v1), F1: 2, F2: (("foo bar" + ("a" + ("a" + "foo bar"))) + "a")}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 uint = 487
    var v9 float64 = (0.0 / 1e10)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 = (v8 - v8)
    fmt.Println(v1, v8, v8)
    v9 -= (v9 + v9)
    if false {
        if true {
            for fi10 := 0; fi10 < 15; fi10++ {
                switch (fi10 / ((fi10) | 1)) {
                case 127:
                    for fi11 := 0; fi11 < 16; fi11++ {
                        switch -2147483648 {
                        case -127:
                            fmt.Println(v9, fi11, fi10)
                        default:
                            v8 -= (v8 / ((v8) | 1))
                        }
                        _ = fi11
                    }
                case -627:
                    v9 = v9
                default:
                    v8 = v8
                }
                fmt.Println(v1)
                _ = fi10
            }
            fmt.Println(v9, v1)
        } else {
            var v12 string = "foo bar"
            _ = v12
        }
        _ = len(v2)
        fmt.Println(v9, v1)
    } else {
        switch "hello" {
        case "hello":
            v9 = 0.0
        default:
            switch "\n\t" {
            case "foo bar":
                var v13 int32 = -921
                _ = v13
            default:
                if (((0.1 * (3.14159 * -2.25)) + (float32((v9 + v9)) - 0.1)) < 1.5) {
                    var v14 string = ("" + (("" + ("\n\t" + "foo bar")) + (("foo bar" + "\n\t") + "")))
                    if ((906 > 18446744073709551615) && false) {
                        if true {
                            v15 := GenS1{F0: 32767}
                            var v16 uint8 = 42
                            for fi17 := 0; fi17 < 8; fi17++ {
                                _ = len(v3)
                                _ = fi17
                            }
                            _ = v15
                            _ = v16
                        } else {
                            v3 = append(v3, 2147483648)
                        }
                    } else {
                        v2[32767] = (v9 <= v9)
                        fmt.Println(v1)
                    }
                    _ = v14
                } else {
                    switch int(v1) {
                    case -2147483648:
                        _ = len(v3)
                    case 32768:
                        v1 = v1
                    default:
                        v1 = (v1 >> ((v8) & 63))
                    }
                }
            }
        }
        var v18 string = (("hello" + ("" + ("foo bar" + (("a" + ("" + ("a" + ("" + ("\n\t" + "\n\t"))))) + "a")))) + (("" + "foo bar") + "a"))
        _ = v18
    }
    var v19 int8 = 122
    for fi20 := 0; fi20 < 12; fi20++ {
        v2[-9223372036854775808] = false
        var v21 float64 = float64(3.14159)
        var v22 int64 = -209
        _ = fi20
        _ = v21
        _ = v22
    }
    v19 += v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
