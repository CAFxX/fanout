package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int8; F2 int32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint8, p2 uint) uint {
    if false {
        var v5 uint64 = 28
        _ = v5
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return (p2 + (^(p2)))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 float32 = (((1e10 - (float32((-2.25 - (float64((0.0 * 3.14159)) - 1.5))) + 1.5)) - 1e10) - 1.5)
    v4 := make(map[uint32]int)
    v6 := []int{350, 9223372036854775807, 32768}
    v4[958] = 42
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 2; fi8++ {
        delete(v2, (fi8 + fi8))
        v9 := GenS1{F0: 3, F1: -789, F2: 2147483647}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 31, F1: 662}
    switch "foo bar" {
    case "a":
        switch 2147483647 {
        case 2147483647:
            switch "a" {
            case "a":
                var v11 float64 = 3.14159
                _ = v11
            case "":
                fmt.Println(v3)
            }
        case 127:
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 16; fi13++ {
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                    _ = fi13
                }
                fmt.Println(v7)
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14, fi12, v3)
                    v7 = 0.1
                    switch "hello" {
                    case "a":
                        switch ("" + "hello") {
                        case "":
                            v6 = append(v6, (fi14 & 4294967295))
                        case "hello":
                            switch "foo bar" {
                            case "a":
                                v15 := GenS2{F0: 42, F1: 1}
                                _ = v15
                            }
                        }
                    default:
                        v3 += v3
                        v7 += v3
                        v16 := GenS1{F0: 1, F1: 42, F2: -934}
                        var v17 string = "hello"
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                _ = fi12
            }
        case 950:
            for fi18 := 0; fi18 < 13; fi18++ {
                v7 = v3
                switch ("" + ("hello" + "hello")) {
                case "\n\t":
                    if ((true && (42 == 2)) || (0.1 == float64(v7))) {
                        switch "a" {
                        case "hello":
                            v3 -= (0.1 + ((3.14159 / float32(1e10)) * v3))
                        case "a":
                            v4[922] = fi18
                        case "\n\t":
                            if false {
                                v7 = v7
                                v7 += 3.14159
                                fmt.Println(v3)
                            } else {
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    fmt.Println(v7)
                                    if (-2.25 >= 1e10) {
                                        switch (fi18 - (^(fi18))) {
                                        case 1:
                                            if ((0 < 255) || (2 >= 128)) {
                                                switch fi18 {
                                                case 0:
                                                    delete(v4, 1)
                                                default:
                                                    v20 := GenS2{F0: -1, F1: 255}
                                                    _ = v20
                                                }
                                            } else {
                                                v1[65535] = ""
                                            }
                                        case 119:
                                            fmt.Println(fi19)
                                        }
                                        fmt.Println(fi19, v3, fi19)
                                    }
                                    _ = fi19
                                }
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    var v22 uint32 = 4294967295
                                    delete(v2, (fi18 >> ((2) & 63)))
                                    _ = fi21
                                    _ = v22
                                }
                                v7 = float32(-2.25)
                            }
                        default:
                            v3 = v7
                        }
                        v23 := GenS3{F0: 4294967295}
                        _ = v23
                    }
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "foo bar":
                    var v24 string = "foo bar"
                    _ = v24
                default:
                    var v25 int32 = 388
                    _ = v25
                }
                _ = fi18
            }
        default:
            v3 = (((((3.14159 + v3) + (v3 - v7)) + (v7 - ((0.1 / v3) + v3))) + 0.0) / v7)
        }
    default:
        fmt.Println(v7, v7)
    }
    v26 := GenS4{F0: 2, F1: 127, F2: -2147483648}
    v7 = (v3 * (v7 / float32(((0.0 * 1.5) / (3.14159 / (((((3.14159 - 0.0) + ((3.14159 / (1.5 - 3.14159)) - ((1.5 - (1.5 + (1.5 / (((-2.25 / 1e10) - 1e10) / 0.1)))) - 1.5))) * ((-2.25 + (((((1e10 + -2.25) * 1e10) / 1.0) - (((0.0 * 1e10) / ((3.14159 - (-2.25 / 1.5)) / 1.5)) * 1.5)) - 0.1)) * ((((1.5 * float64(float32(0.0))) + 0.0) - 0.0) - float64((v3 / 1.0))))) + float64(3.14159)) - 1.5))))))
    v3 = v7
    v27 := GenS3{F0: 2}
    fmt.Println(v7, v7)
    v3 -= (float32((float64(v3) + (0.1 / 0.1))) / -2.25)
    v7 = v3
    var v28 int16 = 32767
    v3 = v3
    v29 := GenS5{F0: -1}
    fmt.Println(v28, v3)
    delete(v1, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
