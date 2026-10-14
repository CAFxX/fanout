package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[uint32]int)
    var v7 int32 = 310
    var v8 int = -580
    v6[462] = (v8 + v8)
    _ = len(v1)
    var v9 bool = false
    v10 := []int{}
    if v9 {
        v11 := v2[(v8 * v8)]
        _ = v11
    }
    v8 = v8
    var v12 uint32 = 2147483648
    var v13 uint = 0
    v14 := GenS1{F0: (v12 * 2147483647), F1: 42}
    var v15 uint = (v13 & v13)
    var v16 float32 = float32((0.1 - 1e10))
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int32 = v7
            _ = fi18
            _ = v19
        }
        if v9 {
            v9 = (!(-1 >= v7))
            v12 -= (^(v12))
            v7 = (v7 + (v7 % ((v7) | 1)))
        } else {
            v20 := v5[(fi17 ^ 127)]
            v9 = v9
            if (0 != 306) {
                switch "\n\t" {
                case "a":
                    if len(v10) > 0 {
                        v10[0] = v20
                    }
                default:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if v9 {
                            var v22 int16 = 717
                            _ = v22
                        } else {
                            v23 := GenS1{F0: 2147483648, F1: (^(v12))}
                            fmt.Println(v12)
                            _ = v23
                        }
                        _ = fi21
                    }
                }
                v20 += v8
            } else {
                if ((!(-9223372036854775808 > 319)) && (-848 == -1)) {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if (-32768 > -1) {
                            v8 = v8
                            v25 := GenS1{F0: (v12 + (v12 * (v12 << (((v13 | (v13 + (v13 - v15)))) & 63)))), F1: v12}
                            v15 = v15
                            _ = v25
                        }
                        _ = len(v6)
                        var v26 int8 = -1
                        _ = fi24
                        _ = v26
                    }
                    delete(v6, v12)
                    _ = len(v10)
                }
            }
            _ = v20
        }
        _ = fi17
    }
    delete(v5, v8)
    fmt.Println(v12, v12)
    if v9 {
        v6[v12] = v8
        if v9 {
            var v27 uint32 = 65535
            v9 = false
            for fi28 := 0; fi28 < 8; fi28++ {
                if (((false && v9) || (v7 == v7)) || v9) {
                    switch -2147483648 {
                    case 2147483647:
                        v29 := GenS1{F0: 42, F1: (^(v12))}
                        _ = v29
                    case 128:
                        fmt.Println(v13)
                    }
                }
                _ = fi28
            }
            _ = v27
        }
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 int8 = 127
            _ = fi30
            _ = v31
        }
    }
    var v32 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
