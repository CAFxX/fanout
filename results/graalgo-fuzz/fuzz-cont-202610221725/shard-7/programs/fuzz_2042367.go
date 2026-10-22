package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint16, p2 uint8) int8 {
    p0 = (p0 - p0)
    var v5 uint32 = 1
    var v6 uint8 = p2
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v5, p1)
        _ = fi7
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1() int16 {
    v8 := GenS1{F0: 419, F1: 0, F2: 9223372036854775807}
    var v9 float32 = 1.5
    var v10 int16 = -32768
    v11 := GenS1{F0: (v10 * (v10 % ((169) | 1))), F1: -128, F2: 44}
    var v12 string = "hello"
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return (v10 + (v10 / ((v10) | 1)))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int16 = -1
    v4 := make(map[uint32]uint32)
    fmt.Println(v3)
    for fi13 := 0; fi13 < 9; fi13++ {
        if false {
            v3 = (v3 * v3)
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: (v3 + (v3 << ((255) & 63))), F1: genFn0((-(v3)), 1, 1), F2: (fi14 * (fi14 - 9223372036854775807))}
                _ = fi14
                _ = v15
            }
            fmt.Println(v3, fi13)
        } else {
            switch (^(fi13)) {
            case -554:
                v16 := []int{32768, (fi13 % ((fi13) | 1))}
                _ = v16
            }
            v17 := GenS2{F0: -115}
            v3 -= v3
            _ = v17
        }
        _ = fi13
    }
    v3 = (v3 + v3)
    var v18 float32 = (0.1 + 3.14159)
    delete(v1, 127)
    var v19 bool = (166 == 4294967295)
    v20 := GenS1{F0: (v3 >> ((176) & 63)), F1: genFn0(genFn1(), 0, 128), F2: -2147483648}
    if v19 {
        v21 := GenS3{F0: 4294967295}
        var v22 uint16 = 24
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS3{F0: 255}
            _ = fi23
            _ = v24
        }
        _ = v21
        _ = v22
    } else {
        v19 = (!v19)
        v19 = (32767 > 0)
    }
    if ((!false) && ((v19 && (("\n\t" + "hello") == ("\n\t" + ""))) && v19)) {
        var v25 int64 = -1
        _ = v25
    }
    v26 := GenS4{F0: 125}
    var v27 float32 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
