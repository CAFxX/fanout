package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -763
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: -487}
    v5 := []int{}
    if true {
        var v6 int64 = 0
        var v7 int16 = 1000
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -32768}
    if true {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS2{F0: 0}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 float32 = 1e10
            _ = fi11
            _ = v12
        }
    }
    if (2 == 0) {
        var v13 uint8 = 28
        fmt.Println(v2)
        _ = v13
    }
    var v14 uint32 = 65535
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int64 = 9223372036854775807
        var v17 string = ("a" + "\n\t")
        v5 = append(v5, fi15)
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 += v2
    v5 = append(v5, 451)
    var v18 int16 = 180
    fmt.Println(v2, v14, v2)
    v19 := GenS2{F0: -1}
    v14 -= v14
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float32 = 1e10
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
