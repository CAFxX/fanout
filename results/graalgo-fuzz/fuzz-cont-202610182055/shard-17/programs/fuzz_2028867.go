package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 bool) uint16 {
    var v6 float32 = 1.5
    if ((0.1 + 0.1) >= 0.0) {
        fmt.Println(v6)
    } else {
        v7 := GenS1{F0: -1, F1: -880}
        v8 := []int{}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 622
}
func genFn1(p0 uint8) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS2{F0: "a"}
        _ = fi9
        _ = v10
    }
    if (-1 != 42) {
        var v11 int = 0
        _ = v11
    } else {
        v12 := []int{-2147483648, 0, 827, 127}
        var v13 uint16 = 255
        _ = v12
        _ = v13
    }
    _ = p0
    return (false && true)
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 string = v1
    delete(v4, v5)
    _ = len(v3)
    v3[2] = v5
    v1 = v1
    fmt.Println(v5, v5)
    fmt.Println(v1, v5)
    switch 2147483647 {
    case 1:
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 bool = (227 >= 42)
            _ = fi14
            _ = v15
        }
    case -1:
        fmt.Println(v5, v5)
    }
    v16 := GenS2{F0: v1}
    var v17 uint16 = 255
    _ = len(v4)
    var v18 int64 = -477
    for fi19 := 0; fi19 < 13; fi19++ {
        v4[(v1 + "a")] = 312
        v18 -= -1
        if true {
            var v20 int16 = -32768
            _ = v20
        } else {
            v18 = v18
            v1 = v1
            v21 := []int{fi19, 32767, (-(fi19))}
            _ = v21
        }
        _ = fi19
    }
    v22 := []int{}
    delete(v3, 2147483647)
    fmt.Println(v5, v5, v18)
    if true {
        switch ("a" + v1) {
        case "a":
            var v23 uint16 = (v17 - (^(v17)))
            _ = v23
        case "":
            var v24 uint64 = 18446744073709551615
            _ = v24
        case "hello":
            for fi25 := 0; fi25 < 9; fi25++ {
                v18 += -356
                _ = fi25
            }
        default:
            _ = len(v22)
        }
        var v26 bool = false
        _ = len(v22)
        _ = v26
    } else {
        v18 = (v18 / ((-509) | 1))
        v27 := GenS1{F0: 2147483647, F1: 32767}
        _ = v27
    }
    v22 = append(v22, -720)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
