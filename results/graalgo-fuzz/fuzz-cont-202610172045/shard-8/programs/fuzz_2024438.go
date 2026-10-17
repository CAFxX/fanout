package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 int16, p2 uint16) uint8 {
    v6 := GenS1{F0: 947}
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint32, p1 int16) uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    var v9 int16 = (^(p1))
    _ = p0
    _ = p1
    _ = v9
    return 593
}

func main() {
    v1 := []int{-675, -820, 128, 4294967295}
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 uint = 309
    var v5 uint32 = 926
    v5 += (v5 % (((v5 ^ v5)) | 1))
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 string = "foo bar"
        for fi12 := 0; fi12 < 3; fi12++ {
            v4 = v4
            v5 = (v5 * 862)
            _ = fi12
        }
        var v13 float32 = float32(1e10)
        _ = fi10
        _ = v11
        _ = v13
    }
    if true {
        v14 := GenS2{F0: 476}
        var v15 uint8 = 1
        _ = v14
        _ = v15
    } else {
        v3["foo bar"] = (!true)
        var v16 uint64 = 173
        _ = v16
    }
    var v17 uint32 = v5
    v18 := GenS2{F0: 42}
    fmt.Println(v5, v5, v5)
    if true {
        for fi19 := 0; fi19 < 5; fi19++ {
            if true {
                v17 = v17
            }
            var v20 int64 = 932
            _ = fi19
            _ = v20
        }
        v4 += v4
        fmt.Println(v17, v4, v17)
    } else {
        var v21 uint8 = 165
        _ = v21
    }
    v17 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
