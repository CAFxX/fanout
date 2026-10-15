package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0() int32 {
    if false {
        if ((4294967295 >= 398) || false) {
            v4 := []int{}
            var v5 int = 2147483647
            _ = v4
            _ = v5
        }
    } else {
        var v6 int32 = 42
        v7 := []int{2147483647, 123, 4294967295, 893}
        _ = v6
        _ = v7
    }
    v8 := []int{-1, -1, 449, -9223372036854775808}
    var v9 uint16 = 1
    if (42 == 127) {
        v9 = v9
        if false {
            var v10 bool = (!false)
            v9 = v9
            _ = v10
        }
    }
    _ = v8
    _ = v9
    return -215
}

func main() {
    var v1 int = 273
    v2 := make(map[uint32]uint32)
    var v3 int8 = -1
    var v11 float32 = -2.25
    v11 += ((v11 - float32((1.5 - 0.0))) + 1e10)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 bool = true
        _ = fi12
        _ = v13
    }
    var v14 int64 = 0
    v1 = (v1 + v1)
    switch v1 {
    case 1:
        var v15 float64 = 1e10
        _ = v15
    case -860:
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS1{F0: 42, F1: v1, F2: 255}
            _ = len(v2)
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        switch fi18 {
        case 347:
            v19 := []int{(v1 * 0), fi18, 9223372036854775807, v1}
            _ = v19
        case 42:
            if true {
                v20 := GenS2{F0: "\n\t", F1: 369, F2: 65535}
                _ = v20
            }
        default:
            v21 := GenS1{F0: 92, F1: (-(fi18)), F2: 42}
            _ = v21
        }
        _ = fi18
    }
    v22 := GenS3{F0: 383, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v3
}
