package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() string {
    var v6 int16 = 0
    fmt.Println(v6)
    v7 := GenS1{F0: 4294967295, F1: 127, F2: 796}
    var v8 int8 = 0
    _ = v6
    _ = v7
    _ = v8
    return ("a" + "\n\t")
}
func genFn1(p0 uint8, p1 int64, p2 string) int16 {
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := []int{(fi9 & 32767)}
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: 216, F1: 0}
        _ = fi9
        _ = v12
    }
    v13 := []int{282, 42}
    var v14 uint32 = 2147483648
    var v15 float32 = 0.0
    if (float64(float32((0.0 - 3.14159))) != 0.1) {
        if len(v13) > 0 {
            v13[0] = 128
        }
        fmt.Println(v15)
        v13 = append(v13, 128)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 496
}

func main() {
    var v1 uint32 = 42
    v2 := []int{}
    var v3 uint8 = uint8(v1)
    v4 := []string{"a", ("a" + "foo bar"), "foo bar", ("hello" + ("foo bar" + "a"))}
    v5 := make(map[uint32]string)
    v16 := GenS1{F0: 1, F1: 92, F2: 42}
    var v17 int = 128
    v1 += (v1 ^ (v1 - v1))
    v18 := GenS2{F0: v1, F1: 673}
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: 0, F1: -72, F2: 2147483647}
        _ = fi19
        _ = v20
    }
    var v21 uint8 = v3
    v1 = v1
    fmt.Println(v17, v17, v21)
    v17 = v17
    for fi22 := 0; fi22 < 11; fi22++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
