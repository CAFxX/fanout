package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 uint8) uint8 {
    v5 := []int{-1, 555, -1}
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: true, F1: 132, F2: 492}
        var v8 float32 = (-2.25 + (float32(1e10) / 3.14159))
        fmt.Println(v8, v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]int)
    var v3 uint8 = 255
    var v4 float64 = ((-2.25 / -2.25) / ((-2.25 / -2.25) - (0.0 + 1.5)))
    v3 += v3
    v9 := []int{127}
    switch 988 {
    case -987:
        var v10 uint16 = 41
        _ = v10
    case -927:
        v1 = (v1 & v1)
    case 0:
        v11 := GenS2{F0: false, F1: 474, F2: uint16(v3)}
        _ = v11
    default:
        switch "\n\t" {
        case "hello":
            if false {
                if (1e10 > (1e10 + 1.5)) {
                    var v12 float64 = ((v4 * v4) * float64(((1.5 * (float32(v4) / 1.5)) / ((((1e10 + float32(-2.25)) + 1e10) / 0.1) + 1.5))))
                    v9 = append(v9, 9223372036854775807)
                    _ = v12
                }
                var v13 string = "foo bar"
                var v14 uint32 = uint32(v3)
                _ = v13
                _ = v14
            } else {
                var v15 int8 = -1
                _ = v15
            }
        default:
            v16 := GenS1{F0: (-822 != 258), F1: 18446744073709551615, F2: 42}
            _ = v16
        }
    }
    fmt.Println(v1, v1, v3)
    fmt.Println(v4, v4)
    v2[2147483647] = 2147483647
    _ = len(v9)
    if false {
        v1 = 42
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 int8 = 1
                _ = fi18
                _ = v19
            }
            var v20 float64 = 0.0
            for fi21 := 0; fi21 < 16; fi21++ {
                v1 -= v1
                var v22 uint = 4294967295
                _ = fi21
                _ = v22
            }
            _ = fi17
            _ = v20
        }
        var v23 uint = 42
        _ = v23
    } else {
        v9 = append(v9, 4294967295)
    }
    v24 := GenS2{F0: (280 == 0), F1: 1, F2: uint16(v1)}
    for fi25 := 0; fi25 < 11; fi25++ {
        v3 += (v3 >> ((939) & 63))
        if (65535 <= 4294967295) {
            for fi26 := 0; fi26 < 17; fi26++ {
                switch fi25 {
                case 42:
                    delete(v2, int(v1))
                case 4294967295:
                    var v27 float64 = (v4 - ((float64(-2.25) + (1e10 * v4)) * v4))
                    _ = v27
                case 0:
                    if (!(!(42 == 0))) {
                        if (false || false) {
                            if true {
                                switch (fi25 + fi25) {
                                case 443:
                                    v28 := GenS2{F0: true, F1: 159, F2: 32767}
                                    _ = v28
                                }
                                fmt.Println(v3)
                                v3 = (v3 % (((^(v3))) | 1))
                            }
                            var v29 uint16 = uint16(v1)
                            if false {
                                if false {
                                    var v30 uint = 811
                                    v29 = 1
                                    _ = v30
                                } else {
                                    v1 = (v1 ^ v1)
                                }
                                var v31 int = (-(fi25))
                                v32 := GenS3{F0: (v3 | v3), F1: 1}
                                _ = v31
                                _ = v32
                            }
                            _ = v29
                        }
                        v4 = v4
                        var v33 string = "hello"
                        _ = v33
                    } else {
                        if len(v9) > 0 {
                            v9[0] = (fi25 * 2147483647)
                        }
                    }
                }
                v1 = v1
                _ = fi26
            }
        }
        var v34 uint32 = v1
        _ = fi25
        _ = v34
    }
    switch "a" {
    case "":
        fmt.Println(v4)
    default:
        switch "\n\t" {
        case "foo bar":
            var v35 uint64 = 18446744073709551615
            _ = v35
        case "a":
            fmt.Println(v1, v1)
        default:
            v1 = v1
        }
    }
    v36 := GenS4{F0: ""}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v4 = (((0.1 + v4) / v4) * (v4 + ((v4 * 1.5) / (v4 / v4))))
    var v37 uint16 = 822
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v9
}
