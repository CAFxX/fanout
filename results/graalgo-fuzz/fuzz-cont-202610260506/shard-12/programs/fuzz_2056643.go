package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 680
    var v3 float32 = float32(1.5)
    fmt.Println(v2, v2)
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 uint32 = 1
            v6 += (v6 % ((v6) | 1))
            v6 = 2
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 int32 = 42
    v8 := GenS1{F0: (((^(v2)) < (v2 << ((2) & 63))) && true), F1: 406, F2: v7}
    v1[1] = 310
    var v9 int = 32767
    v10 := GenS1{F0: false, F1: 42, F2: v7}
    v3 = float32((-2.25 * (-2.25 / float64(v3))))
    v9 = int(v7)
    var v11 uint64 = 0
    var v12 int = (v9 | 247)
    var v13 uint8 = 0
    var v14 uint16 = 0
    if true {
        var v15 int32 = (v7 ^ (v7 + 42))
        v14 -= v14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int16 = 1
        v9 = v12
        _ = fi16
        _ = v17
    }
    v18 := []int{v9}
    switch 9223372036854775807 {
    case 32768:
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: true, F1: 372, F2: v7}
            v13 -= (^(v13))
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 int16 = -32768
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
