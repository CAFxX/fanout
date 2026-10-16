package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0(p0 string, p1 string) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint32 = 4294967295
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.0
    _ = p0
    _ = p1
    _ = v8
    return 264
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    v9 := GenS1{F0: 18446744073709551615, F1: -15, F2: 9223372036854775807}
    fmt.Println(v1)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int16 = 0
        for fi12 := 0; fi12 < 10; fi12++ {
            v11 = 1000
            v1 = (v11 + -188)
            v11 += int16(fi12)
            switch (fi10 & fi12) {
            case 42:
                v11 = (v11 ^ (v11 * (v1 * 1000)))
            case 0:
                v3 += v3
            case -2147483648:
                var v13 int64 = 9223372036854775807
                _ = v13
            }
            v14 := GenS1{F0: 18446744073709551615, F1: 1, F2: -9223372036854775808}
            if false {
                v3 = v3
                v15 := GenS1{F0: 293, F1: 1, F2: -1}
                var v16 int64 = 42
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v14
        }
        v17 := []int{-1, fi10}
        _ = fi10
        _ = v11
        _ = v17
    }
    v3 = v3
    fmt.Println(v1, v3)
    fmt.Println(v3)
    if (!true) {
        var v18 int64 = 1
        _ = v18
    }
    v19 := []int{208, 9223372036854775807, 4294967295}
    var v20 float32 = 1e10
    v21 := GenS1{F0: 0, F1: -35, F2: -1}
    fmt.Println(v1, v20, v20)
    fmt.Println(v20, v1, v3)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
