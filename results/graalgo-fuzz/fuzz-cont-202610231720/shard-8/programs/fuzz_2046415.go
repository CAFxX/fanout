package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }

func main() {
    var v1 uint16 = 1
    v2 := []string{("\n\t" + (("\n\t" + "hello") + "foo bar")), ("\n\t" + "hello"), "\n\t"}
    v3 := make(map[uint32]string)
    var v4 bool = (6 > 267)
    var v5 float32 = 1e10
    v4 = v4
    v6 := GenS1{F0: (!(!(v4 && (!(v4 && v4))))), F1: -238, F2: 129}
    v1 -= (v1 >> ((42) & 63))
    var v7 float64 = ((-2.25 + float64(v5)) + -2.25)
    v4 = v4
    v4 = (!false)
    if (42 > 0) {
        v5 = float32(v7)
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint64 = 128
            _ = fi8
            _ = v9
        }
    }
    _ = len(v3)
    var v10 uint64 = 2
    v10 += 84
    v11 := GenS1{F0: (int(v10) > 980), F1: 32768, F2: 286}
    v12 := GenS1{F0: v4, F1: 4294967295, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
