package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int16; F1 int; F2 int }

func main() {
    var v1 bool = (!(true || (!false)))
    var v2 int8 = -1
    var v3 int32 = -644
    v4 := make(map[uint32]int)
    v3 -= v3
    var v5 uint = 255
    v6 := GenS1{F0: 371, F1: 128}
    fmt.Println(v3, v3)
    v7 := []int{-750, 0, 42}
    var v8 int16 = 827
    v3 = (v3 + (v3 ^ (^(v3))))
    var v9 uint8 = 42
    v7 = append(v7, -2147483648)
    for fi10 := 0; fi10 < 1; fi10++ {
        v2 += (^(v2))
        if (v8 <= v8) {
            var v11 float32 = 0.1
            _ = v11
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 int16 = (v8 + (-(v8)))
                var v14 string = (("hello" + "hello") + "a")
                _ = fi12
                _ = v13
                _ = v14
            }
            v2 += v2
        }
        _ = fi10
    }
    var v15 float64 = 3.14159
    v16 := v4[42]
    v17 := GenS1{F0: (v16 * (v16 & 127)), F1: v16}
    if ((v9 | v9) < (v9 & 127)) {
        fmt.Println(v1)
        var v18 uint32 = 0
        _ = v18
    } else {
        if (1 == (v3 << ((v5) & 63))) {
            fmt.Println(v2)
        }
    }
    v19 := GenS2{F0: -1, F1: 1, F2: 592}
    if (!(-268 != v8)) {
        for fi20 := 0; fi20 < 10; fi20++ {
            if v1 {
                var v21 bool = (520 <= 255)
                _ = v21
            } else {
                var v22 float32 = 0.1
                _ = v22
            }
            v16 += v16
            _ = fi20
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            v9 -= v9
            _ = fi23
        }
        var v24 int32 = (v3 | 0)
        _ = v24
    } else {
        if ((v5 & (^(v5))) > 193) {
            v15 = v15
        }
    }
    if (!v1) {
        var v25 bool = v1
        if v25 {
            fmt.Println(v16)
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            if ((!(!false)) || v25) {
                v7 = append(v7, fi26)
                v7 = append(v7, (fi26 + (fi26 ^ -459)))
            } else {
                var v27 float32 = 3.14159
                v8 -= v8
                _ = v27
            }
            fmt.Println(v9, v5, v16)
            _ = fi26
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
