package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64, p1 int32, p2 int16) int8 {
    var v5 int8 = 42
    var v6 uint = 255
    p1 = (p1 | (p1 << ((v6) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int8, p1 int32) uint64 {
    if true {
        v7 := GenS1{F0: 140}
        var v8 float64 = (3.14159 / 0.1)
        _ = v7
        _ = v8
    } else {
        p1 += (p1 | p1)
        if false {
            p1 = (p1 - p1)
            var v9 uint = 18446744073709551615
            if false {
                var v10 int8 = (p0 << ((42) & 63))
                _ = v10
            } else {
                v11 := []int{871, 832, -2147483648}
                v9 = v9
                _ = v11
            }
            _ = v9
        }
    }
    switch "foo bar" {
    case "a":
        var v12 int32 = (p1 - 42)
        _ = v12
    }
    fmt.Println(p0, p0)
    var v13 int32 = -245
    p1 -= (v13 + (v13 - (p1 & -2147483648)))
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 bool = (!false)
    v2 := []int{1}
    v3 := []int{127, -144, -751}
    v4 := make(map[uint32]string)
    fmt.Println(v1)
    v14 := GenS1{F0: 937}
    v1 = v1
    if v1 {
        v1 = true
        v15 := GenS2{F0: 0, F1: 4294967295}
        if (-2.25 > float32(0.0)) {
            v3 = append(v3, -470)
            v16 := GenS1{F0: 2}
            _ = v16
        }
        _ = v15
    }
    fmt.Println(v1)
    var v17 uint8 = 1
    fmt.Println(v1, v17)
    var v18 float64 = (0.0 / (3.14159 / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
