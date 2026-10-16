package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 bool, p1 string) uint16 {
    v4 := []int{-1, 75}
    var v5 uint16 = 527
    var v6 float32 = 1.5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int16) uint16 {
    var v7 int64 = 0
    var v8 int32 = -921
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []string{}
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 = v1
        v1 = v1
        v3 = append(v3, "")
        _ = fi9
    }
    fmt.Println(v1)
    switch "\n\t" {
    case "\n\t":
        var v10 bool = (0.0 > ((1e10 / -2.25) / float64(-2.25)))
        _ = v10
    case "foo bar":
        v11 := GenS1{F0: 4294967295}
        _ = v11
    case "hello":
        v1 = true
    }
    var v12 int = 941
    for fi13 := 0; fi13 < 4; fi13++ {
        if v1 {
            v14 := GenS1{F0: 0}
            _ = v14
        }
        _ = fi13
    }
    v2[(v12 * 4294967295)] = 0
    v15 := GenS2{F0: 10, F1: -907}
    fmt.Println(v1, v12)
    v16 := v2[-2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
