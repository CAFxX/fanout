package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
func genFn0() bool {
    if false {
        var v6 uint32 = 2
        var v7 bool = (-1 == 2147483647)
        var v8 bool = v7
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float32 = (3.14159 - 1e10)
        var v12 float64 = 0.0
        v12 = 3.14159
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return (!false)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 uint32 = 65535
    var v4 int32 = -1
    var v5 uint32 = (v3 - (^(v3)))
    var v13 uint = 126
    v4 = v4
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = 1
        for fi16 := 0; fi16 < 13; fi16++ {
            if true {
                var v17 float64 = ((float64(0.1) - 0.0) * -2.25)
                var v18 int64 = -9223372036854775808
                if (v17 == (v17 + v17)) {
                    v3 = (v3 + (^(v3)))
                } else {
                    var v19 int32 = v4
                    fmt.Println(v13, v19, fi14)
                    _ = v19
                }
                _ = v17
                _ = v18
            } else {
                v3 = v3
                var v20 uint64 = 0
                _ = v20
            }
            v21 := GenS1{F0: 42, F1: -32768, F2: int64(v13)}
            _ = fi16
            _ = v21
        }
        _ = fi14
        _ = v15
    }
    v3 = (v5 >> ((42) & 63))
    v4 -= (v4 << ((255) & 63))
    v5 = v3
    switch 500 {
    case 0:
        for fi22 := 0; fi22 < 13; fi22++ {
            v5 = v5
            v4 = v4
            v5 -= v5
            _ = fi22
        }
    default:
        v23 := []int{-371, 9223372036854775807, 128, 0}
        _ = v23
    }
    v4 = 2147483647
    v5 = v3
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
