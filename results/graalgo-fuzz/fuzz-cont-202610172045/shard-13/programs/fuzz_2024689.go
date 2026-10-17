package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 int8, p1 uint16, p2 int64) uint64 {
    v6 := []int{1, -177, 2147483647, 32767}
    v7 := GenS1{F0: 0, F1: p1}
    p2 = p2
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 4294967295
    var v3 int16 = 32767
    var v4 float32 = (1.5 - (0.0 - (3.14159 / float32(3.14159))))
    v5 := []int{v2, v2, v2}
    for fi9 := 0; fi9 < 7; fi9++ {
        v3 = (v3 | (v3 - v3))
        v2 = 42
        switch "\n\t" {
        case "":
            fmt.Println(v4, v2, v3)
        case "hello":
            v4 = v4
        case "\n\t":
            if (((-546 < 577) || true) && (!(821 == 0))) {
                var v10 int16 = v3
                v10 = v3
                v11 := v1["a"]
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: 4294967295, F1: 255}
                    _ = fi12
                    _ = v13
                }
                var v14 int8 = -117
                if true {
                    if (true || ((false || false) || (genFn0(v14, 255, -877) == 1))) {
                        v14 = (v14 / ((v14) | 1))
                        _ = len(v1)
                        var v15 bool = false
                        _ = v15
                    }
                }
                _ = v14
            }
        }
        _ = fi9
    }
    switch v2 {
    case -27:
        var v16 int32 = -1
        _ = v16
    default:
        var v17 float64 = (1.5 / -2.25)
        _ = v17
    }
    v18 := GenS2{F0: 4294967295, F1: -1, F2: 1}
    v2 -= v2
    _ = len(v1)
    v4 -= float32(3.14159)
    var v19 int32 = 728
    v20 := GenS3{F0: genFn0(42, 255, 443), F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
