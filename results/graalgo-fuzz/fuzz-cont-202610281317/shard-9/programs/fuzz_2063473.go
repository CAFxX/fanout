package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0() int16 {
    var v6 float64 = 0.1
    v6 = ((v6 + v6) / v6)
    v7 := GenS1{F0: 24}
    v6 = ((1.5 * v6) - v6)
    v6 = (v6 - v6)
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{1}
    var v4 int = 21
    var v5 int32 = -37
    v8 := GenS2{F0: (false && ("a" >= "foo bar")), F1: 1}
    var v9 uint = 906
    var v10 uint = (v9 ^ uint(v2))
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int32 = v5
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS3{F0: -1, F1: 4294967295, F2: 0}
    v15 := GenS3{F0: 127, F1: 2, F2: 1}
    if true {
        v5 += (v5 | 2147483647)
    } else {
        fmt.Println(v5)
    }
    v1 = -2.25
    var v16 uint = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
