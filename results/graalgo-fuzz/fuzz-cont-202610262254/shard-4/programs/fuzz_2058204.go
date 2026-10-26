package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int64 = 42
        var v8 float64 = v2
        var v9 uint8 = 0
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{fi10, fi10}
            _ = fi11
            _ = v12
        }
        var v13 float32 = (1e10 / 1e10)
        v13 = v13
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: 0}
    if true {
        var v15 uint = 4294967295
        v5["foo bar"] = false
        _ = v15
    }
    v2 += v2
    switch 127 {
    case 42:
        var v16 int = -1
        _ = v16
    case 1:
        v2 = (float64(((((1.5 + 0.0) + 1e10) - 1e10) * 3.14159)) + v2)
    case 994:
        fmt.Println(v2)
    default:
        fmt.Println(v2)
    }
    var v17 uint = 205
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS2{F0: "a", F1: 0}
            v17 = v17
            var v21 string = "\n\t"
            _ = fi19
            _ = v20
            _ = v21
        }
        fmt.Println(v2, v17, v17)
        switch (fi18 << ((v17) & 63)) {
        case 32768:
            fmt.Println(fi18)
        default:
            var v22 int32 = 2147483647
            _ = v22
        }
        _ = fi18
    }
    if (65535 <= 4294967295) {
        delete(v3, 2)
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v17 += (v17 * (v17 << ((728) & 63)))
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
