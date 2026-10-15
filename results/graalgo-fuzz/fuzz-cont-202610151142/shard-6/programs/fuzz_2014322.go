package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    v4 := make(map[uint32]bool)
    v5 := []uint32{993, 0, 839, 259}
    v3 = "hello"
    fmt.Println(v3, v3)
    v3 = "hello"
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 float32 = float32(1e10)
            var v9 uint64 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            v4[581] = false
            _ = fi10
        }
        v11 := GenS1{F0: (^(fi6)), F1: v1}
        _ = fi6
        _ = v11
    }
    v12 := GenS2{F0: 42}
    _ = len(v4)
    var v13 uint64 = 0
    if false {
        v1 = 1
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := v2["\n\t"]
            _ = fi14
            _ = v15
        }
        fmt.Println(v3, v1)
    }
    v16 := GenS2{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
