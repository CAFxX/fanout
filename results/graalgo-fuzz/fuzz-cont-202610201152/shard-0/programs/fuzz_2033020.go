package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := []string{"a", "foo bar", "foo bar", "a"}
    v2 := make(map[uint32]int)
    var v3 float64 = 3.14159
    v4 := []uint32{701, 13, 1}
    var v5 int16 = 691
    v6 := GenS1{F0: 2, F1: 1, F2: 722}
    var v7 uint16 = 32767
    var v8 uint8 = 255
    v9 := GenS1{F0: uint16(v8), F1: 255, F2: -274}
    v8 = v8
    v3 = float64(1.5)
    if false {
        var v10 float64 = v3
        var v11 int32 = 995
        _ = v10
        _ = v11
    } else {
        v5 -= v5
        v7 = 32767
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int64 = 0
        if (v3 >= 0.0) {
            var v14 uint16 = (v7 & v7)
            _ = v14
        }
        var v15 uint64 = 4294967295
        _ = fi12
        _ = v13
        _ = v15
    }
    var v16 uint8 = (^(v8))
    var v17 float32 = -2.25
    v8 += (v16 | (v16 << ((895) & 63)))
    var v18 uint = 170
    v5 = v5
    var v19 int64 = 42
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 string = "hello"
            if true {
                _ = len(v1)
                v22 = "a"
            } else {
                var v23 uint64 = 234
                _ = v23
            }
            switch "\n\t" {
            case "a":
                v24 := GenS1{F0: (v7 & v7), F1: v18, F2: fi20}
                _ = v24
            default:
                v25 := GenS1{F0: v7, F1: (^(v18)), F2: (fi21 & fi21)}
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v26 uint16 = (v7 / ((0) | 1))
    v27 := GenS2{F0: 2147483647, F1: (v16 + (v16 - v8))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
