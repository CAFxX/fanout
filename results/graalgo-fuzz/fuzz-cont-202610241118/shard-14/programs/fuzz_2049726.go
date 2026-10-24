package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 int8) int8 {
    fmt.Println(p0)
    v7 := GenS1{F0: 0, F1: 112}
    v8 := GenS1{F0: 1, F1: 0}
    var v9 int64 = 42
    v9 = v9
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (p0 << ((297) & 63))
}
func genFn1(p0 uint16) uint32 {
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: -1, F1: -1}
        var v12 uint8 = 2
        var v13 float32 = -2.25
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 string = "hello"
    v15 := GenS1{F0: -1, F1: 9223372036854775807}
    v16 := []int{-79}
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: 535, F1: 42}
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 int = fi17
            p0 -= p0
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v14 = v14
            _ = len(v16)
            _ = len(v16)
            _ = fi21
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-328}
    var v3 int = 4294967295
    var v4 string = "hello"
    v5 := []uint32{2147483648, 477}
    var v6 int64 = 0
    v3 = v3
    v3 = v3
    var v22 uint = 651
    _ = len(v5)
    v3 -= 859
    v22 -= v22
    fmt.Println(v3)
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v6 = (v6 & v6)
            _ = fi24
        }
        switch 32767 {
        case 42:
            var v25 float64 = -2.25
            _ = v25
        case 32768:
            v6 -= v6
        }
        _ = fi23
    }
    v26 := GenS1{F0: v6, F1: (v6 | (v6 | (^(v6))))}
    _ = len(v5)
    var v27 int32 = -707
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v28 uint64 = 42
    v28 = (^(v28))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
