package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }
func genFn0(p0 bool, p1 int) int {
    var v6 int8 = -4
    v7 := []int{-24, -628}
    fmt.Println(p1)
    for fi8 := 0; fi8 < 8; fi8++ {
        p1 -= fi8
        v9 := GenS1{F0: false, F1: 42, F2: (fi8 / (((p1 + fi8)) | 1))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 / ((229) | 1))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    var v10 int16 = -1
    fmt.Println(v10)
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int32 = 42
            var v14 uint32 = 908
            v10 = -1
            v13 = (v13 / ((v13) | 1))
            v14 += 2147483647
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 int32 = 2147483647
    v16 := v1["foo bar"]
    var v17 float64 = ((1e10 - -2.25) * -2.25)
    v10 = (v10 >> ((0) & 63))
    var v18 string = (("a" + (("hello" + ("" + ("hello" + ""))) + "foo bar")) + "foo bar")
    var v19 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
