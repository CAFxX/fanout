package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16) uint8 {
    v4 := []int{4294967295, 32767, 4294967295}
    var v5 int64 = -115
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 -= p0
        p0 = (p0 * p0)
        fmt.Println(p0, fi7, fi7)
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() int32 {
    var v8 uint = 4294967295
    var v9 int8 = 0
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 float32 = ((1e10 * 1.5) * 1e10)
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    var v10 int8 = 1
    if (true || false) {
        v11 := []int{9223372036854775807, -1, 32768}
        _ = v11
    } else {
        v10 = (-(v10))
    }
    v1 -= v1
    var v12 float32 = -2.25
    v13 := GenS1{F0: -954}
    v14 := GenS2{F0: 4294967295}
    v15 := GenS2{F0: 4294967295}
    fmt.Println(v10)
    var v16 int64 = (v2 - 0)
    v17 := GenS1{F0: genFn1()}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
