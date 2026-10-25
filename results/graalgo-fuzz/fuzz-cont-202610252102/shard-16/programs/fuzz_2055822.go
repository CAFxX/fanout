package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint) uint64 {
    p1 = p1
    p1 -= p1
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 uint) bool {
    var v6 int = 2147483647
    var v7 uint8 = 42
    var v8 uint64 = 0
    v9 := GenS1{F0: 2}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (-2.25 != 1e10)
}

func main() {
    var v1 float32 = (0.0 * 1e10)
    v2 := make(map[string]int)
    var v3 int64 = 231
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    switch 2147483647 {
    case 32768:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint32 = 2
            var v12 int16 = 32767
            if (!((v12 != v12) && (!(127 > -51)))) {
                var v13 uint16 = 409
                v14 := []int{9223372036854775807, fi10}
                v1 -= v1
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case 127:
        v15 := GenS1{F0: 1}
        _ = v15
    default:
        v3 = -1
    }
    switch 496 {
    case 4294967295:
        var v16 uint64 = 42
        _ = v16
    }
    v2[("a" + "foo bar")] = 2147483647
    var v17 int = 42
    _ = len(v4)
    if true {
        var v18 uint32 = 0
        fmt.Println(v3, v3, v18)
        _ = v18
    }
    v19 := []int{(v17 / (((-(v17))) | 1))}
    v20 := GenS2{F0: 851}
    delete(v2, "")
    var v21 uint16 = 0
    switch -9223372036854775808 {
    case 4294967295:
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS2{F0: 2147483647}
                if ((1 <= 201) && (42 >= (v3 << ((42) & 63)))) {
                    if true {
                        fmt.Println(fi23)
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = len(v2)
            _ = fi22
        }
    case 997:
        if (!(!true)) {
            v1 = ((1.5 - -2.25) + v1)
        } else {
            var v25 int16 = -149
            _ = v25
        }
    case 0:
        for fi26 := 0; fi26 < 4; fi26++ {
            v1 += v1
            var v27 int = fi26
            var v28 int16 = 388
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        var v29 uint = 4294967295
        _ = v29
    }
    v3 = 1
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
