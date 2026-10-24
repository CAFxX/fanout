package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    v4 := []int{}
    var v5 int64 = 42
    var v6 uint16 = 563
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 0
    }
    v1 = (v1 * (v1 << ((42) & 63)))
    var v7 float64 = (float64(1.5) / 1.0)
    v1 = (v1 >> ((4294967295) & 63))
    var v8 uint16 = 42
    _ = len(v3)
    var v9 int16 = -32768
    if true {
        var v10 bool = true
        for fi11 := 0; fi11 < 15; fi11++ {
            v5 -= int64(fi11)
            if true {
                v5 = 9223372036854775807
            } else {
                if (!true) {
                    var v12 float32 = float32(1e10)
                    var v13 int16 = v9
                    _ = v12
                    _ = v13
                }
                switch fi11 {
                case 1:
                    v14 := GenS1{F0: fi11, F1: (v8 >> ((42) & 63)), F2: 0}
                    _ = v14
                case 2147483647:
                    var v15 int64 = (v5 - (v5 / ((398) | 1)))
                    _ = v15
                case -9223372036854775808:
                    v2 -= v2
                default:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v17 := GenS2{F0: v10, F1: 42}
                        _ = fi16
                        _ = v17
                    }
                }
            }
            _ = fi11
        }
        _ = len(v4)
        _ = v10
    } else {
        v1 += v1
    }
    fmt.Println(v7)
    v2 = -2147483648
    for fi18 := 0; fi18 < 15; fi18++ {
        if false {
            fmt.Println(v5, v2)
            if true {
                var v19 bool = false
                _ = v19
            }
        }
        var v20 int8 = 42
        _ = fi18
        _ = v20
    }
    switch 32767 {
    case 358:
        var v21 uint32 = 639
        _ = v21
    case -427:
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
