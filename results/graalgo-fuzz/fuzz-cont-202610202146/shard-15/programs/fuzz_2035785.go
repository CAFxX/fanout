package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64, p1 int8) int32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7}
        var v9 bool = true
        var v10 float64 = 0.1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int64 = -1
    v11 = -33
    p1 -= p1
    _ = p0
    _ = p1
    _ = v11
    return 0
}
func genFn1(p0 uint32, p1 uint32) uint32 {
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            var v13 uint64 = 4294967295
            _ = v13
        }
        var v14 int = 32768
        _ = fi12
        _ = v14
    }
    var v15 uint8 = 1
    fmt.Println(v15)
    _ = p0
    _ = p1
    _ = v15
    return (p0 & p1)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 int64 = 739
    var v4 float64 = 1e10
    var v5 int = 9223372036854775807
    var v6 uint64 = 401
    var v16 bool = false
    v4 = (float64(0.1) / float64((0.0 + 1.5)))
    var v17 uint64 = (v6 * v6)
    var v18 float32 = (-2.25 - float32(v4))
    if v16 {
        v19 := GenS1{F0: 0}
        var v20 bool = ((0 != -85) && v16)
        _ = v19
        _ = v20
    }
    fmt.Println(v3, v17)
    var v21 int16 = -1
    var v22 uint16 = 65535
    if false {
        for fi23 := 0; fi23 < 16; fi23++ {
            if (!((v22 >> ((1) & 63)) > v22)) {
                if (!(!false)) {
                    v24 := GenS1{F0: 32767}
                    _ = v24
                }
            }
            _ = fi23
        }
    }
    v25 := v2[v5]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
