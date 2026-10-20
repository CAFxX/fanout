package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0(p0 uint64) uint64 {
    var v7 int64 = 0
    var v8 bool = false
    v9 := []int{}
    var v10 int32 = -64
    var v11 uint64 = (p0 % (((^(p0))) | 1))
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int = -349
    v2 := []uint32{}
    var v3 uint64 = 0
    var v4 int8 = -40
    var v5 float64 = 0.1
    var v6 uint8 = 128
    v6 += v6
    v12 := GenS1{F0: 906, F1: v1}
    var v13 uint8 = v6
    var v14 string = ""
    v3 -= (v3 >> ((4294967295) & 63))
    var v15 uint64 = v3
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch v1 {
    case 59:
        for fi16 := 0; fi16 < 18; fi16++ {
            if ((v4 >= v4) || (-353 == 212)) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 float32 = 1e10
                    var v19 uint64 = (^(v3))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint64 = v3
                _ = v20
            }
            fmt.Println(v5, v4)
            _ = fi16
        }
    default:
        v13 += v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
