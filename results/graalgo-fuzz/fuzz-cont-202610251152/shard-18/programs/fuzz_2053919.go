package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 string) uint32 {
    v7 := GenS1{F0: 1, F1: 917, F2: -1}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 bool = (169 < 0)
    v4 := []uint32{880, 310, 319, 2}
    v5 := []int{}
    v6 := make(map[uint32]uint32)
    v2 = ((3.14159 > 1e10) || (0 == 1))
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int64 = 1
        switch (fi9 << ((uint(v8)) & 63)) {
        case -1:
            v10 += (-(v10))
        default:
            var v11 uint64 = 1
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v3 = (1e10 < ((1e10 + 1.5) / 1.0))
    var v12 uint = 255
    v13 := v1[("a" + "a")]
    if (float32(-2.25) == 0.0) {
        var v14 uint64 = 107
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v3 = (842 > 0)
        var v16 string = "hello"
        switch fi15 {
        case 127:
            v1[(v16 + "hello")] = v13
        default:
            var v17 uint16 = 32767
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v13 += (v13 + v13)
    if true {
        var v18 uint32 = (v13 << ((v12) & 63))
        _ = v18
    }
    v4 = append(v4, v13)
    v19 := GenS2{F0: (^(v13)), F1: 0}
    v12 += (^(v12))
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int = (fi20 | fi20)
        _ = fi20
        _ = v21
    }
    switch ("hello" + "hello") {
    case "\n\t":
        v2 = (v3 || true)
    case "foo bar":
        v2 = true
    default:
        var v22 uint = 1
        _ = v22
    }
    v23 := GenS3{F0: 4294967295, F1: 695}
    for fi24 := 0; fi24 < 14; fi24++ {
        if ((uint16(v13) >= uint16(v13)) && v3) {
            var v25 bool = v3
            v1["a"] = v13
            _ = v25
        }
        v2 = v3
        v26 := GenS2{F0: genFn0(""), F1: 42}
        _ = fi24
        _ = v26
    }
    var v27 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
