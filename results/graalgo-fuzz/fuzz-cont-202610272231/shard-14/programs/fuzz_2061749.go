package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 int, p1 uint16) uint64 {
    var v5 int8 = 127
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{-456, p0, 127, p0}
        var v8 int16 = 1000
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (p0 << ((0) & 63)), F1: 2}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 string = "foo bar"
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 42
}

func main() {
    var v1 float64 = 0.0
    var v2 string = ""
    var v3 uint32 = 65535
    var v4 int = -2147483648
    switch ("\n\t" + v2) {
    case "a":
        v4 = 2147483647
    case "foo bar":
        fmt.Println(v3, v1)
    case "\n\t":
        v3 = (v3 * (v3 & 42))
    }
    if (-2147483648 >= -1) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 bool = true
            _ = fi12
            _ = v13
        }
        fmt.Println(v3, v1)
        for fi14 := 0; fi14 < 4; fi14++ {
            fmt.Println(v4, v4, v1)
            _ = fi14
        }
    } else {
        v4 -= -2147483648
        v4 += v4
        var v15 int64 = -1
        _ = v15
    }
    var v16 float32 = 1.5
    v16 = (v16 / (float32((1.5 * (v1 / v1))) - (float32((((v1 * 0.1) * v1) + v1)) + 0.0)))
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v4)
        fmt.Println(v4)
        _ = fi17
    }
    v1 -= (((-2.25 * v1) * v1) / 1.0)
    v18 := GenS1{F0: 127, F1: 416}
    var v19 int64 = -18
    var v20 uint8 = 1
    v16 = v16
    v3 += v3
    v1 += v1
    var v21 int8 = 100
    switch (((v2 + v2) + (v2 + ("a" + v2))) + "") {
    case "":
        if (0 <= (v21 << ((uint(v3)) & 63))) {
            if true {
                var v22 float32 = float32(float64(1.5))
                _ = v22
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                if (4294967295 <= 0) {
                    v24 := []int{v4, (v4 / (((v4 / (((v4 * fi23)) | 1))) | 1)), (fi23 / ((348) | 1)), fi23}
                    v25 := GenS2{F0: (v4 | (v4 & v4)), F1: (v2 + ("\n\t" + v2))}
                    _ = v24
                    _ = v25
                } else {
                    var v26 bool = false
                    _ = v26
                }
                _ = fi23
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := GenS1{F0: fi27, F1: 1}
                _ = fi27
                _ = v28
            }
        }
    }
    fmt.Println(v2)
    var v29 int16 = 722
    var v30 uint64 = 0
    v31 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
