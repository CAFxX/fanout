package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 uint32, p1 int) int32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 704, F1: (fi6 >> ((18446744073709551615) & 63))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint32 = p0
        _ = fi8
        _ = v9
    }
    v10 := []int{p1, p1, 881}
    switch 4294967295 {
    case 128:
        var v11 int = 262
        _ = v11
    case 4294967295:
        if (!(!true)) {
            if true {
                var v12 uint8 = 255
                _ = v12
            }
        } else {
            var v13 uint8 = 255
            _ = v13
        }
    }
    v14 := GenS1{F0: -1, F1: (p1 | 32768)}
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return 138
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 255
    var v3 uint = 18446744073709551615
    var v4 float32 = 0.0
    var v5 int64 = 9223372036854775807
    v15 := GenS1{F0: 1, F1: 76}
    v16 := GenS2{F0: -60, F1: (v5 ^ (v5 / ((v5) | 1))), F2: 150}
    v17 := GenS1{F0: 1, F1: 0}
    v5 = 0
    fmt.Println(v4, v2, v1)
    var v18 string = (v1 + v1)
    v18 = v1
    var v19 int32 = 1
    v20 := []int{4294967295, 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
