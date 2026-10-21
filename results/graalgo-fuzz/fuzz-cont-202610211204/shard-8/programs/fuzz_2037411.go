package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    var v5 int = -1
    for fi6 := 0; fi6 < 7; fi6++ {
        v5 = fi6
        v5 = -580
        _ = fi6
    }
    _ = v5
    return 526
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 int = 20
    v4 := make(map[int]bool)
    v1 += (v1 % ((v1) | 1))
    var v7 uint64 = 0
    var v8 uint64 = v7
    for fi9 := 0; fi9 < 7; fi9++ {
        v3 += -1
        var v10 float64 = (3.14159 * 0.0)
        v10 += (0.0 * ((v10 / v10) - v10))
        _ = fi9
        _ = v10
    }
    if true {
        v3 = v3
        var v11 int = v3
        _ = v11
    } else {
        v12 := GenS1{F0: -413, F1: genFn0(), F2: 42}
        _ = v12
    }
    var v13 uint = 1
    if false {
        if true {
            var v14 int64 = 0
            _ = v14
        }
        delete(v2, uint32(v7))
        v15 := GenS1{F0: -1, F1: (v1 % (((^(v1))) | 1)), F2: 100}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := GenS1{F0: -32768, F1: (v1 & 1), F2: 127}
            var v18 int = v3
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v13)
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := []int{v3, (v3 * 127)}
        v1 += v1
        _ = fi19
        _ = v20
    }
    var v21 int = v3
    v13 += 255
    v21 = -91
    v3 = 191
    v22 := GenS2{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
