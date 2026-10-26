package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 string }

func main() {
    var v1 int8 = -85
    var v2 int32 = 1
    var v3 int8 = v1
    var v4 float64 = float64(1e10)
    v5 := make(map[string]uint32)
    var v6 uint = 952
    v7 := GenS1{F0: 42, F1: 475}
    var v8 int8 = v1
    var v9 int32 = -2147483648
    fmt.Println(v1, v8)
    var v10 int = 0
    v1 = (^(v3))
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS2{F0: 255, F1: (v3 | (v1 % ((42) | 1)))}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS3{F0: "\n\t"}
        var v15 int16 = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    delete(v5, "\n\t")
    v16 := GenS3{F0: ("a" + "")}
    v8 -= v1
    v17 := GenS2{F0: uint8(v6), F1: (v1 + int8(v6))}
    if false {
        if true {
            v18 := []int{v10, int(v6)}
            _ = v18
        } else {
            var v19 int64 = -1
            _ = v19
        }
    } else {
        if false {
            var v20 uint = 132
            _ = v20
        }
        if (v2 > (v2 ^ int32(v3))) {
            var v21 float32 = float32((-2.25 - 1.5))
            v1 += v3
            var v22 uint32 = 32
            _ = v21
            _ = v22
        }
        var v23 bool = (!true)
        _ = v23
    }
    if true {
        var v24 uint8 = 90
        var v25 uint64 = 2
        var v26 float64 = (0.1 / v4)
        _ = v24
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS1{F0: 9223372036854775807, F1: 0}
            v29 := GenS2{F0: 127, F1: int8(v6)}
            _ = fi27
            _ = v28
            _ = v29
        }
        fmt.Println(v9)
        switch 2147483647 {
        case 0:
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := v5[(("a" + ("foo bar" + ("a" + ""))) + "")]
                v8 = v1
                _ = fi30
                _ = v31
            }
        }
    }
    var v32 uint64 = 143
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint = (^(v6))
        _ = fi33
        _ = v34
    }
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
