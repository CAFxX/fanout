package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64) bool {
    v6 := GenS1{F0: 32767, F1: 0, F2: true}
    p0 = (p0 * p0)
    var v7 float64 = 3.14159
    v8 := GenS2{F0: 4294967295, F1: -1}
    v7 = 0.1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return false
}
func genFn1() uint16 {
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int32 = -1
        v10 += (v10 + (-(v10)))
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{fi11, fi9, -2147483648, (fi9 | 32768)}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if true {
        if true {
            var v13 float32 = 0.0
            _ = v13
        }
        var v14 int64 = 1
        _ = v14
    }
    return 32767
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 int16 = -151
    var v5 int64 = -112
    v15 := GenS1{F0: -764, F1: 32767, F2: ((65535 > 2) && (!false))}
    if (false && true) {
        v16 := GenS1{F0: 127, F1: 1, F2: (true && (true || (false || true)))}
        _ = v16
    }
    v17 := GenS3{F0: v1}
    v18 := GenS4{F0: (v1 ^ 4294967295)}
    v19 := []int{2147483647, 1, 32767}
    switch 9223372036854775807 {
    case -2147483648:
        v20 := GenS1{F0: -77, F1: 255, F2: false}
        _ = v20
    case 63:
        _ = len(v2)
    default:
        v1 += v1
    }
    v1 = 1
    var v21 int64 = v5
    var v22 uint32 = 1
    v4 = (v4 + 632)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
