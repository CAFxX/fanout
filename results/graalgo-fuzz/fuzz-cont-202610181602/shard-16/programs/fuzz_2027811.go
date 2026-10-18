package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0() int16 {
    v4 := []int{}
    var v5 uint8 = 42
    var v6 int32 = -1
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 372
    var v2 int64 = -813
    var v3 float64 = 0.0
    var v7 float32 = 1e10
    var v8 uint32 = 42
    var v9 uint32 = v8
    v10 := []int{2147483647, 602}
    if true {
        if false {
            v11 := GenS1{F0: 153}
            switch 42 {
            case 42:
                for fi12 := 0; fi12 < 15; fi12++ {
                    fmt.Println(v1, v8, fi12)
                    fmt.Println(fi12, v3)
                    if (!(42 >= 100)) {
                        if (true || false) {
                            v8 += 2
                            switch fi12 {
                            case 42:
                                v13 := GenS2{F0: false}
                                _ = v13
                            case 2147483647:
                                v7 = ((v7 - float32(v3)) * 1e10)
                            case -9223372036854775808:
                                for fi14 := 0; fi14 < 11; fi14++ {
                                    if false {
                                        var v15 uint16 = 0
                                        v15 = (^(v15))
                                        var v16 bool = (42 >= 153)
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = fi14
                                }
                            default:
                                v1 -= (^(v1))
                            }
                        } else {
                            v17 := GenS3{F0: 354, F1: v2}
                            _ = v17
                        }
                        v7 = -2.25
                        _ = len(v10)
                    } else {
                        _ = len(v10)
                        var v18 float32 = -2.25
                        _ = v18
                    }
                    _ = fi12
                }
            case 249:
                var v19 uint64 = (^(v1))
                _ = v19
            case 9223372036854775807:
                v3 = ((v3 - ((float64(0.1) / ((v3 * (float64((0.1 - 1.5)) * v3)) + v3)) * (v3 - (v3 / v3)))) - -2.25)
            }
            _ = v11
        } else {
            v20 := GenS3{F0: 1, F1: 118}
            v10 = append(v10, 2147483647)
            v3 += 0.1
            _ = v20
        }
        v1 -= 572
        v7 = float32(0.1)
    }
    _ = len(v10)
    if len(v10) > 0 {
        v10[0] = -9223372036854775808
    }
    var v21 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
