package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int16 {
    if ((42 != 1) || (!(2 < 2147483647))) {
        var v4 float64 = -2.25
        var v5 int32 = -865
        v5 -= (v5 % (((v5 << ((811) & 63))) | 1))
        _ = v4
        _ = v5
    }
    v6 := []int{539, -2147483648, 42}
    if true {
        var v7 float32 = -2.25
        _ = v7
    }
    switch 9223372036854775807 {
    case 1:
        var v8 float32 = (-2.25 / 1.0)
        _ = v8
    case 536:
        var v9 int16 = 0
        _ = v9
    default:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint64 = 1
            _ = fi10
            _ = v11
        }
    }
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{65535, 42}
    v2 := make(map[string]int)
    var v3 uint8 = 115
    v3 = (^(v3))
    switch -101 {
    case -9223372036854775808:
        v12 := GenS1{F0: -2147483648}
        _ = v12
    default:
        if ((true && false) && (4294967295 == 923)) {
            v3 = 42
        } else {
            if (!true) {
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 float32 = 1e10
                    if true {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v3 += v3
                            var v16 uint8 = 2
                            _ = fi15
                            _ = v16
                        }
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    fmt.Println(fi13, v14, v3)
                    _ = fi13
                    _ = v14
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
            } else {
                var v18 int = 42
                _ = v18
            }
        }
    }
    fmt.Println(v3)
    v3 = v3
    if true {
        v3 = (^(v3))
        v3 = (v3 / ((2) | 1))
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            v3 = v3
            _ = fi19
        }
    }
    if len(v1) > 0 {
        v1[0] = 432
    }
    v20 := GenS1{F0: 1}
    v3 = (v3 << ((0) & 63))
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
