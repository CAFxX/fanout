package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32; F2 uint8 }
type GenS5 struct { F0 bool; F1 int32; F2 uint32 }
func genFn0(p0 uint64, p1 int) uint16 {
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: uint(v6)}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int16) uint16 {
    v8 := GenS2{F0: (p0 ^ (^(p0)))}
    v9 := GenS3{F0: -53, F1: 61}
    p0 -= 1000
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{(("hello" + "hello") + "hello")}
    var v3 uint32 = 70
    v4 := make(map[string]bool)
    var v5 float32 = 0.1
    v10 := GenS2{F0: -1}
    v5 += (1.5 - (0.0 * (float32(0.1) + (((v5 + -2.25) + (3.14159 / v5)) - v5))))
    v5 = v5
    var v11 int32 = -1
    switch 32767 {
    case -2147483648:
        var v12 uint16 = 32767
        _ = v12
    case 0:
        v5 -= (1e10 * 0.1)
    }
    v3 = (v3 ^ v3)
    if false {
        v3 += v3
        v13 := GenS4{F0: 727, F1: (v11 % ((1) | 1)), F2: 42}
        v3 = v3
        _ = v13
    } else {
        var v14 float64 = (0.0 + 3.14159)
        var v15 uint64 = 42
        var v16 float64 = (v14 + 0.1)
        _ = v14
        _ = v15
        _ = v16
    }
    switch "" {
    case "hello":
        for fi17 := 0; fi17 < 5; fi17++ {
            switch 79 {
            case 0:
                if (2 != 1) {
                    var v18 uint64 = 42
                    _ = v18
                } else {
                    var v19 int32 = (v11 - 936)
                    _ = v19
                }
            case 569:
                v1[fi17] = (v3 + 207)
            case 169:
                v11 = v11
            default:
                v20 := GenS5{F0: (4294967295 > 2), F1: -292, F2: (v3 * 42)}
                _ = v20
            }
            _ = fi17
        }
    case "a":
        var v21 int16 = -696
        _ = v21
    default:
        v3 -= (^(v3))
    }
    v5 -= -2.25
    switch 42 {
    case 0:
        v11 += v11
    default:
        var v22 uint16 = 2
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 string = "a"
        switch 1 {
        case 32768:
            if ((((((v5 / (v5 + (v5 * v5))) / ((v5 + v5) / v5)) + v5) / 0.1) * v5) < ((((-2.25 + ((3.14159 * (1e10 + ((v5 / (((float32((0.0 + 3.14159)) - 0.1) + (float32(float64((v5 - (v5 / float32(0.1))))) - v5)) * v5)) - (((3.14159 * -2.25) - float32(0.0)) * 1e10)))) / v5)) - v5) - float32(((-2.25 + 0.0) + ((-2.25 - 1e10) * (((1.5 / 3.14159) / 3.14159) * (float64(v5) / float64(v5))))))) * v5)) {
                if (!false) {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if true {
                            switch v24 {
                            case "\n\t":
                                var v26 int8 = 1
                                _ = v26
                            default:
                                switch 2147483647 {
                                case 128:
                                    fmt.Println(v11, v24, v24)
                                case 32768:
                                    v5 -= (1.5 / v5)
                                default:
                                    _ = len(v4)
                                    v3 -= 2
                                }
                            }
                            var v27 int = 4294967295
                            _ = v27
                        }
                        _ = fi25
                    }
                    if len(v2) > 0 {
                        v2[0] = v24
                    }
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    v1[fi23] = (v3 + v3)
                    var v29 uint8 = 18
                    var v30 uint = 4294967295
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            } else {
                switch v24 {
                case "":
                    v31 := GenS1{F0: 17}
                    _ = v31
                case "foo bar":
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v32 := GenS2{F0: 0}
                        _ = v32
                    }
                default:
                    var v33 uint32 = v3
                    _ = v33
                }
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v11, v11)
    v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
