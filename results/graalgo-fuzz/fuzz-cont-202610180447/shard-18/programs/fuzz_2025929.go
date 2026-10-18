package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []int{-9223372036854775808, -9223372036854775808, 2147483647, 0}
    v2 = "a"
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    v5 = append(v5, 2147483647)
    v6 := GenS1{F0: -2147483648}
    if ((42 >= 4294967295) && (!(-2147483648 <= -2147483648))) {
        var v7 int32 = 417
        for fi8 := 0; fi8 < 9; fi8++ {
            v7 += (v7 + v7)
            _ = fi8
        }
        if (-9223372036854775808 != 42) {
            var v9 uint8 = 42
            var v10 int = 42
            _ = v10
            _ = v9
        } else {
            if false {
                if (true && (2147483648 > 4294967295)) {
                    var v11 int64 = 0
                    var v12 int16 = 1000
                    var v13 int32 = -1
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v14 uint32 = 0
                    _ = v14
                }
                if false {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            var v17 uint32 = 4294967295
                            _ = fi16
                            _ = v17
                        }
                        var v18 float64 = float64(((-2.25 / 1.0) / (float32(-2.25) - 0.1)))
                        _ = fi15
                        _ = v18
                    }
                    v19 := GenS2{F0: v7, F1: 569}
                    v20 := GenS1{F0: (v7 << ((42) & 63))}
                    _ = v19
                    _ = v20
                } else {
                    v7 += (v7 ^ (v7 * v7))
                    var v21 uint64 = 42
                    _ = v21
                }
                var v22 uint8 = 128
                _ = v22
            } else {
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint16 = 197
                    _ = fi23
                    _ = v24
                }
                v2 = ((("foo bar" + v2) + "hello") + (v2 + v2))
            }
        }
        _ = v7
    } else {
        fmt.Println(v2)
    }
    fmt.Println(v2)
    var v25 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
