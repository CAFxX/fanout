package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 int = -369
    var v2 uint8 = 127
    var v3 float32 = ((3.14159 + (float32(1e10) * 0.1)) / 1.5)
    var v4 float32 = (v3 / v3)
    v5 := make(map[int]uint32)
    if ((42 > 1) || (42 <= 0)) {
        switch (-(v1)) {
        case 643:
            fmt.Println(v4, v2, v1)
        case 127:
            delete(v5, (v1 >> ((492) & 63)))
        default:
            var v6 int = v1
            _ = v6
        }
        switch (v1 >> ((4294967295) & 63)) {
        case 531:
            if false {
                v7 := GenS1{F0: 2147483648}
                var v8 float32 = v3
                _ = v7
                _ = v8
            }
        case 128:
            v9 := []int{}
            _ = v9
        default:
            var v10 uint16 = 42
            _ = v10
        }
        v11 := GenS1{F0: 65535}
        _ = v11
    }
    v4 = (v4 / ((1e10 - v3) + 0.0))
    if false {
        v12 := []int{}
        v13 := GenS1{F0: 966}
        v14 := GenS1{F0: 2147483647}
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 float32 = 0.0
        if ((1000 == 42) || (!false)) {
            v2 = v2
            var v16 uint32 = 65535
            _ = v16
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 2
    v20 := []int{v1, (v1 | 128)}
    v21 := GenS2{F0: (^(v19)), F1: ((true && true) && (v19 >= (v19 ^ (v19 << ((642) & 63))))), F2: 42}
    v22 := GenS1{F0: 358}
    _ = len(v5)
    v23 := GenS3{F0: v19, F1: (v1 | 401), F2: uint64(v2)}
    v24 := GenS3{F0: (^(v19)), F1: v1, F2: v19}
    var v25 int16 = -32768
    var v26 int16 = v25
    switch v1 {
    case 1:
        v25 -= (-(v25))
    case 32768:
        _ = len(v20)
    case -910:
        _ = len(v20)
    default:
        switch v1 {
        case -9223372036854775808:
            var v27 uint = 4294967295
            _ = v27
        case 42:
            v19 = (v19 & 697)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
