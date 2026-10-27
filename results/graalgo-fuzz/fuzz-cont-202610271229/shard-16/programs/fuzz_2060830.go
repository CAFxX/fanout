package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 string) int {
    switch "a" {
    case "hello":
        p2 = "foo bar"
    default:
        var v7 uint8 = p1
        _ = v7
    }
    switch -1 {
    case -681:
        var v8 int8 = 127
        _ = v8
    case 0:
        switch p2 {
        case "a":
            v9 := []int{0, -464, 32767, -822}
            _ = v9
        case "\n\t":
            v10 := []int{128, 188}
            _ = v10
        default:
            v11 := []int{32767, 577, 2147483647}
            _ = v11
        }
    }
    p2 = "hello"
    _ = p0
    _ = p1
    _ = p2
    return 309
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = ((0.1 + ((((((1e10 - 1e10) + 3.14159) * 1e10) - (((-2.25 * ((1.5 / float32(((3.14159 - (float64(float32(1e10)) - ((0.1 + 1e10) + ((((3.14159 / 3.14159) * (0.1 - -2.25)) - 1.5) * 1.5)))) - ((3.14159 / 1e10) + 0.0)))) - 1e10)) - (-2.25 / (1.5 + 3.14159))) + 1.5)) * 0.0) * 3.14159)) + -2.25)
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    v5 := make(map[int]bool)
    v6 := make(map[string]int)
    var v12 float32 = (1e10 * v2)
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: 42}
        v4 = (v4 % (((v4 >> ((65535) & 63))) | 1))
        var v15 int8 = -99
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int = 986
    var v17 float64 = 0.0
    v12 = 1e10
    for fi18 := 0; fi18 < 3; fi18++ {
        switch "a" {
        case "\n\t":
            switch fi18 {
            case 0:
                v12 = 0.0
            case -2147483648:
                v19 := []int{(v16 - 32768), genFn0(949, 255, "foo bar"), (^(fi18)), 4294967295}
                _ = v19
            case 32768:
                var v20 int32 = -2147483648
                _ = v20
            }
        }
        _ = fi18
    }
    if false {
        switch genFn0(845, 127, "hello") {
        case -1:
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(v12, fi21, v12)
                _ = fi21
            }
        default:
            if (1.5 == -2.25) {
                v22 := v6[("a" + "\n\t")]
                v23 := GenS2{F0: 1, F1: 601}
                v12 = -2.25
                _ = v22
                _ = v23
            }
        }
        v12 = v2
    } else {
        if false {
            for fi24 := 0; fi24 < 18; fi24++ {
                v12 += -2.25
                _ = fi24
            }
            v25 := GenS1{F0: 110}
            if (-714 <= -2147483648) {
                switch 644 {
                case 2147483647:
                    var v26 float64 = float64((float32(0.0) * v12))
                    _ = v26
                case 152:
                    v27 := GenS2{F0: 4294967295, F1: 754}
                    _ = v27
                case 128:
                    v2 = v2
                default:
                    if false {
                        v2 += (1.5 / float32(v17))
                        v5[(v16 << ((255) & 63))] = true
                        var v28 uint8 = 128
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 18; fi29++ {
                            switch fi29 {
                            case 42:
                                var v30 string = "a"
                                _ = v30
                            case 2147483647:
                                v31 := []int{fi29, 0}
                                _ = v31
                            case 9223372036854775807:
                                v12 = v2
                            }
                            for fi32 := 0; fi32 < 11; fi32++ {
                                v33 := []int{genFn0(0, 2, "a")}
                                v2 -= float32(float64((v2 / 3.14159)))
                                _ = fi32
                                _ = v33
                            }
                            _ = fi29
                        }
                        switch 2147483647 {
                        case -154:
                            for fi34 := 0; fi34 < 10; fi34++ {
                                for fi35 := 0; fi35 < 9; fi35++ {
                                    v16 -= (^(fi35))
                                    _ = fi35
                                }
                                v36 := []int{genFn0(124, 255, "\n\t"), -9223372036854775808, -926, fi34}
                                _ = fi34
                                _ = v36
                            }
                        }
                        fmt.Println(v2, v12)
                    }
                }
                var v37 bool = false
                for fi38 := 0; fi38 < 12; fi38++ {
                    v39 := []int{genFn0(2147483648, 1, "foo bar")}
                    var v40 uint8 = 0
                    v37 = ((v40 + v40) > v40)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = v37
            } else {
                var v41 int16 = v4
                _ = v41
            }
            _ = v25
        }
    }
    _ = len(v5)
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
