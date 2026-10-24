package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint8) uint {
    v4 := GenS1{F0: 65535}
    v5 := GenS1{F0: 42}
    var v6 float64 = 3.14159
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int16, p1 string) int16 {
    v7 := GenS2{F0: 42, F1: 841}
    var v8 float64 = 0.0
    var v9 bool = false
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 394
    var v2 uint16 = v1
    v3 := make(map[int]uint32)
    v1 = (v2 | 549)
    v1 -= (v1 + (v1 / (((^(v1))) | 1)))
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if (true && (76 < -53)) {
                var v12 int64 = -1
                v2 = v2
                var v13 uint32 = 4294967295
                _ = v12
                _ = v13
            }
            if true {
                v14 := []int{(fi10 | 9223372036854775807), (fi10 | 4294967295)}
                _ = v14
            }
            var v15 int32 = -199
            _ = fi11
            _ = v15
        }
        _ = fi10
    }
    fmt.Println(v2)
    v2 = v1
    v1 = v2
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint64 = 18446744073709551615
        _ = fi16
        _ = v17
    }
    v1 += v2
    v18 := []int{440, 4294967295, 9223372036854775807}
    v19 := GenS3{F0: -9223372036854775808, F1: v1}
    var v20 uint64 = 18446744073709551615
    var v21 float32 = float32(1.5)
    var v22 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
