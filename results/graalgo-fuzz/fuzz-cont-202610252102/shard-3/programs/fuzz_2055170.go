package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint }
func genFn0(p0 int8, p1 int32, p2 uint64) int32 {
    p1 -= p1
    var v4 int32 = p1
    p0 = (p0 >> ((42) & 63))
    var v5 uint = uint(p2)
    fmt.Println(p2, p0, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (^(p1))
}

func main() {
    var v1 int32 = 0
    var v2 int64 = 875
    v3 := make(map[int]int)
    v6 := GenS1{F0: 1, F1: v1, F2: 127}
    v7 := GenS1{F0: 65535, F1: v1, F2: 2}
    v8 := GenS1{F0: 998, F1: (-(v1)), F2: 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (("a" + "\n\t") + "hello") {
        case "hello":
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint = 65535
                v12 := v3[(fi10 / ((fi10) | 1))]
                _ = fi10
                _ = v11
                _ = v12
            }
        case "foo bar":
            var v13 int8 = 127
            _ = v13
        default:
            fmt.Println(v2, fi9)
        }
        _ = fi9
    }
    v14 := []int{1, 32768, 41}
    v1 -= genFn0(0, v1, 0)
    if len(v14) > 0 {
        v14[0] = 127
    }
    fmt.Println(v1, v2)
    v15 := GenS2{F0: false, F1: 33}
    var v16 uint16 = 662
    var v17 bool = (false || (!false))
    var v18 uint8 = 2
    var v19 float64 = 0.1
    var v20 uint32 = 661
    var v21 float32 = (1e10 - 0.1)
    switch "hello" {
    case "":
        fmt.Println(v19, v1)
    case "foo bar":
        v16 = uint16(v20)
    case "hello":
        v3[32768] = 4294967295
    default:
        v22 := GenS3{F0: "foo bar", F1: genFn0(100, (v1 + (v1 & int32(v18))), 2), F2: 0}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
