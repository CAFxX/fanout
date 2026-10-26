package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
func genFn0() int {
    if true {
        var v7 uint16 = 0
        var v8 string = ""
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 50}
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: -1}
            v12 := []int{fi10, -9223372036854775808, (fi10 | fi10)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v13 := []int{-1, 2147483647, -820}
    v14 := GenS2{F0: 18446744073709551615}
    _ = v13
    _ = v14
    return -384
}
func genFn1(p0 int32, p1 uint8, p2 int8) uint8 {
    var v15 bool = false
    var v16 int16 = -184
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return p1
}

func main() {
    v1 := []uint32{158}
    v2 := make(map[string]bool)
    var v3 float64 = float64((0.0 + 3.14159))
    var v4 float32 = float32(v3)
    var v5 string = "\n\t"
    var v6 uint32 = 1
    fmt.Println(v3, v5)
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint8 = genFn1(806, 2, -1)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v4 = 1.5
        var v20 int32 = 2147483647
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v4 = ((v4 / float32(-2.25)) * v4)
        v2[v5] = false
        _ = fi21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 float64 = (v3 / v3)
        var v24 float32 = 3.14159
        v25 := GenS1{F0: 127}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println(v4)
    var v26 string = "\n\t"
    fmt.Println(v4)
    switch 2147483647 {
    case 32768:
        var v27 float32 = v4
        _ = v27
    case -2147483648:
        var v28 bool = (v6 < v6)
        _ = v28
    case 2147483647:
        var v29 int64 = 42
        _ = v29
    }
    var v30 int16 = 797
    v4 -= ((v4 + (float32((1e10 - (v3 + v3))) - v4)) + v4)
    var v31 int8 = 124
    v32 := GenS3{F0: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
