package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 1
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 30
    var v5 int32 = 1
    v6 := make(map[int]bool)
    v7 := GenS1{F0: (v5 << ((42) & 63)), F1: -1}
    var v8 uint32 = 42
    v9 := GenS1{F0: v5, F1: -517}
    v10 := GenS2{F0: (v5 << ((18446744073709551615) & 63)), F1: -802}
    var v11 int16 = -32768
    v12 := []int{32767, 1}
    v13 := GenS2{F0: (v5 * v5), F1: (-(v11))}
    var v14 uint32 = uint32(v3)
    v15 := GenS1{F0: (-(v5)), F1: -774}
    v11 += v11
    fmt.Println(v4, v14, v11)
    switch 321 {
    case 9223372036854775807:
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS3{F0: (!false)}
                fmt.Println(v3)
                _ = fi16
                _ = v17
            }
        }
    default:
        v11 -= (v11 * v11)
    }
    var v18 string = "hello"
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float32 = -2.25
        _ = fi19
        _ = v20
    }
    var v21 int16 = -399
    v22 := GenS4{F0: 2147483647, F1: false}
    var v23 uint32 = (^(v14))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
