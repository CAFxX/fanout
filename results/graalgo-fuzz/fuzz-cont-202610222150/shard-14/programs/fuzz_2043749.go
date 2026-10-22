package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int16, p1 uint16, p2 int16) uint32 {
    var v7 int64 = 0
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 736
}
func genFn1(p0 int8, p1 uint16) int16 {
    v8 := GenS1{F0: true, F1: "hello"}
    p1 -= (p1 | (p1 * p1))
    _ = p0
    _ = p1
    _ = v8
    return -32768
}

func main() {
    var v1 uint64 = 1
    v2 := []int{-2147483648, 127, -2147483648, 1}
    v3 := make(map[int]string)
    var v4 int8 = 10
    v5 := make(map[string]uint32)
    var v6 uint64 = (v1 ^ (v1 >> ((255) & 63)))
    v1 = v1
    v9 := GenS1{F0: (!(0 != 1)), F1: "a"}
    var v10 uint8 = 166
    v10 -= 2
    fmt.Println(v6)
    v10 = (v10 + (v10 ^ (v10 % ((v10) | 1))))
    if false {
        var v11 int16 = -207
        v10 = (v10 >> ((889) & 63))
        _ = v11
    }
    var v12 uint16 = 843
    v4 += (^(v4))
    switch 751 {
    case 32768:
        v12 = (v12 + (v12 >> ((1) & 63)))
    case 42:
        for fi13 := 0; fi13 < 9; fi13++ {
            if true {
                v12 += (v12 % (((v12 >> ((2) & 63))) | 1))
                if true {
                    if (-498 > int32(v4)) {
                        fmt.Println(v10, v10)
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v15 := GenS1{F0: (1 < 42), F1: ("" + "hello")}
                            _ = fi14
                            _ = v15
                        }
                        v6 += (v6 + (^(v6)))
                    }
                    var v16 int32 = 0
                    for fi17 := 0; fi17 < 4; fi17++ {
                        fmt.Println(fi13, v16, fi13)
                        switch "" {
                        case "a":
                            var v18 bool = (true && (v4 >= (v4 / ((v4) | 1))))
                            _ = v18
                        default:
                            if false {
                                v5[((("a" + "hello") + "hello") + "foo bar")] = 403
                            }
                        }
                        _ = fi17
                    }
                    _ = v16
                }
            } else {
                switch "" {
                case "foo bar":
                    if (!(!((v1 ^ (v1 / (((v1 + (v1 >> ((744) & 63)))) | 1))) == (v1 << ((0) & 63))))) {
                        v4 = (v4 << ((4294967295) & 63))
                        if ((^(v12)) != v12) {
                            if len(v2) > 0 {
                                v2[0] = fi13
                            }
                            var v19 int8 = (v4 | (v4 / ((v4) | 1)))
                            _ = v19
                        } else {
                            var v20 uint16 = v12
                            _ = v20
                        }
                        v4 = v4
                    }
                default:
                    v21 := GenS2{F0: (v4 + v4), F1: (v12 & v12), F2: "hello"}
                    _ = v21
                }
                if false {
                    v22 := GenS1{F0: ((true || (0 == 42)) && false), F1: "\n\t"}
                    _ = v22
                }
                _ = len(v2)
            }
            switch ("a" + ("foo bar" + "a")) {
            case "":
                var v23 float32 = 3.14159
                _ = v23
            case "\n\t":
                v12 += v12
            case "foo bar":
                v10 = (v10 ^ v10)
            }
            _ = fi13
        }
    case 0:
        v24 := GenS2{F0: 1, F1: v12, F2: ""}
        _ = v24
    default:
        if (((1.5 >= (((float64(0.0) + (3.14159 / 1.0)) - 3.14159) - float64(float32((0.1 / ((0.1 - 1e10) * (3.14159 + 0.1))))))) || (1 == -2147483648)) || (uint8(v6) <= 119)) {
            if false {
                var v25 string = (("a" + "foo bar") + "")
                switch v25 {
                case "a":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v5[(v25 + ((v25 + v25) + (v25 + (v25 + v25))))] = 1
                        v6 = v6
                        var v27 int8 = v4
                        _ = fi26
                        _ = v27
                    }
                default:
                    switch 4294967295 {
                    case 303:
                        v25 = v25
                    }
                }
                _ = v25
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                v6 = (v6 - v1)
                fmt.Println(v10, v12)
                _ = fi28
            }
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                for fi30 := 0; fi30 < 13; fi30++ {
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                    _ = fi30
                }
                var v31 int32 = -2147483648
                v12 = 255
                _ = fi29
                _ = v31
            }
            v32 := v3[-249]
            v32 = "a"
            _ = v32
        }
    }
    if (!true) {
        fmt.Println(v1, v10, v12)
        v1 = v1
    } else {
        fmt.Println(v10, v12, v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
