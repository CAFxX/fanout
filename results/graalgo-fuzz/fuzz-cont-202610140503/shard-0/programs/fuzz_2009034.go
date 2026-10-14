package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 886
    var v3 uint32 = v2
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: -966, F1: 65535, F2: 751}
            _ = fi7
            _ = v8
        }
        v9 := []int{fi6}
        fmt.Println(fi6)
        _ = fi6
        _ = v9
    }
    v1 = (v1 ^ (v1 * v1))
    var v10 uint64 = 42
    var v11 bool = ((false && false) || true)
    switch "a" {
    case "a":
        v12 := GenS1{F0: 780, F1: 237, F2: 1}
        _ = v12
    case "foo bar":
        fmt.Println(v11)
    case "\n\t":
        v1 = (v1 ^ uint8(v3))
    }
    v13 := GenS2{F0: 1, F1: -895, F2: 2}
    var v14 string = "\n\t"
    v10 -= 42
    var v15 int32 = -1
    fmt.Println(v15)
    for fi16 := 0; fi16 < 14; fi16++ {
        if v11 {
            var v17 float32 = 1.5
            v17 += v17
            _ = v17
        } else {
            v11 = v11
        }
        var v18 uint32 = (v3 + 0)
        v2 -= v18
        _ = fi16
        _ = v18
    }
    v19 := []int{-11, 42, 127}
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            if ((100 < 0) || false) {
                v11 = ((!v11) || (false && true))
                if v11 {
                    v22 := v4[32767]
                    _ = v22
                }
                v1 = v1
            }
            _ = fi21
        }
        fmt.Println(v2, v2, v3)
        v10 = uint64(v3)
        _ = fi20
    }
    if len(v19) > 0 {
        v19[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
