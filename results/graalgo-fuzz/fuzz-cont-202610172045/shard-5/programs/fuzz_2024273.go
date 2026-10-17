package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 int16, p2 int16) uint64 {
    p1 += p1
    p2 += p2
    var v4 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 129
}
func genFn1(p0 uint16, p1 uint8) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = "hello"
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 748, F1: p0, F2: (p1 % ((94) | 1))}
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 float32 = float32(0.0)
    var v2 int32 = 0
    v3 := []int{2147483647}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint = 770
        var v10 int64 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = len(v3)
    var v11 uint8 = 1
    if false {
        v12 := GenS2{F0: 0, F1: v11, F2: 42}
        _ = v12
    } else {
        if true {
            switch "a" {
            case "a":
                if (false && false) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v11 -= v11
                }
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v13 := GenS2{F0: 127, F1: v11, F2: -32768}
                _ = v13
            }
            switch -2147483648 {
            case 94:
                for fi14 := 0; fi14 < 18; fi14++ {
                    v11 -= 255
                    switch fi14 {
                    case -803:
                        if len(v3) > 0 {
                            v3[0] = (fi14 ^ (fi14 & -968))
                        }
                    case 127:
                        _ = len(v3)
                    case 0:
                        fmt.Println(fi14, v1)
                    default:
                        v15 := GenS1{F0: 1, F1: 0, F2: (v11 & 2)}
                        _ = v15
                    }
                    _ = fi14
                }
            }
            fmt.Println(v1, v2, v2)
        }
        v16 := GenS3{F0: 127, F1: 32767, F2: 42}
        _ = v16
    }
    var v17 int16 = -32768
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS4{F0: true}
            if (((fi18 % ((fi18) | 1)) <= fi18) && (v1 > -2.25)) {
                fmt.Println(v17, fi18)
                switch (fi19 | fi19) {
                case 0:
                    if (((float64(0.1) - 0.1) - float64(v1)) == (0.1 * -2.25)) {
                        _ = len(v3)
                    } else {
                        if false {
                            var v21 float64 = float64(v1)
                            var v22 uint8 = v11
                            _ = v21
                            _ = v22
                        } else {
                            if (3.14159 == ((3.14159 + 0.0) + 1e10)) {
                                if true {
                                    if ((^(v11)) >= (v11 % (((v11 ^ v11)) | 1))) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v2 = -370
                                    var v23 uint32 = 770
                                    v11 = 161
                                    _ = v23
                                }
                                fmt.Println(v2)
                            }
                        }
                        v1 -= v1
                    }
                default:
                    var v24 int64 = 42
                    _ = v24
                }
            } else {
                var v25 uint = 1
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v11, v25)
                case "a":
                    var v26 int = (fi19 - fi18)
                    _ = v26
                }
                var v27 uint16 = 1
                _ = v25
                _ = v27
            }
            if (42 < 630) {
                if (!true) {
                    v11 -= (^(v11))
                } else {
                    v17 = (v17 | (v17 & (v17 | 967)))
                }
                if ((v2 | (v2 >> ((1) & 63))) == (v2 >> ((42) & 63))) {
                    v11 = v11
                }
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v28 int64 = 1
    _ = len(v3)
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            v3 = append(v3, fi29)
            if (-861 <= v2) {
                var v31 bool = true
                _ = v31
            }
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println(v1, v11, v28)
    fmt.Println(v17, v2)
    var v32 string = "hello"
    v3 = append(v3, -905)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v32
}
