package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int32; F2 string }
func genFn0() bool {
    var v6 uint8 = 2
    v7 := []int{0}
    if false {
        var v8 string = "foo bar"
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 float64 = (3.14159 + -2.25)
            v6 = v6
            if true {
                var v11 uint16 = 65535
                var v12 int8 = -27
                v13 := GenS1{F0: 967, F1: 42, F2: (fi9 / ((968) | 1))}
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        v14 := GenS2{F0: "a"}
        _ = v14
        _ = v8
    }
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = 1e10
    v5 := make(map[uint32]string)
    fmt.Println(v4)
    var v15 float64 = 3.14159
    v16 := []int{821, 128, -122}
    var v17 uint8 = 2
    var v18 float32 = 0.1
    v4 += -2.25
    v15 -= (float64(v4) - ((v15 * -2.25) + v15))
    var v19 uint16 = 42
    _ = len(v16)
    fmt.Println(v19, v18)
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS2{F0: "\n\t"}
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 float32 = v18
            v4 -= 0.0
            _ = fi22
            _ = v23
        }
        v24 := v5[2147483647]
        _ = fi20
        _ = v21
        _ = v24
    }
    v25 := GenS3{F0: 42, F1: 124, F2: "a"}
    v4 = v4
    fmt.Println(v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
