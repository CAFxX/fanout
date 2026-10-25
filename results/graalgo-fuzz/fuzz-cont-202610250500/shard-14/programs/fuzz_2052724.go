package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 uint32) uint8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        p1 = p1
        var v5 string = ""
        fmt.Println(p1, p2, p0)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int = (fi6 + fi6)
        p1 = p1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 222
}

func main() {
    var v1 int = -415
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v1 += 9223372036854775807
    _ = len(v3)
    v1 -= 1
    v1 += (-(v1))
    v1 -= v1
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v1 = v1
            var v10 uint64 = 4294967295
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if true {
        v1 = (v1 - v1)
        v2 = append(v2, 680)
    }
    if false {
        var v11 int = (v1 ^ v1)
        v11 -= v11
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            _ = len(v2)
            v11 = v11
            if len(v2) > 0 {
                v2[0] = 243
            }
        }
        _ = v11
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = 1
            _ = len(v2)
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 692
        _ = v14
    }
    var v15 uint = 51
    v16 := GenS1{F0: 1, F1: 835}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
