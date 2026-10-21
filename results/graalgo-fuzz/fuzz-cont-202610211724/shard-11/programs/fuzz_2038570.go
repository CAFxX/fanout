package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int8 = 1
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    var v4 uint8 = 42
    v1 -= (v1 * 42)
    var v5 uint32 = 2147483647
    if false {
        var v6 int64 = 148
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: (^(v5))}
    for fi8 := 0; fi8 < 2; fi8++ {
        switch fi8 {
        case -106:
            var v9 float64 = -2.25
            _ = v9
        }
        if false {
            if false {
                var v10 int16 = -32768
                _ = v10
            } else {
                var v11 float64 = (((1e10 - 0.1) * (3.14159 + ((0.0 * 0.0) * 0.0))) + float64(3.14159))
                _ = v11
            }
            if (!false) {
                if false {
                    var v12 uint32 = v5
                    _ = v12
                } else {
                    fmt.Println(v1, fi8)
                    v5 -= v5
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if true {
                            v14 := []int{fi8, fi8}
                            _ = v14
                        }
                        if ((3.14159 - (3.14159 + 3.14159)) < 0.1) {
                            var v15 uint64 = 42
                            v15 = (v15 % (((^(v15))) | 1))
                            var v16 int = (-(fi8))
                            _ = v15
                            _ = v16
                        }
                        switch "\n\t" {
                        case "a":
                            switch "" {
                            case "foo bar":
                                v5 -= 665
                            case "a":
                                v4 = 1
                            default:
                                var v17 int64 = 42
                                _ = v17
                            }
                        default:
                            var v18 int8 = v1
                            _ = v18
                        }
                        _ = fi13
                    }
                }
                var v19 int16 = -476
                _ = v19
            } else {
                v4 = (v4 * (v4 - 127))
            }
            _ = len(v3)
        }
        var v20 int64 = -9223372036854775808
        _ = fi8
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint64 = 36
        _ = fi21
        _ = v22
    }
    var v23 uint = 18446744073709551615
    v24 := GenS1{F0: 32767, F1: (^(v5))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
