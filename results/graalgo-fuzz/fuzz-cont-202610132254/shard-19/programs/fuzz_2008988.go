package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint64 = 103
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    v1 += v1
    if false {
        v1 = (v1 << ((0) & 63))
    }
    v1 -= 2
    var v6 uint8 = 127
    v6 -= (v6 + (v6 * (v6 + v6)))
    var v7 int8 = 38
    var v8 uint = 305
    if false {
        v9 := []int{32768, 726, 32768}
        _ = v9
    } else {
        v6 = 1
    }
    v7 = (v7 / (((v7 % ((v7) | 1))) | 1))
    if true {
        fmt.Println(v8, v6, v6)
    }
    var v10 float64 = 0.0
    var v11 float32 = 1.5
    if true {
        v2 -= (v2 + v2)
        for fi12 := 0; fi12 < 3; fi12++ {
            switch fi12 {
            case 32768:
                v8 -= v8
            default:
                if (true || (!(!(-2147483648 < 2147483647)))) {
                    var v13 string = "\n\t"
                    _ = v13
                }
            }
            _ = fi12
        }
    } else {
        switch -1 {
        case -893:
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS1{F0: 18446744073709551615, F1: 65535}
                    if ("hello" >= "hello") {
                        delete(v5, (fi14 >> ((v8) & 63)))
                    } else {
                        v6 += 128
                        v6 = v6
                        v17 := []int{-250, fi15, fi15, 4294967295}
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                v18 := GenS1{F0: (v1 + 42), F1: 2147483647}
                v19 := GenS2{F0: 861, F1: 42}
                _ = fi14
                _ = v18
                _ = v19
            }
        case 0:
            v2 -= v2
        case -1:
            if false {
                v20 := GenS1{F0: v1, F1: 4294967295}
                v10 = 1e10
                _ = len(v4)
                _ = v20
            } else {
                v1 += 4294967295
                v8 = (v8 ^ uint(v6))
                switch 128 {
                case 9223372036854775807:
                    v1 = (^(v1))
                case 105:
                    v7 += (v7 % ((v7) | 1))
                case -84:
                    if ((true && (!(3.14159 > v10))) && true) {
                        var v21 float32 = (v11 / 1.5)
                        var v22 int64 = 498
                        for fi23 := 0; fi23 < 19; fi23++ {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v2 = (v2 >> ((1) & 63))
                                v25 := []int{(fi23 / (((fi24 * (-(fi24)))) | 1)), fi24, (fi24 * fi23), 32768}
                                _ = fi24
                                _ = v25
                            }
                            _ = fi23
                        }
                        _ = v21
                        _ = v22
                    }
                default:
                    var v26 uint8 = 255
                    _ = v26
                }
            }
        default:
            delete(v5, 2147483647)
        }
    }
    var v27 uint16 = 636
    var v28 int = -459
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
