package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 bool = (!(686 > 343))
    var v2 uint32 = 65535
    var v3 int64 = 0
    var v4 float64 = 1e10
    var v5 uint32 = 821
    var v6 float32 = ((1.5 + -2.25) - 0.1)
    switch "hello" {
    case "":
        if ((v6 / v6) >= ((v6 - ((0.1 - (((0.0 - 0.0) - 0.1) * v6)) * v6)) + float32((float64(float32(v4)) + ((-2.25 / 1e10) + (((((v4 - v4) / ((v4 * v4) + 1.5)) / (v4 + v4)) / v4) - (-2.25 + 1.5))))))) {
            v7 := []int{}
            switch "hello" {
            case "":
                if len(v7) > 0 {
                    v7[0] = -2147483648
                }
            case "\n\t":
                v1 = v1
            case "hello":
                switch 670 {
                case 128:
                    v7 = append(v7, 319)
                }
            default:
                for fi8 := 0; fi8 < 10; fi8++ {
                    switch ("\n\t" + "hello") {
                    case "a":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        v9 := GenS1{F0: (v5 << ((901) & 63)), F1: ((fi8 ^ (fi8 | fi8)) <= fi8), F2: 113}
                        _ = v9
                    }
                    _ = fi8
                }
            }
            _ = v7
        } else {
            v6 = (float32((3.14159 / 1.5)) / v6)
            if (v4 <= v4) {
                for fi10 := 0; fi10 < 1; fi10++ {
                    if (!true) {
                        for fi11 := 0; fi11 < 4; fi11++ {
                            if v1 {
                                v12 := []int{(fi10 - 4294967295), (fi10 * fi11)}
                                fmt.Println(v1)
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                _ = v12
                            } else {
                                v6 = float32(((float64(v6) * 3.14159) * 0.1))
                            }
                            v1 = v1
                            _ = fi11
                        }
                        if false {
                            for fi13 := 0; fi13 < 16; fi13++ {
                                var v14 uint64 = 18446744073709551615
                                v15 := GenS2{F0: false}
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            var v16 int32 = 0
                            _ = v16
                        }
                    }
                    v17 := []int{}
                    v6 = (v6 * v6)
                    _ = fi10
                    _ = v17
                }
                var v18 uint16 = 42
                _ = v18
            } else {
                var v19 int8 = -128
                var v20 bool = v1
                if (18446744073709551615 != 2) {
                    v21 := []int{32767}
                    _ = v21
                } else {
                    var v22 uint8 = 67
                    v4 += (v4 - v4)
                    _ = v22
                }
                _ = v19
                _ = v20
            }
        }
    }
    if v1 {
        v23 := GenS3{F0: -494, F1: v3, F2: v2}
        switch 1 {
        case -562:
            var v24 int = 127
            _ = v24
        case -430:
            v2 += v5
        case 518:
            var v25 uint16 = 255
            _ = v25
        }
        fmt.Println(v1)
        _ = v23
    }
    v6 = v6
    v5 += (v5 >> ((18446744073709551615) & 63))
    var v26 int8 = -108
    fmt.Println(v1)
    var v27 uint64 = 1
    var v28 int8 = -1
    fmt.Println(v6, v4, v27)
    switch "foo bar" {
    case "foo bar":
        v29 := GenS3{F0: -32768, F1: v3, F2: 960}
        _ = v29
    case "":
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 uint8 = 42
            _ = fi30
            _ = v31
        }
    case "\n\t":
        var v32 int8 = (v26 / (((v28 * (v26 | (v26 << ((2) & 63))))) | 1))
        _ = v32
    default:
        fmt.Println(v28, v26)
    }
    var v33 int8 = 127
    var v34 uint16 = 228
    v5 -= (v5 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
