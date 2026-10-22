package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 uint16 = 255
    v4 := make(map[string]bool)
    v5 := GenS1{F0: (0 == 486), F1: v3}
    if ("hello" >= "a") {
        fmt.Println(v2, v2)
        var v6 int16 = 1000
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint8 = 2
            v1[353] = fi7
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 uint64 = 0
    var v10 float64 = -2.25
    v11 := GenS2{F0: 42, F1: ("hello" + "\n\t"), F2: 1}
    v10 -= (v10 / 1.5)
    if (v2 || (v2 || v2)) {
        v12 := []int{int(v9), 42}
        var v13 uint16 = uint16(v9)
        _ = v12
        _ = v13
    }
    if v2 {
        v14 := []int{}
        var v15 uint64 = v9
        _ = v14
        _ = v15
    }
    delete(v4, ("\n\t" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
