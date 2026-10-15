package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    v3 := []int{32767, 0}
    var v4 bool = (true || (449 == 0))
    var v5 uint32 = 248
    v6 := make(map[uint32]int)
    v7 := GenS1{F0: 2, F1: 679, F2: -944}
    v8 := GenS2{F0: 65535, F1: v1}
    if true {
        fmt.Println(v4)
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v3 = append(v3, 32768)
        _ = fi9
    }
    _ = len(v2)
    v5 = v5
    if v4 {
        v10 := GenS1{F0: 697, F1: 1, F2: -1}
        v11 := GenS2{F0: 4294967295, F1: (v1 & v1)}
        fmt.Println(v5, v5, v1)
        _ = v10
        _ = v11
    }
    var v12 float32 = float32(1e10)
    fmt.Println(v12, v12, v1)
    v13 := GenS1{F0: 2, F1: -1, F2: 1}
    var v14 int = 0
    if (((1e10 * (0.1 - 3.14159)) * 3.14159) < 1e10) {
        v2[2147483648] = v14
    } else {
        if (0 > 32767) {
            v15 := GenS1{F0: 65535, F1: -2147483648, F2: -767}
            _ = v15
        } else {
            v14 = v14
        }
        v16 := GenS1{F0: 632, F1: 601, F2: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: 595, F1: (v1 & (^(v1)))}
        var v19 int8 = -128
        _ = fi17
        _ = v18
        _ = v19
    }
    v12 = (float32(1e10) / 1e10)
    var v20 float32 = ((v12 - float32(((1e10 + -2.25) - 0.0))) / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
