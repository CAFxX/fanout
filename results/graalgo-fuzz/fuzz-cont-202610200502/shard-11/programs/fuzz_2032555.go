package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint8, p1 int64) uint64 {
    v4 := []int{1}
    if ((910 == 32768) && false) {
        v5 := GenS1{F0: -81, F1: 284, F2: 1}
        _ = v5
    }
    fmt.Println(p0)
    _ = len(v4)
    p0 = 1
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    v3 := make(map[uint32]string)
    var v6 int8 = -116
    switch 32768 {
    case -796:
        v6 += (^(v6))
    }
    v7 := GenS1{F0: (v6 >> ((309) & 63)), F1: -232, F2: 95}
    if v2 {
        fmt.Println(v2)
        delete(v1, -197)
    } else {
        fmt.Println(v2, v6)
    }
    v6 -= (v6 - 127)
    v8 := GenS2{F0: ((v2 && false) && false), F1: 874}
    var v9 int = 42
    v6 = v6
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS3{F0: 255, F1: 1}
        _ = fi10
        _ = v11
    }
    var v12 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
