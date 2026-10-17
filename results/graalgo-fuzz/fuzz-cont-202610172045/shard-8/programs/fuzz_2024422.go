package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 uint8, p2 uint64) int16 {
    var v6 int16 = -441
    if true {
        v7 := []int{p0, p0, (p0 * (-(p0)))}
        v8 := GenS1{F0: 2147483647}
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                if len(v7) > 0 {
                    v7[0] = -1
                }
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v7 = append(v7, fi10)
                _ = fi10
            }
            p0 -= (p0 + fi9)
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        p0 -= (p0 * p0)
        v11 := GenS2{F0: v6, F1: p0, F2: v6}
        var v12 uint8 = p1
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -32768
}

func main() {
    var v1 uint8 = 2
    v2 := []int{}
    var v3 int = 4294967295
    var v4 float64 = (float64(0.0) * 3.14159)
    var v5 uint32 = 1
    var v13 uint = 255
    var v14 uint64 = 18446744073709551615
    var v15 string = ""
    v3 = v3
    v16 := GenS3{F0: v15, F1: 42}
    fmt.Println(v3, v13, v15)
    if true {
        if false {
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: genFn0(fi17, v1, v14), F1: int(v13), F2: -1}
                v2 = append(v2, 9223372036854775807)
                v4 = v4
                _ = fi17
                _ = v18
            }
        }
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint16 = 255
            fmt.Println(v1)
            var v21 float32 = 0.0
            _ = fi19
            _ = v20
            _ = v21
        }
        if (v3 == v3) {
            v5 = 4294967295
            var v22 string = (v15 + v15)
            _ = v22
        }
    }
    v5 = 4294967295
    v23 := GenS4{F0: v15}
    if (1.5 != 3.14159) {
        var v24 float64 = v4
        _ = len(v2)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
