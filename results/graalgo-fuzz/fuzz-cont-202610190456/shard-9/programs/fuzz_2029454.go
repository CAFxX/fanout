package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }
func genFn0(p0 int) int8 {
    v7 := []int{p0, (p0 / (((^(p0))) | 1)), p0}
    p0 = p0
    v8 := GenS1{F0: 128, F1: 0}
    v7 = append(v7, 127)
    switch (("\n\t" + "a") + "") {
    case "":
        v9 := GenS1{F0: 128, F1: -2147483648}
        _ = v9
    default:
        var v10 int32 = 471
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 1000
    var v3 int16 = v2
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 bool = (!(65535 <= 403))
    v1 += 255
    if false {
        v6 = true
    } else {
        var v11 float32 = 0.0
        if (541 < 0) {
            if false {
                if (!true) {
                    var v12 int8 = 42
                    var v13 bool = true
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 int32 = 206
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v3 += (-(v3))
                            _ = fi16
                        }
                        _ = fi14
                        _ = v15
                    }
                    _ = v12
                    _ = v13
                }
            }
            v3 += (v2 & (v3 ^ v2))
        }
        v17 := GenS1{F0: (v1 << ((893) & 63)), F1: 1}
        _ = v11
        _ = v17
    }
    v18 := GenS1{F0: v1, F1: int32(v2)}
    v6 = (v1 > (^(v1)))
    var v19 uint32 = 65535
    fmt.Println(v3)
    if (!true) {
        v20 := []int{}
        var v21 uint64 = 18446744073709551615
        fmt.Println(v19, v21, v6)
        _ = v20
        _ = v21
    }
    v19 = v19
    v22 := []int{}
    v23 := GenS2{F0: v19, F1: 1, F2: 42}
    v5["\n\t"] = "a"
    var v24 string = "\n\t"
    v2 += v3
    var v25 int32 = 0
    var v26 uint32 = v19
    v4[17] = -491
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
