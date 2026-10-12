package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 bool, p1 int32) int32 {
    v7 := GenS1{F0: 255}
    v8 := []int{-2147483648, 32767, 2147483647, 92}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: 1}
        v11 := GenS1{F0: 255}
        _ = len(v8)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    v3 := []int{2147483647, -775, -902, -2147483648}
    var v4 uint64 = 516
    var v5 uint16 = 1
    v6 := make(map[int]uint32)
    switch "\n\t" {
    case "hello":
        v3 = append(v3, 2147483647)
    }
    fmt.Println(v5, v5, v2)
    var v12 uint64 = 2
    var v13 int16 = 32767
    var v14 float64 = -2.25
    fmt.Println(v5)
    for fi15 := 0; fi15 < 16; fi15++ {
        switch "a" {
        case "":
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 float64 = float64(3.14159)
                var v18 float64 = v17
                var v19 uint32 = 0
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int = fi20
        v12 -= 42
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
