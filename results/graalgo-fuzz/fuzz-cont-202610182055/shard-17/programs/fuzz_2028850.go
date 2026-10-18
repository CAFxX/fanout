package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 19; fi7++ {
        v2 -= v2
        v8 := []int{(fi7 - (fi7 | fi7))}
        _ = fi7
        _ = v8
    }
    v2 += (v2 * v2)
    var v9 int32 = -1
    for fi10 := 0; fi10 < 15; fi10++ {
        v2 = v2
        var v11 float32 = 1e10
        var v12 int64 = v6
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int16 = 52
    v14 := []int{}
    v15 := GenS1{F0: v2}
    fmt.Println(v6)
    _ = len(v1)
    var v16 uint32 = 804
    for fi17 := 0; fi17 < 17; fi17++ {
        v2 = (-(v2))
        var v18 int64 = v6
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi17
        _ = v18
    }
    v13 = (^(v13))
    if len(v14) > 0 {
        v14[0] = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
