package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint64 {
    var v5 uint = 18446744073709551615
    v6 := GenS1{F0: 0, F1: 1}
    v7 := []int{-474, -245, -1}
    v8 := GenS1{F0: 42, F1: -28}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32) bool {
    v9 := GenS1{F0: 1, F1: -1}
    v10 := []int{56, 4294967295, 809}
    _ = p0
    _ = v10
    _ = v9
    return true
}

func main() {
    var v1 int16 = -32768
    var v2 string = "hello"
    var v3 int64 = 859
    v4 := []uint32{4294967295, 1}
    if ((-746 > 9223372036854775807) && false) {
        var v11 uint8 = 0
        _ = v11
    } else {
        var v12 float64 = 0.0
        var v13 int16 = (v1 | (v1 * (v1 & v1)))
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int = fi14
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: -171, F1: int8(v1)}
    switch 127 {
    case 2147483647:
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS1{F0: 2147483647, F1: 0}
                v2 = v2
                _ = fi18
                _ = v19
            }
            v20 := GenS2{F0: 831}
            _ = fi17
            _ = v20
        }
    case 9223372036854775807:
        v2 = (((v2 + v2) + v2) + v2)
    case 94:
        fmt.Println(v1, v2)
    }
    if (0 >= 87) {
        v3 = v3
        for fi21 := 0; fi21 < 6; fi21++ {
            switch fi21 {
            case -9223372036854775808:
                var v22 int64 = v3
                _ = v22
            }
            _ = fi21
        }
        _ = len(v4)
    }
    if false {
        v4 = append(v4, 2147483647)
        v23 := GenS2{F0: 42}
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS1{F0: 1, F1: -1}
            var v26 uint32 = 42
            if (4294967295 != 55) {
                v3 += v3
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    v1 += v1
    v27 := GenS1{F0: -1, F1: 90}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
