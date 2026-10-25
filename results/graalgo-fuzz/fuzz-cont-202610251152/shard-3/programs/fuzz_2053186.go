package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0() int8 {
    var v6 string = ""
    v7 := GenS1{F0: v6, F1: 279, F2: 82}
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int8, p1 uint8) uint32 {
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float32 = (1e10 * float32((1.5 * 0.1)))
        fmt.Println(p1)
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 2
    v11 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 int16 = -823
    var v4 float64 = float64(0.0)
    var v5 float32 = -2.25
    v3 = -1
    v3 = v3
    v12 := GenS3{F0: 240, F1: 0}
    switch -9223372036854775808 {
    case 32768:
        var v13 uint32 = 42
        _ = v13
    default:
        var v14 bool = false
        _ = v14
    }
    var v15 bool = (32767 <= 255)
    var v16 int32 = -2147483648
    v17 := v1["hello"]
    _ = len(v1)
    var v18 float64 = 1e10
    if ((v5 / float32(float64(v5))) > v5) {
        v19 := GenS2{F0: 0}
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
