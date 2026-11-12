package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 float32 = (0.0 - 0.0)
    var v2 uint16 = 2
    var v3 uint16 = 46
    v1 -= (v1 * (v1 - ((0.0 / 1e10) / float32((0.1 * (float64((v1 * -2.25)) + (3.14159 / 1.0)))))))
    var v4 uint8 = 45
    v1 = float32(0.1)
    v5 := GenS1{F0: true, F1: v3}
    v1 = (v1 * -2.25)
    if false {
        v3 = v3
        fmt.Println(v1, v4, v4)
    } else {
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                if true {
                    var v8 uint8 = v4
                    if false {
                        v9 := []int{(fi7 | 2147483647)}
                        _ = v9
                    }
                    var v10 uint16 = (^(v2))
                    v3 -= 0
                    var v11 uint16 = (^(v10))
                    _ = v10
                    _ = v11
                    _ = v8
                }
                v12 := []int{128}
                _ = fi7
                _ = v12
            }
            v4 -= (v4 & v4)
            _ = fi6
        }
        var v13 int32 = 236
        _ = v13
    }
    switch "hello" {
    case "a":
        var v14 uint8 = 0
        _ = v14
    case "hello":
        v4 -= 42
    default:
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint16 = v3
            var v17 uint8 = (^(v4))
            switch 2147483647 {
            case -870:
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 uint = 65535
                    if (v1 == v1) {
                        v20 := []int{fi15}
                        _ = v20
                    }
                    v4 += (v4 * v4)
                    v3 += v2
                    _ = fi18
                    _ = v19
                }
            case 225:
                var v21 uint32 = uint32(v3)
                _ = v21
            case -159:
                if false {
                    v22 := []int{fi15}
                    if true {
                        var v23 uint = 599
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 float64 = 1e10
                            v25 = v25
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        var v26 bool = true
                        var v27 int32 = 643
                        _ = v26
                        _ = v27
                    }
                    var v28 uint64 = 4294967295
                    _ = v22
                    _ = v28
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v29 uint = 0
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint8 = 1
        _ = fi30
        _ = v31
    }
    if true {
        if false {
            v4 += v4
            v29 = 18446744073709551615
        } else {
            var v32 string = "hello"
            var v33 float32 = 1e10
            _ = v32
            _ = v33
        }
    } else {
        var v34 int = -9223372036854775808
        _ = v34
    }
    v2 -= (v3 / ((32767) | 1))
    v35 := []int{-2147483648, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
