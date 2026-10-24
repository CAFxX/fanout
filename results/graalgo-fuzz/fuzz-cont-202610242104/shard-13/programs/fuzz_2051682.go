package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0() int8 {
    if true {
        if false {
            var v5 float32 = 0.0
            _ = v5
        } else {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 float64 = 0.0
                v8 := []int{}
                _ = fi6
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint64 = 4294967295
                _ = fi9
                _ = v10
            }
        }
        var v11 float32 = ((1.5 + (0.0 - (0.1 * 0.1))) + 1e10)
        _ = v11
    } else {
        var v12 uint64 = 864
        v13 := GenS1{F0: -263, F1: 0, F2: 1}
        var v14 int = -2147483648
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: -32768, F1: 1, F2: 0}
    _ = v15
    return 0
}

func main() {
    v1 := []int{2147483647, -1, 1, 2147483647}
    v2 := []int{4294967295, 128}
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    fmt.Println(v4)
    v16 := GenS1{F0: 42, F1: 2, F2: 781}
    _ = len(v3)
    if (774 != 42) {
        var v17 string = "a"
        v18 := GenS1{F0: -1, F1: 127, F2: 0}
        for fi19 := 0; fi19 < 6; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int64 = (v4 & v4)
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint64 = 0
                v24 := GenS1{F0: 42, F1: 0, F2: v23}
                var v25 int8 = 127
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    v4 += v4
    var v26 uint64 = 599
    fmt.Println(v26)
    var v27 string = "a"
    v27 = (v27 + (((v27 + (((v27 + v27) + ("foo bar" + v27)) + v27)) + (v27 + ("foo bar" + v27))) + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
