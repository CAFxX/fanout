package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 int8, p1 uint32) int {
    fmt.Println(p1, p0)
    var v5 float64 = 3.14159
    v6 := []int{120, 4294967295, 127}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -770
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{}
    var v3 uint64 = 341
    v4 := make(map[uint32]int)
    var v7 int64 = 42
    for fi8 := 0; fi8 < 6; fi8++ {
        if false {
            fmt.Println(v7, v7)
        } else {
            fmt.Println(fi8)
            switch "a" {
            case "hello":
                var v9 int = genFn0(15, v1)
                _ = v9
            default:
                v1 = (v1 / ((65535) | 1))
            }
        }
        fmt.Println(v3)
        var v10 float32 = (0.1 / 1e10)
        _ = fi8
        _ = v10
    }
    v3 += v3
    var v11 float64 = (0.0 / 1.0)
    var v12 uint16 = 255
    v1 = (v1 % ((v1) | 1))
    v3 = (^(v3))
    if true {
        if false {
            v12 += v12
        } else {
            v13 := GenS1{F0: v12, F1: 1}
            if (100 >= 100) {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 uint64 = v3
                    _ = fi14
                    _ = v15
                }
                v7 = -9223372036854775808
                if true {
                    v11 -= v11
                    var v16 uint8 = 2
                    for fi17 := 0; fi17 < 12; fi17++ {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            var v19 uint32 = v1
                            _ = fi18
                            _ = v19
                        }
                        if false {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                var v21 int32 = 0
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            fmt.Println(fi17)
                        }
                        v22 := GenS1{F0: 80, F1: int32(v16)}
                        _ = fi17
                        _ = v22
                    }
                    _ = v16
                }
            }
            v1 = (^(v1))
            _ = v13
        }
        v7 = (v7 / (((v7 << ((2) & 63))) | 1))
        var v23 int32 = -1
        _ = v23
    }
    v11 += v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
