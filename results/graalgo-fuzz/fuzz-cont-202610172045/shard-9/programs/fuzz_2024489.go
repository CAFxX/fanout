package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    v1 := []int{492, -2147483648}
    v2 := make(map[uint32]bool)
    var v3 int32 = -885
    var v4 string = "foo bar"
    var v5 float32 = 1.5
    v6 := GenS1{F0: false, F1: 1, F2: 211}
    v2[2] = ((true && true) || true)
    var v7 uint16 = 255
    switch -2147483648 {
    case 0:
        if false {
            delete(v2, 1)
        }
    case -583:
        var v8 uint16 = v7
        _ = v8
    default:
        var v9 bool = (float64(((v5 * 0.0) / (1.5 - ((v5 + v5) - float32(1e10))))) != (0.1 + 1e10))
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        switch ("a" + "a") {
        case "\n\t":
            var v11 uint = 255
            _ = v11
        default:
            v12 := GenS1{F0: true, F1: -128, F2: v7}
            _ = v12
        }
        if (0.1 >= 0.0) {
            switch fi10 {
            case -1:
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: false, F1: 42, F2: v7}
                    switch (v4 + v4) {
                    case "":
                        v7 += v7
                    default:
                        v7 -= (^(v7))
                        v7 = v7
                        v2[42] = (v4 < "hello")
                        v1 = append(v1, 32768)
                    }
                    _ = fi13
                    _ = v14
                }
            case 2147483647:
                fmt.Println(v5, fi10, fi10)
            case -2147483648:
                if ("\n\t" == v4) {
                    var v15 uint = 255
                    _ = v15
                } else {
                    switch v4 {
                    case "":
                        v5 -= ((((1e10 / v5) / ((v5 - v5) + ((v5 + v5) * ((0.1 * ((v5 + (float32(3.14159) + v5)) / v5)) + (0.0 + 0.1))))) / v5) * (1.5 * (float32(0.1) - (v5 / v5))))
                    case "\n\t":
                        var v16 uint8 = 255
                        _ = v16
                    }
                }
            }
        } else {
            switch fi10 {
            case -1:
                if (-945 < int64(v7)) {
                    if false {
                        v7 += (^(v7))
                        v17 := GenS1{F0: false, F1: -121, F2: v7}
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if (1 <= -640) {
                                if (18446744073709551615 <= 80) {
                                    switch fi18 {
                                    case 128:
                                        v4 = v4
                                    case -176:
                                        var v19 float64 = 0.0
                                        _ = v19
                                    }
                                }
                            }
                            fmt.Println(fi10, v5, v7)
                            _ = fi18
                        }
                        _ = v17
                    }
                    v20 := GenS1{F0: (true && false), F1: 42, F2: v7}
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v5 = 3.14159
                        _ = fi21
                    }
                    _ = v20
                } else {
                    if (0 <= 247) {
                        fmt.Println(v4, v5, v7)
                    }
                    var v22 bool = false
                    _ = v22
                }
            case 2147483647:
                if len(v1) > 0 {
                    v1[0] = fi10
                }
            case 32768:
                v23 := GenS1{F0: true, F1: 1, F2: (v7 | 42)}
                _ = v23
            }
        }
        v7 -= v7
        _ = fi10
    }
    v1 = append(v1, 48)
    v2[65535] = false
    if (!false) {
        v7 = (^(v7))
    }
    var v24 int32 = 42
    v5 -= 0.1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 -= 981
    v25 := GenS2{F0: false, F1: 2147483647}
    v24 = (v3 / (((v24 + (v24 & -1))) | 1))
    v7 -= (v7 - (v7 | 42))
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    switch 2147483647 {
    case 4294967295:
        v3 += (v24 << ((433) & 63))
    case 175:
        v1 = append(v1, 0)
    case 723:
        fmt.Println(v24)
    default:
        if ((162 >= 127) && false) {
            _ = len(v1)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
