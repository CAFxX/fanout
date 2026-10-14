package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint, p2 int32) uint8 {
    var v5 uint16 = 255
    p1 = p1
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 << ((p0) & 63))}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := []string{"a", "\n\t", "foo bar"}
    var v2 int32 = 0
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    if false {
        v8 := GenS1{F0: 267, F1: 42, F2: 4294967295}
        v2 += v2
        fmt.Println(v4, v4)
        _ = v8
    }
    var v9 float32 = 1.5
    switch "a" {
    case "":
        v4 = (v4 >> ((0) & 63))
    default:
        v2 += (v2 & (v2 & 0))
    }
    v2 = v2
    var v10 uint = 0
    delete(v3, 2147483648)
    var v11 uint = 42
    if false {
        var v12 uint = (^(v10))
        v1 = append(v1, "foo bar")
        v13 := GenS2{F0: 42}
        _ = v12
        _ = v13
    }
    var v14 float64 = (-2.25 + (1e10 + 1.5))
    v15 := GenS3{F0: v4}
    var v16 float32 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
