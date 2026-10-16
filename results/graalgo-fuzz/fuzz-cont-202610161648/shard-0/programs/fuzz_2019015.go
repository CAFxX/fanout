package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0() uint8 {
    var v4 float32 = 3.14159
    v5 := GenS1{F0: 65535, F1: -124}
    var v6 uint16 = 65535
    _ = v4
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 18446744073709551615
    v3 := []int{}
    var v7 uint16 = 663
    if len(v3) > 0 {
        v3[0] = 128
    }
    delete(v1, 1)
    if false {
        var v8 float32 = 0.1
        _ = v8
    }
    v7 = 2
    for fi9 := 0; fi9 < 19; fi9++ {
        v7 -= 833
        var v10 uint = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi9
        _ = v10
    }
    if (!false) {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint = (v2 * (v2 - v2))
            v13 := GenS1{F0: 758, F1: 127}
            _ = fi11
            _ = v12
            _ = v13
        }
        if false {
            v2 += (v2 >> ((v2) & 63))
        }
    }
    v2 += (v2 & v2)
    fmt.Println(v2)
    v7 += 32767
    v2 = (v2 << (((v2 >> ((949) & 63))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
