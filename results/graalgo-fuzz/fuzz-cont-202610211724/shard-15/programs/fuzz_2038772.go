package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8, p1 int8) int {
    var v4 uint16 = 882
    fmt.Println(p0)
    p1 += -91
    v5 := GenS1{F0: 0, F1: 612}
    v6 := GenS2{F0: v4, F1: (!(false && ((1 < 1) && true)))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -87
}

func main() {
    var v1 int8 = 42
    v2 := []int{-327}
    v3 := make(map[uint32]int)
    if true {
        v1 -= (-(v1))
        for fi7 := 0; fi7 < 12; fi7++ {
            v3[1] = 902
            _ = fi7
        }
    } else {
        var v8 float64 = (0.1 / 1.5)
        v1 -= 62
        fmt.Println(v1)
        _ = v8
    }
    var v9 int16 = 1000
    v1 = v1
    var v10 int8 = -128
    v1 += v1
    v1 += (v10 % ((v1) | 1))
    fmt.Println(v1)
    var v11 int16 = v9
    switch genFn0(128, (v10 * (v1 >> ((42) & 63)))) {
    case 0:
        switch 2147483647 {
        case 128:
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
        }
    case 128:
        var v12 uint64 = 4294967295
        _ = v12
    default:
        v11 -= 313
    }
    v1 -= (v10 % (((v1 * (v1 / (((v1 | (v1 - v1))) | 1)))) | 1))
    var v13 int32 = -2147483648
    var v14 uint64 = 200
    v3[2] = 1
    for fi15 := 0; fi15 < 9; fi15++ {
        v10 = (v1 >> ((65535) & 63))
        var v16 uint8 = 162
        _ = fi15
        _ = v16
    }
    fmt.Println(v11, v10)
    v10 = (v10 * (-(v1)))
    for fi17 := 0; fi17 < 5; fi17++ {
        v3[uint32(v14)] = fi17
        if ((!true) && ((!false) || ((v14 >> ((0) & 63)) >= v14))) {
            _ = len(v3)
            switch "a" {
            case "a":
                if (genFn0(2, (^(v1))) <= (fi17 / ((32767) | 1))) {
                    v13 -= (v13 | v13)
                } else {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v10 = v10
                        _ = fi18
                    }
                    switch fi17 {
                    case -987:
                        for fi19 := 0; fi19 < 10; fi19++ {
                            v2 = append(v2, fi17)
                            _ = fi19
                        }
                    case 128:
                        v20 := GenS2{F0: 32767, F1: true}
                        _ = v20
                    }
                }
            case "hello":
                switch "a" {
                case "hello":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        fmt.Println(v10)
                        _ = fi21
                    }
                }
            }
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                v11 = (v11 / ((v11) | 1))
                v9 = v11
                var v23 float32 = (((((1.5 - 1.5) * 1.5) - ((((1e10 * (1e10 + 1e10)) - 0.0) - ((-2.25 * 1.5) / (-2.25 + 0.1))) / 1.0)) - float32(0.1)) - float32(1e10))
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 13; fi24++ {
                switch "" {
                case "foo bar":
                    delete(v3, 42)
                case "\n\t":
                    switch fi24 {
                    case 4294967295:
                        if (((1e10 * -2.25) * ((3.14159 * 0.0) / (0.1 - (1.5 * (((1.5 / 1.5) - 0.1) / 1e10))))) < 1e10) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                case "a":
                    v11 = int16(v10)
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v13 += v13
                v11 += 744
                v9 += v11
                _ = fi24
            }
        }
        _ = fi17
    }
    switch 128 {
    case 9223372036854775807:
        v25 := GenS2{F0: 65535, F1: ("foo bar" > (("hello" + ("" + "\n\t")) + "hello"))}
        _ = v25
    case 32768:
        v13 += v13
    case -9223372036854775808:
        _ = len(v2)
    default:
        v26 := GenS1{F0: -1, F1: (v14 ^ v14)}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
