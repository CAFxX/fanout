package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 bool) uint8 {
    if (1000 == -1) {
        p0 = p0
    }
    fmt.Println(p0, p0)
    v7 := []int{-1}
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 bool = true
    var v6 bool = (v5 && (802 <= 4294967295))
    var v8 uint64 = 42
    switch 32767 {
    case 42:
        var v9 int16 = 1000
        _ = v9
    case 0:
        v10 := []int{32767, 9223372036854775807, -1}
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: (2 == 83)}
    v8 -= v8
    var v12 int8 = 127
    var v13 uint = 4294967295
    if false {
        switch "\n\t" {
        case "\n\t":
            if (!(!(1 == 593))) {
                _ = len(v1)
                var v14 float32 = (3.14159 * float32(float64(0.1)))
                v13 = v13
                _ = v14
            }
        case "foo bar":
            v13 = 2
        case "hello":
            v5 = ((-(v12)) > (v12 >> ((v13) & 63)))
        }
        fmt.Println(v6)
    }
    v15 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
