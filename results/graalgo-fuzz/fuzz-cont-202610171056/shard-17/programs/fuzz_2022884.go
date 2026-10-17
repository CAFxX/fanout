package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    var v3 uint16 = 423
    v4 := make(map[int]bool)
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{fi6, (fi6 - (^(fi6)))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v5, v3, v2)
                _ = fi9
            }
            v2 -= v2
        }
        v10 := v1[936]
        _ = fi8
        _ = v10
    }
    v3 = 32767
    v3 = (v3 >> ((184) & 63))
    v5 = v5
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    v12 := []int{128}
    var v13 int64 = v2
    fmt.Println(v2, v3)
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint = 1
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v3 += 525
        v13 += v13
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: 0, F1: 0}
            v13 -= v2
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v13 -= v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
