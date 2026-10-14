package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 309
    var v3 uint8 = 42
    var v4 int16 = -972
    for fi5 := 0; fi5 < 8; fi5++ {
        v2 = (v2 >> ((0) & 63))
        _ = fi5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: 42, F1: 65535, F2: 18446744073709551615}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 bool = false
    switch 9223372036854775807 {
    case 4294967295:
        var v10 int32 = -2147483648
        _ = v10
    case 1:
        v11 := []int{}
        _ = v11
    case -452:
        v12 := GenS2{F0: 42}
        _ = v12
    default:
        var v13 uint32 = uint32(v3)
        _ = v13
    }
    v14 := []int{128}
    v9 = v9
    var v15 float32 = -2.25
    var v16 int8 = 127
    v17 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: 42}
    v18 := GenS1{F0: -297, F1: 4294967295, F2: 42}
    v14 = append(v14, 127)
    fmt.Println(v9)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
