package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64, p1 bool, p2 uint) int32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 1, F1: (fi6 + fi6), F2: fi6}
        _ = fi6
        _ = v7
    }
    if true {
        v8 := GenS1{F0: 834, F1: 42, F2: -495}
        switch 0 {
        case 42:
            var v9 uint16 = 332
            _ = v9
        case 0:
            v10 := GenS2{F0: 1, F1: 103, F2: 1000}
            _ = v10
        default:
            var v11 int32 = 657
            _ = v11
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return -527
}

func main() {
    var v1 int64 = 303
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 int = 128
    v5 := make(map[uint32]bool)
    delete(v2, 42)
    v4 -= v4
    v12 := GenS1{F0: 2147483647, F1: v4, F2: v4}
    v13 := GenS2{F0: 1, F1: 225, F2: -32768}
    var v14 uint8 = 42
    fmt.Println(v1, v14, v4)
    v15 := GenS3{F0: 42, F1: 2147483648}
    for fi16 := 0; fi16 < 13; fi16++ {
        v14 -= v14
        var v17 bool = (1 > 195)
        v18 := v5[907]
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v1 -= v1
        _ = fi19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint8 = 1
        if len(v3) > 0 {
            v3[0] = (fi20 * 127)
        }
        var v22 int16 = 1000
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int64 = (^(v1))
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 uint16 = 928
        v26 := v2[65535]
        var v27 uint = 65535
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint8 = v14
    for fi29 := 0; fi29 < 16; fi29++ {
        if (1e10 == float32(((float64(0.0) * 1.5) + -2.25))) {
            _ = len(v3)
        }
        var v30 uint32 = 496
        v2[42] = (v4 ^ int(v23))
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
