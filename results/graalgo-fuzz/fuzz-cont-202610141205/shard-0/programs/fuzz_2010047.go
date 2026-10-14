package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := []uint32{574, 42, 362}
    var v2 string = "\n\t"
    var v3 uint = 65535
    v4 := []uint32{4294967295, 4294967295}
    if true {
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS1{F0: v2, F1: 180}
            var v7 uint = 0
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 2; fi9++ {
        switch (fi9 >> ((v3) & 63)) {
        case -842:
            v3 = v3
        default:
            v2 = (v2 + ("foo bar" + v2))
        }
        var v10 uint = (v3 % ((v3) | 1))
        _ = fi9
        _ = v10
    }
    if len(v1) > 0 {
        v1[0] = (v8 ^ v8)
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: v2, F1: v8}
        _ = len(v4)
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 669, F1: -718, F2: -174}
    var v14 int32 = -1
    var v15 uint16 = 2
    var v16 int16 = 0
    var v17 float32 = float32(1.5)
    v3 += 889
    for fi18 := 0; fi18 < 8; fi18++ {
        if true {
            v19 := GenS1{F0: v2, F1: (^(v8))}
            _ = v19
        }
        _ = fi18
    }
    v15 = 432
    v20 := GenS2{F0: -186, F1: int64(v14), F2: 629}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
