package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := []string{}
    v4 := make(map[uint32]int)
    var v5 int32 = 0
    v6 := []uint32{4294967295, 65535}
    var v7 float64 = 1e10
    if len(v6) > 0 {
        v6[0] = 2
    }
    var v8 int8 = -128
    var v9 int16 = -1
    delete(v1, 709)
    if (-1 == v8) {
        v2[42] = -1
    }
    v10 := GenS1{F0: (-(v8))}
    switch "\n\t" {
    case "a":
        v11 := GenS1{F0: (v8 | 89)}
        _ = v11
    case "hello":
        if ((402 == 411) || ((!true) && (0.0 == (1.5 - 0.0)))) {
            for fi12 := 0; fi12 < 8; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint16 = 2
                    _ = fi13
                    _ = v14
                }
                v5 = v5
                var v15 float32 = 1.5
                _ = fi12
                _ = v15
            }
            var v16 uint16 = 1
            _ = v16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
