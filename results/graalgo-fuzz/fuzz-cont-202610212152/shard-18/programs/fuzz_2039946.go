package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 int }
func genFn0(p0 int8) uint64 {
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 uint64 = 842
            _ = fi5
            _ = v6
        }
        p0 -= p0
        _ = fi4
    }
    fmt.Println(p0)
    _ = p0
    return 60
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{0}
    v3 := []string{"hello", "\n\t"}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int32 = -546
        var v9 uint = 755
        v10 := GenS1{F0: 0, F1: -1, F2: (int8(v9) > 127)}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint16 = 255
    v11 = v11
    for fi12 := 0; fi12 < 1; fi12++ {
        v11 = 1
        if (!(!(-1 < 1))) {
            fmt.Println(v11, fi12)
            var v13 uint = 42
            var v14 int8 = -1
            _ = v13
            _ = v14
        } else {
            var v15 uint16 = 273
            v11 = (v11 - v15)
            var v16 int64 = 0
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: 18446744073709551615, F1: 100, F2: (2147483647 != 65535)}
        _ = fi12
        _ = v17
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        delete(v1, 4294967295)
    }
    v11 = (v11 << ((850) & 63))
    v18 := GenS2{F0: -665}
    if false {
        v11 -= (v11 / (((v11 << ((1) & 63))) | 1))
        v19 := GenS1{F0: 1, F1: -128, F2: false}
        _ = v19
    } else {
        var v20 float32 = ((-2.25 * ((((float32(1e10) / 1e10) / (float32(-2.25) - 1e10)) / 1.0) - float32((1e10 + 1.5)))) / 1.5)
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int64 = -9223372036854775808
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
}
