package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 1
    var v3 int64 = 173
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            switch int(v3) {
            case 0:
                fmt.Println(v2)
            default:
                v3 -= (v3 * v3)
            }
            var v7 float64 = (3.14159 - (1e10 + (-2.25 + ((1e10 / 1.0) / 1.0))))
            _ = fi6
            _ = v7
        }
        var v8 int16 = 1
        v9 := []int{int(v2)}
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 uint8 = (v1 * v1)
    switch "hello" {
    case "a":
        var v11 uint16 = 32767
        _ = v11
    default:
        v1 = v10
    }
    var v12 uint64 = 875
    v12 += (v12 - (v12 % ((228) | 1)))
    var v13 int32 = -414
    if ((v1 / (((v10 - 2)) | 1)) == 107) {
        if true {
            v14 := []int{32768, -2147483648}
            _ = v14
        }
    }
    var v15 float64 = (((-2.25 * 1.5) + (float64(0.0) + 0.1)) + 3.14159)
    v16 := GenS1{F0: int8(v10), F1: int(v2)}
    v17 := GenS2{F0: true, F1: "a", F2: 1}
    if true {
        var v18 uint16 = 2
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS3{F0: -2147483648, F1: 103}
            v21 := v4[uint32(v12)]
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    v22 := GenS4{F0: 448}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
