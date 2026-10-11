package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 bool = true
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 2147483647, F1: -128}
    v11 := GenS2{F0: 32767, F1: 0, F2: -32768}
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := []int{-309, 127}
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint8 = 120
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    v4 += (v4 * v4)
    var v12 float32 = 0.0
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: 65535, F1: 42}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 int32 = -548
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: -511, F1: (!(((1.5 * 3.14159) / (1.5 + (1.5 + 1.5))) != ((-2.25 - (1.5 / (1.5 * 0.1))) + 1e10))), F2: ("" + "foo bar")}
    var v17 uint16 = 42
    var v18 float32 = v12
    for fi19 := 0; fi19 < 2; fi19++ {
        if true {
            fmt.Println(v18, v4, v18)
            v12 += v12
        } else {
            var v20 uint32 = uint32(v17)
            _ = v20
        }
        v18 += v18
        _ = fi19
    }
    var v21 uint32 = 42
    v4 += 0
    v22 := GenS1{F0: v21, F1: 21}
    v23 := GenS4{F0: 1, F1: 128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
