package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool }

func main() {
    v1 := []uint32{431}
    v2 := make(map[int]bool)
    v3 := []int{}
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    v6 := []int{777, 2147483647, 589}
    if (42 == 4294967295) {
        var v7 uint64 = 2
        _ = v7
    } else {
        if false {
            var v8 uint64 = 1
            var v9 float32 = ((1e10 * (float32(-2.25) - (1e10 * 0.1))) - (1.5 / (3.14159 + 1e10)))
            v10 := GenS1{F0: -198, F1: uint(v8)}
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                if (1 > -24) {
                    switch fi11 {
                    case 2147483647:
                        var v12 uint64 = 0
                        _ = v12
                    }
                }
                var v13 bool = (524 >= 230)
                _ = fi11
                _ = v13
            }
        }
        var v14 uint64 = 0
        _ = v14
    }
    var v15 uint16 = 32767
    v15 = v15
    v15 += v15
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v15, fi16, fi16)
            _ = fi17
        }
        var v18 float32 = 3.14159
        _ = fi16
        _ = v18
    }
    v15 -= v15
    v19 := GenS2{F0: true}
    var v20 int32 = int32(v15)
    var v21 int8 = -1
    fmt.Println(v15, v20, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
