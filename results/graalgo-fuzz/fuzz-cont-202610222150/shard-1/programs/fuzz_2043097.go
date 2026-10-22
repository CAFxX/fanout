package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() uint32 {
    v5 := GenS1{F0: 32767}
    switch -1 {
    case 128:
        var v6 uint64 = 656
        _ = v6
    default:
        var v7 int32 = -2147483648
        _ = v7
    }
    var v8 uint = 65535
    v9 := GenS1{F0: 42}
    _ = v5
    _ = v8
    _ = v9
    return 672
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483648, 464}
    var v4 float32 = v1
    var v10 int64 = -951
    fmt.Println(v1, v4, v1)
    switch 4294967295 {
    case -371:
        v11 := GenS2{F0: 1}
        _ = v11
    default:
        v10 += v10
    }
    if false {
        v1 = ((float32(float64(v1)) * v1) * (((v4 + v4) * 3.14159) * -2.25))
        var v12 uint64 = 4294967295
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := v2[42]
            _ = fi13
            _ = v14
        }
        var v15 int = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float64 = (0.0 + (-2.25 + (1e10 * float64(v1))))
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 2}
    v2[731] = 128
    var v19 int8 = -19
    if true {
        v20 := GenS3{F0: 33, F1: 231}
        for fi21 := 0; fi21 < 17; fi21++ {
            v1 = (3.14159 * v4)
            _ = fi21
        }
        _ = v20
    }
    v19 = v19
    var v22 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
