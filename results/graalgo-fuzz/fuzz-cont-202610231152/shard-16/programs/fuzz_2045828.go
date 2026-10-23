package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int32 = 1
    v2 := []string{("" + "\n\t"), "\n\t"}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 187
        var v9 bool = true
        var v10 int8 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v2 = append(v2, ("hello" + "\n\t"))
    v11 := GenS2{F0: 42}
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(fi12, v1)
        fmt.Println(v1, v1)
        v4[65535] = 789
        _ = fi12
    }
    v1 = (v1 - 1)
    delete(v3, 65535)
    if (!true) {
        var v13 bool = (2 >= 0)
        _ = v13
    }
    delete(v3, 4294967295)
    v14 := GenS3{F0: false, F1: 451, F2: 255}
    fmt.Println(v1)
    v1 = (v1 + 42)
    v15 := GenS4{F0: 732}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
