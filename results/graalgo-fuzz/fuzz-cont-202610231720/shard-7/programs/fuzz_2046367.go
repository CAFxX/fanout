package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float64 = ((1.5 * (0.1 * ((((1e10 + ((1.5 + -2.25) - float64(3.14159))) + -2.25) + 0.0) * (1.5 + -2.25)))) * -2.25)
    var v2 uint = 405
    var v3 uint8 = 2
    v4 := make(map[int]int)
    var v5 int32 = -767
    v2 -= (v2 + v2)
    v6 := GenS1{F0: 4294967295, F1: 42}
    if true {
        switch -353 {
        case -895:
            for fi7 := 0; fi7 < 14; fi7++ {
                fmt.Println(v5, v1)
                var v8 int = (fi7 ^ fi7)
                switch v8 {
                case 295:
                    if (int16(v2) > -105) {
                        if (!(9223372036854775807 < 602)) {
                            switch "a" {
                            case "\n\t":
                                v9 := GenS2{F0: v5, F1: fi7}
                                _ = v9
                            }
                            if (-1 == 1) {
                                v4[fi7] = (fi7 * fi7)
                                var v10 string = "a"
                                v2 -= 18446744073709551615
                                _ = v10
                            } else {
                                v2 += 65535
                                v8 = (fi7 << ((v2) & 63))
                            }
                        }
                        v11 := GenS2{F0: (v5 + v5), F1: v8}
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v8 = fi12
                            _ = fi12
                        }
                        var v13 uint8 = (v3 << ((uint(v3)) & 63))
                        _ = v13
                    }
                case -1:
                    if true {
                        if ((v5 >= 161) || false) {
                            for fi14 := 0; fi14 < 14; fi14++ {
                                v15 := GenS1{F0: 42, F1: 2147483647}
                                _ = fi14
                                _ = v15
                            }
                        } else {
                            var v16 uint64 = 2
                            v17 := []int{int(v2), (fi7 % ((fi7) | 1)), -264}
                            if false {
                                v18 := GenS3{F0: (^(v16))}
                                v16 = 389
                                if true {
                                    for fi19 := 0; fi19 < 1; fi19++ {
                                        var v20 float64 = (float64(1.5) * (v1 + v1))
                                        if (!((v16 % (((^(v16))) | 1)) == 2)) {
                                            fmt.Println(v16, v2, v16)
                                            var v21 float64 = v1
                                            _ = v21
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        }
                    } else {
                        v1 = float64(float32(v1))
                    }
                }
                _ = fi7
                _ = v8
            }
        case 4294967295:
            v22 := v4[42]
            _ = v22
        default:
            fmt.Println(v5, v5)
        }
    }
    v2 -= 306
    var v23 bool = false
    fmt.Println(v1, v23, v5)
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v2, v5, fi24)
        _ = fi24
    }
    v23 = false
    var v25 uint = uint(v3)
    v1 = 1.5
    if v23 {
        var v26 float32 = (float32(((v1 * 1e10) - v1)) + 1.5)
        _ = v26
    }
    v3 = (v3 & v3)
    var v27 uint16 = 2
    var v28 bool = (-9223372036854775808 < -1)
    if false {
        fmt.Println(v27)
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 int8 = 42
            v31 := GenS2{F0: -180, F1: (fi29 << ((42) & 63))}
            v30 -= 77
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := []int{fi32, fi32, fi32}
            _ = fi32
            _ = v33
        }
    } else {
        v34 := GenS1{F0: 679, F1: 2147483648}
        v23 = (0 != uint32(v3))
        v35 := GenS2{F0: (v5 ^ v5), F1: 32768}
        _ = v34
        _ = v35
    }
    var v36 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
