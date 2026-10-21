package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int }
func genFn0() int64 {
    var v5 float64 = 3.14159
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 += v5
        _ = fi6
    }
    v5 -= ((v5 + v5) * v5)
    v7 := GenS1{F0: 0, F1: 814, F2: 237}
    var v8 uint8 = 127
    _ = v5
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint8 = 2
    var v9 uint32 = 2147483648
    var v10 string = "\n\t"
    _ = len(v3)
    if true {
        fmt.Println(v4, v9)
        switch 32768 {
        case 42:
            switch "foo bar" {
            case "a":
                v11 := GenS2{F0: 1, F1: -238}
                _ = v11
            case "foo bar":
                var v12 uint64 = 2
                _ = v12
            default:
                if (0 == 42) {
                    v9 = (v9 ^ uint32(v4))
                    switch "hello" {
                    case "":
                        if (0 > 42) {
                            v13 := GenS1{F0: 160, F1: 0, F2: 74}
                            _ = v13
                        }
                    case "hello":
                        v9 += 2147483648
                    case "a":
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if (false || (682 < (v9 << ((255) & 63)))) {
                                var v15 float64 = 1.5
                                _ = v15
                            } else {
                                fmt.Println(v9, v4)
                                v4 -= 0
                            }
                            v9 -= 0
                            _ = fi14
                        }
                    default:
                        v16 := []int{179, int(v9)}
                        _ = v16
                    }
                    v10 = v10
                } else {
                    var v17 float64 = 0.1
                    _ = v17
                }
            }
        case 32767:
            v18 := []int{-91, 32768, 2147483647}
            _ = v18
        default:
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 float32 = 3.14159
                _ = fi19
                _ = v20
            }
        }
    }
    fmt.Println(v4)
    if (true || (32767 >= 0)) {
        var v21 uint16 = 1
        if true {
            _ = len(v3)
            v22 := GenS1{F0: 18446744073709551615, F1: -1, F2: v4}
            _ = v22
        }
        v23 := GenS1{F0: 65535, F1: -1, F2: 128}
        _ = v21
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        _ = len(v3)
        switch fi24 {
        case -9223372036854775808:
            var v25 bool = (-2.25 >= 3.14159)
            _ = v25
        case 128:
            for fi26 := 0; fi26 < 1; fi26++ {
                if false {
                    var v27 uint64 = 251
                    v1[2147483648] = fi24
                    _ = v27
                } else {
                    v28 := GenS2{F0: 9223372036854775807, F1: fi24}
                    if (!(((false || false) || (!(float64(((float32(1.5) / 1.0) * -2.25)) > 0.0))) && false)) {
                        fmt.Println(v4, fi24)
                    } else {
                        fmt.Println(fi26, v4)
                        v9 = 2147483648
                        var v29 int64 = 9223372036854775807
                        _ = v29
                    }
                    _ = v28
                }
                _ = fi26
            }
        case 0:
            var v30 int64 = 42
            _ = v30
        default:
            v9 = v9
        }
        v31 := GenS1{F0: 0, F1: -1, F2: (v4 + v4)}
        _ = fi24
        _ = v31
    }
    v32 := GenS1{F0: 18446744073709551615, F1: 661, F2: (v4 ^ (v4 / ((v4) | 1)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
