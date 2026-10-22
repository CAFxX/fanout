package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 int = -142
    _ = len(v2)
    var v5 uint16 = 255
    fmt.Println(v5, v5)
    fmt.Println(v5)
    v5 += 0
    var v6 uint64 = uint64(v5)
    var v7 int = -742
    v8 := GenS1{F0: 42}
    var v9 int32 = 42
    for fi10 := 0; fi10 < 12; fi10++ {
        _ = len(v3)
        _ = fi10
    }
    var v11 string = (("\n\t" + "\n\t") + ("" + "\n\t"))
    for fi12 := 0; fi12 < 12; fi12++ {
        v9 = (v9 + (v9 >> ((4294967295) & 63)))
        var v13 float32 = (float32(((0.0 + -2.25) * 1e10)) - (-2.25 / 3.14159))
        v14 := GenS1{F0: 204}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
