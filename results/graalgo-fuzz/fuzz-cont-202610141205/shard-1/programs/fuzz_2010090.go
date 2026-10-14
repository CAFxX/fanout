package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64 }
func genFn0() int64 {
    var v7 uint32 = 677
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (!(true && (32767 >= 2))) {
                v7 = (^(v7))
                v10 := GenS1{F0: 1, F1: 255}
                fmt.Println(fi9)
                _ = v10
            }
            var v11 uint8 = 62
            _ = fi9
            _ = v11
        }
        v12 := GenS2{F0: 127, F1: (^(v7))}
        v13 := GenS1{F0: 0, F1: 1}
        _ = fi8
        _ = v12
        _ = v13
    }
    var v14 int16 = -180
    var v15 float32 = -2.25
    _ = v14
    _ = v15
    _ = v7
    return -929
}
func genFn1(p0 bool) int16 {
    var v16 uint16 = 2
    var v17 float64 = 3.14159
    _ = p0
    _ = v16
    _ = v17
    return -32768
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 float64 = (((3.14159 - 3.14159) - 0.0) + 1.5)
    v4 := []int{4294967295, 42, 127}
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    v4 = append(v4, 4294967295)
    if ((0 < 42) && true) {
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(fi18)
            _ = fi18
        }
    }
    delete(v6, "hello")
    var v19 uint32 = 157
    v3 = float64(3.14159)
    var v20 string = "\n\t"
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS3{F0: -95}
        var v23 bool = true
        _ = fi21
        _ = v22
        _ = v23
    }
    if (718 < 32767) {
        v24 := GenS2{F0: 42, F1: (v19 % ((v19) | 1))}
        _ = v24
    } else {
        _ = len(v6)
        var v25 uint16 = 255
        for fi26 := 0; fi26 < 9; fi26++ {
            v19 -= (v19 * (^(v19)))
            fmt.Println(v3, fi26)
            _ = fi26
        }
        _ = v25
    }
    var v27 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
