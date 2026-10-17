package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
func genFn0() int {
    var v6 uint = 2
    var v7 int = 32768
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 uint16 = 328
    var v2 uint = 42
    var v3 int32 = 1
    v4 := []string{"hello"}
    var v5 uint32 = 2147483647
    if (float64(1e10) == 0.1) {
        var v8 uint16 = (^(v1))
        _ = v8
    }
    v5 -= v5
    var v9 bool = (3.14159 < (0.0 * 3.14159))
    if v9 {
        if false {
            if false {
                v2 = 42
                v10 := GenS1{F0: v3}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v10
            }
        }
        v1 += v1
    }
    fmt.Println(v3, v3, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 float32 = 1.5
            v14 := GenS2{F0: v2, F1: v3, F2: -128}
            var v15 uint = (v2 / ((0) | 1))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    if true {
        var v16 uint32 = v5
        fmt.Println(v9)
        _ = len(v4)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(v2, v5)
            v18 := GenS1{F0: v3}
            fmt.Println(v5)
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS2{F0: 2, F1: v3, F2: 100}
    v3 = 1
    var v20 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
