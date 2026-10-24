package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
func genFn0() int16 {
    var v6 int8 = 0
    v7 := []int{-570, -339}
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 string = "a"
    v2 := []string{("hello" + "")}
    var v3 float32 = -2.25
    var v4 string = (v1 + "")
    v5 := make(map[int]uint32)
    v3 = 1e10
    var v8 uint32 = 0
    v9 := GenS1{F0: 582, F1: -2147483648, F2: v4}
    v4 = ""
    for fi10 := 0; fi10 < 11; fi10++ {
        if (false || true) {
            fmt.Println(v8)
            if false {
                if true {
                    if true {
                        var v11 int = fi10
                        fmt.Println(v4)
                        _ = v11
                    }
                }
                fmt.Println(v3, fi10)
            }
            v1 = v1
        } else {
            v12 := v5[fi10]
            _ = v12
        }
        var v13 float64 = -2.25
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: 1, F1: -1, F2: (v4 + (("hello" + v4) + (v4 + v1)))}
    var v15 uint32 = 0
    var v16 uint = 4294967295
    for fi17 := 0; fi17 < 12; fi17++ {
        if (!((v16 / (((v16 << (((v16 << ((v16) & 63))) & 63))) | 1)) >= 1)) {
            v15 -= v8
            var v18 float64 = 1e10
            _ = v18
        }
        v8 += (v15 & v15)
        var v19 int64 = 9223372036854775807
        _ = fi17
        _ = v19
    }
    v8 += (v8 | 1)
    v1 = "a"
    if true {
        fmt.Println(v4, v16)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v20 float64 = ((3.14159 - 0.0) / 1.0)
    v16 += (v16 ^ v16)
    if (!true) {
        v4 = v1
        v3 += v3
    } else {
        var v21 int = -583
        _ = v21
    }
    v5[9223372036854775807] = (^(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
