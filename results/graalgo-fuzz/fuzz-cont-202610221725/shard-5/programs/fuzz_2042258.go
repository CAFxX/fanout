package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 int64) uint16 {
    if (4294967295 < 1) {
        p0 = (^(p0))
        fmt.Println(p0)
        p0 = p0
    }
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(fi4, p0)
        _ = fi4
    }
    fmt.Println(p0)
    _ = p0
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 42
    var v3 int8 = 100
    for fi5 := 0; fi5 < 1; fi5++ {
        v2 = v2
        fmt.Println(v2, v2)
        v3 += v3
        _ = fi5
    }
    var v6 bool = false
    var v7 int = -862
    var v8 uint8 = 176
    var v9 uint = 173
    fmt.Println(v8)
    var v10 int16 = -32768
    v6 = (!("a" < ""))
    if (v10 >= (v10 | (-(v10)))) {
        if (v6 && v6) {
            if false {
                var v11 float32 = (float32((-2.25 / 1.0)) / 1e10)
                _ = v11
            }
            v10 = v10
            var v12 uint8 = (v8 >> (((v9 | v9)) & 63))
            _ = v12
        }
        var v13 int = (v7 << ((42) & 63))
        _ = v13
    }
    var v14 int16 = (v10 * 32767)
    v15 := GenS1{F0: "\n\t", F1: 1, F2: v8}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
