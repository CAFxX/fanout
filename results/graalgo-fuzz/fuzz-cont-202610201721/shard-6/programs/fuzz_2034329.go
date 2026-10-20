package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 string) bool {
    var v4 uint32 = 164
    v5 := GenS1{F0: 1, F1: 677}
    var v6 uint16 = 42
    p0 = p0
    fmt.Println(p0, v6, v4)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (2 > 108)
}
func genFn1() uint32 {
    var v7 uint = 42
    var v8 uint32 = 2
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{2147483647, 0}
    v3 := []string{}
    if false {
        v1 = v1
    } else {
        _ = len(v3)
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v2 = append(v2, 784)
    v1 += 18446744073709551615
    v1 = (v1 - v1)
    v1 -= v1
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 float32 = 0.1
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int64 = 1
            _ = fi11
            _ = v12
        }
        switch fi9 {
        case -1:
            fmt.Println(v10, v10)
        default:
            if false {
                v10 -= v10
                v1 = (v1 * v1)
                var v13 int32 = -2147483648
                _ = v13
            } else {
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 int64 = -9223372036854775808
                    _ = fi14
                    _ = v15
                }
                var v16 float32 = v10
                _ = v16
            }
        }
        _ = fi9
        _ = v10
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint8 = 42
        _ = fi17
        _ = v18
    }
    var v19 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
